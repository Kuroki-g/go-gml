package gml

import "fmt"

// Point is a coordinate tuple whose length equals the srsDimension of the source
// GML element: [x, y] for 2D or [x, y, z] for 3D. Use len(p) to check dimension.
type Point []float64

// LineString is an ordered sequence of Points.
type LineString []Point

// Ring is a closed sequence of Points.
// The first and last point should be equal per GML convention.
type Ring []Point

// Polygon is a surface defined by Rings.
// rings[0] is the exterior boundary; rings[1:] are interior holes.
type Polygon []Ring

// MultiPoint is a collection of Points.
type MultiPoint []Point

// MultiLineString is a collection of LineStrings.
type MultiLineString []LineString

// MultiPolygon is a collection of Polygons.
type MultiPolygon []Polygon

// Bound is an axis-aligned bounding box.
type Bound struct {
	Min, Max Point
}

// -- helpers: flat []float64 → geometry types --

// PointFromFlat builds a Point from a flat coord slice.
// Takes the first dim values. If dim <= 0 it defaults to 2.
func PointFromFlat(coords []float64, dim int) (Point, error) {
	if dim <= 0 {
		dim = 2
	}
	if len(coords) < dim {
		return Point{}, fmt.Errorf("gml: need at least %d values for a point, got %d", dim, len(coords))
	}
	return append(Point(nil), coords[:dim]...), nil
}

// LineStringFromFlat builds a LineString from a flat coord slice.
func LineStringFromFlat(coords []float64, dim int) (LineString, error) {
	pts, err := ToPoints(coords, dim)
	if err != nil {
		return nil, err
	}
	return LineString(pts), nil
}

// RingFromFlat builds a Ring from a flat coord slice.
func RingFromFlat(coords []float64, dim int) (Ring, error) {
	pts, err := ToPoints(coords, dim)
	if err != nil {
		return nil, err
	}
	return Ring(pts), nil
}

// -- helpers: string → geometry types (wraps ParsePosList + above) --

// PointFromPosString parses a gml:pos chardata string and returns a Point.
// e.g. "139.7 35.6" or "139.7 35.6 10.5"
func PointFromPosString(s string, dim int) (Point, error) {
	coords, err := ParsePosList(s)
	if err != nil {
		return Point{}, err
	}
	return PointFromFlat(coords, effectiveDim(dim, len(coords)))
}

// LineStringFromPosListString parses a gml:posList chardata string.
func LineStringFromPosListString(s string, dim int) (LineString, error) {
	coords, err := ParsePosList(s)
	if err != nil {
		return nil, err
	}
	return LineStringFromFlat(coords, effectiveDim(dim, len(coords)))
}

// RingFromPosListString parses a gml:posList chardata string into a Ring.
func RingFromPosListString(s string, dim int) (Ring, error) {
	coords, err := ParsePosList(s)
	if err != nil {
		return nil, err
	}
	return RingFromFlat(coords, effectiveDim(dim, len(coords)))
}

// RingFromCoordinatesString parses a deprecated gml:coordinates string into a Ring.
func RingFromCoordinatesString(s, cs, ts string) (Ring, error) {
	coords, err := ParseCoordinates(s, cs, ts)
	if err != nil {
		return nil, err
	}
	return RingFromFlat(coords, 2)
}

// effectiveDim returns dim if > 0, otherwise infers from total coord count.
// Falls back to 2 if inference is ambiguous.
func effectiveDim(dim, total int) int {
	if dim > 0 {
		return dim
	}
	if total%3 == 0 && total > 0 {
		return 3
	}
	return 2
}
