package resolve

import (
	"testing"

	"xsd2go-lite/internal/parse"
)

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
	fields := []parse.Field{
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
	if got := deduplicateFields([]parse.Field{}); len(got) != 0 {
		t.Errorf("empty input should return empty")
	}
}

// ---- deduplicateFields: attr vs element priority ----

// Bug 1: inherited attribute + own element → element should win.
func TestDeduplicateFields_inheritedAttrOwnElement_elementWins(t *testing.T) {
	fields := []parse.Field{
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
	fields := []parse.Field{
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
