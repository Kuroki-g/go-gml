package xsdlite

import (
	"strings"
	"testing"
)

func TestGenerate_sortedOutput(t *testing.T) {
	// Output should be sorted by type name for determinism.
	types := []*ComplexType{
		{Name: "Zebra", Fields: []Field{{GoName: "V", GoType: "string", XMLTag: "v"}}},
		{Name: "Alpha", Fields: []Field{{GoName: "V", GoType: "string", XMLTag: "v"}}},
		{Name: "Mango", Fields: []Field{{GoName: "V", GoType: "string", XMLTag: "v"}}},
	}
	src, err := Generate(types, "pkg", false, false, nil, "", "")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	alphaPos := strings.Index(src, "Alpha")
	mangoPos := strings.Index(src, "Mango")
	zebraPos := strings.Index(src, "Zebra")
	if !(alphaPos < mangoPos && mangoPos < zebraPos) {
		t.Errorf("types not sorted: Alpha=%d Mango=%d Zebra=%d", alphaPos, mangoPos, zebraPos)
	}
}

func TestGenerate_pointerField(t *testing.T) {
	types := []*ComplexType{
		{
			Name: "PolygonType",
			Fields: []Field{
				{GoName: "Exterior", GoType: "*RingType", XMLTag: "http://example.com exterior", Omit: true},
			},
		},
	}
	src, err := Generate(types, "gml", false, false, nil, "", "")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(src, "*RingType") {
		t.Error("pointer type should appear")
	}
	if !strings.Contains(src, "omitempty") {
		t.Error("optional pointer should have omitempty")
	}
}

func TestGenerate_namespaceInTag(t *testing.T) {
	types := []*ComplexType{
		{
			Name: "FooType",
			Fields: []Field{
				{GoName: "Bar", GoType: "string", XMLTag: "http://www.opengis.net/gml/3.2 bar"},
			},
		},
	}
	src, err := Generate(types, "gml", false, false, nil, "", "")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(src, `xml:"http://www.opengis.net/gml/3.2 bar"`) {
		t.Errorf("namespace URI should be in tag, src:\n%s", src)
	}
}
