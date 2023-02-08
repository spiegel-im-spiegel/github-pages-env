+++
title = "goark/errs パッケージ v1.2.1 をリリースした"
date =  "2023-02-08T12:45:17+09:00"
description = "Go 1.20 でマルチエラーをサポートするようになったため，その対応"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "package", "module", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

おそらく私しか使ってないであろう，エラーハンドリング用の [`/github.com/goark/errs`][`errs`] パッケージ v1.2.1 をリリースした。

- [Release v1.2.0 · goark/errs · GitHub](https://github.com/goark/errs/releases/tag/v1.2.0)
- [Release v1.2.1 · goark/errs · GitHub](https://github.com/goark/errs/releases/tag/v1.2.1)

実は v1.2.0 でチョンボして v1.2.1 として出し直している。
ごめんペコン。

[Go] 1.20 で `Unwrap() []error` メソッドを持つマルチエラーをサポートするようになったため，その対応。
たとえば，複数の原因エラーがある場合に，標準の [`errors`]`.Join()` 関数などと組み合わせて

```go {hl_lines=[13,18]}
package main

import (
    "errors"
    "fmt"
    "io"
    "os"

    "github.com/goark/errs"
)

func generateMultiError() error {
    return errs.New("error with multiple causes", errs.WithCause(errors.Join(os.ErrInvalid, io.EOF)))
}

func main() {
    err := generateMultiError()
    fmt.Println(errors.Is(err, io.EOF)) // true
}
```

てな感じに書くことができる。
ちなみに

```go {hl_lines=[3]}
func main() {
    err := generateMultiError()
    fmt.Printf("%+v\n",  err)
}
```

と書き換えて実行すると

```text
$ go run sample4.go | jq .
{
  "Type": "*errs.Error",
  "Err": {
    "Type": "*errors.errorString",
    "Msg": "error with multiple causes"
  },
  "Context": {
    "function": "main.generateMultiError"
  },
  "Cause": {
    "Type": "*errors.joinError",
    "Msg": "invalid argument\nEOF",
    "Cause": [
      {
        "Type": "*errors.errorString",
        "Msg": "invalid argument"
      },
      {
        "Type": "*errors.errorString",
        "Msg": "EOF"
      }
    ]
  }
}
```

マルチエラー部分もいい感じに JSON 出力できる。

以前の [`errs`]`.Cause()` 関数には `Deprecated` マークを付けた（マルチエラーに対応できないため）。

```go
// Deprecated: should not be used
func Cause(err error) error {
    for err != nil {
        unwraped := errors.Unwrap(err)
        if unwraped == nil {
            return err
        }
        err = unwraped
    }
    return err
}
```

`Deprecated` マークを付けた型や関数を使用すると lint などで注意喚起してくれる。
便利。

[`errs`]`.Cause()` 関数の代わりと言ってはナニだが [`errs`]`.Unwraps()` 関数を用意した。

```go
func Unwraps(err error) []error {
    if err != nil {
        if es, ok := err.(interface {
            Unwrap() []error
        }); ok {
            return es.Unwrap()
        }
    }
    e := errors.Unwrap(err)
    if e != nil {
        return []error{e}
    }
    return nil
}
```

この関数を使えば原因エラーを `[]error` で返してくれる（原因エラーが単一の場合もスライスで返すので注意）。

マルチエラーを格納する独自の型を作ることも一瞬考えたが，標準の [`errors`]`.Join()` 関数などで作れば必要十分だと思うし，足りなければ改めて考えるか。
ということで今回はここまで。

## ブックマーク

- [【Go 1.20 の予習】複数 error を返す Unwrap メソッドについて]({{< ref "/golang/wrapping-multiple-errors.md" >}})
- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})
- [Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang) : Zenn 本書きました

[Go]: https://go.dev/
[`errs`]: https://github.com/goark/errs "goark/errs: Error handling for Golang"
[`errors`]: https://pkg.go.dev/errors "errors · pkg.go.dev"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
