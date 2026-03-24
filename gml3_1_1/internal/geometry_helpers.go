package internal

import (
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
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

// PointFromPosString parses a gml:pos chardata string and returns a Point.
func PointFromPosString(s string, dim int) (core.Point, error) {
	coords, err := core.ParsePosList(s)
	if err != nil {
		return core.Point{}, err
	}
	if dim <= 0 {
		dim = len(coords)
	}
	return PointFromFlat(coords, dim)
}

// LineStringFromPosListString parses a gml:posList chardata string.
func LineStringFromPosListString(s string, dim int) (core.LineString, error) {
	coords, err := core.ParsePosList(s)
	if err != nil {
		return nil, err
	}
	return LineStringFromFlat(coords, effectiveDim(dim))
}

// RingFromPosListString parses a gml:posList chardata string into a Ring.
func RingFromPosListString(s string, dim int) (core.Ring, error) {
	coords, err := core.ParsePosList(s)
	if err != nil {
		return nil, err
	}
	return RingFromFlat(coords, effectiveDim(dim))
}

// RingFromCoordinatesString parses a deprecated gml:coordinates string into a Ring.
func RingFromCoordinatesString(s, cs, ts string) (core.Ring, error) {
	coords, err := core.ParseCoordinates(s, cs, ts)
	if err != nil {
		return nil, err
	}
	return RingFromFlat(coords, 2)
}

// effectiveDim returns dim if > 0, otherwise defaults to 2.
func effectiveDim(dim int) int {
	if dim > 0 {
		return dim
	}
	return 2
}
