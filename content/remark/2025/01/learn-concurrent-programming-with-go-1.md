+++
title = "goroutine はグリーンスレッドではない（『Go言語で学ぶ並行プログラミング』読書会1回目）"
date =  "2025-01-25T23:27:33+09:00"
description = "「並行処理」と「並列処理」を使って文を書け"
image = "/images/attention/go-logo_blue.png"
tags = [ "book", "golang", "engineering", "programming", "concurrency" ]
pageType = "text"

[canonical]
  url = "/golang/learn-concurrent-programming-with-go-1/"

[scripts]
  mathjax = false
  mermaidjs = false
+++

{{< div-box type="markdown" >}}
[プログラミング言語 Go]({{< rlnk "/golang/" >}})」セクションの[同名記事]({{< ref "/golang/learn-concurrent-programming-with-go-1.md" >}})に移動しました。

<script>
location.href = "/golang/learn-concurrent-programming-with-go-1/";
</script>
{{< /div-box >}}

今日は「[第1回『Go言語で学ぶ並行プログラミング』オンライン読書会]」だった。

版元から PDF 版を[購入](https://book.impress.co.jp/books/1123101144 "Go言語で学ぶ並行プログラミング　他言語にも適用できる原則とベストプラクティス - インプレスブックス")した。
インプレス社って絶版でもないのにいつの間にか Kindle 版を引っ込めたりするそうで，デジタル版を買うなら PDF 版を買うのがオススメらしい。
まぁ，技術参考書は PDF 版のほうが取り回ししやすいからな。
紙はかさばるし。

事前にざっと斜め読みしたのだが，解説が丁寧という印象。
タイトル通り [Go] 言語を前提に書かれているが C/C++ や Java などの構文型のプログラミング言語に慣れているのなら難しくないと思う。

まずは1章から順に読むのがおすすめ。
「並行処理，[完全に理解した](https://zenn.dev/activecore/articles/9862409de182c7 "エンジニア完全に理解した")」という人も復習のつもりで順に読んでいくのがいいだろう。

## 並行と並列

以前，読書感想文で『[Go言語による並行処理]』を[大絶賛]({{< ref "/remark/2018/11/concurrency-in-go.md" >}} "『Go 言語による並行処理』は Go 言語プログラマ必読書だろう")したのだが，その中の「2.1 並行性と並列性の違い」の中で

{{< fig-quote type="markdown" title="『Go言語による並行処理』 p.23" link="https://www.amazon.co.jp/dp/4873118468?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
並行性はコードの性質を指し、並列性は動作しているプログラムの性質を指します。
{{< /fig-quote >}}

と書かれている。
『[Go言語で学ぶ並行プログラミング]』では更に具体的に

{{< fig-quote type="markdown" title="『Go言語で学ぶ並行プログラミング』 p.39" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
並行性とは、多くのタスクを同時にどのように実行するかを計画する（planning）ことです。並列性とは、多くのタスクを同時に実行する（performing）ことです。
{{< /fig-quote >}}

と書かれている。
たとえば

{{< div-box type="markdown" >}}
*問. 「並行処理」と「並列処理」を使って文を書け*
{{< /div-box >}}

みたいな問題があったら

{{< div-box type="markdown" >}}
並行処理を組み込んだプログラムを書いて実行したところ正しく並列処理が行われた。
{{< /div-box >}}

みたいに回答するのがいいだろうか（笑）

今回の[読書会][第1回『Go言語で学ぶ並行プログラミング』オンライン読書会]は第3章にかかった辺りまでだったが，今のところ『[Go言語による並行処理]』よりは『[Go言語で学ぶ並行プログラミング]』のほうがオススメな感じ。
まぁ，あとから出た本だからねぇ。

## goroutine はグリーンスレッドではない

並行処理の仕組みとしてカーネルが用意している機能にはプロセスとスレッドがある。

{{< fig-quote type="markdown" title="『Go言語で学ぶ並行プログラミング』 p.20" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
オペレーティングシステムの各プロセスは、他のプロセスから隔離された独自のメモリ空間を持っています。通常、プロセスは独立して動作し、他のプロセスとの相互作用は最小限です。プロセスは、多くの資源を消費する代償として、隔離を提供します。たとえば、あるプロセスがエラーでクラッシュしても、そのプロセスは独自のメモリ空間を持っているため、他のプロセスに影響を与えることはありません。この隔離の欠点は、多くのメモリを消費してしまうことです。さらに、メモリ領域やその他のシステム資源を確保する必要があるため、プロセスの起動には（スレッドに比べて）少し時間がかかります。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『Go言語で学ぶ並行プログラミング』 p.25" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
新たなスレッドを作成する場合、オペレーティングシステムが作成する必要があるのは、スタック領域、レジスタ、プログラムカウンタを管理するのに十分な資源だけです。新たなスレッドは、同じプロセスのコンテキスト内で実行されます。対照的に、新たなプロセスを作成する場合、オペレーティングシステムはそれに完全に新たなメモリ空間を割り当てる必要があります。この理由から、スレッドはプロセスよりもはるかに軽量であり、システムが資源を使い果たすまでに、プロセスよりも多くのスレッドを作成することが通常可能です。さらに、新たに割り当てる資源がかなり少ないため、スレッドの起動はプロセスの起動よりもはるかに速いです。

[...]

ただし、この追加の性能には、ある代償が伴います。同じメモリ空間で作業するということは、プロセスが提供する隔離が得られないことを意味します。そのため、あるスレッドが他のスレッドの作業に干渉して被害を与える可能性があります。
{{< /fig-quote >}}

各スレッドがどの CPU コアに割り当てられるかはカーネルが決定する。
このため上述のスレッドのことを特に「カーネルレベルスレッド」と呼ぶ。

カーネルレベルスレッドがあるということはユーザー空間で完全にアプリケーションが制御する「ユーザーレベルスレッド」もあるわけだ。

{{< fig-quote type="markdown" title="『Go言語で学ぶ並行プログラミング』 p.33-34" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
ユーザーレベルスレッドの主な利点は性能です。ユーザーレベルスレッドのコンテキストスイッチは、カーネルレベルスレッドのコンテキストスイッチよりも高速です。なぜなら、カーネルレベルのコンテキストスイッチでは、オペレーティングシステムが介入して次に実行するスレッドを選択する必要があるからです。カーネルを呼び出さずに実行を切り替えることができれば、実行中のプロセスは、キャッシュをフラッシュして処理速度を低下させる必要がなく、CPU を保持し続けられます。

ユーザーレベルスレッドを使うことの不都合な点は、ブロッキング I/O 呼び出しを行うコードを実行するときに現れます。ファイルからの読み込みが必要な状況を考えてみましょう。オペレーティングシステムはプロセスを 1 つの実行のスレッドと見なすため、ユーザーレベルスレッドがこのブロッキング読み込み呼び出しを実行すると、プロセス全体がスケジュールから外されます。同じプロセス内に他のユーザーレベルスレッドが存在する場合、読み込み操作が完了するまで、それらのスレッドは実行されません。
{{< /fig-quote >}}

で， [Go] の goroutine はどれやねん？ という話だが，実はカーネルレベルスレッドとユーザーレベルスレッドを組み合わせたハイブリッドな構成となっている。

{{< fig-quote type="markdown" title="『Go言語で学ぶ並行プログラミング』 p.34-35" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
Go はハイブリッドシステムを採用することで、ユーザーレベルスレッドの優れた性能を提供し、不都合な点がほとんどありません。これは、カーネルレベルスレッドの集まりを使い、それぞれがゴルーチンのキューを管理することで実現しています。複数のカーネルレベルスレッドがあるので、複数のプロセッサが利用可能な場合、複数のプロセッサを利用できます。
{{< /fig-quote >}}

これを「M:N スレッディングモデル」と呼ぶ。

{{< fig-quote type="markdown" title="『Go言語で学ぶ並行プログラミング』 p.35" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
Go がゴルーチンに使っているシステムは、M:N スレッディングモデル（threading model）と呼ばれることがあります。M:N とするのは M 個のユーザーレベルスレッド（ゴルーチン）がN 個のカーネルレベルスレッドにマッピングされている場合です。これは、通常のユーザーレベルスレッドとは対照的です。通常のユーザーレベルスレッドは、N:1 スレッディングモデルと呼ばれ、1 つのカーネルレベルスレッドに対して N 個のユーザーレベルスレッドという意味です。M:N モデルのランタイムを実装するには、カーネルレベルスレッドの集まり上でユーザーレベルスレッドを移動させてバランスを取るための多くの技法が必要なため、他のモデルよりも複雑になります。
{{< /fig-quote >}}

一方で「グリーンスレッド」という言葉があるのだが，これは Java の用語らしい。

{{< fig-quote type="markdown" title="『Go言語で学ぶ並行プログラミング』 p.34" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
グリーンスレッド（green thread）という用語は、プログラミング言語 Java のバージョン 1.1で生まれました 。Java のオリジナルのグリーンスレッドは、ユーザーレベルスレッドの実装でした。単一のコア上でのみ実行され、JVM によって完全に管理されていました 。Javaバージョン 1.3 では、グリーンスレッドはカーネルレベルスレッドに取って代わられました。それ以来、多くの開発者がユーザーレベルスレッドの他の実装を指すために用語「グリーンスレッド」を使っています。後述するように、Go のランタイムはゴルーチンが複数の CPU を十分に活用できるようにしているので、 Go のゴルーチンをグリーンスレッドと呼ぶのはおそらく不正確です。
{{< /fig-quote >}}

へー。
よその記事で goroutine をグリーンスレッドと書いてる気がするなぁ。
直すの面倒だし，過去記事は放置で，これから気をつけることにしよう。

ちなみに Java 21 からは [Go] に似たスレッドモデルが正式に導入されていて，これは「仮想スレッド（virtual thread）」と呼ぶそうな。

## Go のメモリモデル

翻訳者で[読書会][第1回『Go言語で学ぶ並行プログラミング』オンライン読書会]の主催者である[柴田芳樹](https://note.com/yoshiki_shibata "柴田 芳樹｜note")さんが今回おっしゃっていたが（実は訳者あとがきでも言及されている），『[Go言語で学ぶ並行プログラミング]』にはメモリモデルについての解説がない。
並行処理とメモリ共有は密接な関係があるのでメモリモデルについても解説があればよかったのに，という感じ。

[Go] のメモリモデルについては公式のドキュメントがある。

- [The Go Memory Model - The Go Programming Language](https://go.dev/ref/mem)

このドキュメントの冒頭に

{{< fig-quote type="markdown" title="The Go Memory Model" link="https://go.dev/ref/mem" lang="en" >}}
If you must read the rest of this document to understand the behavior of your program, you are being too clever.
{{< /fig-quote >}}

とか皮肉（？）な文章が書かれていて笑っちまったよ。

それはともかく，中身をちょろんと紹介すると，たとえば

```go
var a string

func f() {
    print(a)
}

func hello() {
    a = "hello, world"
    go f()
}
```

というコードがあったとき “`hello, world`” と出力されることが保証されるというもの。

{{< fig-quote type="markdown" title="The Go Memory Model" link="https://go.dev/ref/mem" lang="en" >}}
The go statement that starts a new goroutine is synchronized before the start of the goroutine's execution.
{{< /fig-quote >}}

「当たり前やんけ！」と思うかもしれないが `hello()` 関数と `f()` 関数が異なる CPU コアで並列に実行される場合，この「保証」はかなり重要である。

マルチコアプロセッサでは（システムバスを通じて）メモリに直接アクセスするのではなく，メモリキャッシュを挟んだ間接的なアクセスになる。
この際に複数の CPU コアとメモリとの間にデータ競合が起きないよう「キャッシュ・コヒーレンシー・プロトコル（cache-coherency protocols）」が走るらしい。

このプロトコルが正しく働いて CPU レベルでのデータ競合が起きないことが保証できないと，並行処理下で「何も信用できない」ことになってしまう。

キャッシュ・コヒーレンシー・プロトコル周りの話を始めるとそれだけで本が書けるそうで，実際に紹介してもらったが，かなりの分量があって理解するのも大変とのこと。
たぶん『[Go言語による並行処理]』にメモリモデルの解説を入れると分量が1.5倍とか2倍とかになるんじゃないのかな（笑）

## というわけで

次回以降の読書会も楽しみである。
記事にするかどうかは不明。

## ブックマーク

- [GitHub - cutajarj/ConcurrentProgrammingWithGo: Listings from manning book](https://github.com/cutajarj/ConcurrentProgrammingWithGo)

[Go]: https://go.dev/
[Go言語で学ぶ並行プログラミング]: https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Go言語で学ぶ並行プログラミング　他言語にも適用できる原則とベストプラクティス impress top gearシリーズ | James Cutajar, 柴田 芳樹 | 工学 | Kindleストア | Amazon"
[Go言語による並行処理]: https://www.amazon.co.jp/dp/4873118468?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Go言語による並行処理 | Katherine Cox-Buday, 山口 能迪 |本 | 通販 | Amazon"
[第1回『Go言語で学ぶ並行プログラミング』オンライン読書会]: https://technical-book-reading-2.connpass.com/event/337562/ "第1回『Go言語で学ぶ並行プログラミング』オンライン読書会 - connpass"

## 参考図書

{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B0CX1NVW3F" %}} <!-- へぇ～ボタン へーボタン -->

## 作業中の BGV (メン限配信以外)

- [【MV】また、おかえり。／猫又おかゆ　ノベルゲーム『おかゆにゅ～～む！』OPテーマ - YouTube](https://www.youtube.com/watch?v=kFH-KEZGYRk)
- [ファッとして桃源郷 - 尾丸ポルカ(cover) - YouTube](https://www.youtube.com/watch?v=a9kO512t5Vs)
- [【民俗学 / 解説】蛇と関係しているってほんと？七福神「弁財天」の民俗学【VTuber/ #諸星めぐる 】 - YouTube](https://www.youtube.com/watch?v=DHAH4NBz8c4)
- [【 #大猫百桜 】俺たちの大富豪はこれからだ🤝🔥【 ホロライブ 】 - YouTube](https://www.youtube.com/watch?v=cCI8YBhPsdo)
