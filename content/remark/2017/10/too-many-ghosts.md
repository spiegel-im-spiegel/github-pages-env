+++
title = "今こそ「グリゴリの捕縛」を読め！ または遍在する草薙素子"
date =  "2017-10-14T01:12:32+09:00"
description = "折角なので私も何か書いてみる。オチはないです。駄文ですみません。"
tags = [
  "code",
  "law",
  "hacker-ethic",
  "internet",
  "politics",
  "artificial-intelligence",
  "blockchain",
  "grigori"
]

[scripts]
  mathjax = false
  mermaidjs = false
+++

（なんか Twitter で呼ばれたようなので）

久しぶりの [ced] さんの文章だ。

- [草薙素子は何人いるか？ - Is a ghost countable or uncountable? - 雑記帳](http://d.hatena.ne.jp/ced/20171010/1507615755)

折角なので私も何か書いてみる。
オチはないです。
駄文ですみません。

## 今こそ「[グリゴリの捕縛]」を読め！

- [グリゴリの捕縛 あるいは 情報時代の憲法について](http://orion.mt.tama.hosei.ac.jp/hideaki/kenporon.htm)

いやぁ，久しぶりに白田秀彰さん[^t1] の「[グリゴリの捕縛]」を読んだですよ。
青空文庫に登録されているとは気づかなかったです。
ついでに MOBI 版も Kindle に入れてしまいました。
Amazon Drive の “send-to-kindle” フォルダにぶち込めば楽々♪

[^t1]: えーっと，私は（一方的なものも含めて）師事している人や業務上利害関係のある人以外で相手を「先生」呼ばわりすることはしないので悪しからず。尊敬するしないとは別ものです。

- [図書カード：グリゴリの捕縛](http://www.aozora.gr.jp/cards/000021/card4307.html)
- [白田 秀彰先生のePub・Kindle本まとめ - P2P today ダブルスラッシュ](https://wslash.com/?page_id=4729)
    - [白田先生の「グリゴリの捕縛 あるいは 情報時代の憲法について」のePub版を無料配布します。 - P2P today ダブルスラッシュ](https://wslash.com/?p=4478)
    - [リクエストがあったので、白田先生の「グリゴリの捕縛 あるいは 情報時代の憲法について」のKindle版を無料配布します。 - P2P today ダブルスラッシュ](https://wslash.com/?p=4482)

15年以上経った今読んでも十分通用する内容。
総選挙期間中のこの時期に読んだのはよかった。
初心にかえった気分だよ。
用語や歴史イベントについても中学・高校の子なら分かるんじゃないかな。
つか18歳から選挙権がもらえる（あるいはもらえた）高校生には一般知識として是非読んでほしい。

内容は ~~怪獣大決戦~~ おっと憲法（基本法）についてのお話。

古代社会 → 中世社会 → 近代社会 → 現代社会 と順を追って法と慣習そして力（power）との関係について解説し，その中で憲法（基本法）がどのように望まれ実装されていったか指摘してる。
その後，現代社会の次のパラダイム（paradigm; つまり現在と未来）に顕現する「情報力」と社会との関係に言及していくわけだ。

注意すべきなのは1点だけ。
この論文が2001年に書かれたものだということ。
ちょうど 9.11 が起きた年。
当然ながら「9.11 以後」については言及がない。
「9.11 以後」の国際政治および軍事の風景はガラッと変わった。
近代先進国と新興国以外の第3国も核を持てるに至り，核による抑止力が緩みつつある。
中東を中心とした地域はまた[100年前]({{< ref "/remark/2016/07/02-stories.md#sp" >}})を繰り返そうとしている。

でも「9.11 以後」を盛り込んだとしても論旨は殆ど変わらないと思う。
改憲云々の議論が盛んだが，まずはこれを読んで，これから先の未来における power balance をどう設計していくのか考えていけばいいと思う。

あー。
アニメの「[絶対防衛レヴィアタン](http://leviathan-anime.net/)」が観たくなった。

ちなみに今回はアウトライン・エディタ [Dynalist] を使ってメモをとりながら読んだ。
こういう用途なら使えるな。
先に目次をざざーっと攫っていって，それに（読んで気がついたこととかの）メモを追記していく感じ。
余計な DRM とかないからこそ，こういうことができるのですよ。

## 遍在する草薙素子

さて，それを踏まえて，改めまして。

- [草薙素子は何人いるか？ - Is a ghost countable or uncountable? - 雑記帳](http://d.hatena.ne.jp/ced/20171010/1507615755)

### スマートな機械は「後悔」するか

この話に出てくる人工知能「A」は私がよく言う「スマートな機械」である。

{{< fig-quote title="草薙素子は何人いるか？" link="http://d.hatena.ne.jp/ced/20171010/1507615755" >}}
<q>「A」は診断の結果、「Aは他者とのコミュニケーションを行う手段を持たない」ことが判明する。「A」は他者という存在を理解することができない。これは重大な問題を発生させる。<br>
これは異なる人という意味での、つまり物理的に別の身体を持つ他人とのコミュニケーションができないだけではなく、今現在の自分以外の過去の自分という他人ともコミュニケーションができないことを意味する。 [...] 「A」には、いつもリアルタイムの、最新の自分しかいない。</q>
{{< /fig-quote >}}

たとえば「トロッコ問題」というのがある。

{{< fig-quote title="完全自動運転自動車とトロッコ問題について" link="http://blogos.com/article/142284/" >}}
<q>トロッコ問題とは、線路を走るトロッコが制御不能で止まれなくなり、そのまま走ると線路の先の5人の作業員を轢いてしまうが、線路の分岐点で進路を変えると、その先の1人の作業員を轢いてしまうとき、進路を変えることが正しいか否か、といった思考実験だ。</q>
{{< /fig-quote >}}

トロッコ問題のポイントは，どちらを選択するのが正しいかではなく，「人間はどちらを選んでも後悔する」ということだ。
でも「A」は「トロッコ問題」に後悔することはない。
どちらを選ぶにせよ，その時の判断と判断に至るアルゴリズムが全てだから。

もし「A」がこの時の記録を記した「本」を発見し，同じ状況をシミュレーションした上で当時とは異なる判断に至った場合，はたしてスマートな機械「A」は「後悔」するのだろうか。
もし後悔するのであれば「A」は，少なくとも自らの時間軸において，「自己」を獲得したと言える。

### 「ヒトは誰かがいてはじめて自分を証明できる」（By “BOOM TOWN”）

人工知能「B」は哲学者が言うところの「[強い人工知能]({{< ref "/remark/2017/09/the-myth-of-the-singularity.md" >}} "『シンギュラリティの神話』を読む")」である[^sm1]。

[^sm1]: 「[Spiegelさんへの回答。](http://d.hatena.ne.jp/ced/20171015/1508027814 "Spiegelさんへの回答。 - 雑記帳")」で指摘していただいて気がついたけど，この記事では「スマートな機械」と「強い人工知能」を対置するように記述してしまっているが，哲学的には「強い人工知能」に対置されるのは「弱い人工知能（＝そもそもの意味での人工知能）」となる。私が「スマートな機械」と書くのはビジネス用語としての「AI」に対する強烈な皮肉である（参考：[「AIを搭載」は「全て自然」同様の技術的ナンセンスだ | TechCrunch Japan](http://jp.techcrunch.com/2017/01/11/20170110ai-powered-is-techs-meaningless-equivalent-of-all-natural/)）

{{< fig-quote title="草薙素子は何人いるか？" link="http://d.hatena.ne.jp/ced/20171010/1507615755" >}}
<q>「B」には日本国籍と住民票が紐づけられ、人間と同様に所有権を持ち、納税の義務がある。 [...] この観点では、「個人」としての「B」は国家があるから定義されるのであって、国家の先に(a prioriに)「個人」があるわけではない。</q>
{{< /fig-quote >}}

ところで，動物愛護家が近年よく使うキーワードに “non-human person” というのがある。
すなわち，特定の動物も「（人間ではない）人」とみなして一定の「人権」を付与すべき，という考え方である。
しかもこの言葉，政治的に効果を上げているのだ。

- [non-human person の「人権」 — Baldanders.info](https://baldanders.info/blog/000788/)

実はこれと似たような話で機械（＝人工知能？）も人のように見なすべきという議論がある。

- [Robots in Europe May Become "Electronic Persons"](https://futurism.com/robots-in-europe-may-become-electronic-persons/)

もちろんこれは「動物愛護精神」を機械にまで適用しようとかいう与太話ではない。
機械が人の社会に（たとえば失業者が増えるような）インパクトを与えるのなら，機械にも「権利と義務」を負わせるべきだ，という主張である[^pl1]。

[^pl1]: たとえば「人工知能に製造物責任を負わせるべきか」といった議論にも通じると思う。

機械が単なる「道具」としての領分（domain）を超えようとするとき，社会はそれを「道具」に押しとどめようとするのか。
それとも「個人」として遇するのか。
いずれにしろ「[シンギュラリティの神話]」が夢想する「強い人工知能」は（仮に実現するとしても）技術要件を満たせば自然発生するというわけではなく，あくまでも社会との関係の中で「構築」されていくものだ，ということだろう。

### N人の草薙素子

{{< fig-quote title="草薙素子は何人いるか？" link="http://d.hatena.ne.jp/ced/20171010/1507615755" >}}
<q>義体①に紐づくゴースト①が、米国では義体②に紐づくゴースト①だった場合、草薙素子は1人の場合と、2人の場合がある。草薙素子の主観上では、草薙素子は１人となる。日本とアメリカの国家が、国民の管理用に、義体側に対してユニークなIDを振るなら、草薙素子は２人になるが、国際条約を結び、IDをゴーストの側に振ることができれば、日本とアメリカの国家にとっても、草薙素子は１人となる。では、米国にいるゴーストが②だった場合は？</q>
{{< /fig-quote >}}

これって寧ろタチコマ（フチコマ）を連想してしまうのだが（笑）

{{< fig-img src="https://photo.baldanders.info/flickr/image/36958536904_m.jpg" title="天然オイルについて語るタチコマ" link="https://photo.baldanders.info/flickr/36958536904/" >}}

タチコマは記憶をリンクし個体差がなくなるよう運用されているが，全体で統合されたひとつの人格（？）を持っているというわけではない。
まぁ，後にエージェント機能により個体差も発生するのだが。

似たような存在としては竹本泉作品に時々登場する執事ロボット「アレックス」がいる。
アレックスは外見も中身も全く同一の機体として製造され，個々を識別する製造番号もない[^tkpk1]。
とはいえタチコマのように記憶をリンクするわけではない。
そういう意味じゃ手塚治虫さんの「ロビタ」もコピーだよな。

[^tkpk1]: アレックスの設定は『[てきぱきワーキン♥ラブ](https://www.amazon.co.jp/exec/obidos/ASIN/4757700423/baldandersinf-22/)』で語られている。

### コピーは許可，Forkは不許可

ここ数年で最大の技術トピックは何といっても Blockchain だろう[^bc1]。
Blockchain の最大の特徴は P2P (Peer-to-Peer) ネットワーク上に遍在しながらも唯一の存在であることだ。
これは Blockchain が誰の所有物でもなく誰も（国家も）制御できないことを意味する。

[^bc1]: ここでは Bitcoin の Blockchain を想定する。 Blockchain 自体はすでに Bitcoin 固有の技術というわけではないためこの想定には無理があるが，やはり Blockchain を最大限に特徴づけている実装は Bitcoin だろうということで。

- [ギリシャがユーロを捨ててBitcoinに切り替えてはいけない理由 | TechCrunch Japan](http://jp.techcrunch.com/2015/03/03/20150228why-greece-should-not-switch-to-bitcoin/)

「[グリゴリの捕縛]」ではインターネットについて

{{< fig-quote title="グリゴリの捕縛" link="http://orion.mt.tama.hosei.ac.jp/hideaki/kenporon.htm" >}}
<q>インターネットは、初め国家が使うためのネットワークとして開発が開始されました。 [...] ところが、実際にインターネットを作り上げた技術者(やはり彼らも「ハッカー」と呼ばれていました)たちは、 このインターネットの構造をとても中央集権的にコントロールしにくいように設計し ました。</q>
{{< /fig-quote >}}

と記述されているが，実際にはインターネットは（「[黄金時代](https://www.technologyreview.jp/s/22020/the-internet-is-sick/ "MIT Tech Review: インターネットの黄金時代をどう次世代に残せばいいのか？")」と呼ばれる1990年代には既に）そのような構成になっていない。
これが最悪の形で顕在化したのが「アラブの春」である。
当時，市民運動で揺れるエジプトではたった5本の電話で8000万人のインターネット・アクセスが遮断された。

{{< fig-quote title="『日経サイエンス』2012年6月号 介入されないもうひとつのウェブ" >}}
<q>もし今日のインターネットが実際にこの理論に近い状態であれば，メッシュネットワークは余計ものだったろう。だがインターネットが当初の学術目的から踏み出して現在のような誰でも使える商業サービスになってから20年以上が経つうちに，そうした蓄積伝送の原理が果たす役割は，一貫して縮小していった。<br>
この間，ネットワークに加わる新たなノードの圧倒的多数はISPを介してネットに接続する家庭や企業のコンピューターだった。ISPの接続モデルでは，利用者のコンピューターはデータの中継はしない。それはネットワークの端末，つまりデータの送受信だけを，常にISPのコンピューターを介して行うターミナル・ノードだ。言い換えれば，インターネットの爆発的な成長はネットワーク地図に行き止まりのルートを増やしただけで，新たなルートを加えることはほとんどなかった。<br>
そしてISPなど大量の情報ルートを持つ者は，彼らがルートを提供している何百万ものノードを支配下におくこととなった。これらのノードは，もしISPがダウンしたり，ネットから遮断されたりすると，その障害を回避する方法がない。ISPはインターネットが停止しないようにするどころか，実効上は停止スイッチになってしまった。</q>
{{< /fig-quote >}}

このことへの抵抗のひとつが P2P ネットワークである。
決して（ゼロ年代に批判されたような）脱著作権を目論んで考えられたものではないのだ。

P2P ネットワークは「中央集権的ネットワーク」であるインターネットの上に実質的（virtual[^v1]）な分散ネットワークを構築する。
そして P2P ネットワークのもうひとつの特徴はネットワークの各ノードにコピーを置くことである。
コピーを置くことで障害耐性（fault tolerance）を上げているのだが，このことにより P2P ネットワーク上の「もの」の**場所**が無意味化された。
遍在してしまったわけだ。

[^v1]: “virtual” を「仮想的」と訳したのが[誤訳](http://d.hatena.ne.jp/oraccha/20140926/1411694081 "「Virtualを仮想と誤訳した責任は我々にあります」 - Plan9日記")だったというのはもはや有名な話なので，敢えてこう訳してみた（笑）

### 遍在する草薙素子

もしn人の草薙素子が統合されたひとつの人格を持っているのなら，彼女が何処にいるかという問いは意味を成さなくなる。
それは当初インターネットを（既存の power の及ばない）分散ネットワークとして構築しようとした人たちが夢見たことだ。

遍在するということは誰にも（国家にも）所有されないということであり，言い換えれば「何処にもいない」ことと同義でもある。
国家が「個人」という概念を規定するのなら，遍在してしまった彼女は「個人」として規定できないことになる。

## 【追記】 レスポンスをいただきました

[ced] さんからレスポンスをいただきました。

- [Spiegelさんへの回答。 - 雑記帳](http://d.hatena.ne.jp/ced/20171015/1508027814)

ありがたいやら，こんな駄文で申し訳ないやら。
多分これから「GHOST IN THE SHELL」でバトーさんが叫ぶ度に涙腺が緩むと思います（笑）

## ブックマーク

- [「グリゴリの捕縛」に関する公開往復書簡](http://orion.mt.tama.hosei.ac.jp/hideaki/kenpoletter.htm)

- [AI は「トロッコ問題」をどう解くか]({{< ref "/remark/2016/10/artificial-intelligence.md#trolley-problem" >}})
- [『シンギュラリティの神話』を読む]({{< ref "/remark/2017/09/the-myth-of-the-singularity.md" >}})
- [『法のデザイン』を斜め読みした]({{< ref "/remark/2017/05/legal-design-book.md" >}})
- [孵卵器の中のインターネット]({{< ref "/remark/2017/01/internet-in-the-incubator.md" >}})
- [“The Shadow Web” — Baldanders.info](https://baldanders.info/blog/000599/)


[グリゴリの捕縛]: http://orion.mt.tama.hosei.ac.jp/hideaki/kenporon.htm "グリゴリの捕縛 あるいは 情報時代の憲法について"
[シンギュラリティの神話]: {{< ref "/remark/2017/09/the-myth-of-the-singularity.md" >}} "『シンギュラリティの神話』を読む"
[ced]: http://d.hatena.ne.jp/ced/ "雑記帳"
[Dynalist]: https://dynalist.io/

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.aozora.gr.jp/cards/000021/card4307.html"><img src="https://text.baldanders.info/images/aozora/card4307.svg" width="110" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.aozora.gr.jp/cards/000021/card4307.html">グリゴリの捕縛</a></dt>
    <dd>白田 秀彰</dd>
    <dd> 2001-11-26 (Release 2014-09-17)</dd>
    <dd>青空文庫</dd>
    <dd>4307 (図書カードNo.)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">白田秀彰さんの「<a href="http://orion.mt.tama.hosei.ac.jp/hideaki/kenporon.htm">グリゴリの捕縛</a>」が青空文庫に収録されていた。
内容は <delete>怪獣大決戦</delete> おっと憲法（基本法）についてのお話。
古代社会 → 中世社会 → 近代社会 → 現代社会 と順を追って法と慣習そして力（power）との関係について解説し，その中で憲法（基本法）がどのように望まれ実装されていったか指摘してる。
その後，現代社会の次のパラダイムに顕現する「情報力」と社会との関係に言及していくわけだ。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-30">2019-03-30</abbr> (powered by <a href="https://aozorahack.org/">aozorahack</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/CODE-VERSION-2-0-%E3%83%AD%E3%83%BC%E3%83%AC%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%AC%E3%83%83%E3%82%B7%E3%82%B0-ebook/dp/B01CYDGUV8?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01CYDGUV8"><img src="https://images-fe.ssl-images-amazon.com/images/I/31Q2jh%2B5SgL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/CODE-VERSION-2-0-%E3%83%AD%E3%83%BC%E3%83%AC%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%AC%E3%83%83%E3%82%B7%E3%82%B0-ebook/dp/B01CYDGUV8?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01CYDGUV8">CODE VERSION 2.0</a></dt>
	<dd>ローレンス・レッシグ</dd>
	<dd>山形浩生 (翻訳)</dd>
    <dd>翔泳社 2007-12-19 (Release 2016-03-14)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B01CYDGUV8</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">前著『CODE』改訂版。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-11-17">2018-11-17</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%8D%E3%83%83%E3%83%88%E3%81%AE%E6%B3%95%E3%81%A8%E6%85%A3%E7%BF%92-%E3%81%8B%E3%81%AA%E3%82%8A%E5%A5%87%E5%A6%99%E3%81%AA%E6%B3%95%E5%AD%A6%E5%85%A5%E9%96%80-%E3%82%BD%E3%83%95%E3%83%88%E3%83%90%E3%83%B3%E3%82%AF%E6%96%B0%E6%9B%B8-%E7%99%BD%E7%94%B0-%E7%A7%80%E5%BD%B0/dp/4797334673?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4797334673"><img src="https://images-fe.ssl-images-amazon.com/images/I/41tEJvyOO2L._SL160_.jpg" width="97" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%8D%E3%83%83%E3%83%88%E3%81%AE%E6%B3%95%E3%81%A8%E6%85%A3%E7%BF%92-%E3%81%8B%E3%81%AA%E3%82%8A%E5%A5%87%E5%A6%99%E3%81%AA%E6%B3%95%E5%AD%A6%E5%85%A5%E9%96%80-%E3%82%BD%E3%83%95%E3%83%88%E3%83%90%E3%83%B3%E3%82%AF%E6%96%B0%E6%9B%B8-%E7%99%BD%E7%94%B0-%E7%A7%80%E5%BD%B0/dp/4797334673?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4797334673">インターネットの法と慣習　かなり奇妙な法学入門 [ソフトバンク新書]</a></dt>
	<dd>白田 秀彰</dd>
    <dd>ソフトバンククリエイティブ 2006-07-15</dd>
    <dd>Book 新書</dd>
    <dd>ASIN: 4797334673, EAN: 9784797334678</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ライアカ本。 Web 2.0 真っ只中に書かれた本だが，時事的な部分を除けば古びてはいないと思う。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-11-20">2018-11-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%A4%E3%83%B3%E3%83%86%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%B3%E3%83%BB%E3%82%A8%E3%82%B3%E3%83%8E%E3%83%9F%E3%83%BC%EF%BD%9E%E9%A1%A7%E5%AE%A2%E3%81%8C%E6%94%AF%E9%85%8D%E3%81%99%E3%82%8B%E7%B5%8C%E6%B8%88-Doc-Searls-ebook/dp/B00DIM6BE6?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00DIM6BE6"><img src="https://images-fe.ssl-images-amazon.com/images/I/519%2BkIHb71L._SL160_.jpg" width="111" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%A4%E3%83%B3%E3%83%86%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%B3%E3%83%BB%E3%82%A8%E3%82%B3%E3%83%8E%E3%83%9F%E3%83%BC%EF%BD%9E%E9%A1%A7%E5%AE%A2%E3%81%8C%E6%94%AF%E9%85%8D%E3%81%99%E3%82%8B%E7%B5%8C%E6%B8%88-Doc-Searls-ebook/dp/B00DIM6BE6?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00DIM6BE6">インテンション・エコノミー～顧客が支配する経済</a></dt>
	<dd>Doc Searls</dd>
	<dd>栗原潔 (翻訳)</dd>
    <dd>翔泳社 2013-03-14 (Release 2013-06-20)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00DIM6BE6</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">時代はソーシャル CRM から VRM へ。<a href='https://baldanders.info/blog/000794/'>俺達がインターネットだ！</a></p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-04-26">2015-04-26</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%80%90%E4%B8%AD%E6%9D%B1%E5%A4%A7%E6%B7%B7%E8%BF%B7%E3%82%92%E8%A7%A3%E3%81%8F%E3%80%91-%E3%82%B5%E3%82%A4%E3%82%AF%E3%82%B9-%E3%83%94%E3%82%B3%E5%8D%94%E5%AE%9A-%E7%99%BE%E5%B9%B4%E3%81%AE%E5%91%AA%E7%B8%9B-%E6%96%B0%E6%BD%AE%E9%81%B8%E6%9B%B8/dp/4106037866?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4106037866"><img src="https://images-fe.ssl-images-amazon.com/images/I/51QsC2WBr5L._SL160_.jpg" width="106" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%80%90%E4%B8%AD%E6%9D%B1%E5%A4%A7%E6%B7%B7%E8%BF%B7%E3%82%92%E8%A7%A3%E3%81%8F%E3%80%91-%E3%82%B5%E3%82%A4%E3%82%AF%E3%82%B9-%E3%83%94%E3%82%B3%E5%8D%94%E5%AE%9A-%E7%99%BE%E5%B9%B4%E3%81%AE%E5%91%AA%E7%B8%9B-%E6%96%B0%E6%BD%AE%E9%81%B8%E6%9B%B8/dp/4106037866?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4106037866">【中東大混迷を解く】 サイクス=ピコ協定 百年の呪縛 (新潮選書)</a></dt>
	<dd>池内 恵</dd>
    <dd>新潮社 2016-05-27</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4106037866, EAN: 9784106037863</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">欧州および中東の近代および現代を「サイクス=ピコ協定」を特異点として網羅的に解説していいる。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-02">2016-07-02</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%9D%E3%82%8D%E3%81%9D%E3%82%8D%E3%80%81%E4%BA%BA%E5%B7%A5%E7%9F%A5%E8%83%BD%E3%81%AE%E7%9C%9F%E5%AE%9F%E3%82%92%E8%A9%B1%E3%81%9D%E3%81%86-%E6%97%A9%E5%B7%9D%E6%9B%B8%E6%88%BF-%E3%82%B8%E3%83%A3%E3%83%B3%EF%BC%9D%E3%82%AC%E3%83%96%E3%83%AA%E3%82%A8%E3%83%AB-%E3%82%AC%E3%83%8A%E3%82%B7%E3%82%A2-ebook/dp/B071FHBGW8?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B071FHBGW8"><img src="https://images-fe.ssl-images-amazon.com/images/I/51cD7DR87IL._SL160_.jpg" width="111" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%9D%E3%82%8D%E3%81%9D%E3%82%8D%E3%80%81%E4%BA%BA%E5%B7%A5%E7%9F%A5%E8%83%BD%E3%81%AE%E7%9C%9F%E5%AE%9F%E3%82%92%E8%A9%B1%E3%81%9D%E3%81%86-%E6%97%A9%E5%B7%9D%E6%9B%B8%E6%88%BF-%E3%82%B8%E3%83%A3%E3%83%B3%EF%BC%9D%E3%82%AC%E3%83%96%E3%83%AA%E3%82%A8%E3%83%AB-%E3%82%AC%E3%83%8A%E3%82%B7%E3%82%A2-ebook/dp/B071FHBGW8?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B071FHBGW8">そろそろ、人工知能の真実を話そう (早川書房)</a></dt>
	<dd>ジャン＝ガブリエル ガナシア</dd>
	<dd>小林 重裕・他 (翻訳), 伊藤 直子 (監修)</dd>
    <dd>早川書房 2017-05-25 (Release 2017-05-31)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B071FHBGW8</dd>
    <dd>評価<abbr class="rating fa-sm" title="3">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">シンギュラリティは起きない。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-02">2016-07-02</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/BOOM-TOWN-TRIP-30-%E5%86%85%E7%94%B0-%E7%BE%8E%E5%A5%88%E5%AD%90-ebook/dp/B00NWQI4N4?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00NWQI4N4"><img src="https://images-fe.ssl-images-amazon.com/images/I/51Ia%2B77IpiL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/BOOM-TOWN-TRIP-30-%E5%86%85%E7%94%B0-%E7%BE%8E%E5%A5%88%E5%AD%90-ebook/dp/B00NWQI4N4?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00NWQI4N4">BOOM TOWN　TRIP.30</a></dt>
	<dd>内田 美奈子</dd>
    <dd> 2014-09-26 (Release 2014-09-26)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00NWQI4N4</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">掲載誌「コミックガンマ」が休刊になって単行本収録できなかった<del>まるぼし</del>まぼろしの30話。これが Kindle で読めるとはいい時代になったものです。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-06-05">2016-06-05</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%A6%E3%81%8D%E3%81%B1%E3%81%8D%E3%83%AF%E3%83%BC%E3%82%AD%E3%83%B3%E3%80%93%E3%83%A9%E3%83%96-5-%E3%83%93%E3%83%BC%E3%83%A0%E3%82%B3%E3%83%9F%E3%83%83%E3%82%AF%E3%82%B9-%E7%AB%B9%E6%9C%AC-%E6%B3%89/dp/4757700423?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4757700423"><img src="https://images-fe.ssl-images-amazon.com/images/I/518JGPT6DEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%A6%E3%81%8D%E3%81%B1%E3%81%8D%E3%83%AF%E3%83%BC%E3%82%AD%E3%83%B3%E3%80%93%E3%83%A9%E3%83%96-5-%E3%83%93%E3%83%BC%E3%83%A0%E3%82%B3%E3%83%9F%E3%83%83%E3%82%AF%E3%82%B9-%E7%AB%B9%E6%9C%AC-%E6%B3%89/dp/4757700423?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4757700423">てきぱきワーキン〓ラブ (5) (ビームコミックス)</a></dt>
	<dd>竹本 泉</dd>
    <dd>エンターブレイン 2000-05-01</dd>
    <dd>Book コミック</dd>
    <dd>ASIN: 4757700423, EAN: 9784757700420</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ついに（「さよパラ」にも出てきた）アレックスの謎が解ける？</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-10-13">2017-10-13</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
