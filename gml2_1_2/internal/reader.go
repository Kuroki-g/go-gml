package internal

import (
	"encoding/xml"
	"fmt"
	"io"

	core "github.com/Kuroki-g/go-gml/core"
)

// GML 2.1.2 namespace URI.
const gmlNS2 = "http://www.opengis.net/gml"

func isGMLNS(ns string) bool { return ns == gmlNS2 }

// Reader scans a GML 2.1.2 document for geometry elements.
type Reader struct {
	src              io.ReadSeeker
	dec              *xml.Decoder
	charsetFn        func(string, io.Reader) (io.Reader, error)
	resolver         *resolver
	prescanned       bool
	OnUnknownElement func(name string)
}

// NewReader creates a Reader that streams geometry elements from r.
func NewReader(r io.ReadSeeker) *Reader {
	charsetFn := func(charset string, input io.Reader) (io.Reader, error) {
		return nil, fmt.Errorf("unsupported charset %q: call SetCharsetReader to handle non-UTF-8 files", charset)
	}
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charsetFn
	return &Reader{src: r, dec: dec, charsetFn: charsetFn, resolver: newResolver()}
}

// SetCharsetReader installs a charset conversion function on the underlying decoder.
func (r *Reader) SetCharsetReader(fn func(charset string, input io.Reader) (io.Reader, error)) {
	r.charsetFn = fn
	r.dec.CharsetReader = fn
}

func (r *Reader) runPreScan() error {
	preDec := xml.NewDecoder(r.src)
	preDec.CharsetReader = r.charsetFn
	if err := preScanGeometries(preDec, r.resolver); err != nil {
		return fmt.Errorf("gml: prescan: %w", err)
	}
	if _, err := r.src.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("gml: seek after prescan: %w", err)
	}
	r.dec = xml.NewDecoder(r.src)
	r.dec.CharsetReader = r.charsetFn
	r.prescanned = true
	return nil
}

type handlerFunc func(*xml.Decoder, xml.StartElement) (core.Geometry, error)

var handlers = map[string]handlerFunc{
	gmlPoint:      decodePointElement,
	gmlLineString: decodeLineStringElement,
	gmlPolygon:    decodePolygonElement,
	gmlBox:        decodeBoxElement,
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
			if r.OnUnknownElement != nil {
				r.OnUnknownElement(se.Name.Space + ":" + se.Name.Local)
			}
			continue
		}
		switch se.Name.Local {
		case gmlMultiPoint:
			return decodeMultiPointElement(r.dec, se, r.resolver)
		case gmlMultiLineString:
			return decodeMultiLineStringElement(r.dec, se, r.resolver)
		case gmlMultiPolygon:
			return decodeMultiPolygonElement(r.dec, se, r.resolver)
		}
		h, ok := handlers[se.Name.Local]
		if !ok {
			if r.OnUnknownElement != nil {
				r.OnUnknownElement(se.Name.Local)
			}
			continue
		}
		return h(r.dec, se)
	}
}
