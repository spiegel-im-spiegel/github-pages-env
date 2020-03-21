+++
title = "GnuPG 2.2.20 がリリースされた"
date =  "2020-03-21T19:27:24+09:00"
description = "セキュリティ・アップデートはなし。"
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

[GnuPG] 2.2.20 がリリースされた。

- [[Announce] GnuPG 2.2.20 released](https://lists.gnupg.org/pipermail/gnupg-announce/2020q1/000444.html)

メンテナンス・リリース。
セキュリティ・アップデートはなし。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.20 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2020q1/000444.html" lang="en" >}}
* Protect the error counter against overflow to guarantee that the tools can't be tricked into returning success after an error.
* gpg: Make really sure that `--verify-files` always returns an error.
* gpg: Fix key listing `--with-secret` if a pattern is given.  [#4061]
* gpg: Fix detection of certain keys used as default-key.  [#4810]
* gpg: Fix default-key selection when a card is available.  [#4850]
* gpg: Fix key expiration and key usage for keys created with a creation date of zero.  [#4670]
* gpgsm: Fix import of some CR,LF terminated certificates.  [#4847]
* gpg: New options `--include-key-block` and `--auto-key-import` to allow encrypted replies after an initial signed message.  [#4856]
* gpg: Allow the use of a fingerprint with `--trusted-key`. [#4855]
* gpg: New property `"fpr"` for use by `--export-filter`.
* scdaemon: Disable the pinpad if a KDF DO is used.  [#4832]
* dirmngr: Improve finding OCSP certificates.  [#4536]
* Avoid build problems with LTO or gcc-10. [#4831]

Release-info: https://dev.gnupg.org/T4860
{{< /fig-quote >}}

アップデートは計画的に。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
