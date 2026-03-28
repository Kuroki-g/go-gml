package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
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
		if m.Point.SrsDimension == nil {
			m.Point.SrsDimension = x.SrsDimension
		}
		pt, err := pointFromXML(m.Point)
		if err != nil {
			return core.Geometry{}, err
		}
		pts = append(pts, pt)
	}
	return core.Geometry{Value: pts, SRSName: x.SrsName}, nil
}

// ---- multi-curve / multi-linestring ----

func (r *Reader) handleMultiCurve(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.MultiCurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: %s: %w", se.Name.Local, err)
	}
	dim := preferDim(derefDim(x.SrsDimension), r.globalDim)
	var lines core.MultiLineString
	for i := range x.CurveMember {
		ls, err := lineStringFromCurveProperty(&x.CurveMember[i], dim, r.resolver)
		if err != nil {
			return core.Geometry{}, fmt.Errorf("gml: %s curveMember[%d]: %w", se.Name.Local, i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	if x.CurveMembers != nil {
		extra, err := lineStringsFromCurveArrayProperty(x.CurveMembers, dim, r.resolver)
		if err != nil {
			return core.Geometry{}, fmt.Errorf("gml: %s curveMembers: %w", se.Name.Local, err)
		}
		lines = append(lines, extra...)
	}
	return core.Geometry{Value: lines, SRSName: x.SrsName}, nil
}

// ---- multi-surface / multi-polygon ----

func (r *Reader) handleMultiSurface(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.MultiSurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: %s: %w", se.Name.Local, err)
	}
	dim := preferDim(derefDim(x.SrsDimension), r.globalDim)
	var polys core.MultiPolygon
	for i := range x.SurfaceMember {
		mp, err := multiPolygonFromSurfaceProperty(&x.SurfaceMember[i], dim, r.resolver)
		if err != nil {
			return core.Geometry{}, fmt.Errorf("gml: %s surfaceMember[%d]: %w", se.Name.Local, i, err)
		}
		polys = append(polys, mp...)
	}
	if x.SurfaceMembers != nil {
		extra, err := polygonsFromSurfaceArrayProperty(x.SurfaceMembers, dim, r.resolver)
		if err != nil {
			return core.Geometry{}, fmt.Errorf("gml: %s surfaceMembers: %w", se.Name.Local, err)
		}
		polys = append(polys, extra...)
	}
	return core.Geometry{Value: polys, SRSName: x.SrsName}, nil
}

// ---- envelope ----

func decodeEnvelopeElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.EnvelopeType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Envelope: %w", err)
	}
	b, err := boundFromXML(&x)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: b, SRSName: x.SrsName}, nil
}

func lineStringsFromCurveArrayProperty(a *gen.CurveArrayPropertyType, inheritDim int, resolver *curveResolver) (core.MultiLineString, error) {
	var lines core.MultiLineString
	for i := range a.Curve {
		cm := gen.CurvePropertyType{Curve: &a.Curve[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("Curve[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.LineString {
		cm := gen.CurvePropertyType{LineString: &a.LineString[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("LineString[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.CompositeCurve {
		cm := gen.CurvePropertyType{CompositeCurve: &a.CompositeCurve[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("CompositeCurve[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.OrientableCurve {
		cm := gen.CurvePropertyType{OrientableCurve: &a.OrientableCurve[i]}
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

func polygonsFromSurfaceArrayProperty(a *gen.SurfaceArrayPropertyType, inheritDim int, resolver *curveResolver) (core.MultiPolygon, error) {
	var polys core.MultiPolygon
	for i := range a.Polygon {
		sm := gen.SurfacePropertyType{Polygon: &a.Polygon[i]}
		poly, err := polygonFromSurfaceProperty(&sm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("Polygon[%d]: %w", i, err)
		}
		if len(poly) > 0 {
			polys = append(polys, poly)
		}
	}
	for i := range a.Surface {
		sm := gen.SurfacePropertyType{Surface: &a.Surface[i]}
		poly, err := polygonFromSurfaceProperty(&sm, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("Surface[%d]: %w", i, err)
		}
		if len(poly) > 0 {
			polys = append(polys, poly)
		}
	}
	for i := range a.CompositeSurface {
		mp, err := multiPolygonFromCompositeSurface(&a.CompositeSurface[i], resolver, inheritDim)
		if err != nil {
			return nil, fmt.Errorf("CompositeSurface[%d]: %w", i, err)
		}
		polys = append(polys, mp...)
	}
	for i := range a.OrientableSurface {
		sm := gen.SurfacePropertyType{OrientableSurface: &a.OrientableSurface[i]}
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

func boundFromXML(x *gen.EnvelopeType) (core.Bound, error) {
	dim := derefDim(x.SrsDimension)
	if x.LowerCorner != nil && x.UpperCorner != nil {
		d := preferDim(dim, derefDim(x.LowerCorner.SrsDimension))
		lo, err := core.PointFromPosString(x.LowerCorner.Value, d)
		if err != nil {
			return core.Bound{}, fmt.Errorf("gml: Envelope lowerCorner: %w", err)
		}
		hi, err := core.PointFromPosString(x.UpperCorner.Value, d)
		if err != nil {
			return core.Bound{}, fmt.Errorf("gml: Envelope upperCorner: %w", err)
		}
		return core.Bound{Min: lo, Max: hi}, nil
	}
	if len(x.Pos) >= 2 {
		d := preferDim(dim, derefDim(x.Pos[0].SrsDimension))
		lo, err := core.PointFromPosString(x.Pos[0].Value, d)
		if err != nil {
			return core.Bound{}, fmt.Errorf("gml: Envelope pos[0]: %w", err)
		}
		hi, err := core.PointFromPosString(x.Pos[1].Value, d)
		if err != nil {
			return core.Bound{}, fmt.Errorf("gml: Envelope pos[1]: %w", err)
		}
		return core.Bound{Min: lo, Max: hi}, nil
	}
	if x.Coordinates != nil {
		coords, err := core.ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return core.Bound{}, fmt.Errorf("gml: Envelope coordinates: %w", err)
		}
		if len(coords) < 4 {
			return core.Bound{}, fmt.Errorf("gml: Envelope coordinates: need at least 2 points")
		}
		return core.Bound{Min: core.Point{coords[0], coords[1]}, Max: core.Point{coords[2], coords[3]}}, nil
	}
	return core.Bound{}, fmt.Errorf("gml: Envelope has no coordinate data")
}
