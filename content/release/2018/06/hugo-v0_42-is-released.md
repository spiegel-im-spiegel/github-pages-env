+++
title = "Hugo v0.42 がリリース"
date = "2018-06-24T18:50:58+09:00"
description = "テーマをコンポーネント化できるようになったらしい。"
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

このブログのサイト・ジェネレータでもある [Hugo] v0.42 がリリースされた。

- [Release v0.42 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.42)

{{< fig-quote title="Release v0.42" link="https://github.com/gohugoio/hugo/releases/tag/v0.42" lang="en" >}}
<q>Hugo 0.42 adds Theme Components. This is a really powerful new way of building your Hugo sites with reusable components. This is both Theme Composition and Theme Inheritance.</q>
{{< /fig-quote >}}

というわけで，テーマをコンポーネント化できるようになったらしい。
こんな感じに複数のテーマを指定できる（従来の単一指定でも問題ないようだ）。

```toml
theme = ["my-shortcodes", "base-theme", "hyde"]
```

確かに shortcodes だけを集めたテーマとか作ったら面白いかな？

詳しくは以下のドキュメントを参考のこと。

- [Theme Components | Hugo](https://gohugo.io/themes/theme-components/)

## 【追記 2018-06-24】 [Hugo] v0.42 がリリース

- [Release v0.42.1 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.42.1)

不具合修正版。

{{% fig-quote-md title="Release v0.42.1 · gohugoio/hugo" link="https://github.com/gohugoio/hugo/releases/tag/v0.42.1" lang="en" %}}
- Reset the global pages cache on server rebuilds 128f14ef @bep #4845
- Do not fail server build when /static is missing 34ee27a7 @bep #4846
{{% /fig-quote-md %}}

だそうな。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
