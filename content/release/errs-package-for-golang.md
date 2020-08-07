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

本パッケージは [Go 言語]によるプログラミングに於いて標準の [`errors`] パッケージを補完し，構造化されたエラーハンドリングを行うことができる。

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/errs.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/errs)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/errs/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/errs.svg)](https://github.com/spiegel-im-spiegel/errs/releases/latest)

なお [`errs`] パッケージは [Go] 1.13 以上を要求する。
ご注意を。

## インポート

```go
import "github.com/spiegel-im-spiegel/errs"
```

## 簡単な使い方

### error インスタンスの生成

まず [`errs`]`.New()` 関数は標準の [`errors`]`.New()` を置き換え可能である。

```go
err := errs.New("file open error")
```

その上で [`errs`]`.WithContext()` 関数を使ってコンテキスト情報を付加することができる。

```go {hl_lines=[3]}
err := errs.New(
    "file open error",
    errs.WithContext("path", path),
)
```

[`errs`]`.WithContext()` 関数は引数に複数セット可能で `map[string]interface{}` 形式の連想配列に格納される。

更に [`errs`]`.WithCause()` 関数を使って原因エラーを付加することもできる。

```go {hl_lines=[4]}
err := errs.New(
    "file open error",
    errs.WithContext("path", path),
    errs.WithCause(cause),
)
```

[`errs`]`.WithCause()` 関数も引数に複数セットできるが，最後にセットしたインスタンスのみが有効となる。

以上を踏まえて，ファイルをオープンするだけの関数を考えてみよう。
こんな感じ。

```go {hl_lines=["4-8"]}
func checkFileOpen(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return errs.New(
            "file open error",
            errs.WithContext("path", path),
            errs.WithCause(err),
        )
    }
    defer file.Close()

    return nil
}
```

この `checkFileOpen()` 関数の返り値を評価する。

最初は簡単に

```go {hl_lines=[3]}
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        fmt.Printf("%v\n", err)
    }
}
```

とする。
これを実行すると

```text
$ go run sample1a.go
file open error: open not-exist.txt: no such file or directory
```

と，生成したエラーのメッセージと原因エラーのメッセージが連結されて表示される。

ここで書式 `%#v` を使ってエラー内容を表示してみる。

```go {hl_lines=[3]}
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        fmt.Printf("%#v\n", err)
    }
}
```

これを実行すると

```text
$ go run sample1b.go
*errs.Error{Err:&errors.errorString{s:"file open error"}, Cause:&os.PathError{Op:"open", Path:"not-exist.txt", Err:0x2}, Context:map[string]interface {}{"function":"main.checkFileOpen", "path":"not-exist.txt"}}
```

と，エラーの内部構造を出力してくれる。

更に 書式 `%+v` を使って

```go {hl_lines=[3]}
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        fmt.Printf("%+v\n", err)
    }
}
```

とすると，実行結果は

```json
$ go run sample1c.go | jq .
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

と JSON 形式で出力される。

この出力を見ると分かるように [`errs`] パッケージでは `error` インスタンス生成時にエラーが発生した関数をコンテキスト情報として自動的に付加している。
書式を使い分けて上手く利用して欲しい。

### error インスタンスのラッピング

[`errs`]`.Wrap()` 関数を使って，他の関数・メソッドが出力した `error` インスタンスにコンテキスト情報を付加することもできる。
先程の `checkFileOpen()` 関数であれば

```go {hl_lines=["4-7"]}
func checkFileOpen(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return errs.Wrap(
            err,
            errs.WithContext("path", path),
        )
    }
    defer file.Close()

    return nil
}
```

のように書ける。
この `checkFileOpen()` 関数の返り値を書式 `%v` で出力すると

```text
$ go run sample2a.go
open not-exist.txt: no such file or directory
```

と，元の `error` インスタンスと同じ結果になるが，書式 `%+v` で出力すると


```json
$ go run sample2c.go | jq .
{
  "Type": "*errs.Error",
  "Err": {
    "Type": "*os.PathError",
    "Msg": "open not-exist.txt: no such file or directory",
    "Cause": {
      "Type": "syscall.Errno",
      "Msg": "no such file or directory"
    }
  },
  "Context": {
    "function": "main.checkFileOpen",
    "path": "not-exist.txt"
  }
}
```

と，コンテキスト情報が付加されているのが分かる。

[`errs`]`.Wrap()` 関数にも [`errs`]`.WithCause()` 関数を使って原因エラーを付加することができる。

たとえば

```go
var ErrCheckFileOpen = errors.New("file open error")
```

などと，あらかじめ `error` インスタンス `ErrCheckFileOpen` を定義しておいて

```go {hl_lines=["5-6"]}
func checkFileOpen(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return errs.Wrap(
            ErrCheckFileOpen,
            errs.WithCause(err),
            errs.WithContext("path", path),
        )
    }
    defer file.Close()

    return nil
}
```

と `ErrCheckFileOpen` をラップし原因エラーやコンテキスト情報を付加することができる。
これを使えば

```go {hl_lines=[3]}
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        if errors.Is(err, ErrCheckFileOpen) {
            fmt.Printf("%+v\n", err)
        } else {
            fmt.Printf("Other: %v\n", err)
        }
        return
    }
}
```

と [`errors`]`.Is()` 関数等を使って比較的簡単にエラーハンドリングを行うことができる。

### その他のハンドリング関数

おまけの機能として [`errs`]`.Cause()` 関数も用意してみた。

```go {hl_lines=[3]}
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        fmt.Printf("%v\n", errs.Cause(err))
    }
    // Output:
    // no such file or directory
}
```

このように [`errs`]`.Cause()` 関数では `error` の構造を遡って大元の  `error` インスタンスを抽出することができる。

さらに，標準の [`errors`]`.As()`, [`errors`]`.Is()`, [`errors`]`.Unwrap()` 各関数の互換となる [`errs`]`.As()`, [`errs`]`.Is()`, [`errs`]`.Unwrap()` 関数も用意した。
まぁ，内部で [`errors`] の各関数を呼び出しているだけだけど。
でも，これで標準の [`errors`] パッケージを [`errs`] パッケージに置き換えて使うことができると思う。

さらにさらに，  [`errs`]`.EncodeJSON()` 関数を使うと，通常の `error` インスタンスでも可能な限り構造を辿って JSON 形式で出力する。
たとえば

```go {hl_lines=[5]}
func main() {
	if err := checkFileOpen("not-exist.txt"); err != nil {
		var pathError *os.PathError
		if errs.As(err, &pathError) {
			fmt.Printf("%v\n", errs.EncodeJSON(pathError))
		} else {
			fmt.Println(err)
		}
		return
	}
}
```

のように書けば

```json
$ go run sample/sample2d.go | jq .
{
  "Type": "*os.PathError",
  "Msg": "open not-exist.txt: no such file or directory",
  "Cause": {
    "Type": "syscall.Errno",
    "Msg": "no such file or directory"
  }
}

```

などと出力される。

## ブックマーク

- [Go 1.13 のエラー・ハンドリング]({{< ref "/golang/error-handling-in-go-1_3.md" >}})
- [書式 %v のカスタマイズ]({{< ref "/golang/formatting.md" >}})
- [構造化エラーをログ出力する]({{< ref "/golang/logging-error.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"
[`errs`]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
