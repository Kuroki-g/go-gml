package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
)

// srsNamePtr returns a non-nil *string only when s is non-empty.
// Multi* types declare srsName as required (string, not *string), so an empty
// value means the attribute was omitted; in that case, SRSName should be nil.
func srsNamePtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// ---- multi-point ----

func decodeMultiPointElement(dec *xml.Decoder, se xml.StartElement, res *resolver) (core.Geometry, error) {
	var x gen.MultiPointType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: MultiPoint: %w", err)
	}
	var pts core.MultiPoint
	for i := range x.PointMember {
		pt, err := fromPointMember(&x.PointMember[i], res)
		if err != nil {
			return core.Geometry{}, fmt.Errorf("gml: MultiPoint pointMember[%d]: %w", i, err)
		}
		pts = append(pts, pt)
	}
	return core.Geometry{Value: pts, SRSName: srsNamePtr(x.SrsName)}, nil
}

// fromPointMember extracts a Point from a PointMemberType.
// If m.Href is set, the resolver is consulted to find the referenced Point.
func fromPointMember(m *gen.PointMemberType, res *resolver) (core.Point, error) {
	_ = m.TypeField
	_ = m.Role
	_ = m.Arcrole
	_ = m.Title
	_ = m.Show
	_ = m.Actuate
	_ = m.RemoteSchema
	if m.Href != "" {
		id := stripHash(m.Href)
		if pt, ok := res.pointByID[id]; ok {
			return pt, nil
		}
		return core.Point{}, fmt.Errorf("gml: unresolved xlink:href %q", m.Href)
	}
	if m.Point == nil {
		return core.Point{}, fmt.Errorf("gml: missing Point element")
	}
	return pointFromXML(m.Point)
}

// ---- multi-linestring ----

func decodeMultiLineStringElement(dec *xml.Decoder, se xml.StartElement, res *resolver) (core.Geometry, error) {
	var x gen.MultiLineStringType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: MultiLineString: %w", err)
	}
	var lines core.MultiLineString
	for i := range x.LineStringMember {
		ls, err := fromLineStringMember(&x.LineStringMember[i], res)
		if err != nil {
			return core.Geometry{}, fmt.Errorf("gml: MultiLineString lineStringMember[%d]: %w", i, err)
		}
		lines = append(lines, ls)
	}
	return core.Geometry{Value: lines, SRSName: srsNamePtr(x.SrsName)}, nil
}

// fromLineStringMember extracts a LineString from a LineStringMemberType.
// If m.Href is set, the resolver is consulted to find the referenced LineString.
func fromLineStringMember(m *gen.LineStringMemberType, res *resolver) (core.LineString, error) {
	_ = m.TypeField
	_ = m.Role
	_ = m.Arcrole
	_ = m.Title
	_ = m.Show
	_ = m.Actuate
	_ = m.RemoteSchema
	if m.Href != "" {
		id := stripHash(m.Href)
		if ls, ok := res.lineStringByID[id]; ok {
			return ls, nil
		}
		return nil, fmt.Errorf("gml: unresolved xlink:href %q", m.Href)
	}
	if m.LineString == nil {
		return nil, fmt.Errorf("gml: missing LineString element")
	}
	return lineStringFromXML(m.LineString)
}

// ---- multi-polygon ----

func decodeMultiPolygonElement(dec *xml.Decoder, se xml.StartElement, res *resolver) (core.Geometry, error) {
	var x gen.MultiPolygonType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: MultiPolygon: %w", err)
	}
	var polys core.MultiPolygon
	for i := range x.PolygonMember {
		poly, err := fromPolygonMember(&x.PolygonMember[i], res)
		if err != nil {
			return core.Geometry{}, fmt.Errorf("gml: MultiPolygon polygonMember[%d]: %w", i, err)
		}
		polys = append(polys, poly)
	}
	return core.Geometry{Value: polys, SRSName: srsNamePtr(x.SrsName)}, nil
}

// fromPolygonMember extracts a Polygon from a PolygonMemberType.
// If m.Href is set, the resolver is consulted to find the referenced Polygon.
func fromPolygonMember(m *gen.PolygonMemberType, res *resolver) (core.Polygon, error) {
	_ = m.TypeField
	_ = m.Role
	_ = m.Arcrole
	_ = m.Title
	_ = m.Show
	_ = m.Actuate
	_ = m.RemoteSchema
	if m.Href != "" {
		id := stripHash(m.Href)
		if poly, ok := res.polygonByID[id]; ok {
			return poly, nil
		}
		return nil, fmt.Errorf("gml: unresolved xlink:href %q", m.Href)
	}
	if m.Polygon == nil {
		return nil, fmt.Errorf("gml: missing Polygon element")
	}
	return polygonFromXML(m.Polygon)
}

// ---- box (GML 2.x bounding box) ----

func decodeBoxElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.BoxType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Box: %w", err)
	}
	b, err := boundFromBoxXML(&x)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: b, SRSName: x.SrsName}, nil
}

func boundFromBoxXML(x *gen.BoxType) (core.Bound, error) {
	if x.Coordinates != nil {
		coords, err := core.ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return core.Bound{}, fmt.Errorf("gml: Box coordinates: %w", err)
		}
		if len(coords) < 4 {
			return core.Bound{}, fmt.Errorf("gml: Box coordinates: need at least 2 points")
		}
		return core.Bound{Min: core.Point{coords[0], coords[1]}, Max: core.Point{coords[2], coords[3]}}, nil
	}
	if len(x.Coord) >= 2 {
		y0 := 0.0
		if x.Coord[0].Y != nil {
			y0 = *x.Coord[0].Y
		}
		y1 := 0.0
		if x.Coord[1].Y != nil {
			y1 = *x.Coord[1].Y
		}
		return core.Bound{
			Min: core.Point{x.Coord[0].X, y0},
			Max: core.Point{x.Coord[1].X, y1},
		}, nil
	}
	return core.Bound{}, fmt.Errorf("gml: Box has no coordinate data")
}
