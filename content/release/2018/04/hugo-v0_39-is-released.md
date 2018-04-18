+++
title = "Hugo v0.39 がリリース"
date =  "2018-04-18T09:26:31+09:00"
description = "「XXX」他"
image = "/images/attention/tools.png"
tags  = [ "tools" ]
draft = true

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

このブログのサイト・ジェネレータでもある [Hugo] v0.39 がリリースされた。

- [Release v0.39 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.39)

このバージョンでは `commands` パッケージを中心にかなりのリファクタリングが行われたようだ。
そのほかにも

{{< fig-quote title="Release v0.39" link="https://github.com/gohugoio/hugo/releases/tag/v0.39" lang="en" >}}
<q>But this release isn't all boring and technical: It includes several important bug fixes, several useful new template functions, and Resource.Content allows you to get any resource's content without having to fiddle with file paths and readFile.<br>
The main.Execute function now returns a Response object and the global Hugo variable is removed. This is only relevant for people building some kind of API around Hugo.</q>
{{< /fig-quote >}}

という感じでいろいろ変わっている。

## ブックマーク

- [ゼロから始める Hugo](/hugo/)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
