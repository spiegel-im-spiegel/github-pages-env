+++
title = "暗号鍵関連の各種変数について"
date =  "2017-10-17T16:32:03+09:00"
update = "2018-12-16T14:07:16+09:00"
description = "この記事は将来の記事で再利用するための snippet 置き場として使うことにする。"
tags = ["security", "cryptography", "hash", "risk", "management"]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = true
  mermaidjs = false
+++

「[ECDSA鍵をGitHubで使う方法](https://qiita.com/darai0512/items/c7b47d1b3fe06c4dea7d)」で SP 800-56 が Rev.4 になってるのを見て一瞬「ふぁ！」となったが，よく考えたら今年始めに書いた「[最初の SHA-1 衝突例]({{< ref "/remark/2017/02/sha-1-collision.md" >}})」では既に Rev.4 を参照していたのだった。
でも古い記事を見返したら結構 Rev.3 のままになってるので，このさい全部アップデートすることにした。

この記事は将来の記事で再利用するための snippet 置き場として使うことにする。
表のレイアウトの関係で携帯端末で見ている人は見づらいかもしれないけど，そこはご容赦。

## 参照資料

ここで参照する資料は，米国 [NIST] の Special Publication (SP) 800-57 Part 1 で正式タイトルは「Recommendation for Key Management Part 1: General （鍵管理における推奨事項 第一部：一般事項）」となっている（日本語訳は IPA によるもの）。

- [ SP 800-57 Part 1 Rev. 4 Recommendation for Key Management, Part 1: General](https://csrc.nist.gov/publications/detail/sp/800-57-part-1/rev-4/final)
    - {{< pdf-file title="NIST Special Publication 800-57 Part 1; Revision 4; Recommendation for Key Management Part 1: General" link="https://doi.org/10.6028/NIST.SP.800-57pt1r4" >}} （{{< pdf-file title="IPA による日本語版" link="https://www.ipa.go.jp/files/000055490.pdf" >}}）

## セキュリティ強度と鍵長の関係

最初はセキュリティ強度と鍵長の関係を示す表。
単位は全て bit である。

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
<tr><td>$\le 80$</td><td>2TDEA</td><td>$L=1024$<br>$N=160$</td><td>$k=1024$</td> <td>$f = 160\text{ - }223$</td></tr>
<tr><td>$112$</td><td>3TDEA</td><td>$L=2048$<br>$N=224$</td> <td>$k=2048$</td> <td>$f = 224\text{ - }255$</td></tr>
<tr><td>$128$</td><td>AES-128</td><td>$L=3072$<br>$N=256$</td> <td>$k=3072$</td> <td>$f = 256\text{ - }383$</td></tr>
<tr><td>$192$</td><td>AES-192</td><td>$L=7680$<br>$N=384$</td> <td>$k=7680$</td> <td>$f = 384\text{ - }511$</td></tr>
<tr><td>$256$</td><td>AES-256</td><td>$L=15360$<br>$N=512$</td><td>$k=15360$</td><td>$f=512+$</td></tr>
</tbody>
</table>
<figcaption>Comparable strengths (via <q><a href='https://doi.org/10.6028/NIST.SP.800-57pt1r4'>SP800-57 Part 1 Revision 4 <sup><i class='far fa-file-pdf'></i></sup></a></q>)</figcaption>
</figure>
{{< /div-gen >}}

Symmetric key algorithms は共通鍵暗号アルゴリズム全般を指す。
たとえば AES とか。
IFC (Integer Factorization Cryptosystems) は素因数分解問題ベースの公開鍵暗号アルゴリズムで RSA がこれに該当する。
FFC (Finite Field Cryptosystems) は離散対数問題ベースの公開鍵暗号アルゴリズムで Diffie-Hellman や ElGamal, DSA などがこれに該当する。
ECC (Elliptic Curve Cryptosystems) は離散対数問題でも特に楕円曲線上の離散対数問題ベースの公開鍵暗号アルゴリズムを指す。
たとえば ECDH や ECDSA など。

IFC では $k$，FFC では $L$，ECC では $f$ が鍵長を示す。
たとえばセキュリティ強度が 128bit なら

- AES 128bit
- ElGamal, DSA 3072bit
- RSA 3072bit
- ECDH, ECDSA 256bit

の組み合わせで「[ベストマッチ](https://dic.pixiv.net/a/%E3%83%93%E3%83%AB%E3%83%89%E3%83%89%E3%83%A9%E3%82%A4%E3%83%90%E3%83%BC) キター！」となる。

## セキュリティ強度と Hash 関数の関係

次はセキュリティ強度とHash 関数の関係を示す表。

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
<td>SHA-256, SHA-512/256, SHA3-256</td>
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

考え方は先程の暗号鍵長のときと同じ。
ただし Hash 関数の場合は使用目的ごとに要求されるアルゴリズムが異なるので注意が必要である。

## セキュリティ強度と有効期限

こちらはセキュリティ強度の有効期限を表したものだ。

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

各用語はそれぞれ

| 用語       | 意味                     |
|:---------- |:------------------------ |
| Applying   | 適用                     |
| Processing | 処理                     |
| Acceptable | 許容                     |
| Legacy-use | 許容（レガシー使用のみ） |
| Disallowed | 禁止                     |

という意味だ。
例を挙げると，セキュリティ強度 112bit の暗号スイート（Cipher Suites）を適用する場合は2030年までは許容するけど2031年以降は禁止。
すでに暗号化されているデータを復号したい場合でも，2031年以降はレガシー・システムしか許容しない，ということになる。

たとえば ssh 認証は「適用」なので多くの人が使ってる RSA 2048bit の鍵は2031年以降は使用禁止となるわけだ。
まぁ，そんな先まで同じシステムで同じ鍵を使い続けるかどうかは分からないが（なので今使ってる鍵を慌てて新調する必要はない。新規に作成するなら 128bit 強度の鍵をお勧めするが）。

なおこれは各アルゴリズムに危殆化要因となる脆弱性等がない場合の話である。
したがって暗号製品を使うシステムの管理者やセキュリティ管理者は常に暗号関係のトピックに耳を澄ませておくべきであろう。

## OpenPGP で利用可能なアルゴリズム

この項は「[OpenPGP で利用可能なアルゴリズム]({{< ref "/openpgp/algorithms-for-openpgp.md" >}})」に移動した。

## パスワードの強度

これは [NIST] ではなく IPA の資料だが，文字種と文字数の組み合わせによるパスワードの強度についても上げておこう。
出典は以下。

- [情報漏えいを防ぐためのモバイルデバイス等設定マニュアル：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ipg/documents/dev_setting_crypt.html)
    - {{< pdf-file title="情報漏えいを防ぐためのモバイルデバイス等設定マニュアル 解説編" link="https://www.ipa.go.jp/files/000026760.pdf" >}}

{{< div-gen >}}
<figure>
<table>
<thead>
<tr>
<th colspan='4'>利用する文字種類数と内訳</th>
<th colspan='4'>パスワード長</th>
</tr>
<tr>
<th>種類数</th>
<th>数字</th>
<th>文字</th>
<th>シンボル</th>
<th>4文字</th>
<th>8文字</th>
<th>12文字</th>
<th>16文字</th>
</tr>
</thead>
<tbody>
<tr><td>10種</td><td>0-9</td><td>なし</td>      <td>なし</td><td>1円未満（$2^{13.3}$）</td><td>1円未満（$2^{26.6}$）</td>  <td>約35円（$2^{39.9}$）</td>     <td>約35万円（$2^{53.2}$）</td></tr>
<tr><td>36種</td><td>0-9</td><td>a-z</td>       <td>なし</td><td>1円未満（$2^{20.7}$）</td><td>約100円（$2^{41.4}$）</td>  <td>約1.65億円（$2^{62.0}$）</td> <td>約276兆円（$2^{82.7}$）</td></tr>
<tr><td>62種</td><td>0-9</td><td>a-z<br>A-Z</td><td>なし</td><td>1円未満（$2^{23.8}$）</td><td>約7,500円（$2^{47.6}$）</td><td>約1,120億円（$2^{71.5}$）</td><td>約165京円（$2^{95.3}$）</td></tr>
<tr><td>94種</td><td>0-9</td><td>a-z<br>A-Z</td><td><code style='font-size:smaller;'>! " # $ %<br>&amp; ' ( ) =<br>~ | - ^ `<br>¥ { @ [<br>+ * ] ; :<br>} &lt; &gt; ? _<br>, . /</code></td>
                                                             <td>1円未満（$2^{26.2}$）</td><td>約21万円（$2^{52.4}$）</td> <td>約16.5兆円（$2^{78.7}$）</td> <td>約129,000京円（$2^{104.9}$）</td></tr>
</tbody>
</table>
<figcaption>パスワード解読の想定コスト例（<q><a href='https://www.ipa.go.jp/files/000026760.pdf'>情報漏えいを防ぐためのモバイルデバイス等設定マニュアル 解説編 <sup><i class='far fa-file-pdf'></i></sup></a></q> 2.4.2.2項より）</figcaption>
</figure>
{{< /div-gen >}}

測定基準は以下の通り。

{{< div-gen >}}
<figure>
<blockquote>
<q>利用できる文字種類すべてを完全にランダムに選択して作ったパスワードを一つ一つ調べる全数探索により1日で解読しようとした際にかかるおおまかな想定攻撃コストを示しています。ここでは、全数探索(暗号鍵の総数256)でDES10を1日で解読するためのコストを約250万円と仮定します。また、パスワードを1つ検査するのとDESの暗号鍵を1つ検査するコストは同じであるとし、パスワードを求めるのに必要な計算量(検査する個数)が半分になればコストも半分、2倍になればコストも2倍になるものとしています。</q>
</blockquote>
<figcaption><q><a href='https://www.ipa.go.jp/files/000026760.pdf'>情報漏えいを防ぐためのモバイルデバイス等設定マニュアル 解説編 <sup><i class='far fa-file-pdf'></i></sup></a></q> 2.4.2.2項より</figcaption>
</figure>
{{< /div-gen >}}


## ブックマーク

- [CRYPTREC Report 2013 — Baldanders.info](https://baldanders.info/spiegel/log2/000740.shtml)
- [Prohibiting RC4 — Baldanders.info](https://baldanders.info/spiegel/log2/000810.shtml)
- [CRYPTREC | SHA-1の安全性について](http://www.cryptrec.go.jp/topics/cryptrec_20151218_sha1_cryptanalysis.html)
    - {{< pdf-file title="CRYPTREC暗号技術ガイドライン(SHA-1)" link="http://www.cryptrec.go.jp/report/c13_tech_guideline_SHA-1_web.pdf" >}}
- [scryptがGPUに破られる時 | びりあるの研究ノート](https://blog.visvirial.com/articles/519)

- [最初の SHA-1 衝突例]({{< ref "/remark/2017/02/sha-1-collision.md" >}})
- [「パスワードのベストプラクティス」が変わる]({{< ref "/remark/2017/10/changes-in-password-best-practices.md" >}})

[NIST]: https://www.nist.gov/ "National Institute of Standards and Technology | NIST"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071"><img src="https://images-fe.ssl-images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071">暗号化 プライバシーを救った反乱者たち</a></dt>
	<dd>スティーブン・レビー</dd>
	<dd>斉藤 隆央 (翻訳)</dd>
    <dd>紀伊國屋書店 2002-02-16</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4314009071, EAN: 9784314009072</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-16">2018-12-16</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
