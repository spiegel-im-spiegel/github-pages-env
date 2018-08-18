+++
date = "2015-09-19T23:43:25+09:00"
update = "2018-05-15T20:42:53+09:00"
description = "パッケージ化したのならドキュメントを書きましょう。"
draft = false
tags = ["golang", "godoc"]
title = "Go 言語のドキュメント・フレームワーク"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（初出： [はじめての Go 言語 (on Windows) その8 - Qiita](http://qiita.com/spiegel-im-spiegel/items/5f9e96f226f46089388f)）

[前回]の続き。
パッケージ化したのならドキュメントを書きましょう。

## godoc のインストール

[`godoc`] は [Go 言語]のドキュメント化ツールです。
`go get` コマンドで導入できます。（`go get` コマンドについては「[go get コマンドでパッケージを管理する]({{< relref "go-get-package.md" >}})」を参照）

```
C:>mkdir C:\workspace\godoc

C:>cd C:\workspace\godoc

C:\workspace\godoc>SET GOPATH=C:\workspace\godoc

C:\workspace\godoc>go get -v golang.org/x/tools/cmd/godoc
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
golang.org/x/tools/blog
golang.org/x/tools/godoc/vfs
golang.org/x/tools/godoc/redirect
golang.org/x/tools/godoc/static
golang.org/x/tools/playground
golang.org/x/tools/godoc/util
golang.org/x/tools/go/types/typeutil
golang.org/x/tools/go/loader
golang.org/x/tools/godoc/vfs/httpfs
golang.org/x/tools/godoc/vfs/gatefs
golang.org/x/tools/godoc/vfs/mapfs
golang.org/x/tools/godoc/vfs/zipfs
golang.org/x/tools/go/ssa
golang.org/x/tools/go/callgraph
golang.org/x/tools/go/ssa/ssautil
golang.org/x/tools/go/pointer
golang.org/x/tools/godoc/analysis
golang.org/x/tools/godoc
golang.org/x/tools/cmd/godoc
```

ちなみに `godoc` の実行モジュールは `%GOPATH\bin` フォルダではなく `%GOROOT%\bin` フォルダに格納されます。
これで

```
C:>godoc time
```

とかやるとパッケージ（この場合は [`time`] パッケージ）のドキュメントが表示されるのですが，さすがにコマンドプロンプトでこれを見るのは辛いので， HTTP サービスを起動します。

```
C:>godoc -http=":3000"
```

これでブラウザで http://localhost:3000/ にアクセスするとドキュメントを見ることができます。

{{< fig-img src="https://farm9.staticflickr.com/8788/18026303435_7b136c64bb.jpg" alt="godoc" title="godoc" caption="godoc" link="https://www.flickr.com/photos/spiegel/18026303435/" >}}

たとえば [`time`] パッケージはこんなふうに表示されます。

{{< fig-img src="https://farm9.staticflickr.com/8847/18023061102_e5474f1ddc.jpg" alt="godoc: time package" title="godoc: time package" caption="godoc: time package" link="https://www.flickr.com/photos/spiegel/18023061102/" >}}

本家サイトと同じですね。

### godoc で modjulian パッケージを見てみる

「[機能のパッケージ化]({{< relref "packaging.md" >}})」で作った [`github.com/spiegel-im-spiegel/astrocalc`]`/modjulian` はどうなっているでしょう。

{{< fig-img src="https://farm9.staticflickr.com/8831/18023689372_08795d4e8e.jpg" alt="godoc: package list" title="godoc: package list" caption="godoc: package list" link="https://www.flickr.com/photos/spiegel/18023689372/" >}}

{{< fig-img src="https://farm9.staticflickr.com/8835/17839022348_4315878c95.jpg" alt="godoc: modjulian package" title="godoc: modjulian package" caption="godoc: modjulian package" link="https://www.flickr.com/photos/spiegel/17839022348/" >}}

全くコメントがないので，さすがに一覧には何もないですが，個別ページには最小限の情報が載っています。凄いなぁ。

## modjulian パッケージにドキュメント用のコメントを追記する

では，ソースコードを少しいじってドキュメント用のコードを追記してみましょう。

```go
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

{{< fig-img src="https://farm8.staticflickr.com/7732/18002607746_9990483503.jpg" alt="godoc: package list" title="godoc: package list" caption="godoc: package list" link="https://www.flickr.com/photos/spiegel/18002607746/" >}}

{{< fig-img src="https://farm9.staticflickr.com/8897/17408544863_f0f5649e32.jpg" alt="godoc: modjulian 2" title="godoc: modjulian 2" caption="godoc: modjulian 2" link="https://www.flickr.com/photos/spiegel/17408544863/" >}}

日本語ですみません。
英語不得手なもので。

1. パッケージのコメントは `package` 指定の直前のコメントが有効になる。（ファイル先頭のコメントは反映されない）
1. パッケージリストの説明はパッケージ・コメントの最初の1文のみ表示される（日本語の句読点も理解しているらしい）
1. 関数等のコメントはそれぞれの記述の直前のコメントが有効になる。
1. 基本的に改行は無視される。ただし空行があれば別のパラグラフと理解しているようだ。
1. 空白文字2文字のインデントでコード記述領域（HTML 的には `<pre>` 要素）とみなしているらしい。コードを書く必要はないけど。

上の例では説明のためにコメント内にサンプルコードを載せましたが，サンプルコードを記述するのであればもっとスマートな方法があります。
それはテストにサンプルコードを含める方法です。

```go
package modjulian_test

import (
    "fmt"
    "time"

    "github.com/spiegel-im-spiegel/astrocalc/modjulian"
)

func ExampleDayNumber() {
    t := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
    fmt.Print(modjulian.DayNumber(t))
    // Output:
    // 57023
}
```

このようなテスト `example_test.go` を作ると，ドキュメントが以下のようになります。

{{< fig-img src="https://farm6.staticflickr.com/5327/17843944479_024f2f4073.jpg" alt="godoc: modjulian 3" title="godoc: modjulian 3" caption="godoc: modjulian 3" link="https://www.flickr.com/photos/spiegel/17843944479/" >}}


もちろん，テストもできます。

```shell
C:\workspace\jd>go test -v github.com/spiegel-im-spiegel/astrocalc/modjulian
=== RUN   TestDayNumber
--- PASS: TestDayNumber (0.00s)
=== RUN   ExampleDayNumber
--- PASS: ExampleDayNumber (0.00s)
PASS
ok      github.com/spiegel-im-spiegel/astrocalc/modjulian       2.755s
```

この仕組みを使えばサンプルコードを常に最新の仕様にマッチさせることが可能になります。
プログラマにとってドキュメントで一番欲しいのはサンプルコードなので，サンプルコードさえ正しければ，他はそれほど詳細に書かなくても推測できます。
そういう意味で，このようなテストと連動したドキュメント・フレームワークはなかなかおもしろいと思います。

## ブックマーク{#bookmark}

- [Go で godoc を使えるようにする〜godoc のインストール〜 - Qiita](http://qiita.com/megu_ma/items/2066aef2f8c7f0ce2cc3)
- [Go言語のコードレビュー | SOTA](http://deeeet.com/writing/2014/05/26/go-code-review/)
- [Go コードのレビュー時によくされるコメント - build error](http://papaeye.tumblr.com/post/92328649161/go)
- [testingパッケージのExamplesについて - taknb2nchのメモ](http://d.hatena.ne.jp/taknb2nch/20131101/1383285018)
- [GoのExampleテストが便利 : swdyh](http://swdyh.tumblr.com/post/55771477125/go-example)
- [godoc.org への掲載方法を調べた - taknb2nchのメモ](http://d.hatena.ne.jp/taknb2nch/20140225/1393302743)
- [GoDocドキュメントで知っていると便利な機能 - Qiita](https://qiita.com/lufia/items/97acb391c26f967048f1) : よくまとまっている

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< relref "testing.md" >}} "Go 言語のテスト・フレームワーク"
[`godoc`]: http://godoc.org/golang.org/x/tools/cmd/godoc "godoc - GoDoc"
[`time`]: http://golang.org/pkg/time/
[`github.com/spiegel-im-spiegel/astrocalc`]: https://github.com/spiegel-im-spiegel/astrocalc
