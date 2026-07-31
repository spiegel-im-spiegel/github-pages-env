+++
title = "Thunderbird 78 で署名・暗号化を行うのは少し待ったほうがいいらしい"
date =  "2020-07-05T12:41:15+09:00"
description = "ここは大人しく 78.2 が出るまで待つか。 それともチャレンジするか（笑）"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "security", "cryptography", "mua", "thunderbird" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 Firefox 78 がリリースされ ESR 版も 78 ベースにアップグレードされた。

- [Firefox  78.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/78.0/releasenotes/)

これに伴い [Thunderbird] も今後は 78 ベースになる予定だが，このアップグレードを前に [Enigmail] から以下の注意喚起が出た。

{{< fig-img src="./caution-by-enigmail.png" link="./caution-by-enigmail.png" width="746" >}}

実は [OpenPGP] 暗号化機能を提供する拡張機能 [Enigmail] が [Thunderbird] 78 から標準機能として組み込まれることが（昨年時点で）決まっていた。

- [Thunderbird, Enigmail and OpenPGP | The Mozilla Thunderbird Blog](https://blog.mozilla.org/thunderbird/2019/10/thunderbird-enigmail-and-openpgp/)

のだが，どうやら [Thunderbird] 78.0 リリース時点では安定した機能を提供というわけにはいかないようだ。
ここは大人しく 78.2 が出るまで待つか。
それともチャレンジするか（笑）

## ブックマーク

- [Thunderbird でメール暗号化]({{< ref "/openpgp/encrypted-mail-with-thunderbird.md" >}})

[OpenPGP]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[Thunderbird]: https://www.thunderbird.net/ "Thunderbird — Software made to make email easier. — Mozilla"
[Enigmail]: https://addons.mozilla.org/thunderbird/addon/enigmail/ "Enigmail :: Add-ons for Thunderbird"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
