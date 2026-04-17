# CLAUDE.md

Guide for Claude Code when working on this repository.

## [Most Important Principle] Read the XSD Before Implementing

**Always verify the XSD before implementing, modifying, or making any decisions.**
!!! Recurring problem: only checking v3.1.1 and missing other versions. When a missing implementation is found for one version, verify ALL versions!!

- v2.1.2 → `docs/schemas/gml/2.1.2/`
- v3.1.1 → `docs/schemas/gml/3.1.1/`
- v3.2.1 → `docs/schemas/gml/3.2.1/`
- CityGML → `docs/schemas/citygml/` (1.0 / 2.0 / 3.0)
- CRS/srsName → `SRSReferenceGroup` in GML XSD

**Anti-patterns:** declaring "non-compliant" without reading XSD / hardcoded string mappings / ad-hoc symptom fixes.

---

## Project

**go-gml** — Pure Go GML (ISO 19136) parser SDK for Japanese government data (KSJ, PLATEAU, etc.). Core policy: XSD compliance.

**Reference specs:**
- CityGML 2.0: OGC 12-019 (`testdata/12-019_OGC_City_Geography_Markup_Language_CityGML_Encoding_Standard.pdf`) — LoD: §6.2, Building: §10.3
- CityGML XSD: `docs/schemas/citygml/` (v1.0 / v2.0 / v3.0)

### Milestones

**GML Parser**

| Level | Status |
|---|---|
| **SF-0** | ✓ Done — Point / LineString / Polygon / Multi* / Envelope |
| **SF-1** | ✓ Done — Curve / Surface / OrientableCurve / CompositeCurve / CompositeSurface / OrientableSurface |
| **SF-2** | Not implemented — Arc / Circle interpolation (low priority) |

**CityGML 2.0**

| LoD | Status |
|---|---|
| **LoD0** | ✓ Done — `lod0FootPrint` / `lod0RoofEdge` (MultiSurface) |
| **LoD1** | ✓ Done — `bldg:lod1Solid` → `gml:Solid` |
| **LoD2** | ✓ Done — `lod2Solid` / `lod2MultiSurface` / `lod2MultiCurve` / `lod2TerrainIntersection` / `boundedBy` |
| **LoD3** | Not implemented — `bldg:lod3Solid` |

### Unimplemented / Remaining Issues

All tracked in `docs/issues/`. See files there for details.

---

## Code Structure

Multi-module monorepo managed locally with `go.work`.

**Module layout (gml* / citygml* common):**
- `reader.go` — `NewReader()` / `Next()` entry points
- `generated/` — xsd2go-lite generated structs. Do not write by hand.
- `internal/decode_*.go` — per-element decode implementations
- `internal/reader.go` switch — index of supported elements

### Module Dependencies

| Module | Depends on |
|---|---|
| `core` | none |
| `gml2_1_2` / `gml3_1_1` / `gml3_2_1` | core |
| `gml` | core + gml2_1_2 + gml3_1_1 + gml3_2_1 |
| `citygml2_0` | core + gml3_1_1 |
| `cmd/gml-parser` | gml |
| `cmd/citygml-parser` | citygml2_0 + gml3_1_1 |
| `citygml3_0` (future) | core + gml3_2_1 |

### Naming Conventions

- Generated packages (`v3_2_1`, `v3_1_1`, `v2_1_2`): underscores OK
- Handwritten parser packages (`gml3_1_1`, `gml3_2_1`): no underscores

---

## Release Procedure

See [`docs/release.md`](../docs/release.md).

---

## Development Commands

Always run via Makefile (`go test` / `go run` direct fails on WSL `/tmp noexec`). See `make help`.

When PDF analysis is needed: use `pypdf` via `uv` — get user approval first.

---

## File Exploration Rules

**Glob results may be truncated.** Never conclude "file does not exist" from truncated results.
Always start with `ls` at the top level before drilling down.

---

## Testing Policy

Run all tests with `make test`.

**Fuzz tests:** `fuzz_test.go` (`FuzzReader`) in `gml2_1_2/`, `gml3_1_1/`, `gml3_2_1/`. Regenerate seed corpus with `make fuzz-gen`. Run: `go -C gml3_2_1 test -fuzz=FuzzReader -fuzztime=60s`

---

## Implementation Rules

### Forbidden: Handwriting XML Unmarshal Structs

Never write `type xmlXxx struct { ... }` by hand inside `internal/`. Always generate via `xsd2go-lite` (`make xsd2go-gen`). Fix `xsd2go-lite` if generation fails, then regenerate.

### File Size and Structure

- **~150 lines max per file**
- New elements: add `internal/decode_*.go` + one case in `internal/reader.go` switch

### PropertyType Owning Function Naming

Functions taking `*gen.XxxPropertyType` directly must include `XxxProperty` in their name.

- Use `FromXxxProperty` grep to find owning functions
- Fields must only be accessed inside the owning function
- Violations reported by `make check-coverage`

### Dependencies

- Parser modules: `encoding/xml` (stdlib) only.
- CLIs (`cmd/*`): external libraries allowed.

---

## Known Caveats

- **GML coordinates:** `gml:pos` / `gml:posList` / `gml:coordinates`; space/comma/mixed delimiters; 2D/3D
- **CRS/srsName:** `EPSG:4326` / `urn:ogc:def:crs:EPSG::4326` / `http://www.opengis.net/def/crs/EPSG/0/4326` / `urn:ogc:def:crs:EPSG:6.18:4326`
- **GML namespaces:** 2.x = `http://www.opengis.net/gml`; 3.x = `http://www.opengis.net/gml/3.2`
- **cobra-cli:** conflicts with go.work → use `GOWORK=off cobra-cli init`
