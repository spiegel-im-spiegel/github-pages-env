+++
title = "Go 1.11.5 のリリース 【セキュリティ・アップデート】"
date = "2019-01-26T19:08:12+09:00"
description = "今回は楕円曲線暗号まわりの脆弱性が修正されている。アップデートは計画的に。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "security", "vulnerability" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]コンパイラの 1.11.5 および 1.10.8 がリリースされた。
今回はセキュリティ・アップデートを含んでいるので必ずアップデートすること。

- [[security] Go 1.11.5 and Go 1.10.8 are released - Google group](https://groups.google.com/forum/#%21topic/golang-announce/mVeX35iXuSw)
- [crypto/elliptic: CPU DoS vulnerability affecting P-521 and P-384 · Issue #29903 · golang/go · GitHub](https://github.com/golang/go/issues/29903)

以下のインシデントに対して修正が行われている。

{{< fig-quote title="Go 1.11.5 and Go 1.10.8 are released" link="https://groups.google.com/forum/#%21topic/golang-announce/mVeX35iXuSw" lang="en" >}}
<q>This DoS vulnerability in the crypto/elliptic implementations of the P-521 and P-384 elliptic curves may let an attacker craft inputs that consume excessive amounts of CPU.<br>
These inputs might be delivered via TLS handshakes, X.509 certificates, JWT tokens, ECDH shares or ECDSA signatures. In some cases, if an ECDH private key is reused more than once, the attack can also lead to key recovery.</q>
{{< /fig-quote >}}

## 想定される影響

### [CVE-2019-6486](https://nvd.nist.gov/vuln/detail/CVE-2019-6486)

`CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H` (深刻度7.5) : 

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更なし     |
| 機密性への影響   | なし         |
| 完全性への影響   | なし         |
| 可用性への影響   | 高           |

可用性への影響が高いとしているが，とりあえず鍵の漏洩とかは心配しなくていい？

アップデートは計画的に。

[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
