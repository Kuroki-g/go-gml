package citygml2_0

import (
	"encoding/xml"
	"io"

	"github.com/Kuroki-g/go-gml/core"
)

const nsCityGML = "http://www.opengis.net/citygml/2.0"

// Reader streams Building features from a CityGML 2.0 document.
type Reader struct {
	dec    *xml.Decoder
	gmlDec core.Decoder
}

// NewReader creates a Reader from r. dec decodes GML geometry subtrees.
// Pass gml3_1_1.NewReader(r) as dec for PLATEAU/standard CityGML 2.0 data.
func NewReader(r io.ReadSeeker, dec core.Decoder) *Reader {
	return &Reader{
		dec:    xml.NewDecoder(r),
		gmlDec: dec,
	}
}

// Next returns the next Building in the stream, or io.EOF when exhausted.
func (r *Reader) Next() (*Building, error) {
	for {
		tok, err := r.dec.Token()
		if err != nil {
			return nil, err
		}
		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		if se.Name.Space == nsBldg && se.Name.Local == "Building" {
			return decodeBuilding(r.dec, se, r.gmlDec)
		}
	}
}
