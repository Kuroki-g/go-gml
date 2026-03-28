package core

// Point is a coordinate tuple whose length equals the srsDimension of the source
// GML element: [x, y] for 2D or [x, y, z] for 3D.
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

// Solid is a 3D volumetric geometry bounded by surfaces.
// Exterior is the outer shell; Interior holds inner shells (voids).
// Each shell is a MultiPolygon because a shell may consist of multiple
// surface members (e.g. walls, roof, floor of a LoD1 building).
type Solid struct {
	Exterior MultiPolygon
	Interior []MultiPolygon
}

// GridCoverage holds the decoded content of a gml:Grid/gml:RectifiedGrid coverage.
// Low and High are the integer grid index bounds (one value per axis).
// Tuples is the raw CSV content of gml:DataBlock/gml:tupleList.
type GridCoverage struct {
	Low    []int
	High   []int
	Tuples string
}
