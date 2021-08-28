+++
title = "GnuPG 2.3.2 がリリースされた"
date =  "2021-08-25T21:23:52+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.3.2 がリリースされた。

- [[Announce] GnuPG 2.3.2 released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q3/000462.html)

セキュリティ・アップデートはなし。
詳細はこちら。
かなりあるぞ。

{{< fig-quote type="markdown" title="GnuPG 2.3.2 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q3/000462.html" lang="en" >}}

* gpg: Allow fingerprint based lookup with `--locate-external-key`. [ec36eca08c]
* gpg: Allow decryption w/o public key but with correct card inserted.  [50293ec2eb]
* gpg: Auto import keys specified with `--trusted-keys`.  [100037ac0f]
* gpg: Do not use `import-clean` for LDAP keyserver imports.  [#5387]
* gpg: Fix mailbox based search via AKL keyserver method.  [4fcfac6feb]
* gpg: Fix memory corruption with `--clearsign` introduced with 2.3.1. [#5430]
* gpg: Use a more descriptive prompt for symmetric decryption. [6dfae2f402]
* gpg: Improve speed of secret key listing.  [40da61b89b]
* gpg: Support keygrip search with traditional keyring.  [#5469]
* gpg: Let `--fetch-key` return an exit code on failure.  [#5376]
* gpg: Emit the `NO_SECKEY` status again for decryption.  [#5562]
* gpgsm: Support decryption of password based encryption (pwri). [eeb65d3bbd]
* gpgsm: Support AES-GCM decryption.  [4980fb3c6d]
* gpgsm: Let `--dump-cert` `--show-cert` also print an OpenPGP fingerprint.  [52bbdc731f]
* gpgsm: Fix finding of issuer in use-keyboxd mode.  [6b76693ff5]
* gpgsm: New option `--ldapserver` as an alias for `--keyserver`. [89df86157e]
* agent: Use SHA-256 for SSH fingerprint by default.  [#5434]
* agent: Fix calling `handle_pincache_put`.  [#5436]
* agent: Fix importing protected secret key.  [#5122]
* agent: Fix a regression in `agent_get_shadow_info_type`.  [#5393]
* agent: Add translatable text for Caps Lock hint.  [#4950]
* agent: New option `--pinentry-formatted-passphrase`.  [#5517]
* agent: Add checkpin inquiry for pinentry.  [#5517,#5532]
* agent: New option `--check-sym-passphrase-pattern`.  [#5517]
* agent: Use the sysconfdir for a pattern file.
* agent: Make `QT_QPA_PLATFORMTHEME=qt5ct` work for the pinentry. [1305baf099]
* dirmngr: LDAP search by a mailbox now ignores revoked keys. [1406f551f1]
* dirmngr: For `KS_SEARCH` return the fingerprint also with LDAP. [#5441]
* dirmngr: Allow for non-URL specified ldap keyservers.  [#5405,#5452]
* dirmngr: New option `--ldapserver`.  [52cf32ce2f]
* dirmngr: Fix regression in `KS_GET` for mail address pattern. [#5497]
* card: New option `--shadow` for the list command.  [2fce99d73a]
* tests: Make sure the built keyboxd is used.  [#5406]
* scd: Fix computing shared secrets for 512 bit curves. [9e24f2a45c]
* scd: Fix unblock PIN by a Reset Code with KDF.  [#5413]
* scd: Fix PC/SC removed card problem.  [8d81fd7c01]
* scd: Recover the partial match for PORTSTR for PC/SC. [53bdc6288f]
* scd: Make sure to release the PC/SC context.  [#5416]
* scd: Fix zero-byte handling in ECC.  [#5163]
* scd: Fix serial number detection for Yubikey 5.  [#5442]
* scd: Add basic support for AET JCOP cards.  [544ec7872a]
* scd: Detect external interference when `--pcsc-shared` is in use. [#5484]
* scd: Fix access to the list of cards.  [#5524]
* gpgconf: Do not list a disabled tpm2d.  [#5408]
* gpgconf: Make runtime changes with different homedir work. [31c0aa2ff3]
* keyboxd: Fix searching for exact mail adddress.  [f79e9540ca]
* keyboxd: Fix searching with multiple patterns.  [101ba4f18a]
* gpgtar: Fix file size computation under Windows.  [14e36bdbe1]
* tools: Extend gpg-check-pattern.  [73c03e0232]
* wkd: Fix client issue with leading or trailing spaces in user-ids.  [b4345f7521]
* Under Windows add a fallback in case the console can't cope with Unicode.  [#5491]
* Under Windows use `LOCAL_APPDATA` for the socket directory.  [#5537]
* Pass `XDG_SESSION_TYPE` and `QT_QPA_PLATFORM` envvars to Pinentry. [#3659]
* Change the default keyserver to `keyserver.ubuntu.com`.  This is a temporary change due to the shutdown of the SKS keyserver pools. [55b5928099]

Release-info: https://dev.gnupg.org/T5405
{{< /fig-quote >}}

2.3 系も既定の鍵サーバが `keyserver.ubuntu.com` に一時変更されたみたいだねぇ。

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.42         | 2021-03-22 |                       |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.8 (LTS)  | 2021-06-02 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.9.4        | 2021-08-22 | {{< icons "check" >}} |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.0        | 2021-06-10 |                       |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0        | 2020-08-27 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.29 (LTS) | 2021-07-04 |                       |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.2        | 2021-08-24 | {{< icons "check" >}} |

現在 [GnuPG] には2.2系と2.3系があり[^gpg14]，2.2系は LTS 版に位置付けられている。
2.3系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されているので，最新機能を試したいのであればこちらを入れるとよいだろう。
なお2.2系は少なくとも2024年末まではサポートが続けられる予定である。
通常運用であれば，当面は2.2系でも問題ない（[ECC も対応]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}} "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな")してるよ）。

[^gpg14]: 厳密には1.4系もあるが，これは legacy 版と位置付けられており，よほどのバグか脆弱性がない限りは更新されない。もし今だに1.4系（あるいは既にサポートされていない2.0/2.1系）を使っているのなら2.2系以降にアップグレードすることを強くお勧めする。

アップデートは計画的に。

## ブックマーク

- [Using a TPM with GnuPG 2.3](https://gnupg.org/blog/20210315-using-tpm-with-gnupg-2.3.html)
- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
