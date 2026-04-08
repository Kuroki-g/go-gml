package parse

import "testing"

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
