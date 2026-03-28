package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

// handlePolygon decodes a gml:Polygon, caches it by gml:id for xlink:href resolution, and returns it.
func (r *Reader) handlePolygon(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	id := extractGMLID(se)
	g, err := decodePolygonElement(dec, se)
	if err != nil {
		return core.Geometry{}, err
	}
	if id != "" {
		if poly, ok := g.Value.(core.Polygon); ok {
			r.resolver.polygonByID[id] = poly
		}
	}
	return g, err
}

func decodePolygonElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.PolygonType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Polygon: %w", err)
	}
	poly, err := polygonFromXML(&x)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: poly, SRSName: x.SrsName}, nil
}

// ringFromLinearRing builds a Ring from a decoded LinearRingType.
// inheritDim is the srsDimension from the enclosing Polygon element.
// NOTE: PLATEAU CityGML 2.0 files set srsDimension="3" only on the root
// gml:Envelope; individual Polygon/LinearRing/posList elements carry no
// srsDimension. Since this reader receives inheritDim from the Polygon only,
// that top-level value is never propagated here, causing dim=0 to fall back
// to effectiveDim heuristics (odd count → 3D; even count divisible by 6 → 2D
// and potentially wrong). True fix requires threading srsDimension across the
// full element stack, which is an architecture change.
// TODO: inherit srsDimension from document root / parent CRS context.
func ringFromLinearRing(lr *gen.LinearRingType, inheritDim int) (core.Ring, error) {
	if lr == nil {
		return nil, fmt.Errorf("gml: nil LinearRing")
	}
	dim := preferDim(inheritDim, derefDim(lr.SrsDimension))
	if lr.PosList != nil {
		return core.RingFromPosListString(lr.PosList.Value, preferDim(dim, derefDim(lr.PosList.SrsDimension)))
	}
	if len(lr.Pos) > 0 {
		var flat []float64
		for _, p := range lr.Pos {
			vals, err := core.ParsePosList(p.Value)
			if err != nil {
				return nil, err
			}
			flat = append(flat, vals...)
		}
		d := preferDim(dim, derefDim(lr.Pos[0].SrsDimension))
		if d == 0 {
			d = len(strings.Fields(lr.Pos[0].Value))
			if d < 2 {
				d = 2
			}
		}
		return core.RingFromFlat(flat, d)
	}
	if lr.Coordinates != nil {
		return core.RingFromCoordinatesString(lr.Coordinates.Value, derefStrOr(lr.Coordinates.Cs, ","), derefStrOr(lr.Coordinates.Ts, " "))
	}
	return nil, fmt.Errorf("gml: LinearRing has no coordinate data")
}

func polygonFromXML(x *gen.PolygonType) (core.Polygon, error) {
	dim := derefDim(x.SrsDimension)
	var rings []core.Ring
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
	return core.Polygon(rings), nil
}
