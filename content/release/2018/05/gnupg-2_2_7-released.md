+++
title = "GnuPG 2.2.7 がリリース"
date =  "2018-05-07T09:27:47+09:00"
description = "今回もセキュリティ・アップデートはなし。"
image = "/images/attention/tools.png"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]
draft = true

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.7 がリリースされた。

- [[Announce] GnuPG 2.2.7 released](https://lists.gnupg.org/pipermail/gnupg-announce/2018q2/000424.html)

今回もセキュリティ・アップデートはなし。
主な修正点は以下の通り。

{{% fig-quote title="GnuPG 2.2.6 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2018q2/000421.html" lang="en" %}}
* gpg: New option `--no-symkey-cache` to disable the passphrase cache for symmetrical en- and decryption.
* gpg: The ERRSIG status now prints the fingerprint if that is part of the signature.
* gpg: Relax emitting of FAILURE status lines
* gpg: Add a status flag to "sig" lines printed with `--list-sigs`.
* gpg: Fix "Too many open files" when using `--multifile`.  [#3951]
* ssh: Return an error for unknown `ssh-agent` flags.  [#3880]
* dirmngr: Fix a regression since 2.1.16 which caused corrupted CRL caches under Windows.  [#2448,#3923]
* dirmngr: Fix a CNAME problem with pools and TLS.  Also use a fixed mapping of keys.gnupg.net to sks-keyservers.net.  [#3755]
* dirmngr: Try resurrecting dead hosts earlier (from 3 to 1.5 hours).
* dirmngr: Fallback to CRL if no default OCSP responder is configured.
* dirmngr: Implement CRL fetching via https.  Here a redirection to http is explictly allowed.
* dirmngr: Make LDAP searching and CRL fetching work under Windows.  This stopped working with 2.1.  [#3937]
* agent,dirmngr: New sub-command "getenv" for "getinfo" to ease debugging.
{{% /fig-quote %}}

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.7
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

なお [Gpg4win] についてもアップデートが行われている。

- [[Gpg4win-users-en] [Gpg4win-announce] Gpg4win 3.1.1 released](http://lists.wald.intevation.org/pipermail/gpg4win-users-en/2018-May/001493.html)
- [English README file for Gpg4win](https://files.gpg4win.org/README-3.1.1.en.txt)

[Gpg4win] 3.1.1 の構成は以下の通り。

- GnuPG:          2.2.7
- Kleopatra:      3.1.1
- GPA:            0.9.10
- GpgOL:          2.1.1
- GpgEX:          1.0.6

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装](/openpgp/)

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
