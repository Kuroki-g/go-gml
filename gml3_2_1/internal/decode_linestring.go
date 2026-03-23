package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

func decodeLineStringElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.LineStringType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: LineString: %w", err)
	}
	ls, err := lineStringFromXML(&x)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: ls, SRSName: x.SrsName}, nil
}

func lineStringFromXML(x *gen.LineStringType) (core.LineString, error) {
	dim := derefDim(x.SrsDimension)
	if x.PosList != nil {
		return LineStringFromPosListString(x.PosList.Value, preferDim(dim, derefDim(x.PosList.SrsDimension)))
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
		d := preferDim(dim, derefDim(x.Pos[0].SrsDimension))
		if d == 0 {
			d = len(strings.Fields(x.Pos[0].Value))
			if d < 2 {
				d = 2
			}
		}
		return LineStringFromFlat(flat, d)
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return nil, err
		}
		return LineStringFromFlat(coords, 2)
	}
	return nil, fmt.Errorf("gml: LineString has no coordinate data")
}
