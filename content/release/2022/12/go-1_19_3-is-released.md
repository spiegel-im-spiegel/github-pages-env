+++
title = "Go 1.19.4 のリリース【セキュリティ・アップデート】"
date =  "2022-12-07T08:17:43+09:00"
description = "今回は2件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/v2KvT18_yCg "[security] Go 1.19.4 and Go 1.18.9 pre-announcement")通り 2022-12-07 (日本時間) に [Go] 1.19.4 がリリースされた。

- [[security] Go 1.19.4 and Go 1.18.9 are released](https://groups.google.com/g/golang-announce/c/L_3rmdT0BMU)

今回は2件の脆弱性修正を含んでいる。

## [CVE-2022-41720] os, net/http: avoid escapes from os.DirFS and http.Dir on Windows

{{< fig-quote type="markdown" title="Go 1.19.4 and Go 1.18.9 are released" link="https://groups.google.com/g/golang-announce/c/L_3rmdT0BMU" lang="en" >}}
The `os.DirFS` function and `http.Dir` type provide access to a tree of files rooted at a given directory. These functions permitted access to Windows device files under that root. For example, `os.DirFS("C:/tmp").Open("COM1")` would open the `COM1` device. Both `os.DirFS` and `http.Dir` only provide read-only filesystem access.

In addition, on Windows, an `os.DirFS` for the directory `\`(the root of the current drive) can permit a maliciously crafted path to escape from the drive and access any path on the system.
{{< /fig-quote >}}

（以下未稿）

## [CVE-2022-41717] net/http: limit canonical header cache by bytes, not entries

{{< fig-quote type="markdown" title="Go 1.19.4 and Go 1.18.9 are released" link="https://groups.google.com/g/golang-announce/c/L_3rmdT0BMU" lang="en" >}}
An attacker can cause excessive memory growth in a Go server accepting HTTP/2 requests.

HTTP/2 server connections contain a cache of HTTP header keys sent by the client. While the total number of entries in this cache is capped, an attacker sending very large keys can cause the server to allocate approximately 64 MiB per open connection.

This issue is also fixed in [`golang.org/x/net/http2`](http://golang.org/x/net/http2) `vX.Y.Z`, for users manually configuring HTTP/2.


The `os.DirFS` function and `http.Dir` type provide access to a tree of files rooted at a given directory. These functions permitted access to Windows device files under that root. For example, `os.DirFS("C:/tmp").Open("COM1")` would open the `COM1` device. Both `os.DirFS` and `http.Dir` only provide read-only filesystem access.

In addition, on Windows, an `os.DirFS` for the directory `\`(the root of the current drive) can permit a maliciously crafted path to escape from the drive and access any path on the system.
{{< /fig-quote >}}

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.19.4.linux-amd64.tar.gz`](https://go.dev/dl/go1.19.4.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.19.4.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.19.4.linux-amd64.tar.gz
$ sudo mv go go1.19.4
$ sudo ln -s go1.19.4 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.19.4 linux/amd64
```

Windows は [Scoop] 経由で OK

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2022-41720]: https://nvd.nist.gov/vuln/detail/CVE-2022-41720
[CVE-2022-41717]: https://nvd.nist.gov/vuln/detail/CVE-2022-41717

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
