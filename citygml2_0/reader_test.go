package citygml2_0_test

import (
	"io"
	"os"
	"strings"
	"testing"

	citygml2_0 "github.com/Kuroki-g/go-gml/citygml2_0"
	"github.com/Kuroki-g/go-gml/core"
	"github.com/Kuroki-g/go-gml/gml3_1_1"
)

const plateauGML = "../testdata/13105_bunkyo-ku_pref_2023_citygml_2_op/udx/bldg/53394548_bldg_6697_op.gml"

// newGMLDecoder returns a core.Decoder backed by gml3_1_1.
// The Reader's Decode method creates a fresh sub-reader per call, so the
// underlying io.ReadSeeker is never accessed after construction.
func newGMLDecoder() core.Decoder {
	return gml3_1_1.NewReader(strings.NewReader(""))
}

func openGML(t *testing.T, path string) *os.File {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	t.Cleanup(func() { f.Close() })
	return f
}

func TestNext_Lod0RoofEdge(t *testing.T) {
	f := openGML(t, plateauGML)
	r := citygml2_0.NewReader(f, newGMLDecoder())

	b, err := r.Next()
	if err != nil {
		t.Fatalf("Next: %v", err)
	}
	if b.ID == "" {
		t.Error("Building.ID is empty")
	}
	if b.Lod0RoofEdge == nil {
		t.Fatal("Lod0RoofEdge is nil")
	}
	mp, ok := b.Lod0RoofEdge.Value.(core.MultiPolygon)
	if !ok {
		t.Fatalf("Lod0RoofEdge.Value type = %T, want core.MultiPolygon", b.Lod0RoofEdge.Value)
	}
	if len(mp) == 0 {
		t.Error("MultiPolygon is empty")
	}
	t.Logf("Building ID=%s, Lod0RoofEdge polygons=%d", b.ID, len(mp))
}

func TestNext_Lod1Solid(t *testing.T) {
	f := openGML(t, plateauGML)
	r := citygml2_0.NewReader(f, newGMLDecoder())

	var b *citygml2_0.Building
	for {
		building, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Next: %v", err)
		}
		if building.Lod1Solid != nil {
			b = building
			break
		}
	}
	if b == nil {
		t.Fatal("no building with Lod1Solid found")
	}
	solid, ok := b.Lod1Solid.Value.(core.Solid)
	if !ok {
		t.Fatalf("Lod1Solid.Value type = %T, want core.Solid", b.Lod1Solid.Value)
	}
	if len(solid.Exterior) == 0 {
		t.Error("Solid.Exterior is empty")
	}
	t.Logf("Building ID=%s, Lod1Solid Exterior polygons=%d", b.ID, len(solid.Exterior))
}

func TestNext_AllBuildings(t *testing.T) {
	f := openGML(t, plateauGML)
	r := citygml2_0.NewReader(f, newGMLDecoder())

	count := 0
	for {
		b, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Next (building %d): %v", count, err)
		}
		if b.ID == "" {
			t.Errorf("building %d: ID is empty", count)
		}
		count++
	}
	if count == 0 {
		t.Error("no buildings found")
	}
	t.Logf("found %d buildings", count)
}
