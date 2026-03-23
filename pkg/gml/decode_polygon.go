package gml

import (
	"encoding/xml"
	"fmt"

	v3_1_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_1_1"
	v3_2_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

// handlePolygon decodes a gml:Polygon, caches it by gml:id for xlink:href resolution, and returns it.
func (r *Reader) handlePolygon(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	id := extractGMLID(se)
	g, err := decodePolygonElement(dec, se)
	if err != nil {
		return Geometry{}, err
	}
	if id != "" {
		if poly, ok := g.Value.(Polygon); ok {
			r.resolver.polygonByID[id] = poly
		}
	}
	return g, err
}

func decodePolygonElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	if se.Name.Space == gmlNS2 {
		var x v3_1_1.PolygonType
		if err := dec.DecodeElement(&x, &se); err != nil {
			return Geometry{}, fmt.Errorf("gml: Polygon(ns2): %w", err)
		}
		poly, err := polygonFromV31(&x)
		if err != nil {
			return Geometry{}, err
		}
		return Geometry{Value: poly, SRSName: x.SrsName}, nil
	}
	var x v3_2_1.PolygonType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Polygon: %w", err)
	}
	poly, err := polygonFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: poly, SRSName: x.SrsName}, nil
}

func ringFromLinearRing(lr *v3_2_1.LinearRingType, inheritDim int) (Ring, error) {
	if lr == nil {
		return nil, fmt.Errorf("gml: nil LinearRing")
	}
	dim := preferDim(inheritDim, derefDim(lr.SrsDimension))
	if lr.PosList != nil {
		return RingFromPosListString(lr.PosList.Value, preferDim(dim, derefDim(lr.PosList.SrsDimension)))
	}
	if lr.Coordinates != nil {
		return RingFromCoordinatesString(lr.Coordinates.Value, derefStrOr(lr.Coordinates.Cs, ","), derefStrOr(lr.Coordinates.Ts, " "))
	}
	return nil, fmt.Errorf("gml: LinearRing has no coordinate data")
}

func polygonFromXML(x *v3_2_1.PolygonType) (Polygon, error) {
	dim := derefDim(x.SrsDimension)
	var rings []Ring
	if x.Exterior != nil && x.Exterior.LinearRing != nil {
		r, err := ringFromLinearRing(x.Exterior.LinearRing, dim)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon exterior: %w", err)
		}
		rings = append(rings, r)
	}
	for i, ir := range x.Interior {
		if ir.LinearRing == nil {
			continue
		}
		r, err := ringFromLinearRing(ir.LinearRing, dim)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon interior[%d]: %w", i, err)
		}
		rings = append(rings, r)
	}
	return Polygon(rings), nil
}

// polygonFromV31 decodes a Polygon from a v3_1_1.PolygonType, covering GML 3.1.1
// primary fields (exterior/interior) and GML 2.x deprecated fields
// (outerBoundaryIs/innerBoundaryIs).
//
// Priority:
//  1. gml:exterior / gml:interior        — GML 3.1.1 primary
//  2. gml:outerBoundaryIs / gml:innerBoundaryIs — GML 2.x deprecated
func polygonFromV31(x *v3_1_1.PolygonType) (Polygon, error) {
	dim := derefDim(x.SrsDimension)
	var rings []Ring

	extRing := x.Exterior
	if extRing == nil {
		extRing = x.OuterBoundaryIs
	}
	if extRing != nil && extRing.LinearRing != nil {
		r, err := ringFromV31LinearRing(extRing.LinearRing, dim)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon(ns2) exterior: %w", err)
		}
		rings = append(rings, r)
	}

	intRings := x.Interior
	if len(intRings) == 0 {
		intRings = x.InnerBoundaryIs
	}
	for i, ir := range intRings {
		if ir.LinearRing == nil {
			continue
		}
		r, err := ringFromV31LinearRing(ir.LinearRing, dim)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon(ns2) interior[%d]: %w", i, err)
		}
		rings = append(rings, r)
	}
	return Polygon(rings), nil
}

// ringFromV31LinearRing builds a Ring from a v3_1_1.LinearRingType.
//
// Priority:
//  1. gml:posList     — GML 3.1.1 primary
//  2. gml:pos (list)  — GML 3.1.1 alternative
//  3. gml:coordinates — GML 2.x / GML 3.1.1 deprecated
//  4. gml:coord (list) — GML 2.x deprecated
func ringFromV31LinearRing(lr *v3_1_1.LinearRingType, inheritDim int) (Ring, error) {
	if lr == nil {
		return nil, fmt.Errorf("gml: nil LinearRing(ns2)")
	}
	dim := preferDim(inheritDim, derefDim(lr.SrsDimension))
	if lr.PosList != nil {
		return RingFromPosListString(lr.PosList.Value, preferDim(dim, derefDim(lr.PosList.SrsDimension)))
	}
	if len(lr.Pos) > 0 {
		flat, d, err := flatFromPosSlice(lr.Pos, dim)
		if err != nil {
			return nil, err
		}
		return RingFromFlat(flat, d)
	}
	if lr.Coordinates != nil {
		return RingFromCoordinatesString(lr.Coordinates.Value, derefStrOr(lr.Coordinates.Cs, ","), derefStrOr(lr.Coordinates.Ts, " "))
	}
	if len(lr.Coord) > 0 {
		flat := coordSliceToFlat(lr.Coord)
		return RingFromFlat(flat, 2)
	}
	return nil, fmt.Errorf("gml: LinearRing(ns2) has no coordinate data")
}
