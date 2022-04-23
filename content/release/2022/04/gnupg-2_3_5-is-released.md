+++
title = "GnuPG 2.3.5 のリリース"
date =  "2022-04-23T13:03:03+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

毎度遅まきながらで申し訳ないが [GnuPG] 2.3.5 がリリースされている。

- [[Announce] GnuPG 2.3.5 released](https://lists.gnupg.org/pipermail/gnupg-announce/2022q2/000472.html)

セキュリティ・アップデートはなし。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.3.5 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2022q2/000472.html" lang="en" >}}
* gpg: Up to five times faster verification of detached signatures. Doubled detached signing speed.  [T5826,rG4e27b9defc,rGf8943ce098]
* gpg: Threefold decryption speedup for large files.  [T5820,rGab177eed51]
* gpg: Nearly double the AES256.OCB encryption speed.  [rG99e2c178c7]
* gpg: Removed EAX from the preference list.  [rG253fcb9777]
* gpg: Allow `--dearmor` to decode all kinds of armor files.  [rG34ea19aff9]
* gpg: Remove restrictions for the name part of a `user-id`.  [rG8945f1aedf]
* gpg: Allow decryption of symmetric encrypted data even for non-compliant cipher.  [rG8631d4cfe2]
* gpg,gpgsm: New option `--require-compliance`.  [rGee013c5350]
* gpgsm: New option -`-ignore-cert-with-oid`.  [rGe23dc755fa]
* gpgtar: Create and handle extended headers to support long file names.  [T5754]
* gpgtar: Support file names longer than `MAX_PATH` on Windows.  [rG70b738f93f]
* gpgtar: Use a pipe for decryption and thus avoid memory exhaustion.  [rGe5ef5e3b91]
* gpgtar: New option `--with-log`.  [rGed53d41b4c]
* agent: New flag "qual" for the trustlist.txt.  [rG7c8c606061]
* scdaemon: Add support for GeNUA cards.  [rG0dcc249852]
* scdaemon: Add `--challenge-response` option to `PK_AUTH` for OpenPGP cards.  [T5862]
* dirmngr: Support the use of ECDSA for CRLs and OCSP.  [rGde87c8e1ea,rG890e9849b5]
* dirmngr: Map all gnupg.net addresses to the Ubuntu keyserver.  [T5751]
* ssh: Return a faked response for the new session-bind extension.  [T5931]
* gpgconf: Add command aliases `-L` `-K` `-R`.  [rGec4a1cffb8]
* gpg: Request keygrip of key to add via command interface.  [T5771]
* gpg: Print Yubikey version correctly.  [T5787]
* gpg: Always use version >= 4 to generate key signature.  [T5809]
* gpg: Fix generating AEAD packet.  [T5853]
* gpg: Fix version on symmetric encrypted AEAD files if the force option is used.  [T5856]
* gpg: Fix adding the list of ultimate trusted keys.  [T5742]
* gpgsm: Fix parsing of certain PKCS#12 files.  [T5793]
* gpgsm: Print diagnostic about CRL problems due to Tor mode.  [rG137e59a6a5]
* agent: Use "`Created:`" field for creation time.  [T5538]
* scdaemon Fix error handling for a PC/SC reader selected with `reader-port`.  [T5758]
* scdaemon: Fix `DEVINFO` with no `--watch`.  [rGc6dd9ff929]
* scdaemon: Fix socket resource leak on Windwos.  [T5029]
* scdaemon: Use extended mode for pkcs#15 already for rsa2048.  [rG597253ca17]
* scdaemon: Enhance `PASSWD` command to accept `KEYGRIP` optionally.  [T5862]
* scdaemon: Fix memory leak in ccid-driver.  [rG8ac92f0e80]
* tpm: Always use hexgrip when storing a key password.  [rGaf2fbd9b01]
* dirmngr: Make WKD lookups work for resolvers not handling SRV records.  [T4729]
* dirmngr: Avoid initial delay on the first keyserver access in presence of `--no-use-tor`.  [rG57d546674d]
* dirmngr: Workaround for a certain broken LDAP URL.  [rG90caa7ad59]
* dirmngr: Escape more characters in WKD requests.  [T5902]
* dirmngr: Suppress error message on trial reading as PEM format.  [T5531]
* gpgconf: Fix component table when not building without TPM support.  [T5701]
* gpgconf: Silence warnings from parsing the option files.  [T5874]
* gpgconf: Do not list ignored options and mark forced options as `read-only`.  [rG42785d7c8a]
* gpgconf: Tweak the use of the ldapserver option.  [T5801]
* ssh: Fix adding an `ed25519` key with a zero length comment.  [T5794]
* kbx: Fix searching for FPR20 in version 2 blob.  [T5888]
* Fix early homedir creation.  [T5895]
* Improve removing of stale lockfiles under Unix.  [T5884]

Release-info: https://dev.gnupg.org/T5743
{{< /fig-quote >}}

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.45         | 2022-04-21 | {{< icons "check" >}} |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.9 (LTS)  | 2022-02-07 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.10.1       | 2022-03-28 | {{< icons "check" >}} |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.0        | 2021-06-10 |                       |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.3.1        | 2022-04-07 | {{< icons "check" >}} |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.34 (LTS) | 2022-02-07 |                       |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.5        | 2022-04-21 | {{< icons "check" >}} |

現在 [GnuPG] には2.2系と2.3系があり[^gpg14]，2.2系は LTS 版に位置付けられている。
2.3系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されているので，最新機能を試したいのであればこちらを入れるとよいだろう。
なお2.2系は少なくとも2024年末まではサポートが続けられる予定である。
通常運用であれば，当面は2.2系でも問題ない（[ECC も対応]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}} "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな")してるよ）。

[^gpg14]: 厳密には1.4系もあるが，これは legacy 版と位置付けられており，よほどのバグか脆弱性がない限りは更新されない。もし今だに1.4系（あるいは既にサポートされていない2.0/2.1系）を使っているのなら2.2系以降にアップグレードすることを強くお勧めする。

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"
[OpenPGP]: http://openpgp.org/
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
