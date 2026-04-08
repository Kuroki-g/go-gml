package gen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xsd2go-lite/internal/resolve"
)

// ---- Correctness regression tests for bugs found ----

// Bug1: resolveTypeName returned "" for simpleType with restriction of non-builtin base.
func TestBug1_simpleTypeRestrictionNonBuiltinBase(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <simpleType name="BaseCode">
    <restriction base="string"/>
  </simpleType>
  <simpleType name="DerivedCode">
    <restriction base="tns:BaseCode"/>
  </simpleType>
  <element name="code" type="tns:DerivedCode"/>
  <complexType name="ItemType">
    <sequence>
      <element ref="tns:code" minOccurs="0"/>
    </sequence>
  </complexType>
</schema>`

	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(xsd), 0644); err != nil {
		t.Fatal(err)
	}

	r := resolve.NewResolver()
	if _, err := r.Load(path); err != nil {
		t.Fatal(err)
	}
	types := r.ResolveAll("http://example.com/test")

	src, err := Generate(types, "test", false, false, nil, "", "")
	if err != nil {
		t.Fatalf("Generate failed (likely invalid Go type): %v", err)
	}
	if strings.Contains(src, " * `") || strings.Contains(src, " [] `") {
		t.Errorf("generated code contains invalid bare pointer/slice type:\n%s", src)
	}
}

// Bug2: element ref path did not guard against empty GoType before pointer/slice wrapping.
func TestBug2_elementRefWithEmptyGoType(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <simpleType name="AliasCode">
    <restriction base="tns:NonExistent"/>
  </simpleType>
  <element name="alias" type="tns:AliasCode"/>
  <complexType name="ContainerType">
    <sequence>
      <element ref="tns:alias" maxOccurs="unbounded"/>
    </sequence>
  </complexType>
</schema>`

	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(xsd), 0644); err != nil {
		t.Fatal(err)
	}

	r := resolve.NewResolver()
	if _, err := r.Load(path); err != nil {
		t.Fatal(err)
	}
	types := r.ResolveAll("http://example.com/test")

	src, err := Generate(types, "test", false, false, nil, "", "")
	if err != nil {
		t.Fatalf("Generate produced invalid Go (bug2): %v", err)
	}
	if strings.Contains(src, " [] `") {
		t.Errorf("generated slice field with empty element type:\n%s", src)
	}
}

// Bug3: __base__:simpleType path passed baseST.GoType="" directly into Field.GoType.
func TestBug3_baseSimpleTypeWithEmptyGoType(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <simpleType name="CustomCode">
    <restriction base="tns:Unknown"/>
  </simpleType>
  <complexType name="CodedValueType">
    <simpleContent>
      <extension base="tns:CustomCode">
        <attribute name="codeSpace" type="anyURI"/>
      </extension>
    </simpleContent>
  </complexType>
</schema>`

	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(xsd), 0644); err != nil {
		t.Fatal(err)
	}

	r := resolve.NewResolver()
	if _, err := r.Load(path); err != nil {
		t.Fatal(err)
	}
	types := r.ResolveAll("http://example.com/test")

	src, err := Generate(types, "test", false, false, nil, "", "")
	if err != nil {
		t.Fatalf("Generate produced invalid Go (bug3): %v", err)
	}
	if strings.Contains(src, "Value  `") {
		t.Errorf("generated chardata field with empty type:\n%s", src)
	}
}
