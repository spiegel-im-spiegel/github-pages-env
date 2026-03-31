# Document Environment for GitHub Pages

GitHub Page [spiegel-im-spiegel.github.io](https://github.com/spiegel-im-spiegel/spiegel-im-spiegel.github.io)（[text.Baldanders.info](https://text.baldanders.info/)）の作業環境である。

- コンテンツの生成には [Hugo](https://gohugo.io/) を使用している。
- [Hugo](https://gohugo.io/) の theme として [Baldanders.info] 専用 theme [spiegel-im-spiegel/hugo-theme-baldanders-info](https://github.com/spiegel-im-spiegel/hugo-theme-baldanders-info) を使用している。

[Baldanders.info]: https://baldanders.info/

## ライセンス

コンテンツのライセンスは基本的に [Creative Commons Liscense Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/) で公開している。利用はライセンスの許諾範囲内で自由である。

また Web フォントや外部から呼び出している CSS や JavaScript ファイル等のライセンスについてはそれぞれの発行元を確認すること。

## 文章スタイル

- 日本語の文章では，読点に「，」を，句点に「。」を使用する。

## 新規投稿の作成（ローカル補助スクリプト）

- **スクリプト**: `new.sh`（`text` ディレクトリ直下）を使うとテンプレート（`archetypes`）を自動検出して `hugo new` を実行する。
- **使い方例**:

```bash
./new.sh remark                # content/remark/<YYYY>/<MM>/<DD>-stories.md を作成（archetype: remark があれば使用）
./new.sh bookmarks             # content/bookmarks/<YYYY>/<MM>/<DD>-bookmarks.md を作成（archetype: bookmarks があれば使用）
./new.sh release my-release.md # content/release/<YYYY>/<MM>/my-release.md を作成
./new.sh some/path.md          # 任意のパスでファイルを作成（archetype が無ければ default を使用）
```

- スクリプトは `archetypes/*.md` を参照して利用可能な archetype を検出する。
- `hugo new --kind=<kind> <path>` を実行するため，カスタム `archetypes` を追加するとその種類を利用できる。

### 運用メモ（remark で slug 指定）

- タイトル「aptitude をインストールする」の英語 slug には `install-aptitude` を採用した。
- `remark` セクションで slug をファイル名に使って作成する場合は次を実行する。

```bash
./new.sh remark install-aptitude.md
```

- 生成先: `content/remark/<YYYY>/<MM>/install-aptitude.md`
- 今回は front matter の追加・変更は行わない（生成された内容をそのまま使用する）。

変更したスクリプト: `new.sh` を安全に実行するために `set -euo pipefail` を有効にしている。

## 公開（ビルドと GitHub へのアップ）

- **スクリプト**: `publish.sh` は，サイトのビルドと公開用リポジトリへの push をまとめて実行する。
- **処理内容**:
	- 先に `./build.sh` を実行し，失敗時はその場で終了する。
	- `../text-publishd` へ移動して，`git add --all`，`git commit`，`git push -u origin master` を実行する。
- **使い方例**:

```bash
./publish.sh
./publish.sh "your commit message"
```

- 引数を省略した場合は，`Auto commit in 2026-03-31T03:00:00+00:00` のような UTC 時刻ベースのメッセージを使う。

### 運用メモ（最新コミットメッセージをそのまま使う）

- 直近コミットの件名を publish 用コミットメッセージとして使う場合は次を実行する。

```bash
./publish.sh "$(git log -1 --pretty=%s)"
```

- この手順で，サイトのビルド，`../text-publishd` でのコミット，`origin/master` への push まで実行できる。

## タグ集計（ローカル補助スクリプト）

- **スクリプト**: `tagslist.sh` は，`content/**/*.md` の front matter から `tags` を収集し，`.github/workflows/tagslist.csv` を再生成する。
- 出力形式は `tag,count` の CSV であり，カウント降順，カウント同値時は tag 名のアルファベット順でソートする。

- **スクリプト**: `toptags.sh` は，front matter の `date` が直近 1 年以内の記事を対象に `tags` を集計し，`data/toptags.json` を出力する。
- 出力形式は tag 名のみの JSON 配列であり，配列内はアルファベット順である。
- 既定では上位 15 件を出力し，`TOP_N` 環境変数で件数を変更できる。

## ビルド時の実行順序

- `build.sh` は先頭で `./toptags.sh` を実行し，失敗時はその場で中断する。
- 続いて `./tagslist.sh` を実行し，失敗時はその場で中断する。
- 最後に `hugo --gc --cleanDestinationDir --destination=../text-publishd` を実行する。
