+++
title = "Go 1.19.2 のリリース【セキュリティ・アップデート】"
date =  "2022-10-07T20:08:24+09:00"
description = "今回は3件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

遅まきながらで申し訳ないが，[予告](https://groups.google.com/g/golang-announce/c/lx8kR68jRXE "[security] Go 1.19.2 and Go 1.18.7 pre-announcement")通りに先日 [Go] 1.19.2 がリリースされた。

- [[security] Go 1.19.2 and Go 1.18.7 are released](https://groups.google.com/g/golang-announce/c/xtuG5faxtaU)

今回は3件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://go.dev/doc/devel/release#go1.19.minor" lang="en" >}}
go1.19.2 (released 2022-10-04) includes security fixes to the `archive/tar`, `net/http/httputil`, and `regexp` packages, as well as bug fixes to the compiler, the linker, the runtime, and the `go/types` package. See the [Go 1.19.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.19.2+label%3ACherryPickApproved) on our issue tracker for details.
{{< /fig-quote >}}

## [CVE-2022-2879] archive/tar: unbounded memory consumption when reading headers

{{< fig-quote type="markdown" title="Go 1.19.2 and Go 1.18.7 are released" link="https://groups.google.com/g/golang-announce/c/xtuG5faxtaU" lang="en" >}}
`Reader.Read` did not set a limit on the maximum size of file headers.  A maliciously crafted archive could cause Read to allocate unbounded amounts of memory, potentially causing resource exhaustion or panics.  `Reader.Read` now limits the maximum size of header blocks to 1 MiB.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H`
- 深刻度: 重要 (Score: 7.5)

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

## [CVE-2022-2880] net/http/httputil: ReverseProxy should not forward unparseable query parameters

{{< fig-quote type="markdown" title="Go 1.19.2 and Go 1.18.7 are released" link="https://groups.google.com/g/golang-announce/c/xtuG5faxtaU" lang="en" >}}
Requests forwarded by `ReverseProxy` included the raw query parameters from the inbound request, including unparseable parameters rejected by `net/http`. This could permit query parameter smuggling when a Go proxy forwards a parameter with an unparseable value.

`ReverseProxy` will now sanitize the query parameters in the forwarded query when the outbound request's Form field is set after the `ReverseProxy.Director` function returns, indicating that the proxy has parsed the query parameters.  Proxies which do not parse query parameters continue to forward the original query parameters unchanged.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:N`
- 深刻度: 重要 (Score: 7.5)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | なし |
| 完全性への影響 | 高 |
| 可用性への影響 | なし |

## [CVE-2022-41715] regexp/syntax: limit memory used by parsing regexps

{{< fig-quote type="markdown" title="Go 1.19.2 and Go 1.18.7 are released" link="https://groups.google.com/g/golang-announce/c/xtuG5faxtaU" lang="en" >}}
The parsed `regexp` representation is linear in the size of the input, but in some cases the constant factor can be as high as 40,000, making relatively small regexps consume much larger amounts of memory.

Each `regexp` being parsed is now limited to a 256 MB memory footprint. Regular expressions whose representation would use more space than that are now rejected. Normal use of regular expressions is unaffected.
{{< /fig-quote >}}

$ go run main.go CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H
- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H`
- 深刻度: 重要 (Score: 7.5)

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

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.19.2.linux-amd64.tar.gz`](https://go.dev/dl/go1.19.2.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.19.2.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.19.2.linux-amd64.tar.gz
$ sudo mv go go1.19.2
$ sudo ln -s go1.19.2 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.19.2 linux/amd64
```

Windows は [Scoop] 経由で OK

アップデートは計画的に。

## 【2022-10-13 追記】 govulncheck コマンドが役に立った

最近 [govulncheck] を GitHub Actions のワークフローに仕込んでいるのだが，今回の脆弱性をちゃんと検出してくれたようだ。

{{< fig-img src="./github-actions.png" link="./github-actions.png" width="992" >}}

よーし，うむうむ，よーし。

- [Go 公式の脆弱性管理システム](https://zenn.dev/spiegel/articles/20220811-go-vulnerability-management)

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2022-2879]: https://nvd.nist.gov/vuln/detail/CVE-2022-2879
[CVE-2022-2880]: https://nvd.nist.gov/vuln/detail/CVE-2022-2880
[CVE-2022-41715]: https://nvd.nist.gov/vuln/detail/CVE-2022-41715
[govulncheck]: https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck "govulncheck command - golang.org/x/vuln/cmd/govulncheck - Go Packages"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
