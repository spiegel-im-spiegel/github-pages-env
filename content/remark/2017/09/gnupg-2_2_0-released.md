+++
title = "GnuPG 2.2.0 がリリース（脆弱性の修正あり）"
date =  "2017-09-14T04:11:20+09:00"
description = "バージョン 2.2 は stable 版 2.0.x の置き換えになるようだ。なお 2.0.x は今年いっぱいでサポートを終了する。"
tags = [
  "security",
  "cryptography",
  "vulnerability",
  "tools",
  "openpgp",
  "gnupg",
]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
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
+++

個人的に色々あって更新が滞ってます。
ゴメン，ペコン。

- [[Announce] GnuPG 2.2.0 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q3/000413.html)

バージョン 2.2 は stable 版 2.0.x の置き換えになるようだ。
なお 2.0.x は今年いっぱいでサポートを終了する。

{{< fig-quote title="GnuPG 2.2.0 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2017q3/000413.html" lang="en" >}}
<q>This release marks the start of a new long term support series to replace the 2.0.x series which will reach end-of-life on 2017-12-31.</q>
{{< /fig-quote >}}

今までのように modarn 版と stable 版に分けて開発を進めるのかどうかについては記述がないので分からない。
2.1.23 からの変更点は以下の通り。

* gpg: Reverted change done in 2.1.23 so that `--no-auto-key-retrieve` is again the default.
* Fixed a few minor bugs.

ところで GnuPG が内部で使っている Libgcrypt に脆弱性が発見されている。

- [[Announce] Libgcrypt 1.8.1 and 1.7.9 to fix CVE-2017-0379](https://lists.gnupg.org/pipermail/gnupg-announce/2017q3/000414.html)
- [May the Fourth Be With You: A Microarchitectural Side Channel Attack on Several Real-World Applications of Curve25519](https://eprint.iacr.org/2017/806)

（スター・ウォーズかよ！）

{{< fig-quote title="A Microarchitectural Side Channel Attack on Several Real-World Applications of Curve25519" link="https://eprint.iacr.org/2017/806" lang="en" >}}
<q>We demonstrate the effect of this vulnerability on three software applications---encrypted git, email and messaging---that use Libgcrypt. In each case, we show how to craft malicious OpenPGP files that use the Curve25519 point of order 4 as a chosen ciphertext to the ECDH encryption scheme. We find that the resulting interactions of the point at infinity, order-2, and order-4 elements in the Montgomery ladder scalar-by-point multiplication routine create side channel leakage that allows us to recover the private key in as few as 11 attempts to access such malicious files.</q>
{{< /fig-quote >}}

たとえば暗号化メールを使った攻撃では

{{< fig-quote title="A Microarchitectural Side Channel Attack on Several Real-World Applications of Curve25519" link="https://eprint.iacr.org/2017/806" lang="en" >}}
<q>For an attack on encrypted email we use the Thunderbird plugin Enigmail. As Genkin et al. observe, Enigmail automatically decrypts incoming emails by passing them to GnuPG, which uses Libgcrypt as its cryptographic engine. To attack Enigmail, we inject an element of order 4 into Libgcrypt we send the victim a PGP/MIME-encoded e-mail, with the element of order-4 as the ciphertext.</q>
{{< /fig-quote >}}

といったシナリオが挙げられている。

この問題を軽減（mitigate）したバージョン 1.8.1 および 1.7.9 がリリースされている。
なお Windows 用の GnuPG 2.2.0 バイナリには Libgcrypt 1.8.1 が同梱されている。

影響度は以下の通り。

- [NVD - CVE-2017-0379](https://nvd.nist.gov/vuln/detail/CVE-2017-0379)

**CVSSv3 基本評価値 7.5 (`CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 低（L）           |
| 必要な特権レベル（PR）                  | 不要（N）         |
| ユーザ関与レベル（UI）                  | 不要（N）         |
| スコープ（S）                           | 変更なし（U）     |
| 情報漏えいの可能性（機密性への影響, C） | 高（H）           |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）         |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |

一般にサイドチャネル攻撃はあまり影響度が高くないが [NVD] はかなり高めに見積もっているようである[^rh1]。

[^rh1]: 一方， [Red Hat の評価](https://access.redhat.com/security/cve/cve-2017-0379)は `CVSS:3.0/AV:L/AC:H/PR:L/UI:R/S:U/C:H/I:N/A:N` で基本評価値 4.4 になっている。ただし[論文]を読むと攻撃シナリオとして Enigmail, Git-crypt, Pidgin-OpenPGP を使った方法が挙げられているので， [NVD] の評価のほうが近いかも。

CVSS については[解説ページ]({{< ref "/remark/2015/cvss-v3-metrics-in-jvn.md" >}})を参照のこと。

インストールが上手くいけば以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.0
libgcrypt 1.8.1
Copyright (C) 2017 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: ********
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
    CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

アップデートは計画的に。

[NVD]: https://nvd.nist.gov/ "National Vulnerability Database"
[論文]: https://eprint.iacr.org/2017/806 "May the Fourth Be With You: A Microarchitectural Side Channel Attack on Several Real-World Applications of Curve25519"
