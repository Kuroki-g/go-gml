package xsdlite

import "testing"

// ---- buildXMLTag ----

func TestBuildXMLTag(t *testing.T) {
	tests := []struct {
		ns, local string
		isAttr    bool
		want      string
	}{
		{"", "name", false, "name"},
		{"http://example.com", "name", false, "http://example.com name"},
		{"", "srsName", true, "srsName,attr"},
		{"http://example.com", "id", true, "http://example.com id,attr"},
		{"", ",chardata", false, ",chardata"},
	}
	for _, tt := range tests {
		got := buildXMLTag(tt.ns, tt.local, tt.isAttr)
		if got != tt.want {
			t.Errorf("buildXMLTag(%q, %q, %v) = %q, want %q", tt.ns, tt.local, tt.isAttr, got, tt.want)
		}
	}
}

// ---- deduplicateFields ----

func TestDeduplicateFields(t *testing.T) {
	fields := []Field{
		{GoName: "Id", GoType: "string"},
		{GoName: "Name", GoType: "string"},
		{GoName: "Id", GoType: "int"}, // duplicate — should be dropped
		{GoName: "Value", GoType: "float64"},
	}
	got := deduplicateFields(fields)
	if len(got) != 3 {
		t.Fatalf("expected 3 unique fields, got %d", len(got))
	}
	// First occurrence wins
	if got[0].GoType != "string" {
		t.Errorf("first Id should be string, got %q", got[0].GoType)
	}
}

func TestDeduplicateFields_empty(t *testing.T) {
	if got := deduplicateFields(nil); got != nil {
		t.Errorf("nil input should return nil, got %v", got)
	}
	if got := deduplicateFields([]Field{}); len(got) != 0 {
		t.Errorf("empty input should return empty")
	}
}

// ---- deduplicateFields: attr vs element priority ----

// Bug 1: inherited attribute + own element → element should win.
func TestDeduplicateFields_inheritedAttrOwnElement_elementWins(t *testing.T) {
	fields := []Field{
		{GoName: "AxisLabels", GoType: "string", IsAttr: true},  // inherited attr (comes first)
		{GoName: "AxisLabels", GoType: "string", IsAttr: false}, // own element (comes after)
	}
	got := deduplicateFields(fields)
	if len(got) != 1 {
		t.Fatalf("expected 1 field, got %d: %+v", len(got), got)
	}
	if got[0].IsAttr {
		t.Error("own element should win over inherited attribute, but got attr")
	}
}

// Regression: inherited element + own attribute → attribute should still win.
func TestDeduplicateFields_inheritedElementOwnAttr_attrWins(t *testing.T) {
	fields := []Field{
		{GoName: "SrsName", GoType: "string", IsAttr: false}, // inherited element (comes first)
		{GoName: "SrsName", GoType: "string", IsAttr: true},  // own attribute (comes after)
	}
	got := deduplicateFields(fields)
	if len(got) != 1 {
		t.Fatalf("expected 1 field, got %d: %+v", len(got), got)
	}
	if !got[0].IsAttr {
		t.Error("own attribute should win over inherited element, but got element")
	}
}

// ---- isHTTP ----

func TestIsHTTP(t *testing.T) {
	if !isHTTP("http://www.opengis.net/gml/3.2") {
		t.Error("http:// should be HTTP")
	}
	if !isHTTP("https://example.com/schema.xsd") {
		t.Error("https:// should be HTTP")
	}
	if isHTTP("geometryBasic0d1d.xsd") {
		t.Error("relative path should not be HTTP")
	}
	if isHTTP("../schemas/gml.xsd") {
		t.Error("relative path with .. should not be HTTP")
	}
	if isHTTP("") {
		t.Error("empty string should not be HTTP")
	}
}

// ---- isBuiltinGoType ----

func TestIsBuiltinGoType(t *testing.T) {
	for _, bt := range []string{"string", "bool", "int", "float64", "float32"} {
		if !isBuiltinGoType(bt) {
			t.Errorf("isBuiltinGoType(%q) = false, want true", bt)
		}
	}
	for _, ct := range []string{"PointType", "[]string", "*Foo", "", "int64", "uint"} {
		if isBuiltinGoType(ct) {
			t.Errorf("isBuiltinGoType(%q) = true, want false", ct)
		}
	}
}

// ---- isStructType ----

func TestIsStructType(t *testing.T) {
	if isStructType("") {
		t.Error("empty string should not be struct type")
	}
	if isStructType("string") {
		t.Error("string is not a struct type")
	}
	if !isStructType("PointType") {
		t.Error("PointType should be a struct type")
	}
	if !isStructType("*DirectPositionType") {
		t.Error("pointer to struct should be a struct type")
	}
}
