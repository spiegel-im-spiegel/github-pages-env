+++
title = "『Go言語で学ぶ並行プログラミング』の練習問題より"
date =  "2025-04-17T19:47:01+09:00"
description = "今回は読書会じゃなくて Bluesky からのネタ"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "book", "concurrency" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は読書会じゃなくて Bluesky からのネタ[^x1]。

[^x1]: {{< emoji "X" >}} の TL はホンマにウザいので，エンジニアの方々は Bluesky に来て欲しい，マジで。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:fuzh3iojtr6xbixqug5izlqf/app.bsky.feed.post/3lmwu2pptqc2n" data-bluesky-cid="bafyreifc5mlf7ywtra4ur3ml4733vqtnp5kol22ibmjlnuhbzoc4wuwbta" data-bluesky-embed-color-mode="system"><p lang="ja">Go言語で学ぶ並行プログラミングの4章の hello world みたいな練習問題、すでにめっちゃ難しいんだけど、この答えの 26行目は lock とらないでいいの？ (4.3 練習問題 1)
github.com/cutajarj/Con...<br><br><a href="https://bsky.app/profile/did:plc:fuzh3iojtr6xbixqug5izlqf/post/3lmwu2pptqc2n?ref_src=embed">[image or embed]</a></p>&mdash; ikawaha (<a href="https://bsky.app/profile/did:plc:fuzh3iojtr6xbixqug5izlqf?ref_src=embed">@ikawaha.bsky.social</a>) <a href="https://bsky.app/profile/did:plc:fuzh3iojtr6xbixqug5izlqf/post/3lmwu2pptqc2n?ref_src=embed">April 16, 2025 at 11:57 PM</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

「4.3 練習問題」の1というのがこれ。

{{< fig-quote type="markdown" title="『Go言語で学ぶ並行プログラミング』4.3節より" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
リスト 4.15（もともとは第 3 章から）は、共有変数へのアクセスを保護するためのミューテックスを使っていません。これは悪い習慣です。このプログラムを変更して、共有変数 seconds へのアクセスをミューテックスで保護するようにしてください。ヒント：変数をコピーする必要があるかもしれません。
{{< /fig-quote >}}

んで「リスト 4.15」というのがこちら：

```go
package main

import (
    "fmt"
    "time"
)
func countdown(seconds *int) {
    for *seconds > 0 {
        time.Sleep(1 * time.Second)
        *seconds -= 1
    }
}

func main() {
    count := 5
    go countdown(&count)
    for count > 0 {
        time.Sleep(500 * time.Millisecond)
        fmt.Println(count)
    }
}
```

簡単なカウントダウンのプログラムだけど，カウントダウンを goroutine で実行しているのがポイント。
`main()` の `count` 変数が2つの goroutine で使われているの（`countdown()` 関数側は `seconds`）がお分かりだろうか。

で，この問題の回答は [GitHub リポジトリ](https://github.com/cutajarj/ConcurrentProgrammingWithGo "cutajarj/ConcurrentProgrammingWithGo: Listings from manning book")で公開されている。
こんな感じ：

```go {hl_lines=[26]}
package main

import (
    "fmt"
    "sync"
    "time"
)

func countdown(seconds *int, mutex *sync.Mutex) {
    mutex.Lock()
    remaining := *seconds
    mutex.Unlock()
    for remaining > 0 {
        time.Sleep(1 * time.Second)
        mutex.Lock()
        *seconds -= 1
        remaining = *seconds
        mutex.Unlock()
    }
}

func main() {
    mutex := sync.Mutex{}
    count := 5
    go countdown(&count, &mutex)
    remaining := count
    for remaining > 0 {
        time.Sleep(500 * time.Millisecond)
        mutex.Lock()
        fmt.Println(count)
        remaining = count
        mutex.Unlock()
    }
}
```

上で強調している行はロック取ってないよね，というのが最初のポストの内容というわけだ。

ちょっとややこしいのだが，もともとこの練習問題の回答としては

```go {hl_lines=[5]}
func main() {
    mutex := sync.Mutex{}
    count := 5
    go countdown(&count, &mutex)
    for count  > 0 {
        time.Sleep(500 * time.Millisecond)
        mutex.Lock()
        fmt.Println(count)
        mutex.Unlock()
    }
}
```

が提示されていたらしい（`countdown()` 関数は省略して `main()` 関数のみ挙げている）。
これだと `for` ループの条件文で `count` を直接参照しているので拙いよね，という[指摘]があったのだが，この指摘に対応しようとして失敗しているように見える。

おそらく `main()` 関数を修正するのであれば以下が妥当だと思う。

```go {hl_lines=[4,7]}
func main() {
    mutex := sync.Mutex{}
    count := 5
    remaining := count
    go countdown(&count, &mutex)
    for remaining > 0 {
        fmt.Println(remaining)
        time.Sleep(500 * time.Millisecond)
        mutex.Lock()
        remaining = count
        mutex.Unlock()
    }
}
```

まず `remaining` 変数の初期化宣言を goroutine 起動前に行う。
さらに `for` ループ内の [`fmt`]`.Println()` を `remaining` 変数を参照するよう変更する（こうすることで [`fmt`]`.Println()` をロックの外に出せる。なお，上のコードではカウント 0 を出力しないよう [`fmt`]`.Println()` 関数の位置を変えている）。

といったところだろうか。

問題は最初の回答コードでも，その後の不完全修正コードでも，たぶん支障なく動いてしまうところなんだよね。
だってカウントダウンさせてる int 型の共有変数を参照してるだけだもん。
参照時に多少の不整合が起きても，おそらく見た目では分からない。

著者の方も[指摘]に対して

{{< fig-quote type="markdown" title="GitHub Issue #2" link="https://github.com/cutajarj/ConcurrentProgrammingWithGo/issues/2" lang="en" >}}
Thank you for spotting this. This is proof that race conditions are easy to miss!
{{< /fig-quote >}}

と返しておられる通り，並行処理のデバッグがいかに難しいか分かる。
ただ，共有変数 `count` へアクセスする際に「不変式が真[^i1]」であることをきちんと保証するためには，この実装をサボってはいけない。

[^i1]: 不変式の話は『[プログラミング言語Go]』に出てくる。また『[Go言語で学ぶ並行プログラミング]』の「訳者あとがき」でも『[プログラミング言語Go]』の内容を紹介する形で説明がある。でも『[プログラミング言語Go]』は絶版状態なんだよなぁ `orz`

[Go] に詳しい方なら「`for` 文で回さなくても channel を使えばいいし，なんなら [`atomic`] パッケージも使えるぢゃん」と思うだろう（そしてそれは正しい）が『[Go言語で学ぶ並行プログラミング]』では4章でようやく [`sync`]`.Mutex` が登場したところで， channel は7章， [`atomic`] パッケージは12章で登場する。
あしからず。
なお『[Go言語で学ぶ並行プログラミング]』本編には言及が見当たらないが [`sync`]`.Mutex` は再入不可である[^r1]。

[^r1]: 再入云々については『[Go言語で学ぶ並行プログラミング]』の「訳者あとがき」に解説がある。たとえば Java のミューテックスは再入可能なため，その辺の差異について認識しておく必要がある。

[Go]: https://go.dev/
[`fmt`]: https://pkg.go.dev/fmt "fmt package - fmt - Go Packages"
[`sync`]: https://pkg.go.dev/sync "sync package - sync - Go Packages"
[`atomic`]: https://pkg.go.dev/sync/atomic "atomic package - sync/atomic - Go Packages"

[指摘]: https://github.com/cutajarj/ConcurrentProgrammingWithGo/issues/2 "Ex. 4.1 · Issue #2 · cutajarj/ConcurrentProgrammingWithGo"
[Go言語で学ぶ並行プログラミング]: https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: Go言語で学ぶ並行プログラミング 他言語にも適用できる原則とベストプラクティス impress top gearシリーズ eBook : James Cutajar, 柴田 芳樹: Kindleストア"
[プログラミング言語Go]: https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES) | Alan A.A. Donovan, Brian W. Kernighan, 柴田 芳樹 |本 | 通販 | Amazon"

## 参考図書

{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
