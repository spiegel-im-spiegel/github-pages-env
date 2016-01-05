+++
date = "2016-01-05T22:06:41+09:00"
description = "ファサード・パターンは DDD (Domain-Driven Design) と相性がよい。普通は Web アプリケーションのような多様なサブシステムを持つシステムを設計する際に導入する考え方だが， CLI の場合でもサブコマンドを構成するのであればファサード・パターンがよいだろう。"
draft = false
tags = ["golang", "cli", "facade"]
title = "コマンドライン・インタフェースとファサード・パターン"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

[Go 言語]コンパイラには [`flag`] パッケージが標準で提供されており，いわゆるコマンドライン・インタフェース（Command line interface; CLI）の操作はこれでまかなうことができる。
ただし [`flag`] パッケージではサブコマンドをサポートしていないためサブコマンドを構成したい場合は少し工夫が必要となる。
ちなみにサブコマンドとは，以下のようなコマンドラインの構成になっているアプリケーションである。

```
$ command [golabal options] <sub-command> [sub-options] [arguments]
```

たとえば [Go 言語]コンパイラの `go run` もサブコマンドだし， [git] の `git commit` とかもサブコマンドである。

## コマンドライン・インタフェースと UNIX Philosophy

ところで CLI でよく引き合いに出されるのが “[UNIX Philosophy]” と呼ばれるアプリケーションを作る際の哲学というか指針のようなものである。
曰く

1. Small is beautiful. （小さいものは美しい）
2. Make each program do one thing well. （各プログラムが一つのことをうまくやるようにせよ）
3. Build a prototype as soon as possible. （できる限り早くプロトタイプを作れ）
4. Choose portability over efficiency. （効率よりも移植しやすさを選べ）
5. Store data in flat text files. （単純なテキストファイルにデータを格納せよ）
6. Use software leverage to your advantage. （ソフトウェアの効率を優位さとして利用せよ）
7. Use shell scripts to increase leverage and portability. （効率と移植性を高めるためにシェルスクリプトを利用せよ）
8. Avoid captive user interfaces. （拘束的なユーザーインターフェースは作るな）
9. Make every program a Filter. （全てのプログラムはフィルタとして振る舞うようにせよ）

の9項目[^up]。
昨今は UNIX 互換環境でも GUI が普通になってきたので対話型のインタフェースも増えてきたが，それでも従来の CUI shell 上で動作するアプリケーションの需要が減ったわけではなく，サーバサイドではむしろ需要は大きくなっていると言ってもいい。

[^up]: 翻訳は [Wikipedia の記事](https://ja.wikipedia.org/wiki/UNIX%E5%93%B2%E5%AD%A6)から拝借させてもらった。ちなみに [Wikipedia のコンテンツは基本的には by-sa ライセンスで公開](https://ja.wikipedia.org/wiki/Wikipedia:Text_of_Creative_Commons_Attribution-ShareAlike_3.0_Unported_License)されている。

[Go 言語]で CLI アプリケーションを作る際に気をつける点としては

- 他のツールと shell を介して連携できるよう標準入出力を使ったフィルタプログラムとする
- 外部データの入出力は JSON, YAML, TOML といったテキストを用い UTF-8 文字エンコーディングに統一する
- コードの可搬性（または移植性）を考慮し，プラットフォーム依存を避けるようにする

といったところだろうか。
もともと [Go 言語]はクロスプラットフォーム開発に強いため，それほど難しい要件ではないはずである。

## サブコマンドとファサード・パターン

サブコマンド方式は一見 “[UNIX Philosophy]” に反しているように見えるが， [Go 言語]の場合は全てのパッケージをひとつの実行モジュールに結合してしまうため，関連する機能をサブコマンドとして組み込むのは悪くないやりかたである。

サブコマンドを構成する場合は「ファサード・パターン（facade pattern）」で考えるとよい。
「ファサード」は「建物の正面」という意味だそうで，システム内の各機能（サブシステム）の窓口のように機能する。

{{< fig-img src="/images/facade-pattern.svg" title="Facade Pattern" link="/images/facade-pattern.svg" >}}

この図のようにファサード・パターンは DDD (Domain-Driven Design) と相性がよい。
普通は Web アプリケーションのような多様なサブシステムを含むシステムを設計する際に導入する考え方だが， CLI の場合でもサブコマンドを構成するのであればファサード・パターンがよいだろう。

## mitchellh/cli パッケージ

CLI をサポートするパッケージはいくつか公開されているのだが，この中で今回は [mitchellh/cli] を紹介する。
[mitchellh/cli] はサブコマンドをファサード・パターンで実装するのに便利な機能を実装している。

### Command インタフェース

まずは `Command` インタフェース。

```go
// A command is a runnable sub-command of a CLI.
type Command interface {
	// Help should return long-form help text that includes the command-line
	// usage, a brief few sentences explaining the function of the command,
	// and the complete list of flags the command accepts.
	Help() string

	// Run should run the actual command with the given CLI instance and
	// command-line arguments. It should return the exit status when it is
	// finished.
	Run(args []string) int

	// Synopsis should return a one-line, short synopsis of the command.
	// This should be less than 50 characters ideally.
	Synopsis() string
}
```

`Command` インタフェースはサブコマンドの context を構成するのに使う。
[mitchellh/cli] は `Command` インタフェースに適合する型（[type]）のインスタンスを受け取ってサブコマンドの制御を行う[^t]。
さらに以下の関数値（function value）を示す型 `CommandFactory` も用意されている。

[^t]: 型（[type]）については「[Go 言語における「オブジェクト」]({{< relref "golang/object-oriented-programming.md" >}})」を参照のこと。

```go
// CommandFactory is a type of function that is a factory for commands.
// We need a factory because we may need to setup some state on the
// struct that implements the command itself.
type CommandFactory func() (Command, error)
```

このように `Command` 型のインスタンスを返す関数を型として定義し，この型のリストを作成するのである。

### CLI 構造体

[mitchellh/cli] に渡す context は `CLI` 構造体にまとめられている。

```go
// CLI contains the state necessary to run subcommands and parse the
// command line arguments.
type CLI struct {
	// Args is the list of command-line arguments received excluding
	// the name of the app. For example, if the command "./cli foo bar"
	// was invoked, then Args should be []string{"foo", "bar"}.
	Args []string

	// Commands is a mapping of subcommand names to a factory function
	// for creating that Command implementation. If there is a command
	// with a blank string "", then it will be used as the default command
	// if no subcommand is specified.
	Commands map[string]CommandFactory

	// Name defines the name of the CLI.
	Name string

	// Version of the CLI.
	Version string

	// HelpFunc and HelpWriter are used to output help information, if
	// requested.
	//
	// HelpFunc is the function called to generate the generic help
	// text that is shown if help must be shown for the CLI that doesn't
	// pertain to a specific command.
	//
	// HelpWriter is the Writer where the help text is outputted to. If
	// not specified, it will default to Stderr.
	HelpFunc   HelpFunc
	HelpWriter io.Writer

	once           sync.Once
	isHelp         bool
	subcommand     string
	subcommandArgs []string
	topFlags       []string

	isVersion bool
}
```

構造体の中に `CommandFactory` のリストが含まれていることがお分かりだろうか。

```go
Commands map[string]CommandFactory
```

これによってサブコマンド名と対応する処理を関連付けている。

### Ui インタフェース

入出力関数群を持つ `Ui` インタフェースは以下のように定義されている。

```go
// Ui is an interface for interacting with the terminal, or "interface"
// of a CLI. This abstraction doesn't have to be used, but helps provide
// a simple, layerable way to manage user interactions.
type Ui interface {
	// Ask asks the user for input using the given query. The response is
	// returned as the given string, or an error.
	Ask(string) (string, error)

	// AskSecret asks the user for input using the given query, but does not echo
	// the keystrokes to the terminal.
	AskSecret(string) (string, error)

	// Output is called for normal standard output.
	Output(string)

	// Info is called for information related to the previous output.
	// In general this may be the exact same as Output, but this gives
	// Ui implementors some flexibility with output formats.
	Info(string)

	// Error is used for any error messages that might appear on standard
	// error.
	Error(string)

	// Warn is used for any warning messages that might appear on standard
	// error.
	Warn(string)
}
```

更に `Ui` の特化クラスとして `BasicUi` や `PrefixedUi` や `ColoredUi` が定義されている。
`ColoredUi` は出力をカラーにできるが，残念ながら Windows のコマンドプロンプトには対応していないようだ。

`Ui` インタフェースは `Command` インタフェースと組み合わせてサブコマンド側の context を構成するのに使う。

### mitchellh/cli パッケージのメリット

上述したように [mitchellh/cli] はサブコマンドをファサード・パターンで実装するのに便利な機能を実装している。
なおかつ [mitchellh/cli] ではファサード・パターンを入れ子にすることができる。
たとえばサブコマンドのサブコマンドを構成することもできるのだ。

## mitchellh/cli を使ってファサード・パターンを組んでみる

[mitchellh/cli] をファサード・パターンとして組みやすくするためのラッパーとして [spiegel-im-spiegel/gofacade] パッケージを作ってみた[^li]。

[^li]: [spiegel-im-spiegel/gofacade] は [CC0](https://creativecommons.org/publicdomain/zero/1.0/) で公開している。個人的には実証コードの扱いなので，（著作権情報の書き換えも含めて）自由に利用して 構わない。

まず，入出力の Context を定義するためのクラスとして `Context` 構造体を作った。
中身は [`cli`].`BasicUi` 構造体を埋め込んでいるだけである[^ebd1]。

[^ebd1]: なんでこんな回りくどいことをしているかというと， [mitchellh/cli] パッケージをカプセル化したかったから。

```go
//Context inheritance cli.BasicUi
type Context struct {
	//Embedded BasicUi
	*cli.BasicUi
}
```

更に `Context` 構造体を包含する `Facade` 構造体を定義する。

```go
// Facade is context of facade
type Facade struct {
	//UI defines user interface of the Cli
	Cxt *Context
	// commands is a mapping of subcommand names to a factory function
	commands map[string]cli.CommandFactory
}
```

`Facade` 構造体には [`cli`].`CommandFactory` のリストを含んでいる。
このリストに [`cli`].`Command` インタフェースに適合するインスタンスを追加するための関数がこれ[^cl]。

[^cl]: [Go 言語]では関数は全て関数閉包（closure）として機能する。

```go
// AddCommand add command
func (f *Facade) AddCommand(name string, command cli.Command) {
	f.commands[name] = func() (cli.Command, error) {
		return command, nil
	}
}
```

実際にファサードを実行するには以下の関数を起動する。

```go
// Run facade
func (f *Facade) Run(appName, version string, args []string) (int, error) {
	c := cli.NewCLI(appName, version)
	c.Args = args
	c.Commands = f.commands
	c.HelpWriter = f.Cxt.Writer
	return c.Run()
}
```

他に細かい道具はあるが，まぁこんなもんだろう。

### spiegel-im-spiegel/gofacade の実装例

[spiegel-im-spiegel/gofacade] パッケージの実装例として [spiegel-im-spiegel/astrocalc] パッケージに CLI ツールを追加してみた。
こんな感じのコマンドラインを構成してみる。

```
$ astrocalc [-v | -h] mjdn <year> <month> <day>
```

まず `astrocalc mjdn` サブコマンドを以下のように定義する。

```go
package mjdnCmd

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spiegel-im-spiegel/astrocalc/mjdn"
	"github.com/spiegel-im-spiegel/gofacade"
)

// Name は mjdn コマンド名を定義する
const Name string = "mjdn"

// Context は mjdn コマンドのコンテキストを定義する
type Context struct {
	//Embedded gofacade.Context
	*gofacade.Context
	//AppName にはアプリケーション名を格納する
	AppName string
}

// Command は Context のインスタンスを返す
func Command(cxt *gofacade.Context, appName string) *Context {
	return &Context{Context: cxt, AppName: appName}
}

// Synopsis は mjdn コマンドの概要を返す
func (c Context) Synopsis() string {
	return "Calculation of Modified Julian Day"
}

// Help は mjdn コマンドのヘルプを返す
func (c Context) Help() string {
	helpText := `
Usage: astrocalc mjdn <year> <month> <day>
`
	return fmt.Sprintln(strings.TrimSpace(helpText))
}

// Run は mjdn コマンドを実行する
func (c Context) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.Usage = func() {
		c.Error(c.Help())
	}
	// Parse commandline flag
	if err := flags.Parse(args); err != nil {
		return gofacade.ExitCodeError
	}
	if flags.NArg() != 3 {
		c.Error(fmt.Sprintf("年月日を指定してください\n\n%s", c.Help()))
		return gofacade.ExitCodeError
	}
	argsStr := flags.Args()
	var ymd = make([]int, 3)
	for i, arg := range argsStr {
		num, err := strconv.Atoi(arg)
		if err != nil {
			c.Error(fmt.Sprintln(err))
			return gofacade.ExitCodeError
		}
		ymd[i] = num
	}
	tm := time.Date(ymd[0], time.Month(ymd[1]), ymd[2], 0, 0, 0, 0, time.UTC)
	c.Output(fmt.Sprint(mjdn.DayNumber(tm)))
	return gofacade.ExitCodeOK
}
```

ポイントは `astrocalc mjdn` サブコマンド用の context として `Context` 構造体を定義しているところ。

```go
// Context は mjdn コマンドのコンテキストを定義する
type Context struct {
	//Embedded gofacade.Context
	*gofacade.Context
	//AppName にはアプリケーション名を格納する
	AppName string
}
```

[`gofacade`].`Context` 構造体を埋め込みフィールドで定義しているのがお分かりだろうか。
[`gofacade`].`Context` はさらに [`cli`].`BasicUi` 構造体を埋め込んでいる。
また `Context` 構造体は [`cli`].`Command` インタフェースの特化クラスとして実装している。

では，この `Context` 構造体を使ってアプリケーションの起動部分を書いてみよう。

```go
package main

import (
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/astrocalc/internal/mjdnCmd"
	"github.com/spiegel-im-spiegel/gofacade"
)

const (
	// Name はアプリケーション名を定義する
	Name string = "astrocalc"
	// Version はアプリケーションのバージョン番号を定義する
	Version string = "0.1.0"
)

func setupFacade(cxt *gofacade.Context) *gofacade.Facade {
	fcd := gofacade.NewFacade(cxt)
	fcd.AddCommand(mjdnCmd.Name, mjdnCmd.Command(cxt, Name))
	return fcd
}

func main() {
	cxt := gofacade.NewContext(os.Stdin, os.Stdout, os.Stderr)
	fcd := setupFacade(cxt)
	rtn, err := fcd.Run(Name, Version, os.Args[1:])
	if err != nil {
		cxt.Error(fmt.Sprintln(err))
	}
	os.Exit(rtn)
}
```

`setupFacade()` 関数でファサードを作成し， `main()` 関数で実行しているのが分かると思う。
では実際に compile & run してみよう。

```
C:\workspace\astrocalc> pushd C:\workspace\astrocalc\src\github.com\spiegel-im-spiegel\astrocalc

C:\workspace\astrocalc\src\github.com\spiegel-im-spiegel\astrocalc> glide up
[INFO] Fetching updates for github.com/spiegel-im-spiegel/gofacade.
[INFO] Found glide.yaml in C:\workspace\astrocalc\src\github.com\spiegel-im-spiegel\astrocalc\vendor/github.com/spiegel-im-spiegel/gofacade/glide.yaml
[INFO] Scanning github.com/mitchellh/cli for dependencies.
[INFO] Scanning github.com/bgentry/speakeasy for dependencies.
[INFO] Scanning github.com/mattn/go-isatty for dependencies.
[INFO] Project relies on 4 dependencies.
[INFO] Writing glide.lock file

C:\workspace\astrocalc\src\github.com\spiegel-im-spiegel\astrocalc> popd

C:\workspace\astrocalc> go install -v ./...
github.com/spiegel-im-spiegel/astrocalc/mjdn
github.com/spiegel-im-spiegel/astrocalc/vendor/github.com/bgentry/speakeasy
github.com/spiegel-im-spiegel/astrocalc/vendor/github.com/mattn/go-isatty
github.com/spiegel-im-spiegel/astrocalc/era
github.com/spiegel-im-spiegel/astrocalc/vendor/github.com/mitchellh/cli
github.com/spiegel-im-spiegel/astrocalc/vendor/github.com/bgentry/speakeasy/example
github.com/spiegel-im-spiegel/astrocalc/vendor/github.com/spiegel-im-spiegel/gofacade
github.com/spiegel-im-spiegel/astrocalc/internal/mjdnCmd
github.com/spiegel-im-spiegel/astrocalc

C:\workspace\astrocalc> bin\astrocalc.exe -h
usage: astrocalc [--version] [--help] <command> [<args>]

Available commands are:
    mjdn    Calculation of Modified Julian Day

C:\workspace\astrocalc> bin\astrocalc.exe -h mjdn
Usage: astrocalc mjdn <year> <month> <day>

C:\workspace\astrocalc> bin\astrocalc.exe mjdn 2015 1 1
57023 (2015-01-01)
```

よしよし。
うまくいった。
なお [glide] については「[Glide で Vendoring]({{< relref "golang/vendoring-with-glide.md" >}})」を参考にどうぞ。

## ブックマーク

- [Go言語によるCLIツール開発とUNIX哲学について - ゆううきブログ](http://yuuki.hatenablog.com/entry/go-cli-unix)
- [開発者から見た UNIX 哲学とコマンドラインツールと Go言語 - TELLME.TOKYO](http://tellme.tokyo/post/2015/06/23/unix_cli_tool_go/)
- [Go言語のflagパッケージを使う - uragami note](http://ryochack.hatenablog.com/entry/2013/04/17/232753)
- [Go言語のCLIツールのpanicをラップしてクラッシュレポートをつくる | SOTA](http://deeeet.com/writing/2015/04/17/panicwrap/)
- [flag 並にシンプルでより強力な CLI パーサ kingpin の紹介 - Qiita](http://qiita.com/kumatch/items/258d7984c0270f6dd73a)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`flag`]: https://golang.org/pkg/flag/ "flag - The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[type]: https://golang.org/ref/spec#Properties_of_types_and_values "Properties of types and values"
[interface]: https://golang.org/ref/spec#Interface_types "Interface types"
[git]: https://git-scm.com/ "Git"
[UNIX Philosophy]: http://www.ru.j-npcs.org/usoft/WWW/LJ/Articles/unixtenets.html "Tenets of the UNIX Philosophy"
[mitchellh/cli]: https://github.com/mitchellh/cli "mitchellh/cli"
[`cli`]: https://github.com/mitchellh/cli "mitchellh/cli"
[spiegel-im-spiegel/gofacade]: https://github.com/spiegel-im-spiegel/gofacade "spiegel-im-spiegel/gofacade"
[`gofacade`]: https://github.com/spiegel-im-spiegel/gofacade "spiegel-im-spiegel/gofacade"
[spiegel-im-spiegel/astrocalc]: https://github.com/spiegel-im-spiegel/astrocalc "spiegel-im-spiegel/astrocalc"
[glide]: https://github.com/Masterminds/glide "Masterminds/glide"
