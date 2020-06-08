+++
title = "Edwards-curve Digital Signature Algorithm"
date =  "2020-06-07T17:33:37+09:00"
description = "SP 800-57 Part 1 Rev.5 が正式リリースした記念に EdDSA に関する情報を覚え書きの形で記しておく。"
image = "/images/attention/kitten.jpg"
tags = [
  "security",
  "cryptography",
  "ecc",
  "nist",
  "openpgp",
  "gnupg",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ネットであちこち眺めていて気がついたのだが SP 800-57 第一部の Rev.5 最終版が2020年5月にリリースされていた。

- [SP 800-57 Part 1 Rev. 5, Recommendation for Key Management: Part 1 – General | CSRC](https://csrc.nist.gov/publications/detail/sp/800-57-part-1/rev-5/final)

私はセキュリティや暗号技術の専門家ではないし Rev.5 の変更点が（軽微なものも合わせて）67項目もあって全部は紹介しきれないので，ひとつだけ

{{< fig-quote title="SP 800-57 Part 1 Revision 5" link="https://doi.org/10.6028/NIST.SP.800-57pt1r5" lang="en" >}}
{{% quote %}}In Section 2.2, EdDSA was added. Modified ECDSA{{% /quote %}}.
{{< /fig-quote >}}

について関連情報を覚え書きの形で記しておく。

## [RFC 8032]: Edwards-Curve Digital Signature Algorithm (EdDSA)

EdDSA のオリジナルは2011年に公開された “{{< pdf-file title="High-speed high-security signatures" link="http://ed25519.cr.yp.to/ed25519-20110926.pdf" >}}” らしい。
その後も改良版が出たり色々あって，2017年に [RFC 8032] として標準化された。

EdDSA は以下の特徴を持っている。

{{< fig-quote type="markdown" title="RFC 8032: Edwards-Curve Digital Signature Algorithm (EdDSA)" link="https://www.rfc-editor.org/rfc/rfc8032.html" lang="en" >}}
1.  EdDSA provides high performance on a variety of platforms;
2.  The use of a unique random number for each signature is not required;
3.  It is more resilient to side-channel attacks;
4.  EdDSA uses small public keys (32 or 57 bytes) and signatures (64 or 114 bytes) for Ed25519 and Ed448, respectively;
5.  The formulas are "complete", i.e., they are valid for all points on the curve, with no exceptions.  This obviates the need for EdDSA to perform expensive point validation on untrusted public values; and
6.  EdDSA provides collision resilience, meaning that hash-function collisions do not break this system (only holds for PureEdDSA).
{{< /fig-quote >}}

特に2番目が重要。

これまでの NIST 標準の（ECDSA を含む） DSA は署名の度にランダムな値をひとつ決めないといけないのだが，ここの実装をサボると，最悪の場合，秘密鍵の漏洩に繋がる。
更に言うと，かつて SP 800-90A に載っていた疑似乱数生成器 Dual_EC_DRBG に NSA の関与が疑われる欠陥が発覚し [SP 800-90A を改訂](https://csrc.nist.gov/publications/detail/sp/800-90a/rev-1/final "SP 800-90A Rev. 1, Random Number Generation Using Deterministic RBGs | CSRC")する騒ぎにまで発展したこともある。

このように DSA の具体的な実装について常に懸念が付きまとっているため，より安全性が高いとされる EdDSA が注目されることとなった[^risk1]。

[^risk1]: DSA に対する懸念は，実装に必要な技術要素が多すぎる点にあるかもしれない。如何にシンプルな設計で要件を満たすかってのは安全性を考える上でも重要であるというよい事例になっていると思う。

EdDSA で使える楕円曲線には{{< pdf-file title="かなりのバリエーションがある" link="http://ed25519.cr.yp.to/eddsa-20150704.pdf" >}} のだが， [RFC 8032] に記載されているのは以下の2つである[^cv25519]。

[^cv25519]: edwards25519 は [Curve25519] と双有理同値である。ちなみに [Curve25519] は ECDH 用の楕円曲線およびそのライブラリで，公有（public domain）のソフトウェアとして公開されている。

| 楕円曲線名   | 鍵長（bytes） | 強度（bits） |
| ------------ | -------------:| ------------:|
| edwards25519 |            32 |          128 |
| edwards448   |            57 |          224 |

生成した ECC 鍵を2031年以降も使い続けるにはセキュリティ強度にして128ビット以上必要だが，この2つの楕円曲線であれば十分であることが分かる。

## NIST 標準としての EdDSA

EdDSA の RFC 化に伴い，以下の NIST 標準文書にも EdDSA が追加されることになった。
ただし今のところはまだドラフト版である。

- [FIPS 186-5 (Draft), Digital Signature Standard (DSS) | CSRC](https://csrc.nist.gov/publications/detail/fips/186/5/draft)
- [SP 800-186 (Draft), Discrete Logarithm-Based Crypto: Elliptic Curve Parameters | CSRC](https://csrc.nist.gov/publications/detail/sp/800-186/draft)

ちなみに両方共パブリックコメントは2020年初に締め切られている。

FIPS 186-5 および SP 800-186 が正式リリースされれば，政府調達でもなんでも，大手を振って EdDSA を使えるようになる（笑）

## EdDSA の実装

### [OpenPGP] と [GnuPG]

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 6637]: https://tools.ietf.org/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

[OpenPGP] は [RFC 6637] で正式に ECC を組み込んだが，この中に EdDSA は含まれていない。
ただし，次期 [OpenPGP] となる [RFC 4880bis] では EdDSA を組み込み済みで [GnuPG] の最新版では既に EdDSA 鍵を生成し使用することができる。

- [そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< ref "http://localhost:1313/openpgp/using-ecc-with-gnupg.md" >}})

なお，最新の [GnuPG] では edwards25519 のみサポートしているようだ。

### [OpenSSH]

[OpenSSH]: https://www.openssh.com/

[OpenSSH] では EdDSA/edwards25519 鍵を生成・使用できる。
鍵生成は

```text
$ ssh-keygen -t ed25519
```

で可能。

### [OpenSSL]

[OpenSSL]: https://www.openssl.org/

現在の [OpenSSL] は EdDSA をサポートしている。
edwards25519 および edwards448 を指定可能。

## 【おまけ】 量子コンピュータ耐性

現在，公開鍵暗号の主流である IFC (Integer Factorization Cryptosystems) および FFC (Finite Field Cryptosystems) の各アルゴリズムは量子コンピュータによる攻略法が既にあり，十分な性能を獲得すれば短時間で攻略可能になると考えられている。
ECC も FFC のバリエーションであり，これに含まれる。

しかし，2020年時点の性能では128ビット以上のセキュリティ強度であれば現実的な脅威には至っていないようだ。

{{< fig-quote type="markdown" title="現在の量子コンピュータによる暗号技術の安全性への影響" link="https://www.cryptrec.go.jp/topics/cryptrec-er-0001-2019.html" >}}
{{% quote %}}例えば、量子コンピュータを用いて2048ビットRSA合成数の素因数分解を行う場合には、量子誤りが一切ないという理想的な環境下でも、4098量子ビットが必要であり、1012～1013回のゲート演算が必要であると見積もられています。また、量子誤りがあるという現実的な環境下では、2000万量子ビットが必要であるという見積もりもあります{{% /quote %}}。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="現在の量子コンピュータによる暗号技術の安全性への影響" link="https://www.cryptrec.go.jp/topics/cryptrec-er-0001-2019.html" >}}
{{% quote %}}量子コンピュータの性能を測る上での指標（量子ビット数、量子誤りの大きさ、演算可能回数など）や、量子コンピュータの開発状況もあわせて考慮にいれると、近い将来に、2048ビットの素因数分解や256ビットの楕円曲線上の離散対数問題が解かれる可能性は低いと考えます{{% /quote %}}。
{{< /fig-quote >}}

ただし，量子コンピュータ開発は成長が著しい分野でもあり，今後も成り行きを注視していく必要がある。

- {{< pdf-file title="CRYPTREC Report 2018: 暗号技術評価委員会報告" link="https://www.cryptrec.go.jp/report/cryptrec-rp-2000-2018.pdf" >}} : 素因数分解問題および楕円曲線上の離散対数問題の困難性に関する計算量評価のレポートあり
- {{< pdf-file title="耐量子計算機暗号の研究動向調査報告書" link="https://www.cryptrec.go.jp/report/cryptrec-tr-2001-2018.pdf" >}}

## ブックマーク

- [Ed25519: high-speed high-security signatures](http://ed25519.cr.yp.to/)
- [擬似乱数生成アルゴリズム Dual_EC_DRBG について](https://www.cryptrec.go.jp/topics/cryptrec-er-0001-2013.html)

- [暗号鍵関連の各種変数について]({{< ref "/remark/2017/10/key-parameters.md" >}})
- [OpenSSH 鍵をアップグレードする（さようなら SHA-1）]({{< ref "/remark/2020/06/upgrade-openssh-key.md" >}})

[RFC 8032]: https://www.rfc-editor.org/rfc/rfc8032.html "RFC 8032: Edwards-Curve Digital Signature Algorithm (EdDSA)"
[Curve25519]: http://cr.yp.to/ecdh.html "Curve25519: high-speed elliptic-curve cryptography"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
