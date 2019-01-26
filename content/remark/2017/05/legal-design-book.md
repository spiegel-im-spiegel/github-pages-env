+++
date = "2017-05-08T02:43:31+09:00"
description = "GW の読書の課題として『[法のデザイン]』を読み始めたのだが，既に第一部で「えー？ うーん？」な感じになり挫折してしまった。"
draft = false
tags = ["book", "code", "law", "intellectual-property", "hacker-ethic", "artificial-intelligence", "creative-commons"]
title = "『法のデザイン』を斜め読みした"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

GW の読書の課題として『[法のデザイン]』を読み始めたのだが，既に第一部で「えー？ うーん？」な感じになり挫折してしまった。

最初はいちいちメモしてたんだけど早々に諦めて以降は斜め読み。
なので，『[法のデザイン]』についての感想というより読みながら頭の中で考えたことをつらつらと吐き出してみる。

あらかじめ予防線を張っておくと，私は法の専門家ではないしそれを目指しているわけでもない。
その辺を割り引いて読んでいただければ幸いである。

## 制約は構造を生む

{{< fig-quote title="数学ガール／フェルマーの最終定理" link="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/baldandersinf-22/" >}}
<q>公理によって与えられる暗黙の制約。この制約が集合の要素同士をしっかり結びつける。単純にしばるのではない、相互に秩序ある関係を結ぶ。言い換えれば――公理によって与えられる制約が構造を生み出しているのだ</q>
{{< /fig-quote >}}

『[法のデザイン]』を一言で要約するなら「制約は構造を生む」ということに尽きる。
その制約とは法（Law）・規範（Norm）・市場（Market）・アーキテクチャ（Architecture）からなる「4つの規制[^cd1]」を指す。

[^cd1]: 日本の場合，「4つの規制」は Lawrence Lessig 氏の『[CODE]』で紹介されたことで一般に認知されたと言っていいだろう。しかし背景となる新シカゴ学派の理論まで認知されているとまでは言えない状況である。

「4つの規制」は新シカゴ学派の理論大系で出てくる概念である。
中でも法は「規制を規制するもの（meta-regulator）」であり，他の3規制を法によって規制することで「主体」を規制することが可能であるとされている。
これを「共同規制（co-reguration）」と呼ぶらしい。
本のタイトルにもなっている「法のデザイン（legal design）」もこの共同規制の延長線上にあると考えていいかもしれない。

## Copyleft という法的 Hack

「自由なソフトウェア（Free Software）」運動における [copyleft] の考え方（およびそれを実装した [GNU GPL]）が優れた法的[^law1] hack である，というのは有名な話である。

[^law1]: [GNU GPL] は法律ではないが一種の契約であり「4つの規制」のうちの法に属するものと考えられる。余談だが，何故か最近になって日本で話題の [Mastodon] は [AGPL](https://www.gnu.org/licenses/agpl.html "GNU Affero General Public License - GNU Project - Free Software Foundation") で公開されている。これは SaaS 向けのソフトウェアに対しても改変バージョンの公開を求めるものである。

この [copyleft] が Linux のような製品の台頭を促し，オープンソース・ソフトウェア（Open-Source Software; OSS）という paradigm shift を引き起こし，更にはコンピュータ・ネットワークの産業構造をも変えていく。

こうした一連の過程は法そのものにも影響を与える。
たとえば米国の DMCA (Digital Millennium Copyright Act) は現在も批判の多い著作権関連法だが，適用除外既定については定期的に見直されており，ここ10年くらいでは規制が緩められる傾向にある。

- [速報！　米著作権庁、合法利用の範囲を大幅拡大―iPhone脱獄など全6項目 | TechCrunch Japan](http://jp.techcrunch.com/2010/07/27/20100726now-legal-in-the-u-s-jailbreaking-your-iphone-ripping-a-dvd-for-educational-purposes/)
- [EFF、JailbreakやSIMロック解除の合法化を勝ち取る - P2Pとかその辺のお話](http://peer2peer.blog79.fc2.com/blog-entry-1694.html)

たしかに法には「規制を規制するもの」としての側面があるが（数学の公理のように）絶対不変ではなく，「4つの規制」はお互いに影響を与えながら変質する。
「法のデザイン」は単なる設計ではなく，こうした相互作用を測定しながらチューニングを繰り返すプロセスが重要になってくる。

## 規制は規制に汚染される

「4つの規制」はお互いに影響を与えながら変質する，ということで言うと近年のバズワードのひとつになっている人工知能（AI）が挙げられるだろう。

機械学習なんてのは体のいい洗脳である。

たとえば AI が人類を皆殺しにしようとしたり[「ハイル・ヒットラー」と叫んだり]({{< ref "/remark/2015/artificial-intelligence.md" >}} "人工知能は「ハイル・ヒトラー」と叫ぶか")したら大抵の人は「これは拙い」と思うだろうが，過激でない程度にリベラル寄りの知識体系を得たり，あるいは保守的なそれを得た場合，それは拙いことなのだろか正しいことなのだろうか。
または AI 技術が商業的に使われる際，ユーザに対してサービス・ベンダに有利なマーケティング情報操作がないと証明できるのだろうか。

たとえばビッグ・データを AI 技術の助けを借りて解析し非常に綿密な消費者情報を得たとする。
ビッグ・データ自体は個人を特定する情報を排除した行動パターンのみであれば個人情報保護法には抵触しないしプライバシーの観点からの問題もないだろう。
でも解析された情報を使って特定の誰かの行動を精確に予測し先回りできるとしたら，そこに問題はないと言えるのだろうか。

機械が単なる道具であることを超え個人の活動や考え方に深く関与するようになったとき（現実にそうなりつつある）私たちはそれをどう測定し評価・改善していけばいいのだろうか。

## 空き地はコモンズか

「コモンズ（commons）」には様々なシチュエーションがある。
Google に[機械翻訳](https://translate.google.com/)させると大抵「下院」と訳される。
また辞書を引くと「入会地[^cms1]」などと説明されることが多い。

[^cms1]: 「{{< ruby "いりあいち" >}}入会地{{< /ruby >}}」は Wikipedia では「村や部落などの村落共同体で総有した土地で、薪炭・用材・肥料用の落葉を採取した山林である入会山と、まぐさや屋根を葺くカヤなどを採取した原野・川原である草刈場の2種類に大別される」などと[説明](https://ja.wikipedia.org/wiki/%E5%85%A5%E4%BC%9A%E5%9C%B0 "入会地 - Wikipedia")されている。

しかし， FLOSS や [Creative Commons] をはじめ，ネット活動におけるコモンズは「入会地」よりもう一歩踏み込んだところにあり，その考え方の源流は言うまでもなく「[伽藍とバザール]」である。
まずバザール型のコミュニティがあって，それを駆動するための方便としてコモンズという概念がある。
これは例えば企業同士が集まって行うカテドラル型の patent pool のようなものとは根本的に異なる。

一方で（ネットに限らず）日本におけるコミュニティ活動は「入会地」ですらなく，喩えるなら「空き地で遊ぶ子供[^cms2]」のようなものである。
空き地で遊ぶ子供は（公有か私有かに関わらず）その所有者を含む大人から「お目こぼし」されている状態である。
しかし当の子供はそれが本来は不法侵入であるとは微塵も思わない。
そして所有者が空き地への侵入を禁止しても別の空き地に移るだけだ。

[^cms2]: 「空き地で遊ぶ子供」といっても今時の人はピンとこないかもしれない。「ドラえもん」など昭和の作品で描写されるくらいか。

これの大規模なものが「同人市場[^cms3]」である。
[Creative Commons] について[日本で活発に議論された2003年から2004年](https://baldanders.info/spiegel/log/cclog.html "日記／Weblog でみる Creative Commons -- Spiegel's Trunk")頃にかけて同人市場についてもよく議論に上ったが，当事者たちの考え方は決まっていて「怒られたら逃げればいい」というものだった。

[^cms3]: 同人活動と言っても本来的な意味での「同好の士による閉じた活動」もあれば「パクリ上等」なものまで様々なレベルがあるので一絡げに論じるのは無茶だと分かっているのだが今回はご容赦。今では大規模なイベントにおいては国際的な認知度も上がったため「自主規制」されているし，職業作家や出版社も出品するので煩く言われることも少なくなっているように見えるが，結局は職業作家や出版社の「お目こぼし」で回っているという事実には変わらない。これは TPP 知財部門に対して[「二次創作」の部分だけに反応する](http://internet.watch.impress.co.jp/docs/news/730105.html "TPPで“違法ダウンロード”適用拡大も、文化庁の審議会で再び検討か -INTERNET Watch Watch")という点からも窺い知れる。

『[法のデザイン]』ではコモンズを法的な「余白」であるとしている。
たとえば米国の fair use 規定などは「余白」と言えるかもしれないが，日本の法律でこのような「余白」があるものは少ないし，仮にあっても行政機関が出す「命令」や「ガイドライン」で埋め尽くされてしまう。
そもそも日本の個人にとって法は（社会契約とかではなく）上から下知される「ご法度」の延長線上にあり[^law2]，その運用のされかたは主に「お目こぼし」と「見せしめ」という恣意的なものである。

[^law2]: 私は「悪法も法」とは思わないけどね。むしろ「ルールが守られないのはルール自体に問題がある」と考える。実際にルール・メイキングを行う際にポイントになるのはそこだと思う。

このような不公正な環境で（トロルならともかく）イノベーションなど生まれるはずもなく，あえてそれを狙うなら（海外に税金を払うことになっても）活動拠点（および準拠法）を海外に置くことが「法のデザイン」を考える上で賢明な選択になってしまうのだろう。

## [CCjp] という温度差

私が『[法のデザイン]』を読もうと思ったのは著者の水野祐さんが [CCjp の理事](https://creativecommons.jp/about/people/#mizuno "CCJPのメンバー | クリエイティブ・コモンズ・ジャパン")であるというのがきっかけである。

[CCjp] については，ライセンスや “[State of the Commons](https://stateof.creativecommons.org/)" の翻訳活動に対しては敬意を表すものだが，それ以外の活動や組織としてのビジョンが全く見えない。
これは本家 [Creative Commons] が明確に [Free Culture] に舵を切ったのとは対照的である。
この温度差の原因を『[法のデザイン]』から窺い知ることができないかと思ったのである。

最終的にこの本の印象は「広告記事」である。
第二部の各論はとても参考になるが参考になるだけ。
あれを読んで「よし！ 私も！」とか思う人はいないだろう。
この辺は同じ [CCjp] の理事の方が書かれた『[フリーカルチャーをつくるためのガイドブック](https://www.amazon.co.jp/exec/obidos/ASIN/4845911744/baldandersinf-22/ "フリーカルチャーをつくるためのガイドブック クリエイティブ・コモンズによる創造の循環 | ドミニク・チェン |本 | 通販 | Amazon")』とはだいぶ異なっている。
SIer とか企業の法務部向けの本って感じかな。

[CCjp] がこの本と同じ方向を向いているかどうかは分からないのだが「何処を向いているか分からない」という点では共通しているかもしれない。

## ブックマーク

- [レッシグの4つの制約要素 - 法音](http://d.hatena.ne.jp/ho-on/20110626/p1)
- [伽藍とバザール（The Cathedral and the Bazaar）](http://cruel.org/freeware/cathedral.html)
- [改訂3版： CC Licenses について](/cc-licenses/)

[法のデザイン]: http://filmart.co.jp/books/society/business/legaldesign/ "法のデザイン | 動く出版社 フィルムアート社"
[CODE]: https://www.amazon.co.jp/exec/obidos/ASIN/B01CYDGUV8/baldandersinf-22/ "CODE VERSION 2.0 | ローレンス・レッシグ, 山形浩生 | コンピュータ・IT | Kindleストア | Amazon"
[copyleft]: http://www.gnu.org/licenses/copyleft.html "コピーレフトって何? - GNUプロジェクト - フリーソフトウェアファウンデーション"
[GNU GPL]: https://www.gnu.org/copyleft/gpl.html "The GNU General Public License v3.0 - GNU Project - Free Software Foundation"
[Creative Commons]: https://creativecommons.org/ "Creative Commons"
[伽藍とバザール]: http://cruel.org/freeware/cathedral.html "伽藍とバザール（The Cathedral and the Bazaar）"
[Mastodon]: https://github.com/tootsuite/mastodon "tootsuite/mastodon: A GNU Social-compatible microblogging server"
[CCjp]: https://creativecommons.jp/ "クリエイティブ・コモンズ・ジャパン"
[Free Culture]: https://creativecommons.org/share-your-work/public-domain/freeworks/ "Understanding Free Cultural Works - Creative Commons"

## 参考図書

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B071V8J53D/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/41df%2BhdoLfL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B071V8J53D/baldandersinf-22/">法のデザイン</a></dt><dd>水野祐 </dd><dd>フィルムアート社 2017-02-28</dd><dd>評価<abbr class="rating" title="3"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-3-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00W1FH0AU/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00W1FH0AU.09._SCTHUMBZZZ_.jpg"  alt="知ろうとすること。（新潮文庫）"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B06XDM85QK/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B06XDM85QK.09._SCTHUMBZZZ_.jpg"  alt="遙かなる他者のためのデザイン ―久保田晃弘の思索と実装"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B06ZYQRMM1/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B06ZYQRMM1.09._SCTHUMBZZZ_.jpg"  alt="ウェルビーイングの設計論 ―人がよりよく生きるための情報技術"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01N9ZT9NP/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B01N9ZT9NP.09._SCTHUMBZZZ_.jpg"  alt="未来を築くデザインの思想 ―ポスト人間中心デザインへ向けて読むべき24のテキスト"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B06Y1G4X1F/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B06Y1G4X1F.09._SCTHUMBZZZ_.jpg"  alt="ブリッジング"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01N4S2LMB/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B01N4S2LMB.09._SCTHUMBZZZ_.jpg"  alt="これからの僕らの働き方　次世代のスタンダードを創る10人に聞く (早川書房)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B06XTR4GFP/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B06XTR4GFP.09._SCTHUMBZZZ_.jpg"  alt="逆説のスタートアップ思考 (中公新書ラクレ)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B071XSG4XS/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B071XSG4XS.09._SCTHUMBZZZ_.jpg"  alt="Processing クリエイティブ・コーディング入門 ―コードが生み出す創造表現"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B06XY66JS9/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B06XY66JS9.09._SCTHUMBZZZ_.jpg"  alt="完全対応 新個人情報保護法－Q＆Aと書式例－"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01N3ABWZ9/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B01N3ABWZ9.09._SCTHUMBZZZ_.jpg"  alt="ケース・スタディ ネット権利侵害対応の実務－発信者情報開示請求と削除請求－"  /></a> </p>
<p class="description">Kindle 版出てたのかよ待てばよかった orz 強いて言うなら第二部の各論は参考になる部分も多い。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-04-07">2017-04-07</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
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
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/FREE-CULTURE-%E3%83%AD%E3%83%BC%E3%83%AC%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%AC%E3%83%83%E3%82%B7%E3%82%B0-ebook/dp/B01DJ5VE0W?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01DJ5VE0W"><img src="https://images-fe.ssl-images-amazon.com/images/I/51d5%2BE9anaL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/FREE-CULTURE-%E3%83%AD%E3%83%BC%E3%83%AC%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%AC%E3%83%83%E3%82%B7%E3%82%B0-ebook/dp/B01DJ5VE0W?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01DJ5VE0W">FREE CULTURE</a></dt>
	<dd>ローレンス・レッシグ</dd>
	<dd>山形浩生 (翻訳), 守岡桜 (翻訳)</dd>
    <dd>翔泳社 2004-07-22 (Release 2016-03-28)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B01DJ5VE0W</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">Free Culture の原典。白田秀彰さんの「<a href='http://orion.mt.tama.hosei.ac.jp/hideaki/freeannotation.htm'>FREE ANNOTATION</a>」も併せてどうぞ。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-11-17">2018-11-17</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%95%E3%82%A3%E3%83%AB%E3%82%BF%E3%83%BC%E3%83%90%E3%83%96%E3%83%AB%E2%94%80%E2%94%80%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%8D%E3%83%83%E3%83%88%E3%81%8C%E9%9A%A0%E3%81%97%E3%81%A6%E3%81%84%E3%82%8B%E3%81%93%E3%81%A8-%E3%83%8F%E3%83%A4%E3%82%AB%E3%83%AF%E6%96%87%E5%BA%ABNF-%E3%82%A4%E3%83%BC%E3%83%A9%E3%82%A4%E3%83%BB%E3%83%91%E3%83%AA%E3%82%B5%E3%83%BC/dp/4150504598?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4150504598"><img src="https://images-fe.ssl-images-amazon.com/images/I/41UdjkE4OpL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%95%E3%82%A3%E3%83%AB%E3%82%BF%E3%83%BC%E3%83%90%E3%83%96%E3%83%AB%E2%94%80%E2%94%80%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%8D%E3%83%83%E3%83%88%E3%81%8C%E9%9A%A0%E3%81%97%E3%81%A6%E3%81%84%E3%82%8B%E3%81%93%E3%81%A8-%E3%83%8F%E3%83%A4%E3%82%AB%E3%83%AF%E6%96%87%E5%BA%ABNF-%E3%82%A4%E3%83%BC%E3%83%A9%E3%82%A4%E3%83%BB%E3%83%91%E3%83%AA%E3%82%B5%E3%83%BC/dp/4150504598?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4150504598">フィルターバブル──インターネットが隠していること (ハヤカワ文庫NF)</a></dt>
	<dd>イーライ・パリサー</dd>
	<dd>井口耕二 (翻訳)</dd>
    <dd>早川書房 2016-03-09</dd>
    <dd>Book 文庫</dd>
    <dd>ASIN: 4150504598, EAN: 9784150504595</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ネットにおいて私たちは自由ではなく，それと知らず「フィルターバブル」に捕らわれている。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-31">2018-12-31</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E8%91%97%E4%BD%9C%E6%A8%A9%EF%BC%92%EF%BC%8E%EF%BC%90-%E3%82%A6%E3%82%A7%E3%83%96%E6%99%82%E4%BB%A3%E3%81%AE%E6%96%87%E5%8C%96%E7%99%BA%E5%B1%95%E3%82%92%E3%82%81%E3%81%96%E3%81%97%E3%81%A6-NTT%E5%87%BA%E7%89%88%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E3%83%BC%E2%80%95%E3%83%AC%E3%82%BE%E3%83%8A%E3%83%B3%E3%83%88-%E5%90%8D%E5%92%8C-%E5%B0%8F%E5%A4%AA%E9%83%8E/dp/4757102852?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4757102852"><img src="https://images-fe.ssl-images-amazon.com/images/I/41YkbcP5IyL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E8%91%97%E4%BD%9C%E6%A8%A9%EF%BC%92%EF%BC%8E%EF%BC%90-%E3%82%A6%E3%82%A7%E3%83%96%E6%99%82%E4%BB%A3%E3%81%AE%E6%96%87%E5%8C%96%E7%99%BA%E5%B1%95%E3%82%92%E3%82%81%E3%81%96%E3%81%97%E3%81%A6-NTT%E5%87%BA%E7%89%88%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E3%83%BC%E2%80%95%E3%83%AC%E3%82%BE%E3%83%8A%E3%83%B3%E3%83%88-%E5%90%8D%E5%92%8C-%E5%B0%8F%E5%A4%AA%E9%83%8E/dp/4757102852?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4757102852">著作権２．０ ウェブ時代の文化発展をめざして (NTT出版ライブラリー―レゾナント)</a></dt>
	<dd>名和 小太郎</dd>
    <dd>エヌティティ出版 2010-06-24</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4757102852, EAN: 9784757102859</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">名著です。今すぐ買うべきです。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-11-13">2018-11-13</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
