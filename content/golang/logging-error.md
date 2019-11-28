+++
title = "構造化エラーをログ出力する"
date =  "2019-11-24T14:30:43+09:00"
description = "ログ出力も JSON 形式にしたいよね。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "engineering", "programming", "logger", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.13 のリリースに合わせて [spiegel-im-spiegel/errs] パッケージを公開したのだが，このパッケージで構成した構造化エラーをログ出力することを考える。

まぁ標準の [`log`] パッケージでエラーメッセージを出力してもいいのだが，せっかく JSON 形式で出力できるようにしたんだから，ログ出力も JSON 形式にしたいよね。

ちうわけで，今回はこれを使います。
てってれー

- [rs/zerolog: Zero Allocation JSON Logger](https://github.com/rs/zerolog)

まずは準備として以下の関数を考える。

```go
import (
    "os"

    "github.com/spiegel-im-spiegel/errs"
)

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
```

ファイルをオープンするだけの関数で，オープンに失敗すると error を返す。
[spiegel-im-spiegel/errs] パッケージの使い方は

- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})

を参考にしてね。

この関数を使った `main()` 関数を書いてみよう。
まずは標準出力に対して普通に

```go
func main() {
    if err := checkFileOpen("not-exist.txt"); err != nil {
        fmt.Printf("%+v\n", err)
    }
}
```

と出力する。
これの実行結果は

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

てな感じになる。
うむうむ。

ここからが本題。

[`fmt`].`Printf()` の部分を [rs/zerolog] によるログ出力に置き換えてみよう。
とりあえず logger インスタンスの生成はこんな感じかな。

```go
logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().
    Timestamp().
    Str("role", "logger-sample").
    Logger()
```

まずは普通に error インスタンスをログ出力してみる。

```go
func main() {
    logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().
        Timestamp().
        Str("role", "logger-sample").
        Logger()

    if err := checkFileOpen("not-exist.txt"); err != nil {
        logger.Error().Err(err).Send()
    }
}
```

これの実行結果は以下の通り。

```text
$ go run sample.go
{"level":"error","role":"logger-sample","error":"file open error: open not-exist.txt: no such file or directory","time":"2009-11-10T23:00:00Z"}
```

更に [jq] コマンドを噛ませるとこんな感じになる。

```json
$ go run sample.go | jq .
{
  "level": "error",
  "role": "logger-sample",
  "error": "file open error: open not-exist.txt: no such file or directory",
  "time": "2009-11-10T23:00:00Z"
}
```

見ての通り [`zerolog`].`Event.Err()` メソッドでは単純なエラーメッセージしか出力されない（当たり前だが）。
通常の error ならこれで十分だが [`errs`].`Wrap()` 関数で生成した error では不十分である。

そこで [`zerolog`].`Event.Interface()` メソッドのほうを使ってみる。

```go {hl_lines=[8]}
func main() {
    logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().
        Timestamp().
        Str("role", "logger-sample").
        Logger()

    if err := checkFileOpen("not-exist.txt"); err != nil {
        logger.Error().Interface("error", err).Send()
    }
}
```

これでログ出力は

```json
$ go run sample.go | jq .
{
  "level": "error",
  "role": "logger-sample",
  "error": {
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
  },
  "time": "2009-11-10T23:00:00Z"
}
```

てな感じになった。
よーし，うむうむ，よーし。

[rs/zerolog] はパフォーマンスもよく使い勝手のいいパッケージで，しかも JSON 出力できるので加工しやすいだろう。
オススメである。

## ブックマーク

- [Go 1.13 のエラー・ハンドリング]({{< ref "/golang/error-handling-in-go-1_3.md" >}})
- [Log パッケージで遊ぶ]({{< relref "./logger.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`log`]: https://golang.org/pkg/log/ "log - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"
[spiegel-im-spiegel/errs]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"
[`errs`]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"
[rs/zerolog]: https://github.com/rs/zerolog "rs/zerolog: Zero Allocation JSON Logger"
[`zerolog`]: https://github.com/rs/zerolog "rs/zerolog: Zero Allocation JSON Logger"
[jq]: https://stedolan.github.io/jq/

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
