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
// Value is one of: Point, LineString, Polygon, MultiPoint, MultiLineString, MultiPolygon, Bound.
type Geometry struct {
	Value   interface{}
	SRSName string // raw srsName attribute value from the GML element; empty if absent
}

// Reader scans a GML document for geometry elements.
type Reader struct {
	dec        *xml.Decoder
	defaultSRS string // last srsName seen on any element in the stream
}

// NewReader creates a Reader that streams geometry elements from r.
// For non-UTF-8 encoded files (e.g. Shift-JIS), call SetCharsetReader after creation.
func NewReader(r io.Reader) *Reader {
	dec := xml.NewDecoder(newNSNormReader(r))
	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return nil, fmt.Errorf("unsupported charset %q: call SetCharsetReader to handle non-UTF-8 files", charset)
	}
	return &Reader{dec: dec}
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

type handlerFunc func(*xml.Decoder, xml.StartElement) (Geometry, error)

var handlers = map[string]handlerFunc{
	// SF-0 elements
	"Point":      decodePointElement,
	"LineString": decodeLineStringElement,
	"Polygon":    decodePolygonElement,
	// Real-data elements (gml:Curve and gml:Surface with patch/segment pattern)
	"Curve":   decodeCurveElement,
	"Surface": decodeSurfaceElement,
	// Multi-geometry
	"MultiPoint":      decodeMultiPointElement,
	"MultiCurve":      decodeMultiCurveElement,
	"MultiLineString": decodeMultiCurveElement,
	"MultiSurface":    decodeMultiSurfaceElement,
	"MultiPolygon":    decodeMultiSurfaceElement,
	// Bounding box
	"Envelope": decodeEnvelopeElement,
}

// Next returns the next geometry found in the stream.
// Returns io.EOF when the document is exhausted.
func (r *Reader) Next() (Geometry, error) {
	for {
		tok, err := r.dec.Token()
		if err != nil {
			return Geometry{}, err // includes io.EOF
		}
		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		// Track srsName from any element encountered in the stream.
		// GML commonly declares srsName once on a bounding box or feature collection
		// element; individual geometry elements then inherit it implicitly.
		for _, attr := range se.Attr {
			if attr.Name.Local == "srsName" && attr.Name.Space == "" && attr.Value != "" {
				r.defaultSRS = attr.Value
			}
		}
		if !isGMLNS(se.Name.Space) {
			continue
		}
		h, ok := handlers[se.Name.Local]
		if !ok {
			continue
		}
		g, err := h(r.dec, se)
		// Inherit document-level srsName when the geometry element has none.
		if g.SRSName == "" && r.defaultSRS != "" {
			g.SRSName = r.defaultSRS
		}
		return g, err
	}
}
