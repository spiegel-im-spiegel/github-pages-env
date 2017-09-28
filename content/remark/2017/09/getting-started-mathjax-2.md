+++
title = "ちょこっと MathJax： 基本的な数式表現"
date =  "2017-09-28T22:25:54+09:00"
description = "今回は基本的な数式の書き方を説明していこう。"
tags        = [ "math", "tex", "mathjax" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = true
  mermaidjs = false
+++

さて，初期設定は[前回]で完了したので，今回は基本的な数式の書き方を説明していこう。

1. [ちょこっと MathJax： 初期設定]({{< relref "remark/2017/09/getting-started-mathjax-1.md" >}})
2. [ちょこっと MathJax： 基本的な数式表現]({{< relref "remark/2017/09/getting-started-mathjax-2.md" >}}) ← イマココ

## 数式表現の「お約束」 {#rule}

初っ端からナニだが，数式を書く際には昔からある「お約束」が存在する。

まず最初に，大原則として「数式で使う文字はイタリック体[^it1]で書く」というのがある。
たとえば，こんな感じだ。

[^it1]: $\\mathrm{\TeX}$ ではイタリック体（italics）と斜体（slanted）は異なる字体である。また斜体はしばしばローマン体（Roman）から合成される場合がある。

> エネルギーと質量には $$E=mc^2$$ の関係がある。

$E$ や $m$ や $c$ の文字がイタリック体になっているのが分かるだろうか。
この大原則をベースに以下の3つの例外が存在する（『[LaTeX2ε美文書作成入門]』より引用）。

- 数字はローマン体にする（$\mathit{3.14}$ ではなく $3.14$）
- 複数文字からなる名前はローマン体にする（$sin\,x$ ではなく $\sin x$）
- 単位記号はローマン体にする（$3\,m$ ではなく $3\,\mathrm{m}$）

最初の大原則と併せて，以上の4つのルールを守れば簡単な数式表現は問題ない。
$\\mathrm{\TeX}$ ではこのルールを加味した組版を半自動で行い（万能ではない）， [MathJax] についても $\\mathrm{\TeX}$ と同様の処理ができる。

たとえば数字に関しては普通に `1+2=3` と書けば $1+2=3$ と表示される。
数字部分がローマン体になっているのが分かるだろう。

$\sin$ 関数のように有名な名前については `\sin` のようにあらかじめコマンド[^cmd1] が登録されていて， `\sin x` と書けば $\sin x$ と表示される。
コマンドがないものについては `\mathrm` コマンドで文字をローマン体に変更できる。
たとえば $3\,\mathrm{m}$ と表示するには `3\,\mathrm{m}` と書けばよい。

[^cmd1]: $\\mathrm{\TeX}$ では `\` から始まる文字列をコマンドという。 `\sin` は $\sin$ 関数を表すコマンドである。主な数式コマンドについては『[LaTeX2ε美文書作成入門]』の第5章と第6章に解説があるので，とりあえずは参考になるだろう。

### その他の数式組版規則

実は ISO や JIS で決められている数式の組版規則はもう少し複雑である。

たとえば円周率 $\pi$ は1文字だけど定数の名前なので

$$
    \mathrm{\pi}=3.1415\dots
$$

とローマン体にする。
また以下の式について

$$
    F(x)=\int\,f(x)\,\mathrm{d}x
$$

$\mathrm{d}x$ の d はイタリック体 $d$ ではなくローマン体 $\mathrm{d}$ にする，といった決まりがあるそうだ。

更に数学論文や科学論文では所属する学会や研究会などによってローカル・ルールが存在し，名前についても特定のものについては $\mathrm{Spin}(n)$ ではなく $Spin(n)$ と書きなさい，とか色々あるらしい。
数式を論文などで記述する際は，このような細かいルールにも注意する必要がある。

## 数式は「空き」が重要

少し遡って，先程 $3\,\mathrm{m}$ と表示するのに `3\,\mathrm{m}` と記述したことを覚えているだろうか。
この中の `\,` は「空き」を意味するコマンドである。

明示的な「空き」は見た目を分かりやすく調整するために重要である。

たとえば $\sqrt{2} \times x$ を表現したい場合，素朴に “`\sqrt{2}x`” と記述すると $\sqrt{2}x$ となり， $\sqrt{2x}$ なのか $\sqrt{2}\,x$ なのか見た目でわかり辛い。
そこで明示的に「空き」を入れて “`\sqrt{2}\,x`” と記述することで，はっきりと $\sqrt{2}\,x$ であることを示せる[^sp1]。

[^sp1]: `\sin` のようなコマンドは演算子と同じ扱いなので “`\sin x`” と記述しても結果は $\sin x$ となる。ちなみに $x$ を括弧で囲むと $\sin (x)$ となる。

「空き」を示すコマンドは他にもあるのだが，とりあえず簡単な数式を書く場合には `\,` コマンドだけ覚えておけばOKである。


## [MathJax] の制限事項 {#limit}

残念ながら [MathJax] には $\\mathrm{\TeX}$ にない制限が存在する。
ここに2つほど紹介しよう。

### Neo-Euler にはイタリック体がない

まず [MathJax] で Web フォント `Neo-Euler` を利用するために `HTML-CSS` オプションを以下のように設定する[^chtml1]。

[^chtml1]: Web フォントの指定は `TeX-AMS_HTML` コンフィギュレーションでのみ可能である。 `TeX-AMS_CHTML` では今のところ `"TeX"` 固定になっている。

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  "HTML-CSS": {
    availableFonts: [],
    preferredFont: null,
    webFont: "Neo-Euler"
  }
});
</script>
```

この設定で以下の数式

```html
$$
  \left( \int_0^\infty \frac{\sin x}{\sqrt x}\,\mathrm{d}x \right)^2 =
  \prod_{k=1}^\infty \frac{4k^2}{4k^2-1} \neq \frac{\pi}{2010}
$$
```

を処理すると以下のような表示になる。

{{< fig-img src="https://farm5.staticflickr.com/4361/37112149590_888aff535c_o.png" title="Neo-Euler web fonts"  link="https://www.flickr.com/photos/spiegel/37112149590/" >}}

一方，同じものを $\mathrm{\LaTeX}$ で処理すると以下のようになる。
（「[コンクリートなフォントとか ～Computer Modern の兄弟たち～](http://zrbabbler.sp.land.to/concrete.html)」より引用）

{{< fig-img src="https://farm5.staticflickr.com/4390/36657659994_89e7c9a957_o.png" title="Euler fonts"  link="https://www.flickr.com/photos/spiegel/36657659994/" >}}

両者の違いが分かるだろうか（文字の大きさとかは無視してね）。
ポイントは数字と $\sin$ 関数である。

実は，現行の [MathJax] では `Neo-Euler` にイタリック体が用意されていない[^euler1]。
このため名前も変数も数字も同じ字体で表示されてしまうのだ[^euler2]。

[^euler1]: $\mathrm{\LaTeX}$ での処理結果と比較するかぎり `Neo-Euler` の文字はイタリック体に見えるのだが，[リポジトリに格納されているフォントセット](https://github.com/mathjax/MathJax/tree/master/fonts/HTML-CSS/Neo-Euler)のファイル名を見るとローマン体の扱いになっているようである。いずれにしても Euler 書体は手書き文字が元になっているためローマン体やイタリック体といった字体の区別はないと思われる。
[^euler2]: $\mathrm{\LaTeX}$ でも Euler を Concrete と組み合わせるのが普通みたいなので， [MathJax] でも Web フォントを複数指定できるようにするか Concrete ＋ Euler の組み合わせをひとつのフォントセットとして定義しないとダメだと思う。

{{< fig-quote title="The HTML-CSS output processor" link="http://docs.mathjax.org/en/latest/options/output-processors/HTML-CSS.html" lang="en" >}}
<q>Note that not all mathematical characters are available in all fonts (e.g., Neo-Euler does not include italic characters), so some mathematics may work better in some fonts than in others.</q>
{{< /fig-quote >}}

$\sin$ の見た目だけでも変えたいのであれば，ちょっと裏技的だが以下のようにする。

まず `HTML-CSS` オプションに `mtextFontInherit` を追加し true をセットする。

{{< highlight html "hl_lines=7" >}}
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  "HTML-CSS": {
    availableFonts: [],
    preferredFont: null,
    webFont: "Neo-Euler",
    mtextFontInherit: true
  }
});
</script>
{{< /highlight >}}

そして `\sin` コマンドを `\text{sin}` に書き換える。
以下がその結果である。

{{< fig-img src="https://farm5.staticflickr.com/4432/36698887063_0e218e1705_o.png" title="Neo-Euler web fonts 2"  link="https://www.flickr.com/photos/spiegel/36698887063/" >}}

どうやったかというと `\text{sin}` によって `sin` を（数式フォントではなく）本文のフォントに書き換えたのだ。
`mtextFontInherit` を true にすると `\text` コマンドで指定した文字を本文のフォントにする効果がある。

まぁ，これで見た目はそれっぽくなったが，やはり無理矢理な感じが強いので， [MathJax] で `Neo-Euler` を使うのは（今のところ）お薦めできない。 

### 不等号に注意

$\mathrm{\TeX}$ 記法では不等号記号はそのまま `<` や `>` 文字を使うのだが， HTML ではこれらの文字は要素タグ（`<body>` など）を指す記号として使われるため，そのまま数式に使うとブラウザによっては誤動作を起こす可能性がある。
そのため [MathJax] では不等号記号を示すコマンドが用意されている。

たとえば $a \lt b$ は `$a \lt b$` と記述する。
同様に $a \gt b$ は `$a \gt b$` と記述する。

`\lt` や `\gt` は標準の $\mathrm{\TeX}$ 記法には存在しないコマンドなので，数式を流用する場合には注意が必要である。
あるいは $\mathrm{\LaTeX}$ 側で

```text
\newcommand{\lt}{<}
\newcommand{\gt}{>}
```

のようにコマンドを定義しておく手もある。

## TeX 記法のメリットは可搬性にある {#benefit}

こうした細かい面倒がありつつも $\mathrm{\TeX}$ 記法を使う理由は，記述の可搬性にある。
数式を $\mathrm{\TeX}$ 記法で書けば，それはそのまま他のブログや論文に流用できる。
これは MathML のような見た目だけを制御する記述形式にはない利点と言える[^acm]。

[^acm]: [本ブログ]では触れないが， [MathJax] は [AsciiMath] にも対応している。 [AsciiMath] にも $\mathrm{\TeX}$ 記法と同様のメリットがある。 [AsciiMath] と $\mathrm{\TeX}$ 記法の比較記事は時々目にするが，個人的には使いやすい方を使えばいいと思う。

というわけで， [MathJax] で積極的に数式を書いて読んで流用していくのがいいと思う。

## ブックマーク {#bookmark}

- [斜体とイタリック体 – Pineapple Blog](https://pineapple.blog/%E6%96%9C%E4%BD%93%E3%81%A8%E3%82%A4%E3%82%BF%E3%83%AA%E3%83%83%E3%82%AF%E4%BD%93-68dda513eca2)
- [MathJaxでEuler(オイラー)フォントを使ったときの不具合](http://www.math.sci.hokudai.ac.jp/~yano/memo/mathjax_euler.html) ： `mtextFontInherit` には true または false が入るので，文字列 `"false"` をセットするのは間違い。おそらく文字列を内部で無理やり true に評価してるんだと思う。この辺は流石 JavaScript というところか（笑）
- [MathJaxによる数式表示](https://oku.edu.mie-u.ac.jp/~okumura/javascript/mathjax.html)

[本ブログ]: / "text.Baldanders.info"
[MathJax]: https://www.mathjax.org/
[AsciiMath]: http://asciimath.org/
[前回]: {{< relref "remark/2017/09/getting-started-mathjax-1.md" >}} "ちょこっと MathJax： 初期設定"
[LaTeX2ε美文書作成入門]: http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/ "Amazon | [改訂第7版]LaTeX2ε美文書作成入門 | 奥村 晴彦, 黒木 裕介 通販"

## 参考図書 {#books}

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51E5K7B53aL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/">[改訂第7版]LaTeX2ε美文書作成入門</a></dt><dd>奥村 晴彦 黒木 裕介 </dd><dd>技術評論社 2017-01-24</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798118141/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798118141.09._SCTHUMBZZZ_.jpg"  alt="LaTeX2e辞典~用法・用例逆引きリファレンス (DESKTOP REFERENCE)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4535558752/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4535558752.09._SCTHUMBZZZ_.jpg"  alt="公共政策入門 ミクロ経済学的アプローチ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320112415/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320112415.09._SCTHUMBZZZ_.jpg"  alt="Rで楽しむ統計 (Wonderful R 1)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298550/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298550.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.5"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797391383/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797391383.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/積分を見つめて (数学ガールの秘密ノートシリーズ)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298569/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298569.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.6"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798115363/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798115363.09._SCTHUMBZZZ_.jpg"  alt="独習 LaTeX2ε"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4785315717/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4785315717.09._SCTHUMBZZZ_.jpg"  alt="具体例から学ぶ 多様体"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774193046/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774193046.09._SCTHUMBZZZ_.jpg"  alt="【改訂第3版】基礎からわかる情報リテラシー"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4768704700/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4768704700.09._SCTHUMBZZZ_.jpg"  alt="はじめて学ぶリー群 ―線型代数から始めよう"  /></a> </p>
<p class="description">ついに第7版が登場。紙の本で買って常に側に置いておくのが吉。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-27">2017-09-27</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
