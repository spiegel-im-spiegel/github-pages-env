+++
title = "Go 1.13 のエラー・ハンドリング"
date =  "2019-08-24T12:06:25+09:00"
description = "golang.org/x/xerrors パッケージの仕様とはかなり異なっているので注意が必要である。 "
image = "/images/attention/go-logo_blue.png"
tags = [
  "golang",
  "error",
  "programming",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

さて， [Go 言語]コンパイラの [1.13 がリリース]({{< ref "/release/2019/09/go-1_13-is-released.md" >}})された。

[Go] 1.13 の変更点は色々あるが，エラー・ハンドリングに関しては「エラーの構造化」が正式に組み込まれたことだろう。
この記事では「エラーの構造化」に絞って紹介する。

## errors.Unwrap, errors.Is, errors.As 関数の追加 

[`errors`] パッケージにおいては `Unwrap()`, `Is()`, `As()` 関数が追加された。

[`errors`]`.Unwrap()` 関数はシンプルで，引数の error インスタンスが `Unwrap()` メソッドを持っていればその結果を返すというものだ。

```go
// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func Unwrap(err error) error {
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return nil
	}
	return u.Unwrap()
}
```

[golang.org/x/xerrors] パッケージでは [`xerrors`]`.Wrapper` interface 型が定義されていたが，まぁ `Unwrap()` 関数以外で `Wrapper` 型を使う局面はないので，これでもいいっちゃあいいのかな。

[`errors`]`.Is()` は2つの error インスタンスの同値性（equality）を検査する[^eq1]。
[`errors`]`.As()` 関数は error インスタンスから指定した型へ変換または抽出する。
先ほどの [`errors`]`.Unwrap()` 関数はこれらの関数内で呼び出される。

[^eq1]: 比較対象の error インスタンス（第2引数）と被検査対象の error インスタンス（第1引数）内にラッピングされている error インスタンスのいずれかが同値であるなら両インスタンは同値であると見做す。

両関数の中身は [golang.org/x/xerrors] パッケージのものと同じ（筈）。

```go
// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
func Is(err, target error) bool {
	if target == nil {
		return err == target
	}

	isComparable := reflectlite.TypeOf(target).Comparable()
	for {
		if isComparable && err == target {
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

// As finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// As will panic if target is not a non-nil pointer to either a type that implements
// error, or to any interface type. As returns false if err is nil.
func As(err error, target interface{}) bool {
	if target == nil {
		panic("errors: target cannot be nil")
	}
	val := reflectlite.ValueOf(target)
	typ := val.Type()
	if typ.Kind() != reflectlite.Ptr || val.IsNil() {
		panic("errors: target must be a non-nil pointer")
	}
	if e := typ.Elem(); e.Kind() != reflectlite.Interface && !e.Implements(errorType) {
		panic("errors: *target must be interface or implement error")
	}
	targetType := typ.Elem()
	for err != nil {
		if reflectlite.TypeOf(err).AssignableTo(targetType) {
			val.Elem().Set(reflectlite.ValueOf(err))
			return true
		}
		if x, ok := err.(interface{ As(interface{}) bool }); ok && x.As(target) {
			return true
		}
		err = Unwrap(err)
	}
	return false
}

var errorType = reflectlite.TypeOf((*error)(nil)).Elem()
```

コードが微妙にダサいのは [Go 言語]が[総称型を持っていない]({{< relref "./generics-in-go-2.md" >}} "次期 Go 言語で導入される（かもしれない）総称型について予習する")ため。
逆に言うと，総称型がなくともこの程度はできる，ということで（笑）

ここまでが準備運動。

## fmt.Errorf 関数による error のラッピング

[`fmt`]`.Errorf()` 関数の書式で `%w` が使えるようになった。
`%w` を使うことで，対応する error インスタンスをラッピングする `wrapError` 型のインスタンスを生成する。

```go
// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
//
// If the format specifier includes a %w verb with an error operand,
// the returned error will implement an Unwrap method returning the operand. It is
// invalid to include more than one %w verb or to supply it with an operand
// that does not implement the error interface. The %w verb is otherwise
// a synonym for %v.
func Errorf(format string, a ...interface{}) error {
	p := newPrinter()
	p.wrapErrs = true
	p.doPrintf(format, a)
	s := string(p.buf)
	var err error
	if p.wrappedErr == nil {
		err = errors.New(s)
	} else {
		err = &wrapError{s, p.wrappedErr}
	}
	p.free()
	return err
}
```

ちなみに `wrapError` 型は以下のように定義されている。

```go
type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}
```

シンプルで結構。

[golang.org/x/xerrors] パッケージの仕様とはかなり異なっているので注意が必要である。
開発しているシステム/アプリケーションが [`xerrors`]`.Errorf()` 関数の仕様に依存しているなら置き換えは難しいかも。

## 標準パッケージへの Unwrap() メソッドの組み込み

標準パッケージのソースコードに対して [`jvgrep`] `Unwrap src/**/*.go` とかやると分かるが，いくつかのパッケージで定義されている error 派生型にも `Unwrap()` メソッドが組み込まれているようだ。

たとえばファイル操作失敗時に吐かれる [`os`]`.PathError` 型は以下のように定義されている。

```go
// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

func (e *PathError) Unwrap() error { return e.Err }
```

これを踏まえて

```go
package main

import (
	"fmt"
	"os"
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
	if err := checkFileOpen("not-exist.txt"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
        //Outpout:
        //open not-exist.txt: no such file or directory
	}
}
```

を以下のように書き換えてみる。

{{< highlight go "hl_lines=22-26" >}}
package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
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
	if err := checkFileOpen("not-exist.txt"); err != nil {
		var errno syscall.Errno
		if errors.As(err, &errno) {
			fmt.Fprintln(os.Stderr, errno)
			return
		}
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
{{< /highlight >}}

これを実行すると

```text
$ go run sample2.go 
no such file or directory
```

と一発で [`syscall`]`.Errno` のインスタンスが抽出されていることが分かる。
これで標準パッケージのエラーの取り回しが楽になるだろう。

## そろそろ真面目にエラー・ハンドリングを設計しないと

自作ツールでもそろそろ真面目にエラー・ハンドリングを設計しないといけないかなぁ。

エラー・ハンドリングで難しいのはシステム/アプリケーションの「利用者」が欲しいエラー情報と「運用・開発者」が欲しいエラー情報とが微妙に異なる点だろう。
ビジネス用語のいわゆる 5W1H で考えるなら「利用者」が欲しいのは主に “What” と “Why” だろうが，「運用・開発者」は寧ろ残りの “When”, “Where”, “Who”, “How” の情報が重要だよね。
まぁ “When” や “Who” は logger の担当だろうけど。

幸いなことに [`fmt`] パッケージでは `%v`, `%#v`, `%+v` で情報の詳細度を変えられるので，この辺を上手く使ってどうにか，というところだろうか。

ふむむむむー。

## ブックマーク

- [Working with Errors in Go 1.13 - The Go Blog](https://blog.golang.org/go1.13-errors)

- [エラー・ハンドリングについて]({{< relref "error-handling.md" >}})
- [Error の構造化]({{< relref "error-handling2.md" >}})
- [階層化 Error パッケージ “xerrors” を試してみる]({{< relref "./xerrors.md" >}})
- [Go 1.13 と 1.14 （Go 2 へ向けて）]({{< ref "/release/2019/06/next-steps-toward-go-2.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`syscall`]: https://golang.org/pkg/syscall/ "syscall - The Go Programming Language"
[golang.org/x/xerrors]: https://godoc.org/golang.org/x/xerrors "xerrors - GoDoc"
[`xerrors`]: https://godoc.org/golang.org/x/xerrors "xerrors - GoDoc"
[`jvgrep`]: https://mattn.kaoriya.net/software/lang/go/20110819203649.htm "Big Sky :: 日本語grepが出来るjvgrepというのを作った。"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
