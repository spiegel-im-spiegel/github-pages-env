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

[`errs`]`.Wrap()` 関数では元になる `error` インスタンスと追加のメッセージ，および[関数オプション]({{< ref "/golang/functional-options-pattern.md" >}})として [`errs`]`.WithContext()` を引数とする。

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
&errs.Error{Msg:"file open error", Context:map[string]interface {}{"function":"main.checkFileOpen", "path":"not-exist.txt"}, Cause:&os.PathError{Op:"open", Path:"not-exist.txt", Err:0x2}}
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
{"Type":"*errs.Error","Msg":"file open error: open not-exist.txt: no such file or directory","Context":{"function":"main.checkFileOpen","path":"not-exist.txt"},"Cause":{"Type":"*os.PathError","Msg":"open not-exist.txt: no such file or directory","Cause":{"Type":"syscall.Errno","Msg":"no such file or directory"}}}
```

と JSON フォーマットで出力される。
これなら

```json
$ go run sample.go | jq .
{
  "Type": "*errs.Error",
  "Msg": "file open error: open not-exist.txt: no such file or directory",
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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
