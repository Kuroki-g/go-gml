package resolve

import (
	"fmt"

	"xsd2go-lite2/internal/parse"
)

// Resolver loads XSD schemas and resolves ComplexType fields.
// This is a stub; full implementation is in Step 3/4.
type Resolver struct {
	catalog map[string]string
}

func NewResolver() *Resolver {
	return &Resolver{catalog: make(map[string]string)}
}

func (r *Resolver) AddCatalogEntry(ns, path string) {
	r.catalog[ns] = path
}

func (r *Resolver) Load(filename string) (*parse.Schema, error) {
	return nil, fmt.Errorf("resolve.Resolver.Load: not yet implemented")
}

func (r *Resolver) ResolveAll(namespace string) []*parse.ComplexType {
	return nil
}
