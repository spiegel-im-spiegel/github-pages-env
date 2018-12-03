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
  flattr    = "spiegel"

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

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
