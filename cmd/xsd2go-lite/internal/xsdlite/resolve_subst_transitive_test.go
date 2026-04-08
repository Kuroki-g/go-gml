package xsdlite

import (
	"os"
	"path/filepath"
	"testing"
)

// Bug 2: sortedSubstMembers must expand transitively.
// Chain: AbstractGeom → AbstractImplicitGeom (1-level) → Grid (2-level) → RectifiedGrid (3-level).
func TestResolver_ResolveAll_substitutionGroup_transitive(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <element name="AbstractGeom" type="tns:AbstractGeomType" abstract="true"/>
  <element name="AbstractImplicitGeom" type="tns:AbstractGeomType" abstract="true" substitutionGroup="tns:AbstractGeom"/>
  <element name="Grid" type="tns:GridType" substitutionGroup="tns:AbstractImplicitGeom"/>
  <element name="RectifiedGrid" type="tns:RectifiedGridType" substitutionGroup="tns:Grid"/>
  <complexType name="AbstractGeomType" abstract="true">
    <sequence/>
  </complexType>
  <complexType name="GridType">
    <sequence>
      <element name="limits" type="string"/>
    </sequence>
  </complexType>
  <complexType name="RectifiedGridType">
    <complexContent>
      <extension base="tns:GridType">
        <sequence>
          <element name="origin" type="string"/>
        </sequence>
      </extension>
    </complexContent>
  </complexType>
  <complexType name="DomainSetType">
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
	byName := make(map[string]*ComplexType)
	for _, ct := range types {
		byName[ct.Name] = ct
	}

	ds, ok := byName["DomainSetType"]
	if !ok {
		t.Fatal("DomainSetType not found")
	}

	fieldMap := make(map[string]Field)
	for _, f := range ds.Fields {
		fieldMap[f.GoName] = f
	}

	if _, ok := fieldMap["AbstractGeom"]; !ok {
		t.Error("AbstractGeom field missing")
	}
	if _, ok := fieldMap["AbstractImplicitGeom"]; !ok {
		t.Error("AbstractImplicitGeom (1-level transitive) field missing")
	}
	if gf, ok := fieldMap["Grid"]; !ok {
		t.Error("Grid (2-level transitive) field missing")
	} else if gf.GoType != "*GridType" {
		t.Errorf("Grid GoType = %q, want *GridType", gf.GoType)
	}
	if rf, ok := fieldMap["RectifiedGrid"]; !ok {
		t.Error("RectifiedGrid (3-level transitive) field missing")
	} else if rf.GoType != "*RectifiedGridType" {
		t.Errorf("RectifiedGrid GoType = %q, want *RectifiedGridType", rf.GoType)
	}
}
