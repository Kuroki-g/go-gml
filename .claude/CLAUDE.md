# CLAUDE.md

このリポジトリでコード作業を行う際の Claude Code へのガイド。

## 【最重要原則】XSD仕様を読んでから実装する

**実装・修正・判断を行う前に必ず XSD を確認すること。**
！！！v3.1.1のみ確認し、他のバージョンの確認漏れが頻発している！！特定のバージョンの実装漏れが判明した場合、すべてのバージョンの実装を確認せよ！！
- GML 要素の挙動: バージョン毎に必ず要確認
  - v2.1.2 → `docs/schemas/gml/2.1.2/` の XSD
  - v3.1.1 → `docs/schemas/gml/3.1.1/` の XSD
  - v3.2.1 → `docs/schemas/gml/3.2.1/` の XSD
- CityGML 要素の挙動 → `docs/schemas/citygml/` の XSD (1.0 / 2.0 / 3.0)
- CRS/srsName → GML XSD の `SRSReferenceGroup` 定義

**NG パターン:**

- XSD を読まずに「非準拠」と断言する
- 特定データの文字列にハードコードマッピングを追加する
- 症状だけ見て場当たり的に fix する

---

## プロジェクト

**go-gml**

- gml package: GML (ISO 19136) パーサの Pure Go 実装。国土数値情報・基盤地図情報・PLATEAU など実際の日本政府 GML データが読める SDK。

**基本方針:** XSD準拠に従い実装する。

**参照仕様書:**

- CityGML 2.0: OGC 12-019 (`testdata/12-019_OGC_City_Geography_Markup_Language_CityGML_Encoding_Standard.pdf`)
  - LoD 定義: Section 6.2 (p.11–12), Tab. 3 (p.12)
  - Building モデル: Section 10.3 (p.63–)
- CityGML XSD: `docs/schemas/citygml/` (v1.0 / v2.0 / v3.0)

### マイルストーン (GML パーサ)

| レベル | 状態 | 説明 |
|---|---|---|
| **SF-0** | ✓ 完了 | Point / LineString / Polygon / Multi* / Envelope |
| **SF-1** | ✓ 完了 | Curve / Surface / OrientableCurve + CompositeCurve / CompositeSurface / OrientableSurface |
| **SF-2** | 未実装 | Arc / Circle 等の曲線補間 (日本政府データでは不要なため低優先) |

### マイルストーン (CityGML 2.0)

| LoD | 状態 | 説明 |
|---|---|---|
| **LoD0** | ✓ 完了 | `reader.go` / `building.go` / `internal/subtree.go` 実装済み。`lod0FootPrint` / `lod0RoofEdge` (MultiSurface) をストリーム読み取り |
| **LoD1** | ✓ 完了 | ブロックモデル (`bldg:lod1Solid` → `gml:Solid`) |
| **LoD2** | ✓ 完了 | `bldg:lod2Solid` / `lod2MultiSurface` / `lod2MultiCurve` / `lod2TerrainIntersection` |
| **LoD3** | 未実装 | 建築詳細モデル (`bldg:lod3Solid`) |

### 未実装・残課題

- **CityGML 2.0 LoD3**: 次の実装対象 (`bldg:lod3Solid` → `gml:Solid`)
- **check-coverage 残課題**: `docs/issues/check-coverage-bug-c-xlink-meta-attrs.md` (xlink meta attrs)
- **gml3_2_1 nilReason 未チェック**: `docs/issues/nil-reason-gml3_2_1.md` — AssociationAttributeGroup の nilReason 属性が未処理
- **gml3_1_1 / gml3_2_1 ストリームテストなし**: N03 旧形式データ取得困難。`Decode` メソッドテストは追加済み。CityGML 実装で代替する方針
- **SF-2**: `docs/issues/sf2-curves.md` — Arc/Circle 等の曲線補間・SurfacePatchArray 内の Cone/Cylinder/Sphere/Rectangle 未実装。低優先

---

## コード構造

マルチモジュールモノレポ。`go.work` でローカル管理。

**各モジュールの構成パターン (gml*/citygml*共通):**
- `reader.go` — `NewReader()` / `Next()` エントリポイント。`internal/reader.go` の switch が対応済み要素の索引
- `generated/` — xsd2go-lite 生成 struct。手書き禁止
- `internal/decode_*.go` — 要素別デコード実装
- `internal/subtree.go` (citygml2_0のみ) — GML サブツリー読み取りユーティリティ
- `internal/ns_norm.go` (gml3_2_1のみ) — KSJ データバグ対応

**citygml-parser のサブコマンド:** `inspect` (Building 数・LoD カウント) / `version`

### モジュール依存関係

`→` は「依存する」を示す。

```
cmd/gml-parser      → gml → core
                          → gml2_1_2  → core
                          → gml3_1_1  → core
                          → gml3_2_1  → core

cmd/citygml-parser  → citygml2_0 → core
                                 → gml3_1_1 → core
                    → gml3_1_1   → core
```

各モジュールの依存まとめ:

| モジュール | 依存 |
|---|---|
| `core` | なし |
| `gml2_1_2` | core |
| `gml3_1_1` | core |
| `gml3_2_1` | core |
| `gml` | core + gml2_1_2 + gml3_1_1 + gml3_2_1 (全バージョン re-export) |
| `citygml2_0` | core + gml3_1_1 |
| `cmd/gml-parser` | gml |
| `cmd/citygml-parser` | citygml2_0 + gml3_1_1 |
| `citygml3_0` (将来) | core + gml3_2_1 |

※ `citygml-core` (将来): `citygml3_0` 実装時に `citygml2_0` から共通型を切り出して追加予定

`go.work` でローカル管理。

**CityGML モジュール設計方針:**

- `citygml2_0`: `NewReader` は `core.Decoder` を受け取る形 (将来の GML バージョン差し替え口)
- `citygml-core` (将来): `citygml3_0` 実装時に共通型 (`CityObject`, `CityModel`, `Building`) を切り出す予定
- バージョン固有の namespace URI・要素名 const は各バージョンモジュールで定義し re-export

### 命名規則

- 生成パッケージ (`v3_2_1`, `v3_1_1`, `v2_1_2`): アンダースコアOK (Google style guide generated code 例外)
- 手書きパーサパッケージ (`gml3_1_1`, `gml3_2_1`): アンダースコアなし

---

## リリース手順

詳細は [`docs/release.md`](../docs/release.md) を参照。

- ローカル開発中はバージョン bump・タグ不要 (`go.work` で解決)
- bump が必要なのはコードが変わったモジュールだけ (下流は MVS が自動選択)
- 依存順: `core → gml3_2_1/gml3_1_1/gml2_1_2 → gml → cmd/gml-parser`
- `citygml2_0` リリース時: `core` を先に bump (v0.1.2 以降) → `citygml2_0` タグ → `cmd/citygml-parser` タグ
- tidy は `GONOSUMDB='*'` を付けて実行する

---

## 開発コマンド

コマンド一覧は `make help` で確認。**必ず Makefile 経由で実行すること** (直接 `go test` / `go run` は WSL `/tmp noexec` 制約で失敗する)。

**PDF 解析が必要な場合:** `uv` 経由で `pypdf` を使う。事前にユーザーの許可を取ること。

```bash
uv venv /tmp/pdf-env
uv pip install --system-certs --python /tmp/pdf-env/bin/python pypdf
/tmp/pdf-env/bin/python3 -c "import pypdf; ..."
```

---

## ファイル探索ルール

**Glob の結果は truncated される場合がある。** truncated された結果から「ファイルが存在しない」と結論づけてはならない。

- ディレクトリ構造を確認するときは **まず `ls` でトップレベルを確認**してから掘り下げる
- `**/*.xsd` のような深いパターンを使う前に `ls <dir>/` で上位構造を把握する

---

## テスト方針

`make test` で全テストを実行する。

**fuzz テスト:** `gml2_1_2/`, `gml3_1_1/`, `gml3_2_1/` に `fuzz_test.go` (`FuzzReader`) 追加済み。
seed corpus は `make fuzz-gen` (→ `scripts/gen_fuzz_seeds.py`) で都度生成する。
corpus ファイルは gitignore 済み (`**/testdata/fuzz/`)。
実行: `go -C gml3_2_1 test -fuzz=FuzzReader -fuzztime=60s`

---

## 実装ルール

### 禁止: XML unmarshal 用 struct の手書き

`internal/` 内で `type xmlXxx struct { ... }` を手書きしてはならない。

**新しい XML 要素に対応する手順:**

1. `xsd2go-lite` で生成する (`make xsd2go-gen`)
2. 生成できない場合は `xsd2go-lite` を修正してから再生成する
3. 生成 struct を各パーサモジュールの `generated/` から import して decode に使う

### ファイルサイズと構造

- **1ファイル ~150行以下** を厳守 (AI コンテキスト効率のため)
- `internal/reader.go` の switch 文が対応済み要素の索引 → 新要素は `internal/decode_*.go` 追加 + switch に1 case 追加

### 依存関係

- 各パーサモジュールは `encoding/xml` (stdlib) のみ。外部ライブラリを追加しない
- CLI (`cmd/*`) は独立モジュールなので外部ライブラリを使ってよい

---

## 既知の注意点

### GML 座標表現のバリエーション

- 形式: `gml:pos` / `gml:posList` / `gml:coordinates`
- 区切り: スペース (標準)、カンマ、混合
- 次元: 2D / 3D; srsName が軸順を指定する場合あり

### CRS / srsName のバリエーション

- `EPSG:4326`
- `urn:ogc:def:crs:EPSG::4326`
- `http://www.opengis.net/def/crs/EPSG/0/4326`
- `urn:ogc:def:crs:EPSG:6.18:4326`

### GML 名前空間

- GML 2.x: `http://www.opengis.net/gml`
- GML 3.x: `http://www.opengis.net/gml/3.2` (または 3.1.1)

### cobra-cli

呼び出しの際にはgo.workと干渉するので注意:
`GOWORK=off cobra-cli init`