# リリース手順 (マルチモジュールモノレポ)

## 原則

**ローカル開発中はバージョン bump もタグも不要。**
`go.work` の `use` ディレクティブがローカルコードを解決するため、`make build` / `make test` はそのまま動く。

バージョン bump・タグ・push が必要なのは **外部に publish するときだけ**。

---

## どのモジュールを bump するか

**bump が必要:** コードが変わったモジュール、またはそのモジュールが「新バージョンの上流を最低要件にしたい」場合。

**bump 不要:** コードが変わっていない下流モジュール。Go の MVS が自動的に上位バージョンを選ぶ。

例: `core` だけ変更 → `core/vX.Y.Z` のみ bump。`gml3_2_1` 等はコードが変わっていなければそのまま。

---

## リリース手順

### 1. 変更モジュールを特定する

```
git diff --name-only <前回タグ>..HEAD
```

変更ファイルが属するモジュールを確認する。

### 2. go.mod の require を更新する

変更モジュールの `go.mod` 内で、上流モジュールの require を新バージョンに上げる（新バージョンの機能・修正が必要な場合のみ）。

### 3. コミットする

```bash
git add <変更ファイル>
git commit -m "chore: bump <module> to vX.Y.Z"
```

**タグはコミット後に打つこと。**

### 4. go mod tidy を実行する

依存順に実行する（上流から下流へ）:

```bash
GONOSUMDB='*' GOTMPDIR=/workspace/go-gml/.tmp go mod tidy -C <module>
```

`GONOSUMDB='*'` は private repo の checksum DB 検索をスキップするため必須。

go.sum が新規生成・更新された場合はコミットする:

```bash
git add <module>/go.sum
git commit -m "chore(<module>): update go.sum"
```

### 5. タグを打って push する

依存順（上流から下流）に push する:

```
core → gml3_2_1 / gml3_1_1 / gml2_1_2 → gml → cmd/gml-parser
```

```bash
git tag <module>/vX.Y.Z
git push origin <module>/vX.Y.Z
```

**proxy.golang.org への反映待ちが発生することがある。**
タグ push 後すぐに下流の tidy が失敗する場合は、ブラウザで以下を開いてプロキシにインデックスさせる:

```
https://proxy.golang.org/github.com/Kuroki-g/go-gml/<module>/@v/vX.Y.Z.info
```

JSON が返れば反映済み。その後に下流モジュールの tidy を実行する。

---

## モジュール依存順

```
core
  ├── gml3_2_1
  ├── gml3_1_1
  └── gml2_1_2
        └── gml  (core + gml3_2_1 + gml3_1_1 + gml2_1_2)
              └── cmd/gml-parser
```

---

## よくある失敗

| 症状 | 原因 | 対処 |
|---|---|---|
| `git ls-remote HTTP 405` | HTTPS git が WSL でブロックされている | `GOPROXY=https://proxy.golang.org` (デフォルト) を使う。`GOPROXY=direct` は使わない |
| `reading go.mod at revision vX.Y.Z` | プロキシ未インデックス or タグ未 push | タグを push → ブラウザでプロキシ URL を開く |
| ローカルビルドが壊れた | 存在しないバージョンを require に書いた | require は最後に publish 済みのバージョンのままにしておく |
