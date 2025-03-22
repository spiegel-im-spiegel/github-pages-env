+++
title = "条件変数とセマフォ（『Go言語で学ぶ並行プログラミング』読書会3回目）"
date =  "2025-03-22T23:49:15+09:00"
description = "条件変数 / 書き込み優先な RWMutex / セマフォ"
image = "/images/attention/go-logo_blue.png"
tags = [ "book", "golang", "engineering", "programming", "concurrency" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

「[第3回『Go言語で学ぶ並行プログラミング』オンライン読書会]」の話題より。
ちなみに第2回は記事にしていない。
あしからず。

## 条件変数

条件変数（condition variable）は，ミューテックス（mutex）と組み合わせて，並行処理が特定の条件が発生するまで待たせたいときに使う。

[Go] では標準パッケージの [`sync`]`.Cond` 型で提供されている。
こんな感じで使う。

```go
cond := sync.NewCond(&sync.Mutex{})
```

ちなみに [Java](https://docs.oracle.com/en/java/javase/24/docs/api/java.base/java/util/concurrent/locks/Condition.html "Condition (Java SE 24 & JDK 24)") では

```java
final Lock lock = new ReentrantLock();
final Condition condLock  = lock.newCondition();
```

みたいな感じにミューテックス・インスタンスから生成する。
おそらく言語やライブラリやフレームワークによって色々あるだろう。

それはさておき『[Go言語で学ぶ並行プログラミング]』に出ていた例題は

1. Stingy と Spendy の２人の登場人物（goroutine）
2. ２人は共通の銀行口座を持っていて現在100ドルの残高がある
3. Stingy は毎回10ドルずつ100万回入金する
4. Spendy は毎回50ドルずつ20万回出金する
5. 銀行口座の残高をマイナスにしてはいけない（← ここが条件変数の条件）

という要件で入出金の様子をコードで書くというもの。
過程をすっ飛ばして最終形を挙げておく。

{{<fig-quote class="nobox" type="markdown" title="『Go言語で学ぶ並行プログラミング』第5章" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
```go {hl_lines=[24,"26-27", "34-37",43]}
package main

import (
    "fmt"
    "os"
    "sync"
    "time"
)

func main() {
    money := 100
    mutex := sync.Mutex{}
    cond := sync.NewCond(&mutex)
    go stingy(&money, cond)
    go spendy(&money, cond)
    time.Sleep(2 * time.Second)
    mutex.Lock()
    fmt.Println("Money in bank account: ", money)
    mutex.Unlock()
}

func stingy(money *int, cond *sync.Cond) {
    for i := 0; i < 1000000; i++ {
        cond.L.Lock()
        *money += 10
        cond.Signal()
        cond.L.Unlock()
    }
    fmt.Println("Stingy Done")
}

func spendy(money *int, cond *sync.Cond) {
    for i := 0; i < 200000; i++ {
        cond.L.Lock()
        for *money < 50 {
            cond.Wait()
        }
        *money -= 50
        if *money < 0 {
            fmt.Println("Money is negative!")
            os.Exit(1)
        }
        cond.L.Unlock()
    }
    fmt.Println("Spendy Done")
}
```
{{</fig-quote >}}

`spendy()` 関数内 `cond.Wait()` の動作は分かりにくいかもしれないが，内部ではミューテックスを解放して待ち状態に入っている。
その後 `Signal()` を受け取ったらミューテックスを再取得して続きの処理を再開する。
ここで `*money < 50` の条件に合致すればループを抜けて出金操作を行うわけだ（合致しなければ再び `Wait()`）。

このコードの注意点は2つ。
ひとつは `Signal()` や `Wait()` の前後は必ず `Lock() ... Unlock()` で囲っておくこと。
もうひとつは `Wait()` 付近で条件を記述する際は `if` 文ではなく `for` 文でループを構成すること。
後者はわかりやすいよね。
前者は，たとえば以下のようなコードがあって

{{<fig-quote class="nobox" type="markdown" title="『Go言語で学ぶ並行プログラミング』第5章" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
```go {hl_lines=[15]}
package main

import (
    "fmt"
    "sync"
)

/*
Note: this program has a bug for demonstration purposes
We demonstrate how to fix this problem in the next listing
*/
func doWork(cond *sync.Cond) {
    fmt.Println("Work started")
    fmt.Println("Work finished")
    cond.Signal()
}

func main() {
    cond := sync.NewCond(&sync.Mutex{})
    cond.L.Lock()
    for i := 0; i < 50000; i++ {
        go doWork(cond)
        fmt.Println("Waiting for child goroutine")
        cond.Wait()
        fmt.Println("Child goroutine finished")
    }
    cond.L.Unlock()
}
```
{{</fig-quote >}}

これを実行すると大抵は（`Wait()` で受ける相手が起動する前に `Signal()` を投げちゃうので）デッドロック状態になるのだが，状態によっては上手く動いてしまうことがあるのよ。
動いたり動かなかったりするってのが一番厄介。

ちなみに Java の条件変数の場合は，ロックをせずにシグナルを投げようとすると例外が発生するらしい。

そういえばシグナルを投げるメソッドには `Broadcast()` もあり，こちらは `Wait()` で一時停止している全ての goroutine が起こされる。
`Signal()` と `Broadcast()` のどちらを使うかは状況によると思うが，読書会主催の[柴田芳樹](https://connpass.com/user/yoshiki_shibata/)さんは対象の goroutine のうちどれが起きるのか明確に分からない場合は（ひとつだけ起こしたい場合でも）なるべく `Signal()` ではなく `Broadcast()` を使うべきだと言っておられた。

## 書き込み優先な RWMutex

[Go] では，共有資源に対して

1. 書き込みを行うスレッド（goroutine）はひとつだけ
2. 読み込み専用スレッドは複数ある

場合に読み込み側と書き込み側で優先度の異なるロックをかける [`sync`]`.RWMutex` 型ってのがある。
『[Go言語で学ぶ並行プログラミング]』の第4章では，標準の [`sync`]`.RWMutex` を使わずリーダー・ライター・ミューテックスを自作している。

ただし第4章で作ったのは読み込み優先の実装だったので，読み込み側が次々とロックを取ってしまうと，いつまで経っても書き込み側がロックを取れない事態になる[^ws1]。
この場合はデッドロックにもならないので，これはこれで厄介である。
これを解消するためには書き込み側を優先するロックにする必要がある。
第5章では，条件変数を用いて，自作リーダー・ライター・ミューテックスを書き込み優先に修正するくだりがある。

[^ws1]: リーダー・ライター・ミューテックスで書き込み側がいつまでもロックを獲得できない状態を writer-starvation と呼ぶ。『[Go言語で学ぶ並行プログラミング]』における “starvation” の定義は「他の貪欲な実行によって資源が長時間（または無期限に）利用できなくなり、共有資源へのアクセスを得ることができない状況を指します」としている。書き込み側が starvation になるから writer-starvation ということらしい。辞書的な “starvation” の意味は「食糧の不足によって栄養失調が続き、体調の維持が困難になっている状態」みたいな意味だそうな。継続的な飢餓状態ってことやね。

『[Go言語で学ぶ並行プログラミング]』は比較的簡単なサンプルコードで仕組みを紹介しているのがとてもいいと思う。

ちなみに標準の [`sync`]`.RWMutex` は書き込み優先である。

{{< fig-quote type="markdown" title="" link="https://pkg.go.dev/sync#RWMutex" lang="en" >}}
If any goroutine calls [`RWMutex.Lock`](https://pkg.go.dev/sync#RWMutex.Lock) while the lock is already held by one or more readers, concurrent calls to [`RWMutex.RLock`](https://pkg.go.dev/sync#RWMutex.RLock) will block until the writer has acquired (and released) the lock, to ensure that the lock eventually becomes available to the writer. Note that this prohibits recursive read-locking. A [`RWMutex.RLock`](https://pkg.go.dev/sync#RWMutex.RLock) cannot be upgraded into a [`RWMutex.Lock`](https://pkg.go.dev/sync#RWMutex.Lock), nor can a [`RWMutex.Lock`](https://pkg.go.dev/sync#RWMutex.Lock) be downgraded into a [`RWMutex.RLock`](https://pkg.go.dev/sync#RWMutex.RLock).
{{< /fig-quote >}}

だから第4章では標準を使わなかったんだねぇ。
ここで繋がったよ（笑）

## セマフォ

セマフォ（semaphore）は並行処理の最大数を制御する仕組みである。
たとえば計算資源をフルに使わせたくない場合に，ある時点の並行処理の最大数を（おそらくプロセッサ数以下に）抑えることで実現できる。

『[Go言語で学ぶ並行プログラミング]』では条件変数を使ってセマフォを自作している。
一方 [Go] の標準パッケージにはセマフォはないが [`golang.org/x/sync/semaphore`] パッケージを使うことでセマフォを扱うことができる。
こちらはほぼ標準パッケージに近いのでバンバン使っていいだろう。

この記事では [`golang.org/x/sync/semaphore`] パッケージに載っているサンプルコードを挙げておく。

{{<fig-quote class="nobox" type="markdown" title="semaphore package - golang.org/x/sync/semaphore - Go Packages" link="https://pkg.go.dev/golang.org/x/sync/semaphore" lang="en" >}}
```go {hl_lines=[22,30,36,45]}
package main

import (
    "context"
    "fmt"
    "log"
    "runtime"

    "golang.org/x/sync/semaphore"
)

// Example_workerPool demonstrates how to use a semaphore to limit the number of
// goroutines working on parallel tasks.
//
// This use of a semaphore mimics a typical “worker pool” pattern, but without
// the need to explicitly shut down idle workers when the work is done.
func main() {
    ctx := context.TODO()

    var (
        maxWorkers = runtime.GOMAXPROCS(0)
        sem        = semaphore.NewWeighted(int64(maxWorkers))
        out        = make([]int, 32)
    )

    // Compute the output using up to maxWorkers goroutines at a time.
    for i := range out {
        // When maxWorkers goroutines are in flight, Acquire blocks until one of the
        // workers finishes.
        if err := sem.Acquire(ctx, 1); err != nil {
            log.Printf("Failed to acquire semaphore: %v", err)
            break
        }

        go func(i int) {
            defer sem.Release(1)
            out[i] = collatzSteps(i + 1)
        }(i)
    }

    // Acquire all of the tokens to wait for any remaining workers to finish.
    //
    // If you are already waiting for the workers by some other means (such as an
    // errgroup.Group), you can omit this final Acquire call.
    if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
        log.Printf("Failed to acquire semaphore: %v", err)
    }

    fmt.Println(out)

}

// collatzSteps computes the number of steps to reach 1 under the Collatz
// conjecture. (See https://en.wikipedia.org/wiki/Collatz_conjecture.)
func collatzSteps(n int) (steps int) {
    if n <= 0 {
        panic("nonpositive input")
    }

    for ; n > 1; steps++ {
        if steps < 0 {
            panic("too many steps")
        }

        if n%2 == 0 {
            n /= 2
            continue
        }

        const maxInt = int(^uint(0) >> 1)
        if n > (maxInt-1)/3 {
            panic("overflow")
        }
        n = 3*n + 1
    }

    return steps
}
```
{{</fig-quote >}}

[コラッツ予想](https://ja.wikipedia.org/wiki/%E3%82%B3%E3%83%A9%E3%83%83%E3%83%84%E3%81%AE%E5%95%8F%E9%A1%8C "コラッツの問題 - Wikipedia")？ このサンプルでは，メイン以外の goroutine の間では共有資源の異なるアドレスにアクセスするためデータ競合（data race）は起きないことに注意。
データ競合が発生し得るのであればミューテックスを使ってアクセス制御を行う必要がある。

[`golang.org/x/sync/semaphore`] パッケージで実装されているセマフォは「重み付きセマフォ（weighted semaphore）」と呼ぶらしい。
`Acquire()` および `Release()` で獲得数・解放数を指定できる。
上のサンプルコードでは最後の `sem.Acquire(ctx, int64(maxWorkers))` でメイン以外の goroutine が稼働していないことを確認している（稼働していれば待ちになる）。
なお `Release()` 時にカウントがマイナスになると `panic` を吐くようだ。

## 今回はここまで

[次回](https://technical-book-reading-2.connpass.com/event/350009/ "第4回『Go言語で学ぶ並行プログラミング』オンライン読書会 - connpass")は第6章の「バリア（barrier）」の話からである。
楽しみ！

## ブックマーク

- [cutajarj/ConcurrentProgrammingWithGo: Listings from manning book](https://github.com/cutajarj/ConcurrentProgrammingWithGo) : 『[Go言語で学ぶ並行プログラミング]』に出てくるサンプルコード

[Go]: https://go.dev/
[`sync`]: https://pkg.go.dev/sync "sync package - sync - Go Packages"
[`golang.org/x/sync/semaphore`]: https://pkg.go.dev/golang.org/x/sync/semaphore "semaphore package - golang.org/x/sync/semaphore - Go Packages"

[Go言語で学ぶ並行プログラミング]: https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Go言語で学ぶ並行プログラミング　他言語にも適用できる原則とベストプラクティス impress top gearシリーズ | James Cutajar, 柴田 芳樹 | 工学 | Kindleストア | Amazon"
[第3回『Go言語で学ぶ並行プログラミング』オンライン読書会]: https://technical-book-reading-2.connpass.com/event/347110/ "第3回『Go言語で学ぶ並行プログラミング』オンライン読書会 - connpass"

## 参考図書

{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
