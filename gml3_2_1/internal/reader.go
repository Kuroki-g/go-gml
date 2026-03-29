package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"

	core "github.com/Kuroki-g/go-gml/core"
)

// GML 3.2.1 namespace URI.
const gmlNS32 = "http://www.opengis.net/gml/3.2"

func isGMLNS(ns string) bool { return ns == gmlNS32 }

// Reader scans a GML 3.2.1 document for geometry elements.
type Reader struct {
	src         io.ReadSeeker
	dec         *xml.Decoder
	charsetFn   func(string, io.Reader) (io.Reader, error)
	resolver    *curveResolver
	prescanned  bool
	pendingGrid *gridBounds
	globalDim   int // srsDimension captured from root gml:Envelope; 0 if not yet seen
}

// NewReader creates a Reader that streams geometry elements from r.
func NewReader(r io.ReadSeeker) *Reader {
	charsetFn := func(charset string, input io.Reader) (io.Reader, error) {
		return nil, fmt.Errorf("unsupported charset %q: call SetCharsetReader to handle non-UTF-8 files", charset)
	}
	dec := xml.NewDecoder(newNSNormReader(r))
	dec.CharsetReader = charsetFn
	return &Reader{src: r, dec: dec, charsetFn: charsetFn, resolver: newCurveResolver()}
}

func (r *Reader) runPreScan() error {
	preDec := xml.NewDecoder(newNSNormReader(r.src))
	preDec.CharsetReader = r.charsetFn
	if err := preScanGeometries(preDec, r.resolver); err != nil {
		return fmt.Errorf("gml: prescan: %w", err)
	}
	if _, err := r.src.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("gml: seek after prescan: %w", err)
	}
	r.dec = xml.NewDecoder(newNSNormReader(r.src))
	r.dec.CharsetReader = r.charsetFn
	r.prescanned = true
	return nil
}

// SetCharsetReader installs a charset conversion function on the underlying decoder.
func (r *Reader) SetCharsetReader(fn func(charset string, input io.Reader) (io.Reader, error)) {
	r.charsetFn = fn
	r.dec.CharsetReader = fn
}

// extractGMLID returns the value of the gml:id attribute from a StartElement, or "".
func extractGMLID(se xml.StartElement) string {
	for _, a := range se.Attr {
		if a.Name.Local == "id" && isGMLNS(a.Name.Space) {
			return a.Value
		}
	}
	return ""
}

type handlerFunc func(*xml.Decoder, xml.StartElement) (core.Geometry, error)

var handlers = map[string]handlerFunc{
	gmlPoint:      decodePointElement,
	gmlLineString: decodeLineStringElement,
	gmlMultiPoint: decodeMultiPointElement,
}

// Next returns the next geometry found in the stream.
// Returns io.EOF when the document is exhausted.
func (r *Reader) Next() (core.Geometry, error) {
	if !r.prescanned {
		if err := r.runPreScan(); err != nil {
			return core.Geometry{}, err
		}
	}
	for {
		tok, err := r.dec.Token()
		if err != nil {
			return core.Geometry{}, err
		}
		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		if !isGMLNS(se.Name.Space) {
			continue
		}
		switch se.Name.Local {
		case gmlCurve:
			return r.handleCurve(r.dec, se)
		case gmlOrientableCurve:
			if err := r.cacheOrientableCurve(r.dec, se); err != nil {
				return core.Geometry{}, err
			}
			continue
		case gmlEnvelope:
			for _, a := range se.Attr {
				if a.Name.Local == "srsDimension" {
					if d, err2 := strconv.Atoi(a.Value); err2 == nil && d > 0 {
						r.globalDim = d
					}
				}
			}
			return decodeEnvelopeElement(r.dec, se)
		case gmlPolygon:
			return r.handlePolygon(r.dec, se)
		case gmlSurface:
			return r.handleSurface(r.dec, se)
		case gmlCompositeCurve:
			return r.handleCompositeCurve(r.dec, se)
		case gmlCompositeSurface:
			return r.handleCompositeSurface(r.dec, se)
		case gmlOrientableSurface:
			return r.handleOrientableSurface(r.dec, se)
		case gmlMultiCurve, gmlMultiLineString:
			return r.handleMultiCurve(r.dec, se)
		case gmlMultiSurface, gmlMultiPolygon:
			return r.handleMultiSurface(r.dec, se)
		case gmlTin:
			return r.handleTin(r.dec, se)
		case gmlTriangulatedSurface:
			return r.handleTriangulatedSurface(r.dec, se)
		case gmlSolid:
			return r.handleSolid(r.dec, se)
		case gmlCompositeSolid:
			return r.handleCompositeSolid(r.dec, se)
		case gmlMultiSolid:
			return r.handleMultiSolid(r.dec, se)
		case gmlDomainSet:
			if err := r.handleDomainSet(r.dec, se); err != nil {
				return core.Geometry{}, err
			}
			continue
		case gmlRangeSet:
			if r.pendingGrid == nil {
				if err := r.dec.Skip(); err != nil {
					return core.Geometry{}, err
				}
				continue
			}
			return r.handleRangeSet(r.dec, se)
		}
		h, ok := handlers[se.Name.Local]
		if !ok {
			continue
		}
		g, err := h(r.dec, se)
		return g, err
	}
}
