+++
title = "並行処理を待ち合わせる（WaitGroup と Barrier パターン）"
date =  "2025-11-09T19:39:03+09:00"
description = "この記事では比較的簡単な実装である，複数の goroutine を待ち合わせる方法を紹介する。"
image = "/images/attention/go-logo_blue.png"
tags = [ "concurrency", "programming", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] の goroutine は起動すると外側から直接制御する方法がない。
とはいっても goroutine 間で同期を取ることは多く，そのための仕組みがいくつか提供されている。
この記事では比較的簡単な実装である，複数の goroutine を待ち合わせる方法を紹介する。

## タスクの完了を待ち合わせる

まずは5つの goroutine を起動し，それぞれ某かのタスクを実行するプログラムを考える。
まずは素朴に以下のコードで[^l1]。

```go
package main
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	tasks := 5 // Number of tasks
	for t := range tasks {
		go func() {
			fmt.Println("Start Task:", t+1)
			time.Sleep(time.Duration(rand.IntN(5)+1) * time.Second) // Random duration between 1 and 5 seconds
			fmt.Println("End Task:", t+1)
		}()
	}
	fmt.Println("All Complete")
}
```

（[Go] の関数は全て関数閉包（closure）として機能するため goroutine として定義した関数はループ変数 `t` を close over している。この `for-range` ブロックのループ変数 `t` は，ループするごとに新たにインスタンスが生成されている点に注意）

実際には，このコードはうまく動かない。
実行するとこんな感じ：

```text
$ go run sample1.go
All Complete
```

実はメイン goroutine が終了する際に他の goroutine も強制終了されてしまうため，このような結果になる。
これを安直に解決する手段として，メイン goroutine を適当な時間スリープさせる方法がある。
たとえばこんな感じ[^l1]：

[^l1]: `for i := range n { ... }` は [Go] 1.22 で導入された range over int と呼ばれる構文で ，`for i := 0; i < n; i++ { ... }` と同じ意味になる。糖衣構文（syntactic sugar）の一種ですな。

```go {hl_lines=[18]}
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	tasks := 5 // Number of tasks
	for t := range tasks {
		go func() {
			fmt.Println("Start Task:", t+1)
			time.Sleep(time.Duration(rand.IntN(5)+1) * time.Second) // Random duration between 1 and 5 seconds
			fmt.Println("End Task:", t+1)
		}()
	}
	time.Sleep(10 * time.Second)
	fmt.Println("All Complete")
}
```

これを実行すると確かに

```text
$ go run sample1b.go
Start Task: 5
Start Task: 4
Start Task: 2
Start Task: 1
Start Task: 3
End Task: 3
End Task: 5
End Task: 4
End Task: 1
End Task: 2
All Complete
```

などと一見うまく動いているようだが，このスリープ時間には根拠がないため（10秒もあれば全タスク完了してるやろ？），条件が変わればうまく動かなくなる可能性がある。

複数のタスクの完了を待ち合わせるために，標準パッケージには [`sync`]`.WaitGroup` が用意されている。
使い方はこんな感じ：

```go {hl_lines=["12-13",16,22]}
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	tasks := 5             // Number of tasks
	wg := sync.WaitGroup{} // WaitGroup to synchronize goroutines
	wg.Add(tasks)          // Add the number of tasks to the WaitGroup
	for t := range tasks {
		go func() {
			defer wg.Done() // Ensure Done is called when the goroutine completes
			fmt.Println("Start Task:", t+1)
			time.Sleep(time.Duration(rand.IntN(5)+1) * time.Second) // Random duration between 1 and 5 seconds
			fmt.Println("End Task:", t+1)
		}()
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All Complete")
}
```

これは goroutine の総数があらかじめ分かっている場合だが goroutine 数が動的に変化する場合でも使える。

```go {hl_lines=[11,13,15,24]}
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{} // WaitGroup to synchronize goroutines
	for t := 1; ; t++ {
		wg.Add(1) // Add a task to the WaitGroup
		go func() {
			defer wg.Done() // Ensure Done is called when the goroutine completes
			fmt.Println("Start Task:", t)
			time.Sleep(time.Duration(rand.IntN(5)+1) * time.Second) // Random duration between 1 and 5 seconds
			fmt.Println("End Task:", t)
		}()
		if rand.IntN(6) > 4 { // Randomly decide to stop adding tasks
			break
		}
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All Complete")
}
```

ポイントは goroutine を起動する直前に `wg.Add(1)` を呼び出すことである。
実行結果はこんな感じ。

```text
$ go run sample1d.go
Start Task: 3
Start Task: 1
Start Task: 2
End Task: 2
End Task: 1
End Task: 3
All Complete
```

### for ブロックのループ変数について

先ほどのサンプルコードでは `for` ブロック

```go
for t := 1; ; t++ {
    ...
}
```

で宣言したループ変数 `t` をそのまま goroutine 内で参照しているが，古いバージョンの [Go] では意図しない出力になる。
これはループ変数 `t` が `for` ブロック内で単一のインスタンスになっていて，ループするごとに値が上書きされるためである。
この問題を回避するために以下のように `t` を shadowing する方法が推奨されていた。

```go {hl_lines=[2]}
for t := 1; ; t++ {
    t := t // Shadowing to capture the current value of t
    ...
}
```

最近のバージョン（1.22 以降）では `for` ブロック内でループ変数が内部の関数（閉包）から直接参照されている場合は暗黙的に shadowing が行われる。
上述のような記述は不要になったわけだ。

### WaitGroup.Go メソッド

[Go] 1.25 で [`sync`]`.WaitGroup` に `Go()` メソッドが追加された。
実装はこんな感じ。

```go
// Go calls f in a new goroutine and adds that task to the [WaitGroup].
// When f returns, the task is removed from the WaitGroup.
//
// The function f must not panic.
//
// If the WaitGroup is empty, Go must happen before a [WaitGroup.Wait].
// Typically, this simply means Go is called to start tasks before Wait is called.
// If the WaitGroup is not empty, Go may happen at any time.
// This means a goroutine started by Go may itself call Go.
// If a WaitGroup is reused to wait for several independent sets of tasks,
// new Go calls must happen after all previous Wait calls have returned.
//
// In the terminology of [the Go memory model], the return from f
// "synchronizes before" the return of any Wait call that it unblocks.
//
// [the Go memory model]: https://go.dev/ref/mem
func (wg *WaitGroup) Go(f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}
```

つまり，`Go()` メソッド内部で `Add(1)` と `Done()` が呼び出され，かつ引数の関数 `f` を goroutine として起動する。
これを使うと，以下のようにコードが簡潔になる。

```go {hl_lines=["14-18"]}
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	tasks := 5             // Number of tasks
	wg := sync.WaitGroup{} // WaitGroup to synchronize goroutines
	for t := range tasks {
		wg.Go(func() {
			fmt.Println("Start Task:", t+1)
			time.Sleep(time.Duration(rand.IntN(5)+1) * time.Second) // Random duration between 1 and 5 seconds
			fmt.Println("End Task:", t+1)
		})
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All Complete")
}
```

この辺も定番の慣用句（idiom）だったので，`Go()` メソッドの追加は嬉しい改善点と言えるだろう。

## タスクの開始を待ち合わせる

前節とは異なり，特定の条件が成立するまで各 goroutine を停止し，条件が成立したら一斉に実行を再開したい場合がある。
条件が成立するまでタスクを止めるため，この実装パターンをバリア（Barrier）と呼ぶ。

[Go] の標準パッケージにはバリア・パターンを直接提供する仕組みはないが，条件変数を制御する [`sync`]`.Cond` を用いて比較的簡単に実装できる。

前節のコードを元に，指定した数の goroutine の起動が揃うまで待機し，揃った時点で一斉にタスクを開始するコードを実装してみよう。
まずはバリアパターンを駆動する構造体 `Barrier` を定義する。

```go
type Barrier struct {
	total int // Number of goroutines to wait for
	count int // Current count of arrived goroutines
	cond  *sync.Cond // Condition variable to manage waiting and signaling
}
```

この構造体の初期化関数を以下のように定義する。

```go
func NewBarrier(total int) *Barrier {
	return &Barrier{
		total: total,
		count: 0,
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}
```

さらに `Barrier` に対して `Wait()` メソッドを定義する。

```go
func (b *Barrier) Wait() {
	b.cond.L.Lock()         // Acquire the lock before modifying count
	defer b.cond.L.Unlock() // Ensure the lock is released when done
	b.count++               // Increment the count of arrived goroutines
	if b.count >= b.total {
		b.count = 0
		b.cond.Broadcast() // Wake up all waiting goroutines
	} else {
		b.cond.Wait() // Wait until the barrier is released
	}
}
```

`Wait()` メソッドが呼ばれるたびに `count` がインクリメントされ，`count` が `total` に達したら `cond.Broadcast()` メソッドで全ての待機中 goroutine を起こす。

`Barrier` を使った `main()` 関数はこんな感じ：

```go {hl_lines=[3,"6-8"]}
func main() {
	tasks := 5                   // Number of tasks
	barrier := NewBarrier(tasks) // Create a barrier for the tasks
	for t := range tasks {
		go func() {
			time.Sleep(time.Duration(rand.IntN(3)+1) * time.Second) // Simulate preparation time
			fmt.Println("Ready Task:", t+1)
			barrier.Wait() // Wait at the barrier
			fmt.Println("Start Task:", t+1)
			time.Sleep(time.Duration(rand.IntN(5)+1) * time.Second) // Random duration between 1 and 5 seconds
			fmt.Println("End Task:", t+1)
		}()
	}
	time.Sleep(10 * time.Second)
	fmt.Println("All Complete")
}
```

これを実行するとこんな感じになる。

```text
$ go run sample2.go
Ready Task: 2
Ready Task: 3
Ready Task: 1
Ready Task: 4
Ready Task: 5
Start Task: 5
Start Task: 4
Start Task: 3
Start Task: 2
Start Task: 1
End Task: 2
End Task: 1
End Task: 3
End Task: 4
End Task: 5
All Complete
```

各 goroutine とも `Ready` 状態が揃ってからタスクを開始しているのが分かる[^c1]。

[^c1]: 今回の記事は『[Go言語で学ぶ並行プログラミング](https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: Go言語で学ぶ並行プログラミング　他言語にも適用できる原則とベストプラクティス impress top gearシリーズ eBook : James Cutajar, 柴田 芳樹: Kindleストア")』を参考にしている。 [`sync`]`.Cond` の `Broadcast()` メソッドは待機中の全ての goroutine を起こすため，実際には起こされたタイミングで本当に条件を満たしているかどうか確認する処理が必要かもしれない。もし条件を満たしていなければ再び `Wait()` する。

さらにこれに前節の完了待ち合わせ機能を入れたコード全体は以下の通り：

```go
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Barrier struct {
	total int        // Number of goroutines to wait for
	count int        // Current count of arrived goroutines
	cond  *sync.Cond // Condition variable to manage waiting and signaling
}

func NewBarrier(total int) *Barrier {
	return &Barrier{
		total: total,
		count: 0,
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()         // Acquire the lock before modifying count
	defer b.cond.L.Unlock() // Ensure the lock is released when done
	b.count++               // Increment the count of arrived goroutines
	if b.count >= b.total {
		b.count = 0
		b.cond.Broadcast() // Wake up all waiting goroutines
	} else {
		b.cond.Wait() // Wait until the barrier is released
	}
}

func main() {
	tasks := 5                   // Number of tasks
	barrier := NewBarrier(tasks) // Create a barrier for the tasks
	wg := sync.WaitGroup{}       // WaitGroup to synchronize goroutines
	wg.Add(tasks)                // Add the number of tasks to the WaitGroup
	for t := range tasks {
		go func() {
			defer wg.Done()                                         // Ensure Done is called when the goroutine completes
			time.Sleep(time.Duration(rand.IntN(3)+1) * time.Second) // Simulate preparation time
			fmt.Println("Ready Task:", t+1)
			barrier.Wait() // Wait at the barrier
			fmt.Println("Start Task:", t+1)
			time.Sleep(time.Duration(rand.IntN(5)+1) * time.Second) // Random duration between 1 and 5 seconds
			fmt.Println("End Task:", t+1)
		}()
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All Complete")
}
```

起動条件は何でもいい。
今回のように決められた数の goroutine が揃ったら即起動でもいいし，特定の時刻が来たら待機中の goroutine を起こすでもいいだろう。
Channel と組み合わせる方法もあるかも知れない。

## ブックマーク

- [Go 1.22リリース連載始まります & ループの変化とTinyGo 0.31 | フューチャー技術ブログ](https://future-architect.github.io/articles/20240129a/)
- [[Go 1.25] WaitGroup.Go()を使って既存コードを書き換える際の注意点](https://zenn.dev/kudotaka0421/articles/9659653cae84f0)

[Go]: https://go.dev/
[`sync`]: https://pkg.go.dev/sync "sync package - sync - Go Packages"

## 参考図書

{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4814401191" %}} <!-- 初めてのGo言語 第2版 -->
