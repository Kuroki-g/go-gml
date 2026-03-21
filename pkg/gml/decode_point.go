package gml

import (
	"encoding/xml"
	"fmt"
)

type xmlPoint struct {
	SrsName      string          `xml:"srsName,attr,omitempty"`
	SrsDimension int             `xml:"srsDimension,attr,omitempty"`
	Pos          *xmlDirectPos   `xml:"pos"`
	Coordinates  *xmlCoordinates `xml:"coordinates"`
}

func decodePointElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x xmlPoint
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Point: %w", err)
	}
	pt, err := pointFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: pt, EPSG: EPSGFromSRSName(x.SrsName)}, nil
}

func pointFromXML(x *xmlPoint) (Point, error) {
	if x.Pos != nil {
		return PointFromPosString(x.Pos.Value, preferDim(x.SrsDimension, x.Pos.SrsDimension))
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, x.Coordinates.Cs, x.Coordinates.Ts)
		if err != nil {
			return Point{}, err
		}
		return PointFromFlat(coords, 2)
	}
	return Point{}, fmt.Errorf("gml: Point has no coordinate data")
}

// preferDim returns a if non-zero, otherwise b.
func preferDim(a, b int) int {
	if a > 0 {
		return a
	}
	return b
}
