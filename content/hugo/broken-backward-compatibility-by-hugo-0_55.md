+++
title = "Hugo 0.55 リリースでまた後方互換性が壊れた"
date = "2019-04-13T19:05:15+09:00"
description = "しょうがないので shortcode を設計し直したですよ orz"
image = "/images/attention/hugo.jpg"
tags = [ "hugo", "shortcodes" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## Shortcode の仕様変更

[Hugo] 0.45 のときは組み込み [shortcode] である [`ref`/`relref` の仕様変更]({{< relref "./cross-references-with-shortcodes.md" >}})だったが，今回は [shortcode] そのものの仕様が変更になった。

- [Hugo 0.55.0: The early Easter Egg Edition! | Hugo](https://gohugo.io/news/0.55.0-relnotes/)

{{% fig-quote type="md" title="Hugo 0.55.0: The early Easter Egg Edition!" link="https://gohugo.io/news/0.55.0-relnotes/" lang="en" %}}
{{% quote %}}Shortcodes using the `{{%/* */%}}` as the outer-most delimiter will now be fully rendered when sent to the content renderer (e.g. Blackfriday for Markdown), meaning they can be part of the generated table of contents, footnotes, etc.{{% /quote %}}
{{% /fig-quote %}}

具体的な例を挙げたほうが分かりやすいだろう。

たとえば記述内容を `<div>` 要素で囲むだけの簡単な [shortcode] “`div-box`” を作ってみる。
コードはこんな感じ。

```html
<div class="box">{{ .Inner }}</div>
```

この `div-box` を使って

```text
{{</* div-box */>}}**Hello world!**{{</* /div-box */>}}
```
 
と書けば `.Inner` 変数の内容がそのまま展開されて

```html
<div class="box">**Hello world!**</div>
```

となる。
ここで

```text
{{%/* div-box */%}}**Hello world!**{{%/* /div-box */%}}
```

と書き換えると，以前のバージョンでは

```html
<div class="box"><strong>Hello world!</strong></div>
```

のように `.Inner` 変数の内容が markdown の文法に従って変換されていたのが， 0.55 からは `{{</* */>}}` と同じように

```html
<div class="box">**Hello world!**</div>
```

と素通ししてしまうのだ。
もし markdown として処理したければ [shortcode] 側で

```html
<div class="box">{{ .Inner | markdownify }}</div>
```

と `.Inner` 変数の内容を `markdownify` 関数に渡して明示的に変換する必要がある[^c1]。

[^c1]: もうひとつの方法としてはテンプレート・ファイルの先頭で ``{{ $_hugo_config := `{ "version": 1 }` }}`` と呪文を唱えることで 0.55 以前の動作に戻る。が，これ将来バージョンで無効になるよなぁ，多分。

なんでこんなことになったかというと `{{%/* */%}}` での処理は地の記述と連動しているらしい。
たとえば `span` という名前で以下の内容の [shortcode] をつくり

```html
<span>{{ .Inner }}</span>
```

これを markdown の表の中で使うと

```markdown
| 強調したい |
| ---------- |
| {{</* span */>}}**強調**{{</* /span */>}} |
| {{%/* span */%}}**強調**{{%/* /span */%}} |
```

| 強調したい                         |
| ---------------------------------- |
| {{< span >}}**強調**{{< /span >}} |
| {{% span %}}**強調**{{% /span %}}  |

という感じに `{{</* */>}}` と `{{%/* */%}}` で違いが生じる。

[Hugo] のテンプレート処理は文脈依存になっていて，同じ記述でもどの要素の中で書かれるか（`<head>` 要素か `<body>` 要素か，あるいは JavaScript か CSS か）で出力が違ったりする。
おそらく [shortcode] の `{{%/* */%}}` 記述でも同じように文脈依存で出力が変わるようにしたかったのだろう。

でも，私は `.Inner` 変数の展開を `{{</* */>}}` か `{{%/* */%}}` かで使い分けていたので，今回のアップデートで大ダメージを食らってしまったですよ。
しょうがないので [shortcode] を設計し直したけどね `orz`

今後の [shortcode] の運用方針は以下のような感じだろうか。

- `.Inner` の内容を markdown ドキュメントとして処理するかどうか明示的に制御する場合は `{{</* */>}}` を使う
- `.Inner` の処理を文脈依存で行う（[shortcode] を入れ子にするなどの）場合は `{{%/* */%}}` を使う

後方互換性が壊れる変更は，少なくとも最初はオプトアウトではなくオプトインで組み込めるようにして欲しい。

## Shortcode の入れ子ができてる

今回いろいろ弄っていて気がついたのだが，いつの間にか [shortcode] の入れ子ができるようになってたんだね。
いつのバージョンからだろう。

たとえば

```text
{{%/* div-box */%}}{{%/* ruby "Hello world!" */%}}こんにちは世界{{%/* /ruby */%}}{{%/* /div-box */%}}
```

と書くと

```html
<div class="box"><ruby><rb>こんにちは世界</rb><rp> (</rp><rt>Hello world!</rt><rp>) </rp></ruby></div>
```

{{% div-box %}}{{% ruby "Hello world!" %}}こんにちは世界{{% /ruby %}}{{% /div-box %}}

などとできるようになっていた。
入れ子記述は `{{</* */>}}` でも `{{%/* */%}}` でもできる。
`{{</* */>}}` と `{{%/* */%}}` 混在でもOK。

ちなみに `ruby` は自作の [shortcode] で中身はこんな感じ。

```html
<ruby><rb>{{ .Inner }}</rb><rp> (</rp><rt>{{ index .Params 0 }}</rt><rp>) </rp></ruby>
```

特に権利は主張しないので（するほどの内容じゃないし）自由に使ってください。

## Deprecated な変数・関数

[Hugo] 0.55 を起動すると以下のワーニングが出るようになった。

```text
$ hugo
WARN 2019/04/13 09:00:00 Page's .Hugo is deprecated and will be removed in a future release. Use the global hugo function.
WARN 2019/04/13 09:00:00 Page's .RSSLink is deprecated and will be removed in a future release. Use the Output Format's link, e.g. something like:
    {{ with .OutputFormats.Get "RSS" }}{{ .RelPermalink }}{{ end }}.
WARN 2019/04/13 09:00:00 Page's .GetParam is deprecated and will be removed in a future release. Use .Param or .Params.myParam.
```

テンプレート内で `.Hugo`, `.RSSLink` 変数および `.GetParam` 関数が使われていると上記ワーニングが出る。

### .Hugo の廃止と代替手段

（以前の内容を書き換えた）

`.Hugo` 変数は将来バージョンで削除されるようだ。
代替として（変数ではなく） `hugo` 関数が用意されている（Thanx [@peaceiris](https://twitter.com/piris314/status/1117092823159853057)）。
出力例は以下の通り。

{{< hugo-info >}}

この手が使えるなら `.Site` 変数も関数で用意してほしいものである。

### .RSSLink の廃止と代替手段

`.RSSLink` 変数も将来バージョンで削除されるようだ。
代替として `.AlternativeOutputFormats` および `.OutputFormats` 変数が使える。
たとえば `<head>` 要素内なら

```html
{{ with .Site.Home.AlternativeOutputFormats.Get "RSS" }}
  <link rel="alternate" href="{{ .Permalink }}" type="application/rss+xml" title="{{ $.Site.Title | plainify }}">
{{ end }}
```

と書けばいいし `<head>` 要素以外なら

```html
{{ with .OutputFormats.Get "RSS" }}
  <a href='{{ .RelPermalink }}' title='Feed'>Feed</a>
{{ end }}
```

などと書ける。
`.AlternativeOutputFormats` および `.OutputFormats` 変数はかなり応用範囲が広くて，たとえば私は[フィードを JSON 形式でも用意]({{< relref "./feed-with-json-format.md" >}} "JSON 形式で Feed を出力する")しているが，

```html
{{ with .Site.Home.AlternativeOutputFormats.Get "JSON" }}
  <link rel="alternate" href="{{ .Permalink }}" type="application/json" title="{{ $.Site.Title | plainify }}">
{{ end }}
```

と書けば簡単に `<head>` 要素に組み込める。

### .GetParam の廃止と代替手段

`.GetParam` 関数については随分前からアナウンスがあったので使っている人はいないと思うが `.Param` 関数で代替できる。

## ブックマーク

- [0.55.1: 3 Bug Fixes | Hugo](https://gohugo.io/news/0.55.1-relnotes/)
- [Hugo 0.55.2: Some Important Bug Fixes | Hugo](https://gohugo.io/news/0.55.2-relnotes/)
- [Hugo 0.55.3: A Few More Bug Fixes! | Hugo](https://gohugo.io/news/0.55.3-relnotes/)
- [Hugo 0.55.4: Some Bug Fixes | Hugo](https://gohugo.io/news/0.55.4-relnotes/)

- [.Hugo.Generator は廃止されるので hugo.Generator を使おう - Qiita](https://qiita.com/peaceiris/items/b6d611e184b2f28cc0ab)

- [Shortcodes で HTML コードを埋め込む]({{< relref "./shortcodes.md">}})

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[shortcode]: https://gohugo.io/extras/shortcodes/ "Shortcodes | Hugo"
