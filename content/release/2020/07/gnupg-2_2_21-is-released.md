+++
title = "GnuPG 2.2.20 および Libgcrypt 1.8.6 がリリースされた"
date =  "2020-07-10T11:13:57+09:00"
description = "AEAD Encrypted Data Packet の復号に対応したか。"
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

[GnuPG] 2.2.20 および [Libgcrypt] 1.8.6 がリリースされた。

- [[Announce] GnuPG 2.2.21 released](https://lists.gnupg.org/pipermail/gnupg-announce/2020q3/000446.html)
- [[Announce] Libgcrypt 1.8.6 released](https://lists.gnupg.org/pipermail/gnupg-announce/2020q3/000445.html)

メンテナンス・リリース。
両者ともセキュリティ・アップデートはなし。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.21 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2020q3/000446.html" lang="en" >}}
* gpg: Improve symmetric decryption speed by about 25%. See commit 144b95cc9d.
* gpg: Support decryption of AEAD encrypted data packets.
* gpg: Add option `--no-include-key-block`. [#4856]
* gpg: Allow for extra padding in ECDH.  [#4908]
* gpg: Only a single pinentry is shown for symmetric encryption if the pinentry supports this.  [#4971]
* gpg: Print a note if no keys are given to `--delete-key`.  [#4959]
* gpg,gpgsm: The ridiculous passphrase quality bar is not anymore shown.  [#2103]
* gpgsm: Certificates without a CRL distribution point are now considered valid without looking up a CRL.  The new option `--enable-issuer-based-crl-check` can be used to revert to the former behaviour.
* gpgsm: Support rsaPSS signature verification.  [#4538]
* gpgsm: Unless CRL checking is disabled lookup a missing issuer certificate using the certificate's authorityInfoAccess.  [#4898]
* gpgsm: Print the certificate's serial number also in decimal notation.
* gpgsm: Fix possible NULL-deref in messages of `--gen-key`.  [#4895]
* scd: Support the CardOS 5 based D-Trust Card 3.1.
* dirmngr: Allow http URLs with "`LOOKUP --url`".
* wkd: Take name of sendmail from configure.  Fixes an OpenBSD specific bug.  [#4886]

Release-info: https://dev.gnupg.org/T4897
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Libgcrypt 1.8.6 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2020q3/000445.html" lang="en" >}}
* Bug fixes:
   - Fix build problems on OpenIndiana et al. [#4818]
   - Fix GCM bug on arm64 which troubles for example OMEMO.  [#4986]
   - Fix wrong code execution in Poly1305 ARM/NEON implementation. [#4833]
   - Detect a div-by-zero in a debug helper tool.  [#4868]
   - Use a constant time mpi_inv in some cases and change the order mpi_invm is called.  [#4869]
   - Fix mpi_copy to correctly handle flags of opaque MPIs.
   - Fix mpi_cmp to consider +0 and -0 the same.
 * Other features:
   - Add OIDs from [RFC-8410](https://tools.ietf.org/html/rfc8410) as aliases for Ed25519 and Curve25519.

 Release-info: https://dev.gnupg.org/T4985
{{< /fig-quote >}}

ふむふむ。
AEAD Encrypted Data Packet (tag 20) の復号に対応したか。

{{% quote lang="en" %}}AEAD Encrypted Data Packet{{% /quote %}} は次期 OpenPGP ([RFC 4880bis]) で導入されるフォーマットで， AEAD (Authenticated Encryption with Associated Data; 認証付き暗号) に対応している。
利用可能な暗号モードは以下の通り。

{{< div-gen >}}
<figure>
<table>
<thead>
<tr><th>ID</th><th>暗号モード</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>1</td>
<td class="nowrap">EAX</td>
<td><a href="https://eprint.iacr.org/2003/069">EAX: A Conventional Authenticated-Encryption Mode</a></td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">OCB</td>
<td><a href="https://tools.ietf.org/html/rfc7253">RFC7253</a></td>
</tr><tr>
<td class='right'>100-110</td>
<td colspan="2">Private/Experimental algorithm</td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な認証付き暗号アルゴリズム一覧</figcaption>
</figure>
{{< /div-gen >}}

[RFC 4880bis] では EAX が MUST になっている。
ちなみに，拙作の [gpgpdump] では {{% quote lang="en" %}}AEAD Encrypted Data Packet{{% /quote %}} を一応識別可能である。

アップデートは計画的に。

## ブックマーク

- [OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）]({{< ref "/openpgp/algorithms-for-openpgp.md" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Libgcrypt]: https://gnupg.org/software/libgcrypt/
[OpenPGP]: http://openpgp.org/
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
