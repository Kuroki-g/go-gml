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
		// GML 3.1.1 abstract element names: "_Xxx" → "AbstractXxx"
		{"_Curve", "AbstractCurve"},
		{"_Surface", "AbstractSurface"},
		{"_Ring", "AbstractRing"},
		{"_CurveSegment", "AbstractCurveSegment"},
		{"_Geometry", "AbstractGeometry"},
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
	src, err := Generate(nil, "mypkg", false, false, nil, "", "")
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
	src, err := Generate(types, "gml", false, false, nil, "", "")
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
	src, err := Generate(types, "gml", true, false, nil, "", "")
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
	src, err := Generate(types, "pkg", false, false, nil, "", "")
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
	src, err := Generate(types, "gml", false, false, nil, "", "")
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
	src, err := Generate(types, "gml", false, true, nil, "", "")
	if err != nil {
		t.Fatalf("Generate(withDoc=true): %v", err)
	}
	if !strings.Contains(src, "// Direct position of the point.") {
		t.Errorf("expected field doc comment, got:\n%s", src)
	}

	// withDoc=false: no comment
	src2, err := Generate(types, "gml", false, false, nil, "", "")
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
	src, err := Generate(types, "pkg", false, true, nil, "", "")
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

// ---- --map-namespace ----

func TestGenerate_mapNS_fieldKept(t *testing.T) {
	// A field whose TypeNS is in mapNS should be kept and qualified with the alias.
	types := []*ComplexType{
		{
			Name: "WallSurfaceType",
			Fields: []Field{
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
	// Field must be present (not dropped)
	if !strings.Contains(src, "Lod2MultiSurface") {
		t.Errorf("field should be kept, got:\n%s", src)
	}
	// Type must be qualified with alias
	if !strings.Contains(src, "gml3_1_1gen.MultiSurfacePropertyType") {
		t.Errorf("expected qualified type gml3_1_1gen.MultiSurfacePropertyType, got:\n%s", src)
	}
}

func TestGenerate_mapNS_importAdded(t *testing.T) {
	// When a field uses a mapped namespace, the import block should be generated.
	types := []*ComplexType{
		{
			Name: "FooType",
			Fields: []Field{
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
	// Import should not appear when no field actually uses the mapped namespace.
	types := []*ComplexType{
		{
			Name: "FooType",
			Fields: []Field{
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
	// Two different namespaces mapping to aliases that collide should return an error.
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

// ---- deriveAlias ----

func TestDeriveAlias(t *testing.T) {
	tests := []struct {
		pkgPath string
		want    string
	}{
		// last segment is "generated" → use preceding segment + "gen", keep underscores
		{"github.com/Kuroki-g/go-gml/gml3_1_1/generated", "gml3_1_1gen"},
		{"github.com/Kuroki-g/go-gml/gml3_2_1/generated", "gml3_2_1gen"},
		{"github.com/Kuroki-g/go-gml/gml2_1_2/generated", "gml2_1_2gen"},
		// last segment is not "generated" → use last segment
		{"github.com/foo/mypkg", "mypkg"},
		// hyphens and dots stripped
		{"github.com/foo/my-pkg", "mypkg"},
	}
	for _, tt := range tests {
		got := deriveAlias(tt.pkgPath)
		if got != tt.want {
			t.Errorf("deriveAlias(%q) = %q, want %q", tt.pkgPath, got, tt.want)
		}
	}
}

// ---- generated header ----

func TestGenerate_header(t *testing.T) {
	src, err := Generate(nil, "pkg", false, false, nil, "path/to/geometry.xsd", "v1.2.3")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	want := "// Code generated by xsd2go-lite v1.2.3 from geometry.xsd. DO NOT EDIT."
	if !strings.Contains(src, want) {
		t.Errorf("expected header %q, got:\n%s", want, src)
	}
}

func TestGenerate_noHeaderWhenEmpty(t *testing.T) {
	src, err := Generate(nil, "pkg", false, false, nil, "", "")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if strings.Contains(src, "DO NOT EDIT") {
		t.Errorf("unexpected DO NOT EDIT when xsdSource and version are empty:\n%s", src)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
