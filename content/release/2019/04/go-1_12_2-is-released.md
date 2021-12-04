+++
title = "Go 1.12.2 がリリースされた"
date = "2019-04-07T14:15:13+09:00"
description = "コンパイラ本体は通常のアップデート。外部パッケージで脆弱性の情報あり。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.12.2 がリリースされた。
セキュリティ・アップデートはなし。

- [Go 1.12.2 and Go 1.11.7 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/z9eTD34GEIs)

{{< fig-quote title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.12.minor" lang="en" >}}
<q>go1.12.2 (released 2019/04/05) includes fixes to the compiler, the go command, the runtime, and the <code>doc</code>, <code>net</code>, <code>net/http/httputil</code>, and <code>os</code> packages. See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.2">Go 1.12.2 milestone</a> on our issue tracker for details.</q>
{{< /fig-quote >}}

また，このリリースより少し前に [`golang.org/x/crypto/salsa20`](https://golang.org/x/crypto/salsa20) パッケージに関する脆弱性情報がアナウンスされている。

- [[security] Vulnerability in golang.org/x/crypto/salsa20 - Google Group](https://groups.google.com/forum/#!topic/golang-announce/tjyNcJxb2vQ)

{{< fig-quote title="Vulnerability in golang.org/x/crypto/salsa20" link="https://groups.google.com/forum/#!topic/golang-announce/tjyNcJxb2vQ" lang="en" >}}
<q>If more than 256 GiB of keystream is generated, or if the counter otherwise grows greater than 32 bits, the amd64 implementation will first generate incorrect output, and then cycle back to previously generated keystream. Repeated keystream bytes can lead to loss of confidentiality in encryption applications, or to predictability in CSPRNG applications.</q>
{{< /fig-quote >}}

けっこうヤバい脆弱性なので，ご利用の方はアップデートを。
それ以外にも（メール本文にも目を通して）必要な措置があれば行うこと。

アップデートは計画的に。

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
