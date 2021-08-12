+++
title = "紙芝居用の簡易サーバを書く【Go 1.16 版】"
date =  "2021-02-24T19:53:56+09:00"
description = "Go 1.16 で追加された embed および io/fs 標準パッケージを使う。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "golang", "programming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ずいぶん前に

- [紙芝居用の簡易 Web サーバを Go 言語で書く]({{< ref "/remark/2018/02/simple-web-server-with-golang.md" >}})

という記事を書いたが，これで作ったコードは任意のディレクトリをドキュメント・ルートにできるという結構ヤバい代物である。

ところで [Go 1.16 リリース]({{< ref "/release/2021/02/go-1_16-is-released.md" >}} "Go 1.16 がリリースされた")で真っ先に思ったのは「これで安全に紙芝居ができるぢゃん」だった。
[Go] 1.16 で追加された [`embed`] および [`io/fs`][`fs`] 標準パッケージを使えば，サードパーティの外部パッケージを使わずとも，簡単に「紙芝居」が作れる。

さっそく試してみよう。
こんな感じかな。

```go {hl_lines=[10,11,18]}
package main

import (
    "embed"
    "fmt"
    "net/http"
    "os"
)

//go:embed html
var assets embed.FS

func main() {
    addr := "localhost:3000"
    fmt.Printf("Open http://%s/\n", addr)
    fmt.Println("Press ctrl+c to stop")

    http.Handle("/", http.FileServer(http.FS(assets)))
    if err := http.ListenAndServe(addr, nil); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
```

ここで `//go:embed` ディレクティブがコンパイラに `html` の埋め込みを指示している部分である。
`html` は実在のディレクトリで，以下の内容の `index.html` ファイルを格納している[^embd1]。

[^embd1]: `//go:embed` ディレクティブはあくまでもファイルに対してディレクトリ構造ごと埋め込むものなので，空のディレクトリは埋め込めない。

```html
<!DOCTYPE html>
<html>
    <head>
        <title>Hello</title>
    </head>
    <body>
        <p>Hello World!</p>
    </body>
</html>
```

[`embed`]`.FS` 型のインスタンスはそのままでは [`http`]`.FileSystem` 型として使えないので [`http`]`.FS()` 関数で [`http`]`.FileSystem` 型に適合するよう変換している。

説明はこのくらいにして，とりあえず動かしてみよう。

```text
$ go run sample1.go 
Open http://localhost:3000/
Press ctrl+c to stop
```

これでブラウザで `http://localhost:3000/` にアクセスしてみると。

{{< fig-img src="./browser1.png" title="localhost:3000/" link="./browser1.png" >}}

ありゃりゃーん。
ちょっと切ない感じになってしまった。

`html` ディレクトリをドキュメント・ルートに出来ないかな，と思って色々眺めてたら [`fs`]`.Sub()` 関数が使えるっぽい。

というわけで，さきほどのコードを少し書き換える。

```go {hl_lines=[19,20]}
package main

import (
    "embed"
    "fmt"
    "io/fs"
    "net/http"
    "os"
)

//go:embed html
var assets embed.FS

func main() {
    addr := "localhost:3000"
    fmt.Printf("Open http://%s/\n", addr)
    fmt.Println("Press ctrl+c to stop")

    root, _ := fs.Sub(assets, "html")
    http.Handle("/", http.FileServer(http.FS(root)))
    if err := http.ListenAndServe(addr, nil); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
```

エラー処理をサボっているがご容赦。
では，このコードを起動してみる。

```text
$ go run sample2.go 
Open http://localhost:3000/
Press ctrl+c to stop
```

ブラウザでアクセスしてみると。

{{< fig-img src="./browser2.png" title="Hello" link="./browser2.png" >}}

よーし，うむうむ，よーし。
これで安全に紙芝居ができるな。
`go build` でシングル・バイナリにまとめれば持ち運びもOK（笑）

`//go:embed` ディレクティブはかなり便利に使えるようで

```go
package main

import (
    _ "embed"
    "fmt"
)

//go:embed html/index.html
var hello string

func main() {
    fmt.Println(hello)
}
```

てな感じに，ファイル内容を `[]byte` または `string` 型のデータに展開してくれる。

```text
$ go run sample3.go 
<!DOCTYPE html>
<html>
    <head>
        <title>Hello</title>
    </head>
    <body>
        <p>Hello World!</p>
    </body>
</html>
```

こりゃあ，ひょっとしてテストが捗るねぇ。
データの読み込みを `//go:embed` ディレクティブで簡略化できるもんね。

## ブックマーク

- [Big Sky :: Go に go:embed が入った。](https://mattn.kaoriya.net/software/lang/go/20201030092005.htm)
- [go:embed 詳解 - 使用編 -](https://zenn.dev/koya_iwamura/articles/53a4469271022e)

[Go]: https://golang.org/ "The Go Programming Language"
[`embed`]: https://golang.org/pkg/embed/ "embed - The Go Programming Language"
[`fs`]: https://golang.org/pkg/io/fs/ "fs - The Go Programming Language"
[`http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
