+++
title = "OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）"
date =  "2017-12-01T17:47:50+09:00"
update =  "2017-12-03T15:12:36+09:00"
description = "RFC 4880bis は，名前の通り，ドラフト段階なので今後変わる可能性がある。正式な RFC 番号が振られた段階でこの記事の最終稿とする予定である。"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "cryptography",
  "openpgp",
]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

(move from [{{< ref "remark/2017/11/algorithms-for-openpgp.md" >}}]({{< ref "remark/2017/11/algorithms-for-openpgp.md" >}} "OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）"))

[gpgpdump] を [RFC 4880bis] に[対応させていく作業]({{< relref "remark/2017/11/gpgpdump-0_3_0-released.md" >}} "gpgpdump 0.3.0 をリリースした")の中で「改めて『[OpenPGP] で利用可能なアルゴリズム』をまとめておいたほうがいいかなぁ」と感じたので，覚え書きとして記しておく。

なお [RFC 4880bis] は，名前の通り，ドラフト段階なので今後変わる可能性がある。
アルゴリズム関連で変更があった場合は随時この記事に加筆・修正していって，正式な RFC 番号が振られた段階でこの記事の最終稿とする予定である。

## 共通鍵暗号アルゴリズム（Symmetric-Key Encryption Algorithms）

[OpenPGP] で利用可能な共通鍵暗号は以下の通り。
なお「鍵長」項目の括弧内はブロック長を指す。
いずれも単位は “bit”。

{{< div-gen >}}
<figure lang="en">
<style>
main table.openpgp-sym td  {
  vertical-align:middle;
}
</style>
<table class="openpgp-sym">
<thead>
<tr><th>ID</th><th>アルゴリズム</th><th>鍵長</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>1</td>
<td class="nowrap">IDEA</td>
<td class='right nowrap'>128 (64)</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">TripleDES</td>
<td class='right nowrap'>168 (64)</td>
<td><a href="https://doi.org/10.6028/NIST.SP.800-67r2">SP800-67 Rev.2 <sup><i class='far fa-file-pdf'></i></sup></a></td>
</tr><tr>
<td class='right'>3</td>
<td class="nowrap">CAST5</td>
<td class='right nowrap'>128 (64)</td>
<td><a href="http://tools.ietf.org/html/rfc2144">RFC2144</a></td>
</tr><tr>
<td class='right'>4</td>
<td class="nowrap">Blowfish</td>
<td class='right nowrap'>128 (64)</td>
<td><q><a href="http://www.schneier.com/paper-blowfish-fse.html">Description of a New Variable-Length Key, 64-Bit Block Cipher (Blowfish)</a></q></td>
</tr><tr>
<td class='right'>5</td>
<td class="nowrap">(Reserved)</td>
<td class='right nowrap'>--</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>6</td>
<td class="nowrap">(Reserved)</td>
<td class='right nowrap'>--</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>7</td>
<td class="nowrap">AES with 128-bit key</td>
<td class='right nowrap'>128 (128)</td>
<td rowspan="3"><a href="https://dx.doi.org/10.6028/NIST.FIPS.197">FIPS PUB 197 <sup><i class='far fa-file-pdf'></i></sup></a></td>
</tr><tr>
<td class='right'>8</td>
<td class="nowrap">AES with 192-bit key</td>
<td class='right nowrap'>192 (128)</td>
</tr><tr>
<td class='right'>9</td>
<td class="nowrap">AES with 256-bit key</td>
<td class='right nowrap'>256 (128)</td>
</tr><tr>
<td class='right'>10</td>
<td class="nowrap">Twofish with 256-bit key</td>
<td class='right nowrap'>256 (128)</td>
<td></td>
</tr><tr>
<td class='right'>11</td>
<td class="nowrap">Camellia with 128-bit key</td>
<td class='right nowrap'>128 (128)</td>
<td rowspan="3"><a href="http://tools.ietf.org/html/rfc3713">RFC3713</a>, <a href="https://tools.ietf.org/html/rfc5581">RFC5581</a></td>
</tr><tr>
<td class='right'>12</td>
<td class="nowrap">Camellia with 192-bit key</td>
<td class='right nowrap'>192 (128)</td>
</tr><tr>
<td class='right'>13</td>
<td class="nowrap">Camellia with 256-bit key</td>
<td class='right nowrap'>256 (128)</td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な共通鍵暗号アルゴリズム一覧</figcaption>
</figure>
{{< /div-gen >}}

ID は [OpenPGP] で定義されるもので ”sym 1” のように表記する。

- sym 2 の TripleDES は "MUST implement” である。また sym 3 の CAST5 と sym 7 の AES-128 は "SHOULD implement” である（AES は MUST でいいと思うんだけどなぁ）
- 古い PGP（2.6 以前）の暗号データを復号するためには sym 1 の IDEA が必要となる

ちなみに sym 11 から sym 13 の [Camellia 暗号は日本製](http://www.baldanders.info/spiegel/log2/000451.shtml "The Camellia Cipher in OpenPGP — Baldanders.info")である。

## 公開鍵暗号アルゴリズム（Public-Key Encryption Algorithms）

[OpenPGP] で利用可能な公開鍵暗号は以下の通り。

{{< div-gen >}}
<figure lang="en">
<style>
main table.openpgp-pub td  {
  vertical-align:middle;
}
</style>
<table class="openpgp-pub">
<thead>
<tr><th>ID</th><th>アルゴリズム</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>1</td>
<td class="nowrap">RSA (Encrypt or Sign)</td>
<td rowspan="3"><a href="http://tools.ietf.org/html/rfc3447">RFC3447</a></td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">RSA Encrypt-Only</td>
</tr><tr>
<td class='right'>3</td>
<td class="nowrap">RSA Sign-Only</td>
</tr><tr>
<td class='right'>16</td>
<td class="nowrap">Elgamal<br>(Encrypt-Only)</td>
<td><q><a href="http://crypto.csail.mit.edu/classes/6.857/papers/elgamal.pdf">A public key cryptosystem and a signature scheme based on discrete logarithms <sup><i class='far fa-file-pdf'></i></sup></a></q></td>
</tr><tr>
<td class='right'>17</td>
<td class="nowrap">DSA</td>
<td><a href="http://doi.org/10.6028/NIST.FIPS.186-4">FIPS PUB 186-4 <sup> <i class='far fa-file-pdf'></i></sup></a></td>
</tr><tr>
<td class='right'>18</td>
<td class="nowrap">ECDH public key algorithm</td>
<td><a href="http://doi.org/10.6028/NIST.SP.800-56Ar2">SP800-56A Revision 2 <sup><i class='far fa-file-pdf'></i></sup></a>, <a href="https://tools.ietf.org/html/rfc6637">RFC6637</a></td>
</tr><tr>
<td class='right'>19</td>
<td class="nowrap">ECDSA public key algorithm</td>
<td><a href="http://doi.org/10.6028/NIST.FIPS.186-4">FIPS PUB 186-4 <sup><i class='far fa-file-pdf'></i></sup></a>, <a href="https://tools.ietf.org/html/rfc6090">RFC6090</a>, <a href="https://tools.ietf.org/html/rfc6637">RFC6637</a></td>
</tr><tr>
<td class='right'>20</td>
<td class="nowrap">(Reserved)</td>
<td>(formerly Elgamal Encrypt or Sign)</td>
</tr><tr>
<td class='right'>21</td>
<td class="nowrap" colspan="2">(Reserved for Diffie-Hellman)</td>
</tr><tr>
<td class='right'>22</td>
<td class="nowrap">EdDSA</td>
<td><a href="http://link.springer.com/article/10.1007%2Fs13389-012-0027-1">High-speed high-security signatures</a>, <a href="https://tools.ietf.org/html/rfc8032">RFC8032</a></td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な公開鍵暗号アルゴリズム一覧</figcaption>
</figure>
{{< /div-gen >}}

ID は [OpenPGP] で定義されるもので ”pub 1” のように表記する。

- pub 2 の RSA Encrypt-Only および pub 3 の RSA Sign-Only は deprecated なので，これらの鍵は新たに作成するべきではない（SHOULD NOT be generated）
- 表には書いてないが pub 4 から pub 15 は予約済み（Reserved）である
- [RFC 4880bis] では pub 16 の ElGamal, pub 17 の DSA, pub 18 の ECDH, pub 19 の ECDSA が "MUST implement” となる（現行の [RFC 4880] では ElGamal と DSA が "MUST implement”）
- pub 20 の ElGamal は，元々暗号化と署名の両方できるものだったが，脆弱性が見つかったため [OpenPGP] では使用禁止になった
- pub 22 の EdDSA は2017年1月に [RFC 8032] として正式に RFC 化された。ただし [OpenPGP] ではまだドラフト段階である
-  pub 18 の ECDH, pub 19 の ECDSA, pub 22 の EdDSA で利用可能な楕円曲線については「[そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< relref "openpgp/using-ecc-with-gnupg.md" >}})」で紹介している

## 一方向ハッシュ関数アルゴリズム（Hash Algorithms）

[OpenPGP] で利用可能なハッシュ関数は以下の通り。

{{< div-gen >}}
<figure lang="en">
<style>
main table.openpgp-hash td  {
  vertical-align:middle;
}
</style>
<table class="openpgp-hash">
<thead>
<tr><th>ID</th><th>アルゴリズム</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>1</td>
<td class="nowrap">MD5</td>
<td><a href="https://tools.ietf.org/html/rfc1321">RFC1321</a></td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">SHA-1</td>
<td><a href="http://doi.org/10.6028/NIST.FIPS.186-4">FIPS PUB 186-4 <sup><i class='far fa-file-pdf'></i></sup></a></td>
</tr><tr>
<td class='right'>3</td>
<td class="nowrap">RIPE-MD/160</td>
<td><q><a href="http://homes.esat.kuleuven.be/~bosselae/ripemd160.html">The hash function RIPEMD-160</a></q></td>
</tr><tr>
<td class='right'>4</td>
<td class="nowrap">(Reserved)</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>5</td>
<td class="nowrap">(Reserved)</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>6</td>
<td class="nowrap">(Reserved)</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>7</td>
<td class="nowrap">(Reserved)</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>8</td>
<td class="nowrap">SHA2-256</td>
<td rowspan="4"><a href="http://doi.org/10.6028/NIST.FIPS.180-4">FIPS PUB 180-4 <sup><i class='far fa-file-pdf'></i></sup></a></td>
</tr><tr>
<td class='right'>9</td>
<td class="nowrap">SHA2-384</td>
</tr><tr>
<td class='right'>10</td>
<td class="nowrap">SHA2-512</td>
</tr><tr>
<td class='right'>11</td>
<td class="nowrap">SHA2-224</td>
</tr><tr>
<td class='right'>12</td>
<td class="nowrap">SHA3-256</td>
<td><a href="http://doi.org/10.6028/NIST.FIPS.202">FIPS PUB 202 <sup><i class='far fa-file-pdf'></i></sup></a></td>
</tr><tr>
<td class='right'>13</td>
<td class="nowrap">(Reserved)</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>14</td>
<td class="nowrap">SHA3-512</td>
<td><a href="http://doi.org/10.6028/NIST.FIPS.202">FIPS PUB 202 <sup><i class='far fa-file-pdf'></i></sup></a></td>
</tr></tbody>
</table>
<figcaption>OpenPGP で使用可能な一方向ハッシュ関数一覧</figcaption>
</figure>
{{< /div-gen >}}

ID は [OpenPGP] で定義されるもので ”hash 1” のように表記する。

- [RFC 4880bis] では hash 8 の SHA2-256 が "MUST implement” となる（現行の [RFC 4880] では SHA-1 が "MUST implement”）。鍵指紋（key fingerprint）についても V5 では SHA2-256 を使用することになる（現行は V4）
- [RFC 4880bis] では hash 1 の MD5 と hash 3 の RIPE-MD/160 は "SHOULD NOT use” となる。これらは古い PGP による鍵や暗号データの互換性のためにのみ残されるだろう
- [RFC 4880bis] では hash 2 の SHA-1 も基本的に非推奨になる（SHOULD NOT create messages）。ただし V4 の鍵指紋や MDC (Modification Detection Code) 用には引き続き使われる

## その他のアルゴリズム

### S2K (String-to-Key)

S2K はパスフレーズからセッション鍵を生成するためのハッシュ化の手順である。

<figure>
<table>
<thead>
<tr><th>ID</th><th>S2K タイプ</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>0</td>
<td>Simple S2K</td>
</tr><tr>
<td class='right'>1</td>
<td>Salted S2K</td>
</tr><tr>
<td class='right'>2</td>
<td>Reserved value</td>
</tr><tr>
<td class='right'>3</td>
<td>Iterated and Salted S2K</td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な S2K アルゴリズム一覧</figcaption>
</figure>

ID は [OpenPGP] で定義されるもので ”s2k 1” のように表記する。

[OpenPGP] では，パスフレーズ自体はいかなる形（ハッシュ値を含む）でも保存しない。
このため，パスフレーズを紛失してしまった場合は復元できない[^pp1]。

[^pp1]: ただし [OpenPGP] では試行回数によるロックアウトは定義されないため無限にパスフレーズ解読を試みることができる。

### 乱数生成器（Random Number Generator）

{{< fig-quote title="draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format" link="https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/" lang="en" >}}
<q>Certain operations in this specification involve the use of random numbers.  An appropriate entropy source should be used to generate these numbers (see [<a href="https://tools.ietf.org/html/rfc4086">RFC4086</a>]).</q>
{{< /fig-quote >}}

（リンクは私によるもの）

### データ圧縮（Compression Algorithms）

暗号化メッセージや電子署名を圧縮するためのアルゴリズムである。

<figure>
<table>
<thead>
<tr><th>ID</th><th>アルゴリズム</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>0</td>
<td class="nowrap">Uncompressed</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>1</td>
<td class="nowrap">ZIP</td>
<td><a href="https://tools.ietf.org/html/rfc1951">RFC1951</a></td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">ZLIB</td>
<td><a href="https://tools.ietf.org/html/rfc1950">RFC1950</a></td>
</tr><tr>
<td class='right'>3</td>
<td class="nowrap">BZip2</td>
<td><a href="http://www.bzip.org/">bzip2</a></td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能なデータ圧縮アルゴリズム一覧</figcaption>
</figure>

ID は [OpenPGP] で定義されるもので ”comp 1” のように表記する。

## アルゴリズムの選択

暗号関連のアルゴリズムや鍵長は組み合わせが重要で，あるアルゴリズムのみ強くても効率が悪くなるだけである。
たとえば，2030年以降も使える組み合わせが必要なら

- AES 128bit
- ElGamal, DSA 3072bit
- RSA 3072bit
- ECDH, ECDSA 256bit
- SHA2-256, SHA3-256

の中から組み合わせるのが「[ベストマッチ！](https://dic.pixiv.net/a/%E3%83%93%E3%83%AB%E3%83%89%28%E4%BB%AE%E9%9D%A2%E3%83%A9%E3%82%A4%E3%83%80%E3%83%BC%29)」である[^key1]。
詳しくは「[暗号鍵関連の各種変数について]({{< relref "remark/2017/10/key-parameters.md" >}})」を参照のこと。

[^key1]: もちろんアルゴリズムの危殆化が起きない前提での話である。こればっかりは予測しようもないし（笑） 危殆化があり得ることを前提にするなら，ひとつのアルゴリズムに固定するのではなく，常に代替えを用意する（または用意できるよう準備する）ことが大事である。

## ブックマーク

- [わかる！ OpenPGP 暗号](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)
- [OpenPGP: First RFC4880bis Draft]({{< relref "remark/2015/openpgp-draft-rfc4880bis-first.md" >}})
- [OpenPGP に関する話題]({{< relref "remark/2017/03/topics-on-openpgp.md" >}})
- [OpenPGP 鍵管理に関する考察]({{< relref "remark/2017/11/openpgp-key-management.md" >}})
- [Issuer Fingerprint Signature Subpacket in Next OpenPGP]({{< relref "remark/2017/11/issuer-fingerprint-signature-subpacket-in-next-openpgp.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[RFC 5581]: https://tools.ietf.org/html/rfc5581 "RFC 5581 - The Camellia Cipher in OpenPGP"
[RFC 6637]: https://tools.ietf.org/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[RFC 8032]: https://tools.ietf.org/html/rfc8032 "RFC 8032 - Edwards-Curve Digital Signature Algorithm (EdDSA)"
[RFC 4086]: https://tools.ietf.org/html/rfc4086 "RFC 4086 - Randomness Requirements for Security"
[RFC 1951]: https://tools.ietf.org/html/rfc1951 "RFC 1951 - DEFLATE Compressed Data Format Specification version 1.3"
[RFC 1951]: https://tools.ietf.org/html/rfc1951 "RFC 1951 - DEFLATE Compressed Data Format Specification version 1.3"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/">暗号化 プライバシーを救った反乱者たち</a></dt><dd>スティーブン・レビー 斉藤 隆央 </dd><dd>紀伊國屋書店 2002-02-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487593100X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/487593100X.09._SCTHUMBZZZ_.jpg"  alt="ハッカーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4105393022/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4105393022.09._SCTHUMBZZZ_.jpg"  alt="暗号解読―ロゼッタストーンから量子暗号まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4484111160/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4484111160.09._SCTHUMBZZZ_.jpg"  alt="グーグル ネット覇者の真実 追われる立場から追う立場へ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/410215972X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/410215972X.09._SCTHUMBZZZ_.jpg"  alt="暗号解読〈上〉 (新潮文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4102159738/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4102159738.09._SCTHUMBZZZ_.jpg"  alt="暗号解読 下巻 (新潮文庫 シ 37-3)"  /></a> </p>
<p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-03-09">2015/03/09</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
