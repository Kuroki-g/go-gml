package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

// curveResolver resolves gml:id references to geometry elements.
type curveResolver struct {
	curves           map[string]*gen.CurveType
	orientable       map[string]*gen.OrientableCurveType
	polygonByID      map[string]core.Polygon
	multiPolygonByID map[string]core.MultiPolygon
	lineStringByID   map[string]core.LineString
	gridByID         map[string]*gridBounds
	solidByID        map[string]core.Solid
}

func newCurveResolver() *curveResolver {
	return &curveResolver{
		curves:           make(map[string]*gen.CurveType),
		orientable:       make(map[string]*gen.OrientableCurveType),
		polygonByID:      make(map[string]core.Polygon),
		multiPolygonByID: make(map[string]core.MultiPolygon),
		lineStringByID:   make(map[string]core.LineString),
		gridByID:         make(map[string]*gridBounds),
		solidByID:        make(map[string]core.Solid),
	}
}

// resolve returns the CurveType for the given id (without leading '#').
func (cr *curveResolver) resolve(id string) *gen.CurveType {
	if c, ok := cr.curves[id]; ok {
		return c
	}
	if oc, ok := cr.orientable[id]; ok && oc.BaseCurve != nil {
		if oc.BaseCurve.Curve != nil {
			return oc.BaseCurve.Curve
		}
		if oc.BaseCurve.Href != "" {
			return cr.curves[strings.TrimPrefix(oc.BaseCurve.Href, "#")]
		}
	}
	return nil
}

func decodeCurveType(dec *xml.Decoder, se xml.StartElement) (*gen.CurveType, error) {
	var x gen.CurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return nil, fmt.Errorf("gml: Curve: %w", err)
	}
	return &x, nil
}

// handleCurve decodes a gml:Curve, caches it for xlink:href resolution, and returns a LineString.
func (r *Reader) handleCurve(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	x, err := decodeCurveType(dec, se)
	if err != nil {
		return core.Geometry{}, err
	}
	if x.Id != "" {
		r.resolver.curves[x.Id] = x
	}
	ls, err := lineStringFromCurve(x, 0)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: ls, SRSName: x.SrsName}, nil
}

// cacheOrientableCurve decodes a gml:OrientableCurve and caches it by gml:id.
func (r *Reader) cacheOrientableCurve(dec *xml.Decoder, se xml.StartElement) error {
	var x gen.OrientableCurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return fmt.Errorf("gml: OrientableCurve: %w", err)
	}
	if x.Id != "" {
		r.resolver.orientable[x.Id] = &x
	}
	return nil
}

// lineStringFromCurve converts a gml:Curve to a LineString.
func lineStringFromCurve(x *gen.CurveType, inheritDim uint) (core.LineString, error) {
	if x.Segments == nil || len(x.Segments.LineStringSegment) == 0 {
		return nil, fmt.Errorf("gml: Curve has no LineStringSegment")
	}
	var result core.LineString
	for i, seg := range x.Segments.LineStringSegment {
		var ls core.LineString
		var err error
		curveDim := derefDim(x.SrsDimension)
		if seg.PosList != nil {
			dim := preferDim(preferDim(inheritDim, curveDim), derefDim(seg.PosList.SrsDimension))
			ls, err = core.LineStringFromPosListString(seg.PosList.Value, dim)
		} else if seg.Coordinates != nil {
			var coords []float64
			coords, err = core.ParseCoordinates(seg.Coordinates.Value, derefStrOr(seg.Coordinates.Cs, ","), derefStrOr(seg.Coordinates.Ts, " "))
			if err == nil {
				ls, err = core.LineStringFromFlat(coords, 2)
			}
		} else if len(seg.Pos) > 0 {
			ls, err = lineStringFromPosSlice(seg.Pos, preferDim(inheritDim, curveDim))
		} else if len(seg.PointProperty) > 0 {
			for j := range seg.PointProperty {
				pt, ptErr := fromPointProperty(&seg.PointProperty[j], j, preferDim(inheritDim, curveDim))
				if ptErr != nil {
					err = ptErr
					break
				}
				ls = append(ls, pt)
			}
		} else if len(seg.PointRep) > 0 {
			for j := range seg.PointRep {
				pt, ptErr := fromPointProperty(&seg.PointRep[j], j, preferDim(inheritDim, curveDim))
				if ptErr != nil {
					err = ptErr
					break
				}
				ls = append(ls, pt)
			}
		} else {
			return nil, fmt.Errorf("gml: LineStringSegment[%d] has no coordinate data", i)
		}
		if err != nil {
			return nil, err
		}
		if len(result) > 0 && len(ls) > 0 {
			result = append(result, ls[1:]...)
		} else {
			result = append(result, ls...)
		}
	}
	return result, nil
}

func lineStringFromPosSlice(poses []gen.DirectPositionType, inheritDim uint) (core.LineString, error) {
	var result core.LineString
	for j, p := range poses {
		dim := preferDim(inheritDim, derefDim(p.SrsDimension))
		pt, err := core.PointFromPosString(p.Value, dim)
		if err != nil {
			return nil, fmt.Errorf("pos[%d]: %w", j, err)
		}
		result = append(result, pt)
	}
	return result, nil
}

// fromPointProperty extracts a Point from a PointPropertyType.
// Returns an error if the property uses xlink:href (unresolved reference) or has no Point element.
func fromPointProperty(pp *gen.PointPropertyType, j int, inheritDim uint) (core.Point, error) {
	_ = pp.TypeField
	_ = pp.Role
	_ = pp.Arcrole
	_ = pp.Title
	_ = pp.Show
	_ = pp.Actuate
	_ = pp.RemoteSchema
	_ = pp.Owns
	if pp.NilReason != nil {
		return core.Point{}, fmt.Errorf("pointProperty[%d]: nil point (nilReason=%s)", j, *pp.NilReason)
	}
	if pp.Href != "" {
		return core.Point{}, fmt.Errorf("pointProperty[%d]: unresolved xlink:href %q", j, pp.Href)
	}
	if pp.Point == nil {
		return core.Point{}, fmt.Errorf("pointProperty[%d]: missing Point element", j)
	}
	p := pp.Point
	if p.SrsDimension == nil && inheritDim > 0 {
		p.SrsDimension = &inheritDim
	}
	pt, err := pointFromXML(p)
	if err != nil {
		return core.Point{}, fmt.Errorf("pointProperty[%d]: %w", j, err)
	}
	return pt, nil
}
