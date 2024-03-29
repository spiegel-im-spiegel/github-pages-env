+++
title = "Go 言語でデータ圧縮と解凍"
date =  "2017-12-03T20:07:25+09:00"
description = "最近データ圧縮と解凍でちょっと悩んだので覚え書きとしてまとめておく。"
tags        = [ "golang", "compress", "zlib", "gzip", "bzip2" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

最近データ圧縮と解凍で[ちょっと悩んだ]({{< ref "/remark/2017/11/gpgpdump-0_3_0-released.md" >}} "gpgpdump 0.3.0 をリリースした")ので覚え書きとしてまとめておく。

## ZLIB 圧縮と解凍（[RFC 1950]）

[RFC 1950] で定義される圧縮方式。

[Go 言語]では標準パッケージの [`compress/zlib`] で提供される。
コードで書くとこんな感じ。

```go
package main

import (
    "bytes"
    "compress/zlib"
    "fmt"
    "io"
    "os"
)

func compress(r io.Reader) (*bytes.Buffer, error) {
    buf := new(bytes.Buffer)
    zw := zlib.NewWriter(buf)
    defer zw.Close()

    if _, err := io.Copy(zw, r); err != nil {
        return buf, err
    }
    return buf, nil
}

func extract(zr io.Reader) (io.Reader, error) {
    return zlib.NewReader(zr)
}

func main() {
    content := "Hello world\n" //raw data
    //compress raw data
    zr, err := compress(bytes.NewBufferString(content))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    //display compressed data
    b := zr.Bytes()
    fmt.Printf("%d bytes: %v\n", len(b), b)
    //extract from compressed data
    r, err := extract(zr)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    //display extracted data
    io.Copy(os.Stdout, r)
}
```

実行結果はこんな感じになる。

```text
$ go run zlib.go
24 bytes: [120 156 242 72 205 201 201 87 40 207 47 202 73 225 2 4 0 0 255 255 28 242 4 71]
Hello world
```

## DEFLATE 圧縮と解凍（[RFC 1951]）

[RFC 1951] で定義される圧縮方式。

[Go 言語]では標準パッケージの [`compress/flate`] で提供される。
コードで書くとこんな感じ。

```go
package main

import (
    "bytes"
    "compress/flate"
    "fmt"
    "io"
    "os"
)

func compress(r io.Reader) (*bytes.Buffer, error) {
    buf := new(bytes.Buffer)
    zw, err := flate.NewWriter(buf, flate.BestCompression)
    if err != nil {
        return buf, err
    }
    defer zw.Close()

    if _, err := io.Copy(zw, r); err != nil {
        return buf, err
    }
    return buf, nil
}

func extract(zr io.Reader) (io.Reader, error) {
    return flate.NewReader(zr), nil
}

func main() {
    content := "Hello world\n" //raw data
    //compress raw data
    zr, err := compress(bytes.NewBufferString(content))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    //display compressed data
    b := zr.Bytes()
    fmt.Printf("%d bytes: %v\n", len(b), b)
    //extract from compressed data
    r, err := extract(zr)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    //display extracted data
    io.Copy(os.Stdout, r)
}
```

分かりやすいように関数のインタフェースを合わせてある。

実行結果はこんな感じになる。

```text
$ go run zip.go
18 bytes: [242 72 205 201 201 87 40 207 47 202 73 225 2 4 0 0 255 255]
Hello world
```

実は C 言語などでは， [RFC 1951] は [RFC 1950] と同じ [zlib] で提供される。
だから私はてっきり [RFC 1951] も [`compress/zlib`] で操作できると思いこんでいて，途方に暮れていたのだ。
うん。
次からは間違えないぞ！

ついでに他の方式も紹介しておこう。

## Gzip 圧縮と解凍（[RFC 1952]）

[RFC 1952] で定義される圧縮方式。
これも C 言語などでは [zlib] で提供されている。
データというよりファイル圧縮に使うものだが，データ圧縮にも使えないことはない。

[Go 言語]では標準パッケージの [`compress/gzip`] で提供されるのだが，実は以前に紹介している。

- [Gzip 操作について覚え書き]({{< relref "gzip-operation.md" >}})

今回は上述のコードに合わせて書き直してみる。

```go
package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
    "os"
)

func compress(r io.Reader) (*bytes.Buffer, error) {
    buf := new(bytes.Buffer)
    zw, err := gzip.NewWriterLevel(buf, gzip.BestCompression)
    if err != nil {
        return buf, err
    }
    defer zw.Close()

    if _, err := io.Copy(zw, r); err != nil {
        return buf, err
    }
    return buf, nil
}

func extract(zr io.Reader) (io.Reader, error) {
    return gzip.NewReader(zr)
}

func main() {
    content := "Hello world\n" //raw data
    //compress raw data
    zr, err := compress(bytes.NewBufferString(content))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    //display compressed data
    b := zr.Bytes()
    fmt.Printf("%d bytes: %v\n", len(b), b)
    //extract from compressed data
    r, err := extract(zr)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    //display extracted data
    io.Copy(os.Stdout, r)
}
```

実行結果はこんな感じになる。

```text
$ go run gzip.go
36 bytes: [31 139 8 0 0 0 0 0 2 255 242 72 205 201 201 87 40 207 47 202 73 225 2 4 0 0 255 255 213 224 57 183 12 0 0 0]
Hello world
```

[RFC 1950], [RFC 1951], [RFC 1952] はデータのフォーマットが違うだけで圧縮アルゴリズムは同じ。
なので，単純にデータを圧縮するだけなら [RFC 1951] でいいという話になる。

## [Bzip2] 圧縮と解凍

これは [zlib] 系の圧縮アルゴリズムとは異なる系統。
[zlib] よりも圧縮率がいいと言われている。
これもデータというよりファイル圧縮に使うものだが，データ圧縮にも使えないことはない。

[Go 言語]では標準パッケージの [`compress/bzip2`] で提供されるのだが，解凍しかサポートされてない。

しょうがないので解凍のみ書いてみる。

```go
package main

import (
    "bytes"
    "compress/bzip2"
    "fmt"
    "io"
    "os"
)

func extract(zr io.Reader) (io.Reader, error) {
    return bzip2.NewReader(zr), nil
}

func main() {
    cdata := []byte{66, 90, 104, 57, 49, 65, 89, 38, 83, 89, 176, 48, 136, 246, 0, 0, 1, 85, 128, 0, 16, 64, 0, 0, 64, 6, 4, 144, 128, 32, 0, 34, 40, 246, 166, 244, 8, 6, 4, 105, 205, 172, 132, 162, 238, 72, 167, 10, 18, 22, 6, 17, 30, 192}
    //display compressed data
    fmt.Printf("%d bytes: %v\n", len(cdata), cdata)
    //extract from compressed data
    r, err := extract(bytes.NewReader(cdata))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    //display extracted data
    io.Copy(os.Stdout, r)
}
```

実行結果はこんな感じになる。

```text
$ go run bzip2.go
54 bytes: [66 90 104 57 49 65 89 38 83 89 176 48 136 246 0 0 1 85 128 0 16 64 0 0 64 6 4 144 128 32 0 34 40 246 166 244 8 6 4 105 205 172 132 162 238 72 167 10 18 22 6 17 30 192]
Hello world
```

とはいえ「圧縮機能がないのは困るなぁ」と思って探したら，ありました。

- [dsnet/compress: Collection of compression related Go packages.](https://github.com/dsnet/compress)

これを使ってこれまでと同じように圧縮と解凍の処理を書き直してみる。

```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"

    "github.com/dsnet/compress/bzip2"
)

func compress(r io.Reader) (*bytes.Buffer, error) {
    buf := new(bytes.Buffer)
    zw, err := bzip2.NewWriter(buf, &bzip2.WriterConfig{Level: bzip2.BestCompression})
    if err != nil {
        return buf, err
    }
    defer zw.Close()

    if _, err := io.Copy(zw, r); err != nil {
        return buf, err
    }
    return buf, nil
}

func extract(zr io.Reader) (io.Reader, error) {
    return bzip2.NewReader(zr, new(bzip2.ReaderConfig))
}

func main() {
    content := "Hello world\n" //raw data
    //compress raw data
    zr, err := compress(bytes.NewBufferString(content))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    //display compressed data
    b := zr.Bytes()
    fmt.Printf("%d bytes: %v\n", len(b), b)
    //extract from compressed data
    r, err := extract(zr)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    //display extracted data
    io.Copy(os.Stdout, r)
}
```

実行結果はこんな感じになる。

```text
$ go run bzip2.go
54 bytes: [66 90 104 57 49 65 89 38 83 89 176 48 136 246 0 0 1 85 128 0 16 64 0 0 64 6 4 144 128 32 0 34 40 246 166 244 8 6 4 105 205 172 132 162 238 72 167 10 18 22 6 17 30 192]
Hello world
```

めでたし。

## ブックマーク

今回は言及しなかった `archive/*` パッケージ周りの記事を中心にブックマークしてみた。

- [Go言語でzipファイル解凍 - Qiita](https://qiita.com/hnaohiro/items/e6fd3154bd277bc6302e)
- [Big Sky :: バイナリ一つで zip, tar.gz, tar.bz2, tar.xz が開けるコマンド「archiver」(と go1.8 への対応方法)](https://mattn.kaoriya.net/software/lang/go/20161202095532.htm)
    - [Big Sky :: Golang の archive/zip でタイムゾーンの問題とファイル名の問題が解決した。](https://mattn.kaoriya.net/software/lang/go/20171111154359.htm)
- [golang標準ライブラリから学ぶzipファイルの構造](https://blog.freedom-man.com/zip-structure-golang/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`compress/flate`]: https://golang.org/pkg/compress/flate/ "flate - The Go Programming Language"
[`compress/zlib`]: https://golang.org/pkg/compress/zlib/ "zlib - The Go Programming Language"
[`compress/gzip`]: https://golang.org/pkg/compress/gzip/ "gzip - The Go Programming Language"
[`compress/bzip2`]: https://golang.org/pkg/compress/bzip2/ "bzip2 - The Go Programming Language"
[RFC 1950]: https://tools.ietf.org/html/rfc1950 "RFC 1950 - ZLIB Compressed Data Format Specification version 3.3"
[RFC 1951]: https://tools.ietf.org/html/rfc1951 "RFC 1951 - DEFLATE Compressed Data Format Specification version 1.3"
[RFC 1952]: https://tools.ietf.org/html/rfc1952 "RFC 1952 - GZIP file format specification version 4.3"
[zlib]: https://zlib.net/
[Bzip2]: http://www.bzip.org/

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
