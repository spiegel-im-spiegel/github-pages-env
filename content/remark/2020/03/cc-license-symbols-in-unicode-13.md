+++
title = "Unicode 13 に CC Licenses シンボルが追加された"
date =  "2020-03-20T22:04:19+09:00"
description = "実際にはフォントも対応しないといけないのだが，それはしばらくかかるかな。"
image = "/images/attention/kitten.jpg"
tags = ["creative-commons", "license", "unicode", "character"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2020-03-10 にリリースされた [Unicode 13](https://unicode.org/versions/Unicode13.0.0/ "Unicode 13.0.0") に Creative Commons Licenses のシンボルが追加されたらしい。

- [The Unicode Standard Now Includes CC License Symbols - Creative Commons](https://creativecommons.org/2020/03/18/the-unicode-standard-now-includes-cc-license-symbols/)

追加された CC Licenses シンボルは以下の通り。

| 符号点    |         文字          | 名前                                        |
| --------- |:---------------------:| ------------------------------------------- |
| `U+1F16D` |  {{< cc-sym "cc" >}}  | CIRCLED CC                                  |
| `U+1F16E` |  {{< cc-sym "pd" >}}  | CIRCLED C WITH OVERLAID BACKSLASH           |
| `U+1F16F` |  {{< cc-sym "by" >}}  | CIRCLED HUMAN FIGURE                        |
| `U+1F10D` | {{< cc-sym "zero" >}} | CIRCLED ZERO WITH SLASH                     |
| `U+1F10E` |  {{< cc-sym "sa" >}}  | CIRCLED ANTICLOCKWISE ARROW                 |
| `U+1F10F` |  {{< cc-sym "nc" >}}  | CIRCLED DOLLAR SIGN WITH OVERLAID BACKSLASH |

ちなみに「改変禁止」を表す <i class="fab fa-creative-commons-nd"></i> は数学の演算子 CIRCLED EQUALS (`U+229C`) を流用するようだ。
なんだかなぁ。

これらのシンボルを使って CC Licenses の各条件の組み合わせは以下のように表現できる。

| 記号                                                        | 意味                                                                       |
| ----------------------------------------------------------- | -------------------------------------------------------------------------- |
| {{< cc-sym "pd" >}}                                         | 公有（public domain）                                                      |
| {{< cc-sym "zero" >}}                                       | [CC0](https://creativecommons.org/publicdomain/zero/1.0/)              |
| {{< cc-sym "by" >}}                                         | [表示](https://creativecommons.org/licenses/by/4.0/)                       |
| {{< cc-sym "by" >}} {{< cc-sym "sa" >}}                     | [表示-継承](https://creativecommons.org/licenses/by-sa/4.0/)               |
| {{< cc-sym "by" >}} {{< cc-sym "nc" >}}                     | [表示-非営利](https://creativecommons.org/licenses/by-nc/4.0/)             |
| {{< cc-sym "by" >}} {{< cc-sym "nc" >}} {{< cc-sym "sa" >}} | [表示-非営利-継承](https://creativecommons.org/licenses/by-nc-sa/4.0/)     |
| {{< cc-sym "by" >}} {{< cc-sym "nd" >}}                     | [表示-改変禁止](https://creativecommons.org/licenses/by-nd/4.0/)           |
| {{< cc-sym "by" >}} {{< cc-sym "nd" >}} {{< cc-sym "nc" >}} | [表示-非営利-改変禁止](https://creativecommons.org/licenses/by-nc-nd/4.0/) |

実際にはフォントも対応しないといけないのだが，それはしばらくかかるかな。

## ブックマーク

- [改訂3版： CC Licenses について]({{< rlnk "/cc-licenses/" >}})
- [Font Awesome 5.0.11 で Creative Commons アイコンに完全対応した]({{< ref "/release/2018/05/creative-commons-icons-by-font-awesome.md" >}})

## 参考図書

{{% review-paapi "475710152X" %}} <!-- クリエイティブ・コモンズ―デジタル時代の知的財産権 -->
{{% review-paapi "B00DI8TMPU" %}} <!-- オープン化する創造の時代 -->
{{% review-paapi "4757102852" %}} <!-- 著作権２．０ ウェブ時代の文化発展をめざして -->
