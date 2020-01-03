+++
title = "Go 1.13.2 および Go 1.13.3 のリリース【セキュリティ・アップデート】"
date =  "2019-10-18T05:00:51+09:00"
description = "今回のはちょーっとヤバめかも。あと Go 1.13.3 のリリースを追記した。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.13.2 がリリースされた。

- [[security] Go 1.13.2 and Go 1.12.11 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/lVEm7llp0w0)

今回のはちょーっとヤバめかも。

- [CVE-2019-17596](https://nvd.nist.gov/vuln/detail/CVE-2019-17596)

HTTPS サーバを構成している場合は要注意。

{{< fig-quote type="markdown" title="Go 1.13.2 and Go 1.12.11 are released" link="https://groups.google.com/forum/#!topic/golang-announce/lVEm7llp0w0" lang="en" >}}
{{< quote >}}Invalid DSA public keys can cause a panic in `dsa.Verify`. In particular, using `crypto/x509.Verify` on a `crafted X.509` certificate chain can lead to a panic, even if the certificates don’t chain to a trusted root. The chain can be delivered via a `crypto/tls` connection to a client, or to a server that accepts and verifies client certificates. `net/http` clients can be made to crash by an HTTPS server, while net/http servers that accept client certificates will recover the panic and are unaffected{{< /quote >}}.
{{< /fig-quote >}}

他にも OpenPGP, OTR, SSH あたりも影響を受けるそうな。

{{< fig-quote type="markdown" title="Go 1.13.2 and Go 1.12.11 are released" link="https://groups.google.com/forum/#!topic/golang-announce/lVEm7llp0w0" lang="en" >}}
{{< quote >}}Moreover, an application might crash invoking `crypto/x509.(*CertificateRequest).CheckSignature` on an X.509 certificate request, parsing a [golang.org/x/crypto/openpgp](http://golang.org/x/crypto/openpgp) Entity, or during a [golang.org/x/crypto/otr](http://golang.org/x/crypto/otr) conversation. Finally, a [golang.org/x/crypto/ssh](http://golang.org/x/crypto/ssh) client can panic due to a malformed host key, while a server could panic if either `PublicKeyCallback` accepts a malformed public key, or if `IsUserAuthority` accepts a certificate with a malformed public key{{< /quote >}}.
{{< /fig-quote >}}

更に更に slice に関して

{{< fig-quote type="markdown" title="Go 1.13.2 and Go 1.12.11 are released" link="https://groups.google.com/forum/#!topic/golang-announce/lVEm7llp0w0" lang="en" >}}
{{< quote >}}The Go 1.13.2 release also includes a fix to the compiler that prevents improper access to negative slice indexes in rare cases. Affected code, in which the compiler can prove that the index is zero or negative, would have resulted in a panic in Go 1.12, but could have led to arbitrary memory read and writes in Go 1.13 and Go 1.13.1{{< /quote >}}.
{{< /fig-quote >}}

とあり， [Go] 1.13.x でバイナリを提供している場合はもれなくリコンパイルする必要があるかもねぇ。

## 【追記】 [Go] 1.13.3 のリリース

まさか5時間足らずでアップデートするとは思わざりき。
別途記事を起こすのはアレなので，この記事に追記する。
フットワークの軽い言語だとポジティブに考えることにしよう（笑）

というわけで [Go] 1.13.3 がリリースされた。

- [Go 1.13.3 and Go 1.12.12 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/R3XK-Wf-Mtk)

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://github.com/golang/go/issues?q=milestone%3AGo1.13.3" lang="en" >}}
{{< quote >}}go1.13.3 (released 2019/10/17) includes fixes to the go command, the toolchain, the `runtime`, `syscall`, `net`, `net/http`, and `crypto/ecdsa` packages. See the [Go 1.13.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.3) on our issue tracker for details{{< /quote >}}.
{{< /fig-quote >}}

1.13.2 のアップデートがまだの人はまとめてやってしまおう。
私のように既にアップデートしてしまった人は，ごくろーさん（泣）

[Ubuntu] の APT は相変わらずサポートから外れた 1.10.x しか対応していないので[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.13.3.linux-amd64.tar.gz`](https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.13.3.linux-amd64.tar.gz
$ sudo mv go go1.13.3
$ sudo ln -s go1.13.3 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.13.3 linux/amd64
```

アップデートは計画的に。

## ブックマーク

- [crypto/dsa: invalid public key causes panic in dsa.Verify · Issue #34960 · golang/go · GitHub](https://github.com/golang/go/issues/34960)
- [cmd/compile: access to negative slice indices improperly permitted · Issue #34802 · golang/go · GitHub](https://github.com/golang/go/issues/34802)

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

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
