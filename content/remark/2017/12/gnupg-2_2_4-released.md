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

Home: C:/home/spiegel/keyring
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

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/">暗号化 プライバシーを救った反乱者たち</a></dt><dd>スティーブン・レビー 斉藤 隆央 </dd><dd>紀伊國屋書店 2002-02-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487593100X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/487593100X.09._SCTHUMBZZZ_.jpg"  alt="ハッカーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4105393022/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4105393022.09._SCTHUMBZZZ_.jpg"  alt="暗号解読―ロゼッタストーンから量子暗号まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4484111160/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4484111160.09._SCTHUMBZZZ_.jpg"  alt="グーグル ネット覇者の真実 追われる立場から追う立場へ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/410215972X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/410215972X.09._SCTHUMBZZZ_.jpg"  alt="暗号解読〈上〉 (新潮文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4102159738/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4102159738.09._SCTHUMBZZZ_.jpg"  alt="暗号解読 下巻 (新潮文庫 シ 37-3)"  /></a> </p>
<p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-03-09">2015/03/09</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
