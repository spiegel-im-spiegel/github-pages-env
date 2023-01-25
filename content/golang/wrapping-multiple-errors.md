+++
title = "【Go 1.20 の予習】複数 error を返す Unwrap メソッドについて"
date =  "2023-01-25T12:57:01+09:00"
description = "Go 1.20 で errors パッケージの仕様が変わるみたいなので予習しておく。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

おそらく 2022-02 にリリースされる [Go] 1.20 で errors パッケージの仕様が変わるみたいなので予習しておく。

{{< fig-quote type="markdown" title="Go 1.20 Release Notes - The Go Programming Language" link="https://tip.golang.org/doc/go1.20#errors" lang="en" >}}
**Wrapping multiple errors**

Go 1.20 expands support for error wrapping to permit an error to wrap multiple other errors.

An error e can wrap more than one error by providing an `Unwrap` method that returns a `[]error`.

The [`errors.Is`](https://tip.golang.org/pkg/errors/#Is) and [`errors.As`](https://tip.golang.org/pkg/errors/#As) functions have been updated to inspect multiply wrapped errors.

The [`fmt.Errorf`](https://tip.golang.org/pkg/fmt/#Errorf) function now supports multiple occurrences of the `%w` format verb, which will cause it to return an error that wraps all of those error operands.

The new function [`errors.Join`](https://tip.golang.org/pkg/errors/#Join) returns an error wrapping a list of errors. 
{{< /fig-quote >}}

現行の [`errors`]`.Is()` および [`errors`]`.As()` 各関数では 対象となる error インスタンスについて型アサーションを行い `Unwrap() error` メソッドを含む型か否かで再帰的に処理を行っているが， [Go] 1.20 からは，この評価に `Unwrap() []error` メソッドが加わる。

具体的にはこんな感じらしい。

```go { hl_lines=["37-43", "97-103"]}
// Is reports whether any error in err's tree matches target.
//
// The tree consists of err itself, followed by the errors obtained by repeatedly
// calling Unwrap. When err wraps multiple errors, Is examines err followed by a
// depth-first traversal of its children.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
//
// An error type might provide an Is method so it can be treated as equivalent
// to an existing error. For example, if MyError defines
//
//	func (m MyError) Is(target error) bool { return target == fs.ErrExist }
//
// then Is(MyError{}, fs.ErrExist) returns true. See syscall.Errno.Is for
// an example in the standard library. An Is method should only shallowly
// compare err and the target and not call Unwrap on either.
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
		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return false
			}
		case interface{ Unwrap() []error }:
			for _, err := range x.Unwrap() {
				if Is(err, target) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}

// As finds the first error in err's tree that matches target, and if one is found, sets
// target to that error value and returns true. Otherwise, it returns false.
//
// The tree consists of err itself, followed by the errors obtained by repeatedly
// calling Unwrap. When err wraps multiple errors, As examines err followed by a
// depth-first traversal of its children.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// An error type might provide an As method so it can be treated as if it were a
// different error type.
//
// As panics if target is not a non-nil pointer to either a type that implements
// error, or to any interface type.
func As(err error, target any) bool {
	if err == nil {
		return false
	}
	if target == nil {
		panic("errors: target cannot be nil")
	}
	val := reflectlite.ValueOf(target)
	typ := val.Type()
	if typ.Kind() != reflectlite.Ptr || val.IsNil() {
		panic("errors: target must be a non-nil pointer")
	}
	targetType := typ.Elem()
	if targetType.Kind() != reflectlite.Interface && !targetType.Implements(errorType) {
		panic("errors: *target must be interface or implement error")
	}
	for {
		if reflectlite.TypeOf(err).AssignableTo(targetType) {
			val.Elem().Set(reflectlite.ValueOf(err))
			return true
		}
		if x, ok := err.(interface{ As(any) bool }); ok && x.As(target) {
			return true
		}
		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return false
			}
		case interface{ Unwrap() []error }:
			for _, err := range x.Unwrap() {
				if As(err, target) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}
```

ちょっと長くて申し訳ないが，各関数の型 switch 文のなかで `Unwrap() []error` メソッドを含む型を評価しているのがおわかりだろうか。

標準パッケージでは新設の [`errors`]`.Join()` 関数や [`fmt`]`.Errorf()` 関数の拡張でマルチエラーに対応するようだが，自前で error 型を作る場合でも `Unwrap() []error` メソッドを追加することで [`errors`]`.Is()` 関数や [`errors`]`.As()` 関数による評価が可能になるわけだ。

## ブックマーク

- [Go 1.20 Wrapping multiple errors | フューチャー技術ブログ](https://future-architect.github.io/articles/20230126a/)

- [Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang)

[Go]: https://go.dev/
[`errors`]: https://pkg.go.dev/errors "errors package - errors - Go Packages"
[`fmt`]: https://pkg.go.dev/fmt "fmt package - errors - Go Packages"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
