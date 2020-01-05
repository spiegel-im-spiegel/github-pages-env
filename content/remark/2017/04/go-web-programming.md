+++
title = "『Goプログラミング実践入門』を眺める"
date = "2017-04-23T19:08:38+09:00"
description = "個人的には context パッケージについて解説があるとなおよかったが， Go 言語 の 1.6 をベースに書かれているので無理か。"
tags = ["book", "golang", "web", "programming"]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[原書のタイトルが “Go Web Programming"](https://www.manning.com/books/go-web-programming "Manning | Go Web Programming") となっている通り，『[Goプログラミング実践入門]』は「Web アプリケーションまたはサービスのプログラミング」について [Go 言語]のコードを使って解説している本である。
しかも「実践入門」というよりは基礎学習に近い内容。
したがって既に現場でばりばりコードを書いてる人には物足りないだろう。
そういう人は（少し前に出た本だけど）オライリー・ジャパンの『[Go言語によるWebアプリケーション開発](https://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/ "Go言語によるWebアプリケーション開発 | Mat Ryer, 鵜飼 文敏, 牧野 聡 |本 | 通販 | Amazon")』のほうがいいかもしれない。

特徴的なのが， [Echo](https://echo.labstack.com/ "Echo - High performance, minimalist Go web framework") のような有名どころのフレームワークは使わず， [`net`]/[`http`] や [`html`]/[`template`] といった標準パッケージのみで解説しているところ[^pkg]。
大昔によくあった， TCP/IP の解説を C 言語コードで行ったり CGI (Common Gateway Interface) の解説を Perl のコードで行ってた技術解説本のようなノリがあってなかなか楽しく読めた。

[^pkg]: テストフレームワークや ORM (Object-Relational Mapping) についてはサードパーティのパッケージも紹介している。また RDBMS のドライバは標準パッケージとしては提供されないので，サードパーティのパッケージが使われている（『[Goプログラミング実践入門]』では PostgreSQL なので [`github.com/lib/pq`] を使用）。

Web アプリケーション以外でもそうだけど，フレームワークって「中身」がちゃんと分かってないと適切に使えないよね。
そういう意味ではよく出来てると思う。
個人的には [`context`] パッケージについて解説があるとなおよかったが， [Go 言語] の 1.6 をベースに書かれているので無理か。

[Goプログラミング実践入門]: http://book.impress.co.jp/books/1115101145 "Goプログラミング実践入門 標準ライブラリでゼロからWebアプリを作る - インプレスブックス"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`net`]: https://golang.org/pkg/net/ "net - The Go Programming Language"
[`http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"
[`html`]: https://golang.org/pkg/html/ "html - The Go Programming Language"
[`template`]: https://golang.org/pkg/html/template/ "template - The Go Programming Language"
[`github.com/lib/pq`]: https://github.com/lib/pq "lib/pq: Pure Go Postgres driver for database/sql"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"

## 参考図書

{{% review-paapi "B06XKPNVWV" %}} <!-- Goプログラミング実践入門　標準ライブラリでゼロからWebアプリを作る -->

{{% review-paapi "4873117526" %}} <!-- Go言語によるWebアプリケーション開発 -->
