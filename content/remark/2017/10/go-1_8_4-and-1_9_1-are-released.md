+++
title = "Go 1.8.4 および 1.9.1 がリリース（セキュリティ・アップデート）"
date =  "2017-10-06T17:09:52+09:00"
description = "今回は2件のセキュリティ・アップデートがある。アップデートは計画的に。"
tags = [
  "security",
  "vulnerability",
  "golang",
]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]コンパイラの 1.8.4 および 1.9.1 がリリースされている。

- [[security] Go 1.8.4 and Go 1.9.1 are released](https://groups.google.com/forum/#!topic/golang-announce/1hZYiemnkdE)

今回は2件のセキュリティ・アップデートがある。
ひとつは `go get` 時の問題で

{{< fig-quote  title="Go 1.8.4 and Go 1.9.1 are released" link="https://groups.google.com/forum/#!topic/golang-announce/1hZYiemnkdE" lang="en" >}}
<q>By nesting a git checkout inside another version control repository, it was possible for an attacker to trick the “go get” command into executing arbitrary code. The go command now refuses to use version control checkouts found inside other version control systems, with an exception for git submodules (git inside git).</q>
{{< /fig-quote >}}

これって CVE-2017-1000117 と似たような問題かなぁ？

- [SCM ツールの脆弱性]({{< ref "/remark/2017/08/vulnerabilities-in-scm-tools.md" >}})

バージョン管理ツールの方は既に改修版が出てるんだけど， [Go 言語]コンパイラも対応しましたよ，ってことなんだろうか。
（git の submodule などとは別に）こういう入れ子構造のリポジトリの問題って根が深そうだなぁ。

もうひとつは [`smtp`] パッケージに関する問題だ。

{{< fig-quote  title="Go 1.8.4 and Go 1.9.1 are released" link="https://groups.google.com/forum/#!topic/golang-announce/1hZYiemnkdE" lang="en" >}}
<q>In the smtp package, PlainAuth is documented as sending credentials only over authenticated, encrypted TLS connections, but it was changed in Go 1.1 to also send credentials on non-TLS connections when the remote server advertises that PLAIN authentication is supported. The change was meant to allow use of PLAIN authentication on localhost, but it has the effect of allowing a man-in-the-middle attacker to harvest credentials. PlainAuth now requires either TLS or a localhost connection before sending credentials, regardless of what the remote server claims.</q>
{{< /fig-quote >}}

アップデートは計画的に。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`smtp`]: https://golang.org/pkg/net/smtp/ "smtp - The Go Programming Language"
