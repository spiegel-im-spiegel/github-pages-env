+++
title = "GnuPG 2.2.11 がリリース"
date = "2018-11-11T22:48:47+09:00"
description = "おおう。 2.2.10 リリース時のアナウンスを忘れてたよ。今回もセキュリティ・アップデートはなし。"
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

[scripts]
  mathjax = false
  mermaidjs = false
+++

おおう。
2.2.10 リリース時のアナウンスを忘れてたよ（手元のはちゃんと 2.2.10 にアップデートしてたのに）。

というわけで，2.2.10 すっ飛ばして [GnuPG] 2.2.11 リリースのお知らせです。

- [[Announce] GnuPG 2.2.11 released](https://lists.gnupg.org/pipermail/gnupg-announce/2018q4/000432.html)

今回もセキュリティ・アップデートはなし。
主な修正点は以下の通り（2.2.10 と併せて載せておきます）。

{{% fig-quote title="GnuPG 2.2.10 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2018q3/000428.html" lang="en" %}}
- gpg: Refresh expired keys originating from the WKD.  [#2917]
- gpg: Use a 256 KiB limit for a WKD imported key.
- gpg: New option `--known-notation`.  [#4060]
- scd: Add support for the Trustica Cryptoucan reader.
- agent: Speed up starting during on-demand launching.  [#3490]
- dirmngr: Validate SRV records in WKD queries.
{{% /fig-quote %}}

{{% fig-quote title="GnuPG 2.2.11 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2018q4/000432.html" lang="en" %}}
* gpgsm: Fix CRL loading when intermediate certicates are not yet trusted.
* gpgsm: Fix an error message about the digest algo.  [#4219]
* gpg: Fix a wrong warning due to new sign usage check introduced with 2.2.9.  [#4014]
* gpg: Print the "data source" even for an unsuccessful keyserver query.
* gpg: Do not store the TOFU trust model in the trustdb.  This allows to enable or disable a TOFO model without triggering a trustdb rebuild.  [#4134]
* scd: Fix cases of "Bad PIN" after using "forcesig".  [#4177]
* agent: Fix possible hang in the ssh handler.  [#4221]
* dirmngr: Tack the unmodified mail address to a WKD request.  See commit `a2bd4a64e5b057f291a60a9499f881dd47745e2f` for details.
* dirmngr: Tweak diagnostic about missing LDAP server file.
* dirmngr: In verbose mode print the OCSP responder id.
* dirmngr: Fix parsing of the LDAP port.  [#4230]
* wks: Add option `--directory`/`-C` to the server.  Always build the server on Unix systems.
* wks: Add option `--with-colons` to the client.  Support sites which use the policy file instead of the submission-address file.
* Fix EBADF when gpg et al. are called by broken CGI scripts.
* Fix some minor memory leaks and bugs.
{{% /fig-quote %}}


最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.11
libgcrypt 1.8.4
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

[Libgcrypt] のバージョンが 1.8.4 に上がっているのでご注意を。
こちらもセキュリティ・アップデートはなし。

- [[Announce] Libgcrypt 1.8.4 released](https://lists.gnupg.org/pipermail/gnupg-announce/2018q4/000431.html)

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< ref "/openpgp" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Libgcrypt]: https://gnupg.org/software/libgcrypt/
