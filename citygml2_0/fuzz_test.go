package citygml2_0_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	citygml2_0 "github.com/Kuroki-g/go-gml/citygml2_0"
	"github.com/Kuroki-g/go-gml/gml3_1_1"
)

func FuzzReader(f *testing.F) {
	f.Add([]byte(`<bldg:Building xmlns:bldg="http://www.opengis.net/citygml/building/2.0" xmlns:gml="http://www.opengis.net/gml" gml:id="b1"><bldg:lod0RoofEdge><gml:MultiSurface srsDimension="3"><gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing><gml:posList>0 0 0 1 0 0 1 1 0 0 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember></gml:MultiSurface></bldg:lod0RoofEdge></bldg:Building>`))
	f.Add([]byte(`<bldg:Building xmlns:bldg="http://www.opengis.net/citygml/building/2.0" xmlns:gml="http://www.opengis.net/gml" gml:id="b2"><bldg:lod1Solid><gml:Solid srsDimension="3"><gml:exterior><gml:CompositeSurface><gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing><gml:posList>0 0 0 1 0 0 1 1 0 0 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember></gml:CompositeSurface></gml:exterior></gml:Solid></bldg:lod1Solid></bldg:Building>`))
	f.Add([]byte(`<bldg:Building xmlns:bldg="http://www.opengis.net/citygml/building/2.0" xmlns:gml="http://www.opengis.net/gml" gml:id="b3"><bldg:lod2Solid><gml:Solid srsDimension="3"><gml:exterior><gml:CompositeSurface><gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing><gml:posList>0 0 0 1 0 0 1 1 0 0 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember></gml:CompositeSurface></gml:exterior></gml:Solid></bldg:lod2Solid></bldg:Building>`))
	f.Add([]byte(`<bldg:Building xmlns:bldg="http://www.opengis.net/citygml/building/2.0" xmlns:gml="http://www.opengis.net/gml" gml:id="b4"><bldg:lod2MultiSurface><gml:MultiSurface srsDimension="3"><gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing><gml:posList>0 0 0 1 0 0 1 1 0 0 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember></gml:MultiSurface></bldg:lod2MultiSurface></bldg:Building>`))
	f.Add([]byte(`<root/>`))
	f.Add([]byte(``))

	f.Fuzz(func(t *testing.T, data []byte) {
		r := citygml2_0.NewReader(bytes.NewReader(data), gml3_1_1.NewReader(strings.NewReader("")))
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
