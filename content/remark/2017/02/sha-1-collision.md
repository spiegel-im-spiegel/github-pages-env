+++
date = "2017-02-24T13:38:36+09:00"
title = "最初の SHA-1 衝突例"
draft = true
tags = ["security", "cryptography", "risk", "hash", "sha-1", "collision"]
description = "もうみんな SHA-1 とはオサラバしてるよね（笑）"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flickr = "spiegel"
  instagram = "spiegel_2007"
  tumblr = "spiegel-im-spiegel"
  github = "spiegel-im-spiegel"
  flattr = "spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  license = "by-sa"
+++

いやぁ，ついにこの日が来たようです。

- [Google Online Security Blog: Announcing the first SHA1 collision](https://security.googleblog.com/2017/02/announcing-first-sha1-collision.html)
- [SHAttered](https://shattered.it/) : SHA-1 の衝突例
- [SHA-1 Collision Found - Schneier on Security](https://www.schneier.com/blog/archives/2017/02/sha-1_collision.html)
- [SHA-1衝突攻撃がついに現実に、Google発表　90日後にコード公開 - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1702/24/news067.html)
- [The end of SHA-1 on the Public Web | Mozilla Security Blog](https://blog.mozilla.org/security/2017/02/23/the-end-of-sha-1-on-the-public-web/)

{{< fig-quote title="SHA-1 Collision Found" link="https://www.schneier.com/blog/archives/2017/02/sha-1_collision.html" lang="en" >}}
<q>This is not a surprise. We've all expected this for over a decade, watching computing power increase. This is why NIST standardized SHA-3 in 2012.</q>
{{< /fig-quote >}}

SHA-1 衝突問題については以下を参照のこと。
NIST などでは2014年以降 SHA-1 を電子署名等に使わないよう勧告している。

- [SHA-1 衝突問題： 廃止の前倒し]({{< relref "remark/2015/problem-of-sha1-collision.md" >}})
- [CRYPTREC | SHA-1の安全性について](http://www.cryptrec.go.jp/topics/cryptrec_20151218_sha1_cryptanalysis.html)

現時点で主要な CA では証明書に SHA-1 は使っていないはずである。
また，主要なブラウザについても SHA-1 を使う証明書に対して警告を発するようになっている。

- [SHA-1 ウェブサーバー証明書は 2017 年２月から警告！ウェブサイト管理者は影響の最終確認を – 日本のセキュリティチーム](https://blogs.technet.microsoft.com/jpsecurity/2016/11/25/sha1countdown/)
- [「Google Chrome」の閲覧画面にエラーが！ ～“https://”のサイトにアクセスできない - やじうまの杜 - 窓の杜](http://forest.watch.impress.co.jp/docs/serial/yajiuma/1041798.html)

もうみんな SHA-1 とはオサラバしてるよね（笑）

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
