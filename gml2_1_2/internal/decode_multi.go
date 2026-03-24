package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
)

// ---- multi-point ----

func decodeMultiPointElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.MultiPointType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: MultiPoint: %w", err)
	}
	var pts core.MultiPoint
	for _, m := range x.PointMember {
		if m.Point == nil {
			continue
		}
		pt, err := pointFromXML(m.Point)
		if err != nil {
			return core.Geometry{}, err
		}
		pts = append(pts, pt)
	}
	return core.Geometry{Value: pts, SRSName: x.SrsName}, nil
}

// ---- multi-linestring ----

func decodeMultiLineStringElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.MultiLineStringType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: MultiLineString: %w", err)
	}
	var lines core.MultiLineString
	for _, m := range x.LineStringMember {
		if m.LineString == nil {
			continue
		}
		ls, err := lineStringFromXML(m.LineString)
		if err != nil {
			return core.Geometry{}, err
		}
		lines = append(lines, ls)
	}
	return core.Geometry{Value: lines, SRSName: x.SrsName}, nil
}

// ---- multi-polygon ----

func decodeMultiPolygonElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.MultiPolygonType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: MultiPolygon: %w", err)
	}
	var polys core.MultiPolygon
	for _, m := range x.PolygonMember {
		if m.Polygon == nil {
			continue
		}
		poly, err := polygonFromXML(m.Polygon)
		if err != nil {
			return core.Geometry{}, err
		}
		polys = append(polys, poly)
	}
	return core.Geometry{Value: polys, SRSName: x.SrsName}, nil
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
		return core.Bound{
			Min: core.Point{x.Coord[0].X, x.Coord[0].Y},
			Max: core.Point{x.Coord[1].X, x.Coord[1].Y},
		}, nil
	}
	return core.Bound{}, fmt.Errorf("gml: Box has no coordinate data")
}
