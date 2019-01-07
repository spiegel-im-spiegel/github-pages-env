+++
title = "Issuer Fingerprint Signature Subpacket in Next OpenPGP"
date =  "2017-12-01T17:52:08+09:00"
update =  "2018-01-28T18:00:19+09:00"
description = "署名パケットに関して，次期 OpenPGP (RFC 4880bis) の実装が既に GnuPG に一部入っているようである。"
image = "/images/attention/openpgp.png"
tags        = [ "openpgp", "gnupg", "cryptography" ]

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
  mathjax = false
  mermaidjs = false
+++

(move from [{{< ref "/remark/2017/11/issuer-fingerprint-signature-subpacket-in-next-openpgp.md" >}}]({{< ref "/remark/2017/11/issuer-fingerprint-signature-subpacket-in-next-openpgp.md" >}} "Issuer Fingerprint Signature Subpacket in Next OpenPGP"))

「[OpenPGP 鍵管理に関する考察]({{< ref "/openpgp/openpgp-key-management.md" >}})」を書いていて気づいたのだが，署名パケット（signature packet）に関して，次期 OpenPGP ([RFC 4880bis]) の実装が既に [GnuPG] に一部入っているようである[^gpg1]。

[^gpg1]: 先行してドラフト仕様が [GnuPG] に組み込まれるのは珍しいことではない。

署名パケットの中に署名サブパケット（signature subpacket; 鍵や署名に関する属性情報が入っている）というのがあって [RFC 4880] では sub 32 まで ID が振られているのだけど（プライベート用は別）， [RFC 4880bis] で sub 33 が追加された。
それが “Issuer Fingerprint” で以下の内容になっている。

{{< fig-quote title="draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format" link="https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/" lang="en" >}}
<q>The OpenPGP Key fingerprint of the key issuing the signature.  This subpacket SHOULD be included in all signatures.  If the version of the issuing key is 4 and an Issuer subpacket is also included in the signature, the key ID of the Issuer subpacket MUST match the low 64 bits of the fingerprint.<br>
Note that the length N of the fingerprint for a version 4 key is 20 octets; for a version 5 key N is 32.
</q>
{{< /fig-quote >}}

まぁ，要するに “Issuer Fingerprint” には署名を行う鍵の鍵指紋（key fingerprint）が入りますよ，ということのようだ。
ちょっと試してみよう。

たとえば “Hello world” の文字列に電子署名してみる。

```text
$ echo Hello world | gpg -u 0x7E20B81C --clear-sign
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA256

Hello world
-----BEGIN PGP SIGNATURE-----

iHUEAREIAB0WIQQbUgLbSj7HdvHgrRi02juufiC4HAUCWhfcngAKCRC02juufiC4
HCPUAP4npfesVUOXy/RbFn65Ci6rgtvrFNeNIfmFjYit/xMEywD/eHflgDJQWY+Y
7g7btse2kmbJvRwaKUf1QMgntzgo07E=
=k8ok
-----END PGP SIGNATURE-----
```

ちなみに `0x7E20B81C` は[私の鍵](https://baldanders.info/spiegel/pubkeys/ "OpenPGP 公開鍵リスト — Baldanders.info")の鍵 ID である。
`--clear-sign` コマンドは署名対象のテキストとその電子署名を ASCII armor で出力する。
これをそのまま [pgpdump] にかけてみよう。

こんな感じになる。

{{< highlight text "hl_lines=7-8" >}}
$ echo Hello world | gpg -u 0x7E20B81C --clear-sign | pgpdump -u
Old: Signature Packet(tag 2)(117 bytes)
        Ver 4 - new
        Sig type - Signature of a canonical text document(0x01).
        Pub alg - DSA Digital Signature Algorithm(pub 17)
        Hash alg - SHA256(hash 8)
        Hashed Sub: issuer fingerprint(sub 33)(21 bytes)
         v4 -   Fingerprint - 1b 52 02 db 4a 3e c7 76 f1 e0 ad 18 b4 da 3b ae 7e 20 b8 1c
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Fri Nov 24 08:56:21 UTC 2017
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0xB4DA3BAE7E20B81C
        Hash left 2 bytes - 23 54
        DSA r(255 bits) - ...
        DSA s(256 bits) - ...
                -> hash(DSA q bits)
{{< /highlight >}}

なお [pgpdump] はバージョン 0.32 で “Issuer Fingerprint” に対応した[^gpd]。

[^gpd]: 私の [gpgpdump] も 0.2.1 以降で対応させている。

[RFC 4880bis] では V5 フォーマットの公開鍵パケットや鍵指紋[^v5] の仕様が出てきている。
これらの仕様を取り込んだ [GnuPG] 2.3 あたりがそのうち出てくるんじゃないかと期待している。

[^v5]: V5 フォーマットの鍵指紋は SHA-3 ではなく SHA256 (SHA-2) を使うようだ。まぁドラフト段階なので変わるかもだけど。 SHA-3 自体は既に ID が振られているので組み込みはやれないこともない。

## ブックマーク

- [OpenPGP: First RFC4880bis Draft]({{< ref "/remark/2015/openpgp-draft-rfc4880bis-first.md" >}})
- [OpenPGP に関する話題]({{< ref "/remark/2017/03/topics-on-openpgp.md" >}})

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
