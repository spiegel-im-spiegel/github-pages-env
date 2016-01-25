+++
date = "2016-01-25T14:54:34+09:00"
description = "description"
draft = true
tags = ["remark"]
title = "二次的著作物と CC-Licenses"

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
+++

今回は「二次的著作物」について。

## 「二次的著作物」とは

[著作権法]第2条では「二次的著作物」は以下のように定義されている。

{{< fig-quote title="著作権法" link="http://law.e-gov.go.jp/htmldata/S45/S45HO048.html" >}}
<q>著作物を翻訳し、編曲し、若しくは変形し、又は脚色し、映画化し、その他翻案することにより創作した著作物をいう。</q>
{{< /fig-quote >}}

「創作」的であることが重要で，「既存の著作物の修正増減に創作性が認められるが、原著作物の表現形式の本質的な特徴が失われるに至っていない場合[^ak]」に二次的著作物と見なされる。
たとえばフォーマット変換や機械翻訳（点字などへの置き換え）などは「逐語的コピー」と呼ばれ複製と見なされる[^cpy]。

[^ak]: {{< pdf-file title="「アンコウ行灯事件」判決文" link="http://www.courts.go.jp/app/files/hanrei_jp/827/013827_hanrei.pdf" >}}より。
[^cpy]: 他にも漫画のキャラクタなどを似せて描いた場合も，二次的著作物ではなく，複製とみなされる場合がある。文脈の問題とか，どの程度似てるかとかにもよるだろうけど。

具体的には

- 翻訳
- 編曲
- 変形（美術、写真、建築物、地図・図形の著作物で用いられることが多い）
- 脚色
- 映画化
- 翻案（上述した以外の全て。コミカライズやノベライズ，文章の要約，あるいはプログラムのバージョンアップや他言語への移植なども含まれる）

といった感じに分類されるが，一絡げに「翻案」または「改変」とすることが多い。
[CC Licenses] では二次的著作物に相当するものを「翻案物（Adapted Material）[^m]」としている。

[^m]: [CC Licenses] は著作隣接権で保護されるものも含む形で許諾しているためこのような表現になる。詳しくは「[CCライセンス・バージョン4.0 日本語版の公開 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2015/07/15/cc%E3%83%A9%E3%82%A4%E3%82%BB%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%90%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B34-0-%E6%97%A5%E6%9C%AC%E8%AA%9E%E7%89%88%E3%81%AE%E5%85%AC%E9%96%8B/)」を参照のこと。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="http://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
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

つまり二次的著作物を利用したいと考えるなら，二次的著作物の著作（権）者と原著作（権）者の双方から許可を得る必要がある。

## CC Licenses による改変の許諾

[CC Licenses] では「改変禁止 <i class="cc cc-nd"></i>」条件がなければ「翻案物を作成、複製および共有すること」を許諾する。
また翻案物を受け取ったユーザは原著作者の許諾を（提示されている [CC Licenses] の条件に従って）自動的に得る[^sl]。

[^sl]: [CC Licenses] ではサブライセンスを禁止しているため，翻案物に対する許諾を原著作者から得る，という形になっている。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="http://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>あなたから翻案物を受領した者は、あなたが適用した翻案者のライセンスの条件にしたがった翻案物におけるライセンスされた権利を行使できるという申出を自動的に許諾者から受け取ります。</q>
{{< /fig-quote >}}

つまり，原マテリアルとその翻案物の双方が [CC Licenses] で許諾されているなら，双方のライセンス条件に従って翻案物も利用することができる。

### 原著作者のクレジットも表示する必要あり

[CC Licenses] では「表示 <i class="cc cc-by"></i>」条件が必須になっているため，翻案物およびその複製や翻案物に対しても原著作者のクレジット表示が必要になる[^attr]。

[^attr]: クレジット表示については前に書いた「[Creative Commons Licenses]({{< relref "cc-licenses/02-creative-commons-licenses.md" >}})」を参照のこと。

## Copyleft のすすめ

二次的著作物については原著作物の許諾条件と二次的著作物の許諾条件の両方に従う必要がある。
原著作物が [CC Licenses] で許諾されている場合でも，その翻案物が [CC Licenses] で許諾されていない場合は利用条件が著しく制限される場合がある（まぁ原著作物から派生させればいいのだが）。
また原著作物とその翻案物の双方が [CC Licenses] で許諾されているとしても両者の条件が異なれば，やはりそれも制限になってしまう。

もし翻案物も含めてマテリアルを広く共有したいと望むのであれば「継承 <i class="cc cc-sa"></i>」条件を付加することをお勧めする。
「継承 <i class="cc cc-sa"></i>」条件が付加されている場合は，翻案物に対しても同等のライセンス[^cmp] を付加することが求められる。

[^cmp]: 同等のライセンスとして同じ条件の [CC Licenses] の他バージョンが挙げられる。また「[表示-継承]」については [Free Art License](http://artlibre.org/licence/lal/en/) 1.3 や [GNU GPL](https://www.gnu.org/copyleft/gpl.html)v3 も互換性のあるライセンスとして認められている。

このようなライセンスの仕組みは [Copyleft] と呼ばれている。
[Copyleft] の起源は [GNU GPL（GNU General Public License）](https://www.gnu.org/copyleft/gpl.html) であるが， [CC Licenses] においても「継承 <i class="cc cc-sa"></i>」条件よって [Copyleft] もしくはそれに近いライセンスを構成している[^cl]。

[^cl]: ただし厳密に [Copyleft] と言えるのは「[表示-継承]」条件のみである。

[Copyleft] はプログラミングの世界でも「ウイルス的」と揶揄されることもあるが，共有を維持するには必要な仕掛けだと思う。
[CC Licenses] の使われ方を見ても，2015年には「[表示-継承]」条件が37%で最も多い。
この機会に検討してみてはいかがだろうか。

## パロディについて


## 2次創作のみを許可したい場合



## 参考になる（かもしれない） Web ページ

- [クリエイティブ・コモンズ・ライセンスのブログ翻訳のススメ](http://www.yamdas.org/column/technique/cctrans.html)
- [Creative Commons — State of the Commons 2015 — It's been a remarkable year, most notably for the more than 1.1 billion works under one of the CC licenses, CC0, or the public domain mark.](https://stateof.creativecommons.org/2015/)

[本シリーズ]: /cc-licenses "改訂3版： CC-License について — text.Baldanders.info"
[著作権法]: http://law.e-gov.go.jp/htmldata/S45/S45HO048.html "著作権法"
[CC Licenses]: http://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"
[表示-継承]: http://creativecommons.org/licenses/by-sa/4.0/ "Creative Commons — Attribution-ShareAlike 4.0 International — CC BY-SA 4.0"
[Copyleft]: http://www.gnu.org/licenses/copyleft.html "コピーレフトって何? - GNUプロジェクト - フリーソフトウェアファウンデーション"
