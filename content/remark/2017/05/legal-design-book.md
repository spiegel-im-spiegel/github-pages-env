+++
title = "『法のデザイン』を斜め読みした"
date = "2017-05-08T02:43:31+09:00"
description = "GW の読書の課題として『[法のデザイン]』を読み始めたのだが，既に第一部で「えー？ うーん？」な感じになり挫折してしまった。"
tags = ["book", "code", "law", "intellectual-property", "hacker-ethic", "artificial-intelligence", "creative-commons"]

[scripts]
  mathjax = false
  mermaidjs = false
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

- [速報！　米著作権庁、合法利用の範囲を大幅拡大―iPhone脱獄など全6項目 | TechCrunch Japan](https://techcrunch.com/2010/07/26/now-legal-in-the-u-s-jailbreaking-your-iphone-ripping-a-dvd-for-educational-purposes/)
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

{{% review-paapi "B071V8J53D" %}} <!-- 法のデザイン -->
{{% review-paapi "B01CYDGUV8" %}} <!-- CODE VERSION 2.0 -->
{{% review-paapi "B01DJ5VE0W" %}} <!-- FREE CULTURE -->
{{% review-paapi "4150504598" %}} <!-- フィルターバブル -->
{{% review-paapi "4757102852" %}} <!-- 著作権２．０ ウェブ時代の文化発展をめざして -->
