package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
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
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return nil, err
		}
		return LineStringFromFlat(coords, 2)
	}
	if len(x.Coord) > 0 {
		return LineStringFromFlat(coordSliceToFlat(x.Coord), 2)
	}
	return nil, fmt.Errorf("gml: LineString has no coordinate data")
}
