package parse

import (
	"strings"
	"testing"
)

// ---- collectGroupFields ----

func TestCollectGroupFields_sequence(t *testing.T) {
	g := xsdGroup{
		Name: "GeomGroup",
		Sequence: &xsdSequence{
			Elements: []xsdElement{
				{Name: "x", Type: "string"},
				{Name: "y", Type: "double"},
			},
		},
	}
	fields := collectGroupFields(g)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
}

func TestCollectGroupFields_choice(t *testing.T) {
	g := xsdGroup{
		Name: "ChoiceGroup",
		Choice: &xsdSequence{
			Elements: []xsdElement{
				{Name: "a", Type: "string"},
			},
		},
	}
	fields := collectGroupFields(g)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
}

func TestCollectGroupFields_all(t *testing.T) {
	g := xsdGroup{
		Name: "AllGroup",
		All: &xsdSequence{
			Elements: []xsdElement{
				{Name: "m", Type: "integer"},
			},
		},
	}
	fields := collectGroupFields(g)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
}

func TestCollectGroupFields_empty(t *testing.T) {
	g := xsdGroup{Name: "Empty"}
	fields := collectGroupFields(g)
	if len(fields) != 0 {
		t.Errorf("empty group should have 0 fields, got %d", len(fields))
	}
}

// ---- convertComplexType branches ----

func TestConvertComplexType_directSequence(t *testing.T) {
	ct := xsdComplexType{
		Name: "DirectType",
		Sequence: &xsdSequence{
			Elements: []xsdElement{
				{Name: "value", Type: "string"},
			},
		},
		Attributes: []xsdAttribute{
			{Name: "id", Type: "string", Use: UseRequired},
		},
	}
	result := convertComplexType(ct, "http://example.com")

	if result.Name != "DirectType" {
		t.Errorf("Name = %q", result.Name)
	}
	// Should have 1 element + 1 attribute raw field
	if len(result.RawFields) != 2 {
		t.Errorf("expected 2 raw fields, got %d: %+v", len(result.RawFields), result.RawFields)
	}
}

func TestConvertComplexType_complexContentRestriction(t *testing.T) {
	ct := xsdComplexType{
		Name: "RestrictedType",
		ComplexContent: &xsdComplexContent{
			Restriction: &xsdRestriction{
				Base: "gml:BaseType",
				Sequence: &xsdSequence{
					Elements: []xsdElement{
						{Name: "pos", Type: "string"},
					},
				},
			},
		},
	}
	result := convertComplexType(ct, "http://example.com")

	// For restriction: own sequence fields come first, base placeholder last.
	if len(result.RawFields) == 0 {
		t.Fatal("no raw fields")
	}
	last := result.RawFields[len(result.RawFields)-1]
	if !strings.HasPrefix(last.Ref, "__base__:") {
		t.Errorf("last field should be __base__ placeholder, got %+v", last)
	}
}

func TestConvertComplexType_simpleContentRestriction(t *testing.T) {
	ct := xsdComplexType{
		Name: "SCRestricted",
		SimpleContent: &xsdSimpleContent{
			Restriction: &xsdRestriction{
				Base: "xs:string",
				Attributes: []xsdAttribute{
					{Name: "lang", Type: "string"},
				},
			},
		},
	}
	result := convertComplexType(ct, "http://example.com")

	if len(result.RawFields) == 0 {
		t.Fatal("no raw fields")
	}
	// First should be chardata
	if !result.RawFields[0].IsChar {
		t.Errorf("first field should be chardata: %+v", result.RawFields[0])
	}
}

func TestConvertComplexType_abstract(t *testing.T) {
	ct := xsdComplexType{
		Name:     "AbstractBase",
		Abstract: "true",
	}
	result := convertComplexType(ct, "http://example.com")
	if !result.Abstract {
		t.Error("Abstract should be true")
	}
}

func TestConvertComplexType_mixed(t *testing.T) {
	ct := xsdComplexType{
		Name:  "MixedType",
		Mixed: "true",
	}
	result := convertComplexType(ct, "http://example.com")
	if !result.Mixed {
		t.Error("Mixed should be true")
	}
}
