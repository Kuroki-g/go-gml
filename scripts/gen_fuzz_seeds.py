#!/usr/bin/env python3
"""Generate Go fuzz seed corpus files for gml2_1_2, gml3_1_1, and gml3_2_1.

Output directories (always regenerated, gitignored):
  gml2_1_2/testdata/fuzz/FuzzReader/
  gml3_1_1/testdata/fuzz/FuzzReader/
  gml3_2_1/testdata/fuzz/FuzzReader/

Go fuzz corpus format:
  go test fuzz v1
  []byte("...")
"""

import os
import re

REPO_ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

# ---------------------------------------------------------------------------
# Go fuzz corpus writer
# ---------------------------------------------------------------------------

def _go_quote(s: str) -> str:
    """Escape a string for Go []byte("...") corpus format."""
    out = []
    for ch in s:
        if ch == '"':
            out.append('\\"')
        elif ch == '\\':
            out.append('\\\\')
        elif ch == '\n':
            out.append('\\n')
        elif ch == '\r':
            out.append('\\r')
        elif ch == '\t':
            out.append('\\t')
        else:
            out.append(ch)
    return ''.join(out)


def write_corpus(module: str, func_name: str, seeds: list[str]) -> None:
    out_dir = os.path.join(REPO_ROOT, module, "testdata", "fuzz", func_name)
    os.makedirs(out_dir, exist_ok=True)
    for i, seed in enumerate(seeds, start=1):
        path = os.path.join(out_dir, f"{i:03d}")
        content = f'go test fuzz v1\n[]byte("{_go_quote(seed)}")\n'
        with open(path, "w", encoding="utf-8") as f:
            f.write(content)
    print(f"  {module}: {len(seeds)} seeds → {out_dir}")


# ---------------------------------------------------------------------------
# GML 2.1.2  (namespace: http://www.opengis.net/gml, coordinates/coord)
# ---------------------------------------------------------------------------

NS2 = 'xmlns:gml="http://www.opengis.net/gml"'

GML2_SEEDS = [
    # Point — coordinates
    f'<gml:Point {NS2}><gml:coordinates>139.7,35.6</gml:coordinates></gml:Point>',
    # Point — coord
    f'<gml:Point {NS2}><gml:coord><gml:X>139.7</gml:X><gml:Y>35.6</gml:Y></gml:coord></gml:Point>',
    # Point 3D
    f'<gml:Point {NS2}><gml:coordinates>139.7,35.6,10.5</gml:coordinates></gml:Point>',
    # LineString
    f'<gml:LineString {NS2}><gml:coordinates>0,0 1,1 2,0</gml:coordinates></gml:LineString>',
    # LineString — coord elements
    f'<gml:LineString {NS2}>'
    f'<gml:coord><gml:X>0</gml:X><gml:Y>0</gml:Y></gml:coord>'
    f'<gml:coord><gml:X>1</gml:X><gml:Y>1</gml:Y></gml:coord>'
    f'</gml:LineString>',
    # Polygon — outerBoundaryIs
    f'<gml:Polygon {NS2}>'
    f'<gml:outerBoundaryIs><gml:LinearRing>'
    f'<gml:coordinates>0,0 10,0 10,10 0,10 0,0</gml:coordinates>'
    f'</gml:LinearRing></gml:outerBoundaryIs>'
    f'</gml:Polygon>',
    # Polygon with hole
    f'<gml:Polygon {NS2}>'
    f'<gml:outerBoundaryIs><gml:LinearRing><gml:coordinates>0,0 10,0 10,10 0,10 0,0</gml:coordinates></gml:LinearRing></gml:outerBoundaryIs>'
    f'<gml:innerBoundaryIs><gml:LinearRing><gml:coordinates>2,2 8,2 8,8 2,8 2,2</gml:coordinates></gml:LinearRing></gml:innerBoundaryIs>'
    f'</gml:Polygon>',
    # MultiPoint
    f'<gml:MultiPoint {NS2}>'
    f'<gml:pointMember><gml:Point><gml:coordinates>0,0</gml:coordinates></gml:Point></gml:pointMember>'
    f'<gml:pointMember><gml:Point><gml:coordinates>1,1</gml:coordinates></gml:Point></gml:pointMember>'
    f'</gml:MultiPoint>',
    # MultiLineString
    f'<gml:MultiLineString {NS2}>'
    f'<gml:lineStringMember><gml:LineString><gml:coordinates>0,0 1,1</gml:coordinates></gml:LineString></gml:lineStringMember>'
    f'<gml:lineStringMember><gml:LineString><gml:coordinates>2,2 3,3</gml:coordinates></gml:LineString></gml:lineStringMember>'
    f'</gml:MultiLineString>',
    # MultiPolygon
    f'<gml:MultiPolygon {NS2}>'
    f'<gml:polygonMember><gml:Polygon>'
    f'<gml:outerBoundaryIs><gml:LinearRing><gml:coordinates>0,0 5,0 5,5 0,5 0,0</gml:coordinates></gml:LinearRing></gml:outerBoundaryIs>'
    f'</gml:Polygon></gml:polygonMember>'
    f'<gml:polygonMember><gml:Polygon>'
    f'<gml:outerBoundaryIs><gml:LinearRing><gml:coordinates>5,0 10,0 10,5 5,5 5,0</gml:coordinates></gml:LinearRing></gml:outerBoundaryIs>'
    f'</gml:Polygon></gml:polygonMember>'
    f'</gml:MultiPolygon>',
    # Box — coordinates
    f'<gml:Box {NS2} srsName="EPSG:4326"><gml:coordinates>139.0,35.0 140.0,36.0</gml:coordinates></gml:Box>',
    # Box — coord
    f'<gml:Box {NS2}>'
    f'<gml:coord><gml:X>139.0</gml:X><gml:Y>35.0</gml:Y></gml:coord>'
    f'<gml:coord><gml:X>140.0</gml:X><gml:Y>36.0</gml:Y></gml:coord>'
    f'</gml:Box>',
    # Large coordinate list
    f'<gml:LineString {NS2}><gml:coordinates>'
    + ' '.join(f'{i},{i}' for i in range(100))
    + f'</gml:coordinates></gml:LineString>',
    # Mixed geometry stream
    f'<root {NS2}>'
    f'<gml:Point><gml:coordinates>1,2</gml:coordinates></gml:Point>'
    f'<gml:LineString><gml:coordinates>0,0 1,1</gml:coordinates></gml:LineString>'
    f'</root>',
    # Empty root
    f'<root/>',
    # Empty input
    '',
    # No namespace (parser should handle gracefully)
    '<gml:Point><gml:coordinates>1,2</gml:coordinates></gml:Point>',
]

# ---------------------------------------------------------------------------
# GML 3.1.1  (namespace: http://www.opengis.net/gml, pos/posList)
# ---------------------------------------------------------------------------

NS311 = 'xmlns:gml="http://www.opengis.net/gml"'

GML311_SEEDS = [
    # Point
    f'<gml:Point {NS311}><gml:pos>139.7 35.6</gml:pos></gml:Point>',
    # Point 3D
    f'<gml:Point {NS311} srsDimension="3"><gml:pos>139.7 35.6 10.5</gml:pos></gml:Point>',
    # LineString — posList
    f'<gml:LineString {NS311} srsDimension="2"><gml:posList>0 0 1 1 2 0</gml:posList></gml:LineString>',
    # LineString — pos elements
    f'<gml:LineString {NS311} srsDimension="2">'
    f'<gml:pos>0 0</gml:pos><gml:pos>1 1</gml:pos>'
    f'</gml:LineString>',
    # Polygon
    f'<gml:Polygon {NS311} srsDimension="2">'
    f'<gml:exterior><gml:LinearRing><gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList></gml:LinearRing></gml:exterior>'
    f'</gml:Polygon>',
    # Polygon with hole
    f'<gml:Polygon {NS311} srsDimension="2">'
    f'<gml:exterior><gml:LinearRing><gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList></gml:LinearRing></gml:exterior>'
    f'<gml:interior><gml:LinearRing><gml:posList>2 2 8 2 8 8 2 8 2 2</gml:posList></gml:LinearRing></gml:interior>'
    f'</gml:Polygon>',
    # MultiSurface
    f'<gml:MultiSurface {NS311} srsDimension="2">'
    f'<gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing><gml:posList>0 0 5 0 5 5 0 5 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember>'
    f'</gml:MultiSurface>',
    # Envelope
    f'<gml:Envelope {NS311} srsName="EPSG:4326" srsDimension="2">'
    f'<gml:lowerCorner>139.0 35.0</gml:lowerCorner>'
    f'<gml:upperCorner>140.0 36.0</gml:upperCorner>'
    f'</gml:Envelope>',
    # Large coordinate list
    f'<gml:Polygon {NS311} srsDimension="2">'
    f'<gml:exterior><gml:LinearRing><gml:posList>'
    + ' '.join(f'{i} {i}' for i in range(50)) + ' 0 0'
    + f'</gml:posList></gml:LinearRing></gml:exterior>'
    f'</gml:Polygon>',
    # Mixed stream
    f'<root {NS311}>'
    f'<gml:Point><gml:pos>1 2</gml:pos></gml:Point>'
    f'<gml:LineString srsDimension="2"><gml:posList>0 0 1 1</gml:posList></gml:LineString>'
    f'</root>',
    # Empty root
    f'<root/>',
    # Empty input
    '',
]

# ---------------------------------------------------------------------------
# GML 3.2.1  (namespace: http://www.opengis.net/gml/3.2, pos/posList + xlink)
# ---------------------------------------------------------------------------

NS32 = 'xmlns:gml="http://www.opengis.net/gml/3.2"'
XLINK = 'xmlns:xlink="http://www.w3.org/1999/xlink"'

GML32_SEEDS = [
    # Point
    f'<gml:Point {NS32}><gml:pos>139.7 35.6</gml:pos></gml:Point>',
    # Point 3D
    f'<gml:Point {NS32} srsDimension="3"><gml:pos>139.7 35.6 10.5</gml:pos></gml:Point>',
    # LineString — posList
    f'<gml:LineString {NS32} srsDimension="2"><gml:posList>0 0 1 1 2 0</gml:posList></gml:LineString>',
    # LineString — pos elements
    f'<gml:LineString {NS32} srsDimension="2">'
    f'<gml:pos>0 0</gml:pos><gml:pos>1 1</gml:pos>'
    f'</gml:LineString>',
    # Curve / LineStringSegment
    f'<gml:Curve {NS32} srsDimension="2">'
    f'<gml:segments><gml:LineStringSegment><gml:posList>0 0 5 5 10 0</gml:posList></gml:LineStringSegment></gml:segments>'
    f'</gml:Curve>',
    # Polygon
    f'<gml:Polygon {NS32} srsDimension="2">'
    f'<gml:exterior><gml:LinearRing><gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList></gml:LinearRing></gml:exterior>'
    f'</gml:Polygon>',
    # Polygon with hole
    f'<gml:Polygon {NS32} srsDimension="2">'
    f'<gml:exterior><gml:LinearRing><gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList></gml:LinearRing></gml:exterior>'
    f'<gml:interior><gml:LinearRing><gml:posList>2 2 8 2 8 8 2 8 2 2</gml:posList></gml:LinearRing></gml:interior>'
    f'</gml:Polygon>',
    # Surface / PolygonPatch
    f'<gml:Surface {NS32} srsDimension="2">'
    f'<gml:patches><gml:PolygonPatch>'
    f'<gml:exterior><gml:LinearRing><gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList></gml:LinearRing></gml:exterior>'
    f'</gml:PolygonPatch></gml:patches>'
    f'</gml:Surface>',
    # MultiPoint
    f'<gml:MultiPoint {NS32} srsDimension="2">'
    f'<gml:pointMember><gml:Point><gml:pos>0 0</gml:pos></gml:Point></gml:pointMember>'
    f'<gml:pointMember><gml:Point><gml:pos>1 1</gml:pos></gml:Point></gml:pointMember>'
    f'</gml:MultiPoint>',
    # MultiCurve — curveMember
    f'<gml:MultiCurve {NS32} srsDimension="2">'
    f'<gml:curveMember><gml:LineString><gml:posList>0 0 1 1</gml:posList></gml:LineString></gml:curveMember>'
    f'<gml:curveMember><gml:LineString><gml:posList>2 2 3 3</gml:posList></gml:LineString></gml:curveMember>'
    f'</gml:MultiCurve>',
    # MultiSurface — surfaceMember
    f'<gml:MultiSurface {NS32} srsDimension="2">'
    f'<gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing><gml:posList>0 0 5 0 5 5 0 5 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember>'
    f'<gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing><gml:posList>5 0 10 0 10 5 5 5 5 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember>'
    f'</gml:MultiSurface>',
    # CompositeCurve
    f'<gml:CompositeCurve {NS32} srsDimension="2">'
    f'<gml:curveMember><gml:Curve><gml:segments><gml:LineStringSegment><gml:posList>0 0 5 0</gml:posList></gml:LineStringSegment></gml:segments></gml:Curve></gml:curveMember>'
    f'<gml:curveMember><gml:Curve><gml:segments><gml:LineStringSegment><gml:posList>5 0 10 0</gml:posList></gml:LineStringSegment></gml:segments></gml:Curve></gml:curveMember>'
    f'</gml:CompositeCurve>',
    # CompositeSurface
    f'<gml:CompositeSurface {NS32} srsDimension="2">'
    f'<gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing><gml:posList>0 0 5 0 5 5 0 5 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember>'
    f'</gml:CompositeSurface>',
    # OrientableSurface
    f'<gml:OrientableSurface {NS32} orientation="+" srsDimension="2">'
    f'<gml:baseSurface><gml:Polygon><gml:exterior><gml:LinearRing><gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></gml:baseSurface>'
    f'</gml:OrientableSurface>',
    # Envelope
    f'<gml:Envelope {NS32} srsName="EPSG:6668" srsDimension="2">'
    f'<gml:lowerCorner>139.0 35.0</gml:lowerCorner>'
    f'<gml:upperCorner>140.0 36.0</gml:upperCorner>'
    f'</gml:Envelope>',
    # xlink:href — forward reference
    f'<root {NS32} {XLINK}>'
    f'<gml:CompositeSurface gml:id="CS1" srsDimension="2"><gml:surfaceMember xlink:href="#P1"/></gml:CompositeSurface>'
    f'<gml:Polygon gml:id="P1"><gml:exterior><gml:LinearRing><gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon>'
    f'</root>',
    # xlink:href — back reference
    f'<root {NS32} {XLINK}>'
    f'<gml:Polygon gml:id="P1"><gml:exterior><gml:LinearRing><gml:posList>0 0 10 0 10 10 0 10 0 0</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon>'
    f'<gml:CompositeSurface gml:id="CS1" srsDimension="2"><gml:surfaceMember xlink:href="#P1"/></gml:CompositeSurface>'
    f'</root>',
    # Large coordinate list
    f'<gml:Polygon {NS32} srsDimension="2">'
    f'<gml:exterior><gml:LinearRing><gml:posList>'
    + ' '.join(f'{i} {i}' for i in range(50)) + ' 0 0'
    + f'</gml:posList></gml:LinearRing></gml:exterior>'
    f'</gml:Polygon>',
    # Mixed geometry stream
    f'<root {NS32}>'
    f'<gml:Point><gml:pos>1 2</gml:pos></gml:Point>'
    f'<gml:LineString srsDimension="2"><gml:posList>0 0 1 1</gml:posList></gml:LineString>'
    f'<gml:Envelope><gml:lowerCorner>0 0</gml:lowerCorner><gml:upperCorner>1 1</gml:upperCorner></gml:Envelope>'
    f'</root>',
    # Empty root
    '<root/>',
    # Empty input
    '',
]

# ---------------------------------------------------------------------------
# CityGML 2.0  (bldg namespace, GML 3.1.1 geometry)
# ---------------------------------------------------------------------------

NS_BLDG = 'xmlns:bldg="http://www.opengis.net/citygml/building/2.0"'
NS_GML311_BLDG = f'{NS_BLDG} xmlns:gml="http://www.opengis.net/gml"'

_SOLID3D = (
    '<gml:Solid srsDimension="3">'
    '<gml:exterior><gml:CompositeSurface>'
    '<gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing>'
    '<gml:posList>0 0 0 1 0 0 1 1 0 0 0 0</gml:posList>'
    '</gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember>'
    '</gml:CompositeSurface></gml:exterior>'
    '</gml:Solid>'
)

_MULTISURFACE3D = (
    '<gml:MultiSurface srsDimension="3">'
    '<gml:surfaceMember><gml:Polygon><gml:exterior><gml:LinearRing>'
    '<gml:posList>0 0 0 1 0 0 1 1 0 0 0 0</gml:posList>'
    '</gml:LinearRing></gml:exterior></gml:Polygon></gml:surfaceMember>'
    '</gml:MultiSurface>'
)

CITYGML20_SEEDS = [
    # lod0RoofEdge (MultiSurface)
    f'<bldg:Building {NS_GML311_BLDG} gml:id="b1">'
    f'<bldg:lod0RoofEdge>{_MULTISURFACE3D}</bldg:lod0RoofEdge>'
    f'</bldg:Building>',
    # lod0FootPrint (MultiSurface)
    f'<bldg:Building {NS_GML311_BLDG} gml:id="b2">'
    f'<bldg:lod0FootPrint>{_MULTISURFACE3D}</bldg:lod0FootPrint>'
    f'</bldg:Building>',
    # lod1Solid
    f'<bldg:Building {NS_GML311_BLDG} gml:id="b3">'
    f'<bldg:lod1Solid>{_SOLID3D}</bldg:lod1Solid>'
    f'</bldg:Building>',
    # lod1MultiSurface
    f'<bldg:Building {NS_GML311_BLDG} gml:id="b3b">'
    f'<bldg:lod1MultiSurface>{_MULTISURFACE3D}</bldg:lod1MultiSurface>'
    f'</bldg:Building>',
    # lod2Solid
    f'<bldg:Building {NS_GML311_BLDG} gml:id="b4">'
    f'<bldg:lod2Solid>{_SOLID3D}</bldg:lod2Solid>'
    f'</bldg:Building>',
    # lod2MultiSurface
    f'<bldg:Building {NS_GML311_BLDG} gml:id="b5">'
    f'<bldg:lod2MultiSurface>{_MULTISURFACE3D}</bldg:lod2MultiSurface>'
    f'</bldg:Building>',
    # lod2MultiCurve
    f'<bldg:Building {NS_GML311_BLDG} gml:id="b5b">'
    f'<bldg:lod2MultiCurve>{_MULTISURFACE3D}</bldg:lod2MultiCurve>'
    f'</bldg:Building>',
    # lod2Solid + lod2MultiSurface (両方)
    f'<bldg:Building {NS_GML311_BLDG} gml:id="b6">'
    f'<bldg:lod2Solid>{_SOLID3D}</bldg:lod2Solid>'
    f'<bldg:lod2MultiSurface>{_MULTISURFACE3D}</bldg:lod2MultiSurface>'
    f'</bldg:Building>',
    # Building without geometry
    f'<bldg:Building {NS_GML311_BLDG} gml:id="b7"></bldg:Building>',
    # CityModel wrapper
    f'<core:CityModel xmlns:core="http://www.opengis.net/citygml/2.0" {NS_GML311_BLDG}>'
    f'<core:cityObjectMember>'
    f'<bldg:Building gml:id="b8"><bldg:lod2Solid>{_SOLID3D}</bldg:lod2Solid></bldg:Building>'
    f'</core:cityObjectMember>'
    f'</core:CityModel>',
    # Empty root
    '<root/>',
    # Empty input
    '',
]

# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

def main() -> None:
    print("Generating fuzz seed corpus...")
    write_corpus("gml2_1_2", "FuzzReader", GML2_SEEDS)
    write_corpus("gml3_1_1", "FuzzReader", GML311_SEEDS)
    write_corpus("gml3_2_1", "FuzzReader", GML32_SEEDS)
    write_corpus("citygml2_0", "FuzzReader", CITYGML20_SEEDS)
    print("Done.")


if __name__ == "__main__":
    main()
