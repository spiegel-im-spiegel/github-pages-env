+++
date = "2016-02-11T03:58:42+09:00"
update = "2016-02-15T07:48:52+09:00"
description = "セキュリティ・アップデートを含むため，関係しているアプリケーション（GnuPG を含む）を使用している場合はアップデートする必要がある。"
draft = false
tags = ["security", "vulnerability", "cryptography", "libgcrypt", "gnupg", "openpgp"]
title = "Libgcrypt 1.6.5 with security fix released"

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

[Libgcrypt] の 1.6.5 がリリースされた。
セキュリティ・アップデートを含むため，関係している暗号製品（[GnuPG] を含む）を使用している場合はアップデートする必要がある。

- [[Announce] Libgcrypt 1.6.5 with security fix released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q1/000384.html)

## 脆弱性の内容

[Libgcrypt] の 1.6.5 の変更内容は以下の通り。

* Mitigate side-channel attack on ECDH with Weierstrass curves [CVE-2015-7511].  See [http://www.cs.tau.ac.IL/~tromer/ecdh/](http://www.cs.tau.ac.IL/~tromer/ecdh/) for details.
* Fix build problem on Solaris.

ECDH (Elliptic Curve Diffie–Hellman key exchange) に対する side-channel attack に対応している。
ただし低減（mitigate）レベル。
一般的に side-channel attack は完全な対処が難しいためこのようなことになる。
詳しくは以下を参照のこと。

- [ECDH Key-Extraction via Low-Bandwidth Electromagnetic Attacks on PCs](http://www.cs.tau.ac.il/~tromer/ecdh/)
- [Cryptology ePrint Archive: Report 2016/129](http://eprint.iacr.org/2016/129)

## 影響度（CVSS）

CVE-2015-7511 （[CVE-2015-7511 - Red Hat Customer Portal](https://access.redhat.com/security/cve/CVE-2015-7511)）

CVSSv2 基本値 4.3 (`AV:A/AC:M/Au:N/C:P/I:P/A:N`)

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | 隣接ネットワーク（A） |
| 攻撃条件の複雑さ（AC）                  | 中（M）           |
| 攻撃前の認証要否（Au）                  | 不要（N）         |
| 情報漏えいの可能性（機密性への影響, C） | 部分的（P）       |
| 情報改ざんの可能性（完全性への影響, I） | 部分的（P）       |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |

CVSS については[デモページ](http://www.baldanders.info/spiegel/archive/cvss/cvss2.html)を参照のこと。

## 影響を受ける実装

- [Libgcrypt] を利用する [GnuPG] 2.x (stable および modern バージョン)

Windows 版 [GnuPG] については [2.1.11]({{< relref "remark/2016/01/30-stories.md#gpg" >}}) の 20160209 版がリリースされているので入れ替えること。

[GnuPG] や [Libgcrypt] はよく研究対象にされるので，今回のような脆弱性報告は割とあるのだが，他の暗号実装（たとえば OpenSSL とか）はどうなんだろうと心配になる。
実際どうなんだろう。

[Libgcrypt]: https://www.gnu.org/software/libgcrypt/ "Libgcrypt - GNU Project - Free Software Foundation (FSF)"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
