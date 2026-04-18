package core

// nonEPSGDimTable maps known non-EPSG srsName strings to their spatial dimension.
// Covers Japanese government data formats that do not embed EPSG codes in the URI.
var nonEPSGDimTable = map[string]uint{
	"JGD2000 / (B, L)":  2,
	"JGD2000/(B,L)":     2,
	"JGD2011 / (B, L)":  2,
	"JGD2011/(B,L)":     2,
	"TD / (B, L)":       2,
	"TD/(B,L)":          2,
	"fguuid:jgd2000.bl": 2,
	"fguuid:jgd2011.bl": 2,
	"fguuid:jgd2024.bl": 2,
}
