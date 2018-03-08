+++
title = "Hugo v0.37.1 Released"
date =  "2018-03-08T09:59:38+09:00"
description = "今回は v0.37 で発生した regression の修正。"
image = "/images/attention/tools.png"
tags  = [ "tools", "hugo" ]
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

[Hugo] v0.37.1 がリリースされた。

- [Release v0.37.1 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.37.1)

今回は v0.37 で発生した regression の修正で

{{< fig-quote title="Release v0.37.1" link="https://github.com/gohugoio/hugo/releases/tag/v0.37.1" lang="en" >}}
<q>Image content such as SVG cannot be scaled with the built-in image processing methods, but it should still be possible to use them as page resources. This was a regression in Hugo 0.37 and is now fixed.</q>
{{< /fig-quote >}}

とのことなので，該当する方はアップデートを。
私は自前で画像表示の shotcode を書いてるので無問題。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
