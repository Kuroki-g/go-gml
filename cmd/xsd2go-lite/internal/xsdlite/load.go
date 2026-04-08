package xsdlite

import (
	"fmt"
	"os"
	"path/filepath"
)

// Load recursively parses a schema and all its includes/imports.
func (r *Resolver) Load(filename string) (*Schema, error) {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	if s, ok := r.schemas[abs]; ok {
		return s, nil // already loaded
	}

	s, err := parseSchema(abs)
	if err != nil {
		return nil, err
	}
	// Register immediately to break cycles.
	r.schemas[abs] = s

	// Load includes/imports first (so we can resolve refs later).
	xsdRaw, err := rawIncludesImports(abs)
	if err != nil {
		return nil, err
	}

	baseDir := filepath.Dir(abs)
	for _, inc := range xsdRaw.includes {
		if isHTTP(inc) {
			fmt.Fprintf(os.Stderr, "warn: skipping external include %s\n", inc)
			continue
		}
		path := filepath.Join(baseDir, inc)
		if _, err := r.Load(path); err != nil {
			fmt.Fprintf(os.Stderr, "warn: failed to load include %s: %v\n", path, err)
		}
	}
	for _, imp := range xsdRaw.imports {
		var path string
		if imp.loc != "" && !isHTTP(imp.loc) {
			path = filepath.Join(baseDir, imp.loc)
		} else if imp.ns != "" {
			if mapped, ok := r.catalog[imp.ns]; ok {
				path = mapped
			}
		}
		if path == "" {
			if imp.ns != "" {
				fmt.Fprintf(os.Stderr, "warn: skipping external import ns=%s\n", imp.ns)
			}
			continue
		}
		if _, err := r.Load(path); err != nil {
			fmt.Fprintf(os.Stderr, "warn: failed to load import %s: %v\n", path, err)
		}
	}

	// Register types into global maps.
	ns := s.TargetNamespace

	// Merge this schema's namespace prefix declarations into nsMaps[ns].
	if _, ok := r.nsMaps[ns]; !ok {
		r.nsMaps[ns] = make(map[string]string)
	}
	for prefix, uri := range s.NSMap {
		// First-write wins: if multiple schemas with the same targetNS define the
		// same prefix differently, keep the first (earliest loaded) mapping.
		if _, exists := r.nsMaps[ns][prefix]; !exists {
			r.nsMaps[ns][prefix] = uri
		}
	}

	for i := range s.ComplexTypes {
		ct := &s.ComplexTypes[i]
		key := ns + " " + ct.Name
		r.allComplexTypes[key] = ct
	}
	for i := range s.SimpleTypes {
		st := &s.SimpleTypes[i]
		key := ns + " " + st.Name
		r.allSimpleTypes[key] = st
	}
	for name, typ := range s.Elements {
		key := ns + " " + name
		r.allElements[key] = typ
	}
	for name, doc := range s.ElementDocs {
		key := ns + " " + name
		r.allElementDocs[key] = doc
	}
	for name, ag := range s.AttributeGroups {
		agCopy := ag
		key := ns + " " + name
		r.allAttrGroups[key] = &agCopy
	}
	for name, fields := range s.Groups {
		key := ns + " " + name
		r.allGroups[key] = fields
	}

	// Build inverse substitutionGroup map: abstract head → concrete members.
	for elemName, headRef := range s.SubstGroups {
		headNS, headName := r.resolveQName(headRef, ns)
		headKey := headNS + " " + headName
		r.substHeads[headKey] = append(r.substHeads[headKey], substMember{
			NS:   ns,
			Name: elemName,
		})
	}

	return s, nil
}

// rawIncludeImport is a lightweight holder.
type rawIncludesImportsResult struct {
	includes []string
	imports  []struct{ ns, loc string }
}

// rawIncludesImports re-reads only the top-level include/import locations.
func rawIncludesImports(filename string) (rawIncludesImportsResult, error) {
	s, err := parseSchema(filename)
	if err != nil {
		return rawIncludesImportsResult{}, err
	}
	// We already have the info in the schema — but parseSchema doesn't store
	// include/import locations. Re-parse quickly using a simpler approach.
	// Actually we embed the info in Schema at load time below. Here we just
	// re-parse the raw xsdSchema. This is a bit wasteful but keeps things simple.
	_ = s
	// Re-open to get raw xml includes.
	return rawIncludesImportsFromFile(filename)
}

func rawIncludesImportsFromFile(filename string) (rawIncludesImportsResult, error) {
	f, err := os.Open(filename)
	if err != nil {
		return rawIncludesImportsResult{}, err
	}
	defer f.Close()

	var raw xsdSchema
	if err := decodeXML(f, &raw); err != nil {
		return rawIncludesImportsResult{}, fmt.Errorf("decode %s: %w", filename, err)
	}

	var res rawIncludesImportsResult
	for _, inc := range raw.Includes {
		res.includes = append(res.includes, inc.SchemaLocation)
	}
	for _, imp := range raw.Imports {
		res.imports = append(res.imports, struct{ ns, loc string }{imp.Namespace, imp.SchemaLocation})
	}
	return res, nil
}
