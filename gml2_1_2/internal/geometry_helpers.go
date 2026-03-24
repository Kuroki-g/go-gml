package internal

import (
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
)

// coordSliceToFlat converts a []CoordType to a flat []float64 (X, Y pairs).
func coordSliceToFlat(coords []gen.CoordType) []float64 {
	flat := make([]float64, 0, len(coords)*2)
	for _, c := range coords {
		flat = append(flat, c.X, c.Y)
	}
	return flat
}

// derefStrOr dereferences a *string attribute; nil → def (the XSD default value).
func derefStrOr(p *string, def string) string {
	if p == nil {
		return def
	}
	return *p
}
