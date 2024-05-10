+++
title = "Log パッケージで遊ぶ"
date = "2018-02-27T01:32:43+09:00"
description = "log パッケージの欠点は出力のフィルタリングができないことである。せっかく「できる子」なんだから，なるべく生かした形でカスタマイズしてみる。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "engineering", "programming", "logger" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は [Go 言語]の標準パッケージに含まれる [`log`] について。

[`log`] パッケージは，名前の通り，ログ出力を行うパッケージである。
例えば以下のように書ける。

```go
package main

import (
    "log"
)

func main() {
    log.Println("Hello, World")
}
```

これを実行すると以下のフォーマットで標準エラー出力に出力される。

```text
2009/11/10 23:00:00 Hello, World
```

## Log パッケージの中身

具体的に [`log`] パッケージを見てみよう。

```go
// A Logger represents an active logging object that generates lines of
// output to an io.Writer. Each logging operation makes a single call to
// the Writer's Write method. A Logger can be used simultaneously from
// multiple goroutines; it guarantees to serialize access to the Writer.
type Logger struct {
    mu     sync.Mutex // ensures atomic writes; protects the following fields
    prefix string     // prefix to write at beginning of each line
    flag   int        // properties
    out    io.Writer  // destination for output
    buf    []byte     // for accumulating text to write
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
    l.Output(2, fmt.Sprintf(format, v...))
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is used to recover the PC and is
// provided for generality, although at the moment on all pre-defined
// paths it will be 2.
func (l *Logger) Output(calldepth int, s string) error {
    now := time.Now() // get this early.
    var file string
    var line int
    l.mu.Lock()
    defer l.mu.Unlock()
    if l.flag&(Lshortfile|Llongfile) != 0 {
        // Release lock while getting caller info - it's expensive.
        l.mu.Unlock()
        var ok bool
        _, file, line, ok = runtime.Caller(calldepth)
        if !ok {
            file = "???"
            line = 0
        }
        l.mu.Lock()
    }
    l.buf = l.buf[:0]
    l.formatHeader(&l.buf, now, file, line)
    l.buf = append(l.buf, s...)
    if len(s) == 0 || s[len(s)-1] != '\n' {
        l.buf = append(l.buf, '\n')
    }
    _, err := l.out.Write(l.buf)
    return err
}
```

ログ出力時は [`log`]`.Logger` 内部のバッファを使って整形するため，複数の [goroutine] から非同期に呼ばれる可能性を考慮し， [`sync`]`.Mutex` を使って排他的に処理している。

## Log のカスタマイズ

[`log`] パッケージではログの出力先や出力フォーマットをある程度カスタマイズできる。
たとえば，こんな感じ。

```go
package main

import (
    "log"
    "os"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
    log.SetOutput(os.Stdout)
    log.SetPrefix("[Hello] ")

    log.Println("Hello, World")
}
```

この場合，以下の形式で標準出力に出力される。

```text
[Hello] 2009/11/10 23:00:00.000000 main.go:13: Hello, World
```

また [`log`]`.New()` 関数で新たな [`log`]`.Logger` インスタンスを作成することもできる。

```go
package main

import (
    "log"
    "os"
)

func main() {
    l := log.New(os.Stdout, "[Hello] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

    l.Println("Hello, World")
}
```

このコードのログ出力は先ほどと同じになる。

## 出力レベルとフィルタリング

[`log`] パッケージの欠点は出力のフィルタリングができないことである。
ログメッセージ毎に ERROR, WARN, INFO, DEBUG といった出力レベルの設定を行い，あらかじめ指定したレベル以下のメッセージについてはフィルタリングできる仕掛けが欲しいところである。

サードパーティーではフィルタリング機能を持つ logger が色々あるが，せっかく [`log`] パッケージが「できる子」なんだから，なるべく生かした形で実装を考えてみる。

### フィルタリング付き Writer

まずは，出力時にフィルタリングを行う Writer を考えてみる。
有り難いことに既にこの機能を持ったパッケージが存在する。

- [hashicorp/logutils: Utilities for slightly better logging in Go (Golang).](https://github.com/hashicorp/logutils)

これは [`log`]`.Logger` インスタンスにセットする Writer をラップしてフィルタリング機能を付加するものである。
こんな感じで使うらしい。

```go
package main

import (
    "log"
    "os"

    "github.com/hashicorp/logutils"
)

func main() {
    filter := &logutils.LevelFilter{
        Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
        MinLevel: logutils.LogLevel("WARN"),
        Writer:   os.Stderr,
    }
    log.SetOutput(filter)

    log.Print("[DEBUG] Debugging")         // this will not print
    log.Print("[WARN] Warning")            // this will
    log.Print("[ERROR] Erroring")          // and so will this
    log.Print("Message I haven't updated") // and so will this
}
```

これを実行するとこんな感じになる。

```text
2009/11/10 23:00:00 [WARN] Warning
2009/11/10 23:00:00 [ERROR] Erroring
2009/11/10 23:00:00 Message I haven't updated
```

あらかじめキーワード（`"DEBUG"` など）を決めておいて，ログ出力する文字列にこのキーワードを埋め込むことで出力レベルの設定とフィルタリングを行うようだ。
実際には埋め込むキーワードの typo を防ぐために何らかの仕掛けを作る必要があると思うが，とてもシンプルな作りなので，気軽に導入できそうである。

### Logger に渡す前にフィルタリングする

[hashicorp/logutils] は Writer への書き込み時にフィルタリングを行うものだったが，別のアプローチで [`log`]`.Logger` に渡す前にフィルタリングすることを考えてみる。
というわけで作ってみた。

- [spiegel-im-spiegel/logf: Simple logging package by Golang](https://github.com/spiegel-im-spiegel/logf)

[`logf`]`.Logger` はこんな感じで [`log`]`.Logger` インスタンス（へのポインタ）を内部に持つ。

```go
//Logger is logger class
type Logger struct {
    lg   *log.Logger // logger
    mu   sync.Mutex  // ensures atomic writes; protects the following fields
    flag int         // properties
    min  Level
}
```

実際の出力はこんな感じにしてみた。

```go
//Output writes the output for a logging event.
func (l *Logger) Output(lv Level, calldepth int, s string) error {
    if lv >= l.min {
        if (l.flag & Llevel) != 0 {
            return l.lg.Output(calldepth, fmt.Sprintf("[%v] %s", lv, s))
        }
        return l.lg.Output(calldepth, s)
    }
    return nil
}
```

出力レベルのチェック時に排他処理を行っていないが，値の参照だけで状態を変えるわけではないので，まぁいいかなと。

これを使ってログ出力のコードを書いてみる。

```go
package main

import (
    "os"

    "github.com/spiegel-im-spiegel/logf"
)

func main() {
    logf.SetOutput(os.Stdout)
    logf.SetMinLevel(logf.WARN)

    logf.Debug("Debugging")   // this will not print
    logf.Print("Information") // this will not print
    logf.Warn("Warning")      // this will
    logf.Error("Erroring")    // and so will this
}
```

これを実行するとこんな感じになる。

```text
2009/11/10 23:00:00 [WARN] Warning
2009/11/10 23:00:00 [ERROR] Erroring
```

んー。
杜撰なコードだが，取り敢えずこんな感じかな。

## 【追記】ログファイルの分割とローテーション

もうひとつ [`log`] パッケージにない機能としてログファイルの分割とローテーション機能がある。
たとえば [Apache HTTP Server](http://httpd.apache.org/ "Welcome! - The Apache HTTP Server Project") の [rotatelogs] のような機能があると便利である。

と思って探したらありましたよ。

- [lestrrat-go/file-rotatelogs: Port of perl5 File::RotateLogs to Go](https://github.com/lestrrat-go/file-rotatelogs) [^rl1]

[^rl1]: [lestrrat-go/file-rotatelogs] パッケージは Perl の [File::RotateLogs](https://metacpan.org/release/File-RotateLogs "File-RotateLogs-0.08 - File logger supports log rotation - metacpan.org") から [Go 言語]へのポーティングのようだ。

[`rotatelogs`] パッケージと先ほどの [`logf`] パッケージを組み合わせれば，こんな感じに書ける。

```go
package main

import (
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    "github.com/spiegel-im-spiegel/logf"
)

func main() {
    logf.SetMinLevel(logf.WARN)
    rl, err := rotatelogs.New("./log.%Y%m%d%H%M.txt")
    if err != nil {
        logf.Fatal(err)
        return
    }
    logf.SetOutput(rl)

    logf.Debug("Debugging")   // this will not print
    logf.Print("Information") // this will not print
    logf.Warn("Warning")      // this will
    logf.Error("Erroring")    // and so will this
}
```

よーし，うむうむ，よーし。

## ブックマーク

### 構造化ログ

- [rs/zerolog: Zero Allocation JSON Logger](https://github.com/rs/zerolog) : JSON 形式でログを吐く。おススメ
- [hnakamur/ltsvlog: a minimalist LTSV logging library in Go](https://github.com/hnakamur/ltsvlog) : LTSV 形式でログを吐く
    - [GoでLTSV形式でログ出力するライブラリを書いた](https://hnakamur.github.io/blog/2016/06/13/wrote_go_ltsvlog_library/)
    - [zerologを参考にしてltsvlogを改良してみた](https://hnakamur.github.io/blog/2017/05/28/improve-ltsvlog-with-referring-to-zerolog/)

### その他

- [golangでlogを標準出力とテキストファイルの2箇所の出力する - Qiita](http://qiita.com/74th/items/441ffcab80a6a28f7ee3)
- [Goのバッチで統計を取得するAPIを用意しておくと便利 - Qiita](http://qiita.com/sudix/items/c542e1b59bc94dc741e3)
- [ええっ！？　文字列で書くの！？　ログレベル付きロガーhashicorp/logutilsのご紹介 - Qiita](https://qiita.com/mackee_w/items/3c0940733b6c0922554c)
- [標準出力に表示をしながらファイルにも保存して、かつローテーションもする - Qiita](https://qiita.com/saka1_p/items/5e37fafb35b10bb3d5a3)
- [GitHub - spf13/jWalterWeatherman: So you always leave a note](https://github.com/spf13/jwalterweatherman)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`log`]: https://golang.org/pkg/log/ "log - The Go Programming Language"
[`sync`]: https://golang.org/pkg/sync/ "sync - The Go Programming Language"
[hashicorp/logutils]: https://github.com/hashicorp/logutils "hashicorp/logutils: Utilities for slightly better logging in Go (Golang)."
[`logutils`]: https://github.com/hashicorp/logutils "hashicorp/logutils: Utilities for slightly better logging in Go (Golang)."
[`logf`]: https://github.com/spiegel-im-spiegel/logf "spiegel-im-spiegel/logf: Simple logging package by Golang"
[goroutine]: http://golang.org/ref/spec#Go_statements "The Go Programming Language Specification - The Go Programming Language"
[rotatelogs]: https://httpd.apache.org/docs/2.4/programs/rotatelogs.html "rotatelogs - Piped logging program to rotate Apache logs - Apache HTTP Server Version 2.4"
[lestrrat-go/file-rotatelogs]: https://github.com/lestrrat-go/file-rotatelogs "lestrrat-go/file-rotatelogs: Port of perl5 File::RotateLogs to Go"
[`rotatelogs`]: https://github.com/lestrrat-go/file-rotatelogs "lestrrat-go/file-rotatelogs: Port of perl5 File::RotateLogs to Go"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
