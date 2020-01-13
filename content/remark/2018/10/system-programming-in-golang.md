+++
title = "『Go ならわかるシステムプログラミング』を眺める"
date = "2018-10-19T15:31:01+09:00"
description = "Go 言語でシステム寄りのプログラミングを行うのであれば，是非とも目を通しておくべきだろう。"
image = "/images/attention/go-logo_blue.png"
tags = [ "book", "golang", "engineering", "programming", "concurrency" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]の言語仕様について勉強するなら真っ先に『[プログラミング言語 Go]』を推すが，もう少し実装よりの話であれば『[Go ならわかるシステムプログラミング]』を推す。
特に [Go 言語]でシステム寄りのプログラミングを行うのであれば，是非とも目を通しておくべきだろう。

この本は [ASCII.jp](http://ascii.jp/) で連載されていた「[Goならわかるシステムプログラミング](http://ascii.jp/elem/000/001/235/1235262/ "ASCII.jp：Goならわかるシステムプログラミング")」を書籍用に再構成したものらしい。
なので，まずは Web サイトの方を見て紙の本で買うべきか判断するのがいいと思う。

[Go 言語]の特徴のひとつがクロス・コンパイルの容易さである。
いくつか制約はあるが，基本的には同じコードで Windows も UNIX 系プラットフォームも対応していて，これを実現するために標準の [`syscall`] や [`runtime`] パッケージ周辺を巧妙にカプセル化している。
しかもこれらのソースコードが公開されているため[^src1] システム・プログラミングの学習教材としても使える[^sp1]。

[^src1]: [Go 言語]の標準パッケージは [MIT ライセンス]で公開されている。
[^sp1]: ちなみに『[Go ならわかるシステムプログラミング]』では，各章の最後に演習問題がある。

特に『[Go ならわかるシステムプログラミング]』では，ファイルやソケットなどに代表される順次アクセスの汎化である [`io`]`.Reader` / [`io`]`.Writer` およびその派生・特化クラス，またプロセスやスレッドに関する解説が秀逸だと思う。
さらに Docker コアの [libcontainer] についても解説があったりする（自前で [libcontainer] を直に触る人はあまりいないかも知れないが）。

個人的によく出来てると思うのが平行（concurrent）/並列（parallel）処理について解説している13章と14章だ。
プロセスやスレッド（更にはガベージコレクション）と [goroutine] の関係について日本語で分かりやすく解説している本は少ないと思うので，これだけで『[Go ならわかるシステムプログラミング]』を買っておく価値があると思う[^c1]。

[^c1]: ただし並行処理のデザインパターン等，もう少し踏み込んだ内容については『[Go 言語による並行処理]』のほうがいいかも知れない。

[goroutine] と channel の組み合わせは並行処理におけるパラダイムシフトとなる可能性がある。
それくらい高いポテンシャルを持っているのだ。
そのための基礎学習を『[Go ならわかるシステムプログラミング]』でやっておくのがいいんじゃないかな。

## ブックマーク

- [『Go 言語による並行処理』は Go 言語プログラマ必読書だろう]({{< ref "/remark/2018/11/concurrency-in-go.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`syscall`]: https://golang.org/pkg/syscall/ "syscall - The Go Programming Language"
[`runtime`]: https://golang.org/pkg/runtime/ "runtime - The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[goroutine]: https://golang.org/doc/effective_go.html#concurrency "Effective Go - The Go Programming Language"
[libcontainer]: https://github.com/docker/libcontainer "docker/libcontainer: PROJECT MOVED TO RUNC"
[プログラミング言語 Go]: {{< ref "/remark/2016/07/go-programming-language.md" >}} "『プログラミング言語 Go』を眺める"
[Go ならわかるシステムプログラミング]: https://www.amazon.co.jp/exec/obidos/ASIN/4908686033/baldandersinf-22 "Goならわかるシステムプログラミング | 渋川 よしき, ごっちん |本 | 通販 | Amazon"
[Go 言語による並行処理]: https://www.amazon.co.jp/exec/obidos/ASIN/4873118468/baldandersinf-22/ "Go言語による並行処理 | Katherine Cox-Buday, 山口 能迪 |本 | 通販 | Amazon"
[MIT ライセンス]: https://opensource.org/licenses/mit-license.php "The MIT License | Open Source Initiative"

## 参考図書

{{% review-paapi "4908686033" %}} <!-- Goならわかるシステムプログラミング -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
