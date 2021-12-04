+++
title = "GNKF: NKF ぽいなにか の v0.6.3 をリリースした"
date =  "2021-09-18T18:40:49+09:00"
description = "全角・半角変換とかなカナ変換に対応できてない仮名文字があることに気づいたので修正した。"
image = "/images/attention/tools.png"
tags  = [ "tools", "gnkf", "golang", "unicode", "character" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] における文字列処理の習作 [gnkf] の v0.6.3 をリリースした。

- [Release v0.6.3 · spiegel-im-spiegel/gnkf · GitHub](https://github.com/spiegel-im-spiegel/gnkf/releases/tag/v0.6.3)

さっき上げたブログ記事

- [Unicode のカタカナを調べる]({{< ref "/golang/unicode-katakana.md" >}})

から，全角・半角変換とかなカナ変換に対応できてない仮名文字があることに気づいたので修正した。

なお [gnkf] では変体仮名は無視することにした。
アレは（歴史的経緯はともかく文字処理としては）異体字というより「そういう文字」と考えたほうがよさそう。
Unicode 正規化でも通常の仮名文字への変換はしないみたいだし。

まぁ，個人的に需要が発生すれば対応するということで。

## ブックマーク

- [GNKF: Network Kanji Filter by Golang]({{< ref "/release/gnkf.md" >}})

[Go]: https://go.dev/
[gnkf]: https://github.com/spiegel-im-spiegel/gnkf "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
