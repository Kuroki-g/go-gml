package gml

import "fmt"

// Point is a 2D coordinate pair [x, y] (longitude/easting, latitude/northing).
// Memory layout is identical to [2]float64, so users can cast to
// github.com/paulmach/orb.Point without copying.
type Point [2]float64

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

// PointFromFlat builds a Point from a flat coord slice of length >= 2.
// If dim > 2, extra ordinates (Z, M) are ignored.
func PointFromFlat(coords []float64, dim int) (Point, error) {
	if dim <= 0 {
		dim = 2
	}
	if len(coords) < dim {
		return Point{}, fmt.Errorf("gml: need at least %d values for a point, got %d", dim, len(coords))
	}
	return Point{coords[0], coords[1]}, nil
}

// LineStringFromFlat builds a LineString from a flat coord slice.
func LineStringFromFlat(coords []float64, dim int) (LineString, error) {
	pts, err := toPoints(coords, dim)
	if err != nil {
		return nil, err
	}
	return LineString(pts), nil
}

// RingFromFlat builds a Ring from a flat coord slice.
func RingFromFlat(coords []float64, dim int) (Ring, error) {
	pts, err := toPoints(coords, dim)
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

// -- internal --

// toPoints converts []float64 → []Point, reusing ToPoints from coords.go.
func toPoints(coords []float64, dim int) ([]Point, error) {
	raw, err := ToPoints(coords, dim)
	if err != nil {
		return nil, err
	}
	pts := make([]Point, len(raw))
	for i, p := range raw {
		pts[i] = Point(p)
	}
	return pts, nil
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
