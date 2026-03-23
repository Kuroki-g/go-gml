package gml2_1_2_test

import (
	"io"
	"strings"
	"testing"

	"github.com/Kuroki-g/go-gml/core"
	gml "github.com/Kuroki-g/go-gml/gml2_1_2"
)

const gml2NS = `xmlns:gml="http://www.opengis.net/gml"`

func readAll(t *testing.T, xml string) []core.Geometry {
	t.Helper()
	r := gml.NewReader(strings.NewReader(xml))
	var out []core.Geometry
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		out = append(out, g)
	}
	return out
}

func pointEq(a, b core.Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ---- Point ----

func TestParsePoint_coordinates(t *testing.T) {
	xml := `<gml:Point ` + gml2NS + `>
		<gml:coordinates>139.7,35.6</gml:coordinates>
	</gml:Point>`
	gs := readAll(t, xml)
	if len(gs) != 1 {
		t.Fatalf("expected 1 geometry, got %d", len(gs))
	}
	pt := gs[0].Value.(core.Point)
	if pt[0] != 139.7 || pt[1] != 35.6 {
		t.Errorf("got %v", pt)
	}
}

func TestParsePoint_coord(t *testing.T) {
	xml := `<gml:Point ` + gml2NS + `>
		<gml:coord><gml:X>139.7</gml:X><gml:Y>35.6</gml:Y></gml:coord>
	</gml:Point>`
	gs := readAll(t, xml)
	if len(gs) != 1 {
		t.Fatalf("expected 1 geometry, got %d", len(gs))
	}
	pt := gs[0].Value.(core.Point)
	if pt[0] != 139.7 || pt[1] != 35.6 {
		t.Errorf("got %v", pt)
	}
}

// ---- LineString ----

func TestParseLineString_coordinates(t *testing.T) {
	xml := `<gml:LineString ` + gml2NS + `>
		<gml:coordinates>139.7,35.6 139.8,35.7</gml:coordinates>
	</gml:LineString>`
	gs := readAll(t, xml)
	ls := gs[0].Value.(core.LineString)
	if len(ls) != 2 {
		t.Fatalf("len=%d want 2", len(ls))
	}
}

func TestParseLineString_coord(t *testing.T) {
	xml := `<gml:LineString ` + gml2NS + `>
		<gml:coord><gml:X>0</gml:X><gml:Y>0</gml:Y></gml:coord>
		<gml:coord><gml:X>1</gml:X><gml:Y>1</gml:Y></gml:coord>
	</gml:LineString>`
	gs := readAll(t, xml)
	ls := gs[0].Value.(core.LineString)
	if len(ls) != 2 {
		t.Fatalf("len=%d want 2", len(ls))
	}
}

// ---- Polygon ----

func TestParsePolygon_outerBoundaryIs(t *testing.T) {
	xml := `<gml:Polygon ` + gml2NS + `>
		<gml:outerBoundaryIs>
			<gml:LinearRing>
				<gml:coordinates>0,0 10,0 10,10 0,10 0,0</gml:coordinates>
			</gml:LinearRing>
		</gml:outerBoundaryIs>
	</gml:Polygon>`
	gs := readAll(t, xml)
	poly := gs[0].Value.(core.Polygon)
	if len(poly) != 1 {
		t.Fatalf("rings=%d want 1", len(poly))
	}
	if len(poly[0]) != 5 {
		t.Fatalf("exterior points=%d want 5", len(poly[0]))
	}
}

func TestParsePolygon_withHole(t *testing.T) {
	xml := `<gml:Polygon ` + gml2NS + `>
		<gml:outerBoundaryIs>
			<gml:LinearRing>
				<gml:coordinates>0,0 10,0 10,10 0,10 0,0</gml:coordinates>
			</gml:LinearRing>
		</gml:outerBoundaryIs>
		<gml:innerBoundaryIs>
			<gml:LinearRing>
				<gml:coordinates>2,2 8,2 8,8 2,8 2,2</gml:coordinates>
			</gml:LinearRing>
		</gml:innerBoundaryIs>
	</gml:Polygon>`
	gs := readAll(t, xml)
	poly := gs[0].Value.(core.Polygon)
	if len(poly) != 2 {
		t.Fatalf("rings=%d want 2 (1 exterior + 1 hole)", len(poly))
	}
}

// ---- MultiPoint ----

func TestParseMultiPoint(t *testing.T) {
	xml := `<gml:MultiPoint ` + gml2NS + `>
		<gml:pointMember><gml:Point><gml:coordinates>139.7,35.6</gml:coordinates></gml:Point></gml:pointMember>
		<gml:pointMember><gml:Point><gml:coordinates>139.8,35.7</gml:coordinates></gml:Point></gml:pointMember>
	</gml:MultiPoint>`
	gs := readAll(t, xml)
	mp := gs[0].Value.(core.MultiPoint)
	if len(mp) != 2 {
		t.Fatalf("len=%d want 2", len(mp))
	}
	if !pointEq(mp[0], core.Point{139.7, 35.6}) {
		t.Errorf("mp[0]=%v", mp[0])
	}
}

// ---- MultiLineString ----

func TestParseMultiLineString(t *testing.T) {
	xml := `<gml:MultiLineString ` + gml2NS + `>
		<gml:lineStringMember>
			<gml:LineString><gml:coordinates>0,0 1,1</gml:coordinates></gml:LineString>
		</gml:lineStringMember>
		<gml:lineStringMember>
			<gml:LineString><gml:coordinates>2,2 3,3</gml:coordinates></gml:LineString>
		</gml:lineStringMember>
	</gml:MultiLineString>`
	gs := readAll(t, xml)
	ml := gs[0].Value.(core.MultiLineString)
	if len(ml) != 2 {
		t.Fatalf("len=%d want 2", len(ml))
	}
}

// ---- MultiPolygon ----

func TestParseMultiPolygon(t *testing.T) {
	ring := `<gml:coordinates>0,0 10,0 10,10 0,10 0,0</gml:coordinates>`
	xml := `<gml:MultiPolygon ` + gml2NS + `>
		<gml:polygonMember>
			<gml:Polygon>
				<gml:outerBoundaryIs><gml:LinearRing>` + ring + `</gml:LinearRing></gml:outerBoundaryIs>
			</gml:Polygon>
		</gml:polygonMember>
		<gml:polygonMember>
			<gml:Polygon>
				<gml:outerBoundaryIs><gml:LinearRing>` + ring + `</gml:LinearRing></gml:outerBoundaryIs>
			</gml:Polygon>
		</gml:polygonMember>
	</gml:MultiPolygon>`
	gs := readAll(t, xml)
	mp := gs[0].Value.(core.MultiPolygon)
	if len(mp) != 2 {
		t.Fatalf("len=%d want 2", len(mp))
	}
}

// ---- Box ----

func TestParseBox_coordinates(t *testing.T) {
	xml := `<gml:Box ` + gml2NS + ` srsName="EPSG:4326">
		<gml:coordinates>139.0,35.0 140.0,36.0</gml:coordinates>
	</gml:Box>`
	gs := readAll(t, xml)
	b := gs[0].Value.(core.Bound)
	if !pointEq(b.Min, core.Point{139.0, 35.0}) {
		t.Errorf("Min=%v", b.Min)
	}
	if !pointEq(b.Max, core.Point{140.0, 36.0}) {
		t.Errorf("Max=%v", b.Max)
	}
}

func TestParseBox_coord(t *testing.T) {
	xml := `<gml:Box ` + gml2NS + `>
		<gml:coord><gml:X>139.0</gml:X><gml:Y>35.0</gml:Y></gml:coord>
		<gml:coord><gml:X>140.0</gml:X><gml:Y>36.0</gml:Y></gml:coord>
	</gml:Box>`
	gs := readAll(t, xml)
	b := gs[0].Value.(core.Bound)
	if !pointEq(b.Min, core.Point{139.0, 35.0}) || !pointEq(b.Max, core.Point{140.0, 36.0}) {
		t.Errorf("got %v", b)
	}
}

// ---- streaming ----

func TestParseStream_empty(t *testing.T) {
	gs := readAll(t, `<root/>`)
	if len(gs) != 0 {
		t.Errorf("expected 0 geometries, got %d", len(gs))
	}
}
