+++
title = "Hugo で Git 情報を取得する"
date = "2019-02-16T22:15:35+09:00"
lastmod = "2019-02-16T22:15:35+09:00"
description = "Hugo の環境を git で管理している場合はコミット情報等をテンプレートに組み込むことができる。"
image = "/images/attention/hugo.jpg"
tags = [ "hugo", "site", "git" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[本家サイト]を改装するにあたって [Hugo] の環境を見直している。
この作業でやったことをこのセクションで少しずつ紹介していこうと思う。

今回は [git] との連携の話。

## .GitInfo

[Hugo] の環境を [git] で管理している場合は以下のコミット情報等をテンプレートに組み込むことができる。

| テンプレート変数           | 内容                   |
| -------------------------- | ---------------------- |
| `.GitInfo.AuthorName`      | ユーザ名               |
| `.GitInfo.AuthorEmail`     | ユーザのメールアドレス |
| `.GitInfo.AuthorDate`      | コミット時刻           |
| `.GitInfo.Hash`            | リビジョン値           |
| `.GitInfo.AbbreviatedHash` | リビジョン値（短縮値） |
| `.GitInfo.Subject`         | コミット・メッセージ   |

`.GitInfo` を有効にするには [Hugo] コマンド起動時に `--enableGitInfo` オプションを付加するか設定ファイル（`config.toml` など）に

```toml
enableGitInfo = true
```

と設定すればよい。

## ページ更新時刻 .Lastmod をコミット時刻にする

`.GitInfo` が有効であればページの更新時刻を示すテンプレート変数 `.Lastmod` の値がコミット時刻になる。

テンプレート変数 `.Lastmod` は，既定ではページの front matter 内の以下の変数にマッピングされている。

```toml
[frontmatter]
  lastmod = [":git", "lastmod", "date", "publishDate"]
```

配列の順番がそのまま優先順位になっている。
つまり `":git"` がコミット時刻 `.GitInfo.AuthorDate` に対応する。

この設定は設定ファイルで変更可能なので，たとえば `.GitInfo` を有効にしたいがコミット時刻を更新時刻に使いたくない場合は

```toml
[frontmatter]
  lastmod = ["lastmod", "date", "publishDate"]
```

などとしてしまってもよい。

ちなみに入力ファイルのタイムスタンプを更新時刻にする場合は

```toml
[frontmatter]
  lastmod = [":fileModTime"]
```

などとする。
まぁ今どきファイルのタイムスタンプを更新時刻にするのはオススメではないが。


[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[本家サイト]: https://baldanders.info/ "Baldanders.info"
[git]: https://git-scm.com/
[Git]: https://git-scm.com/
