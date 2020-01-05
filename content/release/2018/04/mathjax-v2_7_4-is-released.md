+++
title = "MathJax v2.7.4 is Released"
date = "2018-04-08T10:14:06+09:00"
description = "不具合の修正がメインのようだ。"
image = "/images/attention/tools.png"
tags = [ "tools", "math", "mathjax", "javascript" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

Web ブラウザ上で高品質な数式表現を行うための JavaScript パッケージである [MathJax] v2.7.4 がリリースされた。

- [MathJax v2.7.4 now available | MathJax](https://www.mathjax.org/mathjax-v2-7-4-now-available/)
- [MathJax v2.7.4 Milestone](https://github.com/mathjax/MathJax/milestone/18?closed=1)

不具合の修正がメインのようだ。

{{% fig-quote type="markdown" title="MathJax v2.7.4 now available" link="https://www.mathjax.org/mathjax-v2-7-4-now-available/" lang="en" %}}
- Prevent infinite loop if an autoloaded component fails to load. (#1936)
- Always set movablelimits to false in `\overset` and `\underset`. (#1929)
- CSS reset for box-sizing in HTML-CSS output. (#1942)
- Add `px` to `max-width` for SVG output containing tags. (#1950)
- Properly handle namespaces starting with `math` in MathML input. (#1951)
- Make `tex2jax` and `asciimath2jax` rescan after unmatched delimiter. (#1960)
- Fix minimum height of accents in scripts. (#1956)
- Make monospaced non-breaking space be of correct width. (#1953)
- Handle size of centered large operators correctly in mrows. (#1933)
{{% /fig-quote %}}

アップデートは計画的に。

## ブックマーク

- [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})

[MathJax]: https://www.mathjax.org/ "MathJax | Beautiful math in all browsers."

## 参考図書

{{% review-paapi "4774187054" %}} <!-- [改訂第7版]LaTeX2ε美文書作成入門 -->
