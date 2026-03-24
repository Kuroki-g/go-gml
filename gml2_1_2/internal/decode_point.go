package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
)

func decodePointElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.PointType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Point: %w", err)
	}
	pt, err := pointFromXML(&x)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: pt, SRSName: x.SrsName}, nil
}

func pointFromXML(x *gen.PointType) (core.Point, error) {
	if x.Coordinates != nil {
		coords, err := core.ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return core.Point{}, err
		}
		return core.PointFromFlat(coords, 2)
	}
	if x.Coord != nil {
		return core.Point{x.Coord.X, x.Coord.Y}, nil
	}
	return core.Point{}, fmt.Errorf("gml: Point has no coordinate data")
}
