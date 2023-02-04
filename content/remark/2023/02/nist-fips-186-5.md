+++
title = "NIST FIPS 186-5 および SP 800-186 正式版がリリースされた"
date =  "2023-02-04T09:30:31+09:00"
description = "DSA が標準から外れる / ECDSA 自体を定義し直す / EdDSA が標準として承認された"
image = "/images/attention/kitten.jpg"
tags = [ "security", "cryptography", "ecc", "nist", "dsa" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

長い間ドラフト状態だった FIPS 186-5 および SP 800-186 の最終版が出たようだ。

- [FIPS 186-5, Digital Signature Standard (DSS) | CSRC](https://csrc.nist.gov/publications/detail/fips/186/5/final)
- [SP 800-186, Discrete Logarithm-Based Crypto: Elliptic Curve Parameters | CSRC](https://csrc.nist.gov/publications/detail/sp/800-186/final)

この記事では [FIPS 186-5] に注目して書いてみる。

大きな変更は3つ。

## DSA が標準から外れる

{{< fig-quote type="markdown" title="“Digital Signature Standard (DSS)” section 4" link="https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf" lang="en" >}}
Prior versions of this standard specified the DSA. This standard no longer approves the DSA for digital signature generation. However, the DSA may be used to verify signatures generated prior to the implementation date of this standard. See FIPS 186-4 [7] for the specifications for the DSA.
{{< /fig-quote >}}

ということで予定通り，電子署名生成用途としては， DSA は NIST 標準から外れることになった。
ただし，以前に DSA で作成された署名を検証する必要があるため，アプリケーションによっては実装を残しておく必要がある。
DSA 実装に関しては旧版の FIPS 186-4 を読めってあるな。

## ECDSA 自体を定義し直す

今まで ECDSA の実装については，以前の FIPS 186-4 だけでは完結してなくて， ANSI X9.62 なんかも併せて参照する必要があった。
特に ANSI X9.62 は所謂 paywalled document って奴で，誰でも気軽に見れるもんじゃないのね。

まだ中身をちゃんと読んでないのだが，この版ではそういった外部参照を取り払って [FIPS 186-5] のみで実装を定義しているようだ。
めでたい！ 言い方を変えると，今後 ECDSA を実装する際は [FIPS 186-5] を見ろってことでもある。
まぁ，私を含めて殆どの人は偉い人が作ったライブラリを利用するだけで，中身を気にする人は少ないだろうけど（笑）

## EdDSA が標準として承認された

そして最大のトピックは EdDSA が NIST 標準として承認されたことだろう。

ちなみに EdDSA については，日本の [CRYPTREC でも評価](https://www.cryptrec.go.jp/topics/cryptrec_20211012_c20report.html "CRYPTREC | CRYPTREC Report 2020")が行われていて

{{< fig-quote type="markdown" title="CRYPTREC Report 2020 暗号技術評価委員会報告" link="https://www.cryptrec.go.jp/report/cryptrec-rp-2000-2020.pdf" >}}
- 総評：下記の観点から、EdDSA の構成に関わる安全性において、EdDSA が ECDSA に劣ると考えられる点は無いと思われると述べている。
  - Schnorr 署名をもとに EdDSA は構成されているため、ランダムオラクルモデルで安全性が証明されている Schnorr 署名に対する安全性評価を参考にすることができる。
  - Schnorr 署名との大きな違いはノンスの生成方法であるが、EdDSA におけるノンスの生成方法は、署名の内部乱数を弱い疑似乱数生成器に委ねることによる危険を排除し、現実的な脅威を回避するための配慮が施されている。
  - 比較対象となる ECDSA については、既存結果として generic group model でのみ安全性が証明されている。
{{< /fig-quote >}}

と書かれている。
まぁ，これで政府調達でも何でも大手を振って EdDSA を使えるというものである。

めでたい！ （大事なことなので2度言いました）

## ブックマーク

- [Edwards-curve Digital Signature Algorithm]({{< ref "/remark/2020/06/eddsa.md" >}})
- [DSA は NIST 電子署名標準から外れるようだ]({{< ref "/remark/2021/02/dsa-deprecation.md" >}})
- [量子コンピュータで256ビット楕円曲線暗号は破れるか]({{< ref "/remark/2022/02/breaking-256-bit-elliptic-curve-encryption-with-a-quantum-computer.md" >}})

[FIPS 186-5]: https://csrc.nist.gov/publications/detail/fips/186/5/final "FIPS 186-5, Digital Signature Standard (DSS) | CSRC"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
