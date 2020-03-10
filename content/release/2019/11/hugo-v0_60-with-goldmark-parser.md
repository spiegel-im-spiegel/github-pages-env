+++
title = "Hugo v0.60 から既定の Markdown パーサが Goldmark になったようだ"
date =  "2019-11-28T22:45:04+09:00"
description = "今回の変更に伴い色々と変わった部分もあるので，覚え書きとして記しておく。"
image = "/images/attention/tools.png"
tags  = [ "hugo", "markdown", "site", "syntax-highlight", "shortcodes" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Hugo] v0.60 がリリースされた。

- [Release v0.60.0 · gohugoio/hugo · GitHub](https://github.com/gohugoio/hugo/releases/tag/v0.60.0)
    - [Now CommonMark Compliant! | Hugo](https://gohugo.io/news/0.60.0-relnotes/)
- [Release v0.60.1 · gohugoio/hugo · GitHub](https://github.com/gohugoio/hugo/releases/tag/v0.60.1)
    - [Hugo 0.60.1: A couple of Bug Fixes | Hugo](https://gohugo.io/news/0.60.1-relnotes/)

[Hugo] v0.60 から既定の Markdown パーサが [yuin/goldmark] になったようだ。

{{< fig-quote type="markdown" title="Now CommonMark Compliant!" link="https://gohugo.io/news/0.60.0-relnotes/" lang="en" >}}
{{< quote >}}[Goldmark](https://github.com/yuin/goldmark/) by [@yuin](https://github.com/yuin) is now the new default library used for Markdown in Hugo. It's CommonMark compliant and GitHub flavored, and both fast and flexible{{< /quote >}}.
{{< /fig-quote >}}

このサイトで試してみたが確かに速くなった！
こうやって柔軟に機能を入れ替えるフットワークの軽さが [Hugo] のいいところだよねぇ（たまに後方互換性が壊れるのが問題だが）。
まぁ，コミュニティが活況だからできるんだろうけど。

今回の変更に伴い色々と変わった部分もあるので，以降に覚え書きとして記しておく。

## パーサの選択

Markdown パーサについては，グローバル設定で新しい [yuin/goldmark] （既定）と旧来の [russross/blackfriday] のいずれかを選択できる。
旧来の [russross/blackfriday] を使いたいなら

```toml
[markup]
  defaultMarkdownHandler = "blackfriday"
```

とすればよい（`config.toml` に TOML 形式で記述する場合）。

## パーサの設定

以降は [yuin/goldmark] パーサを前提に話を進める。

グローバル設定でパーサの設定を記述できる。
既定値は以下の通り。

```toml
[markup]
  defaultMarkdownHandler = "goldmark"
  [markup.goldmark]
    [markup.goldmark.extensions]
      definitionList = true
      footnote = true
      linkify = true
      strikethrough = true
      table = true
      taskList = true
      typographer = true
    [markup.goldmark.parser]
      attribute = true
      autoHeadingID = true
    [markup.goldmark.renderer]
      hardWraps = false
      unsafe = false
      xHTML = false
```

通常はこのままで問題ないが，いくつか注意点がある。

## HTML 記述が混在する場合

[russross/blackfriday] では HTML 形式の記述を許容していたが [yuin/goldmark] では HTML の記述は削除されるようだ[^bf1]。
しかし，これでは複雑なレイアウトの表などを書きたい場合に困るので，設定を以下のように変更する。

[^bf1]: [russross/blackfriday] を使う場合でもグローバル設定で `skipHTML` 項目を `true` にすれば HTML 記述が削除されるようだ。

```toml {hl_lines=[2]}
[markup.goldmark.renderer]
  unsafe = true
```

これで HTML 形式の記述を許容してくれる。

この設定は入れ子の [shortcode] にも影響してくるので，これまでのバージョンに近い出力が欲しいなら `unsafe = true` にしておくのがいいかもしれない。

## Inline Code の挙動に注意

{{< div-box type="markdown" >}}
**【2020-01-02 追記】**
この件は [Hugo] 0.61 で修正されたらしい。 [yuin/goldmark] のバグだったようだ。
なので，また元に戻した。
なんだかなぁ `orz`

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[yuin/goldmark]: https://github.com/yuin/goldmark/ "yuin/goldmark: A markdown parser written in Go. Easy to extend, standard(CommonMark) compliant, well structured."
{{< /div-box >}}

[CommonMark] に準拠したせいかどうかは分からないが inline code の挙動が微妙に変わっているので注意が必要である。

たとえば以下のような記述があるとする。

```markdown
[`fmt`]`.println()` 関数で “Hello world” と出力する。

[`fmt`]: https://golang.org/pkg/fmt/
```

これを処理して

```html
<a href="https://golang.org/pkg/fmt/"><code>fmt</code></a><code>.println()</code> 関数で “Hello world” と出力する。
```

となってくれることを期待するのだが（つか以前のバージョンではそうなっていた），実際には

```html
[<code>fmt`]</code>.println()` 関数で “Hello world” と出力する。
```

となる。
なんでやねん `orz`

しょうがないので

```markdown
[`fmt`].`println()` 関数で “Hello world” と出力する。
```

と書き直しましたよ。
これで

```html
<a href="https://golang.org/pkg/fmt/"><code>fmt</code></a>.<code>println()</code> 関数で “Hello world” と出力する。
```

と（ちょっと不格好だが）なった。
[ATOM] で複数ファイルの内容を一括置換できてよかったね（泣）

## フェンス付きコード・ブロックのハイライト表示

[Hugo] ではフェンス付きコード・ブロックのハイライト表示には [alecthomas/chroma] を使っているが，グローバル設定でより細かい設定が可能になった。
以下は既定値。

```toml
[markup.highlight]
  codeFences = true
  hl_Lines = ""
  lineNoStart = 1
  lineNos = false
  lineNumbersInTable = true
  noClasses = true
  style = "monokai"
  tabWidth = 4
```

従来の `pygmentsUseClasses` や `pygmentsCodeFences` 等も効いているようだが，こちらの記述に切り替えたほうがいいだろう。
`style` のサンプルは以下をどうぞ。

- [Short snippets](https://xyproto.github.io/splash/docs/all.html)
- [Long snippets](https://xyproto.github.io/splash/docs/longer/all.html)

`noClasses` を `false` にすればスタイル指定を（タグへの直書きではなく）クラスで行うようになる。
スタイル指定用の CSS ファイルの作成はこんな感じでできる。

```text
$ hugo gen chromastyles --style=tango > chroma-styles.css
```

行のハイライト等を行うには `highlight` の組み込み [shortcode] を使う必要があったが， v0.60 からは

{{< highlight text >}}
```go {hl_lines=[1, "5-7"]}
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```
{{< /highlight >}}

などと書けば

```go  {hl_lines=[1, "5-7"]}
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```

と表示されるようになった。
やっほい！

## 脚注について

以前は脚注部分について

```html
<div class="footnotes">
...
</div>
```

と `<div>` 要素になっていたが， v0.60 からは

```html
<section class="footnotes" role="doc-endnotes">
...
</section>
```

と `<section>` 要素になっている。
こちらのほうが合理的だよね。
ただし，脚注のスタイルを CSS 等で制御している場合は注意が必要である。

## ブックマーク

- [Configure Markup | Hugo](https://gohugo.io/getting-started/configuration-markup/)
- [Syntax Highlighting | Hugo](https://gohugo.io/content-management/syntax-highlighting/)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[CommonMark]: https://commonmark.org/
[yuin/goldmark]: https://github.com/yuin/goldmark/ "yuin/goldmark: A markdown parser written in Go. Easy to extend, standard(CommonMark) compliant, well structured."
[russross/blackfriday]: https://github.com/russross/blackfriday "russross/blackfriday: Blackfriday: a markdown processor for Go"
[alecthomas/chroma]: https://github.com/alecthomas/chroma "alecthomas/chroma: A general purpose syntax highlighter in pure Go"
[shortcode]: https://gohugo.io/extras/shortcodes/ "Shortcodes | Hugo"
[ATOM]: https://atom.io/
