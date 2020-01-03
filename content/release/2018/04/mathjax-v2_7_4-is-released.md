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
