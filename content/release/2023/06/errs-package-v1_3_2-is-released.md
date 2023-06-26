+++
title = "goark/errs パッケージ v1.3.2 をリリースした"
date =  "2023-06-26T21:22:20+09:00"
description = "concurrency-safe なマルチエラー型を作った。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "package", "module", "error", "concurrency" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[この前書いた記事]({{< ref "/remark/2023/06/benchmark-for-errs-package.md" >}} "拙作 gorak/errs パッケージの性能評価をしてもらった")で，最後の方に「そうそう [`errors`]`.Join` 互換の関数ってあったほうがいいのかなぁ」と漠然と書いたのだが，その後，複数の goroutine を跨いでエラーハンドリングを行う必要ができたので， concurrency-safe なマルチエラー型 [`errs`]`.Errors` を作ってリリースした。

- [Release v1.3.2 · goark/errs · GitHub](https://github.com/goark/errs/releases/tag/v1.3.2)

いやぁ，並行処理書くの久しぶりで，色々とバグを潰してたら v1.3.2 まで上がってしまった（笑）

たとえば，こんな感じに使える。

```go {hl_lines=[11,18,22]}
package main

import (
    "fmt"
    "sync"

    "github.com/goark/errs"
)

func generateMultiError() error {
    errlist := &errs.Errors{}
    var wg sync.WaitGroup
    for i := 1; i <= 2; i++ {
        i := i
        wg.Add(1)
        go func() {
            defer wg.Done()
            errlist.Add(fmt.Errorf("error %d", i))
        }()
    }
    wg.Wait()
    return errlist.ErrorOrNil()
}

func main() {
    err := generateMultiError()
    fmt.Printf("%+v\n", err) // {"Type":"*errs.Errors","Errs":[{"Type":"*errors.errorString","Msg":"error 2"},{"Type":"*errors.errorString","Msg":"error 1"}]}
}
```

また [`errors`]`.Join()` 互換の関数も用意してみた。
こんな感じ。

```go {hl_lines=[13]}
package main

import (
    "errors"
    "fmt"
    "io"
    "os"

    "github.com/goark/errs"
)

func generateMultiError() error {
    return errs.Join(os.ErrInvalid, io.EOF)
}

func main() {
    err := generateMultiError()
    fmt.Printf("%+v\n", err)            // {"Type":"*errs.Errors","Errs":[{"Type":"*errors.errorString","Msg":"invalid argument"},{"Type":"*errors.errorString","Msg":"EOF"}]}
    fmt.Println(errors.Is(err, io.EOF)) // true
}
```

今回新たに作った [`errs`]`.Errors` 型は `Unwrap() []error` メソッドを持っているため， [Go] 1.20 以降であれば [`errors`]`.Is()` および [`errors`]`.As()` 関数を使って評価することができる。

これで [`errors`] 標準パッケージと互換性ありと言えるようになったかな？

## ブックマーク

- [次なる`pkg/errors`を探して - カンムテックブログ](https://tech.kanmu.co.jp/entry/2023/06/19/150000)

- [【Go 1.20 の予習】複数 error を返す Unwrap メソッドについて]({{< ref "/golang/wrapping-multiple-errors.md" >}})
- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})
- [Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang) : Zenn 本書きました

[Go]: https://go.dev/
[`errs`]: https://github.com/goark/errs "goark/errs: Error handling for Golang"
[`errors`]: https://pkg.go.dev/errors "errors · pkg.go.dev"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
