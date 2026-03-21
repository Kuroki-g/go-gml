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

**基本方針:** 実データで使われる GML 要素を順次対応する。

| ジオメトリ要素 | 変換先 | 状態 |
|---|---|---|
| `gml:Point` | `Point` | ✓ |
| `gml:LineString` | `LineString` | ✓ |
| `gml:Polygon` | `Polygon` | ✓ |
| `gml:MultiPoint` | `MultiPoint` | ✓ |
| `gml:MultiCurve` / `gml:MultiLineString` | `MultiLineString` | ✓ (curveMember が LineString の場合のみ) |
| `gml:MultiSurface` / `gml:MultiPolygon` | `MultiPolygon` | ✓ (surfaceMember が Polygon の場合のみ) |
| `gml:Envelope` | `Bound` | ✓ |
| `gml:Surface` + `PolygonPatch` | `Polygon` | 未実装 (N03 新形式) |
| `gml:Curve` + `LineStringSegment` | `LineString` | 未実装 (N03 旧形式) |
| `gml:CompositeCurve` | `LineString` | 未実装 (W09 湖沼) |
| Arc/Circle 等の曲線補間 | — | 未実装 |
| Topology / Coverage | — | 未実装 |

---

## コード構造

```
go-gml/                          # module: github.com/Kuroki-g/go-gml
├── pkg/
│   ├── gml/                     # メイン GML パーサ
│   │   ├── parser.go            # Reader, Next(), handlers マップ (対応済み要素の索引)
│   │   ├── geometry.go          # ジオメトリ型定義
│   │   ├── coords.go            # 座標文字列パース
│   │   ├── crs.go               # EPSG 抽出
│   │   └── decode/              # decode 関数群 (package gml のまま)
│   └── gml/v3/geometry.go       # xsd2go-lite 生成 struct
├── docs/go/
│   ├── xsd2go-lite/             # XSD → Go struct ジェネレータ
│   │   └── schemas/gml/3.2.1/   # GML XSD (仕様参照先)
│   └── gml-parser/              # CLI (cobra、convert/inspect)
└── testdata/                    # 実 GML サンプル (CC BY 4.0)
```

`go.work` でローカル管理。現在の use: `. / docs/go/gml-parser`。

CityGML は同一リポジトリ内の別モジュール (`extensions/citygml/`) として管理する。モジュール名は `github.com/Kuroki-g/go-citygml`、go-gml に依存する。

---

## 開発コマンド

**必ず Makefile 経由で実行すること。** 直接 `go test` / `go run` は WSL `/tmp noexec` 制約で失敗する。Bash コマンドは必ず1つずつ実行すること (`&&` / `||` 禁止)。

```bash
make build          # go build ./...
make test           # Small テストのみ
make test-medium    # Medium も含む
make test-large     # 全サイズ
make cover          # カバレッジレポート

make xsd2go-build   # xsd2go-lite バイナリビルド
make xsd2go-test    # xsd2go-lite テスト
make xsd2go-gen     # GML XSD → pkg/gml/v3/geometry.go 生成

make gml-parser-build
make gml-parser-run
```

---

## テスト方針 (Google Test Sizes)

| サイズ | build tag | 対象 |
|---|---|---|
| Small | なし | 純粋関数テスト (coords, crs 等) |
| Medium | `//go:build medium` | `testdata/` の実 GML ファイルを読むテスト |
| Large | `//go:build large` | 外部サーバー・大容量ファイル |

---

## 実装ルール

### 禁止: XML unmarshal 用 struct の手書き

`decode/` 内で `type xmlXxx struct { ... }` を手書きしてはならない。

**新しい GML 要素に対応する手順:**
1. `docs/go/xsd2go-lite/schemas/gml/3.2.1/` の XSD で要素定義を確認する
2. `xsd2go-lite` が該当 struct を生成できるか確認する (`make xsd2go-gen`)
3. 生成できない場合は `xsd2go-lite` を修正してから再生成する
4. 生成 struct を `pkg/gml/v3/` から import して `decode/` で使う

### ファイルサイズと構造

- **1ファイル ~150行以下** を厳守 (AI コンテキスト効率のため)
- `decode/` は `package gml` のまま (サブパッケージにしない)
- `parser.go` の `handlers` マップが対応済み要素の索引 → 新要素は `decode/` にファイル追加 + `handlers` に1行追加

### 依存関係

- `pkg/gml/` コアは `encoding/xml` (stdlib) のみ。外部ライブラリを追加しない
- CLI (`docs/go/gml-parser/`) は独立モジュールなので外部ライブラリを使ってよい

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
