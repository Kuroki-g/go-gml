package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"iter"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
)

// Feature represents a decoded GML 2.1.2 AbstractFeatureType instance.
type Feature struct {
	XMLName     xml.Name
	Fid         *string
	Description *string
	Name        *string
	BoundedBy   *core.Bound
}

// Features returns an iterator over Feature elements in the document.
// GML namespace elements (featureMember, FeatureCollection, etc.) are traversed transparently.
// Any non-GML namespace element is yielded as a Feature.
func (r *Reader) Features() iter.Seq2[Feature, error] {
	return func(yield func(Feature, error) bool) {
		if !r.prescanned {
			if err := r.runPreScan(); err != nil {
				yield(Feature{}, err)
				return
			}
		}
		scanFeatures(r.dec, yield)
	}
}

func scanFeatures(dec *xml.Decoder, yield func(Feature, error) bool) {
	for {
		tok, err := dec.Token()
		if err != nil {
			if err != io.EOF {
				yield(Feature{}, err)
			}
			return
		}
		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		if isGMLNS(se.Name.Space) {
			continue
		}
		f, err := decodeFeatureElement(dec, se)
		if !yield(f, err) {
			return
		}
		if err != nil {
			return
		}
	}
}

func decodeFeatureElement(dec *xml.Decoder, se xml.StartElement) (Feature, error) {
	f := Feature{XMLName: se.Name}
	for _, a := range se.Attr {
		if a.Name.Local == "fid" {
			s := a.Value
			f.Fid = &s
		}
	}
	for {
		tok, err := dec.Token()
		if err != nil {
			return f, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if err := consumeFeatureChild(dec, t, &f); err != nil {
				return f, err
			}
		case xml.EndElement:
			return f, nil
		}
	}
}

func consumeFeatureChild(dec *xml.Decoder, se xml.StartElement, f *Feature) error {
	if !isGMLNS(se.Name.Space) {
		return dec.Skip()
	}
	switch se.Name.Local {
	case "description":
		var s string
		if err := dec.DecodeElement(&s, &se); err != nil {
			return fmt.Errorf("gml: Feature description: %w", err)
		}
		f.Description = &s
	case "name":
		var s string
		if err := dec.DecodeElement(&s, &se); err != nil {
			return fmt.Errorf("gml: Feature name: %w", err)
		}
		f.Name = &s
	case "boundedBy":
		b, err := decodeBoundedByElement(dec)
		if err != nil {
			return err
		}
		f.BoundedBy = b
	default:
		return dec.Skip()
	}
	return nil
}

// decodeBoundedByElement reads the content of a <gml:boundedBy> that has already
// been opened. Returns nil when the element contains a null instead of a Box.
func decodeBoundedByElement(dec *xml.Decoder) (*core.Bound, error) {
	var bound *core.Bound
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if isGMLNS(t.Name.Space) && t.Name.Local == "Box" {
				var x gen.BoxType
				if err := dec.DecodeElement(&x, &t); err != nil {
					return nil, fmt.Errorf("gml: boundedBy Box: %w", err)
				}
				b, err := boundFromBoxXML(&x)
				if err != nil {
					return nil, err
				}
				bound = &b
			} else {
				if err := dec.Skip(); err != nil {
					return nil, err
				}
			}
		case xml.EndElement:
			return bound, nil
		}
	}
}
