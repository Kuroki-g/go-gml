package parse

import "testing"

const testNS = "http://www.opengis.net/gml/3.2"

func TestConvertComplexType_directGroupInExtension(t *testing.T) {
	ct := xsdComplexType{
		Name: "DynamicFeatureType",
		ComplexContent: &xsdComplexContent{
			Extension: &xsdExtension{
				Base:   "gml:AbstractFeatureType",
				Groups: []xsdGroupRef{{Ref: "gml:dynamicProperties", MinOccurs: "0"}},
			},
		},
	}
	result := convertComplexType(ct, testNS)
	if result.Content == nil {
		t.Fatal("Content must not be nil")
	}
	if len(result.Content.Groups) != 1 {
		t.Fatalf("Groups len: want 1, got %d", len(result.Content.Groups))
	}
	if result.Content.Groups[0].Ref != "gml:dynamicProperties" {
		t.Errorf("Ref: want gml:dynamicProperties, got %s", result.Content.Groups[0].Ref)
	}
	if result.Derivation == nil || result.Derivation.Kind != "extension" {
		t.Error("Derivation must be extension")
	}
}

func TestConvertComplexType_directGroupInRestriction(t *testing.T) {
	ct := xsdComplexType{
		Name: "RestrictedType",
		ComplexContent: &xsdComplexContent{
			Restriction: &xsdRestriction{
				Base:   "gml:AbstractType",
				Groups: []xsdGroupRef{{Ref: "gml:someGroup"}},
			},
		},
	}
	result := convertComplexType(ct, testNS)
	if result.Content == nil {
		t.Fatal("Content must not be nil")
	}
	if len(result.Content.Groups) != 1 {
		t.Fatalf("Groups len: want 1, got %d", len(result.Content.Groups))
	}
	if result.Derivation == nil || result.Derivation.Kind != "restriction" {
		t.Error("Derivation must be restriction")
	}
}

func TestConvertComplexType_directGroupTopLevel(t *testing.T) {
	ct := xsdComplexType{
		Name:   "DirectGroupType",
		Groups: []xsdGroupRef{{Ref: "gml:someGroup", MaxOccurs: "unbounded"}},
	}
	result := convertComplexType(ct, testNS)
	if result.Content == nil {
		t.Fatal("Content must not be nil")
	}
	if len(result.Content.Groups) != 1 {
		t.Fatalf("Groups len: want 1, got %d", len(result.Content.Groups))
	}
	if result.Content.Groups[0].MaxOccurs != "unbounded" {
		t.Errorf("MaxOccurs: want unbounded, got %s", result.Content.Groups[0].MaxOccurs)
	}
}

// Compositor (sequence) takes priority when present — direct Groups are ignored per XSD spec
// since (group | all | choice | sequence)? are mutually exclusive.
func TestConvertComplexType_compositorWhenBothPresent(t *testing.T) {
	ct := xsdComplexType{
		Name: "MixedType",
		ComplexContent: &xsdComplexContent{
			Extension: &xsdExtension{
				Base:     "gml:AbstractType",
				Sequence: &xsdSequence{Elements: []xsdElement{{Name: "foo", Type: "xs:string"}}},
				Groups:   []xsdGroupRef{{Ref: "gml:shouldNotAppear"}},
			},
		},
	}
	result := convertComplexType(ct, testNS)
	if result.Content == nil {
		t.Fatal("Content must not be nil")
	}
	if len(result.Content.Elems) != 1 {
		t.Errorf("Elems len: want 1 (from sequence), got %d", len(result.Content.Elems))
	}
	if len(result.Content.Groups) != 0 {
		t.Errorf("Groups len: want 0 when sequence present, got %d", len(result.Content.Groups))
	}
}
