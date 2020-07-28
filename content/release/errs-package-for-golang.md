+++
title = "Go 言語用エラーハンドリング・パッケージ"
date =  "2019-09-05T23:54:45+09:00"
description = "標準の errors パッケージと組み合わせてエラーハンドリングの助けとなれば幸いである。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [spiegel-im-spiegel/errs: Error handling for Golang](https://github.com/spiegel-im-spiegel/errs)

本パッケージは [Go 言語]によるプログラミングに於いて標準の [`errors`] パッケージを補完しエラーハンドリングを行うことができる。

なお [`errs`] パッケージは [Go] 1.13 以上を要求する。
ご注意を。

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/errs.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/errs)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/errs/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/errs.svg)](https://github.com/spiegel-im-spiegel/errs/releases/latest)

## インポート

```go
import "github.com/spiegel-im-spiegel/errs"
```

## 簡単な使い方

たとえば，以下のようなファイルをオープンするだけの関数を考えてみる。

```go
func checkFileOpen(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    return nil
}
```

[`os`]`.Open()` 関数の実行時に吐き出されるエラー・インスタンス `err` を [`errs`]`.Wrap()` 関数でラッピングする。
こんな感じ。

{{< highlight go "hl_lines=4-8" >}}
func checkFileOpen(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return errs.Wrap(
            err,
            "file open error",
            errs.WithContext("path", path),
        )
    }
    defer file.Close()

    return nil
}
{{< /highlight >}}

[`errs`]`.Wrap()` 関数では元になる `error` インスタンスと追加のメッセージ，および[関数オプション]({{< ref "/golang/functional-options-pattern.md" >}})として0個以上（複数可）の [`errs`]`.WithContext()` を引数とする。

では実際に `checkFileOpen()` 関数を動かしてみよう。

```go
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        fmt.Printf("%v\n", err)
    }
}
```

この場合の実行結果は以下の通り。

```text
$ go run sample.go 
file open error: open not-exist.txt: no such file or directory
```

まぁ [Go 言語]ではありふれた出力形式だ。

ここで [`fmt`]`.Printf()` の書式を `%v` から `%#v` に変えてみる。

{{< highlight go "hl_lines=3" >}}
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        fmt.Printf("%#v\n", err)
    }
}
{{< /highlight >}}

すると実行結果は

```text
$ go run sample.go 
*errs.Error{Err:&errors.errorString{s:"file open error"}, Cause:&os.PathError{Op:"open", Path:"not-exist.txt", Err:0x2}, Context:map[string]interface {}{"function":"main.checkFileOpen", "path":"not-exist.txt"}}
```

という感じに構造体のダンプ表示ぽい出力になる。

ちなみに [`errs`]`.Error.Context` 要素は `map[string]interface{}` 型の連想配列になっているが，既定でエラーが発生した関数名を格納している。
これでエラーを追いやすくなるだろう。 

更に書式を `%+v` に変える。

{{< highlight go "hl_lines=3" >}}
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        fmt.Printf("%+v\n", err)
    }
}
{{< /highlight >}}

この場合の実行結果は

```text
$ go run sample.go 
{"Type":"*errs.Error","Err":{"Type":"*errors.errorString","Msg":"file open error"},"Context":{"function":"main.checkFileOpen","path":"not-exist.txt"},"Cause":{"Type":"*os.PathError","Msg":"open not-exist.txt: no such file or directory","Cause":{"Type":"syscall.Errno","Msg":"no such file or directory"}}}
```

と JSON フォーマットで出力される。
これなら

```json
$ go run sample.go | jq .
{
  "Type": "*errs.Error",
  "Err": {
    "Type": "*errors.errorString",
    "Msg": "file open error"
  },
  "Context": {
    "function": "main.checkFileOpen",
    "path": "not-exist.txt"
  },
  "Cause": {
    "Type": "*os.PathError",
    "Msg": "open not-exist.txt: no such file or directory",
    "Cause": {
      "Type": "syscall.Errno",
      "Msg": "no such file or directory"
    }
  }
}
```

といった感じに他ツールと組み合わせて [`errs`]`.Error` インスタンスの中身を検証することができる。

おまけの機能として [`errs`]`.Cause()` 関数も用意してみた。

```go
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        fmt.Printf("%v\n", errs.Cause(err))
    }
    // Output:
    // no such file or directory
}
```

このように [`errs`]`.Cause()` 関数では階層化 `error` を遡って大元の  `error` インスタンスを抽出することができる。

[`errs`] パッケージと標準の [`errors`] パッケージを組み合わせることでエラーハンドリングの助けとなれば幸いである。

## ブックマーク

- [Go 1.13 のエラー・ハンドリング]({{< ref "/golang/error-handling-in-go-1_3.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"
[`errs`]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
