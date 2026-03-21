package gml

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type xmlLineString struct {
	SrsName      string          `xml:"srsName,attr,omitempty"`
	SrsDimension int             `xml:"srsDimension,attr,omitempty"`
	PosList      *xmlPosList     `xml:"posList"`
	Pos          []xmlDirectPos  `xml:"pos"`
	Coordinates  *xmlCoordinates `xml:"coordinates"`
}

func decodeLineStringElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x xmlLineString
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: LineString: %w", err)
	}
	ls, err := lineStringFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: ls, EPSG: EPSGFromSRSName(x.SrsName)}, nil
}

func lineStringFromXML(x *xmlLineString) (LineString, error) {
	dim := x.SrsDimension
	if x.PosList != nil {
		return LineStringFromPosListString(x.PosList.Value, preferDim(dim, x.PosList.SrsDimension))
	}
	if len(x.Pos) > 0 {
		var flat []float64
		for _, p := range x.Pos {
			vals, err := ParsePosList(p.Value)
			if err != nil {
				return nil, err
			}
			flat = append(flat, vals...)
		}
		d := preferDim(dim, x.Pos[0].SrsDimension)
		if d == 0 {
			d = len(strings.Fields(x.Pos[0].Value))
			if d < 2 {
				d = 2
			}
		}
		return LineStringFromFlat(flat, d)
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, x.Coordinates.Cs, x.Coordinates.Ts)
		if err != nil {
			return nil, err
		}
		return LineStringFromFlat(coords, 2)
	}
	return nil, fmt.Errorf("gml: LineString has no coordinate data")
}
