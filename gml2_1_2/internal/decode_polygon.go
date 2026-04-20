package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
)

func decodePolygonElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.PolygonType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Polygon: %w", err)
	}
	poly, err := polygonFromXML(&x)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: poly, SRSName: x.SrsName}, nil
}

func ringFromLinearRing(lr *gen.LinearRingType) (core.Ring, error) {
	if lr == nil {
		return nil, fmt.Errorf("gml: nil LinearRing")
	}
	if lr.Coordinates != nil {
		return core.RingFromCoordinatesString(lr.Coordinates.Value, derefStrOr(lr.Coordinates.Cs, ","), derefStrOr(lr.Coordinates.Ts, " "), derefStrOr(lr.Coordinates.Decimal, "."))
	}
	if len(lr.Coord) > 0 {
		flat, dim := coordSliceToFlat(lr.Coord)
		return core.RingFromFlat(flat, dim)
	}
	return nil, fmt.Errorf("gml: LinearRing has no coordinate data")
}

// fromLinearRingMember extracts the LinearRingType from a LinearRingMemberType.
// Returns nil when the member is an xlink:href reference — gml2_1_2 has no
// prescan resolver, so href references are silently skipped.
func fromLinearRingMember(m *gen.LinearRingMemberType) *gen.LinearRingType {
	_ = m.RemoteSchema
	_ = m.TypeField
	_ = m.Role
	_ = m.Arcrole
	_ = m.Title
	_ = m.Show
	_ = m.Actuate
	if m.Href != "" {
		return nil
	}
	return m.LinearRing
}

func polygonFromXML(x *gen.PolygonType) (core.Polygon, error) {
	var rings []core.Ring
	if x.OuterBoundaryIs != nil && x.OuterBoundaryIs.LinearRing != nil {
		r, err := ringFromLinearRing(x.OuterBoundaryIs.LinearRing)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon outerBoundaryIs: %w", err)
		}
		rings = append(rings, r)
	}
	for i := range x.InnerBoundaryIs {
		lr := fromLinearRingMember(&x.InnerBoundaryIs[i])
		if lr == nil {
			continue
		}
		r, err := ringFromLinearRing(lr)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon innerBoundaryIs[%d]: %w", i, err)
		}
		rings = append(rings, r)
	}
	return core.Polygon(rings), nil
}
