// Package gml is the all-in-one entry point for the go-gml SDK.
// It re-exports geometry types from core and provides a GML 3.2.1 streaming reader.
package gml

import (
	"io"

	"github.com/Kuroki-g/go-gml/core"
	"github.com/Kuroki-g/go-gml/gml3_2_1"
)

// Geometry types (re-exported from core).
type (
	Point           = core.Point
	LineString      = core.LineString
	Ring            = core.Ring
	Polygon         = core.Polygon
	MultiPoint      = core.MultiPoint
	MultiLineString = core.MultiLineString
	MultiPolygon    = core.MultiPolygon
	Bound           = core.Bound
	GridCoverage    = core.GridCoverage
	Geometry        = core.Geometry
)

// Reader is a GML 3.2.1 streaming geometry reader.
type Reader = gml3_2_1.Reader

// NewReader creates a Reader that streams geometry elements from r.
// For non-UTF-8 encoded files (e.g. Shift-JIS), call SetCharsetReader after creation.
func NewReader(r io.ReadSeeker) *Reader {
	return gml3_2_1.NewReader(r)
}
