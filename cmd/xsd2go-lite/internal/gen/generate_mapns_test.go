package gen

import (
	"strings"
	"testing"

	"xsd2go-lite/internal/parse"
)

// ---- --map-namespace ----

func TestGenerate_mapNS_fieldKept(t *testing.T) {
	// A field whose TypeNS is in mapNS should be kept and qualified with the alias.
	types := []*parse.ComplexType{
		{
			Name: "WallSurfaceType",
			Fields: []parse.Field{
				{GoName: "Lod2MultiSurface", GoType: "*MultiSurfacePropertyType",
					XMLTag: "http://www.opengis.net/citygml/building/2.0 lod2MultiSurface",
					TypeNS: "http://www.opengis.net/gml", Omit: true},
			},
		},
	}
	mapNS := map[string]string{
		"http://www.opengis.net/gml": "github.com/Kuroki-g/go-gml/gml3_1_1/generated",
	}
	src, err := Generate(types, "generated", false, false, mapNS, "", "")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(src, "Lod2MultiSurface") {
		t.Errorf("field should be kept, got:\n%s", src)
	}
	if !strings.Contains(src, "gml3_1_1gen.MultiSurfacePropertyType") {
		t.Errorf("expected qualified type gml3_1_1gen.MultiSurfacePropertyType, got:\n%s", src)
	}
}

func TestGenerate_mapNS_importAdded(t *testing.T) {
	types := []*parse.ComplexType{
		{
			Name: "FooType",
			Fields: []parse.Field{
				{GoName: "Geom", GoType: "*PointType",
					XMLTag: "http://www.opengis.net/gml/3.2 geom",
					TypeNS: "http://www.opengis.net/gml/3.2", Omit: true},
			},
		},
	}
	mapNS := map[string]string{
		"http://www.opengis.net/gml/3.2": "github.com/Kuroki-g/go-gml/gml3_2_1/generated",
	}
	src, err := Generate(types, "pkg", false, false, mapNS, "", "")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(src, `gml3_2_1gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"`) {
		t.Errorf("expected import statement, got:\n%s", src)
	}
}

func TestGenerate_mapNS_noImportWhenUnused(t *testing.T) {
	types := []*parse.ComplexType{
		{
			Name: "FooType",
			Fields: []parse.Field{
				{GoName: "Name", GoType: "string", XMLTag: "name", TypeNS: ""},
			},
		},
	}
	mapNS := map[string]string{
		"http://www.opengis.net/gml": "github.com/Kuroki-g/go-gml/gml3_1_1/generated",
	}
	src, err := Generate(types, "pkg", false, false, mapNS, "", "")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if strings.Contains(src, "import") {
		t.Errorf("unexpected import when no field uses mapped namespace:\n%s", src)
	}
}

func TestGenerate_mapNS_aliasCollision(t *testing.T) {
	mapNS := map[string]string{
		"http://ns1": "github.com/foo/mypkg/generated",
		"http://ns2": "github.com/bar/mypkg/generated",
	}
	_, err := Generate(nil, "pkg", false, false, mapNS, "", "")
	if err == nil {
		t.Error("expected alias collision error, got nil")
	}
	if !strings.Contains(err.Error(), "alias collision") {
		t.Errorf("expected 'alias collision' in error message, got: %v", err)
	}
}
