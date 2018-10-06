+++
title = "Go 1.11 のリリースと「モジュール」機能の実験的サポート"
date = "2018-09-26T13:53:04+09:00"
update = "2018-10-06T13:40:55+09:00"
description = "「モジュール」とは，これまで vgo (Versioned Go) として開発が進められてきたものである。この機能について簡単に紹介する。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "golang", "engineering", "module", "vendoring", "versioning" ]

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

1ヶ月も前の話で申し訳ないが，先月末に [Go 言語]コンパイラ 1.11 がリリースされた。

- [Go 1.11 is released - The Go Blog](https://blog.golang.org/go1.11)
- [Go 1.11 Release Notes - The Go Programming Language](https://golang.org/doc/go1.11)

主な変更点は以下の2つ。

- WebAssembly へのコンパイルをサポート
- 「モジュール」機能の実験的サポート

このうち，今回は「モジュール」機能のサポートについて簡単に紹介する。

「モジュール」とは，これまで vgo (Versioned Go) として開発が進められてきたもので

- [vgo (Versioned Go) に関する覚え書き]({{< ref "/golang/go-and-versioning.md" >}})

この成果が正式に組み込まれる。
バージョン 1.11 で「モジュール」機能を有効にするには環境変数 `GO111MODULE` を `on` または `auto` にセットする。

ちなみに，環境変数 `GO111MODULE` を `auto` にセットした場合， GOPATH 下のコードについては以前と同じように動作するが，それ以外の場所では「モジュール」機能が有効になる。
1.11 では環境変数 `GO111MODULE` が既定で auto になっている。

試しに以下のソースファイル `hello.go` を適当なフォルダに作成する。

```go
package main

import (
    "fmt"

    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
```

これを実行しても `rsc.io/quote` パッケージがないと怒られてコンパイルエラーになる。
ここまでは従来どおりの動作。

次に以下のコマンドを実行する。

```text
$ go mod init hello
go: creating new go.mod: module hello
```

このコマンドで `go.mod` ファイルが生成される。
中身は以下の通り。

```text
module hello
```

これで `hello.go` ファイルはモジュール `hello` のコードとして定義された。

モジュール名はパッケージのインポート・パスと同じく[リポジトリのパスと連動](https://golang.org/cmd/go/#hdr-Remote_import_paths)している。
たとえば `hello.go` ファイルをリポジトリ `https://github.com/spiegel-im-spiegel/hello` で公開・管理するなら，モジュール名も `github.com/spiegel-im-spiegel/hello` となる。

では，この状態でコードを実行してみよう。

```text
$ go run hello.go
go: finding rsc.io/quote v1.5.2
go: downloading rsc.io/quote v1.5.2
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
Hello, world.
```

自動的に `rsc.io/quote` およびその依存パッケージがモジュールとして読み込まれていることが分かると思う。
このとき `go.mod` ファイルを見ると

```text
module hello

require rsc.io/quote v1.5.2
```

依存モジュールおよびそのバージョンの記述が追加されていることが分かる。
モジュール間の依存関係を見るには以下のコマンドを実行する[^dot1]。

[^dot1]: 是非とも dot 言語で出力してほしい。

```text
$ go mod graph
hello rsc.io/quote@v1.5.2
rsc.io/quote@v1.5.2 rsc.io/sampler@v1.3.0
rsc.io/sampler@v1.3.0 golang.org/x/text@v0.0.0-20170915032832-14c0d48ead0c
```

ちなみに，読み込まれたモジュールの実体は `$GOPATH/mod` フォルダ以下に格納されている。

更に `go.mod` ファイルを以下のように書き換えてみよう。

```text
module hello

require (
	golang.org/x/text v0.3.0
	rsc.io/quote v1.5.2
)
```

この状態で，再び実行してみる。

```text
$ go run hello.go
go: finding golang.org/x/text v0.3.0
go: downloading golang.org/x/text v0.3.0
Hello, world.
```

`golang.org/x/text` モジュールの v0.3.0 が読み込まれていることが分かる。

既存のパッケージをモジュールとして定義する際， [dep] で管理していると移行がスムーズになる。
自作の [gpgpdump] で試してみよう。

まずはリポジトリの内容を適当なフォルダに展開し `go.mod` ファイルを作成する。

```text
$ go mod init github.com/spiegel-im-spiegel/gpgpdump
go: creating new go.mod: module github.com/spiegel-im-spiegel/gpgpdump
go: copying requirements from Gopkg.lock
```

依存情報を [dep] の `Gopkg.lock` ファイルから取得しているのが分かるだろう。
生成された `go.mod` ファイルの内容は以下の通り。

```text
module github.com/spiegel-im-spiegel/gpgpdump

require (
	github.com/BurntSushi/toml v0.3.0
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/pkg/errors v0.8.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.2
	github.com/spiegel-im-spiegel/gocli v0.8.0
	golang.org/x/crypto v0.0.0-20180816225734-aabede6cba87
)
```

新しいバージョンがリリースされていないか調べる場合は以下のコマンドを実行する。

```text
$ go list -m -u all
go: finding github.com/pkg/errors v0.8.0
go: finding github.com/spf13/pflag v1.0.2
go: finding github.com/spf13/cobra v0.0.3
go: finding github.com/BurntSushi/toml v0.3.0
go: finding github.com/spiegel-im-spiegel/gocli v0.8.0
go: finding github.com/inconshreveable/mousetrap v1.0.0
go: finding golang.org/x/crypto v0.0.0-20180816225734-aabede6cba87
go: finding golang.org/x/crypto latest
go: finding github.com/BurntSushi/toml v0.3.1
github.com/spiegel-im-spiegel/gpgpdump
github.com/BurntSushi/toml v0.3.0 [v0.3.1]
github.com/inconshreveable/mousetrap v1.0.0
github.com/pkg/errors v0.8.0
github.com/spf13/cobra v0.0.3
github.com/spf13/pflag v1.0.2
github.com/spiegel-im-spiegel/gocli v0.8.0
golang.org/x/crypto v0.0.0-20180816225734-aabede6cba87 [v0.0.0-20180910181607-0e37d006457b]
```

この情報を元に `go.mod` ファイルを書き換えてみる。

{{< highlight text "hl_lines=4 10" >}}
module github.com/spiegel-im-spiegel/gpgpdump

require (
	github.com/BurntSushi/toml v0.3
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/pkg/errors v0.8.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.2
	github.com/spiegel-im-spiegel/gocli v0.8.0
	golang.org/x/crypto latest
)
{{< /highlight >}}

これで最新モジュールを取得できる。

```text
$ go mod download
go: finding golang.org/x/crypto latest
```

取得後の `go.mod` ファイルは以下のように書き換えられる。

{{< highlight text "hl_lines=4 10" >}}
module github.com/spiegel-im-spiegel/gpgpdump

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/pkg/errors v0.8.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.2
	github.com/spiegel-im-spiegel/gocli v0.8.0
	golang.org/x/crypto v0.0.0-20180910181607-0e37d006457b
)
{{< /highlight >}}


んー。
これで [dep] からの置き換えができそうかな。
ようやく「[GOPATH 汚染問題]({{< ref "/golang/gopath-pollution.md" >}})」が解消されるかねぇ。

## 【2018-10-04 追記】 Go 1.11.1 がリリース

{{% fig-quote title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.11.minor" %}}
“go1.11.1 (released 2018/10/01) includes fixes to the compiler, documentation, go command, runtime, and the `crypto/x509`, `encoding/json`, `go/types`, `net`, `net/http`, and `reflect` packages. See the [Go 1.11.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.11.1) on our issue tracker for details.”
{{% /fig-quote %}}

## ブックマーク

- [Go 1.11 リリースノート（和訳） - Qiita](https://qiita.com/pokeh/items/c6511ca15c9a33b48fcc)
- [Using Go modules with Travis CI | Dave Cheney](https://dave.cheney.net/2018/07/16/using-go-modules-with-travis-ci)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[dep]: https://golang.github.io/dep/ "dep · Dependency management for Go"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
