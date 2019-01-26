+++
description = "今回はセキュリティ・アップデートはなし。"
tags = ["security", "cryptography", "tools", "gnupg", "openpgp"]
draft = false
date = "2016-11-20T18:02:54+09:00"
title = "GnuPG 2.1.16 released"

[author]
  name = "Spiegel"
  avatar = "/images/avatar.jpg"
  flattr = ""
  url = "https://baldanders.info/spiegel/profile/"
  instagram = "spiegel_2007"
  twitter = "spiegel_2007"
  flickr = "spiegel"
  facebook = "spiegel.im.spiegel"
  license = "by-sa"
  tumblr = ""
  linkedin = "spiegelimspiegel"
  github = "spiegel-im-spiegel"
+++

- [[Announce] GnuPG 2.1.16 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q4/000398.html)

今回はセキュリティ・アップデートはなし。

* gpg: New algorithm for selecting the best ranked public key when using a mail address with `-r`, `-R`, or `--locate-key`.
* gpg: New option `--with-tofu-info` to print a new "`tfs`" record in colon formatted key listings.
* gpg: New option `--compliance` as an alternative way to specify options like `--rfc2440`, `--rfc4880`, et al.
* gpg: Many changes to the TOFU implementation.
* gpg: Improve usability of `--quick-gen-key`.
* gpg: In `--verbose` mode print a diagnostic when a pinentry is launched.
* gpg: Remove code which warns for old versions of gnome-keyring.
* gpg: New option `--override-session-key-fd`.
* gpg: Option `--output` does now work with `--verify`.
* gpgv: New option `--output` to allow saving the verified data.
* gpgv: New option `--enable-special-filenames`.
* agent, dirmngr: New `--supervised` mode for use by systemd and alike.
* agent: By default listen on all available sockets using standard names.
* agent: Invoke scdaemon with `--homedir`.
* dirmngr: On Linux now detects the removal of its own socket and terminates.
* scd: Support ECC key generation.
* scd: Support more card readers.
* dirmngr: New option `--allow-version-check` to download a software version database in the background.
* dirmngr: Use system provided CAs if no `--hkp-cacert` is given.
* dirmngr: Use a default keyserver if none is explicitly set
* gpgconf: New command `--query-swdb` to check software versions against an copy of an online database.
* gpgconf: Print the socket directory with `--list-dirs`.
* tools: The WKS tools now support draft version -02.
* tools: Always build gpg-wks-client and install under libexec.
* tools: New option` --supported` for gpg-wks-client.
* The log-file option now accepts a value "`socket://`" to log to the socket named "S.log" in the standard socket directory.
* Provide fake pinentries for use by tests cases of downstream developers.
* Fixed many bugs and regressions.
* Many changes and improvements for the test suite.

パッと見， [TOFU (Trust On First Use)](https://en.wikipedia.org/wiki/Trust_on_first_use "Trust on first use - Wikipedia") 周りをかなりいじってる感じかねぇ。

Windows 版を起動するとこんな感じになる。

```text
$ gpg --version
gpg (GnuPG) 2.1.16
libgcrypt 1.7.3
Copyright (C) 2016 Free Software Foundation, Inc.
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

いつもの感じ。

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
