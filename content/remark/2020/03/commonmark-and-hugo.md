+++
title = "CommonMark と Hugo 0.68"
date =  "2020-03-29T15:04:41+09:00"
description = "てっきり yuin/goldmark パーサのバグだと思っていたのだが，どうやら CommonMark の仕様らしい。"
image = "/images/attention/kitten.jpg"
tags = [ "hugo", "markdown" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 [Hugo] 0.68.x がリリースされたのだが

- [Minify config and more! | Hugo](https://gohugo.io/news/0.68.0-relnotes/)
- [Hugo 0.68.1: 1 bug fix | Hugo](https://gohugo.io/news/0.68.1-relnotes/)
- [Hugo 0.68.2: A couple of Bug Fixes | Hugo](https://gohugo.io/news/0.68.2-relnotes/)
- [Hugo 0.68.3: A couple of Bug Fixes | Hugo](https://gohugo.io/news/0.68.3-relnotes/)

どうもまた markdown 周りの挙動が変わったようだ。
具体的には

```markdown
[string] を [][rune] に変換する。

[string]: http://golang.org/ref/spec#String_types
[rune]: http://blog.golang.org/strings "Strings, bytes, runes and characters in Go - The Go Blog"
```

をレンダリングすると

```html
<a href="http://golang.org/ref/spec#String_types">string</a> を <a href="http://blog.golang.org/strings" title="Strings, bytes, runes and characters in Go - The Go Blog"></a> に変換する。
```

{{< fig-quote >}}
<a href="http://golang.org/ref/spec#String_types">string</a> を <a href="http://blog.golang.org/strings" title="Strings, bytes, runes and characters in Go - The Go Blog"></a> に変換する。
{{< /fig-quote >}}

てな感じに `[]` 記述がまるっと抜け落ちて `<a>` 要素の中身も空になってしまう。
てっきり [yuin/goldmark] パーサのバグだと思っていたのだが，どうやらこれは [CommonMark] の仕様のようだ。

具体的には

```markdown
[foo][bar]

[bar]: http://bar
```

と記述すると

```html
<a href="http://bar">foo</a>
```

のように `[bar]` で生成される `<a>` 要素の中身を `foo` に置き換えるらしい。
じゃあ今までがバグだったってこと？

これを回避するには

```markdown {hl_lines=[1]}
\[foo][bar]

[bar]: http://bar
```

とエスケープすれば

```html
[foo]<a href="http://bar">bar</a>
```

{{< fig-quote >}}
[foo]<a href="http://bar">bar</a>
{{< /fig-quote >}}

と意図通りにレンダリングしてくれる。
なので最初の記述も

```markdown {hl_lines=[1]}
[string] を \[][rune] に変換する。

[string]: http://golang.org/ref/spec#String_types
[rune]: http://blog.golang.org/strings "Strings, bytes, runes and characters in Go - The Go Blog"
```

とすれば

```html
<a href="http://golang.org/ref/spec#String_types">string</a> を []<a href="http://blog.golang.org/strings" title="Strings, bytes, runes and characters in Go - The Go Blog">rune</a> に変換する。
```

{{< fig-quote >}}
<a href="http://golang.org/ref/spec#String_types">string</a> を []<a href="http://blog.golang.org/strings" title="Strings, bytes, runes and characters in Go - The Go Blog">rune</a> に変換する。
{{< /fig-quote >}}

とできた。

私のブログでは影響が出たのが（結果的には）一箇所のみだったので直すのは簡単だったが，特に [Go 言語]では `[]` はスライスを意味する記述なので，今後は気をつけないとなぁ。
まぁ `<code>` 要素（\``...`\`）で囲むのが無難か。

改めて [CommonMark] の仕様を眺めてみたが，不可思議な記述が沢山ある。
たとえば

```markdown
[bar][]

[bar]: http://bar
```

と書くと

```html
<a href="http://bar">bar</a>
```

てな感じに，これまた `[]` が吸い込まれる。
しかも

```markdown
[bar][foo]

[bar]: http://bar
```

とすると，今度は `<a>` 要素が外れてただの

```html
[bar][foo]
```

となる。
なんだよそれ `orz`

雑すぎんだろ！ 一度ちゃんと文法を整理したほうがいいんじゃないのか？

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[CommonMark]: https://commonmark.org/
[yuin/goldmark]: https://github.com/yuin/goldmark/ "yuin/goldmark: A markdown parser written in Go. Easy to extend, standard(CommonMark) compliant, well structured."
[Go 言語]: https://golang.org/ "The Go Programming Language"
