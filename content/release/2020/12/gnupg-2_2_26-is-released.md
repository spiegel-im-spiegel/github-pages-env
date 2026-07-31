+++
title = "GnuPG 2.2.26 がリリースされた"
date =  "2020-12-22T19:21:54+09:00"
description = "LDAP のサポートが改善され， Active Directory の基本的なサポートが追加された。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.26 がリリースされた。

このバージョンでは

{{< fig-quote type="markdown" title="[Announce] GnuPG 2.2.26 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2020q4/000451.html" lang="en" >}}
{{% quote %}} This is a maintenance release improving support for LDAP keyservers and enterprise use{{% /quote %}}.
{{< /fig-quote >}}

とのことで，さらに Twitter では

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">The <a href="https://twitter.com/hashtag/EU_Commission?src=hash&amp;ref_src=twsrc%5Etfw">#EU_Commission</a> might not like such software but here is another GnuPG release. This time with improved support for LDAP and basic support for Active Directory : <a href="https://t.co/QWlOjAwALO">https://t.co/QWlOjAwALO</a></p>&mdash; GNU Privacy Guard (@gnupg) <a href="https://twitter.com/gnupg/status/1341089812988768259?ref_src=twsrc%5Etfw">December 21, 2020</a></blockquote>
{{< /fig-gen >}}

とか tweet されていて，ちょっと笑ってしまった。
まぁ，今の OpenPGP 鍵サーバは明らかに設計限界に来てるからねぇ。
今後も注意して見ていこう。

その他の詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.26 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2020q4/000451.html" lang="en" >}}
* gpg: New AKL method "ntds".
* gpg: Fix --trusted-key with fingerprint arg.
* scd: Fix writing of ECC keys to an OpenPGP card.  [#5163]
* scd: Make an USB error fix specific to SPR532 readers.  [#5167]
* dirmngr: With new LDAP keyservers store the new attributes.  Never store the useless pgpSignerID.  Fix a long standing bug storing me keys on an ldap server.
* dirmngr: Support the new Active Direcory LDAP schema for keyservers.
* dirmngr: Allow LDAP OpenPGP searches via fingerprint.
* dirmngr: Do not block other threads during keyserver LDAP calls.
* Support global configuration files.  [#4788]
* Fix the iconv fallback handling to UTF-8.  [#5038]

Release-info: https://dev.gnupg.org/T5153
{{< /fig-quote >}}

## ビルド対象パッケージ

[自前でビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")する際の対象パッケージは以下の通り。

|   # | パッケージ名                                             | バージョン | 公開日     |         更新          |
| ---:| -------------------------------------------------------- | ---------- | ---------- |:---------------------:|
|   1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.41       | 2020-12-21 | {{< icons "check" >}} |
|   2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.7      | 2020-10-23 |                       |
|   3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.4      | 2020-10-23 |                       |
|   4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.5.0      | 2020-11-18 |                       |
|   5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6        | 2018-07-16 |                       |
|   6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0      | 2020-08-27 |                       |
|   7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.26     | 2020-12-21 | {{< icons "check" >}} |


アップデートは計画的に。

## ブックマーク

- [GnuPG and LDAP](https://gnupg.org/blog/20201018-gnupg-and-ldap.html) : [GnuPG] と LDAP との連携方法

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
