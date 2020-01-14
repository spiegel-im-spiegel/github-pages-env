+++
title = "Libgcrypt 1.8.5 がリリース【セキュリティ・アップデート】"
date =  "2019-08-30T21:16:14+09:00"
description = "一般的にサイドチャネル攻撃は成立条件が複雑になることが多いので深刻度は高くならないのだが，計画的にアップデートを行ってほしい。"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "cryptography",
  "library",
  "gnupg",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] でも使われている GNU の暗号ライブラリ [Libgcrypt] の 1.8.5 がリリースされた。

- [[Announce] Libgcrypt 1.8.5 released](https://lists.gnupg.org/pipermail/gnupg-announce/2019q3/000440.html)

今回はセキュリティ・アップデートとなる。

- [CVE-2019-13627](https://nvd.nist.gov/vuln/detail/CVE-2019-13627)

{{< fig-quote type="markdown" title="Libgcrypt 1.8.5 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q3/000440.html" lang="en" >}}
{{< quote >}}This release fixes an ECDSA side-channel attack.{{< /quote >}}
{{< /fig-quote >}}

以上を含む主な変更点は以下のとおりだ。

{{< fig-quote type="markdown" title="Libgcrypt 1.8.5 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q3/000440.html" lang="en" >}}
- Bug fixes:
   - Add mitigation against an ECDSA timing attack. [T4626,CVE-2019-13627]
   - Improve ECDSA unblinding.
-  Other features:
   - Provide a pkg-config file for libgcrypt.

Release-info: https://dev.gnupg.org/T4683
{{< /fig-quote >}}

一般的にサイドチャネル攻撃は成立条件が複雑になることが多いので深刻度は高くならないのだが，計画的にアップデートを行ってほしい。

Linux 等では [Libgcrypt] のみを入れ替えれば済む（筈）なのだが， Windows 版の [GnuPG] などはパッケージに [Libgcrypt] のバイナリが含まれているので，今後のリリース情報に注意すること。

さて [Ubuntu] はどうするんだろうねぇ。
[GnuPG] 2.2.17 の[セキュリティ・アップデート]({{< ref "/release/2019/07/gnupg-2_2_17-is-released.md" >}} "GnuPG 2.2.17 リリース： 公開鍵サーバ・アクセスに関する過激な変更あり")に関してもまるっと無視している。
ホンマ [Ubuntu] ってセキュリティに無頓着だよなぁ。

[GnuPG] や [Libgcrypt] をビルドする環境を早急に整える必要があるのだが，今はちょっと無理なんだよね。
もうしばらくは我慢我慢我慢。

## 追記 2020-0114

[Ubuntu] がようやく今回の脆弱性に対応したようだが...

- [USN-4236-1: Libgcrypt vulnerability | Ubuntu security notices](https://usn.ubuntu.com/4236-1/)

[Libgcrypt] 1.8.4 へのバックポート・パッチかよ `orz`

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Libgcrypt]: https://gnupg.org/software/libgcrypt/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
