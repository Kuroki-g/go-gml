package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

// ---- multi-point ----

func decodeMultiPointElement(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.MultiPointType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: MultiPoint: %w", err)
	}
	var pts core.MultiPoint
	dim := preferDim(x.SrsDimension, nil)
	for i := range x.PointMember {
		pt, err := fromPointProperty(&x.PointMember[i], i, dim)
		if err != nil {
			continue // xlink:href or missing Point — skip
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
	dim := preferDim(x.SrsDimension, r.globalDim)
	var lines core.MultiLineString
	for i := range x.CurveMember {
		ls, err := lineStringFromCurveProperty(&x.CurveMember[i], dim, x.SrsName, r.resolver)
		if err != nil {
			return core.Geometry{}, fmt.Errorf("gml: %s curveMember[%d]: %w", se.Name.Local, i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	if x.CurveMembers != nil {
		extra, err := lineStringsFromCurveArrayProperty(x.CurveMembers, dim, x.SrsName, r.resolver)
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
	dim := preferDim(x.SrsDimension, r.globalDim)
	var polys core.MultiPolygon
	for i := range x.SurfaceMember {
		mp, err := multiPolygonFromSurfaceProperty(&x.SurfaceMember[i], dim, x.SrsName, r.resolver)
		if err != nil {
			return core.Geometry{}, fmt.Errorf("gml: %s surfaceMember[%d]: %w", se.Name.Local, i, err)
		}
		polys = append(polys, mp...)
	}
	if x.SurfaceMembers != nil {
		extra, err := polygonsFromSurfaceArrayProperty(x.SurfaceMembers, dim, x.SrsName, r.resolver)
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

func lineStringsFromCurveArrayProperty(a *gen.CurveArrayPropertyType, inheritDim *uint, inheritSrsName *string, resolver *curveResolver) (core.MultiLineString, error) {
	// xlink ownership attribute — not used for curve collection resolution.
	_ = a.Owns
	var lines core.MultiLineString
	for i := range a.Curve {
		cm := gen.CurvePropertyType{Curve: &a.Curve[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, inheritSrsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("Curve[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.LineString {
		cm := gen.CurvePropertyType{LineString: &a.LineString[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, inheritSrsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("LineString[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.CompositeCurve {
		cm := gen.CurvePropertyType{CompositeCurve: &a.CompositeCurve[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, inheritSrsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("CompositeCurve[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.OrientableCurve {
		cm := gen.CurvePropertyType{OrientableCurve: &a.OrientableCurve[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, inheritSrsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("OrientableCurve[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.LinearRing {
		cm := gen.CurvePropertyType{LinearRing: &a.LinearRing[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, inheritSrsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("LinearRing[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	for i := range a.Ring {
		cm := gen.CurvePropertyType{Ring: &a.Ring[i]}
		ls, err := lineStringFromCurveProperty(&cm, inheritDim, inheritSrsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("Ring[%d]: %w", i, err)
		}
		if ls != nil {
			lines = append(lines, ls)
		}
	}
	return lines, nil
}

func polygonsFromSurfaceArrayProperty(a *gen.SurfaceArrayPropertyType, inheritDim *uint, inheritSrsName *string, resolver *curveResolver) (core.MultiPolygon, error) {
	// xlink ownership attribute — not used for surface collection resolution.
	_ = a.Owns
	var polys core.MultiPolygon
	for i := range a.Polygon {
		sm := gen.SurfacePropertyType{Polygon: &a.Polygon[i]}
		poly, err := polygonFromSurfaceProperty(&sm, inheritDim, inheritSrsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("Polygon[%d]: %w", i, err)
		}
		if len(poly) > 0 {
			polys = append(polys, poly)
		}
	}
	for i := range a.Surface {
		sm := gen.SurfacePropertyType{Surface: &a.Surface[i]}
		poly, err := polygonFromSurfaceProperty(&sm, inheritDim, inheritSrsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("Surface[%d]: %w", i, err)
		}
		if len(poly) > 0 {
			polys = append(polys, poly)
		}
	}
	for i := range a.CompositeSurface {
		mp, err := multiPolygonFromCompositeSurface(&a.CompositeSurface[i], resolver, inheritDim, inheritSrsName)
		if err != nil {
			return nil, fmt.Errorf("CompositeSurface[%d]: %w", i, err)
		}
		polys = append(polys, mp...)
	}
	for i := range a.OrientableSurface {
		sm := gen.SurfacePropertyType{OrientableSurface: &a.OrientableSurface[i]}
		poly, err := polygonFromSurfaceProperty(&sm, inheritDim, inheritSrsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("OrientableSurface[%d]: %w", i, err)
		}
		if len(poly) > 0 {
			polys = append(polys, poly)
		}
	}
	for i := range a.Tin {
		t := &a.Tin[i]
		mp, err := multiPolygonFromSurfacePatchArrayProperty(t.TrianglePatches, inheritDim, preferSrsName(t.SrsName, inheritSrsName), resolver)
		if err != nil {
			return nil, fmt.Errorf("Tin[%d]: %w", i, err)
		}
		polys = append(polys, mp...)
	}
	for i := range a.TriangulatedSurface {
		ts := &a.TriangulatedSurface[i]
		mp, err := multiPolygonFromSurfacePatchArrayProperty(ts.Patches, inheritDim, preferSrsName(ts.SrsName, inheritSrsName), resolver)
		if err != nil {
			return nil, fmt.Errorf("TriangulatedSurface[%d]: %w", i, err)
		}
		polys = append(polys, mp...)
	}
	for i := range a.PolyhedralSurface {
		ps := &a.PolyhedralSurface[i]
		mp, err := multiPolygonFromSurfacePatchArrayProperty(ps.Patches, inheritDim, preferSrsName(ps.SrsName, inheritSrsName), resolver)
		if err != nil {
			return nil, fmt.Errorf("PolyhedralSurface[%d]: %w", i, err)
		}
		polys = append(polys, mp...)
	}
	for i := range a.Shell {
		mp, err := multiPolygonFromShell(&a.Shell[i], inheritDim, inheritSrsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("Shell[%d]: %w", i, err)
		}
		polys = append(polys, mp...)
	}
	return polys, nil
}

func boundFromXML(x *gen.EnvelopeType) (core.Bound, error) {
	resolvedDim := preferDim(x.SrsDimension, nil)
	if x.LowerCorner != nil && x.UpperCorner != nil {
		d := preferDim(x.LowerCorner.SrsDimension, resolvedDim)
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
		d := preferDim(x.Pos[0].SrsDimension, resolvedDim)
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
