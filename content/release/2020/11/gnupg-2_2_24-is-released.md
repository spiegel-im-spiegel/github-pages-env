+++
title = "GnuPG 2.2.24 がリリースされた"
date =  "2020-11-18T18:53:19+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.24 がリリースされた。

- [[Announce] GnuPG 2.2.24 released](https://lists.gnupg.org/pipermail/gnupg-announce/2020q4/000449.html)

メンテナンス・リリース。
セキュリティ・アップデートはなし。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.24 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2020q4/000449.html" lang="en" >}}
* Allow Unicode file names on Windows almost everywhere.  Note that it is still not possible to use Unicode strings on the command line.  This change also fixes a regression in 2.2.22 related to non-ascii file names.  [#5098]
* Fix localized time printing on Windows.  [#5073]
* gpg: New command `--quick-revoke-sig`.  [#5093]
* gpg: Do not use weak digest algos if selected by recipient preference during sign+encrypt.  [4c181d51a6]
* gpg: Switch to AES256 for symmetric encryption in de-vs mode. [166e779634]
* gpg: Silence weak digest warnings with `--quiet`.  [#4893]
* gpg: Print new status line CANCELED_BY_USER for a cancel during symmetric encryption.  [f05d1772c4]
* gpg: Fix the encrypt+sign hash algo preference selection for ECDSA.  This is in particular needed for keys created from existing smartcard based keys.  [aeed0b93ff]
* agent: Fix secret key import of GnuPG 2.3 generated Ed25519 keys. [#5114]
* agent: Keep some permissions of private-keys-v1.d.  [#2312]
* dirmngr: Align sks-keyservers.netCA.pem use between ntbtls and gnutls builds.  [e4f3b74c91]
* dirmngr: Fix the pool keyserver case for a single host in the pool.  [72e04b03b1a7]
* scd: Fix the use case of verify_chv2 by CHECKPIN.  [61aea64b3c]
* scd: Various improvements to the ccid-driver.  [#4616,#5065]
* scd: Minor fixes for Yubikey [25bec16d0b]
* gpgconf: New option `--show-versions`.
* w32: Install gpg-check-pattern and example profiles.  Install Windows subsystem variant of gpgconf (gpgconf-w32).
* i18n: Complete overhaul and completion of the Italian translation. Thanks to Denis Renzi.
* Require Libgcrypt 1.8 because 1.7 has long reached end-of-life.

Release-info: https://dev.gnupg.org/T5052
{{< /fig-quote >}}

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
