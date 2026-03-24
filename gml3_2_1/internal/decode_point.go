package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
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
	if x.Pos != nil {
		return PointFromPosString(x.Pos.Value, preferDim(derefDim(x.SrsDimension), derefDim(x.Pos.SrsDimension)))
	}
	if x.Coordinates != nil {
		coords, err := core.ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return core.Point{}, err
		}
		return PointFromFlat(coords, 2)
	}
	return core.Point{}, fmt.Errorf("gml: Point has no coordinate data")
}

// preferDim returns a if non-zero, otherwise b.
func preferDim(a, b int) int {
	if a > 0 {
		return a
	}
	return b
}

// derefDim dereferences a *int dimension attribute; nil → 0.
func derefDim(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

// derefStrOr dereferences a *string attribute; nil → def (the XSD default value).
func derefStrOr(p *string, def string) string {
	if p == nil {
		return def
	}
	return *p
}
