+++
title = "ノーコード・プラットフォームとトレーディングカード"
date =  "2020-07-07T14:48:52+09:00"
description = "プレイヤーに与えられた「自由」はゲームという箱庭の内部にしか存在しない。箱庭自体を弄ることはできないし，箱庭の外に出る意義もない。"
image = "/images/attention/kitten.jpg"
tags = [ "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

面白い記事を見つけたので，私もポエムを書いてみる（笑）

- [2030年 「エンジニアです。コードは書けません。」｜__shinji__｜note](https://note.com/__shinji__/n/nde03573dc3a4)
- [プログラミングをする必要がなくなった後に人間に何が残るのか？ - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20200706/the-world-with-no-programming)

ちなみに私の親父が現役会社員だったあたりまでの時代にはプログラマとは別に「コーダー」という職業があって「[2030年 「エンジニアです。コードは書けません。](https://note.com/__shinji__/n/nde03573dc3a4)」で言うところの「パンチャー」に近い，ひたすら「コードを書く」だけの仕事だったそうな。

もっと言うと，私がペーペーの新人だった頃の某大手プロジェクトでは詳細設計を「馬鹿でもコードが書ける」くらいまでブレイクダウンさせることが要求されていた。
まぁ，私がちょっと昔に助っ人で入った Java 系プロジェクトでも詳細設計に SQL 文が書いてあって「この通りに入力しろ」と言われたことがあるが（笑）

実は「ノーコード・プラットフォーム」のようなものは1990年代に一大ブームがあって，当時は CASE と呼ばれていたが[^case1]，私もその手のプロジェクトに参加したことがある。
1990年代のブームには「産業界におけるオブジェクト指向の台頭」という背景があって，つまり「ノーコード・プラットフォーム」ってのはオブジェクト指向の究極というか「夢」なんだよね。

[^case1]: いま世の中にある CASE ツールは，当時の「夢」のかけらもしくはサブセットのようなものなので，実質は違うものと考えていただければ。

さて，最初の記事でも紹介されているように「ノーコード・プラットフォーム」が特に AI 分野で注目されている。

- [Googleはなぜノーコード開発ツールのAppSheetを買収し、1年半前に正式版になったばかりのApp Makerを終了させるのか？ － Publickey](https://www.publickey1.jp/blog/20/googleappsheet1app_maker.html)
- [AWSのノーコードツール「Amazon Honeycode」はなぜ生まれたのか？  |  TechCrunch Japan](https://techcrunch.com/2020/06/24/why-aws-built-a-no-code-tool/)

私のようなロートルから見れば「この道はいつか来た道♪」とか思ってしまうのだが，将来どうなるかは（私には）予想できない。
本当に「エンジニアです。コードは書けません」という未来がやってくるかもしれない。

「プログラミング」について物凄くものすごーく簡単に言うと「データと機能の状態を表現する」ことに尽きる（[フローチャート]({{< ref "/remark/2018/10/object-oriented-design.md" >}} "「オブジェクト指向」の黒歴史")は忘れよう`w`）。
「状態を表現」できるのであれば，穴ボコだろうが文字コードだろうが，もっと別の手段（電子ブロックとか`w`）でも構わないわけだ。
ただ，まぁ，本当に AI 技術が情報処理の主役になるのなら，ノイマン型のプラットフォームは衰退するだろうけどね。

20世紀末に夢見た CASE と比べて「ノーコード・プラットフォーム」は既に特定の「何か」から強い統制を受けている点が異なる。
これは「トレーディングカード」の比喩で考えると分かりやすい。

トレーディングカードの主な遊び方はデッキを構築して対戦することだ。
デッキの組み方は（定石やローカルルールはあるが）基本的に自由で，人によって無限の組み合わせがあると言っていいだろう。
しかしカードの種類とレアリティは配給企業がコントロールしていて「持たざるもの」に不利に働くようチューニングされている（故に皆「持」とうとする）。
言い換えればプレイヤーに与えられた「自由」はゲームという箱庭の内部にしか存在しない。
箱庭自体を弄ることはできないし，箱庭の外に出る意義もない。

「ノーコード・プラットフォーム」に限らずクラウド上の XaaS は等しく「箱庭」の境界問題を抱えている。
FOSS を通じて「コードの自由」を知っている私達が「ノーコード・プラットフォーム」の内側だけで満足できるのか。

まぁ「[インターネットの黄金時代]({{< ref "/remark/2017/01/internet-in-the-incubator.md" >}} "孵卵器の中のインターネット")」が云々とか言いつつ，結局は「中央集権型のインターネット」に甘んじている現状を見れば，トレーディングカード・ゲームと同じく「箱庭の外に出る意義もない」のかもしれないが（笑）

## ブックマーク

- [ブログ: ローコードとノーコードプラットフォームはプログラミングに革命をもたらすか?](https://okuranagaimo.blogspot.com/2020/02/blog-post_24.html)

## 参考図書

{{% review-paapi "B000ALJ18G" %}} <!-- 北原白秋の童謡集 -->
{{% review-paapi "B01G12HBDY" %}} <!-- 問題児たちが異世界から来るそうですよ？ -->
