package gml_test

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Kuroki-g/go-gml/gml"
)

// ExampleNewReader321 demonstrates streaming GML 3.2.1 geometry from a file.
// It counts each geometry type encountered in a 国土数値情報 N03 file.
func ExampleNewReader321() {
	f, err := os.Open("../testdata/N03/N03-20240101_13.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := gml.NewReader321(f)
	var polygons, others int
	for {
		g, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		switch g.Value.(type) {
		case gml.Polygon:
			polygons++
		default:
			others++
		}
	}
	fmt.Printf("polygons=%d others=%d\n", polygons, others)
}

// ExampleEPSGFromSRSName demonstrates extracting EPSG codes from srsName strings.
func ExampleEPSGFromSRSName() {
	formats := []string{
		"EPSG:4326",
		"urn:ogc:def:crs:EPSG::6668",
		"http://www.opengis.net/def/crs/EPSG/0/6697",
		"urn:ogc:def:crs:EPSG:6.18:4326",
	}
	for _, s := range formats {
		fmt.Printf("%s -> %d\n", s, gml.EPSGFromSRSName(s))
	}
	// Output:
	// EPSG:4326 -> 4326
	// urn:ogc:def:crs:EPSG::6668 -> 6668
	// http://www.opengis.net/def/crs/EPSG/0/6697 -> 6697
	// urn:ogc:def:crs:EPSG:6.18:4326 -> 4326
}
