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

[前回の記事]({{< relref "./walking-with-multi-core.md" >}} "saracen/walker で歩いてみる")を書いたあと，ちょっと思いついたので手遊びに [`filepath`].`Glob()` 関数の拡張版を書いてみた。

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

もちろんオリジナルの [`filepath`].`Glob()` 関数の [syntax](https://golang.org/pkg/path/filepath/#Match "filepath - The Go Programming Language") も使える。

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

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`filepath`]: https://golang.org/pkg/path/filepath/ "filepath - The Go Programming Language"
[`file`]: https://github.com/spiegel-im-spiegel/file "spiegel-im-spiegel/file: Extend filepath.Glob function"
[saracen/walker]: https://github.com/saracen/walker "saracen/walker: walker is a faster, parallel version, of filepath.Walk"

## ブックマーク

- [singleton method Dir.[] (Ruby 2.6.0)](https://docs.ruby-lang.org/ja/latest/method/Dir/s/glob.html)

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4873118468?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51pUKQajnaL._SL160_.jpg" width="125" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4873118468?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">Go言語による並行処理</a></dt>
    <dd>Katherine Cox-Buday (著), 山口 能迪 (翻訳)</dd>
    <dd>オライリージャパン 2018-10-26</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4873118468 (ASIN), 9784873118468 (EAN), 4873118468 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">この本を読んで「よっしゃ，明日から立派な goroutine 使いだ！」とはならないと思うけど，有象無象なコピペ・プログラマじゃなく，きちんと Go 言語のプログラミングを勉強したいのであれば，この本は必読書になると思う。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-11-03">2018-11-03</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
