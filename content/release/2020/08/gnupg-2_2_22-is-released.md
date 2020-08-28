+++
title = "GnuPG 2.2.22 がリリースされた"
date =  "2020-08-28T11:38:35+09:00"
description = "メンテナンス・リリース。 セキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
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

[GnuPG] 2.2.22 がリリースされた。
ゾロ目。
にゃん×4（笑）

- [[Announce] GnuPG 2.2.22 released](https://lists.gnupg.org/pipermail/gnupg-announce/2020q3/000447.html)

メンテナンス・リリース。
セキュリティ・アップデートはなし。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.22 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2020q3/000447.html" lang="en" >}}
* gpg: Change the default key algorithm to rsa3072.
* gpg: Add regular expression support for Trust Signatures on all platforms.  [#4843]
* gpg: Fix regression in 2.2.21 with non-default `--passphrase-repeat` option.  [#4991]
* gpg: Ignore `--personal-digest-prefs` for ECDSA keys.  [#5021]
* gpgsm: Make rsaPSS a de-vs compliant scheme.
* gpgsm: Show also the SHA256 fingerprint in key listings.
* gpgsm: Do not require a default keyring for `--gpgconf-list`.  [#4867]
* gpg-agent: Default to extended key format and record the creation time of keys.  Add new option `--disable-extended-key-format`.
* gpg-agent: Support the `WAYLAND_DISPLAY` envvar.  [#5016]
* gpg-agent: Allow using `--gpgconf-list` even if HOME does not exist.  [#4866]
* gpg-agent: Make the Pinentry work even if the envvar TERM is set to the empty string.  [#4137]
* scdaemon: Add a workaround for Gnuk tokens <= 2.15 which wrongly incremented the error counter when using the "`verify`" command of "`gpg --edit-key`" with only the signature key being present.
* dirmngr: Better handle systems with disabled IPv6.  [#4977]
* gpgpslit: Install tool.  It was not installed in the past to avoid conflicts with the version installed by GnuPG 1.4.  [#5023]
* gpgtar: Handle Unicode file names on Windows correctly (requires libgpg-error 1.39).  [#4083]
* gpgtar: Make `--files-from` and `--null` work as documented.  [#5027]
* Build the Windows installer with the new Ntbtls 0.2.0 so that TLS connections succeed for servers demanding GCM.

Release-info: https://dev.gnupg.org/T5030
{{< /fig-quote >}}

あれ？ 既定のアルゴリズムって既に RSA 3072 ビットじゃなかったっけ？ まぁいいや。
どうせ[新規で鍵を作るなら ECC]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}} "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな") だろうし。

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
