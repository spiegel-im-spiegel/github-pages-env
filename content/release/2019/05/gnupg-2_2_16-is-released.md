+++
title = "GnuPG 2.2.16 がリリースされた"
date =  "2019-05-29T19:11:01+09:00"
description = "今回もセキュリティ・アップデートはなし。"
image = "/images/attention/tools.png"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.16 がリリースされた。

- [[Announce] GnuPG 2.2.16 released](https://lists.gnupg.org/pipermail/gnupg-announce/2019q2/000438.html)

今回もセキュリティ・アップデートはなし。
主な機能追加・修正点は以下の通り。

{{< fig-quote type="markdown" title="GnuPG 2.2.16 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q2/000438.html" lang="en" >}}
* gpg,gpgsm: Fix deadlock on Windows due to a keybox sharing violation.  [#4505]
* gpg: Allow deletion of subkeys with `--delete-key`.  This finally makes the bang-suffix work as expected for that command.  [#4457]
* gpg: Replace SHA-1 by SHA-256 in self-signatures when updating them with `--quick-set-expire` or `--quick-set-primary-uid`. [#4508]
* gpg: Improve the photo image viewer selection.  [#4334]
* gpg: Fix decryption with `--use-embedded-filename`.  [#4500]
* gpg: Remove hints on using the `--keyserver` option.  [#4512]
* gpg: Fix export of certain secret keys with comments.  [#4490]
* gpg: Reject too long user-ids in `--quick-gen-key`.  [#4532]
* gpg: Fix a double free in the best key selection code.  [#4462]
* gpg: Fix the key generation dialog for switching back from EdDSA to ECDSA.
* gpg: Use AES-192 with SHA-384 to comply with RFC-6637.
* gpg: Use only the addrspec from the Signer's UID subpacket to mitigate a problem with another implementation.
* gpg: Skip invalid packets during a keyring listing and sync diagnostics with the output.
* gpgsm: Avoid confusing diagnostic when signing with the default key.  [#4535]
* agent: Do not delete any secret key in `--dry-run mode`.
* agent: Fix failures on 64 bit big-endian boxes related to URIs in a keyfile.  [#4501]
* agent: Stop scdaemon after a reload with disable-scdaemon newly configured.  [#4326]
* dirmngr: Improve caching algorithm for WKD domains.
* dirmngr: Support other hash algorithms than SHA-1 for OCSP.  [#3966]
* gpgconf: Make `--homedir` work for `--launch`.  [#4496]
* gpgconf: Before `--launch` check for a valid config file.  [#4497]
* wkd: Do not import more than 5 keys from one WKD address.
* wkd: Accept keys which are stored in armored format in the directory.
* The installer for Windows now comes with signed binaries.

Release-info: https://dev.gnupg.org/T4509
{{< /fig-quote >}}

着々と SHA-1 からの置き換えが進んでる感じだねぇ。

[Ubuntu] で最新版バイナリを入手する方法を考えないと。
やっぱ自前でビルドするしかないのかなぁ。

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< ref "/openpgp" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Libgcrypt]: https://gnupg.org/software/libgcrypt/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
