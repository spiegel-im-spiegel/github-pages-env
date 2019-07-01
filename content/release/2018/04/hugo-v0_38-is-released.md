+++
title = "Hugo v0.38 がリリース"
date = "2018-04-02T21:03:37+09:00"
description = "Chroma Syntax Highlight のアップデートに対応したのは嬉しい。"
image = "/images/attention/tools.png"
tags  = [ "tools", "hugo" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

このブログのサイト・ジェネレータでもある [Hugo] v0.38 がリリースされた。

- [Release v0.38 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.38)

変更は多岐にわたるが

{{< fig-quote title="Release v0.38" link="https://github.com/gohugoio/hugo/releases/tag/v0.38" lang="en" >}}
<q>Hugo 0.38 is an Easter egg filled with good stuff. We now support fetching date and slug from the content filename, making the move from Jekyll even easier. And you can now set contentDir per language with intelligent merging, and themes can now provide configuration ... Also worth mentioning is several improvements in the Chroma highlighter, most notable support for Go templates.</q>
{{< /fig-quote >}}

なんだだそうだ。

[Chroma Syntax Highlight]({{< ref "/hugo/chroma-syntax-highlighting.md" >}} "Chroma Syntax Highlighting") のアップデートに対応したのは嬉しい。
Jekyll は使ったことないので今回の対応が嬉しいことなのか分からない（笑）

昔に比べて Jekyll のパフォーマンスはかなりよくなったと聞いた気がするのだが，気のせいかな。
実際どうなんだろう。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"

## ブックマーク

- [ゼロから始める Hugo](/hugo/)
