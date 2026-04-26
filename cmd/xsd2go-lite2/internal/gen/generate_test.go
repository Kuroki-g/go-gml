package gen_test

import (
	"strings"
	"testing"

	"xsd2go-lite2/internal/gen"
	"xsd2go-lite2/internal/parse"
)

const testNS = "http://www.opengis.net/gml/3.2"

func agKey(ns, name string) string { return ns + " " + name }

// TestGenerate_attrGroupStructGenerated verifies that a referenced attributeGroup
// produces a named struct in the output.
func TestGenerate_attrGroupStructGenerated(t *testing.T) {
	ag := &parse.AttrGroup{
		Name:   "SRSReferenceGroup",
		Source: testNS,
		Fields: []parse.Field{
			{GoName: "SrsName", XMLTag: "srsName,attr", GoType: "*string", IsAttr: true, Omit: true},
		},
	}
	ct := &parse.ComplexType{
		Name:   "AbstractGeometryType",
		Source: testNS,
		Fields: []parse.Field{},
		Embeds: []parse.EmbedRef{
			{XSDName: "SRSReferenceGroup", NS: testNS, Kind: "attributeGroup"},
		},
	}
	attrGroups := map[string]*parse.AttrGroup{agKey(testNS, "SRSReferenceGroup"): ag}

	got, err := gen.Generate([]*parse.ComplexType{ct}, "v3_2_1", false, false, nil, "", "", attrGroups)
	if err != nil {
		t.Fatalf("Generate: %v\n%s", err, got)
	}
	if !strings.Contains(got, "type SRSReferenceGroup struct") {
		t.Errorf("expected SRSReferenceGroup struct; got:\n%s", got)
	}
	if !strings.Contains(got, `xml:"srsName,attr,omitempty"`) {
		t.Errorf("expected srsName attr field in SRSReferenceGroup; got:\n%s", got)
	}
}

// TestGenerate_complexTypeEmbedsAttrGroup verifies that a ComplexType with an
// attributeGroup in Embeds has an embedded field (no xml tag) in the generated struct.
func TestGenerate_complexTypeEmbedsAttrGroup(t *testing.T) {
	ag := &parse.AttrGroup{
		Name:   "SRSReferenceGroup",
		Source: testNS,
		Fields: []parse.Field{
			{GoName: "SrsName", XMLTag: "srsName,attr", GoType: "*string", IsAttr: true, Omit: true},
		},
	}
	ct := &parse.ComplexType{
		Name:   "AbstractGeometryType",
		Source: testNS,
		Fields: []parse.Field{
			{GoName: "Gid", XMLTag: "gid", GoType: "*string", IsAttr: true, Omit: true},
		},
		Embeds: []parse.EmbedRef{
			{XSDName: "SRSReferenceGroup", NS: testNS, Kind: "attributeGroup"},
		},
	}
	attrGroups := map[string]*parse.AttrGroup{agKey(testNS, "SRSReferenceGroup"): ag}

	got, err := gen.Generate([]*parse.ComplexType{ct}, "v3_2_1", false, false, nil, "", "", attrGroups)
	if err != nil {
		t.Fatalf("Generate: %v\n%s", err, got)
	}

	// AbstractGeometryType must embed SRSReferenceGroup without an xml tag.
	// The line should be just "SRSReferenceGroup" with no backtick.
	lines := strings.Split(got, "\n")
	found := false
	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed == "SRSReferenceGroup" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected bare 'SRSReferenceGroup' embed line in AbstractGeometryType; got:\n%s", got)
	}
}

// TestGenerate_attrGroupGeneratedOnce verifies that when two ComplexTypes reference
// the same attributeGroup, its struct is generated exactly once.
func TestGenerate_attrGroupGeneratedOnce(t *testing.T) {
	ag := &parse.AttrGroup{
		Name:   "SRSReferenceGroup",
		Source: testNS,
		Fields: []parse.Field{
			{GoName: "SrsName", XMLTag: "srsName,attr", GoType: "*string", IsAttr: true, Omit: true},
		},
	}
	embed := parse.EmbedRef{XSDName: "SRSReferenceGroup", NS: testNS, Kind: "attributeGroup"}
	ct1 := &parse.ComplexType{Name: "TypeA", Source: testNS, Embeds: []parse.EmbedRef{embed}}
	ct2 := &parse.ComplexType{Name: "TypeB", Source: testNS, Embeds: []parse.EmbedRef{embed}}
	attrGroups := map[string]*parse.AttrGroup{agKey(testNS, "SRSReferenceGroup"): ag}

	got, err := gen.Generate([]*parse.ComplexType{ct1, ct2}, "v3_2_1", false, false, nil, "", "", attrGroups)
	if err != nil {
		t.Fatalf("Generate: %v\n%s", err, got)
	}
	count := strings.Count(got, "type SRSReferenceGroup struct")
	if count != 1 {
		t.Errorf("expected SRSReferenceGroup struct exactly once; got %d occurrences:\n%s", count, got)
	}
}

// TestGenerate_noAttrGroups verifies backward compatibility: passing a nil attrGroups
// map produces valid output with no embed structs.
func TestGenerate_noAttrGroups(t *testing.T) {
	ct := &parse.ComplexType{
		Name:   "PointType",
		Source: testNS,
		Fields: []parse.Field{
			{GoName: "Pos", XMLTag: "http://www.opengis.net/gml/3.2 pos", GoType: "string"},
		},
	}

	got, err := gen.Generate([]*parse.ComplexType{ct}, "v3_2_1", false, false, nil, "", "", nil)
	if err != nil {
		t.Fatalf("Generate: %v\n%s", err, got)
	}
	if !strings.Contains(got, "type PointType struct") {
		t.Errorf("expected PointType struct; got:\n%s", got)
	}
}
