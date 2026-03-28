package core

import "fmt"

// PointFromFlat builds a Point from a flat coord slice.
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

// RingFromCoordinatesString parses a deprecated gml:coordinates string into a Ring.
func RingFromCoordinatesString(s, cs, ts string) (Ring, error) {
	coords, err := ParseCoordinates(s, cs, ts)
	if err != nil {
		return nil, err
	}
	return RingFromFlat(coords, 2)
}

// PointFromPosString parses a gml:pos chardata string and returns a Point.
func PointFromPosString(s string, dim int) (Point, error) {
	coords, err := ParsePosList(s)
	if err != nil {
		return Point{}, err
	}
	if dim <= 0 {
		dim = len(coords)
	}
	return PointFromFlat(coords, dim)
}

// LineStringFromPosListString parses a gml:posList chardata string.
func LineStringFromPosListString(s string, dim int) (LineString, error) {
	coords, err := ParsePosList(s)
	if err != nil {
		return nil, err
	}
	d, err := effectiveDim(dim, len(coords))
	if err != nil {
		return nil, err
	}
	return LineStringFromFlat(coords, d)
}

// RingFromPosListString parses a gml:posList chardata string into a Ring.
func RingFromPosListString(s string, dim int) (Ring, error) {
	coords, err := ParsePosList(s)
	if err != nil {
		return nil, err
	}
	d, err := effectiveDim(dim, len(coords))
	if err != nil {
		return nil, err
	}
	return RingFromFlat(coords, d)
}

// effectiveDim resolves the coordinate dimension.
// If dim is explicitly provided (2 or 3), it is returned as-is.
// If dim is 0 (srsDimension omitted), the dimension is inferred from nValues:
// an odd value count cannot be 2D, so dim=3 is assumed.
// When nValues is divisible by both 2 and 3 (e.g. 6, 12, 18), 2D is assumed
// because the true dimension cannot be determined without the parent CRS.
// Dimensions other than 0, 2, or 3 are not supported and return an error.
func effectiveDim(dim, nValues int) (int, error) {
	switch dim {
	case 0:
		if nValues%2 != 0 {
			return 3, nil
		}
		return 2, nil
	case 2, 3:
		return dim, nil
	default:
		return 0, fmt.Errorf("gml: unsupported srsDimension %d", dim)
	}
}
