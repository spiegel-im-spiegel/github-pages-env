+++
title = "Markdown パーサ blackfriday.v2 で遊ぶ"
date = "2018-02-17T15:44:07+09:00"
description = "調子に乗ってプレビュー・ツールも作ってみた。"
image = "/images/attention/go-code2.png"
tags        = [ "golang", "programming", "markdown", "template", "html", "tools" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = true
  mermaidjs = false
+++

この前 [markdown テキスト変換ツールの物色]({{< relref "remark/2018/02/from-markdown-to-html.md" >}})を行っていたのだが，この中で [blackfriday] パッケージがなかなか面白そうなのでちょっと遊んでみることにした。

## [blackfriday].v2

現在 [blackfriday] は v2 系が最新版で，作者も v2 を推奨しているみたいなのだが，軽くググってみるかぎり v2 を使っているところを見かけない。
このブログのサイト生成ツールである [Hugo] も 1.x 系みたいだし。
おそらくインタフェースが違うだけで（HTML に変換する限りは）性能的にはあまり変わらないため需要がないのかもしれない。

v2 系は [GitHub] ではなく以下から取得するのがいいらしい。

```text
$ go get -u gopkg.in/russross/blackfriday.v2
```

[dep] を使うなら `dep ensure -add` コマンドで取り込むか [`Gopkg.toml`] ファイルに以下の記述を追加する。

```toml
[[constraint]]
  name = "gopkg.in/russross/blackfriday.v2"
  version = "~2.0.0"
```

HTML への変換はこんな感じに書ける。

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	md, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

    //HTMLFlags and Renderer
	htmlFlags := blackfriday.CommonHTMLFlags         //UseXHTML | Smartypants | SmartypantsFractions | SmartypantsDashes | SmartypantsLatexDashes
	htmlFlags |= blackfriday.FootnoteReturnLinks     //Generate a link at the end of a footnote to return to the source
	htmlFlags |= blackfriday.SmartypantsAngledQuotes //Enable angled double quotes (with Smartypants) for double quotes rendering
	htmlFlags |= blackfriday.SmartypantsQuotesNBSP   //Enable French guillemets 損 (with Smartypants)
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: htmlFlags, Title: "", CSS: ""})

	//Extensions
	extFlags := blackfriday.CommonExtensions //NoIntraEmphasis | Tables | FencedCode | Autolink | Strikethrough | SpaceHeadings | HeadingIDs | BackslashLineBreak | DefinitionLists
	extFlags |= blackfriday.Footnotes        //Pandoc-style footnotes
	extFlags |= blackfriday.HeadingIDs       //specify heading IDs  with {#id}
	extFlags |= blackfriday.Titleblock       //Titleblock ala pandoc
	extFlags |= blackfriday.DefinitionLists  //Render definition lists

	html := blackfriday.Run(md, blackfriday.WithExtensions(extFlags), blackfriday.WithRenderer(renderer))
	fmt.Println(string(html))
}
```

[`blackfriday`]`.WithExtensions()` 関数および [`blackfriday`]`.WithRenderer()` 関数は [Functional Options パターン]({{< relref "golang/functional-options-pattern.md" >}} "インスタンスの生成と Functional Options パターン")の応用で任意に設定できる。

v1.x 系に比べて HTML レンダリング・オプションの指定が面倒くさい感じになっているが，これは [`blackfriday`]`.Renderer` インタフェースに合わせた別のレンダリング・パッケージを使えるようにするためらしい。
HTML 変換以外のレンダリング・パッケージとしては $\mathrm{\LaTeX}$ への変換パッケージがあるようだ。

- [Ambrevar/blackfriday-latex: A LaTeX renderer for the Blackfriday Markdown Processor](https://github.com/Ambrevar/blackfriday-latex)

数式表現はできれば [MathJax] 互換にしてほしかったが，まぁいいか。
これもそのうち試してみたい。

## 調子に乗ってプレビュー・ツールを作ってみた

- [spiegel-im-spiegel/markdown-preview: Markdown Preview Tool](https://github.com/spiegel-im-spiegel/markdown-preview)

一応[バイナリも用意](https://github.com/spiegel-im-spiegel/markdown-preview/releases/latest)している。
使い方はこんな感じ。

```text
$ markdown-preview -h
Processing Markdown by Golang

Usage:
  markdown-preview [flags]
  markdown-preview [command]

Available Commands:
  help        Help about any command
  proc        Processing Markdown
  version     Print the version number of markdown-preview

Flags:
  -h, --help   help for markdown-preview

Use "markdown-preview [command] --help" for more information about a command.
```

今のところは HTML への変換のみサポートしている。

```text
$ markdown-preview proc -h
Processing Markdown

Usage:
  markdown-preview proc [flags] [markdown file]

Flags:
  -c, --css string      CSS file URL (with --page option)
  -g, --github          use GitHub Markdown API
  -h, --help            help for proc
  -l, --line-break      translate newlines into line breaks
  -o, --output string   output file path
  -p, --page            generate a complete HTML page
  -s, --sanitize        sanitize untrusted content
```

### [GitHub] Markdown API

`--github` オプションを使うと [blackfriday] パッケージではなく [GitHub] Markdown API を使う。

- [Markdown | GitHub Developer Guide](https://developer.github.com/v3/markdown/)

[GitHub] Markdown API を利用するには Google による [go-github]/github パッケージが便利である。
こんな感じに書ける。

```go
import (
	"context"

	"github.com/google/go-github/github"
)

func renderWithGitHub(md []byte) ([]byte, error) {
	client := github.NewClient(nil)
	opt := &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}
	body, _, err := client.Markdown(context.Background(), string(md), opt)
	return []byte(body), err
}
```

### テンプレートを使ったページの出力

`--page` オプションを使うと完全なページを出力する。
 [blackfriday] パッケージなら [`blackfriday`]`.CompletePage` フラグを付加することで完全なページを出力してくれるが，今回はテンプレートを使ってページを出力するようにした。

[Go 言語]標準のテンプレートパッケージには [`text/template`] と [`html/template`] の2つがある。
[`html/template`] は `<` や `>` などの特殊文字を適切に変換してくれるので良いのだが，今回は HTML テキストをまるっと埋め込むので（勝手に sanitizing されては困るので） [`text/template`] のほうを使うことにした。

また，テンプレートファイルは [go-assets を使って]({{< relref "golang/using-go-assets.md" >}} "go-assets でシングルバイナリにまとめる")コードに埋め込むことにした。
いやぁ，[勉強して]({{< relref "golang/using-go-assets.md" >}} "go-assets でシングルバイナリにまとめる")おいてよかった。

用意したテンプレートファイルはこんな感じ。

```html
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="generator" content="markdown-preview">{{with .CSS }}
  <link rel="stylesheet" type="text/css" href="{{ . }}">{{ end }}
  <title>{{ .Title }}</title>
</head>
<body>

{{ .Body }}

</body>
</html>
```

[Hugo] でテンプレートの扱いにはすっかり慣れてしまったので，このくらいは楽勝である（笑）

テンプレートに埋め込むためのデータセットはこんな感じで用意した。

```go
type PageData struct {
	Title string
	CSS   string
	Body  string
}
```

これで，テンプレート処理は以下のように書ける。

```go
//Render returns HTML page text with template
func (p *PageData) Render() ([]byte, error) {
	f, err := data.Assets.Open("/template.html")
	if err != nil {
		return nil, err
	}
	tmpData := &bytes.Buffer{}
	io.Copy(tmpData, f)

	t, err := template.New("Markdown Processing").Parse(tmpData.String())
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, p)
	return buf.Bytes(), err
}
```

### 今後の予定

というほどではないが，簡易 Web サービスでプレビューできるようにしたいな，と。
[mattn/mkup] みたいな感じ。
ただ [Go 言語]による Web アプリケーションは最近勉強を始めたばかりなので当分先だろうけど（その前に [gpgpdump] の [RFC 4880bis] 対応しろって）。

個人的に CLI のフィルタ・コマンド大好きなので，そういうものばっかり（主に自分用に）作ってるが，監視系のツールや繰り返し処理が多いものについては簡易 Web サービスを立ててブラウザ上で作業するのもアリなんじゃないかと思ったりしている。

## ブックマーク

- [Glide から Dep への移行を検討する]({{< relref "golang/consider-switching-from-glide-to-dep.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[blackfriday]: https://github.com/russross/blackfriday "russross/blackfriday: Blackfriday: a markdown processor for Go"
[`blackfriday`]: https://github.com/russross/blackfriday "russross/blackfriday: Blackfriday: a markdown processor for Go"
[blackfriday-tool]: https://github.com/russross/blackfriday-tool "russross/blackfriday-tool: Blackfriday command-line tool"
[go-github]: https://github.com/google/go-github "google/go-github: Go library for accessing the GitHub API"
[GitHub]: https://github.com/
[dep]: https://golang.github.io/dep/ "dep · Dependency management for Go"
[`Gopkg.toml`]: https://golang.github.io/dep/docs/Gopkg.toml.html "Gopkg.toml · dep"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[MathJax]: https://www.mathjax.org/ "MathJax | Beautiful math in all browsers."
[mattn/mkup]: https://github.com/mattn/mkup "mattn/mkup: Portable Markdown Previewer"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[`text/template`]: https://golang.org/pkg/text/template/ "template - The Go Programming Language"
[`html/template`]: https://golang.org/pkg/html/template/ "template - The Go Programming Language"
