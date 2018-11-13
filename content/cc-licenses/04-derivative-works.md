+++
date = "2016-01-26T20:29:34+09:00"
update = "2018-11-13T13:34:22+09:00"
description = "CC Licenses では「改変禁止」条件がなければ「翻案物を作成、複製および共有すること」を許諾する。また翻案物を受け取ったユーザは原著作者の許諾を（提示されている CC Licenses の条件に従って）自動的に得る。"
draft = false
tags = ["derivative-works", "creative-commons", "copyright", "license", "tpp"]
title = "二次的著作物と CC Licenses"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  
[scripts]
  athjax = false
  ermaidjs = false
+++

今回は「二次的著作物」について。

1. [「二次的著作物」とは]({{< relref "#about" >}})
1. [CC Licenses による改変の許諾]({{< relref "#ccl" >}})
1. [Copyleft のすすめ]({{< relref "#cl" >}})
1. [パロディについて]({{< relref "#parody" >}})
1. [二次創作のみを許可したい]({{< relref "#niji" >}})
1. [「改変禁止」が意図するもの]({{< relref "#nd" >}})

## 「二次的著作物」とは{#about}

[著作権法]第2条では「二次的著作物」は以下のように定義されている。

{{< fig-quote title="著作権法" link="http://law.e-gov.go.jp/htmldata/S45/S45HO048.html" >}}
<q>著作物を翻訳し、編曲し、若しくは変形し、又は脚色し、映画化し、その他翻案することにより創作した著作物をいう。</q>
{{< /fig-quote >}}

「創作」的であることが重要で，「既存の著作物の修正増減に創作性が認められるが、原著作物の表現形式の本質的な特徴が失われるに至っていない場合[^ak]」に二次的著作物と見なされる。
たとえばフォーマット変換や機械翻訳（点字などへの置き換え）などは「逐語的コピー」と呼ばれ複製と見なされる[^cpy]。

[^ak]: {{< pdf-file title="「アンコウ行灯事件」判決文" link="http://www.courts.go.jp/app/files/hanrei_jp/827/013827_hanrei.pdf" >}}より。
[^cpy]: 他にも漫画のキャラクタなどを似せて描いた場合も，翻案ではなく，複製と見なされる場合がある。文脈とか，どの程度似てるかとかにもよるだろうけど。

具体的には

- 翻訳
- 編曲
- 変形（美術、写真、建築物、地図・図形の著作物で用いられることが多い）
- 脚色
- 映画化
- 翻案（上述した以外の全て。コミカライズやノベライズ，文章の要約，あるいはプログラムのバージョンアップや他言語への移植なども含まれる）

といった感じに分類されるが，一絡げに「翻案」または「改変」と表記されることが多い。
[CC Licenses] では二次的著作物に相当するものを「翻案物（Adapted Material）[^m]」としている。

[^m]: [CC Licenses] は著作隣接権で保護されるものも含む形で許諾しているためこのような表現になる。詳しくは「[CCライセンス・バージョン4.0 日本語版の公開 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2015/07/15/cc%E3%83%A9%E3%82%A4%E3%82%BB%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%90%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B34-0-%E6%97%A5%E6%9C%AC%E8%AA%9E%E7%89%88%E3%81%AE%E5%85%AC%E9%96%8B/)」を参照のこと。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q><strong>「翻案物」</strong>とは、著作権およびそれに類する権利の対象となり、ライセンス対象物について許諾者が有する著作権およびそれに類する権利に基づく許諾が必要とされるような形で、翻訳され、改変され、編集され、変形され、またはその他の方法により変更されたマテリアルで、ライセンス対象物から派生したか、またはライセンス対象物に基づくものを意味します。本パブリック・ライセンスにおいては、ライセンス対象物が音楽作品、実演または録音物で、これらが動画と同期させられる場合には、翻案物が常に作成されることになります。</q>
{{< /fig-quote >}}

著作（権）者は自身の著作物に対して

- 二次的著作物の作成に関する権利（[著作権法] 第27条）
- 二次的著作物の利用に関する原著作者の権利（[著作権法] 第28条）

の2つの権利を持っている。
ポイントは二次的著作物の原著作者（元の著作物の著作者）も二次的著作物に対して一連の著作財産権を持っていることである。

{{< fig-quote title="著作権法" link="http://law.e-gov.go.jp/htmldata/S45/S45HO048.html" >}}
<q>二次的著作物の原著作物の著作者は、当該二次的著作物の利用に関し、この款に規定する権利で当該二次的著作物の著作者が有するものと同一の種類の権利を専有する。</q>
{{< /fig-quote >}}

つまり，ユーザが二次的著作物を利用したいと考えるなら，二次的著作物の著作（権）者と原著作（権）者の双方から許可を得る必要がある。

## CC Licenses による改変の許諾{#ccl}

[CC Licenses] では「改変禁止 <i class="fab fa-creative-commons-nd"></i>」条件がなければ「翻案物を作成、複製および共有すること」を許諾する。
また翻案物を受け取ったユーザは原著作者の許諾を（提示されている [CC Licenses] の条件に従って）自動的に得る[^sl]。

[^sl]: [CC Licenses] ではサブライセンスを禁止しているため，翻案物に対する許諾の一方を直に原著作者から得る，という形になっている。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>あなたから翻案物を受領した者は、あなたが適用した翻案者のライセンスの条件にしたがった翻案物におけるライセンスされた権利を行使できるという申出を自動的に許諾者から受け取ります。</q>
{{< /fig-quote >}}

つまり，原マテリアルとその翻案物の双方が [CC Licenses] で許諾されているなら，双方のライセンス条件に従って翻案物も利用することができる。

なお [CC Licenses] では「表示 <i class="fab fa-creative-commons-by"></i>」条件が必須になっているため，翻案物およびその複製や翻案物に対しても原著作者のクレジット表示が必要になる[^attr]。

[^attr]: クレジット表示については[前に書いた記事]({{< ref "/cc-licenses/02-creative-commons-licenses.md#attribution" >}})を参照のこと。

## Copyleft のすすめ{#cl}

二次的著作物については原著作物の許諾条件と二次的著作物の許諾条件の両方に従う必要がある。
原著作物が [CC Licenses] で許諾されている場合でも，その翻案物が [CC Licenses] で許諾されていない場合は利用条件が著しく制限される場合がある（まぁ原著作物から派生させる手もあるのだが）。
また原著作物とその翻案物の双方が [CC Licenses] で許諾されているとしても両者の条件が異なれば，やはりそれも制限になってしまう。

もし翻案物も含めてマテリアルを広く共有したいと望むのであれば「継承 <i class="fab fa-creative-commons-sa"></i>」条件を付加することをお勧めする。
「継承 <i class="fab fa-creative-commons-sa"></i>」条件が付加されている場合は，翻案物に対しても同等のライセンス[^cmp] を付加することが求められる。

[^cmp]: 「同等のライセンス」として同じ条件の [CC Licenses] （他バージョンを含む）が挙げられる。また「[表示-継承]」については [Free Art License](http://artlibre.org/licence/lal/en/) 1.3 や [GNU GPL](https://www.gnu.org/copyleft/gpl.html)v3 も互換性のあるライセンスとして認められている。

このようなライセンスの仕組みは [Copyleft] と呼ばれている。
[Copyleft] の起源は [GNU GPL（General Public License）](https://www.gnu.org/copyleft/gpl.html) であるが， [CC Licenses] においても「継承 <i class="fab fa-creative-commons-sa"></i>」条件よって [Copyleft] もしくはそれに近いライセンスを構成できる[^cl]。

[^cl]: ただし厳密に [Copyleft] と言えるのは「[表示-継承]」条件のみである。

[CC Licenses] で「継承 <i class="fab fa-creative-commons-sa"></i>」条件が付くのは以下の2つである。

{{< fig-gen title="Creative Commons Licenses Version 4.0 International" >}}
<table>
<tbody>
<tr><th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i></a>   </th><th class='left'>表示-継承           </th><td><a href="https://creativecommons.org/licenses/by-sa/4.0/">コモンズ証</a>（<a href="https://creativecommons.org/licenses/by-sa/4.0/deed.ja">日本語</a>） <a href="https://creativecommons.org/licenses/by-sa/4.0/legalcode">法的条項</a>（<a href="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja">日本語</a>） <a href="https://creativecommons.org/licenses/by-sa/4.0/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i></th><th class='left'>表示-非営利-継承    </th><td><a href="https://creativecommons.org/licenses/by-nc-sa/4.0/">コモンズ証</a>（<a href="https://creativecommons.org/licenses/by-nc-sa/4.0/deed.ja">日本語</a>） <a href="https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode">法的条項</a>（<a href="https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode.ja">日本語</a>） <a href="https://creativecommons.org/licenses/by-nc-sa/4.0/rdf">メタデータ</a></td></tr>
</tbody>
</table>
{{< /fig-gen >}}

[Copyleft] はプログラミングの世界でも「ウイルス的」と揶揄されることもあるが，共有を維持するには必要な仕掛けだと思う。
[CC Licenses] の使われ方を見ても，[2015年時点](https://stateof.creativecommons.org/2015/)では「[表示-継承]」条件が37%で最も多い。
この機会に検討してみてはいかがだろうか。

なお [CC Licenses] は[「公正な利用（fair use）」や「著作権の制限」]({{< ref "/cc-licenses/01-copyright.md#fair-use" >}})として利用が認められていることに対しては効力が及ばないため， [CC Licenses] の条件に関わらず利用可能である。

## パロディについて{#parody}

海外では，パロディ・風刺について「公正な利用（fair use）」としてある程度認められている国もあるが，日本ではパロディに関する規定そのものがない。
そのため現状では「引用」か「翻案」かで線引されることになる[^pd]。

[^pd]: パロディはいわゆる「パクり」（一般的には「剽窃」）とは異なる。「替え歌」や短歌の「本歌取り」もパロディの一種と考えられる。米国における parody と fair use の関係については「{{< pdf-file title="米国著作権法におけるパロディとフェア・ユース/差止め請求 −パロディに関する裁判例と，小説の続編出版が問題とされた最近の事例から−" link="http://www.law.nihon-u.ac.jp/publication/pdf/chizai/4/04.pdf" >}}」が参考になる。

商業作品の場合は（防衛のために）あらかじめオリジナル作品の著作（権）者から許可を得ていることもあるようだが，そうでない場合は特に悪質なものでない限り「黙認」されているのが現状のようだ。
しかし，いったん訴訟になった場合，パロディが「引用」として認められるのはかなり難しいと思われる。

## 二次創作のみを許可したい{#niji}

[CC Licenses] では，どの条件の組み合わせでもマテリアルの複製や配布を（条件に従う限りは）制限しない。
これは [CC Licenses] を適用したマテリアルやその翻案物がインターネット上に置かれることを前提にしたものだからである[^in]。
しかしインターネットに乗らないマテリアル（紙の書籍， CD や DVD/BD でパッケージされた楽曲や映像など）は旧来の流通経路でコントロールする必要があるため [CC Licenses] とは馴染まない側面があるのも確かだ。

[^in]: そもそもインターネットはコピーの連鎖で成り立っているものだし，無理に制限しようとすればユーザ側の「使用」や「公正な利用」まで侵害しかねない。

「複製・配布は許可できないが二次創作は許可したい」という需要に対応するため，かつて [Creative Commons] では翻案のみを許可する「サンプリング・ライセンス（sampling license）」の作成が試みられた。
ただ，この試みはうまくいかなかったようで，現在サンプリング・ライセンスは retire している。

少なくとも [Creative Commons] では翻案のみを許可するライセンス・ツールは存在しない。
このような要件がある場合は独自にライセンスを構築するしかない。

### 同人マーク・ライセンス

「二次創作のみを許可したい」という要件に対して，日本での試みとして「[同人マーク・ライセンス]」がある。
これは日本の「同人市場」という特殊商慣行に特化したライセンスで，原著作物の複製・配布を禁止する代わりに，二次的著作物の作成とその複製・配布を許可している。
ただしいくつか条件があって

{{< fig-quote title="同人マーク・ライセンス 1.0" link="http://commonsphere.jp/doujin-license-1/" >}}
<q>許諾者は、利用者に対し、本作品について、全世界において、本作品の著作権の存続期間中、本ライセンスの各条項に従い、利用者自身が、二次創作同人誌を作成し、同人誌即売会において配布すること（有償および無償の場合双方を含みます。また、インターネット配信やCD、DVD等のデジタルメディアでの配布などのデジタルデータによる配布は除きます。以下同じ）を非独占的に許諾します。</q>
{{< /fig-quote >}}

とあるとおり，媒体は「二次創作同人誌」に限られ，配布経路も「同人誌即売会」に限られる。
たとえば，同人誌即売会に出品しない（本来的な意味での）二次創作同人誌は原著作物に「[同人マーク・ライセンス]」があろうと（著作（権）者から別途許諾がない限り） NG である。

最近（2013年）になって「[同人マーク・ライセンス]」が提唱された背景には「環太平洋パートナーシップ（Trans-Pacific Partnership; TPP）協定」の知財分野における「著作権の非親告罪化」が挙げられる。
もし「著作権の非親告罪化」が日本で成立すれば同人市場は壊滅状態になると予想されている[^mkt]。
いわばこれは作家たちによる自己防衛的ライセンスであると言える。

[^mkt]: さる調査によると[2015年の同人誌市場は757億円規模](http://www.gamespark.jp/article/2016/01/19/63159.html)だそうだ。そりゃあ狙われるよねぇ（笑） まぁ同人市場は同人誌だけじゃないし同人誌と言ってもいろいろあるので，二次創作が実際にどの程度かは分からないけど数十億から数百億の規模でもおかしくないだろう。

同人市場に限らず日本では事後承諾的に著作物を利用する慣習があり（たとえば[パロディ]({{< relref "#parody" >}})や勝手翻訳など），こういったものも軒並み「著作権の非親告罪化」に引っかかることになるだろう。
しかもこれに（同じく TPP で決まった）「法定賠償制度」も加われば巨額の賠償金を払わされる可能性がある。
これにより日本においても「[著作権トロル](https://www.eff.org/issues/copyright-trolls)」が台頭することになる。
むしろこちらの方が深刻である。
これについて文化庁は

{{< fig-quote title="TPPで“違法ダウンロード”適用拡大も、文化庁の審議会で再び検討か -INTERNET Watch" link="http://internet.watch.impress.co.jp/docs/news/20151111_730105.html" >}}
<q>著作権等侵害罪の一部非親告罪化については、TPP協定において非親告罪化が義務づけられている範囲及びその趣旨を踏まえつつ、我が国の二次創作文化への影響に十分配慮し、適切に非親告罪の範囲を定めること</q>
{{< /fig-quote >}}

などとしているが，翻案は「二次創作文化」だけではなく日常生活の広い範囲に関わる「活動」である。
「[著作権と著作権法]({{< ref "/cc-licenses/01-copyright.md" >}})」でも書いたが，ユーザを無視した知財政策を続けていると本当に「ガラパゴス」になっちゃうよ。

## 「改変禁止」が意図するもの{#nd}

[Creative Commons] は名前の通り「創造性の共有地[^cc]」を創り広げることを目的（のひとつ）としている。
ならば [CC Licenses] に「改変禁止 <i class="fab fa-creative-commons-nd"></i>」条件を加えることは [Creative Commons] の目的に反するのではないだろうか。

[^cc]: 翻訳は「[XML的思索：創造性の共有　機械可読ライセンスによるクリエイティブな作品の育成と流通](http://www.ibm.com/developerworks/jp/xml/library/x-think18/)」のタイトルから拝借した。リンク切れててゴメン（[Internet Archive](https://web.archive.org/web/20091202161418/http://www.ibm.com/developerworks/jp/xml/library/x-think18/)）。

これに関しては以下の記事が参考になる。

- [by-nc-nd - 雑記帳](http://d.hatena.ne.jp/ced/20060720/1153344179)

この記事によると [CC Licenses] の機能は以下のふたつに分類できる。

1. making derives (creation：創造)
2. dissemination (access：アクセス)

この分類で考えるなら「改変禁止 <i class="fab fa-creative-commons-nd"></i>」条件は主に dissemination (access) にフォーカスしたものであると解釈することができる。

{{< fig-quote title="by-nc-nd - 雑記帳" link="http://d.hatena.ne.jp/ced/20060720/1153344179" >}}
<q>ここからが思考のアクロバットを必要とする部分。インターネットは「技術的」にはコピーが自由。けれども、インターネットとは別の次元でこれまで発展してきた著作権法は、このような「コピー自由」な環境が存在することを想定していなかった。で、現時点でも「コピー自由」な環境に著作権法は対応していない。だから、「技術的」にはインターネット上で簡単にコピー出来てしまう素材も、現実世界の著作権法の元ではその利用は違反となってしまう。それを解決するのがCCのby-nc-ndで、手を加えないならコピー自由、という行為をライセンスというかたちで保証している。これにより、素材のインターネット上での頒布が可能となる。</q>
{{< /fig-quote >}}

確かに「改変禁止 <i class="fab fa-creative-commons-nd"></i>」条件を加えると [Creative Commons] が本来欲しい making derives (creation) としての機能は弱くなる。
しかし「改変禁止 <i class="fab fa-creative-commons-nd"></i>」条件を加えることによって，少なくとも（“all rights reserved” な状態と比較すれば） dissemination (access) は担保できる。

これも大事な「共有」の形である。
逃す手はないだろう。

## 参考になる（かもしれない） Web ページ

- [キャラクターの著作権法上の取扱いについて | 大島・西村・宮永商標特許事務所](http://onm-tm.jp/news/%EF%BD%91%EF%BC%97%EF%BC%8E%E3%82%AD%E3%83%A3%E3%83%A9%E3%82%AF%E3%82%BF%E3%83%BC%E3%81%AE%E8%91%97%E4%BD%9C%E6%A8%A9%E6%B3%95%E4%B8%8A%E3%81%AE%E5%8F%96%E6%89%B1%E3%81%84%E3%81%AB%E3%81%A4%E3%81%84)
- [スクウェア・エニックスの著作権侵害の可能性はグレー!?　『ハイスコアガール』問題について福井健策弁護士に話をうかがってみた｜おたぽる](http://otapol.jp/2014/09/post-1542.html)
- {{< pdf-file title="米国著作権法におけるパロディとフェア・ユース/差止め請求 −パロディに関する裁判例と，小説の続編出版が問題とされた最近の事例から−" link="http://www.law.nihon-u.ac.jp/publication/pdf/chizai/4/04.pdf" >}}
- {{< pdf-file title="二次創作における「意思表示システム」の提唱" link="http://www.bunka.go.jp/seisaku/chosakuken/seminar/contents_symposium/08/pdf/akamatsu.pdf" >}}
- [クリエイティブ・コモンズ・ライセンスのブログ翻訳のススメ](http://www.yamdas.org/column/technique/cctrans.html)
- [CC BY-SA 4.0 now one-way compatible with GPLv3 - Creative Commons Blog - Creative Commons](https://blog.creativecommons.org/2015/10/08/cc-by-sa-4-0-now-one-way-compatible-with-gplv3/) （[日本語訳](http://creativecommons.jp/2016/01/25/cc-by-sa-%EF%BC%88%E8%A1%A8%E7%A4%BA-%E7%B6%99%E6%89%BF%EF%BC%89-4-0%E3%81%8B%E3%82%89gpl-v3%E3%81%B8%E3%81%AE%E4%B8%80%E6%96%B9%E5%90%91%E3%81%AE%E4%BA%92%E6%8F%9B%E3%81%8C%E5%AE%9F%E7%8F%BE/)）
- [Creative Commons — State of the Commons 2015 — It's been a remarkable year, most notably for the more than 1.1 billion works under one of the CC licenses, CC0, or the public domain mark.](https://stateof.creativecommons.org/2015/)
- [TPPで“違法ダウンロード”適用拡大も、文化庁の審議会で再び検討か -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20151111_730105.html)
- [Trans-Pacific Partnership Would Harm User Rights and the Commons - Creative Commons - Creative Commons](https://creativecommons.org/campaigns/trans-pacific-partnership-would-harm-user-rights-and-the-commons/)

[本シリーズ]: /cc-licenses "改訂3版： CC-License について — text.Baldanders.info"
[著作権法]: http://elaws.e-gov.go.jp/search/elawsSearch/elaws_search/lsg0500/detail?lawId=345AC0000000048 "著作権法"
[Creative Commons]: https://creativecommons.org/ "Creative Commons"
[CC Licenses]: https://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"
[表示-継承]: https://creativecommons.org/licenses/by-sa/4.0/ "Creative Commons — Attribution-ShareAlike 4.0 International — CC BY-SA 4.0"
[Copyleft]: http://www.gnu.org/licenses/copyleft.html "コピーレフトって何? - GNUプロジェクト - フリーソフトウェアファウンデーション"
[同人マーク・ライセンス]: http://commonsphere.jp/doujin-license-1/ "同人マーク・ライセンス 1.0 | commonsphere"
