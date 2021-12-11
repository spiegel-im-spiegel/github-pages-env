+++
title = "Go 1.17.5 のリリース【セキュリティ・アップデート】"
date =  "2021-12-11T18:20:44+09:00"
description = "今回は2件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.17.5 がリリースされた。

- [[security] Go 1.17.5 and Go 1.16.12 are released](https://groups.google.com/g/golang-announce/c/hcmEScgc00k)

今回は2件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release#go1.17.minor" lang="en" >}}
{{% quote %}}go1.17.5 (released 2021-12-09) includes security fixes to the `syscall` and `net/http` packages. See the [Go 1.17.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.5+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2021-44716]: net/http: limit growth of header canonicalization cache

{{< fig-quote type="markdown" title="Go 1.17.5 and Go 1.16.12 are released" link="https://groups.google.com/g/golang-announce/c/hcmEScgc00k" lang="en" >}}
An attacker can cause unbounded memory growth in a Go server accepting HTTP/2 requests.

For users who cannot immediately update to the new release, setting the `GODEBUG=http2server=0` environment variable before calling Serve will disable HTTP/2 unless it was manually configured through the [`golang.org/x/net/http2`](http://golang.org/x/net/http2) package.

This issue is also fixed in [`golang.org/x/net/http2`](http://golang.org/x/net/http2) `v0.0.0-20211209124913-491a49abca63`, for users manually configuring HTTP/2.
{{< /fig-quote >}}

というわけで [`golang.org/x/net/http2`](http://golang.org/x/net/http2) パッケージを使っている場合は，こちらも要アップデートだな。

（以下未稿）

## [CVE-2021-44717]: syscall: don’t close fd 0 on ForkExec error

{{< fig-quote type="markdown" title="Go 1.17.5 and Go 1.16.12 are released" link="https://groups.google.com/g/golang-announce/c/hcmEScgc00k" lang="en" >}}
When a Go program running on a Unix system is out of file descriptors and calls `syscall.ForkExec` (including indirectly by using the `os/exec` package), `syscall.ForkExec` can close file descriptor 0 as it fails. If this happens (or can be provoked) repeatedly, it can result in misdirected I/O such as writing network traffic intended for one connection to a different connection, or content intended for one file to a different one.

For users who cannot immediately update to the new release, the bug can be mitigated by raising the per-process file descriptor limit.
{{< /fig-quote >}}

これはちょっとヤバいかな。

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.17.5.linux-amd64.tar.gz`](https://go.dev/dl/go1.17.5.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.17.5.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.17.5.linux-amd64.tar.gz
$ sudo mv go go1.17.5
$ sudo ln -s go1.17.5 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.17.5 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-44716]: https://nvd.nist.gov/vuln/detail/CVE-2021-44716
[CVE-2021-44717]: https://nvd.nist.gov/vuln/detail/CVE-2021-44717

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
