+++
date = "2017-01-31T16:51:01+09:00"
title = "NIST SP800-52 Rev.1 の邦訳が登場"
draft = true
tags = ["security", "cryptography", "tls", "nist"]
description = "NIST SP800-52 Rev.1 の邦訳が登場したようだ。"

[author]
  github = "spiegel-im-spiegel"
  license = "by-sa"
  url = "http://www.baldanders.info/spiegel/profile/"
  flattr = "spiegel"
  name = "Spiegel"
  linkedin = "spiegelimspiegel"
  facebook = "spiegel.im.spiegel"
  flickr = "spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  instagram = "spiegel_2007"
  avatar = "/images/avatar.jpg"
+++

IPA は [NIST のセキュリティ関連文書の邦訳を積極的に行っている](http://www.ipa.go.jp/security/publications/nist/ "セキュリティ関連NIST文書：IPA 独立行政法人 情報処理推進機構")が，SP800-52 Rev.1 の邦訳が登場したようだ。

- [Guidelines for the Selection, Configuration, and Use of Transport Layer Security (TLS) Implementations | NIST](https://www.nist.gov/node/562891?pub_id=915295)
- {{< pdf-file title="NIST Special Publication 800-52 Revision 1 トランスポート層セキュリティ (TLS) 実装の選択、設定、および使用のためのガイドライン" link="http://www.ipa.go.jp/files/000057084.pdf" >}}

古い話になるが，2013年までに [RC4 の危殆化](http://www.baldanders.info/spiegel/log2/000626.shtml "RC4 終了のお知らせ — Baldanders.info")や SSL/TLS の攻略コードがいくつか「開発」されたことにより TLS 1.2 への移行が強く推奨されることになった。
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

- [CRYPTREC Report 2013 — Baldanders.info](http://www.baldanders.info/spiegel/log2/000740.shtml)
- [パスワード変更は計画的に — Baldanders.info](http://www.baldanders.info/spiegel/log2/000682.shtml)
- [Prohibiting RC4 — Baldanders.info](http://www.baldanders.info/spiegel/log2/000810.shtml)

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
