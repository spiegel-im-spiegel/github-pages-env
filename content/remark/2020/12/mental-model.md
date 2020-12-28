+++
title = "「無人化システム」とメンタル・モデル"
date =  "2020-12-28T21:31:19+09:00"
description = "20世紀なコードはもうケッコウ"
image = "/images/attention/kitten.jpg"
tags = [ "engineering" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

先週のことで恐縮だが

- [COBOLのコードは未だに我々の金を握っており、バリバリ現役である - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20201221/cobol-controls-your-money)

という記事が公開されていて，これを読んで思い出したのが20世紀に参加した某プロジェクトでの世間話。

もう四半世紀前のうろ覚えでしかも{{% ruby "ひとづて" %}}人伝{{% /ruby %}}に聞いた内容だが，当時は意図的に難読化コードを書くプログラマも多かったらしい。
理由は単純で，コードを「属人化」してしまえば少なくともそのプロジェクトに関してはクビにならないから。

というわけで COBOL プログラマはきっと「[まだだ！ まだ終わらんよ！](https://dic.pixiv.net/a/%E3%81%BE%E3%81%A0%E3%81%A0%21%E3%81%BE%E3%81%A0%E7%B5%82%E3%82%8F%E3%82%89%E3%82%93%E3%82%88%21)」と思っているだろう（笑）

## 「無人化システム」

一方，上の記事が出る少し前に

- [「無人化システム」を駆逐する組織マネジメントとエンジニアリング](https://zenn.dev/tmknom/articles/93f227ad5e55aa)

というのが Zenn で公開されていて，あまりの納得感に{{% ruby "サポート" %}}投げ銭{{% /ruby %}}してしまったのだが（笑），記事では「無人化システム」を

{{< fig-quote type="markdown" title="「無人化システム」を駆逐する組織マネジメントとエンジニアリング" link="https://zenn.dev/tmknom/articles/93f227ad5e55aa" >}}
システム運用が属人化し、かつその運用者が退職するとシステムが無人化します。我々の会社ではこのようなシステムを『**無人化システム**』と呼んでいます。

無人化システムは「**誰も詳細は知らないが、なぜか動いているシステム**」です。
{{< /fig-quote >}}

と定義しているようだ。

「傭兵」時代は「汎用機＋COBOL」からのリプレイス案件を時々受けていたが，何が困るって，ドキュメント化されない「誰も知らないコード」が平気で紛れ込んでいて[^revise1]，しかもそのコードに手を出すとどんな影響が出るか予測できないという事態にホンマに困っていた（あと，どうやっても正規化できないデータベースとか`w`）。

[^revise1]: 勿論ちゃんとしてる企業もあるよ。コードを1行修正するのにも2重3重のレビューを行って，変更申請書が受理されないと変更できない，みたいなガチガチの企業もあったな。

まっ，要するに，時代や言語に関係なく，この手の話は割と普遍的に観測できるということなんだろう。

## メンタル・モデル

2017年に公開された記事だが

- [Design Philosophy On Data And Semantics](https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html)

というのを最近読んだ。
特に [Go] でコードを書く人はこの記事は必読だろう。

いくつか拾い読みしてみる。

{{< fig-quote type="markdown" title="Design Philosophy On Data And Semantics" link="https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html" lang="en" >}}
{{% quote %}}A consistent use of value/pointer semantics, for a given type of data, is critical if you want to maintain integrity and readability throughout your software. Why? Because, if you are changing the semantics for a piece of data as it is passed between functions, you are making it difficult to maintain a clear and consistent mental model of the code. The larger the code base and the team becomes, the more bugs, data races and side effects will creep unseen into the code base{{% /quote %}}.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Design Philosophy On Data And Semantics" link="https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html" lang="en" >}}
{{% quote %}}Tom has also mentioned that a box of copy paper can hold 100k lines of code. Take a second to let that sink in. For what percentage of the code in that box could you maintain a mental model of?{{% /quote %}}
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Design Philosophy On Data And Semantics" link="https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html" lang="en" >}}
{{% quote %}}“The hardest bugs are those where your mental model of the situation is just wrong, so you can’t see the problem at all” - Brian Kernighan{{% /quote %}}
{{< /fig-quote >}}

これ以降は [Go] をターゲットにした具体的な話に入っていくのだが，これが [Go] に限る話ではないということはお分かりいただけるだろう。

プログラマにとって最も信頼できるドキュメントは動いているプログラムコードである。
だからこそコードは「文芸的[^literate1]」であるべきだし，プログラマは要件定義の段階から積極的にコードを書くべきだと思う。
リファクタリングは何時でもできるのだから[^refact1]。

[^literate1]: 「コードはもっと文芸的であるべき」というのはクヌース博士のいわゆる「文芸的プログラミング（literate programming）」とはちょっと違う。ごめんペコン。
[^refact1]: というか，これからの時代はリファクタリングに厚い言語を選択すべき。

書いた人にしか分からない20世紀なコードはもうケッコウである。

## ブックマーク

- [技術的負債とハッカー]({{< ref "/remark/2020/12/technical-debt-and-hacker.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
<!-- eof -->
