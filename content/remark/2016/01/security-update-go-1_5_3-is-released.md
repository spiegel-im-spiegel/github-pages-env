+++
date = "2016-01-24T15:24:23+09:00"
update = "2016-01-24T16:29:20+09:00"
description = "リスクとしては大したことはないですが， Web アプリケーションまたは TLS の実装を Go 言語で行っている方は要更新です。"
draft = false
tags = ["golang", "security", "vulnerability", "math", "tls"]
title = "[security] Go 1.5.3 is released"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

2週間前の話ですみません。
1月14日に [Go 言語]コンパイラ 1.5.3 が出ています。
メインは [`math/big`] パッケージの不具合修正です。

リスクとしては大したことはないですが， Web アプリケーションまたは TLS の実装を [Go 言語]で行っている方は要更新です。

{{< fig-quote title="[security] Go 1.5.3 is released - Google グループ" link="https://groups.google.com/forum/#!topic/golang-announce/MEATuOi_ei4" lang="en" >}}
<q>This issue can affect RSA computations in crypto/rsa, which is used by crypto/tls. TLS servers on 32-bit systems could plausibly leak their RSA private key due to this issue. Other protocol implementations that create many RSA signatures could also be impacted in the same way.<br>
Specifically, incorrect results in one part of the RSA Chinese Remainder computation can cause the result to be incorrect in such a way that it leaks one of the primes. While RSA blinding should prevent an attacker from crafting specific inputs that trigger the bug, on 32-bit systems the bug can be expected to occur at random around one in 2^26 times. Thus collecting around 64 million signatures (of known data) from an affected server should be enough to extract the private key used.<br>
On 64-bit systems, the frequency of the bug is so low (less than one in 2^50) that it would be very difficult to exploit. Nonetheless, everyone is strongly encouraged to upgrade.</q>
{{< /fig-quote >}}

## CVE-2015-8618 Carry propagation in Int.Exp Montgomery code in math/big library

**CVSSv2 基本評価値 2.6 (`AV:N/AC:H/Au:N/C:P/I:N/A:N`)**
（[Redhat による評価](https://access.redhat.com/security/cve/cve-2015-8618)）

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 高（H）           |
| 攻撃前の認証要否（Au）                  | 不要（N）         |
| 情報漏えいの可能性（機密性への影響, C） | 部分的（P）       |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）         |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`math/big`]: https://golang.org/pkg/math/big/ "big - The Go Programming Language"
