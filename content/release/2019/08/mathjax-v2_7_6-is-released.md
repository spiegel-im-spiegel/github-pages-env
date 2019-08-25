+++
title = "MathJax v2.7.6 がリリースされた"
date =  "2019-08-25T10:25:16+09:00"
description = "つか v3 の（βが取れた）正式版が8月末（日本では9月初？）にようやく出るのか。"
image = "/images/attention/tools.png"
tags = [ "tools", "math", "mathjax", "javascript" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Web ブラウザ上で高品質な数式表現を行うための JavaScript パッケージである [MathJax] v2.7.6 がリリースされた。

- [MathJax v2.7.6 now available | MathJax](https://www.mathjax.org/mathjax-v2-7-6-now-available/)

今回はメンテナンス・リリースで [MathJax] v3 正式リリースに向けた修正のようだ。

{{< fig-quote type="md" title="MathJax v2.7.6 now available" link="https://www.mathjax.org/mathjax-v2-7-6-now-available/" lang="en" >}}
{{< quote >}}This is a maintenance release that fixes an issue with the latest.js file that is used to obtain the most current 2.x version of MathJax from one of the CDNs that serves MathJax. The problem is that the most current version is only obtained if the highest version on the CDN is version 2.x.y for some x and y, so when MathJax goes to version 3.0.0 (scheduled for August 31st), latest.js will find that the current CDN version is 3.0.0 and (correctly) will not switch to that, but instead will (incorrectly) use the version from which latest.js was loaded rather than the highest 2.x.y available. This means that when version 3.0 is released, sites using latest.js will fall back from version 2.7.5 to the version that they specified for latests.js. MathJax will still run on those pages, but it may be an earlier version than you have been getting in the past.{{< /quote >}}
{{< /fig-quote >}}

つか v3 の（βが取れた）正式版が8月末（日本では9月初？）にようやく出るのか。
[MathJax] v3 は configuration が大きく変わると聞いているので，しばらくは要注意だな。

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
