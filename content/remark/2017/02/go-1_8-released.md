+++
title = "Go 言語 1.8 がリリース"
tags = ["golang", "programming", "language"]
description = "Go 言語コンパイラのバージョン 1.8 がリリースされた。"
date = "2017-02-19T15:45:53+09:00"
image = "/images/attention/go-logo_blue.png"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]コンパイラのバージョン 1.8 がリリースされた。

- [Go 1.8 is released - The Go Blog](https://blog.golang.org/go1.8)

主な改善点を上げると

- コンパイル時間の短縮
- ガベージコレクションの改善（通常で 100μsec 未満，しばしば 10μsec 未満）
- HTTP/2 Push の追加
- 標準ライブラリの [`context`] パッケージについてキャンセルとタイムアウトの仕組みを追加
    - HTTP サーバのシャットダウンの改善など
- [`sort`].[`Slice`](https://golang.org/pkg/sort/#Slice) 関数の追加。 [slice] のソートが簡単になった

その他，詳しくは以下のリリースノートを参照のこと。

- [Go 1.8 Release Notes - The Go Programming Language](https://golang.org/doc/go1.8)

## 戯れ言

2015年頃から（仕事以外でだが） [Go 言語]で遊ぶようになって2年半近くが過ぎた。

仕事においては，業務システムでは相変わらず Java への replace 仕事ばっかりだし（私のようなロートルエンジニアは過去の[技術的負債](https://ja.wikipedia.org/wiki/%E6%8A%80%E8%A1%93%E7%9A%84%E8%B2%A0%E5%82%B5)の後始末をするのがお役目），組み込みでは C/C++ がメインなので， [Go 言語]を使う機会がないのだが，恐ろしいことに私の中で [Go 言語]が「[母国語]」になりつつある。
つまり，あるロジックをプログラム・コードに「翻訳」する際に，まず [Go 言語]のコードが思い浮かぶようになってきた。

この業界に四半世紀以上も足を突っ込んでるが脳内の[母国語]が変わるという経験は初めてで，まるで転生物のラノベ作品を読むがごとく，年甲斐もなく「[わーい！ たのしー！](https://nijipi.com/it-news/kemono-lang_ruby-brainfuck/)」な気分でコードを眺める日々である。

もっとも，有り余る計算資源を持つクラウド環境ならともかくリソースの限られた RTOS (Real-Time Operating System) 環境下では息を吸うようにヒープを使いまくる [Go 言語]での実装は向いてない気がするので，「これは言語のチョイスを間違えたかなぁ」とも思わないでもない。
まぁでもそれならそれで C/C++ を使えばいいので困ることでもないんだけどね。

でも  [Go 言語]が[母国語]になると（アセンブラに近い C 言語はともかく） C++ って本当に面倒くさい言語だったんだなぁ，と涙が出ちゃう。
だってエンジニアだもん。

## ブックマーク

- [Go言語がダメな理由 | プログラミング | POSTD](http://postd.cc/why-go-is-not-good/)
- [Go言語のリアルタイムGC　理論と実践 | プログラミング | POSTD](http://postd.cc/golangs-real-time-gc-in-theory-and-practice/)

- [プログラミング言語 Go](/golang/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[母国語]: {{< ref "/remark/2015/programming-language.md" >}} "プログラミング言語との付き合い方"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"
[`sort`]: https://golang.org/pkg/sort/ "sort - The Go Programming Language"
[slice]: http://golang.org/ref/spec#Slice_types

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->

{{% review-paapi "4320026926" %}} <!-- プログラミング言語C -->
