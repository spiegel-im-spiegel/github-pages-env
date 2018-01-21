+++
title = "『シンギュラリティの神話』を読む"
date =  "2017-09-18T21:45:03+09:00"
update = "2018-01-21T11:19:13+09:00"
description = "もし『そろそろ、人工知能の真実を話そう』が日本の読者になにがしかの衝撃を与えるとするなら，自らの稚拙さの方だろう。"
tags        = [ "book", "artificial-intelligence", "singularity" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"
+++

（この記事は「[AI と哲学？]({{< relref "remark/2017/09/artificial-intelligence.md" >}})」の続編である）

『[そろそろ、人工知能の真実を話そう]』読んだ。
めっさ疲れた。
疲労感つか徒労感？ が半端ない。

原書のタイトルは “Le mythe de la Singularité" であり， Google 先生の力を借りて訳すなら「シンギュラリティの神話」といったところだろうか。
なんで翻訳本ってタイトルのセンスが壊滅してるんだろうねぇ。
まぁ副題が “Faut-il craindre l'intelligence artificielle ?" なので人工知能の話がメインになると言えなくもないけど，実際には人工知能に限る話ではないのだ。 

『[そろそろ、人工知能の真実を話そう]』は8章からなる本文と東京大学名誉教授の西垣通さんによる解説で構成されているが，作者の主張は一貫していて

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/" >}}
<q>シンギュラリティの提唱者が主張するような断絶が起こることを証明するものは何もない。
おそらく、これからも進歩は加速し続け、その大きな渦でわれわれを飲みこもうとするだろう。
しかし、その現実に目をふさいでしまうのではなく、行動することこそが求められている。</q>
{{< /fig-quote >}}

というものである。

しかし，この主張の前準備のために延々7章を費やしているのである。
ちなみに7章まで読んだ私の感想は「呆れてものが言えない」だった。
あぁ，誤解のないように言っておくと作者のロジックや文章が酷いという意味じゃないからね（翻訳の妥当性は私には分からないけど）。

本当に面白いのは8章と解説なのだが，8章に出てくる用語をきちんと理解するには7章までを読み込んでおかなければならない。
実に面倒くさい構成である。
これだから哲学者ってやつは...

というわけで，以降では『[そろそろ、人工知能の真実を話そう]』に登場する用語についていくつか紹介してみたいと思う。
これらを踏まえて読めば多少は理解が進むかもしれない。

## 用語について

### シンギュラリティ（Singularity）

タイトルにも出てくる「シンギュラリティ」という言葉だが，これはもともと数学用語らしい。
そののち科学の分野でも使われるようになったようだ。
例えばビッグバンやブラックホールは古典物理学が破綻する「特異点（singularity）」である。

この「シンギュラリティ」の概念を人文科学（言語学や認識論など）の分野に取り入れようという動きが1970年代にあり，それが今回のメインテーマである「シンギュラリティ仮説」の下地になっているように見える。

ちなみに「シンギュラリティ仮説」とは以下のような内容である（「解説」より）。

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/" >}}
<q>シンギュラリティ仮説とは、二〇四五年あたりに AI の能力が人間を凌ぎ、機械支配が進んで世界のありさまが大きく変容してしまうという予測だ。</q>
{{< /fig-quote >}}

AI の能力が人間を凌ぐ2045年が「特異点」というわけだ。

『[そろそろ、人工知能の真実を話そう]』を読んでて「この流れって既視感があるなぁ」と思ったが，よく考えたらこれって「[ソーカル事件]」ぢゃん。

### ムーアの法則

いや，もう，いまさら[ムーアの法則]に言及するのは止めてほしいのだが，この[ムーアの法則]が「シンギュラリティ仮説」の屋台骨のひとつになっているというのだから笑ってしまう。

ちなみに元々の[ムーアの法則]は以下の文章から来ている。

{{< fig-quote  title="Cramming more components onto integrated circuits （和訳は Wikipedia より）" link="https://ja.wikipedia.org/wiki/%E3%83%A0%E3%83%BC%E3%82%A2%E3%81%AE%E6%B3%95%E5%89%87" >}}
<q>部品あたりのコストが最小になるような複雑さは、毎年およそ2倍の割合で増大してきた。短期的には、この増加率が上昇しないまでも、現状を維持することは確実である。
より長期的には、増加率はやや不確実であるとはいえ、少なくとも今後10年間ほぼ一定の率を保てないと信ずべき理由は無い。
すなわち、1975年までには、最小コストで得られる集積回路の部品数は65,000に達するであろう。
私は、それほどにも大規模な回路が1個のウェハー上に構築できるようになると信じている。</q>
{{< /fig-quote >}}

[ムーアの法則]はゴードン・ムーア氏が1965年の論文で述べたもので，1975年までの短期的予測を述べたものに過ぎない。
なのに，たまたまその後も半導体部品の集積度が[ムーアの法則]に追従するように進歩していったため，ほとんど御神託のようになってしまった。

しかし，元々の[ムーアの法則]は今世紀早々に破綻しているのだ。
その後は微妙に解釈を変えて如何にも[ムーアの法則]に追従しているように見せかけているが，はっきり言って「粉飾決算」である。

### 自立と自律

いやぁ。
私は「自立」と「自律」について曖昧にしていたよ。
これからは気をつけよう。

両者の違いは以下のようなものらしい。

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/" >}}
<q>自立とは、仮想代理人ソフトウェアであるところのエージェントが自ら動き、誰の力も借りずに意思決定できることを言う。
[...]
一方、自律というのは哲学的な意味であり、自らが行動する際の基準と目的を明確を持ち、自ら規範を作り出すことができることをいう。</q>
{{< /fig-quote >}}

この定義によれば，現在この世界に「自律機械」は存在しないし，現在の延長線上の未来においても「自律機械」は登場しそうもない。

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/" >}}
<q>今、世の中で懸念されているのは、自立ではなく自律の方だが、学習能力を与えられ、自らのプログラムを改善できるようになっても、機械が自律することは考えられない。
なぜなら、機械は結局、人間に与えられた理論やルールにのっとって行動することになるからである。</q>
{{< /fig-quote >}}

そのとーり！ （by 財津一郎）

### 「強い人工知能」と「弱い人工知能」

「人工知能（artificial intelligence）」という言葉が登場したのは1955年のことらしい。

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/" >}}
<q>若干はSFの世界に影響されていたかもしれないが、マッカーシーらの狙いは謙虚なものだった。
創造主に成りかわるつもりはなく、人間のコピーや超人を創り出そうというのでもない。
彼等の目的は、あくまでも実証的で現実的なものだった。
動物の知能であれ人間の知能であれ、認知機能を機械で模倣することで知能をもっとよく理解したいと考えたのだ。</q>
{{< /fig-quote >}}

しかし，ここでも哲学者共が横槍を入れる。

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/" >}}
<q>哲学者が問題にしたのは、実際に研究者やエンジニアが取り組んでいる研究活動というより、その向こうにある哲学的な思想だった。
サールが強い人工知能という言葉を説明の道具として使ったのは、一九八〇年代の初頭だった。
この強い人工知能こそが、先程定義した人工知能の仮像である。
なぜなら、強い人工知能とそもそもの意味での人工知能は、名前の上では似通っていても、目標も方法もまったく異なるものだからだ。
かつて、コンピュータでシミュレーションを行ない、実験によって検証することに基礎を置いた科学の一分野であったものが、今では論証だけに基づいた哲学的アプローチとなっている。</q>
{{< /fig-quote >}}

こうした「脱線」が「シンギュラリティ仮説」の下地になっているのだから，笑うしかない。

ちなみにサールが唱えた「強い人工知能」は「[中国語の部屋]」に登場する。

#### 中国語の部屋

「[中国語の部屋]」は人工知能を模した思考実験の論文だが，ほぼ同じ機能を持つプログラムは既に存在する。
「人工無脳」（現在はチャットボット（chatbot）などと呼ばれている）である。

人工無脳は相手と擬似的なコミュニケーションを行うが，人工無脳はそのやり取りの意味を知っているわけではない。
入力に対して何らかのアルゴリズムを介し，バックエンドにある語彙を組み合わせて応答を返しているだけである。
これを以って人工無脳がそのやり取りを「理解」していると言えるか，というのが「[中国語の部屋]」の命題である。

### 可能性、蓋然性、信憑性

これも『[そろそろ、人工知能の真実を話そう]』では重要な言葉である。

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/" >}}
<q>近頃は、われわれを翻弄するような科学的な言い回しが非常に多く、未来に対して異なるアプローチをとる三つの概念がよく混同されている。すなわち、可能性（possibilité）、蓋然性（probabilité）、信憑性（plausibilité）―― この三つが区別されていないために混乱が生じているのである。</q>
{{< /fig-quote >}}

信憑性についてはもう少し説明が必要だろう。

plausibilité (plausibility) の意味を Google 先生に訊くと「{{< ruby "ゆうど" >}}尤度{{< /ruby >}}」と答えが返ってきた。
尤度とは統計学の用語らしい。

- [【統計学】尤度って何？をグラフィカルに説明してみる。 - Qiita](http://qiita.com/kenmatsu4/items/b28d1b3b3d291d0cc698)

ただし，この本では「信憑性（plausibilité）」をそのような意味では使ってなくて，語源に近いニュアンスで

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/" >}}
<q>つまりそれは一般受けが良く、多くの人が起こると思っているということだ。だが、実際には可能性も蓋然性も保証されていない。</q>
{{< /fig-quote >}}

という感じで使っている。
要するに「シンギュラリティ仮説」には可能性も蓋然性もなく，ただ根拠のない信憑性のみで強引に推し進められているというわけだ。

### 仮像（Pseudomorphose）

これまでの記述でも度々出てきたが，「仮像」というのは『[そろそろ、人工知能の真実を話そう]』では最重要の言葉である。

もともと「仮像」は科学用語で，たとえば鉱物が外形を保ったままで別の鉱物に置換されることを指したりするらしい。
そういや化石も仮像の一種だよね。

この本ではこの「仮像」をもっと広い範囲に拡張する。
たとえば，「強い人工知能」は「弱い人工知能（＝そもそもの意味での人工知能）」の仮像である，といった具合にだ。
そして「シンギュラリティ仮説」を巡る言説を，仮像をキーワードにしてグノーシス主義と比較し，「シンギュラリティ仮説」の問題点を炙りだしている。

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/" >}}
<q>結局、このシンギュラリティ自身も啓蒙主義におけるヒューマニズムの仮像なのだ。
[...]
啓蒙主義には、ヒューマニズムの名のもとに進歩を限りなく続けていこうという理想がある。
そして、そのためには外の世界へ自らを無限に開放していかなければならない。
だがシンギュラリティは、完璧に作り上げられた結末の中に、未来を閉じ込めてしまうのである。</q>
{{< /fig-quote >}}


正直に言って，この手の意味の拡張はものすごく危険な行為だと思うのだけど，私が気にしてもしょうがないし，まぁいいか。

## 本当に面白くなるのは第八章から（笑）

そろそろまとめに入ろうか。

『[そろそろ、人工知能の真実を話そう]』の論点は2つあると思う。
ひとつは「シンギュラリティ仮説」に対して無邪気に礼賛なり嫌悪なりする人たちで，もうひとつは「シンギュラリティ仮説」に乗っかる「市場」である。

前者については多くの（キリスト教圏外の）日本人は「ネタにマジレス（笑）」としか返せないだろう。
解説で西垣通さんは

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/" >}}
<q>AI と言えば、技術的改良と経済効果の話題だけ、あとはせいぜい幼稚な夢物語というのが、情けないことにこの国の常識だ。
本書がそんな常識に衝撃を与えることを切に願う。</q>
{{< /fig-quote >}}

と書かれているが，「シンギュラリティ仮説」を真面目に叫ぶ輩などカルト教祖[^aum] か詐欺師か，さもなくば厨二病患者に決まってる。
いまどき子供向けのラノベでも「ない」わ。

[^aum]: 私は「シンギュラリティ仮説」とグノーシス主義との比較部分を読んで，かつてのオウム真理教を連想してしまった。オウム真理教は医者や研究者などのいわゆる「理系」の信者が多かったことも特徴で（それ故にサリンを使ったテロにまで走ってしまうのだが），そういう意味でも類似性はあるのかな，などと思ったり。

後者に関してこの本では Google を例に挙げている。
確かに Google は「放火魔の消防士[^mp]」だが，それは Google に限る話ではなく，そこらじゅうに転がっている。

[^mp]: 日本語で言えば「マッチポンプ」。今まで知らなかったのだが「マッチポンプ」って和製外来語なんだねぇ。

Google が誠実なのはユーザや消費者に対してではなく「市場」に対してである。
Google が言う “don't be evil" も市場に対するメッセージであると解釈すれば納得だろう。

例えば Google は「スノーデン」以前において[「完全なプライバシーは存在しない」と公言](http://www.thesmokinggun.com/documents/internet/google-complete-privacy-does-not-exist)し「[Google はプライバシーに敵対的](http://cloud.watch.impress.co.jp/epw/cda/infostand/2007/06/18/10531.html)」と大きな批判を浴びていた。
それが「スノーデン」以後は手のひらを返すようにプライバシー擁護を謳っている。

この事実だけでも Google という企業の本質が垣間見えるだろう。
そしてそれは Google に限った話ではなく Amazon だって Facebook だって Microsoft だって Tesla だって同じ穴の狢なのである。

そして「賢い企業」はどちらか一方に bet することはない。
掛けで必ず勝つには胴元になることである。

自社の計算資源を「クラウド」として切り売りする企業にとって「「シンギュラリティ仮説」に対して無邪気に礼賛なり嫌悪なりする人たち」は大事なお客様である。
とくに Google は AI の計算資源すら対象にしているのである。
ならば，「放火魔の消防士」と呼ばれようが，「シンギュラリティ仮説」に群がる全ての人達を顧客として迎い入れるのが「正解」である。

この点において日本の「IT 業界」は， Google などと比較すれば，全くもって稚拙であると言わざるをえない。
もし『[そろそろ、人工知能の真実を話そう]』が日本の読者になにがしかの衝撃を与えるとするなら，自らの稚拙さの方だろう。

**結論：** シンギュラリティを口にする連中と哲学者は信用してはいけない

## ブックマーク

- [『そろそろ、人工知能の真実を話そう』シンギュラリティ仮説の背後にうごめくもの - HONZ](http://honz.jp/articles/-/44063)
- [AI時代に「哲学」は何を果たせるか？ 『そろそろ、人工知能の真実を話そう』著者に訊く｜WIRED.jp](https://wired.jp/2017/07/04/jean-gabriel-ganascia/)
- [GoogleのAIのトップは曰く、人工知能という言葉自体が間違っている、誇大宣伝を生む温床だ | TechCrunch Japan](http://jp.techcrunch.com/2017/09/20/20170919googles-ai-chief-thinks-reports-of-the-ai-apocalypse-are-greatly-exaggerated/)
- [「マスターアルゴリズム」著者が預言する人工知能の未来 | ROBOTEER](https://roboteer-tokyo.com/archives/6203)
    - [10年後の未来、面接やデートはオンライン分身がこなす！？ | ROBOTEER](https://roboteer-tokyo.com/archives/2155)
- [「人工知能で神を」 元Googleエンジニアが宗教団体を創立](http://www.huffingtonpost.jp/amp/2017/10/12/anthony-levandowski_a_23241928/) : ネタがマジに！ つか，ホンマにネタじゃないのか？
- [MIT Tech Review: シンギュラリティは来ない、AIの未来予想でよくある7つの勘違い](https://www.technologyreview.jp/s/58257/the-seven-deadly-sins-of-ai-predictions/)
- ~~[人型ロボットに市民権を与えた最初の国家が登場 - GIGAZINE](http://gigazine.net/news/20171027-citizenship-humanoid-robot/)~~ : なんか fakenews ? という話もあるようだ。判断できないので，いったん保留
- [役に立つAIシステムを作ることは、まだまだ難しい  |  TechCrunch Japan](http://jp.techcrunch.com/2018/01/05/2018-01-01-building-ai-systems-that-work-is-still-hard/)
- [シリコンバレーが警告するAIの恐怖、その本質を「メッセージ」原作者が分析](https://www.buzzfeed.com/jp/tedchiang/the-real-danger-to-civilization-isnt-ai-its-runaway-1)

[そろそろ、人工知能の真実を話そう]: http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/ "そろそろ、人工知能の真実を話そう (早川書房) | ジャン＝ガブリエル ガナシア, 小林 重裕・他, 伊藤 直子 | コンピュータサイエンス | Kindleストア | Amazon"
[ソーカル事件]: https://ja.wikipedia.org/wiki/%E3%82%BD%E3%83%BC%E3%82%AB%E3%83%AB%E4%BA%8B%E4%BB%B6 "ソーカル事件 - Wikipedia"
[ムーアの法則]: https://ja.wikipedia.org/wiki/%E3%83%A0%E3%83%BC%E3%82%A2%E3%81%AE%E6%B3%95%E5%89%87 "ムーアの法則 - Wikipedia"
[中国語の部屋]: https://ja.wikipedia.org/wiki/%E4%B8%AD%E5%9B%BD%E8%AA%9E%E3%81%AE%E9%83%A8%E5%B1%8B "中国語の部屋 - Wikipedia"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51cD7DR87IL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B071FHBGW8/baldandersinf-22/">そろそろ、人工知能の真実を話そう (早川書房)</a></dt><dd>ジャン＝ガブリエル ガナシア 小林 重裕・他 </dd><dd>早川書房 2017-05-25</dd><dd>評価<abbr class="rating" title="3"><img src="http://g-images.amazon.com/images/G/01/detail/stars-3-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B075842TRD/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B075842TRD.09._SCTHUMBZZZ_.jpg"  alt="なぜ人工知能は人と会話ができるのか (マイナビ新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B073W485F6/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B073W485F6.09._SCTHUMBZZZ_.jpg"  alt="AIが神になる日"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B072LD494B/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B072LD494B.09._SCTHUMBZZZ_.jpg"  alt="人工知能はどのようにして　「名人」を超えたのか？―――最強の将棋ＡＩポナンザの開発者が教える機械学習・深層学習・強化学習の本質"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B074PQ1KYJ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B074PQ1KYJ.09._SCTHUMBZZZ_.jpg"  alt="ＡＩが人間を殺す日　車、医療、兵器に組み込まれる人工知能 (集英社新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B073S45MCC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B073S45MCC.09._SCTHUMBZZZ_.jpg"  alt="９プリンシプルズ　加速する未来で勝ち残るために (早川書房)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B072FMZL49/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B072FMZL49.09._SCTHUMBZZZ_.jpg"  alt="隷属なき道 AIとの競争に勝つ ベーシックインカムと一日三時間労働 (文春e-book)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0744G41NQ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0744G41NQ.09._SCTHUMBZZZ_.jpg"  alt="アマゾノミクス　データ・サイエンティストはこう考える (文春e-book)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B075D4YLGG/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B075D4YLGG.09._SCTHUMBZZZ_.jpg"  alt="インダストリーX.0"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B071P89MYL/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B071P89MYL.09._SCTHUMBZZZ_.jpg"  alt="ペンタゴンの頭脳　世界を動かす軍事科学機関DARPA"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B074H9SHPM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B074H9SHPM.09._SCTHUMBZZZ_.jpg"  alt="理系脳で考える　AI時代に生き残る人の条件 (朝日新書)"  /></a> </p>
<p class="description">シンギュラリティは起きない。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-18">2017-09-18</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4309410391/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51sgo2CPdpL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4309410391/baldandersinf-22/">社会は情報化の夢を見る---［新世紀版］ノイマンの夢・近代の欲望 (河出文庫)</a></dt><dd>佐藤 俊樹 </dd><dd>河出書房新社 2010-09-03</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4004304652/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4004304652.09._SCTHUMBZZZ_.jpg"  alt="現代社会の理論―情報化・消費化社会の現在と未来 (岩波新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4641150370/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4641150370.09._SCTHUMBZZZ_.jpg"  alt="質的社会調査の方法 -- 他者の合理性の理解社会学 (有斐閣ストゥディア)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4062203855/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4062203855.09._SCTHUMBZZZ_.jpg"  alt="人間と機械のあいだ 心はどこにあるのか"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4480097813/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4480097813.09._SCTHUMBZZZ_.jpg"  alt="社会学的想像力 (ちくま学芸文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4121015371/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4121015371.09._SCTHUMBZZZ_.jpg"  alt="不平等社会日本―さよなら総中流 (中公新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4140910844/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4140910844.09._SCTHUMBZZZ_.jpg"  alt="ウェブ社会の思想 〈遍在する私〉をどう生きるか (NHKブックス)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4140911360/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4140911360.09._SCTHUMBZZZ_.jpg"  alt="社会学入門 〈多元化する時代〉をどう捉えるか (NHKブックス)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4326550775/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4326550775.09._SCTHUMBZZZ_.jpg"  alt="選択しないという選択: ビッグデータで変わる「自由」のかたち"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4121022033/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4121022033.09._SCTHUMBZZZ_.jpg"  alt="集合知とは何か - ネット時代の「知」のゆくえ (中公新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4480069232/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4480069232.09._SCTHUMBZZZ_.jpg"  alt="モテる構造: 男と女の社会学 ((ちくま新書 1216))"  /></a> </p>
<p class="description">1996年に出版された『ノイマンの夢・近代の欲望―情報化社会を解体する』の改訂新装版。しかし内容はこれまでと変わりなく，繰り返し語られる技術決定論を前提とする安易な未来予測を「情報化」社会論だとして批判する。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-19">2017-09-19</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4488711022/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51J3DEJJ1TL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4488711022/baldandersinf-22/">万物理論 (創元SF文庫)</a></dt><dd>グレッグ・イーガン 山岸 真 </dd><dd>東京創元社 2004-10-28</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4488711014/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4488711014.09._SCTHUMBZZZ_.jpg"  alt="宇宙消失 (創元SF文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4150115311/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4150115311.09._SCTHUMBZZZ_.jpg"  alt="ディアスポラ (ハヤカワ文庫 SF)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4150113378/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4150113378.09._SCTHUMBZZZ_.jpg"  alt="祈りの海 (ハヤカワ文庫SF)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4150112908/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4150112908.09._SCTHUMBZZZ_.jpg"  alt="順列都市〈下〉 (ハヤカワ文庫SF)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4150118264/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4150118264.09._SCTHUMBZZZ_.jpg"  alt="プランク・ダイヴ (ハヤカワ文庫SF)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4150112894/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4150112894.09._SCTHUMBZZZ_.jpg"  alt="順列都市〈上〉 (ハヤカワ文庫SF)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/415011594X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/415011594X.09._SCTHUMBZZZ_.jpg"  alt="ひとりっ子 (ハヤカワ文庫SF)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4122056810/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4122056810.09._SCTHUMBZZZ_.jpg"  alt="六本指のゴルトベルク (中公文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/415011451X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/415011451X.09._SCTHUMBZZZ_.jpg"  alt="しあわせの理由 (ハヤカワ文庫SF)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4582760465/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4582760465.09._SCTHUMBZZZ_.jpg"  alt="文字逍遥 (平凡社ライブラリー)"  /></a> </p>
<p class="description">グレッグ・イーガンの名作。これも singularity を巡る物語だな。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-18">2017-09-18</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
