+++
title = "FOSS とジョナサン"
date =  "2020-09-14T21:06:21+09:00"
description = "そして Open Source は「強者の武器」となる。"
image = "/images/attention/kitten.jpg"
tags = [ "code", "grigori", "politics", "internet", "surveillance-capitalism" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近 Facebook や Twitter の TL で FOSS (Free/Open Source Software) に関する総括みたいな記事をよく見かけるのだが，気のせいだろうか。
まぁ TL って自身の性癖の暴露だから，そういう風に見えてしまうこともあるのだろう。

## 「自由」は主張である

人とは本来，不自由な存在である。
何故なら，人は独りでは自身を証明できないから。
故に人が「自由」を求めるのなら「対象」が必ずあるし，求めるからには「対価」が必要である。

Free Software は以下の4つの自由を求めている。

{{< fig-quote type="markdown" title="自由ソフトウェアとは? - GNUプロジェクト - フリーソフトウェアファウンデーション" link="https://www.gnu.org/philosophy/free-sw.html" >}}
- どんな目的に対しても、プログラムを望むままに実行する自由 (第零の自由)。
- プログラムがどのように動作しているか研究し、必要に応じて改造する自由 (第一の自由)。
- ほかの人を助けられるよう、コピーを再配布する自由 (第二の自由)。
- 改変した版を他に配布する自由 (第三の自由)。
{{< /fig-quote >}}

物凄く簡単に言うと Free Software は「『政治から自由である』という政治主張」を持つソフトウェア（活動）なのである。
そして「政治主張」を実践するからには，責任やら義務やらが発生する。
その実装例が GNU GPL (General Public License; 一般公衆ライセンス) における “copyleft” という法的 hack である，と考えれば分かりやすいだろう。

つまり Free Software は「政治から自由である」ために政治的な責任・義務を負うという Anarchism 特有のレトリックを抱えているわけだ。

## Open Source as a Jonathan

インターネットの老害達（笑）が懐かしむ「インターネットの黄金時代」をヒッピー文化に擬えることがある。
で，私の中でヒッピー文化の典型は『[かもめのジョナサン](https://www.amazon.co.jp/dp/B01916B8V8?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』だったりする。

かの作品に対する感想や評価は色々あるだろうが，私の中でジョナサンは「『飛ぶ』ために全てから逃げた存在」だ。
まぁ，かつてのオウム信者にはアレを読んで入信した人とかもいるそうなので，それほど的はずれではないだろう（笑）

（最初からそう呼ばれていたわけではないが） Open Source は Free Software 運動から「政治主張」を洗い落としたものである。
Open Source の枠組みでは “copyleft” もオプションのひとつに過ぎない。
おおっ！ これぞまさしくポリシー・ロンダリング（違う[^pl1]）。

[^pl1]: 本当の policy laundering は国際的な議論や規制を都合よく利用して国内での立法プロセスを迂回することをさす： {{% quote lang="en" %}}In their review of global business regulation, Braithwaite and Drahos find that some countries (notably the U.S. and the UK) push for certain regulatory standards in international bodies and then bring those regulations home under the requirement of harmonization and the guise of multilateralism; this is what we refer to as policy laundering{{% /quote %}}. (via [Policy laundering - Wikipedia](https://en.wikipedia.org/wiki/Policy_laundering))

Open Source がそうした理由は単純だ。
「自由」の対価を払いたくなかったから。
だから「自由を得る」のではなく「不自由から逃げる」のだ，ジョナサンのように，「それがぼくには楽しかったから」。

## 囚人のジレンマと搾取

すずきひろのぶさんが最近「{{< pdf-file title="フリーソフトウェアと それを取巻く世界の状況" link="file:///tmp/mozilla_spiegel0/FreeSoftware20200905.pdf" >}}」という2007年当時のスライドを公開されていて，この中に「囚人のジレンマ」についての言及がある。

一応「囚人のジレンマ」について簡単に解説しておこう。

ある犯罪を行った囚人 A, B に対し

{{< fig-quote type="markdown" title="Prisoner's dilemma - Wikipedia" link="https://en.wikipedia.org/wiki/Prisoner%27s_dilemma" lang="en" >}}
- If A and B each betray the other, each of them serves two years in prison
- If A betrays B but B remains silent, A will be set free and B will serve three years in prison
- If A remains silent but B betrays A, A will serve three years in prison and B will be set free
- If A and B both remain silent, both of them will serve only one year in prison (on the lesser charge)
{{< /fig-quote >}}

と持ちかける。
表にするとこんな感じ。

{{< fig-quote class="nobox" title="Prisoner's dilemma - Wikipedia" link="https://en.wikipedia.org/wiki/Prisoner%27s_dilemma" lang="en" >}}
<table><tbody>
<tr>
  <th style="vertical-align:middle"></th>
  <th style="vertical-align:middle">Prisoner B stays silent<br>(cooperates)</th>
  <th style="vertical-align:middle">Prisoner B betrays<br>(defects)</th>
</tr><tr>
  <th style="vertical-align:middle">Prisoner A stays silent<br>(cooperates)</th>
  <td style="vertical-align:middle">Each serves 1 year</td>
  <td style="vertical-align:middle">Prisoner A: 3 years<br>Prisoner B: goes free</td>
</tr><tr>
  <th style="vertical-align:middle">Prisoner A betrays<br>(defects)</th>
  <td style="vertical-align:middle">Prisoner A: goes free<br>Prisoner B: 3 years</td>
  <td style="vertical-align:middle">Each serves 2 years<br>(lesser charge)</td>
</tr>
</tbody></table>
{{< /fig-quote >}}

このとき囚人 A, B はそれぞれ黙秘するか裏切るか？ という「非協力ゲーム」の一種である。
この条件下では「両者とも黙秘（協力）する」がパレート効率的であるにも関わらず，「単独のゲーム」または「有限繰り返しゲーム」においては「両者とも裏切る」がナッシュ均衡となることが分かっている。

では「無期限繰り返しゲーム」による囚人のジレンマはどうなるのか。
これについては様々なモデルで研究が行われている。
たとえば

{{< fig-quote title="MIT Tech Review: なぜ搾取は起きるのか？東大研究者が「囚人のジレンマ」で解明" link="https://www.technologyreview.jp/s/144680/prisoners-dilemma-shows-how-exploitation-is-a-basic-property-of-human-society/" >}}
{{% quote %}}明白な答えの1つは、力のある人はその力を使って力の弱い人を利用できるということだ{{% /quote %}}
{{< /fig-quote >}}

は典型的な「搾取」のパターンだ。

さらに最近では「[戦略的力関係とゲームのルールが両者に対称的である場合](https://www.technologyreview.jp/s/144680/prisoners-dilemma-shows-how-exploitation-is-a-basic-property-of-human-society/ "MIT Tech Review: なぜ搾取は起きるのか？東大研究者が「囚人のジレンマ」で解明")」でも

{{< fig-quote type="markdown" title="MIT Tech Review: なぜ搾取は起きるのか？東大研究者が「囚人のジレンマ」で解明" link="https://www.technologyreview.jp/s/144680/prisoners-dilemma-shows-how-exploitation-is-a-basic-property-of-human-society/" >}}
アリスがボブの戦略を知った場合に、ボブの行動を利用し、アリス自身の結果をより良くできることを、藤本と金子教授は示している。

しかし、この戦略がボブにとってより良い結果を約束するとアリスが保証すれば、アリスはボブの協力を確実なものにできる。たとえばある状況では、両方のプレイヤーが裏切る場合の結果よりも、ボブが犠牲になった方がボブにとって良い結果になるとアリスが保証できるのだ。

このため、たとえアリスがより有利になったとしても、ボブには搾取を受け入れる動機がある。「このように、両方のプレイヤーが搾取関係を安定にしています」。
{{< /fig-quote >}}

といった「搾取」のパターンも発表されている。

## そして Open Source は「強者の武器」となる

すずきひろのぶさんの「{{< pdf-file title="フリーソフトウェアと それを取巻く世界の状況" link="file:///tmp/mozilla_spiegel0/FreeSoftware20200905.pdf" >}}」では， GPL 下でソフトウェア開発を行うことでパレート効率的である「両者とも協力する」関係に固定できると主張している（もちろん喩え話ですよ）。

しかし GPLv3 アップデートを巡る騒動を見る限り，誰もがパレート効率的な「両者とも協力する」関係を望んでいるわけではないように見える。

{{< fig-quote type="markdown" title="Post-Open Source | boringcactus" link="https://www.boringcactus.com/2020/08/13/post-open-source.html" lang="en" >}}
{{% quote %}}the GPLv2 was pretty popular at the time, but there were a couple notable loopholes some big corporations had been taking advantage of, which the free software people wanted to close. a whole bunch of people thought the GPLv2 was fine the way it was, though - closing the loopholes as aggressively as the GPLv3 did cut off some justifiable security measures, and some people said that it could do more harm than good. the linux kernel, along with a lot more stuff, declared it was sticking with the GPLv2 and not moving to the GPLv3{{% /quote %}}.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Post-Open Source | boringcactus" link="https://www.boringcactus.com/2020/08/13/post-open-source.html" lang="en" >}}
{{% quote %}}and LLVM became at least as good as GCC, and a less risky decision for big companies, and easier to use to build new languages. so the free software movement’s last technical advantage was gone. its social advantages also kinda went up in flames with the GPLv3, too: the software that was the foundation for the GPL enforcement lawsuits stuck with the GPLv2{{% /quote %}}.
{{< /fig-quote >}}

「自由を得る」 Free Software (運動) より「不自由から逃げる」 Open Source を選好することによって何が起きるか。
「[戦略的力関係とゲームのルールが両者に対称的である場合でも、搾取的な関係が構築できる](https://www.technologyreview.jp/s/144680/prisoners-dilemma-shows-how-exploitation-is-a-basic-property-of-human-society/ "MIT Tech Review: なぜ搾取は起きるのか？東大研究者が「囚人のジレンマ」で解明")」のであれば，それはどのような形を取るのか。

ジョナサンは「ここではない」世界へと逃げ込めたが，かもめならぬ我が身では何があろうとこの世界この社会で生きていかなければならない。

Open Source に関する批判は色々見受けるが，要するに Open Source が「強者の武器」となり，情報力を背景にした搾取の構造を構成していることが問題なのだと思う。
しかし Free Software ではその非対称性を壊せない。

故に

{{< fig-quote type="markdown" title="Post-Open Source | boringcactus" link="https://www.boringcactus.com/2020/08/13/post-open-source.html" lang="en" >}}
{{% quote %}}tl;dr: say no to licenses, say yes to norms{{% /quote %}}.
{{< /fig-quote >}}

なんてな話も出るのかもしれない。

## ブックマーク

- [クーリエ連載；エコノミスト紹介、自由のためなら人が死んでもいい](https://cruel.org/economist/courier200712.html)
- [フリーソフトウェアとそれを取巻く世界の状況 (2007) – とりあえずノートがわりに書いてます](https://mynotebook.h2np.net/post/5666)
- [Post-Open Source | boringcactus](https://www.boringcactus.com/2020/08/13/post-open-source.html)
    - [ブログ: ポスト・オープンソース](https://okuranagaimo.blogspot.com/2020/09/blog-post_13.html)
- [孵卵器の中のインターネット]({{< ref "/remark/2017/01/internet-in-the-incubator.md" >}})
- [搾取と狂狷]({{< ref "/remark/2019/06/kyoken.md" >}})

## 参考図書

{{% review-paapi "B01916B8V8" %}} <!-- かもめのジョナサン -->
{{% review-aozora "4307" %}} <!-- グリゴリの捕縛 -->
{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
{{% review-paapi "4796880011" %}} <!-- それがぼくには楽しかったから -->
{{% review-paapi "B00FU3P37W" %}} <!-- かもめが翔んだ日 -->
