+++
title = "『Go言語で学ぶ並行プログラミング』は必読書である（夏の読書感想文2）"
date =  "2025-08-23T23:16:37+09:00"
description = "特に第1部と第3部は全てのプログラマが読むべき内容だと思う"
image = "/images/attention/kitten.jpg"
tags = [ "book", "golang", "engineering", "programming", "concurrency" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

夏の読書感想文。
『[にぎやかな未来]({{< ref "/remark/2025/07/noisy-future.md" >}} "『にぎやかな未来』を読む（夏の読書感想文1）")』に続いて今回は『[Go言語で学ぶ並行プログラミング]』の感想文。

今年の1月から始まったオンライン読書会は[第7回](https://technical-book-reading-2.connpass.com/event/363980/ "第7回『Go言語で学ぶ並行プログラミング』オンライン読書会 - connpass")で無事に終了。
この記事ではあまり深堀りせず「訳者あとがき」やオンライン読書会の話などを中心にさらっと書いてみる。
なお『[Go言語で学ぶ並行プログラミング]』は[版元][Go言語で学ぶ並行プログラミング]で PDF 版が買える。
もしデジタル版をご所望であれば PDF 版を購入することを強くお勧めする。

オンライン読書会は翻訳者の[柴田芳樹]さんが主催者なのだが，[柴田芳樹]さんといえば Java の技術参考書の翻訳本でも有名な方で，今回の読書会でも Java との対比で解説されることも多々あり，非常に勉強になった。
サンプルコードなどは，タイトル通り [Go] 言語が前提になっているが，内容としては特定のプログラミング言語に限ることなく普遍的な内容になっている。

『[Go言語で学ぶ並行プログラミング]』は大きく三部構成になっている。

- 第1部 並行プログラミングの基礎
- 第2部 メッセージパッシング
- 第3部 並行処理のさらなるトピック

大雑把な内容はこんな感じ。

{{< fig-quote type="markdown" title="『Go言語で学ぶ並行プログラミング』訳者あとがき" link="https://book.impress.co.jp/books/1123101144" >}}
「第1部 並行プログラミングの基礎」では、Go 言語に限定せず、並行プログラミングについての多くの基礎的な事柄が説明されています。したがって、他のプログラミング言語での開発においても役立つ知識となるでしょう。実際、私自身も Go 言語を学ぶまでは、この第1部で紹介される技法だけで長年プログラミングを行ってきました。

「第2部 メッセージパッシング」では、Go 言語が提供するチャネルを中心に解説しています。そのため、内容の多くは Go 言語に特化しており、 Go でプログラミングを行う開発者にとって必須の知識となります。

「第3部 並行処理のさらなるトピック」は高度なトピックを扱っていますが、開発者として知っておくべき重要な領域です。
{{< /fig-quote >}}

基本的に最初から順に読んでいくのがオススメだが，特に第1部と第3部は全てのプログラマが読むべき内容だと思う。
第2部は CSP (communicating sequential processes) と呼ばれる並行処理モデルについて詳しく解説している。
具体的には [Go] のチャネルを使ったデザインパターンを紹介している。

並行プログラミングを [Go] で書く強みは，カーネルレベルスレッドとユーザーレベルスレッドを組み合わせた極軽量な M:N スレッディングモデルと第一級オブジェクト（first-class object）として実装されているチャネルにある。
例えば，大量の軽量スレッドとチャネルを組み合わせたパイプライン・パターンを使った素数探索コードを紹介している。
“[Go Playground](https://go.dev/play/ "Go Playground - The Go Programming Language")” でもサンプルコードとして “[Concurrent Prime Sieve](https://go.dev/play/p/iN6HCp_e91p)” が例示されているので見比べてみると面白いかもしれない[^cps1]。

[^cps1]: パイプライン・パターンを使った素数探索は，残念ながら速くない。これは素数探索に使える最適化アルゴリズムが使えず，発見した素数の数だけフィルタ処理を行う goroutine が生成されてしまうため。 “[Concurrent Prime Sieve](https://go.dev/play/p/iN6HCp_e91p)” は `main()` 関数で指定した数だけ素数を探すのだが，たとえば100個の素数を探す場合は最大で100個の goroutine が生成され goroutine 間を繋ぐチャネルも同じ数だけ生成される。

『[Go言語で学ぶ並行プログラミング]』には [Go] では今や一般的になった [`context`] パッケージについての言及がない[^ctx1]。
その代わりもう少し簡単な quit チャネルを使ったパターンを紹介している。
[Go] 以外の言語で考える場合や [`context`] パッケージがとっつきにくいと感じる人は9章で登場する quit チャネルを使ったパターンで考えるほうが分かりやすいかもしれない。

[^ctx1]: [`context`] パッケージについては『[Go言語 100Tips]』の8章〜9章で触れられている。

「訳者あとがき」に書かれているが『[Go言語で学ぶ並行プログラミング]』にはメモリモデルについての言及がない。
[Go] の場合は公式ドキュメントとして “[The Go Memory Model](https://go.dev/ref/mem "The Go Memory Model - The Go Programming Language")” が公開されているので，一度は目を通しておくとよいだろう。
でもドキュメントの冒頭には

{{< fig-quote type="markdown" title="The Go Memory Model" link="https://go.dev/ref/mem" lang="en" >}}
If you must read the rest of this document to understand the behavior of your program, you are being too clever.
{{< /fig-quote >}}

とか書いてあったりするので，沼らない程度にさらっと見るくらいでいいだろう[^mm1]（笑）

[^mm1]: [Go] のメモリモデルについては『[Go言語 100Tips]』の8.4.2節に言及がある。

他にも「訳者あとがき」には本編にはない有用な情報が書かれているので最後まで余さず読んでみてほしい。

以前は[『Go 言語による並行処理』をお勧め]({{< ref "/remark/2018/11/concurrency-in-go.md" >}} "『Go 言語による並行処理』は Go 言語プログラマ必読書だろう")していたが，今なら断然『[Go言語で学ぶ並行プログラミング]』のほうをお勧めする。
まぁ，あとから出た本のほうがよく出来てるのは当然なんだろうけど。
斬新な内容というわけではないけど，並行プログラミングを書く際の基本を押さえた堅実な内容だと思う。

とあるオンライン読書会の余談で出てきた話題なのだが「[食事する哲学者の問題](https://ja.wikipedia.org/wiki/%E9%A3%9F%E4%BA%8B%E3%81%99%E3%82%8B%E5%93%B2%E5%AD%A6%E8%80%85%E3%81%AE%E5%95%8F%E9%A1%8C "食事する哲学者の問題 - Wikipedia")（[dining philosophers problem](https://en.wikipedia.org/wiki/Dining_philosophers_problem "Dining philosophers problem - Wikipedia")）」というのがあるそうな。

簡単に言うとこんな感じに円卓に複数の哲学者がいて食事をするのだが

{{< fig-img-quote src="./Dining_philosophers_diagram.jpg" title="File:Dining philosophers diagram.jpg - Wikimedia Commons" link="https://commons.wikimedia.org/wiki/File:Dining_philosophers_diagram.jpg" lang="en" >}}

食事にはひとりあたり2つのフォークが必要なのに人数分しか用意されていない。
そのままだとフォークの取り合いによるデッドロックが発生してしまうわけだ。
デッドロックによる飢餓（starvation）が発生しないようにするにはどうするか，という問題である。

並行プログラミングの練習問題として丁度いいよね。
これを [Go] で解くことを「夏休みの自由研究」にしようかと思っていたのだが，もたもたしてるうちに夏が終わりそうである。

## ブックマーク

- [goroutine はグリーンスレッドではない（『Go言語で学ぶ並行プログラミング』読書会1回目）]({{< ref "/golang/learn-concurrent-programming-with-go-1.md" >}})
- [条件変数とセマフォ（『Go言語で学ぶ並行プログラミング』読書会3回目）]({{< ref "/golang/learn-concurrent-programming-with-go-3.md" >}})
- [『Go言語で学ぶ並行プログラミング』の練習問題より]({{< ref "/golang/learn-concurrent-programming-with-go-2x.md" >}})

[Go]: https://go.dev/
[Go言語で学ぶ並行プログラミング]: https://book.impress.co.jp/books/1123101144 "Go言語で学ぶ並行プログラミング　他言語にも適用できる原則とベストプラクティス - インプレスブックス"
[Go言語 100Tips]: https://book.impress.co.jp/books/1122101133 "Go言語 100Tips ありがちなミスを把握し、実装を最適化する - インプレスブックス"
[柴田芳樹]: https://note.com/yoshiki_shibata "柴田 芳樹｜note"
[`context`]: https://pkg.go.dev/context "context package - context - Go Packages"

## 参考図書

{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
