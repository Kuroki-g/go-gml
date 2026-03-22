package gml

import (
	"encoding/xml"
	"fmt"

	v3 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

func decodePointElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	if se.Name.Space == gmlNS2 {
		return decodePointV2(dec, se)
	}
	var x v3.PointType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Point: %w", err)
	}
	pt, err := pointFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: pt, SRSName: x.SrsName}, nil
}

func pointFromXML(x *v3.PointType) (Point, error) {
	if x.Pos != nil {
		return PointFromPosString(x.Pos.Value, preferDim(derefDim(x.SrsDimension), derefDim(x.Pos.SrsDimension)))
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
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
