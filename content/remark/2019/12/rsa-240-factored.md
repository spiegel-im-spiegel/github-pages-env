+++
title = "RSA-240 が解けたらしい"
date =  "2019-12-07T14:35:56+09:00"
description = "2030年以降を見据えるならそろそろ RSA や ElGamal/DSA 等の古い公開鍵暗号について見直しを始めるべきなんだろうね。"
image = "/images/attention/kitten.jpg"
tags = [ "cryptography", "rsa" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

“[RSA-240 Factored - Schneier on Security](https://www.schneier.com/blog/archives/2019/12/rsa-240_factore.html)” 経由：

RSA challenge list のうち RSA-240 が解けたようだ。

{{< fig-gen type="markdown" title="795-bit factoring and discrete logarithms" link="https://listserv.nodak.edu/cgi-bin/wa.exe?A2=NMBRTHRY;fd743373.1912&FT=M&P=T&H=&S=" >}}

```
We are pleased to announce the factorization of RSA-240, from RSA's challenge
list, and the computation of a discrete logarithm of the same size (795 bits):

RSA-240 = 124620366781718784065835044608106590434820374651678805754818788883289666801188210855036039570272508747509864768438458621054865537970253930571891217684318286362846948405301614416430468066875699415246993185704183030512549594371372159029236099
        = 509435952285839914555051023580843714132648382024111473186660296521821206469746700620316443478873837606252372049619334517
        * 244624208838318150567813139024002896653802092578931401452041221336558477095178155258218897735030590669041302045908071447

[...]

The sum of the computation time for both records is roughly 4000
core-years, using Intel Xeon Gold 6130 CPUs as a reference (2.1GHz).
A rough breakdown of the time spent in the main computation steps is as
follows.
    RSA-240 sieving:  800 physical core-years
    RSA-240 matrix:   100 physical core-years
    DLP-240 sieving: 2400 physical core-years
    DLP-240 matrix:   700 physical core-years
```

{{< /fig-gen >}}

なお

{{< fig-quote type="markdown" title="795-bit factoring and discrete logarithms" link="https://listserv.nodak.edu/cgi-bin/wa.exe?A2=NMBRTHRY;fd743373.1912&FT=M&P=T&H=&S=" lang="en" >}}
The previous records were RSA-768 (768 bits) in December 2009[^rsa2], and a 768-bit prime discrete logarithm in June 2016[^rsa3].

It is the first time that two records for integer factorization and discrete
logarithm are broken together, moreover with the same hardware and software.

[^rsa2]: https://documents.epfl.ch/users/l/le/lenstra/public/papers/rsa768.txt
[^rsa3]: https://listserv.nodak.edu/cgi-bin/wa.exe?A2=NMBRTHRY;a0c66b63.1606
{{< /fig-quote >}}

だそうな

まぁ，クラウド等を使った安価な分散コンピューティングや実用化されつつある量子コンピュータの台頭により，これから状況は変わっていくだろうけど一応の指標にはなると思う。

ちなみに，セキュリティ強度と鍵長の関係は以下の表の通り（単位は全て bit）。

{{< div-gen >}}
<figure lang="en">
<style>
main table.nist2 th  {
  vertical-align:middle;
  text-align: center;
}
main table.nist2 td  {
  vertical-align:middle;
  text-align: center;
}
</style>
<table class="nist2">
<thead>
<tr>
<th>Security<br>Strength</th>
<th>Symmetric<br> key<br> algorithms</th>
<th>FFC<br>(e.g., DSA, D-H)</th>
<th>IFC<br>(e.g., RSA)</th>
<th>ECC<br>(e.g., ECDSA)</th>
</tr>
</thead>
<tbody>
<tr><td> $\le 80$ </td><td>2TDEA</td><td> $L=1024$ <br> $N=160$ </td><td> $k=1024$ </td> <td> $f = 160\text{ - }223$ </td></tr>
<tr><td> $112$ </td><td>3TDEA</td><td> $L=2048$ <br> $N=224$ </td><td>$k=2048$</td> <td>$f = 224\text{ - }255$</td></tr>
<tr><td> $128$ </td><td>AES-128</td><td> $L=3072$ <br> $N=256$ </td><td>$k=3072$</td> <td>$f = 256\text{ - }383$</td></tr>
<tr><td> $192$ </td><td>AES-192</td><td> $L=7680$ <br> $N=384$ </td><td>$k=7680$</td> <td>$f = 384\text{ - }511$</td></tr>
<tr><td> $256$ </td><td>AES-256</td><td> $L=15360$ <br> $N=512$ </td><td>$k=15360$</td><td>$f=512+$</td></tr>
</tbody>
</table>
<figcaption>Comparable strengths (via <q><a href='https://doi.org/10.6028/NIST.SP.800-57pt1r4'>SP800-57 Part 1 Revision 4 <sup><i class='far fa-file-pdf'></i></sup></a></q>)</figcaption>
</figure>
{{< /div-gen >}}

更に各セキュリティ強度の有効期限は以下のとおりだ。

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

まぁ今どき1,024ビット以下の鍵長で運用している馬鹿者はおらんじゃろうけど，2030年以降を見据えるならそろそろ RSA や ElGamal/DSA 等の古い公開鍵暗号について見直しを始めるべきなんだろうね。

## ブックマーク

- [暗号鍵関連の各種変数について]({{< ref "/remark/2017/10/key-parameters.md" >}})

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/B015643CPE?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/B015643CPE?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">暗号技術入門 第3版　秘密の国のアリス</a></dt>
    <dd>結城 浩 (著)</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>Kindle版</dd>
    <dd>B015643CPE (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4314009071?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4314009071?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">暗号化 プライバシーを救った反乱者たち</a></dt>
    <dd>スティーブン・レビー (著), 斉藤 隆央 (翻訳)</dd>
    <dd>紀伊國屋書店 2002-02-16</dd>
    <dd>単行本</dd>
    <dd>4314009071 (ASIN), 9784314009072 (EAN), 4314009071 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-16">2018-12-16</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
