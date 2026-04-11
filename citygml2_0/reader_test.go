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

const boundedByXML = `<?xml version="1.0" encoding="UTF-8"?>
<CityModel xmlns="http://www.opengis.net/citygml/2.0"
           xmlns:bldg="http://www.opengis.net/citygml/building/2.0"
           xmlns:gml="http://www.opengis.net/gml">
  <cityObjectMember>
    <bldg:Building gml:id="bld1">
      <bldg:boundedBy>
        <bldg:WallSurface gml:id="ws1">
          <bldg:lod2MultiSurface>
            <gml:MultiSurface>
              <gml:surfaceMember>
                <gml:Polygon>
                  <gml:exterior>
                    <gml:LinearRing>
                      <gml:posList srsDimension="3">0 0 0 1 0 0 1 1 0 0 0 0</gml:posList>
                    </gml:LinearRing>
                  </gml:exterior>
                </gml:Polygon>
              </gml:surfaceMember>
            </gml:MultiSurface>
          </bldg:lod2MultiSurface>
        </bldg:WallSurface>
      </bldg:boundedBy>
      <bldg:boundedBy>
        <bldg:RoofSurface gml:id="rs1">
          <bldg:lod2MultiSurface>
            <gml:MultiSurface>
              <gml:surfaceMember>
                <gml:Polygon>
                  <gml:exterior>
                    <gml:LinearRing>
                      <gml:posList srsDimension="3">0 0 1 1 0 1 1 1 1 0 0 1</gml:posList>
                    </gml:LinearRing>
                  </gml:exterior>
                </gml:Polygon>
              </gml:surfaceMember>
            </gml:MultiSurface>
          </bldg:lod2MultiSurface>
        </bldg:RoofSurface>
      </bldg:boundedBy>
    </bldg:Building>
  </cityObjectMember>
</CityModel>`

func TestNext_BoundedBy(t *testing.T) {
	r := citygml2_0.NewReader(strings.NewReader(boundedByXML), newGMLDecoder())
	b, err := r.Next()
	if err != nil {
		t.Fatalf("Next: %v", err)
	}
	if len(b.BoundedBy) != 2 {
		t.Fatalf("BoundedBy count = %d, want 2", len(b.BoundedBy))
	}

	wall := b.BoundedBy[0]
	if wall.ID != "ws1" {
		t.Errorf("BoundedBy[0].ID = %q, want %q", wall.ID, "ws1")
	}
	if wall.SurfaceType != citygml2_0.SurfaceTypeWallSurface {
		t.Errorf("BoundedBy[0].SurfaceType = %q, want %q", wall.SurfaceType, citygml2_0.SurfaceTypeWallSurface)
	}
	if wall.Lod2MultiSurface == nil {
		t.Error("BoundedBy[0].Lod2MultiSurface is nil")
	}

	roof := b.BoundedBy[1]
	if roof.ID != "rs1" {
		t.Errorf("BoundedBy[1].ID = %q, want %q", roof.ID, "rs1")
	}
	if roof.SurfaceType != citygml2_0.SurfaceTypeRoofSurface {
		t.Errorf("BoundedBy[1].SurfaceType = %q, want %q", roof.SurfaceType, citygml2_0.SurfaceTypeRoofSurface)
	}
	if roof.Lod2MultiSurface == nil {
		t.Error("BoundedBy[1].Lod2MultiSurface is nil")
	}
	t.Logf("BoundedBy[0]: %s/%s Lod2MultiSurface=%T", wall.SurfaceType, wall.ID, wall.Lod2MultiSurface.Value)
	t.Logf("BoundedBy[1]: %s/%s Lod2MultiSurface=%T", roof.SurfaceType, roof.ID, roof.Lod2MultiSurface.Value)
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
