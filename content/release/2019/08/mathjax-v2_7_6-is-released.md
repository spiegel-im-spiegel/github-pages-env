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

{{< fig-quote type="markdown" title="MathJax v2.7.6 now available" link="https://www.mathjax.org/mathjax-v2-7-6-now-available/" lang="en" >}}
{{< quote >}}This is a maintenance release that fixes an issue with the latest.js file that is used to obtain the most current 2.x version of MathJax from one of the CDNs that serves MathJax. The problem is that the most current version is only obtained if the highest version on the CDN is version 2.x.y for some x and y, so when MathJax goes to version 3.0.0 (scheduled for August 31st), latest.js will find that the current CDN version is 3.0.0 and (correctly) will not switch to that, but instead will (incorrectly) use the version from which latest.js was loaded rather than the highest 2.x.y available. This means that when version 3.0 is released, sites using latest.js will fall back from version 2.7.5 to the version that they specified for latests.js. MathJax will still run on those pages, but it may be an earlier version than you have been getting in the past.{{< /quote >}}
{{< /fig-quote >}}

つか v3 の（βが取れた）正式版が8月末（日本では9月初？）にようやく出るのか。
[MathJax] v3 は configuration が大きく変わると聞いているので，しばらくは要注意だな。

アップデートは計画的に。

## ブックマーク

- [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})

[MathJax]: https://www.mathjax.org/ "MathJax | Beautiful math in all browsers."

## 参考図書

{{% review-paapi "4774187054" %}} <!-- [改訂第7版]LaTeX2ε美文書作成入門 -->
