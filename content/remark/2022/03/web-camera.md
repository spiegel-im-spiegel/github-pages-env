+++
title = "リモート会議用 Web カメラを買い替えた"
date =  "2022-03-06T18:06:07+09:00"
description = "マイクなんて飾りですよ。エラい人には分からんのです"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先月（2022-02）の話なのだが，仕事が佳境に入り始め，さらに[引っ越し]({{< ref "/remark/2022/02/moving-2022.md" >}})でプライベートも忙しくなったというのに，例の感染症絡みで自宅待機する羽目になりまして。
かといって仕事を休むわけにもいかないので，急遽リモートワークとなった。

まぁ，1日程度なら社内のリソースにアクセスする必要もなくセキュリティ的にもシビアな状況ではなかったので仕事自体は何とかなったんだけどね。
打ち合わせのために[自宅機に Teams を入れて]({{< ref "/remark/2022/02/installing-teams-in-ubuntu.md" >}} "Ubuntu に Microsoft Teams をインストールする")セットアップしたんだけど，カメラをONにすると音声がうまく入らない（またはノイズが酷い）。
結局，画面共有ができて音声がやり取りできればよかったので打ち合わせ自体は無問題だったのだが，どうも Web カメラが駄目ぽいことに気がついた。
そういえば読書会とかのオンラインイベントでも似たような現象が起きてたな。
Zoom でも Teams でも駄目ってなると[^tms1] やっぱ Web カメラだよなぁ...

[^tms1]: Teams はネット越しに音声をエコーバックして状態をテストする機能があって，そこで初めて自分の状態に気がついた。 Zoom はデバイスチェックはできるけどネット越しにエコーバック試験する機能はないよね？ この機能があったらもっと早く分かったのに...

手持ちの Web カメラのスペックをよく見たらマイク内蔵ぢゃん。お前かお前かお前か！
全く気が付かなかったよ `orz`

というわけで，マイクが付属しない Web カメラを物色し始めた。
最近の Web カメラって大抵マイクが付いてるんだな。
要らんじゃろ！

最終的に以下の Web カメラを発注

{{% review-paapi "B08TC3NR9L" %}} <!-- Web カメラ マイクなし -->

したのだが，自宅待機は（対象者の検査結果がシロだったため）一日で解除となり，以後は試す機会もなく昨日（2022-03-05）の[読書会](https://gpl-reading.connpass.com/event/238850/ "第21回『プログラミング言語Go』オンライン読書会（第2サイクル：8回目） - connpass")に臨んだのだが，他の人からは「聞こえにくい」とか「ノイズが酷い」といった苦情もなく上手くいったようだ。
よーし，うむうむ，よーし。

ただし，画角が狭いせいか顔がどアップになるんだよね。
自分で自分の映像を見て仰け反ってしまったよ。
雰囲気を和らげようとバーチャル背景を（私が普段壁紙で使ってる）天体写真にしたら絵面が[宇宙猫](https://dic.nicovideo.jp/a/%E5%AE%87%E5%AE%99%E7%8C%AB "宇宙猫とは (ウチュウネコとは) [単語記事] - ニコニコ大百科")みたいになってしまったので断念した（笑）

{{< fig-img-quote src="./431789693_096297b7d2_e.jpg" title="Boris in space by Andrew | Flickr (licensed CC-BY)" link="https://www.flickr.com/photos/polandeze/431789693/" lang="en" >}}

以前からネットメディアの記事とか動画とか見て思っているのだが，自分の顔のどアップなんてキツいと思わないのだろうか。
いや，まぁ，人それぞれか。

今の勤務先は様々な勤務形態に柔軟に対応できる職場で在宅勤務者も多いのだが，今回のように短期的にリモートワークにする想定はさすがになかったようで，今後をにらんでルール化しワークフローに落とし込んだ。
今回はいいテストケースになったのだろう。

## ブックマーク

- [また散財してしまった...]({{< ref "/remark/2021/08/web-meeting.md" >}})

## 参考

{{% review-paapi "B07RRQ59JR" %}} <!-- AfterShokz Aeropex 骨伝導ヘッドセット -->
{{% review-paapi "B08ND2M821" %}} <!-- Shokz OpenComm -->
{{% review-paapi "B09NRQNJ6R" %}} <!-- Daylight -->