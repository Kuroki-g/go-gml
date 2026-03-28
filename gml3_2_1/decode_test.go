package gml3_2_1_test

import (
	"strings"
	"testing"

	"github.com/Kuroki-g/go-gml/core"
	gml "github.com/Kuroki-g/go-gml/gml3_2_1"
)

func TestDecode_Point(t *testing.T) {
	const src = `<gml:Point xmlns:gml="http://www.opengis.net/gml/3.2"><gml:pos>1 2</gml:pos></gml:Point>`
	r := gml.NewReader(strings.NewReader(""))
	g, err := r.Decode(strings.NewReader(src))
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	pt, ok := g.Value.(core.Point)
	if !ok {
		t.Fatalf("expected Point, got %T", g.Value)
	}
	if len(pt) != 2 || pt[0] != 1 || pt[1] != 2 {
		t.Fatalf("unexpected point %v", pt)
	}
}

func TestDecode_Polygon(t *testing.T) {
	const src = `<gml:Polygon xmlns:gml="http://www.opengis.net/gml/3.2">` +
		`<gml:exterior><gml:LinearRing>` +
		`<gml:posList>0 0 1 0 1 1 0 1 0 0</gml:posList>` +
		`</gml:LinearRing></gml:exterior></gml:Polygon>`
	r := gml.NewReader(strings.NewReader(""))
	g, err := r.Decode(strings.NewReader(src))
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if _, ok := g.Value.(core.Polygon); !ok {
		t.Fatalf("expected Polygon, got %T", g.Value)
	}
}
