package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// substMember holds a concrete element that is a member of a substitutionGroup.
type substMember struct {
	NS   string // target namespace URI of the concrete element
	Name string // local name of the concrete element
}

// Resolver holds all loaded schemas and resolves cross-schema references.
type Resolver struct {
	// All loaded schemas keyed by absolute filename.
	schemas map[string]*Schema
	// All complexTypes across all schemas, keyed by "targetNS localName".
	allComplexTypes map[string]*ComplexType
	// All simpleTypes across all schemas, keyed by "targetNS localName".
	allSimpleTypes map[string]*SimpleType
	// All global elements across all schemas, keyed by "targetNS localName" → type string.
	allElements map[string]string
	// All attributeGroups across all schemas, keyed by "targetNS localName".
	allAttrGroups map[string]*AttrGroup
	// All groups across all schemas, keyed by "targetNS localName".
	allGroups map[string][]rawField
	// substHeads maps abstract element key ("ns localName") → concrete members.
	substHeads map[string][]substMember
	// nsMaps maps targetNamespace → {prefix → uri}, built from each loaded schema's
	// xmlns declarations. Used by resolveQName to deterministically resolve prefixes
	// against the schema that defined the type being resolved.
	nsMaps map[string]map[string]string
}

func newResolver() *Resolver {
	return &Resolver{
		schemas:         make(map[string]*Schema),
		allComplexTypes: make(map[string]*ComplexType),
		allSimpleTypes:  make(map[string]*SimpleType),
		allElements:     make(map[string]string),
		allAttrGroups:   make(map[string]*AttrGroup),
		allGroups:       make(map[string][]rawField),
		substHeads:      make(map[string][]substMember),
		nsMaps:          make(map[string]map[string]string),
	}
}

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
		if imp.loc == "" || isHTTP(imp.loc) {
			if imp.ns != "" {
				fmt.Fprintf(os.Stderr, "warn: skipping external import ns=%s\n", imp.ns)
			}
			continue
		}
		path := filepath.Join(baseDir, imp.loc)
		if _, err := r.Load(path); err != nil {
			fmt.Fprintf(os.Stderr, "warn: failed to load import %s: %v\n", path, err)
		}
	}

	// Register types into global maps.
	ns := s.TargetNamespace

	// Merge this schema's namespace prefix declarations into nsMaps[ns].
	// This allows resolveQName to look up the correct URI for a prefix by
	// using the declaring schema's namespace map, avoiding map-iteration
	// non-determinism when multiple schemas share the same targetNamespace.
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

// ResolveAll resolves all loaded schemas: flattens inheritance, inlines attributeGroups.
func (r *Resolver) ResolveAll(targetNS string) []*ComplexType {
	// Resolve each complex type's fields.
	for _, ct := range r.allComplexTypes {
		if len(ct.Fields) == 0 {
			r.resolveComplexType(ct, nil)
		}
	}

	// Return only types belonging to the target namespace.
	var result []*ComplexType
	for key, ct := range r.allComplexTypes {
		if strings.HasPrefix(key, targetNS+" ") {
			result = append(result, ct)
		}
	}
	return result
}

// resolveComplexType builds the final Fields slice (depth-first, handles cycles).
func (r *Resolver) resolveComplexType(ct *ComplexType, visiting map[string]bool) {
	if visiting == nil {
		visiting = make(map[string]bool)
	}
	key := ct.Source + " " + ct.Name
	if visiting[key] {
		return // cycle guard
	}
	visiting[key] = true
	defer delete(visiting, key)

	if len(ct.Fields) > 0 {
		return // already resolved
	}

	var fields []Field

	for _, rf := range ct.RawFields {
		resolved := r.resolveRawField(rf, ct.Source, visiting)
		fields = append(fields, resolved...)
	}

	// Inline attributeGroups
	for _, agRef := range ct.AttrGroups {
		agFields := r.resolveAttrGroupRef(agRef, ct.Source)
		fields = append(fields, agFields...)
	}

	ct.Fields = deduplicateFields(fields)
}

// resolveRawField converts a single rawField to zero or more Fields.
func (r *Resolver) resolveRawField(rf rawField, schemaNS string, visiting map[string]bool) []Field {
	// Base type inheritance placeholder
	if rf.Ref != "" && strings.HasPrefix(rf.Ref, "__base__:") {
		baseRef := strings.TrimPrefix(rf.Ref, "__base__:")
		baseNS, baseName := r.resolveQName(baseRef, schemaNS)
		baseKey := baseNS + " " + baseName
		if baseCT, ok := r.allComplexTypes[baseKey]; ok {
			r.resolveComplexType(baseCT, visiting)
			return baseCT.Fields
		}
		// If base is a simpleType, add a chardata field.
		if baseST, ok := r.allSimpleTypes[baseKey]; ok {
			goType := baseST.GoType
			if goType == "" {
				goType = "string"
			}
			return []Field{{GoName: "Value", GoType: goType, IsChar: true, XMLTag: ",chardata"}}
		}
		// If it's an XSD built-in
		if goType := xsBuiltinGoType(baseRef); goType != "" {
			return []Field{{GoName: "Value", GoType: goType, IsChar: true, XMLTag: ",chardata"}}
		}
		fmt.Fprintf(os.Stderr, "warn: base type not found: %s\n", baseRef)
		return nil
	}

	// Element ref (e.g. ref="gml:pointMember")
	if rf.Ref != "" {
		refNS, refName := r.resolveQName(rf.Ref, schemaNS)
		refKey := refNS + " " + refName
		elemType := r.allElements[refKey]
		if elemType == "" {
			fmt.Fprintf(os.Stderr, "warn: element ref not found: %s\n", rf.Ref)
			elemType = "string"
		}
		resolvedType := r.resolveTypeName(elemType, refNS)
		if resolvedType == "" {
			resolvedType = "string"
		}
		f := Field{
			GoName: goName(refName),
			XMLTag: buildXMLTag(refNS, refName, false),
			GoType: resolvedType,
			Omit:   rf.MinOccurs == "0" || rf.MinOccurs == "",
			Slice:  rf.MaxOccurs == "unbounded" || (rf.MaxOccurs != "" && rf.MaxOccurs != "1"),
		}
		if f.Slice {
			f.GoType = "[]" + strings.TrimPrefix(f.GoType, "[]")
		} else {
			f.GoType = "*" + strings.TrimPrefix(f.GoType, "*")
		}
		result := []Field{f}
		// Expand substitutionGroup: add a parallel field for each concrete member.
		for _, member := range r.sortedSubstMembers(refKey) {
			memberElemType := r.allElements[member.NS+" "+member.Name]
			memberGoType := r.resolveTypeName(memberElemType, member.NS)
			if memberGoType == "" {
				memberGoType = "string"
			}
			memberGoType = "*" + strings.TrimPrefix(memberGoType, "*")
			result = append(result, Field{
				GoName: goName(member.Name),
				XMLTag: buildXMLTag(member.NS, member.Name, false),
				GoType: memberGoType,
				Omit:   true,
				Slice:  false,
			})
		}
		return result
	}

	// Attribute ref
	if rf.IsAttr && rf.AttrRef != "" {
		// Nested attributeGroup ref (encoded as "__group__:ref")
		if strings.HasPrefix(rf.AttrRef, "__group__:") {
			agRef := strings.TrimPrefix(rf.AttrRef, "__group__:")
			return r.resolveAttrGroupRef(agRef, schemaNS)
		}
		refNS, refName := r.resolveQName(rf.AttrRef, schemaNS)
		refKey := refNS + " " + refName
		goType := "string"
		if ga, ok := r.allAttrGroups[refKey]; ok {
			_ = ga // attribute group as attribute ref is unusual; treat as string
		}
		return []Field{{
			GoName: goName(refName),
			XMLTag: buildXMLTag("", refName, true),
			GoType: goType,
			IsAttr: true,
			Omit:   true,
		}}
	}

	// Group ref
	if rf.GroupRef != "" {
		groupNS, groupName := r.resolveQName(rf.GroupRef, schemaNS)
		groupKey := groupNS + " " + groupName
		groupFields, ok := r.allGroups[groupKey]
		if !ok {
			fmt.Fprintf(os.Stderr, "warn: group ref not found: %s\n", rf.GroupRef)
			return nil
		}
		var result []Field
		for _, gf := range groupFields {
			result = append(result, r.resolveRawField(gf, groupNS, visiting)...)
		}
		return result
	}

	// chardata field
	if rf.IsChar {
		goType := r.resolveTypeName(rf.TypeRef, schemaNS)
		// chardata for simpleContent is always string-like
		if goType == "" || goType == "string" || isStructType(goType) {
			goType = "string"
		}
		return []Field{{
			GoName: "Value",
			XMLTag: ",chardata",
			GoType: goType,
			IsChar: true,
		}}
	}

	// Normal named element or attribute
	if rf.LocalName == "" || rf.LocalName == "__chardata__" {
		return nil
	}

	goType := r.resolveTypeName(rf.TypeRef, schemaNS)
	if goType == "" {
		goType = "string"
	}

	isSlice := rf.MaxOccurs == "unbounded" || (rf.MaxOccurs != "" && rf.MaxOccurs != "1" && rf.MaxOccurs != "0")
	isOmit := rf.IsAttr && rf.MinOccurs != "required"
	if !rf.IsAttr {
		isOmit = rf.MinOccurs == "0" || rf.MinOccurs == ""
	}

	xmlTag := buildXMLTag(schemaNS, rf.LocalName, rf.IsAttr)
	if rf.IsAttr {
		// Attributes don't use namespace URI in tag unless explicitly from another NS.
		xmlTag = buildXMLTag("", rf.LocalName, true)
	}

	f := Field{
		GoName: goName(rf.LocalName),
		XMLTag: xmlTag,
		GoType: goType,
		IsAttr: rf.IsAttr,
		Omit:   isOmit,
		Slice:  isSlice,
	}

	if isSlice && !rf.IsAttr {
		f.GoType = "[]" + strings.TrimPrefix(goType, "[]")
	} else if !isSlice && !rf.IsAttr && !isBuiltinGoType(goType) {
		f.GoType = "*" + strings.TrimPrefix(goType, "*")
	}

	return []Field{f}
}

// resolveAttrGroupRef resolves an attributeGroup ref to Fields.
func (r *Resolver) resolveAttrGroupRef(ref, schemaNS string) []Field {
	agNS, agName := r.resolveQName(ref, schemaNS)
	agKey := agNS + " " + agName
	ag, ok := r.allAttrGroups[agKey]
	if !ok {
		fmt.Fprintf(os.Stderr, "warn: attributeGroup not found: %s\n", ref)
		return nil
	}

	var fields []Field
	for _, rf := range ag.RawFields {
		// Nested group ref in attrGroup
		if rf.IsAttr && strings.HasPrefix(rf.AttrRef, "__group__:") {
			nestedRef := strings.TrimPrefix(rf.AttrRef, "__group__:")
			fields = append(fields, r.resolveAttrGroupRef(nestedRef, agNS)...)
			continue
		}
		f := r.resolveRawField(rf, agNS, nil)
		fields = append(fields, f...)
	}
	return fields
}

// resolveTypeName converts a QName type reference to a Go type string.
func (r *Resolver) resolveTypeName(typeRef, schemaNS string) string {
	if typeRef == "" || typeRef == "string" {
		return "string"
	}
	if typeRef == "__inline__" {
		return "string"
	}

	// Check XSD built-in first
	if goType := xsBuiltinGoType(typeRef); goType != "" {
		return goType
	}

	ns, name := r.resolveQName(typeRef, schemaNS)
	key := ns + " " + name

	// simpleType?
	if st, ok := r.allSimpleTypes[key]; ok {
		// GoType can be "" when the simpleType is a restriction of a non-builtin
		// base (e.g. restriction base="someCustomType"). Fall back to string so
		// callers never receive an empty type string that would generate invalid Go.
		if st.GoType == "" {
			return "string"
		}
		return st.GoType
	}
	// complexType?
	if _, ok := r.allComplexTypes[key]; ok {
		return goTypeName(name)
	}

	// Fallback: use the local name as a struct type
	fmt.Fprintf(os.Stderr, "warn: type not found: %s (ns=%s)\n", typeRef, schemaNS)
	return "string"
}

// resolveQName resolves a possibly-prefixed name ("gml:Foo" or "Foo") to (ns, localName).
func (r *Resolver) resolveQName(qname, schemaNS string) (string, string) {
	idx := strings.Index(qname, ":")
	if idx < 0 {
		return schemaNS, qname
	}
	prefix := qname[:idx]
	local := qname[idx+1:]

	// First, try the namespace map for the schema that defines the type being
	// resolved (identified by schemaNS). This is deterministic and correct:
	// the prefix must be declared in the schema that uses it.
	if nsMap, ok := r.nsMaps[schemaNS]; ok {
		if uri, ok := nsMap[prefix]; ok {
			return uri, local
		}
	}
	// Fall back: search other schemas in deterministic (sorted) order.
	// This handles cases where the type's schema imports another schema's
	// prefix without re-declaring it (unusual but valid).
	keys := make([]string, 0, len(r.nsMaps))
	for k := range r.nsMaps {
		keys = append(keys, k)
	}
	sortStrings(keys)
	for _, k := range keys {
		if k == schemaNS {
			continue // already tried above
		}
		if uri, ok := r.nsMaps[k][prefix]; ok {
			return uri, local
		}
	}
	// Unknown prefix — return as-is with original schemaNS.
	fmt.Fprintf(os.Stderr, "warn: unknown prefix %q in %s\n", prefix, qname)
	return schemaNS, local
}

// buildXMLTag constructs the xml struct tag value.
func buildXMLTag(ns, localName string, isAttr bool) string {
	var sb strings.Builder
	if ns != "" {
		sb.WriteString(ns)
		sb.WriteString(" ")
	}
	sb.WriteString(localName)
	if isAttr {
		sb.WriteString(",attr")
	}
	return sb.String()
}

// deduplicateFields removes fields with duplicate GoNames, keeping the first.
func deduplicateFields(fields []Field) []Field {
	seen := make(map[string]bool)
	var result []Field
	for _, f := range fields {
		if !seen[f.GoName] {
			seen[f.GoName] = true
			result = append(result, f)
		}
	}
	return result
}

// sortedSubstMembers returns the substitution group members for a head element key,
// sorted deterministically by (NS, Name).
func (r *Resolver) sortedSubstMembers(headKey string) []substMember {
	members := r.substHeads[headKey]
	if len(members) == 0 {
		return nil
	}
	sorted := make([]substMember, len(members))
	copy(sorted, members)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].NS != sorted[j].NS {
			return sorted[i].NS < sorted[j].NS
		}
		return sorted[i].Name < sorted[j].Name
	})
	return sorted
}

func isHTTP(loc string) bool {
	return strings.HasPrefix(loc, "http://") || strings.HasPrefix(loc, "https://")
}

func isBuiltinGoType(t string) bool {
	switch t {
	case "string", "bool", "int", "float64", "float32":
		return true
	}
	return false
}

func isStructType(t string) bool {
	return !isBuiltinGoType(t) && t != ""
}

func sortStrings(s []string) {
	sort.Strings(s)
}
