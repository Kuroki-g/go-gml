// Package gml provides a pure Go parser for GML (ISO 19136 / OGC GML 3.2.1).
//
// It is designed to read real-world Japanese government GML datasets such as
// National Land Numerical Information (国土数値情報), Fundamental Geospatial Data
// (基盤地図情報), and PLATEAU city models.
//
// # Basic usage
//
// Create a [Reader] from any [io.ReadSeeker] (e.g. *os.File) and call [Reader.Next]
// in a loop until [io.EOF]:
//
//	r := gml.NewReader(f)
//	for {
//	    g, err := r.Next()
//	    if err == io.EOF {
//	        break
//	    }
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    switch v := g.Value.(type) {
//	    case gml.Polygon:
//	        // ...
//	    case gml.LineString:
//	        // ...
//	    }
//	}
//
// # Geometry types
//
// All GML geometry elements are decoded into one of the following Go types:
//   - [Point]
//   - [LineString]
//   - [Polygon]
//   - [MultiPoint]
//   - [MultiLineString]
//   - [MultiPolygon]
//   - [Bound]
//
// Complex GML elements (gml:Curve, gml:Surface, gml:CompositeCurve, etc.) are
// transparently flattened into the above types.
//
// # CRS
//
// The srsName attribute of each geometry is available via [Geometry].SRSName.
// Use [EPSGFromSRSName] to extract the numeric EPSG code from the various URI
// formats used in practice (e.g. "EPSG:6668", "urn:ogc:def:crs:EPSG::6668").
package gml
