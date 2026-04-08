package xsdlite

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ---- full pipeline: Load → ResolveAll → Generate ----

func TestEndToEnd_basicPipeline(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com/geo"
        xmlns="http://www.w3.org/2001/XMLSchema"
        xmlns:geo="http://example.com/geo">
  <simpleType name="DoubleList">
    <list itemType="double"/>
  </simpleType>
  <complexType name="DirectPositionType">
    <simpleContent>
      <extension base="geo:DoubleList">
        <attribute name="srsName" type="anyURI"/>
        <attribute name="srsDimension" type="positiveInteger"/>
      </extension>
    </simpleContent>
  </complexType>
  <complexType name="PointType">
    <sequence>
      <element name="pos" type="geo:DirectPositionType" minOccurs="0"/>
    </sequence>
    <attribute name="id" type="string"/>
    <attribute name="srsName" type="anyURI"/>
  </complexType>
  <element name="Point" type="geo:PointType"/>
</schema>`

	dir := t.TempDir()
	path := filepath.Join(dir, "geo.xsd")
	if err := os.WriteFile(path, []byte(xsd), 0644); err != nil {
		t.Fatal(err)
	}

	r := NewResolver()
	if _, err := r.Load(path); err != nil {
		t.Fatal(err)
	}

	types := r.ResolveAll("http://example.com/geo")
	src, err := Generate(types, "geo", false, false, nil, "", "")
	if err != nil {
		t.Fatalf("Generate: %v\n%s", err, src)
	}

	if !strings.HasPrefix(src, "package geo") {
		t.Errorf("missing package declaration")
	}
	if !strings.Contains(src, "type DirectPositionType struct") {
		t.Error("DirectPositionType missing")
	}
	if !strings.Contains(src, "type PointType struct") {
		t.Error("PointType missing")
	}
	if !strings.Contains(src, `xml:",chardata"`) {
		t.Error("chardata field missing")
	}
	if !strings.Contains(src, "srsName") {
		t.Error("srsName attribute missing")
	}
	if !strings.Contains(src, "*string") {
		t.Errorf("optional string attribute should be *string, got:\n%s", src)
	}
}

// ---- Load: nonexistent file ----

func TestResolver_Load_notFound(t *testing.T) {
	r := NewResolver()
	_, err := r.Load("/nonexistent/schema.xsd")
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}
