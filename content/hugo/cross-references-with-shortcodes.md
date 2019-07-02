+++
title = "Hugo におけるサイト内相互参照"
date = "2018-08-18T14:46:11+09:00"
lastmod = "2018-08-18T14:46:11+09:00"
description = "Hugo 0.45 から ref/relref の仕様が変わったので紹介する。"
image = "/images/attention/hugo.jpg"
tags = [ "hugo", "shortcodes" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Hugo] ではサイト内の相互参照（リンク）を表現するために組み込み [shortcode] である `ref` や `relref` を使うのだが [Hugo] 0.45 から `ref`/`relref` の仕様が変わったので紹介する。

例えば参照したい記事を `content/blog/entry.md` とする。
他の記事からこの記事を参照したい場合， [Hugo] 0.45 より前のバージョンでは

```text
{{</* ref "blog/entry.md" */>}}
```

または

```text
{{</* relref "blog/entry.md" */>}}
```

のように作業環境の `content` フォルダ以下の相対パスを指定すればよかった。
これにより `ref` コマンドなら `http(s)://site.name/blog/entry/` に， `relref` コマンドなら `/blog/entry/` に展開されていた。

[Hugo] 0.45 からは ビルド後のディレクトリ構成で考える。
たとえば

```text
{{</* ref "/blog/entry.md" */>}}
```

あるいは

```text
{{</* relref "entry.md" */>}}
```

といった感じ（サイトのルート・ディレクトリからの絶対パスを指定する場合 `blog/...` ではなく `/blog/...` となる点に注目）。
したがって [Hugo] 0.45 より前の古い表記では warning が発生するので注意が必要である。
また  `relref` コマンドを使う際に絶対パスを指定した場合も warning になる。

なお，このルールはテンプレートにも適用される。
たとえば

```text
{{ with .Site.GetPage "/blog/entry.md" }}{{ .Title }}{{ end }}
```

などとすれば任意の記事の情報を取得できる。

## ブックマーク

- [Hugo 0.45: Revival of ref, relref and GetPage | Hugo](https://gohugo.io/news/0.45-relnotes/)
- [Links and Cross References | Hugo](https://gohugo.io/content-management/cross-references/)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[shortcode]: https://gohugo.io/extras/shortcodes/ "Shortcodes | Hugo"
