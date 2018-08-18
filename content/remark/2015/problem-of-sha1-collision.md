+++
date = "2015-10-22T14:35:22+09:00"
update = "2017-10-28T21:26:57+09:00"
description = "GPU をふんだんに使った専用ハードウェアやクラウド・サービスなどを組み合わせることにより，近い将来に実用的なコストで SHA-1 攻略が可能になると指摘されている。"
draft = false
tags = ["security", "cryptography", "hash", "sha-1", "collision", "risk"]
title = "SHA-1 衝突問題： 廃止の前倒し"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

SHA-1 の廃止プロセスが前倒しになるかもしれない。

## おさらい： SHA-1 衝突問題

そもそもの発端は，2004年に複数の hash 関数において高い確率で hash 値を衝突（collision）させる攻略法が発表されたことだった。

- [Collisions for Hash Functions MD4, MD5, HAVAL-128 and RIPEMD](http://eprint.iacr.org/2004/199) : 発端となった論文。この中で SHA-0 も攻略可能であることが示されている

その後の研究で SHA-1 も攻略可能であることが分かってきて暗号技術の周辺は大騒動になった。

- [SHA-1 Broken - Schneier on Security](https://www.schneier.com/blog/archives/2005/02/sha1_broken.html)
- [デファクトスタンダード暗号技術の大移行（1）：すべてはここから始まった～SHA-1の脆弱化 (1/2) - ＠IT](http://www.atmarkit.co.jp/ait/articles/0603/09/news117.html)

### Hash 値の衝突問題

暗号技術における hash 関数とは，以下の機能を持ったアルゴリズムである

1. 任意のデータ列を一定の長さのデータ列（hash 値）に「要約」する
1. Hash 値から元のデータ列を推測できない
1. ひとつの hash 値に対して複数のデータ列が（実時間で）見つからない

Hash 関数はメッセージ認証符号（Message Authentication Code; MAC）や電子署名（digital signature）の中核技術のひとつであり，データの「完全性（Integrity）」を担保する重要な要素である。
特に3番目の「ひとつの hash 値に対して複数のデータ列が（実時間で）見つからない」という機能が破られると，その hash 関数では完全性を担保できなくなってしまう。
これを「Hash 値の衝突問題」という。

よくできた hash 関数であれば， hash 値のビット長を $n$ とすると，衝突の確率は $1 / 2^{\frac{n}{2}}$ であると言われている。
SHA-1 の hash 値の長さは $160\,\mathrm{bits}$ なので，衝突の確率は $1 / 2^{80}$ になるはずだが，実際にはそれよりもずっと大きい $1 / 2^{69}$ で可能，というのが当時の問題だった。

### 2010年問題

これを受けて， SHA-1 を2010年までに廃止し SHA-2 (SHA-224/256/364/512) に移行する措置がとられた。
これが暗号技術における「2010年問題」である[^a]。

[^a]: 他にも $1024\,\mathrm{bits}$ 以下の鍵長の RSA 公開鍵を廃止する，などの措置が盛り込まれていた。

ただ，現状では（特に legacy system において）アルゴリズムの置き換えがなかなか進まなかったことと SHA-1 の攻略があまり進展しなかったことにより，この期限は2013年まで延長された。
現在の SHA アルゴリズムの評価と有効期限は以下のとおり。

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

しかし現状は全くスケジュールどおりではなく， SHA-1 を使った証明書の発行が停止され始めたのはようやく昨年末頃からだ。
一方，主要ブラウザは2017年以降 SHA-1 を使った証明書を無効にする計画を発表している。

## SHA-1 Freestart Collision

ところが最近になって SHA-1 の攻略について進展があった。

- [The Shappening: freestart collisions for SHA-1](https://sites.google.com/site/itstheshappening/)
- [SHA-1 Freestart Collision - Schneier on Security](https://www.schneier.com/blog/archives/2015/10/sha-1_freestart.html)
- [SHA1 algorithm securing e-commerce and software could break by year’s end | Ars Technica](http://arstechnica.com/security/2015/10/sha1-crypto-algorithm-securing-internet-could-break-by-years-end/)
- [「SHA-1の廃止前倒しを」　専門家チームが提言 - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1510/09/news054.html)

これは SHA-1 のアルゴリズム上の更なる危殆化を指すものではないが，最新の計算リソースを効率的に使って，いわば「力任せ」（って言うと御幣があるけど）で攻略している。
GPU をふんだんに使った専用ハードウェアやクラウド・サービスなどを組み合わせることにより，近い将来に実用的なコストで SHA-1 攻略が可能になると指摘されている。

{{< fig-quote lang="en" title="The Shappening: freestart collisions for SHA-1" link="https://sites.google.com/site/itstheshappening/" >}}
<q>Concretely, we estimate the SHA-1 collision cost today (i.e., Fall 2015) between 75K\$ and 120K\$ renting Amazon EC2 cloud computing over a few months. By contrast, security expert Bruce Schneier previously projected (based on calculations from Jesse Walker) the SHA-1 collision cost to be ~173K\$ by 2018. Note that he deems this to be within the resources of a criminal syndicate. Large corporations and governments may possess even greater resources and may not require Amazon EC2.</q>
{{< /fig-quote >}}

Mozilla では SHA-1 を使った証明書の無効化を2016年7月以降に前倒しすることを検討している。

- [Continuing to Phase Out SHA-1 Certificates | Mozilla Security Blog](https://blog.mozilla.org/security/2015/10/20/continuing-to-phase-out-sha-1-certificates/)
- [Firefox、SHA-1対策を前倒し検討 | マイナビニュース](http://news.mynavi.jp/news/2015/10/22/093/)

GPU ベースのシステムやクラウド・サービス等による巨大計算リソースを使った攻撃は，指摘はありつつも，あまり重視されていなかった。
しかし，今回の例に見るように，攻撃手法として現実的な脅威になりつつある。
状況によっては SHA-1 以外でもセキュリティ・リスクの見直しを迫られるかもしれない。

## 関連記事

- [ハッシュ値の衝突問題 -- 戯れ言++](http://www.baldanders.info/spiegel/remark/archives/000048.shtml)
- [暗号の危殆化と新しいアルゴリズム -- 戯れ言++](http://www.baldanders.info/spiegel/remark/archives/000204.shtml)
- [「安全な鍵長の下限」とは -- 戯れ言++](http://www.baldanders.info/spiegel/remark/archives/000210.shtml)
- ["NIST's Plan for New Cryptographic Hash Functions" — Baldanders.info](http://www.baldanders.info/spiegel/log2/000267.shtml)
- [『暗号をめぐる最近の話題』 — Baldanders.info](http://www.baldanders.info/spiegel/log2/000586.shtml)
- [SHA-3 が正式リリース： あれから10年も... — Baldanders.info](http://www.baldanders.info/spiegel/log2/000865.shtml)
- [自堕落な技術者の日記 : 「RFC 7525 TLSとDTLSの安全な利用に関する推奨事項」の公開 - livedoor Blog（ブログ）](http://blog.livedoor.jp/k_urushima/archives/1768181.html)
- [「Y!mobileケータイ」で一部サイトへ接続不能に--サーバ証明書の切り替えで - CNET Japan](http://japan.cnet.com/news/service/35067422/)
- [scryptがGPUに破られる時 | びりあるの研究ノート](https://blog.visvirial.com/articles/519) : GPU 耐性が高いと言われる scrypt も実時間で攻略できるようになってきたという話
- [自堕落な技術者の日記 : SSL Pulseの統計情報で見るSSL/TLS (2015年10月版) - livedoor Blog（ブログ）](http://blog.livedoor.jp/k_urushima/archives/1782546.html) : Alexa 社によるアクセス世界トップ20万サイトを対象にした SSL/TLS 関連の統計情報
- [FAQ: SHA-1 廃止/SHA-2 移行に関するマイクロソフトのポリシー - 日本のセキュリティチーム](http://blogs.technet.com/b/jpsecurity/archive/2015/11/02/faq-microsoft-policy-on-sha1-deprecation.aspx)
- [「SHA-1」SSLサーバ証明書の廃止迫る--「SHA-2」への移行状況とその影響、課題 - ZDNet Japan](http://japan.zdnet.com/article/35072827/)
- [CRYPTREC | SHA-1の安全性について](http://www.cryptrec.go.jp/topics/cryptrec_20151218_sha1_cryptanalysis.html)
    - {{< pdf-file title="CRYPTREC暗号技術ガイドライン(SHA-1)" link="http://www.cryptrec.go.jp/report/c13_tech_guideline_SHA-1_web.pdf" >}}
- [[IT 管理者向け] 残っていませんか? SSL/TLS 証明書の SHA-1 廃止はもうすぐ - 日本のセキュリティチーム - Site Home - TechNet Blogs](http://blogs.technet.com/b/jpsecurity/archive/2016/02/08/sha1-deprecation-tls.aspx)
- [SHA-1 ウェブサーバー証明書は 2017 年２月から警告！ウェブサイト管理者は影響の最終確認を – 日本のセキュリティチーム](https://blogs.technet.microsoft.com/jpsecurity/2016/11/25/sha1countdown/)
- [「Google Chrome」の閲覧画面にエラーが！ ～“https://”のサイトにアクセスできない - やじうまの杜 - 窓の杜](http://forest.watch.impress.co.jp/docs/serial/yajiuma/1041798.html)
- [最初の SHA-1 衝突例]({{< ref "/remark/2017/02/sha-1-collision.md" >}})

## 関連図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">第3版出た！ てか，もう Kindle 版出てるのか。紙の本買うのはやまったかなぁ。 SHA-3 や BitCoin/BlockChain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
