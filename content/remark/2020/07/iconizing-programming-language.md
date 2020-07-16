+++
title = "アイコン化されたプログラミング言語"
date =  "2020-07-16T11:43:19+09:00"
description = "おまけ「人気（？）のプログラミング言語」も追記した。"
image = "/images/attention/kitten.jpg"
tags = [ "font", "programming", "language" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Font Awesome] を{{< ruby "つらつら" >}}倩{{< /ruby >}}と眺めていたらプログラミング言語のアイコンとかもあるんだねぇ。
多分ファイル種別と関連させて使うんだろうけど。

面白いのでプログラミング言語を示す [Font Awesome] アイコンを拾ってみた。
色は気にしないように。

{{< fig-gen type="markdown" >}}
|                             Icon                              | Code                             |
|:-------------------------------------------------------------:|:-------------------------------- |
|  {{% span class="fa-2x" %}}{{% icons "css3" %}}{{% /span %}}  | `<i class="fab fa-css3"></i>`    |
| {{% span class="fa-2x" %}}{{% icons "erlang" %}}{{% /span %}} | `<i class="fab fa-erlang"></i>`  |
| {{% span class="fa-2x" %}}{{% icons "html5" %}}{{% /span %}}  | `<i class="fab fa-html5"></i>`   |
|  {{% span class="fa-2x" %}}{{% icons "java" %}}{{% /span %}}  | `<i class="fab fa-java"></i>`    |
| {{% span class="fa-2x" %}}{{% icons "nodejs" %}}{{% /span %}} | `<i class="fab fa-node-js"></i>` |
|  {{% span class="fa-2x" %}}{{% icons "php" %}}{{% /span %}}   | `<i class="fab fa-php"></i>`     |
| {{% span class="fa-2x" %}}{{% icons "python" %}}{{% /span %}} | `<i class="fab fa-python"></i>`  |
|  {{% span class="fa-2x" %}}{{% icons "rust" %}}{{% /span %}}  | `<i class="fab fa-rust"></i>`    |
| {{% span class="fa-2x" %}}{{% icons "swift" %}}{{% /span %}}  | `<i class="fab fa-swift"></i>`   |
{{< /fig-gen >}}

HTML5 や CSS3 をプログラミング言語と言っていいか微妙だが，ついでということで（笑）

今回は 2020-07-16 時点の [Font Awesome] 5.13.1 で提供されるフリー・アイコンの中から目視で探してみた。
見落とし等あれば指摘していただくと主に私が喜びます。

アイコン・デザインは商標権とか絡んでくるから難しいかもしれないけど，こういうブランド化も面白いかもね。

## 【おまけ】 人気（？）のプログラミング言語

今回は [TIOBE] ランキング50位内にあるメジャーな言語で浚ってみたのだが，このランキングはいつも C と Java が2強で，あとはどんぐりの背比べと代わり映えがない。
強いて言うなら，以前は C++ が3位に絡んでたけど今は Python かな，やっぱり。

まぁ，保守的といえば保守的なんだろうけど...

何故か初夏の季節はプログラミング言語のランキング記事が増えるのだが，その中で個人的なイメージに近いのが以下の記事。

- [The State of Developer Ecosystem in 2020 Infographic | JetBrains: Developer Tools for Professionals and Teams](https://www.jetbrains.com/lp/devecosystem-2020/)
- [どのプログラミング言語が使われているのか、JetBrainsが調査レポートを発表：移行先はGo、Kotlin、Python - ＠IT](https://www.atmarkit.co.jp/ait/articles/2007/14/news138.html)

これは JetBrains 社によるアンケート結果だそうで，なかなか面白い内容である。

たとえば，いわゆる「第1言語」としてよく使われているのが

1. JavaScript
1. Java
1. Python
1. HTML/CSS
1. SQL
1. PHP
1. C++
1. C#
1. TypeScript
1. Go
1. Kotlin

らしい。
まぁ HTML/CSS は洒落だと思うが。
そんで，これらの中で移行を検討している言語が以下の表になっている。

{{< fig-img src="./migrate-to-other-languages.png" title="via “The State of Developer Ecosystem in 2020 Infographic”" link="https://www.jetbrains.com/lp/devecosystem-2020/" width="1048" >}}

まぁ，そもそも「移行したい」と考えるほうが少数派なのだが，その少数派の動向を見ると，  Java や Python などから Go へ移行したいと考えている人が意外にいる一方で Go から Rust へ移行したいと考える人も多いことに気づく。

これって，この数年くらいのトレンドに合ってるような気がするのだが，どうだろう。

ちなみに，上の人気（？）プログラミング言語の用途は以下のようになっている。

{{< fig-img src="./what-types-of-software-do-you-develop.png" title="via “The State of Developer Ecosystem in 2020 Infographic”" link="https://www.jetbrains.com/lp/devecosystem-2020/" width="1048" >}}



## ブックマーク

- [Googleが「Open Usage Commons」設立。オープンソースの商標を自由かつ公正に使用するための支援団体 － Publickey](https://www.publickey1.jp/blog/20/googleopen_usage_commons.html)
- [Font Awesome 5.0.11 で Creative Commons アイコンに完全対応した]({{< ref "/release/2018/05/creative-commons-icons-by-font-awesome.md" >}})
- [Font Awesome 5.13 に COVID-19 関連アイコンが登場]({{< ref "/remark/2020/03/covid-19-icons-in-font-awesome.md" >}})

[Font Awesome]: https://fontawesome.com/
[TIOBE]: https://www.tiobe.com/tiobe-index/ "TIOBE - The Software Quality Company"
