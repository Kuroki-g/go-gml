package resolve

import (
	"testing"

	"xsd2go-lite2/internal/parse"
)

const (
	gmlNS = "http://www.opengis.net/gml/3.2"
	xsNS  = "http://www.w3.org/2001/XMLSchema"
)

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

// TestResolveRestrictionBaseAnyType verifies that own attrs declared inside
// <complexContent><restriction base="anyType"> are emitted as Fields.
// anyType is a built-in XSD type not registered in allComplexTypes,
// so the !ok branch must still process own attrs instead of returning nil.
func TestResolveRestrictionBaseAnyType(t *testing.T) {
	r := NewResolver()
	r.nsMaps[gmlNS] = map[string]string{"xs": xsNS}

	ct := &parse.ComplexType{
		Name:        "AbstractGeometryType",
		Abstract:    true,
		Source:      gmlNS,
		ContentKind: parse.ContentKindComplex,
		Derivation:  &parse.Derivation{Kind: "restriction", Base: "anyType"},
		Attrs: []parse.AttrDecl{
			{Name: "gid", TypeRef: "xs:ID", Use: "optional"},
			{Name: "srsName", TypeRef: "xs:anyURI", Use: "optional"},
		},
	}
	r.allComplexTypes[gmlNS+" AbstractGeometryType"] = ct

	r.resolveComplexType(ct, nil)

	if len(ct.Fields) == 0 {
		t.Fatal("expected Fields to be populated; got none")
	}
	fieldNames := make(map[string]bool)
	for _, f := range ct.Fields {
		fieldNames[f.GoName] = true
	}
	for _, want := range []string{"Gid", "SrsName"} {
		if !fieldNames[want] {
			t.Errorf("missing field %s; got %v", want, ct.Fields)
		}
	}
}

// TestResolveRestrictionBaseMissingExternal verifies that when the restriction
// base references an unresolvable external type, own attrs are still emitted
// as Fields (warning is expected but not asserted here).
func TestResolveRestrictionBaseMissingExternal(t *testing.T) {
	r := NewResolver()
	r.nsMaps[gmlNS] = map[string]string{"ext": "http://example.com/external"}

	ct := &parse.ComplexType{
		Name:        "LocalType",
		Source:      gmlNS,
		ContentKind: parse.ContentKindComplex,
		Derivation:  &parse.Derivation{Kind: "restriction", Base: "ext:UnknownBase"},
		Attrs: []parse.AttrDecl{
			{Name: "id", TypeRef: "xs:string", Use: "optional"},
		},
	}
	r.allComplexTypes[gmlNS+" LocalType"] = ct

	r.resolveComplexType(ct, nil)

	fieldNames := make(map[string]bool)
	for _, f := range ct.Fields {
		fieldNames[f.GoName] = true
	}
	if !fieldNames["Id"] {
		t.Errorf("own attr 'id' missing from Fields after unresolvable restriction base; got %v", ct.Fields)
	}
}
