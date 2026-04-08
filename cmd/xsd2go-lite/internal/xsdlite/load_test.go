package xsdlite

import (
	"os"
	"path/filepath"
	"testing"
)

// ---- rawIncludesImports ----

func TestRawIncludesImportsFromFile(t *testing.T) {
	xsd := `<?xml version="1.0" encoding="UTF-8"?>
<schema targetNamespace="http://example.com"
        xmlns="http://www.w3.org/2001/XMLSchema">
  <include schemaLocation="other.xsd"/>
  <import namespace="http://www.w3.org/1999/xlink" schemaLocation="xlink.xsd"/>
  <import namespace="http://external.example.com"/>
</schema>`

	dir := t.TempDir()
	path := filepath.Join(dir, "test.xsd")
	if err := os.WriteFile(path, []byte(xsd), 0644); err != nil {
		t.Fatal(err)
	}

	res, err := rawIncludesImportsFromFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.includes) != 1 || res.includes[0] != "other.xsd" {
		t.Errorf("includes = %v, want [other.xsd]", res.includes)
	}
	if len(res.imports) != 2 {
		t.Errorf("imports = %d, want 2", len(res.imports))
	}
	if res.imports[0].loc != "xlink.xsd" {
		t.Errorf("import[0].loc = %q, want xlink.xsd", res.imports[0].loc)
	}
	if res.imports[1].loc != "" {
		t.Errorf("import[1].loc should be empty (no schemaLocation)")
	}
}
