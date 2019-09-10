+++
title = "gpgpdump v0.6.2 をリリースした"
date =  "2019-09-09T19:38:32+09:00"
description = "今回は細かい調整のみ。 draft-ietf-openpgp-rfc4880bis-08 への追随も行った。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "golang", "gpgpdump"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.6.2 をリリースした。

- [Release v0.6.2 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.6.2)

といっても今回は細かい調整のみ。
[Go] 1.13 に対応したのと新しいエラーハンドリング・パッケージ [spiegel-im-spiegel/errs] を組み込んだだけ。
でも，これでデバッグが格段にし易くなった。

あ，そうそう。
[draft-ietf-openpgp-rfc4880bis-08](https://tools.ietf.org/html/draft-ietf-openpgp-rfc4880bis-08 "draft-ietf-openpgp-rfc4880bis-08 - OpenPGP Message Format") への追随も行った。
でも試す対象がないからなぁ（笑）

[gpgpdump] の使い方等は以下を参照のこと。

- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

## ブックマーク

- [Go 言語用エラーハンドリング・パッケージをリリースした]({{< ref "/release/2019/09/errs-package-is-released.md" >}})
- [spiegel-im-spiegel/gocli v0.10.1 のリリース]({{< ref "/release/2019/09/gocli-v0_10_1-is-released.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Go]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/errs]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071"><img src="https://images-fe.ssl-images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071">暗号化 プライバシーを救った反乱者たち</a></dt>
    <dd>スティーブン・レビー</dd>
    <dd>斉藤 隆央 (翻訳)</dd>
    <dd>紀伊國屋書店 2002-02-16</dd>
    <dd>単行本</dd>
    <dd>4314009071 (ASIN), 9784314009072 (EAN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-16">2018-12-16</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>Kindle版</dd>
    <dd>B015643CPE (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
