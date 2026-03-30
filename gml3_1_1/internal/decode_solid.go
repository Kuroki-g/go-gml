package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

// handleSolid decodes a gml:Solid and returns a core.Solid.
func (r *Reader) handleSolid(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.SolidType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Solid: %w", err)
	}
	dim := preferDim(derefDim(x.SrsDimension), r.globalDim)
	solid, err := solidFromXML(&x, dim, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: solid, SRSName: x.SrsName}, nil
}

// handleCompositeSolid decodes a gml:CompositeSolid and returns a core.Solid
// with all member exteriors merged. XSD: CompositeSolid is a _Solid substitutionGroup
// member with "all the geometric properties of a (primitive) solid".
func (r *Reader) handleCompositeSolid(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.CompositeSolidType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: CompositeSolid: %w", err)
	}
	dim := preferDim(derefDim(x.SrsDimension), r.globalDim)
	solid, err := solidFromSolidMembers(x.SolidMember, dim, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: solid, SRSName: x.SrsName}, nil
}

// handleMultiSolid decodes a gml:MultiSolid and returns a core.MultiSolid.
// XSD: MultiSolid is substitutionGroup="_GeometricAggregate", not "_Solid",
// so it is a collection rather than a single solid.
func (r *Reader) handleMultiSolid(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.MultiSolidType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: MultiSolid: %w", err)
	}
	dim := preferDim(derefDim(x.SrsDimension), r.globalDim)
	solids, err := solidsFromSolidMembers(x.SolidMember, dim, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	if x.SolidMembers != nil {
		for i := range x.SolidMembers.Solid {
			s, err := solidFromXML(&x.SolidMembers.Solid[i], dim, r.resolver)
			if err != nil {
				return core.Geometry{}, fmt.Errorf("gml: MultiSolid solidMembers[%d]: %w", i, err)
			}
			solids = append(solids, s)
		}
		for i := range x.SolidMembers.CompositeSolid {
			s, err := solidFromCompositeSolid(&x.SolidMembers.CompositeSolid[i], dim, r.resolver)
			if err != nil {
				return core.Geometry{}, fmt.Errorf("gml: MultiSolid solidMembers CompositeSolid[%d]: %w", i, err)
			}
			solids = append(solids, s)
		}
	}
	return core.Geometry{Value: solids, SRSName: x.SrsName}, nil
}

// solidFromSolidMembers merges a slice of SolidPropertyType into one core.Solid.
// Used by handleCompositeSolid (XSD: CompositeSolid has "all properties of a primitive solid").
func solidFromSolidMembers(members []gen.SolidPropertyType, dim int, resolver *curveResolver) (core.Solid, error) {
	var merged core.Solid
	for i := range members {
		m := &members[i]
		var s core.Solid
		var err error
		switch {
		case m.Solid != nil:
			s, err = solidFromXML(m.Solid, dim, resolver)
		case m.CompositeSolid != nil:
			s, err = solidFromCompositeSolid(m.CompositeSolid, dim, resolver)
		case m.Href != "":
			id := strings.TrimPrefix(m.Href, "#")
			var ok bool
			if s, ok = resolver.solidByID[id]; !ok {
				continue
			}
		default:
			continue
		}
		if err != nil {
			return core.Solid{}, fmt.Errorf("gml: solidMember[%d]: %w", i, err)
		}
		merged.Exterior = append(merged.Exterior, s.Exterior...)
		merged.Interior = append(merged.Interior, s.Interior...)
	}
	return merged, nil
}

// solidsFromSolidMembers collects a slice of SolidPropertyType as individual core.Solid values.
// Used by handleMultiSolid (XSD: MultiSolid is a _GeometricAggregate collection).
func solidsFromSolidMembers(members []gen.SolidPropertyType, dim int, resolver *curveResolver) (core.MultiSolid, error) {
	var result core.MultiSolid
	for i := range members {
		m := &members[i]
		var s core.Solid
		var err error
		switch {
		case m.Solid != nil:
			s, err = solidFromXML(m.Solid, dim, resolver)
		case m.CompositeSolid != nil:
			s, err = solidFromCompositeSolid(m.CompositeSolid, dim, resolver)
		case m.Href != "":
			id := strings.TrimPrefix(m.Href, "#")
			var ok bool
			if s, ok = resolver.solidByID[id]; !ok {
				continue
			}
		default:
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("gml: solidMember[%d]: %w", i, err)
		}
		result = append(result, s)
	}
	return result, nil
}

// solidFromCompositeSolid recursively resolves a CompositeSolidType into a core.Solid.
func solidFromCompositeSolid(x *gen.CompositeSolidType, dim int, resolver *curveResolver) (core.Solid, error) {
	return solidFromSolidMembers(x.SolidMember, dim, resolver)
}

func solidFromXML(x *gen.SolidType, dim int, resolver *curveResolver) (core.Solid, error) {
	var s core.Solid
	if x.Exterior != nil {
		mp, err := multiPolygonFromSurfaceProperty(x.Exterior, dim, resolver)
		if err != nil {
			return core.Solid{}, fmt.Errorf("gml: Solid exterior: %w", err)
		}
		s.Exterior = mp
	}
	for i := range x.Interior {
		mp, err := multiPolygonFromSurfaceProperty(&x.Interior[i], dim, resolver)
		if err != nil {
			return core.Solid{}, fmt.Errorf("gml: Solid interior[%d]: %w", i, err)
		}
		s.Interior = append(s.Interior, mp)
	}
	return s, nil
}
