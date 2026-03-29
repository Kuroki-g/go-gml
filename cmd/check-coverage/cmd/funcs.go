package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

// FuncAccess records which fields were accessed within one function.
type FuncAccess struct {
	FuncName string
	Fields   map[string]bool // set of Go field names accessed via selector exprs
}

// analyzeFunctions scans all .go files in dir and returns, for each
// PropertyType name, the list of functions that handle it along with their
// field accesses.
//
// A function "handles" a PropertyType when at least one of its parameters
// (including the receiver) has a base type matching a known PropertyType name.
// Base type resolution strips *, [], and the imported package alias so that
// "*gen.SolidPropertyType", "[]gen.SolidPropertyType", and "SolidPropertyType"
// all resolve to "SolidPropertyType".
//
// Field accesses are all selector expressions (x.Field) found anywhere in the
// function body. Transitive accesses inside called helpers are intentionally
// not followed: each function must directly reference or explicitly blank
// (_ = m.Field) every field it is responsible for.
func analyzeFunctions(dir string, knownTypes map[string][]FieldInfo) (map[string][]FuncAccess, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, err
	}

	result := make(map[string][]FuncAccess)

	for _, path := range files {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return nil, err
		}
		alias := genPackageAlias(f)

		for _, decl := range f.Decls {
			fd, ok := decl.(*ast.FuncDecl)
			if !ok || fd.Body == nil {
				continue
			}
			handled := handledTypes(fd, alias, knownTypes)
			if len(handled) == 0 {
				continue
			}
			accessed := selectorFields(fd.Body)
			name := qualifiedFuncName(fd)
			for _, typeName := range handled {
				result[typeName] = append(result[typeName], FuncAccess{
					FuncName: name,
					Fields:   accessed,
				})
			}
		}
	}
	return result, nil
}

// genPackageAlias returns the import alias used for the "generated" package
// in a file by inspecting its import block. If no generated import is found,
// "gen" is returned as the conventional default so callers can still match.
func genPackageAlias(f *ast.File) string {
	for _, imp := range f.Imports {
		path := strings.Trim(imp.Path.Value, `"`)
		if !strings.HasSuffix(path, "/generated") {
			continue
		}
		if imp.Name != nil {
			return imp.Name.Name
		}
		// No explicit alias: use the last path segment.
		parts := strings.Split(path, "/")
		return parts[len(parts)-1]
	}
	return "gen"
}

// handledTypes returns the names of all PropertyType structs that appear in
// the function's parameter list or receiver, resolving package aliases.
func handledTypes(fd *ast.FuncDecl, alias string, knownTypes map[string][]FieldInfo) []string {
	var types []string
	seen := make(map[string]bool)

	check := func(expr ast.Expr) {
		name := baseTypeName(expr, alias)
		if name == "" || seen[name] {
			return
		}
		if _, ok := knownTypes[name]; ok {
			seen[name] = true
			types = append(types, name)
		}
	}

	if fd.Recv != nil {
		for _, f := range fd.Recv.List {
			check(f.Type)
		}
	}
	if fd.Type.Params != nil {
		for _, f := range fd.Type.Params.List {
			check(f.Type)
		}
	}
	return types
}

// baseTypeName extracts the struct name from a type expression by stripping
// pointer (*), slice ([]), and a leading "alias." qualifier.
// Returns "" for types that don't match these patterns.
func baseTypeName(expr ast.Expr, alias string) string {
	switch t := expr.(type) {
	case *ast.StarExpr:
		return baseTypeName(t.X, alias)
	case *ast.ArrayType:
		return baseTypeName(t.Elt, alias)
	case *ast.SelectorExpr:
		if id, ok := t.X.(*ast.Ident); ok && id.Name == alias {
			return t.Sel.Name
		}
	case *ast.Ident:
		// Unqualified name: matches if alias is empty or the type is local.
		return t.Name
	}
	return ""
}

// selectorFields returns the set of all field names accessed via selector
// expressions (x.FieldName) anywhere in a function body, including blank
// assignments (_ = x.FieldName) used to suppress coverage findings.
func selectorFields(body *ast.BlockStmt) map[string]bool {
	fields := make(map[string]bool)
	ast.Inspect(body, func(n ast.Node) bool {
		sel, ok := n.(*ast.SelectorExpr)
		if ok {
			fields[sel.Sel.Name] = true
		}
		return true
	})
	return fields
}

// qualifiedFuncName returns "ReceiverType.MethodName" for methods and
// "FuncName" for plain functions.
func qualifiedFuncName(fd *ast.FuncDecl) string {
	if fd.Recv != nil && len(fd.Recv.List) > 0 {
		t := fd.Recv.List[0].Type
		if star, ok := t.(*ast.StarExpr); ok {
			if id, ok := star.X.(*ast.Ident); ok {
				return id.Name + "." + fd.Name.Name
			}
		}
		if id, ok := t.(*ast.Ident); ok {
			return id.Name + "." + fd.Name.Name
		}
	}
	return fd.Name.Name
}
