+++
title = "Hugo v0.38.2 がリリース"
date = "2018-04-10T20:28:59+09:00"
description = "不具合の修正がメイン。"
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

このブログのサイト・ジェネレータでもある [Hugo] v0.38.1 がリリースされた。

- [Hugo 0.38.2: Two Bugfixes | Hugo](https://gohugo.io/news/0.38.2-relnotes/)
- [Release v0.38.2 · gohugoio/hugo · GitHub](https://github.com/gohugoio/hugo/releases/tag/v0.38.2)

不具合の修正がメインのようだ。

{{% fig-quote type="markdown" title="Release v0.38.2" link="https://github.com/gohugoio/hugo/releases/tag/v0.38.2" lang="en" %}}
This is a bug-fix release with a couple of important fixes:

- Fix handling of the `--contentDir` and possibly other related flags 080302eb @bep #4589
- Fix handling of content files with "." in them 2817e842 @bep #4559

Also in this release:

- Set .Parent in bundled pages to its owner 6792d86a @bep #4582
{{% /fig-quote %}}

アップデートは計画的に。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
