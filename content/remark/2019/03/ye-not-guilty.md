+++
title = "「自由とはそういうことだ」"
date = "2019-03-17T10:16:57+09:00"
description = "「汝ら罪なし」"
image = "/images/attention/kitten.jpg"
tags = ["moral-rights", "creative-commons", "license", "flickr", "artificial-intelligence", "engineering"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

面白い記事を見つけたので。

- [IBMがAI顔認識のトレーニングにFlickrの写真利用--ユーザーの同意がないとして物議 - CNET Japan](https://japan.cnet.com/article/35134198/)

記事でも書かれているが [CC Licenses] は「表現の自由」を規制する著作権（法）を根拠としたライセンスである。
たとえば同じ知財権でもアイデアは特許権で守られているし名前や意匠は商標権や意匠権で守られている。
人工的に作られた「権利」は注意深く運用すべきであり，守るべき本来のドメインを越えてはいけない[^cr1]。

[^cr1]: これはもちろん昨今の著作権法改訂の「暴走」に対する皮肉だよ。

ようするに

{{< fig-quote title="IBMがAI顔認識のトレーニングにFlickrの写真利用--ユーザーの同意がないとして物議" link="https://japan.cnet.com/article/35134198/" >}}
<q>すでにCreative Commonsのライセンスを選択しているため、彼らはデータセットにオプトインする必要はなかった。彼らはすでにアクションしていたのだから。これがライセンスの仕組みだ。この仕組みのおかげで、世界中のアーティストや科学者がCCライセンスの適用されたものを使って創造したり発明したりすることが可能になっている</q>
{{< /fig-quote >}}

ということである。

## CC Licenses における「非営利」とは（「[Creative Commons Licenses]({{< ref "/cc-licenses/02-creative-commons-licenses.md" >}})」より抜粋）

巷にあるライセンスでは「営利」あるいは「商用」などのキーワードが付くことがあるが，どこまでが「営利」で「商用」なのか分かりにくいことが多い。
[CC Licenses] では「非営利 <i class="fab fa-creative-commons-nc"></i>」について

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-非営利 4.0 国際" link="https://creativecommons.org/licenses/by-nc/4.0/legalcode.ja" >}}
<q>商業的な利得や金銭的報酬を、主たる目的とせず、それらに主に向けられてもいないことを意味します。本パブリック・ライセンスにおいては、デジタル・ファイル共有または類似した手段による、ライセンス対象物と、著作権およびそれに類する権利の対象となるその他のマテリアルとの交換は、その交換に関連して金銭的報酬の支払いがない場合は、非営利に該当します。</q>
{{< /fig-quote >}}

と定義し「商業的な利得や金銭的報酬」が主目的でなければ「非営利 <i class="fab fa-creative-commons-nc"></i>」条件を満たすとしている[^cm2]。
後半の文言は分かりにくいが，これはファイル交換などオンラインでの取引を念頭に置いたもので，この場合でも金銭的報酬がなければ「非営利 <i class="fab fa-creative-commons-nc"></i>」条件を満たすとみなしている。

[^cm2]: 一般的に「営利」「商用」という場合にはかなり広い範囲をさすようだ。たとえば企業における業務利用は全て（直接の金銭的報酬がなくとも）「営利」「商用」と見なされる。面白い例を挙げると， IPA が公開している「[GNU GPL v3 逐条解説書](https://www.ipa.go.jp/osc/license1.html "GNU GPL v3 解説書：IPA 独立行政法人 情報処理推進機構")」は日本版（v2.1）の「[表示-非営利-改変禁止](https://creativecommons.org/licenses/by-nc-nd/2.1/jp/)」ライセンスになっているが「出版等営利目的での利用は禁止しますが、企業・団体等の内部における利用（講習会、勉強会等）を目的とした複製及び翻訳については、無償で許可します」という但し書きがある。おそらく [CC Licenses] の「非営利 <i class="fab fa-creative-commons-nc"></i>」について拡大解釈しすぎないよう配慮したものと思われる（実際には「改変禁止 <i class="fab fa-creative-commons-nd"></i>」条件では翻訳は許諾範囲外だが，利用状況によって個別に許可を与えていると解釈できる。このように [CC Licenses] は許諾範囲を広げる方向に対しては割と柔軟に運用できる）。

ここで微妙なのは，アフィリエイト広告のある Web ページに「非営利 <i class="fab fa-creative-commons-nc"></i>」条件のあるマテリアルを利用することができるかということだが， [Creative Commons] 創設者のひとりである [Lawrence Lessig](http://lessig.tumblr.com/) 教授の2008年当時の発言によると

{{< fig-quote title="Creative Commons “Non-Commercial” Content on Ad-Supported Sites" link="http://blogoscoped.com/archive/2008-02-07-n77.html" lang="en" >}}
<q>I’ve asked the Wikipedia mailing list a while ago, and recently received another confirmation from Creative Commons’ Lawrence Lessig: <strong>yes, the CC organization believes this being OK is the best reading of the license.</strong></q>
{{< /fig-quote >}}

ということで，ひとまず問題ないようだ[^g4]。

[^g4]: [Creative Commons] 側は，マテリアルの取り引きによって金銭的報酬が発生するようなことでなければ「非営利 <i class="fab fa-creative-commons-nc"></i>」条件にかかることはないと考えているようである（あとは企業広告の素材として使う場合とか）。ただ，一時期流行ったアフィリエイト報酬を主目的とする「パクりサイト」の場合はダメな気がする。司法判断を仰ぐ必要があるかもしないけど。ちなみにこのサイトは「[表示-継承](https://creativecommons.org/licenses/by-sa/4.0/)」条件を守っていただけるなら，いくらでもパクって稼いで構いませんよ（笑）

## CC Licenses では（可能なかぎり）人格権は行使されない（「[人格権と CC Licenses]({{< ref "/cc-licenses/03-moral-rights.md" >}})」より抜粋）

人格権について [CC Licenses] はどうしているかというと，許諾範囲を逸脱しない限り許諾者は人格権を行使しないことになっている。

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

## ブックマーク

- [クーリエ連載；エコノミスト紹介、自由のためなら人が死んでもいい](https://cruel.org/economist/courier200712.html)

[Creative Commons]: https://creativecommons.org/ "Creative Commons"
[CC Licenses]: https://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"

## 参考図書

{{% review-paapi "B0732279VT" %}} <!-- THE ビッグオー -->
{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
