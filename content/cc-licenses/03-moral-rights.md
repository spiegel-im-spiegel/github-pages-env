+++
date = "2016-01-20T21:31:46+09:00"
update = "2018-11-13T13:34:22+09:00"
description = "前回紹介した「Creative Commons Licenses」の内容を踏まえた上で今回は「著作者人格権」と CC Licenses との関係について解説する。"
draft = false
tags = ["moral-rights", "creative-commons", "copyright", "license"]
title = "人格権と CC Licenses"

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
  url = "http://www.baldanders.info/spiegel/profile/"
  
[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回紹介した「Creative Commons Licenses」]({{< ref "/cc-licenses/02-creative-commons-licenses.md" >}})の内容を踏まえた上で今回は人格権と [CC Licenses] との関係について解説する。

1. [「著作者人格権」とは]({{< relref "#about" >}})
1. [その他の「人格権」]({{< relref "#other" >}})
1. [CC Licenses では（可能なかぎり）人格権は行使されない]({{< relref "#ccl" >}})

## 「著作者人格権」とは{#about}

まずは「著作者人格権」から。

著作者人格権は[著作権法] 第18条から第20条にかけて定められている。
すなわち以下の3つの権利の総称である。

- 公表権（[著作権法] 第18条）
- 氏名表示権（[著作権法] 第19条）
- 同一性保持権（[著作権法] 第20条）

また「著作者人格権は、著作者の一身に専属し、譲渡することができない」（[著作権法] 第59条）。
つまり著作者以外の人に譲渡できないし，著作者以外の人がこれを行使することもできない[^a]。
また著作者人格権は著作者の死後も機能するとされている[^a2]。

[^a]: 近年はあまり見なくなったが， Web サービスの中にはユーザのコンテンツの著作権をサービス・プロバイダに譲渡し著作者人格権の不行使を要求するものがある。著作者人格権は譲渡できないため，著作者人格権の不行使も無効になると思われる。ただし個別契約の中で著作者人格権の不行使を約束する場合はあり得る。ちなみに「職務著作（[著作権法] 第15条）」の場合は所属する法人が著作者となるため，個々人に著作者人格権は発生しない。
[^a2]: 「著作物を公衆に提供し、又は提示する者は、その著作物の著作者が存しなくなつた後においても、著作者が存しているとしたならばその著作者人格権の侵害となるべき行為をしてはならない」（[著作権法] 第60条）。

このうちもっとも強力な権利が「同一性保持権」である。

{{< fig-quote title="著作権法" link="http://elaws.e-gov.go.jp/search/elawsSearch/elaws_search/lsg0500/detail?lawId=345AC0000000048" >}}
<q>著作者は、その著作物及びその題号の同一性を保持する権利を有し、その意に反してこれらの変更、切除その他の改変を受けないものとする。</q>
{{< /fig-quote >}}

とあるように著作者は著作物の改変をいつでも禁止することができる[^b]。
このような強い権利は他の国にはない（同一性保持権自体は他の国でも見られるが限定的）。
このため日本の著作権ライセンスの中には，許諾者による「同一性保持権」の不行使を明示しているものもある。

[^b]: ただしこれにはいくつか例外がある。学校教育の目的上やむを得ない場合，建築物の模様替えやリフォーム，プログラムの不具合修正や改善，などである。ちなみに「著作権の制限」（[著作権法] 第30-50条）は著作者人格権には及ばないのでご注意を。

## その他の「人格権」{#other}

著作隣接権のひとつである「実演家の権利」には「実演家の人格権」というのがある。
具体的には

- 氏名表示権（[著作権法] 第90条）
- 同一性保持権（[著作権法] 第90条）

の2つである。
著作者人格権と同じく一身専属性を持ち譲渡できないし，実演家の死後も機能する（[著作権法] 第101条）。
ただし同一性保持権については

{{< fig-quote title="著作権法" link="http://elaws.e-gov.go.jp/search/elawsSearch/elaws_search/lsg0500/detail?lawId=345AC0000000048" >}}
<q>実演家は、その実演の同一性を保持する権利を有し、自己の名誉又は声望を害するその実演の変更、切除その他の改変を受けないものとする。</q>
{{< /fig-quote >}}

とあるとおり「自己の名誉又は声望を害する」場合のみという制限がある。

著作権ではないが「パブリシティ権」や「プライバシー権」といったものも人格権の一種として挙げられるだろう[^c]。
著作権法上は問題なくとも，これらの人格権によって著作物が利用できない場合はある，ということは覚えておいたほうがいいと思う。

[^c]: ただし日本には明示的な「プライバシー権」は存在しない。「個人情報保護法」は（人格権ではなく）個人にまつわる**情報**の取り扱いについて定めた法律である。

## CC Licenses では（可能なかぎり）人格権は行使されない{#ccl}

さて，上述した人格権について [CC Licenses] はどうしているかというと，許諾範囲を逸脱しない限り許諾者は人格権を行使しないことになっている。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>同一性保持の権利のような著作者人格権は、本パブリック・ライセンスのもとではライセンスされません。パブリシティ権、プライバシー権、および／または他の類似した人格権も同様です。ただし、可能なかぎり、許諾者は、あなたがライセンスされた権利を行使するために必要とされる範囲内で、また、その範囲内でのみ、許諾者の保持する、いかなるそのような権利を放棄し、および／または主張しないことに同意します。</q>
{{< /fig-quote >}}

人格権の一身専属性を考えればこのような表現になるということだろう。
ただし「可能なかぎり」という条件がついている。

この点について，以前のバージョンでは比較的明確に書かれていて，たとえば 2.1 日本版では

{{< fig-quote title="クリエイティブ・コモンズ リーガル・コード — 表示-継承 2.1 日本版" link="https://creativecommons.org/licenses/by-sa/2.1/jp/legalcode" >}}
<q>原著作者及び実演家の名誉又は声望を害する方法で原著作物を改作、変形もしくは翻案して生じる著作物は、この利用許諾の目的においては、二次的著作物に含まれない。</q>
{{< /fig-quote >}}

とあり，原著作者の名誉・声望を害する改変は二次的著作物に含めない，としている。
また 3.0 Unported でも同様に

{{< fig-quote title="Creative Commons Legal Code — Attribution-ShareAlike 3.0 Unported" link="https://creativecommons.org/licenses/by-sa/3.0/legalcode" lang="en" >}}
<q>Except as otherwise agreed in writing by the Licensor or as may be otherwise permitted by applicable law, if You Reproduce, Distribute or Publicly Perform the Work either by itself or as part of any Adaptations or Collections, You must not distort, mutilate, modify or take other derogatory action in relation to the Work which would be prejudicial to the Original Author's honor or reputation.</q>
{{< /fig-quote >}}

となっている[^c2]。

[^c2]: [CC Licenses] 3.0 では著作者人格権に関する議論が国際的にも行われたようだ。くわしくは「[著作者人格権（同一性保持権）に関する議論](http://creativecommons.jp/2006/11/15/ccplv3-discussion/)」が参考になる。

この辺の文言が 4.0 International でなくなった理由はよく分からないが，“[What's New in 4.0](https://creativecommons.org/Version4/)” には

{{< fig-quote title="What's New in 4.0 - Creative Commons" link="https://creativecommons.org/Version4/" lang="en" >}}
<q>The 4.0 license suite uniformly and explicitly waives moral rights held by the licensor where possible to the limited extent necessary to enable reuse of the content in the manner intended by the license. Publicity, privacy, and personality rights held by the licensor are expressly waived to the same limited extent. While many understand these rights to be waived when held by the licensor in 3.0 and earlier versions, version 4.0’s treatment makes the intended outcome clear.</q>
{{< /fig-quote >}}

との記述があり，あえて「名誉・声望を害する改変」部分の記述を削った可能性もある[^d]。

[^d]: この辺は GPL などの他ライセンスとの互換性を考える上で重要なポイントでもある。

### キャラクタの権利

キャラクタやキャラクタの名前の利用については著作権ではなく商標権（工業デザインの場合は意匠権）で保護されることが多い[^e]。
この場合， [CC Licenses] は商標権や意匠権についてはライセンスしないので，個別に許可を得る必要がある。

 [^e]: 基本的にキャラクタや名前には著作権はない。このため防衛のためにキャラクタやキャラクタの名前を商標登録することが多いらしい。ただし創作上の文脈としてキャラクタや名前が書（描）かれている場合は著作物の一部として認められる場合がある。

実在の人物やその人物の延長上のキャラクタ（「デーモン小暮閣下」など）に対しては「パブリシティ権」が適用される。
パブリシティ権は肖像権の一種と考えられ「氏名・肖像から生じる経済的利益ないし価値を排他的に支配する権利[^ds]」と定義されている[^pr1]。
[CC Licenses] では，これまで述べた通り，パブリシティ権についてもライセンスしないが「可能なかぎり」行使しない，となっている。

[^ds]: {{< pdf-file title="「ダービースタリオン事件」判決文" link="http://www.courts.go.jp/hanrei/pdf/A730EBEA9CA60D6249256C7F0023A16E.pdf" >}}より。
[^pr1]: ただし作品上の架空のキャラクタや無機物やペット等にはパブリシティ権は適用されない。

## 参考になる（かもしれない） Web ページ

- [【CCPLv3.0】著作者人格権（同一性保持権）に関する議論 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2006/11/15/ccplv3-discussion/)
- [What's New in 4.0 - Creative Commons - Creative Commons](https://creativecommons.org/Version4/) （[日本語参考訳](http://qiita.com/nyampire/items/c03904bd27bd8812aad3)）
- [【ネット著作権】 人名・グループ名を作品タイトルに使ってはいけない？　～水曜日のカンパネラ「ヒカシュー」騒動と疑似著作権～ - INTERNET Watch](http://internet.watch.impress.co.jp/docs/special/fukui/20160517_757708.html)

[本シリーズ]: /cc-licenses "改訂3版： CC-License について — text.Baldanders.info"
[著作権法]: http://elaws.e-gov.go.jp/search/elawsSearch/elaws_search/lsg0500/detail?lawId=345AC0000000048 "著作権法"
[CC Licenses]: https://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"
