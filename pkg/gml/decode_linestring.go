package gml

import (
	"encoding/xml"
	"fmt"
	"strings"

	v3 "github.com/Kuroki-g/go-gml/pkg/gml/v3"
)

func decodeLineStringElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	if se.Name.Space == gmlNS2 {
		return decodeLineStringV2(dec, se)
	}
	var x v3.LineStringType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: LineString: %w", err)
	}
	ls, err := lineStringFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: ls, SRSName: x.SrsName}, nil
}

func lineStringFromXML(x *v3.LineStringType) (LineString, error) {
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
