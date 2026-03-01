# Document Environment for GitHub Pages

GitHub Page [spiegel-im-spiegel.github.io](https://github.com/spiegel-im-spiegel/spiegel-im-spiegel.github.io)（[text.Baldanders.info](https://text.baldanders.info/)）の作業環境です。

- コンテンツの生成には [Hugo](https://gohugo.io/) を使用しています。
- [Hugo](https://gohugo.io/) の theme として [Baldanders.info] 専用 theme [spiegel-im-spiegel/hugo-theme-baldanders-info](https://github.com/spiegel-im-spiegel/hugo-theme-baldanders-info) を使用しています。

## ネタ帳

ネタは [Issues](https://github.com/spiegel-im-spiegel/github-pages-env/issues) に上げている。
またネタの進捗管理については「[ブログのネタ帳](https://github.com/spiegel-im-spiegel/github-pages-env/projects/1)」を参照のこと。

## ライセンス

コンテンツのライセンスは基本的に [Creative Commons Liscense Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/) で公開しています。ご利用はライセンスの許諾範囲内でご自由にどうぞ。

また Web フォントや外部から呼び出している CSS や JavaScript ファイル等のライセンスについてはそれぞれの発行元を確認してください。

[Baldanders.info]: https://baldanders.info/

## 新規投稿の作成（ローカル補助スクリプト）

- **スクリプト**: `new.sh`（`text` ディレクトリ直下）を使うとテンプレート（`archetypes`）を自動検出して `hugo new` を実行します。
- **使い方例**:

```bash
./new.sh remark                # content/remark/<YYYY>/<MM>/<DD>-stories.md を作成（archetype: remark があれば使用）
./new.sh bookmarks             # content/bookmarks/<YYYY>/<MM>/<DD>-bookmarks.md を作成（archetype: bookmarks があれば使用）
./new.sh release my-release.md # content/release/<YYYY>/<MM>/my-release.md を作成
./new.sh some/path.md          # 任意のパスでファイルを作成（archetype が無ければ default を使用）
```

- スクリプトは `archetypes/*.md` を参照して利用可能な archetype を検出します。
- `hugo new --kind=<kind> <path>` を実行するため、カスタム `archetypes` を追加するとその種類を利用できます。

変更したスクリプト: `new.sh` を安全に実行するために `set -euo pipefail` を有効にしています。
