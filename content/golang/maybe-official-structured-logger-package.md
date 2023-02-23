+++
title = "公式の構造化 Logger (になるかもしれない) slog パッケージ"
date =  "2022-11-12T21:26:57+09:00"
description = "多少は期待してもいいかもしれない。公式になるといいな。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "logger" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今日の「[第4回『Go言語による分散サービス』オンライン読書会](https://technical-book-reading-2.connpass.com/event/262819/)」で構造化ログについて話題になったのだが（ちなみに『[Go言語による分散サービス]』のサンプルコードでは構造化ロガーとして [`zap`] を利用している），その中で公式の構造化 logger がプロポーザルとして上がってるという話を聞く。
どうもこれのことらしい。

- [proposal: log/slog: structured, leveled logging · Issue #56345 · golang/go · GitHub](https://github.com/golang/go/issues/56345)
- [Proposal: Structured Logging](https://go.googlesource.com/proposal/+/master/design/56345-structured-logging.md)

この [`slog`] パッケージについて[解説記事](https://www.sobyte.net/post/2022-10/go-slog/ "slog: Golang's official structured logging package - SoByte")も見つけたので，これも参考にしながら少し遊んでみようと思う。

2022年10月末時点で，このパッケージについては

{{< fig-quote type="markdown" title="slog: Golang's official structured logging package - SoByte" link="https://www.sobyte.net/post/2022-10/go-slog/" lang="en" >}}
In late August 2022, [Jonathan Amsterdam of the Go team initiated a discussion](https://github.com/jba) with the community about [adding structured, log-level support for logging packages](https://github.com/golang/go/discussions/54763) to the Go standard library, and formed a consensus proposal.

Jonathan Amsterdam has named the structured logging package slog and plans to put it under `log/slog`. He also gave the [initial implementation of slog](https://github.com/golang/exp/slog) under `golang.org/x/exp`, and the Proposal is [officially in review these days](https://go-review.googlesource.com/c/proposal/+/444415/3/design/56345-structured-logging.md). It is not known when it will be officially implemented in the official version of Go.
{{< /fig-quote >}}

という感じらしい。
実際に [design doc](https://github.com/golang/proposal#design-documents) が公開されレビューが行われているということなら多少は期待してもいいかもしれない。
公式になるといいな。

試しに簡単なコードを書いてみよう。

```go
package main

import (
    "os"

    "golang.org/x/exp/slog"
)

func main() {
    slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
    slog.Info("hello world")
}
```

これを実行すると

```text
$ go run sample1.go 
time=2022-11-12T19:00:27.704+09:00 level=INFO msg="hello world"
```

という感じに `KEY=VALUE` スタイルのログを出力する。
またバックエンド側のハンドラを

```go { hl_lines=[10]}
package main

import (
    "os"

    "golang.org/x/exp/slog"
)

func main() {
    slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr)))
    slog.Info("hello world")
}
```

と差し替えることで

```text
$ go run sample2.go 
{"time":"2022-11-12T19:45:51.816529605+09:00","level":"INFO","msg":"hello world"}
```

と JSON スタイルの出力に切り替えることもできる。
この出力ハンドラは自作することもでき，たとえば [`zap`] や [`zerolog`] といったサードパーティの logger に接続することもできるらしい。

{{< fig-img-quote src="./go-slog.png" title="slog: Golang's official structured logging package - SoByte" link="https://www.sobyte.net/post/2022-10/go-slog/" width="1931" lang="en" >}}

もう少し遊んでみよう。
`error` 値を出力するコードを書いてみる。

```go
package main

import (
    "context"
    "os"

    "github.com/goark/errs"
    "golang.org/x/exp/slog"
)

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

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stderr)).WithContext(context.TODO())
    logger.Enabled(slog.DebugLevel)

    if err := checkFileOpen("not-exist.txt"); err != nil {
        logger.Error("open error", err)
    } else {
        logger.Info("no error")
    }
}
```

これを実行すると

```text
$ go run sample3.go 
{"time":"2022-11-12T20:32:38.775783665+09:00","level":"ERROR","msg":"open error","err":"open not-exist.txt: no such file or directory"}
```

となる。
まぁ，予想通り。

でも，拙作の [`errs`] パッケージは [`errs`]`.Error.MarshalJSON()` メソッドを持っているので JSON 形式で出力して欲しい。
苦肉の策で

```go
logger.Error("open error", err)
```

を

```go
logger.Error("open error", err, slog.Any("info", err))
```

としてみたが，結果は

```text
$ go run sample4.go 
{"time":"2022-11-12T20:54:04.358245853+09:00","level":"ERROR","msg":"open error","info":"open not-exist.txt: no such file or directory","err":"open not-exist.txt: no such file or directory"}
```

となる。
ソースコードを見ると `fmt.Sprint(v.any)` なる記述が見れるので error 型は問答無用で `Error()` メソッドが出力する文字列に変換されてしまうようだ。
ふむむー。
ひょっとしたら [`slog`]`.JSONHandler` を参考に独自のハンドラを組めばどうにかなるかもしれんが，今回は止めておこう。
ホンマに公式になるかどうか分からんし（笑）

ともかく，公式の構造化 logger が登場すればログ周りのシーンはかなり変わるだろう。
楽しみなことである。

## ブックマーク

- [Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang)
- [Go公式の構造化ロガー（として提案されている）slogを触ってみたメモ](https://zenn.dev/mizutani/articles/golang-exp-slog)
- [GitHub - chanchal1987/zaphandler: zaphandler will help to create slog handler using zap logger](https://github.com/chanchal1987/zaphandler) : slog 用 zap 出力ハンドラ
- [Playing With Slog, the Proposed Structured Logging Package for the Go Standard Library](https://josephwoodward.co.uk/2022/11/slog-structured-logging-proposal)
- [Goの新しい構造化ロガーを体験しよう | gihyo.jp](https://gihyo.jp/article/2023/02/tukinami-go-04)

[Go]: https://go.dev/
[`slog`]: https://pkg.go.dev/golang.org/x/exp/slog "slog package - golang.org/x/exp/slog - Go Packages"
[`zap`]: https://github.com/uber-go/zap "uber-go/zap: Blazing fast, structured, leveled logging in Go."
[`zerolog`]: https://github.com/rs/zerolog "rs/zerolog: Zero Allocation JSON Logger"
[`errs`]: https://github.com/goark/errs "goark/errs: Error handling for Golang"
[Go言語による分散サービス]: https://www.amazon.co.jp/dp/4873119979?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Go言語による分散サービス ―信頼性、拡張性、保守性の高いシステムの構築 | Travis Jeffery, 柴田 芳樹 |本 | 通販 | Amazon"

## 参考図書

{{% review-paapi "4873119979" %}} <!-- Go言語による分散サービス -->
{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
