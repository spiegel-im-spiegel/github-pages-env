+++
title = "Go 言語用 CLI プログラミング支援パッケージ"
date = "2019-09-07T11:55:58+09:00"
description = "このパッケージをそのまま使うことは想定しておらず，何らかのアレンジを加えた上で，それぞれの CLI ツール用に組み込むことを念頭に置いている。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "cli", "filepath" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface](https://github.com/spiegel-im-spiegel/gocli)

本パッケージ [`gocli`] は [Go 言語] で CLI (Command-Line Interface) を構成する際に必要になるであろう細々とした機能をまとめたライブラリである。
ただし，このパッケージをそのまま使うことは想定しておらず（そのまま使ってもいいけど）何らかのアレンジを加えた上で，それぞれの CLI ツール用に組み込むことを念頭に置いている。

このため [`gocli`] では [Go] コンパイラが提供する標準パッケージ以外の外部パッケージはなるべく使わないようにし，ライセンスも，あらゆる権利を放棄した [CC0] を設定している。

なお [`gocli`] パッケージは [Go] 1.13 以上を要求する。
ご注意を。

[![check vulns](https://github.com/spiegel-im-spiegel/gocli/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/gocli/actions)
[![lint status](https://github.com/spiegel-im-spiegel/gocli/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/gocli/actions)
[![GitHub license](https://img.shields.io/badge/license-CC0-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/gocli/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/gocli.svg)](https://github.com/spiegel-im-spiegel/gocli/releases/latest)

## 標準入出力と終了コード

[`gocli`]`/rwi` パッケージは標準入出力をコンテキスト情報として格納する構造体を提供する。
また [`gocli`]`/exitcode` パッケージは CLI 終了時の終了コードを定義する。

両者は以下のように使う。

```go
package main

import (
    "os"

    "github.com/spiegel-im-spiegel/gocli/exitcode"
    "github.com/spiegel-im-spiegel/gocli/rwi"
)

func run(ui *rwi.RWI) exitcode.ExitCode {
    ui.Outputln("Hello world")
    return exitcode.Normal
}

func main() {
    run(rwi.New(
        rwi.WithReader(os.Stdin),
        rwi.WithWriter(os.Stdout),
        rwi.WithErrorWriter(os.Stderr),
    )).Exit()
}
```

[`gocli`]`/rwi` パッケージを使うメリットはテストで発揮される。
たとえば上述の `run()` 関数をテストするのであれば

```go
outBuf := new(bytes.Buffer)
outErrBuf := new(bytes.Buffer)
code := run(rwi.New(
    rwi.WithWriter(outBuf),
    rwi.WithErrorWriter(outErrBuf),
))
```

として実行結果を `code`, `outBuf` および `outErrBuf` から取り出し評価することができる。

## SIGNAL をハンドリングする

[`gocli`]`/signal` パッケージは標準の [`context`] パッケージと組み合わせて SIGNAL のハンドリングを行う。
たとえば，こんな感じ

{{< highlight go "hl_lines=33" >}}
package main

import (
    "context"
    "fmt"
    "os"
    "time"

    "github.com/spiegel-im-spiegel/gocli/signal"
)

func ticker(ctx context.Context) error {
    t := time.NewTicker(1 * time.Second) // 1 second cycle
    defer t.Stop()

    for {
        select {
        case now := <-t.C: // ticker event
            fmt.Println(now.Format(time.RFC3339))
        case <-ctx.Done(): // cancel event from context
            fmt.Println("Stop ticker")
            return ctx.Err()
        }
    }
}

func Run() error {
    errCh := make(chan error, 1)
    defer close(errCh)

    go func() {
        child, cancelChild := context.WithTimeout(
            signal.Context(context.Background(), os.Interrupt), // cancel event by SIGNAL
            10*time.Second, // timeout after 10 seconds
        )
        defer cancelChild()
        errCh <- ticker(child)
    }()

    err := <-errCh
    fmt.Println("Done")
    return err
}

func main() {
    if err := Run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
{{< /highlight >}}

このコードでは `signal.Context()` 関数で指定した SIGNAL 用の [`context`]`.Context` インスタンスを生成している。
SIGNAL または親 [`context`]`.Context` インスタンスによるキャンセルイベントを受信した場合は，子 [`context`]`.Context` インスタンスにキャンセルが伝搬する。

[`context`] パッケージを使ったキャンセルの伝搬については以下を参照のこと。

- [time.Ticker で遊ぶ]({{< ref "/golang/ticker.md" >}})

## ワイルドカードを含むファイルの検索

[`gocli`]`/file` パッケージを使ったファイル検索は標準の [`filepath`]`.Glob()` 関数を拡張する形で実装している。
こんな感じに使える。

```go
package main

import (
	"fmt"

	"github.com/spiegel-im-spiegel/gocli/file"
)

func main() {
	result := file.Glob("**/*.[ch]", nil)
	fmt.Println(result)
	// Output:
	// [testdata/include/source.h testdata/source.c]
}
```

`file.Glob()` 関数の第2引数には検索時の条件を設定できる。
こんな感じ。

```go
package main

import (
	"fmt"

	"github.com/spiegel-im-spiegel/gocli/file"
)

func main() {
	result := file.Glob(
		"**/*.[ch]",
		file.NewGlobOption(file.WithFlags(file.GlobStdFlags|file.GlobAbsolutePath)))
	fmt.Println(result)
	// Output:
	// [/home/username/work/gocli/file/testdata/include/source.h /home/username/work/gocli/file/testdata/source.c]
}
```

指定できるフラグは以下の通り。

```go
//Operation flag in Glob() function.
const (
	GlobContainsFile GlobFlag = 1 << iota
	GlobContainsDir
	GlobSeparatorSlash
	GlobAbsolutePath
	GlobStdFlags = GlobContainsFile | GlobContainsDir
)
```

`file.Glob()` 関数の第2引数に `nil` をセットするか `file.NewGlobOption()` を引数なしで呼び出した場合は `file.GlobStdFlags` のみがセットされる。

[`gocli`]`/file` パッケージはファイル操作の練習用に作ったもので，それなりには使えるとは思うが，正直に言って素朴すぎて効率はよくない。
実際に使うにはもう少しアレンジが必要になるだろう。

## 設定ファイルのパスを取得する

[Go] 1.13 から [`os`]`.UserConfigDir()` 関数が追加されたので，これを使って設定ファイルのパスを取得するパッケージ [`gocli`]`/config` を作ってみた。
こんな感じで使う。

```go
package main

import (
	"fmt"

	"github.com/spiegel-im-spiegel/gocli/config"
)

func main() {
	path := config.Path("app", "config.json")
	fmt.Println(path)
	// Output:
	// /home/username/.config/app/config.json
}
```

アプリケーション名を指定して設定用ディレクトリのパスを取得する `config.Dir(appName string)` 関数も用意した。

[`os`]`.UserConfigDir()` 関数で取得したパスにアプリケーション名と設定ファイル名をくっ付けただけの簡単なお仕事である。
[Go] 1.13 では [`os`]`.UserConfigDir()` 関数は以下のように記述されている。

```go
// UserConfigDir returns the default root directory to use for user-specific
// configuration data. Users should create their own application-specific
// subdirectory within this one and use that.
//
// On Unix systems, it returns $XDG_CONFIG_HOME as specified by
// https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html if
// non-empty, else $HOME/.config.
// On Darwin, it returns $HOME/Library/Application Support.
// On Windows, it returns %AppData%.
// On Plan 9, it returns $home/lib.
//
// If the location cannot be determined (for example, $HOME is not defined),
// then it will return an error.
func UserConfigDir() (string, error) {
	var dir string

	switch runtime.GOOS {
	case "windows":
		dir = Getenv("AppData")
		if dir == "" {
			return "", errors.New("%AppData% is not defined")
		}

	case "darwin":
		dir = Getenv("HOME")
		if dir == "" {
			return "", errors.New("$HOME is not defined")
		}
		dir += "/Library/Application Support"

	case "plan9":
		dir = Getenv("home")
		if dir == "" {
			return "", errors.New("$home is not defined")
		}
		dir += "/lib"

	default: // Unix
		dir = Getenv("XDG_CONFIG_HOME")
		if dir == "" {
			dir = Getenv("HOME")
			if dir == "" {
				return "", errors.New("neither $XDG_CONFIG_HOME nor $HOME are defined")
			}
			dir += "/.config"
		}
	}

	return dir, nil
}
```

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`gocli`]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[CC0]: http://creativecommons.org/publicdomain/zero/1.0/ "Creative Commons — CC0 1.0 Universal"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"
[`filepath`]: https://golang.org/pkg/path/filepath/ "filepath - The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
