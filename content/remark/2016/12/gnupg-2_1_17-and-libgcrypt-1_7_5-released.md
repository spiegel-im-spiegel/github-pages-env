+++
date = "2016-12-26T20:45:46+09:00"
update = "2016-12-27T06:46:48+09:00"
title = "GnuPG 2.1.17 and Libgcrypt 1.7.5 released"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]
draft = false
description = "Libgcrypt はバグフィックスがメイン， GnuPG 2.1 系 はちょこちょこと機能変更が行われている。"

[author]
  flickr = "spiegel"
  twitter = "spiegel_2007"
  tumblr = ""
  avatar = "/images/avatar.jpg"
  url = "https://baldanders.info/spiegel/profile/"
  license = "by-sa"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  facebook = "spiegel.im.spiegel"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  flattr = ""
+++

Libgcrypt 1.7.5 および GnuPG 2.1.17 がリリースされている。
Libgcrypt はバグフィックスがメイン， GnuPG 2.1 系 はちょこちょこと機能変更が行われている。

- [[Announce] Libgcrypt 1.7.5 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q4/000399.html)
- [[Announce] GnuPG 2.1.17 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q4/000400.html)

以下は GnuPG 2.1.17 の変更履歴。

* gpg: By default new keys expire after 2 years.
* gpg: New command `--quick-set-expire` to conveniently change the expiration date of keys.
* gpg: Option and command names have been changed for easier comprehension.  The old names are still available as aliases.
* gpg: Improved the TOFU trust model.
* gpg: New option `--default-new-key-algo`.
* scd: Support OpenPGP card V3 for RSA.
* dirmngr: Support for the ADNS library has been removed.  Instead William Ahern's Libdns is now source included and used on all platforms.  This enables Tor support on all platforms.  The new option `--standard-resolver` can be used to disable this code at runtime.  In case of build problems the new configure option `--disable-libdns` can be used to build without Libdns.
* dirmngr: Lazily launch ldap reaper thread.
* tools: New options `--check` and `--status-fd` for gpg-wks-client.
* The UTF-8 byte order mark is now skipped when reading conf files.
* Fixed many bugs and regressions.
* Major improvements to the test suite.  For example it is possible to run the external test suite of GPGME.

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
