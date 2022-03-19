+++
title = "golang.org/x/crypto/ssh パッケージをアップデートしましょう【セキュリティ・アップデート】"
date =  "2022-03-19T14:17:32+09:00"
description = "CVE-2022-27191 の改修を含む。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "security", "vulnerability", "openssh", "package", "sha-1" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

`golang.org/x/crypto/ssh` パッケージをアップデートするようアナウンスが出ている。

- [An update of golang.org/x/crypto/ssh might be necessary](https://groups.google.com/g/golang-announce/c/-cp44ypCT5s)

ひとつは SSH 認証時の問題：

{{< fig-quote type="markdown" title="An update of golang.org/x/crypto/ssh might be necessary" link="https://groups.google.com/g/golang-announce/c/-cp44ypCT5s" lang="en" >}}
Version v0.0.0-20220315160706-3147a52a75dd of [`golang.org/x/crypto/ssh`](http://golang.org/x/crypto/ssh) implements client authentication support for signature algorithms based on SHA-2 for use with existing RSA keys.

Previously, a client would fail to authenticate with RSA keys to servers that reject signature algorithms based on SHA-1. This includes [OpenSSH 8.8](https://www.openssh.com/txt/release-8.8) by default and—[starting today March 15, 2022](https://github.blog/changelog/2022-03-15-removed-unencrypted-git-protocol-and-certain-ssh-keys/)—[github.com](http://github.com/) for recently uploaded keys.
{{< /fig-quote >}}

もうひとつは `golang.org/x/crypto/ssh` パッケージを使って SSH サービスを構成する場合の脆弱性で [CVE-2022-27191] が割り当てられている。

{{< fig-quote type="markdown" title="An update of golang.org/x/crypto/ssh might be necessary" link="https://groups.google.com/g/golang-announce/c/-cp44ypCT5s" lang="en" >}}
Version v0.0.0-20220314234659-1baeb1ce4c0b (included in the version above) also fixes a potential security issue where an attacker could cause a crash in a [`golang.org/x/crypto/ssh`](http://golang.org/x/crypto/ssh) server under these conditions:

- The server has been configured by passing a [`Signer`](https://pkg.go.dev/golang.org/x/crypto/ssh#Signer) to [`ServerConfig.AddHostKey`](https://pkg.go.dev/golang.org/x/crypto/ssh#ServerConfig.AddHostKey).
- The Signer passed to `AddHostKey` does not also implement [`AlgorithmSigner`](https://pkg.go.dev/golang.org/x/crypto/ssh#AlgorithmSigner).
- The Signer passed to `AddHostKey` does return a key of type “`ssh-rsa`” from its PublicKey method.

Servers that only use `Signer` implementations provided by the ssh package are unaffected.
{{< /fig-quote >}}

`golang.org/x/crypto/ssh` パッケージを使って SSH サービスを利用している方はアップデートしましょう。

## ブックマーク

- [さようなら SHA-1](https://zenn.dev/spiegel/articles/20201025-sayonara-sha1)

[CVE-2022-27191]: https://nvd.nist.gov/vuln/detail/CVE-2022-27191

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
