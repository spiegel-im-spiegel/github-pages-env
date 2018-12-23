+++
title = "次期 OpenPGP と AEAD"
date =  "2018-01-27T16:21:29+09:00"
update =  "2018-01-28T18:40:55+09:00"
description = "ドラフト 03, 04 の主な変更点としては AEAD (Authenticated Encryption with Associated Data; 認証付き暗号) の仕様が追加されたことだろう。"
image = "/images/attention/remark.jpg"
tags = ["security", "cryptography", "openpgp", "aead"]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
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
  mathjax = false
  mermaidjs = false
+++

次期 OpenPGP ([RFC 4880bis]) のドラフト 03 および 04 が登場している。

- [draft-ietf-openpgp-rfc4880bis-03.txt](https://tools.ietf.org/id/draft-ietf-openpgp-rfc4880bis-03.txt)
- [draft-ietf-openpgp-rfc4880bis-04.txt](https://tools.ietf.org/id/draft-ietf-openpgp-rfc4880bis-04.txt)

ドラフト 04 の有効期限は 2018-07-29 まで。

ドラフト 03, 04 の主な変更点としては AEAD (Authenticated Encryption with Associated Data; 認証付き暗号) の仕様が追加されたことだろう。
[RFC 4880bis] では V5 と呼ばれる仕様[^pct1] が追加されるが，どうやら V5 を AEAD 準拠にする目論見があるようだ。

[^pct1]: [OpenPGP] では暗号鍵や電子署名・暗号文などの情報を「パケット」と呼ばれるブロック単位で管理しているが， V5 というのはこのパケットのバージョンを指している。具体的には旧 PGP 2.6.x ([RFC 1991]) の仕様が V3。それ以降の OpenPGP ([RFC 2440], [RFC 4880]) で追加されたのが V4 となる。

AEAD 的な仕組みの必要性は [OpenPGP] でも随分昔から認識されていたが，現行の [RFC 4880] では MDC (Modification Detection Code) と呼ばれる独自の実装を行っている。
しかし V5 で AEAD を導入することにより MDC は不要になるわけだ（後方互換性確保のための実装のみ残る感じ？）。

{{< fig-quote title="draft-ietf-openpgp-rfc4880bis-04.txt" link="https://tools.ietf.org/id/draft-ietf-openpgp-rfc4880bis-04.txt" lang="en   " >}}
<q>There won't be any future versions of this packet because the MDC system has been superseded by the AEAD Encrypted Data packet.</q>
{{< /fig-quote >}}

AEAD のアルゴリズムとしては EAX mode と OCB mode をサポートするようだ。

<figure>
<table>
<thead>
<tr><th>ID</th><th>AEAD アルゴリズム</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>1</td>
<td class="nowrap">EAX</td>
<td><a href="https://eprint.iacr.org/2003/069">EAX: A Conventional Authenticated-Encryption Mode</a></td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">OCB</td>
<td><a href="https://tools.ietf.org/html/rfc7253">RFC7253</a></td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な AEAD アルゴリズム一覧</figcaption>
</figure>

このうち EAX mode は “MUST implement”。
OCB mode に関しては，特許問題が絡むようで，取り扱いを議論中のようだ。

{{< fig-quote title="draft-ietf-openpgp-rfc4880bis-04.txt" link="https://tools.ietf.org/id/draft-ietf-openpgp-rfc4880bis-04.txt" lang="en   " >}}
<q>The OCB mode is patented and a debate is still underway on whether it can be included in RFC4880bis or needs to be moved to a separate document.  For the sole purpose of experimenting with the Preferred AEAD Algorithms signature subpacket it is has been included in this I-D.</q>
{{< /fig-quote >}}

更に AEAD を導入するにあたり，共通鍵暗号アルゴリズムの AES-128 が “MUST implement” になった。
[RFC 4880] で “MUST implement” だった TeipleDES は後方互換性のためにのみ残される。

[OpenPGP]: http://openpgp.org/
[RFC 1991]: http://tools.ietf.org/html/rfc1991 "RFC 1991 - OpenPGP Message Format"
[RFC 2440]: https://tools.ietf.org/html/rfc2440 "RFC 2440 - OpenPGP Message Format"
[RFC 4880]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## ブックマーク

- [OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）]({{< ref "/openpgp/algorithms-for-openpgp.md" >}})

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
