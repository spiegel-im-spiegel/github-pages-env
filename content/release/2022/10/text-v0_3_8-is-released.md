+++
title = "golang.org/x/text/language のセキュリティ・アップデート"
date =  "2022-10-13T19:38:30+09:00"
description = "1件の脆弱性の修正。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[`golang.org/x/text`] パッケージの v0.3.8 がリリースされている。
[`golang.org/x/text/language`] パッケージの脆弱性が修正されている。

- [[security] Vulnerability in golang.org/x/text/language](https://groups.google.com/g/golang-announce/c/-hjNw559_tE)

{{< fig-quote type="markdown" title="Vulnerability in golang.org/x/text/language" link="https://groups.google.com/g/golang-announce/c/-hjNw559_tE" lang="en" >}}
Version v0.3.8 of [`golang.org/x/text`](http://golang.org/x/text) fixes a vulnerability in the [`golang.org/x/text/language`](http://golang.org/x/text/language) package which could cause a denial of service.

An attacker can craft an `Accept-Language` header which `ParseAcceptLanguage` will take significant time to parse.
{{< /fig-quote >}}

Web 周りの実装をしている人でリクエストヘッダの `Accept-Language` を評価している場合は要注意かな。

## [CVE-2022-32149]

（未稿）

[Go]: https://go.dev/
[`golang.org/x/text`]: https://pkg.go.dev/golang.org/x/text "text package - golang.org/x/text - Go Packages"
[`golang.org/x/text/language`]: http://golang.org/x/text/language "language package - golang.org/x/text/language - Go Packages"
[CVE-2022-32149]: https://nvd.nist.gov/vuln/detail/CVE-2022-32149

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
