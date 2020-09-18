+++
title = "GnuPG 2.2.18 リリース： さようなら SHA-1"
date =  "2019-11-26T22:12:19+09:00"
description = "2019-01-19 以降に鍵に付与された SHA-1 ベースの電子署名を全て削除する（CVE-2019-14855）。"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
  "risk",
  "hash",
  "sha-1",
  "gpgpdump",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.18 がリリースされた。

- [[Announce] GnuPG 2.2.18 released](https://lists.gnupg.org/pipermail/gnupg-announce/2019q4/000442.html)

今回は（[GnuPG] 自体には）脆弱性もなく通常のメンテナンス・リリースなのだが，ひとつ大きな修正というか対策があって

{{< fig-quote type="markdown" title="GnuPG 2.2.18 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q4/000442.html" lang="en" >}}
{{< quote >}}This release also retires the use of SHA-1 key signatures created since this year.{{< /quote >}}
{{< /fig-quote >}}

らしい。
厳密には 2019-01-19 以降に鍵に付与された SHA-1 ベースの電子署名を全て削除するというもの（[CVE-2019-14855]）。

{{< div-box type="markdown" >}}
**【2020-09-18 追記】**
[CVE-2019-14855] の CVSS 評価は以下の通り。

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N`
- 深刻度: 重要 (7.5)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更なし     |
| 機密性への影響   | 高           |
| 完全性への影響   | なし         |
| 可用性への影響   | なし         |

ただし Red Hat はもう少し低めに見積もっているようで

- `CVSS:3.0/AV:N/AC:H/PR:N/UI:R/S:U/C:H/I:N/A:N`
- 深刻度: 警告 (5.3)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 高           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 要           |
| スコープ         | 変更なし     |
| 機密性への影響   | 高           |
| 完全性への影響   | なし         |
| 可用性への影響   | なし         |

としている。

[CVE-2019-14855]: https://nvd.nist.gov/vuln/detail/CVE-2019-14855
{{< /div-box >}}

とはいえ，ずいぶん前から [GnuPG] が生成する電子署名は SHA2-256 ベースが既定なので影響は限定的だと思うが[^sig1]， **わざわざ** SHA-1 ベースの電子署名を鍵に付与している方はご注意を。

[^sig1]: ちなみに私が[普段遣いしている鍵](https://baldanders.info/pubkeys/ "OpenPGP 公開鍵リスト")は2013年に作ったものだが， SHA2-256 ベースの電子署名が付与されている。

一応 `--allow-weak-key-signatues` オプションを付けることで今回の措置を回避できるようだが，腹を括ったほうがいいだろう。
なお，鍵への電子署名にどのようなアルゴリズムが使われているかを調べるために拙作の [gpgpdump] を宣伝しておく（笑）

- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

例えば，こんな感じで鍵を取り出して調べることができる。

```text
$ gpg -a --export alice@example.com | gpgpdump
```

その他，詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.18 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q4/000442.html" lang="en" >}}
* gpg: Changed the way keys are detected on a smartcards; this allows the use of non-OpenPGP cards.  In the case of a not very likely regression the new option `--use-only-openpgp-card` is available.  [#4681]
* gpg: The commands `--full-gen-key` and `--quick-gen-key` now allow direct key generation from supported cards.  [#4681]
* gpg: Prepare against chosen-prefix SHA-1 collisions in key signatures.  This change removes all SHA-1 based key signature newer than 2019-01-19 from the web-of-trust.  Note that this includes all key signature created with dsa1024 keys.  The new option `--allow-weak-key-signatues` can be used to override the new and safer behaviour.  [#4755,CVE-2019-14855]
* gpg: Improve performance for import of large keyblocks.  [#4592]
* gpg: Implement a keybox compression run.  [#4644]
* gpg: Show warnings from dirmngr about redirect and certificate problems (details require `--verbose` as usual).
* gpg: Allow to pass the empty string for the passphrase if the '`--passphase=`' syntax is used.  [#4633]
* gpg: Fix printing of the KDF object attributes.
* gpg: Avoid surprises with `--locate-external-key` and certain `--auto-key-locate` settings.  [#4662]
* gpg: Improve selection of best matching key.  [#4713]
* gpg: Delete key binding signature when deletring a subkey.  [#4665,#4457]
* gpg: Fix a potential loss of key sigantures during import with `self-sigs-only` active.  [#4628]
* gpg: Silence "marked as ultimately trusted" diagnostics if option `--quiet` is used.  [#4634]
* gpg: Silence some diagnostics during in key listsing even with option `--verbose`.  [#4627]
* gpg, gpgsm: Change parsing of agent's pkdecrypt results.  [#4652]
* gpgsm: Support AES-256 keys.
* gpgsm: Fix a bug in triggering a keybox compression run if `--faked-system-time` is used.
* dirmngr: System CA certificates are no longer used for the SKS pool if GNUTLS instead of NTBTLS is used as TLS library.  [#4594]
* dirmngr: On Windows detect usability of IPv4 and IPv6 interfaces to avoid long timeouts.  [#4165]
* scd: Fix BWI value for APDU level transfers to make Gemalto Ezio Shield and Trustica Cryptoucan work.  [#4654,#4566]
* wkd: gpg-wks-client `--install-key` now installs the required policy file.

Release-info: https://dev.gnupg.org/T4684
{{< /fig-quote >}}

アップデートは計画的に。

## 【2020-09-18 追記】 [Ubuntu] 18.04 LTS にてバックポート・パッチのリリースあり

LTS 版 [Ubuntu] 18.04 にて [CVE-2019-14855] に対応したバックポートパッチが出ている。

- [USN-4516-1: GnuPG vulnerability | Ubuntu security notices | Ubuntu](https://ubuntu.com/security/notices/USN-4516-1)

今更どういうつもりなのか。
いや，まぁ，出るだけマシだけど。

ちなみに現時点で LTS 最新版である 20.04 では 2.2.19 が提供されている。

```text
gpg --version
gpg (GnuPG) 2.2.19
libgcrypt 1.8.5
Copyright (C) 2019 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: /home/username/.gnupg
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256,
     TWOFISH, CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

現時点の [GnuPG] 最新は 2.2.23 だけどね（笑）

しかし [Ubuntu] の足の遅さは何とかならないのかねぇ。
APT でも使う中核部品だろうに。

## ブックマーク

- ["SHA-1 is a Shambles" and forging PGP WoT signatures](https://mailarchive.ietf.org/arch/msg/openpgp/h-6vCMDFFKhVXpXLC6gAt9tK7r8)

- [OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）]({{< ref "/openpgp/algorithms-for-openpgp.md" >}})
- [SHA-1 衝突問題： 廃止の前倒し]({{< ref "/remark/2015/problem-of-sha1-collision.md" >}})
- [最初の SHA-1 衝突例]({{< ref "/remark/2017/02/sha-1-collision.md" >}})
- [（何度目かの）さようなら SHA-1]({{< ref "/remark/2020/01/sayonara-sha-1.md" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[CVE-2019-14855]: https://nvd.nist.gov/vuln/detail/CVE-2019-14855
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
