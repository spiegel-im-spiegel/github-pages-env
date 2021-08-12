+++
title = "#shimanego より： 継承できないなら注入すればいいじゃない！"
date =  "2020-01-23T22:49:25+09:00"
description = "今回は僭越ながら LT のひとつをやることになったので，ちょっとだけ頑張ったですよ（笑）"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "communication", "shimane", "engineering", "presentation", "site" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Shimane.go#03] に参加してきた。
今回は僭越ながら LT (Lightning Talk) のひとつをやることになったので，ちょっとだけ頑張ったですよ（笑） 

LT のスライドは以下に置いてある。

- [slide.Baldanders.info](https://slide.baldanders.info/)
    - [継承できないなら注入すればいいじゃない！ | slide.Baldanders.info](https://slide.baldanders.info/shimane-go-2020-01-23/)

サイト全体がスライドになっている。
右下の矢印をクリックするか `[Page Up]`/`[page Down]` または矢印キーの左右でページ送りできる。
スライド自体は [CC BY](https://creativecommons.org/licenses/by/4.0/ "Creative Commons — Attribution 4.0 International   — CC BY 4.0") で公開しているので，再利用は（条件の範囲内で）ご自由にどうぞ。

本当はスライドを書いた [Hugo] 環境も GitHub あたりに公開したかったが，デプロイ用スクリプトとかもリポジトリに含まれているので，公開は見送ることにした。
ゴメンペコン。

まぁ，私のはポエムみたいなものだが，他の方のは面白かった。
やっぱ [Go] をちゃんと仕事で使っている人は要所をキッチリ押さえた内容で，とても勉強になった。

あと，やっぱ「独り遊び」には限界があることを思い知らされた。
たとえ仕事で使ってなくとも，こうやって興味のあるイベントに参加して他の方の「仕事」を見るのはためになるし，インスピレーションも湧いてくる。
「参加することに意義がある」って，多分こういうときに使う言葉なんだろう。

ちうわけで，今後も可処分時間の許す限り参加していこうと思っております，はい。

---

以降，追加のポエム。

私は「数学ガール」シリーズの中の「制約が構造を生む」というフレーズが大好きだ。

{{< fig-quote type="markdown" title="数学ガール／フェルマーの最終定理" link="https://www.amazon.co.jp/dp/B00I8AT1CM?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}公理によって与えられる暗黙の制約。この制約が集合の要素同士をしっかり結びつける。単純にしばるのではない、相互に秩序ある関係を結ぶ。言い換えれば――公理によって与えられる制約が構造を生み出しているのだ{{% /quote %}}
{{< /fig-quote >}}

プログラミング言語としての [Go] で気に入ってることのひとつは，まさに [Go] という制約が興味深い構造を生み出している点である。
いつだって新しい言語は面白いのだ（まぁ [Go 言語]はそろそろ中堅に入るのかもしれないが`w`）。

いや，世の中「構造を守るために制約を課す」ケースが多すぎると思わない？ 逆だろ，ふつう。
少なくとも「開発」においては「構造を守るために制約を課」してはならないと思っている。
それは「停滞」と同義である。

## ブックマーク

- [第3回 Shimane.go の様子 - Togetter](https://togetter.com/li/1459564)

- [Hugo でスライド・サイトを立てる実験]({{< ref "/remark/2019/12/slide-site-by-hugo.md" >}})

[Shimane.go#03]: https://shimane-go.connpass.com/event/159977/ "Shimane.go#03 - connpass"
[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"


## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
