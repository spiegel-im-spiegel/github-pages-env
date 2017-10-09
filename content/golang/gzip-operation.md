+++
title = "Gzip 操作について覚え書き"
date =  "2017-09-19T17:31:49+09:00"
update =  "2017-09-22T16:37:42+09:00"
description = "このようにインスタンスの生存期間を意識することで Go 言語の得意なパターンに嵌めることが容易になる。"
tags        = [ "golang", "gzip", "defer" ]

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
+++

- [gzip のやり方 - Qiita](http://qiita.com/ta_ta_ta_miya/items/3a1cba3a13418b732811)

この記事を見て「んん？」となったので，覚え書きとして [`gzip`] パッケージについて紹介する。

[リンク先の記事]で挙げられているコードは以下の通り。

```go
func makeGzip(body string) []byte {
    var b bytes.Buffer
    gw := gzip.NewWriter(&b)
    _, err := gw.Write([]byte(body)); if err != nil {
        ...
    }
    gw.Flush()
    gw.Close()
    return b.Bytes()
}
```

ここで `gw.Close()` 関数を [defer] 指定すると返ってくるバイト列が不完全なデータになってしまう，という話。
これは，[リンク先の記事]で指摘されている通り， [`gzip`].`Writer.Close()` 関数で gzip のフッタデータをフラッシュしているからである。

{{< fig-quote title="gzip - The Go Programming Language" link="https://golang.org/pkg/compress/gzip/" lang="en" >}}
<q>Close closes the Writer by flushing any unwritten data to the underlying io.Writer and writing the GZIP footer. It does not close the underlying io.Writer.</q>
{{< /fig-quote >}}

つまり [defer] で指定した関数は return 後に駆動するため `b.Bytes()` 関数を呼び出した時点ではまだ不完全なデータということになる[^b1]。

[^b1]: この挙動から分かるとおり， [`bytes`].`Buffer.Bytes()` 関数は，バッファの内容をそのまま返しているのではなく，内容のコピーを返している。

ここでちょっと考える。

関数の再利用性を考えるのなら，関数内でバッファを生成してバッファ処理の結果を返すのはあまり筋がよろしくない。
また圧縮データを書き込む先はメモリバッファじゃなくてファイルかもしれない。

ゆえに関数をこう書き換える。

```go
func makeGzip(dst io.Writer, content []byte) error {
    zw, err := gzip.NewWriterLevel(dst, gzip.BestCompression)
    if err != nil {
        return err
    }
    defer zw.Close()

    if _, err := zw.Write(content); err != nil {
        return err
    }
    return nil
}
```

つまり圧縮データの書き込み先である `Writer` を引数で指定するのである。
これなら生成した [`gzip`].`Writer.Close()` 関数を問題なく [defer] で指定できる。

これを踏まえて完全なコードは以下のようになる。

```go
package main

import (
    "compress/gzip"
    "fmt"
    "io"
    "os"
)

func makeGzip(dst io.Writer, content []byte) error {
    zw, err := gzip.NewWriterLevel(dst, gzip.BestCompression)
    if err != nil {
        return err
    }
    defer zw.Close()

    if _, err := zw.Write(content); err != nil {
        return err
    }
    return nil
}

func main() {
    content := []byte("Hello world\n")

    file, err := os.Create("test.txt.gz")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer file.Close()

    if err := makeGzip(file, content); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
```

このコードでは圧縮データの書き込む先をファイルにしている。
もちろん書き込み先を [`bytes`].`Buffer` に置き換えることもできる。
このようにインスタンスの生存期間を意識することで [Go 言語]の得意なパターンに嵌めることが容易になる。

ついでに対となる読み込み処理のコードも示しておこう。
ここでは復元したデータを標準出力に直接出力している。

```go
package main

import (
    "compress/gzip"
    "fmt"
    "io"
    "os"
)

func readGzip(dst io.Writer, src io.Reader) error {
    zr, err := gzip.NewReader(src)
    if err != nil {
        return err
    }
    defer zr.Close()

    io.Copy(dst, zr)

    return nil
}

func main() {
    file, err := os.Open("test.txt.gz")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer file.Close()

    if err := readGzip(os.Stdout, file); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
```

ところで，ファイル操作では生のデータを直接 gzip 圧縮するシチュエーションは少なく，大抵は tar と組み合わせることになる。
そこで tar と組み合わせ，指定フォルダ直下の複数ファイルを gzip 圧縮するコードも以下に示しておく。

```go
package main

import (
    "archive/tar"
    "compress/gzip"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
)

func makeTarGzip(dst io.Writer, rt string) error {
    zw, err := gzip.NewWriterLevel(dst, gzip.BestCompression)
    if err != nil {
        return err
    }
    defer zw.Close()

    tw := tar.NewWriter(zw)
    defer tw.Close()

    filepath.Walk(rt, func(path string, info os.FileInfo, err error) error {
        if info.IsDir() {
            return nil
        }
        fmt.Println(path)

        hd, e := tar.FileInfoHeader(info, "")
        if e != nil {
            return e
        }
        content, e := ioutil.ReadFile(path)
        if e != nil {
            return e
        }

        if e := tw.WriteHeader(hd); e != nil {
            return e
        }
        if _, e := tw.Write(content); e != nil {
            return e
        }
        return nil
    })
    if err != nil {
        return err
    }

    return nil
}

func main() {
    file, err := os.Create("test.tar.gz")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer file.Close()

    if err := makeTarGzip(file, "./"); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
```

## 【追記】別解あります

実は最初の `makeGzip()` には別解がある。
要するに [`gzip`] 処理部分を関数スコープで囲ってしまえばいいのだ。
実際にはこんな感じ。

```go
func makeGzip(body []byte) ([]byte, error) {
    var b bytes.Buffer
    err := func() error {
        gw := gzip.NewWriter(&b)
        defer gw.Close()

        if _, err := gw.Write(body); err != nil {
            return err
        }
        return nil
    }()
    return b.Bytes(), err
}
```

完全なコードはこんな感じになる。

```go
package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "os"
)

func makeGzip(body []byte) ([]byte, error) {
    var b bytes.Buffer
    err := func() error {
        gw := gzip.NewWriter(&b)
        defer gw.Close()

        if _, err := gw.Write(body); err != nil {
            return err
        }
        return nil
    }()
    return b.Bytes(), err
}

func main() {
    content := []byte("Hello world\n")

    file, err := os.Create("test.txt.gz")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer file.Close()

    z, err := makeGzip(content)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    if _, err := file.Write(z); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
```

このように（[defer] を含む）一連の処理を関数スコープで囲うやり方は，条件分岐や繰り返し処理の中で役に立つこともあるだろう。

## ブックマーク

- [「連結されたgzipを1行ずつ見る」をgolangでやったらハマった - Qiita](http://qiita.com/kroton/items/431e6dad9e5e4dbc44cf)
- [Big Sky :: golang では for ループの中で defer してはいけない。](https://mattn.kaoriya.net/software/lang/go/20151212021608.htm)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[defer]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[`gzip`]: https://golang.org/pkg/compress/gzip/ "gzip - The Go Programming Language"
[`bytes`]: https://golang.org/pkg/bytes/ "bytes - The Go Programming Language"

[リンク先の記事]: http://qiita.com/ta_ta_ta_miya/items/3a1cba3a13418b732811 "gzip のやり方 - Qiita"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
