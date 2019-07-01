+++
date = "2017-04-23T19:08:38+09:00"
update = "2017-04-24T09:26:58+09:00"
description = "個人的には context パッケージについて解説があるとなおよかったが， Go 言語 の 1.6 をベースに書かれているので無理か。"
draft = false
tags = ["book", "golang", "web", "programming"]
title = "『Goプログラミング実践入門』を眺める"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"
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

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B06XKPNVWV/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51dQZeafzvL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B06XKPNVWV/baldandersinf-22/">Goプログラミング実践入門　標準ライブラリでゼロからWebアプリを作る impress top gearシリーズ</a></dt><dd>Sau Sheong Chang 武舎 広幸 阿部 和也 上西 昌弘 </dd><dd>インプレス 2017-03-17</dd><dd>評価<abbr class="rating" title="4"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B06Y3JV86V/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B06Y3JV86V.09._SCTHUMBZZZ_.jpg"  alt="IoTエンジニア養成読本"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B06XTKZS7J/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B06XTKZS7J.09._SCTHUMBZZZ_.jpg"  alt="Electronではじめるアプリ開発 ～JavaScript/HTML/CSSでデスクトップアプリを作ろう"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B06XJ86BFZ/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B06XJ86BFZ.09._SCTHUMBZZZ_.jpg"  alt="プログラミング経験者がGo言語を本格的に勉強する前に読むための本"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B06XNQCW7B/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B06XNQCW7B.09._SCTHUMBZZZ_.jpg"  alt="徹底マスター JavaScriptの教科書　プログラミングの教養から、言語仕様、開発技法までが正しく身につく"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01LMS7B1O/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B01LMS7B1O.09._SCTHUMBZZZ_.jpg"  alt="みんなのGo言語[現場で使える実践テクニック]"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B06X9PL5WD/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B06X9PL5WD.09._SCTHUMBZZZ_.jpg"  alt="Elixir/Phoenix 初級②: データベースとクエリ構造体 (OIAX BOOKS)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01MUS2RP9/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B01MUS2RP9.09._SCTHUMBZZZ_.jpg"  alt="JavaScriptエンジニアが手っ取り早くReactの基礎を理解するための「超」入門書"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01N183E3H/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B01N183E3H.09._SCTHUMBZZZ_.jpg"  alt="nginx実践ガイド impress top gearシリーズ"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01FH3KRTI/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B01FH3KRTI.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01N1GOX62/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B01N1GOX62.09._SCTHUMBZZZ_.jpg"  alt="WebデベロッパーのためのReact開発入門 JavaScript UIライブラリの基本と活用"  /></a> </p>
<p class="description">Web アプリケーションまたはサービスについて Go 言語のコードで解説。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-04-23">2017-04-23</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%AB%E3%82%88%E3%82%8BWeb%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E9%96%8B%E7%99%BA-Mat-Ryer/dp/4873117526?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4873117526"><img src="https://images-fe.ssl-images-amazon.com/images/I/51UoREcNrnL._SL160_.jpg" width="125" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%AB%E3%82%88%E3%82%8BWeb%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E9%96%8B%E7%99%BA-Mat-Ryer/dp/4873117526?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4873117526">Go言語によるWebアプリケーション開発</a></dt>
	<dd>Mat Ryer</dd>
	<dd>鵜飼 文敏 (監修), 牧野 聡 (翻訳)</dd>
    <dd>オライリージャパン 2016-01-22</dd>
    <dd>Book 大型本</dd>
    <dd>ASIN: 4873117526, EAN: 9784873117522</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">日本語監訳者による解説（付録 B）が意外に役に立つ感じ。 Web アプリケーションだけでなく，サーバサイドで動く CLI アプリへの言及もある。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
