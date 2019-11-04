+++
title = "『もうすぐ絶滅するという開かれたウェブについて 続・情報共有の未来』を読む（ボーナストラック以外）"
date = "2019-01-04T19:53:34+09:00"
description = "具体的な話はせず与太話をいくつか書くに留めておくことにするので悪しからずご了承のほどを。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "code", "internet", "generativity", "grigori", "surveillance-capitalism" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

『もうすぐ絶滅するという開かれたウェブについて [続・情報共有の未来]』（以後「[続・情報共有の未来]」と略称する）をよーやく読み終わったですよ。

買ったのも2018年の5月に入ってからだし，私的事情で積ん読期間が結構あったということもあるけど，50章は（長いとか多いとかじゃなくて）デカい！ 読み終わった直後はクラクラしたですよ（笑） ちなみに2018年9月時点で v1.0.1 がリリースされている（追記：2019年3月に [v1.1.1 が出た]({{< ref "/remark/2019/03/infoshare2.md" >}} "『続・情報共有の未来』は付録から読むのがオススメ") ）。

今回も [Dynalist] でメモを取りながら読んだ。
最初に目次を拾ってコピペし，章ごとに思ったことをメモったり，気になるフレーズを引用したりする。
「技術的保護手段」で縛られた本ではできない芸当であり，ホンマ，[達人出版会]様グッジョブ！ って感じである。

[Dynalist] に残したメモを眺めると，思ったより量が多くて，これを全部記事に盛り込んでたら文字数が膨大になりそう。
なので具体的な話はせず与太話をいくつか書くに留めておくことにする。

なお，ボーナストラックの感想は書かない（多分）。
こちらは一言「泣ける！」とだけ記しておこう。

[^b1]: 「数学ガール」シリーズの感想文が例外だが，アレも小説としての感想は一言くらいしか書いてない。

## 「情報力」をめぐるゲーム

他の記事でも書いているが，SF作家（？）の野尻抱介さん曰く「黄色矮星人は2人いれば力比べを始める」のだそうだ。
「情報共有の未来」とは「情報力」という power をめぐるゲームだと考えれば分かりやすい。
これは白田秀彰さんの[「グリゴリの捕縛」を読み直す]({{< ref "/remark/2017/10/too-many-ghosts.md" >}} "今こそ「グリゴリの捕縛」を読め！ または遍在する草薙素子")ことで腑に落ちた。

ゲームであればルールが必要である。
つまりルール（＝法）メイキングがとても重要ということだ。

インターネットは基本的にコピーの連鎖でできている。
コピーによってデータと情報を血液のように循環させるわけだ。
したがってコピーを支配する「もの」がインターネットにおける勝者になれる。
国際的な著作権ルールが[アクセス・コントロールの独占]({{< ref "/remark/2018/11/copy-control-and-access-control.md" >}} "「技術的保護手段」と「技術的利用制限手段」")へと野放図に向かっているのもそれ故だろう。

## ぼくがかんがえたさいきょうのいんたねっと史

私が考えるに，インターネット・ユーザに要求される唯一のルールは「インターネットに参加する」ことであり，そこに付随するコストとリスクと責任を（手の届く範囲で）引き受けることである。
インターネットは（コンビニとかで`w`）買うものではなく，作る（make）ものだ。

本来インターネットに接続するマシンは，網または網間を繋ぐ「ノード」として機能する。
しかしノードを構築・維持するには相応のコストがかかる。
回線だって安くないし，そもそもカセットテープやフロッピー・ディスクで運用しているマイコンやパソコンなんか全くお呼びではなかった。
でも，この「参入障壁」こそが初期インターネットを守っていたのである。

これが崩れたのが1990年代のインターネット商用化で，多くのユーザは ISP にぶら下がる「奴隷端末」になることと引き換えに「インターネットに参加する」ことなくインターネットに接続してデータや情報を得ることが可能になった。

この状況を更に後押しするのがクラウドと SNS を含む XaaS である。
ゼロ年代以降に台頭した各種 XaaS は「独占的なマスター・ノードとそれにぶら下がる数多の奴隷端末」という構造を強化したわけだ。

市場原理に基づいて駆動するインターネットは「情報力」を背景とした新たな階級を生み出し「支配なき統制」とも言うべき構造を構築した。
たとえば Bitcoin/Blockchain は最初こそ脱中央集権的な通貨システムとして期待されたが，結局は[新たな階級社会を駆動するシステム]({{< ref "/remark/2018/05/internet-as-a-class-system.md" >}} "階級社会としてのインターネット")となった。

私を含む多くのユーザはクラウドや XaaS を通じてインターネットに「参加」しているつもりかも知れないが，本当はそれをただ「使用」しているに過ぎないかも知れない。
テレビ・リモコンの4色ボタンを押すだけではテレビというシステムに参加しているとは言わない。

## [エンジニアこそ「狂狷の徒」たれ]({{< ref "/remark/2017/12/hacker-ethic.md" >}}) {#hacker-ethic}

『[はやぶさ―不死身の探査機と宇宙研の物語](https://www.amazon.co.jp/exec/obidos/ASIN/4344980158/baldandersinf-22)』にこんな一節がある。

{{< fig-quote title="はやぶさ―不死身の探査機と宇宙研の物語" link="https://www.amazon.co.jp/exec/obidos/ASIN/4344980158/baldandersinf-22" >}}
<q>理学は、真理の探究であり、工学は善の実現である。そして、藝術は美の表現である－－これで所謂「真美善」が揃う</q>
{{< /fig-quote >}}

何を以って「善」とするかは[非常に難しい](https://baldanders.info/blog/000242/ "本当にヤバいのは「技術者」ではない人達だと思う — Baldanders.info")ところだが，少なくとも職業エンジニアとしての私はこれを自らの矜持として行動してきたつもりである。

『[情報共有の未来]』も，その続編たる『[続・情報共有の未来]』も [generativity] をキーワードにしている。
もし [generativity] を「善」と位置付けるなら，如何にしてそれを実装するか試行錯誤するのが「狂狷の徒」たるエンジニアがインターネットに「参加」する条件だろう。

書籍として『[続・情報共有の未来]』を読み通して最初に感じたのは「私ってインターネットを信用していないんだなぁ」ということだった。

私から見て初期インターネットの背景にはヒッピー文化があると思うが，結果としてその文化は市場経済に塗りつぶされた。
でも初期インターネットを夢想する「老害」たちは当時の「黄金時代」を忘れられず，かといって「ネットのテレビ化」も止められないという状況に陥っている。

これがこの本を読んで感じたインターネットに対する印象だ。
もしインターネットが「幼年期」を抜けて [generativity] に象徴される「成年期」に向かおうというのなら「開かれたウェブ」が「もうすぐ絶滅する」というのは，ひょっとしたら（痛みは伴うとしても）歓迎すべき事態なのかも知れない。

まぁ，そのとき私はもう[ネットにいない]({{< ref "/remark/2016/12/real-is-not-therapy.md" >}} "ネットには居場所がないということ")かも知れないが。

## 参考になる（かもしれない）ページ{#ref}

以下に『[続・情報共有の未来]』の Web 連載時に反応した記事や関連しそうな記事を挙げてみた。
参考になれば（私が）ラッキーということで。

- [本当にヤバいのは「技術者」ではない人達だと思う — Baldanders.info](https://baldanders.info/blog/000242/)
- [ディズニーランド的コモンズと Creative Commons — Baldanders.info](https://baldanders.info/blog/000401/)
- [『情報共有の未来』を読む — Baldanders.info](https://baldanders.info/blog/000611/)
- [『インテンション・エコノミー』を読む — Baldanders.info](https://baldanders.info/blog/000638/)
- [恐るべき結婚 — Baldanders.info](https://baldanders.info/blog/000692/)
- [「ネット上の中傷やヘイトスピーチにどう対処すべきか」より — Baldanders.info](https://baldanders.info/blog/000737/)
- [俺達がインターネットだ！ — Baldanders.info](https://baldanders.info/blog/000794/)

- [『つながりっぱなしの日常を生きる』を読む]({{< ref "/remark/2016/05/its-complicated.md" >}})
- [『フィルターバブル』を読む]({{< ref "/remark/2016/05/filter-bubble.md" >}})
- [“The Shadow Web” （再掲載）]({{< ref "/remark/2016/11/the-shadow-web.md" >}})
- [「さよならはてなダイアリー」 ― 黒 Web 2.0 の終焉]({{< ref "/remark/2017/06/goodbye-hatena-diary.md" >}})
- [ようやく『反知性主義』を読んだ]({{< ref "/remark/2016/09/anti-intellectualism.md" >}})
- [AI は人（の良心）を殺すか？]({{< ref "/remark/2016/10/artificial-intelligence.md" >}})
- [ネットには居場所がないということ]({{< ref "/remark/2016/12/real-is-not-therapy.md" >}})
- [「インターネットと発酵」]({{< ref "/remark/2016/12/fermented-internet.md" >}})
- [今こそ「グリゴリの捕縛」を読め！ または遍在する草薙素子]({{< ref "/remark/2017/10/too-many-ghosts.md" >}})
- [タイムラインの奴隷 - Spiegel's Branch](https://scrapbox.io/spiegel-branch/%E3%82%BF%E3%82%A4%E3%83%A0%E3%83%A9%E3%82%A4%E3%83%B3%E3%81%AE%E5%A5%B4%E9%9A%B7)
- [エンジニアこそ「狂狷の徒」たれ]({{< ref "/remark/2017/12/hacker-ethic.md" >}})
- [「ネットの中立性」と「後出しジャンケン」と「メディアによる検閲を迂回する」]({{< ref "/remark/2017/12/hacker-ethic-2.md" >}})
- [本当に「良心」を無意味化する時代がやってくる？]({{< ref "/remark/2017/12/artificial-intelligence.md" >}})
- [『AI vs. 教科書が読めない子どもたち』を読む]({{< ref "/remark/2018/02/artificial-intelligence-book.md" >}})
- [「仮想通貨」と公開鍵基盤]({{< ref "/remark/2018/02/blockchain-and-pki.md" >}})
- [誰がプライバシーを支配するのか]({{< ref "/remark/2018/04/handling-privacy.md" >}})
- [階級社会としてのインターネット]({{< ref "/remark/2018/05/internet-as-a-class-system.md" >}})
- [「情報弱者」を再々定義する]({{< ref "/remark/2018/05/information-illiterate.md" >}})
- [年末年始に施行される改正著作権法に関する覚え書き]({{< ref "/remark/2018/11/copyright-law-is-revised.md" >}})

## ブックマーク

- [グリゴリの捕縛 あるいは 情報時代の憲法について](http://orion.mt.tama.hosei.ac.jp/hideaki/kenporon.htm)
- [群衆の知恵・集団的知性とWikiコラボレーション](https://www.slideshare.net/tsukamoto/wiki-2984796)
- [MEMO: 集団的知性（Collective Intelligence）と、群衆の知恵（Wisdom of Crowds）の違い](http://kzk-memo.blogspot.com/2010/04/collective-intelligencewisdom-of-crowds.html)
- [総務省、IoT機器に「認証マーク」導入へ　サイバー攻撃急増で - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1712/26/news104.html)
- [クーリエ連載；エコノミスト紹介、自由のためなら人が死んでもいい](https://cruel.org/economist/courier200712.html)
- [『もうすぐ絶滅するという開かれたウェブについて 続・情報共有の未来』への反応 その18 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20190115/openweb) : あざっす！
- [『もうすぐ絶滅するという開かれたウェブについて 続・情報共有の未来』が国会図書館に納本された - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20190609/openweb)
- [サイゾー2019年8月号の『クロサカタツヤのネオ・ビジネス・マイニング』第67回になぜか登場 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20190715/cyzo) : 本の宣伝を兼ねて出演されたらしい（笑）

[情報共有の未来]: https://www.yamdas.org/infoshare/ "『情報共有の未来』サポートページ"
[続・情報共有の未来]: https://www.yamdas.org/infoshare2/ "『もうすぐ絶滅するという開かれたウェブについて 続・情報共有の未来』サポートページ"
[達人出版会]: https://tatsu-zine.com/
[Dynalist]: https://dynalist.io/
[generativity]: http://shift-inc.co.jp/gtl/generativity/ "SHIFT Inc » Generativity"

## 参考図書

<div class="hreview" >
	<div class="photo"><a class="item url" href="https://tatsu-zine.com/books/infoshare2"><img src="https://tatsu-zine.com/images/books/877/cover_s.jpg" alt="photo"></a></div>
    <dl class="fn">
      <dt><a href="https://tatsu-zine.com/books/infoshare2">もうすぐ絶滅するという開かれたウェブについて 続・情報共有の未来</a></dt>
      <dd>yomoyomo</dd>
      <dd>達人出版会 2017-12-25</dd>
      <dd>評価&nbsp;<abbr class="rating fa-sm" title="4">
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="far fa-star"></i>
      </abbr></dd>
    </dl>
    <p class="description"><a href="https://wirelesswire.jp/author/yomoyomo/">WirelessWire News 連載</a>の書籍化。感想は<a href="/remark/2019/01/infoshare2/">こちら</a></p>
	<p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018-12-31</abbr></p>
</div>

<div class="hreview" >
	<div class="photo"><a class="item url" href="https://tatsu-zine.com/books/infoshare"><img src="https://tatsu-zine.com/images/books/15/cover_s.jpg" alt="photo"></a></div>
    <dl class="fn">
      <dt><a href="https://tatsu-zine.com/books/infoshare">情報共有の未来</a></dt>
      <dd>yomoyomo</dd>
      <dd>達人出版会 2011-12-30</dd>
      <dd>評価&nbsp;<abbr class="rating fa-sm" title="4">
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="far fa-star"></i>
      </abbr></dd>
    </dl>
    <p class="description">同名ブログの書籍化。感想は<a href="https://baldanders.info/blog/000611/">こちら</a>。</p>
    <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018-12-31</abbr></p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4344980158?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41pq6GlY+GL._SL160_.jpg" width="100" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4344980158?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">はやぶさ―不死身の探査機と宇宙研の物語 (幻冬舎新書)</a></dt>
    <dd>吉田 武 (著)</dd>
    <dd>幻冬舎 2006-11-01</dd>
    <dd>新書</dd>
    <dd>4344980158 (ASIN), 9784344980150 (EAN), 4344980158 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">宇宙研（ISAS）の歴史とともに日本の宇宙開発について解説する。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-31">2018-12-31</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E5%8F%8D%E7%9F%A5%E6%80%A7%E4%B8%BB%E7%BE%A9%E2%80%95%E3%82%A2%E3%83%A1%E3%83%AA%E3%82%AB%E3%81%8C%E7%94%9F%E3%82%93%E3%81%A0%E3%80%8C%E7%86%B1%E7%97%85%E3%80%8D%E3%81%AE%E6%AD%A3%E4%BD%93%E2%80%95%EF%BC%88%E6%96%B0%E6%BD%AE%E9%81%B8%E6%9B%B8%EF%BC%89-%E6%A3%AE%E6%9C%AC-%E3%81%82%E3%82%93%E3%82%8A-ebook/dp/B012VRLPRG?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B012VRLPRG"><img src="https://images-fe.ssl-images-amazon.com/images/I/41-khbugqTL._SL160_.jpg" width="106" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E5%8F%8D%E7%9F%A5%E6%80%A7%E4%B8%BB%E7%BE%A9%E2%80%95%E3%82%A2%E3%83%A1%E3%83%AA%E3%82%AB%E3%81%8C%E7%94%9F%E3%82%93%E3%81%A0%E3%80%8C%E7%86%B1%E7%97%85%E3%80%8D%E3%81%AE%E6%AD%A3%E4%BD%93%E2%80%95%EF%BC%88%E6%96%B0%E6%BD%AE%E9%81%B8%E6%9B%B8%EF%BC%89-%E6%A3%AE%E6%9C%AC-%E3%81%82%E3%82%93%E3%82%8A-ebook/dp/B012VRLPRG?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B012VRLPRG">反知性主義―アメリカが生んだ「熱病」の正体―（新潮選書）</a></dt>
	<dd>森本 あんり</dd>
    <dd>新潮社 2015-02-20 (Release 2015-08-14)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B012VRLPRG</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">アメリカの近代思想史または宗教史といった位置付けだろうか。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-31">2018-12-31</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4757143044?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/413qoSjODUL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4757143044?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">信頼と裏切りの社会</a></dt>
    <dd>ブルース・シュナイアー (著), 山形 浩生 (翻訳)</dd>
    <dd>NTT出版 2013-12-24</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4757143044 (ASIN), 9784757143043 (EAN), 4757143044 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">社会における「信頼」とは。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-23">2019-03-23</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4150504598?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41Le2d7nWiL._SL160_.jpg" width="109" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4150504598?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">フィルターバブル──インターネットが隠していること (ハヤカワ文庫NF)</a></dt>
    <dd>イーライ・パリサー (著), 井口耕二 (翻訳)</dd>
    <dd>早川書房 2016-03-09</dd>
    <dd>文庫</dd>
    <dd>4150504598 (ASIN), 9784150504595 (EAN), 4150504598 (ISBN), 9784150504595 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ネットにおいて私たちは自由ではなく，それと知らず「フィルターバブル」に捕らわれている。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-31">2018-12-31</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/B0125TZSZ0?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/616sjle5ITL._SL160_.jpg" width="109" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/B0125TZSZ0?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">つながりっぱなしの日常を生きる：ソーシャルメディアが若者にもたらしたもの</a></dt>
    <dd>ダナ・ボイド (著), 野中 モモ (翻訳)</dd>
    <dd>草思社 2014-10-09 (Release 2015-07-21)</dd>
    <dd>Kindle版</dd>
    <dd>B0125TZSZ0 (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">読むのに1年半以上かかってしまった。ネット，特に SNS 上で自身のアイデンティティやプライバシーを保つにはどうすればいいか。豊富な事例を交えて考察する。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-31">2018-12-31</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/B01J1I8PRQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/5132MaXTkGL._SL160_.jpg" width="112" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/B01J1I8PRQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">社会は情報化の夢を見る (河出文庫)</a></dt>
    <dd>佐藤俊樹 (著)</dd>
    <dd>河出書房新社 2010-09-03 (Release 2016-07-29)</dd>
    <dd>Kindle版</dd>
    <dd>B01J1I8PRQ (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">1996年に出版された『ノイマンの夢・近代の欲望―情報化社会を解体する』の改訂新装版。しかし内容はこれまでと変わりなく，繰り返し語られる技術決定論を前提とする安易な未来予測を「情報化」社会論だとして批判する。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-31">2018-12-31</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4000280872?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/31e2h91IUWL._SL160_.jpg" width="103" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4000280872?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">イノベーション 悪意なき嘘 (双書 時代のカルテ)</a></dt>
    <dd>名和 小太郎 (著)</dd>
    <dd>岩波書店 2007-01-11</dd>
    <dd>単行本</dd>
    <dd>4000280872 (ASIN), 9784000280877 (EAN), 4000280872 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">「両用技術とはどのようなものか。その核心には「矛と楯の論理」がある」（まえがき「予断・診断・独断 そんなばかな」より）</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-31">2018-12-31</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

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
