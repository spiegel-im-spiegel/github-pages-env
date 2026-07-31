+++
title = "『ピープルウエア』を借りて流し読む"
date =  "2023-02-12T20:38:11+09:00"
description = "もしこれから買うって方は紙の本をおすすめする。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "communication", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

『[Effective Java](https://www.amazon.co.jp/dp/4621303252?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』や『[プログラミング言語Go](https://www.amazon.co.jp/dp/B099928SJD?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』の翻訳でおなじみ[柴田芳樹](https://yshibata.blog.ss-blog.jp/ "柴田 芳樹 (Yoshiki Shibata)：SSブログ")さん主催の読書会で『[ピープルウエア]』を取り上げることになったそうな。

- [第1回『ピープルウェア第3版』オンライン読書会 - connpass](https://technical-book-reading.connpass.com/event/273601/)

不勉強なことに『[ピープルウエア]』を知らなかったので，まずは[県立図書館][島根県立図書館]で[借りて]({{< ref "/remark/2023/01/castle-in-winter.md" >}} "松江城 冬景色")読んでみることにした。
これが，いざ借りようとしたら[図書館][島根県立図書館]の地下倉庫にあるっていうのよ。
受付で一般の人が借りていいか訊いてしまったよ。
まぁ，でも，無事に借りれた。

もしこれから買うって方は紙の本をおすすめする。
この本は注釈が巻末にまとめて書かれていて（縦書きの人文系の本ではよくあるスタイル） Kindle ではたぶん読みにくい。
PDF なら複数のウィンドウで開いたり，アプリケーションによっては画面分割したりできるんだけどねぇ。
まぁ，私は[日経嫌い](https://baldanders.info/blog/000709/)なので PDF 版出ても買わんだろうけど（笑）

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

『[ピープルウエア]』の問題意識は，プロジェクトで起きるアレコレは大抵，技術的な問題というよりは人間側の問題であると考えている点にある。
この本ではこれを「プロジェクトの社会学」と呼んでいる。
だから “Peopleware” なのだ。

たとえば

{{< fig-quote type="markdown" title="『ピープルウエア』 p.37" link="https://www.amazon.co.jp/dp/4822285243?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
開発者の主な仕事は、ユーザー流の表現で表したユーザー要求を、厳格な処理手順に組み替えるための、人と人とのコミュニケーションである。
{{< /fig-quote >}}

という感じ。
他にも

{{< fig-quote type="markdown" title="『ピープルウエア』 p.4" link="https://www.amazon.co.jp/dp/4822285243?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
本当にハイテクビジネスに身を置いているのは、ハイテクの領域で、基本的な発明発見を成し遂げた研究者だけだ。それ以外の人は他人の研究成果を応用しているに過ぎない。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『ピープルウエア』 p.36" link="https://www.amazon.co.jp/dp/4822285243?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
プログラミング言語は、問題を考える方法に大きな影響を与えるので重要だが、これもコーディング工程にしか影響を与えない。
{{< /fig-quote >}}

など，私のような凡庸なエンジニアには耳の痛い記述もある（笑）

でも『[ピープルウエア]』でメインに dis ってるのは経営者や管理職だろう。

{{< fig-quote type="markdown" title="『ピープルウエア』 p.110" link="https://www.amazon.co.jp/dp/4822285243?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
自信のない権威主義的体制（教区附属学校や軍隊など）のもとでは画一性は極めて重要で、この下ではドレスコードを強いることさえある。スカートの丈やジャケットの色の違いに脅威を感じるから禁止されるのだ。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『ピープルウエア』 p.110" link="https://www.amazon.co.jp/dp/4822285243?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
屍を支える仕事は人に満足を与えない。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『ピープルウエア』 p.143" link="https://www.amazon.co.jp/dp/4822285243?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
人的資本は非常に大きな意味を持ち得る。誤って消えてなくなる経費と考えると、投資した価値が残らないような行動にマネージャーを誘導する危険がある。
{{< /fig-quote >}}

この辺を読んで，某 Twitter のことか？ と一瞬思っちまったよ（笑）

最近 “[Elon Musk fires a top Twitter engineer over his declining view count](https://www.platformer.news/p/elon-musk-fires-a-top-twitter-engineer)” という記事を見かけたが，この中で

{{< fig-quote type="markdown" title="Elon Musk fires a top Twitter engineer over his declining view count" link="https://www.platformer.news/p/elon-musk-fires-a-top-twitter-engineer" lang="en" >}}
The perks that made Twitter an attractive place to work pre-Musk have been eradicated. Food at the office? “Sucks – and now we have to pay for it. And, I know this sounds petty, but they appear to have obtained the absolute worst coffee vendors on earth.”
{{< /fig-quote >}}

とか

{{< fig-quote type="markdown" title="Elon Musk fires a top Twitter engineer over his declining view count" link="https://www.platformer.news/p/elon-musk-fires-a-top-twitter-engineer" lang="en" >}}
“People don’t even chat about work things anymore,” the employee said. “It’s just heartbreaking. I have more conversations with my colleagues on Signal and WhatsApp than I do on Slack. Before the transition, it was not uncommon in the team channel to talk about what everybody did that weekend. There’s none of that anymore.” 
{{< /fig-quote >}}

とか書かれていて，まさに『[ピープルウエア]』に書かれている悪い例が進行しているのを見て「うわぁ」となった。

たぶん『[ユニコーン企業のひみつ](https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』や，最近なら『[ちょうぜつソフトウェア設計入門](https://www.amazon.co.jp/dp/B0BNH1J2W2?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』を読むような人なら，その前説として『[ピープルウエア]』は読んでおいて損はないだろう。
そういえば，オンラインイベント「[ちょうぜつ改め21世紀ふつうのソフトウェア設計入門 - Forkwell Library #14](https://forkwell.connpass.com/event/271212/)」の[めもりー](https://twitter.com/m3m0r7)さんによる[講演](https://speakerdeck.com/memory1994/why-the-application-design-is-breaking-sometimes-at-a-startup-company "CTO から見た，なぜスタートアップ初期のソフトウェア設計は壊れがちなのか - Speaker Deck")の中で，エンジニアの採用事情について言及があったが，『[ピープルウエア]』を読めば，経費ではなく，最も価値ある資源であり投資対象でもある「人材」を如何にして確保し定着させるか，といったことの片鱗が見えるかもしれない。

というわけで，とりあえず第III部の途中まで読んで本は返したが，なかなか面白かったので買うことにした。
[図書館][島根県立図書館]においては，地下倉庫じゃなくて，是非とも一般の書架に置いていただきたいものである（笑） [読書会](https://technical-book-reading.connpass.com/event/273601/ "第1回『ピープルウェア第3版』オンライン読書会 - connpass")も参加してみようかな。
この本を若い人が読んだらどう思うのだろうか。

[ピープルウエア]: https://www.amazon.co.jp/dp/4822285243?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1
[島根県立図書館]: https://www.library.pref.shimane.lg.jp/

## 参考図書

{{% review-paapi "4822285243" %}} <!-- ピープルウエア Peopleware -->
{{% review-paapi "B07FSBHS2V" %}} <!-- Clean Architecture -->
{{% review-paapi "4873119464" %}} <!-- ユニコーン企業のひみつ -->
{{% review-paapi "B0BNH1J2W2" %}} <!-- ちょうぜつエンジニアめもりーちゃん -->
{{% review-paapi "B0BQ6M39HJ" %}} <!-- 小さな会社のスクラム実践講座 -->
