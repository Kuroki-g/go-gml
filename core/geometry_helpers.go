package core

import "fmt"

// PointFromFlat builds a Point from a flat coord slice.
func PointFromFlat(coords []float64, dim uint) (Point, error) {
	if dim == 0 {
		dim = 2
	}
	if uint(len(coords)) < dim {
		return Point{}, fmt.Errorf("gml: need at least %d values for a point, got %d", dim, len(coords))
	}
	return append(Point(nil), coords[:dim]...), nil
}

// LineStringFromFlat builds a LineString from a flat coord slice.
func LineStringFromFlat(coords []float64, dim uint) (LineString, error) {
	pts, err := ToPoints(coords, dim)
	if err != nil {
		return nil, err
	}
	return LineString(pts), nil
}

// RingFromFlat builds a Ring from a flat coord slice.
func RingFromFlat(coords []float64, dim uint) (Ring, error) {
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
// dim is the srsDimension hint; nil means infer from the number of values in s.
func PointFromPosString(s string, dim *uint) (Point, error) {
	coords, err := ParsePosList(s)
	if err != nil {
		return Point{}, err
	}
	var d uint
	if dim != nil {
		d = *dim
	} else {
		d = uint(len(coords))
	}
	return PointFromFlat(coords, d)
}

// LineStringFromPosListString parses a gml:posList chardata string.
func LineStringFromPosListString(s string, dim *uint) (LineString, error) {
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
func RingFromPosListString(s string, dim *uint) (Ring, error) {
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

// effectiveDim resolves the coordinate dimension for a posList.
// srsDimension is positiveInteger and optional per GML XSD; any dim>0 is valid.
// If dim is nil (srsDimension omitted), the dimension is inferred from nValues:
// an odd value count cannot be 2D, so dim=3 is assumed; otherwise 2D is assumed.
func effectiveDim(dim *uint, nValues int) (uint, error) {
	if dim != nil {
		return *dim, nil
	}
	if nValues%2 != 0 {
		return 3, nil
	}
	return 2, nil
}
