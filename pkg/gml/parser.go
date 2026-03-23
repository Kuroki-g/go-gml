package gml

import (
	"encoding/xml"
	"fmt"
	"io"
)

// GML namespace URIs recognised by this parser.
const (
	gmlNS32 = "http://www.opengis.net/gml/3.2"
	gmlNS2  = "http://www.opengis.net/gml"
)

func isGMLNS(ns string) bool { return ns == gmlNS32 || ns == gmlNS2 }

// ---- public types ----

// Geometry wraps a parsed GML geometry element with CRS metadata.
// Value is one of: Point, LineString, Polygon, MultiPoint, MultiLineString, MultiPolygon, Bound, GridCoverage.
type Geometry struct {
	Value   interface{}
	SRSName *string // nil = srsName attribute absent; non-nil = explicitly set in XML
}

// Reader scans a GML document for geometry elements.
type Reader struct {
	src         io.ReadSeeker
	dec         *xml.Decoder
	resolver    *curveResolver // gml:id → Curve/OrientableCurve, for xlink:href resolution
	prescanned  bool
	pendingGrid *gridBounds // set by domainSet handler; consumed by rangeSet handler
}

// NewReader creates a Reader that streams geometry elements from r.
// For non-UTF-8 encoded files (e.g. Shift-JIS), call SetCharsetReader after creation.
// On the first call to Next, the document is pre-scanned to populate the xlink:href
// resolver cache, enabling forward references to be resolved correctly.
func NewReader(r io.ReadSeeker) *Reader {
	dec := xml.NewDecoder(newNSNormReader(r))
	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return nil, fmt.Errorf("unsupported charset %q: call SetCharsetReader to handle non-UTF-8 files", charset)
	}
	return &Reader{src: r, dec: dec, resolver: newCurveResolver()}
}

// runPreScan runs a full document scan to populate the resolver cache before main parsing.
// This enables forward xlink:href references to be resolved in the main pass.
func (r *Reader) runPreScan() error {
	charsetFn := r.dec.CharsetReader
	preDec := xml.NewDecoder(newNSNormReader(r.src))
	preDec.CharsetReader = charsetFn
	if err := preScanGeometries(preDec, r.resolver); err != nil {
		return fmt.Errorf("gml: prescan: %w", err)
	}
	if _, err := r.src.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("gml: seek after prescan: %w", err)
	}
	r.dec = xml.NewDecoder(newNSNormReader(r.src))
	r.dec.CharsetReader = charsetFn
	r.prescanned = true
	return nil
}

// SetCharsetReader installs a charset conversion function on the underlying decoder.
// Use this to support non-UTF-8 encoded GML files (e.g. Shift-JIS found in older 国土数値情報 files).
//
// Example using golang.org/x/text:
//
//	reader.SetCharsetReader(func(charset string, input io.Reader) (io.Reader, error) {
//	    if strings.EqualFold(charset, "shift_jis") {
//	        return transform.NewReader(input, japanese.ShiftJIS.NewDecoder()), nil
//	    }
//	    return nil, fmt.Errorf("unsupported charset: %s", charset)
//	})
func (r *Reader) SetCharsetReader(fn func(charset string, input io.Reader) (io.Reader, error)) {
	r.dec.CharsetReader = fn
}

// ---- handlers map: this is the index of all supported GML geometry elements ----

// extractGMLID returns the value of the gml:id attribute from a StartElement, or "".
func extractGMLID(se xml.StartElement) string {
	for _, a := range se.Attr {
		if a.Name.Local == "id" && isGMLNS(a.Name.Space) {
			return a.Value
		}
	}
	return ""
}

type handlerFunc func(*xml.Decoder, xml.StartElement) (Geometry, error)

var handlers = map[string]handlerFunc{
	// SF-0 elements
	gmlPoint:      decodePointElement,
	gmlLineString: decodeLineStringElement,
	// Polygon is handled via Reader.handlePolygon in Next() for caching support.
	// Curve and Surface are handled via Reader methods in Next() for xlink:href support.
	// Multi-geometry (MultiPoint only; Multi{Curve,Surface} need resolver → handled in Next)
	gmlMultiPoint: decodeMultiPointElement,
	// Bounding box
	gmlEnvelope: decodeEnvelopeElement,
}

// Next returns the next geometry found in the stream.
// Returns io.EOF when the document is exhausted.
// On the first call, the document is pre-scanned to resolve forward xlink:href references.
func (r *Reader) Next() (Geometry, error) {
	if !r.prescanned {
		if err := r.runPreScan(); err != nil {
			return Geometry{}, err
		}
	}
	for {
		tok, err := r.dec.Token()
		if err != nil {
			return Geometry{}, err // includes io.EOF
		}
		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		if !isGMLNS(se.Name.Space) {
			continue
		}
		// Elements that need Reader state for xlink:href resolution or caching.
		switch se.Name.Local {
		case gmlCurve:
			return r.handleCurve(r.dec, se)
		case gmlOrientableCurve:
			if err := r.cacheOrientableCurve(r.dec, se); err != nil {
				return Geometry{}, err
			}
			continue
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
		case gmlDomainSet:
			if err := r.handleDomainSet(r.dec, se); err != nil {
				return Geometry{}, err
			}
			continue
		case gmlRangeSet:
			if r.pendingGrid == nil {
				if err := r.dec.Skip(); err != nil {
					return Geometry{}, err
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
