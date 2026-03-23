package xsdlite

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ---- buildXMLTag ----

func TestBuildXMLTag(t *testing.T) {
	tests := []struct {
		ns, local string
		isAttr    bool
		want      string
	}{
		{"", "name", false, "name"},
		{"http://example.com", "name", false, "http://example.com name"},
		{"", "srsName", true, "srsName,attr"},
		{"http://example.com", "id", true, "http://example.com id,attr"},
		{"", ",chardata", false, ",chardata"},
	}
	for _, tt := range tests {
		got := buildXMLTag(tt.ns, tt.local, tt.isAttr)
		if got != tt.want {
			t.Errorf("buildXMLTag(%q, %q, %v) = %q, want %q", tt.ns, tt.local, tt.isAttr, got, tt.want)
		}
	}
}

// ---- deduplicateFields ----

func TestDeduplicateFields(t *testing.T) {
	fields := []Field{
		{GoName: "Id", GoType: "string"},
		{GoName: "Name", GoType: "string"},
		{GoName: "Id", GoType: "int"}, // duplicate — should be dropped
		{GoName: "Value", GoType: "float64"},
	}
	got := deduplicateFields(fields)
	if len(got) != 3 {
		t.Fatalf("expected 3 unique fields, got %d", len(got))
	}
	// First occurrence wins
	if got[0].GoType != "string" {
		t.Errorf("first Id should be string, got %q", got[0].GoType)
	}
}

func TestDeduplicateFields_empty(t *testing.T) {
	if got := deduplicateFields(nil); got != nil {
		t.Errorf("nil input should return nil, got %v", got)
	}
	if got := deduplicateFields([]Field{}); len(got) != 0 {
		t.Errorf("empty input should return empty")
	}
}

// ---- isHTTP ----

func TestIsHTTP(t *testing.T) {
	if !isHTTP("http://www.opengis.net/gml/3.2") {
		t.Error("http:// should be HTTP")
	}
	if !isHTTP("https://example.com/schema.xsd") {
		t.Error("https:// should be HTTP")
	}
	if isHTTP("geometryBasic0d1d.xsd") {
		t.Error("relative path should not be HTTP")
	}
	if isHTTP("../schemas/gml.xsd") {
		t.Error("relative path with .. should not be HTTP")
	}
	if isHTTP("") {
		t.Error("empty string should not be HTTP")
	}
}

// ---- isBuiltinGoType ----

func TestIsBuiltinGoType(t *testing.T) {
	for _, bt := range []string{"string", "bool", "int", "float64", "float32"} {
		if !isBuiltinGoType(bt) {
			t.Errorf("isBuiltinGoType(%q) = false, want true", bt)
		}
	}
	for _, ct := range []string{"PointType", "[]string", "*Foo", "", "int64", "uint"} {
		if isBuiltinGoType(ct) {
			t.Errorf("isBuiltinGoType(%q) = true, want false", ct)
		}
	}
}

// ---- isStructType ----

func TestIsStructType(t *testing.T) {
	if isStructType("") {
		t.Error("empty string should not be struct type")
	}
	if isStructType("string") {
		t.Error("string is not a struct type")
	}
	if !isStructType("PointType") {
		t.Error("PointType should be a struct type")
	}
	if !isStructType("*DirectPositionType") {
		t.Error("pointer to struct should be a struct type")
	}
}

// ---- Resolver integration tests with temp XSD files ----

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
	byName := make(map[string]*ComplexType)
	for _, ct := range types {
		byName[ct.Name] = ct
	}

	derived, ok := byName["DerivedType"]
	if !ok {
		t.Fatal("DerivedType not found")
	}

	// DerivedType should have inherited fields from BaseType (name, id) + own (value).
	fieldNames := make(map[string]bool)
	for _, f := range derived.Fields {
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

func TestResolver_resolveQName_noPrefix(t *testing.T) {
	r := NewResolver()
	ns, name := r.resolveQName("MyType", "http://example.com")
	if ns != "http://example.com" {
		t.Errorf("ns = %q, want %q", ns, "http://example.com")
	}
	if name != "MyType" {
		t.Errorf("name = %q, want %q", name, "MyType")
	}
}

func TestResolver_resolveQName_withPrefix(t *testing.T) {
	r := NewResolver()
	// Inject namespace map for a schema with targetNS "http://example.com".
	r.nsMaps["http://example.com"] = map[string]string{
		"gml": "http://www.opengis.net/gml/3.2",
	}
	// Resolving from a type defined in "http://example.com" should use that
	// schema's NSMap, returning the correct URI for prefix "gml".
	ns, name := r.resolveQName("gml:PointType", "http://example.com")
	if ns != "http://www.opengis.net/gml/3.2" {
		t.Errorf("ns = %q, want GML namespace", ns)
	}
	if name != "PointType" {
		t.Errorf("name = %q, want PointType", name)
	}
}

func TestResolver_resolveQName_prefixDeterminism(t *testing.T) {
	r := NewResolver()
	// Two different schemas define the same prefix "ns" with different URIs.
	// The schema matching schemaNS should win (deterministic, not map-order dependent).
	r.nsMaps["http://schema-a.com"] = map[string]string{"ns": "http://ns-a.com"}
	r.nsMaps["http://schema-b.com"] = map[string]string{"ns": "http://ns-b.com"}

	// Resolving from schema-a must return ns-a, not ns-b.
	ns, _ := r.resolveQName("ns:Foo", "http://schema-a.com")
	if ns != "http://ns-a.com" {
		t.Errorf("resolveQName should prefer source schema's NSMap, got %q", ns)
	}

	// Resolving from schema-b must return ns-b.
	ns, _ = r.resolveQName("ns:Bar", "http://schema-b.com")
	if ns != "http://ns-b.com" {
		t.Errorf("resolveQName should prefer source schema's NSMap, got %q", ns)
	}
}

func TestResolver_resolveQName_unknownPrefix(t *testing.T) {
	r := NewResolver()
	// Unknown prefix falls back to schemaNS.
	ns, name := r.resolveQName("xyz:Foo", "http://fallback")
	if ns != "http://fallback" {
		t.Errorf("unknown prefix should fall back to schemaNS, got %q", ns)
	}
	if name != "Foo" {
		t.Errorf("name = %q, want Foo", name)
	}
}

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

	fieldMap := make(map[string]Field)
	for _, f := range mt.Fields {
		fieldMap[f.GoName] = f
	}

	// chardata Value field
	vf, ok := fieldMap["Value"]
	if !ok {
		t.Error("Value (chardata) field missing")
	} else if !vf.IsChar {
		t.Error("Value field should be chardata")
	}

	// uom attribute (required → not omit)
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
	byName := make(map[string]*ComplexType)
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
	byName := make(map[string]*ComplexType)
	for _, ct := range types {
		byName[ct.Name] = ct
	}

	prop, ok := byName["AbstractRingPropertyType"]
	if !ok {
		t.Fatal("AbstractRingPropertyType not found")
	}

	fieldMap := make(map[string]Field)
	for _, f := range prop.Fields {
		fieldMap[f.GoName] = f
	}

	// The abstract ref itself should still be present.
	if _, ok := fieldMap["AbstractRing"]; !ok {
		t.Error("AbstractRing field missing")
	}

	// LinearRing should be added as a substitution member.
	lrf, ok := fieldMap["LinearRing"]
	if !ok {
		t.Error("LinearRing substitution member field missing")
	} else {
		if lrf.GoType != "*LinearRingType" {
			t.Errorf("LinearRing GoType = %q, want *LinearRingType", lrf.GoType)
		}
		if !lrf.Omit {
			t.Error("LinearRing field should have omitempty (Omit=true)")
		}
		if lrf.Slice {
			t.Error("LinearRing field should not be a slice")
		}
	}

	// Ring should be added as a substitution member.
	rf, ok := fieldMap["Ring"]
	if !ok {
		t.Error("Ring substitution member field missing")
	} else {
		if rf.GoType != "*RingType" {
			t.Errorf("Ring GoType = %q, want *RingType", rf.GoType)
		}
	}

	// Total fields: AbstractRing + LinearRing + Ring = 3.
	if len(prop.Fields) != 3 {
		t.Errorf("expected 3 fields, got %d: %v", len(prop.Fields), prop.Fields)
	}
}

func TestResolver_ResolveAll_substitutionGroup_noMembers(t *testing.T) {
	// If an abstract element has no substitution group members registered,
	// only the abstract field should appear (no expansion).
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
	byName := make(map[string]*ComplexType)
	for _, ct := range types {
		byName[ct.Name] = ct
	}

	container, ok := byName["ContainerType"]
	if !ok {
		t.Fatal("ContainerType not found")
	}
	if len(container.Fields) != 1 {
		t.Errorf("expected 1 field (no subst members), got %d", len(container.Fields))
	}
	if container.Fields[0].GoName != "AbstractGeom" {
		t.Errorf("field GoName = %q, want AbstractGeom", container.Fields[0].GoName)
	}
}

// ---- deduplicateFields: attr vs element priority ----

// Bug 1: inherited attribute + own element → element should win.
// (e.g. GridType.axisLabels: SRSInformationGroup gives attr, GridType sequence gives element)
func TestDeduplicateFields_inheritedAttrOwnElement_elementWins(t *testing.T) {
	fields := []Field{
		{GoName: "AxisLabels", GoType: "string", IsAttr: true},  // inherited attr (comes first)
		{GoName: "AxisLabels", GoType: "string", IsAttr: false}, // own element (comes after)
	}
	got := deduplicateFields(fields)
	if len(got) != 1 {
		t.Fatalf("expected 1 field, got %d: %+v", len(got), got)
	}
	if got[0].IsAttr {
		t.Error("own element should win over inherited attribute, but got attr")
	}
}

// Regression: inherited element + own attribute → attribute should still win.
// (e.g. GML 3.1.1 srsName: some base gives element, attributeGroup gives attr)
func TestDeduplicateFields_inheritedElementOwnAttr_attrWins(t *testing.T) {
	fields := []Field{
		{GoName: "SrsName", GoType: "string", IsAttr: false}, // inherited element (comes first)
		{GoName: "SrsName", GoType: "string", IsAttr: true},  // own attribute (comes after)
	}
	got := deduplicateFields(fields)
	if len(got) != 1 {
		t.Fatalf("expected 1 field, got %d: %+v", len(got), got)
	}
	if !got[0].IsAttr {
		t.Error("own attribute should win over inherited element, but got element")
	}
}

// ---- substitutionGroup: transitive (multi-level) expansion ----

// Bug 2: sortedSubstMembers must expand transitively.
// Chain: AbstractGeom → AbstractImplicitGeom (1-level) → Grid (2-level) → RectifiedGrid (3-level).
// A container that refs AbstractGeom should get fields for ALL transitive members.
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

	// Direct ref field must be present.
	if _, ok := fieldMap["AbstractGeom"]; !ok {
		t.Error("AbstractGeom field missing")
	}
	// 1-level transitive member must be present.
	if _, ok := fieldMap["AbstractImplicitGeom"]; !ok {
		t.Error("AbstractImplicitGeom (1-level transitive) field missing")
	}
	// 2-level transitive member must be present.
	if gf, ok := fieldMap["Grid"]; !ok {
		t.Error("Grid (2-level transitive) field missing")
	} else if gf.GoType != "*GridType" {
		t.Errorf("Grid GoType = %q, want *GridType", gf.GoType)
	}
	// 3-level transitive member must be present.
	if rf, ok := fieldMap["RectifiedGrid"]; !ok {
		t.Error("RectifiedGrid (3-level transitive) field missing")
	} else if rf.GoType != "*RectifiedGridType" {
		t.Errorf("RectifiedGrid GoType = %q, want *RectifiedGridType", rf.GoType)
	}
}

// ---- rawIncludesImports ----

func TestRawIncludesImportsFromFile(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com"
        xmlns="http://www.w3.org/2001/XMLSchema">
  <include schemaLocation="other.xsd"/>
  <import namespace="http://www.w3.org/1999/xlink" schemaLocation="xlink.xsd"/>
  <import namespace="http://external.example.com"/>
</schema>`

	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(xsd), 0644); err != nil {
		t.Fatal(err)
	}

	res, err := rawIncludesImportsFromFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.includes) != 1 || res.includes[0] != "other.xsd" {
		t.Errorf("includes = %v, want [other.xsd]", res.includes)
	}
	if len(res.imports) != 2 {
		t.Errorf("imports = %d, want 2", len(res.imports))
	}
	if res.imports[0].loc != "xlink.xsd" {
		t.Errorf("import[0].loc = %q, want xlink.xsd", res.imports[0].loc)
	}
	if res.imports[1].loc != "" {
		t.Errorf("import[1].loc should be empty (no schemaLocation)")
	}
}
