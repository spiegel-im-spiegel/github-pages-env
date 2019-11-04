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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/B06XKPNVWV?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51dQZeafzvL._SL160_.jpg" width="125" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/B06XKPNVWV?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">Goプログラミング実践入門　標準ライブラリでゼロからWebアプリを作る impress top gearシリーズ</a></dt>
    <dd>Sau Sheong Chang (著), 武舎 広幸 (著), 阿部 和也 (著), 上西 昌弘 (著)</dd>
    <dd>インプレス 2017-03-17 (Release 2017-03-17)</dd>
    <dd>Kindle版</dd>
    <dd>B06XKPNVWV (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">Web アプリケーションまたはサービスについて Go 言語のコードで解説。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-04-23">2017-04-23</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4873117526?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51UoREcNrnL._SL160_.jpg" width="125" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4873117526?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">Go言語によるWebアプリケーション開発</a></dt>
    <dd>Mat Ryer (著), 鵜飼 文敏 (監修), 牧野 聡 (翻訳)</dd>
    <dd>オライリージャパン 2016-01-22</dd>
    <dd>大型本</dd>
    <dd>4873117526 (ASIN), 9784873117522 (EAN), 4873117526 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="3">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">日本語監訳者による解説（付録 B）が意外に役に立つ感じ。 Web アプリケーションだけでなく，サーバサイドで動く CLI アプリへの言及もある。良書だが今となってはちょっと内容が古い。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-08-12">2019-08-12</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
