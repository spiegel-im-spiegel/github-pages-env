+++
title = "Chroma Syntax Highlighting"
date =  "2017-09-25T22:57:09+09:00"
description = "今回のアップデートの目玉は純正 Golang 製の Chroma を組み込んで Syntax Highlight に対応したことだろう。"
tags        = [ "hugo", "syntax-highlight", "shortcodes", "markdown" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Hugo] 0.28 がリリースされた。

- [Hugo | Hugo 0.28: High-speed Syntax Highlighting!](https://gohugo.io/news/0.28-relnotes/)

今回のアップデートの目玉は [Golang] 製の  [Chroma] を組み込んで Syntax Highlight に対応したことだろう。
それまでは [Pygments] のようなツールを使うか [highlight.js] のようなスクリプトを噛ませるかだったのだが，ようやく [Hugo] が自前で用意できるようになった[^sh1]。

[^sh1]: いや，簡易的なものなら以前からあったのだが，吐き出すコードがイマイチで使う気にならなかったんだよねぇ。

Syntax Highlight の設定方法については[マニュアル](https://gohugo.io/documentation/ "Hugo | Hugo Documentation")の以下のページが参考になる。

- [Hugo | Syntax Highlighting](https://gohugo.io/content-management/syntax-highlighting/)

## Syntax Highlight を有効にする

[Chroma] Syntax Highlight を有効にするには設定ファイル（`config.toml` 等）で以下の項目を有効にすればよい。

{{< highlight ini >}}
pygmentsUseClasses = true
{{< /highlight  >}}

さらにハイライト表示の CSS ファイルを用意する。
これは [Hugo] コマンドで生成できる。

```text
$ hugo gen chromastyles --style=tango > chroma-styles.css
```

指定できる `style` の詳細は以下が参考になる。

- [Pygments style gallery!](https://help.farbox.com/pygments.html)

もちろん，生成した CSS ファイルは HTML の `<head>` 要素内で指定すること。

```html
<link rel="stylesheet" href="/css/chroma-styles.css">
```

## ハイライト表示

[Hugo] でハイライトを指定するには以下の [shortcode] を使う。

```text
{{</* highlight go */>}}
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
{{</* /highlight */>}}
```

実際の表示は以下の通り。

{{< highlight go >}}
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
{{< /highlight >}}

行番号を表示するには以下のように指定する

```text
{{</* highlight go "linenos=inline,linenostart=1" */>}}
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
{{</* /highlight */>}}
```

{{< highlight go "linenos=inline,linenostart=1" >}}
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
{{< /highlight >}}

特定の行を強調するには以下のように指定する。

```text
{{</* highlight go "linenos=inline,linenostart=1,hl_lines=5-7" */>}}
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
{{</* /highlight */>}}
```

{{< highlight go "linenos=inline,linenostart=1,hl_lines=5-7" >}}
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
{{< /highlight >}}

一方，設定ファイルで `pygmentsCodeFences` を有効にすることで GitHub Flavored Markdown の Syntax Highlight も使えるようになる。

{{< highlight text >}}
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```
{{< /highlight >}}

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[Chroma]: https://github.com/alecthomas/chroma "alecthomas/chroma: A general purpose syntax highlighter in pure Go"
[Pygments]: http://pygments.org/
[highlight.js]: https://highlightjs.org/
[Golang]: https://golang.org/ "The Go Programming Language"
[shortcode]: https://gohugo.io/content-management/shortcodes/ "Hugo | Shortcodes"
