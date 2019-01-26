+++
draft = false
date = "2016-11-03T20:51:00+09:00"
update = "2016-11-08T20:41:58+09:00"
title = "Error の構造化"
description = "エラーハンドリングのために error を構造化する手段として github.com/pkg/errors というパッケージがあるそうだ。"
tags = [
  "golang",
  "error",
]

[author]
  name = "Spiegel"
  linkedin = "spiegelimspiegel"
  twitter = "spiegel_2007"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  url = "https://baldanders.info/spiegel/profile/"
  avatar = "/images/avatar.jpg"
  tumblr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr = "spiegel"
  license = "by-sa"
+++

今回は「[エラー・ハンドリングについて]({{< relref "error-handling.md" >}})」の続き。
とりあえず，エラーハンドリングのデモ用にこんなパッケージを考える。

```go
package errdemo1

import "os"

// F returns error.
func F() error {
    file, err := os.Open("not-exist.txt")
    if err != nil {
        return err
    }
    defer file.Close()

    return nil
}
```

呼び出し側の `main()` 関数では，このパッケージを以下のようにハンドリングする。

```go
package main

import (
    "fmt"
    "os"

    "errdemo/errdemo1"
)

func main() {
    if err := errdemo1.F(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    return
}
```

`not-exist.txt` が存在しない場合，実行結果は以下のようになる。

```text
$ go run main1.go
open not-exist.txt: The system cannot find the file specified.
```

まぁハンドリングというか `Error()` 関数が吐くエラー内容（文字列）を垂れ流してるだけだけど。
もし `errdemo1.F()` 関数が返す [error] の [type] を知りたければ [Conversion] 構文を使って

```go
package main

import (
    "fmt"
    "os"

    "errdemo/errdemo1"
)

func main() {
    if err := errdemo1.F(); err != nil {
        switch err.(type) {
        case *os.PathError:
            fmt.Fprintln(os.Stderr, "*os.PathError")
        default:
        }
        fmt.Fprintln(os.Stderr, err)
        return
    }

    return
}
```

のように書ける[^1]。
実行結果は

[^1]: 実際にはこのような [type] の判別はパッケージ側で提供すべきなのだろうが，今回はコードの比較のためにこんな書き方をしている。

```text
$ go run main2.go
*os.PathError
open not-exist.txt: The system cannot find the file specified.
```

となる。

この方法の欠点は大元の [error] を吐いた位置と [error] の伝達経路が分かりにくい点である。
`errdemo1.F()` 関数で [error] を拾ったら `errdemo1` パッケージ専用の [error] に差し替えて呼び出し側に返す方法もあるが，単に [error] を差し替えただけでは大元の [error] 情報が消失してしまう。
このようなことが起きるのは [error] 情報が構造化されていないことに原因がある。

エラーハンドリングのために [error] を構造化する手段として [`github.com/pkg/errors`] というパッケージがあるそうだ。

- [Golangのエラー処理とpkg/errors | SOTA](http://deeeet.com/writing/2016/04/25/go-pkg-errors/)

[`github.com/pkg/errors`] パッケージは `go get` コマンドで導入できる。

```text
$ go get -v github.com/pkg/errors
github.com/pkg/errors (download)
github.com/pkg/errors
```

あるパッケージで大元の [error] 情報を含んだ [error] をセットする場合は `errors.Wrap()` 関数を使う[^2]。

[^2]: 追加する文字列部分を書式文字列とパラメータで指定する `errors.Wrapf()` 関数も用意されている。

```go
package errdemo2

import (
    "os"

    "github.com/pkg/errors"
)

// F returns error.
func F() error {
    file, err := os.Open("not-exist.txt")
    if err != nil {
        return errors.Wrap(err, "Error by F() function")
    }
    defer file.Close()

    return nil
}
```

この `errdemo2` パッケージを呼び出す側を以下のように書けば

```go
package main

import (
    "fmt"
    "os"

    "errdemo/errdemo2"
)

func main() {
    if err := errdemo2.F(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    return
}
```

実行結果は以下のようになる。

```text
$ go run main3.go
Error by F() function: open not-exist.txt: The system cannot find the file specified.
```

また大元の [error] を取得したければ `errors.Cause()` 関数を使う。

```go
package main

import (
    "fmt"
    "os"

    "errdemo/errdemo2"

    "github.com/pkg/errors"
)

func main() {
    if err := errdemo2.F(); err != nil {
        switch errors.Cause(err).(type) {
        case *os.PathError:
            fmt.Fprintln(os.Stderr, "*os.PathError")
        default:
        }
        fmt.Fprintln(os.Stderr, err)
        return
    }

    return
}
```

この場合の実行結果は以下のようになる。

```text
$ go run main4.go
*os.PathError
Error by F() function: open not-exist.txt: The system cannot find the file specified.
```

ちなみに `errors.Cause()` 関数の中身はこんな感じ。

```go
// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
    type causer interface {
        Cause() error
    }

    for err != nil {
        cause, ok := err.(causer)
        if !ok {
            break
        }
        err = cause.Cause()
    }
    return err
}
```

つまり `Cause()` 関数を持つ [error] インスタンスであれば `Cause()` 関数を辿って大元の [error] インスタンスを返すが， `Cause()` 関数がない場合はそのまま引数の [error] インスタンスを返す。

このように [`github.com/pkg/errors`] パッケージを使えば [error] を構造的に，かつ手軽に扱うことができる。
とても便利なパッケージなので是非活用したいところである。

## ブックマーク

- [Golangでエラー時にスタックトレースを表示する - Qiita](http://qiita.com/deeeet/items/dacc71932393ab35d9f8)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[error]: http://blog.golang.org/error-handling-and-go "Error handling and Go - The Go Blog"
[type]: https://golang.org/ref/spec#Properties_of_types_and_values "Properties of types and values"
[Conversion]: https://golang.org/ref/spec#Conversions "The Go Programming Language Specification - The Go Programming Language"
[`github.com/pkg/errors`]: https://github.com/pkg/errors "pkg/errors: Simple error handling primitives"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
