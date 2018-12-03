+++
title = "MathJax v2.7.4 is Released"
date = "2018-04-08T10:14:06+09:00"
description = "不具合の修正がメインのようだ。"
image = "/images/attention/tools.png"
tags = [ "package", "math", "mathjax", "javascript" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Web ブラウザ上で高品質な数式表現を行うための JavaScript パッケージである [MathJax] v2.7.4 がリリースされた。

- [MathJax v2.7.4 now available | MathJax](https://www.mathjax.org/mathjax-v2-7-4-now-available/)
- [MathJax v2.7.4 Milestone](https://github.com/mathjax/MathJax/milestone/18?closed=1)

不具合の修正がメインのようだ。

{{% fig-quote title="MathJax v2.7.4 now available" link="https://www.mathjax.org/mathjax-v2-7-4-now-available/" lang="en" %}}
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

[MathJax]: https://www.mathjax.org/ "MathJax | Beautiful math in all browsers."

## ブックマーク

- [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51E5K7B53aL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/">[改訂第7版]LaTeX2ε美文書作成入門</a></dt><dd>奥村 晴彦 黒木 裕介 </dd><dd>技術評論社 2017-01-24</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798118141/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798118141.09._SCTHUMBZZZ_.jpg"  alt="LaTeX2e辞典~用法・用例逆引きリファレンス (DESKTOP REFERENCE)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4535558752/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4535558752.09._SCTHUMBZZZ_.jpg"  alt="公共政策入門 ミクロ経済学的アプローチ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320112415/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320112415.09._SCTHUMBZZZ_.jpg"  alt="Rで楽しむ統計 (Wonderful R 1)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298550/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298550.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.5"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797391383/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797391383.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/積分を見つめて (数学ガールの秘密ノートシリーズ)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298569/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298569.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.6"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798115363/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798115363.09._SCTHUMBZZZ_.jpg"  alt="独習 LaTeX2ε"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4785315717/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4785315717.09._SCTHUMBZZZ_.jpg"  alt="具体例から学ぶ 多様体"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774193046/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774193046.09._SCTHUMBZZZ_.jpg"  alt="【改訂第3版】基礎からわかる情報リテラシー"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4768704700/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4768704700.09._SCTHUMBZZZ_.jpg"  alt="はじめて学ぶリー群 ―線型代数から始めよう"  /></a> </p>
<p class="description">ついに第7版が登場。紙の本で買って常に側に置いておくのが吉。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-27">2017-09-27</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
