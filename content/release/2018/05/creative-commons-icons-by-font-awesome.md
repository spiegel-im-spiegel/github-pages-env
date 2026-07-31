+++
title = "Font Awesome 5.0.11 で Creative Commons アイコンに完全対応した"
date = "2018-05-13T15:31:54+09:00"
description = "先日リリースされた Font Awesome 5.0.11 で Creative Commons アイコンに完全対応したらしい。"
tags = [ "creative-commons", "web", "font", "site" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日リリースされた [Font Awesome] 5.0.11 で [Creative Commons] アイコンに完全対応したらしい。

- [What’s new in Font Awesome 5.0.11 – Font Awesomeness](https://blog.fontawesome.com/whats-new-in-font-awesome-5-0-11-ff914ce33258)

{{< fig-quote title="What’s new in Font Awesome 5.0.11" link="https://blog.fontawesome.com/whats-new-in-font-awesome-5-0-11-ff914ce33258" lang="en" >}}
<q>The <a href="https://fontawesome.com/icons/creative-commons?style=brands">Creative Commons brand icon</a> has been in Font Awesome for a long time, but we never added <a href="https://fontawesome.com/icons?d=gallery&q=creative%20commons">the rest of the symbols</a> folks commonly use to express what can and can’t be done with a piece of work. Now included with Font Awesome Free, you can easily tell someone exactly where you project stands legally with just a couple more icons.</q>
{{< /fig-quote >}}

既に retire したサンプリング・ライセンスのアイコンも含まれている。
よく使うであろうアイコンはこんな感じ。

|                         | コード                                          | 意味        |
|:-----------------------:| ----------------------------------------------- | ----------- |
|  {{< cc-syms "cc" >}}   | `<i class="fab fa-creative-commons"></i>`       | [CC] マーク |
|  {{< cc-syms "by" >}}   | `<i class="fab fa-creative-commons-by"></i>`    | 表示        |
|  {{< cc-syms "sa" >}}   | `<i class="fab fa-creative-commons-sa"></i>`    | 継承        |
|  {{< cc-syms "nc" >}}   | `<i class="fab fa-creative-commons-nc"></i>`    | 非営利      |
| {{< cc-syms "nc-jp" >}} | `<i class="fab fa-creative-commons-nc-jp"></i>` | 非営利      |
| {{< cc-syms "nc-eu" >}} | `<i class="fab fa-creative-commons-nc-eu"></i>` | 非営利      |
|  {{< cc-syms "nd" >}}   | `<i class="fab fa-creative-commons-nd"></i>`    | 改変禁止    |
| {{< cc-syms "zero" >}}  | `<i class="fab fa-creative-commons-zero"></i>`  | [CC0]       |
|  {{< cc-syms "pd" >}}   | `<i class="fab fa-creative-commons-pd"></i>`    | [公有]      |

このサイトも [Creative Commons Icon Font](http://cc-icons.github.io/) から [Font Awesome] に切り換えた。

## そうそう。

[State of the Commons 2017](https://stateof.creativecommons.org/) が登場しているよ。

## ブックマーク

- [Font Awesome 5 がリリースされた]({{< ref "/remark/2017/12/font-awesome-5-released.md" >}})
- [Web フォントに関する覚え書き]({{< ref "/remark/2015/web-font-family.md" >}})

- [Creative Commons Licenses]({{< ref "/cc-licenses/02-creative-commons-licenses.md" >}}) — [改訂3版： CC Licenses について](/cc-licenses/)

[Font Awesome]: https://fontawesome.com/ "Font Awesome 5 | Font Awesome"
[Creative Commons]: https://creativecommons.org/
[CC]: https://creativecommons.org/
[CC0]: https://creativecommons.org/publicdomain/zero/1.0/ "Creative Commons — CC0 1.0 Universal"
[公有]: https://creativecommons.org/publicdomain/mark/1.0/ "Creative Commons — Public Domain Mark 1.0"
