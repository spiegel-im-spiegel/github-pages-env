+++
title = "20周年記念！ GnuPG 2.2.4 がリリース"
date =  "2017-12-21T15:02:20+09:00"
description = "今回もセキュリティ・アップデートはない。平和なのはよいことである。"
image = "/images/attention/remark.jpg"
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

[GnuPG] 2.2.4 がリリースされた。

- [[Announce] GnuPG 2.2.4 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q4/000419.html)

なんと！

{{< fig-quote title="GnuPG 2.2.4 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2017q4/000419.html" lang="en" >}}
<q>20 years after the first version we are pleased to announce the availability of a new GnuPG release: version 2.2.4.</q>
{{< /fig-quote >}}

なんだってさ。
同梱されている `README.txt` を見ると

{{< highlight text "hl_lines=3" >}}
GNUPG is

  Copyright (C) 1997-2017 Werner Koch
  Copyright (C) 1994-2017 Free Software Foundation, Inc.
  Copyright (C) 2003-2017 g10 Code GmbH
  Copyright (C) 2002 Klarälvdalens Datakonsult AB
  Copyright (C) 1995-1997, 2000-2007 Ulrich Drepper <drepper@gnu.ai.mit.edu>
  Copyright (C) 1994 X Consortium
  Copyright (C) 1998 by The Internet Society.
  Copyright (C) 1998-2004 The OpenLDAP Foundation
  Copyright (C) 1998-2004 Kurt D. Zeilenga.
  Copyright (C) 1998-2004 Net Boolean Incorporated.
  Copyright (C) 2001-2004 IBM Corporation.
  Copyright (C) 1999-2003 Howard Y.H. Chu.
  Copyright (C) 1999-2003 Symas Corporation.
  Copyright (C) 1998-2003 Hallvard B. Furuseth.
  Copyright (C) 1992-1996 Regents of the University of Michigan.
  Copyright (C) 2000 Dimitrios Souflis
  Copyright (C) 2008,2009,2010,2012-2016 William Ahern

  GnuPG is free software; you can redistribute it and/or modify it
  under the terms of the GNU General Public License as published by
  the Free Software Foundation; either version 3 of the License, or
  (at your option) any later version.

  GnuPG is distributed in the hope that it will be useful, but WITHOUT
  ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
  or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public
  License for more details.

  You should have received a copy of the GNU General Public License
  along with this program; if not, see <https://www.gnu.org/licenses/>.
{{< /highlight >}}

となっているので，そういうことなんだろう。

ちなみに，私が最初の [BkGnuPG] を公開[^bgp1] したのが2000年9月12日，らしい（よく憶えてない`w`）。
色んな思惑や成り行きがあり，当時は仕事で使うのは憚れた PGP から [GnuPG] へ切り替えるにあたって当時使ってた Becky! がネックになっていて，誰も [GnuPG] 用のプラグインを作ってくれないからしょうがなしに自分で作ったんだよな。
したら日本語圏より英語圏のほうで需要が高くて，英語不得手なのに大変だったよ。
最後はロシア語にも対応してくれと言われる始末（ロシア語のリソースファイルもらって対応したけどね）。

[^bgp1]: [BkGnuPG] はもうメンテナンスしてないので間違っても使おうとか考えないように。危ないからね。

もし将来また GUI なアプリケーションをフリーで公開することがあるなら今度は最初から英語版を作ろうと心に誓ったのだが，今は「アプリ」の時代で国際化は当たり前だしねぇ。
もう私の出番はないだろう。

与太話はこれくらいにして， [GnuPG] 2.2.4 では今回もセキュリティ・アップデートはない。
平和なのはよいことである。
主な修正点は以下の通り。

* gpg: Change default preferences to prefer SHA512.
* gpg: Print a warning when more than 150 MiB are encrypted using a cipher with 64 bit block size.
* gpg: Print a warning if the MDC feature has not been used for a message.
* gpg: Fix regular expression of domain addresses in trust signatures. [#2923]
* agent: New option `--auto-expand-secmem` to help with high numbers of concurrent connections.  Requires libgcrypt 1.8.2 for having an effect.  [#3530]
* dirmngr: Cache responses of WKD queries.
* gpgconf: Add option `--status-fd`.
* wks: Add commands --check and `--remove-key` to gpg-wks-server.
* Increase the backlog parameter of the daemons to 64 and add option `--listen-backlog`.
* New configure option `--enable-run-gnupg-user-socket` to first try a socket directory which is not removed by systemd at session end.

`--enable-run-gnupg-user-socket` オプションはそのうち試してみたい。
そのうちね。

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.4
libgcrypt 1.8.2
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

## ブックマーク

- [[Gpg4win-announce] Gpg4win 3.0.2 released](http://lists.wald.intevation.org/pipermail/gpg4win-announce/2017-December/000075.html) : こちらはまだ [GnuPG] 2.2.3 ベース
- [[Announce] GPGME 1.10.0 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q4/000418.html)

- [OpenPGP の実装](/openpgp/)

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[BkGnuPG]: https://github.com/spiegel-im-spiegel/BkGnuPG "spiegel-im-spiegel/BkGnuPG: GNU Privacy Guard Plug-in for Becky! 2"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
