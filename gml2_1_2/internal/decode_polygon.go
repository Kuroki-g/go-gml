package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
)

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

func ringFromLinearRing(lr *gen.LinearRingType) (core.Ring, error) {
	if lr == nil {
		return nil, fmt.Errorf("gml: nil LinearRing")
	}
	if lr.Coordinates != nil {
		return RingFromCoordinatesString(lr.Coordinates.Value, derefStrOr(lr.Coordinates.Cs, ","), derefStrOr(lr.Coordinates.Ts, " "))
	}
	if len(lr.Coord) > 0 {
		return RingFromFlat(coordSliceToFlat(lr.Coord), 2)
	}
	return nil, fmt.Errorf("gml: LinearRing has no coordinate data")
}

func polygonFromXML(x *gen.PolygonType) (core.Polygon, error) {
	var rings []core.Ring
	if x.OuterBoundaryIs != nil && x.OuterBoundaryIs.LinearRing != nil {
		r, err := ringFromLinearRing(x.OuterBoundaryIs.LinearRing)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon outerBoundaryIs: %w", err)
		}
		rings = append(rings, r)
	}
	for i, ir := range x.InnerBoundaryIs {
		if ir.LinearRing == nil {
			continue
		}
		r, err := ringFromLinearRing(ir.LinearRing)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon innerBoundaryIs[%d]: %w", i, err)
		}
		rings = append(rings, r)
	}
	return core.Polygon(rings), nil
}
