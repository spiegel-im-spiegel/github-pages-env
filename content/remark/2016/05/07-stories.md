+++
date = "2016-05-07T20:36:22+09:00"
description = "自前でブックマークを運用することにしました / GnuPG 2.1.12 and Libgcrypt 1.7.0 released / その他の気になる記事"
draft = false
tags = ["bookmark", "security", "cryptography", "tools", "gnupg"]
title = "週末スペシャル： 自前でブックマークを運用することにしました"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1. [自前でブックマークを運用することにしました]({{< relref "#bkmk" >}})
1. [GnuPG 2.1.12 and Libgcrypt 1.7.0 released]({{< relref "#gpg" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## 自前でブックマークを運用することにしました{#bkmk}

先月

- [delicious.com が del.icio.us に移行するらしい（出戻り？）]({{< relref "remark/2016/04/delicious-to-del_icio_us.md" >}})

と書いたのだが，その後無事に移行が完了した。

ただねぇ。
Markdown でリンクを列挙するのがメチャメチャ楽ちんなことに気づいてしまった。
もう folksonomy な時代じゃないだろうし，今年に入ってから [del.icio.us](https://del.icio.us/) のデザインが一新したけど，イマイチ使い勝手がよくないんだよな。

というわけで [Bookmarks](/bookmarks/) セクションを追加して，そこにブックマークを蓄積していくことにした。
タグが使えなくなるけどソースコードの markdown ファイルを全文検索できるので，まぁ問題ないかと。

[del.icio.us/spiegel](https://del.icio.us/spiegel) にあるブックマークは「[Export from Delicious.com]({{< relref "bookmarks/2016/04/export-from-delicious.html" >}})」に移した。
ただ [del.icio.us](https://del.icio.us/) の Export 機能がうまく動いてなくて古いブックマークが取れない。
あと「[Export from Delicious.com]({{< relref "bookmarks/2016/04/export-from-delicious.html" >}})」にあるリンクが膨大で表示に結構時間がかかるので迂闊に開かないほうがいいようだ（笑）

[Bookmarks](/bookmarks/) セクションの記事は[メインの RSS](/index.xml) に含めないことにした。
RSS を撮りたいという奇特な方は[こちら](/bookmarks/index.xml)からどうぞ。

## GnuPG 2.1.12 and Libgcrypt 1.7.0 released{#gpg}

- [[Announce] GnuPG 2.1.12 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q2/000387.html)
- [[Announce] Libgcrypt 1.7.0 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q2/000386.html)

セキュリティ・アップデートはなし。

GnuPG 2.1.12 の変更点は以下のとおり。

* gpg: New `--edit-key` sub-command "`change-usage`" for testing purposes.
* gpg: Out of order key-signatures are now systematically detected and fixed by `--edit-key`.
* gpg: Improved detection of non-armored messages.
* gpg: Removed the extra prompt needed to create Curve25519 keys.
* gpg: Improved user ID selection for `--quick-sign-key`.
* gpg: Use the root CAs provided by the system with --fetch-key.
* gpg: Add support for the experimental Web Key Directory key location service.
* gpg: Improve formatting of Tofu messages and emit new Tofu specific status lines.
* gpgsm: Add option `--pinentry-mode` to support a loopback pinentry.
* gpgsm: A new pubring.kbx is now created with the header blob so that gpg can detect that the keybox format needs to be used.
* agent: Add read support for the new private key protection format openpgp-s2k-ocb-aes.
* agent: Add read support for the new extended private key format.
* agent: Default to `--allow-loopback-pinentry` and add option `--no-allow-loopback-pinentry`.
* scd: Changed to use the new libusb 1.0 API for the internal CCID driver.
* dirmngr: The dirmngr-client does now auto-detect the PEM format.
* g13: Add experimental support for dm-crypt.
* w32: Tofu support is now available with the Speedo build method.
* w32: Removed the need for libiconv.dll.
* The man pages for gpg and gpgv are now installed under the correct name (gpg2 or gpg - depending on a configure option).
* Lots of internal cleanups and bug fixes.

ふむ。
Windows 版では libiconv.dll は不要になったらしい。

Libgcrypt 1.7.0 について 1.6 系からの変更点は以下のとおり。

* New algorithms and modes:
    - SHA3-224, SHA3-256, SHA3-384, SHA3-512, and MD2 hash algorithms.
    - SHAKE128 and SHAKE256 extendable-output hash algorithms.
    - ChaCha20 stream cipher.
    - Poly1305 message authentication algorithm
    - ChaCha20-Poly1305 Authenticated Encryption with Associated Data mode.
    - OCB mode.
    - HMAC-MD2 for use by legacy applications.
* New curves for ECC:
    - Curve25519.
    - sec256k1.
    - GOST R 34.10-2001 and GOST R 34.10-2012.
 * Performance:
    - Improved performance of KDF functions.
    - Assembler optimized implementations of Blowfish and Serpent on ARM.
    - Assembler optimized implementation of 3DES on x86.
    - Improved AES using the SSSE3 based vector permutation method by Mike Hamburg.
    - AVX/BMI is used for SHA-1 and SHA-256 on x86.  This is for SHA-1 about 20% faster than SSSE3 and more than 100% faster than the generic C implementation.
    - 40% speedup for SHA-512 and 72% for SHA-1 on ARM Cortex-A8.
    - 60-90% speedup for Whirlpool on x86.
    - 300% speedup for RIPE MD-160.
    - Up to 11 times speedup for CRC functions on x86.
* Other features:
    - Improved ECDSA and FIPS 186-4 compliance.
    - Support for Montgomery curves.
    - gcry_cipher_set_sbox to tweak S-boxes of the gost28147 cipher algorithm.
    - gcry_mpi_ec_sub to subtract two points on a curve.
    - gcry_mpi_ec_decode_point to decode an MPI into a point object.
    - Emulation for broken Whirlpool code prior to 1.6.0.  [from 1.6.1]
    - Flag "pkcs1-raw" to enable PCKS#1 padding with a user supplied hash part.
    - Parameter "saltlen" to set a non-default salt length for RSA PSS.
    - A SP800-90A conforming DRNG replaces the former X9.31 alternative random number generator.
    - Map deprecated RSA algo number to the RSA algo number for better backward compatibility. [from 1.6.2]
    - Use ciphertext blinding for Elgamal decryption [CVE-2014-3591]. See http://www.cs.tau.ac.il/~tromer/radioexp/ for details. [from 1.6.3]
    - Fixed data-dependent timing variations in modular exponentiation [related to CVE-2015-0837, Last-Level Cache Side-Channel Attacks are Practical]. [from 1.6.3]
    - Flag "no-keytest" for ECC key generation.  Due to a bug in the parser that flag will also be accepted but ignored by older version of Libgcrypt. [from 1.6.4]
    - Speed up the random number generator by requiring less extra seeding. [from 1.6.4]
    - Always verify a created RSA signature to avoid private key leaks due to hardware failures. [from 1.6.4]
    - Mitigate side-channel attack on ECDH with Weierstrass curves [CVE-2015-7511].  See http://www.cs.tau.ac.IL/~tromer/ecdh/ for details. [from 1.6.5]
* Internal changes:
    - Moved locking out to libgpg-error.
    - Support of the SYSROOT envvar in the build system.
    - Refactor some code.
    - The availability of a 64 bit integer type is now mandatory.
* Bug fixes:
    - Fixed message digest lookup by OID (regression in 1.6.0).
    - Fixed a build problem on NetBSD
    - Fixed memory leaks in ECC code.
    - Fixed some asm build problems and feature detection bugs.

インストール後はこんな感じ。

```text
$ gpg --version
gpg (GnuPG) 2.1.12
libgcrypt 1.7.0
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

## その他の気になる記事{#other}

- [「出版不況論」をめぐる議論の混乱について « マガジン航[kɔː]](http://magazine-k.jp/2016/05/01/editors-note-12/)
    - [対談「50年後の文芸はどうなっているのか？」藤谷治✕藤井大洋 « マガジン航[kɔː]](http://magazine-k.jp/2016/05/02/literature-in-2066/)
    - [ライトノベル市場は衰退どころか拡大傾向にある ── 『ORICONエンタメ・マーケット白書2015』より:見て歩く者 by 鷹野凌](http://www.wildhawkfield.com/2016/05/real-light-novel-market.html)
- [やましくなければプライバシーは要らない？ nothing to hideを巡って(八田真行) - 個人 - Yahoo!ニュース](http://bylines.news.yahoo.co.jp/hattamasayuki/20160430-00057230/)
- [About the security content of Xcode 7.3.1 - Apple サポート](https://support.apple.com/ja-jp/HT206338)
    - [Vulnerability Summary for CVE-2016-2315](https://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2016-2315)
    - [Vulnerability Summary for CVE-2016-2324](https://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2016-2324)
    - [AppleがXcodeの更新版を公開、Gitの脆弱性を修正 - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1605/06/news042.html)
- [ImageMagick の脆弱性 (CVE-2016-3714) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160021.html)
    - [ImageMagickの脆弱性(CVE-2016-3714他)についてまとめてみた - piyolog](http://d.hatena.ne.jp/Kango/20160504/1462352882)
    - [[重要] ImageMagick の脆弱性 ImageTragick (CVE-2016-3714) への対応について | Movable Type ニュース](http://www.sixapart.jp/movabletype/news/2016/05/06-1452.html)
- [「OpenSSL」に6件の脆弱性、修正を施した最新版v1.0.2h/v1.0.1tが公開 - 窓の杜](http://www.forest.impress.co.jp/docs/news/20160506_756180.html)
    - [OpenSSLの更新版公開、複数の脆弱性を修正 - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1605/06/news046.html)
    - [OpenSSLに2件の深刻な脆弱性--早急にパッチの適用を - ZDNet Japan](http://japan.zdnet.com/article/35082166/)
    - [Node v0.10.45 (Maintenance) | Node.js](https://nodejs.org/en/blog/release/v0.10.45/)
    - [Node v0.12.14 (Maintenance) | Node.js](https://nodejs.org/en/blog/release/v0.12.14/)
    - [Node v4.4.4 (LTS) | Node.js](https://nodejs.org/en/blog/release/v4.4.4/)
    - [Node v5.11.1 (Stable) | Node.js](https://nodejs.org/en/blog/release/v5.11.1/)
    - [Node v6.1.0 (Current) | Node.js](https://nodejs.org/en/blog/release/v6.1.0/)
- [Chrome, Firefox and Safari Block Pirate Bay as "Phishing" Site - TorrentFreak](https://torrentfreak.com/chrome-and-firefox-block-tpb-as-phishing-site-160507/)
- [ローリングストーンズにはドナルドトランプ氏に楽曲の使用中止を命じる権利があるのか(栗原潔) - 個人 - Yahoo!ニュース](http://bylines.news.yahoo.co.jp/kuriharakiyoshi/20160506-00057409/)
