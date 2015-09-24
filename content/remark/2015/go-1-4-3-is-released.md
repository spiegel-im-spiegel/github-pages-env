+++
date = "2015-09-25T00:42:22+09:00"
description = "Go 言語の net/http パッケージに脆弱性が発見された模様。"
draft = false
tags = ["security", "vulnerability", "golang"]
title = "セキュリティ脆弱性を修正した Go 1.4.3 がリリース"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

{{< fig-quote lang="en" title="Go 1.4.3 is released (security fix) - Google グループ" link="https://groups.google.com/forum/#!topic/golang-announce/iSIyW4lM4hY" >}}
The issues were reported in Go's net/http package. They affect programs using that package to proxy HTTP requests. We recommend that all users upgrade to Go 1.5, which fixes these issues. For users unable to upgrade to Go 1.5, we have released version 1.4.3, which is based on Go 1.4.2 plus fixes for these issues. Affected Go programs—those that use the net/http package as a proxy server—must be recompiled with Go 1.5 or Go 1.4.3 to receive the fixes. 
{{< /fig-quote >}}

というわけで [Go 言語]の [`net/http`] パッケージに脆弱性が発見された模様。

[`net/http`] パッケージを使っている製品は最新バージョンで再コンパイルすること（パッケージ間の依存関係に注意）。
可能であれば 1.5 系を使うのが望ましい。
諸事情で 1.5 系が使えない場合は，リリースされた 1.4.3 を使ってもよい。

## 影響度（CVSS）

CVSS 基本評価値 6.8 (AV:N/AC:M/Au:N/C:P/I:P/A:P) （暫定値）

- [access.redhat.com | CVE-2015-5739](https://access.redhat.com/security/cve/CVE-2015-5739)
- [access.redhat.com | CVE-2015-5740](https://access.redhat.com/security/cve/CVE-2015-5740)
- [access.redhat.com | CVE-2015-5741](https://access.redhat.com/security/cve/CVE-2015-5741)

{{% fig-gen title="CVSSv2 基本評価値" %}}
| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 中（M）           |
| 攻撃前の認証要否（Au）                  | 不要（N）         |
| 情報漏えいの可能性（機密性への影響, C） | 部分的（P）       |
| 情報改ざんの可能性（完全性への影響, I） | 部分的（P）       |
| 業務停止の可能性（可用性への影響, A）   | 部分的（P）       |
{{% /fig-gen %}}

CVSS については[デモページ](http://www.baldanders.info/spiegel/archive/cvss/cvss2.html)を参照のこと。

## 参考

- [プログラミング言語 Go - text.Baldanders.info](/golang)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`net/http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"
