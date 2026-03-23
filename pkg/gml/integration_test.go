package gml

import (
	"io"
	"os"
	"testing"
)

// TestN03_2022Format parses a 2022 N03 file that uses standalone gml:Curve/LineStringSegment
// with the non-standard namespace http://schemas.opengis.net/gml/3.2.1/gml.xsd.
// gml:Surface references gml:Curve via xlink:href; Polygons must have non-empty coordinates.
func TestN03_2022Format(t *testing.T) {
	f, err := os.Open("../../testdata/N03/N03-22_13_220101.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := NewReader(f)
	var curves, polygons, emptyPolygons int
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		switch v := g.Value.(type) {
		case LineString:
			curves++
		case Polygon:
			polygons++
			if len(v) == 0 || len(v[0]) == 0 {
				emptyPolygons++
			}
		}
	}
	if curves == 0 {
		t.Errorf("expected LineStrings (from gml:Curve), got 0")
	}
	if polygons == 0 {
		t.Errorf("expected Polygons (from gml:Surface via xlink:href), got 0")
	}
	// N03-22 の大部分の Surface は xlink:href が dangling reference（データバグ）のため
	// 空 Polygon を返す。これは XSD 準拠の正しい挙動。
	// OrientableCurve が存在する sf0/sf128/sf129/sf140 のみ解決される。
	if emptyPolygons == polygons {
		t.Errorf("all %d Polygons are empty: expected at least some to resolve via OrientableCurve", polygons)
	}
	t.Logf("N03 2022 format: %d LineStrings, %d Polygons (%d empty)", curves, polygons, emptyPolygons)
}

// TestN03_2024Format parses a 2024 N03 file that uses gml:Surface/PolygonPatch
// with the non-standard namespace http://schemas.opengis.net/gml/3.2.1/gml.xsd.
func TestN03_2024Format(t *testing.T) {
	f, err := os.Open("../../testdata/N03/N03-20240101_13.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := NewReader(f)
	var polygons, others int
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if _, ok := g.Value.(Polygon); ok {
			polygons++
		} else {
			others++
		}
	}
	if polygons == 0 {
		t.Errorf("expected at least one Polygon (from gml:Surface), got 0")
	}
	t.Logf("N03 2024 format: %d Polygons, %d other geometries", polygons, others)
}

// TestN03_oldFormat parses a pre-2020 N03 file that uses standalone gml:Curve elements.
func TestN03_oldFormat(t *testing.T) {
	f, err := os.Open("../../testdata/N03/N03-001001_13-g.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := NewReader(f)
	var curves, others int
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if _, ok := g.Value.(LineString); ok {
			curves++
		} else {
			others++
		}
	}
	if curves == 0 {
		t.Errorf("expected at least one LineString (from gml:Curve), got 0")
	}
	t.Logf("N03 old format: %d LineStrings, %d other geometries", curves, others)
}

// TestW09 parses a W09 (湖沼) file that uses gml:Curve/LineStringSegment and gml:Polygon.
func TestW09(t *testing.T) {
	f, err := os.Open("../../testdata/W09/W09-05_GML/W09-05-g.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := NewReader(f)
	var curves, polygons int
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		switch g.Value.(type) {
		case LineString:
			curves++
		case Polygon:
			polygons++
		}
	}
	if curves == 0 {
		t.Errorf("expected LineStrings (from gml:Curve), got 0")
	}
	if polygons == 0 {
		t.Errorf("expected Polygons (from gml:Polygon), got 0")
	}
	t.Logf("W09: %d LineStrings, %d Polygons", curves, polygons)
}

// TestG04 parses a G04 (標高・傾斜度メッシュ) file.
// Old format: gml:Grid at document root, domainSet uses xlink:href="#grid".
// Expects GridCoverage with Low=[0,0], High=[9,9] and non-empty tuples.
func TestG04(t *testing.T) {
	f, err := os.Open("../../testdata/G04/G04-a-11_5339-jgd_GML/G04-a-11_5339-jgd.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := NewReader(f)
	var coverages []GridCoverage
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if cov, ok := g.Value.(GridCoverage); ok {
			coverages = append(coverages, cov)
		}
	}
	if len(coverages) == 0 {
		t.Fatal("expected GridCoverage results, got 0")
	}
	first := coverages[0]
	if len(first.Low) != 2 || first.Low[0] != 0 || first.Low[1] != 0 {
		t.Errorf("Low = %v, want [0 0]", first.Low)
	}
	if len(first.High) != 2 || first.High[0] != 9 || first.High[1] != 9 {
		t.Errorf("High = %v, want [9 9]", first.High)
	}
	if first.Tuples == "" {
		t.Error("Tuples is empty")
	}
	t.Logf("G04: %d GridCoverages, first Low=%v High=%v", len(coverages), first.Low, first.High)
}

// TestL03Old parses a L03 旧形式 (〜2006) file.
// Old format: gml:Grid at document root, domainSet uses xlink:href="#grid".
// Expects GridCoverage with Low=[0,0], High=[9,9] and non-empty tuples.
func TestL03Old(t *testing.T) {
	f, err := os.Open("../../testdata/L03/L03-a-06_5339-jgd_GML/L03-a-06_5339-jgd.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := NewReader(f)
	var coverages []GridCoverage
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if cov, ok := g.Value.(GridCoverage); ok {
			coverages = append(coverages, cov)
		}
	}
	if len(coverages) == 0 {
		t.Fatal("expected GridCoverage results, got 0")
	}
	first := coverages[0]
	if len(first.Low) != 2 || first.Low[0] != 0 || first.Low[1] != 0 {
		t.Errorf("Low = %v, want [0 0]", first.Low)
	}
	if len(first.High) != 2 || first.High[0] != 9 || first.High[1] != 9 {
		t.Errorf("High = %v, want [9 9]", first.High)
	}
	if first.Tuples == "" {
		t.Error("Tuples is empty")
	}
	t.Logf("L03 old: %d GridCoverages, first Low=%v High=%v", len(coverages), first.Low, first.High)
}

// TestL03New parses a L03 新形式 (2009〜) file.
// New format: gml:Grid is inline inside gml:domainSet.
// Expects GridCoverage with Low=[0,0], High=[79,79] and non-empty tuples.
func TestL03New(t *testing.T) {
	f, err := os.Open("../../testdata/L03/L03-a-09_5339-jgd_GML/L03-a-09_5339.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := NewReader(f)
	var coverages []GridCoverage
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if cov, ok := g.Value.(GridCoverage); ok {
			coverages = append(coverages, cov)
		}
	}
	if len(coverages) == 0 {
		t.Fatal("expected GridCoverage results, got 0")
	}
	first := coverages[0]
	if len(first.Low) != 2 || first.Low[0] != 0 || first.Low[1] != 0 {
		t.Errorf("Low = %v, want [0 0]", first.Low)
	}
	if len(first.High) != 2 || first.High[0] != 79 || first.High[1] != 79 {
		t.Errorf("High = %v, want [79 79]", first.High)
	}
	if first.Tuples == "" {
		t.Error("Tuples is empty")
	}
	t.Logf("L03 new: %d GridCoverages, first Low=%v High=%v", len(coverages), first.Low, first.High)
}

// TestN03_newFormat parses a 2025 N03 file that uses gml:Surface/PolygonPatch/Ring/curveMember.
func TestN03_newFormat(t *testing.T) {
	f, err := os.Open("../../testdata/N03/N03-20250101_13.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := NewReader(f)
	var polygons, others int
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if _, ok := g.Value.(Polygon); ok {
			polygons++
		} else {
			others++
		}
	}
	if polygons == 0 {
		t.Errorf("expected at least one Polygon (from gml:Surface), got 0")
	}
	t.Logf("N03 new format: %d Polygons, %d other geometries", polygons, others)
}
