package resolve

import (
	"os"
	"path/filepath"
	"testing"
)

// ---- Resolver Load tests ----

// buildTwoFileSchema creates two XSD files: base.xsd (defines BaseType) and
// derived.xsd (includes base.xsd, defines DerivedType extends BaseType).
func buildTwoFileSchema(t *testing.T) (string, string) {
	t.Helper()
	dir := t.TempDir()

	baseXSD := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <complexType name="BaseType">
    <sequence>
      <element name="name" type="string" minOccurs="0"/>
    </sequence>
    <attribute name="id" type="string"/>
  </complexType>
</schema>`

	derivedXSD := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test">
  <include schemaLocation="base.xsd"/>
  <complexType name="DerivedType">
    <complexContent>
      <extension base="tns:BaseType">
        <sequence>
          <element name="value" type="double"/>
        </sequence>
      </extension>
    </complexContent>
  </complexType>
</schema>`

	basePath := filepath.Join(dir, "base.xsd")
	derivedPath := filepath.Join(dir, "derived.xsd")
	if err := os.WriteFile(basePath, []byte(baseXSD), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(derivedPath, []byte(derivedXSD), 0644); err != nil {
		t.Fatal(err)
	}
	return basePath, derivedPath
}

func TestResolver_Load_includeAndInheritance(t *testing.T) {
	_, derivedPath := buildTwoFileSchema(t)

	r := NewResolver()
	if _, err := r.Load(derivedPath); err != nil {
		t.Fatalf("Load: %v", err)
	}

	types := r.ResolveAll("http://example.com/test")
	byName := make(map[string]bool)
	for _, ct := range types {
		byName[ct.Name] = true
		if ct.Name == "DerivedType" {
			// DerivedType should have inherited fields from BaseType (name, id) + own (value).
			fieldNames := make(map[string]bool)
			for _, f := range ct.Fields {
				fieldNames[f.GoName] = true
			}
			if !fieldNames["Name"] {
				t.Error("inherited field 'Name' missing from DerivedType")
			}
			if !fieldNames["Id"] {
				t.Error("inherited field 'Id' missing from DerivedType")
			}
			if !fieldNames["Value"] {
				t.Error("own field 'Value' missing from DerivedType")
			}
		}
	}
	if !byName["DerivedType"] {
		t.Fatal("DerivedType not found")
	}
}

func TestResolver_Load_cycleProtection(t *testing.T) {
	// Two schemas that include each other (cycle).
	dir := t.TempDir()

	aXSD := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/a"
        xmlns="http://www.w3.org/2001/XMLSchema">
  <include schemaLocation="b.xsd"/>
  <complexType name="AType">
    <sequence><element name="x" type="string"/></sequence>
  </complexType>
</schema>`

	bXSD := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/b"
        xmlns="http://www.w3.org/2001/XMLSchema">
  <include schemaLocation="a.xsd"/>
  <complexType name="BType">
    <sequence><element name="y" type="string"/></sequence>
  </complexType>
</schema>`

	if err := os.WriteFile(filepath.Join(dir, "a.xsd"), []byte(aXSD), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "b.xsd"), []byte(bXSD), 0644); err != nil {
		t.Fatal(err)
	}

	r := NewResolver()
	// Should not hang or panic.
	if _, err := r.Load(filepath.Join(dir, "a.xsd")); err != nil {
		t.Fatalf("Load with cycle: %v", err)
	}
}

func TestResolver_Load_alreadyLoadedReturnsCached(t *testing.T) {
	_, derivedPath := buildTwoFileSchema(t)

	r := NewResolver()
	s1, err := r.Load(derivedPath)
	if err != nil {
		t.Fatal(err)
	}
	s2, err := r.Load(derivedPath)
	if err != nil {
		t.Fatal(err)
	}
	// Should return the same pointer (cached).
	if s1 != s2 {
		t.Error("second Load should return cached schema pointer")
	}
}
