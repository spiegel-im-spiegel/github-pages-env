+++
title = "【セキュリティ・アップデート】 golang.org/x/crypto/ssh"
date =  "2020-02-23T21:07:30+09:00"
description = "モジュール管理している場合はバージョンに要注意。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "security", "vulnerability", "package" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[`golang.org/x/crypto/ssh`] パッケージについてアップデートのアナウンスがあった。

- [[security] Vulnerability in golang.org/x/crypto/ssh - Google group](https://groups.google.com/forum/#!topic/golang-announce/3L45YRc91SY)

セキュリティ・アップデートを含んでいるので利用者は要注意である。

{{< fig-quote type="markdown" title="Vulnerability in golang.org/x/crypto/ssh" link="https://groups.google.com/forum/#!topic/golang-announce/3L45YRc91SY" lang="en" >}}
{{< quote >}}An attacker can craft an `ssh-ed25519` or `sk-ssh-ed25519@openssh.com` public key, such that the library will panic when trying to verify a signature with it. Clients can deliver such a public key and signature to any [`golang.org/x/crypto/ssh`](http://golang.org/x/crypto/ssh) server with a `PublicKeyCallback`, and servers can deliver them to any [`golang.org/x/crypto/ssh`](http://golang.org/x/crypto/ssh) client{{< /quote >}}.
{{< /fig-quote >}}

## [CVE-2020-9283](https://nvd.nist.gov/vuln/detail/CVE-2020-9283)

未稿

## go.mod に注意

他パッケージを `go.md` ファイルでモジュールとして管理している場合はバージョンに要注意。
とりあえず `latest` をセットして最新バージョンに差し替えればいいだろう。

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`golang.org/x/crypto/ssh`]: https://pkg.go.dev/golang.org/x/crypto/ssh "ssh package · go.dev"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
