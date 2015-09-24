# はじめての Go 言語 (on Windows) その8

（[これまでの記事の目次](http://qiita.com/spiegel-im-spiegel/items/dca0df389df1470bdbfa#%E7%9B%AE%E6%AC%A1)）

[前回](http://qiita.com/spiegel-im-spiegel/items/64224f22ef17d916dc2d)の続き。パッケージ化したのならドキュメントを書きましょう。

## godoc のインストール

godoc は Go 言語のドキュメント化ツールです。 get コマンドで導入できます。（get コマンドについては「[その3](http://qiita.com/spiegel-im-spiegel/items/a52a47942fd3946bb583)」を参照）

```shell
C:>go get -v golang.org/x/tools/cmd/godoc
Fetching https://golang.org/x/tools/cmd/godoc?go-get=1
Parsing meta tags from https://golang.org/x/tools/cmd/godoc?go-get=1 (status code 200)
get "golang.org/x/tools/cmd/godoc": found meta tag main.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/cmd/godoc?go-get=1
get "golang.org/x/tools/cmd/godoc": verifying non-authoritative meta tag
Fetching https://golang.org/x/tools?go-get=1
Parsing meta tags from https://golang.org/x/tools?go-get=1 (status code 200)
golang.org/x/tools (download)
golang.org/x/tools/blog/atom
golang.org/x/tools/present
golang.org/x/tools/go/ast/astutil
golang.org/x/tools/go/exact
golang.org/x/tools/go/buildutil
golang.org/x/tools/go/types
golang.org/x/tools/container/intsets
golang.org/x/tools/godoc/vfs
golang.org/x/tools/godoc/util
golang.org/x/tools/godoc/vfs/httpfs
golang.org/x/tools/blog
golang.org/x/tools/godoc/redirect
golang.org/x/tools/godoc/static
golang.org/x/tools/godoc/vfs/gatefs
golang.org/x/tools/godoc/vfs/mapfs
golang.org/x/tools/godoc/vfs/zipfs
golang.org/x/tools/playground
golang.org/x/tools/go/types/typeutil
golang.org/x/tools/go/loader
golang.org/x/tools/go/ssa
golang.org/x/tools/go/callgraph
golang.org/x/tools/go/ssa/ssautil
golang.org/x/tools/go/pointer
golang.org/x/tools/godoc/analysis
golang.org/x/tools/godoc
golang.org/x/tools/cmd/godoc
```

ちなみに実行モジュールの `godoc.exe` は `%GOPATH\bin` フォルダではなく `%GOROOT%\bin` フォルダに格納されます。これで

```shell
C:>godoc time
```

とかやるとパッケージ（この場合は [time](http://golang.org/pkg/time/) パッケージ）のドキュメントが表示されるのですが，さすがにコマンドプロンプトでこれを見るのは辛いので， HTTP サービスを起動します。

```shell
C:>godoc -http=":3000"
```

これでブラウザで http://localhost:3000/ にアクセスするとドキュメントを見ることができます。

 ![godoc](https://c1.staticflickr.com/9/8788/18026303435_7b136c64bb.jpg "godoc")

たとえば [time](http://golang.org/pkg/time/) パッケージはこんなふうに表示されます。

 ![godoc: time package](https://c1.staticflickr.com/9/8847/18023061102_e5474f1ddc.jpg "godoc: time package")

本家サイトと同じですね。

### godoc で modjulian パッケージを見てみる

では「[その6](http://qiita.com/spiegel-im-spiegel/items/404871d2bafd22bdbb90)」で作った modjulian パッケージを get して godoc で見てみます。

```shell
C:>go get -v github.com/spiegel-im-spiegel/astrocalc/modjulian
github.com/spiegel-im-spiegel/astrocalc (download)
github.com/spiegel-im-spiegel/astrocalc/modjulian

C:>godoc -http=":3000"
```

 ![godoc; package list](https://c1.staticflickr.com/9/8831/18023689372_08795d4e8e.jpg "godoc; package list")


 ![godoc: modjulian package](https://c1.staticflickr.com/9/8835/17839022348_4315878c95.jpg "godoc: modjulian package")

全くコメントがないので，さすがに一覧には何もないですが，個別ページには最小限の情報が載っています。凄いなぁ。

## modjulian パッケージにドキュメント用のコメントを追記する

では，ソースコードを少しいじってドキュメント用のコードを追記してみましょう。

```go:modjulian.go
/**
 * Astronomical calculation for Golang.
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/deed.ja
 */

// modjulian パッケージは
// 修正ユリウス日（Modified Julian Date）の計算を行います。
package modjulian

import "time"

// DayNumber は
// 日付から修正ユリウス通日を取得します。
//
//   t := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
//   fmt.Print(modjulian.DayNumber(t)) //57023
//
// 時刻（時分秒）は無視します。
// 1970年1月1日以前のグレゴリオ暦では Fliegel の公式を使って計算します。
// 1970年1月1日以降は UNIX Time を用いて通日を取得します。
func DayNumber(t time.Time) int64 {
	if t.Sub(time.Unix(0, 0)) >= 0 {
		return dnUnix(t)
	} else {
		return dnGregorian(t)
	}
}

// dnGregorian は
// Fliegel の公式を使い，日付から修正ユリウス通日を計算します。
//
// 時刻（時分秒）は無視します。
func dnGregorian(t time.Time) int64 {
	y := int64(t.Year())
	m := int64(t.Month())
	if m < 3 {
		y -= 1
		m += 9
	} else {
		m -= 3
	}
	d := int64(t.Day()) - 1
	return (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
}

// dnUnix は
// UNIX Time で1970年1月1日からの通日を取得し，修正ユリウス通日を計算します。
//
// 時刻（時分秒）は無視します。
// 1970年1月1日以前の日付では正しく計算できません。
func dnUnix(t time.Time) int64 {
	const (
		onday   = int64(86400) //seconds
		baseDay = int64(40587) //Modified Julian Date at January 1, 1970
	)
	return (t.Unix())/onday + baseDay
}
```

 ![godoc; package list 2](https://c4.staticflickr.com/8/7732/18002607746_9990483503.jpg "godoc; package list 2")


 ![godoc: modjulian package 2](https://c1.staticflickr.com/9/8897/17408544863_f0f5649e32.jpg "godoc: modjulian package 2")

日本語ですみません。英語不得手なもので。

1. パッケージのコメントは `package` 指定の直前のコメントが有効になる。（ファイル先頭のコメントは反映されない）
1. パッケージリストの説明はパッケージ・コメントの最初の1文のみ表示される（日本語の句読点も理解しているらしい）
1. 関数等のコメントはそれぞれの記述の直前のコメントが有効になる。
1. 基本的に改行は無視される。ただし空行があれば別のパラグラフと理解しているようだ。
1. 空白文字2文字のインデントでコード記述領域（HTML 的には `<pre>` 要素）とみなしているらしい。コードを書く必要はないけど。

上の例では説明のためにコメント内にサンプルコードを載せましたが，サンプルコードを記述するのであればもっとスマートな方法があります。それはテストにサンプルコードを含める方法です。

```go:example_test.go
package modjulian_test

import (
	"fmt"
	"github.com/spiegel-im-spiegel/astrocalc/modjulian"
	"time"
)

func ExampleDayNumber() {
	t := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Print(modjulian.DayNumber(t))
	// Output:
	// 57023
}
```

このようなテストを作ると，ドキュメントが以下のようになります。

 ![godoc: modjulian package 3](https://c2.staticflickr.com/6/5327/17843944479_024f2f4073.jpg "godoc: modjulian package 3")

もちろん，テストもできます。

```shell
C:>go test -v ./...
=== RUN TestDayNumber
--- PASS: TestDayNumber (0.00s)
=== RUN: ExampleDayNumber
--- PASS: ExampleDayNumber (0.00s)
PASS
ok      github.com/spiegel-im-spiegel/astrocalc/modjulian       2.038s
```

この仕組みを使えばサンプルコードを常に最新の仕様にマッチさせることが可能になります。プログラマにとってドキュメントで一番欲しいのはサンプルコードなので，サンプルコードさえ正しければ，他はそれほど詳細に書かなくても推測できます。そういう意味で，このドキュメントシステムはなかなかおもしろいと思います。

## ブックマーク

- [Go で godoc を使えるようにする〜godoc のインストール〜 - Qiita](http://qiita.com/megu_ma/items/2066aef2f8c7f0ce2cc3)
- [Go言語のコードレビュー | SOTA](http://deeeet.com/writing/2014/05/26/go-code-review/)
- [Go コードのレビュー時によくされるコメント - build error](http://papaeye.tumblr.com/post/92328649161/go)
- [testingパッケージのExamplesについて - taknb2nchのメモ](http://d.hatena.ne.jp/taknb2nch/20131101/1383285018)
- [GoのExampleテストが便利 : swdyh](http://swdyh.tumblr.com/post/55771477125/go-example)
- [godoc.org への掲載方法を調べた - taknb2nchのメモ](http://d.hatena.ne.jp/taknb2nch/20140225/1393302743)
