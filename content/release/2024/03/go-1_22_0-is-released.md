+++
title = "Go 1.22 にアップデートする（セキュリティ・アップデートの追記あり）"
date =  "2024-03-03T11:25:30+09:00"
description = "仕事の忙しさにかまけて色々と放っぽり出してたので，少しずつ回復中。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

仕事の忙しさにかまけて色々と放っぽり出してたので，少しずつ回復中。

さて今更だが，先月 2024-02 に [Go] 1.22.0 が予定通りリリースされた。

- [Go 1.22.0 is released](https://groups.google.com/g/golang-announce/c/TpowDYVBMoY)
- [Go 1.22 is released! - The Go Programming Language](https://go.dev/blog/go1.22)
- [Go 1.22 Release Notes - The Go Programming Language](https://go.dev/doc/go1.22)

変化点の解説はページ末のブックマークを参照してもらうとして，とりあえず自機の環境をアップデートしてしまおうそうしよう。

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.22.1.linux-amd64.tar.gz`](https://go.dev/dl/go1.22.1.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.22.1.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.22.1.linux-amd64.tar.gz
$ sudo mv go go1.22.1
$ sudo ln -s go1.22.1 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.22.1 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/go1.22.1@latest
$ go1.22.1 download
$ go1.22.1 version
go version go1.22.1 linux/amd64
```

てな感じに導入できる。

## 【2024-03-06 追記】Go 1.22.1 のリリース【セキュリティ・アップデート】

予定通り [Go] 1.22.1 がリリースされた。
ついでに [`google.golang.org/protobuf`](http://google.golang.org/protobuf) パッケージにもアップデートのアナウンスがあった。
gRPC などを操ってる方は要注意である。

- [[security] Go 1.22.1 and Go 1.21.8 are released](https://groups.google.com/g/golang-announce/c/5pwGVUPoMbg)
- [[security] Vulnerability in google.golang.org/protobuf](https://groups.google.com/g/golang-announce/c/ArQ6CDgtEjY)

脆弱性の内容は以下の通り（面倒になったので CVSS 評価は端折る）。

- [CVE-2024-24783](https://nvd.nist.gov/vuln/detail/CVE-2024-24783) `crypto/x509`: Verify panics on certificates with an unknown public key algorithm
- [CVE-2023-45290](https://nvd.nist.gov/vuln/detail/CVE-2023-45290) `net/http`: memory exhaustion in `Request.ParseMultipartForm`
- [CVE-2023-45289](https://nvd.nist.gov/vuln/detail/CVE-2023-45289) `net/http`, `net/http/cookiejar`: incorrect forwarding of sensitive headers and cookies on HTTP redirect
- [CVE-2024-24785](https://nvd.nist.gov/vuln/detail/CVE-2024-24785) `html/template`: errors returned from `MarshalJSON` methods may break template escaping
- [CVE-2024-24784](https://nvd.nist.gov/vuln/detail/CVE-2024-24784) `net/mail`: comments in display names are incorrectly handled

- [CVE-2024-24786](https://nvd.nist.gov/vuln/detail/CVE-2024-24786) Version v1.33.0 of the [`google.golang.org/protobuf`](http://google.golang.org/protobuf) module fixes a bug in the [`google.golang.org/protobuf/encoding/protojson`](https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson) package which could cause the Unmarshal function to enter an infinite loop when handling some invalid inputs.

## ブックマーク

- [Go 1.22リリース連載始まります & ループの変化とTinyGo 0.31 | フューチャー技術ブログ](https://future-architect.github.io/articles/20240129a/)
- [Go 1.22 リリース連載 slicesのマイナーアップデート | フューチャー技術ブログ](https://future-architect.github.io/articles/20240130a/)
- [Go 1.22リリース連載 archive/tar, archive/zip, bufio, io | フューチャー技術ブログ](https://future-architect.github.io/articles/20240131a/)
- [Go 1.22 リリース連載 encoding, encoding/json | フューチャー技術ブログ](https://future-architect.github.io/articles/20240201a/)
- [Go1.22 リリース連載 HTTPルーティングの強化 | フューチャー技術ブログ](https://future-architect.github.io/articles/20240202a/)
- [Go 1.22リリース連載 vet, log/slog, testing/slogtest | フューチャー技術ブログ](https://future-architect.github.io/articles/20240205a/)
- [30種類のプログラミング言語で、ループ処理を書いてみた | フューチャー技術ブログ](https://future-architect.github.io/articles/20240206a/)
- [Go 1.22リリース連載 net, net/http, net/netip | フューチャー技術ブログ](https://future-architect.github.io/articles/20240214a/)
- [Go 1.22 でも残る、"syscall" を使わないといけないケース（ソース付き） - 標準愚痴出力](https://zetamatta.hatenablog.com/entry/2024/02/07/235657)
- [既存パッケージ(uncozip) を Go1.22 の rangefunc 対応に - 標準愚痴出力](https://zetamatta.hatenablog.com/entry/2024/02/19/095251)

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
