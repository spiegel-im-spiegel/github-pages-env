+++
title = "GnuPG 2.1.23 がリリース"
date =  "2017-08-10T10:43:43+09:00"
description = "今回の目玉はやっぱり --enable-gpg-is-gpg2 オプションかな。"
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
+++

GnuPG 2.1.23 がリリースされた。
セキュリティ・アップデートはなし。

- [[Announce] GnuPG 2.1.23 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q3/000412.html)

修正および変更点は以下の通り。

* gpg: "gpg" is now installed as "gpg" and not anymore as "gpg2".  If needed, the new configure option `--enable-gpg-is-gpg2` can be used to revert this.
* gpg: Options `--auto-key-retrieve` and `--auto-key-locate` "local,wkd" are now used by default.  Note: this enables keyserver and Web Key Directory operators to notice when a signature from a locally non-available key is being verified for the first time or when you intend to encrypt to a mail address without having the key locally.  This new behaviour will eventually make key discovery much easier and mostly automatic.  Disable this by adding `no-auto-key-retrieve` `auto-key-locate` local to your gpg.conf.
* agent: Option `--no-grab` is now the default.  The new option `--grab` allows to revert this.
* gpg: New import option "show-only".
* gpg: New option `--disable-dirmngr` to entirely disable network access for gpg.
* gpg,gpgsm: Tweaked DE-VS compliance behaviour.
* New configure flag `--enable-all-tests` to run more extensive tests during "make check".
* gpgsm: The keygrip is now always printed in colon mode as documented in the man page.
* Fixed connection timeout problem under Windows.

今回の目玉はやっぱり `--enable-gpg-is-gpg2` オプションかな。
でもこのオプションの有無で違いが分からないんだよなぁ。
私の環境は classic 版からの以降なので，そのせいかな？

インストールが上手くいけば以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.1.23
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
