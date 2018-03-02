+++
title = "time.Ticker で遊ぶ"
date = "2018-03-01T20:28:49+09:00"
update = "2018-03-02T19:36:47+09:00"
description = "複数の goroutine が協調して動いている場合は SIGNAL イベントに対して全ての goroutine が適切に処理を行う必要がある。"
image = "/images/attention/go-code2.png"
tags        = [ "golang", "programming", "time", "channel", "context", "goroutine" ]

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

相変わらず小ネタで。
今回の目標は2つ。

1. 一定周期ごとの処理を行う
2. Ctrl+C 等の割り込み処理を行う

## 一定周期ごとの処理を行う

[Go 言語]で一定周期ごとに処理を行うには [`time`]`.Ticker` が使える。
以下は1秒ごとに現在時刻を表示する処理である。

```go
package main

import (
	"fmt"
	"time"
)

func ticker() {
	t := time.NewTicker(1 * time.Second) //1秒周期の ticker
	defer t.Stop()

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		}
	}
}

func main() {
	ticker()
}
```

[`time`]`.Ticker.C` は受信 [channel] で，周期イベント発生時の時刻がセットされる。
[defer] 構文を使って終了時に [`time`]`.Ticker.Stop()` 関数で周期イベントを止めようとしているが，実際には無限ループなので， return まで到達しない（笑）

このコードはちゃんと動くが，終了条件を記述していないので Ctrl+C などで外部から強制的に止めない限り動き続ける。

## SIGNAL を捕まえる

[Go 言語]では SIGINT や SIGTERM といった OS から送信される [SIGNAL] をイベントとして [channel] に送り込む仕掛けがある（ちなみに Ctrl+C は SIGINT として送られる）。
この仕掛けを組み込んだコードが以下である。

{{< highlight go "hl_lines=15-22 28-33" >}}
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ticker() {
	t := time.NewTicker(1 * time.Second) //1秒周期の ticker
	defer t.Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
    defer signal.Stop(sig)

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		case s := <-sig:
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Stop!")
				return
			}
		}
	}
}

func main() {
	ticker()
}
{{< /highlight >}}

このやり方であれば Ctrl+C でも強制終了することなく正しく [`time`]`.Ticker.Stop()` 関数が起動される。

ここまでが main [goroutine] のみの処理の場合。
複数の [goroutine] が協調して動いている場合は [SIGNAL] イベントに対して全ての [goroutine] が適切に処理を行う必要がある。

## キャンセル・イベントの伝搬

[Go 言語] 1.7 から [`context`] が標準パッケージに加わった。
[`context`] パッケージは，名前の通り，異なるレイヤやドメイン間でコンテキスト情報を受け渡しするためのパッケージだが，キャンセル・イベントを扱うことができる。

```go
ctx := context.Background()
parent, cancelParent := context.WithCancel(ctx)
child, cancelChild := context.WithCancel(parent)
```

`cancelParent` および `cancelChild` は関数値で，これをキックすることでそれぞれの [`context`]`.Context` にキャンセル・イベントが発生する。
面白いのは `parent` で発生したイベントは `child` にも伝搬する点である（逆向きには伝搬しない）。
これを利用して，発生した [SIGNAL] に対して親の [`context`]`.Context` にイベントを発生させることによって全ての子  [`context`]`.Context` にイベントを伝搬させることが可能になる。

たとえば，先ほどの [SIGNAL] の処理は以下のように書き直すことができる。

```go
func SignalContext(ctx context.Context) context.Context {
	parent, cancelParent := context.WithCancel(ctx)
	go func() {
		defer cancelParent()

		sig := make(chan os.Signal, 1)
		signal.Notify(sig,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
        defer signal.Stop(sig)

		select {
		case <-parent.Done():
			fmt.Println("Cancel from parent")
			return
		case s := <-sig:
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Stop!")
				return
			}
		}
	}()

	return parent
}
```

`ticker()` 関数のほうも [`context`]`.Context` のイベントを拾えるよう以下のように書き直す。

```go
func ticker(ctx context.Context) error {
	t := time.NewTicker(1 * time.Second) //1秒周期の ticker
	defer t.Stop()

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		case <-ctx.Done():
			fmt.Println("Stop child")
			return ctx.Err()
		}
	}
}
```

周期処理と [SIGNAL] 受信処理を別々の [goroutine] で駆動させて両者を [`context`]`.Context` で繋ぐのである。

完全なコードは以下の通り。

```go
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ticker(ctx context.Context) error {
	t := time.NewTicker(1 * time.Second) //1秒周期の ticker
	defer t.Stop()

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		case <-ctx.Done():
			fmt.Println("Stop child")
			return ctx.Err()
		}
	}
}

func SignalContext(ctx context.Context) context.Context {
	parent, cancelParent := context.WithCancel(ctx)
	go func() {
		defer cancelParent()

		sig := make(chan os.Signal, 1)
		signal.Notify(sig,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
        defer signal.Stop(sig)

		select {
		case <-parent.Done():
			fmt.Println("Cancel from parent")
			return
		case s := <-sig:
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Stop!")
				return
			}
		}
	}()

	return parent
}

func Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	defer close(errCh)

	parent := SignalContext(ctx)
	go func() {
		child, cancelChild := context.WithTimeout(parent, 10*time.Second)
		defer cancelChild()
		errCh <- ticker(child)
	}()

	err := <-errCh
	fmt.Println("Done parent")
	return err
}

func main() {
	ctx := context.Background()
	if err := Run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Done")
}
```

ちなみに [`context`]`.WithTimeout()` 関数は [`context`]`.Context` にタイムアウト・イベントを付加する。
他にもデッドライン日時を指定する [`context`]`.WithDeadline()` 関数がある。

なんか今回は久しぶりに [Go 言語]っぽいコードだったねぇ（笑）

## 【追記】 Windows では SIGNAL を送信できない

今回は [SIGNAL] を受信する場合の話だったが，もちろん任意のプロセスに [SIGNAL] を送信することもできる。
以下は自分自身に SIGINT を投げるコード例である[^syscall1]。

[^syscall1]: [`syscall`] パッケージを使うやり方もあるが，今回は割愛する。つか， Windows 版には [`syscall`]`.Kill()` 関数が存在しないのでやりようがないのだが。

```go
proc, _ := os.FindProcess(os.Getppid())
proc.Signal(os.Interrupt)
```

ただし，このコードは Windows では（コンパイル・エラーにはならないが）動かない。
Windows 用の [`os`]`.Process.Signal()` 関数の実体は以下の通りだが

```go
func (p *Process) signal(sig Signal) error {
	handle := atomic.LoadUintptr(&p.handle)
	if handle == uintptr(syscall.InvalidHandle) {
		return syscall.EINVAL
	}
	if p.done() {
		return errors.New("os: process already finished")
	}
	if sig == Kill {
		err := terminateProcess(p.Pid, 1)
		runtime.KeepAlive(p)
		return err
	}
	// TODO(rsc): Handle Interrupt too?
	return syscall.Errno(syscall.EWINDOWS)
}
```

このように [`os`]`.Kill` (SIGKILL) 以外は効かないようになっている。
理由は不明だが TODO になってるようなので何か理由があるのだろう。

というわけでテストができないのですよ `orz`

## ブックマーク

- [Go Concurrency Patterns: Pipelines and cancellation - The Go Blog](https://blog.golang.org/pipelines)
- [Go1.7のcontextパッケージ | SOTA](https://deeeet.com/writing/2016/07/22/context/)
- [Golangのcontext.Valueの使い方 | SOTA](http://deeeet.com/writing/2017/02/23/go-context-value/)
- [goroutine にシグナルを送信する - Qiita](http://qiita.com/Kei-Kamikawa/items/620f9504daf2ec53f0b5)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`time`]: http://golang.org/pkg/time/ "time - The Go Programming Language"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"
[`syscall`]: https://golang.org/pkg/syscall/ "syscall - The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[channel]: http://golang.org/ref/spec#Channel_types "The Go Programming Language Specification - The Go Programming Language"
[defer]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[goroutine]: http://golang.org/ref/spec#Go_statements "The Go Programming Language Specification - The Go Programming Language"
[SIGNAL]: https://linuxjm.osdn.jp/html/LDP_man-pages/man7/signal.7.html "Man page of SIGNAL"
