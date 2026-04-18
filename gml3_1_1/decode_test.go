package gml3_1_1_test

import (
	"io"
	"strings"
	"testing"

	"github.com/Kuroki-g/go-gml/core"
	gml "github.com/Kuroki-g/go-gml/gml3_1_1"
)

func TestDecode_Point(t *testing.T) {
	const src = `<gml:Point xmlns:gml="http://www.opengis.net/gml"><gml:pos>1 2</gml:pos></gml:Point>`
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
	const src = `<gml:Polygon xmlns:gml="http://www.opengis.net/gml">` +
		`<gml:exterior><gml:LinearRing>` +
		`<gml:posList srsDimension="2">0 0 1 0 1 1 0 1 0 0</gml:posList>` +
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

const gmlNS = `xmlns:gml="http://www.opengis.net/gml" xmlns:xlink="http://www.w3.org/1999/xlink"`

// TestNext_SolidExteriorHrefResolved verifies that a Solid whose exterior references
// a Polygon defined in the same document via xlink:href is resolved correctly.
func TestNext_SolidExteriorHrefResolved(t *testing.T) {
	const src = `<root ` + gmlNS + `>
<gml:Polygon gml:id="P1">
  <gml:exterior><gml:LinearRing><gml:posList srsDimension="2">0 0 1 0 1 1 0 0</gml:posList></gml:LinearRing></gml:exterior>
</gml:Polygon>
<gml:Solid>
  <gml:exterior><gml:CompositeSurface><gml:surfaceMember xlink:href="#P1"/></gml:CompositeSurface></gml:exterior>
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

// TestNext_SolidExteriorHrefUnresolved verifies that a Solid whose exterior references
// a non-existent ID returns an error rather than silently returning nil geometry.
func TestNext_SolidExteriorHrefUnresolved(t *testing.T) {
	const src = `<root ` + gmlNS + `>
<gml:Solid>
  <gml:exterior xlink:href="#NONEXISTENT"/>
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

// TestDecode_LineString_Pos verifies dimension inference for gml:LineString using gml:pos elements.
// Covered code: gml3_1_1/internal/decode_linestring.go (preferDim + field-count inference block).
func TestDecode_LineString_Pos(t *testing.T) {
	const ns = `xmlns:gml="http://www.opengis.net/gml"`
	tests := []struct {
		name    string
		src     string
		wantPts int
		wantDim int
	}{
		{
			name:    "2D inferred from pos field count",
			src:     `<gml:LineString ` + ns + `><gml:pos>1 2</gml:pos><gml:pos>3 4</gml:pos></gml:LineString>`,
			wantPts: 2, wantDim: 2,
		},
		{
			name:    "3D inferred from pos field count",
			src:     `<gml:LineString ` + ns + `><gml:pos>1 2 10</gml:pos><gml:pos>3 4 20</gml:pos></gml:LineString>`,
			wantPts: 2, wantDim: 3,
		},
		{
			name:    "explicit srsDimension=3 on LineString element",
			src:     `<gml:LineString ` + ns + ` srsDimension="3"><gml:pos>1 2 10</gml:pos><gml:pos>3 4 20</gml:pos></gml:LineString>`,
			wantPts: 2, wantDim: 3,
		},
		{
			name:    "explicit srsDimension=3 on pos element",
			src:     `<gml:LineString ` + ns + `><gml:pos srsDimension="3">1 2 10</gml:pos><gml:pos srsDimension="3">3 4 20</gml:pos></gml:LineString>`,
			wantPts: 2, wantDim: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gml.NewReader(strings.NewReader(""))
			g, err := r.Decode(strings.NewReader(tt.src))
			if err != nil {
				t.Fatalf("Decode: %v", err)
			}
			ls, ok := g.Value.(core.LineString)
			if !ok {
				t.Fatalf("expected LineString, got %T", g.Value)
			}
			if len(ls) != tt.wantPts {
				t.Fatalf("point count=%d want %d", len(ls), tt.wantPts)
			}
			for i, pt := range ls {
				if len(pt) != tt.wantDim {
					t.Errorf("ls[%d] dim=%d want %d", i, len(pt), tt.wantDim)
				}
			}
		})
	}
}

// TestDecode_LineString_Coordinates verifies gml:LineString using gml:coordinates.
// The current implementation hard-codes dim=2 for gml:coordinates.
// Covered code: gml3_1_1/internal/decode_linestring.go (coordinates branch).
func TestDecode_LineString_Coordinates(t *testing.T) {
	const ns = `xmlns:gml="http://www.opengis.net/gml"`
	src := `<gml:LineString ` + ns + `><gml:coordinates>1,2 3,4</gml:coordinates></gml:LineString>`
	r := gml.NewReader(strings.NewReader(""))
	g, err := r.Decode(strings.NewReader(src))
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	ls, ok := g.Value.(core.LineString)
	if !ok {
		t.Fatalf("expected LineString, got %T", g.Value)
	}
	if len(ls) != 2 {
		t.Fatalf("point count=%d want 2", len(ls))
	}
	for i, pt := range ls {
		if len(pt) != 2 {
			t.Errorf("ls[%d] dim=%d want 2", i, len(pt))
		}
	}
}

// TestDecode_Polygon_Pos verifies dimension inference for gml:Polygon using gml:pos elements
// in LinearRing. Covered code: gml3_1_1/internal/decode_polygon.go (ringFromLinearRing).
func TestDecode_Polygon_Pos(t *testing.T) {
	const ns = `xmlns:gml="http://www.opengis.net/gml"`
	tests := []struct {
		name    string
		src     string
		wantDim int
	}{
		{
			name: "2D inferred from pos field count",
			src: `<gml:Polygon ` + ns + `><gml:exterior><gml:LinearRing>` +
				`<gml:pos>0 0</gml:pos><gml:pos>1 0</gml:pos><gml:pos>1 1</gml:pos><gml:pos>0 0</gml:pos>` +
				`</gml:LinearRing></gml:exterior></gml:Polygon>`,
			wantDim: 2,
		},
		{
			name: "3D inferred from pos field count",
			src: `<gml:Polygon ` + ns + `><gml:exterior><gml:LinearRing>` +
				`<gml:pos>0 0 10</gml:pos><gml:pos>1 0 10</gml:pos><gml:pos>1 1 10</gml:pos><gml:pos>0 0 10</gml:pos>` +
				`</gml:LinearRing></gml:exterior></gml:Polygon>`,
			wantDim: 3,
		},
		{
			name: "explicit srsDimension=3 on Polygon inherited by pos",
			src: `<gml:Polygon ` + ns + ` srsDimension="3"><gml:exterior><gml:LinearRing>` +
				`<gml:pos>0 0 10</gml:pos><gml:pos>1 0 10</gml:pos><gml:pos>1 1 10</gml:pos><gml:pos>0 0 10</gml:pos>` +
				`</gml:LinearRing></gml:exterior></gml:Polygon>`,
			wantDim: 3,
		},
		{
			name: "explicit srsDimension=3 on LinearRing inherited by pos",
			src: `<gml:Polygon ` + ns + `><gml:exterior><gml:LinearRing srsDimension="3">` +
				`<gml:pos>0 0 10</gml:pos><gml:pos>1 0 10</gml:pos><gml:pos>1 1 10</gml:pos><gml:pos>0 0 10</gml:pos>` +
				`</gml:LinearRing></gml:exterior></gml:Polygon>`,
			wantDim: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gml.NewReader(strings.NewReader(""))
			g, err := r.Decode(strings.NewReader(tt.src))
			if err != nil {
				t.Fatalf("Decode: %v", err)
			}
			poly, ok := g.Value.(core.Polygon)
			if !ok {
				t.Fatalf("expected Polygon, got %T", g.Value)
			}
			if len(poly) == 0 {
				t.Fatal("polygon has no rings")
			}
			for i, ring := range poly {
				for j, pt := range ring {
					if len(pt) != tt.wantDim {
						t.Errorf("ring[%d][%d] dim=%d want %d", i, j, len(pt), tt.wantDim)
					}
				}
			}
		})
	}
}

// TestNext_GlobalDim_PropagatedToPolygon verifies that srsDimension from gml:Envelope
// is captured as globalDim and propagated to subsequent gml:Polygon pos elements.
func TestNext_GlobalDim_PropagatedToPolygon(t *testing.T) {
	const src = `<root xmlns:gml="http://www.opengis.net/gml">
<gml:Envelope srsDimension="3" srsName="urn:ogc:def:crs:EPSG::6697">
  <gml:lowerCorner>35.0 139.0 0</gml:lowerCorner>
  <gml:upperCorner>36.0 140.0 100</gml:upperCorner>
</gml:Envelope>
<gml:Polygon>
  <gml:exterior><gml:LinearRing>
    <gml:posList>0 0 10 1 0 10 1 1 10 0 0 10</gml:posList>
  </gml:LinearRing></gml:exterior>
</gml:Polygon>
</root>`
	r := gml.NewReader(strings.NewReader(src))
	var poly core.Polygon
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Next: %v", err)
		}
		if p, ok := g.Value.(core.Polygon); ok {
			poly = p
		}
	}
	if len(poly) == 0 {
		t.Fatal("no polygon decoded")
	}
	for i, ring := range poly {
		for j, pt := range ring {
			if len(pt) != 3 {
				t.Errorf("ring[%d][%d] dim=%d want 3 (globalDim not propagated)", i, j, len(pt))
			}
		}
	}
}

// TestNext_CompositeSurfaceSubPolygonHref verifies Fix 2: a Polygon with gml:id nested
// inside a CompositeSurface that prescan consumes can be resolved via xlink:href
// from another element in the same document.
func TestNext_CompositeSurfaceSubPolygonHref(t *testing.T) {
	const src = `<root ` + gmlNS + `>
<gml:CompositeSurface gml:id="CS1">
  <gml:surfaceMember>
    <gml:Polygon gml:id="P1">
      <gml:exterior><gml:LinearRing><gml:posList srsDimension="2">0 0 1 0 1 1 0 0</gml:posList></gml:LinearRing></gml:exterior>
    </gml:Polygon>
  </gml:surfaceMember>
</gml:CompositeSurface>
<gml:Solid>
  <gml:exterior>
    <gml:CompositeSurface>
      <gml:surfaceMember xlink:href="#P1"/>
    </gml:CompositeSurface>
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
