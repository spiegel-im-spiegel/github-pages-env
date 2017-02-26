+++
date = "2017-02-25T12:38:07+09:00"
update = "2017-02-26T14:40:38+09:00"
title = "最初の SHA-1 衝突例"
draft = false
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
- [グーグル、SHA-1衝突攻撃に成功--同一ハッシュ値の2つのPDFも公開 - ZDNet Japan](https://japan.zdnet.com/article/35097102/)
- [Re: [openpgp] SHA1 collision found](https://mailarchive.ietf.org/arch/search/?email_list=openpgp&index=AjJ3BHzd2c9K2KQ3DTk9Ry_QVYM)
- [GoogleのSHA-1のはなし](https://www.slideshare.net/herumi/googlesha1) : 分かりやすい解説

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

## 追記というか補足

たとえば git の commit hash 値は SHA-1 で付与されるが大丈夫なのか？ とかいった意見が散見されるが，当面は問題ない。

今回の件はあくまでも電子署名や hash 値そのものを何かの証明に使おうとする場合に問題となる。
git の commit hash 値はあくまで identity として付与されるものである。
改ざんされたかどうかは commit hash 値ではなく差分情報によって容易に知ることができる。

git による悪意のなりすまし等を警戒する必要があるのなら commit hash 値を気にするのではなく commit にきちんと電子署名を行うことをお勧めする（チームで作業する人は是非習慣化するべきである）。

- [Git Commit で OpenPGP 署名を行う]({{< relref "remark/2016/04/git-commit-with-openpgp-signature.md" >}})

ただし，かつて標準として使われていた MD5 が危殆化とともに廃れていったように，今後 SHA-1 は電子署名以外でも使われなくなると思われる。
念のため， NIST による現在の SHA アルゴリズムの評価と有効期限を以下に示す。

<figure lang='en'>
<table>
<thead>
<tr>
<th style="vertical-align:middle;">Security <br>Strength</th>
<th style="vertical-align:middle;">Digital Signatures and <br>hash-only applications</th>
<th style="vertical-align:middle;">HMAC,<br>Key Derivation Functions,<br>Random Number Generation</th>
</tr>
</thead>
<tbody>
<tr>
<td class='right'> ≦ 80</td>
<td>SHA-1</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>112</td>
<td>SHA-224, SHA-512/224, SHA3-224</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>128</td>
<td>SHA-256, SHA-512/256, SHA3-25</td>
<td>SHA-1</td>
</tr><tr>
<td class='right'>192</td>
<td>SHA-384, SHA3-384</td>
<td>SHA-224, SHA-512/224</td>
</tr><tr>
<td class='right'>≧ 256</td>
<td>SHA-512, SHA3-512</td>
<td>SHA-256, SHA-512/256,<br> SHA-384,<br> SHA-512, SHA3-512</td>
</tr>
</tbody>
</table>
<figcaption>Hash functions that can be used to provide the targeted security strengths (via <q><a href='http://dx.doi.org/10.6028/NIST.SP.800-57pt1r4'>SP800-57 Part 1 Revision 4 <sup><i class='fa fa-file-pdf-o'></i></sup></a></q>)</figcaption>
</figure>

<figure lang='en'>
<table>
<thead>
<tr>
<th style="vertical-align:middle;" colspan='2'>Security Strength</th>
<th style="vertical-align:middle;">through<br> 2030</th>
<th style="vertical-align:middle;">2031 and<br> Beyond</th>
</tr>
</thead>
<tbody>
<tr>
<td class='right' rowspan='2'>＜ 112</td>
<td>Applying</td>
<td colspan='2' class='center'>Disallowed</td>
</tr><tr>
<!-- -->
<td>Processing</td>
<td colspan='2' class='center'>Legacy use</td>
</tr><tr>
<td class='right' rowspan='2'>112</td>
<td>Applying</td>
<td rowspan='2' style="vertical-align:middle;">Acceptable</td>
<td>Disallowed</td>
</tr><tr>
<td>Processing</td>
<!-- -->
<td>Legacy use</td>
</tr><tr>
<td class='right'>128</td>
<td rowspan='3' style="vertical-align:middle;">Applying/Processing</td>
<td>Acceptable</td>
<td>Acceptable</td>
</tr><tr>
<td class='right'>192</td>
<!-- -->
<td>Acceptable</td>
<td>Acceptable</td>
</tr><tr>
<td class='right'>256</td>
<!-- -->
<td>Acceptable</td>
<td>Acceptable</td>
</tr>
</tbody>
</table>
<figcaption>Security-strength time frames (via <q><a href='http://dx.doi.org/10.6028/NIST.SP.800-57pt1r4'>SP800-57 Part 1 Revision 4 <sup><i class='fa fa-file-pdf-o'></i></sup></a></q>)</figcaption>
</figure>

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
