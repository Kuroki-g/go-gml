# go-gml

Pure Go implementation of a GML (ISO 19136) parser. Reads real Japanese government GML data — 国土数値情報 (KSJ), 基盤地図情報, and PLATEAU CityGML.

## Install CLI

```bash
go install github.com/Kuroki-g/go-gml/cmd/gml-parser@latest
```

```bash
# Convert GML to GeoJSON
gml-parser convert -i data.gml --swap > out.geojson

# Show geometry statistics
gml-parser inspect -i data.gml
```

`--swap` is typically needed for 国土数値情報 files, which store coordinates as lat/lon (GeoJSON requires lon/lat).

## Use as a library

```bash
go get github.com/Kuroki-g/go-gml/gml
```

```go
import "github.com/Kuroki-g/go-gml/gml"

f, _ := os.Open("data.gml")
r := gml.NewReader321(f)
for {
    g, err := r.Next()
    if errors.Is(err, io.EOF) {
        break
    }
    fmt.Println(g.Value) // Point, LineString, Polygon, ...
}
```

## Supported geometry (GML 3.2.1)

| Element | Output type | Status |
|---|---|---|
| `gml:Point` | `Point` | ✓ |
| `gml:LineString` | `LineString` | ✓ |
| `gml:Polygon` | `Polygon` | ✓ |
| `gml:MultiPoint` | `MultiPoint` | ✓ |
| `gml:MultiCurve` / `gml:MultiLineString` | `MultiLineString` | ✓ |
| `gml:MultiSurface` / `gml:MultiPolygon` | `MultiPolygon` | ✓ |
| `gml:Curve` + `LineStringSegment` | `LineString` | ✓ |
| `gml:Surface` + `PolygonPatch` | `Polygon` | ✓ |
| `gml:CompositeCurve` / `gml:OrientableCurve` | `LineString` | ✓ |
| `gml:CompositeSurface` / `gml:OrientableSurface` | `Polygon` | ✓ |
| `gml:Grid` / `gml:RectifiedGrid` | `GridCoverage` | ✓ |
| `gml:Solid` | `Solid` | ✓ (GML 3.1.1) |

GML 3.1.1 and GML 2.1.2 are also supported via `gml.NewReader311` / `gml.NewReader212`.

## Modules

| Module | Purpose |
|---|---|
| `github.com/Kuroki-g/go-gml/gml` | Main entry point (re-exports all versions) |
| `github.com/Kuroki-g/go-gml/core` | Shared geometry types |
| `github.com/Kuroki-g/go-gml/gml3_2_1` | GML 3.2.1 parser |
| `github.com/Kuroki-g/go-gml/gml3_1_1` | GML 3.1.1 parser |
| `github.com/Kuroki-g/go-gml/gml2_1_2` | GML 2.1.2 parser |
| `github.com/Kuroki-g/go-gml/citygml2_0` | CityGML 2.0 parser (Building / LoD0–1) |
| `github.com/Kuroki-g/go-gml/cmd/gml-parser` | GML CLI tool |
| `github.com/Kuroki-g/go-gml/cmd/citygml-parser` | CityGML CLI tool |

## LICENSES

Apache License Version 2.0. See LICENSE file.

`docs/go/xsd2go-lite/schemas/gml` files follows:

-----------------------------------------------------------------------

Policies, Procedures, Terms, and Conditions of OGC(r) are available
  http://www.opengeospatial.org/ogc/legal/ .

OGC and OpenGIS are registered trademarks of Open Geospatial Consortium.

Copyright (c) 2012, 2018 Open Geospatial Consortium

-----------------------------------------------------------------------

`docs/go/xsd2go-lite/schemas/citygml` files follows:

-----------------------------------------------------------------------

Policies, Procedures, Terms, and Conditions of OGC(r) are available
  https://www.ogc.org/legal/ .

OGC and OpenGIS are registered trademarks of Open Geospatial Consortium.

Copyright (c) 2012, 2023 Open Geospatial Consortium

-----------------------------------------------------------------------

`docs/go/xsd2go-lite/schemas/xlink/xlink.xsd` follows:

-----------------------------------------------------------------------

This software includes material copied from or derived from
XLink Schema (https://www.w3.org/1999/xlink.xsd).
Copyright © 1999 World Wide Web Consortium.

This work is distributed under the W3C Software and Document License.
https://www.w3.org/copyright/software-license-2023/

-----------------------------------------------------------------------
