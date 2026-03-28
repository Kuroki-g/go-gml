package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

// handleSolid decodes a gml:Solid and returns a core.Solid.
func (r *Reader) handleSolid(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.SolidType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Solid: %w", err)
	}
	dim := derefDim(x.SrsDimension)
	solid, err := solidFromXML(&x, dim, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: solid, SRSName: x.SrsName}, nil
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
