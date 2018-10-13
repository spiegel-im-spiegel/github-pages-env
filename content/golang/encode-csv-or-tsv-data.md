+++
title = "CSV/TSV データの読み書き"
date = "2018-10-14T05:41:02+09:00"
description = "関数型言語に慣れている人から見ると Go 言語が標準で提供しているコンテナ操作のアレコレはまだるっこしい感じに見えると思う。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "map", "slice" ]

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

今回は小ネタ。
以下の記事を見て，思いつきで書いてみる。

- [tsvファイルの入出力を簡単にする技](https://qiita.com/kei0425/items/e095bc8435429a22a002)

「Python ならこんなに簡単に書けるのに [Go 言語]で書いたらワケワカメだよ」という，まぁよくある DIS り記事なのだが，関数型言語に慣れている人から見ると [Go 言語]が標準で提供しているコンテナ操作のアレコレはまだるっこしい感じに見えると思う。

特に配列や連想配列については zip/unzip や map のような標準的で気の利いた高階関数は用意されておらず，頑張って汎用パッケージを作ってみたところで実用的なパフォーマンスが得られずに打ち捨てられてしまうのがオチのようである。

リンク先の例にしても，結局 [Go 言語]では for 文で回していかざるを得ないのだから連想配列に格納するという発想自体を捨ててしまったほうが得策である。

というわけで，手遊びで書いてみた。

- [GitHub - spiegel-im-spiegel/csvtable: Demonstration for CSV/TSV Access](https://github.com/spiegel-im-spiegel/csvtable)

CSV/TSV は要するに行・列の2次元配列なんだから，以下のクラスを作って連想配列ではなく普通の配列で管理する。

```go
//CsvTable is CSV/TSV table class
type CsvTable struct {
    header []string
    col    map[string]int
    body   [][]string
}
```

その上でデータの読み込み時にヘッダの列名とカラム位置の関係を `col フィールドに`保持ってしまえばいいのである。

```go
//New returns new CsvTable instance
func New(r *csv.Reader) (*CsvTable, error) {
    ct := &CsvTable{}
    err := ct.readAll(r)
    return ct, err
}

func (ct *CsvTable) readAll(r *csv.Reader) error {
    if ct == nil {
        ct = &CsvTable{}
    }
    ct.col = map[string]int{}

    dt, err := r.ReadAll()
    if err != nil {
        return err
    }
    l := len(dt)
    if l > 0 {
        ct.header = dt[0]
        for i, e := range ct.header {
            ct.col[e] = i
        }

        if l > 1 {
            ct.body = dt[1:]
        }
    }
    return nil
}
```

main 側は以下のようになる。

```go
package main

import (
    "encoding/csv"
    "flag"
    "fmt"
    "io"
    "os"
    "strings"

    "github.com/spiegel-im-spiegel/csvtable"
)

func main() {
    flag.Parse()
    if flag.NArg() > 1 {
        fmt.Fprintln(os.Stderr, os.ErrInvalid)
        return
    }
    var r io.Reader
    if flag.NArg() == 0 {
        r = os.Stdin
    } else {
        f, err := os.Open(flag.Arg(0)) //maybe file path
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
        defer f.Close()
        r = f
    }

    cr := csv.NewReader(r)
    cr.Comma = '\t'
    cr.LazyQuotes = true       // a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
    cr.TrimLeadingSpace = true // leading white space in a field is ignored.

    ct, err := csvtable.New(cr)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    //fmt.Println("cols :", ct.Cols())
    //fmt.Println("rows :", ct.Rows())
    w := csv.NewWriter(os.Stdout)
    w.Comma = cr.Comma
    //w.UseCRLF = true
    //header, body := ct.OutputAll()
    header, body := ct.Output(strings.Split("city/temperature", "/"))
    w.Write(header)
    w.WriteAll(body)
}
```

念の為このパッケージの欠点を挙げておくと， CSV/TSV ファイルの内容の総てを一旦メモリ内に読み込んでいるため，巨大データを扱えないという問題がある。
実際問題として CSV/TSV データは数万行から数十万行の規模になることもザラにあるため，このままでは全く実用に耐えられないだろう。

## ブックマーク

- [文字エンコーディング変換]({{< ref "/golang/transform-character-encoding.md" >}})
- [配列と Slice]({{< ref "/golang/array-and-slice.md" >}})
- [Map の話]({{< ref "/golang/map.md" >}})
- [最大公約数と関数型プログラミング]({{< ref "/golang/greatest-common-divisor.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
