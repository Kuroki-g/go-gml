package main

import (
	"os"
	"path/filepath"
	"testing"
)

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

// ---- collectSequenceFields ----

func TestCollectSequenceFields_nil(t *testing.T) {
	if got := collectSequenceFields(nil); got != nil {
		t.Errorf("nil seq should return nil, got %v", got)
	}
}

func TestCollectSequenceFields_namedElement(t *testing.T) {
	seq := &xsdSequence{
		Elements: []xsdElement{
			{Name: "pos", Type: "gml:DirectPositionType", MinOccurs: "0"},
		},
	}
	fields := collectSequenceFields(seq)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
	f := fields[0]
	if f.LocalName != "pos" {
		t.Errorf("LocalName = %q, want %q", f.LocalName, "pos")
	}
	if f.TypeRef != "gml:DirectPositionType" {
		t.Errorf("TypeRef = %q, want %q", f.TypeRef, "gml:DirectPositionType")
	}
	if f.MinOccurs != "0" {
		t.Errorf("MinOccurs = %q, want %q", f.MinOccurs, "0")
	}
}

func TestCollectSequenceFields_refElement(t *testing.T) {
	seq := &xsdSequence{
		Elements: []xsdElement{
			{Ref: "gml:pointMember", MaxOccurs: "unbounded"},
		},
	}
	fields := collectSequenceFields(seq)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
	f := fields[0]
	if f.Ref != "gml:pointMember" {
		t.Errorf("Ref = %q, want %q", f.Ref, "gml:pointMember")
	}
	if f.MaxOccurs != "unbounded" {
		t.Errorf("MaxOccurs = %q, want %q", f.MaxOccurs, "unbounded")
	}
}

func TestCollectSequenceFields_inlineComplexType(t *testing.T) {
	seq := &xsdSequence{
		Elements: []xsdElement{
			{Name: "nested", ComplexType: &xsdComplexType{Name: ""}},
		},
	}
	fields := collectSequenceFields(seq)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
	if fields[0].TypeRef != "__inline__" {
		t.Errorf("TypeRef = %q, want %q", fields[0].TypeRef, "__inline__")
	}
}

func TestCollectSequenceFields_nestedChoice(t *testing.T) {
	seq := &xsdSequence{
		Choices: []xsdSequence{
			{
				Elements: []xsdElement{
					{Name: "a", Type: "xs:string"},
					{Name: "b", Type: "xs:integer"},
				},
			},
		},
	}
	fields := collectSequenceFields(seq)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields from nested choice, got %d", len(fields))
	}
}

func TestCollectSequenceFields_groupRef(t *testing.T) {
	seq := &xsdSequence{
		Groups: []xsdGroupRef{
			{Ref: "gml:StandardObjectProperties", MinOccurs: "0"},
		},
	}
	fields := collectSequenceFields(seq)
	if len(fields) != 1 {
		t.Fatalf("expected 1 group-ref field, got %d", len(fields))
	}
	if fields[0].GroupRef != "gml:StandardObjectProperties" {
		t.Errorf("GroupRef = %q, want %q", fields[0].GroupRef, "gml:StandardObjectProperties")
	}
}

// ---- parseSchema (integration with temp file) ----

func writeTemp(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

const simpleXSD = `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test"
        elementFormDefault="qualified">

  <simpleType name="CodeType">
    <restriction base="string"/>
  </simpleType>

  <simpleType name="MeasureList">
    <list itemType="double"/>
  </simpleType>

  <complexType name="PointType">
    <sequence>
      <element name="pos" type="tns:CodeType" minOccurs="0"/>
    </sequence>
    <attribute name="srsName" type="anyURI"/>
  </complexType>

  <complexType name="ValueType">
    <simpleContent>
      <extension base="double">
        <attribute name="uom" type="string" use="required"/>
      </extension>
    </simpleContent>
  </complexType>

  <element name="Point" type="tns:PointType"/>
</schema>`

func TestParseSchema_simpleXSD(t *testing.T) {
	path := writeTemp(t, simpleXSD)
	s, err := parseSchema(path)
	if err != nil {
		t.Fatalf("parseSchema: %v", err)
	}

	if s.TargetNamespace != "http://example.com/test" {
		t.Errorf("TargetNamespace = %q", s.TargetNamespace)
	}
	if len(s.ComplexTypes) != 2 {
		t.Errorf("ComplexTypes = %d, want 2", len(s.ComplexTypes))
	}
	if len(s.SimpleTypes) != 2 {
		t.Errorf("SimpleTypes = %d, want 2", len(s.SimpleTypes))
	}

	// Global element registered
	if s.Elements["Point"] != "tns:PointType" {
		t.Errorf("Elements[Point] = %q, want %q", s.Elements["Point"], "tns:PointType")
	}

	// NSMap captures xmlns:tns
	if s.NSMap["tns"] != "http://example.com/test" {
		t.Errorf("NSMap[tns] = %q, want %q", s.NSMap["tns"], "http://example.com/test")
	}
}

func TestParseSchema_simpleContent(t *testing.T) {
	path := writeTemp(t, simpleXSD)
	s, err := parseSchema(path)
	if err != nil {
		t.Fatal(err)
	}

	var valueType *ComplexType
	for i := range s.ComplexTypes {
		if s.ComplexTypes[i].Name == "ValueType" {
			valueType = &s.ComplexTypes[i]
		}
	}
	if valueType == nil {
		t.Fatal("ValueType not found")
	}

	// First raw field should be chardata
	if len(valueType.RawFields) == 0 {
		t.Fatal("no raw fields in ValueType")
	}
	if !valueType.RawFields[0].IsChar {
		t.Errorf("first field should be chardata, got %+v", valueType.RawFields[0])
	}
}

func TestParseSchema_notFound(t *testing.T) {
	_, err := parseSchema("/nonexistent/path/test.xsd")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestParseSchema_badXML(t *testing.T) {
	path := writeTemp(t, `<schema><NOT VALID`)
	_, err := parseSchema(path)
	if err == nil {
		t.Error("expected error for invalid XML")
	}
}

// ---- convertAttributes ----

func TestConvertAttributes_required(t *testing.T) {
	attrs := []xsdAttribute{
		{Name: "srsName", Type: "anyURI", Use: "required"},
	}
	fields := convertAttributes(attrs)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
	f := fields[0]
	if !f.IsAttr {
		t.Error("IsAttr should be true")
	}
	if f.LocalName != "srsName" {
		t.Errorf("LocalName = %q", f.LocalName)
	}
	// Use="required" is stored in MinOccurs
	if f.MinOccurs != "required" {
		t.Errorf("MinOccurs = %q, want %q", f.MinOccurs, "required")
	}
}

func TestConvertAttributes_ref(t *testing.T) {
	attrs := []xsdAttribute{
		{Ref: "xlink:href"},
	}
	fields := convertAttributes(attrs)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
	if fields[0].AttrRef != "xlink:href" {
		t.Errorf("AttrRef = %q, want %q", fields[0].AttrRef, "xlink:href")
	}
	if !fields[0].IsAttr {
		t.Error("IsAttr should be true for ref attribute")
	}
}
