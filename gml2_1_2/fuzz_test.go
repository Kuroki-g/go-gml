package gml2_1_2_test

import (
	"bytes"
	"io"
	"testing"

	gml "github.com/Kuroki-g/go-gml/gml2_1_2"
)

func FuzzReader(f *testing.F) {
	f.Add([]byte(`<gml:Point xmlns:gml="http://www.opengis.net/gml"><gml:coordinates>139.7,35.6</gml:coordinates></gml:Point>`))
	f.Add([]byte(`<gml:LineString xmlns:gml="http://www.opengis.net/gml"><gml:coordinates>0,0 1,1 2,0</gml:coordinates></gml:LineString>`))
	f.Add([]byte(`<gml:Polygon xmlns:gml="http://www.opengis.net/gml"><gml:outerBoundaryIs><gml:LinearRing><gml:coordinates>0,0 10,0 10,10 0,10 0,0</gml:coordinates></gml:LinearRing></gml:outerBoundaryIs></gml:Polygon>`))
	f.Add([]byte(`<gml:MultiPolygon xmlns:gml="http://www.opengis.net/gml"><gml:polygonMember><gml:Polygon><gml:outerBoundaryIs><gml:LinearRing><gml:coordinates>0,0 1,0 1,1 0,0</gml:coordinates></gml:LinearRing></gml:outerBoundaryIs></gml:Polygon></gml:polygonMember></gml:MultiPolygon>`))
	f.Add([]byte(`<root/>`))
	f.Add([]byte(``))

	f.Fuzz(func(t *testing.T, data []byte) {
		r := gml.NewReader(bytes.NewReader(data))
		for {
			_, err := r.Next()
			if err == io.EOF {
				return
			}
			if err != nil {
				return
			}
		}
	})
}
