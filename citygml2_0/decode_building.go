package citygml2_0

import (
	"encoding/xml"

	intl "github.com/Kuroki-g/go-gml/citygml2_0/internal"
	"github.com/Kuroki-g/go-gml/core"
)

const (
	nsBldg = "http://www.opengis.net/citygml/building/2.0"
	nsGML  = "http://www.opengis.net/gml"
)

func gmlID(se xml.StartElement) string {
	for _, a := range se.Attr {
		if a.Name.Local == "id" && a.Name.Space == nsGML {
			return a.Value
		}
	}
	return ""
}

func decodeBuilding(dec *xml.Decoder, se xml.StartElement, gmlDec core.Decoder) (*Building, error) {
	b := &Building{ID: gmlID(se)}
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.EndElement:
			return b, nil
		case xml.StartElement:
			if t.Name.Space != nsBldg {
				if err := dec.Skip(); err != nil {
					return nil, err
				}
				continue
			}
			switch t.Name.Local {
			case "lod0FootPrint":
				g, err := decodeGeometryProp(dec, gmlDec)
				if err != nil {
					return nil, err
				}
				b.Lod0FootPrint = g
			case "lod0RoofEdge":
				g, err := decodeGeometryProp(dec, gmlDec)
				if err != nil {
					return nil, err
				}
				b.Lod0RoofEdge = g
			case "lod1Solid":
				g, err := decodeGeometryProp(dec, gmlDec)
				if err != nil {
					return nil, err
				}
				b.Lod1Solid = g
			case "lod1MultiSurface":
				g, err := decodeGeometryProp(dec, gmlDec)
				if err != nil {
					return nil, err
				}
				b.Lod1MultiSurface = g
			case "lod1TerrainIntersection":
				g, err := decodeGeometryProp(dec, gmlDec)
				if err != nil {
					return nil, err
				}
				b.Lod1TerrainIntersection = g
			case "lod2Solid":
				g, err := decodeGeometryProp(dec, gmlDec)
				if err != nil {
					return nil, err
				}
				b.Lod2Solid = g
			case "lod2MultiSurface":
				g, err := decodeGeometryProp(dec, gmlDec)
				if err != nil {
					return nil, err
				}
				b.Lod2MultiSurface = g
			case "lod2MultiCurve":
				g, err := decodeGeometryProp(dec, gmlDec)
				if err != nil {
					return nil, err
				}
				b.Lod2MultiCurve = g
			case "lod2TerrainIntersection":
				g, err := decodeGeometryProp(dec, gmlDec)
				if err != nil {
					return nil, err
				}
				b.Lod2TerrainIntersection = g
			default:
				if err := dec.Skip(); err != nil {
					return nil, err
				}
			}
		}
	}
}

// decodeGeometryProp decodes the first child GML element of a property wrapper
// and consumes through the property's EndElement.
func decodeGeometryProp(dec *xml.Decoder, gmlDec core.Decoder) (*core.Geometry, error) {
	var result *core.Geometry
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if result == nil {
				sub, err := intl.ExtractSubtree(dec, t)
				if err != nil {
					return nil, err
				}
				g, err := gmlDec.Decode(sub)
				if err != nil {
					return nil, err
				}
				result = &g
			} else {
				if err := dec.Skip(); err != nil {
					return nil, err
				}
			}
		case xml.EndElement:
			return result, nil
		}
	}
}
