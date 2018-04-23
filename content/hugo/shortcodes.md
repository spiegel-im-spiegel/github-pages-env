+++
date = "2015-10-24T16:36:04+09:00"
update = "2017-10-27T17:27:50+09:00"
description = "たとえば記事の中に YouTube の動画や SlideShare の HTML コードを埋め込みたいとする。こういうときは Shortcodes の機能を使うと便利だ。"
draft = false
tags = [ "hugo", "shortcodes" ]
title = "Shortcodes で HTML コードを埋め込む"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

今回は [Shortcodes] の話。

たとえば記事の中に [YouTube] の動画や [SlideShare] の HTML コードを埋め込みたいとする。
[Hugo] では段落単位に HTML コードを記述するとそのまま出力してくれるけど，頻繁に使うコードはテンプレート化して再利用したいだろう。
こういうときは [Shortcodes] の機能を使うと便利だ。

## Shortcodes の使い方

[Shortcodes] は partial template の一種だけど，テンプレートの中に埋め込むのではなく，記事（markdown）ファイルに埋め込むことができる。

### 指定した範囲を加工する

書式は以下のとおりである。

```
{{</* shortcodename */>}} ... {{</* /shortcodename */>}}
```

[Shortcodes] の実体は `layouts/shortcodes` フォルダに `shortcodename.hml` などとコード名をそのままファイル名にしてで配置する。
たとえば `layouts/shortcodes/fig-quote.html` というファイルを作成し，以下のコードを書く。

```html
{{ with .Get "lang" }}<figure lang="{{ . }}">{{ else }}<figure>{{ end }}
<blockquote>{{ .Inner }}</blockquote>
{{ if .Get "title"}}<figcaption>via <q>{{ with .Get "link"}}<a href="{{.}}">{{ end }}{{ .Get "title" }}{{ with .Get "link"}}</a>{{ end }}</q></figcaption>{{ end }}
</figure>
```

これを使って

```html
{{</* fig-quote title="Compatible Licenses - Creative Commons" link="https://creativecommons.org/compatiblelicenses" lang="en" */>}}
<q>The <a href="https://www.gnu.org/copyleft/gpl.html">GNU General Public License version 3 </a> was declared a <q>BY-SA–Compatible License</q> for version 4.0 on 8 October 2015.
Note that compatibility with the GPLv3 is one-way only, which means you may license your contributions to adaptations of BY-SA 4.0 materials under GPLv3, but you may not license your contributions to adaptations of GPLv3 projects under BY-SA 4.0.
Other <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3#Considerations_for_adapters_applying_the_GPLv3" target="_blank">special considerations</a> apply.
See the full <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3">analysis</a> and <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility_analysis:_GPL">comparison</a> for more information.</q>
{{</* /fig-quote */>}}
```

このように記述すると，以下のように展開される。

```html
<figure lang="en">
<blockquote>
<q>The <a href="https://www.gnu.org/copyleft/gpl.html">GNU General Public License version 3 </a> was declared a <q>BY-SA–Compatible License</q> for version 4.0 on 8 October 2015.
Note that compatibility with the GPLv3 is one-way only, which means you may license your contributions to adaptations of BY-SA 4.0 materials under GPLv3, but you may not license your contributions to adaptations of GPLv3 projects under BY-SA 4.0.
Other <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3#Considerations_for_adapters_applying_the_GPLv3" target="_blank">special considerations</a> apply.
See the full <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3">analysis</a> and <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility_analysis:_GPL">comparison</a> for more information.</q>
</blockquote>
<figcaption>via <q><a href="https://creativecommons.org/compatiblelicenses">Compatible Licenses - Creative Commons</a></q></figcaption>
</figure>
```

{{< fig-quote title="Compatible Licenses - Creative Commons" link="https://creativecommons.org/compatiblelicenses" lang="en" >}}
<q>The <a href="https://www.gnu.org/copyleft/gpl.html">GNU General Public License version 3 </a> was declared a <q>BY-SA–Compatible License</q> for version 4.0 on 8 October 2015.
Note that compatibility with the GPLv3 is one-way only, which means you may license your contributions to adaptations of BY-SA 4.0 materials under GPLv3, but you may not license your contributions to adaptations of GPLv3 projects under BY-SA 4.0.
Other <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3#Considerations_for_adapters_applying_the_GPLv3">special considerations</a> apply.
See the full <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3">analysis</a> and <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility_analysis:_GPL">comparison</a> for more information.</q>
{{< /fig-quote >}}

`{{</* fig-quote */>}} ... {{</* /fig-quote */>}}` で囲まれている部分が `{{ .Inner }}` として展開されているのがお分かりだろうか。

[Shortcodes] では `{{ .Inner }}` 以外に任意のパラメータを持たせることができる。
上の例で言うなら `lang`, `title`, `link` が `{{</* fig-quote */>}}` に対するパラメータである。

名前を持たないパラメータでも有効である。
ここではルビ（ruby）を振る [Shortcodes] `layouts/shortcodes/ruby.html` を考える。
こんな感じ。

```html
<ruby><rb>{{ .Inner }}</rb><rp> (</rp><rt>{{ index .Params 0 }}</rt><rp>) </rp></ruby>
```

これで例えば

```html
{{</* ruby "ひらがな" */>}}平仮名{{</* /ruby */>}}
```

と記述すれば

```html
これで例えば <ruby><rb>平仮名</rb><rp> (</rp><rt>ひらがな</rt><rp>) </rp></ruby> と記述すれば
```

> これで例えば {{< ruby "ひらがな" >}}平仮名{{< /ruby >}} と記述すれば

と展開される。
パラメータの順番に意味ができるので，多数のパラメータが必要な場合には向かないかもしれない。

更に `{{%/* fig-quote */%}} ... {{%/* /fig-quote */%}}` と表記すると，囲まれた部分を Markdown と解釈する。
先程の `fig-quote` を例にすると

```html
{{%/* fig-quote title="Compatible Licenses - Creative Commons" link="https://creativecommons.org/compatiblelicenses" lang="en" */%}}
“The [GNU General Public License version 3](https://www.gnu.org/copyleft/gpl.html) was declared a “BY-SA–Compatible License” for version 4.0 on 8 October 2015.
Note that compatibility with the GPLv3 is one-way only, which means you may license your contributions to adaptations of BY-SA 4.0 materials under GPLv3, but you may not license your contributions to adaptations of GPLv3 projects under BY-SA 4.0.
Other [special considerations](https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3#Considerations_for_adapters_applying_the_GPLv3) apply.
See the full [analysis](https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3) and [comparison](https://wiki.creativecommons.org/wiki/ShareAlike_compatibility_analysis:_GPL) for more information.”
{{%/* /fig-quote */%}}
```

と Markdown で記述すれば以下のようになる。

```html
<figure lang="en">
<blockquote><p>“The <a href="https://www.gnu.org/copyleft/gpl.html">GNU General Public License version 3</a> was declared a “BY-SA–Compatible License” for version 4.0 on 8 October 2015.
Note that compatibility with the GPLv3 is one-way only, which means you may license your contributions to adaptations of BY-SA 4.0 materials under GPLv3, but you may not license your contributions to adaptations of GPLv3 projects under BY-SA 4.0.
Other <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3#Considerations_for_adapters_applying_the_GPLv3">special considerations</a> apply.
See the full <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3">analysis</a> and <a href="https://wiki.creativecommons.org/wiki/ShareAlike_compatibility_analysis:_GPL">comparison</a> for more information.”</p>
</blockquote>
<figcaption>via <q><a href="https://creativecommons.org/compatiblelicenses">Compatible Licenses - Creative Commons</a></q></figcaption>
</figure>
```

{{% fig-quote title="Compatible Licenses - Creative Commons" link="https://creativecommons.org/compatiblelicenses" lang="en" %}}
“The [GNU General Public License version 3](https://www.gnu.org/copyleft/gpl.html) was declared a “BY-SA–Compatible License” for version 4.0 on 8 October 2015.
Note that compatibility with the GPLv3 is one-way only, which means you may license your contributions to adaptations of BY-SA 4.0 materials under GPLv3, but you may not license your contributions to adaptations of GPLv3 projects under BY-SA 4.0.
Other [special considerations](https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3#Considerations_for_adapters_applying_the_GPLv3) apply.
See the full [analysis](https://wiki.creativecommons.org/wiki/ShareAlike_compatibility:_GPLv3) and [comparison](https://wiki.creativecommons.org/wiki/ShareAlike_compatibility_analysis:_GPL) for more information.”
{{% /fig-quote %}}

### [Shortcodes] のみで HTML 展開する場合

`{{</* shortcodename */>}}` のみ，または `{{</* shortcodename /*/>}}` で表記すると `{{ .Inner }}` を持たない [Shortcodes] を表現できる。
たとえば `layouts/shortcodes/fig-youtube.html` というファイルを作成し，以下のコードを書く。

```html
<figure style='margin:0 auto;text-align:center;'>
<div style="position: relative; margin: 0 2rem; padding-bottom: 56.25%; padding-top: 30px; height: 0; overflow: hidden;">
	<iframe class="youtube-player" style="position: absolute; top: 0; left: 0; width: 100%; height: 100%;" allowfullscreen frameborder="0" src="https://www.youtube-nocookie.com/embed/{{ .Get "id" }}" allowfullscreen></iframe>
</div>{{ if .Get "title"}}
<figcaption><a href="https://www.youtube.com/watch?v={{ .Get "id" }}">{{ .Get "title" }}</a></figcaption>
{{ end }}</figure>
```

これを使って

```html
{{</* fig-youtube id="Kjqff5bkUrE" title="「はやぶさ2」地球スイングバイ解説CG ／ Hayabusa2's Earth Swing-by CG - YouTube" */>}}
```

と記述すると，以下のように展開される。

```html
<figure style='margin:0 auto;text-align:center;'>
<div style="position: relative; margin: 0 2rem; padding-bottom: 56.25%; padding-top: 30px; height: 0; overflow: hidden;">
	<iframe class="youtube-player" style="position: absolute; top: 0; left: 0; width: 100%; height: 100%;" allowfullscreen frameborder="0" src="https://www.youtube-nocookie.com/embed/Kjqff5bkUrE" allowfullscreen></iframe>
</div>
<figcaption><a href="https://www.youtube.com/watch?v=Kjqff5bkUrE">「はやぶさ2」地球スイングバイ解説CG ／ Hayabusa2&#39;s Earth Swing-by CG - YouTube</a></figcaption>
</figure>
```

{{< fig-youtube id="Kjqff5bkUrE" width="500" height="281" title="「はやぶさ2」地球スイングバイ解説CG ／ Hayabusa2's Earth Swing-by CG - YouTube" >}}

## Shortcodes の例

[spiegel-im-spiegel/hugo-theme-text] theme では [Shortcodes] を収録してないが，この[サイトの作業用リポジトリ](https://github.com/spiegel-im-spiegel/github-pages-env)には[いくつか置いてある](https://github.com/spiegel-im-spiegel/github-pages-env/tree/master/layouts/shortcodes)。
再利用はご自由に。

## 組み込みの Shortcodes

[Hugo] にはあらかじめ組み込まれた [Shortcodes] が幾つかある。
この中で一番よく使うのは記事間の相互リンクを張ることのできる `ref` および `relref` だろう。

たとえばこの記事の「ブックマーク」節にリンクを張る場合は

```markdown
[このページのブックマーク]({{</* relref "#bookmark" */>}})
```

とすると

```html
<a href="#bookmark:10ef41a6c37b90d6a6452868d5ba00ba">このページのブックマーク</a>
```

のように展開される。
ちなみに `ref` は絶対パス， `relref` は相対パスに展開される。

相互リンクは [Hugo] で管理している記事ならどれでも指定できる。
たとえば

```markdown
[「ブックマーク」ページの「公式サイト」節]({{</* ref "hugo/bookmark.md#official" */>}})
```

と記述すると

```html
<a href="http://text.baldanders.info/hugo/bookmark/#official:9bacfa348e5fe42acc9342a16675997d">「ブックマーク」ページの「公式サイト」節</a>
```

のように展開される。

## ブックマーク{#bookmark}

- [Hugoサイト構築 | Watanabe-DENKI Inc. 渡辺電気株式会社](http://wdkk.co.jp/lab/hugo/) : お勧め！

[Hugo に関するブックマークはこちら]({{< ref "hugo/bookmark.md" >}})。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[Shortcodes]: https://gohugo.io/extras/shortcodes/ "Shortcodes"
[YouTube]: https://www.youtube.com/ "YouTube"
[SlideShare]: http://www.slideshare.net/ "Share and Discover Knowledge on LinkedIn SlideShare"
[spiegel-im-spiegel/hugo-theme-text]: https://github.com/spiegel-im-spiegel/hugo-theme-text
