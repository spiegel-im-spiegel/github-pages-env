+++
title = "Log パッケージで遊ぶ"
date =  "2018-02-26T13:19:29+09:00"
description = "description"
image = "/images/attention/go-code2.png"
tags        = [ "golang", "engineering", "programming", "logger" ]
draft = true

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

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

[`log`] パッケージの特徴は複数の [goroutine] を横断して使える，つまり [goroutine] safe であるという点である。
たとえば以下のような並行処理の場合

```go
package main

import (
	"log"
	"sync"
)

func main() {
	wait := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		index := i + 1
		wait.Add(1)
		go func() {
			defer wait.Done()
			log.Printf("Hello, No. %d\n", index)
		}()
	}
	wait.Wait()
}
```

出力結果は以下のようにきれいに出力される。
（並行処理なので順番は不定）

```text
2009/11/10 23:00:00 Hello, No. 10
2009/11/10 23:00:00 Hello, No. 4
2009/11/10 23:00:00 Hello, No. 2
2009/11/10 23:00:00 Hello, No. 5
2009/11/10 23:00:00 Hello, No. 6
2009/11/10 23:00:00 Hello, No. 1
2009/11/10 23:00:00 Hello, No. 7
2009/11/10 23:00:00 Hello, No. 8
2009/11/10 23:00:00 Hello, No. 3
2009/11/10 23:00:00 Hello, No. 9
```

## Log パッケージの中身

具体的に [`log`] パッケージを見てみよう。

{{< highlight go "hl_lines=29-30 33 40" >}}
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
func (l * Logger) Printf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprintf(format, v...))
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is used to recover the PC and is
// provided for generality, although at the moment on all pre-defined
// paths it will be 2.
func (l * Logger) Output(calldepth int, s string) error {
	now := time.Now() // get this early.
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.flag&(Lshortfile|Llongfile) != 0 {
		// Release lock while getting caller info - it's expensive.
		l.mu.Unlock()
		var ok bool
		_ , file, line, ok = runtime.Caller(calldepth)
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
	_ , err := l.out.Write(l.buf)
	return err
}
{{< /highlight >}}

`Logger.mu.Lock()` および `Logger.mu.Unlock()` 関数に注目すると [`sync`]`.Mutex` を使って排他的に処理しているのが分かるだろう。

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

この場合，以下のログが標準出力に出力される。

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

このコードのログ出力は先ほどと同じになる[^log1]。

[^log1]: ログ出力の排他処理は [`log`]`.Logger` インスタンス毎に行われるため，複数の [`log`]`.Logger` インスタンスがある場合は，出力先が競合しないように注意する必要がある。

## 出力レベルとフィルタリング

[`log`] パッケージの欠点は出力のフィルタリングができないことである。
ERROR, WARN, INFO, DEBUG といった出力レベルの設定と指定レベル以下のログについてはフィルタリングできる仕掛けが欲しいところである。

サードパーティーではフィルタリング機能を持つ logger が色々あるが，せっかく [`log`] パッケージが「できる子」なんだから，できるだけ生かす形で実装を考えてみる。

## フィルタリング付き Writer

フィルタリング機能付きの Writer を提供するパッケージがある。

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
	log.Print("[ERROR] Erring")            // and so will this
	log.Print("Message I haven't updated") // and so will this
}
```

これを実行するとこんな感じになる。

```text
2009/11/10 23:00:00 [WARN] Warning
2009/11/10 23:00:00 [ERROR] Erring
2009/11/10 23:00:00 Message I haven't updated
```

あらかじめキーワード（`"DEBUG"` など）を決めておいて，ログ出力する文字列にこのキーワードを埋め込むことで出力レベルの設定とフィルタリングを行うようだ。
実際には埋め込むキーワードの typo を防ぐために何らかの仕掛けを作る必要があると思うが，とてもシンプルな作りなので，気軽に導入できそうである。









## ブックマーク

- [ええっ！？　文字列で書くの！？　ログレベル付きロガーhashicorp/logutilsのご紹介 - Qiita](https://qiita.com/mackee_w/items/3c0940733b6c0922554c)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`log`]: https://golang.org/pkg/log/ "log - The Go Programming Language"
[`sync`]: https://golang.org/pkg/sync/ "sync - The Go Programming Language"
[hashicorp/logutils]: https://github.com/hashicorp/logutils "hashicorp/logutils: Utilities for slightly better logging in Go (Golang)."
[`logutils`]: https://github.com/hashicorp/logutils "hashicorp/logutils: Utilities for slightly better logging in Go (Golang)."

[goroutine]: http://golang.org/ref/spec#Go_statements "The Go Programming Language Specification - The Go Programming Language"

<!-- eof -->
