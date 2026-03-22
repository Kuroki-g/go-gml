package gml

import (
	"encoding/xml"
	"fmt"

	v3 "github.com/Kuroki-g/go-gml/pkg/gml/v3"
)

func decodePolygonElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v3.PolygonType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Polygon: %w", err)
	}
	poly, err := polygonFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: poly, SRSName: x.SrsName}, nil
}

func ringFromLinearRing(lr *v3.LinearRingType, inheritDim int) (Ring, error) {
	if lr == nil {
		return nil, fmt.Errorf("gml: nil LinearRing")
	}
	dim := preferDim(inheritDim, lr.SrsDimension)
	if lr.PosList != nil {
		return RingFromPosListString(lr.PosList.Value, preferDim(dim, lr.PosList.SrsDimension))
	}
	if lr.Coordinates != nil {
		return RingFromCoordinatesString(lr.Coordinates.Value, lr.Coordinates.Cs, lr.Coordinates.Ts)
	}
	return nil, fmt.Errorf("gml: LinearRing has no coordinate data")
}

func polygonFromXML(x *v3.PolygonType) (Polygon, error) {
	dim := x.SrsDimension
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
