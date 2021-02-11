+++
title = "Libgcrypt 1.9 系へのアップグレードはちょっと待ちなはれ"
date =  "2021-02-11T11:24:15+09:00"
description = "GnuPG 2.2 系のみに Libgcrypt を使っているのなら今すぐ 1.9 系にアップグレードする必要はない。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg", "security", "vulnerability"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

なんかメチャメチャ忙しい。
この前うっかり休日出勤したら，その後の体調が急落してしまった。
他のプロジェクトメンバにはホンマに申し訳ないが，祝日と休日はカレンダー通り休ませてもらうことにした。
歳をとるとあちこちガタが来るのよ。
マジすんません {{< emoji "ペコン" >}}

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}。

先月の話で恐縮だが，ついに 1.9 系の [Libgcrypt] が安定版としてリリースされた。

- [[Announce] Libgcrypt 1.9.0 relased](https://lists.gnupg.org/pipermail/gnupg-announce/2021q1/000453.html)

これによると

{{< fig-quote type="markdown" title="Libgcrypt 1.9.0 relased" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q1/000453.html" lang="en" >}}
{{% quote %}}We are pleased to announce the availability of Libgcrypt version 1.9.0. This release starts a new stable branch of Libgcrypt with full API and ABI compatibility to the 1.8 series.  Over the last 3 or 4 years Jussi Kivilinna put a lot of work into speeding up the algorithms for the most commonly used CPUs{{% /quote %}}.
{{< /fig-quote >}}

とのことで，現行の 1.8 系との後方互換性を図りつつ，パフォーマンス改善やアルゴリズム追加等が行われているようだ。
詳しくは以下の通り。

{{< fig-quote type="markdown" title="Libgcrypt 1.9.0 relased" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q1/000453.html" lang="en" >}}
* New and extended interfaces:
    - New curves Ed448, X448, and SM2.
    - New cipher mode EAX.
    - New cipher algo SM4.
    - New hash algo SM3.
    - New hash algo variants SHA512/224 and SHA512/256.
    - New MAC algos for Blake-2 algorithms, the new SHA512 variants, SM3, SM4 and for a GOST variant.
    - New convenience function `gcry_mpi_get_ui`.
    - `gcry_sexp_extract_param` understands new format specifiers to directly store to integers and strings.
    - New function `gcry_ecc_mul_point` and curve constants for Curve448 and Curve25519.  [#4293]
    - New function `gcry_ecc_get_algo_keylen`.
    - New control code `GCRYCTL_AUTO_EXPAND_SECMEM` to allow growing the secure memory area.  Also in 1.8.2 as an undocumented feature.

* Performance:
    - Optimized implementations for Aarch64.
    - Faster implementations for Poly1305 and ChaCha.  Also for PowerPC.  [b9a471ccf5,172ad09cbe,#4460]
    - Optimized implementations of AES and SHA-256 on PowerPC. [#4529,#4530]
    - Improved use of AES-NI to speed up AES-XTS (6 times faster). [a00c5b2988]
    - Improved use of AES-NI for OCB.  [eacbd59b13,e924ce456d]
    - Speedup AES-XTS on ARMv8/CE (2.5 times faster).  [93503c127a]
    - New AVX and AVX2 implementations for Blake-2 (1.3/1.4 times faster).  [af7fc732f9, da58a62ac1]
    - Use Intel SHA extension for SHA-1 and SHA-256 (4.0/3.7 times faster).  [d02958bd30, 0b3ec359e2]
    - Use ARMv7/NEON accelerated GCM implementation (3 times faster). [2445cf7431]
    - Use of i386/SSSE3 for SHA-512 (4.5 times faster on Ryzen 7). [b52dde8609]
    - Use 64 bit ARMv8/CE PMULL for CRC (7 times faster).  [14c8a593ed]
    - Improve CAST5 (40% to 70% faster).  [4ec566b368]
    - Improve Blowfish (60% to 80% faster).  [ced7508c85]
{{< /fig-quote >}}

現行の 1.8 系は LTS バージョンという位置付けで今後もサポートされるようだ。
また [GnuPG] 2.2 系も併せて LTS の位置付けに[なっている](https://gnupg.org/download/ "GnuPG - Download")。
おそらく 2.3 系の一般リリースが近いのであろう。

[Libgcrypt] は汎用の暗号ライブラリだが [GnuPG] 以外で使っている事例を寡聞にして知らない。
ほとんどのユーザはそうだと思うが [GnuPG] 2.2 系のみに [Libgcrypt] を使っているのなら今すぐ 1.9 系にアップグレードする必要はない。
しばらくは推移を見守っていたほうがいいだろう。
当分は頻繁なアップデートが行われるだろうし，重大なバグや脆弱性については 1.8 系にもフィードバックされる筈である[^lts1]。

[^lts1]: [Libgcrypt] 1.8 系がいつまでサポートされるかは明示されていないが， 1.7 系のときは1年前には予告があったので，それまでは気にしなくてもいいと思う。

なお [Libgcrypt] 1.9.0 でさっそく脆弱性が見つかって 1.9.1 がリリースされている。

{{< fig-quote type="markdown" title="[Announce] [Security fix] Libgcrypt 1.9.1 relased" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q1/000456.html" lang="en" >}}
{{% quote %}}There is a heap buffer overflow in libgcrypt due to an incorrect assumption in the block buffer management code. Just decrypting some data can overflow a heap buffer with attacker controlled data, no verification or signature is validated before the vulnerability occurs{{% /quote %}}.
{{< /fig-quote >}}

既に [Libgcrypt] 1.9.0 を運用している場合はアップデートしませう。

## アップデートは計画的に

最近の [GnuPG] 関連のアップデート状況は以下の通り。

|   # | パッケージ名                                             | バージョン | 公開日     |         更新          |
| ---:| -------------------------------------------------------- | ---------- | ---------- |:---------------------:|
|   1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.41       | 2020-12-21 |                       |
|  2a | [Libgcrypt](https://gnupg.org/software/libgcrypt/) (LTS) | 1.8.7      | 2020-10-23 |                       |
|  2b | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.9.1      | 2021-01-29 | {{< icons "check" >}} |
|   3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.4      | 2020-10-23 |                       |
|   4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.5.0      | 2020-11-18 |                       |
|   5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6        | 2018-07-16 |                       |
|   6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0      | 2020-08-27 |                       |
|   7 | [GnuPG](https://gnupg.org/software/) (LTS)               | 2.2.27     | 2021-01-11 |                       |

脆弱性が公表されると碌に中身も見ずに「すぐにアップデートしろ」みたいなことを言う馬鹿メディアもあるようだが，20年前の牧歌的な時代ならともかく今は可用性（availability）も明確にセキュリティ・リスクとして認知されているんだから，考えなしにアップデートできるわけ無いだろ！ セキュリティ部品はシステム規模が大きくなる程いざというときの影響が大きい。
勿論いつまでも放置するのは論外だが。

というわけで，アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/
[Libgcrypt]: https://gnupg.org/software/libgcrypt/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
