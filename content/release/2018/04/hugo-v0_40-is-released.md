+++
title = "Hugo v0.40 がリリース"
date = "2018-04-23T21:45:35+09:00"
description = "ちうわけで，マジお疲れ様です。"
image = "/images/attention/tools.png"
tags  = [ "tools", "hugo" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

このブログのサイト・ジェネレータでもある [Hugo] v0.40 がリリースされた。

- [The Revival of the Shortcodes | Hugo](https://gohugo.io/news/0.40-relnotes/)
- [Release v0.40 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.40)

{{< fig-quote title="Release v0.40" link="https://github.com/gohugoio/hugo/releases/tag/v0.40" lang="en" >}}
<q>Hugo 0.40 is The Revival of the Shortcodes. Shortcodes is one of the prime features in Hugo. Really useful, but it has had some known shortcomings. <a href="https://github.com/bep">@bep</a> has been really busy the last week to address those.</q>
{{< /fig-quote >}}

ちうわけで，マジお疲れ様です。

{{% fig-quote title="Release v0.40" link="https://github.com/gohugoio/hugo/releases/tag/v0.40" lang="en" %}}
- `.Content` for a page retrieved in a query in a shortcode is now almost always available. Note that shortcodes can include content that can include shortcodes that can include content... It is possible to bite your tail. See more below.
- Shortcodes are now processed and rendered in their order of appearance.
- Related to the above, we have now added a zero-based `.Ordinal` to the shortcode.
{{% /fig-quote %}}

アップデートは計画的に。

## ブックマーク

- [ゼロから始める Hugo](/hugo/)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
