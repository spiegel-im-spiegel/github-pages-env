+++
title = "GnuPG 2.2.5 がリリース"
date = "2018-02-24T09:52:21+09:00"
description = "今回もセキュリティ・アップデートはない。"
image = "/images/attention/kitten.jpg"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
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

[GnuPG] 2.2.5 がリリースされた。

- [[Announce] GnuPG 2.2.5 released](https://lists.gnupg.org/pipermail/gnupg-announce/2018q1/000420.html)

今回もセキュリティ・アップデートはなし。
主な修正点は以下の通り。

* gpg: Allow the use of the "cv25519" and "ed25519" short names in addition to the canonical curve names in `--batch --gen-key`.
* gpg: Make sure to print all secret keys with option --list-only and `--decrypt`.  [#3718]
* gpg: Fix the use of future-default with `--quick-add-key` for signing keys.  [#3747]
* gpg: Select a secret key by checking availability under gpg-agent. [#1967]
* gpg: Fix reversed prompt texts for `--only-sign-text-ids`.  [#3787]
* gpg,gpgsm: Fix detection of bogus keybox blobs on 32 bit systems. [#3770]
* gpgsm: Fix regression since 2.1 in `--export-secret-key-raw` which got $d mod (q-1)$ wrong.  Note that most tools automatically fixup that parameter anyway.
* ssh: Fix a regression in getting the client'd PID on `*`BSD and macOS.
* scd: Support the KDF Data Object of the OpenPGP card 3.3.  [#3152]
* scd: Fix a regression in the internal CCID driver for certain card readers.  [#3508]
* scd: Fix a problem on NetBSD killing scdaemon on gpg-agent shutdown.  [#3778]
* dirmngr: Improve returned error description on failure of DNS resolving.  [#3756]
* wks: Implement command `--install-key` for gpg-wks-server.
* Add option STATIC=1 to the Speedo build system to allow a build with statically linked versions of the core GnuPG libraries.  Also use `--enable-wks-tools` by default by Speedo builds for Unix.

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.5
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

## ブックマーク

- [OpenPGP の実装](/openpgp/)

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
