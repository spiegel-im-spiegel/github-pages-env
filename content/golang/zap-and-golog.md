+++
title = "Zap と go-log を試す"
date =  "2023-05-20T18:14:08+09:00"
description = "Bluesky の公式 Go パッケージで go-log が使われてるので試してみる。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "logger" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近 Bluesky で遊んでいるのだが，これの公式 [Go] パッケージってのがあって

- [bluesky-social/indigo: Go source code for Bluesky's atproto services. NOT STABLE (yet)](https://github.com/bluesky-social/indigo)

中を見ると logger として [ipfs/go-log][`log`] パッケージを使ってるみたいなのね。
[ipfs/go-log][`log`] パッケージは [zap][`zap`] のラッパーになっていて，特にログ出力周りの取り回しが楽になるよう設計されているっぽい。

[zap][`zap`] は高速な構造化ロギングを謳っている人気のログ・パッケージである。
といっても，ベンチマークを見る限り

{{< fig-quote class="nobox" type="markdown" title="uber-go/zap: Blazing fast, structured, leveled logging in Go." link="https://github.com/uber-go/zap" lang="en" >}}
| Package | Time | Time % to zap | Objects Allocated |
| :------ | --: | -----------: | ---------------: |
| zap | 1744 ns/op | +0% | 5 allocs/op
| zap (sugared) | 2483 ns/op | +42% | 10 allocs/op
| zerolog | 918 ns/op | -47% | 1 allocs/op
| go-kit | 5590 ns/op | +221% | 57 allocs/op
| slog | 5640 ns/op | +223% | 40 allocs/op
| apex/log | 21184 ns/op | +1115% | 63 allocs/op
| logrus | 24338 ns/op | +1296% | 79 allocs/op
| log15 | 26054 ns/op | +1394% | 74 allocs/op
{{< /fig-quote >}}

スピードだけなら [rs/zerolog] のほうがだいぶ速いように見えるのだが... まぁ，でも， gRPC や分散システムなんかでは事実上の標準みたいな感じになってるし，柔軟なカスタマイズができるためクラウドのログ管理サービスとかとも相性がいいらしい。

私個人は [rs/zerolog] 推しだが [indigo] を使うなら [zap][`zap`] & [go-log][`log`] も使えるようになっておこうというわけで試してみることにした。
今回は特に拙作の [goark/errs][`errs`] との相性という観点で評価してみる。

## [Zap][`zap`] を試してみる

いきなりサンプルコードから。

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

案の定 [`zap`]`.Error()` ではエラーメッセージを吐き出すだけのようだ。
`errorVerbose` 項目は `Error()` メソッドを `%+v` 書式で出力してる感じなのだろうか。
でも，テキストとして出力してるんじゃ「構造化」とは言えない。
困ったね。

### zapcore.ObjectMarshaler

[Zap][`zap`] には [`zap`]`.Object()` 関数があって，これを使えば内部構造を出力することができるのだが，そのためには対象のオブジェクトが [`zapcore`]`.ObjectMarshaler` 型の interface を満たす必要がある。

```go
type ObjectMarshaler interface {
    MarshalLogObject(ObjectEncoder) error
}
```

一瞬 [`errs`]`.Error` 型にこのメソッドを生やすことも考えたのだが，汎用エラー構造体が特定のサードパーティ・パッケージに依存するのは面白くない。

## zapobject モジュールを作った

というわけで，新たに [goark/errs/zapobject][`zapobject`] モジュールを作った。

```go
zapobject.New(err)
```

という感じにエラーインスタンスをラップして使う。

ちなみに [errs][`errs`] と [zapobject][`zapobject`] は同一リポジトリにあるが，モジュールを分けている。
[errs][`errs`] は [zapobject][`zapobject`] に依存しないため [errs][`errs`] の依存関係を汚さずに [`zap`]`.Object()` 関数に対応できる。
ついでに言うと， [`errs`]`.Error` 型以外の error 型についても [`zapobject`]`.New()` でラップすれば `Unknown()` メソッドの挙動に従って可能な限り構造化して出力できるようにしている。

では，先程のコードを書き換えよう（一部省略している）。

```go {hl_lines=[7,21]}
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

となった。
うんうん。
ちゃんとエラー構造が出力されているね。

## [ipfs/go-log][`log`] を試してみる

[Zap][`zap`] で拙作の [`errs`]`.Error` 型の構造を出力できるようになったので，次は [ipfs/go-log][`log`] を試してみる。
こちらもいきなりサンプルコードから。

```go {hl_lines=[8,16,20]}
package main

import (
    "os"

    "github.com/goark/errs"
    "github.com/goark/errs/zapobject"
    "github.com/ipfs/go-log/v2"
)

func checkFileOpen(path string) error {
    ...
}

func main() {
    logger := log.Logger("sample")

    path := "not-exist.txt"
    if err := checkFileOpen("not-exist.txt"); err != nil {
        logger.Errorw("error in checkFileOpen function", "error", zapobject.New(err), "file", path)

    }
}
```

[`log`]`.Logger()` 関数で生成される logger の中身はは以下のようになっている。

```go
// ZapEventLogger implements the EventLogger and wraps a go-logging Logger
type ZapEventLogger struct {
    zap.SugaredLogger
    // used to fix the caller location when calling Warning and Warningf.
    skipLogger zap.SugaredLogger
    system     string
}
```

見ての通り中身は [`zap`]`.SugaredLogger` 型になっている。
このため構造化ログを出力する場合は `Errorw()` などのメソッドを使う必要がある。
 [`zap`]`.SugaredLogger` はちょっと... というのであれば `Deshugar()` すればよい。

```go
logger.Desugar().Error("error in checkFileOpen function", zap.Object("error", zapobject.New(err)), zap.String("file", path))
```

さて，これを実行してみる。

```text
$ go run sample3.go | jq .
2023-05-20T17:04:40.794+0900    ERROR    sample    sample3/sample3.go:30    error in checkFileOpen function    {"error": {"type": "*errs.Error", "msg": "file open error: open not-exist.txt: no such file or directory", "error": {"type": "*errors.errorString", "msg": "file open error"}, "cause": {"type": "*fs.PathError", "msg": "open not-exist.txt: no such file or directory", "cause": {"type": "syscall.Errno", "msg": "no such file or directory"}}, "context": {"function": "main.checkFileOpen", "path": "not-exist.txt"}}, "file": "not-exist.txt"}
```

おぅふ `orz`

[ipfs/go-log][`log`] は出力先や出力フォーマットを環境変数で制御する。
主な環境変数は以下の通り

| 環境変数 | 値 |
| --- | --- |
| `GOLOG_LOG_LEVEL` | `debug`, `info`, `warn`, `error`, `dpanic`, `panic`, `fatal` |
| `GOLOG_FILE` | ファイルに出力する場合はパスをセットする |
| `GOLOG_OUTPUT` | `stdout`, `stderr`, `file` |
| `GOLOG_LOG_FMT` | `color`, `nocolor`, `json` |

`GOLOG_LOG_LEVEL` はシステムごとに指定できる。

```bash
export GOLOG_LOG_LEVEL="error,subsystem1=info,subsystem2=debug"
```

`GOLOG_OUTPUT` は複数の出力先を指定できる。
`+` で区切って指定すればよい。

```bash
export GOLOG_FILE="/path/to/my/file.log"
export GOLOG_OUTPUT="stderr+file"
```

というわけで，環境変数を指定して改めて起動してみる。

```text
$ export GOLOG_LOG_FMT="json"

$ export GOLOG_OUTPUT="stdout"

$ go run sample3.go | jq .
{"level":"error","ts":"2023-05-20T17:35:29.669+0900","logger":"sample","caller":"sample3/sample3.go:30","msg":"error in checkFileOpen function","error":{"type":"*errs.Error","msg":"file open error: open not-exist.txt: no such file or directory","error":{"type":"*errors.errorString","msg":"file open error"},"cause":{"type":"*fs.PathError","msg":"open not-exist.txt: no such file or directory","cause":{"type":"syscall.Errno","msg":"no such file or directory"}},"context":{"function":"main.checkFileOpen","path":"not-exist.txt"}},"file":"not-exist.txt"}
{
  "level": "error",
  "ts": "2023-05-20T17:35:29.669+0900",
  "logger": "sample",
  "caller": "sample3/sample3.go:30",
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

んー。
何故か標準エラー出力にも出るなぁ。
まぁ，とりあえず標準出力に JSON 形式で出力できた。

環境変数ではなくコードで設定したい場合もあるだろう。
この場合は

```go {hl_lines=["2-7"]}
func main() {
    cfg := log.GetConfig()
    cfg.Format = log.JSONOutput
    cfg.Stderr = false
    cfg.Stdout = true
    cfg.Level = log.LevelDebug
    log.SetupLogging(cfg)
    logger := log.Logger("sample")

    path := "not-exist.txt"
    if err := checkFileOpen("not-exist.txt"); err != nil {
        logger.Desugar().Error("error in checkFileOpen function", zap.Object("error", zapobject.New(err)), zap.String("file", path))
    }
}
```

という感じに設定できる。

### "caller" は要らん

[`zap`]`.NewExample()` で logger を生成したときには気づかなかったが，実際にはログ出力に既定で `"caller"` 項目が付くんだねぇ。
ファイル名と行番号が付くのはちょっと... と思って外し方を探してみたが

```go
logger := log.Logger("sample").Desugar().WithOptions(zap.WithCaller(false))
```

という感じに `Desugar()` した上で `WithOptions()` を使って明示的に外さないといけないみたい。
んー。
そこまでするのはなぁ。
`logger` の型が変わっちゃうし。

いや

```go
logger := log.Logger("sample")
logger.SugaredLogger = *(logger.Desugar().WithOptions(zap.WithCaller(false)).Sugar())
```

って感じにすればいいのかな。
これで動かしてみよう。

```text
$ go run sample3b.go | jq .
{
  "level": "error",
  "ts": "2023-05-20T18:01:52.106+0900",
  "logger": "sample",
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

おー。
消えた消えた。

まぁ，既に [zap][`zap`] をバリバリにカスタマイズして使ってる人には [ipfs/go-log][`log`] はあまりメリットなさそうだけど，これから [zap][`zap`] を簡易に使いたいって人にはアリな選択肢かも知れない。

じゃぁ，[自作ツール](https://github.com/goark/toolbox "goark/toolbox: A collection of miscellaneous commands")の logger を載せ換えますかね。

## ブックマーク

- [GoのロギングライブラリzapのTips - Carpe Diem](https://christina04.hatenablog.com/entry/golang-zap-tips)
- [Go のロギングライブラリ zap について](https://zenn.dev/mima/articles/069b223d9b221f)
- [golangの高速な構造化ログライブラリ「zap」の使い方 - Qiita](https://qiita.com/emonuh/items/28dbee9bf2fe51d28153)

- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})

[Go]: https://go.dev/
[indigo]: https://github.com/bluesky-social/indigo "bluesky-social/indigo: Go source code for Bluesky's atproto services. NOT STABLE (yet)"
[rs/zerolog]: https://github.com/rs/zerolog "rs/zerolog: Zero Allocation JSON Logger"
[`zap`]: https://pkg.go.dev/go.uber.org/zap "zap package - go.uber.org/zap - Go Packages"
[`zapcore`]: https://pkg.go.dev/go.uber.org/zap@v1.24.0/zapcore "zapcore package - go.uber.org/zap/zapcore - Go Packages"
[`log`]: https://github.com/ipfs/go-log "ipfs/go-log: A logging library used by go-ipfs"
[`errs`]: https://github.com/goark/errs "goark/errs: Error handling for Golang"
[`zapobject`]: https://pkg.go.dev/github.com/goark/errs/zapobject "zapobject package - github.com/goark/errs/zapobject - Go Packages"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
