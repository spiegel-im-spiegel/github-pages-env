+++
title = "GnuPG 2.1.22 リリースと寄付の募集"
date =  "2017-07-31T20:41:53+09:00"
description = "溜まりに溜まってた更新情報をまとめて挙げておく。脆弱性の習性あり。"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
  "vulnerability",
]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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

溜まりに溜まってた更新情報をまとめて挙げておく。

- [[Announce] GnuPG 2.1.21 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q2/000405.html)
- [[Announce] Libgcrypt 1.7.7 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q2/000406.html)
- [[Announce] Libgcrypt 1.7.8 released to fix CVE-2017-7526](https://lists.gnupg.org/pipermail/gnupg-announce/2017q2/000408.html)
    - [Sliding right into disaster: Left-to-right sliding windows leak](https://eprint.iacr.org/2017/627)
- [[Announce] Scute 1.5.0 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q3/000409.html)
- [[Announce] Libgcrypt 1.8.0 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q3/000410.html)
- [[Announce] GnuPG 2.1.22 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q3/000411.html)

[GnuPG] 2.1.21 リリース時にアナウンスされた注意事項は以下の通り。

{{< fig-quote title="GnuPG 2.1.21 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2017q2/000405.html" lang="en" >}}
<q>This release fixes a keyring corruption bug introduced with last release.  Users of 2.1.20, who are using the old "pubring.gpg" file to store their public keys, are asked to update to this new release.</q>
{{< /fig-quote >}}

また Libgcrypt 1.7.8 で “Sliding right into disaster" 脆弱性を軽減する修正が行われている。

{{< fig-quote title="Libgcrypt 1.7.8 released to fix CVE-2017-7526" link="https://lists.gnupg.org/pipermail/gnupg-announce/2017q2/000408.html" lang="en" >}}
<q>Note that this side-channel attack requires that the attacker can run arbitrary software on the hardware where the private RSA key is used.  Allowing execute access to a box with private keys should be considered as a game over condition, anyway.  Thus in practice there are easier ways to access the private keys than to mount this side-channel attack.  However, on boxes with virtual machines this attack may be used by one VM to steal private keys from another VM.</q>
{{< /fig-quote >}}

{{< fig-quote title="Sliding right into disaster: Left-to-right sliding windows leak" link="https://eprint.iacr.org/2017/627" lang="en" >}}
<q>It is well known that constant-time implementations of modular exponentiation cannot use sliding windows. However, software libraries such as Libgcrypt, used by GnuPG, continue to use sliding windows. It is widely believed that, even if the complete pattern of squarings and multiplications is observed through a side-channel attack, the number of exponent bits leaked is not sufficient to carry out a full key-recovery attack against RSA. Specifically, 4-bit sliding windows leak only 40% of the bits, and 5-bit sliding windows leak only 33% of the bits.</q>
{{< /fig-quote >}}

影響度は以下の通り。

“[CVE-2017-7526 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2017-7526)" より

CVSSv3 基本評価値 6.1 (`CVSS:3.0/AV:N/AC:H/PR:N/UI:R/S:C/C:H/I:N/A:N`)

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 高（H）           |
| 必要な特権レベル（PR）                  | 不要（N）         |
| ユーザ関与レベル（UI）                  | 要（R）           |
| スコープ（S）                           | 変更あり（C）     |
| 情報漏えいの可能性（機密性への影響, C） | 高（H）           |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）         |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |

CVSS については[解説ページ]({{< ref "/remark/2015/cvss-v3-metrics-in-jvn.md" >}})を参照のこと。

[論文]では 1024 bits RSA 鍵の解読を行っている。
今時 1024 bits RSA 鍵で運用している人はいないと思うが，もしまだの方がいたらこの機会に 2048 bits 以上にアップデートすることをお勧めする。
もしくは（可能であれば）楕円暗号に切り替えるか。

その他の修正等については各記事を参照のこと。
最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.1.22
libgcrypt 1.8.0
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

余談だが [GnuPG] では寄付を募っている。

- [GnuPG - Donate](https://gnupg.org/donate/)
- [[Announce] GnuPG Funding Campaign Launched](https://lists.gnupg.org/pipermail/gnupg-announce/2017q2/000407.html)

{{< fig-youtube id="wNHhkntqklg" width="500" height="281" title="GnuPG Fundraising Rally - YouTube" >}}

目標は15,000ユーロ/月。
この記事を書いている時点で目標の3分の1程度のペースか。
興味のある方は是非。

私も寄付したいけど，今は財政状況ががが。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[論文]: https://eprint.iacr.org/2017/627 "Sliding right into disaster: Left-to-right sliding windows leak"
