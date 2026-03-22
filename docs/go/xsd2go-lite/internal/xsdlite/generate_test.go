package xsdlite

import (
	"strings"
	"testing"
)

// ---- goName ----

func TestGoName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", "_"},
		{"pointType", "PointType"},
		{"point_type", "PointType"},
		{"point-type", "PointType"},
		{"PolygonType", "PolygonType"},
		{"multiSurface", "MultiSurface"},
		{"multi_surface_type", "MultiSurfaceType"},
		// digit prefix
		{"3DGeometry", "_3DGeometry"},
		// reserved word
		{"type", "TypeField"},
		{"map", "MapField"},
		{"range", "RangeField"},
		// dot separator
		{"gml.geometry", "GmlGeometry"},
		// already exported
		{"AbstractGeometryType", "AbstractGeometryType"},
		// single char
		{"x", "X"},
		// all underscores → empty parts → "_"
		{"_", "_"},
	}

	for _, tt := range tests {
		got := goName(tt.in)
		if got != tt.want {
			t.Errorf("goName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

// ---- avoidReserved ----

func TestAvoidReserved(t *testing.T) {
	reserved := []string{
		"Break", "Case", "Chan", "Const", "Continue", "Default", "Defer",
		"Else", "Fallthrough", "For", "Func", "Go", "Goto", "If", "Import",
		"Interface", "Map", "Package", "Range", "Return", "Select", "Struct",
		"Switch", "Type", "Var",
	}
	for _, r := range reserved {
		got := avoidReserved(r)
		if got != r+"Field" {
			t.Errorf("avoidReserved(%q) = %q, want %q", r, got, r+"Field")
		}
	}
	// Non-reserved should pass through unchanged.
	for _, ok := range []string{"Point", "Line", "Geometry", "Hello"} {
		if got := avoidReserved(ok); got != ok {
			t.Errorf("avoidReserved(%q) = %q, want %q (unchanged)", ok, got, ok)
		}
	}
}

// ---- buildTagSuffix ----

func TestBuildTagSuffix(t *testing.T) {
	tests := []struct {
		name string
		f    Field
		want string
	}{
		{
			name: "optional element → omitempty",
			f:    Field{Omit: true, IsChar: false, Slice: false},
			want: ",omitempty",
		},
		{
			name: "required element → no suffix",
			f:    Field{Omit: false},
			want: "",
		},
		{
			name: "optional slice → no omitempty (slice)",
			f:    Field{Omit: true, Slice: true},
			want: "",
		},
		{
			name: "chardata → no omitempty",
			f:    Field{Omit: true, IsChar: true},
			want: "",
		},
		{
			name: "optional attr → omitempty",
			f:    Field{Omit: true, IsAttr: true},
			want: ",omitempty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTagSuffix(tt.f)
			if got != tt.want {
				t.Errorf("buildTagSuffix = %q, want %q", got, tt.want)
			}
		})
	}
}

// ---- Generate ----

func TestGenerate_empty(t *testing.T) {
	src, err := Generate(nil, "mypkg", false, false)
	if err != nil {
		t.Fatalf("Generate(nil): %v", err)
	}
	if !strings.HasPrefix(src, "package mypkg") {
		t.Errorf("expected package declaration, got: %q", src[:min(50, len(src))])
	}
}

func TestGenerate_singleStruct(t *testing.T) {
	types := []*ComplexType{
		{
			Name: "PointType",
			Fields: []Field{
				{GoName: "Value", GoType: "string", XMLTag: ",chardata", IsChar: true},
				{GoName: "SrsName", GoType: "string", XMLTag: "srsName,attr", IsAttr: true, Omit: true},
			},
		},
	}
	src, err := Generate(types, "gml", false, false)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	if !strings.Contains(src, "type PointType struct") {
		t.Error("missing struct declaration")
	}
	if !strings.Contains(src, `xml:",chardata"`) {
		t.Error("missing chardata tag")
	}
	if !strings.Contains(src, `xml:"srsName,attr,omitempty"`) {
		t.Error("missing attr omitempty tag")
	}
}

func TestGenerate_skipAbstract(t *testing.T) {
	types := []*ComplexType{
		{Name: "AbstractBase", Abstract: true, Fields: []Field{}},
		{Name: "ConcreteType", Fields: []Field{
			{GoName: "Id", GoType: "string", XMLTag: "id,attr", IsAttr: true},
		}},
	}
	src, err := Generate(types, "gml", true, false)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if strings.Contains(src, "AbstractBase") {
		t.Error("abstract type should be skipped")
	}
	if !strings.Contains(src, "ConcreteType") {
		t.Error("concrete type should be present")
	}
}

func TestGenerate_emptyName(t *testing.T) {
	// Types with empty names should be skipped.
	types := []*ComplexType{
		{Name: "", Fields: []Field{{GoName: "X", GoType: "string", XMLTag: "x"}}},
		{Name: "Valid", Fields: []Field{{GoName: "Y", GoType: "int", XMLTag: "y"}}},
	}
	src, err := Generate(types, "pkg", false, false)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(src, "type Valid struct") {
		t.Error("Valid type should be present")
	}
}

func TestGenerate_sliceField(t *testing.T) {
	types := []*ComplexType{
		{
			Name: "MultiType",
			Fields: []Field{
				{GoName: "Members", GoType: "[]MemberType", XMLTag: "http://example.com member", Slice: true},
			},
		},
	}
	src, err := Generate(types, "gml", false, false)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(src, "[]MemberType") {
		t.Error("slice type should appear in output")
	}
	// Slices should NOT have omitempty even if Omit=true (covered in buildTagSuffix test)
}

func TestGenerate_sortedOutput(t *testing.T) {
	// Output should be sorted by type name for determinism.
	types := []*ComplexType{
		{Name: "Zebra", Fields: []Field{{GoName: "V", GoType: "string", XMLTag: "v"}}},
		{Name: "Alpha", Fields: []Field{{GoName: "V", GoType: "string", XMLTag: "v"}}},
		{Name: "Mango", Fields: []Field{{GoName: "V", GoType: "string", XMLTag: "v"}}},
	}
	src, err := Generate(types, "pkg", false, false)
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
	src, err := Generate(types, "gml", false, false)
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
	src, err := Generate(types, "gml", false, false)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(src, `xml:"http://www.opengis.net/gml/3.2 bar"`) {
		t.Errorf("namespace URI should be in tag, src:\n%s", src)
	}
}

func TestGenerate_withDoc(t *testing.T) {
	types := []*ComplexType{
		{
			Name: "PointType",
			Fields: []Field{
				{GoName: "Pos", GoType: "string", XMLTag: "pos", Doc: "Direct position of the point."},
				{GoName: "Id", GoType: "string", XMLTag: "id,attr", IsAttr: true, Omit: true},
			},
		},
	}

	// withDoc=true: field comment should appear
	src, err := Generate(types, "gml", false, true)
	if err != nil {
		t.Fatalf("Generate(withDoc=true): %v", err)
	}
	if !strings.Contains(src, "// Direct position of the point.") {
		t.Errorf("expected field doc comment, got:\n%s", src)
	}

	// withDoc=false: no comment
	src2, err := Generate(types, "gml", false, false)
	if err != nil {
		t.Fatalf("Generate(withDoc=false): %v", err)
	}
	if strings.Contains(src2, "// Direct position") {
		t.Errorf("expected no doc comment when withDoc=false, got:\n%s", src2)
	}
}

func TestGenerate_withDoc_multiline(t *testing.T) {
	types := []*ComplexType{
		{
			Name: "FooType",
			Fields: []Field{
				{GoName: "Val", GoType: "string", XMLTag: "val", Doc: "Line one.\nLine two.\nLine three."},
			},
		},
	}
	src, err := Generate(types, "pkg", false, true)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(src, "// Line one.") {
		t.Errorf("expected first line of doc, got:\n%s", src)
	}
	if !strings.Contains(src, "// Line two.") {
		t.Errorf("expected second line of doc, got:\n%s", src)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
