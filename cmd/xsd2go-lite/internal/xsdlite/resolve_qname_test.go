package xsdlite

import "testing"

func TestResolver_resolveQName_noPrefix(t *testing.T) {
	r := NewResolver()
	ns, name := r.resolveQName("MyType", "http://example.com")
	if ns != "http://example.com" {
		t.Errorf("ns = %q, want %q", ns, "http://example.com")
	}
	if name != "MyType" {
		t.Errorf("name = %q, want %q", name, "MyType")
	}
}

func TestResolver_resolveQName_withPrefix(t *testing.T) {
	r := NewResolver()
	r.nsMaps["http://example.com"] = map[string]string{
		"gml": "http://www.opengis.net/gml/3.2",
	}
	ns, name := r.resolveQName("gml:PointType", "http://example.com")
	if ns != "http://www.opengis.net/gml/3.2" {
		t.Errorf("ns = %q, want GML namespace", ns)
	}
	if name != "PointType" {
		t.Errorf("name = %q, want PointType", name)
	}
}

func TestResolver_resolveQName_prefixDeterminism(t *testing.T) {
	r := NewResolver()
	r.nsMaps["http://schema-a.com"] = map[string]string{"ns": "http://ns-a.com"}
	r.nsMaps["http://schema-b.com"] = map[string]string{"ns": "http://ns-b.com"}

	ns, _ := r.resolveQName("ns:Foo", "http://schema-a.com")
	if ns != "http://ns-a.com" {
		t.Errorf("resolveQName should prefer source schema's NSMap, got %q", ns)
	}

	ns, _ = r.resolveQName("ns:Bar", "http://schema-b.com")
	if ns != "http://ns-b.com" {
		t.Errorf("resolveQName should prefer source schema's NSMap, got %q", ns)
	}
}

func TestResolver_resolveQName_unknownPrefix(t *testing.T) {
	r := NewResolver()
	ns, name := r.resolveQName("xyz:Foo", "http://fallback")
	if ns != "http://fallback" {
		t.Errorf("unknown prefix should fall back to schemaNS, got %q", ns)
	}
	if name != "Foo" {
		t.Errorf("name = %q, want Foo", name)
	}
}
