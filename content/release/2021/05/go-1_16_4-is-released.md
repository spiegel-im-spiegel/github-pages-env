+++
title = "Go 1.16.4 のリリース【セキュリティ・アップデート】"
date =  "2021-05-09T13:52:11+09:00"
description = "脆弱性 CVE-2021-31525 の修正"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.16.4 がリリースされた。

- [Go 1.16.4 and Go 1.15.12 are released](https://groups.google.com/g/golang-announce/c/cu9SP4eSXMc)

今回は1件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.16.minor" lang="en" >}}
{{% quote %}}go1.16.4 (released 2021/05/06) includes a security fix to the `net/http` package, as well as bug fixes to the runtime, the compiler, and the `archive/zip`, `time`, and `syscall` packages. See the [Go 1.16.4 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.4+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2021-31525]

{{< fig-quote type="markdown" title="Go 1.16.4 and Go 1.15.12 are released" link="https://groups.google.com/g/golang-announce/c/cu9SP4eSXMc" lang="en" >}}
{{% quote %}}`ReadRequest` and `ReadResponse` in `net/http` can hit an unrecoverable panic when reading a very large header (over 7MB on 64-bit architectures, or over 4MB on 32-bit ones). `Transport` and `Client` are vulnerable and the program can be made to crash by a malicious server.  `Server` is not vulnerable by default, but can be if the default max header of 1MB is overridden by setting `Server.MaxHeaderBytes` to a higher value, in which case the program can be made to crash by a malicious client{{% /quote %}}.
{{< /fig-quote >}}

この脆弱性は

- [`golang.org/x/net/http2/h2c`](http://golang.org/x/net/http2/h2c)
- `HeaderValuesContainsToken` in [`golang.org/x/net/http/httpguts`](http://golang.org/x/net/http/httpguts)

にも影響するらしい。
該当のパッケージを利用している場合は最新リビジョンにアップデートすること。

- `CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:H`
- 深刻度: 警告 (Score: 5.9)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 高 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | なし |
| 完全性への影響 | なし |
| 可用性への影響 | 高 |

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.16.4.linux-amd64.tar.gz`](https://golang.org/dl/go1.16.4.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.16.4.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.16.4.linux-amd64.tar.gz
$ sudo mv go go1.16.4
$ sudo ln -s go1.16.4 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.16.1 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-31525]: https://nvd.nist.gov/vuln/detail/CVE-2021-31525

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
