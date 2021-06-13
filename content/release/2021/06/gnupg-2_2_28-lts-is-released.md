+++
title = "GnuPG 2.2.28 (LTS) がリリースされた"
date =  "2021-06-13T11:41:30+09:00"
description = "PGP 30周年記念リリース（笑）"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.28 がリリースされた。

- [[Announce] GnuPG 2.2.28 (LTS) released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q2/000460.html)

Twitter には

{{< fig-quote class="nobox" >}}
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">30 years after <a href="https://twitter.com/hashtag/PGP?src=hash&amp;ref_src=twsrc%5Etfw">#PGP</a> 1.0 we are proud to announce a new GnuPG LTS release: 2.2.28: <a href="https://t.co/5YWuC6wrDH">https://t.co/5YWuC6wrDH</a></p>&mdash; GNU Privacy Guard (@gnupg) <a href="https://twitter.com/gnupg/status/1403066139383566336?ref_src=twsrc%5Etfw">June 10, 2021</a></blockquote>
{{< /fig-quote >}}

とあるので [PGP 30周年]({{< ref "/remark/2021/06/pgp-30th.md" >}} "PGP は30周年だった")記念ということでいいだろう。
ちなみに今回はメンテンナンスのみでセキュリティ・アップデートはない。

詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.28 (LTS) released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q2/000460.html" lang="en" >}}
* gpg: Auto import keys specified with `--trusted-keys`. [e7251be84c79]
* gpg: Allow decryption w/o public key but with correct card inserted.  [e53f6037283e]
* gpg: Allow fingerprint based lookup with `--locate-external-key`. [2af217ecd7e4]
* gpg: Lookup a missing public key of the current card via LDAP. [b59af0e2a05a]
* gpg: New option `--force-sign-key`.  [#4584]
* gpg: Use a more descriptive password prompt for symmetric decryption.  [03f83bcda5d1]
* gpg: Do not use the `self-sigs-only` option for LDAP keyserver imports.  [#5387]
* gpg: Keep temp files when opening images via xdg-open. [0441ed6e1c]
* gpg: Fix mailbox based search via AKL keyserver method. [22fe23f46d31]
* gpg: Fix sending an OpenPGP key with umlaut to an LDAP keyserver. [7bf8530e75d0]
* gpg: Allow ECDH with a smartcard returning only the x-coordinate. [b203325ce1]
* gpgsm: New option `--ldapserver` as an alias for `--keyserver`.  Note that configuring servers in gpgsm and gpg is deprecated; please use the dirmngr configuration options.
* gpgsm: Support AES-GCM decryption.  [b722fd755c77]
* gpgsm: Support decryption of password protected files. [6f31acac767f]
* gpgsm: Lock keyboxes also during a search to fix lockups on Windows.  [#4505]
* agent: Skip unknown unknown ssh curves seen on cards. [bbf4bd3bfcb5]
* scdaemon: New option `--pcsc-shared`.  [5eec40f3d827]
* scdaemon: Backport PKCS#15 card support from GnuPG 2.3 [7637d39fe20e]
* scdaemon: Fix CCID driver for SCM SPR332/SPR532.  [#5297]
* scdaemon: Fix possible PC/SC removed card problem.  [9d83bfb63968]
* scdaemon: Fix unblock PIN by a Reset Code with KDF.  [#5413]
* scdaemon: Support compressed points.  [96577e2e46e4]
* scdaemon: Prettify S/N for Yubikeys and fix reading for early Yubikey 5 tokens.  [f8588369bcb0,#5442]
* dirmngr: New option `--ldapserver` to avoid the need for the separate `dirmngr_ldapservers.conf` file.
* dirmngr: The dirmngr_ldap wrapper has been rewritten to properly support ldap-over-tls and starttls for X.509 certificates and CRLs.  [39815c023f03]
* dirmngr: OpenPGP LDAP keyservers may now also be configured using the same syntax as used for X.509 and CRL LDAP servers.  This avoids the former cumbersome quoting rules and adds a flexible set of flags to control the connection.  [2b4cddf9086f]
* dirmngr: The "ldaps" scheme of an OpenPGP keyserver URL is now interpreted as ldap-with-starttls on port 389.  To use the non-standardized ldap-over-tls the new LDAP configuration method of the new attribute "gpgNtds" needs to be used.  [55f46b33df08]
* dirmngr: Return the fingerprint as search result also for LDAP OpenPGP keyservers.  This requires the modernized LDAP schema. [#5441]
* dirmngr: An OpenPGP LDAP search by a mailbox now ignores revoked keys.  [b6f8cd7eef4b]
* gpgconf: Make runtime changes with non-default homedir work. [c8f0b02936c7]
* gpgconf: Do not translate an empty string to the PO file's meta data.  [#5363]
* gpgconf: Fix argv overflow if `--homedir` is used.  [#5366]
* gpgconf: Return a new pseudo option "`compliance_de_vs`".  [9feffc03f364]
* gpgtar: Fix file size computation under Windows.  [198b240b1955]
* Full Unicode support for the Windows command line.  [#4398]
* Fix problem with Windows Job objects and auto start of our daemons.  [#4333]
* i18n: In German always use "Passwort" instead of "Passphrase" in prompts.

Release-info: https://dev.gnupg.org/T5482
{{< /fig-quote >}}

現在 [GnuPG] には2.2系と2.3系があり，2.2系は LTS 版に位置づけられている。
2.3系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されているので，最新機能を試したいのであればこちらを入れるとよいだろう。
通常運用であれば2.2系で問題ない（[ECC も対応]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}} "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな")してるよ）。

[自前でビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")する際の対象パッケージは以下の通り。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.42         | 2021-03-22 |                       |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.8 (LTS)  | 2021-06-02 | {{< icons "check" >}} |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.9.3        | 2021-04-19 |                       |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.0        | 2021-06-10 | {{< icons "check" >}} |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0        | 2020-08-27 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.28 (LTS) | 2021-06-10 | {{< icons "check" >}} |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.1        | 2021-04-20 |                       |

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
