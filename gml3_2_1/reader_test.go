package gml3_2_1_test

import (
	"io"
	"strings"
	"testing"

	"github.com/Kuroki-g/go-gml/core"
	gml "github.com/Kuroki-g/go-gml/gml3_2_1"
)

const gml3NS = `xmlns:gml="http://www.opengis.net/gml/3.2"`
const xlinkNS = `xmlns:xlink="http://www.w3.org/1999/xlink"`

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

func TestParsePoint_pos(t *testing.T) {
	xml := `<gml:Point ` + gml3NS + ` srsName="EPSG:6668" srsDimension="2">
		<gml:pos>139.691667 35.689722</gml:pos>
	</gml:Point>`
	gs := readAll(t, xml)
	if len(gs) != 1 {
		t.Fatalf("expected 1 geometry, got %d", len(gs))
	}
	pt, ok := gs[0].Value.(core.Point)
	if !ok {
		t.Fatalf("expected Point, got %T", gs[0].Value)
	}
	if pt[0] != 139.691667 || pt[1] != 35.689722 {
		t.Errorf("got %v", pt)
	}
}

func TestParsePoint_3D(t *testing.T) {
	xml := `<gml:Point ` + gml3NS + ` srsDimension="3">
		<gml:pos>139.7 35.6 10.5</gml:pos>
	</gml:Point>`
	gs := readAll(t, xml)
	pt := gs[0].Value.(core.Point)
	if len(pt) != 3 {
		t.Fatalf("expected 3D point (len=3), got len=%d", len(pt))
	}
	if pt[0] != 139.7 || pt[1] != 35.6 || pt[2] != 10.5 {
		t.Errorf("got %v", pt)
	}
}

// ---- LineString ----

func TestParseLineString_posList(t *testing.T) {
	xml := `<gml:LineString ` + gml3NS + ` srsName="EPSG:6668" srsDimension="2">
		<gml:posList>139.7 35.6 139.8 35.7 139.9 35.8</gml:posList>
	</gml:LineString>`
	gs := readAll(t, xml)
	if len(gs) != 1 {
		t.Fatalf("expected 1, got %d", len(gs))
	}
	ls := gs[0].Value.(core.LineString)
	if len(ls) != 3 {
		t.Fatalf("len=%d want 3", len(ls))
	}
	if !pointEq(ls[0], core.Point{139.7, 35.6}) {
		t.Errorf("ls[0]=%v", ls[0])
	}
}

func TestParseLineString_posElements(t *testing.T) {
	xml := `<gml:LineString ` + gml3NS + ` srsDimension="2">
		<gml:pos>139.7 35.6</gml:pos>
		<gml:pos>139.8 35.7</gml:pos>
	</gml:LineString>`
	gs := readAll(t, xml)
	ls := gs[0].Value.(core.LineString)
	if len(ls) != 2 {
		t.Fatalf("len=%d want 2", len(ls))
	}
}

// ---- Polygon ----

func TestParsePolygon_exterior(t *testing.T) {
	xml := `<gml:Polygon ` + gml3NS + ` srsName="EPSG:6668" srsDimension="2">
		<gml:exterior>
			<gml:LinearRing>
				<gml:posList>139.7 35.6 139.8 35.6 139.8 35.7 139.7 35.7 139.7 35.6</gml:posList>
			</gml:LinearRing>
		</gml:exterior>
	</gml:Polygon>`
	gs := readAll(t, xml)
	if len(gs) != 1 {
		t.Fatalf("expected 1, got %d", len(gs))
	}
	poly := gs[0].Value.(core.Polygon)
	if len(poly) != 1 {
		t.Fatalf("rings=%d want 1", len(poly))
	}
	if len(poly[0]) != 5 {
		t.Fatalf("exterior points=%d want 5", len(poly[0]))
	}
}

func TestParsePolygon_withHole(t *testing.T) {
	xml := `<gml:Polygon ` + gml3NS + ` srsDimension="2">
		<gml:exterior>
			<gml:LinearRing>
				<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>
			</gml:LinearRing>
		</gml:exterior>
		<gml:interior>
			<gml:LinearRing>
				<gml:posList>2 2 8 2 8 8 2 8 2 2</gml:posList>
			</gml:LinearRing>
		</gml:interior>
	</gml:Polygon>`
	gs := readAll(t, xml)
	poly := gs[0].Value.(core.Polygon)
	if len(poly) != 2 {
		t.Fatalf("rings=%d want 2 (1 exterior + 1 hole)", len(poly))
	}
}

func TestParsePolygon_linearRing_pos(t *testing.T) {
	xml := `<gml:Polygon ` + gml3NS + ` srsName="EPSG:6668" srsDimension="2">
		<gml:exterior>
			<gml:LinearRing>
				<gml:pos>0 0</gml:pos>
				<gml:pos>10 0</gml:pos>
				<gml:pos>10 10</gml:pos>
				<gml:pos>0 10</gml:pos>
				<gml:pos>0 0</gml:pos>
			</gml:LinearRing>
		</gml:exterior>
	</gml:Polygon>`
	gs := readAll(t, xml)
	poly := gs[0].Value.(core.Polygon)
	if len(poly) != 1 || len(poly[0]) != 5 {
		t.Fatalf("rings=%d, exterior points=%d, want 1 and 5", len(poly), len(poly[0]))
	}
}

// ---- MultiPoint ----

func TestParseMultiPoint(t *testing.T) {
	xml := `<gml:MultiPoint ` + gml3NS + ` srsDimension="2">
		<gml:pointMember><gml:Point><gml:pos>139.7 35.6</gml:pos></gml:Point></gml:pointMember>
		<gml:pointMember><gml:Point><gml:pos>139.8 35.7</gml:pos></gml:Point></gml:pointMember>
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

// ---- MultiCurve / MultiLineString ----

func TestParseMultiCurve(t *testing.T) {
	xml := `<gml:MultiCurve ` + gml3NS + ` srsDimension="2">
		<gml:curveMember>
			<gml:LineString><gml:posList>0 0 1 1 2 0</gml:posList></gml:LineString>
		</gml:curveMember>
		<gml:curveMember>
			<gml:LineString><gml:posList>5 5 6 6</gml:posList></gml:LineString>
		</gml:curveMember>
	</gml:MultiCurve>`
	gs := readAll(t, xml)
	ml := gs[0].Value.(core.MultiLineString)
	if len(ml) != 2 {
		t.Fatalf("len=%d want 2", len(ml))
	}
	if len(ml[0]) != 3 {
		t.Errorf("ml[0] points=%d want 3", len(ml[0]))
	}
}

func TestParseMultiCurve_curveMembers_LineString(t *testing.T) {
	xml := `<gml:MultiCurve ` + gml3NS + ` srsDimension="2">
		<gml:curveMembers>
			<gml:LineString><gml:posList>0 0 1 1 2 0</gml:posList></gml:LineString>
			<gml:LineString><gml:posList>5 5 6 6</gml:posList></gml:LineString>
		</gml:curveMembers>
	</gml:MultiCurve>`
	gs := readAll(t, xml)
	ml := gs[0].Value.(core.MultiLineString)
	if len(ml) != 2 {
		t.Fatalf("lines=%d want 2", len(ml))
	}
}

func TestParseMultiCurve_curveMembers_Curve(t *testing.T) {
	xml := `<gml:MultiCurve ` + gml3NS + ` srsDimension="2">
		<gml:curveMembers>
			<gml:Curve>
				<gml:segments>
					<gml:LineStringSegment><gml:posList>0 0 1 1 2 0</gml:posList></gml:LineStringSegment>
				</gml:segments>
			</gml:Curve>
		</gml:curveMembers>
	</gml:MultiCurve>`
	gs := readAll(t, xml)
	ml := gs[0].Value.(core.MultiLineString)
	if len(ml) != 1 || len(ml[0]) != 3 {
		t.Fatalf("lines=%d pts=%d want 1,3", len(ml), len(ml[0]))
	}
}

func TestParseMultiCurve_mixed_member_and_members(t *testing.T) {
	xml := `<gml:MultiCurve ` + gml3NS + ` srsDimension="2">
		<gml:curveMember>
			<gml:LineString><gml:posList>0 0 1 1</gml:posList></gml:LineString>
		</gml:curveMember>
		<gml:curveMembers>
			<gml:LineString><gml:posList>2 2 3 3</gml:posList></gml:LineString>
			<gml:LineString><gml:posList>4 4 5 5</gml:posList></gml:LineString>
		</gml:curveMembers>
	</gml:MultiCurve>`
	gs := readAll(t, xml)
	ml := gs[0].Value.(core.MultiLineString)
	if len(ml) != 3 {
		t.Fatalf("lines=%d want 3", len(ml))
	}
}

// ---- MultiSurface ----

func TestParseMultiSurface(t *testing.T) {
	ring := `<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>`
	xml := `<gml:MultiSurface ` + gml3NS + ` srsDimension="2">
		<gml:surfaceMember>
			<gml:Polygon>
				<gml:exterior><gml:LinearRing>` + ring + `</gml:LinearRing></gml:exterior>
			</gml:Polygon>
		</gml:surfaceMember>
		<gml:surfaceMember>
			<gml:Polygon>
				<gml:exterior><gml:LinearRing>` + ring + `</gml:LinearRing></gml:exterior>
			</gml:Polygon>
		</gml:surfaceMember>
	</gml:MultiSurface>`
	gs := readAll(t, xml)
	mp := gs[0].Value.(core.MultiPolygon)
	if len(mp) != 2 {
		t.Fatalf("len=%d want 2", len(mp))
	}
}

func TestParseMultiSurface_surfaceMembers_Polygon(t *testing.T) {
	ring := `<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>`
	xml := `<gml:MultiSurface ` + gml3NS + ` srsDimension="2">
		<gml:surfaceMembers>
			<gml:Polygon>
				<gml:exterior><gml:LinearRing>` + ring + `</gml:LinearRing></gml:exterior>
			</gml:Polygon>
			<gml:Polygon>
				<gml:exterior><gml:LinearRing>` + ring + `</gml:LinearRing></gml:exterior>
			</gml:Polygon>
		</gml:surfaceMembers>
	</gml:MultiSurface>`
	gs := readAll(t, xml)
	mp := gs[0].Value.(core.MultiPolygon)
	if len(mp) != 2 {
		t.Fatalf("polys=%d want 2", len(mp))
	}
}

func TestParseMultiSurface_mixed_member_and_members(t *testing.T) {
	ring := `<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>`
	xml := `<gml:MultiSurface ` + gml3NS + ` srsDimension="2">
		<gml:surfaceMember>
			<gml:Polygon>
				<gml:exterior><gml:LinearRing>` + ring + `</gml:LinearRing></gml:exterior>
			</gml:Polygon>
		</gml:surfaceMember>
		<gml:surfaceMembers>
			<gml:Polygon>
				<gml:exterior><gml:LinearRing>` + ring + `</gml:LinearRing></gml:exterior>
			</gml:Polygon>
		</gml:surfaceMembers>
	</gml:MultiSurface>`
	gs := readAll(t, xml)
	mp := gs[0].Value.(core.MultiPolygon)
	if len(mp) != 2 {
		t.Fatalf("polys=%d want 2", len(mp))
	}
}

// ---- Envelope ----

func TestParseEnvelope_lowerUpper(t *testing.T) {
	xml := `<gml:Envelope ` + gml3NS + ` srsName="EPSG:6668" srsDimension="2">
		<gml:lowerCorner>139.0 35.0</gml:lowerCorner>
		<gml:upperCorner>140.0 36.0</gml:upperCorner>
	</gml:Envelope>`
	gs := readAll(t, xml)
	b := gs[0].Value.(core.Bound)
	if !pointEq(b.Min, core.Point{139.0, 35.0}) {
		t.Errorf("Min=%v", b.Min)
	}
	if !pointEq(b.Max, core.Point{140.0, 36.0}) {
		t.Errorf("Max=%v", b.Max)
	}
}

func TestParseEnvelope_pos(t *testing.T) {
	xml := `<gml:Envelope ` + gml3NS + ` srsDimension="2">
		<gml:pos>139.0 35.0</gml:pos>
		<gml:pos>140.0 36.0</gml:pos>
	</gml:Envelope>`
	gs := readAll(t, xml)
	b := gs[0].Value.(core.Bound)
	if !pointEq(b.Min, core.Point{139.0, 35.0}) || !pointEq(b.Max, core.Point{140.0, 36.0}) {
		t.Errorf("got %v", b)
	}
}

// ---- Curve ----

func TestParseCurve_posList(t *testing.T) {
	xml := `<gml:Curve ` + gml3NS + ` srsName="urn:ogc:def:crs:EPSG::6668" srsDimension="2">
		<gml:segments>
			<gml:LineStringSegment>
				<gml:posList>35.74504 139.71950 35.74590 139.72000 35.74596 139.72004</gml:posList>
			</gml:LineStringSegment>
		</gml:segments>
	</gml:Curve>`
	gs := readAll(t, xml)
	if len(gs) != 1 {
		t.Fatalf("expected 1 geometry, got %d", len(gs))
	}
	ls, ok := gs[0].Value.(core.LineString)
	if !ok {
		t.Fatalf("expected LineString, got %T", gs[0].Value)
	}
	if len(ls) != 3 {
		t.Fatalf("len=%d want 3", len(ls))
	}
	if !pointEq(ls[0], core.Point{35.74504, 139.71950}) {
		t.Errorf("ls[0]=%v", ls[0])
	}
}

// ---- Surface ----

func TestParseSurface_linearRing(t *testing.T) {
	xml := `<gml:Surface ` + gml3NS + ` srsName="EPSG:6668" srsDimension="2">
		<gml:patches>
			<gml:PolygonPatch>
				<gml:exterior>
					<gml:LinearRing>
						<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>
					</gml:LinearRing>
				</gml:exterior>
			</gml:PolygonPatch>
		</gml:patches>
	</gml:Surface>`
	gs := readAll(t, xml)
	if len(gs) != 1 {
		t.Fatalf("expected 1 geometry, got %d", len(gs))
	}
	poly := gs[0].Value.(core.Polygon)
	if len(poly) != 1 || len(poly[0]) != 5 {
		t.Fatalf("rings=%d, exterior points=%d", len(poly), len(poly[0]))
	}
}

func TestParseSurface_ringCurveMember(t *testing.T) {
	xml := `<gml:Surface ` + gml3NS + ` srsName="urn:ogc:def:crs:EPSG::6668" srsDimension="2">
		<gml:patches>
			<gml:PolygonPatch>
				<gml:exterior>
					<gml:Ring>
						<gml:curveMember>
							<gml:Curve>
								<gml:segments><gml:LineStringSegment><gml:posList>0 0 5 0 10 0</gml:posList></gml:LineStringSegment></gml:segments>
							</gml:Curve>
						</gml:curveMember>
						<gml:curveMember>
							<gml:Curve>
								<gml:segments><gml:LineStringSegment><gml:posList>10 0 10 5 10 10</gml:posList></gml:LineStringSegment></gml:segments>
							</gml:Curve>
						</gml:curveMember>
						<gml:curveMember>
							<gml:Curve>
								<gml:segments><gml:LineStringSegment><gml:posList>10 10 5 10 0 10</gml:posList></gml:LineStringSegment></gml:segments>
							</gml:Curve>
						</gml:curveMember>
						<gml:curveMember>
							<gml:Curve>
								<gml:segments><gml:LineStringSegment><gml:posList>0 10 0 5 0 0</gml:posList></gml:LineStringSegment></gml:segments>
							</gml:Curve>
						</gml:curveMember>
					</gml:Ring>
				</gml:exterior>
			</gml:PolygonPatch>
		</gml:patches>
	</gml:Surface>`
	gs := readAll(t, xml)
	poly := gs[0].Value.(core.Polygon)
	if len(poly) != 1 {
		t.Fatalf("rings=%d want 1", len(poly))
	}
	// 4 curves × 3 pts each, shared endpoints elided: 3+2+2+2 = 9
	if len(poly[0]) != 9 {
		t.Fatalf("exterior points=%d want 9", len(poly[0]))
	}
}

func TestParseSurface_xlinkHref_via_orientablecurve(t *testing.T) {
	xml := `<root ` + gml3NS + ` ` + xlinkNS + `>
		<gml:Curve gml:id="cv1_0">
			<gml:segments>
				<gml:LineStringSegment>
					<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>
				</gml:LineStringSegment>
			</gml:segments>
		</gml:Curve>
		<gml:OrientableCurve gml:id="_cv1_0" orientation="+">
			<gml:baseCurve xlink:href="#cv1_0"/>
		</gml:OrientableCurve>
		<gml:Surface>
			<gml:patches>
				<gml:PolygonPatch>
					<gml:exterior>
						<gml:Ring>
							<gml:curveMember xlink:href="#_cv1_0"/>
						</gml:Ring>
					</gml:exterior>
				</gml:PolygonPatch>
			</gml:patches>
		</gml:Surface>
	</root>`
	gs := readAll(t, xml)
	// Expect: 1 LineString (from Curve) + 1 Polygon (from Surface)
	if len(gs) != 2 {
		t.Fatalf("expected 2 geometries (LineString+Polygon), got %d", len(gs))
	}
	poly, ok := gs[1].Value.(core.Polygon)
	if !ok {
		t.Fatalf("expected Polygon, got %T", gs[1].Value)
	}
	if len(poly) == 0 || len(poly[0]) == 0 {
		t.Fatalf("Polygon has empty coordinates: xlink:href resolution failed")
	}
	if len(poly[0]) != 5 {
		t.Errorf("exterior points=%d want 5", len(poly[0]))
	}
}

// ---- CompositeCurve ----

func TestParseCompositeCurve_inlineCurves(t *testing.T) {
	xml := `<gml:CompositeCurve ` + gml3NS + ` srsDimension="2">
		<gml:curveMember>
			<gml:Curve>
				<gml:segments><gml:LineStringSegment><gml:posList>0 0 5 0 10 0</gml:posList></gml:LineStringSegment></gml:segments>
			</gml:Curve>
		</gml:curveMember>
		<gml:curveMember>
			<gml:Curve>
				<gml:segments><gml:LineStringSegment><gml:posList>10 0 10 5 10 10</gml:posList></gml:LineStringSegment></gml:segments>
			</gml:Curve>
		</gml:curveMember>
	</gml:CompositeCurve>`
	gs := readAll(t, xml)
	if len(gs) != 1 {
		t.Fatalf("expected 1 geometry, got %d", len(gs))
	}
	ls, ok := gs[0].Value.(core.LineString)
	if !ok {
		t.Fatalf("expected LineString, got %T", gs[0].Value)
	}
	// 2 segments × 3 pts, shared endpoint elided: 3+2 = 5
	if len(ls) != 5 {
		t.Fatalf("points=%d want 5", len(ls))
	}
}

func TestParseCompositeCurve_inlineLineString(t *testing.T) {
	xml := `<gml:CompositeCurve ` + gml3NS + ` srsDimension="2">
		<gml:curveMember>
			<gml:LineString><gml:posList>0 0 5 5 10 10</gml:posList></gml:LineString>
		</gml:curveMember>
	</gml:CompositeCurve>`
	gs := readAll(t, xml)
	ls := gs[0].Value.(core.LineString)
	if len(ls) != 3 {
		t.Fatalf("points=%d want 3", len(ls))
	}
}

func TestParseCompositeCurve_nested(t *testing.T) {
	xml := `<gml:CompositeCurve ` + gml3NS + ` srsDimension="2">
		<gml:curveMember>
			<gml:CompositeCurve>
				<gml:curveMember>
					<gml:Curve>
						<gml:segments><gml:LineStringSegment><gml:posList>0 0 5 0</gml:posList></gml:LineStringSegment></gml:segments>
					</gml:Curve>
				</gml:curveMember>
			</gml:CompositeCurve>
		</gml:curveMember>
		<gml:curveMember>
			<gml:Curve>
				<gml:segments><gml:LineStringSegment><gml:posList>5 0 10 0</gml:posList></gml:LineStringSegment></gml:segments>
			</gml:Curve>
		</gml:curveMember>
	</gml:CompositeCurve>`
	gs := readAll(t, xml)
	ls := gs[0].Value.(core.LineString)
	// inner: [0,0 5,0], outer adds [10,0] → 3 pts
	if len(ls) != 3 {
		t.Fatalf("points=%d want 3", len(ls))
	}
}

// ---- CompositeSurface ----

func TestParseCompositeSurface_twoPolygons(t *testing.T) {
	ring1 := `<gml:posList>0 0 5 0 5 5 0 5 0 0</gml:posList>`
	ring2 := `<gml:posList>5 0 10 0 10 5 5 5 5 0</gml:posList>`
	xml := `<gml:CompositeSurface ` + gml3NS + ` srsDimension="2">
		<gml:surfaceMember>
			<gml:Polygon>
				<gml:exterior><gml:LinearRing>` + ring1 + `</gml:LinearRing></gml:exterior>
			</gml:Polygon>
		</gml:surfaceMember>
		<gml:surfaceMember>
			<gml:Polygon>
				<gml:exterior><gml:LinearRing>` + ring2 + `</gml:LinearRing></gml:exterior>
			</gml:Polygon>
		</gml:surfaceMember>
	</gml:CompositeSurface>`
	gs := readAll(t, xml)
	mp, ok := gs[0].Value.(core.MultiPolygon)
	if !ok {
		t.Fatalf("expected MultiPolygon, got %T", gs[0].Value)
	}
	// Two member polygons → two polygons in the MultiPolygon
	if len(mp) != 2 {
		t.Fatalf("polygons=%d want 2", len(mp))
	}
}

func TestParseCompositeSurface_surfaceMembers(t *testing.T) {
	xml := `<gml:CompositeSurface ` + gml3NS + ` srsDimension="2">
		<gml:surfaceMember>
			<gml:Surface>
				<gml:patches>
					<gml:PolygonPatch>
						<gml:exterior>
							<gml:LinearRing>
								<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>
							</gml:LinearRing>
						</gml:exterior>
					</gml:PolygonPatch>
				</gml:patches>
			</gml:Surface>
		</gml:surfaceMember>
	</gml:CompositeSurface>`
	gs := readAll(t, xml)
	mp, ok := gs[0].Value.(core.MultiPolygon)
	if !ok {
		t.Fatalf("expected MultiPolygon, got %T", gs[0].Value)
	}
	if len(mp) != 1 || len(mp[0]) != 1 || len(mp[0][0]) != 5 {
		t.Fatalf("polygons=%d, rings=%d, exterior points=%d", len(mp), len(mp[0]), len(mp[0][0]))
	}
}

// ---- OrientableSurface ----

func TestParseOrientableSurface_inlinePolygon(t *testing.T) {
	xml := `<gml:OrientableSurface ` + gml3NS + ` orientation="+" srsDimension="2">
		<gml:baseSurface>
			<gml:Polygon>
				<gml:exterior>
					<gml:LinearRing>
						<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>
					</gml:LinearRing>
				</gml:exterior>
			</gml:Polygon>
		</gml:baseSurface>
	</gml:OrientableSurface>`
	gs := readAll(t, xml)
	if len(gs) != 1 {
		t.Fatalf("expected 1 geometry, got %d", len(gs))
	}
	poly := gs[0].Value.(core.Polygon)
	if len(poly) != 1 || len(poly[0]) != 5 {
		t.Fatalf("rings=%d, exterior points=%d", len(poly), len(poly[0]))
	}
}

func TestParseOrientableSurface_inlineSurface(t *testing.T) {
	xml := `<gml:OrientableSurface ` + gml3NS + ` orientation="+" srsDimension="2">
		<gml:baseSurface>
			<gml:Surface>
				<gml:patches>
					<gml:PolygonPatch>
						<gml:exterior>
							<gml:LinearRing>
								<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>
							</gml:LinearRing>
						</gml:exterior>
					</gml:PolygonPatch>
				</gml:patches>
			</gml:Surface>
		</gml:baseSurface>
	</gml:OrientableSurface>`
	gs := readAll(t, xml)
	poly := gs[0].Value.(core.Polygon)
	if len(poly) != 1 || len(poly[0]) != 5 {
		t.Fatalf("rings=%d, exterior points=%d", len(poly), len(poly[0]))
	}
}

// ---- xlink:href forward reference ----

func TestParseCompositeSurface_forwardRef(t *testing.T) {
	xml := `<root ` + gml3NS + ` ` + xlinkNS + `>
		<gml:CompositeSurface gml:id="CS1" srsDimension="2">
			<gml:surfaceMember xlink:href="#P1"/>
		</gml:CompositeSurface>
		<gml:Polygon gml:id="P1">
			<gml:exterior>
				<gml:LinearRing>
					<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>
				</gml:LinearRing>
			</gml:exterior>
		</gml:Polygon>
	</root>`
	gs := readAll(t, xml)
	// Expect: CS1 (MultiPolygon) + P1 (Polygon)
	if len(gs) != 2 {
		t.Fatalf("expected 2 geometries, got %d", len(gs))
	}
	cs1, ok := gs[0].Value.(core.MultiPolygon)
	if !ok {
		t.Fatalf("gs[0]: expected MultiPolygon, got %T", gs[0].Value)
	}
	if len(cs1) == 0 || len(cs1[0]) == 0 {
		t.Fatal("CS1: forward reference to P1 was not resolved (empty MultiPolygon)")
	}
	if len(cs1[0][0]) != 5 {
		t.Errorf("CS1: exterior points=%d want 5", len(cs1[0][0]))
	}
}

func TestParseCompositeSurface_chainForwardRef(t *testing.T) {
	// CS1 -> CS2 (forward ref), CS2 -> P1 (forward ref)
	xml := `<root ` + gml3NS + ` ` + xlinkNS + `>
		<gml:CompositeSurface gml:id="CS1" srsDimension="2">
			<gml:surfaceMember xlink:href="#CS2"/>
		</gml:CompositeSurface>
		<gml:CompositeSurface gml:id="CS2" srsDimension="2">
			<gml:surfaceMember xlink:href="#P1"/>
		</gml:CompositeSurface>
		<gml:Polygon gml:id="P1">
			<gml:exterior>
				<gml:LinearRing>
					<gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList>
				</gml:LinearRing>
			</gml:exterior>
		</gml:Polygon>
	</root>`
	gs := readAll(t, xml)
	if len(gs) != 3 {
		t.Fatalf("expected 3 geometries, got %d", len(gs))
	}
	cs1, ok := gs[0].Value.(core.MultiPolygon)
	if !ok {
		t.Fatalf("gs[0]: expected MultiPolygon, got %T", gs[0].Value)
	}
	if len(cs1) == 0 || len(cs1[0]) == 0 {
		t.Fatal("CS1: chain forward reference to CS2->P1 was not resolved")
	}
	if len(cs1[0][0]) != 5 {
		t.Errorf("CS1: exterior points=%d want 5", len(cs1[0][0]))
	}
}

// ---- streaming ----

func TestParseStream_multipleGeometries(t *testing.T) {
	xml := `<root ` + gml3NS + `>
		<gml:Point><gml:pos>139.7 35.6</gml:pos></gml:Point>
		<gml:LineString srsDimension="2"><gml:posList>0 0 1 1</gml:posList></gml:LineString>
		<gml:Envelope>
			<gml:lowerCorner>0 0</gml:lowerCorner>
			<gml:upperCorner>1 1</gml:upperCorner>
		</gml:Envelope>
	</root>`
	gs := readAll(t, xml)
	if len(gs) != 3 {
		t.Fatalf("expected 3 geometries, got %d", len(gs))
	}
	if _, ok := gs[0].Value.(core.Point); !ok {
		t.Errorf("gs[0] expected Point, got %T", gs[0].Value)
	}
	if _, ok := gs[1].Value.(core.LineString); !ok {
		t.Errorf("gs[1] expected LineString, got %T", gs[1].Value)
	}
	if _, ok := gs[2].Value.(core.Bound); !ok {
		t.Errorf("gs[2] expected Bound, got %T", gs[2].Value)
	}
}

func TestParseStream_empty(t *testing.T) {
	gs := readAll(t, `<root/>`)
	if len(gs) != 0 {
		t.Errorf("expected 0 geometries, got %d", len(gs))
	}
}
