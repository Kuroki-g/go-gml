package core

// Geometry wraps a parsed GML geometry element with CRS metadata.
// Value is one of: Point, LineString, Polygon, MultiPoint, MultiLineString, MultiPolygon, Bound, GridCoverage.
type Geometry struct {
	Value   interface{}
	SRSName *string
}

// Reader scans a GML document for geometry elements.
type Reader interface {
	Next() (Geometry, error)
}
