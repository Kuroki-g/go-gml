package xsdlite

import (
	"os"
	"path/filepath"
	"testing"
)

// ---- resolveRawField: group ref path ----

func TestResolver_groupRef(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <group name="NameProps">
    <sequence>
      <element name="name" type="string" minOccurs="0" maxOccurs="unbounded"/>
      <element name="description" type="string" minOccurs="0"/>
    </sequence>
  </group>
  <complexType name="FeatureType">
    <sequence>
      <group ref="tns:NameProps"/>
      <element name="value" type="double"/>
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

	if len(types) != 1 {
		t.Fatalf("expected 1 type, got %d", len(types))
	}
	ft := types[0]

	fieldMap := make(map[string]Field)
	for _, f := range ft.Fields {
		fieldMap[f.GoName] = f
	}

	if _, ok := fieldMap["Name"]; !ok {
		t.Error("Name (from group ref) missing")
	}
	if _, ok := fieldMap["Description"]; !ok {
		t.Error("Description (from group ref) missing")
	}
	if _, ok := fieldMap["Value"]; !ok {
		t.Error("Value missing")
	}
}

// ---- resolveRawField: base = simpleType ----

func TestResolver_baseSimpleType(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <simpleType name="CoordList">
    <list itemType="double"/>
  </simpleType>
  <complexType name="PosListType">
    <simpleContent>
      <extension base="tns:CoordList">
        <attribute name="count" type="positiveInteger"/>
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
	pt := types[0]

	var hasChardata bool
	for _, f := range pt.Fields {
		if f.IsChar {
			hasChardata = true
			// CoordList is a list → string
			if f.GoType != "string" {
				t.Errorf("chardata GoType = %q, want string", f.GoType)
			}
		}
	}
	if !hasChardata {
		t.Error("PosListType should have chardata field for list base type")
	}
}

// ---- resolveRawField: base = XSD builtin (for complexContent) ----

func TestResolver_baseXSBuiltin(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema">
  <complexType name="ExtendedString">
    <simpleContent>
      <extension base="string">
        <attribute name="lang" type="string"/>
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

	var hasChardata bool
	for _, f := range types[0].Fields {
		if f.IsChar {
			hasChardata = true
			if f.GoType != "string" {
				t.Errorf("chardata for string base should be string, got %q", f.GoType)
			}
		}
	}
	if !hasChardata {
		t.Error("ExtendedString should have chardata field")
	}
}

// ---- resolveTypeName: unknown type fallback ----

func TestResolver_resolveTypeName_unknown(t *testing.T) {
	r := NewResolver()
	got := r.resolveTypeName("ns:UnknownType", "http://example.com")
	if got != "string" {
		t.Errorf("unknown type should fall back to string, got %q", got)
	}
}

func TestResolver_resolveTypeName_inline(t *testing.T) {
	r := NewResolver()
	got := r.resolveTypeName("__inline__", "http://example.com")
	if got != "string" {
		t.Errorf("__inline__ should map to string, got %q", got)
	}
}
