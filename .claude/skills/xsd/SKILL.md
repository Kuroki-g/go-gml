---
name: xsd
description: git diff の変更内容が Schema 準拠かどうか、および他バージョンで実装漏れがないかをレビューするスキル。実装・修正後に使う。Trigger when the user types /xsd or asks to verify schema compliance or check for missing implementations across versions.
disable-model-invocation: true
---

# /xsd — Schema 準拠 & クロスバージョン実装漏れ確認

## Step 1: 変更内容の把握

1. `git diff HEAD` で変更を確認
2. 変更されたモジュール・要素・デコーダ関数を抽出する

## Step 2: Schema 確認

Schema の場所は CLAUDE.md の【最重要原則】セクションを参照。

変更に関係する要素を特定し、**変更されたバージョンの Schema** を Read して確認する:
- 要素定義・型定義・substitutionGroup
- 子要素の minOccurs / maxOccurs
- 属性の必須/任意
- abstract 要素と concrete 要素の関係

実装が Schema 定義と一致しているかを判定する。

## Step 3: クロスバージョン実装漏れ確認

変更されたモジュール以外の全バージョンで、同じ要素・機能の実装状況を確認する:

1. Grep で同等のデコーダ・型が他バージョンに存在するか確認
2. 存在する場合: 同じ修正が必要か (バージョン間の Schema 差異を考慮)
3. 存在しない場合: 他バージョンの Schema でその要素が定義されているか確認

## Step 4: レポート

```
## Schema 準拠チェック結果

### 変更サマリー
- モジュール / 要素:

### Schema 準拠
- [x/] <項目>: <根拠>

### クロスバージョン
- <バージョン>: 問題なし / 要対応: <内容>

### アクション
- <対応が必要なら列挙。漏れは docs/issues/ に記録>
```

問題がなければ「準拠 OK / 漏れなし」と明記する。
