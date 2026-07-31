+++
title = "GnuPG 2.5.17 のリリース【セキュリティ・アップデート】"
date =  "2026-01-28T16:02:53+09:00"
description = "2.5.17 では1件の重要な脆弱性の修正がある"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg", "vulnerability"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

セキュリティ・アップデートを含む [GnuPG] 2.5.17 がリリースされた。

- [[Announce] GnuPG and Gpg4win Security Advisory (T8044)](https://lists.gnupg.org/pipermail/gnupg-announce/2026q1/000501.html)

{{< fig-quote type="markdown" title="[Announce] GnuPG and Gpg4win Security Advisory (T8044)" link="https://lists.gnupg.org/pipermail/gnupg-announce/2026q1/000501.html" lang="en" >}}
A crafted CMS (S/MIME) EnvelopedData message carrying an oversized
wrapped session key can cause a stack buffer overflow in gpg-agent
during the `PKDECRYPT--kem=CMS` handling.  This can easily be used for a
DoS but, worse, the memory corruption can very likley also be used to
mount a remote code execution attack.  The bug was introduced while
changing an internal API to the FIPS required KEM API.

A CVE-id has not been assigned.  We track this bug as T8044 under
https://dev.gnupg.org/T8044.  This vulnerability was discovered by:
OpenAI Security Research.  Their report was received on 2026-01-18;
fixed versions released 2026-01-27.
{{< /fig-quote >}}

2026-01-27 時点ではまだ CVE-ID は割り当てられていない。
脆弱性の影響を受けるバージョンは以下の通り：

{{< fig-quote class="nobox" type="markdown" title="[Announce] GnuPG and Gpg4win Security Advisory (T8044)" link="https://lists.gnupg.org/pipermail/gnupg-announce/2026q1/000501.html" lang="en" >}}
```
These versions are affected:

 - GnuPG 2.5.16          (released 2025-12-30)
 - GnuPG 2.5.15          (released 2025-12-29)
 - GnuPG 2.5.14          (released 2025-11-19)
 - GnuPG 2.5.13          (released 2025-10-22)
 - Gpg4win 5.0.0         (released 2026-01-14)
 - Gpg4win 5.0.0-beta479 (released 2026-01-02)
 - Gpg4win 5.0.0-beta476 (released 2025-12-22)
 - Gpg4win 5.0.0-beta395 (released 2025-10-22)

All other versions are not affected.
```
{{< /fig-quote >}}

2.4 系は脆弱性の影響を受けない。
なお 2.5 系は安定版に昇格している。
さらに 2.4 系は 2026-06-30 でサポートが終了する。

{{< fig-quote type="markdown" title="[Announce] GnuPG 2.5.16 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2025q4/000500.html" lang="en" >}}
Note that the 2.5 series is now declared the stable version of GnuPG.
The oldstable 2.4 series will reach end-of-life in just 6 months.

The main features in the 2.5 series are improvements for 64 bit Windows
and the introduction of Kyber (aka ML-KEM or FIPS-203) as PQC encryption
algorithm.  Other than PQC support the 2.6 series will not differ a lot
from 2.4 because the majority of changes are internal to make use of
newer features from the supporting libraries.
{{< /fig-quote >}}

今回の脆弱性に関連して，以下のドキュメントも公開されている。

- [Responsible Disclosure Requires Accountability: Putting Recent GnuPG Security Reports into Context](https://gnupg.com/20260122-39C3_reply_gpg_fail.html)

これは

{{< fig-quote type="markdown" title="Responsible Disclosure Requires Accountability: Putting Recent GnuPG Security Reports into Context" link="https://gnupg.com/20260122-39C3_reply_gpg_fail.html" lang="en" >}}
A talk at the 39th Chaos Communication Congress (39C3) presented several observations related to signature verification in GnuPG. These observations are documented on the accompanying website gpg.fail.

The way the findings were presented creates the impression that GnuPG is affected by serious security flaws and that we, as maintainers, are failing to meet our responsibilities.

In fact, we carefully reviewed each reported observation and provided timely feedback to the security researchers involved. However, this did not result in a constructive technical exchange.
{{< /fig-quote >}}

ということで gpg.fail に対する公式な応答・評価となっている[^g1]。
詳しくはドキュメントを読んでもらうとして [Kagi Assistant] による要約は以下の通り。

[^g1]: ドメインが `gnupg.com` となっているが，これは [GnuPG] を中核技術とした暗号製品の商用展開を行っている（主にドイツで）サイトである。つまり，このドキュメントは顧客へのメッセージにもなっているわけやね。

{{< fig-quote type="markdown" >}}
- GnuPG側は、39C3での発表に基づく gpg.fail の観察を文脈付きで評価し、指摘された問題の多くは実装上の致命的な脆弱性ではなく、使用方法やソーシャルエンジニアリング、検証不足に起因すると結論付けている
- 実際のバグは限定的で、既知の問題は過去に修正済みのものが多く、一部は互換性やUI安定性といった設計方針によるものであると説明している
- 全体として GnuPG は安全性と安定性を維持しつつ、透明なコミュニケーションと責任ある公開（responsible disclosure）が重要であり、悪用の可能性がある領域については利用者教育と適切な使用方法の周知を重視している
- 結論として、観察リストの大半は実装欠陥ではなく、適切な検証・運用の欠如や社会工学的リスクに起因するものとされている
{{< /fig-quote >}}

ドキュメントで言及されているクリアテキスト署名（Cleartext Signatures）の問題については拙文でも言及しているので参考にどうぞ。

- [クリアテキスト署名の危険性]({{< ref "/openpgp/cleartext-signatures-considered-harmful.md" >}})

最近では [GnuPG] 以外にも様々な OpenPGP 実装が存在するので，それらについて調べてみてもいいかもしれない（というか調べなきゃなぁとは思っている）。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"
[OpenPGP]: http://openpgp.org/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Kagi Assistant]: https://kagi.com/assistant "The Assistant"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
