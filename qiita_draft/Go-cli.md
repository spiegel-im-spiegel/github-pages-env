# tcnksm/gcli を使った golang によるコマンドライン・ツール開発について

今回は「[はじめての Go 言語 (on Windows) その10](http://qiita.com/spiegel-im-spiegel/items/5d2878596360af8dd753)」の余録。

「[はじめての Go 言語 (on Windows) その10](http://qiita.com/spiegel-im-spiegel/items/5d2878596360af8dd753)」で [tcnksm/gcli] を紹介したけど，この  [tcnksm/gcli] を使って cli パッケージの感想（評価までは行かない）を書いてみようと思う。なので [tcnksm/gcli] がサポートしていない cli パッケージは今回は対象外。

[tcnksm/gcli] がサポートする cli パッケージは以下のとおり。

- [flag]
- [codegangsta/cli]
- [mitchellh/cli]
- [Like Go command] pattern

[flag] は標準パッケージである。「[Like Go command] pattern」はパッケージではないのだが， Go コンパイラ実装パターンに似たコードを出力してくれる。

一言感想としてはこんな感じ。

<dl>
<dt>flag</dt>           <dd>サブコマンドがない場合にはお勧め</dd>
<dt>codegangsta/cli</dt><dd>問題外</dd>
<dt>mitchellh/cli</dt>  <dd>サブコマンド形式ならお勧めだが，フラグを扱えない</dd>
<dt>Like Go command</dt><dd>自動生成されるコードからかなりカスタマイズする場合にはお勧め</dd>
</dl>

あくまで個人の感想なので，怒らないでね。では，もう少し詳しく見ていく。

## コマンドライン・ツールの要件

「[はじめての Go 言語 (on Windows) その10](http://qiita.com/spiegel-im-spiegel/items/5d2878596360af8dd753)」でも書いたが，一般的なコマンドライン・ツールの要件は以下のとおり。

- shell（bash やコマンドプロンプトなど）からの起動を前提とする
- コマンドラインの引数または標準入力からデータや条件を入力する
- 結果を標準出力に出力する
- 結果以外の情報は標準エラー出力に出力する
- shell へ実行結果の状態を返す

実はこれにもうひとつ重要な要件というか思想のようなものがあって，それは「コマンドライン・ツールの機能はできるだけ単機能にする」というものだ。コマンドライン・ツールを単独で使うことは少なく，実際には shell script（Windows ならバッチスクリプト等）を使って複数のツールを組み合わせて使うことが多い。

### Dmain-Driven で行こう

DDD（Dmain-Driven Design）というものがある。これは Domain Model をベースにした設計パターンのことで DRY（Don't Repeat Yourself）な構造にするのに適していると言われている。 DDD は10年くらい前から台頭してきたが，思想自体は新しいものではなく，いわゆる「オブジェクト指向」設計のパターンとして長い時間（といっても半世紀は経っていないはずだけど）をかけて醸成された知見の集積と言える（ただし DDD 自体は必ずしも「オブジェクト指向」を要件としない）。

アーキテクチャの概念としては以下のようなな感じだろう。（これ，結構ひとによって違うんだよなぁ）

[![DDD Architecture](http://www.baldanders.info/spiegel/archive/DDD-fig.svg)](http://www.baldanders.info/spiegel/archive/DDD-fig.html)

この中の Domain Layer にビジネスロジックが入る。上下の階層構造ではなく左右の構造になっているのは，データの流れを意識してほしいから。 Iuput (trigger) は左から右に流れ， ouutput (response) は右から左に流れる。

普通 DDD は大規模アプリケーション（またはサービス）の設計で適用するものだが，コマンドライン・ツールでも DDD を意識することは悪くないし訓練にもなる。たとえば，コマンドライン・ツールを呼び出す shell や shell script が Presentation Layer で cli パッケージが Application Layer といった感じで考える。 cli パッケージが呼び出すロジックが Domain Layer だと考えると分かりやすいかも知れない。また，[前回作ったツール](http://qiita.com/spiegel-im-spiegel/items/5d2878596360af8dd753)では Git.io が Data Layer になるわけだ。

## サブコマンドを持たない構成の場合

話が逸れてしまった。 cli パッケージの話に戻ろう。

まずはサブコマンドを持たない構成の場合。例えばこんな感じ。

```bash
$ cmd -flag -parm parameter argument
```

条件を指定する `-flag` や `-parm` オプションがあって，その後ろに引数が続く。こういった用途なら [flag] パッケージが最適である。 [tcnksm/gcli] で実際にコードを生成してみる。

```shell
C:>gcli new -F flag -flag=flag:Bool:"Enable Flag" -flag=parm:String:"Option Parameter" cmd
  Created cmd\main.go
  Created cmd\CHANGELOG.md
  Created cmd\README.md
  Created cmd\cli_test.go
  Created cmd\cli.go
  Created cmd\version.go
====> Successfully generated cmd
```

```go:cmd/main.go
package main

import "os"

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
```

```go:cmd/cli.go
package main

import (
	"flag"
	"fmt"
	"io"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		flag bool
		parm string
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&flag, "flag", false, "Enable Flag")
	flags.StringVar(&parm, "parm", "", "Option Parameter")

	flVersion := flags.Bool("version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if *flVersion {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	_ = flag

	_ = parm

	return ExitCodeOK
}
```

実はこのままでは変数名の `flag` とパッケージ名の `flag` が被るので，ちょっと修正（オプション名をそのまま変数名にするのは何かと拙いと思う）。また，オプション以外の引数を解析するロジックはないが，これは自前で実装する。

```go:cmd/cli.go
package main

import (
	"flag"
	"fmt"
	"io"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		flagFlag  bool
		parmFlag  string
		arguments []string
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&flagFlag, "flag", false, "Enable Flag")
	flags.StringVar(&parmFlag, "parm", "", "Option Parameter")

	flVersion := flags.Bool("version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if *flVersion {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	for 0 < flags.NArg() {
		arguments = append(arguments, flags.Arg(0))
		flags.Parse(flags.Args()[1:])
	}

	fmt.Fprintln(cli.errStream, "flag =", flagFlag)
	fmt.Fprintln(cli.errStream, "parm =", parmFlag)
	fmt.Fprintln(cli.errStream, "arguments =", arguments)

	return ExitCodeOK
}
```

これをコンパイルして動かしてみる。

```shell
C:>go build ./cmd

C:>cmd -version
cmd version 0.1.0

C:>cmd arg1 arg2
flag = false
parm =
arguments = [arg1 arg2]

C:>cmd -flag -parm parm1 arg1 arg2
flag = true
parm = parm1
arguments = [arg1 arg2]
```

生成された `CLI` クラスでは，標準（エラー）出力や引数を入力として情報を整理した上で次の処理（おそらく Domain Layer）に処理を渡すことができる。かなり理想的な構成である。ただし，何故か標準入力は含まれていないため，標準入力も必要な場合は `CLI` クラスや `CLI` クラスを呼び出す `main` 関数を書き換える必要がある。

## サブコマンドを持つ構成の場合

サブコマンドを持つ構成というのは，以下のような引数の構成になっているツールである。

```
command [global options] subcommand [subcommand options] [arguments...]
```

[flag] パッケージはサブコマンドを構成するには不向き（できないことはないが面倒）なので，他のパッケージを使うことになる。

### codegangsta/cli を使用する場合

次のようなツールを考える。

```bash
$ cmd2 sub -flag -parm parameter argument
```

`sub` がサブコマンド。これを [codegangsta/cli] 用に生成してみる。

```shell
C:>gcli new -F codegangsta_cli -c sub:"Sub-command" -flag=flag:Bool:"Enable Flag" -flag=parm:String:"Option Parameter" cmd2
  Created cmd2\CHANGELOG.md
  Created cmd2\main.go
  Created cmd2\version.go
  Created cmd2\command\sub.go
  Created cmd2\command\sub_test.go
  Created cmd2\README.md
  Created cmd2\commands.go
====> Successfully generated cmd2
```

```go:cmd2/main.go
package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Spiegel"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
```

見たらわかると思うが， [`app.Run()`](http://godoc.org/github.com/codegangsta/cli#App.Run) 関数の返り値を全く見ていない。実は `error` 型の値を返すのだが，このエラー情報はメインの処理結果に対するものではない。ならばどうするかというと，サブコマンド内の処理の中で `os.Exit()` するという，かなりえげつない仕様になっている。個人的にはこの一点で [codegangsta/cli] は **問題外** なのだが，もう少しだけ解説する。

```go:cmd2/commands.go
package main

import (
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/cmd2/command"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "ENV_FLAG",
		Name:   "flag",

		Usage: "Enable Flag",
	},
	cli.StringFlag{
		EnvVar: "ENV_PARM",
		Name:   "parm",
		Value:  "",
		Usage:  "Option Parameter",
	},
}

var Commands = []cli.Command{
	{
		Name:   "sub",
		Usage:  "Sub-command",
		Action: command.CmdSub,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
```

残念ながら [tcnksm/gcli] はサブコマンド・オプションをサポートしていない。なので， `-flag` で指定したオプションは全てグローバル・オプションとして定義される。やり方としては必要なオプションを [tcnksm/gcli] に書き出し， `commands.go` の中でグローバル・オプションとして記述されているものをサブコマンド・オプションとして手動で振り分け直すのがいいだろう。今回ならこんな感じになる。

```go:cmd2/commands.go
package main

import (
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/cmd2/command"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "sub",
		Usage:  "Sub-command",
		Action: command.CmdSub,
		Flags:  []cli.Flag{
			cli.BoolFlag{
				//EnvVar: "ENV_FLAG",
				Name:   "flag",

				Usage: "Enable Flag",
			},
			cli.StringFlag{
				//EnvVar: "ENV_PARM",
				Name:   "parm",
				Value:  "",
				Usage:  "Option Parameter",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
```

ところで `cli.Commans` および `cli.Context` は以下のように定義されている。

```go:cli/command.go
type Command struct {
	// The name of the command
	Name string
	// short name of the command. Typically one character (deprecated, use `Aliases`)
	ShortName string
	// A list of aliases for the command
	Aliases []string
	// A short description of the usage of this command
	Usage string
	// A longer explanation of how the command works
	Description string
	// The function to call when checking for bash command completions
	BashComplete func(context *Context)
	// An action to execute before any sub-subcommands are run, but after the context is ready
	// If a non-nil error is returned, no sub-subcommands are run
	Before func(context *Context) error
	// An action to execute after any subcommands are run, but after the subcommand has finished
	// It is run even if Action() panics
	After func(context *Context) error
	// The function to call when this command is invoked
	Action func(context *Context)
	// List of child commands
	Subcommands []Command
	// List of flags to parse
	Flags []Flag
	// Treat all flags as normal arguments if true
	SkipFlagParsing bool
	// Boolean to hide built-in help command
	HideHelp bool

	commandNamePath []string
}
```

```go:cli/context.go
type Context struct {
	App            *App
	Command        Command
	flagSet        *flag.FlagSet
	setFlags       map[string]bool
	globalSetFlags map[string]bool
	parentContext  *Context
}
```

つまりこれが context の全てで，その中には標準入出力やオプション以外の引数情報は入っていない。じゃあ，どうするかというと `Action` property にセットする関数内で標準入出力やオプション以外の引数情報を取得して処理するしかない。そして `Action` property にセットする関数は値を返さない（返せない）。したがって呼び出される関数の内部で `os.Exit()` するしかないわけだ。やれやれ。

### mitchellh/cli を使用する場合

同じ UI を，今度は [mitchellh/cli] を使って実装してみる。

```shell
C:>gcli new -F mitchellh_cli -c sub:"Sub-command" -flag=flag:Bool:"Enable Flag" -flag=parm:String:"Option Parameter" cmd3
  Created cmd3\main.go
  Created cmd3\command\meta.go
  Created cmd3\command\version.go
  Created cmd3\cli.go
  Created cmd3\CHANGELOG.md
  Created cmd3\version.go
  Created cmd3\commands.go
  Created cmd3\command\sub.go
  Created cmd3\README.md
  Created cmd3\command\sub_test.go
====> Successfully generated cmd3
```

```go:cmd3/cli.go
package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/spiegel-im-spiegel/cmd3/command"
)

func Run(args []string) int {

	// Meta-option for executables.
	// It defines output color and its stdout/stderr stream.
	meta := &command.Meta{
		Ui: &cli.ColoredUi{
			InfoColor:  cli.UiColorBlue,
			ErrorColor: cli.UiColorRed,
			Ui: &cli.BasicUi{
				Writer:      os.Stdout,
				ErrorWriter: os.Stderr,
				Reader:      os.Stdin,
			},
		}}

	return RunCustom(args, Commands(meta))
}

func RunCustom(args []string, commands map[string]cli.CommandFactory) int {

	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	cli := &cli.CLI{
		Args:       args,
		Commands:   commands,
		Version:    Version,
		HelpFunc:   cli.BasicHelpFunc(Name),
		HelpWriter: os.Stdout,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute: %s\n", err.Error())
	}

	return exitCode
}
```

```go:cmd3/commands.go
package main

import (
	"github.com/mitchellh/cli"
	"github.com/spiegel-im-spiegel/cmd3/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"sub": func() (cli.Command, error) {
			return &command.SubCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
```


実は， [mitchellh/cli] 自身はフラグをサポートしていない。そこで [tcnksm/gcli] が [mitchellh/cli] をラップするコードを生成することで [mitchellh/cli] の弱点をカバーしている。

最終的に [mitchellh/cli] に渡すインスタンスは以下のインタフェースを満たしていればよい。

```go:cli/command.go
package cli

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

// CommandFactory is a type of function that is a factory for commands.
// We need a factory because we may need to setup some state on the
// struct that implements the command itself.
type CommandFactory func() (Command, error)

```

例えば `sub` サブコマンド `command.SubCommand` は以下のようになっている。

```go:cmd3/command/sub.go
package command

import (
	"strings"
)

type SubCommand struct {
	Meta
}

func (c *SubCommand) Run(args []string) int {
	// Write your code here
	return 0
}

func (c *SubCommand) Synopsis() string {
	return "Sub-command"
}

func (c *SubCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
```

このうち `Run()` 関数が処理の本体である。 [tcnksm/gcli] では空っぽの関数しか生成してくれないが，実際にはこの中で引数の解析を行う。引数解析には [flag] が使える。

```go:cmd3/command/sub.go
func (c *SubCommand) Run(args []string) int {
    var (
		flagFlag  bool
		parmFlag  string
		arguments []string
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
    flags.Usage = func() {}

	flags.BoolVar(&flagFlag, "flag", false, "Enable Flag")
	flags.StringVar(&parmFlag, "parm", "", "Option Parameter")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return 1
	}
	for 0 < flags.NArg() {
		arguments = append(arguments, flags.Arg(0))
		flags.Parse(flags.Args()[1:])
	}

	c.Ui.Output("flag =", flagFlag)
	c.Ui.Output("parm =", parmFlag)
	c.Ui.Output("arguments =", arguments)

	return 0
}
```

DDD 的にみれば，生成された `command` サブパッケージは Application Layer の Facade と見なせる。サブコマンドを呼び出す側（Application Layer）は `command` サブパッケージの各コマンドを定型的に呼び出すだけでよく，サブコマンドを呼び出される側（Domain Layer）は `command` サブパッケージとの間で話が通じていれば「友達の友達」のことは気にしなくてもよい。

#### 標準入出力について

サブコマンド `command.SubCommand` に埋め込（embed）まれている Meta は以下の構造になっている。

```go:cmd3/command/meta.go
package command

import "github.com/mitchellh/cli"

// Meta contain the meta-option that nearly all subcommand inherited.
type Meta struct {
	Ui cli.Ui
}
```

更に `cli.Ui` は interface 型で

```go:cli/ui.go
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

と定義されていて，生の標準入出力を隠ぺいするようになっている。しかし，実際には標準入出力をストリームとして直接制御したい場合がある。また，オプション等で標準入出力をファイルに付け替えたい場合もあるので，標準入出力を隠ぺいするのは必ずしも得策ではない。この場合はいっそ

```go:cmd3/command/meta.go
package command

import "github.com/mitchellh/cli"

// Meta contain the meta-option that nearly all subcommand inherited.
type Meta struct {
	Ui cli.BasicUi
}
```

などと書き換えたほうがいいかもしれない。この辺はお好みで。

### Like Go command

最後に “[Like Go command]” パターンのコードを生成してみる。

```shell
C:>gcli new -F go_cmd -c sub:"Sub-command" -flag=flag:Bool:"Enable Flag" -flag=parm:String:"Option Parameter" cmd4
  Created cmd4\CHANGELOG.md
  Created cmd4\README.md
  Created cmd4\sub.go
  Created cmd4\sub_test.go
  Created cmd4\main.go
====> Successfully generated cmd4
```

このパターンでは `Command` をサブコマンド毎の context として定義している。

```go:cmd4/main.go
// A Command is an implementation of a cmd4 command
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(args []string) int

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'cmd4 help' output.
	Short string

	// Long is the long message shown in the 'cmd4 help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet
}

// Commands lists the available commands and help topics.
// The order here is the order in which they are printed by 'cmd4 help'.
var commands = []*Command{
	cmdSub,
}
```

```go:cmd4/sub.go
package main

var cmdSub = &Command{
	Run:       runSub,
	UsageLine: "sub ",
	Short:     "Sub-command",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdSub.Flag.BoolVar(&flagA, "a", false, "")
}

// runSub executes sub command and return exit code.
func runSub(args []string) int {
	return 0
}
```

Application Layer のロジックが全て `main` パッケージに展開されているためカスタマイズしやすのが特徴。 [mitchellh/cli] で実装しづらい構成なら “[Like Go command]” パターンを試してみるのもいいかもしれない。

[tcnksm/gcli]: https://github.com/tcnksm/gcli "The easy way to build Golang command-line application."
[flag]: https://golang.org/pkg/flag/
[codegangsta/cli]: http://deeeet.com/writing/2014/06/22/cli-init/
[mitchellh/cli]: https://github.com/mitchellh/cli
[Like Go command]: https://github.com/golang/go/blob/master/src/cmd/go/main.go#L30#L51
