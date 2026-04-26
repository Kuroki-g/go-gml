package resolve

import "xsd2go-lite2/internal/parse"

// substMember holds a concrete element that is a member of a substitutionGroup.
type substMember struct {
	NS   string // target namespace URI
	Name string // local element name
}

// Resolver loads XSD schemas and resolves ComplexType fields.
type Resolver struct {
	schemas         map[string]*parse.Schema
	allComplexTypes map[string]*parse.ComplexType
	allSimpleTypes  map[string]*parse.SimpleType
	allElements     map[string]*parse.Element
	allGlobalAttrs  map[string]*parse.Attribute
	allAttrGroups   map[string]*parse.AttrGroup
	allGroups       map[string]*parse.Group
	substHeads      map[string][]substMember
	nsMaps          map[string]map[string]string
	catalog         map[string]string
}

// NewResolver creates a new Resolver.
func NewResolver() *Resolver {
	return &Resolver{
		schemas:         make(map[string]*parse.Schema),
		allComplexTypes: make(map[string]*parse.ComplexType),
		allSimpleTypes:  make(map[string]*parse.SimpleType),
		allElements:     make(map[string]*parse.Element),
		allGlobalAttrs:  make(map[string]*parse.Attribute),
		allAttrGroups:   make(map[string]*parse.AttrGroup),
		allGroups:       make(map[string]*parse.Group),
		substHeads:      make(map[string][]substMember),
		nsMaps:          make(map[string]map[string]string),
		catalog:         make(map[string]string),
	}
}

// AddCatalogEntry registers a namespace URI → local file path mapping.
func (r *Resolver) AddCatalogEntry(namespaceURI, localPath string) {
	r.catalog[namespaceURI] = localPath
}

// AllAttrGroups returns the full map of resolved attributeGroup definitions,
// keyed by "NS LocalName". Intended for use by the gen phase.
func (r *Resolver) AllAttrGroups() map[string]*parse.AttrGroup {
	return r.allAttrGroups
}
