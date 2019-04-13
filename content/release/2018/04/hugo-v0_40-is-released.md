+++
title = "Hugo v0.40 がリリース"
date = "2018-04-23T21:45:35+09:00"
update = "2018-05-13T13:24:15+09:00"
description = "ちうわけで，マジお疲れ様です。"
image = "/images/attention/tools.png"
tags  = [ "tools", "hugo" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
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

このブログのサイト・ジェネレータでもある [Hugo] v0.40 がリリースされた。

- [The Revival of the Shortcodes | Hugo](https://gohugo.io/news/0.40-relnotes/)
- [Release v0.40 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.40)

{{< fig-quote title="Release v0.40" link="https://github.com/gohugoio/hugo/releases/tag/v0.40" lang="en" >}}
<q>Hugo 0.40 is The Revival of the Shortcodes. Shortcodes is one of the prime features in Hugo. Really useful, but it has had some known shortcomings. <a href="https://github.com/bep">@bep</a> has been really busy the last week to address those.</q>
{{< /fig-quote >}}

ちうわけで，マジお疲れ様です。

{{% fig-quote type="md" title="Release v0.40" link="https://github.com/gohugoio/hugo/releases/tag/v0.40" lang="en" %}}
- `.Content` for a page retrieved in a query in a shortcode is now almost always available. Note that shortcodes can include content that can include shortcodes that can include content... It is possible to bite your tail. See more below.
- Shortcodes are now processed and rendered in their order of appearance.
- Related to the above, we have now added a zero-based `.Ordinal` to the shortcode.
{{% /fig-quote %}}

アップデートは計画的に。

## 【追記 2018-04-28】 [Hugo] v0.40.1 がリリース

不具合修正版。

{{< fig-quote title="Release v0.40.1 · gohugoio/hugo" link="https://github.com/gohugoio/hugo/releases/tag/v0.40.1" lang="en" >}}
<q>This release fixes some shortcode vs <code>.Content </code> corner cases introduced in Hugo 0.40</q>
{{< /fig-quote >}}

だそうな。

## 【追記 2018-04-30】 [Hugo] v0.40.2 がリリース

- [Release v0.40.2 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.40.2)
- [Hugo 0.40.2: Two Bug fixes | Hugo](https://gohugo.io/news/0.40.2-relnotes/)

今回も不具合の修正のみ。

{{< fig-quote title="Release v0.40.2" link="https://github.com/gohugoio/hugo/releases/tag/v0.40.2" lang="en" >}}
<q>This release fixes some regressions introduced in Hugo 0.40. See the <a href="https://github.com/gohugoio/hugo/milestone/62?closed=1">milestone</a> for an issue list.</q>
{{< /fig-quote >}}

## 【追記 2018-05-10】 [Hugo] v0.40.3 がリリース

- [Hugo 0.40.3: One Bug Fix | Hugo](https://gohugo.io/news/0.40.3-relnotes/)
- [Release v0.40.3 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.40.3)

今回も不具合の修正。
レアケースのようだ。

{{< fig-quote title="Release v0.40.3" link="https://github.com/gohugoio/hugo/releases/tag/v0.40.3" lang="en" >}}
<q>Hugo 0.40.3 fixes a possible .Content truncation issue introduced in 0.40.1 90d0d830 @bep #4706. This should be very rare. It has been reported by only one user on a synthetic site. We have tested a number of big sites that does not show this problem with 0.40.2, but this is serious enough to warrant a patch release.</q>
{{< /fig-quote >}}

## ブックマーク

- [ゼロから始める Hugo](/hugo/)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
