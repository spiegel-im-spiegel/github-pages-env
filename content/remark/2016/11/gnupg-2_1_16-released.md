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

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
