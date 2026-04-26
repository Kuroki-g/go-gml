package resolve

import (
	"testing"

	"xsd2go-lite2/internal/parse"
)

// TestAttrGroupFieldsNotInComplexTypeFields verifies that after Step 1, fields
// from an attributeGroup ref are NOT inlined into ct.Fields. They are instead
// accessible via ct.Embeds, and the generated embed struct carries the fields.
func TestAttrGroupFieldsNotInComplexTypeFields(t *testing.T) {
	r := NewResolver()
	r.nsMaps[gmlNS] = map[string]string{"gml": gmlNS}

	ag := &parse.AttrGroup{
		Name:  "SRSReferenceGroup",
		Attrs: []parse.AttrDecl{{Name: "srsName", TypeRef: "xs:anyURI", Use: "optional"}},
	}
	r.allAttrGroups[gmlNS+" SRSReferenceGroup"] = ag

	ct := &parse.ComplexType{
		Name:       "SurfaceType",
		Source:     gmlNS,
		AttrGroups: []string{"gml:SRSReferenceGroup"},
	}
	r.allComplexTypes[gmlNS+" SurfaceType"] = ct

	r.resolveComplexType(ct, nil)

	for _, f := range ct.Fields {
		if f.GoName == "SrsName" {
			t.Errorf("SrsName from attributeGroup must not appear in ct.Fields; got %v", ct.Fields)
		}
	}
	if len(ct.Embeds) == 0 || ct.Embeds[0].XSDName != "SRSReferenceGroup" {
		t.Errorf("expected SRSReferenceGroup in ct.Embeds; got %v", ct.Embeds)
	}
}

// TestAttrGroupFieldsPopulated verifies that after resolution, the AttrGroup.Fields
// slice is populated with the resolved fields (used by gen phase to generate the struct).
func TestAttrGroupFieldsPopulated(t *testing.T) {
	r := NewResolver()
	r.nsMaps[gmlNS] = map[string]string{"gml": gmlNS}

	ag := &parse.AttrGroup{
		Name:  "SRSReferenceGroup",
		Attrs: []parse.AttrDecl{{Name: "srsName", TypeRef: "xs:anyURI", Use: "optional"}},
	}
	r.allAttrGroups[gmlNS+" SRSReferenceGroup"] = ag

	ct := &parse.ComplexType{
		Name:       "SurfaceType",
		Source:     gmlNS,
		AttrGroups: []string{"gml:SRSReferenceGroup"},
	}
	r.allComplexTypes[gmlNS+" SurfaceType"] = ct

	r.resolveComplexType(ct, nil)

	if len(ag.Fields) == 0 {
		t.Errorf("expected AttrGroup.Fields to be populated after resolve; got empty")
	}
	found := false
	for _, f := range ag.Fields {
		if f.GoName == "SrsName" {
			found = true
		}
	}
	if !found {
		t.Errorf("SrsName missing from AttrGroup.Fields; got %v", ag.Fields)
	}
}
