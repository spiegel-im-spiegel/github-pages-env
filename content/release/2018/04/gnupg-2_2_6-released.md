+++
title = "GnuPG 2.2.6 がリリース"
date = "2018-04-11T20:13:15+09:00"
update = "2018-04-22T18:15:40+09:00"
description = "今回もセキュリティ・アップデートはなし。"
image = "/images/attention/tools.png"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
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

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.6 がリリースされた。

- [[Announce] GnuPG 2.2.6 released](https://lists.gnupg.org/pipermail/gnupg-announce/2018q2/000421.html)

今回もセキュリティ・アップデートはなし。
主な修正点は以下の通り。

{{% fig-quote type="markdown" title="GnuPG 2.2.6 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2018q2/000421.html" lang="en" %}}
* gpg,gpgsm: New option `--request-origin` to pretend requests coming from a browser or a remote site.
* gpg: Fix race condition on trustdb.gpg updates due to too early released lock.  [#3839]
* gpg: Emit FAILURE status lines in almost all cases.  [#3872]
* gpg: Implement `--dry-run` for `--passwd` to make checking a key's passphrase straightforward.
* gpg: Make sure to only accept a certification capable key for key signatures.  [#3844]
* gpg: Better user interaction in `--card-edit` for the factory-reset sub-command.
* gpg: Improve changing key attributes in `--card-edit` by adding an explicit "key-attr" sub-command.  [#3781]
* gpg: Print the keygrips in the `--card-status`.
* scd: Support KDF DO setup.  [#3823]
* scd: Fix some issues with PC/SC on Windows.  [#3825]
* scd: Fix suspend/resume handling in the CCID driver.
* agent: Evict cached passphrases also via a timer.  [#3829]
* agent: Use separate passphrase caches depending on the request origin.  [#3858]
* ssh: Support signature flags.  [#3880]
* dirmngr: Handle failures related to missing IPv6 support gracefully.  [#3331]
* Fix corner cases related to specified home directory with drive letter on Windows.  [#3720]
* Allow the use of UNC directory names as homedir.  [#3818]
{{% /fig-quote %}}

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.6
libgcrypt 1.8.2
Copyright (C) 2018 Free Software Foundation, Inc.
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

## 追記 2018-04-22

Windows プラットフォーム用 GnuPG 環境である [Gpg4win] の 3.1.0 もリリースされた。

- [[Gpg4win-users-en] [Gpg4win-announce] Gpg4win 3.1.0 released](http://lists.wald.intevation.org/pipermail/gpg4win-users-en/2018-April/001486.html)
- [English README file for Gpg4win](https://files.gpg4win.org/README-3.1.0.en.txt)

[Gpg4win] 3.1.0 の構成は以下の通り。

- GnuPG:          2.2.6
- Kleopatra:      3.1.0
- GPA:            0.9.10
- GpgOL:          2.1.0
- GpgEX:          1.0.6

## ブックマーク

- [OpenPGP の実装](/openpgp/)

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
