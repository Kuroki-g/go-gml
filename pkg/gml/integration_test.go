package gml

import (
	"io"
	"os"
	"testing"
)

// TestN03_2022Format parses a 2022 N03 file that uses standalone gml:Curve/LineStringSegment
// with the non-standard namespace http://schemas.opengis.net/gml/3.2.1/gml.xsd.
func TestN03_2022Format(t *testing.T) {
	f, err := os.Open("../../testdata/N03/N03-22_13_220101.xml")
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
	t.Logf("N03 2022 format: %d LineStrings, %d other geometries", curves, others)
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
