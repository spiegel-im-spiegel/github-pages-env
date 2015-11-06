+++
date = "2015-11-06T17:00:00+09:00"
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
- プログラムの著作物[^0]
- 二次的著作物
- 編集著作物
- データベースの著作物

[^0]: 「プログラムの著作物」については [CC Licenses] は向いてないかもしれない。「プログラムの著作物」については GNU GPL などの FOSS 用のライセンスがおすすめである。

また[著作権法]で定められた以下の「著作隣接権」で保護されるものもライセンスの対象にできる。

- 実演家による実演[^1]
- レコード製作者により製作されたレコード[^2]
- 放送事業者による放送[^3]
- 有線放送事業者による有線放送[^4]

[^1]: [著作権法]では実演を「演劇的に演じ、舞い、演奏し、歌い、口演し、朗詠し、又はその他の方法により演ずること」，実演家を「俳優、舞踊家、演奏家、歌手その他実演を行う者及び実演を指揮し、又は演出する者」と定義している（第2条）。
[^2]: [著作権法]ではレコードを「蓄音機用音盤、録音テープその他の物に音を固定したもの（音を専ら影像とともに再生することを目的とするものを除く）」，レコード製作者を「レコードに固定されている音を最初に固定した者」と定義している（第2条）。 DVD 映像などは「映画の著作物」に含まれるらしく，レコードには含まれない。
[^3]: [著作権法]では放送を「公衆送信のうち、公衆によつて同一の内容の送信が同時に受信されることを目的として行う無線通信の送信」，放送事業者を「放送を業として行う者」と定義している（第2条）。ちなみにインターネット通信は放送・有線放送ではなく「自動公衆送信」に分類される。
[^4]: [著作権法]では有線放送を「公衆送信のうち、公衆によつて同一の内容の送信が同時に受信されることを目的として行う有線電気通信の送信」，有線放送事業者を「有線放送を業として行う者」と定義している（第2条）。

ただし [CC Licenses] では他の知的財産権（特許権や商標権など）は対象外なため，これらの権利で保護されるものには効力が及ばない（または限定的）。

[CC Licenses] ではこれらライセンスの対象になるものを「ライセンス対象物（licensed material）」あるいは「マテリアル（material）」と呼んでいる[^a]。
[本シリーズ]でも「マテリアル」で統一することにする。

[^a]: こう呼ぶようになった経緯については「[CCライセンス・バージョン4.0 日本語版の公開 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2015/07/15/cc%E3%83%A9%E3%82%A4%E3%82%BB%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%90%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B34-0-%E6%97%A5%E6%9C%AC%E8%AA%9E%E7%89%88%E3%81%AE%E5%85%AC%E9%96%8B/)」を参照のこと

## CC Licenses の利用条件

大雑把に言うと [CC Licenses] は「ライセンス利用条件（license conditions）」と条件に対する「ライセンス付与（license grant）」で構成されている。
利用条件はライセンスの「許諾者（licensor）」が指定する。
当然ながら許諾者（達）はマテリアルに対して著作権（またはその類似の権利）を持っていなければならない[^aa]。

[^aa]: ひとつのマテリアルに複数の著作（権）者が関わっている場合，マテリアルに  [CC Licenses] を付与するためには全ての権利者の間で合意が必要である。

[CC Licenses] の利用条件は以下の4つの条件の組み合わせである。

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
        - 著作隣接権に係るマテリアルの場合は（有線放送を含む）放送と再放送，伝達，送信可能化を含む
- **対象となるマテリアルの二次的著作物[^5] を作成，複製および共有すること** （「改変禁止 <i class="cc cc-nd"></i>」条件がない場合）
    - マテリアルを改変する（翻訳・翻案等）
    - マテリアルを改変したものを公表・配布する（二次的著作物の利用に関する原著作者の権利の不行使）

[^5]: [著作権法]では二次的著作物を「著作物を翻訳し、編曲し、若しくは変形し、又は脚色し、映画化し、その他翻案することにより創作した著作物」と定義している（第2条）。別の記事で説明することになると思うが，二次的著作物については「創作的」であることが条件。フォーマット変換や機械翻訳（点字などへの置き換え）などは「逐語的コピー」と呼ばれ複製と見なされる。

なお [CC Licenses] は「公正な利用（fair use）」や「著作権の制限」として利用が認められているものに対しては効力が及ばない。
「公正な利用」や「著作権の制限」で認められる利用については条件に関係なく可能である。

また，以下の行為は条件によらず許可されない。

- マテリアルおよびその複製について，クレジット表示を改変・削除すること
- マテリアルおよびその複製について，ライセンスおよびその条件を変更すること
- マテリアルの正当な使用または許可された利用を制限する「技術的保護手段」を用いること[^d]

[^d]: 厳密にはマテリアルが「技術的保護手段」によって不当に保護されている場合は，ユーザは「技術的保護手段」を迂回することができる（許諾者は迂回を禁止する権利を放棄する）。「技術的保護手段」については別の記事で解説する。

許諾は全世界で適用される。
また，無償（royalty-free）かつ再許諾不可（non-sublicensable）かつ非排他的（non-exclusive）に許諾される。

許諾者はマテリアルに付与したライセンスを取り消しできない。
したがって，少なくともライセンスの条件を守っていればマテリアルの利用を後から取り消されることはない[^c]。
またライセンスの効力は著作権の存続期間で持続する（著作権の存続期間が過ぎれば「パブリックドメイン（public domain）」になる）。

[^c]: つまり許諾者は，ライセンスを付与する場合には妥当な条件であるか予め考慮する必要がある。

## CC Licenses の3つのレイヤ

[CC Licenses] は以下の3つのレイヤで構成されている。

- **コモンズ証（commons deed）** [CC Licenses] の利用条件と許諾について Human-Readable な形で簡潔に記述したもの。
- **法的条項（legal code）** [CC Licenses] の内容を Lawyer-Readable な形で記述した「利用許諾書」。 法的条項によって [CC Licenses] が法的拘束力を持つことを担保する。
- **メタデータ（metadata）** マテリアルに関する情報を Machine-Readable な形で記述したもの。記述形式としては XML, RDFa, XMP などがあるが， Web 上では RDFa がよく用いられる。メタデータは Google などの検索サービスで解釈され，ネット上の数多のマテリアルを利用条件によって検索可能とする。

マテリアルに対して [CC Licenses] を指示する場合は，通常はコモンズ証へのリンクを示せばよい（コモンズ証のページから法的条項へリンクが張られている）。
各レイヤのリンクを以下に示す。

{{< fig-gen title="Creative Commons Licenses Version 4.0 International" >}}
<table>
<tbody>
<tr><th class='left'><i class="cc cc-BY cc-2x"></i>      </th><th class='left'>表示                </th><td><a href="http://creativecommons.org/licenses/by/4.0/">コモンズ証</a>（<a href="http://creativecommons.org/licenses/by/4.0/deed.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by/4.0/legalcode">法的条項</a>（<a href="http://creativecommons.org/licenses/by/4.0/legalcode.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by/4.0/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="cc cc-by-sa cc-2x"></i>   </th><th class='left'>表示-継承           </th><td><a href="http://creativecommons.org/licenses/by-sa/4.0/">コモンズ証</a>（<a href="http://creativecommons.org/licenses/by-sa/4.0/deed.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by-sa/4.0/legalcode">法的条項</a>（<a href="http://creativecommons.org/licenses/by-sa/4.0/legalcode.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by-sa/4.0/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="cc cc-by-nc cc-2x"></i>   </th><th class='left'>表示-非営利         </th><td><a href="http://creativecommons.org/licenses/by-nc/4.0/">コモンズ証</a>（<a href="http://creativecommons.org/licenses/by-nc/4.0/deed.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by-nc/4.0/legalcode">法的条項</a>（<a href="http://creativecommons.org/licenses/by-nc/4.0/legalcode.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by-nc/4.0/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="cc cc-by-nc-sa cc-2x"></i></th><th class='left'>表示-非営利-継承    </th><td><a href="http://creativecommons.org/licenses/by-nc-sa/4.0/">コモンズ証</a>（<a href="http://creativecommons.org/licenses/by-nc-sa/4.0/deed.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by-nc-sa/4.0/legalcode">法的条項</a>（<a href="http://creativecommons.org/licenses/by-nc-sa/4.0/legalcode.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by-nc-sa/4.0/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="cc cc-by-nd cc-2x"></i>   </th><th class='left'>表示-改変禁止       </th><td><a href="http://creativecommons.org/licenses/by-nd/4.0/">コモンズ証</a>（<a href="http://creativecommons.org/licenses/by-nd/4.0/deed.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by-nd/4.0/legalcode">法的条項</a>（<a href="http://creativecommons.org/licenses/by-nd/4.0/legalcode.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by-nd/4.0/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="cc cc-by-nc-nd cc-2x"></i></th><th class='left'>表示-非営利-改変禁止</th><td><a href="http://creativecommons.org/licenses/by-nc-nd/4.0/">コモンズ証</a>（<a href="http://creativecommons.org/licenses/by-nc-nd/4.0/deed.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by-nc-nd/4.0/legalcode">法的条項</a>（<a href="http://creativecommons.org/licenses/by-nc-nd/4.0/legalcode.ja">日本語</a>） <a href="http://creativecommons.org/licenses/by-nc-nd/4.0/rdf">メタデータ</a></td></tr>
</tbody>
</table>
{{< /fig-gen >}}

RDFa を使って Web ページにメタデータを埋め込む方法については拙文「[RDFa 入門](http://www.baldanders.info/spiegel/archive/rdfa/)」が参考になるだろう。
ページ全体にライセンスを指示したいなら `head` 要素に

```html
<link rel='cc:license' href='http://creativecommons.org/licenses/by/4.0/'>
```

と記述するだけでもよい（HTML5 の場合）。

## CC Licenses のバージョン

現在， [CC Licenses] として日本で有効なバージョンは以下のとおり。


- バージョン 1.0 Generic
- バージョン 2.0 Generic
- バージョン 2.1 日本版
- バージョン 2.5 Generic
- バージョン 3.0 Unported
- バージョン 4.0 International

このうち，厳密に日本法を準拠法として明記し，かつ法的に日本に最適化されているバージョンは「バージョン 2.1 日本版」のみである。

{{< fig-gen title="クリエイティブ・コモンズ・ライセンス バージョン 2.1 日本版" >}}
<table>
<tbody>
<tr><th class='left'><i class="cc cc-BY cc-2x"></i>         </th><th class='left'>表示                </th><td><a href="http://creativecommons.org/licenses/by/2.1/jp/">コモンズ証</a> <a href="http://creativecommons.org/licenses/by/2.1/jp/legalcode">法的条項</a> <a href="http://creativecommons.org/licenses/by/2.1/jp/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="cc cc-by-sa cc-2x"></i>      </th><th class='left'>表示-継承           </th><td><a href="http://creativecommons.org/licenses/by-sa/2.1/jp/">コモンズ証</a> <a href="http://creativecommons.org/licenses/by-sa/2.1/jp/legalcode">法的条項</a> <a href="http://creativecommons.org/licenses/by-sa/2.1/jp/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="cc cc-by-nc-jp cc-2x"></i>   </th><th class='left'>表示-非営利         </th><td><a href="http://creativecommons.org/licenses/by-nc/2.1/jp/">コモンズ証</a> <a href="http://creativecommons.org/licenses/by-nc/2.1/jp/legalcode">法的条項</a> <a href="http://creativecommons.org/licenses/by-nc/2.1/jp/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="cc cc-by-nc-sa-jp cc-2x"></i></th><th class='left'>表示-非営利-継承    </th><td><a href="http://creativecommons.org/licenses/by-nc-sa/2.1/jp/">コモンズ証</a> <a href="http://creativecommons.org/licenses/by-nc-sa/2.1/jp/legalcode">法的条項</a> <a href="http://creativecommons.org/licenses/by-nc-sa/2.1/jp/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="cc cc-by-nd cc-2x"></i>      </th><th class='left'>表示-改変禁止       </th><td><a href="http://creativecommons.org/licenses/by-nd/2.1/jp/">コモンズ証</a> <a href="http://creativecommons.org/licenses/by-nd/2.1/jp/legalcode">法的条項</a> <a href="http://creativecommons.org/licenses/by-nd/2.1/jp/rdf">メタデータ</a></td></tr>
<tr><th class='left'><i class="cc cc-by-nc-nd-jp cc-2x"></i></th><th class='left'>表示-非営利-改変禁止</th><td><a href="http://creativecommons.org/licenses/by-nc-nd/2.1/jp/">コモンズ証</a> <a href="http://creativecommons.org/licenses/by-nc-nd/2.1/jp/legalcode">法的条項</a> <a href="http://creativecommons.org/licenses/by-nc-nd/2.1/jp/rdf">メタデータ</a></td></tr>
</tbody>
</table>
{{< /fig-gen >}}

ただし，日本版以外のバージョンも日本発の作品であれば日本法を準拠法として解釈されるため通常は問題ない[^e]。
既に Generic / Unported の各バージョンを適用しているマテリアルについては，可能であれば 4.0 International 以降にアップデートすることを検討してみるべきだろう。
4.0  International では，コモンズ証や法的条項で公式な日本語訳が使えるほか，他の「自由なライセンス」との互換性が図られている。

[^e]: 国際取引の場合はどの国を準拠法とするかで問題になることがある。この場合はライセンスの説明に準拠法を併記するなどして利用者に知らせる必要がある。 

## その他 雑多なこと

### 「クレジット表示」について

未稿。

### CC Licenses における「営利」とは

未稿。

### CC Licenses と互換性のあるライセンス

未稿。

## 参考になる（かもしれない） Web ページ

- [CCライセンス・バージョン4.0 日本語版の公開 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2015/07/15/cc%E3%83%A9%E3%82%A4%E3%82%BB%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%90%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B34-0-%E6%97%A5%E6%9C%AC%E8%AA%9E%E7%89%88%E3%81%AE%E5%85%AC%E9%96%8B/)
- [バーチャルネット法律娘　真紀奈17歳](http://homepage3.nifty.com/machina/) : 古いコンテンツだが，このなかの「著作権法講座」は参考になる
- [著作権Q&A | 公益社団法人著作権情報センター CRIC](http://www.cric.or.jp/qa/index.html)

[本シリーズ]: /cc-licenses "改訂3版： CC-License について — text.Baldanders.info"
[著作権法]: http://law.e-gov.go.jp/htmldata/S45/S45HO048.html "著作権法"
[前回]: {{< relref "cc-licenses/01-copyright.md" >}} "著作権と著作権法 — 改訂3版： CC Licenses について"
[Creative Commons]: http://creativecommons.org/ "Creative Commons"
[CC Licenses]: http://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"