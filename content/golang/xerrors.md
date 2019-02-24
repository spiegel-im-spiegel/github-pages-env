+++
title = "階層化 Error パッケージ “xerrors” を試してみる"
date = "2019-02-23T22:31:17+09:00"
description = "これは “Go 2 Draft” の “Error Inspection” を実装したもので Go 1.13 以降で既存の標準 errors パッケージに組み込む計画があるらしい。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

新しい error パッケージ [golang.org/x/xerrors] がリリースされたそうな。
これはいわゆる “[Go 2 Draft](https://go.googlesource.com/proposal/+/master/design/go2draft.md)” の “[Error Inspection](https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md)” を実装したもので，なんと Go 1.13 以降で既存の標準 [`errors`] パッケージに組み込む計画があるらしい。

{{< fig-quote title="xerrors - GoDoc" link="https://godoc.org/golang.org/x/xerrors" lang="en" >}}
<q>Most of the functions and types in this package will be incorporated into the standard library's errors package in Go 1.13; the behavior of this package's Errorf function will be incorporated into the standard library's fmt.Errorf. Use this package to get equivalent behavior in all supported Go versions.</q>
{{< /fig-quote >}}

これは朗報！ というわけで早速試してみることにした。

1. [Error インスタンスの生成]({{< relref "#new" >}})
1. [Error の階層化]({{< relref "#wrap" >}})
1. [Error の等値性]({{< relref "#is" >}})
1. [階層化された Error を検索する]({{< relref "#as" >}})
1. [独自の階層化 error 型を作成する]({{< relref "#custom" >}})

## Error インスタンスの生成{#new}

さっそく簡単なコードを書いてみよう。

```go
package main

import (
    "fmt"
    "os"

    "golang.org/x/xerrors"
)

func foo() error {
    return xerrors.New("an error instance")
}
func main() {
    if err := foo(); err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }
}
```

これを実行すると

```text
$ go run demo1/demo1.go
an error instance
```

となる。
ここまでは普通。
ここで `fmt.Fprintf()` 関数のフォーマット文字列を以下のように書き換えてみる。

{{< highlight go "hl_lines=15" >}}
package main

import (
    "fmt"
    "os"

    "golang.org/x/xerrors"
)

func foo() error {
    return xerrors.New("an error instance")
}
func main() {
    if err := foo(); err != nil {
        fmt.Fprintf(os.Stderr, "%+v\n", err)
    }
}
{{< /highlight  >}}

これ実行すると

```text
$ go run demo1/demo1.go
an error instance:
    main.foo
        /tmp/xerrors/demo1/demo1.go:11
```

となり error の発生箇所が表示されるようになった。

ちなみに [`xerrors`]`.New()` 関数で生成される error インスタンスの構造は以下の通り。

```go
// errorString is a trivial implementation of error.
type errorString struct {
    s     string
    frame Frame
}

// New returns an error that formats as the given text.
//
// The returned error contains a Frame set to the caller's location and
// implements Formatter to show this information when printed with details.
func New(text string) error {
    return &errorString{text, Caller(1)}
}
```

`frame` フィールドに [`xerrors`]`.New()` 関数の呼び出し情報が格納されているのが分かると思う。
これでデバッグ作業がかなり楽になるだろう。

## Error の階層化{#wrap}

たとえば以下のようなファイルをオープンするだけの簡単なコードを書いてみる。

```go
package main

import (
    "fmt"
    "os"
)

func fileOpen(fname string) error {
    file, err := os.Open(fname)
    if err != nil {
        switch e := err.(type) {
        case *os.PathError:
            return fmt.Errorf("Error in fileOpen(\"%v\"): %v", e.Path, e.Err)
        default:
            return fmt.Errorf("Error in fileOpen(): %v", err)
        }
    }
    defer file.Close()
    return nil
}

func main() {
    fmt.Print("Result: ")
    if err := fileOpen("null.txt"); err != nil {
        fmt.rintf("%+v\n", err)
    } else {
        fmt.Println("OK")
    }
}
```

このとき `null.txt` が存在しないなら実行結果は

```text
$ go run demo2a/demo2a.go
Result: Error in fileOpen("null.txt"): The system cannot find the file specified.
```

となる。
パッと見は問題なさそうだが `fileOpen()` 関数が error を返す際に [`os`]`.Open()` 関数が吐き出した error インスタンスが捨てられてしまうため，エラーの追跡が難しくなる。

そこで [`xerrors`]`.Errorf()` 関数を使って error のラッピングを行う。
コードはこんな感じ。

{{< highlight go "hl_lines=7 15 17" >}}
package main

import (
    "fmt"
    "os"

    "golang.org/x/xerrors"
)

func fileOpen(fname string) error {
    file, err := os.Open(fname)
    if err != nil {
        switch e := err.(type) {
        case * os.PathError:
            return xerrors.Errorf("Error in fileOpen(\"%v\"): %w", e.Path, e.Err)
        default:
            return xerrors.Errorf("Error in fileOpen(): %w", err)
        }
    }
    defer file.Close()
    return nil
}

func main() {
    fmt.Print("Result: ")
    if err := fileOpen("null.txt"); err != nil {
        fmt.Printf("%+v\n", err)
    } else {
        fmt.Println("OK")
    }
}
{{< /highlight  >}}

少し解説すると [`xerrors`]`.Errorf()` 関数の第1引数のフォーマット文字列の末尾が `": %w"` になっていて，かつ対応する値が error インタフェースを備えていれば [`xerrors`]`.wrapError` 型のインスタンスを返す。
ただの `"%w"` では [`xerrors`]`.wrapError` 型を返さない点に注意[^wrap1]。

[^wrap1]: なんでこんなヘンテコな仕様になっているかというと Go 1.13 で  [`xerrors`]`.Errorf()` 関数を  [`fmt`]`.Errorf()` 関数に統合する予定があるからだそうだ。なお [`xerrors`]`.Errorf()` 関数では error インスタンスをひとつしかラッピングできない。複数の error インスタンスをまとめてラッピングしたいなら独自の型（クラス）を定義する必要があるだろう。それでも独自型をフルスクラッチで組むよりは簡単だろうけど。

このコードを実行すると以下のような結果になる。

```text
$ go run demo2b/demo2b.go
Result: Error in fileOpen("null.txt"):
    main.fileOpen
        /tmp/xerrors/demo2b/demo2b.go:15
  - The system cannot find the file specified.
```

error が連結され階層構造になっているのが分かると思う。

error を階層化するためには [`xerrors`]`.Wrapper` インタフェースを実装する必要がある。
[`xerrors`]`.Wrapper` インタフェースの定義は以下の通り。

```go
// A Wrapper provides context around another error.
type Wrapper interface {
    // Unwrap returns the next error in the error chain.
    // If there is no next error, Unwrap returns nil.
    Unwrap() error
}
```

UML で描くとこんな感じかな。

{{< fig-img src="./wraperror.png" title="wraperror.png" link="./wraperror.puml" width="971" >}}

ちなみに [`xerrors`]`.wrapError` 型では以下のような実装になっている。

```go
type wrapError struct {
    msg   string
    err   error
    frame Frame
}

func (e *wrapError) Error() string {
    return fmt.Sprint(e)
}

func (e *wrapError) Unwrap() error {
    return e.err
}
```

## Error の等値性{#is}

上述した [`xerrors`]`.Wrapper` インタフェースがなんの役に立つかというと error インスタンスの等値性（equality）をチェックするのに役立つのだ。

error インスタンスの等値性を調べるには [`xerrors`]`.Is()` 関数を使う。
[`xerrors`]`.Is()` 関数の中身はこんな感じ。

```go
// Is reports whether any error in err's chain matches target.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
func Is(err, target error) bool {
    if target == nil {
        return err == target
    }
    for {
        if err == target {
            return true
        }
        if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
            return true
        }
        // TODO: consider supporing target.Is(err). This would allow
        // user-definable predicates, but also may allow for coping with sloppy
        // APIs, thereby making it easier to get away with them.
        if err = Unwrap(err); err == nil {
            return false
        }
    }
}
```

`err` に対して連結されている error インスタンスを遡っていき `target` と等値[^is1] なインスタンスを探している。
連結されている error チェインの中にひとつでも等値な error インスタンスがあるなら対象のインスタンスは等値であるとみなすわけだ。

[^is1]: [Go 言語]では `==` は等値演算子（equality operator）だが error インスタンスはポインタ値で表すことが多く，その場合は `==` 演算子もポインタ値を比較することになり，実質的に error インスタンスの同一性（identity）を調べていることになる。 [Go 言語]の interface 型は変数がインスタンスそのものかインスタンスへの参照（ポインタ値）かを隠蔽してしまうため，等値か同一かを問題にする際には注意が必要である。

先程のファイルをオープンするコードを [`xerrors`]`.Is()` 関数を使って少し書き直してみよう。

{{< highlight go "hl_lines=28-32" >}}
package main

import (
    "fmt"
    "os"
    "syscall"

    "golang.org/x/xerrors"
)

func fileOpen(fname string) error {
    file, err := os.Open(fname)
    if err != nil {
        switch e := err.(type) {
        case * os.PathError:
            return xerrors.Errorf("Error in fileOpen(\"%v\"): %w", e.Path, e.Err)
        default:
            return xerrors.Errorf("Error in fileOpen(): %w", err)
        }
    }
    defer file.Close()
    return nil
}

func main() {
    fmt.Print("Result: ")
    if err := fileOpen("null.txt"); err != nil {
        if xerrors.Is(err, syscall.ENOENT) {
            fmt.Println("ファイルが存在しない。")
        } else {
            fmt.Println("その他のエラー:", err)
        }
    } else {
        fmt.Println("OK")
    }
}
{{< /highlight  >}}

これの実行結果は以下の通り。

```text
$ go run demo3/demo3.go
Result: ファイルが存在しない。
```

## 階層化された Error を検索する{#as}

 [`xerrors`]`.As()` 関数を使うと階層化された Error の中から指定した型の error インスタンスを抽出できる。
 これも [`xerrors`]`.Wrapper` インタフェースが実装されていることが前提となる。

```go
// As finds the first error in err's chain that matches the type to which target
// points, and if so, sets the target to its value and returns true. An error
// matches a type if it is assignable to the target type, or if it has a method
// As(interface{}) bool such that As(target) returns true. As will panic if target
// is not a non-nil pointer to a type which implements error or is of interface type.
//
// The As method should set the target to its value and return true if err
// matches the type to which target points.
func As(err error, target interface{}) bool {
    if target == nil {
        panic("errors: target cannot be nil")
    }
    val := reflect.ValueOf(target)
    typ := val.Type()
    if typ.Kind() != reflect.Ptr || val.IsNil() {
        panic("errors: target must be a non-nil pointer")
    }
    if e := typ.Elem(); e.Kind() != reflect.Interface && !e.Implements(errorType) {
        panic("errors: *target must be interface or implement error")
    }
    targetType := typ.Elem()
    for {
        if reflect.TypeOf(err).AssignableTo(targetType) {
            val.Elem().Set(reflect.ValueOf(err))
            return true
        }
        if x, ok := err.(interface{ As(interface{}) bool }); ok && x.As(target) {
            return true
        }
        if err = Unwrap(err); err == nil {
            return false
        }
    }
}

var errorType = reflect.TypeOf((*error)(nil)).Elem()
```

“[Error Inspection](https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md)” では総称型（Generics）を前提とした実装が提案されていたが，まだ[総称型は実現されていない]({{< relref "./generics-in-go-2.md" >}} "次期 Go 言語で導入される（かもしれない）総称型について予習する")ので，ちょっとアレな感じがするのは致し方ない（笑）

[`xerrors`]`.As()` 関数を使ってコードを書き直してみよう。

{{< highlight go "hl_lines=28-38" >}}
package main

import (
    "fmt"
    "os"
    "syscall"

    "golang.org/x/xerrors"
)

func fileOpen(fname string) error {
    file, err := os.Open(fname)
    if err != nil {
        switch e := err.(type) {
        case * os.PathError:
            return xerrors.Errorf("Error in fileOpen(\"%v\"): %w", e.Path, e.Err)
        default:
            return xerrors.Errorf("Error in fileOpen(): %w", err)
        }
    }
    defer file.Close()
    return nil
}

func main() {
    fmt.Print("Result: ")
    if err := fileOpen("null.txt"); err != nil {
        var errno syscall.Errno
        if xerrors.As(err, &errno) {
            switch errno {
            case syscall.ENOENT:
                fmt.Println("ファイルが存在しない。")
            default:
                fmt.Println("Errno =", errno)
            }
        } else {
            fmt.Println("その他のエラー:", err)
        }
    } else {
        fmt.Println("OK")
    }
}
{{< /highlight  >}}

これの実行結果は以下の通り。

```text
$ go run demo3/demo3.go
Result: ファイルが存在しない。
```

## 独自の階層化 error 型を作成する{#custom}

ここまで分かったので，次は [`xerrors`] 互換の階層化 error 型を自作してみよう。
具体的には [`os`]`.PathError` をカスタマイズした型を書いてみる。

```go
type CustomPathError struct {
    Op    string
    Path  string
    Err   error
    frame xerrors.Frame
}

func (e *CustomPathError) Error() string {
    return "error in " + e.Op + " \"" + e.Path + "\""
}

func (e *CustomPathError) Unwrap() error {
    return e.Err
}

func (e *CustomPathError) Format(s fmt.State, v rune) {
    xerrors.FormatError(e, s, v)
}

func (e *CustomPathError) FormatError(p xerrors.Printer) error {
    p.Print(e.Error())
    e.frame.Format(p)
    return e.Err
}
```

`CustomPathError` が [`os`]`.PathError` の階層化 error 版で，これに `Error()`, `Unwrap()`, `Format()`, `FormatError()` の各関数が実装されている。
`Format()` および `FormatError()` 関数は [`xerrors`]`.Formatter` および [`fmt`]`.Formatter` インタフェースの実装で [`fmt`]`.Printf()` などの関数で呼び出される。

[`fmt`]`.Formatter` は以下のように定義されている。

```go
// Formatter is the interface implemented by values with a custom formatter.
// The implementation of Format may call Sprint(f) or Fprint(f) etc.
// to generate its output.
type Formatter interface {
    Format(f State, c rune)
}
```

また [`xerrors`]`.Formatter` は以下のように定義されている。

```go
// A Formatter formats error messages.
type Formatter interface {
    error

    // FormatError prints the receiver's first error and returns the next error in
    // the error chain, if any.
    FormatError(p Printer) (next error)
}
```

`CustomPathError` を UML で描くとこんな感じ。

{{< fig-img src="./custumerror.png" title="custumerror.png" link="./custumerror.puml" width="1912" >}}

この `CustomPathError` を使って [`os`]`.Open()` 関数の返り値の error をラップする。

```go
func OpenWrapper(name string) (*os.File, error) {
    file, err := os.Open(name)
    if err != nil {
        switch e := err.(type) {
        case *os.PathError:
            return file, &CustomPathError{Op: e.Op, Path: e.Path, Err: e.Err, frame: xerrors.Caller(0)}
        default:
            return file, &CustomPathError{Op: "open", Path: name, Err: err, frame: xerrors.Caller(0)}
        }
    }
    return file, nil
}
```

さらに `OpenWrapper()` を使って先程のファイルをオープンするコードを書き直す。

```go
func fileOpen(fname string) error {
    file, err := OpenWrapper(fname)
    if err != nil {
        return xerrors.Errorf("Error in fileOpen(): %w", err)
    }
    defer file.Close()
    return nil
}

func main() {
    fmt.Print("Result: ")
    if err := fileOpen("null.txt"); err != nil {
        fmt.Printf("%v\n", err)
        fmt.Printf("%+v\n", err)
        if xerrors.Is(err, syscall.ENOENT) {
            fmt.Println("ファイルが存在しない。")
        } else {
            fmt.Println("その他のエラー:", err)
        }
    } else {
        fmt.Println("OK")
    }
}
```

これを実行してみよう。

```text
$ go run demo5/demo5.go
Result: Error in fileOpen(): error in open "null.txt": The system cannot find the file specified.
Error in fileOpen():
    main.fileOpen
        /tmp/xerrors/demo5/demo5.go:52
  - error in open "null.txt":
    main.OpenWrapper
        /tmp/xerrors/demo5/demo5.go:41
  - The system cannot find the file specified.
ファイルが存在しない。
```

ちゃんと error が連結されていることが分かると思う。

はっきり言って `Format()` と `FormatError()` の両関数は [golang.org/x/xerrors] のソースコードからパクっているが，細かいチューニングが必要でないのなら，このまま snippet として使い回せるんじゃないだろうか[^license1]。

[^license1]: ちなみに [golang.org/x/xerrors] は BSD ライセンス下で利用できる。

## ブックマーク

- [xerrors - 関連情報 - Qiita](https://qiita.com/sonatard/items/802db82e7275f17fe702)
    - [Goの新しいerrors パッケージ xerrors(Go 1.13からは標準のerrorsパッケージに入る予定) - Qiita](https://qiita.com/sonatard/items/9c9faf79ac03c20f4ae1)
    - [xerrors パッケージ - 独自に定義したエラー型はIsメソッドとAsメソッドでデフォルトの振る舞いを変更可能 - Qiita](https://qiita.com/sonatard/items/c6da1a895b2ca7e0f13f)
    - [xerrors パッケージで途中に独自定義したエラー型をラップする方法 - Qiita](https://qiita.com/sonatard/items/e0f866f44717c0501f7a)

- [エラー・ハンドリングについて]({{< relref "./error-handling.md" >}})
- [Error の構造化]({{< relref "./error-handling2.md" >}})
- [次期 Go 言語で導入される（かもしれない）新しいエラー・ハンドリングについて予習する]({{< relref "./error-handling-in-go-2.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[golang.org/x/xerrors]: https://godoc.org/golang.org/x/xerrors "xerrors - GoDoc"
[`xerrors`]: https://godoc.org/golang.org/x/xerrors "xerrors - GoDoc"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"

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
