# tagslist/toptags Go移行計画

## 目的

- 現在の `tagslist.sh` と `toptags.sh` の機能を Go バイナリへ移行し，運用の再現性と保守性を高める。
- 既存運用（`build.sh` -> `publish.sh`）を壊さず，段階的に切り替える。
- 出力互換性を最優先し，既存ファイル差分を最小化する。

## 非目的

- front matter 形式の変更（TOML -> YAML など）は行わない。
- 現時点でのタグ意味付けルール（`means` の運用方針）を変更しない。

## 採用ライブラリ

- エラーハンドリング: `github.com/goark/errs`
- コマンドラインパラメータ制御: `github.com/goark/struct2pflag`
- CSV 読み込み: `github.com/goark/csvdata`

## 開発基準ディレクトリ

- 本リポジトリ内に `tagtools/` ディレクトリを作成し，Go 実装一式の基準ディレクトリとする。
- 既存の Hugo コンテンツや運用スクリプトと責務を分離し，段階導入とロールバックを容易にする。

## 現行仕様（要件化）

### tagslist.sh

- 入力:
  - `content/**/*.md`
  - 既存 `.github/workflows/tagslist.csv`（`means` の引き継ぎ元）
- 抽出対象:
  - TOML front matter（`+++ ... +++`）内の `tags = [ ... ]`
  - 複数行 `tags` に対応
- 出力:
  - `.github/workflows/tagslist.csv`
  - ヘッダ: `tag,count,means`
  - 並び順: `count` 降順，`tag` 昇順
  - `means` は既存CSVから同一tagの値を再利用
- CSV要件:
  - ダブルクォートのエスケープ（`"` -> `""`）
  - 読み込み時にCSVパースを厳密に実施

### toptags.sh

- 入力:
  - `content/**/*.md`
- 抽出条件:
  - front matter の `date = "YYYY-MM-DD..."`
  - 期間: `cutoff = 今日 - 1年` から `today` まで
- 集計:
  - タグ頻度を降順，タグ名昇順で並べる
  - 上位 `TOP_N` 件を抽出（既定 `15`，環境変数で変更可）
  - 最終出力はタグ名のみを昇順配列化
- 出力:
  - `data/toptags.json`（JSON文字列配列）

## 提案CLI仕様

単一バイナリ案（推奨）:

- コマンド名: `tagtools`
- サブコマンド:
  - `tagtools tagslist`
  - `tagtools toptags`
  - `tagtools all`（`toptags` -> `tagslist` の順）
  - `tagtools verify`（互換性検証用）

### tagtools tagslist オプション

- `--content-dir`（既定: `content`）
- `--out`（既定: `.github/workflows/tagslist.csv`）
- `--inherit-means-from`（既定: 出力先と同じ）

### tagtools toptags オプション

- `--content-dir`（既定: `content`）
- `--out`（既定: `data/toptags.json`）
- `--top-n`（既定: `15`）
- `--today`（テスト用日付固定）
- `--window`（既定: `1y`）

### tagtools verify オプション

- `--content-dir`（既定: `content`）
- `--expected-tagslist`（既定: `.github/workflows/tagslist.csv`）
- `--expected-toptags`（既定: `data/toptags.json`）
- `--inherit-means-from`（既定: `.github/workflows/tagslist.csv`）
- `--top-n`（既定: `15`）
- `--today`（テスト用日付固定）
- `--window`（既定: `1y`）
- `--debug`（JSON出力を有効化）

## verify 出力仕様

### 通常モード

- 成功時: `Verification passed`
- 失敗時: 以下を含むテキストを出力
  - 失敗ステージ（`tagslist` または `toptags`）
  - expected/actual ファイル
  - 最初の差分行番号
  - expected/actual の該当行
  - `reproduce: ...`（再現コマンド）

### `--debug` モード（JSON）

- 成功時

```json
{
  "status": "ok",
  "reproduce_command": "./bin/tagtools verify ... --debug"
}
```

- 失敗時

```json
{
  "status": "error",
  "stage": "toptags",
  "message": "toptags verification failed (first diff line: 1)",
  "expected_path": "testdata/golden/toptags.json",
  "actual_path": "/tmp/tagtools-verify-xxx/toptags.json",
  "line": 1,
  "expected_line": "...",
  "actual_line": "...",
  "reproduce_command": "./bin/tagtools verify ... --debug"
}
```

- JSON は機械可読性を優先し，テキストの整形情報は含めない。

## 実装方針

- Front matter 解析は TOML 専用の軽量パーサを実装し，既存挙動を優先して再現する。
- Markdown 本文は走査対象外。front matter 範囲のみ解析。
- 既存出力との差分ゼロを目標にゴールデンテストを作成する。
- ファイル列挙順に依存しないため，最終ソートを必ず実施する。
- CLI エントリポイントは `tagtools/main.go` に置く。

## 互換性テスト計画

### 1. ゴールデン比較

- 現行 `tagslist.sh` 実行結果を `testdata/golden/tagslist.csv` として保存。
- 現行 `toptags.sh` 実行結果を `testdata/golden/toptags.json` として保存。
- Go版出力とバイト比較（改行コード含む）を行う。

### 2. エッジケース

- `tags` 複数行配列
- `tags` 空配列
- `date` 欠落記事
- `date` が期間外
- CSVの `means` にカンマや引用符を含むケース

### 3. 回帰確認

- `build.sh` 実行後の差分確認
- `publish.sh` 実行後の `text` 側残差分確認

## 段階的移行

### Phase 1: 並走導入

- `tagtools/` を追加（`main.go` を含む）
- 既存 `.sh` は維持
- CIまたは手元で `sh` 出力と Go 出力を比較

### Phase 2: ラッパー化

- `tagslist.sh` / `toptags.sh` を薄いラッパーへ変更
  - 例: `./bin/tagtools tagslist`
- 問題時は即座にシェル版へ戻せる構成にする

### Phase 3: 本切替

- `build.sh` から Go バイナリ直接呼び出し
- README と運用ドキュメントを更新

## 想定ディレクトリ構成

```text
tagtools/
  go.mod
  main.go
  internal/frontmatter/toml.go
  internal/tagslist/tagslist.go
  internal/toptags/toptags.go
  internal/csvutil/csvutil.go
  testdata/golden/tagslist.csv
  testdata/golden/toptags.json
```

## build.sh への最終反映イメージ

```sh
#!/bin/sh
./bin/tagtools toptags || exit 1
./bin/tagtools tagslist || exit 1
hugo --gc --cleanDestinationDir --destination=../text-publishd || exit 1
```

## 先行タスク（小さく始める）

1. `tagtools/main.go` + `tagtools tagslist` を実装（means引き継ぎ含む）
2. `github.com/goark/csvdata` で既存CSV読み込みを実装
3. ゴールデン比較を通す
4. `tagtools toptags` 実装
5. `all` と `verify` を追加
6. ラッパー切替

## 成功条件

- `tagslist.csv` と `toptags.json` が既存シェル版と一致する。
- `build.sh` / `publish.sh` の日常運用で追加差分・不整合が発生しない。
- エラー時の原因特定がシェル版より容易になる。
