+++
title = "Rust 勉強会，面白かった！"
date =  "2020-02-24T21:58:44+09:00"
description = "しばらくはマジで Rust の勉強をしようかなぁ，とか思ったり。"
image = "/images/attention/kitten.jpg"
tags  = [ "programming", "language", "rust", "shimane", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

以前から気になっていた [Rust] の勉強会を松江でやるというので参加してきた。

- [east-shimane.rs #0 - connpass](https://connpass.com/event/163790/)

いやぁ，面白かった！
C/C++ のコードを [Rust] に置き換えようという動きがあるのは知っていたが，マジで「アリ」だわ，これ。
やはり時代は [Rust] か。

実は今回，どうにかコンパイル環境のインストールまでは済ませたが事前学習が全くできなくて，初心者講座でイチから教えてもらうことにした。
以下に[勉強会](https://connpass.com/event/163790/ "east-shimane.rs #0 - connpass")の内容を一部紹介する。

[Rust] の最大の特徴は GC (Garbage Collection) をランタイム・モジュール内に持たないことだろう。
かといって C/C++ のように自前で明示的に alloc/free するわけでもない。 

どうするのかというと， [Rust] にはインスタンスに対して「所有権（ownership）」や「ライフタイム（lifetime）」といった概念があり，野放図にインスタンスを参照することを防いでいる。
また変数は基本的に不変（immutable）な値として宣言され，可変（mutable）にしたいのであれば可変であることを明示して宣言する必要がある。
これらの仕組みによってコンパイラが alloc/free のタイミングを決めることができる（または決められない処理についてはコンパイル・エラーにできる）わけだ。

[Go] の実装にも見られるように「GC ＝ 遅い」とは言えなくなりつつあるのだが，ランタイム・モジュールに GC を組み込まなくていいというのはそれだけで設計の自由度が上がる。
それこそ，その言語で OS を組もうかと考えたくなるほどに。

コンパイル（VM 用の JIT コンパイルや他言語へのトランス・コンパイルではなく，いわゆる AOT コンパイル）言語の評価ポイントは，個人的には以下の2つと思っている[^l1]。

[^l1]: オブジェクト指向がどうとか関数型がどうとかいうのはジオングの脚のような些末な話である。

1. その言語自身でコンパイラが書けるか
2. その言語を使って OS が書けるか

たとえば [Go] は1番目は満たすが2番目は難しい。
メモリ管理や並列処理等についてランタイム・モジュールに強く依存しているからだ。
[Rust] は両方を満たす[^lnk1]。

[^lnk1]: ただし [Rust] はリンカについては GCC 等の外部のツールチェーンに依存しているようだ。なのでクロス・コンパイルは（[Go] に比べると）少し面倒くさい。

他にも enum と match 式を組み合わせたエラー・ハンドリングとか興味深いものが色々あった。
しばらくはマジで [Rust] の勉強をしようかなぁ，とか思ったり。

## ブックマーク

- [The Rust Programming Language](https://doc.rust-lang.org/book/)
    - [The Rust Programming Language](https://doc.rust-jp.rs/book/second-edition/) : 日本語版
- [Rust by Example](https://doc.rust-jp.rs/rust-by-example-ja/) : 日本語版
- [Rust Playground](https://play.rust-lang.org/)

[Rust]: https://www.rust-lang.org/ "Rust Programming Language"
[Go]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
{{% review-paapi "B07QVQ7RDG" %}} <!-- 実践Rust入門 -->
