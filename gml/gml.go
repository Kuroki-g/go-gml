// Package gml is the all-in-one entry point for the go-gml SDK.
// It re-exports geometry types from core and provides streaming readers for
// GML 3.2.1, GML 3.1.1, and GML 2.1.2.
package gml

import (
	"io"

	"github.com/Kuroki-g/go-gml/core"
	"github.com/Kuroki-g/go-gml/gml2_1_2"
	"github.com/Kuroki-g/go-gml/gml3_1_1"
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

// Reader321 is a GML 3.2.1 streaming geometry reader.
type Reader321 = gml3_2_1.Reader

// Reader311 is a GML 3.1.1 streaming geometry reader.
type Reader311 = gml3_1_1.Reader

// Reader212 is a GML 2.1.2 streaming geometry reader.
type Reader212 = gml2_1_2.Reader

// NewReader321 creates a GML 3.2.1 Reader that streams geometry elements from r.
// For non-UTF-8 encoded files (e.g. Shift-JIS), call SetCharsetReader after creation.
func NewReader321(r io.ReadSeeker) *Reader321 {
	return gml3_2_1.NewReader(r)
}

// NewReader311 creates a GML 3.1.1 Reader that streams geometry elements from r.
// For non-UTF-8 encoded files (e.g. Shift-JIS), call SetCharsetReader after creation.
func NewReader311(r io.ReadSeeker) *Reader311 {
	return gml3_1_1.NewReader(r)
}

// NewReader212 creates a GML 2.1.2 Reader that streams geometry elements from r.
// For non-UTF-8 encoded files (e.g. Shift-JIS), call SetCharsetReader after creation.
func NewReader212(r io.Reader) *Reader212 {
	return gml2_1_2.NewReader(r)
}
