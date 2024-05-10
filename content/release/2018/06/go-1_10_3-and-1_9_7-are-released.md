+++
title = "Go 1.10.3 および 1.9.7 がリリース"
date = "2018-06-24T16:55:53+09:00"
description = "TLS および X.509 関連の改修と vgo 以降のための前準備？ かな。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "engineering", "versioning" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

3週間前の話でゴメンペコン。

[Go 言語]コンパイラの 1.10.3 と 1.9.7 がリリースされている。

- [Go 1.10.3 and Go 1.9.7 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/_S9YQriFKuU)
    - [Go 1.10.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.10.3)

{{< fig-quote title="Go 1.10.3 and Go 1.9.7 are released" link="https://groups.google.com/forum/#!topic/golang-announce/_S9YQriFKuU" lang="en" >}}
<q>These releases include fixes to the go command, and to the crypto/tls, crypto/x509, and strings packages. In particular, they add minimal support to the go command for the vgo transition.</q>
{{< /fig-quote >}}

ちうわけで Web 関連や暗号関連の製品を手がけている人はご注意を。
vgo サポートについての詳細は以下を参考にするといいだろう。

- [cmd/go: add minimal module-awareness for legacy operation](https://go.googlesource.com/go/+/d4e21288e444d3ffd30d1a0737f15ea3fc3b8ad9)

なお vgo の考え方については以下を参考にどうぞ。

- [vgo (Versioned Go) に関する覚え書き]({{< ref "/golang/go-and-versioning.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
