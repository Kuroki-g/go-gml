package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

func decodeLineStringElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.LineStringType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: LineString: %w", err)
	}
	ls, err := lineStringFromXML(&x, nil, nil)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: ls, SRSName: x.SrsName}, nil
}

func lineStringFromXML(x *gen.LineStringType, inheritDim *uint, inheritSrsName *string) (core.LineString, error) {
	resolvedDim := preferDim(x.SrsDimension, inheritDim)
	resolvedSrsName := preferSrsName(x.SrsName, inheritSrsName)
	if x.PosList != nil {
		return core.LineStringFromPosListString(x.PosList.Value, preferDim(x.PosList.SrsDimension, resolvedDim), preferSrsName(x.PosList.SrsName, resolvedSrsName))
	}
	if len(x.Pos) > 0 {
		var flat []float64
		for _, p := range x.Pos {
			vals, err := core.ParsePosList(p.Value)
			if err != nil {
				return nil, err
			}
			flat = append(flat, vals...)
		}
		d := uint(len(strings.Fields(x.Pos[0].Value)))
		if d == 0 {
			return nil, fmt.Errorf("gml: pos has no values")
		}
		return core.LineStringFromFlat(flat, d)
	}
	if x.Coordinates != nil {
		coords, err := core.ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return nil, err
		}
		return core.LineStringFromFlat(coords, 2)
	}
	return nil, fmt.Errorf("gml: LineString has no coordinate data")
}
