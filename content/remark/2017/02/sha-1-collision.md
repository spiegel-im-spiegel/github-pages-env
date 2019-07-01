+++
date = "2017-02-25T12:38:07+09:00"
update = "2017-10-17T16:44:55+09:00"
title = "最初の SHA-1 衝突例"
draft = false
tags = ["security", "cryptography", "risk", "hash", "sha-1", "collision"]
description = "もうみんな SHA-1 とはオサラバしてるよね（笑）"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flickr = "spiegel"
  instagram = "spiegel_2007"
  tumblr = ""
  github = "spiegel-im-spiegel"
  flattr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"
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
- [Re: [openpgp] SHA1 collision found](https://mailarchive.ietf.org/arch/msg/openpgp/AjJ3BHzd2c9K2KQ3DTk9Ry_QVYM)
    - [[openpgp] V5 Fingerprint again](https://mailarchive.ietf.org/arch/msg/openpgp/_uV_coJ0CYayv_2ptJMuSraJhws)
- [GoogleのSHA-1のはなし](https://www.slideshare.net/herumi/googlesha1) : 分かりやすい解説
- [SHA-1 collision detection on GitHub.com](https://github.com/blog/2338-sha-1-collision-detection-on-github-com)
    - [GitHub Enterprise、SHA-1衝突を実行不能にするパッチを適用へ -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/1050486.html)

{{< fig-quote title="SHA-1 Collision Found" link="https://www.schneier.com/blog/archives/2017/02/sha-1_collision.html" lang="en" >}}
<q>This is not a surprise. We've all expected this for over a decade, watching computing power increase. This is why NIST standardized SHA-3 in 2012.</q>
{{< /fig-quote >}}

SHA-1 衝突問題については以下を参照のこと。
NIST などでは2014年以降 SHA-1 を電子署名等に使わないよう勧告している。

- [SHA-1 衝突問題： 廃止の前倒し]({{< ref "/remark/2015/problem-of-sha1-collision.md" >}})
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

- [Git Commit で OpenPGP 署名を行う]({{< ref "/remark/2016/04/git-commit-with-openpgp-signature.md" >}})

ただし，かつて標準として使われていた MD5 が危殆化とともに廃れていったように，今後 SHA-1 は電子署名以外でも使われなくなると思われる。
念のため， NIST による現在の SHA アルゴリズムの評価と有効期限を以下に示す。

{{< div-gen >}}
<figure lang='en'>
<style>
main table.nist3 th  {
  vertical-align:middle;
  text-align: center;
}
main table.nist3 td  {
  //vertical-align:middle;
  text-align: center;
}
</style>
<table class="nist3">
<thead>
<tr>
<th>Security <br>Strength</th>
<th>Digital Signatures and <br>hash-only applications</th>
<th>HMAC,<br>Key Derivation Functions,<br>Random Number Generation</th>
</tr>
</thead>
<tbody>
<tr>
<td> $\le 8$0</td>
<td>SHA-1</td>
<td>&nbsp;</td>
</tr><tr>
<td>$112$</td>
<td>SHA-224, SHA-512/224, SHA3-224</td>
<td>&nbsp;</td>
</tr><tr>
<td>$128$</td>
<td>SHA-256, SHA-512/256, SHA3-25</td>
<td>SHA-1</td>
</tr><tr>
<td>$192$</td>
<td>SHA-384, SHA3-384</td>
<td>SHA-224, SHA-512/224</td>
</tr><tr>
<td>$\ge 256$</td>
<td>SHA-512, SHA3-512</td>
<td>SHA-256, SHA-512/256,<br> SHA-384,<br> SHA-512, SHA3-512</td>
</tr>
</tbody>
</table>
<figcaption>Hash functions that can be used to provide the targeted security strengths (via <q><a href='https://doi.org/10.6028/NIST.SP.800-57pt1r4'>SP800-57 Part 1 Revision 4 <sup><i class='far fa-file-pdf'></i></sup></a></q>)</figcaption>
</figure>
{{< /div-gen >}}

{{< div-gen >}}
<figure lang='en'>
<style>
main table.nist4 th  {
  vertical-align:middle;
  text-align: center;
}
main table.nist4 td  {
  vertical-align:middle;
  text-align: center;
}
</style>
<table class="nist4">
<thead>
<tr>
<th colspan='2'>Security Strength</th>
<th>Through<br> 2030</th>
<th>2031 and<br> Beyond</th>
</tr>
</thead>
<tbody>
<tr><td rowspan='2'>$\lt 112$</td><td>Applying</td>  <td colspan='2'>Disallowed</td></tr>
<tr>                              <td>Processing</td><td colspan='2'>Legacy-use</td></tr>
<tr><td rowspan='2'>$112$</td>    <td>Applying</td>  <td rowspan='2'>Acceptable</td><td>Disallowed</td></tr>
<tr>                              <td>Processing</td>                               <td>Legacy use</td></tr>

<tr><td>$128$</td>                <td rowspan='3'>Applying/Processing</td><td>Acceptable</td><td>Acceptable</td></tr>
<tr><td>$192$</td>                                   <td>Acceptable</td><td>Acceptable</td></tr>
<tr><td>$256$</td>                                   <td>Acceptable</td><td>Acceptable</td></tr>
</tbody>
</table>
<figcaption>Security-strength time frames (via <q><a href='https://doi.org/10.6028/NIST.SP.800-57pt1r4'>SP800-57 Part 1 Revision 4 <sup><i class='far fa-file-pdf'></i></sup></a></q>)</figcaption>
</figure>
{{< /div-gen >}}

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
