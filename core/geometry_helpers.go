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
	if dim == nil {
		d = uint(len(coords))
	} else {
		d = *dim
	}
	return PointFromFlat(coords, d)
}

// LineStringFromPosListString parses a gml:posList chardata string.
// dim is the resolved srsDimension (from the element itself or inherited from parent context).
// If dim is nil, the dimension is derived from the CRS definition via srsName.
// Returns an error if neither dim nor srsName can resolve the dimension.
func LineStringFromPosListString(s string, dim *uint, srsName *string) (LineString, error) {
	if dim == nil {
		dim = DimFromSRSName(srsName)
	}
	if dim == nil {
		return nil, fmt.Errorf("gml: cannot parse gml:posList: srsDimension absent and CRS dimension unknown")
	}
	coords, err := ParsePosList(s)
	if err != nil {
		return nil, err
	}
	return LineStringFromFlat(coords, *dim)
}

// RingFromPosListString parses a gml:posList chardata string into a Ring.
// dim is the resolved srsDimension (from the element itself or inherited from parent context).
// If dim is nil, the dimension is derived from the CRS definition via srsName.
// Returns an error if neither dim nor srsName can resolve the dimension.
func RingFromPosListString(s string, dim *uint, srsName *string) (Ring, error) {
	if dim == nil {
		dim = DimFromSRSName(srsName)
	}
	if dim == nil {
		return nil, fmt.Errorf("gml: cannot parse gml:posList: srsDimension absent and CRS dimension unknown")
	}
	coords, err := ParsePosList(s)
	if err != nil {
		return nil, err
	}
	return RingFromFlat(coords, *dim)
}
