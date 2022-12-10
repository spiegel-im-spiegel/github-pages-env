+++
title = "いまさら『Clean Architecture』の感想文"
date =  "2022-11-23T17:05:54+09:00"
description = "私はちょいちょい読み返している。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

手遊びで Go でバッチ処理を書いてるのだが，ちょっと作業に飽きてきたので，前々から書こうと思っていた『[Clean Architecture]』の感想文をやっつけてしまおう。

Robert C. Martin（通称ボブおじさん）による『Clean ◯◯』シリーズのひとつが『[Clean Architecture]』である。
この本の中身については以下の解説記事が分かりやすい。

- [世界一わかりやすいClean Architecture - nuits.jp blog](https://www.nuits.jp/entry/easiest-clean-architecture-2019-09)

本記事を読むくらいなら上のリンク先記事を読むことを強くお勧めする。
以上！

では面白くないので，続きを。

『[Clean Architecture]』の面白いところは，各トピックが非常に短い文に凝縮されているところ。
たとえば「構造化プログラミング」「オブジェクト指向プログラミング」「関数型プログラミング」の3つのプログラミング・パラダイムについて

{{< fig-quote type="markdown" title="『Clean Architecture』第3章" link="https://www.amazon.co.jp/dp/B07FSBHS2V?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
構造化プログラミングは、直接的な制御の移行に規律を課すものである。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『Clean Architecture』第3章" link="https://www.amazon.co.jp/dp/B07FSBHS2V?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
オブジェクト指向プログラミングは、間接的な制御の移行に規律を課すものである。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『Clean Architecture』第3章" link="https://www.amazon.co.jp/dp/B07FSBHS2V?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
関数型プログラミングは、代入に規律を課すものである。
{{< /fig-quote >}}

などと，ざっくりした文で締めくくっている。
いや，ざっくりしすぎだろう（笑） こういう解説を読んでしまうと， GOTO は悪とか継承がどうのとか多態化だの単態化だの不変だのそれっぽい用語を並べて語るのが馬鹿らしくなる。

世の中に山ほどのプログラミング言語があるが，すべてこの3つのパラダイムの組み合わせで成り立っていて，これらが課す規律または制約がプログラム・コード→コンポーネント→アーキテクチャにどう影響を及ぼすのか，というのが言語を語る（騙る？）際の最重要ポイントだと思う。

『[Clean Architecture]』の主題である「アーキテクチャ」にしてもそう。
この本ではアーキテクチャを「システムに与えた「形状」である」とし，アーキテクチャの形状の目的は「ソフトウェアシステムの開発・デプロイ・運用・保守を容易にすることである」と述べ，さらに

{{< fig-quote type="markdown" title="『Clean Architecture』第15章" link="https://www.amazon.co.jp/dp/B07FSBHS2V?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
それらを容易にするための戦略は、できるだけ長い期間、できるだけ多くの選択肢を残すことである。
{{< /fig-quote >}}

と締めくくる。
ここで私はあの名作を思い出すのである。

{{< fig-quote title="数学ガール／フェルマーの最終定理" link="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/baldandersinf-22/" >}}
<q>公理によって与えられる暗黙の制約。この制約が集合の要素同士をしっかり結びつける。単純にしばるのではない、相互に秩序ある関係を結ぶ。言い換えれば――公理によって与えられる制約が構造を生み出しているのだ</q>
{{< /fig-quote >}}

つまり，プログラミング・パラダイムによって与えられる「制約」が巡り巡ってアーキテクチャという「構造」を生み出す。
おー，円環が繋がったぞ（笑）

個人的には「第32章 フレームワークは詳細」の

{{< fig-quote type="markdown" title="『Clean Architecture』第32章" link="https://www.amazon.co.jp/dp/B07FSBHS2V?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
フレームワークなんかと結婚するな！
{{< /fig-quote >}}

でバカウケしてしまった。

以前に『ユニコーン企業のひみつ』の[読書会の話を書いた]({{< ref "/remark/2022/05/competing-with-unicorns.md" >}} "『ユニコーン企業のひみつ』読書会（1）")が，私は『[Clean Architecture]』も同時に読み返していた。
プロダクトのイテレーションを継続するためにはアーキテクチャの「できるだけ多くの選択肢を残す」戦略が必須だと思ったからだ。

こんな感じで，私はちょいちょい読み返している。
特に設計に迷ったときとか，決して「正解」を教えてくれるわけではないけど，プロダクトやシステムにとって何がコアで何が「仔細」なのか，設計が外道に落ちないようストッパーとしてこの本は効いている。

## 【2022-12-10 追記】「ちょうぜつエンジニアめもりーちゃん」を読め！

祝♪ [#ちょうぜつエンジニアめもりーちゃん](https://twitter.com/search?q=%23%E3%81%A1%E3%82%87%E3%81%86%E3%81%9C%E3%81%A4%E3%82%A8%E3%83%B3%E3%82%B8%E3%83%8B%E3%82%A2%E3%82%81%E3%82%82%E3%82%8A%E3%83%BC%E3%81%A1%E3%82%83%E3%82%93) 単行本化！

というわけで，『[Clean Architecture]』じゃフワッとしていてよく分からんって方やこの記事の最初に挙げたリンク先よりもう少し詳細な解説書はないのか？ って方には田中ひさてるさんの『[ちょうぜつソフトウェア設計入門――PHPで理解するオブジェクト指向の活用](https://www.amazon.co.jp/dp/B0BNH1J2W2?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』を併せて読むことをお勧めする。

特に第9章コラムにある

{{< fig-quote type="markdown" title="『ちょうぜつソフトウェア設計入門』第9章" link="https://www.amazon.co.jp/dp/B0BNH1J2W2?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
パラダイムがじっくり待ってくれているこの時代は、学習者にとっては大きなチャンスです。変化が激しすぎてついていけないなんてことは、もうありません。必要なことは出そろっていて、いい意味で枯れたものを学べます。
{{< /fig-quote >}}

というのは本当にそのとおりだと思う。
こういった良書を時間をかけて理解していくことが大事である。

## ブックマーク

- [プログラミングを独習するには10年かかる（Teach Yourself Programming in Ten Years 日本語訳）](https://www.yamdas.org/column/technique/21-daysj.html)

[Clean Architecture]: https://www.amazon.co.jp/dp/B07FSBHS2V?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Clean Architecture　達人に学ぶソフトウェアの構造と設計 (アスキードワンゴ) | Ｒｏｂｅｒｔ Ｃ．Ｍａｒｔｉｎ, 角 征典, 高木 正弘 | 工学 | Kindleストア | Amazon"

## 参考文献

{{% review-paapi "B07FSBHS2V" %}} <!-- Clean Architecture -->
{{% review-paapi "B0BNH1J2W2" %}} <!-- ちょうぜつエンジニアめもりーちゃん -->
{{% review-paapi "4873119464" %}} <!-- ユニコーン企業のひみつ -->
{{% review-paapi "B00I8AT1CM" %}} <!-- 数学ガール／フェルマーの最終定理 -->
