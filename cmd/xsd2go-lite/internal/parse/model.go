package parse

// ---- Internal representation ----

// Schema is the fully-loaded (but not yet resolved) schema data.
type Schema struct {
	TargetNamespace string
	Filename        string
	NSMap           map[string]string // prefix → URI
	ComplexTypes    []ComplexType
	SimpleTypes     []SimpleType
	Elements        map[string]string    // name → type (global elements)
	ElementDocs     map[string]string    // name → xs:annotation/xs:documentation (global elements)
	SubstGroups     map[string]string    // concrete element name → abstract element head (QName)
	GlobalAttrs     map[string]Attribute // name → Attribute
	AttributeGroups map[string]AttrGroup
	Groups          map[string][]RawField
}

// RawField holds unresolved field info (before resolve phase).
type RawField struct {
	LocalName string // element/attribute local name (may be empty if ref)
	Ref       string // ref="ns:name" (for element refs)
	AttrRef   string // ref="ns:name" (for attribute refs)
	TypeRef   string // type="ns:name"
	IsAttr    bool
	IsChar    bool // simpleContent chardata
	MinOccurs string
	MaxOccurs string
	GroupRef  string // group ref="ns:name"
	Doc       string // xs:annotation/xs:documentation text
}

// ComplexType is the resolved complex type definition.
type ComplexType struct {
	Name       string
	Abstract   bool
	Mixed      bool
	Doc        string // xs:annotation/xs:documentation text
	RawFields  []RawField
	AttrGroups []string // list of attributeGroup ref names
	Source     string   // targetNamespace of the defining schema
	// After resolve:
	Fields []Field
}

// Field is a single resolved struct field.
type Field struct {
	GoName string
	XMLTag string
	GoType string
	TypeNS string // namespace URI of the referenced Go type (for omit filtering)
	IsAttr bool
	IsChar bool
	Omit   bool
	Slice  bool
	Doc    string // xs:annotation/xs:documentation text
}

// SimpleType is the resolved simple type definition.
type SimpleType struct {
	Name   string
	Kind   string // "list" | "restriction" | "union" | "alias"
	GoType string
}

// AttrGroup is an attribute group definition.
type AttrGroup struct {
	Name      string
	RawFields []RawField
}

// Attribute is a global attribute definition.
type Attribute struct {
	Name    string
	Type    string
	Default string
}
