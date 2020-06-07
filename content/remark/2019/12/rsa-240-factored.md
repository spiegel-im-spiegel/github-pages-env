+++
title = "RSA-240 が解けたらしい"
date =  "2019-12-07T14:35:56+09:00"
description = "2030年以降を見据えるならそろそろ RSA や ElGamal/DSA 等の古い公開鍵暗号について見直しを始めるべきなんだろうね。"
image = "/images/attention/kitten.jpg"
tags = [ "cryptography", "rsa" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

“[RSA-240 Factored - Schneier on Security](https://www.schneier.com/blog/archives/2019/12/rsa-240_factore.html)” 経由：

RSA challenge list のうち RSA-240 が解けたようだ。

{{< fig-gen type="markdown" title="795-bit factoring and discrete logarithms" link="https://listserv.nodak.edu/cgi-bin/wa.exe?A2=NMBRTHRY;fd743373.1912&FT=M&P=T&H=&S=" >}}

```
We are pleased to announce the factorization of RSA-240, from RSA's challenge
list, and the computation of a discrete logarithm of the same size (795 bits):

RSA-240 = 124620366781718784065835044608106590434820374651678805754818788883289666801188210855036039570272508747509864768438458621054865537970253930571891217684318286362846948405301614416430468066875699415246993185704183030512549594371372159029236099
        = 509435952285839914555051023580843714132648382024111473186660296521821206469746700620316443478873837606252372049619334517
        * 244624208838318150567813139024002896653802092578931401452041221336558477095178155258218897735030590669041302045908071447

[...]

The sum of the computation time for both records is roughly 4000
core-years, using Intel Xeon Gold 6130 CPUs as a reference (2.1GHz).
A rough breakdown of the time spent in the main computation steps is as
follows.
    RSA-240 sieving:  800 physical core-years
    RSA-240 matrix:   100 physical core-years
    DLP-240 sieving: 2400 physical core-years
    DLP-240 matrix:   700 physical core-years
```

{{< /fig-gen >}}

なお

{{< fig-quote type="markdown" title="795-bit factoring and discrete logarithms" link="https://listserv.nodak.edu/cgi-bin/wa.exe?A2=NMBRTHRY;fd743373.1912&FT=M&P=T&H=&S=" lang="en" >}}
The previous records were RSA-768 (768 bits) in December 2009[^rsa2], and a 768-bit prime discrete logarithm in June 2016[^rsa3].

It is the first time that two records for integer factorization and discrete
logarithm are broken together, moreover with the same hardware and software.

[^rsa2]: https://documents.epfl.ch/users/l/le/lenstra/public/papers/rsa768.txt
[^rsa3]: https://listserv.nodak.edu/cgi-bin/wa.exe?A2=NMBRTHRY;a0c66b63.1606
{{< /fig-quote >}}

だそうな

まぁ，クラウド等を使った安価な分散コンピューティングや実用化されつつある量子コンピュータの台頭により，これから状況は変わっていくだろうけど一応の指標にはなると思う。

ちなみに，セキュリティ強度と鍵長の関係は以下の表の通り（単位は全てビット）。

{{< comparable-security-strengths >}} <!-- 要 MathJax -->

更に各セキュリティ強度の有効期限は以下のとおりだ。

{{< security-strength-time-frames >}} <!-- 要 MathJax -->

まぁ今どき1,024ビット以下の鍵長で運用している馬鹿者はおらんじゃろうけど，2030年以降を見据えるならそろそろ RSA や ElGamal/DSA 等の古い公開鍵暗号について見直しを始めるべきなんだろうね。

## ブックマーク

- [暗号鍵関連の各種変数について]({{< ref "/remark/2017/10/key-parameters.md" >}})

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
