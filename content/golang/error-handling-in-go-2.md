+++
title = "次期 Go 言語で導入される（かもしれない）新しいエラー・ハンドリングについて予習する"
date = "2018-11-13T22:48:47+09:00"
description = "総称型なんか後回しにしてこっちを先に実現してほしい。"
image = "/images/attention/go-logo_blue.png"
tags = [
  "golang",
  "error",
  "programming",
  "engineering",
]

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は「[次期 Go 言語で導入される（かもしれない）総称型について予習する]({{< relref "generics-in-go-2.md" >}})」の続き。

次期 [Go 言語]で追加される（かもしれない）仕様についてもう一度挙げておこう。

- 総称型（generics）
    - [Generics — Problem Overview](https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md)
    - [Contracts — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-contracts.md)
- エラー・ハンドリングに関するもの
    - [Error Handling — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)
    - [Error Inspection — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md)
    - [Error Printing — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-error-printing.md)

この記事ではエラー・ハンドリングについて予習してみる。
はっきり言って私は物凄く期待している。
総称型なんか後回しにしてこっちを先に実現してほしい。

なお “Go 2” の提案はまだドラフト段階なので大幅に変更になったり場合によっては立ち消えになる可能性もある。
なので，この記事では深いところまで踏み込まずフワっとした説明になるけど，あしからずご了承の程を。

{{< div-box type="markdown" >}}
## 追記 2019-08-24

[Go] 1.13 で後半の “[Wrapper interface]({{< relref "#wrapper" >}})” に関連する仕様追加が行われた。
詳しくは以下の記事を参照のこと。

- [Go 1.13 のエラー・ハンドリング]({{< relref "./error-handling-in-go-1_3.md" >}})

[Go]: https://go.dev/
{{< /div-box >}}

## Check 式（Check Expression）と Handle 構文（Handle Statement）

まずはファイルをコピーする簡単なコマンドを書いてみよう。
ちなみにこれは完全に動くコードである。

```go
package main

import (
    "flag"
    "fmt"
    "io"
    "os"
)

func copyFile(src, dst string) error {
    r, err := os.Open(src)
    if err != nil {
        return err
    }
    defer r.Close()

    w, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer w.Close()

    if _, err := io.Copy(w, r); err != nil {
        return err
    }
    return nil
}

func main() {
    flag.Parse()
    if flag.NArg() != 2 {
        fmt.Println(os.ErrInvalid)
        return
    }
    if err := copyFile(flag.Arg(0), flag.Arg(1)); err != nil {
        fmt.Println(err)
        return
    }
    return
}
```

この中で

```go
if err != nil {
    ...
}
```

という記述が多数見られるのが分かると思う。
[Go 言語]は C++ や Java で言うところの例外処理の仕組みを持っていない[^ec1] ためにこのような記述になるのだが，こうした単純な繰り返しの記述は [Go 言語]プログラマの間でも不評のようだ。

[^ec1]: [Go 言語]の panic は例外処理に似た大域脱出の機能を持っているが，本来はリカバリ不能なエラーや障害が発生した際に迅速にプロセスを終了させるための仕組みなので，例外処理のような使い方をすべきではないとされている。

これを解消するのが check 式と handle 構文である。

たとえば，関数の返り値に error を含む場合

```go
v1, v2, ..., vn, err := foo()
```

check 式を使って error を検知し残りの返り値を返す事ができる。

```go
v1, v2, ..., vn := chack foo()
```

検知した error はどうなるかというと直近の handle 構文で指定された処理へ飛ぶ[^hdl1]。

[^hdl1]: 直近に handle 構文がない場合には error を吸い込んだまま何もせずにスルーするようだ。

```go
handle err {
    ...
}
```

では check と handle を使って先程の `copyFile()` 関数を書き直してみよう。

```go
func copyFile(src, dst string) error {
    handle err {
        return err
    }

    r := check os.Open(src)
    defer r.Close()

    w := check os.Create(dst)
    defer w.Close()

    check io.Copy(w, r)
    return nil
}
```

随分とスッキリした。
[こうかはばつぐんだ！](https://dic.pixiv.net/a/%E3%81%93%E3%81%86%E3%81%8B%E3%81%AF%E3%81%B0%E3%81%A4%E3%81%90%E3%82%93%E3%81%A0%21)

Check は式なので

```go
res := check foo(check bar())
```

といった書き方もできる。
関数の返り値が error とタプルになっている場合は（いったん変数に流し込んだり）スマートでない記述になっているので，これは嬉しい。

Handle 構文はいくつでも書くことができる。
たとえば

```go
func process(user string, files chan string) (n int, err error) {
    handle err { return 0, fmt.Errorf("process: %v", err)  }      // handler A
    for i := 0; i < 3; i++ {
        handle err { err = fmt.Errorf("attempt %d: %v", i, err) } // handler B
        handle err { err = moreWrapping(err) }                    // handler C

        check do(something())  // check 1: handler chain C, B, A
    }
    check do(somethingElse())  // check 2: handler chain A
}
```

のように書けるらしい。
Handle 構文の処理はスタック状に積まれていく感じかな。

## Wrapper interface{#wrapper}

たとえば [`os`]`.PathError` は以下のように内部に error 情報を持っている。

```go
// PathError records an error and the operation and file path that caused it.
type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
```

エラーハンドリングを行う際は，この内部の error を見て原因となるエラー情報を取得することができる。
このような構造になっている error オブジェクトは多そうである。
そこで [`errors`] パッケージに Wrapper interface を追加することを考える。

```go
package errors

// A Wrapper is an error implementation
// wrapping context around another error.
type Wrapper interface {
    // Unwrap returns the next error in the error chain.
    // If there is no next error, Unwrap returns nil.
    Unwrap() error
}
```

error オブジェクトが `Unwrap()` 関数を用意していれば，この関数を使って原因となる error オブジェクトを取得できるというわけだ。

```go
func (e *PathError) Unwrap() error { return e.Err }
```

これを踏まえて以下の関数も用意する。

```go
func As(type E)(err error) (e E, ok bool) {
    for {
        if e, ok := err.(E); ok {
            return e, true
        }
        wrapper, ok := err.(Wrapper)
        if !ok {
            return e, false
        }
        err = wrapper.Unwrap()
        if err == nil {
            return e, false
        }
    }
}
```

この `As()` 関数は[総称型] `E` を含んでいる点に注目。
これを使えば

```go
if pe, ok := errors.As(*os.PathError)(err); ok {
    if errno, ok := errors.As(syscall.Errno)(pe.Err); ok {
        switch errno {
        case syscall.ENOENT:
            fmt.Fprintln(os.Stderr, "ファイルが存在しない")
        case syscall.ENOTDIR:
            fmt.Fprintln(os.Stderr, "ディレクトリが存在しない")
        default:
            fmt.Fprintln(os.Stderr, "Errno =", errno)
        }
    } else {
        fmt.Fprintln(os.Stderr, "その他の PathError")
    }
}
```

という感じでハンドリングできるだろう。

ホンマこれ早めに実現しないかなぁ。

## ブックマーク

- [エラー・ハンドリングについて]({{< relref "error-handling.md" >}})
- [Error の構造化]({{< relref "error-handling2.md" >}})
- [Go 1.13 と 1.14 （Go 2 へ向けて）]({{< ref "/release/2019/06/next-steps-toward-go-2.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[総称型]: {{< relref "generics-in-go-2.md" >}} "次期 Go 言語で導入される（かもしれない）新しいエラー・ハンドリングについて予習する"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
