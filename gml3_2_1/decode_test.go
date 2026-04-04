package gml3_2_1_test

import (
	"io"
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

const gmlNS = `xmlns:gml="http://www.opengis.net/gml/3.2" xmlns:xlink="http://www.w3.org/1999/xlink"`

// TestNext_SolidExteriorHrefResolved verifies that a Solid whose Shell surfaceMember
// references a Polygon defined in the same document via xlink:href is resolved correctly.
func TestNext_SolidExteriorHrefResolved(t *testing.T) {
	const src = `<root ` + gmlNS + `>
<gml:Polygon gml:id="P1">
  <gml:exterior><gml:LinearRing><gml:posList>0 0 1 0 1 1 0 0</gml:posList></gml:LinearRing></gml:exterior>
</gml:Polygon>
<gml:Solid>
  <gml:exterior><gml:Shell><gml:surfaceMember xlink:href="#P1"/></gml:Shell></gml:exterior>
</gml:Solid>
</root>`
	r := gml.NewReader(strings.NewReader(src))
	var solid core.Solid
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Next: %v", err)
		}
		if s, ok := g.Value.(core.Solid); ok {
			solid = s
		}
	}
	if len(solid.Exterior) == 0 {
		t.Fatal("Solid.Exterior is empty: xlink:href not resolved")
	}
}

// TestNext_SolidExteriorHrefUnresolved verifies that a Solid whose Shell surfaceMember
// references a non-existent ID returns an error rather than silently returning nil geometry.
func TestNext_SolidExteriorHrefUnresolved(t *testing.T) {
	const src = `<root ` + gmlNS + `>
<gml:Solid>
  <gml:exterior><gml:Shell><gml:surfaceMember xlink:href="#NONEXISTENT"/></gml:Shell></gml:exterior>
</gml:Solid>
</root>`
	r := gml.NewReader(strings.NewReader(src))
	for {
		_, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return // expected: error for unresolved href
		}
	}
	t.Fatal("expected error for unresolved xlink:href, got none")
}

// TestNext_CompositeSurfaceSubPolygonHref verifies Fix 2: a Polygon with gml:id nested
// inside a CompositeSurface that prescan consumes can be resolved via xlink:href
// from another element in the same document.
func TestNext_CompositeSurfaceSubPolygonHref(t *testing.T) {
	const src = `<root ` + gmlNS + `>
<gml:CompositeSurface gml:id="CS1">
  <gml:surfaceMember>
    <gml:Polygon gml:id="P1">
      <gml:exterior><gml:LinearRing><gml:posList>0 0 1 0 1 1 0 0</gml:posList></gml:LinearRing></gml:exterior>
    </gml:Polygon>
  </gml:surfaceMember>
</gml:CompositeSurface>
<gml:Solid>
  <gml:exterior>
    <gml:Shell>
      <gml:surfaceMember xlink:href="#P1"/>
    </gml:Shell>
  </gml:exterior>
</gml:Solid>
</root>`
	r := gml.NewReader(strings.NewReader(src))
	var solid core.Solid
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Next: %v", err)
		}
		if s, ok := g.Value.(core.Solid); ok {
			solid = s
		}
	}
	if len(solid.Exterior) == 0 {
		t.Fatal("Solid.Exterior is empty: sub-polygon xlink:href not resolved after Fix 2")
	}
}
