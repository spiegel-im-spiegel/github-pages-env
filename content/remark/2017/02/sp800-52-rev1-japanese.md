+++
date = "2017-02-01T23:06:06+09:00"
title = "NIST SP800-52 Rev.1 の邦訳が登場"
tags = ["security", "cryptography", "tls", "nist"]
description = "NIST SP800-52 Rev.1 の邦訳が登場したようだ。"

[scripts]
  mathjax = false
  mermaidjs = false
+++

IPA は [NIST のセキュリティ関連文書の邦訳を積極的に行っている](http://www.ipa.go.jp/security/publications/nist/ "セキュリティ関連NIST文書：IPA 独立行政法人 情報処理推進機構")が，SP800-52 Rev.1 の邦訳が登場したようだ。

- [Guidelines for the Selection, Configuration, and Use of Transport Layer Security (TLS) Implementations | NIST](https://www.nist.gov/node/562891?pub_id=915295)
- {{< pdf-file title="NIST Special Publication 800-52 Revision 1 トランスポート層セキュリティ (TLS) 実装の選択、設定、および使用のためのガイドライン" link="http://www.ipa.go.jp/files/000057084.pdf" >}}

古い話になるが，2013年までに [RC4 の危殆化](https://baldanders.info/blog/000626/ "RC4 終了のお知らせ — Baldanders.info")や SSL/TLS の攻略コードがいくつか「開発」されたことにより TLS 1.2 への移行が強く推奨されることになった。
それを受けての SP800-52 改訂だったのだが，その後の SSL/TLS やその実装である OpenSSL 等のソフトウェアへの攻撃の激しさはみなさんご存じのとおりである。

そうそう。
IPA と言えば最近になってヤバい注意喚起が上がっている。

- [【注意喚起】SQLインジェクションをはじめとしたウェブサイトの脆弱性の再点検と速やかな改修を：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/announce/website_vuln.html)

また2014年に大騒ぎになった Heartbleed 脆弱性をいまだに放置しているサイトもあるようだ。

- [「Heartbleed」脆弱性、多くのサイトやサーバでいまだに存在--Shodan Report - ZDNet Japan](https://japan.zdnet.com/article/35095570/)

攻撃者は既知の攻撃は当然のように試す。
先延ばししていいことは何もない。
いや，マジでお願いしますよ，サイト運用者の方々。

## ブックマーク

- [CRYPTREC Report 2013 — Baldanders.info](https://baldanders.info/blog/000740/)
- [パスワード変更は計画的に — Baldanders.info](https://baldanders.info/blog/000682/)
- [Prohibiting RC4 — Baldanders.info](https://baldanders.info/blog/000810/)

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
