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

## 独立リポジトリ化チェックリスト

### 1. リポジトリ設計

- `tagtools` 専用リポジトリを作成する。
- `go.mod` の module path を実リポジトリ名に合わせる（例: `github.com/<owner>/tagtools`）。
- CLI 専用運用か，ライブラリ公開も含むかを先に決める。

### 2. リリース方針

- SemVer でタグ運用する（`v0.x` -> `v1.0.0`）。
- 破壊的変更時のルール（メジャー更新）を明文化する。
- 配布形式（GitHub Releases のバイナリ，`go install`，ソース利用）を決める。

### 2.1 破壊的変更とメジャー更新ルール

- 破壊的変更とみなす条件:
  - CLI サブコマンド名の変更または削除（例: `tagslist` / `toptags` / `verify`）。
  - 既存オプションの名前変更，削除，既定値変更（運用結果に影響するもの）。
  - 既存環境変数互換の破壊（例: `TOP_N` / `TOPTAGS_WINDOW` の無効化，`TAGTOOLS_*` 解釈変更）。
  - 出力フォーマット変更（`tagslist.csv` 列定義，`toptags.json` 形式，`verify --debug` JSON キー）。
  - 非互換な実行終了コードの変更（運用スクリプト連携に影響するもの）。

- メジャー更新ルール:
  - 上記の破壊的変更を含む場合は必ずメジャーバージョンを上げる。
  - メジャー更新時はリリースノート先頭に「Breaking Changes」を明示する。
  - `github-pages-env` 側の `TAGTOOLS_VERSION` 既定値は，互換検証完了後にのみ更新する。

- マイナー/パッチ更新ルール:
  - 後方互換を維持する機能追加はマイナー更新。
  - バグ修正・セキュリティ修正・内部改善のみはパッチ更新。

- 導入前チェック（既存リポジトリ側）:
  1. `TAGTOOLS_VERSION=latest` で試験実行。
  2. `tagslist.csv` / `toptags.json` の差分確認。
  3. `./build.sh` と `./publish.sh` のスモーク確認。
  4. 問題なければ固定版（pin）を更新。

### 3. セキュリティ運用

- 依存更新を自動化する（Dependabot または Renovate）。
- CI に `go test` / `go vet` / `govulncheck` を組み込む。
- リリース時にチェックサムを公開する。
- チェックサム公開を完全性保証の必須要件とし，署名（GPG/cosign）は将来拡張の任意要件とする。

### 4. テスト資産の分離

- 現在のサイト実データ依存テストと，独立リポジトリ内 fixture テストを分離する。
- 独立リポジトリ側には最小 fixture を持たせる。
- 既存リポジトリ側は統合テスト（実コンテンツ前提）を維持する。

### 5. 既存リポジトリとの接続

- `tagslist.sh` / `toptags.sh` の互換インターフェースを維持する。
- `TOP_N` など既存環境変数との互換を壊さない。
- `build.sh` で実行元を切り替え可能にして段階導入する。

### 6. 移行手順（推奨）

1. 新規リポジトリ作成と初期 CI 設定。
2. 現行 `tagtools/` を移植し，module path を更新。
3. `v0.x` で試験リリースし，既存リポジトリで検証。
4. 問題がなければ `v1.0.0` を作成。
5. 既存リポジトリのラッパーを外部配布版参照に切り替え。

### 7. ロールバック方針

- 問題発生時は `tagtools/orgsh/` の旧シェル実装へ即時切り戻し可能な状態を維持する。
- 切り戻し手順（どのファイルを戻すか，どのコマンドで検証するか）を運用手順書に明記する。

### 7.1 ロールバック発動条件

- `build.sh`，`tagslist.sh`，`toptags.sh` のいずれかで外部配布版 `tagtools` 実行に失敗し，軽微修正で即時復旧できない場合。
- 生成物差分が想定外で，`tagslist.csv` または `toptags.json` の互換性が維持できない場合。
- 本番運用（記事公開前チェックを含む）で再現性のない失敗が継続する場合。

### 7.2 ロールバック責任者

- 一次判断者: リポジトリ運用者（owner）。
- 実行担当: 当該作業ブランチの変更実施者。
- 承認ルール: 一次判断者が「外部配布版利用を一時停止」と判断した時点で即時実施。

### 7.3 ロールバック手順（即時）

1. 外部配布版利用を停止し，旧シェル実装に戻す。
  - `tagtools/orgsh/tagslist.sh` を `tagslist.sh` に復元。
  - `tagtools/orgsh/toptags.sh` を `toptags.sh` に復元。
  - `build.sh` の `TAGTOOLS_SOURCE=external` 前提ロジックを戻す（必要に応じて直前コミットを revert）。
2. 実行権限を確認する。
  - `chmod +x tagslist.sh toptags.sh`
3. 生成処理とビルドを再実行する。
  - `./toptags.sh`
  - `./tagslist.sh`
  - `./build.sh`

### 7.4 ロールバック後の検証

- `git diff -- .github/workflows/tagslist.csv data/toptags.json` で生成差分を確認する。
- `./publish.sh` 前提の通常運用フローが通ることを確認する。
- 問題が解消したら，ロールバック理由・再現手順・暫定対処を Issue に記録する。

### 7.5 復帰条件（再切替）

- 原因が特定され，外部配布版で同一条件の再現試験が通ること。
- `TAGTOOLS_VERSION=latest` 試験と固定版試験の両方で互換性確認が完了していること。
- 変更を段階導入（`TAGTOOLS_EXEC=go` のスモーク試験 -> 通常運用）して問題が再発しないこと。
