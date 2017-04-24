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
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

[原書のタイトルが “Go Web Programming"](https://www.manning.com/books/go-web-programming "Manning | Go Web Programming") となっている通り，『[Goプログラミング実践入門]』は「Web アプリケーションまたはサービスのプログラミング」について [Go 言語]のコードを使って解説している本である。
しかも「実践入門」というよりは基礎学習に近い内容。
したがって既に現場でばりばりコードを書いてる人には物足りないだろう。
そういう人は（少し前に出た本だけど）オライリー・ジャパンの『[Go言語によるWebアプリケーション開発](http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/ "Go言語によるWebアプリケーション開発 | Mat Ryer, 鵜飼 文敏, 牧野 聡 |本 | 通販 | Amazon")』のほうがいいかもしれない。

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

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B06XKPNVWV/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51dQZeafzvL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B06XKPNVWV/baldandersinf-22/">Goプログラミング実践入門　標準ライブラリでゼロからWebアプリを作る impress top gearシリーズ</a></dt><dd>Sau Sheong Chang 武舎 広幸 阿部 和也 上西 昌弘 </dd><dd>インプレス 2017-03-17</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B06Y3JV86V/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B06Y3JV86V.09._SCTHUMBZZZ_.jpg"  alt="IoTエンジニア養成読本"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B06XTKZS7J/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B06XTKZS7J.09._SCTHUMBZZZ_.jpg"  alt="Electronではじめるアプリ開発 ～JavaScript/HTML/CSSでデスクトップアプリを作ろう"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B06XJ86BFZ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B06XJ86BFZ.09._SCTHUMBZZZ_.jpg"  alt="プログラミング経験者がGo言語を本格的に勉強する前に読むための本"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B06XNQCW7B/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B06XNQCW7B.09._SCTHUMBZZZ_.jpg"  alt="徹底マスター JavaScriptの教科書　プログラミングの教養から、言語仕様、開発技法までが正しく身につく"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LMS7B1O/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01LMS7B1O.09._SCTHUMBZZZ_.jpg"  alt="みんなのGo言語[現場で使える実践テクニック]"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B06X9PL5WD/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B06X9PL5WD.09._SCTHUMBZZZ_.jpg"  alt="Elixir/Phoenix 初級②: データベースとクエリ構造体 (OIAX BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01MUS2RP9/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01MUS2RP9.09._SCTHUMBZZZ_.jpg"  alt="JavaScriptエンジニアが手っ取り早くReactの基礎を理解するための「超」入門書"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01N183E3H/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01N183E3H.09._SCTHUMBZZZ_.jpg"  alt="nginx実践ガイド impress top gearシリーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01FH3KRTI/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01FH3KRTI.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01N1GOX62/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01N1GOX62.09._SCTHUMBZZZ_.jpg"  alt="WebデベロッパーのためのReact開発入門 JavaScript UIライブラリの基本と活用"  /></a> </p>
<p class="description">Web アプリケーションまたはサービスについて Go 言語のコードで解説。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-04-23">2017-04-23</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51UoREcNrnL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/">Go言語によるWebアプリケーション開発</a></dt><dd>Mat Ryer 鵜飼 文敏 </dd><dd>オライリージャパン 2016-01-22</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621300253.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語Go"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117607/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117607.09._SCTHUMBZZZ_.jpg"  alt="マイクロサービスアーキテクチャ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774178667/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774178667.09._SCTHUMBZZZ_.jpg"  alt="nginx実践入門 (WEB+DB PRESS plus)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4863541783/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4863541783.09._SCTHUMBZZZ_.jpg"  alt="改訂2版 基礎からわかる Go言語"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774179930/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774179930.09._SCTHUMBZZZ_.jpg"  alt="サーバ/インフラエンジニア養成読本 DevOps編 [Infrastructure as Code を実践するノウハウが満載! ] (Software Design plus)"  /></a> </p>
<p class="description">日本語監訳者による解説（付録 B）が意外に役に立つ感じ。 Web アプリケーションだけでなく，サーバサイドで動く CLI アプリへの言及もある。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-03-13">2016-03-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
