package gml2_1_2_test

import (
	"strings"
	"testing"

	"github.com/Kuroki-g/go-gml/core"
	gml "github.com/Kuroki-g/go-gml/gml2_1_2"
)

func TestDecode_Point(t *testing.T) {
	const src = `<gml:Point xmlns:gml="http://www.opengis.net/gml"><gml:coordinates>1,2</gml:coordinates></gml:Point>`
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
