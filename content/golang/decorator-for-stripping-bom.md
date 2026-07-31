+++
title = "BOM を除去する io.ReadCloser Decorator を作ってみた"
date =  "2022-11-05T13:56:19+09:00"
description = "こうやって混沌としたコードを整理していくんですねぇ。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 Zenn で UTF-8 BOM (Byte Order Mark) を除去する Decorator を紹介したのだが

- [Decorator Pattern で BOM を除去する](https://zenn.dev/spiegel/articles/20221029-remove-bom)

最後の

```go
package main

import (
    "encoding/csv"
    "errors"
    "fmt"
    "io"
    "os"

    "github.com/spkg/bom"
)

func main() {
    file, err := os.Open("./sample3.csv")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer file.Close()

    r := csv.NewReader(bom.NewReader(file))
    for {
        row, err := r.Read()
        if err != nil {
            if !errors.Is(err, io.EOF) {
                fmt.Fprintln(os.Stderr, err)
                return
            }
            break
        }
        fmt.Println(row)
    }
}
```

は [`os`]`.File` 型， [`bom`]`.NewReader()` 関数が返す [`io`]`.Reader` 抽象型とその実体である [`bufio`]`.Reader`，そして [`csv`]`.Reader` 型の3つを意識する必要がある。
もう少し詳しく言うと `Close()` メソッドを持っているのはベースの [`os`]`.File` 型のみなのでファイルを閉じるためには [`os`]`.File` 型の変数を保持っておく必要があるのだ。
これは上手いやり方ではないなぁ，というのが記事を書いた後の感想だった。

そこで [`github.com/spkg/bom`][`bom`] パッケージを参考に，自前で [`github.com/goark/utf8bom`][`utf8bom`] パッケージを作ってみた。
このパッケージの `Reader` 型は

```go
type Reader struct {
    *bufio.Reader
    closer func() error
}
```

という構成になっていて，埋め込みの [`bufio`]`.Reader` フィールド以外に `closer` を持っている。
初期化時に

```go { hl_lines=["3-6"]}
func Strip(r io.Reader) *Reader {
    closer := func() error { return nil }
    if c, ok := r.(io.Closer); ok {
        closer = c.Close
    }
    br := &Reader{Reader: bufio.NewReader(r), closer: closer}
    b, err := br.Peek(3)
    if err != nil {
        return br
    }
    if bytes.Equal(b, []byte{0xef, 0xbb, 0xbf}) { // compare BOM
        _, _ = br.Discard(3)
    }
    return br
}
```

てな感じに `closer` フィールドにメソッド値[^mv1] をセットすることで [`utf8bom`]`.Reader.Close()` メソッド

[^mv1]: 「メソッド値」については拙文「[#golang メソッド式とメソッド値](https://zenn.dev/spiegel/articles/20201212-method-value-and-expression)」を参考にどうぞ。

```go
func (r *Reader) Close() error {
    return r.closer()
}
```

起動時にベースの `Close()` メソッドへ処理を委譲できるようにした。
こういうときに変数の生存期間とか考えなくていい [Go] は便利だよねぇ。

これを使って最初のサンプルコードを少し書き換えてみる。

```go { hl_lines=[10,"13-19",22]}
package main

import (
    "encoding/csv"
    "errors"
    "fmt"
    "io"
    "os"

    "github.com/goark/utf8bom"
)

func openCsvFile(path string) (io.ReadCloser, error) {
    file, err := os.Open("./sample3.csv")
    if err != nil {
        return nil, err
    }
    return utf8bom.Strip(file), nil
}

func main() {
    file, err := openCsvFile("./sample3.csv")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer file.Close()

    r := csv.NewReader(file)
    for {
        row, err := r.Read()
        if err != nil {
            if !errors.Is(err, io.EOF) {
                fmt.Fprintln(os.Stderr, err)
                return
            }
            break
        }
        fmt.Println(row)
    }
}
```

これで CSV ファイル・オープン時の仔細を掃き出して [`io`]`.ReadCloser` 型で扱えるようになった。
ここまで来れば [`csv`]`.NewReader()` 関数もどうにかして `openCsvFile()` 関数に掃き出したいよね。

ちうわけで，拙作の [`github.com/goark/csvdata`][`csvdata`] パッケージも導入する。
こんな感じ。

```go { hl_lines=[9,13,18,30,37]}
package main

import (
    "errors"
    "fmt"
    "io"
    "os"

    "github.com/goark/csvdata"
    "github.com/goark/utf8bom"
)

func openCsvFile(path string) (*csvdata.Rows, error) {
    file, err := os.Open("./sample3.csv")
    if err != nil {
        return nil, err
    }
    return csvdata.NewRows(csvdata.New(utf8bom.Strip(file)), true), nil
}

func main() {
    file, err := openCsvFile("./sample3.csv")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer file.Close()

    for {
        if err := file.Next(); err != nil {
            if !errors.Is(err, io.EOF) {
                fmt.Fprintln(os.Stderr, err)
                return
            }
            break
        }
        fmt.Println(file.Row())
    }
}
```

これでメイン側ではデータの行・列構造のみ注視すればよくなった[^gocsv]。

[^gocsv]: CSV データを扱う便利パッケージとしては [`github.com/gocarina/gocsv`](https://github.com/gocarina/gocsv "gocarina/gocsv: The GoCSV package aims to provide easy CSV serialization and deserialization to the golang programming language") が有名なのだが，あれって Unmarshal 時に CSV データ全体を構造体の配列にしてしまうのが気に食わないんだよなぁ。 CSV データが百万レコードあったら百万個の配列を作ってしまう。まぁ，巨大データを扱う前提ではないということなんだろうけど。

ちなみに [`github.com/goark/csvdata`][`csvdata`] パッケージは [Excel ファイルにも対応](https://zenn.dev/spiegel/articles/20211003-excel-as-a-csv "Excel も CSV みたいに扱いたい")していて，同じ [`csvdata`]`.Rows` 型に落とし込んで扱えるようになっている。
つまり上のサンプルコードの `openCsvFile()` 関数の中身をまるっと Excel 用に置き換えることができるのだ。

こうやって混沌としたコードを整理していくんですねぇ。

[Go]: https://go.dev/
[`bufio`]: https://pkg.go.dev/bufio "bufio package - bufio - Go Packages"
[`io`]: https://pkg.go.dev/io "io package - io - Go Packages"
[`os`]: https://pkg.go.dev/os "os package - os - Go Packages"
[`csv`]: https://pkg.go.dev/encoding/csv "csv package - encoding/csv - Go Packages"
[`bom`]: https://github.com/spkg/bom "spkg/bom: Strip UTF-8 byte order marks"
[`utf8bom`]: https://github.com/goark/utf8bom "goark/utf8bom: Strip leading UTF-8 BOM"
[`csvdata`]: https://github.com/goark/csvdata "goark/csvdata: Reading CSV Data"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
{{% review-paapi "B07FSBHS2V" %}} <!-- Clean Architecture -->
