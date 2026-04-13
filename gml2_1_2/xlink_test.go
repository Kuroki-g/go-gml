package gml2_1_2_test

// Tests for xlink:href resolution in Multi* geometry types.
// GML 2.1.2 uses gid (no namespace) as the element identifier.
// The prescan pass caches gid-carrying elements before the main decode,
// enabling both backward and forward xlink:href references to be resolved.

import (
	"io"
	"strings"
	"testing"

	"github.com/Kuroki-g/go-gml/core"
	gml "github.com/Kuroki-g/go-gml/gml2_1_2"
)

const xlinkNS = `xmlns:xlink="http://www.w3.org/1999/xlink"`

// readFirst reads only the first geometry; returns error if Next fails.
func readFirst(t *testing.T, xml string) (core.Geometry, error) {
	t.Helper()
	r := gml.NewReader(strings.NewReader(xml))
	return r.Next()
}

// readAllOrError collects all geometries, returning the first error encountered.
func readAllOrError(xml string) ([]core.Geometry, error) {
	r := gml.NewReader(strings.NewReader(xml))
	var out []core.Geometry
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return out, err
		}
		out = append(out, g)
	}
	return out, nil
}

// ---- MultiPoint xlink:href ----

func TestMultiPoint_xlink_backward(t *testing.T) {
	// Point (gid="p1") is defined BEFORE MultiPoint — backward reference.
	src := `<root ` + gml2NS + ` ` + xlinkNS + `>
		<gml:Point gid="p1"><gml:coordinates>1,2</gml:coordinates></gml:Point>
		<gml:MultiPoint srsName="EPSG:4326">
			<gml:pointMember xlink:href="#p1"/>
		</gml:MultiPoint>
	</root>`
	gs, err := readAllOrError(src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var mp core.MultiPoint
	for _, g := range gs {
		if v, ok := g.Value.(core.MultiPoint); ok {
			mp = v
			break
		}
	}
	if mp == nil {
		t.Fatal("no MultiPoint found")
	}
	if len(mp) != 1 {
		t.Fatalf("len=%d want 1", len(mp))
	}
	if !pointEq(mp[0], core.Point{1, 2}) {
		t.Errorf("mp[0]=%v want [1 2]", mp[0])
	}
}

func TestMultiPoint_xlink_forward(t *testing.T) {
	// MultiPoint appears BEFORE the Point definition — requires prescan.
	src := `<root ` + gml2NS + ` ` + xlinkNS + `>
		<gml:MultiPoint srsName="EPSG:4326">
			<gml:pointMember xlink:href="#p1"/>
		</gml:MultiPoint>
		<gml:Point gid="p1"><gml:coordinates>3,4</gml:coordinates></gml:Point>
	</root>`
	gs, err := readAllOrError(src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var mp core.MultiPoint
	for _, g := range gs {
		if v, ok := g.Value.(core.MultiPoint); ok {
			mp = v
			break
		}
	}
	if mp == nil {
		t.Fatal("no MultiPoint found")
	}
	if len(mp) != 1 {
		t.Fatalf("len=%d want 1", len(mp))
	}
	if !pointEq(mp[0], core.Point{3, 4}) {
		t.Errorf("mp[0]=%v want [3 4]", mp[0])
	}
}

func TestMultiPoint_xlink_unresolved(t *testing.T) {
	src := `<gml:MultiPoint ` + gml2NS + ` ` + xlinkNS + ` srsName="EPSG:4326">
		<gml:pointMember xlink:href="#missing"/>
	</gml:MultiPoint>`
	_, err := readFirst(t, src)
	if err == nil {
		t.Fatal("expected error for unresolved xlink:href, got nil")
	}
}

// ---- MultiLineString xlink:href ----

func TestMultiLineString_xlink_forward(t *testing.T) {
	// MultiLineString appears BEFORE the LineString definition — requires prescan.
	src := `<root ` + gml2NS + ` ` + xlinkNS + `>
		<gml:MultiLineString srsName="EPSG:4326">
			<gml:lineStringMember xlink:href="#ls1"/>
		</gml:MultiLineString>
		<gml:LineString gid="ls1"><gml:coordinates>0,0 1,1</gml:coordinates></gml:LineString>
	</root>`
	gs, err := readAllOrError(src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var ml core.MultiLineString
	for _, g := range gs {
		if v, ok := g.Value.(core.MultiLineString); ok {
			ml = v
			break
		}
	}
	if ml == nil {
		t.Fatal("no MultiLineString found")
	}
	if len(ml) != 1 || len(ml[0]) != 2 {
		t.Fatalf("got %v", ml)
	}
}

func TestMultiLineString_xlink_unresolved(t *testing.T) {
	src := `<gml:MultiLineString ` + gml2NS + ` ` + xlinkNS + ` srsName="EPSG:4326">
		<gml:lineStringMember xlink:href="#missing"/>
	</gml:MultiLineString>`
	_, err := readFirst(t, src)
	if err == nil {
		t.Fatal("expected error for unresolved xlink:href, got nil")
	}
}

// ---- MultiPolygon xlink:href ----

func TestMultiPolygon_xlink_forward(t *testing.T) {
	// MultiPolygon appears BEFORE the Polygon definition — requires prescan.
	ring := `<gml:coordinates>0,0 10,0 10,10 0,10 0,0</gml:coordinates>`
	src := `<root ` + gml2NS + ` ` + xlinkNS + `>
		<gml:MultiPolygon srsName="EPSG:4326">
			<gml:polygonMember xlink:href="#poly1"/>
		</gml:MultiPolygon>
		<gml:Polygon gid="poly1">
			<gml:outerBoundaryIs><gml:LinearRing>` + ring + `</gml:LinearRing></gml:outerBoundaryIs>
		</gml:Polygon>
	</root>`
	gs, err := readAllOrError(src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var mp core.MultiPolygon
	for _, g := range gs {
		if v, ok := g.Value.(core.MultiPolygon); ok {
			mp = v
			break
		}
	}
	if mp == nil {
		t.Fatal("no MultiPolygon found")
	}
	if len(mp) != 1 || len(mp[0]) != 1 || len(mp[0][0]) != 5 {
		t.Fatalf("got %v", mp)
	}
}

func TestMultiPolygon_xlink_unresolved(t *testing.T) {
	src := `<gml:MultiPolygon ` + gml2NS + ` ` + xlinkNS + ` srsName="EPSG:4326">
		<gml:polygonMember xlink:href="#missing"/>
	</gml:MultiPolygon>`
	_, err := readFirst(t, src)
	if err == nil {
		t.Fatal("expected error for unresolved xlink:href, got nil")
	}
}
