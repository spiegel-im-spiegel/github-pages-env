+++
title = "Font Awesome 5 がリリースされた"
date =  "2017-12-12T20:22:50+09:00"
update = "2018-02-14T23:26:48+09:00"
description = "5 系では， Web フォントではなく， SVG と JavaScript を組み合わせた方法がオススメらしい。"
image = "/images/attention/remark.jpg"
tags        = [ "font", "svg", "javascript", "site" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

現在 Font Awesome のバージョンは 4.x 系と新しくリリースされた 5 系のふたつがあるようで， 5 系ではサポート・サイトも新しくなっている。

| Version | URL |
|:-------:|:----|
| 4.7 | [fontawesome.io](http://fontawesome.io/) |
| 5.0 | [fontawesome.com](https://fontawesome.com/) |

5 系では従来のフリー版の他に Pro 版も用意されていて， Pro 版では千個以上のアイコンや関連ツールが用意されている。

- [License | Font Awesome](https://fontawesome.com/license)
- [Font Awesome Pro | Font Awesome](https://fontawesome.com/pro)

フリー版では CDN による配信もあるようで， [BootstrapCDN](https://www.bootstrapcdn.com/fontawesome/ "Font Awesome · BootstrapCDN by MaxCDN") を使わなくともよくなった。

- [Get started with Font Awesome](https://fontawesome.com/get-started/web-fonts-with-css)

CSS はこんな感じで指定できる。

```html
<link href="https://use.fontawesome.com/releases/v5.0.1/css/all.css" rel="stylesheet">
```

5 系で最も大きく変わったのが HTML 上のアイコン名。
たとえば 4.x 系で PDF ファイルを表すアイコンは

```html
<i class="fa file-pdf-o"></i>
```

だったが， 5 系では

```html
<i class="far file-pdf"></i>
```

になった。
4.x 系から 5 系への変更については，以下が参考になる。

- [Upgrading from Version 4](https://fontawesome.com/how-to-use/upgrading-from-4)

アイコンのサイズ指定やアニメーション機能などについては従来のやり方がそのまま使えるようだが， 5 系では， Web フォントではなく， SVG と JavaScript を組み合わせた方法がオススメらしい。

- [SVG with JavaScript](https://fontawesome.com/how-to-use/svg-with-js)

アイコン制御用の JavaScript も CDN で用意されている。

```html
<script defer src="https://use.fontawesome.com/releases/v5.0.1/js/all.js"></script>
```

アイコンの指定の仕方は Web フォント版と同じ。

```html
<i class="fas fa-sync fa-2x fa-spin"></i>
```

{{< fig-quote >}}
<i class="fas fa-sync fa-2x fa-spin"></i>
{{< /fig-quote >}}

このサイトでも JavaScript 版を導入した。
[本家サイト]は... しばらく放置の方向で（笑）

## ブックマーク

- [【保存版】Font Awesomeの使い方：Webアイコンフォントを使おう](https://saruwakakun.com/html-css/basic/font-awesome)
- [アイコンフォントか？ SVG アイコンか？ — Website Usability Info](https://website-usability.info/2015/12/entry_151217.html)
- [Font Awesome 5 Freeで疑似要素(:after,:before)のcontent指定する場合 - Qiita](https://qiita.com/Garyuten/items/6d68da5cdac6dab9ba26)

[本家サイト]: http://www.baldanders.info/ "Baldanders.info"
[Font Awesome]: https://fontawesome.com/ "Font Awesome 5 | Font Awesome"
