+++
tags = ["drm", "creative-commons", "copyright", "license"]
draft = false
description = "CC Licenses v4 では利用者側の技術的保護手段回避を許諾するという形に大きく変わった。"
date = "2016-12-15T21:17:29+09:00"
update = "2016-12-16T09:23:23+09:00"
title = "CC Licenses における「技術的保護手段」の扱い"

[author]
  facebook = "spiegel.im.spiegel"
  url = "http://www.baldanders.info/spiegel/profile/"
  license = "by-sa"
  flattr = "spiegel"
  tumblr = "spiegel-im-spiegel"
  flickr = "spiegel"
  instagram = "spiegel_2007"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  avatar = "/images/avatar.jpg"
  twitter = "spiegel_2007"
  github = "spiegel-im-spiegel"
+++

1. [「技術的保護手段」とは]({{< relref "#abt" >}})
1. [技術的保護手段の回避]({{< relref "#avoid" >}})
1. [CC Licenses による技術的保護手段回避の許諾]({{< relref "#grant" >}})
    1. [「技術的保護手段回避の許諾」に至る経緯]({{< relref "#proc" >}})
1. [その他，雑多なこと]({{< relref "#other" >}})
    1. [DRM の変遷]({{< relref "#drm" >}})


## 「技術的保護手段」とは{#abt}

[著作権法]第2条では「技術的保護手段」は以下のように定義されている。
（あまりに長ったらしい文なので一部端折っている）

{{< fig-quote title="著作権法" link="http://law.e-gov.go.jp/htmldata/S45/S45HO048.html" >}}
<q>電子的方法、磁気的方法その他の人の知覚によつて認識することができない方法により、著作権等を侵害する行為の防止又は抑止をする手段であつて、著作物等の利用に際し、これに用いられる機器が特定の反応をする信号を著作物等に係る音若しくは影像とともに記録媒体に記録し、若しくは送信する方式又は当該機器が特定の変換を必要とするよう著作物等に係る音若しくは影像を変換して記録媒体に記録し、若しくは送信する方式によるものをいう。</q>
{{< /fig-quote >}}

なげーよ（笑） 要するに「著作物等[^m1]」を何らかの形で変換し，特定の手段を用いなければ復元できないような方式を指す。
通常は「変換」や「復元」には暗号技術が用いられる。
ぶっちゃけて言うなら「技術的保護手段」とは「著作物等を暗号化する手段」と思っていただいて構わない。

[^m1]: [著作権法]第2条では「著作物、実演、レコード、放送又は有線放送」としている。ちなみに [CC Licenses] では同様のものを「マテリアル（material）」と定義している。これについては「[Creative Commons Licenses]({{< relref "cc-licenses/02-creative-commons-licenses.md" >}})」を参照のこと。

具体的には，「著作物等」を変換（暗号化）するのは著作（権）者またはコンテンツ・ホルダー（contents holder）で，彼等が許可した（つまり正規ルートでコンテンツを購入した）ユーザのみが専用の再生装置を使って著作物等を復元（復号）し使用できる，という仕組みだ。

## 技術的保護手段の回避{#avoid}

この方式の問題のひとつはユーザは「著作物等」の「利用」のみならず[「公正な利用（fair use）」や「著作権の制限」]({{< relref "cc-licenses/01-copyright.md#fair-use" >}})更には（著作権の範囲外である）「使用」まで制限してしまうことである。

たとえば DVD や BD は対応する再生装置でなければ再生（つまり復元して使用）することができないが，市場から再生装置が無くなれば DVD や BD などベランダに吊るしてカラス除けにするくらいにしか役に立たない（実際そうやって消えていく機器は山ほどある）。

もっと深刻な問題がある。
いささか逆説的かもしれないが。

先ほど述べたように技術的保護手段には暗号技術が用いられる。
ユーザが「著作物等」を復元（復号）するためには復号鍵が必要だが，実際には鍵の生成方法が杜撰だったり（特にスタンドアロンの再生装置では）鍵の配送手段が貧弱または存在しないことが多い。
したがって再生装置にある復号鍵を全て解読すればその鍵情報を使って再生装置なしでも「著作物等」を復元できることになる（実際にはこんなに簡単ではないが）。
こういった行為を「技術的保護手段の回避」と呼ぶ。

「技術的保護手段の回避」は（仕様や実装によるが）技術的にものすごく難しいというわけではない。
しかも「技術的保護手段の回避」は非常にカジュアルな動機で行われる。
何故なら，上述したように技術的保護手段は「[「公正な利用」や「著作権の制限」]({{< relref "cc-licenses/01-copyright.md#fair-use" >}})更には「使用」まで制限してしまう」からだ。

これに対して著作権システムを守りたい側はあまり賢明でない方法をとった。
[著作権法]では「私的使用のための複製（第30条）」を「[著作権の制限]({{< relref "cc-licenses/01-copyright.md#fair-use" >}})」として認めているが，「技術的保護手段の回避」を除外することにより事実上禁止してしまったのだ。

たしかに多くのユーザはこれを律儀に守って無理に技術的保護手段の回避を行うことはないだろうが，一部の悪質なユーザには（罰則を強めにしたところで）抑止効果はない。
一方で技術的保護手段の仕様や実装に欠陥があったとしても，これを報告する動機は抑止される。
結局「正直者が馬鹿を見る」ことになるのだ。

## CC Licenses による技術的保護手段回避の許諾{#grant}

[CC Licenses] では技術的保護手段を以下のように定義している。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q><strong>「効果的な技術的保護手段」</strong>とは、1996年12月20日に採択されたWIPO著作権条約第11条、および／または類似の国際協定の義務を満たす諸法規の下で、正当な権限なしに回避されてはならないものとされる諸手段をいいます。</q>
{{< /fig-quote >}}

簡単に書かれてるが，[著作権法]の定義と概ね同じだと思っていい。
その上で，利用条件に従う限り，技術的保護手段の回避を許諾している。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>許諾者は、あなたに対し、あらゆる媒体や形式（現在知られているか、または今後作られるか否かを問いません）において、ライセンスされた権利を行使する権限、およびその行使に必要とされる技術的な改変を行う権限を付与します。許諾者は、あなたが、ライセンスされた権利を行使するために必要とされる技術的な改変（効果的な技術的保護手段を回避するために必要とされる技術的な改変を含みます）を禁止するいかなる権利または権限を放棄し、および／またはこれらの権利または権限を行使しないことに同意します。</q>
{{< /fig-quote >}}

さらに [CC Licenses] では技術的保護手段の回避を翻案と見なさないことで「改変禁止 <i class="cc cc-nd"></i>」条件でも技術的保護手段の回避を許諾している。

またこれとは別に，下流側（downstream）へ再配布を行う場合は技術的保護手段を適用してはならないともある。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>あなたは、ライセンス対象物の受領者がライセンスされた権利を行使するのを制限されることになる場合には、ライセンス対象物に対して、いかなる追加条項または異なる条項も提案または課してはならず、あるいは、いかなる効果的な技術的保護手段も適用してはなりません。</q>
{{< /fig-quote >}}

[CC Licenses] 下のマテリアルを再配布する場合には注意が必要である[^sl]。

[^sl]: [CC Licenses] ではライセンスの再許諾（sublicense）を許可していない。詳しくは「[Creative Commons Licenses]({{< relref "cc-licenses/02-creative-commons-licenses.md" >}})」を参照のこと。

### 「技術的保護手段回避の許諾」に至る経緯{#proc}

何度も言うように技術的保護手段には「[「公正な利用」や「著作権の制限」]({{< relref "cc-licenses/01-copyright.md#fair-use" >}})更には「使用」まで制限してしまう」問題があり，さらに言えば [CC Licenses] で許諾している利用も阻まれてしまうため [Creative Commons] としてはこれを看過するわけにはいかなかった。
しかし [CC Licenses] v2.x までは技術的保護手段について消極的な文言しかなかった。

{{< fig-quote title="クリエイティブ・コモンズ リーガル・コード — 表示-継承 2.1 日本版" link="https://creativecommons.org/licenses/by-sa/2.1/jp/legalcode" >}}
<q>あなたは、この利用許諾条項と矛盾する方法で本著作物へのアクセス又は使用をコントロールするような技術的保護手段を用いて、本作品又はその二次的著作物を利用してはならない。</q>
{{< /fig-quote >}}

[CC Licenses] v3 作成時には「もっと強い文言にすべきではないか」という意見もあったが，結局は

{{< fig-quote title="Creative Commons Legal Code — Attribution-ShareAlike 3.0 Unported" link="https://creativecommons.org/licenses/by-sa/3.0/legalcode" lang="en" >}}
<q>When You Distribute or Publicly Perform the Work, You may not impose any effective technological measures on the Work that restrict the ability of a recipient of the Work from You to exercise the rights granted to that recipient under the terms of the License.</q>
{{< /fig-quote >}}

{{< fig-quote title="【CCPLv3.0】ＤＲＭ条項の改正に関する議論 | クリエイティブ・コモンズ・ジャパン" link="https://creativecommons.jp/2006/11/15/discussion-drm/" >}}
<q>あなたがこの作品を頒布、譲渡、公衆送信その他の方法で公衆に提供するにあたっては、本作品の受領者がこのライセンスの中で与えられている権利の行使を制限されるような技術的手段を本作品に適用してはならない。</q>
{{< /fig-quote >}}

というレベルにとどまったようである（v4 の「ダウンストリーム（下流側）の受領者」の文言とほぼ同じ）。

これが v4 では利用者側の技術的保護手段回避を許諾するという形で更に踏み込んでいる[^g1]。
「技術的保護手段の回避」は著作（権）者によって許諾できる，というのがポイントだろう。
「技術的保護手段の回避」は複製権（[著作権法] 第21条）の侵害と見なしうるので，「技術的保護手段の回避」を含めて複製を（ [CC Licenses] が要求する条件下で）認めてしまえばいいのである。

[^g1]: この辺の経緯について書かれているまとまった文書が見つからないのだが，おそらく他の「自由なライセンス」との互換性を考えた結果ではないかと推測する。

## その他，雑多なこと{#other}

### DRM の変遷{#drm}

これまで述べてきたように，いわゆる技術的保護手段が技術的に不十分でユーザの使用を不当に制限するものであることは明らかである。
これはコンテンツ・ホルダーの側も認識しているようで， DRM (Digital Rights Management) の形も変わってきた。

ひとつは定額制のストリーミング・サービスの台頭である。
ストリーミング・サービスの利点は，コンテンツ・ホルダーもユーザも「違法コピー」を考えなくていいことである。
ユーザはその場でコンテンツを「消費」するだけ。
定額制で何時でも何度でも見れるなら「コピっておく」必要がない。

欠点はコンテンツ・ホルダーやサービス・プロバイダの力が大きくなりすぎることで，両者が癒着するとかなり酷いことになる。
実際，ストリーミング・サービスでの売り上げがクリエイターやアーティストに公平に分配されないなどの話もチラホラある。

もうひとつは DRM が「監視型」に移行したことである。
「電子透かし」や「電子指紋」を使ってネット上に流通するコンテンツを比較的容易に監視できるようになった。
これがうまく機能すれば一般ユーザのネット上での活動を妨げることなく悪質なものだけを排除することができる。

- [コンテンツの識別で流通の価値をお金に変える](http://techon.nikkeibp.co.jp/article/NEWS/20080812/156366/ "コピーに自由を ―生まれ変わるDRM―（第3回） - 日経テクノロジーオンライン")

しかし，一方でこれもコンテンツ・ホルダーやサービス・プロバイダの力が大きくなりすぎる傾向がある。
コンテンツ・ホルダーやサービス・プロバイダはボットなどを使って機械的に監視と排除を行うが，判断を間違えたり「[公正な利用]({{< relref "cc-licenses/01-copyright.md#fair-use" >}})」まで排除される例があるようだ。

このような感じで技術的保護手段が時代遅れになる一方で新しい形の DRM に対する問題も出てきている。
DRM についてはこれからも注視していく必要がある。

## 参考になる（かもしれない） Web ページ

- [【CCPLv3.0】ＤＲＭ条項の改正に関する議論 | クリエイティブ・コモンズ・ジャパン](https://creativecommons.jp/2006/11/15/discussion-drm/)
- [コピーに自由を ―生まれ変わるDRM―（目次） - 日経テクノロジーオンライン](http://techon.nikkeibp.co.jp/article/NEWS/20080812/156363/)
- [「オープンソースDRM」の不可能性について | OSDN Magazine](https://mag.osdn.jp/08/10/26/2237236)
- [小寺信良：「補償金もDRMも必要ない」――音楽家 平沢進氏の提言 (1/4) - ITmedia LifeStyle](http://www.itmedia.co.jp/lifestyle/articles/0606/12/news005.html)
- [Problem on anti-circumvention provision in Copyright protection プロテクト破り規制法案の問題点](http://www-vacia.media.is.tohoku.ac.jp/member/o/s-yamane2004/articles/crypto/circumvention.html)
- [DRM に関する覚え書き — Baldanders.info](http://www.baldanders.info/spiegel/log2/000295.shtml)
- [監視をコントロールする — Baldanders.info](http://www.baldanders.info/spiegel/log2/000490.shtml)
- [『デジタル音楽の行方』から10年経って - WirelessWire News（ワイヤレスワイヤーニュース）](https://wirelesswire.jp/2015/06/32173/)
- [テイラー効果広がる―プリンスもストリーミング条件に反発してSpotify他から楽曲を引き上げ | TechCrunch Japan](http://jp.techcrunch.com/2015/07/03/20150702prince-wont-stream-4-u/)
- [YouTubeが著作者をあべこべに判断　音源を無断使用された音楽家の動画が再生できなくなるトラブルが発生 - ねとらぼ](http://nlab.itmedia.co.jp/nl/articles/1609/02/news120.html)
- [「フェアユースでも使用料を払え」というソニーミュージックの横暴と、それを許すYouTubeのコンテンツID – P2Pとかその辺のお話R](http://p2ptk.org/copyright/350)
- [Fox、番組にYouTube動画を無断使用した挙句に元動画を権利者削除 – P2Pとかその辺のお話R](http://p2ptk.org/copyright/346)

[本シリーズ]: /cc-licenses "改訂3版： CC-License について — text.Baldanders.info"
[著作権法]: http://law.e-gov.go.jp/htmldata/S45/S45HO048.html "著作権法"
[Creative Commons]: https://creativecommons.org/ "Creative Commons"
[CC Licenses]: https://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"
[表示-継承]: https://creativecommons.org/licenses/by-sa/4.0/ "Creative Commons — Attribution-ShareAlike 4.0 International — CC BY-SA 4.0"
