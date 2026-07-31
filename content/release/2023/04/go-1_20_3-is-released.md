+++
title = "Go 1.20.3 のリリース【セキュリティ・アップデート】"
date =  "2023-04-15T10:39:04+09:00"
description = "今回は4件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

毎度遅まきながらで申し訳ない。
[予告](https://groups.google.com/g/golang-announce/c/71Wg3N0IZk0 "[security] Go 1.20.3 and Go 1.19.8 pre-announcement")通り [Go] 1.20.3 がリリースされた。

- [[security] Go 1.20.3 and Go 1.19.8 are released](https://groups.google.com/g/golang-announce/c/Xdv6JL9ENs8)

今回は4件の脆弱性修正を含んでいる。

## [CVE-2023-24537] go/parser: infinite loop in parsing

{{< fig-quote type="markdown" title="Go 1.20.3 and Go 1.19.8 are released" link="https://groups.google.com/g/golang-announce/c/Xdv6JL9ENs8" lang="en" >}}
Calling any of the Parse functions on Go source code which contains //line directives with very large line numbers can cause an infinite loop due to integer overflow.
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

## [CVE-2023-24538] html/template: backticks not treated as string delimiters

{{< fig-quote type="markdown" title="Go 1.20.3 and Go 1.19.8 are released" link="https://groups.google.com/g/golang-announce/c/Xdv6JL9ENs8" lang="en" >}}
Templates did not properly consider backticks (`\`` ) as Javascript string delimiters, and as such did not escape them as expected. Backticks are used, since ES6, for JS template literals. If a template contained a Go template action within a Javascript template literal, the contents of the action could be used to terminate the literal, injecting arbitrary Javascript code into the Go template.

As ES6 template literals are rather complex, and themselves can do string interpolation, we've decided to simply disallow Go template actions from being used inside of them (e.g. "`var a = {{.}}`"), since there is no obviously safe way to allow this behavior. This takes the same approach as `github.com/google/safehtml`. Template.Parse will now return an Error when it encounters templates like this, with a currently unexported ErrorCode with a value of 12. This ErrorCode will be exported in the next major release.
{{< /fig-quote >}}

うわぁ，なにそれ面倒くさい。

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H`
- 深刻度: 緊急 (Score: 9.8)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |

## [CVE-2023-24534] net/http, net/textproto: denial of service from excessive memory allocation

{{< fig-quote type="markdown" title="Go 1.20.3 and Go 1.19.8 are released" link="https://groups.google.com/g/golang-announce/c/Xdv6JL9ENs8" lang="en" >}}
HTTP and MIME header parsing could allocate large amounts of memory, even when parsing small inputs.

Certain unusual patterns of input data could cause the common function used to parse HTTP and MIME headers to allocate substantially more memory than required to hold the parsed headers. An attacker can exploit this behavior to cause an HTTP server to allocate large amounts of memory from a small request, potentially leading to memory exhaustion and a denial of service.

Header parsing now correctly allocates only the memory required to hold parsed headers.
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

## [CVE-2023-24536] net/http, net/textproto, mime/multipart: denial of service from excessive resource consumption

ちょっと長いけど，そのまま載せるね。

{{< fig-quote type="markdown" title="Go 1.20.3 and Go 1.19.8 are released" link="https://groups.google.com/g/golang-announce/c/Xdv6JL9ENs8" lang="en" >}}
Multipart form parsing can consume large amounts of CPU and memory when processing form inputs containing very large numbers of parts. This stems from several causes:

`mime/multipart.Reader.ReadForm` limits the total memory a parsed multipart form can consume. `ReadForm` could undercount the amount of memory consumed, leading it to accept larger inputs than intended.
Limiting total memory does not account for increased pressure on the garbage collector from large numbers of small allocations in forms with many parts.
`ReadForm` could allocate a large number of short-lived buffers, further increasing pressure on the garbage collector.
The combination of these factors can permit an attacker to cause an program that parses multipart forms to consume large amounts of CPU and memory, potentially resulting in a denial of service. This affects programs that use `mime/multipart.Reader.ReadForm`, as well as form parsing in the `net/http` package with the `Request` methods `FormFile`, `FormValue`, `ParseMultipartForm`, and `PostFormValue`.

`ReadForm` now does a better job of estimating the memory consumption of parsed forms, and performs many fewer short-lived allocations.

In addition, `mime/multipart.Reader` now imposes the following limits on the size of parsed forms:

`Forms` parsed with `ReadForm` may contain no more than 1000 parts. This limit may be adjusted with the environment variable `GODEBUG=multipartmaxparts=`.
Form parts parsed with `NextPart` and `NextRawPart` may contain no more than 10,000 header fields. In addition, forms parsed with `ReadForm` may contain no more than 10,000 header fields across all parts. This limit may be adjusted with the environment variable `GODEBUG=multipartmaxheaders=`.
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

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.20.3.linux-amd64.tar.gz`](https://go.dev/dl/go1.20.3.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.20.3.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.20.3.linux-amd64.tar.gz
$ sudo mv go go1.20.3
$ sudo ln -s go1.20.3 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.20.3 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/go1.20.3@latest
$ go1.20.3 download
$ go1.20.3 version
go version go1.20.3 linux/amd64
```

てな感じで導入できる。

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2023-24534]: https://nvd.nist.gov/vuln/detail/CVE-2023-24534
[CVE-2023-24536]: https://nvd.nist.gov/vuln/detail/CVE-2023-24536
[CVE-2023-24537]: https://nvd.nist.gov/vuln/detail/CVE-2023-24537
[CVE-2023-24538]: https://nvd.nist.gov/vuln/detail/CVE-2023-24538

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
