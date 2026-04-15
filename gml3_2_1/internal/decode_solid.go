package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
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

// handleCompositeSolid decodes a gml:CompositeSolid and returns a core.Solid.
// XSD: CompositeSolid is a AbstractSolid substitutionGroup member with
// "all the geometric properties of a (primitive) solid".
func (r *Reader) handleCompositeSolid(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.CompositeSolidType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: CompositeSolid: %w", err)
	}
	dim := preferDim(derefDim(x.SrsDimension), r.globalDim)
	solid, err := solidFromSolidPropertyMembers(x.SolidMember, dim, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: solid, SRSName: x.SrsName}, nil
}

// handleMultiSolid decodes a gml:MultiSolid and returns a core.MultiSolid.
// XSD: MultiSolid is substitutionGroup="AbstractGeometricAggregate", not AbstractSolid.
func (r *Reader) handleMultiSolid(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.MultiSolidType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: MultiSolid: %w", err)
	}
	dim := preferDim(derefDim(x.SrsDimension), r.globalDim)
	solids, err := solidsFromSolidPropertyMembers(x.SolidMember, dim, r.resolver)
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

func solidFromXML(x *gen.SolidType, dim uint, resolver *curveResolver) (core.Solid, error) {
	var s core.Solid
	if x.Exterior != nil && x.Exterior.Shell != nil {
		mp, err := multiPolygonFromShell(x.Exterior.Shell, dim, resolver)
		if err != nil {
			return core.Solid{}, fmt.Errorf("gml: Solid exterior: %w", err)
		}
		s.Exterior = mp
	}
	for i := range x.Interior {
		if x.Interior[i].Shell == nil {
			continue
		}
		mp, err := multiPolygonFromShell(x.Interior[i].Shell, dim, resolver)
		if err != nil {
			return core.Solid{}, fmt.Errorf("gml: Solid interior[%d]: %w", i, err)
		}
		s.Interior = append(s.Interior, mp)
	}
	return s, nil
}

// multiPolygonFromShell converts a ShellType (set of surface members) to a MultiPolygon.
func multiPolygonFromShell(shell *gen.ShellType, dim uint, resolver *curveResolver) (core.MultiPolygon, error) {
	var result core.MultiPolygon
	for i := range shell.SurfaceMember {
		mp, err := multiPolygonFromSurfaceProperty(&shell.SurfaceMember[i], dim, resolver)
		if err != nil {
			return nil, fmt.Errorf("gml: Shell surfaceMember[%d]: %w", i, err)
		}
		result = append(result, mp...)
	}
	return result, nil
}

// solidFromSolidPropertyMembers merges SolidPropertyType members into one core.Solid.
// Used by handleCompositeSolid (XSD: CompositeSolid has "all properties of a primitive solid").
func solidFromSolidPropertyMembers(members []gen.SolidPropertyType, dim uint, resolver *curveResolver) (core.Solid, error) {
	var merged core.Solid
	for i := range members {
		m := &members[i]
		if m.NilReason != nil {
			continue
		}
		// xlink metadata attributes — not used for solid member resolution.
		_ = m.RemoteSchema
		_ = m.TypeField
		_ = m.Role
		_ = m.Arcrole
		_ = m.Title
		_ = m.Show
		_ = m.Actuate
		_ = m.Owns
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
				return core.Solid{}, fmt.Errorf("gml: unresolved xlink:href %q", m.Href)
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

// solidsFromSolidPropertyMembers collects SolidPropertyType members as individual core.Solid values.
// Used by handleMultiSolid (XSD: MultiSolid is a _GeometricAggregate collection).
func solidsFromSolidPropertyMembers(members []gen.SolidPropertyType, dim uint, resolver *curveResolver) (core.MultiSolid, error) {
	var result core.MultiSolid
	for i := range members {
		m := &members[i]
		if m.NilReason != nil {
			continue
		}
		// xlink metadata attributes — not used for solid member resolution.
		_ = m.RemoteSchema
		_ = m.TypeField
		_ = m.Role
		_ = m.Arcrole
		_ = m.Title
		_ = m.Show
		_ = m.Actuate
		_ = m.Owns
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
				return nil, fmt.Errorf("gml: unresolved xlink:href %q", m.Href)
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
func solidFromCompositeSolid(x *gen.CompositeSolidType, dim uint, resolver *curveResolver) (core.Solid, error) {
	return solidFromSolidPropertyMembers(x.SolidMember, dim, resolver)
}
