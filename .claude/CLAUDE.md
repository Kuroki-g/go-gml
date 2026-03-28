# CLAUDE.md

このリポジトリでコード作業を行う際の Claude Code へのガイド。

## 【最重要原則】XSD仕様を読んでから実装する

**実装・修正・判断を行う前に必ず XSD を確認すること。**

- GML 要素の挙動 → `docs/go/xsd2go-lite/schemas/gml/3.2.1/` の XSD
- CityGML 要素の挙動 → `docs/go/xsd2go-lite/schemas/citygml/` の XSD (1.0 / 2.0 / 3.0)
- CRS/srsName → GML XSD の `SRSReferenceGroup` 定義

**禁止 (Stop フックでブロック):** 「スコープ外」「仕様外」「対応しない」
→ GML 3.2.1 XSD に定義されている要素はすべて**仕様内**。XSD で確認してから判断すること。

**使ってよい表現:** 「未実装」「実装済み」

**NG パターン:**
- XSD を読まずに「非準拠」と断言する
- 特定データの文字列にハードコードマッピングを追加する
- 症状だけ見て場当たり的に fix する

---

## プロジェクト

**go-gml** — GML (ISO 19136) パーサの Pure Go 実装。国土数値情報・基盤地図情報・PLATEAU など実際の日本政府 GML データが読める SDK。

**基本方針:** 実データで使われる GML 要素を順次対応する。対応済み要素の一覧は README を参照。

**参照仕様書:**
- CityGML 2.0: OGC 12-019 (`testdata/12-019_OGC_City_Geography_Markup_Language_CityGML_Encoding_Standard.pdf`)
  - LoD 定義: Section 6.2 (p.11–12), Tab. 3 (p.12)
  - Building モデル: Section 10.3 (p.63–)
- CityGML XSD: `docs/go/xsd2go-lite/schemas/citygml/` (v1.0 / v2.0 / v3.0)

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
| **LoD2** | 未実装 | 屋根形状あり (`bldg:lod2Solid`) |
| **LoD3** | 未実装 | 建築詳細モデル (`bldg:lod3Solid`) |

### 未実装・残課題

- **CityGML 2.0 LoD2**: 次の実装対象 (`bldg:lod2Solid` → `gml:Solid`)
- **gml3_1_1 srsDimension 継承** (`internal/decode_polygon.go:48`): ルート `gml:Envelope` の srsDimension が個々の Polygon に伝播しない。アーキテクチャ変更が必要
- **gml3_1_1 ストリームテストなし**: N03 旧形式データ取得困難。`Decode` メソッドテストは追加済み。CityGML 実装で代替する方針
- **SF-2**: Arc / Circle 等の曲線補間 — 低優先

---

## コード構造

**マルチモジュールモノレポ (Google Cloud Go パターン)**

```
go-gml/                              # module root
├── core/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/core
│   ├── geometry.go                  # 公開型: Point, LineString, Polygon, Multi*, Solid, Bound, GridCoverage
│   ├── reader.go                    # Geometry struct, Reader interface
│   └── decoder.go                   # Decoder interface (Decode メソッド)
├── gml2_1_2/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/gml2_1_2
│   ├── reader.go                    # NewReader(); Decode() — core.Decoder 実装
│   ├── generated/
│   │   └── geometry.go              # xsd2go-lite 生成 (このモジュール専用)
│   └── internal/
│       └── decode_*.go
├── gml3_1_1/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/gml3_1_1
│   ├── reader.go                    # NewReader(); Decode() — core.Decoder 実装
│   ├── generated/
│   │   └── geometry.go              # xsd2go-lite 生成 (このモジュール専用)
│   └── internal/
│       └── decode_*.go
├── gml3_2_1/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/gml3_2_1
│   ├── reader.go                    # NewReader(); Decode() — core.Decoder 実装
│   ├── generated/
│   │   └── geometry.go              # xsd2go-lite 生成 (このモジュール専用)
│   └── internal/
│       ├── decode_*.go
│       └── ns_norm.go               # KSJ データバグ対応 (3.2.1 固有)
├── gml/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/gml  ← GML ユーザーの入口 (re-export)
│   ├── gml.go                       # core 型を type alias re-export + gml3_2_1.NewReader ラップ
│   └── crs.go                       # EPSGFromSRSName 公開
├── citygml2_0/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/citygml2_0 (require core, gml3_1_1)
│   ├── reader.go                    # NewReader(r io.ReadSeeker, dec core.Decoder) *Reader; Next() (*Building, error)
│   ├── building.go                  # Building (lod0FootPrint/RoofEdge/lod1Solid を直接定義)
│   ├── decode_building.go           # bldg:Building デコード (LoD0/LoD1)
│   ├── generated/
│   │   ├── building.go              # xsd2go-lite 生成 (citygml/building/2.0)
│   │   ├── core.go                  # xsd2go-lite 生成 (citygml/2.0)
│   │   └── xal.go                   # xsd2go-lite 生成 (xAL 2.0)
│   └── internal/
│       └── subtree.go               # GML サブツリー読み取りユーティリティ
├── cmd/
│   ├── gml-parser/                  # GML CLI (github.com/Kuroki-g/go-gml/cmd/gml-parser)
│   └── citygml-parser/              # CityGML CLI (github.com/Kuroki-g/go-gml/cmd/citygml-parser)
│       └── cmd/
│           ├── root.go              # rootCmd + openInput
│           ├── inspect.go           # inspect サブコマンド (Building 数・LoD カウント)
│           └── version.go           # version サブコマンド (vcs.revision から自動取得)
├── waterml/
│   └── go.mod                       # github.com/Kuroki-g/go-gml/waterml (将来)
├── scripts/
│   ├── e2e_random.sh                # E2E スモークテスト (gml-parser でランダム順に inspect/convert)
│   └── gen_fuzz_seeds.py            # fuzz seed corpus 生成 (gml*/testdata/fuzz/FuzzReader/)
├── pyproject.toml                   # uv プロジェクト定義 (scripts/ 用; 現在は stdlib のみ)
├── docs/                            # 内部用ツール (xsd2go-lite, gml-parser)
└── testdata/                        # 実 GML サンプル (CC BY 4.0)
```

### モジュール依存関係
```
go-gml/core  (core.Decoder interface を定義)
  ↑
  ├── go-gml/citygml2_0               # v2.0 パーサ; gml3_1_1 に直接依存
  │     ↑
  │     └── go-gml/citygml3_0         # (将来) v3.0 パーサ (同様に core.Decoder DI)
  ├── go-gml/gml3_1_1                 # core.Decoder を実装 (Decode メソッド追加)
  ├── go-gml/gml3_2_1                 # core.Decoder を実装 (Decode メソッド追加)
  │     ↑
  │     └── go-gml/gml  ← ユーザー入口 (re-export)
  └── go-gml/gml2_1_2                 # core.Decoder を実装 (Decode メソッド追加)
```
※ `citygml-core` (将来): `citygml3_0` 実装時に `citygml2_0` から共通型を切り出して追加予定

`go.work` でローカル管理。

### 拡張モジュール

| モジュール | 場所 | 依存 |
|---|---|---|
| `github.com/Kuroki-g/go-gml/citygml2_0` | `citygml2_0/` | core + gml3_1_1 |
| `github.com/Kuroki-g/go-gml/waterml` | `waterml/` | core + gml3_2_1 |
| `github.com/Kuroki-g/go-gml/citygml-core` | `citygml-core/` (将来) | core のみ |

**CityGML モジュール設計方針:**
- `citygml2_0`: `gml3_1_1` に直接依存。`NewReader` は `core.Decoder` を受け取る形 (将来の GML バージョン差し替え口)。`Building` 型を直接定義 (`lod0FootPrint`/`lod0RoofEdge`)
- `citygml-core` (将来): `citygml3_0` 実装時に共通型 (`CityObject`, `CityModel`, `Building`) を切り出す予定
- `citygml3_0` (将来): 同様に `core.Decoder` DI。GML 3.2.1 ベース
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

**必ず Makefile 経由で実行すること。** 直接 `go test` / `go run` は WSL `/tmp noexec` 制約で失敗する。Bash コマンドは必ず1つずつ実行すること (`&&` / `||` 禁止)。

**PDF 解析が必要な場合:** `uv` 経由で `pypdf` を使う。事前にユーザーの許可を取ること。

```bash
uv venv /tmp/pdf-env
uv pip install --system-certs --python /tmp/pdf-env/bin/python pypdf
/tmp/pdf-env/bin/python3 -c "import pypdf; ..."
```

```bash
make build          # go build ./...
make test           # 全テスト実行
make cover          # カバレッジレポート

make xsd2go-build   # xsd2go-lite バイナリビルド
make xsd2go-test    # xsd2go-lite テスト
make xsd2go-gen     # GML XSD → gml3_2_1/generated/geometry.go 生成 (GML_VERSION=3.2.1 デフォルト)
make citygml2_0-gen # CityGML 2.0 XSD → citygml2_0/generated/{building,core,xal}.go 生成

make fuzz-gen       # fuzz seed corpus 生成 (gml*/testdata/fuzz/FuzzReader/) ← 毎回再生成

make gml-parser-build
make gml-parser-run

make citygml-parser-build
make citygml-parser-run
```

---

## ファイル探索ルール

**Glob の結果は truncated される場合がある。** truncated された結果から「ファイルが存在しない」と結論づけてはならない。

- ディレクトリ構造を確認するときは **まず `ls` でトップレベルを確認**してから掘り下げる
- `**/*.xsd` のような深いパターンを使う前に `ls <dir>/` で上位構造を把握する
- 「見つからなかった」と言う前に `ls` で確認すること

---

## テスト方針

`make test` で全テストを実行する。build tag による分離は現状未使用。

**fuzz テスト:** `gml2_1_2/`, `gml3_1_1/`, `gml3_2_1/` に `fuzz_test.go` (`FuzzReader`) 追加済み。
seed corpus は `make fuzz-gen` (→ `scripts/gen_fuzz_seeds.py`) で都度生成する。corpus ファイルは gitignore 済み (`**/testdata/fuzz/`)。
実行: `go -C gml3_2_1 test -fuzz=FuzzReader -fuzztime=60s`

---

## 実装ルール

### 禁止: XML unmarshal 用 struct の手書き

`internal/` 内で `type xmlXxx struct { ... }` を手書きしてはならない。

**新しい GML 要素に対応する手順:**
1. `docs/go/xsd2go-lite/schemas/gml/<version>/` の XSD で要素定義を確認する
2. `xsd2go-lite` が該当 struct を生成できるか確認する (`make xsd2go-gen`)
3. 生成できない場合は `xsd2go-lite` を修正してから再生成する
4. 生成 struct を各パーサモジュールの `generated/` から import して decode に使う

### ファイルサイズと構造

- **1ファイル ~150行以下** を厳守 (AI コンテキスト効率のため)
- `reader.go` の `handlers` マップが対応済み要素の索引 → 新要素は `internal/decode_*.go` 追加 + `handlers` に1行追加

### 依存関係

- 各パーサモジュールは `encoding/xml` (stdlib) のみ。外部ライブラリを追加しない
- CLI (`cmd/gml-parser/`, `cmd/citygml-parser/`) は独立モジュールなので外部ライブラリを使ってよい

---

## 既知の注意点

### GML 座標表現のバリエーション
- 形式: `gml:pos` / `gml:posList` / `gml:coordinates` (非推奨)
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
