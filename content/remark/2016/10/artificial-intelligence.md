+++
title = "AI は人（の良心）を殺すか？"
date = "2016-10-12T00:39:24+09:00"
description = "「正義」とか「良心」とかいうのは人の不完全さの象徴である。その不完全さが人のアイデンティティだというのなら，そんな種は進化の階梯を機械に譲ったほうがいい。でも実際には，私たちは種としてもっと先に進めるはずである，と私は思う。"
tags = ["artificial-intelligence", "engineering", "politics"]

[scripts]
  mathjax = false
  mermaidjs = false
+++

（[yomoyomo] さんの最近の記事は文章や記事の間の「行間」が絶妙に繋がってる感じで色々と触発されます。
まぁ私のは妄言・妄想ですけどw）

1. [良心について]({{< relref "#conscience" >}})
1. [「良心」の無意味化]({{< relref "#nothing" >}})
1. [「信頼に値するアルゴリズム」とは（追記：2016-10-15）]({{< relref "#alg" >}})
1. [AI は「トロッコ問題」をどう解くか（再掲載）]({{< relref "#trolley-problem" >}})
1. [IV 型ウイルス（追記：2016-10-18）]({{< relref "#virus" >}})
1. [ブックマーク]({{< relref "#bookmark" >}})
1. [参考図書]({{< relref "#reference" >}})

## 良心について{#conscience}

ふだん私たちは何の前提もなく「良心」という言葉を使う。
そして人には必ず「良心」が存在すると思っている。
で，「良心」のない（あるいはないように見える）人に対して「ひとでなし」とか「冷血漢」とか言って排除しようとする。
でも，排除される側からすれば「他人を「ひとでなし」と言う貴方の「良心」って何なの？」と思ってしまうわけだ。

実際には「良心」というのは人の中にあるのではなく人と人の間にあるものである。
そしてこれを「人と機械」という関係に拡張したらどうなるのだろう。

「良心」は哲学の問題であり社会なんとか学の問題だった。
でも「人と機械」の間に良心を組み込もうとするなら，それは工学の問題にもなってくる。
はたして工学は「良心」を定義できるのだろうか。

## 「良心」の無意味化{#nothing}

- [ユートピアのキモさと人工知能がもたらす不気味の谷 - WirelessWire News（ワイヤレスワイヤーニュース）](https://wirelesswire.jp/2016/06/54411/)
- [人工知能は人間を人間でなくすのか？ - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20161010/kobayashihideo)

感情とか表情とかいうものは「他者」を認識して初めて意味を成す。
世界の始まりから自分一人しか存在しないとして，喜怒哀楽を感じたり，ましてやそれを表現することに意味があるとは思えない。

人は言葉のみにて話すわけではない。
身振り手振りや顔の表情，匂いや触覚など五感を駆使して「会話」し，お互いに文脈を形成・共有する。
機械が HMI (Human Machine Interface) を備えるなら人の表情（顔の表情だけとは限らない）やそれに伴う感情を読み取ろうとするのは自然な流れだし，たぶん必要なことである。
そうしなければ機械と人との間で文脈を形成できない。
機械が人の感情に対してどう「思う」かは別にして。

問題があるとするならそれは機械の側にではなく，あくまで機械を含むシステム（もっと言えば社会）を設計する側にある[^a]。

[^a]: もちろん機械が自律的にものを創造できるようになれば話は変わってくるが，とりあえず今はそんな心配はない（笑） 技術の両義性の問題は昔から言われていることではあるが，少なくとも「技術が世界を変える」という幻想はそろそろ捨てるべきだと思う。それは近代の「科学万能主義」とさして変わらない発想である。人の社会を動かし変えるのはあくまで人の意思と行動であり，インターネットや SNS といったものは「背中をもうちょっと押してくれる」程度に考えるべきだろう。それはそれで重要なことなんだけど。

たとえば，秘密を持つことや嘘をつくことがある文脈において必要な情報コントロール（嘘も方便）であるなら，技術的に可能だからといってわざわざそれを暴き立てることがシステムにとってどのような結果をもたらすか，といったことである。
あるユーザに関する evidence のみが必要なのであれば，そのユーザが嘘をついているかどうか推測するのは「過剰」だ。

確かに小林秀雄さんの「「良心」の無意味化」の話[^b] は考えさせられるところもあるが，現代は（もちろんインターネットも含めて）既にそういう時代になっている。照魔鏡ほど威圧的ではないかもしれないが。

[^b]: 小林秀雄著『[考えるヒント](https://www.amazon.co.jp/exec/obidos/ASIN/4167107120/baldandersinf-22/)』の「良心」より：  「もし、嘘発見機に止まらず、これが人間観察装置として、例えば、閻魔の持っている照魔鏡のような性能を備えるに至ったならどうなるだろうか。その威力に屈服しない人間はいなくなるだろう。誰にも悪い事は出来なくなるだろうが、その理由はただ為（し）ようにも出来ないからに過ぎず、良心を持つ事は、誰にも無意味な事になろう。」

{{< fig-quote title="排除型社会" link="https://www.amazon.co.jp/exec/obidos/ASIN/4903127044/baldandersinf-22/" >}}
<q>後期近代社会における社会統制の基調にあるもの、それは「保険統計主義」である。
すでにみたように〔第2章の表2-2（一一九頁）〕、ここでは正義を追求することよりも被害を最小限にすることが求められている。
そして犯罪や逸脱の原因を探ったところで犯罪という社会問題は解決しないとみなされている。
保険統計主義の中心にあるのはリスク計算である。
それは精度の高い確率論的解析であり、そこで注意が向けられるのは問題の原因ではなく、その問題が起こる蓋然性である。
保険統計主義にとって重要なのは、正義ではなく、被害の最小化である。
それが目的とするのは、世界から犯罪をなくすことではなく、損害を最小限にする効果的手段である。
それが追求するのは、ユートピアをつくりだすことではなく、敵意に満ちたこの世界に塀で囲まれた小さな楽園をできるだけ多くつくりだすことである。</q>
（p.170）
{{< /fig-quote >}}

この社会のどこに「良心」の余地があるというのだろう。
人が「性善」だろうが「性悪」だろうが関係ない。
私たちは「保険統計主義」が示す「正解」のとおりに行動しなければ排除されてしまうのだから。
そしてそれは近代文明が求めた結果なのである。
良いか悪いかは別にして。

AI (Artificial Intelligence) はたまたま再流行してこれから何か影響を及ぼすのではなく「保険統計主義」な時代の要請で必然的に台頭してきている[^c] （もちろん AI が台頭することによって他の様々なことが変化するだろうけど，それは別の話）。
ある分野やシーンで機械が台頭してくればそこで職を失う人は当然出るだろうし（それは私かもしれない），「「良心」の無意味化」は AI の登場に関わらずとっくに起きている。

[^c]: 「[頼むからプログラミングを学ばないでくれ | TechCrunch Japan](http://jp.techcrunch.com/2016/05/17/20160510please-dont-learn-to-code/)」 より： 「プログラミングが注目されたことにより，問題を理解することよりも，それを「正しい方法で」解決することに重きが置かれるようになった。」

## 「信頼に値するアルゴリズム」とは（追記：2016-10-15）{#alg}

上で

> 「問題があるとするならそれは機械の側にではなく，あくまで機械を含むシステム（もっと言えば社会）を設計する側にある。」

などと無責任に書いたが，実際には並大抵のことではない。

- [我々は信頼に足るアルゴリズムを見極められるのか？ - WirelessWire News（ワイヤレスワイヤーニュース）](https://wirelesswire.jp/2016/10/56935/)
- {{< pdf-file title="Equality of Opportunity in Supervised Learning" link="https://drive.google.com/file/d/0B-wQVEjH9yuhanpyQjUwQS1JOTQ/view" >}}
    - [Google、AIから偏見を排除する方法を研究中 | TechCrunch Japan](http://jp.techcrunch.com/2016/10/10/20161007google-aims-to-prevent-discriminatory-ai-with-equality-of-opportunity-method/)
- [機械と人間の役割分担を見つめ直してみよう | TechCrunch Japan](http://jp.techcrunch.com/2017/01/16/20170114putting-the-intelligent-machine-in-its-place/)

機械学習なんて体のいい「洗脳」であるが，その洗脳から必然的に生まれる「偏見」を数学的手法で排除しようとするのがいかにも Google らしい。
でも，そうしたロジックも含めて実装し評価するのは人なのである（今のところ）。

[yomoyomo] さんの「[我々は信頼に足るアルゴリズムを見極められるのか？](https://wirelesswire.jp/2016/10/56935/)」ではオライリーによる「特定のアルゴリズムが信頼に値するか評価する指針」を紹介している。
曰く，以下の4つだ。

1. アルゴリズムの作者がどんな出力結果を求めているか明確にし、外部の観察者がその出力結果を検証できる。
2. そのアルゴリズムが成功しているか測定可能である。
3. アルゴリズムの作者の目指すものが、そのアルゴリズムの利用者の目指すものと足並みを揃えている。
4. そのアルゴリズムは、その作者と利用者をより良い長期的な意思決定に導くか？

1 から 3 は要するに「基準（criteria）」からの「逸脱（deviation）」がないかということなのだが，そもそもその criteria が正しいものなのか評価することが難しい。
テストすることは設計することと同じなのである（TDD 的にはね）。

そして 4 を満たす製品・サービスはほとんどど無いと思われる。
何故なら長期に渡って要求が変わらないシステムなどありえないし，まず前提として「信頼」の評価は過去の事象に対してのみ可能であり未来については「予測」しかできない。
予測であれば必ず「予断」が混入するものである。
バグのないプログラムがほとんど不可能であるように予断のない予測もほぼ不可能だ[^ana]。

[^ana]: 予断だらけの予測を「{{< ruby "うらない" >}}天数演繹{{< /ruby >}}」と言う（笑） 私は「アナリスト」は「競馬の予想屋」と同業種だと思っている人なのであしからず。

私なら以下の5つ目を加える。

- 失敗に対する対処が比較的容易である。

長く運用していれば必ず何かの「失敗」が起きる。
失敗が起きた時にベンダは指を咥えて見ているだけでユーザは泣き寝入りするしかないシステムは許容できない。

## AI は「トロッコ問題」をどう解くか（再掲載）{#trolley-problem}

（[以前書いた記事]({{< ref "/remark/2015/1103-diary.md#trolley-problem" >}})をちょっと弄って再掲載[^replay]）

[^replay]: 自己防衛的な措置として，このサイトでは安易に「知性（intellect）」という言葉を使わないことにした。今まで私は知性は「知能（intelligence）」の延長線上にあるものとしてあまり区別してなかった（哲学（つか神学）上の intellectus と intelligentia の違いがピンとこなくて）。が，どうもそういう解釈をしている人は少なそうである。なかには知能の反対が知性だという人もいるっぽい。そういう前提条件が異なる状態では議論にも参考にもならない（議論するためには前提となるものの擦りあわせが必要）。このサイトの記事についても遡って修正することにする。本家サイトは（若気の至りということで）とりあえず放置。まぁ「知能」の概念が大きく乖離することはあるまい。「知能＝学力」みたいに勘違いしてる人はいるかもだが。

- [完全自動運転自動車とトロッコ問題について](http://blogos.com/article/142284/)
- [自動運転車の法律問題を総括すると見えてくる難解な課題 - 風観羽　情報空間を羽のように舞い本質を観る](http://d.hatena.ne.jp/ta26/20151104)

これって「[われはロボット](https://www.amazon.co.jp/exec/obidos/ASIN/B00O1VK072/baldandersinf-22/)」だよね。

明らかに「正しい解」がない場合，いくつかの近似解の中から妥当と思われるものを選ぶしかない[^b2]。
「緊急避難」というのは「正しい解」が存在しない場合に「近似解でいいんだよ」ということを法的に担保するものだ[^c2]。
でも，どの解を選んでも結局「正しい解」ではないのだ。
だから人は葛藤し，さらに後悔する。

[^b2]: もちろん「何も選ばない」というのも選択肢のひとつである。
[^c2]: しかし，法的に担保されているからといって倫理・道徳的に問題がないとは限らない。しかも倫理観・道徳観念というのは，特に個人主義が進んだ現代では，かつての「大きな物語（meta-narrative）」ほどには機能しない。

AI は（今のところ）近似解に葛藤したりしない。
もちろん後悔だってしない。
その解に辿り着いたのは，機械が自ら考えたのではなく，あくまでも構築された論理と学習に沿って必然的に導かれたものなのだから（たとえ解への道筋が人には理解できないものだとしても）。
じゃあ機械が導き出した解を実行した結果の責任は誰が取るの？ ってことである。

## IV 型ウイルス（追記：2016-10-18）{#virus}

毎度『BOOM TOWN』ネタで申し訳ないが，この作品の中に AI に感染するウイルスが登場する。
このウイルスは AI の学習パターンを少しずつ変質させるため発見が遅れることが多く，ある日突然 AI が使い物にならなくなるという厄介なものだ。

まぁ，ここまで極端でなくても AI に干渉する要素というのはいくらでもあるものだ。
しかもそこに悪意の有無は関係ない。

- [MicrosoftがAIチャットボット、Tayを停止―人種差別ジョークで機械学習の問題点が明らかに | TechCrunch Japan](http://jp.techcrunch.com/2016/03/25/20160324microsoft-silences-its-new-a-i-bot-tay-after-twitter-users-teach-it-racism/)

これは機械が人とのコミュニケーションを行うようになれば必ず起こり得る問題であり，それを「修正」しようとすれば更に恣意が入り込むことになる。
もし仮に人と機械との間に「良心」が組み込まれたとしても，それが書き換えられる可能性は常にあるのだ。（でもそれは人だって同じことだけどね）

## ブックマーク{#bookmark}

- [人工知能は「ハイル・ヒトラー」と叫ぶか]({{< ref "/remark/2015/artificial-intelligence.md" >}})
- [「ピノキオ」と AI]({{< ref "/remark/2016/08/pinocchio.md" >}})
- [コンピュータが次々と間違える時代。](http://www.ne.jp/asahi/comp/tarusan/main260.htm)
- [なぜいま「ロボット倫理学」が必要か〜問題はすでに起きている（岡本 慎平） | 現代ビジネス | 講談社](http://gendai.ismedia.jp/articles/-/50660)
- [人類と人工知能が共存する未来への第一歩--IBMが「AI倫理」を解説 - ZDNet Japan](https://japan.zdnet.com/article/35108815/)
- [フェイクニュースより恐ろしい、アルゴリズムの「偏見」とは何か？(塚越健司) - 個人 - Yahoo!ニュース](https://news.yahoo.co.jp/byline/tsukagoshikenji/20171013-00076828/)

[yomoyomo]: http://d.hatena.ne.jp/yomoyomo/ "YAMDAS現更新履歴"

## 参考図書{#reference}

{{% review-paapi "4903127044" %}} <!-- 排除型社会 -->
{{% review-paapi "4000280872" %}} <!-- イノベーション 悪意なき嘘 -->

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E8%80%83%E3%81%88%E3%82%8B%E3%83%92%E3%83%B3%E3%83%88-%E5%B0%8F%E6%9E%97-%E7%A7%80%E9%9B%84-ebook/dp/B009HO50XC?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B009HO50XC"><img src="https://images-fe.ssl-images-amazon.com/images/I/41LkjyVMoKL._SL160_.jpg" width="111" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E8%80%83%E3%81%88%E3%82%8B%E3%83%92%E3%83%B3%E3%83%88-%E5%B0%8F%E6%9E%97-%E7%A7%80%E9%9B%84-ebook/dp/B009HO50XC?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B009HO50XC">考えるヒント</a></dt>
	<dd>小林 秀雄</dd>
    <dd>文藝春秋 1974-06-25 (Release 2012-09-20)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B009HO50XC</dd>
  </dl>
  <p class="description">学生時代に受験対策で読んでるはずなんだけどなぁ。あんまり憶えてない。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-10-12">2016-10-12</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%BE%E3%81%A8%E3%82%81%E8%B2%B7%E3%81%84-%E3%82%AD%E3%82%AB%E3%82%A4%E3%83%80%E3%83%BC02%EF%BC%88%E8%A7%92%E5%B7%9D%E3%82%B3%E3%83%9F%E3%83%83%E3%82%AF%E3%82%B9%E3%83%BB%E3%82%A8%E3%83%BC%E3%82%B9%EF%BC%89/dp/B013UQUH80?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B013UQUH80"><img src="https://images-fe.ssl-images-amazon.com/images/I/C1RBaQMQg4S._SL160_.png" width="160" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%BE%E3%81%A8%E3%82%81%E8%B2%B7%E3%81%84-%E3%82%AD%E3%82%AB%E3%82%A4%E3%83%80%E3%83%BC02%EF%BC%88%E8%A7%92%E5%B7%9D%E3%82%B3%E3%83%9F%E3%83%83%E3%82%AF%E3%82%B9%E3%83%BB%E3%82%A8%E3%83%BC%E3%82%B9%EF%BC%89/dp/B013UQUH80?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B013UQUH80">[まとめ買い] キカイダー02（角川コミックス・エース）</a></dt>
	<dd>ＭＥＩＭＵ</dd>
	<dd>石ノ森 章太郎 (その他)</dd>
    <dd></dd>
    <dd>Ebook Bundle Kindle版</dd>
    <dd>ASIN: B013UQUH80</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">「キカイダー」を神秘学的視点で再解釈する（笑）</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-08-18">2016-08-18</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%8F%E3%82%8C%E3%81%AF%E3%83%AD%E3%83%9C%E3%83%83%E3%83%88%E3%80%94%E6%B1%BA%E5%AE%9A%E7%89%88%E3%80%95-%E3%82%A2%E3%82%A4%E3%82%B6%E3%83%83%E3%82%AF-%E3%82%A2%E3%82%B7%E3%83%A2%E3%83%95-ebook/dp/B00O1VK072?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00O1VK072"><img src="https://images-fe.ssl-images-amazon.com/images/I/51UzGYXJ70L._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%8F%E3%82%8C%E3%81%AF%E3%83%AD%E3%83%9C%E3%83%83%E3%83%88%E3%80%94%E6%B1%BA%E5%AE%9A%E7%89%88%E3%80%95-%E3%82%A2%E3%82%A4%E3%82%B6%E3%83%83%E3%82%AF-%E3%82%A2%E3%82%B7%E3%83%A2%E3%83%95-ebook/dp/B00O1VK072?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00O1VK072">われはロボット〔決定版〕</a></dt>
	<dd>アイザック アシモフ</dd>
	<dd>小尾芙佐 (翻訳)</dd>
    <dd>早川書房 2014-04-25 (Release 2014-09-30)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00O1VK072</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ロボットや AI の SF ならこれが古典で定番か？ 面白かったら続けて『<a href="https://www.amazon.co.jp/%E9%8B%BC%E9%89%84%E9%83%BD%E5%B8%82-%E3%82%A2%E3%82%A4%E3%82%B6%E3%83%83%E3%82%AF-%E3%82%A2%E3%82%B7%E3%83%A2%E3%83%95-ebook/dp/B00O2O7JFY?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00O2O7JFY">鋼鉄都市</a>』も読むとよい。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-11-03">2015-11-03</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

{{% review-paapi "B00NWQI4N4" %}} <!-- BOOM TOWN TRIP.30 -->
