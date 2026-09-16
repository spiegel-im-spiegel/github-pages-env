+++
title = "テキストから HTML タグを除去する"
date =  "2026-09-16T12:17:57+09:00"
description = "正規表現でタグを除去する / Tokenizer を使ってタグを除去する"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "html", "regular-expression" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

たとえば

```html
This is a <strong>sample</strong> text
with <a href="#">HTML</a> tags.
```

というテキストデータが与えられた場合に `<strong>` や `<a>` といった HTML タグを除去することを考える。

## 正規表現でタグを除去する

単純に `<...>` の文字列パターンを除去できればいいのであれば正規表現が使える。
こんな感じ。

```go
package striptag

import (
    "html"
    "regexp"
)

// StripTagsRegexp removes all HTML tags from the input string and unescapes any HTML entities.
func StripTagsRegexp(s string) string {
    re := regexp.MustCompile(`<[^>]*>`)
    return html.UnescapeString(re.ReplaceAllString(s, ""))
}
```

実際に動かしてみよう。
入力テキストを

```go
package example

const (
    Text1 = `This is a <strong>sample</strong> text
with <a href="#">HTML</a> tags.`
)
```

と定義しておいて，次のように `main()` 関数を定義する。

```go
//go:build ignore

package main

import (
    "striptag"
    "striptag/example"
)

func main() {
    println(striptag.StripTagsRegexp(example.Text1))
}
```

これを実行すると

```text
$ go run sample1a.go
This is a sample text
with HTML tags.
```

と出力される。
ここまでは問題ない。
問題は入力テキストに異物が混ざってた場合。
たとえば

```go
const (
    Text2 = `This is a <strong>sample</strong> text
with <a href="#">HTML</a> tags.<script>alert("XSS")</script>`
)
```

のような感じ（末尾に `<script>` タグがある）。
これを `StripTagsRegexp()` 関数で処理すると

```text
This is a sample text
with HTML tags.alert("XSS")
```

のように `<script>` タグの中身が露出してしまう。

## Tokenizer を使ってタグを除去する

これを正規表現で対処するのは（やれないことはないだろうが）かなり面倒な気がするので，やり方を変えて [`golang.org/x/net/html`] パッケージの `Tokenizer` を使うことにする。
こんな感じかな。

```go
package striptag

import (
	"errors"
	"html"
	"io"
	"strings"

	ghtml "golang.org/x/net/html"
)

// skipTags defines the HTML tags whose content should be skipped when stripping tags.
var skipTags = map[string]bool{
	"script":   true,
	"style":    true,
	"noscript": true,
	"iframe":   true,
	"object":   true,
	"embed":    true,
	"textarea": true,
	"title":    true,
}

// StripTagsTokenizer removes HTML tags from the input string while skipping
// the content of certain tags defined in skipTags.
func StripTagsTokenizer(s string) (string, error) {
	tokenizer := ghtml.NewTokenizer(strings.NewReader(s))
	var b strings.Builder
	b.Grow(len(s))
	skipDepth := 0

	for {
		tt := tokenizer.Next() // get the next token type
		switch tt {
		case ghtml.ErrorToken: // handle error token
			if errors.Is(tokenizer.Err(), io.EOF) {
				return b.String(), nil // return the accumulated text at the end of input
			}
			return "", tokenizer.Err()
		case ghtml.StartTagToken: // handle start tag token
			t := tokenizer.Token()
			if skipTags[t.Data] {
				skipDepth++ // start skipping content of this tag
			}
		case ghtml.EndTagToken: // handle end tag token
			t := tokenizer.Token()
			if skipDepth > 0 && skipTags[t.Data] {
				skipDepth-- // stop skipping one level of skipped content
			}
		case ghtml.TextToken: // handle text token
			if skipDepth == 0 {
				b.WriteString(html.UnescapeString(string(tokenizer.Text())))
			}
		}
	}
}
```

これで `<script> ... </script>` の中身をスキップできる筈。
試してみよう。
`main()` 関数を

```go { hl_lines=["13-18"]}
//go:build ignore

package main

import (
	"fmt"
	"os"
	"striptag"
	"striptag/example"
)

func main() {
	out, err := striptag.StripTagsTokenizer(example.Text2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(out)
}
```

と書き直して実行する。

```text
$ go run sample2b.go
This is a sample text
with HTML tags.
```

おー。
上手く行ったかな。

### 改行タグを改行コードに変換する

応用として，前節の `StripTagsTokenizer()` 関数に `<br>` タグを改行コードに変換する機能を追加する。

```go {hl_lines=["1-7", "30-37"]}
// changeBrToNewline handles the <br> tag by converting it to a newline character in the output.
func changeBrToNewline(t ghtml.Token, b *strings.Builder) {
	if t.Data == "br" {
		// Handle <br> as a line break.
		b.WriteByte('\n')
	}
}

// StripTagsTokenizer removes HTML tags from the input string while skipping
// the content of certain tags defined in skipTags.
func StripTagsTokenizer(s string) (string, error) {
	tokenizer := ghtml.NewTokenizer(strings.NewReader(s))
	var b strings.Builder
	b.Grow(len(s))
	skipDepth := 0

	for {
		tt := tokenizer.Next() // get the next token type
		switch tt {
		case ghtml.ErrorToken: // handle error token
			if errors.Is(tokenizer.Err(), io.EOF) {
				return b.String(), nil // return the accumulated text at the end of input
			}
			return "", tokenizer.Err()
		case ghtml.StartTagToken: // handle start tag token
			t := tokenizer.Token()
			if skipTags[t.Data] {
				skipDepth++ // start skipping content of this tag
			}
			if skipDepth == 0 {
				changeBrToNewline(t, &b)
			}
		case ghtml.SelfClosingTagToken: // handle self-closing tag token
			t := tokenizer.Token()
			if skipDepth == 0 {
				changeBrToNewline(t, &b)
			}
		case ghtml.EndTagToken: // handle end tag token
			t := tokenizer.Token()
			if skipDepth > 0 && skipTags[t.Data] {
				skipDepth-- // stop skipping one level of skipped content
			}
		case ghtml.TextToken: // handle text token
			if skipDepth == 0 {
				b.WriteString(html.UnescapeString(string(tokenizer.Text())))
			}
		}
	}
}
```

修正した `StripTagsTokenizer()` 関数を使って以下のテキストを処理する。

```go
const (
	Text3 = `This is a <strong>sample</strong> text<br>with <a href="#">HTML</a> tags.<script>alert("XSS")</script>`
)
```

処理結果はこんな感じ。

```text
This is a sample text
with HTML tags.
```

`text<br>with` の `<br>` タグは改行に変換されていることが確認できた。

ここまできたら `<p>` のようなブロック要素にも対応したいところだけど，実際にブロック要素か否かはスタイル情報に依存するため，今回は割愛する。

[Go]: https://go.dev/
[`golang.org/x/net/html`]: https://pkg.go.dev/golang.org/x/net/html "html package - golang.org/x/net/html - Go Packages"

## 参考図書

{{< linkcard "9e8d8f717b1d1d85b0da7f07cb10a83b78d4227f" >}} <!-- プログラミング言語Go -->
{{< review-paapi "B0CFL1DK8Q" >}} <!-- Go言語 100Tips -->
{{< review-paapi "B0DNYMMBBQ" >}} <!-- Go言語で学ぶ並行プログラミング -->
