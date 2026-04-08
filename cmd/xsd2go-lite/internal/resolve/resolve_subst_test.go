package resolve

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolver_ResolveAll_substitutionGroup(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <element name="AbstractRing" type="tns:AbstractRingType" abstract="true"/>
  <element name="LinearRing" type="tns:LinearRingType" substitutionGroup="tns:AbstractRing"/>
  <element name="Ring" type="tns:RingType" substitutionGroup="tns:AbstractRing"/>
  <complexType name="AbstractRingType" abstract="true">
    <sequence/>
  </complexType>
  <complexType name="LinearRingType">
    <sequence>
      <element name="posList" type="string"/>
    </sequence>
  </complexType>
  <complexType name="RingType">
    <sequence>
      <element name="curveMember" type="string"/>
    </sequence>
  </complexType>
  <complexType name="AbstractRingPropertyType">
    <sequence>
      <element ref="tns:AbstractRing"/>
    </sequence>
  </complexType>
</schema>`

	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(xsd), 0644); err != nil {
		t.Fatal(err)
	}

	r := NewResolver()
	if _, err := r.Load(path); err != nil {
		t.Fatalf("Load: %v", err)
	}

	types := r.ResolveAll("http://example.com/test")
	byName := make(map[string]int)
	for i, ct := range types {
		byName[ct.Name] = i
	}

	idx, ok := byName["AbstractRingPropertyType"]
	if !ok {
		t.Fatal("AbstractRingPropertyType not found")
	}
	prop := types[idx]

	fieldMap := make(map[string]bool)
	for _, f := range prop.Fields {
		fieldMap[f.GoName] = true
	}

	if !fieldMap["AbstractRing"] {
		t.Error("AbstractRing field missing")
	}
	if !fieldMap["LinearRing"] {
		t.Error("LinearRing substitution member field missing")
	}
	if !fieldMap["Ring"] {
		t.Error("Ring substitution member field missing")
	}

	for _, f := range prop.Fields {
		switch f.GoName {
		case "LinearRing":
			if f.GoType != "*LinearRingType" {
				t.Errorf("LinearRing GoType = %q, want *LinearRingType", f.GoType)
			}
			if !f.Omit {
				t.Error("LinearRing field should have omitempty (Omit=true)")
			}
			if f.Slice {
				t.Error("LinearRing field should not be a slice")
			}
		case "Ring":
			if f.GoType != "*RingType" {
				t.Errorf("Ring GoType = %q, want *RingType", f.GoType)
			}
		}
	}

	if len(prop.Fields) != 3 {
		t.Errorf("expected 3 fields, got %d: %v", len(prop.Fields), prop.Fields)
	}
}

func TestResolver_ResolveAll_substitutionGroup_noMembers(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <element name="AbstractGeom" type="tns:AbstractGeomType" abstract="true"/>
  <complexType name="AbstractGeomType" abstract="true">
    <sequence/>
  </complexType>
  <complexType name="ContainerType">
    <sequence>
      <element ref="tns:AbstractGeom" minOccurs="0"/>
    </sequence>
  </complexType>
</schema>`

	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(xsd), 0644); err != nil {
		t.Fatal(err)
	}

	r := NewResolver()
	if _, err := r.Load(path); err != nil {
		t.Fatalf("Load: %v", err)
	}

	types := r.ResolveAll("http://example.com/test")
	byName := make(map[string]int)
	for i, ct := range types {
		byName[ct.Name] = i
	}

	idx, ok := byName["ContainerType"]
	if !ok {
		t.Fatal("ContainerType not found")
	}
	container := types[idx]
	if len(container.Fields) != 1 {
		t.Errorf("expected 1 field (no subst members), got %d", len(container.Fields))
	}
	if container.Fields[0].GoName != "AbstractGeom" {
		t.Errorf("field GoName = %q, want AbstractGeom", container.Fields[0].GoName)
	}
}
