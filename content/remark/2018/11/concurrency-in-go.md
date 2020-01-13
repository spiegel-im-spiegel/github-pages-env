+++
title = "『Go 言語による並行処理』は Go 言語プログラマ必読書だろう"
date = "2018-11-03T09:50:43+09:00"
description = "並行処理プログラミングが難しいのは，デザイン・パターンの熟成がまだ若いことと，パターンの組み合わせの選択が複雑な点にあると思う。"
image = "/images/attention/go-logo_blue.png"
tags = [ "book", "golang", "programming", "concurrency", "goroutine", "channel", "context", "message-passing" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

やぁ！ ついに “[Concurrency in Go](https://katherine.cox-buday.com/concurrency-in-go/)” の邦訳版が登場したですよ。
めでたい！

[Go 言語]の特徴はいくつかあるが，大きなものとして「並行処理[^cm1]」を前提とした言語設計が挙げられる。

[^cm1]: 「同時に複数の動作が行われていること」を意味する言葉として「並行（concurrent）」と「並列（parallel）」が混同されることがあるが，プログラミングにおいては，書かれたコードの性質をを表すものを「並行」，動作するランタイム・プログラム（群）の性質を表すものを「並列」と呼んで区別している。両者は密接に関連しているが等価ではない。

並行処理では [CSP (Communicating Sequential Processes)](https://dl.acm.org/citation.cfm?doid=359576.359585) の成果が取り入れられている。
これが goroutine と channel である。

Goroutine は軽量スレッドなどと呼ばれることもあるが[^thrd1]，厳密には coroutine の一種である。
Main goroutine を含む各 goroutine はランタイム・プログラムに組み込まれたスケジューラによって制御される。
つまり実行時の並列処理の詳細はコードレベルでは隠蔽されているのだ。
故にプログラマはコード上の並行処理にのみ注力して記述すればいいことになる。
OS スレッドの仕様がどうなってるとか無駄にデカいスレッドをプールの中でどう使い回すかとか，生産性の欠片もないようなことは考えなくていいわけだ[^grt1]。

[^thrd1]: OS スレッドに比べて goroutine を駆動する（ランタイム・プログラムが展開する）グリーン・スレッドは非常に軽量らしい。初期のメモリ割当で数キロバイト程度で，実行時に自動的に増減するようだ。 Goroutine のリソースの概算については3.1章で言及されている。
[^grt1]: ある規模以上のシステムの場合は流量制限等をかける必要があるかも知れない。大規模システム開発での注意点については5章で言及されている。

Channel は入出力プリミティブで，これによって goroutine 間のメッセージ・パッシング（message passing）を実現している。
たとえば以下のような感じで記述する[^ch1]。

[^ch1]: 受信（`<-chan`）は単項演算子で，送信（`chan<-`）は送信構文（send statement）で記述する。

```go
package main

import "fmt"

//main goroutine
func main() {
    hello := make(chan interface{}) //create channel
    go func() { //sub goroutine
        hello <- "Hello, world!" //send
    }()
    fmt.Println(<-hello) //receive
}
```

（`go` 構文（[go statement](https://golang.org/ref/spec#Go_statements "The Go Programming Language Specification - The Go Programming Language")）でキックされた関数が sub goroutine となる。関数閉包（closure）になっている点に注目）

これらに加えて，同期をとるための [`sync`] パッケージやコールグラフの各枝をキャンセルするための [`context`] パッケージが標準ライブラリで用意されている。

メッセージ・パッシングを構成するか [`sync`] パッケージ等を用いて legacy なメモリアクセス同期を構成するか（あるいはそれらを組み合わせるか）の判断は難しいが，2.4章に大まかな指針が挙げられているので参考になるだろう。
ただし，今まで無理やりメモリアクセス同期で運用していたもののうちかなりのものを軽くて（スレッドセーフという意味で）安全なメッセージ・パッシングに置き換えられるのは確かである。

そうそう。
『[Go 言語による並行処理]』では [`context`] パッケージについてページを割いて紹介されているのだ[^cxt1]（4.12章）。
[`context`] パッケージをデザイン・パターンとして日本語できちんと紹介している Web 記事や書籍はあまり見かけないので，単純に嬉しい。

[^cxt1]: [`context`] パッケージはバージョン 1.7 から標準ライブラリに組み込まれたため，古い参考書には載っていない。なお [`context`] パッケージにはキャンセルの伝搬以外にもコールグラフ間でデータを受け渡す機能もある。もっともこちらは濫用するとヤバいけど。

並行処理プログラミングが難しいのは，デザイン・パターン（4章で言及されている）の熟成がまだ若いこと，パターンの組み合わせ最適化が複雑なこと，もっと言うなら並行処理の設計は context driven であること，にあると思う（だからこそ [`context`] パッケージは秀逸なのよ）。
デザイン・パターンは（数学の公式と同じく）思考のショートカットなので，そのパターンを構成することの意味を分かった上で適用しないと失敗することが多いんじゃないだろうか。

この本を読んで「よっしゃ，明日から立派な goroutine 使いだ！」とはならないと思うけど，有象無象なコピペ・プログラマじゃなく，きちんと [Go 言語]のプログラミングを勉強したいのであれば，この本は必読書になると思う。
少なくとも（立ち読みででも）2章までは熟読すべき。

ところで余談だが，補遺Bは日本語版オリジナルの章だそうだけど，その中に [Go 2 ドラフト](https://blog.golang.org/go2draft "Go 2 Draft Designs - The Go Blog")で提案されている Generics についてしれっと書かれている。

{{< fig-youtube id="6wIP3rO6On8" title="Go 2 Drafts Announcement - YouTube" >}}

次期 [Go 言語]に搭載される（かもしれない） Generics についての解説は以下を参考にどうぞ。

- [Go 2のgenerics/contract簡易まとめ](https://qiita.com/lufia/items/242d25e8c93d88e22a2e)

## ブックマーク

- [Big Sky :: 書評「Go言語による並行処理」](https://mattn.kaoriya.net/software/lang/go/concurrency-in-go.htm)
- [Big Sky :: golang の channel を使ったテクニックあれこれ](http://mattn.kaoriya.net/software/lang/go/20160706165757.htm)
- [Big Sky :: Go 言語の非同期パターン](https://mattn.kaoriya.net/software/lang/go/20180531104907.htm)
- [Go言語の並行性を映像化する | プログラミング | POSTD](http://postd.cc/go_concurrency_visualize/)
- [Go言語の並行処理デザインパターン by Rob Pike 前編 - Qiita](http://qiita.com/tfutada/items/a289628d8b2d0af6152d)
    - [Go言語の並行処理デザインパターン by Rob Pike 後編 - Qiita](http://qiita.com/tfutada/items/dc8db894ac270a79ef2b)

- [time.Ticker で遊ぶ]({{< ref "/golang/ticker.md" >}})
- [『Go 言語による並行処理』はEブック版がオススメ]({{< ref "/remark/2020/01/concurrency-in-go-digital.md" >}})

[Go 言語による並行処理]: https://www.amazon.co.jp/exec/obidos/ASIN/4873118468/baldandersinf-22/ "Go言語による並行処理 | Katherine Cox-Buday, 山口 能迪 |本 | 通販 | Amazon"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`sync`]: https://golang.org/pkg/sync/ "sync - The Go Programming Language"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"

## 参考図書

{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
