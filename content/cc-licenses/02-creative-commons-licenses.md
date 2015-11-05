+++
date = "2015-11-05T21:00:25+09:00"
description = "description"
draft = true
tags = ["creative-commons", "copyright", "license"]
title = "Creative Commons Licenses"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

[前回]は著作権についてざくっと説明したので，今回はいよいよ Creative Commons Licenses （以降 “[CC Licenses]” と略称する）について解説する。

## CC Licenses の対象となるもの

[CC Licenses] は著作権およびその類似の権利で保護されるあらゆるものが対象になる。
たとえば[著作権法]で保護される以下の著作物も対象になる（[著作権法] 第10条から12条）。

- 言語の著作物
- 音楽の著作物
- 舞踊又は無言劇の著作物
- 美術の著作物
- 建築の著作物
- 地図又は図形の著作物
- 映画の著作物
- 写真の著作物
- プログラムの著作物
- 二次的著作物
- 編集著作物
- データベースの著作物

また[著作権法]で定められた以下の「著作隣接権」で保護されるものも対象である。

- 実演家の権利
- レコード製作者の権利
- 放送事業者の権利
- 有線放送事業者の権利

ただし [CC Licenses] では他の知的財産権（特許権や商標権など）は対象外なため，これらの権利で保護されるものには法的効力がない（または限定的）。

[CC Licenses] ではこれらライセンスの対象になるものを「ライセンス対象物（licensed material）」あるいは「マテリアル（material）」と呼んでいる[^a]。
[本シリーズ]でも「マテリアル」で統一することにする。

[^a]: こう呼ぶようになった経緯については「[CCライセンス・バージョン4.0 日本語版の公開 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2015/07/15/cc%E3%83%A9%E3%82%A4%E3%82%BB%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%90%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B34-0-%E6%97%A5%E6%9C%AC%E8%AA%9E%E7%89%88%E3%81%AE%E5%85%AC%E9%96%8B/)」を参照のこと

## CC Licenses の利用条件

CC Licenses は，大雑把に「ライセンス利用条件（license conditions）」と条件に対する「ライセンス付与（license grant）」で構成されている。
利用条件はライセンスの「許諾者（licensor）」が指定する。
当然ながら許諾者（達）はマテリアルに対して著作権（またはその類似の権利）を持っていなければならない。

CC Licenses の利用条件は以下の4つの条件の組み合わせである。

{{< fig-gen >}}
<table>
<tbody>
<tr><th class='left'><i class="cc cc-by cc-2x"></i></th><th class='left'>表示</th><td>マテリアルのクレジットを表示すること</td></tr>
<tr><th class='left'><i class="cc cc-sa cc-2x"></i></th><th class='left'>継承</th><td>改変したマテリアルについて，元のマテリアルのライセンスと同等のライセンスで公開すること</td></tr>
<tr><th class='left'><i class="cc cc-nc cc-2x"></i>
                     <i class="cc cc-nc-eu cc-2x"></i>
                     <i class="cc cc-nc-jp cc-2x"></i></th><th class='left'>非営利</th><td>営利目的での利用をしないこと</td></tr>
<tr><th class='left'><i class="cc cc-nd cc-2x"></i></th><th class='left'>改変禁止</th><td>元のマテリアルを改変しないこと</td></tr>
</tbody>
</table>
{{< /fig-gen >}}

実際には「表示 <i class="cc cc-by"></i>」は必須の条件[^b] になっていて，さらに「継承 <i class="cc cc-sa"></i>」と「改変禁止 <i class="cc cc-nd"></i>」は条件が衝突するため同時に指定できない。
結局以下に示す6つの組み合わせのみが有効になる。

[^b]: バージョン1では「表示 <i class="cc cc-by"></i>」は任意だったが，実際には殆どの人が「表示 <i class="cc cc-by"></i>」条件を指定していたため，（ライセンス選択の煩雑さを防ぐ意味もあり）バージョン2以降は必須になった。なお共有されるマテリアル（改変されたものを含む）に対してクレジット表示を削除するよう要請することもできる。

{{< fig-gen >}}
<table>
<tbody>
<tr><th class='left'><i class="cc cc-BY cc-2x"></i>      </th><td>表示</td></tr>
<tr><th class='left'><i class="cc cc-by-sa cc-2x"></i>   </th><td>表示-継承</td></tr>
<tr><th class='left'><i class="cc cc-by-nc cc-2x"></i>   </th><td>表示-非営利</td></tr>
<tr><th class='left'><i class="cc cc-by-nc-sa cc-2x"></i></th><td>表示-非営利-継承</td></tr>
<tr><th class='left'><i class="cc cc-by-nd cc-2x"></i>   </th><td>表示-改変禁止</td></tr>
<tr><th class='left'><i class="cc cc-by-nc-nd cc-2x"></i></th><td>表示-非営利-改変禁止</td></tr>
</tbody>
</table>
{{< /fig-gen >}}

## CC Licenses が許諾するもの

マテリアルを利用したいユーザは，ライセンスで指定された上述の条件の組み合わせを守る限り，以下の利用が許諾される。

- **対象となるマテリアルの全部または一部を複製および共有すること**
    - マテリアルを複製する（複製）
        - 著作隣接権に係るマテリアルの場合は録音・録画を含む
    - マテリアルを公表・配布する（上演，演奏，上映，公衆送信等，口述，展示，頒布（映画の著作物のみ），譲渡・貸与（映画の著作物以外））
        - 著作隣接権に係るマテリアルの場合は（有線を含む）放送・再放送，送信可能化，伝達を含む
- **対象となるマテリアルの改変物を作成，複製および共有すること** （「改変禁止 <i class="cc cc-nd"></i>」条件がない場合）
    - マテリアルを改変する（翻訳・翻案等）
    - マテリアルを改変したものを公表・配布する（二次的著作物の利用に関する原著作者の権利の不行使）

許諾は全世界で適用される。
また，無償（royalty-free）かつ再許諾不可（non-sublicensable）かつ非排他的（non-exclusive）に許諾される。

許諾者はマテリアルに付与したライセンスを取り消しできない。
したがって，少なくともライセンスの条件を守っていればマテリアルの利用を取り消されることはない[^c]。
またライセンスの効力は著作権の存続期間で持続する（著作権の存続期間が過ぎれば「パブリックドメイン（public domain）」になる）。

[^c]: つまり許諾者は，ライセンスを付与する場合には妥当な条件であるかあらかじめ考慮する必要がある。

なお以下の行為は条件によらず許可されない。

- マテリアルおよびその複製について，クレジット表示を改変・削除すること
- マテリアルおよびその複製について，ライセンスおよびその条件を変更すること
- マテリアルの正当な使用または許可された利用を制限する「技術的保護手段」を用いること[^d]

[^d]: 厳密にはマテリアルが「技術的保護手段」によって不当に保護されている場合は，ユーザは「技術的保護手段」を迂回することができる（許諾者は迂回を禁止する権利を放棄する）。「技術的保護手段」については別の記事で解説する。

## CC Licenses の3つのレイヤ

未稿。

## CC Licenses のバージョン

未稿。

### CC Licenses と互換性のあるライセンス

未稿。

## 全ての権利を放棄（または不行使）する CC0

未稿。

## Free Culture Licenses

未稿。

[本シリーズ]: /cc-licenses "改訂3版： CC-License について — text.Baldanders.info"
[著作権法]: http://law.e-gov.go.jp/htmldata/S45/S45HO048.html "著作権法"
[前回]: {{< relref "cc-licenses/01-copyright.md" >}} "著作権と著作権法 — 改訂3版： CC Licenses について"
[Creative Commons]: http://creativecommons.org/ "Creative Commons"
[CC Licenses]: http://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"