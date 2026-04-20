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
		return core.PointFromCoordinatesString(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "), derefStrOr(x.Coordinates.Decimal, "."))
	}
	if x.Coord != nil {
		y := 0.0
		if x.Coord.Y != nil {
			y = *x.Coord.Y
		}
		if x.Coord.Z != nil {
			return core.Point{x.Coord.X, y, *x.Coord.Z}, nil
		}
		return core.Point{x.Coord.X, y}, nil
	}
	return core.Point{}, fmt.Errorf("gml: Point has no coordinate data")
}
