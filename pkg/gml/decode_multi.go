package gml

import (
	"encoding/xml"
	"fmt"

	v3_2_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

// ---- multi-point ----

func decodeMultiPointElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v3_2_1.MultiPointType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: MultiPoint: %w", err)
	}
	var pts MultiPoint
	for _, m := range x.PointMember {
		if m.Point == nil {
			continue
		}
		if m.Point.SrsDimension == nil {
			m.Point.SrsDimension = x.SrsDimension
		}
		pt, err := pointFromXML(m.Point)
		if err != nil {
			return Geometry{}, err
		}
		pts = append(pts, pt)
	}
	return Geometry{Value: pts, SRSName: x.SrsName}, nil
}

// ---- multi-curve / multi-linestring ----

func (r *Reader) handleMultiCurve(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	if se.Name.Space == gmlNS2 {
		return decodeMultiLineStringV2(dec, se)
	}
	var x v3_2_1.MultiCurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: %s: %w", se.Name.Local, err)
	}
	dim := derefDim(x.SrsDimension)
	var lines MultiLineString
	for i := range x.CurveMember {
		ls, err := lineStringFromCurveProperty(&x.CurveMember[i], dim, r.resolver)
		if err != nil {
			return Geometry{}, fmt.Errorf("gml: %s curveMember[%d]: %w", se.Name.Local, i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	if x.CurveMembers != nil {
		extra, err := lineStringsFromCurveArrayProperty(x.CurveMembers, dim, r.resolver)
		if err != nil {
			return Geometry{}, fmt.Errorf("gml: %s curveMembers: %w", se.Name.Local, err)
		}
		lines = append(lines, extra...)
	}
	return Geometry{Value: lines, SRSName: x.SrsName}, nil
}

// ---- multi-surface / multi-polygon ----

func (r *Reader) handleMultiSurface(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	if se.Name.Space == gmlNS2 {
		return decodeMultiPolygonV2(dec, se)
	}
	var x v3_2_1.MultiSurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: %s: %w", se.Name.Local, err)
	}
	dim := derefDim(x.SrsDimension)
	var polys MultiPolygon
	for i := range x.SurfaceMember {
		poly, err := polygonFromSurfaceProperty(&x.SurfaceMember[i], dim, r.resolver)
		if err != nil {
			return Geometry{}, fmt.Errorf("gml: %s surfaceMember[%d]: %w", se.Name.Local, i, err)
		}
		if len(poly) > 0 {
			polys = append(polys, poly)
		}
	}
	if x.SurfaceMembers != nil {
		extra, err := polygonsFromSurfaceArrayProperty(x.SurfaceMembers, dim, r.resolver)
		if err != nil {
			return Geometry{}, fmt.Errorf("gml: %s surfaceMembers: %w", se.Name.Local, err)
		}
		polys = append(polys, extra...)
	}
	return Geometry{Value: polys, SRSName: x.SrsName}, nil
}

// ---- envelope ----

func decodeEnvelopeElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v3_2_1.EnvelopeType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Envelope: %w", err)
	}
	b, err := boundFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: b, SRSName: x.SrsName}, nil
}

// lineStringsFromCurveArrayProperty converts a CurveArrayPropertyType (curveMembers)
// to a slice of LineStrings by wrapping each element in a CurvePropertyType and
// calling lineStringFromCurveProperty.
func lineStringsFromCurveArrayProperty(a *v3_2_1.CurveArrayPropertyType, inheritDim int, resolver *curveResolver) (MultiLineString, error) {
	var lines MultiLineString
	for i := range a.Curve {
		cm := v3_2_1.CurvePropertyType{Curve: &a.Curve[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("Curve[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.LineString {
		cm := v3_2_1.CurvePropertyType{LineString: &a.LineString[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("LineString[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.CompositeCurve {
		cm := v3_2_1.CurvePropertyType{CompositeCurve: &a.CompositeCurve[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("CompositeCurve[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.OrientableCurve {
		cm := v3_2_1.CurvePropertyType{OrientableCurve: &a.OrientableCurve[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("OrientableCurve[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	return lines, nil
}

// polygonsFromSurfaceArrayProperty converts a SurfaceArrayPropertyType (surfaceMembers)
// to a slice of Polygons by wrapping each element in a SurfacePropertyType and
// calling polygonFromSurfaceProperty.
func polygonsFromSurfaceArrayProperty(a *v3_2_1.SurfaceArrayPropertyType, inheritDim int, resolver *curveResolver) (MultiPolygon, error) {
	var polys MultiPolygon
	for i := range a.Polygon {
		sm := v3_2_1.SurfacePropertyType{Polygon: &a.Polygon[i]}
		poly, err := polygonFromSurfaceProperty(&sm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("Polygon[%d]: %w", i, err)
		}
		if len(poly) > 0 {
			polys = append(polys, poly)
		}
	}
	for i := range a.Surface {
		sm := v3_2_1.SurfacePropertyType{Surface: &a.Surface[i]}
		poly, err := polygonFromSurfaceProperty(&sm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("Surface[%d]: %w", i, err)
		}
		if len(poly) > 0 {
			polys = append(polys, poly)
		}
	}
	for i := range a.CompositeSurface {
		sm := v3_2_1.SurfacePropertyType{CompositeSurface: &a.CompositeSurface[i]}
		poly, err := polygonFromSurfaceProperty(&sm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("CompositeSurface[%d]: %w", i, err)
		}
		if len(poly) > 0 {
			polys = append(polys, poly)
		}
	}
	for i := range a.OrientableSurface {
		sm := v3_2_1.SurfacePropertyType{OrientableSurface: &a.OrientableSurface[i]}
		poly, err := polygonFromSurfaceProperty(&sm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("OrientableSurface[%d]: %w", i, err)
		}
		if len(poly) > 0 {
			polys = append(polys, poly)
		}
	}
	return polys, nil
}

func boundFromXML(x *v3_2_1.EnvelopeType) (Bound, error) {
	dim := derefDim(x.SrsDimension)
	if x.LowerCorner != nil && x.UpperCorner != nil {
		d := preferDim(dim, derefDim(x.LowerCorner.SrsDimension))
		lo, err := PointFromPosString(x.LowerCorner.Value, d)
		if err != nil {
			return Bound{}, fmt.Errorf("gml: Envelope lowerCorner: %w", err)
		}
		hi, err := PointFromPosString(x.UpperCorner.Value, d)
		if err != nil {
			return Bound{}, fmt.Errorf("gml: Envelope upperCorner: %w", err)
		}
		return Bound{Min: lo, Max: hi}, nil
	}
	if len(x.Pos) >= 2 {
		d := preferDim(dim, derefDim(x.Pos[0].SrsDimension))
		lo, err := PointFromPosString(x.Pos[0].Value, d)
		if err != nil {
			return Bound{}, fmt.Errorf("gml: Envelope pos[0]: %w", err)
		}
		hi, err := PointFromPosString(x.Pos[1].Value, d)
		if err != nil {
			return Bound{}, fmt.Errorf("gml: Envelope pos[1]: %w", err)
		}
		return Bound{Min: lo, Max: hi}, nil
	}
	return Bound{}, fmt.Errorf("gml: Envelope has no coordinate data")
}
