+++
title = "Markdown テキスト変換ツールの物色"
date = "2018-02-15T22:40:48+09:00"
update = "2018-02-17T13:42:28+09:00"
description = "どうせならビルドも導入も簡単な Go 言語製のツールがいいな，と思ってたのだが，結構あってビックリした。"
image = "/images/attention/kitten.jpg"
tags        = [ "tools", "markdown", "html", "golang" ]

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
  mathjax = false
  mermaidjs = false
+++

Markdown テキストを HTML 等の他のドキュメント形式に変換するツールは色々あるけれど[^cnv1]，今のところは大げさなドキュメントフレームワークを組むつもりは全くなくて（このサイトで使ってる [Hugo] で必要十分だし）「もっとカジュアルに変換できるツールがないかなぁ」とググってみたのですよ。
どうせならビルドも導入も簡単な [Go 言語]製のツールがいいな，と思ってたのだが，結構あってビックリした。
まぁみんな同じようなことを考えてるってことかねぇ。

[^cnv1]: オリジナルの markdown パーサって Perl だったっけ？ うろ覚えだけど。以前に [node.js 上で markdown から PDF に変換する方法は調べた](https://qiita.com/spiegel-im-spiegel/items/cdea4a17b44fb59d6320 "Markdown ファイルを PDF に変換する。 - Qiita")けどスクリプト言語は環境を整えるのが面倒なのとモジュール間の依存関係が煩わしいのであまり使いたくない。仕事ならやるけど。とりあえず HTML 形式に変換しておけば，あとは LaTeX でも PDF でもどうとでもなるし，そもそも今回はそこまでやる気はない。

## [russross/blackfriday]

まずは [russross/blackfriday]。

名前はアレだけど，割とちゃんとしたパーサみたい[^hg1]。
同じ作者によるコマンドライン・ツール [russross/blackfriday-tool] も提供されているのですぐ使える。

[^hg1]: このブログでも利用している [Hugo] も markdown パーサとして [russross/blackfriday] を使っているようである。

[russross/blackfriday-tool] はバイナリ提供はされてないようなので（つか `main.go` ひとつだけだし`w`），素直に `go get` コマンドでインストールする。

```text
$ go get -v github.com/russross/blackfriday-tool
```

ヘルプを見るとこんな感じ。

```text
$ blackfriday-tool -h
Blackfriday Markdown Processor v1.5
Available at http://github.com/russross/blackfriday

Copyright © 2011 Russ Ross <russ@russross.com>
Distributed under the Simplified BSD License
See website for details

Usage:
  blackfriday-tool [options] [inputfile [outputfile]]

Options:
  -cpuprofile string
        Write cpu profile to a file
  -css string
        Link to a CSS stylesheet (implies -page)
  -fractions
        Use improved fraction rules for smartypants (default true)
  -latex
        Generate LaTeX output instead of HTML
  -latexdashes
        Use LaTeX-style dash rules for smartypants (default true)
  -page
        Generate a standalone HTML page (implies -latex=false)
  -repeat int
        Process the input multiple times (for benchmarking) (default 1)
  -smartypants
        Apply smartypants-style substitutions (default true)
  -toc
        Generate a table of contents (implies -latex=false)
  -toconly
        Generate a table of contents only (implies -toc)
  -xhtml
        Use XHTML-style tags in HTML output (default true)
```

入力データは以下のものを用意してみた。

```markdown
# Go 1.8.7 および 1.9.4 がリリース

[Go 言語]コンパイラの 1.8.7 および 1.9.4 がリリースされている。
このバージョンでは脆弱性 [CVE-2018-6574] の修正が行われている。

## 脆弱性の内容

`#cgo` ディレクティブを含む [Go 言語]コードをビルドする際に任意の処理を実行される可能性がある。

> When cgo is enabled, the build step during “go get” invokes the host C compiler, gcc or clang, adding compiler flags specified in the Go source files.
> Both gcc and clang support a plugin mechanism in which a shared-library plugin is loaded into the compiler, as directed by compiler flags.
> This means that a Go package repository can contain an attack.so file along with a Go source file that says (for example) `// #cgo CFLAGS: -fplugin=attack.so`, causing the attack plugin to be loaded into the host C compiler during the build.
> Gcc and clang plugins are completely unrestricted in their access to the host system.

## 影響度（CVSS）

「[CVE-2018-6574 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2018-6574)」より。

**CVSSv3 基本評価値 8.3 (`CVSS:3.0/AV:N/AC:H/PR:N/UI:R/S:C/C:H/I:H/A:H`)**

|                            基本評価基準 | 評価値            |
| ---------------------------------------:|:----------------- |
|                        攻撃元区分（AV） | ネットワーク（N） |
|                  攻撃条件の複雑さ（AC） | 高（H）           |
|                  必要な特権レベル（PR） | 不要（N）         |
|                  ユーザ関与レベル（UI） | 要（R）           |
|                           スコープ（S） | 変更あり（C）     |
| 情報漏えいの可能性（機密性への影響, C） | 高（H）           |
| 情報改ざんの可能性（完全性への影響, I） | 高（H）           |
|   業務停止の可能性（可用性への影響, A） | 高（H）           |

CVSS については[解説ページ](http://text.baldanders.info/remark/2015/cvss-v3-metrics-in-jvn/ "JVN が CVSSv3 による脆弱性評価を開始 — しっぽのさきっちょ | text.Baldanders.info")を参照のこと。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[CVE-2018-6574]: https://cve.mitre.org/cgi-bin/cvename.cgi?name=2018-6574
```

これを変換する。

```text
$ blackfriday-tool -page -xhtml=false input.md > output.html
```

結果はこんな感じになった。

```html
<!DOCTYPE html>
<html>
<head>
  <title>Go 1.8.7 および 1.9.4 がリリース</title>
  <meta name="GENERATOR" content="Blackfriday Markdown Processor v1.5">
  <meta charset="utf-8">
</head>
<body>

<h1>Go 1.8.7 および 1.9.4 がリリース</h1>

<p><a href="https://golang.org/" title="The Go Programming Language">Go 言語</a>コンパイラの 1.8.7 および 1.9.4 がリリースされている。
このバージョンでは脆弱性 <a href="https://cve.mitre.org/cgi-bin/cvename.cgi?name=2018-6574">CVE-2018-6574</a> の修正が行われている。</p>

<h2>脆弱性の内容</h2>

<p><code>#cgo</code> ディレクティブを含む <a href="https://golang.org/" title="The Go Programming Language">Go 言語</a>コードをビルドする際に任意の処理を実行される可能性がある。</p>

<blockquote>
<p>When cgo is enabled, the build step during “go get” invokes the host C compiler, gcc or clang, adding compiler flags specified in the Go source files.
Both gcc and clang support a plugin mechanism in which a shared-library plugin is loaded into the compiler, as directed by compiler flags.
This means that a Go package repository can contain an attack.so file along with a Go source file that says (for example) <code>// #cgo CFLAGS: -fplugin=attack.so</code>, causing the attack plugin to be loaded into the host C compiler during the build.
Gcc and clang plugins are completely unrestricted in their access to the host system.</p>
</blockquote>

<h2>影響度（CVSS）</h2>

<p>「<a href="https://access.redhat.com/security/cve/cve-2018-6574">CVE-2018-6574 - Red Hat Customer Portal</a>」より。</p>

<p><strong>CVSSv3 基本評価値 8.3 (<code>CVSS:3.0/AV:N/AC:H/PR:N/UI:R/S:C/C:H/I:H/A:H</code>)</strong></p>

<table>
<thead>
<tr>
<th align="right">基本評価基準</th>
<th align="left">評価値</th>
</tr>
</thead>

<tbody>
<tr>
<td align="right">攻撃元区分（AV）</td>
<td align="left">ネットワーク（N）</td>
</tr>

<tr>
<td align="right">攻撃条件の複雑さ（AC）</td>
<td align="left">高（H）</td>
</tr>

<tr>
<td align="right">必要な特権レベル（PR）</td>
<td align="left">不要（N）</td>
</tr>

<tr>
<td align="right">ユーザ関与レベル（UI）</td>
<td align="left">要（R）</td>
</tr>

<tr>
<td align="right">スコープ（S）</td>
<td align="left">変更あり（C）</td>
</tr>

<tr>
<td align="right">情報漏えいの可能性（機密性への影響, C）</td>
<td align="left">高（H）</td>
</tr>

<tr>
<td align="right">情報改ざんの可能性（完全性への影響, I）</td>
<td align="left">高（H）</td>
</tr>

<tr>
<td align="right">業務停止の可能性（可用性への影響, A）</td>
<td align="left">高（H）</td>
</tr>
</tbody>
</table>

<p>CVSS については<a href="http://text.baldanders.info/remark/2015/cvss-v3-metrics-in-jvn/" title="JVN が CVSSv3 による脆弱性評価を開始 — しっぽのさきっちょ | text.Baldanders.info">解説ページ</a>を参照のこと。</p>

</body>
</html>
```

まぁ，概ね良い感じではないだろうか。
ちなみに `-css` オプションで CSS ファイルへのパスを指定できるのでレイアウトについては多少調整できるだろう。

[russross/blackfriday] は v2 が最新版でかなり頻繁にメンテナンスされているようだが， [GitHub] リポジトリの master ブランチは v1 系なので注意が必要である。
そのうち弄って遊びたい。

## [GitHub] の Markdown API を利用する。

[GitHub] が Markdown 変換用の REST API を提供しているらしい。

- [Markdown | GitHub Developer Guide](https://developer.github.com/v3/markdown/)

更にこの API を利用するためのパッケージを Google が提供している。

- [google/go-github: Go library for accessing the GitHub API](https://github.com/google/go-github)

では，このパッケージを使った変換コードを書いてみよう。
こんな感じかな。

```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/go-github/github"
)

func main() {
	md, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	client := github.NewClient(nil)
	opt := &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}

	html, _, err := client.Markdown(string(md), opt)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	io.Copy(os.Stdout, strings.NewReader(html))
}
```

実際に動かしてみる。

```text
$ go run githubapi.go input.md > output2.html
```

出力はこんな感じになった。

```html
<h1>Go 1.8.7 および 1.9.4 がリリース</h1>
<p><a href="https://golang.org/" title="The Go Programming Language" rel="nofollow">Go 言語</a>コンパイラの 1.8.7 および 1.9.4 がリリースされている。<br>
このバージョンでは脆弱性 <a href="https://cve.mitre.org/cgi-bin/cvename.cgi?name=2018-6574" rel="nofollow">CVE-2018-6574</a> の修正が行われている。</p>
<h2>脆弱性の内容</h2>
<p><code>#cgo</code> ディレクティブを含む <a href="https://golang.org/" title="The Go Programming Language" rel="nofollow">Go 言語</a>コードをビルドする際に任意の処理を実行される可能性がある。</p>
<blockquote>
<p>When cgo is enabled, the build step during “go get” invokes the host C compiler, gcc or clang, adding compiler flags specified in the Go source files.<br>
Both gcc and clang support a plugin mechanism in which a shared-library plugin is loaded into the compiler, as directed by compiler flags.<br>
This means that a Go package repository can contain an attack.so file along with a Go source file that says (for example) <code>// #cgo CFLAGS: -fplugin=attack.so</code>, causing the attack plugin to be loaded into the host C compiler during the build.<br>
Gcc and clang plugins are completely unrestricted in their access to the host system.</p>
</blockquote>
<h2>影響度（CVSS）</h2>
<p>「<a href="https://access.redhat.com/security/cve/cve-2018-6574" rel="nofollow">CVE-2018-6574 - Red Hat Customer Portal</a>」より。</p>
<p><strong>CVSSv3 基本評価値 8.3 (<code>CVSS:3.0/AV:N/AC:H/PR:N/UI:R/S:C/C:H/I:H/A:H</code>)</strong></p>
<table>
<thead>
<tr>
<th align="right">基本評価基準</th>
<th align="left">評価値</th>
</tr>
</thead>
<tbody>
<tr>
<td align="right">攻撃元区分（AV）</td>
<td align="left">ネットワーク（N）</td>
</tr>
<tr>
<td align="right">攻撃条件の複雑さ（AC）</td>
<td align="left">高（H）</td>
</tr>
<tr>
<td align="right">必要な特権レベル（PR）</td>
<td align="left">不要（N）</td>
</tr>
<tr>
<td align="right">ユーザ関与レベル（UI）</td>
<td align="left">要（R）</td>
</tr>
<tr>
<td align="right">スコープ（S）</td>
<td align="left">変更あり（C）</td>
</tr>
<tr>
<td align="right">情報漏えいの可能性（機密性への影響, C）</td>
<td align="left">高（H）</td>
</tr>
<tr>
<td align="right">情報改ざんの可能性（完全性への影響, I）</td>
<td align="left">高（H）</td>
</tr>
<tr>
<td align="right">業務停止の可能性（可用性への影響, A）</td>
<td align="left">高（H）</td>
</tr></tbody></table>
<p>CVSS については<a href="http://text.baldanders.info/remark/2015/cvss-v3-metrics-in-jvn/" title="JVN が CVSSv3 による脆弱性評価を開始 — しっぽのさきっちょ | text.Baldanders.info" rel="nofollow">解説ページ</a>を参照のこと。</p>
```

出力されるのは `<body>` 要素内の部分のみなので，ページとして体裁を整えるには別途 `<head>` 要素などを補う必要がある。

それにしても `<th>` 要素や `<td>` 要素内の `align` 属性をいい加減なんとかしてくれないものだろうか。
これ [Validator](https://validator.w3.org/ "The W3C Markup Validation Service") にものごっつ怒られるんだよね。

## プレビュー用簡易 Web サーバ・ツール

最後は，変換ツールとはちょっと違うけど，プレビュー用簡易 Web サーバ・ツールを紹介する。

- [mattn/mkup: Portable Markdown Previewer](https://github.com/mattn/mkup)

[リリースページ](https://github.com/mattn/mkup/releases/latest)にバイナリが置いてあるので，ありがたく使わせていただく。
先程の `input.md` が置いてあるフォルダで

```text
$ mkup
Lisning at :8000
```

と起動する。
この状態で `http://localhost:8000/input.md` にアクセスすれば HTML 変換された内容が表示される。
パーサは [russross/blackfriday] を使っているようだ。
ファイルの内容を変更すれば自動的にリロードしてくれる。

例えば，このブログの記事は markdown で入力してるけど， [Hugo] のサーバ・モードでリアルタイムに確認しながら作業できるのがとても便利である。
そういう意味で簡易サーバ型のプレビュー・ツールは使い勝手がいいと思う。

## ブックマーク

- [ひろぽっぽれす - いくらが大好きです。Webの開発やってます - GolangでBlackfridayを使ってMarkdownをHTMLに変換する](http://hiropo.co.uk/archives/985.html)
- [Markdown記法で書かれたものをHTMLに変換するGo言語コード](https://ja.louisaandlily.com/markdown-to-html-with-golang)
- [Golang で GitHub の Markdown API をたたく · m0t0k1ch1st0ry](https://m0t0k1ch1st0ry.com/blog/2014/06/03/github-api/)
- [Big Sky :: markdown を何時でもチャラーっとプレビューしたい。](https://mattn.kaoriya.net/software/lang/go/20150703125421.htm)


[Go 言語]: https://golang.org/ "The Go Programming Language"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[russross/blackfriday]: https://github.com/russross/blackfriday "russross/blackfriday: Blackfriday: a markdown processor for Go"
[russross/blackfriday-tool]: https://github.com/russross/blackfriday-tool "russross/blackfriday-tool: Blackfriday command-line tool"
[GitHub]: https://github.com/

<!-- eof -->
