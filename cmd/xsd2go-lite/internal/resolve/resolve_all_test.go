package resolve

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xsd2go-lite/internal/parse"
)

func TestResolver_ResolveAll_simpleContent(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <complexType name="MeasureType">
    <simpleContent>
      <extension base="double">
        <attribute name="uom" type="string" use="required"/>
        <attribute name="srs" type="anyURI"/>
      </extension>
    </simpleContent>
  </complexType>
</schema>`

	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(xsd), 0644); err != nil {
		t.Fatal(err)
	}

	r := NewResolver()
	if _, err := r.Load(path); err != nil {
		t.Fatal(err)
	}

	types := r.ResolveAll("http://example.com/test")
	if len(types) != 1 {
		t.Fatalf("expected 1 type, got %d", len(types))
	}

	mt := types[0]
	if mt.Name != "MeasureType" {
		t.Errorf("type name = %q", mt.Name)
	}

	fieldMap := make(map[string]parse.Field)
	for _, f := range mt.Fields {
		fieldMap[f.GoName] = f
	}

	vf, ok := fieldMap["Value"]
	if !ok {
		t.Error("Value (chardata) field missing")
	} else if !vf.IsChar {
		t.Error("Value field should be chardata")
	}

	uom, ok := fieldMap["Uom"]
	if !ok {
		t.Error("Uom attribute missing")
	} else if !uom.IsAttr {
		t.Error("Uom should be attr")
	}
}

func TestResolver_ResolveAll_attributeGroup(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <attributeGroup name="SRSGroup">
    <attribute name="srsName" type="anyURI"/>
    <attribute name="srsDimension" type="positiveInteger"/>
  </attributeGroup>
  <complexType name="GeomType">
    <sequence>
      <element name="pos" type="string"/>
    </sequence>
    <attributeGroup ref="tns:SRSGroup"/>
  </complexType>
</schema>`

	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(xsd), 0644); err != nil {
		t.Fatal(err)
	}

	r := NewResolver()
	if _, err := r.Load(path); err != nil {
		t.Fatal(err)
	}

	types := r.ResolveAll("http://example.com/test")
	if len(types) != 1 {
		t.Fatalf("expected 1 type, got %d", len(types))
	}

	gt := types[0]
	fieldMap := make(map[string]bool)
	for _, f := range gt.Fields {
		fieldMap[f.GoName] = true
	}

	if !fieldMap["Pos"] {
		t.Error("Pos element missing")
	}
	if !fieldMap["SrsName"] {
		t.Error("SrsName (from attributeGroup) missing")
	}
	if !fieldMap["SrsDimension"] {
		t.Error("SrsDimension (from attributeGroup) missing")
	}
}

func TestResolver_ResolveAll_elementRef(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <element name="pointMember" type="tns:PointType"/>
  <complexType name="PointType">
    <sequence>
      <element name="pos" type="string"/>
    </sequence>
  </complexType>
  <complexType name="MultiPointType">
    <sequence>
      <element ref="tns:pointMember" minOccurs="0" maxOccurs="unbounded"/>
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
		t.Fatal(err)
	}

	types := r.ResolveAll("http://example.com/test")
	byName := make(map[string]*parse.ComplexType)
	for _, ct := range types {
		byName[ct.Name] = ct
	}

	mp, ok := byName["MultiPointType"]
	if !ok {
		t.Fatal("MultiPointType not found")
	}

	if len(mp.Fields) != 1 {
		t.Fatalf("expected 1 field (pointMember ref), got %d: %+v", len(mp.Fields), mp.Fields)
	}
	f := mp.Fields[0]
	if f.GoName != "PointMember" {
		t.Errorf("GoName = %q, want PointMember", f.GoName)
	}
	if !f.Slice {
		t.Error("pointMember with maxOccurs=unbounded should be slice")
	}
	if !strings.HasPrefix(f.GoType, "[]") {
		t.Errorf("slice field GoType should start with [], got %q", f.GoType)
	}
}
