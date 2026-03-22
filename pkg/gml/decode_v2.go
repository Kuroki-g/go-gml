package gml

import (
	"encoding/xml"
	"fmt"

	v2 "github.com/Kuroki-g/go-gml/pkg/gml/v2"
)

// GML 2.x decode functions.
// Called when se.Name.Space == gmlNS2 ("http://www.opengis.net/gml").

func decodePointV2(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v2.PointType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Point(v2): %w", err)
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, x.Coordinates.Cs, x.Coordinates.Ts)
		if err != nil {
			return Geometry{}, err
		}
		pt, err := PointFromFlat(coords, 2)
		if err != nil {
			return Geometry{}, err
		}
		return Geometry{Value: pt, SRSName: x.SrsName}, nil
	}
	return Geometry{}, fmt.Errorf("gml: Point(v2) has no coordinate data")
}

func decodeLineStringV2(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v2.LineStringType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: LineString(v2): %w", err)
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, x.Coordinates.Cs, x.Coordinates.Ts)
		if err != nil {
			return Geometry{}, err
		}
		ls, err := LineStringFromFlat(coords, 2)
		if err != nil {
			return Geometry{}, err
		}
		return Geometry{Value: ls, SRSName: x.SrsName}, nil
	}
	return Geometry{}, fmt.Errorf("gml: LineString(v2) has no coordinate data")
}

func decodeMultiLineStringV2(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v2.MultiLineStringType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: MultiLineString(v2): %w", err)
	}
	var lines MultiLineString
	for _, m := range x.LineStringMember {
		if m.LineString == nil {
			continue
		}
		ls := m.LineString
		if ls.Coordinates != nil {
			coords, err := ParseCoordinates(ls.Coordinates.Value, ls.Coordinates.Cs, ls.Coordinates.Ts)
			if err != nil {
				return Geometry{}, err
			}
			line, err := LineStringFromFlat(coords, 2)
			if err != nil {
				return Geometry{}, err
			}
			lines = append(lines, line)
		}
	}
	return Geometry{Value: lines, SRSName: x.SrsName}, nil
}

func decodeMultiPolygonV2(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v2.MultiPolygonType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: MultiPolygon(v2): %w", err)
	}
	var polys MultiPolygon
	for _, m := range x.PolygonMember {
		if m.Polygon == nil {
			continue
		}
		poly, err := polygonFromV2(m.Polygon)
		if err != nil {
			return Geometry{}, err
		}
		polys = append(polys, poly)
	}
	return Geometry{Value: polys, SRSName: x.SrsName}, nil
}

func polygonFromV2(x *v2.PolygonType) (Polygon, error) {
	var rings []Ring
	if x.OuterBoundaryIs != nil && x.OuterBoundaryIs.LinearRing != nil {
		r, err := ringFromV2LinearRing(x.OuterBoundaryIs.LinearRing)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon(v2) exterior: %w", err)
		}
		rings = append(rings, r)
	}
	for i, ib := range x.InnerBoundaryIs {
		if ib.LinearRing == nil {
			continue
		}
		r, err := ringFromV2LinearRing(ib.LinearRing)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon(v2) interior[%d]: %w", i, err)
		}
		rings = append(rings, r)
	}
	return Polygon(rings), nil
}

func ringFromV2LinearRing(lr *v2.LinearRingType) (Ring, error) {
	if lr.Coordinates != nil {
		return RingFromCoordinatesString(lr.Coordinates.Value, lr.Coordinates.Cs, lr.Coordinates.Ts)
	}
	return nil, fmt.Errorf("gml: LinearRing(v2) has no coordinate data")
}
