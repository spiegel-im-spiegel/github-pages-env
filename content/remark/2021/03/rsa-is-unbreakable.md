+++
title = "RSA は砕けない（たぶん？）"
date =  "2021-03-06T09:33:55+09:00"
description = "まだあわてるような時間じゃない（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "cryptography", "security", "risk", "rsa" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

こっち方面のアンテナがすっかり鈍ってしまって，気付いたのがやっと昨日なのだが

- [Cryptology ePrint Archive: Report 2021/232 - Fast Factoring Integers by SVP Algorithms](https://eprint.iacr.org/2021/232)

という論文を巡ってちょっとした騒ぎになっていたらしい。
いや abstract に {{< quote lang="en" >}}This destroys the RSA cryptosystem{{< /quote >}} とか書いてあるもんで，その筋の方々が色めき立っちゃたようだ（笑）

これについて Bruce Schneier 先生も

{{< fig-quote type="markdown" title="No, RSA Is Not Broken - Schneier on Security" link="https://www.schneier.com/blog/archives/2021/03/no-rsa-is-not-broken.html" lang="en" >}}
{{% quote %}}It does not. At best, it’s an improvement in factoring — and I’m not sure it’s even that{{% /quote %}}.
{{< /fig-quote >}}

とおっしゃっている。
論文自体の査読もまだだし，これから要検証という感じ。
少なくとも[2020年時点で RSA-250 の解読には成功]({{< ref "/remark/2020/04/rsa-250-factored.md" >}})してるんだから，それとの比較で評価できるだろう。

要するに「[まだあわてるような時間じゃない](https://dic.nicovideo.jp/a/%E3%81%BE%E3%81%A0%E3%81%82%E3%82%8F%E3%81%A6%E3%82%8B%E3%82%88%E3%81%86%E3%81%AA%E6%99%82%E9%96%93%E3%81%98%E3%82%83%E3%81%AA%E3%81%84)」ということだろう。
まぁ，でも，これから鍵を替えるなら ECC 鍵がいいんじゃないかな（笑）

- [そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}})
- [SSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.d" >}})
- [Edwards-curve Digital Signature Algorithm]({{< ref "/remark/2020/06/eddsa.md" >}})

## ブックマーク

- [Does Schnorr's 2021 factoring method show that the RSA cryptosystem is not secure? - Cryptography Stack Exchange](https://crypto.stackexchange.com/questions/88582/does-schnorrs-2021-factoring-method-show-that-the-rsa-cryptosystem-is-not-secur)
- [Schnorr's factorization algorithm - Issuance Policy - Let's Encrypt Community Support](https://community.letsencrypt.org/t/schnorrs-factorization-algorithm/146550)

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
