package gml3_2_1_test

import (
	"io"
	"os"
	"testing"

	"github.com/Kuroki-g/go-gml/core"
	gml "github.com/Kuroki-g/go-gml/gml3_2_1"
)

// TestN03_2024Format parses a 2024 N03 file that uses gml:Surface/PolygonPatch.
func TestN03_2024Format(t *testing.T) {
	f, err := os.Open("../testdata/N03/N03-20240101_13.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := gml.NewReader(f)
	var polygons, others int
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if _, ok := g.Value.(core.Polygon); ok {
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

// TestN03_newFormat parses a 2025 N03 file that uses gml:Surface/PolygonPatch/Ring/curveMember.
func TestN03_newFormat(t *testing.T) {
	f, err := os.Open("../testdata/N03/N03-20250101_13.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := gml.NewReader(f)
	var polygons, others int
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if _, ok := g.Value.(core.Polygon); ok {
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

// TestG04 parses a G04 (標高・傾斜度メッシュ) file.
// Old format: gml:Grid at document root, domainSet uses xlink:href="#grid".
func TestG04(t *testing.T) {
	f, err := os.Open("../testdata/G04/G04-a-11_5339-jgd_GML/G04-a-11_5339-jgd.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := gml.NewReader(f)
	var coverages []core.GridCoverage
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if cov, ok := g.Value.(core.GridCoverage); ok {
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
func TestL03Old(t *testing.T) {
	f, err := os.Open("../testdata/L03/L03-a-06_5339-jgd_GML/L03-a-06_5339-jgd.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := gml.NewReader(f)
	var coverages []core.GridCoverage
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if cov, ok := g.Value.(core.GridCoverage); ok {
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
func TestL03New(t *testing.T) {
	f, err := os.Open("../testdata/L03/L03-a-09_5339-jgd_GML/L03-a-09_5339.xml")
	if err != nil {
		t.Skip("testdata not available:", err)
	}
	defer f.Close()

	r := gml.NewReader(f)
	var coverages []core.GridCoverage
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if cov, ok := g.Value.(core.GridCoverage); ok {
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
