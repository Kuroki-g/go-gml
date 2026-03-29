package cmd

import (
	"os"
	"path/filepath"
)

// Finding represents one unhandled field in a specific function.
type Finding struct {
	Module  string
	Struct  string
	Func    string
	Field   string
	XMLName string
	IsAttr  bool
}

// checkModule analyzes one parser module directory for functions that handle
// a PropertyType struct but do not access all of its XML-tagged fields.
//
// Every field of a PropertyType must be either directly accessed or explicitly
// acknowledged in the function body via a blank assignment:
//
//	_ = m.FieldName // reason or TODO
//
// There is no hardcoded skip list in this tool: intentional non-handling must
// be expressed in the implementation code, not here.
func checkModule(dir string) ([]Finding, error) {
	genDir := filepath.Join(dir, "generated")
	intDir := filepath.Join(dir, "internal")

	if _, err := os.Stat(genDir); err != nil {
		return nil, nil
	}
	if _, err := os.Stat(intDir); err != nil {
		return nil, nil
	}

	propTypes, err := parsePropertyTypes(genDir)
	if err != nil {
		return nil, err
	}

	funcAccesses, err := analyzeFunctions(intDir, propTypes)
	if err != nil {
		return nil, err
	}

	var findings []Finding
	for typeName, fields := range propTypes {
		accesses, ok := funcAccesses[typeName]
		if !ok {
			continue // type not used as a parameter in internal/ — out of scope
		}
		for _, fa := range accesses {
			for _, f := range fields {
				if !fa.Fields[f.Name] {
					findings = append(findings, Finding{
						Module:  dir,
						Struct:  typeName,
						Func:    fa.FuncName,
						Field:   f.Name,
						XMLName: f.XMLName,
						IsAttr:  f.IsAttr,
					})
				}
			}
		}
	}
	return findings, nil
}
