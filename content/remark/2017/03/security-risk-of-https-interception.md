+++
date = "2017-03-21T20:32:28+09:00"
update = "2017-03-21T22:35:46+09:00"
title = "HTTPS 通信監視機器のリスク"
draft = false
tags = ["security", "risk", "x509", "pki", "ssl", "tls"]
description = "2015年の CERT/CC ブログ記事「The Risks of SSL Inspection」に関する注意喚起が今更ながら出ている。"

[author]
  linkedin = "spiegelimspiegel"
  facebook = "spiegel.im.spiegel"
  github = "spiegel-im-spiegel"
  license = "by-sa"
  flattr = "spiegel"
  avatar = "/images/avatar.jpg"
  url = "http://www.baldanders.info/spiegel/profile/"
  flickr = "spiegel"
  instagram = "spiegel_2007"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
+++

2015年の CERT/CC ブログ記事 “[The Risks of SSL Inspection]” に関する注意喚起が今更ながら出ている。

- [The Risks of SSL Inspection]
    - {{< pdf-file title="The Security Impact of HTTPS Interception" link="https://jhalderm.com/pub/papers/interception-ndss17.pdf" >}}
    - [HTTPS Interception Weakens TLS Security | US-CERT](https://www.us-cert.gov/ncas/alerts/TA17-075A)
    - [JVNTA#96603741: HTTPS 通信監視機器によるセキュリティ強度低下の問題](http://jvn.jp/ta/JVNTA96603741/)

「HTTPS 通信監視機器」というのは，ぶっちゃけていうと， HTTPS 暗号通信[^https] に「中間者攻撃（man-in-the-middle attack）」を仕掛けて通信を傍受し malware 等を検出・排除する「セキュリティ製品」である。

[^https]: 念のため簡単に説明しておくと， HTTPS (Hypertext Transfer Protocol Secure) 暗号通信は WWW (World Wide Web) におけるクライアント-サーバ間の通信経路を暗号化する仕組みである。具体的には TLS (Transport Layer Security) 等のプロトコルを用い公開鍵暗号方式を使ってセッション鍵を生成する。また公開鍵暗号方式の公開鍵は X.509 方式の公開鍵基盤（Public Key Infrastructure; PKI）によって管理される。

HTTPS 通信監視機器のいくつかにはセキュリティ上の問題が存在する。
“[The Risks of SSL Inspection]” から抜き出してみよう。

1. Incomplete validation of upstream certificate validity
2. Not conveying validation of upstream certificate to the client
3. Overloading of certificate Canonical Name (CN) field
4. Use of the application layer to convey certificate validity
5. Use of a User-Agent HTTP header to determine when to validate a certificate
6. Communication before warning
7. Same root CA certificate

これらの問題があると推測される製品のリストが “[The Risks of SSL Inspection]” に挙がっているので該当者は確認してみるとよいだろう。
また以下のサイトからも確認できる。

- [badssl.com](https://badssl.com/)

“[The Risks of SSL Inspection]” では以下のように結論付けている。

{{< fig-quote title="The Risks of SSL Inspection" link="http://insights.sei.cmu.edu/cert/2015/03/the-risks-of-ssl-inspection.html" lang="en" >}}
<q>SSL and TLS do not provide the level of end-to-end security that users may expect. Even in absence of SSL inspection, there are problems with how well browsers are conveying SSL information to users. The fact that "SSL inspection" is a phrase that exists, should be a blazing red flag that what you think SSL is doing for you is fundamentally broken.</q>
{{< /fig-quote >}}

[以前も書いた](http://www.baldanders.info/spiegel/log2/000812.shtml "HTTPS Deep Inspection — Baldanders.info")が，HTTPS 通信監視機器（あるいは HTTPS Deep Inspection）の存在自体がインターネットの “End to End” 原則を崩すものであり，ひいては「ネットの中立性」に楔を入れるものである。
しかし「[馬も鹿も暗号化する時代]({{< relref "remark/2016/03/vulnerability-cross-protocol-attack-on-tls-using-sslv2.md" >}} "SSLv2 を有効にしている TLS 実装の脆弱性 ― 馬も鹿も暗号化する時代のセキュリティ")」にこの原則は風前の灯である。
たとえば [CMS の面倒すらろくすっぽ見られない]({{< relref "remark/2016/07/cms.md">}} "「自分で面倒見られる子」だけが CMS を導入しなさい")ユーザが「うちも [Let's la Encrypt]」とか言い出して脆弱性だらけのサイトを暗号化したらどうなるのか。

ネットワーク・セキュリティ専門家から企業あるいは私たち個人に至るまで，場当たりな対処に満足するのではなく，この「現実」にきちんと向き合うべきだと思うのだが，どうだろう。

## ブックマーク

- [Malware Spoofing HTTPS（3月2日，追記あり） — Baldanders.info](http://www.baldanders.info/spiegel/log2/000809.shtml)
- [HTTPS Deep Inspection — Baldanders.info](http://www.baldanders.info/spiegel/log2/000812.shtml)
- [HTTPS監視装置にセキュリティ低下の危険性--日米機関で注意喚起 - ZDNet Japan](https://japan.zdnet.com/article/35098402/)

[The Risks of SSL Inspection]: http://insights.sei.cmu.edu/cert/2015/03/the-risks-of-ssl-inspection.html
[Let's la Encrypt]: https://letsencrypt.org/ "Let's Encrypt - Free SSL/TLS Certificates"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
