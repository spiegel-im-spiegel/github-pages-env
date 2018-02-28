+++
title = "Go 言語による Lua 実装を試してみた"
date = "2018-02-28T19:00:05+09:00"
update = "2018-02-28T21:26:17+09:00"
description = "引数を伴う場合はどうするんだろう。 arg みたいなのはないっぽいのでグローバル変数として自前でセットしないといけないかな？"
image = "/images/attention/go-code2.png"
tags        = [ "golang", "lua", "engineering", "programming", "language" ]

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
  mathjax = true
  mermaidjs = false
+++

[Lua] 言語について調べてたら，たまたま [Go 言語]による VM とコンパイラの実装を見つけた。

- [milochristiansen/lua: A Lua 5.3 VM and compiler written in Go.](https://github.com/milochristiansen/lua)

面白そうなのでちょろんと試してみた。

## みんな大好き Hello World

まずは，みんな大好き Hello World から。

```go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/milochristiansen/lua"
	"github.com/milochristiansen/lua/lmodbase"
)

func main() {
	l := lua.NewState()

	err := l.Protect(func() {
		l.Push(lmodbase.Open)
		l.Call(0, 0)
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := l.LoadText(strings.NewReader("print(\"Hello Lua!\")"), "", 0); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := l.PCall(0, 0); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
```

これを実行すると [Lua] のスクリプト “`print("Hello Lua!")`” が実行され結果が標準出力に出力される。

```text
$ go run hello.go
Hello Lua!
```

### lua.State の中身

[`lua`]`.NewState()` 関数で生成される [`lua`]`.State` インスタンスの中身はこんな感じ。

```go
// State is the central arbitrator of all Lua operations.
type State struct {
	// Output should be set to whatever writer you want to use for logging.
	// This is where the standard script functions like print will write their output.
	// If nil defaults to os.Stdout.
	Output io.Writer

	// Add a native stack trace to errors that have attached stack traces.
	NativeTrace bool

	registry *table
	global   *table // _G
	metaTbls [typeCount]*table

	stack *stack
}
```

[`lua`]`.State.Output` に適切な Writer をセットすることで `print` 関数の出力先を変更できる（`nil` なら標準出力に出力する）。

### モジュールの読み込み

以下の部分でモジュールを読み込んでいる。

```go
err := l.Protect(func() {
    l.Push(lmodbase.Open)
    l.Call(0, 0)
})
```

[`lua`]`.State.Protect()` 関数は，関数の実行で panic が発生した際に panic を止めて error を返している。

```go
// Protect calls f inside an error handler. Use when you need to use API functions that may "raise errors" outside of
// other error handlers (such as PCall).
//
// Protect does the same cleanup PCall does, so it is safe to run code with Call inside a Protected function.
func (l *State) Protect(f func()) (err error) {
	defer l.Recover(0, false)(&err)

	f()

	return nil
}
```

[`lua`]`/lmodbase` はベースモジュールを格納するサブパッケージである。
ちなみに `print` 関数はベース・モジュールに含まれるため，読み込まないと実行時エラーになる。

```text
Value is not a function and has no __call meta method.
```

[`lua`] パッケージには他にも [`lua`]`/lmodmath`, [`lua`]`/lmodpackage`, [`lua`]`/lmodstring`, [`lua`]`/lmodtable` が用意されている。
また，これらのサブパッケージを参考にすれば独自の拡張も可能になるだろう。

### [Lua] スクリプトの読み込みと実行

以下の部分では [Lua] スクリプトの読み込みと実行を行っている。

```go
if err := l.LoadText(strings.NewReader("print(\"Hello Lua!\")"), "", 0); err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
if err := l.PCall(0, 0); err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

[`lua`]`.State.PCall()` 関数は [`lua`]`.State.Call()` 関数をラップして panic 時の処理を加えている。

```go
// PCall is exactly like Call, except instead of panicking when it encounters an error the
// error is cleanly recovered and returned.
//
// On error the stack is reset to the way it was before the call minus the function and it's arguments,
// the State may then be reused.
func (l *State) PCall(args, rtns int) (err error) {
	defer l.Recover(args+1, true)(&err)

	l.Call(args, rtns)
	return nil
}
```

[`lua`]`.State.LoadText()` 関数の代わりに [`lua`]`.State.LoadBinary()` 関数を使えばコンパイル済みのバイナリを読み込ませることもできるようだ。

## [Lua] スクリプトファイルを読み込んで実行する

大体分かったところで [Lua] スクリプトファイルを読み込んで実行してみる。
用意したファイルはフィボナッチ数を求めるスクリプト。

```lua
function fibonacciNumber(n)
	if n < 2 then
		return n
	end
	return fibonacciNumber(n - 2) + fibonacciNumber(n - 1)
end

print(fibonacciNumber(38))
```

スクリプトを読み込んで実行するコードはこんな感じ。

```go
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/milochristiansen/lua"
	"github.com/milochristiansen/lua/lmodbase"
	"github.com/milochristiansen/lua/lmodmath"
	"github.com/milochristiansen/lua/lmodpackage"
	"github.com/milochristiansen/lua/lmodstring"
	"github.com/milochristiansen/lua/lmodtable"
)

func Run(r io.Reader) error {
	l := lua.NewState()

	err := l.Protect(func() {
		l.Push(lmodbase.Open)
		l.Call(0, 0)

		l.Push(lmodmath.Open)
		l.Call(0, 0)

		l.Push(lmodpackage.Open)
		l.Call(0, 0)

		l.Push(lmodstring.Open)
		l.Call(0, 0)

		l.Push(lmodtable.Open)
		l.Call(0, 0)
	})
	if err != nil {
		return err
	}

	if err := l.LoadText(r, "", 0); err != nil {
		return err
	}
	if err := l.PCall(0, 0); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	if err := Run(file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
```

これを実行するとこんな感じになった。

```text
$ go run glua.go fibonacci.lua
39088169
```

うーん。
引数を伴う場合はどうするんだろう。
`arg` みたいなのはないっぽいのでグローバル変数として自前でセットしないといけないかな？

まぁ，気が向いたらまた今度。

## ブックマーク

- [inforno :: LuaのGo言語実装を公開しました](http://inforno.net/articles/2015/02/15/gopher-lua-released) : こちらは [Lua] 5.1 を [Go 言語]で実装している
    - [zetamatta/expect: expect for Command Prompt powered by GopherLua](https://github.com/zetamatta/expect)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[Lua]: https://www.lua.org/ "The Programming Language Lua"
[`lua`]: https://github.com/milochristiansen/lua "milochristiansen/lua: A Lua 5.3 VM and compiler written in Go."

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
