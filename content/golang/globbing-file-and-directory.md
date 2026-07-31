+++
title = "拡張版 Glob 関数を書いてみた"
date =  "2019-10-27T18:08:38+09:00"
description = "ディレクトリを再帰的に検索するワイルドカードが使えるのが特徴で，内部で saracen/walker パッケージを使っている。とはいえ，全体的に素朴な作りになっているため，スピードは期待しないで欲しい（笑）"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "filepath" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回の記事]({{< relref "./walking-with-multi-core.md" >}} "saracen/walker で歩いてみる")を書いたあと，ちょっと思いついたので手遊びに [`filepath`]`.Glob()` 関数の拡張版を書いてみた。

- [spiegel-im-spiegel/file: Extend filepath.Glob function](https://github.com/spiegel-im-spiegel/file)

ディレクトリを再帰的に検索する `**/` ワイルドカードが使えるのが特徴で，内部で件の [saracen/walker] パッケージを使っている。
とはいえ，全体的に素朴な作りになっているため，スピードは期待しないで欲しい（笑）

たとえば以下のようなコードを書いてみる。

{{< highlight go "hl_lines=8 22" >}}
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/file"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if err := fs.Parse(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
    if fs.NArg() < 2 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}

	matches, err := file.Glob(fs.Arg(1), file.WithFlags(file.ContainsFile))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	fmt.Println("Count:", len(matches))
	for _ , path := range matches {
		fmt.Println(path)
	}
}
{{< /highlight >}}

これに対して `/usr/local/go/src/` ディレクトリ以下のファイルを取得したいなら，コマンドラインの引数に `/usr/local/go/src/**/*` を指定すればよい。
結果は以下の通り。

```text
$ go run work/sample.go "/usr/local/go/src/**/*"
Count: 6008
/usr/local/go/src/Make.dist
/usr/local/go/src/README.vendor
/usr/local/go/src/all.bash
/usr/local/go/src/all.bat
/usr/local/go/src/all.rc
/usr/local/go/src/archive/tar/common.go
/usr/local/go/src/archive/tar/example_test.go
/usr/local/go/src/archive/tar/format.go
/usr/local/go/src/archive/tar/reader.go
/usr/local/go/src/archive/tar/reader_test.go
/usr/local/go/src/archive/tar/stat_actime1.go
/usr/local/go/src/archive/tar/stat_actime2.go
...
```

もちろんオリジナルの [`filepath`]`.Glob()` 関数の [syntax](https://golang.org/pkg/path/filepath/#Match "filepath - The Go Programming Language") も使える。

```text
$ go run work/sample.go "/usr/local/go/src/**/*.[ch]"
Count: 77
/usr/local/go/src/cmd/go/testdata/src/badc/x.c
/usr/local/go/src/cmd/internal/goobj/testdata/mycgo/c1.c
/usr/local/go/src/cmd/internal/goobj/testdata/mycgo/c2.c
/usr/local/go/src/cmd/vendor/golang.org/x/sys/unix/gccgo_c.c
/usr/local/go/src/debug/dwarf/testdata/cycle.c
/usr/local/go/src/debug/dwarf/testdata/line1.c
/usr/local/go/src/debug/dwarf/testdata/line1.h
/usr/local/go/src/debug/dwarf/testdata/line2.c
/usr/local/go/src/debug/dwarf/testdata/ranges.c
/usr/local/go/src/debug/dwarf/testdata/split.c
/usr/local/go/src/debug/dwarf/testdata/typedef.c
/usr/local/go/src/debug/elf/testdata/hello.c
...
```

まぁ，私自身が趣味で書くのはフィルタ・プログラムが殆どで，ファイル検索機能にはあまりお世話にならないのだが（普段は shell スクリプト任せ），使えそうならご自由にどうぞ。
そうそう，「こうすれば効率よくなるよ」みたいな PR は大歓迎である。

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`filepath`]: https://golang.org/pkg/path/filepath/ "filepath - The Go Programming Language"
[`file`]: https://github.com/spiegel-im-spiegel/file "spiegel-im-spiegel/file: Extend filepath.Glob function"
[saracen/walker]: https://github.com/saracen/walker "saracen/walker: walker is a faster, parallel version, of filepath.Walk"

## ブックマーク

- [singleton method Dir.[] (Ruby 2.6.0)](https://docs.ruby-lang.org/ja/latest/method/Dir/s/glob.html)

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->

{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
