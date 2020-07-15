+++
title = "Go 1.13.7 のリリース【セキュリティ・アップデート】"
date =  "2020-01-29T15:21:25+09:00"
description = "今回は2件のセキュリティ・アップデートを含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[先週の予告]({{< ref "/remark/2020/01/pre-announcement-of-go-1_13_7.md" >}} "Go 1.13.7 リリース予告と CVE-2020-0601")どおり [Go] 1.13.7 がリリースされた。

- [[security] Go 1.13.7 and Go 1.12.16 are released - Google group](https://groups.google.com/forum/#!topic/golang-announce/Hsw4mHYc470)

2件のセキュリティ・アップデートを含んでいる。

## 【[CVE-2020-0601]】 X.509 certificate validation bypass on Windows 10

{{< fig-quote type="markdown" title="[security] Go 1.13.7 and Go 1.12.16 are released" link="https://groups.google.com/forum/#!topic/golang-announce/Hsw4mHYc470" lang="en" >}}
A Windows vulnerability allows attackers to spoof valid certificate chains when the system root store is in use. These releases include a mitigation for Go applications, but it’s strongly recommended that affected users install the Windows security update to protect their system.<br>
This issue is CVE-2020-0601 and Go issue [golang.org/issue/36834](http://golang.org/issue/36834).
{{< /fig-quote >}}

と書かれている通り，今回の [Go] 側の措置は「低減（mitigation）」に過ぎないので，必ず Windows Update を適用すること。
Win7 ユーザはとっとと [Win10 に乗り換えるか有償サポートを受けるか](https://news.mynavi.jp/article/20200128-961593/ "Windows 7のIEに脆弱性も無料のセキュリティパッチ提供なし、対処を | マイナビニュース") Windows を捨てるかしましょう。

{{< fig-gen type="markdown" title="CVSS:3.1/AV:N/AC:L/PR:N/UI:R/S:U/C:H/I:H/A:N  (基本スコア:8.1 [重要])" >}}
| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 要           |
| スコープ         | 変更なし     |
| 機密性への影響   | 高           |
| 完全性への影響   | 高           |
| 可用性への影響   | なし         |
{{< /fig-gen >}}

## 【[CVE-2020-7919]】 Panic in crypto/x509 certificate parsing and golang.org/x/crypto/cryptobyte

{{< fig-quote type="markdown" title="[security] Go 1.13.7 and Go 1.12.16 are released" link="https://groups.google.com/forum/#!topic/golang-announce/Hsw4mHYc470" lang="en" >}}
On 32-bit architectures, a malformed input to `crypto/x509` or the ASN.1 parsing functions of [`golang.org/x/crypto/cryptobyte`](http://golang.org/x/crypto/cryptobyte) can lead to a panic.
The malformed certificate can be delivered via a `crypto/tls` connection to a client, or to a server that accepts client certificates. `net/http` clients can be made to crash by an HTTPS server, while `net/http` servers that accept client certificates will recover the panic and are unaffected.<br>
Thanks to [Project Wycheproof](https://github.com/google/wycheproof) for providing the test cases that led to the discovery of this issue.<br>
The issue is CVE-2020-7919 and Go issue [golang.org/issue](http://golang.org/issue/)/36837.<br>
This is also fixed in version v0.0.0-20200124225646-8b5121be2f68 of [`golang.org/x/crypto/cryptobyte`](http://golang.org/x/crypto/cryptobyte).
{{< /fig-quote >}}

[Go] に限らず，今後32ビット・アーキテクチャにおける不備は増えていくんだろうねぇ。

{{< fig-gen type="markdown" title="CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H (基本スコア:7.5 [重要])" >}}
| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | なし |
| 完全性への影響 | なし |
| 可用性への影響 | 高 |
{{< /fig-gen >}}

## 例によって

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.13.7.linux-amd64.tar.gz`](https://dl.google.com/go/go1.13.7.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.13.7.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.13.7.linux-amd64.tar.gz
$ sudo mv go go1.13.7
$ sudo ln -s go1.13.7 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.13.7 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[CVE-2020-0601]: https://nvd.nist.gov/vuln/detail/CVE-2020-0601
[CVE-2020-7919]: https://nvd.nist.gov/vuln/detail/CVE-2020-7919
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
