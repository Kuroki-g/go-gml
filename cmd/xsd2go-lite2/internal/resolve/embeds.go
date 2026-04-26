package resolve

import "xsd2go-lite2/internal/parse"

// appendEmbed appends ref to embeds unless an entry with the same NS+XSDName already exists.
func appendEmbed(embeds []parse.EmbedRef, ref parse.EmbedRef) []parse.EmbedRef {
	for _, e := range embeds {
		if e.NS == ref.NS && e.XSDName == ref.XSDName {
			return embeds
		}
	}
	return append(embeds, ref)
}

// copyBaseEmbeds returns a copy of the resolved base type's Embeds.
// The base type must already be resolved before this is called.
func (r *Resolver) copyBaseEmbeds(base, schemaNS string) []parse.EmbedRef {
	baseNS, baseName := r.resolveQName(base, schemaNS)
	baseCT, ok := r.allComplexTypes[baseNS+" "+baseName]
	if !ok || len(baseCT.Embeds) == 0 {
		return nil
	}
	result := make([]parse.EmbedRef, len(baseCT.Embeds))
	copy(result, baseCT.Embeds)
	return result
}

// addGroupEmbeds records top-level model group refs from ct.Content into ct.Embeds.
func (r *Resolver) addGroupEmbeds(ct *parse.ComplexType) {
	if ct.Content == nil {
		return
	}
	for _, gr := range ct.Content.Groups {
		groupNS, groupName := r.resolveQName(gr.Ref, ct.Source)
		ct.Embeds = appendEmbed(ct.Embeds, parse.EmbedRef{XSDName: groupName, NS: groupNS, Kind: "group"})
	}
}
