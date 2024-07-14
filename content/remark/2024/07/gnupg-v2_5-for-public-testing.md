+++
title = "公開テスト用の Gnupg v2.5 が登場"
date =  "2024-07-14T10:29:46+09:00"
description = "v2.6 系が出るまでは現行の v2.4 系が安定版としてメンテナンスされる筈なので，しばらくは現状の v2.4 系のままでも大丈夫。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

少し前の話で恐縮だが，公開テスト用の [GnuPG] 2.5.0 が登場した。

- [[Announce] GnuPG 2.5.0 released for public testing](https://lists.gnupg.org/pipermail/gnupg-announce/2024q3/000484.html)

これまでと同じく v2.5 系で検証と改良を行い，安定版として v2.6 系がリリースされる流れだと思う。

現行の v2.4 系との差異は以下の通り。

{{< fig-quote type="markdown" title="[Announce] GnuPG 2.5.0 released for public testing" link="https://lists.gnupg.org/pipermail/gnupg-announce/2024q3/000484.html" lang="en" >}}
[compared to version 2.4.5]

* gpg: Support composite Kyber+ECC public key algorithms.  This is experimental due to the yet outstanding FIPS-203 specification.  [T6815]
* gpg: Allow algo string "`pqc`" for `--quick-gen-key`.  [rG12ac129a70]
* gpg: New option `--show-only-session-key`.  [rG1695cf267e]
* gpg: Print designated revokers also in non-colon listing mode.  [rG9d618d1273]
* gpg: Make `--with-sig-check` work with `--show-key` in non-colon listing mode.  [rG0c34edc443]
* tpm: Rework error handling and fix key import [T7129, T7186]
* Varous fixes to improve robustness on 64 bit Windows.  [T7139]

Changes which will also show up in the firthcoming 2.4.6:

* gpg: New command `--quick-set-ownertrust`.  [rG967678d972]
* gpg: Indicate disabled keys in key listings and add list option "`show-ownertrust`".  [rG2a0a706eb2]
* gpg: Make sure a `DECRYPTION_OKAY` is never issued for a bad OCB tag.  [T7042]
* gpg: Do not allow to accidently set the RENC usage.  [T7072]
* gpg: Accept armored files without CRC24 checksum.  [T7071]
* gpg: New `--import-option` "`only-pubkeys`".  [T7146]
* gpg: Repurpose the AKL mechanism "`ldap`" to work like the keyserver mechnism but only for LDAP keyservers.  [rG068ebb6f1e]
* gpg: ADSKs are now configurable for new keys.  [T6882]
* gpgsm: Emit user IDs with an empty Subject also in colon mode.  [T7171]
* agent: Consider an empty pattern file as valid.  [rGc27534de95]
* agent: Fix error handling of `READKEY`.  [T6012]
* agent: Avoid random errors when storing key in ephemeral mode.  [T7129, rGfdc5003956]
* agent: Make "`SCD DEVINFO --watch`" more robust.  [T7151]
* scd: Improve KDF data object handling for OpenPGP cards.  [T7058]
* scd: Avoid buffer overrun with more than 16 PC/SC readers.  [T7129, rG4c1b007035]
* scd: Fix how the scdaemon on its pipe connection finishes.  [T7160]
* gpgconf: Check readability of some files with -X and change its output format.  [rG98e287ba6d]
* gpg-mail-tube: New tool to apply PGP/MIME encryption to a mail.  [rG28a080bc9f]
* Fix some uninitialized variables and double frees in error code paths.  [T7129]
{{< /fig-quote >}}

v2.6 系が出るまでは現行の v2.4 系が安定版としてメンテナンスされる筈なので，しばらくは現状の v2.4 系のままでも大丈夫。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
