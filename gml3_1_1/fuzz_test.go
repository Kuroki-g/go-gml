package gml3_1_1_test

import (
	"bytes"
	"io"
	"testing"

	gml "github.com/Kuroki-g/go-gml/gml3_1_1"
)

func FuzzReader(f *testing.F) {
	f.Add([]byte(`<gml:Point xmlns:gml="http://www.opengis.net/gml"><gml:pos>139.7 35.6</gml:pos></gml:Point>`))
	f.Add([]byte(`<gml:LineString xmlns:gml="http://www.opengis.net/gml" srsDimension="2"><gml:posList>0 0 1 1 2 0</gml:posList></gml:LineString>`))
	f.Add([]byte(`<gml:Polygon xmlns:gml="http://www.opengis.net/gml"><gml:exterior><gml:LinearRing><gml:posList>0 0 1 0 1 1 0 1 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon>`))
	f.Add([]byte(`<gml:MultiSurface xmlns:gml="http://www.opengis.net/gml" srsDimension="2"><gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing><gml:posList>0 0 1 0 1 1 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember></gml:MultiSurface>`))
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
