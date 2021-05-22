+++
title = "GNKF 最初のリリース "
date =  "2020-07-30T22:49:26+09:00"
description = "元のコードの70%くらいは捨てました（笑）"
image = "/images/attention/go-logo_blue.png"
tags  = [ "tools", "gnkf", "golang", "character", "encoding", "unicode", "normalization" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

貴方は3年前に書いた自分のコードにウンザリしたことはありませんか。
私はあります。
それもしょっちゅう（笑） まぁ，3年前と今とで同じコードを書いてたら，それはそれで進歩がないっちう話だけど。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}。

3年前に最初のリリースを公開した「[nkf っぽいなにか]({{< ref "/remark/2017/12/like-nkf.md" >}} "「nkf っぽいなにか」を作った")」はあまりにコードの出来が悪すぎてメンテする気も起こらなくなったので，[リポジトリを放棄](https://github.com/spiegel-im-spiegel/text "spiegel-im-spiegel/text: Encoding/Decoding Text Package by Golang")して新たに作り直すことにした（☆を付けてくださってる方，ごめんなさい）。

- [spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang](https://github.com/spiegel-im-spiegel/gnkf)

元のコードの70%くらいは捨てました（笑）

詳しい使い方は以下を参考にどうぞ。

- [GNKF: Network Kanji Filter by Golang]({{< ref "/release/gnkf.md" >}})

これでようやく「[かなカナ変換]({{< ref "/golang/kana-conversion.md" >}})」や「[康熙部首の正規化]({{< ref "/golang/unicode-kangxi-radical.md" >}} "こんな埼「玉」修正してやるぅ")」を組み込めた。

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
