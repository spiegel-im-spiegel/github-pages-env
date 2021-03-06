+++
date = "2016-09-19T15:26:36+09:00"
description = "乱数生成器の脆弱性のようだ。CVSS 値は要注意レベル。計画的にアップデートすることをお薦めする。"
draft = false
tags = ["security", "cryptography", "vulnerability", "libgcrypt", "gnupg", "openpgp"]
title = "Security fixes for Libgcrypt and GnuPG 1.4"

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
  url = "https://baldanders.info/profile/"
+++

いやぁ，先月の話でゴメン。
ここ1ヶ月ほど余裕がなくて，ようやく個人メールの処理が完了したよ。

で，一応これは書いておかないとね。

- [[Announce] Security fixes for Libgcrypt and GnuPG 1.4 [CVE-2016-6316]](https://lists.gnupg.org/pipermail/gnupg-announce/2016q3/000395.html) [^cve]
- [[Announce] GnuPG 2.1.15 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q3/000396.html)

[^cve]: タイトルの CVE 項番は間違いで正しくは [CVE-2016-6313](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2016-6313)。

> Felix Dörre and Vladimir Klebanov from the Karlsruhe Institute of Technology found a bug in the mixing functions of Libgcrypt's random number generator: An attacker who obtains 4640 bits from the RNG can trivially predict the next 160 bits of output.  This bug exists since 1998 in all GnuPG and Libgcrypt versions.

- [Libgcrypt ・GnuPG 1.4 の 脆弱性 ( CVE-2016-6313 ) — | サイオスOSS | サイオステクノロジー](https://oss.sios.com/security/general-security-20160818)

ちうわけで乱数生成器の脆弱性のようだ。
この脆弱性の影響度は以下の通り。

**CVSSv2 基本評価値 4.0 (`AV:N/AC:H/Au:N/C:P/I:P/A:N`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 高（H）           |
| 攻撃前の認証要否（Au）                  | 不要（N）         |
| 情報漏えいの可能性（機密性への影響, C） | 部分的（P）       |
| 情報改ざんの可能性（完全性への影響, I） | 部分的（P）       |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |


**CVSSv3 基本評価値 4.8 (`CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:N`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 高（H）           |
| 必要な特権レベル（PR）                  | 不要（N）         |
| ユーザ関与レベル（UI）                  | 不要（N）         |
| スコープ（S）                           | 変更なし (U)      |
| 情報漏えいの可能性（機密性への影響, C） | 低（L）           |
| 情報改ざんの可能性（完全性への影響, I） | 低（L）           |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |

（[CVE-2016-6313 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2016-6313)）

ちうわけで要注意レベルなので，計画的にアップデートすることをお薦めする（急がなくてよい）。

[GnuPG] 1.4 (classic) 系については 1.4.21 へアップデートすること。
[GnuPG] 2.0 (stable) および 2.1 (modern) 系については [Libgcrypt] を 1.7.3, 1.6.6, および 1.5.6 へアップデートすること。
Windows 版については [GnuPG] 2.1.15 にアップデートすることをお薦めする。

[GnuPG] 2.1.15 自体は通常のアップデート。

* gpg: Remove the `--tofu-db-format` option and support for the split TOFU database.
* gpg: Add option `--sender` to prepare for coming features.
* gpg: Add option `--input-size-hint` to help progress indicators.
* gpg: Extend the PROGRESS status line with the counted unit.
* gpg: Avoid publishing the GnuPG version by default with `--armor`.
* gpg: Properly ignore legacy keys in the keyring cache.
* gpg: Always print fingerprint records in `--with-colons` mode.
* gpg: Make sure that keygrips are printed for each subkey in `--with-colons` mode.
* gpg: New import filter "`drop-sig`".
* gpgsm: Fix a bug in the machine-readable key listing.
* gpg,gpgsm: Block signals during keyring updates to limits the effects of a Ctrl-C at the wrong time.
* g13: Add command `--umount` and other fixes for dm-crypt.
* agent: Fix regression in SIGTERM handling.
* agent: Cleanup of the ssh-agent code.
* agent: Allow import of overly long keys.
* scd: Fix problems with card removal.
* dirmngr: Remove all code for running as a system service.
* tools: Make gpg-wks-client conforming to the specs.
* tests: Improve the output of the new regression test tool.
* tests: Distribute the standalone test runner.
* tests: Run each test in a clean environment.
* Spelling and grammar fixes.

[Libgcrypt]: https://www.gnu.org/software/libgcrypt/ "Libgcrypt - GNU Project - Free Software Foundation (FSF)"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考文献

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
