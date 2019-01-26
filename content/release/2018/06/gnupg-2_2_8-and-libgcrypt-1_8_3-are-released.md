+++
title = "GnuPG 2.2.8 および Libgcrypt 1.8.3 がリリース（セキュリティ・アップデート）"
date = "2018-06-24T18:16:57+09:00"
description = "今回はキュリティ・アップデートを含む。深刻度が高いので早めのアップデートを推奨する。"
image = "/images/attention/tools.png"
tags = [
  "security",
  "vulnerability",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
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

[GnuPG] 2.2.8 および Libgcrypt 1.8.3 がリリースされた。

- [[Announce] [security fix] GnuPG 2.2.8 released (CVE-2018-12020)](https://lists.gnupg.org/pipermail/gnupg-announce/2018q2/000425.html)
- [[Announce] Libgcrypt 1.8.3 and 1.7.10 to fix CVE-2018-0495](https://lists.gnupg.org/pipermail/gnupg-announce/2018q2/000426.html)

今回はキュリティ・アップデートを含む。

## [CVE-2018-12020]

### 脆弱性の概要

{{< fig-quote title="NVD - CVE-2018-12020" link="https://nvd.nist.gov/vuln/detail/CVE-2018-12020" lang="en" >}}
<q>mainproc.c in GnuPG before 2.2.8 mishandles the original filename during decryption and verification actions, which allows remote attackers to spoof the output that GnuPG sends on file descriptor 2 to other programs that use the "--status-fd 2" option. For example, the OpenPGP data might represent an original filename that contains line feed characters in conjunction with GOODSIG or VALIDSIG status codes.</q>
{{< /fig-quote >}}

{{< fig-quote title="GnuPG 2.2.8 released (CVE-2018-12020)" link="https://lists.gnupg.org/pipermail/gnupg-announce/2018q2/000425.html" lang="en" >}}
<q>The OpenPGP protocol allows to include the file name of the original
input file into a signed or encrypted message.  During decryption and
verification the GPG tool can display a notice with that file name.  The
displayed file name is not sanitized and as such may include line feeds
or other control characters.  This can be used inject terminal control
sequences into the out and, worse, to fake the so-called status
messages.  These status messages are parsed by programs to get
information from gpg about the validity of a signature and an other
parameters.  Status messages are created with the option "--status-fd N"
where N is a file descriptor.  Now if N is 2 the status messages and the
regular diagnostic messages share the stderr output channel.  By using a
made up file name in the message it is possible to fake status messages.
Using this technique it is for example possible to fake the verification
status of a signed mail.</q>
{{< /fig-quote >}}

### 想定される影響

GPGME ライブラリを使用しない全てのメールクライアント等のアプリケーションに影響を及ぼす。
[Red Hat のレポート](https://access.redhat.com/security/cve/cve-2018-12020)ではこの件（[CVE-2018-12020]）の深刻度を「重要（7.5）」レベルと見積っている。

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更なし     |
| 機密性への影響   | なし         |
| 完全性への影響   | 高           |
| 可用性への影響   | なし         |

### 対策

[GnuPG] 2.2.8 にアップデートする。
深刻度が高いので早めのアップデートを推奨する。

アップデートが間に合わない場合は，アプリケーションに応じて，以下の対策を行うこと。

{{< fig-quote title="GnuPG 2.2.8 released (CVE-2018-12020)" link="https://lists.gnupg.org/pipermail/gnupg-announce/2018q2/000425.html" lang="en" >}}
<q>If your application uses GPGME your application is safe.  Fortunately most modern mail readers use GPGME, including GpgOL and KMail.  Mutt users should make sure to use "set crypt_use_gpgme".<br>
If you are parsing GnuPG status output and you use a dedicated file descriptor with --status-fd you are safe.  A dedicated file descriptor is one that is not shared with the log output.  The log output defaults to stderr (2) but may be a different if the option --logger-fd is used.<br>
If you are not using --verbose you are safe.  But take care: --verbose might be specified in the config file.  As a short term mitigation or if you can't immediately upgrade to the latest versions, you can add --no-verbose to the invocation of gpg.<br>
Another short term mitigation is to redirect the log output to a different file: For example "--log-file /dev/null".</q>
{{< /fig-quote >}}

`verbose` オプションを外すのが手っ取り早いかな。

## [CVE-2018-0495]

### 脆弱性の概要

{{< fig-quote title="NVD - CVE-2018-0495" link="https://nvd.nist.gov/vuln/detail/CVE-2018-0495" lang="en" >}}
<q>Libgcrypt before 1.7.10 and 1.8.x before 1.8.3 allows a memory-cache side-channel attack on ECDSA signatures that can be mitigated through the use of blinding during the signing process in the `_gcry_ecc_ecdsa_sign` function in cipher/ecc-ecdsa.c, aka the Return Of the Hidden Number Problem or ROHNP. To discover an ECDSA key, the attacker needs access to either the local machine or a different virtual machine on the same physical host.</q>
{{< /fig-quote >}}

### 想定される影響

GPGME ライブラリを使用しない全てのメールクライアント等のアプリケーションに影響を及ぼす。
[Red Hat のレポート](https://access.redhat.com/security/cve/cve-2018-0495)ではこの件（[CVE-2018-0495]）の深刻度を「警告（5.1）」レベルと見積っている。

| 基本評価基準     | 評価値   |
| ---------------- | -------- |
| 攻撃元区分       | ローカル |
| 攻撃条件の複雑さ | 高       |
| 必要な特権レベル | 不要     |
| ユーザ関与レベル | 不要     |
| スコープ         | 変更なし |
| 機密性への影響   | 高       |
| 完全性への影響   | なし     |
| 可用性への影響   | なし     |


### 対策

Libgcrypt 1.8.3 または 1.7.10 にアップデートする。

一般的に side channel 攻撃はそれほど高い深刻度にはならない。
計画的なアップデートを。

## [GnuPG] アップデート後の状態

```text
$ gpg --version
gpg (GnuPG) 2.2.8
libgcrypt 1.8.3
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

## [Gpg4win] も 3.1.2 にアップデート

- [[Gpg4win-announce] Gpg4win 3.1.2 released](http://lists.wald.intevation.org/pipermail/gpg4win-announce/2018-June/000079.html)
- [English README file for Gpg4win](https://files.gpg4win.org/README-3.1.2.en.txt)

[Gpg4win] 3.1.2 の構成は以下の通り。

- GnuPG:          2.2.8
- Kleopatra:      3.1.2
- GPA:            0.9.10
- GpgOL:          2.2.0
- GpgEX:          1.0.6

アップデートは計画的に。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[CVE-2018-12020]: https://nvd.nist.gov/vuln/detail/CVE-2018-12020 "NVD - CVE-2018-12020"
[CVE-2018-0495]: https://nvd.nist.gov/vuln/detail/CVE-2018-0495 "NVD - CVE-2018-0495"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"

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
