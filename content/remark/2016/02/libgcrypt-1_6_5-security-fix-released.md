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
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
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

CVSS については[デモページ](https://baldanders.info/spiegel/archive/cvss/cvss2.html)を参照のこと。

## 影響を受ける実装

- [Libgcrypt] を利用する [GnuPG] 2.x (stable および modern バージョン)

Windows 版 [GnuPG] については [2.1.11]({{< ref "/remark/2016/01/30-stories.md#gpg" >}}) の 20160209 版がリリースされているので入れ替えること。

[GnuPG] や [Libgcrypt] はよく研究対象にされるので，今回のような脆弱性報告は割とあるのだが，他の暗号実装（たとえば OpenSSL とか）はどうなんだろうと心配になる。
実際どうなんだろう。

[Libgcrypt]: https://www.gnu.org/software/libgcrypt/ "Libgcrypt - GNU Project - Free Software Foundation (FSF)"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

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
