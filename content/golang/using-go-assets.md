+++
title = "go-assets でシングルバイナリにまとめる"
date = "2018-02-13T23:31:39+09:00"
update =  "2018-02-15T07:19:42+09:00"
description = "実行モジュール以外の外部ファイルをソースコードに取り込んでマージすることにより，全部ひっくるめてシングルバイナリにする方法を考える。"
image = "/images/attention/go-code2.png"
tags = [ "golang", "tools", "engineering" ]

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
  mathjax = true
  mermaidjs = false
+++

ここのところ小ネタばかりでゴメンペコン。

[Go 言語]の利点のひとつは単一の実行モジュール（バイナリ・ファイル）でアプリケーションを deploy できる点にある。
しかし，実行モジュール以外のファイル（たとえば固定の辞書ファイルやテンプレート・ファイルなど）がセットになっている場合はそうもいかなくなる。

そこで，実行モジュール以外の外部ファイルをソースコードに取り込んでマージすることにより，全部ひっくるめてシングルバイナリにする方法が考えられた。

ファイル内容を取り込めばその分バイナリのサイズが大きくなってしまうが，元々 [Go 言語]の実行モジュールはモノリシックな構造でファイルサイズが大きいので，[組み込み]({{< relref "embedded-engineering-with-golang.md" >}} "組込みで Go")など計算リソースに制限がある場合を除けば，大した問題にはならないと思われる。

外部ファイルをマージする方法として，以下のパッケージが有名なようだ。

- [jessevdk/go-assets]
- [jteeuwen/go-bindata]

ただし [jteeuwen/go-bindata] のほうは [awesome-go から削除される](https://pinzolo.github.io/2017/11/13/go-bindata-was-remove-from-awesome-go.html "go-bindata が awesome-go から削除された | tail -f pinzo.log")などちょっとアレな感じになってるみたいで，今後使うなら [jessevdk/go-assets] のほうがいいかもしれない。

というわけで，この記事では [jessevdk/go-assets] の使い方を簡単に紹介する。

## まずはファイルを用意

今回のフォルダ・ファイル構成はこんな感じ。

```text
hello/
│  hello.go
│
└─html/
        index.html
```

このうち `html/index.html` をマージしたい。
中身はこんな感じ。

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Hello World!</title>
</head>
<body>
    <p>Hello World!</p>
</body>
</html>
```

## go-assets-builder のインストール

 `index.html` ファイルを [Go 言語]コードに変換するために [jessevdk/go-assets-builder] をインストールする。
 バイナリは提供されていないので，ここは素直に `go get` コマンドを使う。

```text
$ go get -v github.com/jessevdk/go-assets-builder
```

`-h` オプションで見るとこんな感じ。

```
$ go-assets-builder -h
Usage:
  go-assets-builder.exe [OPTIONS] FILES...

Application Options:
  -p, -package:       The package name to generate the assets for (default: main)
  -v, -variable:      The name of the generated asset tree (default: Assets)
  -s, -strip-prefix:  Strip the specified prefix from all paths
  -o, -output:        File to write output to, or - to write to stdout (default: -)

Help Options:
  -?                  Show this help message
  -h, -help           Show this help message
```

## ソース・コードの生成

早速 `hello/` フォルダ直下で `go-assets-builder` を動かしてみる。

```text
$  go-assets-builder html/ > assets.go
```

これで `assets.go` ファイルが生成された。
中身はこんな感じ。

```go
package main

import (
    "time"

    "github.com/jessevdk/go-assets"
)

var _Assets950a25363eee220f7d8ce234bcc0b349e4ea9072 = "<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset=\"utf-8\">\n\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n\t<title>Hello World!</title>\n</head>\n<body>\n\t<p>Hello World!</p>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/html": []string{"index.html"}, "/": []string{"html"}}, map[string]*assets.File{
    "/": &assets.File{
        Path:     "/",
        FileMode: 0x800001ff,
        Mtime:    time.Unix(1518523695, 1518523695933225200),
        Data:     nil,
    }, "/html": &assets.File{
        Path:     "/html",
        FileMode: 0x800001ff,
        Mtime:    time.Unix(1518522609, 1518522609437368800),
        Data:     nil,
    }, "/html/index.html": &assets.File{
        Path:     "/html/index.html",
        FileMode: 0x1b6,
        Mtime:    time.Unix(1518523737, 1518523737979915400),
        Data:     []byte(_Assets950a25363eee220f7d8ce234bcc0b349e4ea9072),
    }}, "")
```

パッケージ名（既定は `main`）は `-p` オプションで，変数名（既定は `Assets`）は `-v` オプションで変更できる。

上の例では `index.html` ファイルのパスは実際のフォルダ構成のまま `/html/index.html` となっているが， `html/` フォルダ直下をドキュメント・ルートにするのであれば `-s="/html"` とする。

```
$ go-assets-builder -p data -v Docs -s="/html" html
package data

import (
    "time"

    "github.com/jessevdk/go-assets"
)

var _Docs950a25363eee220f7d8ce234bcc0b349e4ea9072 = "<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset=\"utf-8\">\n\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n\t<title>Hello World!</title>\n</head>\n<body>\n\t<p>Hello World!</p>\n</body>\n</html>"

// Docs returns go-assets FileSystem
var Docs = assets.NewFileSystem(map[string][]string{"/": []string{"index.html"}}, map[string]*assets.File{
    "/": &assets.File{
        Path:     "/",
        FileMode: 0x800001ff,
        Mtime:    time.Unix(1518522609, 1518522609437368800),
        Data:     nil,
    }, "/index.html": &assets.File{
        Path:     "/index.html",
        FileMode: 0x1b6,
        Mtime:    time.Unix(1518523737, 1518523737979915400),
        Data:     []byte(_Docs950a25363eee220f7d8ce234bcc0b349e4ea9072),
    }}, "")
```

## 組み込んだファイルを読む

生成したソース・コードで定義した `Assets` 変数を使って `hello.go` ファイルの `main()` 関数は以下のように書ける。

{{< highlight go "hl_lines=10" >}}
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    f, err := Assets.Open("/html/index.html")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    io.Copy(os.Stdout, f)
}
{{< /highlight >}}

では，実行してみよう。
こんな感じになる。

```text
$ go run hello.go assets.go
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Hello World!</title>
</head>
<body>
    <p>Hello World!</p>
</body>
</html>
```

まぁ，テキストを表示してるだけだけど。

ちなみに [`assets`]`.FileSystem` は [`http`]`.FileSystem` インタフェースと互換がある。
以下は [`http`]`.FileSystem` インタフェースの定義。

```go
type FileSystem interface {
    Open(name string) (File, error)
}
```

従って [`assets`]`.FileSystem` を使って以下のような簡易 Web サーバも組める。

{{< highlight go "hl_lines=13" >}}
package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    fmt.Println("Open http://localhost:3000/html/")
    fmt.Println("Press ctrl+c to stop")

    http.Handle("/", http.FileServer(Assets))
    if err := http.ListenAndServe(":3000", nil); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
{{< /highlight >}}

## go generate コマンドによるソース・コードの生成

 `//` コメントに `go:generate` マーカを使うことにより `go generate` コマンドで `go-assets-builder` コマンドを呼び出せる。

 {{< highlight go "hl_lines=9" >}}
 package main

 import (
     "fmt"
     "net/http"
     "os"
 )

 //go:generate go-assets-builder -s="/html" -o assets.go html/

 func main() {
     fmt.Println("Open http://localhost:3000/")
     fmt.Println("Press ctrl+c to stop")

     http.Handle("/", http.FileServer(Assets))
     if err := http.ListenAndServe(":3000", nil); err != nil {
         fmt.Fprintln(os.Stderr, err)
     }
 }
 {{< /highlight >}}

```text
$ go generate

$ go run hello.go assets.go
Open http://localhost:3000/
Press ctrl+c to stop
```

`go generate` コマンドは明示的に行う必要があるので注意。
ビルドでソース・コード生成を含めて自動化したいなら make とか導入する必要があるかも。

## ブックマーク

- [Generating code - The Go Blog](https://blog.golang.org/generate)

- [go-assets 使い方 - トミールの技術系日記](http://tomi-ru.hatenablog.com/entry/2016/09/22/go-assets_%E4%BD%BF%E3%81%84%E6%96%B9)
- [Big Sky :: Re: Go でシングルバイナリな Web アプリを開発しているときに webpack --watch をうまいところやる](https://mattn.kaoriya.net/software/lang/go/20170119180147.htm)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"
[jessevdk/go-assets]: https://github.com/jessevdk/go-assets "jessevdk/go-assets: Simple embedding of assets in go"
[`assets`]: https://github.com/jessevdk/go-assets "jessevdk/go-assets: Simple embedding of assets in go"
[jessevdk/go-assets-builder]: https://github.com/jessevdk/go-assets-builder "jessevdk/go-assets-builder: Simple assets builder program for go-assets"
[jteeuwen/go-bindata]: https://github.com/jteeuwen/go-bindata "jteeuwen/go-bindata: Hard fork of jteeuwen/go-bindata because it disappeared, Please refer to issues#5 for details."

## 参考図書

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/477418392X/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/61EL3Dc95dL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/477418392X/baldandersinf-22/">みんなのGo言語【現場で使える実践テクニック】</a></dt><dd>松木雅幸 mattn 藤原俊一郎 中島大一 牧 大輔 鈴木健太 稲葉貴洋 </dd><dd>技術評論社 2016-09-09</dd><dd>評価<abbr class="rating" title="4"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4774184322/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4774184322.09._SCTHUMBZZZ_.jpg"  alt="WEB+DB PRESS Vol.95"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4621300253.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4274219151/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4274219151.09._SCTHUMBZZZ_.jpg"  alt="プログラミングElixir"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4798147400/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4798147400.09._SCTHUMBZZZ_.jpg"  alt="詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド (NEXT ONE)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4774182869/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4774182869.09._SCTHUMBZZZ_.jpg"  alt="WEB+DB PRESS Vol.94"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873117763/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4873117763.09._SCTHUMBZZZ_.jpg"  alt="Docker"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/477418361X/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/477418361X.09._SCTHUMBZZZ_.jpg"  alt="オブジェクト指向設計実践ガイド ~Rubyでわかる 進化しつづける柔軟なアプリケーションの育て方"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4295000337/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4295000337.09._SCTHUMBZZZ_.jpg"  alt="WebデベロッパーのためのReact開発入門 JavaScript UIライブラリの基本と活用"  /></a> </p>
<p class="description">リファレンス本なのに索引が貧弱。これなら Kindle で買ってもよかったか。 1.7 への言及があるのはよかった。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-11-17">2016-11-17</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
