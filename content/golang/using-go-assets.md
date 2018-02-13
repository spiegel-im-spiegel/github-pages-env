+++
title = "go-assets でシングルバイナリにまとめる"
date =  "2018-02-13T16:49:36+09:00"
description = "description"
image = "/images/attention/go-code2.png"
tags        = [ "golang" ]
draft = true

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

[Go 言語]の利点のひとつは単一の実行モジュール（バイナリ・ファイル）でアプリケーションを提供できる点にある。
ただし，実行モジュール以外のファイルがセットになる場合はそうもいかなくなる。
たとえば固定の辞書ファイルやテンプレート・ファイルなどを含む場合だ。

そこで，これらの外部ファイルをソースコードに取り込んでマージすることで全部ひっくるめてシングルバイナリとして提供する方法が考えられた。
ファイルデータを取り込めばその分バイナリのサイズが大きくなってしまうが，元々 [Go 言語]の実行モジュールはモノリシックな構造でファイルサイズが大きいので，[組み込み]({{< relref "http://golang/embedded-engineering-with-golang.md" >}})など，計算リソースに制限が大きい場合を除けば，大きな問題にはならないと思われる。

外部ファイルをマージする方法として，以下のパッケージが有名なようだ。

- [jessevdk/go-assets]
- [jteeuwen/go-bindata]

[jteeuwen/go-bindata] のほうは [awesome-go から削除された](https://pinzolo.github.io/2017/11/13/go-bindata-was-remove-from-awesome-go.html)そうで，今後使うなら
[jessevdk/go-assets] のほうがいいかもしれない。

というわけで，この記事では [jessevdk/go-assets] の使い方を簡単に紹介する。

## まずはファイルを用意

今回のフォルダ・ファイル構成はこんな感じ。

```text
hello/
│  hello.go
│
└─data/
        hello.txt
```

このうち `hello.txt` をマージしたい。
中身はこんな感じ。

```text
Hello World!
```

## go-assets-builder のインストール

 `hello.txt` ファイルを [Go 言語]コードに変換するために [jessevdk/go-assets-builder] をインストールする。
 バイナリは提供されていないので，ここは素直に `go get` コマンドを使う。

```text
$ go get -v github.com/jessevdk/go-assets-builder
```

`-h` オプションで見るとこんな感じ。

``text
$ go-assets-builder -h
Usage:
  go-assets-builder.exe [OPTIONS] FILES...

Application Options:
  -p, -package:       The package name to generate the assets for (default: main)
  -v, -variable:      The name of the generated asset tree (default: Assets)
  -s, -strip-prefix:  Strip the specified prefix from all paths
  -o, -output:        File to write output to, or - to write to stdout
                      (default: -)

Help Options:
  -?                  Show this help message
  -h, -help           Show this help message
```

早速 `hello/` フォルダ直下で `go-assets-builder` を動かしてみる。

```text
$  go-assets-builder data/ > asset.go
```

これで `asset.go` ファイルが生成される。
中身はこんな感じ。

```go
package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets1d921dc85677928c7457b1a7687f82618b5cfd6b = "Hello World!\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"data"}, "/data": []string{"hello.txt"}}, map[string]*assets.File{
	"/data": &assets.File{
		Path:     "/data",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1518509705, 1518509705225360300),
		Data:     nil,
	}, "/data/hello.txt": &assets.File{
		Path:     "/data/hello.txt",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1518512089, 1518512089919917400),
		Data:     []byte(_Assets1d921dc85677928c7457b1a7687f82618b5cfd6b),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1518507214, 1518507214658243100),
		Data:     nil,
	}}, "")
```





[Go 言語]: https://golang.org/ "The Go Programming Language"
[jessevdk/go-assets]: https://github.com/jessevdk/go-assets "jessevdk/go-assets: Simple embedding of assets in go"
[jessevdk/go-assets-builder]: https://github.com/jessevdk/go-assets-builder "jessevdk/go-assets-builder: Simple assets builder program for go-assets"
[jteeuwen/go-bindata]: https://github.com/jteeuwen/go-bindata "jteeuwen/go-bindata: Hard fork of jteeuwen/go-bindata because it disappeared, Please refer to issues#5 for details."

<!-- eof -->
