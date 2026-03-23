package gml

import (
	"encoding/xml"
	"fmt"

	v3_1_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_1_1"
	v3_2_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

func decodePointElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	if se.Name.Space == gmlNS2 {
		var x v3_1_1.PointType
		if err := dec.DecodeElement(&x, &se); err != nil {
			return Geometry{}, fmt.Errorf("gml: Point(ns2): %w", err)
		}
		pt, err := pointFromV31(&x)
		if err != nil {
			return Geometry{}, err
		}
		return Geometry{Value: pt, SRSName: x.SrsName}, nil
	}
	var x v3_2_1.PointType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Point: %w", err)
	}
	pt, err := pointFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: pt, SRSName: x.SrsName}, nil
}

func pointFromXML(x *v3_2_1.PointType) (Point, error) {
	if x.Pos != nil {
		return PointFromPosString(x.Pos.Value, preferDim(derefDim(x.SrsDimension), derefDim(x.Pos.SrsDimension)))
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return Point{}, err
		}
		return PointFromFlat(coords, 2)
	}
	return Point{}, fmt.Errorf("gml: Point has no coordinate data")
}

// pointFromV31 decodes a Point from a v3_1_1.PointType, covering GML 3.1.1 primary
// fields (pos) and GML 2.x deprecated fields (coordinates, coord).
//
// Priority:
//  1. gml:pos        — GML 3.1.1 primary
//  2. gml:coordinates — GML 2.x / GML 3.1.1 deprecated
//  3. gml:coord      — GML 2.x deprecated (removed in GML 3.0)
func pointFromV31(x *v3_1_1.PointType) (Point, error) {
	dim := preferDim(derefDim(x.SrsDimension), 0)
	if x.Pos != nil {
		return PointFromPosString(x.Pos.Value, preferDim(dim, derefDim(x.Pos.SrsDimension)))
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return Point{}, err
		}
		return PointFromFlat(coords, 2)
	}
	if x.Coord != nil {
		return Point{x.Coord.X, x.Coord.Y}, nil
	}
	return Point{}, fmt.Errorf("gml: Point(ns2) has no coordinate data")
}

// preferDim returns a if non-zero, otherwise b.
func preferDim(a, b int) int {
	if a > 0 {
		return a
	}
	return b
}

// derefDim dereferences a *int dimension attribute; nil → 0.
func derefDim(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

// derefStrOr dereferences a *string attribute; nil → def (the XSD default value).
func derefStrOr(p *string, def string) string {
	if p == nil {
		return def
	}
	return *p
}
