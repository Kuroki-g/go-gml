package parse

// Schema is the fully-loaded (but not yet resolved) schema data.
type Schema struct {
	TargetNamespace string
	Filename        string
	NSMap           map[string]string // prefix → URI
	ComplexTypes    []ComplexType
	SimpleTypes     []SimpleType
	Elements        map[string]Element
	GlobalAttrs     map[string]Attribute
	AttributeGroups map[string]AttrGroup
	Groups          map[string]Group
}

// ComplexType is the (partially) resolved complex type definition.
type ComplexType struct {
	Name     string
	Abstract bool
	Mixed    bool
	Doc      string
	Source   string // targetNamespace of defining schema

	ContentKind string // ContentKindComplex | ContentKindSimple | ""

	Derivation *Derivation   // nil = direct definition
	Content    *ContentModel // element content (nil for simpleContent)
	Attrs      []AttrDecl    // own attribute declarations
	AttrGroups []string      // attributeGroup ref QNames
	CharData   *CharDataDecl // simpleContent chardata (nil for complexContent)

	// After resolve:
	Fields []Field
	Embeds []EmbedRef // attributeGroup/group refs, populated by resolve phase
}

// Derivation holds XSD extension/restriction info.
type Derivation struct {
	Kind string // "extension" | "restriction"
	Base string // base type QName
}

// ContentModel preserves the sequence/choice/all compositor structure.
type ContentModel struct {
	Kind      string // "sequence" | "choice" | "all"
	MinOccurs string
	MaxOccurs string
	Elems     []ElemDecl
	Groups    []GroupRef
	Children  []ContentModel // nested compositors
}

// ElemDecl is a single element declaration inside a content model.
type ElemDecl struct {
	Name      string // local name (empty if ref)
	Ref       string // element ref QName (empty if local)
	TypeRef   string
	MinOccurs string
	MaxOccurs string
	Default   string
	Fixed     string
	Doc       string
}

// GroupRef is a reference to a named model group.
type GroupRef struct {
	Ref       string
	MinOccurs string
	MaxOccurs string
}

// AttrDecl is a single attribute declaration.
type AttrDecl struct {
	Name    string // local name (empty if ref)
	Ref     string // attribute ref QName (empty if local)
	TypeRef string
	Use     string // "required" | "optional" | "prohibited" | ""
	Default string
	Fixed   string
	Doc     string
}

// CharDataDecl represents the text content of a simpleContent type.
type CharDataDecl struct {
	TypeRef string // base type QName
}

// AttrGroup is an attribute group definition.
type AttrGroup struct {
	Name         string
	Attrs        []AttrDecl
	NestedGroups []string // nested attributeGroup ref QNames
}

// Group is a named model group definition.
type Group struct {
	Name    string
	Content *ContentModel
}

// Element is a global element declaration.
type Element struct {
	Name              string
	TypeRef           string
	SubstitutionGroup string // head element QName
	Abstract          bool
	Doc               string
}

// Attribute is a global attribute declaration.
type Attribute struct {
	Name    string
	TypeRef string
	Default string
}

// SimpleType is a simple type definition.
type SimpleType struct {
	Name         string
	Kind         string // "list" | "restriction" | "union" | "alias"
	GoType       string
	Enumerations []string // enumeration facet values
	Base         string   // restriction base type QName
}

// EmbedRef records a named attributeGroup or group for future embed struct generation.
type EmbedRef struct {
	XSDName string // local name
	NS      string // namespace URI
	Kind    string // "attributeGroup" | "group"
}

// Field is a single resolved struct field (gen phase input).
type Field struct {
	GoName string
	XMLTag string
	GoType string
	TypeNS string
	IsAttr bool
	IsChar bool
	Omit   bool
	Slice  bool
	Doc    string
}

// XSD attribute use= values.
const (
	UseRequired   = "required"
	UseOptional   = "optional"
	UseProhibited = "prohibited"
)

// ContentKind values for ComplexType.ContentKind.
const (
	ContentKindSimple  = "simple"
	ContentKindComplex = "complex"
)
