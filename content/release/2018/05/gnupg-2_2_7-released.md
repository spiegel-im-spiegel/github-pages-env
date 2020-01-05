+++
title = "GnuPG 2.2.7 がリリース"
date = "2018-05-07T19:25:25+09:00"
update = "2018-08-18T15:41:48+09:00"
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

[GnuPG] 2.2.7 がリリースされた。

- [[Announce] GnuPG 2.2.7 released](https://lists.gnupg.org/pipermail/gnupg-announce/2018q2/000424.html)

今回もセキュリティ・アップデートはなし。
主な修正点は以下の通り。

{{% fig-quote type="markdown" title="GnuPG 2.2.7 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2018q2/000421.html" lang="en" %}}
* gpg: New option `--no-symkey-cache` to disable the passphrase cache for symmetrical en- and decryption.
* gpg: The ERRSIG status now prints the fingerprint if that is part of the signature.
* gpg: Relax emitting of FAILURE status lines
* gpg: Add a status flag to "sig" lines printed with `--list-sigs`.
* gpg: Fix "Too many open files" when using `--multifile`.  [#3951]
* ssh: Return an error for unknown `ssh-agent` flags.  [#3880]
* dirmngr: Fix a regression since 2.1.16 which caused corrupted CRL caches under Windows.  [#2448,#3923]
* dirmngr: Fix a CNAME problem with pools and TLS.  Also use a fixed mapping of keys.gnupg.net to sks-keyservers.net.  [#3755]
* dirmngr: Try resurrecting dead hosts earlier (from 3 to 1.5 hours).
* dirmngr: Fallback to CRL if no default OCSP responder is configured.
* dirmngr: Implement CRL fetching via https.  Here a redirection to http is explictly allowed.
* dirmngr: Make LDAP searching and CRL fetching work under Windows.  This stopped working with 2.1.  [#3937]
* agent,dirmngr: New sub-command "getenv" for "getinfo" to ease debugging.
{{% /fig-quote %}}

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.7
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

なお [Gpg4win] についてもアップデートが行われている。

- [[Gpg4win-users-en] [Gpg4win-announce] Gpg4win 3.1.1 released](http://lists.wald.intevation.org/pipermail/gpg4win-users-en/2018-May/001493.html)
- [English README file for Gpg4win](https://files.gpg4win.org/README-3.1.1.en.txt)

[Gpg4win] 3.1.1 の構成は以下の通り。

- GnuPG:          2.2.7
- Kleopatra:      3.1.1
- GPA:            0.9.10
- GpgOL:          2.1.1
- GpgEX:          1.0.6

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< ref "/openpgp" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
