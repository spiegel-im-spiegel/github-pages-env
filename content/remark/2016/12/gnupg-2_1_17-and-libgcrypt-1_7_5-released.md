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
  url = "https://baldanders.info/profile/"
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

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
