+++
title = "Go 1.20.1 のリリース【セキュリティ・アップデート】"
date =  "2023-02-15T09:58:08+09:00"
description = "golang.org/x/net/http2 v0.7.0 および golang.org/x/image v0.5.0 でも脆弱性の修正が行われている。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/2zqu2BOnpP8 "[security] Go 1.20.1 and Go 1.19.6 pre-announcement")通り 2023-02-15 (日本時間) に [Go] 1.20.1 がリリースされた。

- [[security] Go 1.20.1 and Go 1.19.6 are released](https://groups.google.com/g/golang-announce/c/V0aBFqaFs_E)

今回は4件の脆弱性修正を含んでいる。

## [CVE-2022-41722] path/filepath: path traversal in filepath.Clean on Windows

{{< fig-quote type="markdown" title="Go 1.20.1 and Go 1.19.6 are released" link="https://groups.google.com/g/golang-announce/c/V0aBFqaFs_E" lang="en" >}}
On Windows, the `filepath.Clean` function could transform an invalid path such as `a/../c:/b` into the valid path `c:\b`. This transformation of a relative (if invalid) path into an absolute path could enable a directory traversal attack. The filepath.Clean function will now transform this path into the relative (but still invalid) path `.\c:\b`.
{{< /fig-quote >}}

相変わらず Windows のパス構成は面倒くさいよな。

（以下未稿）

## [CVE-2022-41725] net/http, mime/multipart: denial of service from excessive resource consumption

文章が長いので，一部端折っている。
ごめんペコン。

{{< fig-quote type="markdown" title="Go 1.20.1 and Go 1.19.6 are released" link="https://groups.google.com/g/golang-announce/c/V0aBFqaFs_E" lang="en" >}}
Multipart form parsing with `mime/multipart.Reader.ReadForm` can consume largely unlimited amounts of memory and disk files. This also affects form parsing in the `net/http` package with the `Request` methods `FormFile`, `FormValue`, `ParseMultipartForm`, and `PostFormValue`.

[...]

`ReadForm` now properly accounts for various forms of memory overhead, and should now stay within its documented limit of 10MB + `maxMemory` bytes of memory consumption. Users should still be aware that this limit is high and may still be hazardous.

`ReadForm` now creates at most one on-disk temporary file, combining multiple form parts into a single temporary file. The `mime/multipart.File` interface type's documentation states, "If stored on disk, the `File`'s underlying concrete type will be an `*os.File`.". This is no longer the case when a form contains more than one file part, due to this coalescing of parts into a single file. The previous behavior of using distinct files for each form part may be reenabled with the environment variable `GODEBUG=multipartfiles=distinct`.

Users should be aware that `multipart.ReadForm` and the `http.Request` methods that call it do not limit the amount of disk consumed by temporary files. Callers can limit the size of form data with `http.MaxBytesReader`.
{{< /fig-quote >}}

とあるように，コンパイラを入れ替えれば OK とは必ずしもならないのでご注意を。

（以下未稿）

## [CVE-2022-41724] crypto/tls: large handshake records may cause panics

{{< fig-quote type="markdown" title="Go 1.20.1 and Go 1.19.6 are released" link="https://groups.google.com/g/golang-announce/c/V0aBFqaFs_E" lang="en" >}}
Both clients and servers may send large TLS handshake records which cause servers and clients, respectively, to panic when attempting to construct responses.

This affects all TLS 1.3 clients, TLS 1.2 clients which explicitly enable session resumption (by setting `Config.ClientSessionCache` to a non-nil value), and TLS 1.3 servers which request client certificates (by setting `Config.ClientAuth` >= `RequestClientCert`).
{{< /fig-quote >}}

（以下未稿）

## [CVE-2022-41723] net/http: avoid quadratic complexity in HPACK decoding

{{< fig-quote type="markdown" title="Go 1.20.1 and Go 1.19.6 are released" link="https://groups.google.com/g/golang-announce/c/V0aBFqaFs_E" lang="en" >}}
A maliciously crafted HTTP/2 stream could cause excessive CPU consumption in the HPACK decoder, sufficient to cause a denial of service from a small number of small requests.

This issue is also fixed in [`golang.org/x/net/http2`](http://golang.org/x/net/http2) v0.7.0, for users manually configuring HTTP/2.
{{< /fig-quote >}}

上述のとおり [`golang.org/x/net/http2`](http://golang.org/x/net/http2) v0.7.0 でも同様の修正がされているそうなので，心当たりの方はアップデートを。

（以下未稿）

## [CVE-2022-41727] Vulnerability in golang.org/x/image/tiff

標準パッケージではないが [`golang.org/x/image`](http://golang.org/x/image) v0.5.0 でも脆弱性の修正が行われている。

{{< fig-quote type="markdown" title="[security] Vulnerability in golang.org/x/image/tiff" link="https://groups.google.com/g/golang-announce/c/ag-FiyjlD5o" lang="en" >}}
Version v0.5.0 of [`golang.org/x/image`](http://golang.org/x/image) fixes a vulnerability in the [`golang.org/x/image/tiff`](http://golang.org/x/image/tiff) package which could cause a denial of service.

An attacker can craft a malformed TIFF image which will consume a significant amount of memory when passed to `DecodeConfig`.
{{< /fig-quote >}}

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.20.1.linux-amd64.tar.gz`](https://go.dev/dl/go1.20.1.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.20.1.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.20.1.linux-amd64.tar.gz
$ sudo mv go go1.20.1
$ sudo ln -s go1.20.1 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.20.1 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2022-41722]: https://nvd.nist.gov/vuln/detail/CVE-2022-41722
[CVE-2022-41725]: https://nvd.nist.gov/vuln/detail/CVE-2022-41725
[CVE-2022-41724]: https://nvd.nist.gov/vuln/detail/CVE-2022-41724
[CVE-2022-41723]: https://nvd.nist.gov/vuln/detail/CVE-2022-41723
[CVE-2022-41727]: https://nvd.nist.gov/vuln/detail/CVE-2022-41727

## ブックマーク

- [「Go 1.20.1」「Go 1.19.6」が公開、4件の脆弱性を修正 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1479118.html)

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
