+++
title = "ちょこっと MathJax： インライン数式と別行立て数式"
date =  "2017-10-27T17:24:58+09:00"
description = "MathJax の数式の表示には2種類ある。"
tags = [ "math", "tex", "mathjax", "javascript" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

[前回]からだいぶ間があいたけど，そろそろ続きを。

1. [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})
2. [ちょこっと MathJax： 基本的な数式表現]({{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}})
3. [ちょこっと MathJax： インライン数式と別行立て数式]({{< ref "/remark/2017/10/getting-started-mathjax-3.md" >}}) ← イマココ
4. [ちょこっと MathJax 番外編： mathcomp パッケージの代替え]({{< ref "/remark/2017/12/mathcomp-in-mathjax.md" >}})

## インライン数式と別行立て数式

さて，[初期設定]の回でも少し言及したが，[MathJax] の数式の表示には2種類ある。
まず本文に埋め込まれる形で表示される数式。
これを「インライン数式（in-line math）」と呼ぶ。
インライン数式は `\(...\)` で囲んで記述する。

たとえば

```tex
エネルギーと質量には \(E=mc^2\) の関係がある。
```

と記述すると

{{< div-box >}}
エネルギーと質量には \(E=mc^2\) の関係がある。
{{< /div-box >}}

のように表示される。

もうひとつは本文とは行を改めて表示される数式。
これを「別行立て数式（displayed equations）」と呼ぶ[^book1]。
別行立て数式は `$$...$$` または `\[...\]` で囲んで記述する。

[^book1]: 「インライン数式」「別行立て数式」という言い回しは『[LaTeX2ε美文書作成入門]』を踏襲している。ちなみに “in-line math”, “displayed equations” は [MathJax] のドキュメントでの言い回し。

たとえば

```tex
エネルギーと質量には \[E=mc^2\] の関係がある。
```

と記述すると

{{< div-box >}}
エネルギーと質量には \[E=mc^2\] の関係がある。
{{< /div-box >}}

のように表示される。

$\mathrm{\LaTeX}$ ではインライン数式については `\(...\)` ではなく `$...$` を使う。
なので設定をいじって `$...$` を有効にしてしまおう。

```js
MathJax = {
  tex: {
    inlineMath: [['$', '$'], ['\\(', '\\)']],
    processEscapes: true // use \$ to produce a literal dollar sign (true is default)
  }
};
```

設定について詳しくは[初期設定]の回を参照のこと[^jm1]。

[^jm1]: [初期設定]の回でも言及しているが `processEscapes` はパラグラフ `<p>...</p>` 内でのみ有効なようだ。つまり `processEscapes` が有効な状態では，パラグラフ内の `\(...\)` や `\[...\]` がエスケープされてしまうので注意すること。

## Textstyle と Displaystyle

先程の例で挙げた数式 $E=mc^2$ ではインライン数式と別行立て数式表示に（表示位置以外は）さしたる違いは見られないが，数式によっては表示が明らかに異なる場合がある。
たとえば `\sum_{k=1}^n a_k` という記述について，インライン数式なら

{{< div-box >}}
たとえば $\sum_{k=1}^n a_k$ という記述について
{{< /div-box >}}

となるが，別行立て数式では

{{< div-box >}}
たとえば \[\sum_{k=1}^n a_k\] という記述について
{{< /div-box >}}

となる。

インライン数式では，できるだけ数式が本文からはみ出ないように自動的に調節してくれているのである。
これを意図的に変えるには `\textstyle`, `\displaystyle`, `\limits`, `\nolimits` といったコマンドを使う。
以下に各コマンドを使った場合の表示の違いを挙げる。

| $\mathrm{TeX}$ 記法                      | 表示                                     |
|:---------------------------------------- |:---------------------------------------- |
| `\textstyle\sum_{k=1}^n a_k`             | $\textstyle\sum_{k=1}^n a_k$             |
| `\displaystyle\sum_{k=1}^n a_k`          | $\displaystyle\sum_{k=1}^n a_k$          |
| `\textstyle\sum\limits_{k=1}^n a_k`      | $\textstyle\sum\limits_{k=1}^n a_k$      |
| `\displaystyle\sum\nolimits_{k=1}^n a_k` | $\displaystyle\sum\nolimits_{k=1}^n a_k$ |

## 分数表記

もうひとつ。
インライン数式と別行立て数式で気をつけるべきなのが分数の表記である。
たとえば `y=1/x` を $\mathrm{TeX}$ 記法で記述する場合は `y=\frac{1}{x}` と書くが，インライン数式の場合は

{{< div-box >}}
たとえば <code>y=1/x</code> は $y=\frac{1}{x}$ と書く
{{< /div-box >}}

となり，別行立て数式の場合は

{{< div-box >}}
たとえば <code>y=1/x</code> は \[y=\frac{1}{x}\] と書く
{{< /div-box >}}

となる。
しかしインライン数式の場合，これでは文字が小さくなりすぎるし，結局は本文からハミ出してしまう。
なのでインライン数式では `\frac` を使うのではなくそのまま `$y=1/x$` と書くのが良いとされている。

{{< div-box >}}
たとえば <code>y=1/x</code> は $y=1/x$ と書く
{{< /div-box >}}

なお分数でも textstyle と displaystyle に相当する `\tfrac` と `\dfrac` がある。

| $\mathrm{TeX}$ 記法 |       表示       |
|:------------------- |:----------------:|
| `y=\tfrac{1}{x}`    | $y=\tfrac{1}{x}$ |
| `y=\dfrac{1}{x}`    | $y=\dfrac{1}{x}$ |

ちなみに [MathJax] でも連分数（`\cfrac`）が使える[^bk2]。

[^bk2]: 出典は『[LaTeX2ε美文書作成入門]』より。

```tex
\[
b_0 + \cfrac{c_1}{b_1 +
  \cfrac{c_2}{b_2 +
  \cfrac{c_3}{b_3 +
  \cfrac{c_4}{b_4 + \cdots}}}}
\]
```

{{< fig-math class="box" >}}
\[
b_0 + \cfrac{c_1}{b_1 +
  \cfrac{c_2}{b_2 +
  \cfrac{c_3}{b_3 +
  \cfrac{c_4}{b_4 + \cdots}}}}
\]
{{< /fig-math >}}

流石にこれはインライン数式じゃ無理だよね（笑）

## インライン数式で高さを揃える。

インライン数式では `$\sqrt{g}$` と `$\sqrt{h}$` のように数式によって高さが不揃いになるものがある。

{{< div-box >}}
インライン数式では $\sqrt{g}$ と $\sqrt{h}$ の高さが不揃い
{{< /div-box >}}

高さを揃えるには `\mathstrut` コマンドを使って以下のように記述する。

```text
\mathstrut コマンドを使って $\sqrt{\mathstrut g}$ と $\sqrt{\mathstrut h}$ の高さを揃えてみる
```

{{< div-box >}}
<code>\mathstrut</code> コマンドを使って $\sqrt{\mathstrut g}$ と $\sqrt{\mathstrut h}$ の高さを揃えてみる
{{< /div-box >}}

さらに `\smash` コマンドとも組み合わせて

```text
\smash コマンドも使って $\sqrt{\smash[b]{\mathstrut g}}$ と $\sqrt{\smash[b]{\mathstrut h}}$ を揃えてみる
```

{{< div-box >}}
<code>\smash</code> コマンドも使って $\sqrt{\smash[b]{\mathstrut g}}$ と $\sqrt{\smash[b]{\mathstrut h}}$ を揃えてみる
{{< /div-box >}}

とするともうちょっとだけいい感じになるようである[^smsh1]。
折角なのでマクロに組み込んでしまおう。

[^smsh1]: これも『[LaTeX2ε美文書作成入門]』を参考にした。いつもお世話になっています。

```js
MathJax = {
  tex: {
    macros: {
      ssqrt: ['\\sqrt{\\smash[b]{\\mathstrut #1}}', 1]
    }
  }
};
```

これで以下のように書けば

```text
\smash コマンドも使って $\ssqrt{g}$ と $\ssqrt{h}$ を揃えてみる
```

同じ結果が得られる。

{{< div-box >}}
<code>\smash</code> コマンドも使って $\ssqrt{g}$ と $\ssqrt{h}$ を揃えてみる
{{< /div-box >}}


## 別行立て数式に番号を振る

別行立て数式では数式に番号を振ることができる。たとえば `\[E=mc^2\]` に (a) をいう番号を振りたければ `\tag` コマンドを使って `\[E=mc^2 \tag{a}\]` とする。

{{< fig-math class="box" >}}
\[E=mc^2 \tag{a}\]
{{< /fig-math >}}

この番号にはラベル `\label` を付けて参照することができる。
たとえば `\[E=mc^2 \tag{b}\label{eq-b}\]` としておいて

{{< fig-math class="box" >}}
\[E=mc^2 \tag{b}\label{eq-b}\]
{{< /fig-math >}}

このラベルを参照するには `\eqref{eq-b}` でこのように →\eqref{eq-b}← 表示できる（`$...$` で囲まなくてもよい）[^anc1]。

[^anc1]: 数式参照用の id は [MathJax] で動的に生成されているので，ページ外からの参照はお勧めできない。

### 数式に通し番号を振る

ページ内で通し番号を振りたい場合は，以下のように設定を既定値（`none`）から変更する（オプション設定については[初期設定]の回を参照）。

```js
MathJax = {
  tex: {
    tags: 'ams'
  }
};
```

実際に番号を振るには数式を `\[...\]` で囲むのではなく `\begin{equation}...\end{equation}` で囲む。
自動で番号が振られるので `\tag` コマンドは不要である。

```tex
エネルギーと質量には
\begin{equation}
  E=mc^2 \label{eq-1st}
\end{equation}
の関係がある。
```

{{< div-box >}}
エネルギーと質量には
\begin{equation}
  E=mc^2 \label{eq-1st}
\end{equation}
の関係がある。
{{< /div-box >}}

参照も同様に →\eqref{eq-1st}← 表示できる。

余談だが `\begin` と `\end` で囲まれた領域を $\mathrm{\LaTeX}$  では「環境」と呼ぶ。
`\begin{foober}...\end{foober}` なら「`foobar` 環境」と呼んだりする。

`equation` にアスタリスクを付けた `equation*` 環境では，環境内の数式を自動採番の対象外にできる[^math1]。

[^math1]: 実は `\[...\]` は `equation*` 環境と等価である。

```tex
エネルギーと質量には
\begin{equation*}
  E=mc^2
\end{equation*}
の関係がある。
```

{{< div-box >}}
エネルギーと質量には
\begin{equation*}
  E=mc^2
\end{equation*}
の関係がある。
{{< /div-box >}}

`equation` 環境以外にも `align` 環境や `gather` 環境なども自動採番の対象となる（自動採番を無効にするアスタリスクも有効）。
`align` 環境などで特定の式に番号を振りたくない場合は `\notag` コマンドを使って

```tex
\begin{align}
  (a+b)^2 &= a^2+2ab+b^2 \\
  (a-b)^2 &= a^2-2ab+b^2  \notag \\
  (a+b)^3 &= a^3+3a^{2}b+3ab^2+b^3
\end{align}
```

{{< div-box >}}
\begin{align}
  (a+b)^2 &= a^2+2ab+b^2 \\
  (a-b)^2 &= a^2-2ab+b^2  \notag \\
  (a+b)^3 &= a^3+3a^{2}b+3ab^2+b^3
\end{align}
{{< /div-box >}}

とすればよい。
また複数の数式を `split` 環境や `aligned` 環境で囲むことで複数の数式の塊に一つの番号を振ることもできる[^math2]。

[^math2]: `split` 環境は自動採番の対象外である。

```tex
\begin{equation}
  \begin{split}
    (a+b)^2 &= a^2+2ab+b^2 \\
    (a-b)^2 &= a^2-2ab+b^2 \\
    (a+b)^3 &= a^3+3a^{2}b+3ab^2+b^3
  \end{split}
\end{equation}
```

{{< div-box >}}
\begin{equation}
\begin{split}
  (a+b)^2 &= a^2+2ab+b^2 \\
  (a-b)^2 &= a^2-2ab+b^2 \\
  (a+b)^3 &= a^3+3a^{2}b+3ab^2+b^3
\end{split}
\end{equation}
{{< /div-box >}}

[MathJax]: https://www.mathjax.org/
[初期設定]: {{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}} "ちょこっと MathJax： 初期設定"
[前回]: {{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}} "ちょこっと MathJax： 基本的な数式表現"
[LaTeX2ε美文書作成入門]: https://www.amazon.co.jp/dp/4774187054?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon | [改訂第7版]LaTeX2ε美文書作成入門 | 奥村 晴彦, 黒木 裕介 通販"

## 参考図書 {#books}

{{% review-paapi "4297117126" %}} <!-- LaTeX2ε美文書作成入門 -->
