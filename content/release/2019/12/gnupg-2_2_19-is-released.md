+++
title = "GnuPG 2.2.19 がリリースされた"
date =  "2019-12-08T10:16:04+09:00"
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

[GnuPG] 2.2.19 がリリースされた。

- [[Announce] GnuPG 2.2.19 released](https://lists.gnupg.org/pipermail/gnupg-announce/2019q4/000443.html)

この前[リリースされた 2.2.18]({{< ref "/release/2019/11/gnupg-2_2_18-is-released.md" >}}) で regression があったようだ。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.19 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q4/000443.html" lang="en" >}}
* gpg: Fix double free when decrypting for hidden recipients. Regression in 2.2.18.  [#4762].
* gpg: Use `auto-key-locate` for encryption even for mail addressed given with angle brackets.  [#4726]
* gpgsm: Add special case for certain expired intermediate certificates.  [#4696]

Release-info: https://dev.gnupg.org/T4768
{{< /fig-quote >}}

アップデートは計画的に。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
