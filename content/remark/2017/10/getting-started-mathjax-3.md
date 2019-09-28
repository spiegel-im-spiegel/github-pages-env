+++
title = "ちょこっと MathJax： インライン数式と別行立て数式"
date =  "2017-10-27T17:24:58+09:00"
description = "MathJax の数式の表示には2種類ある。"
tags        = [ "math", "tex", "mathjax", "javascript" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

[前回]からだいぶ間があいたけど，そろそろ続きを。

1. [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})
2. [ちょこっと MathJax： 基本的な数式表現]({{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}})
3. [ちょこっと MathJax： インライン数式と別行立て数式]({{< ref "/remark/2017/10/getting-started-mathjax-3.md" >}}) ← イマココ
4. [ちょこっと MathJax 番外編： mathcomp パッケージの代替え]({{< ref "/remark/2017/12/mathcomp-in-mathjax.md" >}})

{{< div-box type="md" >}}
## 2019-09-28 追記

[MathJax v3 がリリース]({{< ref "/release/2019/09/mathjax-v3-is-released.md" >}} "MathJax v3 がリリースされていた")された。
v2.7.x までとは大きく変わったためこの記事も改訂する予定である。
{{< /div-box >}}

## インライン数式と別行立て数式

さて，[初期設定]の回でも少し言及したが，[MathJax] の数式の表示には2種類ある。
まず本文に埋め込まれる形で表示される数式。
これを「インライン数式（in-line math）」と呼ぶ。
インライン数式は `\(...\)` で囲んで記述する。

たとえば

```html
エネルギーと質量には \(E=mc^2\) の関係がある。
```

と記述すると

> エネルギーと質量には \\(E=mc^2\\) の関係がある。

のように表示される。

もうひとつは本文とは行を改めて表示される数式。
これを「別行立て数式（displayed equations）」と呼ぶ[^book1]。
別行立て数式は `$$...$$` または `\[...\]` で囲んで記述する。

[^book1]: 「インライン数式」「別行立て数式」という言い回しは『[LaTeX2ε美文書作成入門]』を踏襲している。ちなみに “in-line math”, “displayed equations” は [MathJax] のドキュメントでの言い回し。

たとえば

```html
エネルギーと質量には \[E=mc^2\] の関係がある。
```

と記述すると

{{< fig-quote >}}
エネルギーと質量には \[E=mc^2\] の関係がある。
{{< /fig-quote >}}

のように表示される。

$\mathrm{\LaTeX}$ ではインライン数式については `\(...\)` ではなく `$...$` を使う。
なので [`tex2jax`](http://docs.mathjax.org/en/latest/options/preprocessors/tex2jax.html "The tex2jax Preprocessor — MathJax 2.7 documentation") オプションをいじって `$...$` を有効にしてしまおう。

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  tex2jax: {
    inlineMath: [['$','$'], ['\\(','\\)']],
    processEscapes: true
  }
});
</script>
```

オプション設定について詳しくは[初期設定]の回を参照のこと[^jm1]。

[^jm1]: [初期設定]の回でも言及しているが `processEscapes` はパラグラフ `<p>...</p>` 内でのみ有効なようだ。つまり `processEscapes` が有効な状態では，パラグラフ内の `\(...\)` や `\[...\]` がエスケープされてしまうので注意すること。

## Textstyle と Displaystyle

先程の例で挙げた数式 $E=mc^2$ ではインライン数式と別行立て数式表示に（表示位置以外は）さしたる違いは見られないが，数式によっては表示が明らかに異なる場合がある。
たとえば `\sum_{k=1}^n a_k` という記述について，インライン数式なら

> たとえば $\sum_{k=1}^n a_k$ という記述について

となるが，別行立て数式では

{{< fig-quote >}}
たとえば \[\sum_{k=1}^n a_k\] という記述について
{{< /fig-quote >}}

となる。
インライン数式では，できるだけ数式が本文からはみ出ないように自動的に調節してくれているのである。
これを意図的に変えるには `\textstyle`, `\displaystyle`, `\limits`, `\nolimits` といったコマンドを使う。
以下に各コマンドを使った場合の表示の違いを挙げる。

| $\mathrm{TeX}$ 記法 | 表示 |
|:--------------------|:-----|
| `\textstyle\sum_{k=1}^n a_k`             | $\textstyle\sum_{k=1}^n a_k$ |
| `\displaystyle\sum_{k=1}^n a_k`          | $\displaystyle\sum_{k=1}^n a_k$ |
| `\textstyle\sum\limits_{k=1}^n a_k`      | $\textstyle\sum\limits_{k=1}^n a_k$ |
| `\displaystyle\sum\nolimits_{k=1}^n a_k` | $\displaystyle\sum\nolimits_{k=1}^n a_k$ |

## 分数表記

もうひとつ。
インライン数式と別行立て数式で気をつけるべきなのが分数の表記である。
たとえば `y=1/x` を $\mathrm{TeX}$ 記法で記述する場合は `y=\frac{1}{x}` と書くが，インライン数式の場合は

> たとえば `y=1/x` は $y=\frac{1}{x}$ と書く

となり，別行立て数式の場合は

{{< fig-quote >}}
たとえば <code>y=1/x</code> は \[y=\frac{1}{x}\] と書く
{{< /fig-quote >}}

となる。
しかしインライン数式の場合，これでは文字が小さくなりすぎるし，結局は本文からハミ出してしまう。
なのでインライン数式では `\frac` を使うのではなくそのまま `$y=1/x$` と書くのが良いとされている。

> たとえば `y=1/x` は $y=1/x$ と書く

なお分数でも textstyle と displaystyle に相当する `\tfrac` と `\dfrac` がある。

| $\mathrm{TeX}$ 記法 | 表示 |
|:--------------------|:----:|
| `y=\tfrac{1}{x}` | $y=\tfrac{1}{x}$ |
| `y=\dfrac{1}{x}` | $y=\dfrac{1}{x}$ |

ちなみに [MathJax] でも連分数（`\cfrac`）が使える[^bk2]。

[^bk2]: 出典は『[LaTeX2ε美文書作成入門]』より。

```html
\[
b_0 + \cfrac{c_1}{b_1 +
  \cfrac{c_2}{b_2 +
  \cfrac{c_3}{b_3 +
  \cfrac{c_4}{b_4 + \cdots}}}}
\]
```

{{< fig-quote >}}
\[
b_0 + \cfrac{c_1}{b_1 +
  \cfrac{c_2}{b_2 +
  \cfrac{c_3}{b_3 +
  \cfrac{c_4}{b_4 + \cdots}}}}
\]
{{< /fig-quote >}}

流石にこれはインライン数式じゃ無理だよね（笑）

## インライン数式で高さを揃える。

インライン数式では `$\sqrt{g}$` と `$\sqrt{h}$` のように数式によって高さが不揃いになるものがある。

> インライン数式では $\sqrt{g}$ と $\sqrt{h}$ の高さが不揃い

高さを揃えるには `\mathstrut` コマンドを使って以下のように記述する。

```html
\mathstrut コマンドを使って $\sqrt{\mathstrut g}$ と $\sqrt{\mathstrut h}$ の高さを揃えてみる
```

> `\mathstrut` コマンドを使って $\sqrt{\mathstrut g}$ と $\sqrt{\mathstrut h}$ の高さを揃えてみる

さらに `\smash` コマンドとも組み合わせて

```html
\smash コマンドも使って $\sqrt{\smash[b]{\mathstrut g}}$ と $\sqrt{\smash[b]{\mathstrut h}}$ を揃えてみる
```

> `\smash` コマンドも使って $\sqrt{\smash[b]{\mathstrut g}}$ と $\sqrt{\smash[b]{\mathstrut h}}$ を揃えてみる

とするともうちょっとだけいい感じになるようである[^smsh1]。
折角なのでマクロに組み込んでしまおう。

[^smsh1]: これも『[LaTeX2ε美文書作成入門]』を参考にした。いつもお世話になっています。

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  TeX: {
    Macros: {
      ssqrt: ['\\sqrt{\\smash[b]{\\mathstrut #1}}', 1]
    }
 }
});
</script>
```

これで以下のように書けば

```html
\smash コマンドも使って $\ssqrt{g}$ と $\ssqrt{h}$ を揃えてみる
```

同じ結果が得られる。

> `\smash` コマンドも使って $\ssqrt{g}$ と $\ssqrt{h}$ を揃えてみる


## 別行立て数式に番号を振る

別行立て数式では数式に番号を振ることができる。たとえば `\[E=mc^2\]` に (a) をいう番号を振りたければ `\tag` コマンドを使って `\[E=mc^2 \tag{a}\]` とする。

{{< fig-quote >}}
\[E=mc^2 \tag{a}\]
{{< /fig-quote >}}

この番号にはラベル `\label` を付けて参照することができる。
たとえば `\[E=mc^2 \tag{b}\label{eq-b}\]` としておいて

{{< fig-quote >}}
\[E=mc^2 \tag{b}\label{eq-b}\]
{{< /fig-quote >}}

このラベルを参照するには `\eqref{eq-b}` でこのように →\eqref{eq-b}← できる（`$...$` で囲まなくてもよい）[^anc1]。

[^anc1]: 数式参照用の id は [MathJax] で動的に生成されているので，ページ外からの参照はお勧めできない。

### 数式に通し番号を振る

ページ内で通し番号を振りたい場合は，まず [`TeX`](http://docs.mathjax.org/en/latest/options/input-processors/TeX.html "The TeX input processor — MathJax 2.7 documentation") オプションで設定を変更する（オプション設定については[初期設定]の回を参照）。

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  TeX: {
    equationNumbers: { autoNumber: "AMS" },
 }
});
</script>
```

実際に番号を振るには数式を `\[...\]` で囲むのではなく `\begin{equation}...\end{equation}` で囲む。
自動で番号が振られるので `\tag` コマンドは不要である。

```html
エネルギーと質量には
\begin{equation}
  E=mc^2 \label{eq-1st}
\end{equation}
の関係がある。
```

{{< fig-quote >}}
エネルギーと質量には
\begin{equation}
  E=mc^2 \label{eq-1st}
\end{equation}
の関係がある。
{{< /fig-quote >}}

参照も同様に →\eqref{eq-1st}← できる。

余談だが `\begin` と `\end` で囲まれた領域を $\mathrm{\LaTeX}$  では「環境」と呼ぶ。
`\begin{foober}...\end{foober}` なら「`foobar` 環境」と呼んだりする。

`equation` にアスタリスクを付けた `equation*` 環境では，環境内の数式を自動採番の対象外にできる[^math1]。

[^math1]: 実は `\[...\]` は `equation*` 環境と等価である。

```html
エネルギーと質量には
\begin{equation*}
  E=mc^2
\end{equation*}
の関係がある。
```

{{< fig-quote >}}
エネルギーと質量には
\begin{equation*}
  E=mc^2
\end{equation*}
の関係がある。
{{< /fig-quote >}}

`equation` 環境以外にも `align` 環境や `gather` 環境なども自動採番の対象となる（自動採番を無効にするアスタリスクも有効）。
`align` 環境などで特定の式に番号を振りたくない場合は `\notag` コマンドを使って

```html
\begin{align}
  (a+b)^2 &= a^2+2ab+b^2 \\
  (a-b)^2 &= a^2-2ab+b^2  \notag \\
  (a+b)^3 &= a^3+3a^{2}b+3ab^2+b^3
\end{align}
```

{{< fig-quote >}}
\begin{align}
  (a+b)^2 &= a^2+2ab+b^2 \\
  (a-b)^2 &= a^2-2ab+b^2  \notag \\
  (a+b)^3 &= a^3+3a^{2}b+3ab^2+b^3
\end{align}
{{< /fig-quote >}}

とすればよい。
また複数の数式を `split` 環境や `aligned` 環境で囲むことで複数の数式の塊に一つの番号を振ることもできる[^math2]。

[^math2]: `split` 環境は自動採番の対象外である。

```html
\begin{equation}
  \begin{split}
    (a+b)^2 &= a^2+2ab+b^2 \\
    (a-b)^2 &= a^2-2ab+b^2 \\
    (a+b)^3 &= a^3+3a^{2}b+3ab^2+b^3
  \end{split}
\end{equation}
```

{{< fig-quote >}}
\begin{equation}
\begin{split}
  (a+b)^2 &= a^2+2ab+b^2 \\
  (a-b)^2 &= a^2-2ab+b^2 \\
  (a+b)^3 &= a^3+3a^{2}b+3ab^2+b^3
\end{split}
\end{equation}
{{< /fig-quote >}}

[MathJax]: https://www.mathjax.org/
[初期設定]: {{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}} "ちょこっと MathJax： 初期設定"
[前回]: {{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}} "ちょこっと MathJax： 基本的な数式表現"
[LaTeX2ε美文書作成入門]: https://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/ "Amazon | [改訂第7版]LaTeX2ε美文書作成入門 | 奥村 晴彦, 黒木 裕介 通販"

## 参考図書 {#books}

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%94%B9%E8%A8%82%E7%AC%AC7%E7%89%88-LaTeX2%CE%B5%E7%BE%8E%E6%96%87%E6%9B%B8%E4%BD%9C%E6%88%90%E5%85%A5%E9%96%80-%E5%A5%A5%E6%9D%91-%E6%99%B4%E5%BD%A6/dp/4774187054?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4774187054"><img src="https://images-fe.ssl-images-amazon.com/images/I/51E5K7B53aL._SL160_.jpg" width="127" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%94%B9%E8%A8%82%E7%AC%AC7%E7%89%88-LaTeX2%CE%B5%E7%BE%8E%E6%96%87%E6%9B%B8%E4%BD%9C%E6%88%90%E5%85%A5%E9%96%80-%E5%A5%A5%E6%9D%91-%E6%99%B4%E5%BD%A6/dp/4774187054?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4774187054">[改訂第7版]LaTeX2ε美文書作成入門</a></dt>
    <dd>奥村 晴彦, 黒木 裕介</dd>
    <dd>技術評論社 2017-01-24</dd>
    <dd>大型本</dd>
    <dd>4774187054 (ASIN), 9784774187051 (EAN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ついに第7版が登場。紙の本で買って常に側に置いておくのが吉。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-27">2017-09-27</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
