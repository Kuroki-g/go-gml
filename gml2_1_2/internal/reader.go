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
	dec *xml.Decoder
}

// NewReader creates a Reader that streams geometry elements from r.
func NewReader(r io.Reader) *Reader {
	dec := xml.NewDecoder(r)
	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return nil, fmt.Errorf("unsupported charset %q: call SetCharsetReader to handle non-UTF-8 files", charset)
	}
	return &Reader{dec: dec}
}

// SetCharsetReader installs a charset conversion function on the underlying decoder.
func (r *Reader) SetCharsetReader(fn func(charset string, input io.Reader) (io.Reader, error)) {
	r.dec.CharsetReader = fn
}

type handlerFunc func(*xml.Decoder, xml.StartElement) (core.Geometry, error)

var handlers = map[string]handlerFunc{
	gmlPoint:           decodePointElement,
	gmlLineString:      decodeLineStringElement,
	gmlPolygon:         decodePolygonElement,
	gmlMultiPoint:      decodeMultiPointElement,
	gmlMultiLineString: decodeMultiLineStringElement,
	gmlMultiPolygon:    decodeMultiPolygonElement,
	gmlBox:             decodeBoxElement,
}

// Next returns the next geometry found in the stream.
// Returns io.EOF when the document is exhausted.
func (r *Reader) Next() (core.Geometry, error) {
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
		h, ok := handlers[se.Name.Local]
		if !ok {
			continue
		}
		return h(r.dec, se)
	}
}
