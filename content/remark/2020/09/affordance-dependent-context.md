+++
title = "それはコミュニケーションぢゃない，グルーミングだ！"
date =  "2020-09-03T12:27:26+09:00"
description = "コミュニケーションの巧拙が問題なのではなく，そもそもコミュニケーションを始めるに至ってないということだ。"
image = "/images/attention/kitten.jpg"
tags = [ "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近チラホラと「ハイコンテキスト文化」「ローコンテキスト文化」みたいなフレーズを聞くので「？？？」と思っていたが，いわゆる「リモートワーク」が日本で上手く行かない言い訳として使われているようだ。

- [「在宅勤務は生産性ダウン」と感じる人、日本はトップ　10カ国平均大きく上回る　レノボ調査で明らかに - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/2007/16/news080.html)
- [日本人が「在宅勤務は生産性ダウン」と感じる理由 - hogehoge, world.](https://tomoto335.hatenablog.com/entry/Japanese-work-from-home)

最初に「ハイコンテキスト文化／ローコンテキスト文化」なるフレーズを唱えた人の意図は知らないが（興味ない），そもそも「文脈」のない対話や集団間コミュニケーションは存在しない。
とすれば，そのコミュニケーションに横たわる「文脈」はどこに存在するのか，という話になる。

ところで，私が実家に帰郷って一番イライラするのは家族と「文脈」を共有できないことである。
まぁ，30年以上も離れて暮らしてたのだから当たり前っちゃあ当たり前なのだが。

駄菓子菓子！

{{< div-gen type="markdown" class="center" >}}
*！主語を省くな！目的語を省くな！述語を省くな！*
{{< /div-gen >}}

あー，イライラする（笑）

...気を取り直して。

日本語は構造上「文脈」を言外に置かざるを得ない。
たとえば年寄りの話す出雲弁では一人称も二人称も「わー」である。
「わー」が誰を指すのかは言外の文脈に依存する。
日本語におけるこの手の話は，それこそゴマンと存在している。

それじゃあ，継続的なコミュニケーションにおいて言外の「文脈」はどこに置かれるかというと，空間的・心理的な「場」である。
そして「文脈」が共有される「場」は「意味」を得てアフォーダンス（affordance）となる。

だから「ハイコンテキスト」というフレーズは恐らく適当ではなくて，むしろ「アフォーダンス依存コンテキスト」とでも言うべきだろう。
アフォーダンスは境界が明確であるほど，そして内部に横たわる文脈が曖昧であるほど同調圧力も依存性も強くなる。
秘密結社的（笑）

しかし，インターネットのようなフラットな空間ではアフォーダンスを形成しにくい。
このような状況で人が執る行動は概ね以下の2つだ。

1. 「文脈」をコード化・明文化して共有する
2. 擬似的なアフォーダンスを形成して内外を区別（包摂と排除）する

まぁ，大抵の集団は後者を選択するのだろう。
そもそも SNS 等のサービスはそのために存在するんだし。
その様子は例えば『[つながりっぱなしの日常を生きる]』あたりを読むと分かりやすいかもしれない。

で，日本のビジネスでは，アフォーダンスを形成するためにまず「電子メールに時候の挨拶を書く」「Zoom 画面の上座を決める」等々のグルーミングから始めるらしい[^g1]。

[^g1]: 日本ではこうしたグルーミングを「マナー」と呼んでる（笑）

そんな{{< ruby "グルーミング" >}}茶番{{< /ruby >}}を見せられれば「在宅勤務は生産性ダウン」と感じてもしょうがないよね。
オフラインでは「何となく」行われているグルーミングがオンラインで「見える化」しちゃうわけだから。
つまり，コミュニケーションの巧拙が問題なのではなく，そもそもコミュニケーションを始めるに至ってないということだ。

一方，インターネットはそもそも前者の「『文脈』をコード化・明文化して共有する」ことを前提に設計されている。

そういえばシステム・プログラミング設計の「オブジェクト指向」の派生に「コンテキスト指向」というのがある。
簡単に言うと，オブジェクト間・ドメイン間・レイヤ間を跨ぐコンテキストを明文化・コード化する設計指針だ。
コンテキストを明示することでオブジェクト・ドメイン・レイヤといった境界を定義し，それぞれの独立性を担保できる。

逆に言うとコンテキストが曖昧な設計は「境界」も「関係」も曖昧になり，最悪「動かないシステム」になりかねない。
まぁ，設計の初歩だよね（笑）

てな風に考えると「アフォーダンス依存コンテキスト」に慣れすぎている日本人ってめっさ不利なんじゃないだろうか。
前々から「[プログラミング教育よりもっとすることがあるだろう]({{< ref "/remark/2016/09/programming.md" >}} "プログラミングで「計算論的思考」は身につかない")」と思っていたが，今回のような観点からも考えてみるべきかもしれない。




[つながりっぱなしの日常を生きる]: https://www.amazon.co.jp/exec/obidos/ASIN/B0125TZSZ0/baldandersinf-22/ "Amazon.co.jp: つながりっぱなしの日常を生きる：ソーシャルメディアが若者にもたらしたもの 電子書籍: ダナ・ボイド, 野中 モモ: Kindleストア"

## 参考文献

{{% review-paapi "B0125TZSZ0" %}} <!-- つながりっぱなしの日常を生きる -->
{{% review-paapi "B00FW5SSCK" %}} <!-- ソーシャル・ネットワーク -->
{{% review-paapi "B073PT6WDB" %}} <!-- 秘密結社の世界史 -->