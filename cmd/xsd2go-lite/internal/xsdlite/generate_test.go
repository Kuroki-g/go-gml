package xsdlite

import (
	"strings"
	"testing"
)

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
