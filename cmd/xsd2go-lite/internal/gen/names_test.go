package gen

import "testing"

// ---- deriveAlias ----

func TestDeriveAlias(t *testing.T) {
	tests := []struct {
		pkgPath string
		want    string
	}{
		// last segment is "generated" → use preceding segment + "gen", keep underscores
		{"github.com/Kuroki-g/go-gml/gml3_1_1/generated", "gml3_1_1gen"},
		{"github.com/Kuroki-g/go-gml/gml3_2_1/generated", "gml3_2_1gen"},
		{"github.com/Kuroki-g/go-gml/gml2_1_2/generated", "gml2_1_2gen"},
		// last segment is not "generated" → use last segment
		{"github.com/foo/mypkg", "mypkg"},
		// hyphens and dots stripped
		{"github.com/foo/my-pkg", "mypkg"},
	}
	for _, tt := range tests {
		got := deriveAlias(tt.pkgPath)
		if got != tt.want {
			t.Errorf("deriveAlias(%q) = %q, want %q", tt.pkgPath, got, tt.want)
		}
	}
}
