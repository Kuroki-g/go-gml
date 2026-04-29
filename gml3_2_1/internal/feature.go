package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"iter"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

// Feature holds the GML AbstractFeatureType properties parsed from a concrete feature element.
type Feature struct {
	ElementName          xml.Name
	ID                   *string
	Description          *gen.StringOrRefType
	DescriptionReference *gen.ReferenceType
	Identifier           *gen.CodeWithAuthorityType
	Names                []gen.CodeType
	BoundedBy            *core.Bound
	// gml:location is deprecated in GML 3.2.1; not implemented.
}

// hasGMLProperties reports whether any GML AbstractFeatureType property was found.
// Elements with no GML properties (e.g. plain application property elements) are skipped.
func (f *Feature) hasGMLProperties() bool {
	return f.ID != nil || f.Description != nil || f.DescriptionReference != nil ||
		f.Identifier != nil || len(f.Names) > 0 || f.BoundedBy != nil
}

// Features returns an iterator that yields one Feature per non-GML element found in the stream
// that carries at least one GML AbstractFeatureType property (gml:id, gml:description,
// gml:identifier, gml:name, or gml:boundedBy). Plain application property elements with no
// GML content are silently skipped.
//
// Do not call Features and Next/Geometries on the same Reader concurrently or interleaved;
// they share the underlying decoder and will corrupt each other's stream position.
func (r *Reader) Features() iter.Seq2[*Feature, error] {
	return func(yield func(*Feature, error) bool) {
		stack := make([]*Feature, 0, 8)
		for {
			tok, err := r.dec.Token()
			if err == io.EOF {
				return
			}
			if err != nil {
				yield(nil, fmt.Errorf("gml: Features: %w", err))
				return
			}
			switch t := tok.(type) {
			case xml.StartElement:
				if !isGMLNS(t.Name.Space) {
					f := &Feature{ElementName: t.Name}
					if id := extractGMLID(t); id != "" {
						f.ID = &id
					}
					stack = append(stack, f)
					continue
				}
				if len(stack) == 0 {
					if err := r.dec.Skip(); err != nil {
						yield(nil, fmt.Errorf("gml: Features: skip: %w", err))
						return
					}
					continue
				}
				cur := stack[len(stack)-1]
				if err := r.handleFeatureProperty(cur, t); err != nil {
					yield(nil, err)
					return
				}
			case xml.EndElement:
				if !isGMLNS(t.Name.Space) && len(stack) > 0 {
					f := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					if f.hasGMLProperties() {
						if !yield(f, nil) {
							return
						}
					}
				}
			}
		}
	}
}

func (r *Reader) handleFeatureProperty(cur *Feature, t xml.StartElement) error {
	switch t.Name.Local {
	case gmlMetaDataProperty:
		return r.dec.Skip()
	case gmlDescription:
		var v gen.StringOrRefType
		if err := r.dec.DecodeElement(&v, &t); err != nil {
			return fmt.Errorf("gml: Feature description: %w", err)
		}
		cur.Description = &v
	case gmlDescriptionReference:
		var v gen.ReferenceType
		if err := r.dec.DecodeElement(&v, &t); err != nil {
			return fmt.Errorf("gml: Feature descriptionReference: %w", err)
		}
		cur.DescriptionReference = &v
	case gmlIdentifier:
		var v gen.CodeWithAuthorityType
		if err := r.dec.DecodeElement(&v, &t); err != nil {
			return fmt.Errorf("gml: Feature identifier: %w", err)
		}
		cur.Identifier = &v
	case gmlName:
		var v gen.CodeType
		if err := r.dec.DecodeElement(&v, &t); err != nil {
			return fmt.Errorf("gml: Feature name: %w", err)
		}
		cur.Names = append(cur.Names, v)
	case gmlBoundedBy:
		b, err := decodeBoundedByForFeature(r.dec, t)
		if err != nil {
			return err
		}
		cur.BoundedBy = b
	default:
		if err := r.dec.Skip(); err != nil {
			return fmt.Errorf("gml: Feature skip %s: %w", t.Name.Local, err)
		}
	}
	return nil
}

func decodeBoundedByForFeature(dec *xml.Decoder, se xml.StartElement) (*core.Bound, error) {
	var bs gen.BoundingShapeType
	if err := dec.DecodeElement(&bs, &se); err != nil {
		return nil, fmt.Errorf("gml: Feature boundedBy: %w", err)
	}
	if bs.Null != nil || bs.Envelope == nil {
		return nil, nil
	}
	b, err := boundFromXML(bs.Envelope)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
