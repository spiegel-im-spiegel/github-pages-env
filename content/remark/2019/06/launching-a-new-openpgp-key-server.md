+++
title = "新しい OpenPGP 鍵サーバが Launch したらしい"
date =  "2019-06-13T22:19:09+09:00"
description = "これで思い出すのが，かつての OpenPKSD だが，今回の keys.openpgp.org がその二の舞にならないことを祈るばかりである。"
image = "/images/attention/openpgp.png"
tags = [
  "cryptography",
  "openpgp",
  "management",
  "pki",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP のメーリングリスト](https://www.ietf.org/mailman/listinfo/openpgp)より。

- [Launching a new keyserver on keys.openpgp.org!](https://mailarchive.ietf.org/arch/msg/openpgp/1cQeIV8s81lhwG_FQtMuc2JbRSk)
- [Launching a new keyserver! - keys.openpgp.org](https://keys.openpgp.org/about/news#2019-06-12-launch)
- [sequoia-pgp / hagrid · GitLab](https://gitlab.com/sequoia-pgp/hagrid)

{{< fig-quote type="markdown" title="Launching a new keyserver!" link="https://keys.openpgp.org/about/news#2019-06-12-launch" lang="en" >}}
* Fast and reliable. No wait times, no downtimes, no inconsistencies.
* Precise. Searches return only a single key, which allows for easy key discovery.
* Validating. Identities are only published with consent, while non-identity information is freely distributed.
* Deletable. Users can delete personal information with a simple e-mail confirmation.
* Built on Rust, powered by [Sequoia PGP](https://sequoia-pgp.org/) - free and open source, running AGPLv3.
{{< /fig-quote >}}

ほほう。
Rust ベースなのか。
面白いな。

現行の [OpenPGP] 鍵サーバ同士は peer-to-peer で同期しているが，新しい [keys.openpgp.org] はこれらとは別のネットワークを形成するようだ。

{{< fig-quote title="Launching a new keyserver!" link="https://keys.openpgp.org/about/news#2019-06-12-launch" lang="en" >}}
<q>We created keys.openpgp.org to provide an alternative to the SKS Keyserver pool, which is the default in many applications today. This distributed network of keyservers has been struggling with <a href="https://medium.com/@mdrahony/are-sks-keyservers-safe-do-we-need-them-7056b495101c">abuse</a>, <a href="https://en.wikipedia.org/wiki/Key_server_(cryptographic)#Problems_with_keyservers">performance</a>, as well as <a href="http://www.openwall.com/lists/oss-security/2017/12/10/1">privacy issues</a>, and more recently also <a href="http://nongnu.13855.n7.nabble.com/SKS-apocalypse-mitigation-td228252.html">GDPR</a> compliance questions. Kristian Fiskerstrand has done a stellar job maintaining the pool for <a href="https://blog.sumptuouscapital.com/2016/12/10-year-anniversary-for-sks-keyservers-net/">more than ten years</a>, but at this point development activity seems to have <a href="https://bitbucket.org/skskeyserver/sks-keyserver/pull-requests/60/clean-build-with-405">mostly ceased</a>.<br>
We thought it time to consider a fresh approach to solve these problems.</q>
{{< /fig-quote >}}

さらに

{{< fig-quote title="Launching a new keyserver!" link="https://keys.openpgp.org/about/news#2019-06-12-launch" lang="en" >}}
<q>The keys.openpgp.org keysever will receive first-party support in upcoming releases of <a href="https://enigmail.net/">Enigmail</a> for Thunderbird, as well as <a href="https://play.google.com/store/apps/details?id=org.sufficientlysecure.keychain&hl=en">OpenKeychain</a> on Android. This means users of those implementations will benefit from the faster response times, and improved key discovery by e-mail address. We hope that this will also give us some momentum to build this project into a bigger community effort. </q>
{{< /fig-quote >}}

ということで [Enigmail] や [OpenKeychain] と連携するのであれば期待できそうな感じである。
ただし，いまのところ [GnuPG] との相性がイマイチのようなので，もう少し様子を見たいところではある。

これで思い出すのが，かつての [OpenPKSD] だが，日本主導でフル Ruby で組まれていて割と期待してたんだが，世界的にはあまり注目されないまま閉鎖されたんだよなぁ。
今回の [keys.openpgp.org] がそうならないことを祈るばかりである。

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[OpenPGP.org]: https://www.openpgp.org/
[keys.openpgp.org]: https://keys.openpgp.org
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Enigmail]: https://enigmail.net/
[OpenKeychain]: https://play.google.com/store/apps/details?id=org.sufficientlysecure.keychain "OpenKeychain: Easy PGP - Google Play"
[OpenPKSD]: https://www.ipa.go.jp/security/fy16/development/openPKSD/ "信頼できるOpenPGP公開鍵を提供する公開鍵サーバOpenPKSD Trusted Keyserver：IPA 独立行政法人 情報処理推進機構"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
