package citygml2_0

import (
	"encoding/xml"

	"github.com/Kuroki-g/go-gml/core"
)

// decodeBoundedBy decodes a <bldg:boundedBy> property element and returns
// the concrete BoundarySurface it wraps.
func decodeBoundedBy(dec *xml.Decoder, _ xml.StartElement, gmlDec core.Decoder) (*BoundarySurface, error) {
	var bs *BoundarySurface
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.EndElement:
			return bs, nil
		case xml.StartElement:
			if t.Name.Space != nsBldg {
				if err := dec.Skip(); err != nil {
					return nil, err
				}
				continue
			}
			switch t.Name.Local {
			case SurfaceTypeRoofSurface,
				SurfaceTypeWallSurface,
				SurfaceTypeGroundSurface,
				SurfaceTypeClosureSurface,
				SurfaceTypeOuterFloorSurface,
				SurfaceTypeOuterCeilingSurface,
				SurfaceTypeInteriorWallSurface,
				SurfaceTypeFloorSurface,
				SurfaceTypeCeilingSurface:
				var decErr error
				bs, decErr = decodeBoundarySurface(dec, t, gmlDec)
				if decErr != nil {
					return nil, decErr
				}
			default:
				if err := dec.Skip(); err != nil {
					return nil, err
				}
			}
		}
	}
}

// decodeBoundarySurface decodes a concrete BoundarySurface element
// (e.g. <bldg:WallSurface>) and its lod2MultiSurface geometry.
func decodeBoundarySurface(dec *xml.Decoder, se xml.StartElement, gmlDec core.Decoder) (*BoundarySurface, error) {
	bs := &BoundarySurface{
		ID:          gmlID(se),
		SurfaceType: se.Name.Local,
	}
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.EndElement:
			return bs, nil
		case xml.StartElement:
			if t.Name.Space != nsBldg {
				if err := dec.Skip(); err != nil {
					return nil, err
				}
				continue
			}
			switch t.Name.Local {
			case "lod2MultiSurface":
				g, err := decodeGeometryProp(dec, t, gmlDec)
				if err != nil {
					return nil, err
				}
				bs.Lod2MultiSurface = g
			default:
				if err := dec.Skip(); err != nil {
					return nil, err
				}
			}
		}
	}
}
