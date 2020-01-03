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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B015643CPE</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
