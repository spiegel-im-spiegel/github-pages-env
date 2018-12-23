+++
title = "GnuPG 2.2.12 のリリース"
date = "2018-12-15T18:01:01+09:00"
description = "今回もセキュリティ・アップデートはなし。機能追加がいくつか。"
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
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.12 がリリースされた。

- [[Announce] GnuPG 2.2.12 released](https://lists.gnupg.org/pipermail/gnupg-announce/2018q4/000433.html)

今回もセキュリティ・アップデートはなし。
主な機能追加・修正点は以下の通り。

{{% fig-quote title="GnuPG 2.2.12 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2018q4/000433.html" lang="en" %}}
* tools: New commands `--install-key` and `--remove-key` for gpg-wks-client.  This allows to prepare a Web Key Directory on a local file system for later upload to a web server.
* gpg: New `--list-option` "show-only-fpr-mbox".  This makes the use of the new gpg-wks-client `--install-key` command easier on Windows.
* gpg: Improve processing speed when `--skip-verify` is used.
* gpg: Fix a bug where a LF was accidentally written to the console.
* gpg: --card-status now shwos whether a card has the new KDF feature enabled.
* agent: New runtime option `--s2k-calibration=MSEC`.  New configure option `--with-agent-s2k-calibration=MSEC`. [https://dev.gnupg.org/T3399]
* dirmngr: Try another keyserver from the pool on receiving a 502, 503, or 504 error.  [https://dev.gnupg.org/T4175]
* dirmngr: Avoid possible CSRF attacks via http redirects.  A HTTP query will not anymore follow a 3xx redirect unless the Location header gives the same host.  If the host is different only the host and port is taken from the Location header and the original path and query parts are kept.
* dirmngr: New command FLUSHCRL to flush all CRLS from disk and memory.  [https://dev.gnupg.org/T3967]
* New simplified Chinese translation (zh_CN).
{{% /fig-quote %}}

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.12
libgcrypt 1.8.4
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
[Libgcrypt]: https://gnupg.org/software/libgcrypt/

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" height="160" alt="暗号技術入門 第3版　秘密の国のアリス"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22">暗号技術入門 第3版　秘密の国のアリス</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ</dd>
    <dd>評価&nbsp;<abbr class="rating fa-sm" title="5">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
    </abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.12.15</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>
