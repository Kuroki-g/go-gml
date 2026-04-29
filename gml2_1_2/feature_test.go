package gml2_1_2_test

import (
	"strings"
	"testing"

	gml "github.com/Kuroki-g/go-gml/gml2_1_2"
)

func TestFeatures_basic(t *testing.T) {
	const src = `<gml:FeatureCollection xmlns:gml="http://www.opengis.net/gml" xmlns:app="urn:app">
		<gml:featureMember>
			<app:City fid="F1">
				<gml:name>Tokyo</gml:name>
				<gml:description>Capital</gml:description>
			</app:City>
		</gml:featureMember>
	</gml:FeatureCollection>`
	r := gml.NewReader(strings.NewReader(src))
	var features []gml.Feature
	for f, err := range r.Features() {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		features = append(features, f)
	}
	if len(features) != 1 {
		t.Fatalf("expected 1 feature, got %d", len(features))
	}
	f := features[0]
	if f.Fid == nil || *f.Fid != "F1" {
		t.Errorf("Fid: got %v, want F1", f.Fid)
	}
	if f.Name == nil || *f.Name != "Tokyo" {
		t.Errorf("Name: got %v, want Tokyo", f.Name)
	}
	if f.Description == nil || *f.Description != "Capital" {
		t.Errorf("Description: got %v, want Capital", f.Description)
	}
}

func TestFeatures_boundedBy(t *testing.T) {
	const src = `<gml:FeatureCollection xmlns:gml="http://www.opengis.net/gml" xmlns:app="urn:app">
		<gml:featureMember>
			<app:Area>
				<gml:boundedBy>
					<gml:Box srsName="EPSG:4326">
						<gml:coordinates>139.0,35.0 140.0,36.0</gml:coordinates>
					</gml:Box>
				</gml:boundedBy>
			</app:Area>
		</gml:featureMember>
	</gml:FeatureCollection>`
	r := gml.NewReader(strings.NewReader(src))
	var features []gml.Feature
	for f, err := range r.Features() {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		features = append(features, f)
	}
	if len(features) != 1 {
		t.Fatalf("expected 1 feature, got %d", len(features))
	}
	b := features[0].BoundedBy
	if b == nil {
		t.Fatal("BoundedBy is nil")
	}
	if b.Min[0] != 139.0 || b.Min[1] != 35.0 {
		t.Errorf("Min: %v", b.Min)
	}
	if b.Max[0] != 140.0 || b.Max[1] != 36.0 {
		t.Errorf("Max: %v", b.Max)
	}
}

func TestFeatures_featureCollection(t *testing.T) {
	const src = `<gml:FeatureCollection xmlns:gml="http://www.opengis.net/gml">
		<gml:featureMember>
			<app:City xmlns:app="urn:app" fid="A">
				<gml:name>Osaka</gml:name>
			</app:City>
		</gml:featureMember>
		<gml:featureMember>
			<app:City xmlns:app="urn:app" fid="B">
				<gml:name>Kyoto</gml:name>
			</app:City>
		</gml:featureMember>
	</gml:FeatureCollection>`
	r := gml.NewReader(strings.NewReader(src))
	var features []gml.Feature
	for f, err := range r.Features() {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		features = append(features, f)
	}
	if len(features) != 2 {
		t.Fatalf("expected 2 features, got %d", len(features))
	}
	if features[0].Name == nil || *features[0].Name != "Osaka" {
		t.Errorf("features[0].Name: got %v", features[0].Name)
	}
	if features[1].Name == nil || *features[1].Name != "Kyoto" {
		t.Errorf("features[1].Name: got %v", features[1].Name)
	}
}

func TestFeatures_noGMLProperties(t *testing.T) {
	const src = `<gml:FeatureCollection xmlns:gml="http://www.opengis.net/gml" xmlns:app="urn:app">
		<gml:featureMember>
			<app:Thing>
				<app:value>42</app:value>
			</app:Thing>
		</gml:featureMember>
	</gml:FeatureCollection>`
	r := gml.NewReader(strings.NewReader(src))
	var features []gml.Feature
	for f, err := range r.Features() {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		features = append(features, f)
	}
	if len(features) != 1 {
		t.Fatalf("expected 1 feature, got %d", len(features))
	}
	if features[0].XMLName.Local != "Thing" {
		t.Errorf("XMLName.Local: got %q", features[0].XMLName.Local)
	}
}
