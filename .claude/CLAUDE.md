# CLAUDE.md

このリポジトリでコード作業を行う際の Claude Code へのガイド。

## 【最重要原則】XSD仕様を読んでから実装する

**実装・修正・判断を行う前に必ず XSD を確認すること。**

- GML 要素の挙動 → `docs/go/xsd2go-lite/schemas/gml/3.2.1/` の XSD
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

**基本方針:** 実データで使われる GML 要素を順次対応する。OGC GML Simple Features Profile (OGC 10-100r3) の SF-1 完了を当面のマイルストーンとする。

### マイルストーン

| レベル | 状態 | 説明 |
|---|---|---|
| **SF-0** | ✓ 完了 | Point / LineString / Polygon / Multi* / Envelope (線形補間のみ) |
| **SF-1** | △ 実装中 | Curve / Surface / OrientableCurve + CompositeCurve / CompositeSurface / OrientableSurface |
| **SF-2** | 未実装 | Arc / Circle 等の曲線補間 (日本政府データでは不要なため低優先) |

### 要素別実装状況

| ジオメトリ要素 | 変換先 | SF レベル | 状態 |
|---|---|---|---|
| `gml:Point` | `Point` | SF-0 | ✓ |
| `gml:LineString` | `LineString` | SF-0 | ✓ |
| `gml:Polygon` | `Polygon` | SF-0 | ✓ |
| `gml:MultiPoint` | `MultiPoint` | SF-0 | ✓ |
| `gml:MultiCurve` / `gml:MultiLineString` | `MultiLineString` | SF-0 | ✓ (curveMember が LineString の場合のみ) |
| `gml:MultiSurface` / `gml:MultiPolygon` | `MultiPolygon` | SF-0 | ✓ (surfaceMember が Polygon の場合のみ) |
| `gml:Envelope` | `Bound` | SF-0 | ✓ |
| `gml:Curve` + `LineStringSegment` | `LineString` | SF-1 | ✓ (N03 旧形式: 〜2023年 KsjAppSchema-N03-v3_x 以前) |
| `gml:Surface` + `PolygonPatch` | `Polygon` | SF-1 | ✓ (N03 新形式: 2024年〜 KsjAppSchema-N03-v4_0 以降) |
| `gml:OrientableCurve` | `LineString` | SF-1 | ✓ (xlink:href 解決) |
| `gml:CompositeCurve` | `LineString` | SF-1 | ✓ (inline Curve/LineString/OrientableCurve/CompositeCurve + xlink:href) |
| `gml:CompositeSurface` | `Polygon` | SF-1 | ✓ (rings from all surfaceMembers collected; xlink:href 未対応) |
| `gml:OrientableSurface` | `Polygon` | SF-1 | ✓ (baseSurface inline Polygon/Surface/CompositeSurface) |
| Arc/Circle 等の曲線補間 | — | SF-2 | 未実装 |
| `gml:Grid` / `gml:RectifiedGrid` | `GridCoverage` | — | ✓ (domainSet/rangeSet; inline + xlink:href) |
| Topology | — | — | 未実装 |

---

## コード構造

**マルチモジュールモノレポ (Google Cloud Go パターン)**

```
go-gml/                              # module root
├── core/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/core
│   ├── geometry.go                  # 公開型: Point, LineString, Polygon, Multi*, Bound, GridCoverage
│   └── reader.go                    # Geometry struct, Reader interface
├── gml2_1_2/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/gml2_1_2
│   ├── reader.go                    # NewReader(), DecodeGeometry()
│   ├── generated/
│   │   └── geometry.go              # xsd2go-lite 生成 (このモジュール専用)
│   └── internal/
│       └── decode_*.go
├── gml3_1_1/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/gml3_1_1
│   ├── reader.go                    # NewReader(), DecodeGeometry()
│   ├── generated/
│   │   └── geometry.go              # xsd2go-lite 生成 (このモジュール専用)
│   └── internal/
│       └── decode_*.go
├── gml3_2_1/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/gml3_2_1
│   ├── reader.go                    # NewReader(), DecodeGeometry()
│   ├── generated/
│   │   └── geometry.go              # xsd2go-lite 生成 (このモジュール専用)
│   └── internal/
│       ├── decode_*.go
│       └── ns_norm.go               # KSJ データバグ対応 (3.2.1 固有)
├── gml/
│   ├── go.mod                       # github.com/Kuroki-g/go-gml/gml  ← GML ユーザーの入口 (re-export)
│   ├── gml.go                       # core 型を type alias re-export + gml3_2_1.NewReader ラップ
│   └── crs.go                       # EPSGFromSRSName 公開
├── citygml/
│   └── go.mod                       # github.com/Kuroki-g/go-gml/citygml (将来)
├── waterml/
│   └── go.mod                       # github.com/Kuroki-g/go-gml/waterml (将来)
├── docs/                            # 内部用ツール (xsd2go-lite, gml-parser)
└── testdata/                        # 実 GML サンプル (CC BY 4.0)
```

### モジュール依存関係
```
go-gml/core
  ↑
  ├── go-gml/gml3_1_1  ─→  go-citygml (gml3_2_1 は入らない)
  └── go-gml/gml3_2_1  ─→  WaterML 等
        ↑
        └── go-gml/gml  ← ユーザー入口 (re-export)
```

`go.work` でローカル管理。

### 拡張モジュール

| モジュール | 場所 | GML バージョン |
|---|---|---|
| `github.com/Kuroki-g/go-gml/citygml` | `citygml/` | GML 3.1.1 (CityGML 2.0) |
| `github.com/Kuroki-g/go-gml/waterml` | `waterml/` | GML 3.2.1 |

**CityGML ターゲットバージョン: 2.0**
- 対象スキーマ: `http://www.opengis.net/citygml/2.0`
- 依存: `go-gml/core` + `go-gml/gml3_1_1` のみ (`gml3_2_1` は不要)
- CityGML 3.0 は後から対応できる設計にしておく (コア抽象化を壊さない)

### 命名規則
- 生成パッケージ (`v3_2_1`, `v3_1_1`, `v2_1_2`): アンダースコアOK (Google style guide generated code 例外)
- 手書きパーサパッケージ (`gml3_1_1`, `gml3_2_1`): アンダースコアなし

---

## リリース手順

詳細は [`docs/release.md`](../docs/release.md) を参照。

- ローカル開発中はバージョン bump・タグ不要 (`go.work` で解決)
- bump が必要なのはコードが変わったモジュールだけ (下流は MVS が自動選択)
- 依存順: `core → gml3_2_1/gml3_1_1/gml2_1_2 → gml → cmd/gml-parser`
- tidy は `GONOSUMDB='*'` を付けて実行する

---

## 開発コマンド

**必ず Makefile 経由で実行すること。** 直接 `go test` / `go run` は WSL `/tmp noexec` 制約で失敗する。Bash コマンドは必ず1つずつ実行すること (`&&` / `||` 禁止)。

```bash
make build          # go build ./...
make test           # 全テスト実行
make cover          # カバレッジレポート

make xsd2go-build   # xsd2go-lite バイナリビルド
make xsd2go-test    # xsd2go-lite テスト
make xsd2go-gen     # GML XSD → gml3_2_1/generated/geometry.go 生成 (GML_VERSION=3.2.1 デフォルト)

make gml-parser-build
make gml-parser-run
```

---

## テスト方針

`make test` で全テストを実行する。build tag による分離は現状未使用。

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
- CLI (`cmd/gml-parser/`) は独立モジュールなので外部ライブラリを使ってよい

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
