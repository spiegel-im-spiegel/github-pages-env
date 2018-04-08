+++
title = "Hugo v0.38.1 がリリース"
date = "2018-04-08T09:08:21+09:00"
description = "不具合の修正がメインのようだ。"
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

このブログのサイト・ジェネレータでもある [Hugo] v0.38.1 がリリースされた。

- [Release v0.38.1 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.38.1)

不具合の修正がメインのようだ。

{{% fig-quote title="Release v0.38.1" link="https://github.com/gohugoio/hugo/releases/tag/v0.38.1" lang="en" %}}
This is a bug-fix that is mainly motivated by some issues with server live reloading introduced in Hugo 0.38.

- Fix livereload for the home page bundle f87239e4 @bep #4576
- Fix empty BuildDate in "hugo version" 294c0f80 @anthonyfok
- Fix some livereload content regressions a4deaeff @bep #4566
- Update github.com/bep/gitmap to fix snap build 4d115c56 @anthonyfok #4538
- Fix two tests that are broken on Windows 26f34fd5 @neurocline

This release also contains some improvements:

- Add bash completion 874159b5 @anthonyfok
- Handle mass content etc. edits in server mode 730b66b6 @bep #4563
{{% /fig-quote %}}

アップデートは計画的に。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
