package xsdlite

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ---- collectGroupFields ----

func TestCollectGroupFields_sequence(t *testing.T) {
	g := xsdGroup{
		Name: "GeomGroup",
		Sequence: &xsdSequence{
			Elements: []xsdElement{
				{Name: "x", Type: "string"},
				{Name: "y", Type: "double"},
			},
		},
	}
	fields := collectGroupFields(g)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
}

func TestCollectGroupFields_choice(t *testing.T) {
	g := xsdGroup{
		Name: "ChoiceGroup",
		Choice: &xsdSequence{
			Elements: []xsdElement{
				{Name: "a", Type: "string"},
			},
		},
	}
	fields := collectGroupFields(g)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
}

func TestCollectGroupFields_all(t *testing.T) {
	g := xsdGroup{
		Name: "AllGroup",
		All: &xsdSequence{
			Elements: []xsdElement{
				{Name: "m", Type: "integer"},
			},
		},
	}
	fields := collectGroupFields(g)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
}

func TestCollectGroupFields_empty(t *testing.T) {
	g := xsdGroup{Name: "Empty"}
	fields := collectGroupFields(g)
	if len(fields) != 0 {
		t.Errorf("empty group should have 0 fields, got %d", len(fields))
	}
}

// ---- convertComplexType branches ----

func TestConvertComplexType_directSequence(t *testing.T) {
	ct := xsdComplexType{
		Name: "DirectType",
		Sequence: &xsdSequence{
			Elements: []xsdElement{
				{Name: "value", Type: "string"},
			},
		},
		Attributes: []xsdAttribute{
			{Name: "id", Type: "string", Use: "required"},
		},
	}
	result := convertComplexType(ct, "http://example.com")

	if result.Name != "DirectType" {
		t.Errorf("Name = %q", result.Name)
	}
	// Should have 1 element + 1 attribute raw field
	if len(result.RawFields) != 2 {
		t.Errorf("expected 2 raw fields, got %d: %+v", len(result.RawFields), result.RawFields)
	}
}

func TestConvertComplexType_complexContentRestriction(t *testing.T) {
	ct := xsdComplexType{
		Name: "RestrictedType",
		ComplexContent: &xsdComplexContent{
			Restriction: &xsdRestriction{
				Base: "gml:BaseType",
				Sequence: &xsdSequence{
					Elements: []xsdElement{
						{Name: "pos", Type: "string"},
					},
				},
			},
		},
	}
	result := convertComplexType(ct, "http://example.com")

	// For restriction: own sequence fields come first, base placeholder last.
	if len(result.RawFields) == 0 {
		t.Fatal("no raw fields")
	}
	last := result.RawFields[len(result.RawFields)-1]
	if !strings.HasPrefix(last.Ref, "__base__:") {
		t.Errorf("last field should be __base__ placeholder, got %+v", last)
	}
}

func TestConvertComplexType_simpleContentRestriction(t *testing.T) {
	ct := xsdComplexType{
		Name: "SCRestricted",
		SimpleContent: &xsdSimpleContent{
			Restriction: &xsdRestriction{
				Base: "xs:string",
				Attributes: []xsdAttribute{
					{Name: "lang", Type: "string"},
				},
			},
		},
	}
	result := convertComplexType(ct, "http://example.com")

	if len(result.RawFields) == 0 {
		t.Fatal("no raw fields")
	}
	// First should be chardata
	if !result.RawFields[0].IsChar {
		t.Errorf("first field should be chardata: %+v", result.RawFields[0])
	}
}

func TestConvertComplexType_abstract(t *testing.T) {
	ct := xsdComplexType{
		Name:     "AbstractBase",
		Abstract: "true",
	}
	result := convertComplexType(ct, "http://example.com")
	if !result.Abstract {
		t.Error("Abstract should be true")
	}
}

func TestConvertComplexType_mixed(t *testing.T) {
	ct := xsdComplexType{
		Name:  "MixedType",
		Mixed: "true",
	}
	result := convertComplexType(ct, "http://example.com")
	if !result.Mixed {
		t.Error("Mixed should be true")
	}
}

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
	// complexContent extending a builtin is unusual but we test the fallback.
	// This exercises the xsBuiltinGoType path in resolveRawField.
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
	// Unknown type with no schemas loaded → fallback to string.
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

	// Must compile as valid Go.
	if !strings.HasPrefix(src, "package geo") {
		t.Errorf("missing package declaration")
	}
	if !strings.Contains(src, "type DirectPositionType struct") {
		t.Error("DirectPositionType missing")
	}
	if !strings.Contains(src, "type PointType struct") {
		t.Error("PointType missing")
	}
	// DirectPositionType must have chardata + srsName attr as *string (optional).
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

// ---- Correctness regression tests for bugs found ----

// Bug1: resolveTypeName returned "" for simpleType with restriction of non-builtin base.
// That "" propagated into element ref field GoType → "*" or "[]" (invalid Go).
// Fix: resolveTypeName now returns "string" when simpleType.GoType == "".
func TestBug1_simpleTypeRestrictionNonBuiltinBase(t *testing.T) {
	// Schema with a simpleType that restricts a non-builtin (so GoType would be ""),
	// and an element using that type. The generated Go must not have an empty type.
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

	r := NewResolver()
	if _, err := r.Load(path); err != nil {
		t.Fatal(err)
	}
	types := r.ResolveAll("http://example.com/test")

	src, err := Generate(types, "test", false, false, nil, "", "")
	// Before the fix, Generate would produce `Code * \`xml:...\`` (invalid Go)
	// and go/format would fail. After the fix it must compile cleanly.
	if err != nil {
		t.Fatalf("Generate failed (likely invalid Go type): %v", err)
	}

	// The element ref field must have a valid type, not bare "*" or "[]".
	if strings.Contains(src, " * `") || strings.Contains(src, " [] `") {
		t.Errorf("generated code contains invalid bare pointer/slice type:\n%s", src)
	}
}

// Bug2: element ref path did not guard against empty GoType before pointer/slice wrapping.
// Same root cause as Bug1 but tested via element ref directly.
func TestBug2_elementRefWithEmptyGoType(t *testing.T) {
	// Minimal: element whose type is a simpleType restriction of non-builtin.
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

	r := NewResolver()
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
// This produced `Value  \`xml:",chardata"\" which is invalid Go.
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

	r := NewResolver()
	if _, err := r.Load(path); err != nil {
		t.Fatal(err)
	}
	types := r.ResolveAll("http://example.com/test")

	src, err := Generate(types, "test", false, false, nil, "", "")
	if err != nil {
		t.Fatalf("Generate produced invalid Go (bug3): %v", err)
	}
	// The chardata field must have a valid type.
	if strings.Contains(src, "Value  `") {
		t.Errorf("generated chardata field with empty type:\n%s", src)
	}
}
