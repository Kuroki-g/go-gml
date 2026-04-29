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
	g, err := decodePolygonElement(dec, se, r.globalDim, r.globalSrsName)
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

func decodePolygonElement(dec *xml.Decoder, se xml.StartElement, fallbackDim *uint, fallbackSrsName *string) (core.Geometry, error) {
	var x gen.PolygonType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Polygon: %w", err)
	}
	poly, err := polygonFromXML(&x, fallbackDim, fallbackSrsName)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: poly, SRSName: x.SRSReferenceGroup.SrsName}, nil
}

// ringFromLinearRing builds a Ring from a decoded LinearRingType.
// inheritDim carries srsDimension from the enclosing Polygon (which in turn
// inherits from the document-root gml:Envelope via Reader.globalDim).
func ringFromLinearRing(lr *gen.LinearRingType, inheritDim *uint, inheritSrsName *string) (core.Ring, error) {
	if lr == nil {
		return nil, fmt.Errorf("gml: nil LinearRing")
	}
	dim := preferDim(lr.SrsDimension, inheritDim)
	srsName := preferSrsName(lr.SRSReferenceGroup.SrsName, inheritSrsName)
	if lr.PosList != nil {
		return core.RingFromPosListString(lr.PosList.Value, preferDim(lr.PosList.SrsDimension, dim), preferSrsName(lr.PosList.SrsName, srsName))
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
		d := uint(len(strings.Fields(lr.Pos[0].Value)))
		if d == 0 {
			return nil, fmt.Errorf("gml: pos has no values")
		}
		return core.RingFromFlat(flat, d)
	}
	if lr.Coordinates != nil {
		return core.RingFromCoordinatesString(lr.Coordinates.Value, derefStrOr(lr.Coordinates.Cs, ","), derefStrOr(lr.Coordinates.Ts, " "), derefStrOr(lr.Coordinates.Decimal, "."))
	}
	return nil, fmt.Errorf("gml: LinearRing has no coordinate data")
}

func polygonFromXML(x *gen.PolygonType, fallbackDim *uint, fallbackSrsName *string) (core.Polygon, error) {
	dim := preferDim(x.SrsDimension, fallbackDim)
	srsName := preferSrsName(x.SRSReferenceGroup.SrsName, fallbackSrsName)
	var rings []core.Ring
	if x.Exterior != nil && x.Exterior.LinearRing != nil {
		r, err := ringFromLinearRing(x.Exterior.LinearRing, dim, srsName)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon exterior: %w", err)
		}
		rings = append(rings, r)
	}
	for i, ir := range x.Interior {
		if ir.LinearRing == nil {
			continue
		}
		r, err := ringFromLinearRing(ir.LinearRing, dim, srsName)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon interior[%d]: %w", i, err)
		}
		rings = append(rings, r)
	}
	return core.Polygon(rings), nil
}
