+++
title = "Hugo 0.122 で TeX/LaTeX 数式表現ができるようになった（いまさら？）"
date =  "2024-01-27T08:52:26+09:00"
description = "そういえば別行立ての数式は shortcode とか使って細工する必要があったな"
image = "/images/attention/kitten.jpg"
tags = [ "math", "mathjax", "hugo", "tex" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[Hugo] 0.122.0 がリリースされた。

- [Release v0.122.0 · gohugoio/hugo · GitHub](https://github.com/gohugoio/hugo/releases/tag/v0.122.0)

{{< fig-gen >}}
<iframe src="https://fosstodon.org/@gohugoio/111823216188959656/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe><script src="https://fosstodon.org/embed.js" async="async"></script>
{{< /fig-gen >}}

今回の目玉機能は $\mathrm{\TeX}$/$\mathrm{\LaTeX}$ 記法による数式表現に対応したことらしい。
えっ，いまさら？ あー，いや，そういえば別行立ての数式は shortcode とか使って細工する必要があったな，そういえば。

[Hugo] で $\mathrm{\TeX}$/$\mathrm{\LaTeX}$ 記法による数式表現に対応するには

1. サイト設定ファイル `hugo.toml` (または `config.toml`) で数式表現用のデリミタ文字列を指定する
2. 数式変換用のパッケージ（[MathJax] または [KaTeX]）を導入する

といった手順が必要。

サイト設定ファイルの内容はこんな感じ。

```toml
[markup]
  defaultMarkdownHandler = "goldmark"
  [markup.goldmark]
    [markup.goldmark.extensions]
      [markup.goldmark.extensions.passthrough]
        enable = true
        [markup.goldmark.extensions.passthrough.delimiters]
          block = [['\[', '\]'], ['$$', '$$']]
```

YAML や JSON で書く場合は [Hugo のドキュメント](https://gohugo.io/content-management/mathematics/ "Mathematics in markdown | Hugo")を参照のこと。

[KaTeX] は使ったことがないが [MathJax] については以前に記事にしたことがあるので参考にどうぞ。

1. [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})
2. [ちょこっと MathJax： 基本的な数式表現]({{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}})
3. [ちょこっと MathJax： インライン数式と別行立て数式]({{< ref "/remark/2017/10/getting-started-mathjax-3.md" >}})
4. [ちょこっと MathJax 番外編： mathcomp パッケージの代替え]({{< ref "/remark/2017/12/mathcomp-in-mathjax.md" >}})

これでたとえば

```tex
\[
\begin{aligned}
KL(\hat{y} || y) &= \sum_{c=1}^{M}\hat{y}_c \log{\frac{\hat{y}_c}{y_c}} \\
JS(\hat{y} || y) &= \frac{1}{2}(KL(y||\frac{y+\hat{y}}{2}) + KL(\hat{y}||\frac{y+\hat{y}}{2}))
\end{aligned}
\]
```

と書けば

\[
\begin{aligned}
KL(\hat{y} || y) &= \sum_{c=1}^{M}\hat{y}_c \log{\frac{\hat{y}_c}{y_c}} \\
JS(\hat{y} || y) &= \frac{1}{2}(KL(y||\frac{y+\hat{y}}{2}) + KL(\hat{y}||\frac{y+\hat{y}}{2}))
\end{aligned}
\]

と展開される（筈）

いや，待て。
`\[ ... \]` って `\begin{equation*} ... \end{equation*}` と等価だろ。
問題なく処理はされるだろうけど，意味としておかしくないか。

試しに `\[ ... \]` を外して

```tex
\begin{aligned}
KL(\hat{y} || y) &= \sum_{c=1}^{M}\hat{y}_c \log{\frac{\hat{y}_c}{y_c}} \\
JS(\hat{y} || y) &= \frac{1}{2}(KL(y||\frac{y+\hat{y}}{2}) + KL(\hat{y}||\frac{y+\hat{y}}{2}))
\end{aligned}
```

とすると期待通りには変換しない。
Markdown のパーサが上の記述を丸ごと `<p>` 要素で囲ってしまうからのようだ。

```text
エネルギーと質量には $$E=mc^2$$ の関係がある。
```

みたいな記述なら

```html
<p>エネルギーと質量には </p>
$$E=mc^2$$
<p> の関係がある。</p>
```

などとパーサ側でうまく分離してくれるのだが。

`\begin{aligned} ... \end{aligned}` のような $\mathrm{\LaTeX}$ 環境を無駄に `\[ ... \]` で囲むのがどうしても嫌な方は，強制的に `<div>` 要素で囲って

```tex
<div>
\begin{aligned}
KL(\hat{y} || y) &= \sum_{c=1}^{M}\hat{y}_c \log{\frac{\hat{y}_c}{y_c}} \\
JS(\hat{y} || y) &= \frac{1}{2}(KL(y||\frac{y+\hat{y}}{2}) + KL(\hat{y}||\frac{y+\hat{y}}{2}))
\end{aligned}
</div>
```

とすれば markdown のパーサが生の HTML 記述と解釈しスルーしてくれる。
私は数式専用の shortcode を作って利用している。

```tex
{{</* fig-math */>}}
\begin{aligned}
KL(\hat{y} || y) &= \sum_{c=1}^{M}\hat{y}_c \log{\frac{\hat{y}_c}{y_c}} \\
JS(\hat{y} || y) &= \frac{1}{2}(KL(y||\frac{y+\hat{y}}{2}) + KL(\hat{y}||\frac{y+\hat{y}}{2}))
\end{aligned}
{{</* /fig-math */>}}
```

## ブックマーク

- [Mathematics in markdown | Hugo](https://gohugo.io/content-management/mathematics/)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[MathJax]: https://www.mathjax.org/ "MathJax | Beautiful math in all browsers."
[KaTeX]: https://katex.org/ "KaTeX – The fastest math typesetting library for the web"

## 参考図書 {#books}

{{% review-paapi "4297117126" %}} <!-- LaTeX2ε美文書作成入門 -->
