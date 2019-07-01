+++
title = "GnuPG 2.2.9 がリリース"
date = "2018-08-18T15:43:11+09:00"
description = "今回はセキュリティ・アップデートはなし。"
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

1ヶ月以上前の話で恐縮だが [GnuPG] 2.2.9 がリリースされた。

- [[Announce] GnuPG 2.2.9 released](https://lists.gnupg.org/pipermail/gnupg-announce/2018q3/000427.html)

今回はセキュリティ・アップデートはなし。
主な修正点は以下の通り。

{{% fig-quote type="md" title="GnuPG 2.2.9 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2018q3/000427.html" lang="en" %}}
* dirmngr: Fix recursive resolver mode and other bugs in the libdns code.  [#3374,#3803,#3610] 
* dirmngr: When using libgpg-error 1.32 or later a GnuPG build with NTBTLS support (e.g. the standard Windows installer) does not anymore block for dozens of seconds before returning data.  If you still have problems on Windows, please consider to use one of the options `disable-ipv4` or `disable-ipv6`.
* gpg: Fix bug in `--show-keys` which actually imported revocation certificates.  [#4017]
* gpg: Ignore too long user-ID and comment packets.  [#4022]
* gpg: Fix crash due to bad German translation.  Improved printf format compile time check.
* gpg: Handle missing `ISSUER` sub packet gracefully in the presence of the new `ISSUER_FPR`.  [#4046]
* gpg: Allow decryption using several passphrases in most cases. [#3795,#4050]
* gpg: Command `--show-keys` now enables the list options `show-unusable-uids`, `show-unusable-subkeys`, `show-notations` and `show-policy-urls` by default.
* gpg: Command `--show-keys` now prints revocation certificates. [#4018]
* gpg: Add revocation reason to the "rev" and "rvs" records of the option `--with-colons`.  [#1173]
* gpg: Export option `export-clean` does now remove certain expired subkeys; `export-minimal` removes all expired subkeys.  [#3622]
* gpg: New "usage" property for the `drop-subkey` filters.  [#4019]
{{% /fig-quote %}}

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.9
libgcrypt 1.8.3
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

## ブックマーク

- [OpenPGP の実装]({{< ref "/openpgp" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B015643CPE</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
