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

- [goark/errs: Error handling for Golang](https://github.com/goark/errs)

本パッケージは [Go 言語]によるプログラミングに於いて標準の [`errors`] パッケージを補完し，構造化されたエラーハンドリングを行うことができる。

[![check vulns](https://github.com/goark/errs/workflows/vulns/badge.svg)](https://github.com/goark/errs/actions)
[![lint status](https://github.com/goark/errs/workflows/lint/badge.svg)](https://github.com/goark/errs/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/errs/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/errs.svg)](https://github.com/goark/errs/releases/latest)

なお [`errs`] パッケージは [Go] 1.13 以上を要求する。
ご注意を。

## インポート

```go
import "github.com/goark/errs"
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

[`errs`]`.WithContext()` 関数は [`errs`]`.New()` 関数の引数として複数セット可能で[^fop1]，付加されたコンテキスト情報は `map[string]interface{}` 形式の連想配列に格納される。

[^fop1]: 可変引数に関数をセットするプログラミング・パターンは “[Functional Option Pattern]({{< ref "https://text.baldanders.info/golang/functional-options-pattern.md" >}} "インスタンスの生成と Functional Options パターン")” と呼ばれている。

更に [`errs`]`.WithCause()` 関数を使って原因エラーを付加することもできる。

```go {hl_lines=[4]}
err := errs.New(
    "file open error",
    errs.WithContext("path", path),
    errs.WithCause(cause),
)
```

[`errs`]`.WithCause()` 関数も [`errs`]`.New()` 関数の引数として複数セットできるが，最後にセットしたインスタンスのみが有効となる。
なお，複数の原因エラーがある場合は [`errors`]`.Join()` 関数を使うといいだろう（[Go] 1.20 以降）。

```go {hl_lines=[4]}
err := errs.New(
    "file open error",
    errs.WithContext("path", path),
    errs.WithCause(errors.Join(cause1, cause2)),
)
```

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
*errs.Error{Err:&errors.errorString{s:"file open error"}, Cause:&fs.PathError{Op:"open", Path:"not-exist.txt", Err:0x2}, Context:map[string]interface {}{"function":"main.checkFileOpen", "path":"not-exist.txt"}}
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
    "Type": "*fs.PathError",
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
    "Type": "*fs.PathError",
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

### その他のハンドリング関数（2023-02-07 更新）

{{< div-box type="markdown" >}}
**注意：**
[v1.2.1](https://github.com/goark/errs/releases/tag/v1.2.1 "Release v1.2.1 · goark/errs") で [`errs`]`.Cause()` 関数を `Deprecated` とした。
マルチエラーに対応できないため。

[`errs`]: https://github.com/goark/errs "goark/errs: Error handling for Golang"
{{< /div-box >}}

標準の [`errors`]`.As()`, [`errors`]`.Is()`, [`errors`]`.Unwrap()` 各関数の互換となる [`errs`]`.As()`, [`errs`]`.Is()`, [`errs`]`.Unwrap()` 関数も用意した。
まぁ，内部で [`errors`] の各関数を呼び出しているだけだけど。
でも，これで標準の [`errors`] パッケージを [`errs`] パッケージに置き換えて使うことができると思う。

さらに，複数の原因エラーを `[]error` 型で返す [`errs`]`.Unwraps()` 関数を用意した。
こんな感じに使える。

```go {hl_lines=[3]}
func main() {
    err := errors.Join(os.ErrInvalid, io.EOF)
    for _, e := range errs.Unwraps(err) {
        fmt.Println(e)
    }
    //Output:
    //invalid argument
    //EOF
}
```

なお [`errs`]`.Unwraps()` 関数は原因エラーがひとつの場合でも，要素数1の `[]error` 型で返す。

```go
func main() {
    err := errs.Wrap(os.ErrInvalid)
    for _, e := range errs.Unwraps(err) {
        fmt.Println(e)
    }
    //Output:
    //invalid argument
 }
```

さらにさらに， [`errs`]`.EncodeJSON()` 関数を使うと，通常の `error` インスタンスでも可能な限り構造を辿って JSON 形式で出力する。
たとえば

```go {hl_lines=[5]}
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        var pathError *fs.PathError
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
  "Type": "*fs.PathError",
  "Msg": "open not-exist.txt: no such file or directory",
  "Cause": {
    "Type": "syscall.Errno",
    "Msg": "no such file or directory"
  }
}

```

などと出力される。

## Zap にエラーをオブジェクトとして出力する

[Zap][`zap`] は gRPC 関連サービスや分散システムなどで人気の高い logger で，柔軟なカスタマイズができ，かつ高速で JSON 形式の構造化ログを出力できる。
この logger に拙作のパッケージを食わせてみる。
ソースコードはこんな感じ。

```go
package main

import (
    "os"

    "github.com/goark/errs"
    "go.uber.org/zap"
)

func checkFileOpen(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return errs.New(
            "file open error",
            errs.WithCause(err),
            errs.WithContext("path", path),
        )
    }
    defer file.Close()

    return nil
}

func main() {
    logger := zap.NewExample()
    defer logger.Sync()

    path := "not-exist.txt"
    if err := checkFileOpen("not-exist.txt"); err != nil {
        logger.Error("error in checkFileOpen function", zap.Error(err), zap.String("file", path))
    }
}
```

これを実行すると

```text
$ go run sample1.go | jq .
{
  "level": "error",
  "msg": "error in checkFileOpen function",
  "error": "file open error: open not-exist.txt: no such file or directory",
  "errorVerbose": "{\"Type\":\"*errs.Error\",\"Err\":{\"Type\":\"*errors.errorString\",\"Msg\":\"file open error\"},\"Context\":{\"function\":\"main.checkFileOpen\",\"path\":\"not-exist.txt\"},\"Cause\":{\"Type\":\"*fs.PathError\",\"Msg\":\"open not-exist.txt: no such file or directory\",\"Cause\":{\"Type\":\"syscall.Errno\",\"Msg\":\"no such file or directory\"}}}",
  "file": "not-exist.txt"
}
```

となる。
`"error"` 項目も `"errorVerbose"` 項目も文字列として出力されてしまうため構造化されているとは言えない。

[Zap][`zap`] には [`zap`]`.Object()` 関数があって，これを使えば内部構造を出力することができるのだが，そのためには対象のオブジェクトが [`zapcore`]`.ObjectMarshaler` 型の interface を満たす必要がある。

```go
type ObjectMarshaler interface {
    MarshalLogObject(ObjectEncoder) error
}
```

この要件を満たすために [goark/errs/zapobject][`zapobject`] モジュールを作った。
こんな感じに error をラッピングして使う。

```go {hl_lines=[21]}
package main

import (
    "os"

    "github.com/goark/errs"
    "github.com/goark/errs/zapobject"
    "go.uber.org/zap"
)

func checkFileOpen(path string) error {
    ...
}

func main() {
    logger := zap.NewExample()
    defer logger.Sync()

    path := "not-exist.txt"
    if err := checkFileOpen("not-exist.txt"); err != nil {
        logger.Error("error in checkFileOpen function", zap.Object("error", zapobject.New(err)), zap.String("file", path))
    }
}
```

これを実行すると

```text
$ go run sample2.go | jq .
{
  "level": "error",
  "msg": "error in checkFileOpen function",
  "error": {
    "type": "*errs.Error",
    "msg": "file open error: open not-exist.txt: no such file or directory",
    "error": {
      "type": "*errors.errorString",
      "msg": "file open error"
    },
    "cause": {
      "type": "*fs.PathError",
      "msg": "open not-exist.txt: no such file or directory",
      "cause": {
        "type": "syscall.Errno",
        "msg": "no such file or directory"
      }
    },
    "context": {
      "function": "main.checkFileOpen",
      "path": "not-exist.txt"
    }
  },
  "file": "not-exist.txt"
}
```

と，いい感じに構造化されて出力される。

なお [`errs`]`.Error` でラップせず通常のエラーのままでも

```go {hl_lines=[13]}
package main

import (
	"os"

	"github.com/goark/errs/zapobject"
	"go.uber.org/zap"
)

func checkFileOpen(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	path := "not-exist.txt"
	if err := checkFileOpen("not-exist.txt"); err != nil {
		logger.Error("error in checkFileOpen function", zap.Object("error", zapobject.New(err)), zap.String("file", path))
	}
}
```

```text
$ go run sample2b.go | jq .
{
  "level": "error",
  "msg": "error in checkFileOpen function",
  "error": {
    "type": "*fs.PathError",
    "msg": "open not-exist.txt: no such file or directory",
    "cause": {
      "type": "syscall.Errno",
      "msg": "no such file or directory"
    }
  },
  "file": "not-exist.txt"
}
```

という感じに可能な限り構造を辿って出力する。

## ブックマーク

- [Go 1.13 のエラー・ハンドリング]({{< ref "/golang/error-handling-in-go-1_3.md" >}})
- [書式 %v のカスタマイズ]({{< ref "/golang/formatting.md" >}})
- [構造化エラーをログ出力する]({{< ref "/golang/logging-error.md" >}})
- [Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang) : Zenn 本書きました

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"
[`errs`]: https://github.com/goark/errs "goark/errs: Error handling for Golang"
[`zapobject`]: https://pkg.go.dev/github.com/goark/errs/zapobject "zapobject package - github.com/goark/errs/zapobject - Go Packages"
[`zap`]: https://pkg.go.dev/go.uber.org/zap "zap package - go.uber.org/zap - Go Packages"
[`zapcore`]: https://pkg.go.dev/go.uber.org/zap@v1.24.0/zapcore "zapcore package - go.uber.org/zap/zapcore - Go Packages"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
