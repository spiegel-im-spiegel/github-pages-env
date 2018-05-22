+++
date = "2015-11-08T20:10:15+09:00"
update = "2018-05-22T15:40:44+09:00"
description = "今回はいよいよ Creative Commons Licenses について解説する。"
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

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]は著作権についてざくっと説明したので，今回はいよいよ Creative Commons Licenses （以降 “[CC Licenses]” と略称する）について解説する。

1. [CC Licenses の対象となるもの]({{< relref "#scope" >}})
1. [CC Licenses が示す利用条件]({{< relref "#conditions" >}})
1. [CC Licenses が許諾するもの]({{< relref "#grant" >}})
1. [CC Licenses の3つのレイヤ]({{< relref "#layer" >}})
1. [CC Licenses のバージョン]({{< relref "#versions" >}})
1. [その他 雑多なこと]({{< relref "#etc" >}})
    1. [「クレジット表示」について]({{< relref "#attribution" >}})
    1. [CC Licenses における「非営利」とは]({{< relref "#non-commercial" >}})
    1. [CC Licenses と互換性のあるライセンス]({{< relref "#compatible" >}})
    1. [CC Licenses における「無償」とは]({{< relref "#royalty-free" >}})
    1. [CC Licenses における「再許諾不可」とは]({{< relref "#non-sublicensable" >}})
    1. [CC Licenses における「非排他的」とは]({{< relref "#non-exclusive" >}})
    1. [責任の制限と消費者契約法]({{< relref "#limit" >}})
    1. [未成年が許諾者の場合]({{< relref "#agent" >}})

## CC Licenses の対象となるもの{#scope}

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

[^0]: 「プログラムの著作物」については [CC Licenses] は向いてないかもしれない。「プログラムの著作物」については GNU GPL など FLOSS（Free/Libre and Open Source Software）用の「自由なライセンス」がおすすめである（参考： [たくさんあるオープンソースライセンスのそれぞれの特徴のまとめ | コリス](https://coliss.com/articles/build-websites/operation/work/choose-a-license-by-github.html)）。

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

[^a]: こう呼ぶようになった経緯については「[CCライセンス・バージョン4.0 日本語版の公開 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2015/07/15/cc%E3%83%A9%E3%82%A4%E3%82%BB%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%90%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B34-0-%E6%97%A5%E6%9C%AC%E8%AA%9E%E7%89%88%E3%81%AE%E5%85%AC%E9%96%8B/)」を参照のこと。

## CC Licenses が示す利用条件{#conditions}

大雑把に言うと [CC Licenses] は「ライセンス利用条件（license conditions）」と条件に対する「ライセンス付与（license grant）」で構成されている。
利用条件はライセンスの「許諾者（licensor）」が指定する。
当然ながら許諾者（達）はマテリアルに対して許諾可能な権利を持っていなければならない[^aa]。

[^aa]: ひとつのマテリアルに複数の著作（権）者が関わっている場合，マテリアルに  [CC Licenses] を付与するためには全ての権利者の間で合意が必要である。

[CC Licenses] が示す利用条件は以下の4つの条件の組み合わせである。

{{< fig-gen >}}
<table>
<tbody>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i></th>
    <th class='left'>表示</th><td class='left'>マテリアルの<a href="{{< relref "#attribution" >}}">クレジットを表示</a>すること</td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-sa fa-2x"></i></th>
    <th class='left'>継承</th><td class='left'>改変したマテリアルについて，元のマテリアルのライセンスと<a href="{{< relref "#compatible" >}}">同等のライセンス</a>で公開すること</td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-nc fa-2x"></i>
                     <i class="fab fa-creative-commons-nc-eu fa-2x"></i>
                     <i class="fab fa-creative-commons-nc-jp fa-2x"></i></th>
    <th class='left'>非営利</th><td class='left'><a href="{{< relref "#non-commercial" >}}">営利目的での利用をしない</a>こと</td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-nd fa-2x"></i></th>
    <th class='left'>改変禁止</th><td class='left'>元のマテリアルを改変しないこと</td>
</tr>
</tbody>
</table>
{{< /fig-gen >}}

実際には「表示 <i class="fab fa-creative-commons-by"></i>」は必須の条件[^b] になっていて，さらに「継承 <i class="fab fa-creative-commons-sa"></i>」と「改変禁止 <i class="fab fa-creative-commons-nd"></i>」は条件が衝突するため同時に指定できない。
結局以下に示す6つの組み合わせのみが有効になる。

[^b]: バージョン1では「表示 <i class="fab fa-creative-commons-by"></i>」は任意だったが，実際には殆どの人が「表示 <i class="fab fa-creative-commons-by"></i>」条件を指定していたため，（ライセンス選択の煩雑さを防ぐ意味もあり）バージョン2以降は必須になった。

{{< fig-gen >}}
<table>
<tbody>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i></th>
    <td class='left'>表示</td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i></th>
    <td class='left'>表示-継承</td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc fa-2x"></i></th>
    <td class='left'>表示-非営利</td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i></th>
    <td class='left'>表示-非営利-継承</td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nd fa-2x"></i></th>
    <td class='left'>表示-改変禁止</td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nd fa-2x"></i></th>
    <td class='left'>表示-非営利-改変禁止</td>
</tr>
</tbody>
</table>
{{< /fig-gen >}}

## CC Licenses が許諾するもの{#grant}

マテリアルを利用したいユーザは，ライセンスで指定された上述の条件の組み合わせを守る限り，以下の利用が許諾される。

- **対象となるマテリアルの全部または一部を複製および共有すること**
    - マテリアルを複製する（複製）
        - 著作隣接権に係るマテリアルの場合は録音・録画を含む
    - マテリアルを公表・配布する（上演，演奏，上映，公衆送信等，口述，展示，頒布（映画の著作物のみ），譲渡・貸与（映画の著作物以外））
        - 著作隣接権に係るマテリアルの場合は（有線放送を含む）放送と再放送，伝達，送信可能化を含む
- **対象となるマテリアルの二次的著作物[^5] を作成，複製および共有すること** （「改変禁止 <i class="fab fa-creative-commons-nd"></i>」条件がない場合）
    - マテリアルを改変（翻訳・編曲・変形・翻案等）する
    - マテリアルを改変したものを公表・配布する（二次的著作物の利用に関する原著作者の権利の不行使）

[^5]: [著作権法]では二次的著作物を「著作物を翻訳し、編曲し、若しくは変形し、又は脚色し、映画化し、その他翻案することにより創作した著作物」と定義している（第2条）。「[二次的著作物と CC-Licenses]({{< relref "cc-licenses/04-derivative-works.md" >}})」で詳しく解説するが，二次的著作物については「創作的」であることが条件。フォーマット変換や機械翻訳（点字などへの置き換え）などは「逐語的コピー」と呼ばれ複製と見なされる。

なお [CC Licenses] は[「公正な利用（fair use）」や「著作権の制限」]({{< relref "cc-licenses/01-copyright.md#fair-use" >}})として利用が認められているものに対しては効力が及ばない。
[「公正な利用」や「著作権の制限」]({{< relref "cc-licenses/01-copyright.md#fair-use" >}})で認められる利用については条件に関係なく可能である。

許諾は全世界で適用される。
また，[無償（royalty-free）]({{< relref "#royalty-free" >}})かつ[再許諾不可（non-sublicensable）]({{< relref "#non-sublicensable" >}})かつ[非排他的（non-exclusive）]({{< relref "#non-exclusive" >}})に許諾される。

許諾者はマテリアルに付与したライセンスを取り消しできない。
したがって，少なくともライセンスの条件を守っていればマテリアルの利用を後から取り消されることはない[^c]。
またライセンスの効力は，著作権が存続する限り持続する（著作権の存続期間が過ぎればマテリアルは公有（public domain）に帰す）。

[^c]: つまり許諾者は，ライセンスを付与する場合には妥当な条件であるか予め考慮する必要がある。

## CC Licenses の3つのレイヤ{#layer}

[CC Licenses] は以下の3つのレイヤで構成されている。

- **コモンズ証（commons deed）** [CC Licenses] の利用条件と許諾について Human-Readable な形で簡潔に記述したもの。
- **法的条項（legal code）** [CC Licenses] の内容を Lawyer-Readable な形で記述した「利用許諾書」。 法的条項によって [CC Licenses] が法的拘束力を持つことを担保する。
- **メタデータ（metadata）** マテリアルに関する情報を Machine-Readable な形で記述したもの。記述形式としては XML, RDFa, XMP などがあるが， Web 上では RDFa がよく用いられる。メタデータは Google などの検索サービスで解釈され，ネット上の数多のマテリアルを利用条件によって検索可能とする。

マテリアルに対して [CC Licenses] を指示する場合は，通常はコモンズ証へのリンクを示せばよい（コモンズ証のページから法的条項へリンクが張られている）。
各レイヤのリンクを以下に示す。

{{< fig-gen title="Creative Commons Licenses Version 4.0 International" >}}
<table>
<tbody>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i></th>
    <th class='left'>表示</th>
    <td>
        <a href="https://creativecommons.org/licenses/by/4.0/">コモンズ証</a>
        （<a href="https://creativecommons.org/licenses/by/4.0/deed.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by/4.0/legalcode">法的条項</a>
        （<a href="https://creativecommons.org/licenses/by/4.0/legalcode.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by/4.0/rdf">メタデータ</a>
    </td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i></th>
    <th class='left'>表示-継承</th>
    <td>
        <a href="https://creativecommons.org/licenses/by-sa/4.0/">コモンズ証</a>
        （<a href="https://creativecommons.org/licenses/by-sa/4.0/deed.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by-sa/4.0/legalcode">法的条項</a>
        （<a href="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by-sa/4.0/rdf">メタデータ</a>
    </td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc fa-2x"></i></th>
    <th class='left'>表示-非営利</th>
    <td>
        <a href="https://creativecommons.org/licenses/by-nc/4.0/">コモンズ証</a>
        （<a href="https://creativecommons.org/licenses/by-nc/4.0/deed.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by-nc/4.0/legalcode">法的条項</a>
        （<a href="https://creativecommons.org/licenses/by-nc/4.0/legalcode.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by-nc/4.0/rdf">メタデータ</a>
    </td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i></th>
    <th class='left'>表示-非営利-継承</th>
    <td>
        <a href="https://creativecommons.org/licenses/by-nc-sa/4.0/">コモンズ証</a>
        （<a href="https://creativecommons.org/licenses/by-nc-sa/4.0/deed.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode">法的条項</a>
        （<a href="https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by-nc-sa/4.0/rdf">メタデータ</a>
    </td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nd fa-2x"></i></th>
    <th class='left'>表示-改変禁止</th>
    <td>
        <a href="https://creativecommons.org/licenses/by-nd/4.0/">コモンズ証</a>
        （<a href="https://creativecommons.org/licenses/by-nd/4.0/deed.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by-nd/4.0/legalcode">法的条項</a>
        （<a href="https://creativecommons.org/licenses/by-nd/4.0/legalcode.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by-nd/4.0/rdf">メタデータ</a>
    </td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nd fa-2x"></i></th>
    <th class='left'>表示-非営利-改変禁止</th>
    <td>
        <a href="https://creativecommons.org/licenses/by-nc-nd/4.0/">コモンズ証</a>
        （<a href="https://creativecommons.org/licenses/by-nc-nd/4.0/deed.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode">法的条項</a>
        （<a href="https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode.ja">日本語</a>）<br>
        <a href="https://creativecommons.org/licenses/by-nc-nd/4.0/rdf">メタデータ</a>
    </td>
</tr>
</tbody>
</table>
{{< /fig-gen >}}

RDFa を使って Web ページにメタデータを埋め込む方法については拙文「[RDFa 入門](http://www.baldanders.info/spiegel/archive/rdfa/)」が参考になるだろう。
ページ全体にライセンスを指示したいなら `head` 要素に

```html
<link rel='cc:license' href='https://creativecommons.org/licenses/by-sa/4.0/'>
```

と記述するだけでもよい（HTML5 の場合）。

法的条項についてはテキストファイルでも提供されている。

{{< fig-gen title="Creative Commons Licenses Version 4.0 International" >}}
<table>
<tbody>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i></th>
    <th class='left'>表示</th>
    <td><a href="https://creativecommons.org/licenses/by/4.0/legalcode.txt">法的条項（テキスト版）</a></td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i></th>
    <th class='left'>表示-継承</th>
    <td><a href="https://creativecommons.org/licenses/by-sa/4.0/legalcode.txt">法的条項（テキスト版）</a></td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc fa-2x"></i></th>
    <th class='left'>表示-非営利</th>
    <td><a href="https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt">法的条項（テキスト版）</a></td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i></th>
    <th class='left'>表示-非営利-継承</th>
    <td><a href="https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode.txt">法的条項（テキスト版）</a></td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nd fa-2x"></i>
    </th><th class='left'>表示-改変禁止</th>
    <td><a href="https://creativecommons.org/licenses/by-nd/4.0/legalcode.txt">法的条項（テキスト版）</a></td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nd fa-2x"></i></th>
    <th class='left'>表示-非営利-改変禁止</th>
    <td><a href="https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode.txt">法的条項（テキスト版）</a></td>
</tr>
</tbody>
</table>
{{< /fig-gen >}}

## CC Licenses のバージョン{#versions}

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
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i></th>
    <th class='left'>表示</th>
    <td>
        <a href="https://creativecommons.org/licenses/by/2.1/jp/">コモンズ証</a><br>
        <a href="https://creativecommons.org/licenses/by/2.1/jp/legalcode">法的条項</a><br>
        <a href="https://creativecommons.org/licenses/by/2.1/jp/rdf">メタデータ</a>
    </td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i></th>
    <th class='left'>表示-継承</th>
    <td>
        <a href="https://creativecommons.org/licenses/by-sa/2.1/jp/">コモンズ証</a><br>
        <a href="https://creativecommons.org/licenses/by-sa/2.1/jp/legalcode">法的条項</a><br>
        <a href="https://creativecommons.org/licenses/by-sa/2.1/jp/rdf">メタデータ</a>
    </td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc-jp fa-2x"></i></th>
    <th class='left'>表示-非営利</th>
    <td>
        <a href="https://creativecommons.org/licenses/by-nc/2.1/jp/">コモンズ証</a><br>
        <a href="https://creativecommons.org/licenses/by-nc/2.1/jp/legalcode">法的条項</a><br>
        <a href="https://creativecommons.org/licenses/by-nc/2.1/jp/rdf">メタデータ</a>
    </td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc-jp fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i></th>
    <th class='left'>表示-非営利-継承</th>
    <td>
        <a href="https://creativecommons.org/licenses/by-nc-sa/2.1/jp/">コモンズ証</a><br>
        <a href="https://creativecommons.org/licenses/by-nc-sa/2.1/jp/legalcode">法的条項</a><br>
        <a href="https://creativecommons.org/licenses/by-nc-sa/2.1/jp/rdf">メタデータ</a>
    </td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nd fa-2x"></i></th>
    <th class='left'>表示-改変禁止</th>
    <td>
        <a href="https://creativecommons.org/licenses/by-nd/2.1/jp/">コモンズ証</a><br>
        <a href="https://creativecommons.org/licenses/by-nd/2.1/jp/legalcode">法的条項</a><br>
        <a href="https://creativecommons.org/licenses/by-nd/2.1/jp/rdf">メタデータ</a>
    </td>
</tr>
<tr>
    <th class='left'><i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nc-jp fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-nd fa-2x"></i></th>
    <th class='left'>表示-非営利-改変禁止</th>
    <td>
        <a href="https://creativecommons.org/licenses/by-nc-nd/2.1/jp/">コモンズ証</a><br>
        <a href="https://creativecommons.org/licenses/by-nc-nd/2.1/jp/legalcode">法的条項</a><br>
        <a href="https://creativecommons.org/licenses/by-nc-nd/2.1/jp/rdf">メタデータ</a>
    </td>
</tr>
</tbody>
</table>
{{< /fig-gen >}}

ただし，日本版以外のバージョンも日本発の作品であれば日本法を準拠法として解釈されるため通常は問題ない[^e]。

[^e]: 国際取引の場合はどの国を準拠法とするかで問題になることがある。この場合はライセンスの説明に準拠法を併記するなどして利用者に知らせる必要がある。

既に Generic / Unported の各バージョンを適用しているマテリアルについては，可能であれば 4.0 International 以降にアップデートすることを検討してみるべきだろう。
4.0  International では，コモンズ証や法的条項で公式な日本語訳が使えるほか，他の「自由なライセンス」との互換性が図られている。

## その他 雑多なこと{#etc}

### 「クレジット表示」について{#attribution}

[CC Licenses] 下で対象となるマテリアル（二次的著作物を含む）の全部または一部を複製および共有する際にはマテリアルの著作（権）者のクレジットを表示する必要がある。
クレジット表示としては以下のものがある（全部ではないが元のマテリアルで表示されていればそれら全てが対象）。

- 著作（権）者名[^f]
- 著作権表示
- ライセンス表示
- 「無保証」を参照する表示
- マテリアルを示す URI （またはハイパーリンク）

[^f]: 法的条項には「ライセンス対象物の作者その他クレジット表示される者として許諾者によって指定されている者を識別する情報」を「許諾者によってリクエストされた形が合理的である場合はその形で」とある。つまり，できるだけ表記を変えないようにってことらしい。

ただし，この条件は

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>あなたがライセンス対象物を共有する媒体・方法・文脈に照らして、いかなる合理的な方法でも満たすことができます。
例えば、必要とされる情報を含むリソースのURIやハイパーリンクを付すことで条件を満たすことが合理的な場合があります。</q>
{{< /fig-quote >}}

とあり，杓子定規に全部表示しなきゃいけないというわけではない。
要は複製物や改変物を利用する人に元の著作（権）者とライセンスが分かるようになっていればよい。

また，このクレジット表示は元の著作（権）者からの要請があれば削除しなければいけない。
これはちょっと変わっているように見えるが， [CC Licenses] では「表示 <i class="fab fa-creative-commons-by"></i>」条件が必須になっているため，わりと重要な条項である[^g]。

[^g]: これは [CC Licenses] 初期に「嫌煙家の作った [CC Licenses] 作品をタバコ広告で利用できるか」といった議論があり，利用そのものは取り消せないが，クレジット表示を止めさせることはできるんじゃないの，という話になった。これは，その時の議論をライセンスとして明文化したものといえる。

一方，マテリアルを [CC Licenses] で公開する際には，少なくとも「著作（権）者名」「ライセンス」「公表年」を表示し URI で指示できるようにするとよい[^g2]。

[^g2]: &copy; マークなどの「著作権表示」は不要。現在では「著作権表示」を要求する国や地域はほとんどど無い。また「無保証」は [CC Licenses] 内で明記されているので敢えて表示する必要はない。ただし「無保証」の範囲を明示するのはアリ。

### CC Licenses における「非営利」とは{#non-commercial}

巷にあるライセンスでは「営利」あるいは「商用」などのキーワードが付くことがあるが，どこまでが「営利」で「商用」なのか分かりにくいことが多い。
[CC Licenses] では「非営利 <i class="fab fa-creative-commons-nc"></i>」について

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-非営利 4.0 国際" link="https://creativecommons.org/licenses/by-nc/4.0/legalcode.ja" >}}
<q>商業的な利得や金銭的報酬を、主たる目的とせず、それらに主に向けられてもいないことを意味します。本パブリック・ライセンスにおいては、デジタル・ファイル共有または類似した手段による、ライセンス対象物と、著作権およびそれに類する権利の対象となるその他のマテリアルとの交換は、その交換に関連して金銭的報酬の支払いがない場合は、非営利に該当します。</q>
{{< /fig-quote >}}

と定義し「商業的な利得や金銭的報酬」が主目的でなければ「非営利 <i class="fab fa-creative-commons-nc"></i>」条件を満たすとしている[^c2]。
後半の文言は分かりにくいが，これはファイル交換などオンラインでの取引を念頭に置いたもので，この場合でも金銭的報酬がなければ「非営利 <i class="fab fa-creative-commons-nc"></i>」条件を満たすとみなしている。

[^c2]: 一般的に「営利」「商用」という場合にはかなり広い範囲をさすようだ。たとえば企業における業務利用は全て（直接の金銭的報酬がなくとも）「営利」「商用」と見なされる。面白い例を挙げると， IPA が公開している「[GNU GPL v3 逐条解説書](https://www.ipa.go.jp/osc/license1.html "GNU GPL v3 解説書：IPA 独立行政法人 情報処理推進機構")」は日本版（v2.1）の「[表示-非営利-改変禁止](http://creativecommons.org/licenses/by-nc-nd/2.1/jp/)」ライセンスになっているが「出版等営利目的での利用は禁止しますが、企業・団体等の内部における利用（講習会、勉強会等）を目的とした複製及び翻訳については、無償で許可します」という但し書きがある。おそらく [CC Licenses] の「非営利 <i class="fab fa-creative-commons-nc"></i>」について拡大解釈しすぎないよう配慮したものと思われる（実際には「改変禁止 <i class="fab fa-creative-commons-nd"></i>」条件では翻訳は許諾範囲外だが，利用状況によって個別に許可を与えていると解釈できる。このように [CC Licenses] は許諾範囲を広げる方向に対しては割と柔軟に運用できる）。

ここで微妙なのは，アフィリエイト広告のある Web ページに「非営利 <i class="fab fa-creative-commons-nc"></i>」条件のあるマテリアルを利用することができるかということだが， [Creative Commons] 創設者のひとりである [Lawrence Lessig](http://lessig.tumblr.com/) 教授の2008年当時の発言[^g3] によると

{{< fig-quote title="Creative Commons “Non-Commercial” Content on Ad-Supported Sites" link="http://blogoscoped.com/archive/2008-02-07-n77.html" lang="en" >}}
<q>I’ve asked the Wikipedia mailing list a while ago, and recently received another confirmation from Creative Commons’ Lawrence Lessig: <strong>yes, the CC organization believes this being OK is the best reading of the license.</strong></q>
{{< /fig-quote >}}

ということで，ひとまず問題ないようだ[^g4]。

[^g3]: [Lawrence Lessig](http://lessig.tumblr.com/) 教授は現在， [Creative Commons] からは距離を置いて「政治と腐敗」をテーマに研究を行っているが，[2016年の米国大統領選に民主党候補として出馬意向を表明](https://lessig2016.us/)した。
[^g4]: [Creative Commons] 側は，マテリアルの取り引きによって金銭的報酬が発生するようなことでなければ「非営利 <i class="fab fa-creative-commons-nc"></i>」条件にかかることはないと考えているようである（あとは企業広告の素材として使う場合とか）。ただ，一時期流行ったアフィリエイト報酬を主目的とする「パクりサイト」の場合はダメな気がする。司法判断を仰ぐ必要があるかもしないけど。ちなみにこのサイトは「[表示-継承](https://creativecommons.org/licenses/by-sa/4.0/)」条件を守っていただけるなら，いくらでもパクって稼いで構いませんよ（笑）

### CC Licenses と互換性のあるライセンス{#compatible}

「継承 <i class="fab fa-creative-commons-sa"></i>」条件について「元のマテリアルのライセンスと同等のライセンスで公開すること」と書いたが，「同等のライセンス」とはどこまでを指すのか。
具体的には以下のページで示される。

- [Compatible Licenses - Creative Commons](https://creativecommons.org/compatiblelicenses)

基本的には [CC Licenses] の同じ条件であれば他のバージョンであっても「同等のライセンス」とみなされる。
さらに「[表示-継承 <i class="fab fa-creative-commons-by"></i><i class="fab fa-creative-commons-sa"></i>](https://creativecommons.org/licenses/by-sa/4.0/)」については以下のライセンスも「同等のライセンス」と見なすことができる。

- [Free Art License](http://artlibre.org/licence/lal/en/) 1.3
- [GNU GPL (General Public License)](https://www.gnu.org/copyleft/gpl.html) version 3[^h]

[^h]: ただし， GPL の場合は「[表示-継承](https://creativecommons.org/licenses/by-sa/4.0/)」から GPLv3 へ一方向のみの互換性である。 GPLv3 ライセンスのマテリアルを [CC Licenses] の「[表示-継承](https://creativecommons.org/licenses/by-sa/4.0/)」として扱うことはできない。

### CC Licenses における「無償」とは{#royalty-free}

[CC Licenses] は許諾に際して「無償（royalty-free）」であることを保証する。
ただしこれは，利用に対する許諾が無償で行われるという意味で，普通にマテリアルを売り買いすることに関しては「非営利 <i class="fab fa-creative-commons-nc"></i>」条件がなければ問題ない。

これには JASRAC のような著作権管理団体からの royalty 徴収も含むが「可能な限り」の但し書きがつく。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>可能なかぎり、許諾者は、ライセンスされた権利の行使について、直接か、または任意のもしくは放棄可能な法定のもしくは強制的なライセンスに関する仕組みに基づく集中管理団体を介するかを問わず、あなたからライセンス料を得るいかなる権利も放棄します。その他一切の場合において、許諾者はそのようなライセンス料を得るいかなる権利も明確に保持します。</q>
{{< /fig-quote >}}

この点については別の記事で解説する。

### CC Licenses における「再許諾不可」とは{#non-sublicensable}

[CC Licenses] ではマテリアルの再配布[^j] を行う際にライセンスの再許諾（sublicense）を許可していない。
再配布されたマテリアルにも自動的に元のライセンスが付与される。

{{< fig-img src="/images/non-sublicense.svg" title="再配布と許諾の関係" link="/images/non-sublicense.svg" width="567" >}}

これは，たとえばマテリアルを著作（権）者ではない出版者（社）などが配布する際に勝手にライセンスを変えたり [CC Licenses] にない制限を加えたりしてはいけないという意味である。
これにはいわゆる「技術的保護手段」も含まれる[^d]。

[^j]: ここで「再配布」は，マテリアルの利用者がマテリアルまたはマテリアルの複製を更に下流（downstream）へ公表・配布することを指す。
[^d]: 更に言うと，マテリアルが「技術的保護手段」によって不当に保護されている場合は，ユーザは「技術的保護手段」を迂回することができる（許諾者は迂回を禁止する権利を放棄する）。「技術的保護手段」については「[CC Licenses における「技術的保護手段」の扱い]({{< relref "cc-licenses/05-technological-measures.md" >}})」で解説する。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>あなたは、ライセンス対象物の受領者がライセンスされた権利を行使するのを制限されることになる場合には、ライセンス対象物に対して、いかなる追加条項または異なる条項も提案または課してはならず、あるいは、いかなる効果的な技術的保護手段も適用してはなりません。</q>
{{< /fig-quote >}}

### CC Licenses における「非排他的」とは{#non-exclusive}

[CC Licenses] では許諾は「非排他的（non-exclusive）」に行われる。
これは，あるマテリアルについて [CC Licenses] を含む複数のライセンスを付与することが可能であることを指す。
たとえば「非営利 <i class="fab fa-creative-commons-nc"></i>」条件が付くマテリアルを営利目的で利用したい場合に，著作（権）者と別途契約して，報酬を払う代わりに営利目的での利用を許諾してもらう，といったことも（交渉次第では）可能である。

ただし，通常は排他的あるいは独占的ライセンスと [CC Licenses] のようなライセンスは両立しないため，複数ライセンスの付与には注意が必要である。

余談だが， [CC Licenses] の許諾者はライセンス利用条件に縛られない。
たとえば「非営利 <i class="fab fa-creative-commons-nc"></i>」条件を付与したマテリアルを許諾者自身が有料で取り引きする，といったことは自由にできる。

### 責任の制限と消費者契約法{#limit}

[CC Licenses] には「責任制限」の項目があり対象のマテリアルに対していかなる表明も保証も行わないことになっている。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>許諾者が別途合意しない限り、許諾者は可能な範囲において、ライセンス対象物を現状有姿のまま、現在可能な限りで提供し、明示、黙示、法令上、その他に関わらずライセンス対象物について一切の表明または保証をしません。
これには、権利の帰属、商品性、特定の利用目的への適合性、権利侵害の不存在、隠れた瑕疵その他の瑕疵の不存在、正確性または誤りの存在もしくは不存在を含みますが、これに限られず、既知であるか否か、発見可能であるか否かを問いません。
全部または一部の無保証が認められない場合、この無保証はあなたには適用されないこともあります。</q>
{{< /fig-quote >}}

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>可能な範囲において、本パブリック・ライセンスもしくはライセンス対象物の利用によって起きうる直接、特別、間接、偶発、結果的、懲罰的その他の損失、コスト、出費または損害について、例え損失、コスト、出費、損害の可能性について許諾者が知らされていたとしても、許諾者は、あなたに対し、いかなる法理（過失を含みますがこれに限られません）その他に基づいても責任を負いません。
全部または一部の責任制限が認められない場合、この制限はあなたには適用されないこともあります。</q>
{{< /fig-quote >}}

ただし「全部または一部の無保証（責任制限）が認められない場合、この無保証（制限）はあなたには適用されないこともあります」という部分がポイントで，たとえば日本の[消費者契約法](http://law.e-gov.go.jp/htmldata/H12/H12HO061.html)では

{{< fig-quote title="消費者契約法" link="http://law.e-gov.go.jp/htmldata/H12/H12HO061.html" >}}
<q>事業者の損害賠償の責任を免除する条項その他の消費者の利益を不当に害することとなる条項の全部又は一部を無効とする</q>
{{< /fig-quote >}}

とあり（第1条），許諾者が事業者で[^k] かつ免責が無効であると判断された場合は許諾者は従う必要がある。

[^k]: [消費者契約法](http://law.e-gov.go.jp/htmldata/H12/H12HO061.html)における「事業者」は「法人その他の団体及び事業として又は事業のために契約の当事者となる場合における個人」（第2条）と定義されている。

### 未成年が許諾者の場合{#agent}

日本の民法では未成年者（満20歳未満）が許諾を行う場合には法定代理人（親権者など）の同意を得る必要があり，法定代理人の同意がない場合には許諾が取り消されることがある。

許諾者が明らかに未成年であれば，法定代理人の同意の有無を確認したほうが安全である。
一方，未成年の許諾者は法定代理人の存在およびその同意があることを何らかの形で明示したほうがトラブルが少なくて済む。

## 参考文献

### 参考になる（かもしれない） Web ページ

- [バーチャルネット法律娘　真紀奈17歳](http://machina.private.coocan.jp/) : 古いコンテンツだが，このなかの「著作権法講座」は参考になる
- [著作権Q&A | 公益社団法人著作権情報センター CRIC](http://www.cric.or.jp/qa/index.html)
- [クリエイティブ・コモンズ・ライセンスのブログ翻訳のススメ](http://www.yamdas.org/column/technique/cctrans.html)
- [CC Licenses] Version 3 検討時の議論
    - [【CCPLv3.0】著作者人格権（同一性保持権）に関する議論 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2006/11/15/ccplv3-discussion/)
    - [【CCPLv3.0】ＤＲＭ条項の改正に関する議論 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2006/11/15/discussion-drm/)
    - [【CCPLv3.0】著作権管理団体を通じての報酬請求権に関する議論 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2006/11/16/ccplv3/)
- [Creative Commons "Non-Commercial" Content on Ad-Supported Sites](http://blogoscoped.com/archive/2008-02-07-n77.html)
- [What's New in 4.0 - Creative Commons](https://creativecommons.org/Version4/) （[日本語参考訳](http://qiita.com/nyampire/items/c03904bd27bd8812aad3)）
- [CCライセンス・バージョン4.0 日本語版の公開 | クリエイティブ・コモンズ・ジャパン](http://creativecommons.jp/2015/07/15/cc%E3%83%A9%E3%82%A4%E3%82%BB%E3%83%B3%E3%82%B9%E3%83%BB%E3%83%90%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B34-0-%E6%97%A5%E6%9C%AC%E8%AA%9E%E7%89%88%E3%81%AE%E5%85%AC%E9%96%8B/)
- [CC BY-SA 4.0 now one-way compatible with GPLv3 - Creative Commons Blog - Creative Commons](https://blog.creativecommons.org/2015/10/08/cc-by-sa-4-0-now-one-way-compatible-with-gplv3/) （[日本語訳](http://creativecommons.jp/2016/01/25/cc-by-sa-%EF%BC%88%E8%A1%A8%E7%A4%BA-%E7%B6%99%E6%89%BF%EF%BC%89-4-0%E3%81%8B%E3%82%89gpl-v3%E3%81%B8%E3%81%AE%E4%B8%80%E6%96%B9%E5%90%91%E3%81%AE%E4%BA%92%E6%8F%9B%E3%81%8C%E5%AE%9F%E7%8F%BE/)）
- [How should we attribute 3D printed objects? - Creative Commons](https://creativecommons.org/2016/04/19/attribute-3d-printed-objects/)
    - [3Dプリントでできた物にはどのようなライセンス表示をすべきか？ | クリエイティブ・コモンズ・ジャパン](https://creativecommons.jp/2016/07/26/3d%e3%83%97%e3%83%aa%e3%83%b3%e3%83%88%e3%81%a7%e3%81%a7%e3%81%8d%e3%81%9f%e7%89%a9%e3%81%ab%e3%81%af%e3%81%a9%e3%81%ae%e3%82%88%e3%81%86%e3%81%aa%e3%83%a9%e3%82%a4%e3%82%bb%e3%83%b3%e3%82%b9%e8%a1%a8/)
- [【解説】 クリエイティブ・コモンズ・ライセンス入門 【知財管理65巻6号掲載】 | 弁護士 増田雅史の記録帳](https://masudalaw.wordpress.com/2016/05/06/ccl-basics/)
- [クリエイティブ・コモンズと著作権の新しい潮流](http://www.slideshare.net/JEPAslide/ss-68121343)
- [GitHub リポジトリに CC Licenses を設定したい - Qiita](https://qiita.com/spiegel-im-spiegel/items/0997f1693a24e3fd3a74)

[本シリーズ]: /cc-licenses "改訂3版： CC-License について — text.Baldanders.info"
[著作権法]: http://law.e-gov.go.jp/htmldata/S45/S45HO048.html "著作権法"
[前回]: {{< relref "cc-licenses/01-copyright.md" >}} "著作権と著作権法 — 改訂3版： CC Licenses について"
[Creative Commons]: https://creativecommons.org/ "Creative Commons"
[CC Licenses]: https://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"

### 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/475710152X/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41WPNBY7HZL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/475710152X/baldandersinf-22/">クリエイティブ・コモンズ―デジタル時代の知的財産権</a></dt><dd>ローレンス レッシグ 椙山 敬士 上村 圭介 林 紘一郎 若槻 絵美 土屋 大洋 クリエイティブコモンズジャパン </dd><dd>NTT出版 2005-03</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798106801/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798106801.09._SCTHUMBZZZ_.jpg"  alt="Free Culture"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798102040/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798102040.09._SCTHUMBZZZ_.jpg"  alt="コモンズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798119806/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798119806.09._SCTHUMBZZZ_.jpg"  alt="REMIX ハイブリッド経済で栄える文化と商業のあり方"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4087205274/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4087205274.09._SCTHUMBZZZ_.jpg"  alt="著作権の世紀―変わる「情報の独占制度」 (集英社新書 527A)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4480065733/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4480065733.09._SCTHUMBZZZ_.jpg"  alt="デジタル時代の著作権 (ちくま新書)"  /></a> </p>
<p class="description">残念ながら紙の本は実質的に絶版なんですよねぇ。是非デジタル化を希望します。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2014-08-02">2014/08/02</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
