package citygml2_0

import "github.com/Kuroki-g/go-gml/core"

// Building represents a CityGML 2.0 bldg:Building city object.
type Building struct {
	ID            string
	Lod0FootPrint *core.Geometry
	Lod0RoofEdge  *core.Geometry
	Lod1Solid     *core.Geometry
}
