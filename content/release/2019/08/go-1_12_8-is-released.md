+++
title = "Go 1.12.8 がリリース【セキュリティ・アップデート】"
date =  "2019-08-14T09:01:05+09:00"
description = "2つのセキュリティ脆弱性の修正がある。アップデートは計画的に。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.12.8 がリリースされた。

- [[security] Go 1.12.8 and Go 1.11.13 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/65QixT3tcmg)

今回はセキュリティ・アップデートを含んでいる。
[Go 言語]で Web サービスを構築している人は要注意だ。

## net/http: Denial of Service vulnerabilities in the HTTP/2 implementation

- [CVE-2019-9512](https://nvd.nist.gov/vuln/detail/CVE-2019-9512)
- [CVE-2019-9514](https://nvd.nist.gov/vuln/detail/CVE-2019-9514)

{{< fig-quote type="md" title="Go 1.12.8 and Go 1.11.13 are released" link="https://groups.google.com/forum/#!topic/golang-announce/65QixT3tcmg" lang="en" >}}
{{< quote >}}`net/http` and [`golang.org/x/net/http2`](http://golang.org/x/net/http2) servers that accept direct connections from untrusted clients could be remotely made to allocate an unlimited amount of memory, until the program crashes. Servers will now close connections if the send queue accumulates too many control messages.
The issues are CVE-2019-9512 and CVE-2019-9514, and Go issue [golang.org/issue/33606](https://golang.org/issue/33606).{{< /quote >}}
{{< /fig-quote >}}

Netflix の中の人，ありがとう。

## net/url: parsing validation issue

- [CVE-2019-14809](https://nvd.nist.gov/vuln/detail/CVE-2019-14809)

{{< fig-quote type="md" title="Go 1.12.8 and Go 1.11.13 are released" link="https://groups.google.com/forum/#!topic/golang-announce/65QixT3tcmg" lang="en" >}}
{{< quote >}}`url.Parse` would accept URLs with malformed hosts, such that the Host field could have arbitrary suffixes that would appear in neither `Hostname()` nor `Port()`, allowing authorization bypasses in certain applications. Note that URLs with invalid, not numeric ports will now return an error from url.Parse.
The issue is CVE-2019-14809 and Go issue [golang.org/issue/29098](https://golang.org/issue/29098).{{< /quote >}}
{{< /fig-quote >}}

## アップデートは計画的に

[Ubuntu] の APT は相変わらずサポートから外れた 1.10 しか対応していないので[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.12.8.linux-amd64.tar.gz`](https://dl.google.com/go/go1.12.8.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.12.8.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.12.8.linux-amd64.tar.gz
$ sudo mv go go1.12.8
$ sudo ln -s go1.12.8 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.12.8 linux/amd64
```

アップデートは計画的に。
時期的に次は 1.13 のリリースかな。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## ブックマーク

- [VU#605641 - HTTP/2 implementations do not robustly handle abnormal traffic and resource exhaustion](https://www.kb.cert.org/vuls/id/605641/)
    - [[HTTP/2実装に脆弱性、DoS攻撃のおそれ | マイナビニュース](https://news.mynavi.jp/article/20190816-877630/)]

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/B07VPSXF6N?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51jif840ScL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/B07VPSXF6N?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">改訂2版 みんなのGo言語</a></dt>
    <dd>松木 雅幸 (著), mattn (著), 藤原 俊一郎 (著), 中島 大一 (著), 上田 拓也 (著), 牧 大輔 (著), 鈴木 健太 (著)</dd>
    <dd>技術評論社 2019-08-01 (Release 2019-08-01)</dd>
    <dd>Kindle版</dd>
    <dd>B07VPSXF6N (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">改訂2版の目玉は7章の「データベースの扱い方」が追加されたことだろう。他の章では，大まかな構成は1版と同じだが細かい部分が変わっていて Go 1.12 への言及まであるのには驚いた。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-08-12">2019-08-12</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
