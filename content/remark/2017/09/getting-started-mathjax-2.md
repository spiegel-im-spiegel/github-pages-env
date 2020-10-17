+++
title = "ちょこっと MathJax： 基本的な数式表現"
date =  "2017-09-28T22:25:54+09:00"
description = "今回は基本的な数式の書き方を説明していこう。"
tags = [ "math", "tex", "mathjax", "javascript" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

さて，初期設定は[前回]で完了したので，今回は基本的な数式の書き方を説明していこう。

1. [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})
2. [ちょこっと MathJax： 基本的な数式表現]({{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}}) ← イマココ
3. [ちょこっと MathJax： インライン数式と別行立て数式]({{< ref "/remark/2017/10/getting-started-mathjax-3.md" >}})
4. [ちょこっと MathJax 番外編： mathcomp パッケージの代替え]({{< ref "/remark/2017/12/mathcomp-in-mathjax.md" >}})

## 数式表現の「お約束」 {#rule}

初っ端からナニだが，数式を書く際には昔からある「お約束」が存在する。

まず最初に，大原則として「数式で使う文字はイタリック体[^it1]で書く」というのがある。
たとえば，こんな感じだ。

[^it1]: $\\mathrm{\TeX}$ ではイタリック体（italics）と斜体（slanted）は異なる字体である。また斜体はしばしばローマン体（Roman）から合成される場合がある。

{{< div-box >}}
エネルギーと質量には $$E=mc^2$$ の関係がある。
{{< /div-box >}}

$E$ や $m$ や $c$ の文字がイタリック体になっているのが分かるだろうか。
この大原則をベースに以下の3つの例外が存在する（『[LaTeX2ε美文書作成入門]』より引用）。

- 数字はローマン体にする（$\mathit{3.14}$ ではなく $3.14$）
- 複数文字からなる名前はローマン体にする（$sin\\,x$ ではなく $\sin x$）
- 単位記号はローマン体にする（$3\\,m$ ではなく $3\\,\mathrm{m}$）

最初の大原則と併せて，以上の4つのルールを守れば簡単な数式表現は問題ない。
$\\mathrm{\TeX}$ ではこのルールを加味した組版を半自動で行い（万能ではない）， [MathJax] についても $\\mathrm{\TeX}$ と同様の処理ができる。

たとえば数字に関しては普通に `1+2=3` と書けば $1+2=3$ と表示される。
数字部分がローマン体になっているのが分かるだろう。

$\sin$ 関数のように有名な名前については `\sin` のようにあらかじめコマンド[^cmd1] が登録されていて， `\sin x` と書けば $\sin x$ と表示される。
コマンドがないものについては `\mathrm` コマンドで文字をローマン体に変更できる。
たとえば $3\\,\mathrm{m}$ と表示するには `3\\,\mathrm{m}` と書けばよい。

[^cmd1]: $\\mathrm{\TeX}$ では `\` から始まる文字列をコマンド（『[LaTeX2ε美文書作成入門]』では「命令」）という。 `\sin` は $\sin$ 関数を表すコマンドである。主な数式コマンドについては『[LaTeX2ε美文書作成入門]』の第5章と第6章に解説があるので，とりあえずは参考になるだろう。

### その他の数式組版規則

実は ISO や JIS で決められている数式の組版規則はもう少し複雑である。

たとえば円周率 $\mathit{\pi}$ は1文字だけど定数の名前なので

$$
    \mathrm{\pi}=3.1415\dots
$$

とローマン体にする（TeX フォントだと違いが分からないが）。
また以下の式について

$$
    F(x)=\int\\,f(x)\\,\mathrm{d}x
$$

$\mathrm{d}x$ の d はイタリック体 $d$ ではなくローマン体 $\mathrm{d}$ にする，といった決まりがあるそうだ。

更に数学論文や科学論文では所属する学会や研究会などによってローカル・ルールが存在し，名前についても特定のものについては $\mathrm{Spin}(n)$ ではなく $Spin(n)$ と書きなさい，とか色々あるらしい。
数式を論文などで記述する際は，このような細かいルールにも注意する必要がある。

## 数式は「空き」が重要

少し遡って，先程 $3\\,\mathrm{m}$ と表示するのに `3\\,\mathrm{m}` と記述したことを覚えているだろうか。
この中の `\\,` は「空き」を意味するコマンドである。

明示的な「空き」は見た目を分かりやすく調整するために重要である。

たとえば $\sqrt{2} \times x$ を表現したい場合，素朴に “`\sqrt{2}x`” と記述すると $\sqrt{2}x$ となり， $\sqrt{2x}$ なのか $\sqrt{2}\\,x$ なのか見た目でわかり辛い。
そこで明示的に「空き」を入れて “`\sqrt{2}\\,x`” と記述することで，はっきりと $\sqrt{2}\\,x$ であることを示せる[^sp1]。

[^sp1]: `\sin` のようなコマンドは演算子と同じ扱いなので “`\sin x`” と記述しても結果は $\sin x$ となる。ちなみに $x$ を括弧で囲むと $\sin (x)$ となる。

「空き」を示すコマンドは他にもあるのだが，とりあえず簡単な数式を書く場合には `\\,` コマンドだけ覚えておけばOKである。


## [MathJax] の制限事項 {#limit}

残念ながら [MathJax] には $\\mathrm{\TeX}$ にない制限が存在する。

<!-- Mathjax v3 は今のところ Neo-Euler をサポートしないためコメントアウトしておく
### Neo-Euler には斜体がない

数式好きの方に人気がある（らしい） Euler 書体は [MathJax] では Web フォント `Neo-Euler` として提供されている。
`Neo-Euler` を利用するためには `HTML-CSS` オプションを以下のように設定する[^chtml1]。

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
（「[コンクリートなフォントとか ～Computer Modern の兄弟たち～](http://zrbabbler.sp.land.to/concrete.html)」より引用）

```html
$$
  \left( \int_0^\infty \frac{\sin x}{\sqrt x}\\,\mathrm{d}x \right)^2 =
  \prod_{k=1}^\infty \frac{4k^2}{4k^2-1} \neq \frac{\pi}{2010}
$$
```

を処理すると以下のような表示になる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37112149590_m.png" title="Neo-Euler web fonts" link="https://photo.baldanders.info/flickr/37112149590/" >}}

一方，同じものを $\mathrm{\LaTeX}$ で Euler ＋ Concrete を指定して 処理すると以下のようになる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/36657659994_m.png" title="Euler fonts" link="https://photo.baldanders.info/flickr/36657659994/" >}}

両者の違いが分かるだろうか（文字の大きさとかは無視してね）。
ポイントは数字と $\sin$ 関数である。

実は，現行の `Neo-Euler` には文字と記号の一部に斜体（slant）が用意されていない[^euler1]。
このせいで（通常ルールとは異なり）数式表現が「立体（upright）」になるため，名前も変数も数字も同じに見えてしまうのだ（厳密には違うけど）。

[^euler1]: [MathJax] のドキュメントには「イタリック体がない」と書かれているが，厳密には「斜体」のことを言っているようだ。そもそも Euler 書体は手書き文字が元になっていて見た目はイタリック体っぽいデザインをしているため，「イタリック体がない」という言い方は誤解を招く。

{{< fig-quote title="The HTML-CSS output processor" link="http://docs.mathjax.org/en/latest/options/output-processors/HTML-CSS.html" lang="en" >}}
<q>Note that not all mathematical characters are available in all fonts (e.g., Neo-Euler does not include italic characters), so some mathematics may work better in some fonts than in others.</q>
{{< /fig-quote >}}

$\mathrm{\LaTeX}$ っぽく $\sin$ の見た目だけでも変えたいのであれば，ちょっと裏技的だが以下のようにする。

まず `HTML-CSS` オプションに `mtextFontInherit` を追加し true をセットする（`mtextFontInherit` については[前回]の記事を参照）。

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

{{< fig-img src="https://photo.baldanders.info/flickr/image/36698887063_m.png" title="Neo-Euler web fonts 2" link="https://photo.baldanders.info/flickr/36698887063/" >}}

どうやったかというと `\text{sin}` によって `sin` を（数式フォントではなく）本文のフォントに書き換えたのだ。

まぁ，これで見た目はそれっぽくなったが，やはり無理矢理な感じが強いので， [MathJax] で `Neo-Euler` を使うのは（今のところ）お薦めできない[^euler2]。 

[^euler2]: $\mathrm{\LaTeX}$ でも Euler を Concrete と組み合わせるのが普通みたいなので， [MathJax] でも Web フォントを複数指定できるようにするか Concrete ＋ Euler の組み合わせをひとつのフォントセットとして定義しないとダメだと思う。
-->

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
数式を $\mathrm{\TeX}$ 記法で書けば，それはそのまま他のブログや論文に流用できるし，式の構成をちょこちょこっと変えて「翻案」するのも簡単である。
これは MathML のような見た目だけを制御する記述形式にはない利点と言える[^acm]。

[^acm]: [本ブログ]では触れないが， [MathJax] は [AsciiMath] にも[対応](https://docs.mathjax.org/en/latest/input/asciimath.html "AsciiMath Support — MathJax 3.0 documentation")している。 [AsciiMath] にも $\mathrm{\TeX}$ 記法と同様のメリットがある。 [AsciiMath] と $\mathrm{\TeX}$ 記法の比較記事は時々目にするが，個人的には使いやすい方を使えばいいと思う。私が $\mathrm{\TeX}$， というか $\mathrm{\LaTeX}$ を気に入っているのは，見た目の美麗さだけを追求する（出来の悪いペイント・ツールのごとき）ことではなく，かといって届かない理想ばかりを追いかける潔癖症なシステムでもなく，賢さと泥臭さを上手く融合させている点にある。そうでなければ，ひとつのシステムが（中身が変貌しつつも）40年近くも継続して使われ続ける筈がない。

というわけで， [MathJax] で積極的に数式を書いて読んで流用していくのがいいと思う。

## ブックマーク {#bookmark}

- [斜体とイタリック体 – Pineapple Blog](https://pineapple.blog/%E6%96%9C%E4%BD%93%E3%81%A8%E3%82%A4%E3%82%BF%E3%83%AA%E3%83%83%E3%82%AF%E4%BD%93-68dda513eca2)
- [MathJaxによる数式表示](https://oku.edu.mie-u.ac.jp/~okumura/javascript/mathjax.html)
- [数学記号記法一覧](https://zenn.dev/wsuzume/articles/b0b3a51cac5d7fe4555b)

<!--
- [MathJaxでEuler(オイラー)フォントを使ったときの不具合](http://www.math.sci.hokudai.ac.jp/~yano/memo/mathjax_euler.html) ： `mtextFontInherit` には true または false が入るので，文字列 `"false"` をセットするのは間違い。おそらく文字列を内部で無理やり true に評価してるんだと思う。この辺は流石 JavaScript というところか（笑）

- [数式用フォントで遊ぶ]({{< ref "/remark/2017/10/math-fonts.md" >}}) : $\mathrm{TeX}$ における数式表現についてフォントを中心に書いてみた
-->

[本ブログ]: {{< rlnk "/" >}} "text.Baldanders.info"
[MathJax]: https://www.mathjax.org/ "MathJax | Beautiful math in all browsers."
[AsciiMath]: http://asciimath.org/
[前回]: {{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}} "ちょこっと MathJax： 初期設定"
[LaTeX2ε美文書作成入門]: https://www.amazon.co.jp/dp/4774187054?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon | [改訂第7版]LaTeX2ε美文書作成入門 | 奥村 晴彦, 黒木 裕介 通販"

## 参考図書 {#books}

{{% review-paapi "4774187054" %}} <!-- [改訂第7版]LaTeX2ε美文書作成入門 -->
