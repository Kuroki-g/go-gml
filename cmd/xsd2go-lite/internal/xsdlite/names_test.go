package xsdlite

import "testing"

// ---- goName ----

func TestGoName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", "_"},
		{"pointType", "PointType"},
		{"point_type", "PointType"},
		{"point-type", "PointType"},
		{"PolygonType", "PolygonType"},
		{"multiSurface", "MultiSurface"},
		{"multi_surface_type", "MultiSurfaceType"},
		// digit prefix
		{"3DGeometry", "_3DGeometry"},
		// reserved word
		{"type", "TypeField"},
		{"map", "MapField"},
		{"range", "RangeField"},
		// dot separator
		{"gml.geometry", "GmlGeometry"},
		// already exported
		{"AbstractGeometryType", "AbstractGeometryType"},
		// single char
		{"x", "X"},
		// all underscores → empty parts → "_"
		{"_", "_"},
		// GML 3.1.1 abstract element names: "_Xxx" → "AbstractXxx"
		{"_Curve", "AbstractCurve"},
		{"_Surface", "AbstractSurface"},
		{"_Ring", "AbstractRing"},
		{"_CurveSegment", "AbstractCurveSegment"},
		{"_Geometry", "AbstractGeometry"},
	}

	for _, tt := range tests {
		got := goName(tt.in)
		if got != tt.want {
			t.Errorf("goName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

// ---- avoidReserved ----

func TestAvoidReserved(t *testing.T) {
	reserved := []string{
		"Break", "Case", "Chan", "Const", "Continue", "Default", "Defer",
		"Else", "Fallthrough", "For", "Func", "Go", "Goto", "If", "Import",
		"Interface", "Map", "Package", "Range", "Return", "Select", "Struct",
		"Switch", "Type", "Var",
	}
	for _, r := range reserved {
		got := avoidReserved(r)
		if got != r+"Field" {
			t.Errorf("avoidReserved(%q) = %q, want %q", r, got, r+"Field")
		}
	}
	// Non-reserved should pass through unchanged.
	for _, ok := range []string{"Point", "Line", "Geometry", "Hello"} {
		if got := avoidReserved(ok); got != ok {
			t.Errorf("avoidReserved(%q) = %q, want %q (unchanged)", ok, got, ok)
		}
	}
}

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
