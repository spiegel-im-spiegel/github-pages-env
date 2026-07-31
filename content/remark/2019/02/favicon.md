+++
title = "でっかい Favicon を設定する"
date = "2019-02-04T22:33:17+09:00"
description = "16×16, 32×32, 96×96, 192×192 サイズの画像ファイルを用意すればいい。簡単ぢゃん（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "web", "site" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は軽く小ネタで。
つか，覚え書き。

他所のサイトがでっかい favicon を設定してるのを見て「どうやってるんだろう」と思っていたが，なんのことはない， PNG 等の画像ファイルで 16×16, 32×32, 96×96, 192×192 各サイズの画像データを用意して `<link>` 要素で指定すればよかった。
こんな感じ。

```html
<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
<link rel="icon" type="image/png" sizes="96x96" href="/favicon-96x96.png">
<link rel="icon" type="image/png" sizes="192x192"  href="/favicon-192x192.png">
```

簡単ぢゃん（笑）

## ブックマーク

- [Favicon & App Icon Generator](https://www.favicon-generator.org/)
