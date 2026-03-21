package gml

import (
	"encoding/xml"
	"fmt"
)

type xmlLinearRing struct {
	SrsDimension int             `xml:"srsDimension,attr,omitempty"`
	PosList      *xmlPosList     `xml:"posList"`
	Coordinates  *xmlCoordinates `xml:"coordinates"`
}

// xmlRingProp wraps exterior/interior that contain a LinearRing.
// Note: gml:Ring (xlink-based topology) is not handled here.
type xmlRingProp struct {
	LinearRing *xmlLinearRing `xml:"LinearRing"`
}

type xmlPolygon struct {
	SrsName      string        `xml:"srsName,attr,omitempty"`
	SrsDimension int           `xml:"srsDimension,attr,omitempty"`
	Exterior     *xmlRingProp  `xml:"exterior"`
	Interior     []xmlRingProp `xml:"interior"`
}

func decodePolygonElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x xmlPolygon
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Polygon: %w", err)
	}
	poly, err := polygonFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: poly, EPSG: EPSGFromSRSName(x.SrsName)}, nil
}

func ringFromXML(ring *xmlLinearRing, inheritDim int) (Ring, error) {
	if ring == nil {
		return nil, fmt.Errorf("gml: nil LinearRing")
	}
	dim := preferDim(inheritDim, ring.SrsDimension)
	if ring.PosList != nil {
		return RingFromPosListString(ring.PosList.Value, preferDim(dim, ring.PosList.SrsDimension))
	}
	if ring.Coordinates != nil {
		return RingFromCoordinatesString(ring.Coordinates.Value, ring.Coordinates.Cs, ring.Coordinates.Ts)
	}
	return nil, fmt.Errorf("gml: LinearRing has no coordinate data")
}

func polygonFromXML(x *xmlPolygon) (Polygon, error) {
	dim := x.SrsDimension
	var rings []Ring
	if x.Exterior != nil && x.Exterior.LinearRing != nil {
		r, err := ringFromXML(x.Exterior.LinearRing, dim)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon exterior: %w", err)
		}
		rings = append(rings, r)
	}
	for i, ir := range x.Interior {
		if ir.LinearRing == nil {
			continue
		}
		r, err := ringFromXML(ir.LinearRing, dim)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon interior[%d]: %w", i, err)
		}
		rings = append(rings, r)
	}
	return Polygon(rings), nil
}
