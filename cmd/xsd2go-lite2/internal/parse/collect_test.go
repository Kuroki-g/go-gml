package parse

import "testing"

func TestDirectGroupsContentModel_empty(t *testing.T) {
	if directGroupsContentModel(nil) != nil {
		t.Error("expected nil for empty input")
	}
	if directGroupsContentModel([]xsdGroupRef{}) != nil {
		t.Error("expected nil for empty slice")
	}
}

func TestDirectGroupsContentModel_single(t *testing.T) {
	cm := directGroupsContentModel([]xsdGroupRef{
		{Ref: "gml:dynamicProperties", MinOccurs: "0"},
	})
	if cm == nil {
		t.Fatal("expected non-nil ContentModel")
	}
	if cm.Kind != "sequence" {
		t.Errorf("Kind: want sequence, got %s", cm.Kind)
	}
	if len(cm.Groups) != 1 {
		t.Fatalf("Groups len: want 1, got %d", len(cm.Groups))
	}
	g := cm.Groups[0]
	if g.Ref != "gml:dynamicProperties" {
		t.Errorf("Ref: want gml:dynamicProperties, got %s", g.Ref)
	}
	if g.MinOccurs != "0" {
		t.Errorf("MinOccurs: want 0, got %s", g.MinOccurs)
	}
}

func TestDirectGroupsContentModel_unbounded(t *testing.T) {
	cm := directGroupsContentModel([]xsdGroupRef{
		{Ref: "gml:someGroup", MaxOccurs: "unbounded"},
	})
	if cm == nil {
		t.Fatal("expected non-nil ContentModel")
	}
	if cm.Groups[0].MaxOccurs != "unbounded" {
		t.Errorf("MaxOccurs: want unbounded, got %s", cm.Groups[0].MaxOccurs)
	}
}
