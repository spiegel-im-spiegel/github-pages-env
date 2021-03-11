+++
title = "Go 1.16.1 のリリース【セキュリティ・アップデート】"
date =  "2021-03-11T18:52:53+09:00"
description = "今回は複数の脆弱性について改修されている。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/UERZo89zw8o "[security] Go 1.16.1 and Go 1.15.9 pre-announcement")通り， [Go] 1.16.1 がリリースされた。

- [[security] Go 1.16.1 and Go 1.15.9 are released](https://groups.google.com/g/golang-announce/c/MfiLYjG-RAw)

今回は複数の脆弱性について改修されている。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.16.minor" lang="en" >}}
{{% quote %}}go1.16.1 (released 2021/03/10) includes security fixes to the `archive/zip` and `encoding/xml` packages. See the [Go 1.16.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.1+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## encoding/xml: infinite loop when using xml.NewTokenDecoder with a custom TokenReader ([CVE-2021-27918])

{{< fig-quote type="markdown" title="Go 1.16.1 and Go 1.15.9 are released" link="https://groups.google.com/g/golang-announce/c/MfiLYjG-RAw" lang="en" >}}
{{% quote %}}The `Decode`, `DecodeElement`, and Skip methods of an `xml.Decoder` provided by `xml.NewTokenDecoder` may enter an infinite loop when operating on a custom `xml.TokenReader` which returns an EOF in the middle of an open XML element{{% /quote %}}.
{{< /fig-quote >}}

（以下未稿）

## archive/zip: panic when calling Reader.Open ([CVE-2021-27919])

{{< fig-quote type="markdown" title="Go 1.16.1 and Go 1.15.9 are released" link="https://groups.google.com/g/golang-announce/c/MfiLYjG-RAw" lang="en" >}}
{{% quote %}}The `Reader.Open` API, new in Go 1.16, will panic when used on a ZIP archive containing files that start with `"../"`{{% /quote %}}.
{{< /fig-quote >}}

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.16.1.linux-amd64.tar.gz`](https://golang.org/dl/go1.16.1.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.16.1.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.16.1.linux-amd64.tar.gz
$ sudo mv go go1.16.1
$ sudo ln -s go1.16.1 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.16.1 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-27918]: https://nvd.nist.gov/vuln/detail/CVE-2021-27918
[CVE-2021-27919]: https://nvd.nist.gov/vuln/detail/CVE-2021-27919

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
