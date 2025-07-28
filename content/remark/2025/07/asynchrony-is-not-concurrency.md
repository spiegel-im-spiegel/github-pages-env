+++
title = "「並行」と「並列」と「非同期」"
date =  "2025-07-28T11:24:04+09:00"
description = "「非同期」と「並行」は違うよね，という話は確かにあまり聞かない"
image = "/images/attention/kitten.jpg"
tags = [ "concurrency", "engineering", "programming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

面白い記事を見かけたので

- [Asynchrony is not Concurrency | Loris Cro's Blog](https://kristoff.it/blog/asynchrony-is-not-concurrency/)

今回はこの記事を起点につらつらと書いてみる。

## 「並行」と「並列」と「非同期」

プログラミングおよびシステム設計に於いて「並行」と「並列」は意味が違うよ，というのはよく言われるし，私も Go 言語関連の話で紹介したことがある。
その一方で「非同期」と「並行」は違うよね，という話は確かにあまり聞かない。
先程挙げた記事では

{{< fig-quote type="markdown" title="Asynchrony is not Concurrency" link="https://kristoff.it/blog/asynchrony-is-not-concurrency/" lang="en" >}}
**Asynchrony**: the possibility for tasks to run out of order and still be correct.

**Concurrency**: the ability of a system to progress multiple tasks at a time, be it via parallelism or task switching.

**Parallelism**: the ability of a system to execute more than one task simultaneously at the physical level.
{{< /fig-quote >}}

つまり「並行」と「並列」は能力（ability）だが「非同期」は実行順序の可能性（possibility）であると説明されている。
あるいは順序の自由度の問題と言い換えたらいいだろうか。

ここでプログラミング言語 [Zig] のコードが登場する。
たとえば [Zig] による非同期処理は `io.async` を使って以下のように記述するそうなのだが

{{< fig-quote class="nobox" type="markdown" title="Asynchrony is not Concurrency" link="https://kristoff.it/blog/asynchrony-is-not-concurrency/" lang="en" >}}
```zig
// assume that server.listen has already been called
io.async(Server.accept, .{server, io});
io.async(Client.connect, .{client, io});
```
{{< /fig-quote >}}

`io.async` はシングルスレッドのブロッキングモードで動作するため，上のコード例の場合，一方の処理が他方の処理でブロックされてしまい並列処理にはならないらしい。
この挙動はユーザレベルスレッド[^gt1] の特性に近い（[Zig] の実装が実際にどうなっているかは知らないが）。

[^gt1]: “[Asynchrony is not Concurrency](https://kristoff.it/blog/asynchrony-is-not-concurrency/ "Asynchrony is not Concurrency | Loris Cro's Blog")” ではユーザレベルスレッドをグリーンスレッドと呼んでいるように見える。なおグリーンスレッドは元々 Java の用語で，しかもバージョンごとに意味が変遷しているらしいので，この記事では「グリーンスレッド」という用語は使わないことにする。詳しくは拙文「[goroutine はグリーンスレッドではない]」を参考にどうぞ。

{{< fig-quote type="markdown" title="『Go言語で学ぶ並行プログラミング』 p.33-34" link="https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
ユーザーレベルスレッドの主な利点は性能です。ユーザーレベルスレッドのコンテキストスイッチは、カーネルレベルスレッドのコンテキストスイッチよりも高速です。なぜなら、カーネルレベルのコンテキストスイッチでは、オペレーティングシステムが介入して次に実行するスレッドを選択する必要があるからです。カーネルを呼び出さずに実行を切り替えることができれば、実行中のプロセスは、キャッシュをフラッシュして処理速度を低下させる必要がなく、CPU を保持し続けられます。

ユーザーレベルスレッドを使うことの不都合な点は、ブロッキング I/O 呼び出しを行うコードを実行するときに現れます。ファイルからの読み込みが必要な状況を考えてみましょう。オペレーティングシステムはプロセスを 1 つの実行のスレッドと見なすため、ユーザーレベルスレッドがこのブロッキング読み込み呼び出しを実行すると、プロセス全体がスケジュールから外されます。同じプロセス内に他のユーザーレベルスレッドが存在する場合、読み込み操作が完了するまで、それらのスレッドは実行されません。
{{< /fig-quote >}}

正しく並行処理として記述したいなら

{{< fig-quote class="nobox" type="markdown" title="Asynchrony is not Concurrency" link="https://kristoff.it/blog/asynchrony-is-not-concurrency/" lang="en" >}}
```zig
try io.asyncConcurrent(Server.accept, .{server, io});
io.async(Client.connect, .{client, io});
```
{{< /fig-quote >}}

という感じにコードで明示する必要があるそうな。

拙文「[goroutine はグリーンスレッドではない]」でも書いたが [Go] の goroutine はカーネルレベルスレッドとユーザレベルスレッドのハイブリッドになっていて，非同期かつ並行な処理を記述することが出来る。
なので非同期と並行を分離して考えることはないし，その必要もない（そのぶんランタイム側で制御するスケジューリングが複雑になるんだけど）。
強いて言うなら [Go] の[メモリモデル](https://go.dev/ref/mem "The Go Memory Model - The Go Programming Language")について注意が必要といったところか。

[Zig] のように非同期処理と平行処理を明示的に分離して記述するのは面白いと思う[^r1]。
より低レベルな処理に気を配らないといけないので，コード設計が面倒くさいっちゃあ面倒くさいんだけどね。

[^r1]: “[Async Rust Is A Bad Language](https://bitbashing.io/async-rust.html)” では Rust における非同期処理の問題点が挙げられていて，最後に “Maybe Rust isn’t a good tool for massively concurrent, userspace software” とまで書かれていて笑ってしまった。みんな苦労してるんだなぁ。

[Zig]: https://ziglang.org/ "Home ⚡ Zig Programming Language"
[goroutine はグリーンスレッドではない]: {{< ref "/golang/learn-concurrent-programming-with-go-1.md" >}} "goroutine はグリーンスレッドではない（『Go言語で学ぶ並行プログラミング』読書会1回目）"
[Go]: https://go.dev/

## 参考図書

{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
{{% review-paapi "4908686122" %}} <!-- Goならわかるシステムプログラミング 第2版 -->
