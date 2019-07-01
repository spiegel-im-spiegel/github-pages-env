+++
title = "GnuPG 2.2.2 がリリースされた"
date =  "2017-11-07T23:39:38+09:00"
description = "今回はセキュリティ・アップデートはなし。"
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

[GnuPG] 2.2.2 がリリースされた。

- [[Announce] GnuPG 2.2.2 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q4/000416.html)

今回はセキュリティ・アップデートはなし。
主な修正点は以下の通り。

* gpg: Avoid duplicate key imports by concurrently running gpg processes. [#3446]
* gpg: Fix creating on-disk subkey with on-card primary key. [#3280]
* gpg: Fix validity retrieval for multiple keyrings. [Debian#878812]
* gpg: Fix `--dry-run` and import option show-only for secret keys.
* gpg: Print "sec" or "sbb" for secret keys with import option import-show. [#3431]
* gpg: Make import less verbose. [#3397]
* gpg: Add alias "Key-Grip" for parameter "Keygrip" and new parameter "Subkey-Grip" to unattended key generation.  [#3478]
* gpg: Improve "factory-reset" command for OpenPGP cards.  [#3286]
* gpg: Ease switching Gnuk tokens into ECC mode by using the magic keysize value 25519.
* gpgsm: Fix `--with-colon` listing in crt records for fields > 12.
* gpgsm: Do not expect X.509 keyids to be unique.  [#1644]
* agent: Fix stucked Pinentry when using --max-passphrase-days. [#3190]
* agent: New option `--s2k-count`.  [#3276 (workaround)]
* dirmngr: Do not follow https-to-http redirects. [#3436]
* dirmngr: Reduce default LDAP timeout from 100 to 15 seconds. [#3487]
* gpgconf: Ignore non-installed components for commands `--apply-profile` and `--apply-defaults`. [#3313]
* Add configure option `--enable-werror`.  [#2423]


最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.2
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

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

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
