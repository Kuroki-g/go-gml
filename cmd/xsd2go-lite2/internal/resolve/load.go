package resolve

import (
	"fmt"
	"os"
	"path/filepath"

	"xsd2go-lite2/internal/parse"
)

// Load recursively parses a schema and all its includes/imports.
func (r *Resolver) Load(filename string) (*parse.Schema, error) {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	if s, ok := r.schemas[abs]; ok {
		return s, nil
	}

	s, err := parse.Parse(abs)
	if err != nil {
		return nil, err
	}
	r.schemas[abs] = s

	raw, err := rawIncludesImports(abs)
	if err != nil {
		return nil, err
	}

	baseDir := filepath.Dir(abs)
	for _, inc := range raw.includes {
		if isHTTP(inc) {
			fmt.Fprintf(os.Stderr, "warn: skipping external include %s\n", inc)
			continue
		}
		path := filepath.Join(baseDir, inc)
		if _, err := r.Load(path); err != nil {
			fmt.Fprintf(os.Stderr, "warn: failed to load include %s: %v\n", path, err)
		}
	}
	for _, imp := range raw.imports {
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

	ns := s.TargetNamespace

	if _, ok := r.nsMaps[ns]; !ok {
		r.nsMaps[ns] = make(map[string]string)
	}
	for prefix, uri := range s.NSMap {
		if _, exists := r.nsMaps[ns][prefix]; !exists {
			r.nsMaps[ns][prefix] = uri
		}
	}

	for i := range s.ComplexTypes {
		ct := &s.ComplexTypes[i]
		r.allComplexTypes[ns+" "+ct.Name] = ct
	}
	for i := range s.SimpleTypes {
		st := &s.SimpleTypes[i]
		r.allSimpleTypes[ns+" "+st.Name] = st
	}
	for name, elem := range s.Elements {
		elemCopy := elem
		r.allElements[ns+" "+name] = &elemCopy
	}
	for name, attr := range s.GlobalAttrs {
		attrCopy := attr
		r.allGlobalAttrs[ns+" "+name] = &attrCopy
	}
	for name, ag := range s.AttributeGroups {
		agCopy := ag
		agCopy.Source = ns
		r.allAttrGroups[ns+" "+name] = &agCopy
	}
	for name, g := range s.Groups {
		gCopy := g
		r.allGroups[ns+" "+name] = &gCopy
	}

	// Build inverse substitutionGroup map: head element → concrete members.
	for _, elem := range s.Elements {
		if elem.SubstitutionGroup == "" {
			continue
		}
		headNS, headName := r.resolveQName(elem.SubstitutionGroup, ns)
		headKey := headNS + " " + headName
		r.substHeads[headKey] = append(r.substHeads[headKey], substMember{
			NS:   ns,
			Name: elem.Name,
		})
	}

	return s, nil
}

type rawIncludesImportsResult struct {
	includes []string
	imports  []struct{ ns, loc string }
}

func rawIncludesImports(filename string) (rawIncludesImportsResult, error) {
	f, err := os.Open(filename)
	if err != nil {
		return rawIncludesImportsResult{}, err
	}
	defer f.Close()

	type xsdInclude struct {
		SchemaLocation string `xml:"schemaLocation,attr"`
	}
	type xsdImport struct {
		Namespace      string `xml:"namespace,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	}
	type xsdSchema struct {
		Includes []xsdInclude `xml:"include"`
		Imports  []xsdImport  `xml:"import"`
	}

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
