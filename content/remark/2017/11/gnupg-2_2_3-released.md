+++
title = "GnuPG 2.2.3 がリリースされた"
date =  "2017-11-21T17:30:50+09:00"
update =  "2017-11-24T10:01:22+09:00"
image = "/images/attention/remark.jpg"
description = "今回もセキュリティ・アップデートはなし。平和でよい。"
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

[GnuPG] 2.2.3 がリリースされた。

- [[Announce] GnuPG 2.2.3 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q4/000417.html)

今回もセキュリティ・アップデートはなし。
平和でよい。
主な修正点は以下の通り。

* gpgsm: Fix initial keybox creation on Windows. [#3507]
* dirmngr: Fix crash in case of a CRL loading error. [#3510]
* Fix the name of the Windows registry key. [Git#4f5afaf1fd]
* gpgtar: Fix wrong behaviour of `--set-filename`. [#3500]
* gpg: Silence AKL retrieval messages. [#3504]
* agent: Use clock or clock_gettime for calibration. [#3056]
* agent: Improve robustness of the shutdown pending state. [Git#7ffedfab89]

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.3
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

## 【追記】 Gpg4win 3.0.1 もリリース

Gpg4win 3.0.1 もリリースされた。

- [[Gpg4win-announce] Gpg4win 3.0.1 released](http://lists.wald.intevation.org/pipermail/gpg4win-announce/2017-November/000074.html)
    - [English README file for Gpg4win](https://files.gpg4win.org/README-3.0.1.en.txt)

主な修正点は以下の通り。

* GnuPG: Has been updated to version 2.2.3.
* The mkportable process can be used again to create a portable Gpg4win variant.
* GpgOL: A user interface error for Outlook 2010 has been fixed.

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
