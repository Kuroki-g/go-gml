package internal

import (
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
)

// PointFromFlat builds a Point from a flat coord slice.
func PointFromFlat(coords []float64, dim int) (core.Point, error) {
	if dim <= 0 {
		dim = 2
	}
	if len(coords) < dim {
		return core.Point{}, fmt.Errorf("gml: need at least %d values for a point, got %d", dim, len(coords))
	}
	return append(core.Point(nil), coords[:dim]...), nil
}

// LineStringFromFlat builds a LineString from a flat coord slice.
func LineStringFromFlat(coords []float64, dim int) (core.LineString, error) {
	pts, err := core.ToPoints(coords, dim)
	if err != nil {
		return nil, err
	}
	return core.LineString(pts), nil
}

// RingFromFlat builds a Ring from a flat coord slice.
func RingFromFlat(coords []float64, dim int) (core.Ring, error) {
	pts, err := core.ToPoints(coords, dim)
	if err != nil {
		return nil, err
	}
	return core.Ring(pts), nil
}

// RingFromCoordinatesString parses a deprecated gml:coordinates string into a Ring.
func RingFromCoordinatesString(s, cs, ts string) (core.Ring, error) {
	coords, err := core.ParseCoordinates(s, cs, ts)
	if err != nil {
		return nil, err
	}
	return RingFromFlat(coords, 2)
}

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
