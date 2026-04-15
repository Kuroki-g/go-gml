package internal

import (
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
)

// coordSliceToFlat converts a []CoordType to a flat []float64 and returns the stride.
// stride is 3 if any coord has Z != nil, otherwise 2.
func coordSliceToFlat(coords []gen.CoordType) ([]float64, uint) {
	if len(coords) == 0 {
		return nil, 2
	}
	var dim uint = 2
	for _, c := range coords {
		if c.Z != nil {
			dim = 3
			break
		}
	}
	flat := make([]float64, 0, uint(len(coords))*dim)
	for _, c := range coords {
		y := 0.0
		if c.Y != nil {
			y = *c.Y
		}
		if dim == 3 {
			z := 0.0
			if c.Z != nil {
				z = *c.Z
			}
			flat = append(flat, c.X, y, z)
		} else {
			flat = append(flat, c.X, y)
		}
	}
	return flat, dim
}

// derefStrOr dereferences a *string attribute; nil → def (the XSD default value).
func derefStrOr(p *string, def string) string {
	if p == nil {
		return def
	}
	return *p
}
