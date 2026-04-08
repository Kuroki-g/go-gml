package xsdlite

import (
	"os"
	"path/filepath"
	"testing"
)

// ---- parseSchema (integration with temp file) ----

func writeTemp(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

const simpleXSD = `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/test"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:tns="http://example.com/test"
        elementFormDefault="qualified">

  <simpleType name="CodeType">
    <restriction base="string"/>
  </simpleType>

  <simpleType name="MeasureList">
    <list itemType="double"/>
  </simpleType>

  <complexType name="PointType">
    <sequence>
      <element name="pos" type="tns:CodeType" minOccurs="0"/>
    </sequence>
    <attribute name="srsName" type="anyURI"/>
  </complexType>

  <complexType name="ValueType">
    <simpleContent>
      <extension base="double">
        <attribute name="uom" type="string" use="required"/>
      </extension>
    </simpleContent>
  </complexType>

  <element name="Point" type="tns:PointType"/>
</schema>`

func TestParseSchema_simpleXSD(t *testing.T) {
	path := writeTemp(t, simpleXSD)
	s, err := parseSchema(path)
	if err != nil {
		t.Fatalf("parseSchema: %v", err)
	}

	if s.TargetNamespace != "http://example.com/test" {
		t.Errorf("TargetNamespace = %q", s.TargetNamespace)
	}
	if len(s.ComplexTypes) != 2 {
		t.Errorf("ComplexTypes = %d, want 2", len(s.ComplexTypes))
	}
	if len(s.SimpleTypes) != 2 {
		t.Errorf("SimpleTypes = %d, want 2", len(s.SimpleTypes))
	}

	// Global element registered
	if s.Elements["Point"] != "tns:PointType" {
		t.Errorf("Elements[Point] = %q, want %q", s.Elements["Point"], "tns:PointType")
	}

	// NSMap captures xmlns:tns
	if s.NSMap["tns"] != "http://example.com/test" {
		t.Errorf("NSMap[tns] = %q, want %q", s.NSMap["tns"], "http://example.com/test")
	}
}

func TestParseSchema_simpleContent(t *testing.T) {
	path := writeTemp(t, simpleXSD)
	s, err := parseSchema(path)
	if err != nil {
		t.Fatal(err)
	}

	var valueType *ComplexType
	for i := range s.ComplexTypes {
		if s.ComplexTypes[i].Name == "ValueType" {
			valueType = &s.ComplexTypes[i]
		}
	}
	if valueType == nil {
		t.Fatal("ValueType not found")
	}

	// First raw field should be chardata
	if len(valueType.RawFields) == 0 {
		t.Fatal("no raw fields in ValueType")
	}
	if !valueType.RawFields[0].IsChar {
		t.Errorf("first field should be chardata, got %+v", valueType.RawFields[0])
	}
}

func TestParseSchema_notFound(t *testing.T) {
	_, err := parseSchema("/nonexistent/path/test.xsd")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestParseSchema_badXML(t *testing.T) {
	path := writeTemp(t, `<schema><NOT VALID`)
	_, err := parseSchema(path)
	if err == nil {
		t.Error("expected error for invalid XML")
	}
}
