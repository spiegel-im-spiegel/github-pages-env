# gometalinter で楽々 lint

[Go 言語]用の lint ツールはいくつかあるのだが，これらを統合的に扱える [gometalinter] というのがあるらしい。

- [alecthomas/gometalinter: Concurrently run Go lint tools and normalise their output](https://github.com/alecthomas/gometalinter)

[Go 言語]で書かれているので（[Go 言語]の環境があれば）インストールは簡単。

```
$ go get -v github.com/alecthomas/gometalinter
```

[gometalinter] をつかって各種 lint およびフォーマッタ・ツールを一気にインストールできる。

```
$ gometalinter --install --update
Installing:
  structcheck
  interfacer
  goconst
  golint
  goimports
  dupl
  errcheck
  aligncheck
  gocyclo
  ineffassign
  unconvert
  gotype
  varcheck
  deadcode
  lll
```

これらのツールは `$GOPATH/bin` フォルダにインストールされるので `PATH` を通しておくこと。

では早速試してみよう。サンプルコードは，このまえ手遊びで作った [`github.com/spiegel-im-spiegel/zundoko`] パッケージを使おう。

```
$ go get -v github.com/spiegel-im-spiegel/zundoko
```

いきなり lint をかけてみる。

```
$ gometalinter ./...
WARNING: deadline exceeded by linter dupl on sample (try increasing --deadline)
WARNING: deadline exceeded by linter unconvert on . (try increasing --deadline)
WARNING: deadline exceeded by linter unconvert on sample (try increasing --deadline)
WARNING: deadline exceeded by linter aligncheck on . (try increasing --deadline)
WARNING: deadline exceeded by linter aligncheck on sample (try increasing --deadline)
WARNING: deadline exceeded by linter interfacer on . (try increasing --deadline)
WARNING: deadline exceeded by linter interfacer on sample (try increasing --deadline)
```

ありゃりゃ。処理に時間がかかりすぎた？ じゃあ `--fast` オプション[^fast] を付けてもう一度。

[^fast]: `--fast` オプションを付けると deadcode, dupl, goconst, gocyclo, golint, gotype, vet といった処理の速いツールのみでチェックを行う。

```
$ gometalinter --fast ./...
WARNING: deadline exceeded by linter dupl on sample (try increasing --deadline)
```

う～ん。実は私の環境では [gometalinter] が [dupl] の終了を認識してくれないようだ[^a]。なので [dupl] を除外してやってみる。

[^a]: 手動で `dupl -plumbing -threshold 50 zundoko.go` とかやってみたが問題なく動作しているので [gometalinter] 側の問題だろう，多分。

```
$ gometalinter --fast --disable=dupl ./...
```

今度はエラーは出なかった。なお `--deadline=5s` といった感じで deadline 値を調整すれば `--fast` オプションなしでも上手くいく。

では [`github.com/spiegel-im-spiegel/zundoko`] パッケージの `zundoko.go` ファイルを以下のように弄ってみる。

```go
package zundoko

import (
	"math/rand"
	"time"
)

// zundoko-choirs items
const (
	Zun     = "ズン"
	Doko    = "ドコ"
	Kiyoshi = "キ・ヨ・シ！"
)

//Choirs - zundoko-choirs list
type Choirs struct {
	c     []string
    dummy int
}

//Push is append choirs
func (c *Choirs) Push(s string) {
	c.c = append(c.c, s) //maybe panic if c is nil.
}

//CountZunDoko returns count of "ZUN" and "DOKO" choirs
func (c *Choirs) CountZunDoko() (int, int) {
	z := 0
	d := 0
	if c == nil {
		return z, d
	}
	for _, s := range c.c {
		switch s {
		case Zun:
			z++
		case Doko:
			d++
		}
	}
	return z, d
}

func (c *Choirs) String() string {
	if c == nil {
		return ""
	} else {
		content := make([]byte, 0, 128)
		for _, s := range c.c {
			content = append(content, s...)
		}
		return string(content)
	}
	return ""
}

func generate() chan string {
	ch := make(chan string)
	go func() {
		zd := [2]string{Zun, Doko}
		rand.Seed(time.Now().UnixNano())
		for {
			ch <- zd[rand.Intn(2)]
		}
	}()
	return ch
}

//Run zundoko-choirs
func Run() *Choirs {
	zd := generate()
	c := &Choirs{c: make([]string, 0)}
	zcount := 0
	for {
		s := <-zd
		c.Push(s)
		if s == Zun {
			zcount++
		} else if zcount >= 4 {
			break
		} else {
			zcount = 0
		}
	}
	c.Push(Kiyoshi)
	return c
}
```

どこに問題があるかお分かりだろうか。 lint をかけると以下の結果になる。

```
$ gometalinter --deadline=5s --disable=dupl ./...
zundoko.go:72:32:error: too few values in struct literal (gotype)
zundoko.go:54::error: unreachable code (vet)
zundoko.go:47:9:warning: if block ends with a return statement, so drop this else and outdent its block (golint)
zundoko.go:72:32:warning: too few values in struct literal (interfacer)
```

golint, gotype, interfacer, vet のそれぞれのエラー・ワーニングが集約されて出力されているのがわかると思う。ちなみに `--json` オプションを付けると出力が JSON フォーマットになる。

```
$ gometalinter --deadline=5s --disable=dupl --json ./...
[
  {"linter":"golint","severity":"warning","path":"zundoko.go","line":47,"col":9,"message":"if block ends with a return statement, so drop this else and outdent its block"},
  {"linter":"gotype","severity":"error","path":"zundoko.go","line":72,"col":32,"message":"too few values in struct literal"},
  {"linter":"vet","severity":"error","path":"zundoko.go","line":54,"col":0,"message":"unreachable code"},
  {"linter":"interfacer","severity":"warning","path":"zundoko.go","line":72,"col":32,"message":"too few values in struct literal"}
]
```

また `--vendor` オプションを付けると [Go 言語]バージョン 1.5 から採用された Vendoring 機能に対応する。その他，オプション等は以下のとおり。

```
$ gometalinter --help
usage: gometalinter.exe [<flags>] [<path>...]

Aggregate and normalise the output of a whole bunch of Go linters.

Default linters:

deadcode  (github.com/tsenart/deadcode)
      deadcode .
      ^deadcode: (?P<path>[^:]+):(?P<line>\d+):(?P<col>\d+):\s*(?P<message>.*)$
gocyclo  (github.com/alecthomas/gocyclo)
      gocyclo -over {mincyclo} .
      ^(?P<cyclo>\d+)\s+\S+\s(?P<function>\S+)\s+(?P<path>[^:]+):(?P<line>\d+):(\d+)$
gofmt
      gofmt -l -s ./*.go
      ^(?P<path>[^\n]+)$
gotype  (golang.org/x/tools/cmd/gotype)
      gotype -e {tests=-a} .
      ^(?P<path>[^\s][^\r\n:]+?\.go):(?P<line>\d+):(?P<col>\d+):\s*(?P<message>.*)$
ineffassign  (github.com/gordonklaus/ineffassign)
      ineffassign -n .
      ^(?P<path>[^\s][^\r\n:]+?\.go):(?P<line>\d+):(?P<col>\d+):\s*(?P<message>.*)$
vetshadow
      go tool vet --shadow ./*.go
      ^(?P<path>[^\s][^\r\n:]+?\.go):(?P<line>\d+):\s*(?P<message>.*)$
varcheck  (github.com/opennota/check/cmd/varcheck)
      varcheck .
      ^(?:[^:]+: )?(?P<path>[^:]+):(?P<line>\d+):(?P<col>\d+):\s*(?P<message>\w+)$
vet
      go tool vet ./*.go
      ^(?P<path>[^\s][^\r\n:]+?\.go):(?P<line>\d+):\s*(?P<message>.*)$
aligncheck  (github.com/opennota/check/cmd/aligncheck)
      aligncheck .
      ^(?:[^:]+: )?(?P<path>[^:]+):(?P<line>\d+):(?P<col>\d+):\s*(?P<message>.+)$
dupl  (github.com/mibk/dupl)
      dupl -plumbing -threshold {duplthreshold} ./*.go
      ^(?P<path>[^\s][^:]+?\.go):(?P<line>\d+)-\d+:\s*(?P<message>.*)$
golint  (github.com/golang/lint/golint)
      golint -min_confidence {min_confidence} .
      ^(?P<path>[^\s][^\r\n:]+?\.go):(?P<line>\d+):(?P<col>\d+):\s*(?P<message>.*)$
lll  (github.com/walle/lll/cmd/lll)
      lll -g -l {maxlinelength} ./*.go
      ^(?P<path>[^\s][^\r\n:]+?\.go):(?P<line>\d+):\s*(?P<message>.*)$
structcheck  (github.com/opennota/check/cmd/structcheck)
      structcheck {tests=-t} .
      ^(?:[^:]+: )?(?P<path>[^:]+):(?P<line>\d+):(?P<col>\d+):\s*(?P<message>.+)$
testify
      go test
      Location:\s+(?P<path>[^:]+):(?P<line>\d+)$\s+Error:\s+(?P<message>[^\n]+)
errcheck  (github.com/kisielk/errcheck)
      errcheck -abspath .
      ^(?P<path>[^:]+):(?P<line>\d+):(?P<col>\d+)\t(?P<message>.*)$
goconst  (github.com/jgautheron/goconst/cmd/goconst)
      goconst -min-occurrences {min_occurrences} .
      ^(?P<path>[^\s][^\r\n:]+?\.go):(?P<line>\d+):(?P<col>\d+):\s*(?P<message>.*)$
goimports  (golang.org/x/tools/cmd/goimports)
      goimports -l ./*.go
      ^(?P<path>[^\n]+)$
interfacer  (github.com/mvdan/interfacer/cmd/interfacer)
      interfacer ./
      ^(?P<path>[^\s][^\r\n:]+?\.go):(?P<line>\d+):(?P<col>\d+):\s*(?P<message>.*)$
test
      go test
      ^--- FAIL: .*$\s+(?P<path>[^:]+):(?P<line>\d+): (?P<message>.*)$
unconvert  (github.com/mdempsky/unconvert)
      unconvert .
      ^(?P<path>[^\s][^\r\n:]+?\.go):(?P<line>\d+):(?P<col>\d+):\s*(?P<message>.*)$

Severity override map (default is "warning"):

gotype -> error
test -> error
testify -> error
vet -> error

Flags:
      --help                Show context-sensitive help (also try --help-long
                            and --help-man).
      --fast                Only run fast linters.
  -i, --install             Attempt to install all known linters.
  -u, --update              Pass -u to go tool when installing.
  -f, --force               Pass -f to go tool when installing.
  -d, --debug               Display messages for failed linters, etc.
  -j, --concurrency=16      Number of concurrent linters to run.
  -e, --exclude=REGEXP ...  Exclude messages matching these regular expressions.
  -I, --include=REGEXP ...  Include messages matching these regular expressions.
  -s, --skip=DIR... ...     Skip directories with this name when expanding
                            '...'.
      --vendor              Enable vendoring support (skips 'vendor' directories
                            and sets GO15VENDOREXPERIMENT=1).
      --cyclo-over=10       Report functions with cyclomatic complexity over N
                            (using gocyclo).
      --line-length=80      Report lines longer than N (using lll).
      --min-confidence=.80  Minimum confidence interval to pass to golint.
      --min-occurrences=3   Minimum occurrences to pass to goconst.
      --dupl-threshold=50   Minimum token sequence as a clone for dupl.
      --sort=none ...       Sort output by any of none, path, line, column,
                            severity, message, linter.
  -t, --tests               Include test files for linters that support this
                            option
      --deadline=5s         Cancel linters if they have not completed within
                            this duration.
      --errors              Only show errors.
      --json                Generate structured JSON rather than standard
                            line-based output.
  -D, --disable=LINTER ...  List of linters to disable
                            (testify,test,gofmt,goimports,lll).
  -E, --enable=LINTER ...   Enable previously disabled linters.
      --linter=NAME:COMMAND:PATTERN ...
                            Specify a linter.
      --message-overrides=LINTER:MESSAGE ...
                            Override message from linter. {message} will be
                            expanded to the original message.
      --severity=LINTER:SEVERITY ...
                            Map of linter severities.
      --disable-all         Disable all linters.

Args:
  [<path>]  Directory to lint. Defaults to ".". <path>/... will recurse.
```

Vim, emacs, ATOM などでは [gometalinter] を使ったプラグインやパッケージが公開されているようだ。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[gometalinter]: https://github.com/alecthomas/gometalinter "alecthomas/gometalinter: Concurrently run Go lint tools and normalise their output"
[`github.com/spiegel-im-spiegel/zundoko`]: https://github.com/spiegel-im-spiegel/zundoko "spiegel-im-spiegel/zundoko: Zundoko-choirs"
[dupl]: https://github.com/mibk/dupl "mibk/dupl: a tool for code clone detection"

## 脚注
