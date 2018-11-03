+++
title = "『Go 言語による並行処理』は Go 言語プログラマ必読書だろう"
date = "2018-11-03T09:50:43+09:00"
description = "並行処理プログラミングが難しいのは，デザイン・パターンの熟成がまだ若いことと，パターンの組み合わせの選択が複雑な点にあると思う。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "golang", "programming", "concurrency", "goroutine", "channel", "context" ]

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

やぁ！ ついに “Concurrency in Go” の邦訳版が登場したですよ。
めでたい！

[Go 言語]の特徴はいくつかあるが，大きなものとして「並行処理[^cm1]」を前提とした言語設計が挙げられる。

[^cm1]: 「同時に複数の動作が行われていること」を意味する言葉として「並行（concurrent）」と「並列（parallel）」が混同されることがあるが，プログラミングにおいては，書かれたコードの性質をを表すものを「並行」，動作するランタイム・プログラムの性質を表すものを「並列」と呼んで区別している。

並行処理では [CSP (Communicating Sequential Processes)](https://dl.acm.org/citation.cfm?doid=359576.359585) の成果が取り入れられている。
これが goroutine と channel である。

goroutine は軽量スレッドなどと呼ばれることもあるが[^thrd1]，厳密には coroutine の一種である。
Main goroutine を含む各 goroutine はランタイム・プログラムに組み込まれたスケジューラによって制御されている。
つまり実行時の並列処理の詳細はコードレベルでは隠蔽されているわけだ。
これによりプログラマはコード上の並行処理にのみ注力して記述すればいいことになる。
OS スレッドの仕様がどうなってるとかスレッドプールをどう組むとか生産性の欠片もないようなことは考えなくていいわけだ[^grt1]。

[^thrd1]: OS スレッドに比べて goroutine を駆動する（ランタイム・プログラムが展開する）グリーン・スレッドは非常に軽量らしい。
[^grt1]: ある規模以上のシステムの場合は流量制限等をかける必要があるかも知れない。大規模システム開発での注意点については5章で言及されている。

channel は入出力プリミティブで[^ch1]，これによって goroutine 間のメッセージ・パッシングを実現している。

[^ch1]: 受信（`<-chan`）は単項演算子で，送信（`chan<-`）は送信構文（send statement）で記述する。

```go
package main

import "fmt"

func main() {
    hello := make(chan interface{})
    go func() {
        hello <- "Hello, world!"
    }()
    fmt.Println(<-hello)
}
```


これらに加えて，同期をとるための [`sync`] パッケージやコールグラフの各枝をキャンセルするための [`context`] パッケージが標準ライブラリで用意されている。

そうそう。
『[Go 言語による並行処理]』では [`context`] パッケージについてページを割いて紹介されているのよ[^cxt1]。
[`context`] パッケージをデザイン・パターンとして日本語できちんと紹介している Web 記事や書籍はあまり見かけないので，単純に嬉しい。

[^cxt1]: [`context`] パッケージはバージョン 1.7 から標準ライブラリに組み込まれたため，古い参考書には載っていない。

並行処理プログラミングが難しいのは，デザイン・パターンの熟成がまだ若いことと，パターンの組み合わせの選択が複雑な点にあると思う。
デザイン・パターンは（数学の公式と同じく）思考のショートカットなので，そのパターンを構成することの意味を分かった上で適用しないと失敗することが多いんじゃないだろうか。

余談だが，補遺Bは日本語版オリジナルの章だそうだけど，その中に [Go 2 ドラフト](https://blog.golang.org/go2draft "Go 2 Draft Designs - The Go Blog")で提案されている Generics についてしれっと書かれている。

{{< fig-youtube id="6wIP3rO6On8" title="Go 2 Drafts Announcement - YouTube" >}}

この本を読んで「よっしゃ，明日から立派な goroutine 使いだ！」とはならないと思うけど，有象無象なコピペ・プログラマじゃなく，きちんと [Go 言語]のプログラミングを勉強したいのであれば，この本は必読書になると思う。

## ブックマーク

- [Big Sky :: 書評「Go言語による並行処理」](https://mattn.kaoriya.net/software/lang/go/concurrency-in-go.htm)
- [Big Sky :: golang の channel を使ったテクニックあれこれ](http://mattn.kaoriya.net/software/lang/go/20160706165757.htm)
- [Big Sky :: Go 言語の非同期パターン](https://mattn.kaoriya.net/software/lang/go/20180531104907.htm)
- [Go言語の並行性を映像化する | プログラミング | POSTD](http://postd.cc/go_concurrency_visualize/)
- [Go言語の並行処理デザインパターン by Rob Pike 前編 - Qiita](http://qiita.com/tfutada/items/a289628d8b2d0af6152d)
    - [Go言語の並行処理デザインパターン by Rob Pike 後編 - Qiita](http://qiita.com/tfutada/items/dc8db894ac270a79ef2b)
- [Go 2のgenerics/contract簡易まとめ](https://qiita.com/lufia/items/242d25e8c93d88e22a2e)

- [time.Ticker で遊ぶ]({{< ref "/golang/ticker.md" >}})

[Go 言語による並行処理]: https://www.amazon.co.jp/exec/obidos/ASIN/4873118468/baldandersinf-22/ "Go言語による並行処理 | Katherine Cox-Buday, 山口 能迪 |本 | 通販 | Amazon"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`sync`]: https://golang.org/pkg/sync/ "sync - The Go Programming Language"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4873118468/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/51pUKQajnaL._SL160_.jpg" width="125" height="160" alt="Go言語による並行処理"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873118468/baldandersinf-22">Go言語による並行処理</a></dt>
    <dd>Katherine Cox-Buday</dd>
    <dd>オライリージャパン</dd>
    <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
    </span></dd>
  </dl>
  <p class="description">この本を読んで「よっしゃ，明日から立派な goroutine 使いだ！」とはならないと思うけど，有象無象なコピペ・プログラマじゃなく，きちんと Go 言語のプログラミングを勉強したいのであれば，この本は必読書になると思う。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.11.3</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" height="160" alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)"></a></div>
    <dl class="fn">
      <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
      <dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
      <dd>丸善出版</dd>
      <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
      </span></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.10.19</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>
