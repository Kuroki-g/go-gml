package xsdlite

import "testing"

// ---- xsBuiltinGoType ----

func TestXsBuiltinGoType(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// string-like
		{"string", "string"},
		{"anyURI", "string"},
		{"NCName", "string"},
		{"token", "string"},
		{"base64Binary", "string"},
		{"dateTime", "string"},
		{"date", "string"},
		// bool
		{"boolean", "bool"},
		// int family
		{"integer", "int"},
		{"positiveInteger", "int"},
		{"nonNegativeInteger", "int"},
		{"long", "int"},
		{"int", "int"},
		{"unsignedByte", "int"},
		// float
		{"double", "float64"},
		{"decimal", "float64"},
		{"float", "float32"},
		// anyType
		{"anyType", "string"},
		{"anySimpleType", "string"},
		// prefixed (xs:string など)
		{"xs:string", "string"},
		{"xs:integer", "int"},
		{"xs:boolean", "bool"},
		{"xs:double", "float64"},
		// unknown → ""
		{"MyCustomType", ""},
		{"doubleList", ""},
		{"", ""},
	}

	for _, tt := range tests {
		got := xsBuiltinGoType(tt.in)
		if got != tt.want {
			t.Errorf("xsBuiltinGoType(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

// ---- convertSimpleType ----

func TestConvertSimpleType(t *testing.T) {
	tests := []struct {
		name     string
		input    xsdSimpleType
		wantKind string
		wantType string
	}{
		{
			name:     "list → string",
			input:    xsdSimpleType{Name: "doubleList", List: &xsdList{ItemType: "double"}},
			wantKind: "list",
			wantType: "string",
		},
		{
			name:     "union → string",
			input:    xsdSimpleType{Name: "NilReasonType", Union: &xsdUnion{MemberTypes: "string anyURI"}},
			wantKind: "union",
			wantType: "string",
		},
		{
			name:     "restriction builtin",
			input:    xsdSimpleType{Name: "SignType", Restriction: &xsdSRestrict{Base: "xs:string"}},
			wantKind: "restriction",
			wantType: "string",
		},
		{
			name: "restriction unknown base → empty GoType",
			// Issue: non-builtin base yields "" GoType
			input:    xsdSimpleType{Name: "Derived", Restriction: &xsdSRestrict{Base: "gml:SomeType"}},
			wantKind: "restriction",
			wantType: "", // known limitation
		},
		{
			name:     "no variant → alias string",
			input:    xsdSimpleType{Name: "Bare"},
			wantKind: "alias",
			wantType: "string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertSimpleType(tt.input)
			if got.Kind != tt.wantKind {
				t.Errorf("Kind = %q, want %q", got.Kind, tt.wantKind)
			}
			if got.GoType != tt.wantType {
				t.Errorf("GoType = %q, want %q", got.GoType, tt.wantType)
			}
		})
	}
}
