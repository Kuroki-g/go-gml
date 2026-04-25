package resolve

import (
	"testing"

	"xsd2go-lite2/internal/parse"
)

const gmlNS = "http://www.opengis.net/gml/3.2"

// TestResolveDirectGroupRefInExtension verifies that a <group ref> placed directly
// under <complexContent><extension> (without an intervening sequence/choice/all)
// produces the expected Fields after resolution. This covers the DynamicFeatureType
// pattern found in gml/3.1.1 and gml/3.2.1.
func TestResolveDirectGroupRefInExtension(t *testing.T) {
	r := NewResolver()
	r.nsMaps[gmlNS] = map[string]string{"gml": gmlNS}

	// base type: AbstractFeatureType (no fields for simplicity)
	base := &parse.ComplexType{
		Name:   "AbstractFeatureType",
		Source: gmlNS,
		Fields: []parse.Field{},
	}
	r.allComplexTypes[gmlNS+" AbstractFeatureType"] = base

	// group: dynamicProperties → sequence with 2 elements
	group := &parse.Group{
		Name: "dynamicProperties",
		Content: &parse.ContentModel{
			Kind: "sequence",
			Elems: []parse.ElemDecl{
				{Name: "validTime", TypeRef: "gml:TimePrimitivePropertyType", MinOccurs: "0"},
				{Name: "dataSource", TypeRef: "gml:StringOrRefType", MinOccurs: "0"},
			},
		},
	}
	r.allGroups[gmlNS+" dynamicProperties"] = group

	// register the element types so resolveElemDecl can build Field
	r.allComplexTypes[gmlNS+" TimePrimitivePropertyType"] = &parse.ComplexType{
		Name: "TimePrimitivePropertyType", Source: gmlNS,
	}
	r.allComplexTypes[gmlNS+" StringOrRefType"] = &parse.ComplexType{
		Name: "StringOrRefType", Source: gmlNS,
	}

	// DynamicFeatureType: extension of AbstractFeatureType with direct group ref
	dynType := &parse.ComplexType{
		Name:        "DynamicFeatureType",
		Source:      gmlNS,
		ContentKind: parse.ContentKindComplex,
		Derivation:  &parse.Derivation{Kind: "extension", Base: "gml:AbstractFeatureType"},
		Content: &parse.ContentModel{
			Kind: "sequence",
			Groups: []parse.GroupRef{
				{Ref: "gml:dynamicProperties", MinOccurs: "0"},
			},
		},
	}
	r.allComplexTypes[gmlNS+" DynamicFeatureType"] = dynType

	r.resolveComplexType(dynType, nil)

	if len(dynType.Fields) == 0 {
		t.Fatal("expected Fields to be populated")
	}

	fieldNames := make(map[string]bool)
	for _, f := range dynType.Fields {
		fieldNames[f.GoName] = true
	}

	for _, want := range []string{"ValidTime", "DataSource"} {
		if !fieldNames[want] {
			t.Errorf("missing field %s; got %v", want, dynType.Fields)
		}
	}
}
