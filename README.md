# Document Environment for GitHub Pages

GitHub Page [spiegel-im-spiegel.github.io](https://github.com/spiegel-im-spiegel/spiegel-im-spiegel.github.io)（[text.Baldanders.info](https://text.baldanders.info/)）の作業環境である。

- コンテンツの生成には [Hugo](https://gohugo.io/) を使用している。
- [Hugo](https://gohugo.io/) の theme として [Baldanders.info] 専用 theme [spiegel-im-spiegel/hugo-theme-baldanders-info](https://github.com/spiegel-im-spiegel/hugo-theme-baldanders-info) を使用している。

[Baldanders.info]: https://baldanders.info/

## この README の位置づけ

- この README は，人間向けの入口ドキュメントである。
- 日常的な運用ルールや例外判断は `.github/copilot-instructions.md` を正本とする。
- README は，概要，最小手順，参照先の案内に絞る。

## クイックスタート

### 新規投稿の作成

```bash
./new.sh remark
./new.sh release my-release.md
./new.sh some/path.md
```

### ビルド

```bash
./build.sh
```

### 公開

```bash
./publish.sh
./publish.sh "your commit message"
```

### Hugo 更新

```bash
./hugo_inst.sh
hugo version
```

## ブログセクション一覧

- `remark`: 雑記，検証メモ，運用ノウハウ，日々の気づきをまとめる主力セクション。
- `release`: ソフトウェアや規格などのリリース情報を紹介・整理するセクション。
- `bookmarks`: 期間ごとのリンク集や，気になった記事・資料の記録セクション。
- `openpgp`: OpenPGP，GnuPG，暗号運用まわりの話題を扱う技術セクション。
- `golang`: Go 言語と関連ツール，実装メモを扱うセクション。
- `rust-lang`: Rust 言語と関連エコシステムを扱うセクション。
- `hugo`: Hugo 本体やテンプレート運用，サイト構築まわりのセクション（現在休載中）。
- `typst`: Typst と文書作成ワークフローを扱うセクション。
- `cc-licenses`: Creative Commons を中心としたライセンス関連のセクション。

- セクションの運用ルール（front matter，命名，deploy 手順など）は `.github/copilot-instructions.md` を参照すること。

## 運用ドキュメントの地図

- `.github/copilot-instructions.md`
  - Copilot が従う運用ルールの本体。
  - deploy ルーティーン，Hugo 更新手順，後処理ルール，AI ブロック運用を定義。
- `README.md`（このファイル）
  - 作業の入口，最小手順，参照先の案内。

### `.github/copilot-instructions.md` のセクション概要

- `Writing Style`: 日本語の句読点や文体ルール。
- `Local Commands`: ローカル補助コマンドの目的と使い方。
- `Deploy Playbook`: 記事公開の標準手順，事前確認，deploy 後の後処理。
- `Archetypes and Front Matter`: archetype と front matter の運用方針。
- `AI Block Rules`: `div-ai` の使用ルールと既存記事の移行方針。

## ローカル補助スクリプト一覧

- `new.sh`: `archetypes` を見て `hugo new --kind=<kind> <path>` を実行。
- `build.sh`: `toptags.sh` と `tagslist.sh` 実行後に Hugo でビルド。
- `publish.sh`: `build.sh` の後，`../text-publishd` で commit/push。
- `tagslist.sh`: front matter の tags から `.github/workflows/tagslist.csv` を再生成。
- `toptags.sh`: 直近 1 年の記事 tags から `data/toptags.json` を再生成。
- `hugo_inst.sh`: `gohugoio/hugo` の最新 `hugo_extended` `.deb` で Hugo を更新。

## 補足

- `publish.sh` 実行後に `.github/workflows/tagslist.csv` や `data/toptags.json` が source 側に残る場合の扱いは `.github/copilot-instructions.md` に従う。
- 記事作成時の front matter 運用（description，tags，slug など）も `.github/copilot-instructions.md` に従う。

## ライセンス

コンテンツのライセンスは基本的に [Creative Commons Liscense Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/) で公開している。利用はライセンスの許諾範囲内で自由である。

また Web フォントや外部から呼び出している CSS や JavaScript ファイル等のライセンスについてはそれぞれの発行元を確認すること。

## 文章スタイル

- 日本語の文章では，読点に「，」を，句点に「。」を使用する。
