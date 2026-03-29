package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

// FieldInfo holds a pointer field's Go name and XML element name.
type FieldInfo struct {
	Name    string
	XMLElem string
}

// parseGeneratedStructs parses all .go files in dir and returns:
//   - propFields: map of *PropertyType struct name → its pointer fields
//   - referenced: set of all struct names used as *Name pointer field types
func parseGeneratedStructs(dir string) (propFields map[string][]FieldInfo, referenced map[string]bool, err error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return nil, nil, err
	}

	propFields = make(map[string][]FieldInfo)
	referenced = make(map[string]bool)

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				gd, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				for _, spec := range gd.Specs {
					ts, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					st, ok := ts.Type.(*ast.StructType)
					if !ok {
						continue
					}
					visitStruct(ts.Name.Name, st, propFields, referenced)
				}
			}
		}
	}
	return propFields, referenced, nil
}

func visitStruct(
	name string,
	st *ast.StructType,
	propFields map[string][]FieldInfo,
	referenced map[string]bool,
) {
	isProp := strings.HasSuffix(name, "PropertyType")

	for _, field := range st.Fields.List {
		star, ok := field.Type.(*ast.StarExpr)
		if !ok {
			continue
		}
		ident, ok := star.X.(*ast.Ident)
		if !ok {
			continue
		}
		// Every type used as a *Pointer field is considered "referenced".
		referenced[ident.Name] = true

		if !isProp || len(field.Names) == 0 {
			continue
		}
		tag := ""
		if field.Tag != nil {
			tag = field.Tag.Value
		}
		elem := extractXMLElem(tag)
		if elem == "" {
			continue
		}
		propFields[name] = append(propFields[name], FieldInfo{
			Name:    field.Names[0].Name,
			XMLElem: elem,
		})
	}
}

var xmlTagRe = regexp.MustCompile(`xml:"([^"]*)"`)

// extractXMLElem extracts the XML local element name from a struct field tag.
// Returns "" for attribute tags or tags with no element name.
func extractXMLElem(tag string) string {
	m := xmlTagRe.FindStringSubmatch(tag)
	if m == nil {
		return ""
	}
	if strings.Contains(m[1], ",attr") {
		return ""
	}
	v := strings.SplitN(m[1], ",", 2)[0]
	parts := strings.Fields(v)
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}
