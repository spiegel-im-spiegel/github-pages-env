+++
date = "2016-05-04T23:28:56+09:00"
update = "2016-05-06T22:26:53+09:00"
description = "“Satoshi Nakamoto” の OpenPGP 鍵について。"
draft = false
tags = ["cryptography", "openpgp", "certification"]
title = "Satoshi Nakamoto の OpenPGP 鍵"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

色々話題になっているようで。

- [オーストラリア人起業家、Craig WrightがBitcoinの発明者、Satoshi Nakamotoなのか？ | TechCrunch Japan](http://jp.techcrunch.com/2016/05/03/20160502major-questions-arise-over-craig-wrights-claim-to-be-satoshi-nakamoto/)
- [本の虫: Craig WrightがSatoshi Nakamotoだとする証明はない](http://cpplover.blogspot.jp/2016/05/craig-wrightsatoshi-nakamoto.html)
- [Bitcoin’s Creator Satoshi Nakamoto Is Probably This Unknown Australian Genius | WIRED](https://www.wired.com/2015/12/bitcoins-creator-satoshi-nakamoto-is-probably-this-unknown-australian-genius/)
    - [この無名の天才がビットコイン発明者「サトシ・ナカモト」である証拠（1）｜WIRED.jp](http://wired.jp/2016/05/03/bitcoins-creator-satoshi-nakamoto-is-1/)
    - [この無名の天才がビットコイン発明者「サトシ・ナカモト」である証拠（2）｜WIRED.jp](http://wired.jp/2016/05/04/bitcoins-creator-satoshi-nakamoto-is-2/)
    - [この無名の天才がビットコイン発明者「サトシ・ナカモト」である証拠（3）｜WIRED.jp](http://wired.jp/2016/05/05/bitcoins-creator-satoshi-nakamoto-is-3/)

この中で

{{< fig-quote title="この無名の天才がビットコイン発明者「サトシ・ナカモト」である証拠（2）" link="http://wired.jp/2016/05/04/bitcoins-creator-satoshi-nakamoto-is-2/" >}}
<q>保管先のMITサーヴァーのデータベースをチェックしてみると、この公開鍵は〈satoshin@vistomail.com〉というメールアドレスに紐づけられているのがわかる。これはビットコインを紹介する白書を暗号研究メーリングリストに投稿する際にナカモトが使用したアドレス、〈satoshi@vistomail.com〉によく似ている。</q>
{{< /fig-quote >}}

とあるが，おそらく[これ](https://sks-keyservers.net/pks/lookup?op=vindex&search=0x491F9BDF0F7BD4AD)を指していると思われる[^sks]。

[^sks]: OpenPGP 鍵サーバは互いに peer-to-peer で繋がり同期している。今回は MIT の鍵サーバではなく [`sks-keyservers.net`](https://sks-keyservers.net/) から取得した。内容は同じである。

```text
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: SKS 1.1.5+
Comment: Hostname: keys2.kfwebs.net

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
=LSEt
-----END PGP PUBLIC KEY BLOCK-----
```

これを [`pgpdump`] にかけるとこうなる。

```text
Old: Public Key Packet(tag 6)(397 bytes)
        Ver 4 - new
        Public key creation time - Thu Oct 30 11:50:43 東京 (標準時) 2008
        Pub alg - RSA Encrypt or Sign(pub 1)
        RSA n(3072 bits) - ...
        RSA e(17 bits) - ...
Old: User ID Packet(tag 13)(41 bytes)
        User ID - BuZz spacedout   <satoshin@vistomail.com>
Old: Signature Packet(tag 2)(447 bytes)
        Ver 4 - new
        Sig type - Positive certification of a User ID and Public Key packet(0x13).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA1(hash 2)
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Thu Oct 30 11:50:43 東京 (標準時) 2008
        Hashed Sub: key flags(sub 27)(1 bytes)
                Flag - This key may be used to certify other keys
                Flag - This key may be used to sign data
        Hashed Sub: key expiration time(sub 9)(4 bytes)
                Time - Wed Oct 30 10:00:00 東京 (標準時) 2019
        Hashed Sub: preferred symmetric algorithms(sub 11)(6 bytes)
                Sym alg - AES with 256-bit key(sym 9)
                Sym alg - AES with 192-bit key(sym 8)
                Sym alg - AES with 128-bit key(sym 7)
                Sym alg - CAST5(sym 3)
                Sym alg - Triple-DES(sym 2)
                Sym alg - IDEA(sym 1)
        Hashed Sub: preferred hash algorithms(sub 21)(5 bytes)
                Hash alg - SHA256(hash 8)
                Hash alg - SHA1(hash 2)
                Hash alg - SHA384(hash 9)
                Hash alg - SHA512(hash 10)
                Hash alg - SHA224(hash 11)
        Hashed Sub: preferred compression algorithms(sub 22)(3 bytes)
                Comp alg - ZLIB <RFC1950>(comp 2)
                Comp alg - BZip2(comp 3)
                Comp alg - ZIP <RFC1951>(comp 1)
        Hashed Sub: features(sub 30)(1 bytes)
                Flag - Modification detection (packets 18 and 19)
        Hashed Sub: key server preferences(sub 23)(1 bytes)
                Flag - No-modify
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0x491F9BDF0F7BD4AD
        Hash left 2 bytes - f5 97
        RSA m^d mod n(3071 bits) - ...
                -> PKCS-1
Old: User ID Packet(tag 13)(41 bytes)
        User ID - Satoshi Nakamoto <satoshin@vistomail.com>
Old: Signature Packet(tag 2)(447 bytes)
        Ver 4 - new
        Sig type - Positive certification of a User ID and Public Key packet(0x13).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA1(hash 2)
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Thu Oct 30 11:50:43 東京 (標準時) 2008
        Hashed Sub: key flags(sub 27)(1 bytes)
                Flag - This key may be used to certify other keys
                Flag - This key may be used to sign data
        Hashed Sub: key expiration time(sub 9)(4 bytes)
                Time - Wed Oct 30 10:00:00 東京 (標準時) 2019
        Hashed Sub: preferred symmetric algorithms(sub 11)(6 bytes)
                Sym alg - AES with 256-bit key(sym 9)
                Sym alg - AES with 192-bit key(sym 8)
                Sym alg - AES with 128-bit key(sym 7)
                Sym alg - CAST5(sym 3)
                Sym alg - Triple-DES(sym 2)
                Sym alg - IDEA(sym 1)
        Hashed Sub: preferred hash algorithms(sub 21)(5 bytes)
                Hash alg - SHA256(hash 8)
                Hash alg - SHA1(hash 2)
                Hash alg - SHA384(hash 9)
                Hash alg - SHA512(hash 10)
                Hash alg - SHA224(hash 11)
        Hashed Sub: preferred compression algorithms(sub 22)(3 bytes)
                Comp alg - ZLIB <RFC1950>(comp 2)
                Comp alg - BZip2(comp 3)
                Comp alg - ZIP <RFC1951>(comp 1)
        Hashed Sub: features(sub 30)(1 bytes)
                Flag - Modification detection (packets 18 and 19)
        Hashed Sub: key server preferences(sub 23)(1 bytes)
                Flag - No-modify
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0x491F9BDF0F7BD4AD
        Hash left 2 bytes - f5 97
        RSA m^d mod n(3071 bits) - ...
                -> PKCS-1
Old: Public Subkey Packet(tag 14)(397 bytes)
        Ver 4 - new
        Public key creation time - Thu Oct 30 11:50:43 東京 (標準時) 2008
        Pub alg - RSA Encrypt or Sign(pub 1)
        RSA n(3072 bits) - ...
        RSA e(17 bits) - ...
Old: Signature Packet(tag 2)(421 bytes)
        Ver 4 - new
        Sig type - Subkey Binding Signature(0x18).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA1(hash 2)
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Thu Oct 30 11:50:43 東京 (標準時) 2008
        Hashed Sub: key flags(sub 27)(1 bytes)
                Flag - This key may be used to encrypt communications
                Flag - This key may be used to encrypt storage
        Hashed Sub: key expiration time(sub 9)(4 bytes)
                Time - Wed Oct 30 10:00:00 東京 (標準時) 2019
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0x491F9BDF0F7BD4AD
        Hash left 2 bytes - 57 48
        RSA m^d mod n(3072 bits) - ...
                -> PKCS-1
```

というわけで，時期的にはそれっぽいが，よく見ると本人以外の署名がない。
これではこの鍵が本人かどうか確かめようがない。

同じ日に作られた “Satoshi Nakamoto” 名義の鍵は他にもあって， [`satoshi@vistomail.com` の鍵](https://sks-keyservers.net/pks/lookup?op=vindex&search=0xD2D59294CDD2C21C)もあるのだが，これもやはり本人の署名しかない。
更に [`satoshin@gmx.com` の鍵](https://sks-keyservers.net/pks/lookup?op=vindex&search=0x18C09E865EC948A1)もあって，こちらには2009年の署名がいくつかある。

御存知の通り OpenPGP の信用モデルは “web of trust” と呼ばれるもので，ユーザ間の相互署名によって信用度と有効性が担保される[^ksp]。
故に自己署名しかない鍵では有効性を担保できない。
もっとも “Satoshi Nakamoto” は匿名で活動していたらしいので，他者の署名を得られなかった可能性もある。

[^ksp]: 公開鍵の相互署名を行うために [key signing party](https://en.wikipedia.org/wiki/Key_signing_party) を行ったりする。もし，件の鍵が [key signing party](https://en.wikipedia.org/wiki/Key_signing_party) によって署名されていたなら，そのときの証言を以って本人の鍵かどうか証明できたわけだ。

巷では否定的な意見が多いようだが，真相はわからない。
もっとも，今更名乗り出られてもねぇ。
本当に本人なら税務署が押し寄せてきて大変なことになりそうだが。

そういえば昔，電子メールに署名する際に，メッセージに差出人や宛先の情報がない場合，それを使って第三者がいくらでも悪用できるって話題があったな。
この議論以降，電子メールの本文には必ず宛先の名前と自分自身を示す情報を含めるようにしている。

## 参考ページ

- [わかる！ OpenPGP 暗号 — Baldanders.info](https://baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)
- [Craig Wright、「証拠は公表したくない」―オーストラリア人起業家、Bitcoinを発明したとの主張から後退 | TechCrunch Japan](http://jp.techcrunch.com/2016/05/06/20160505craig-wright-backs-out-and-wont-prove-that-he-is-bitcoin-creator-satoshi-nakamoto/) : この話はどうやら収束したらしい。やれやれ

[`pgpdump`]: (http://www.mew.org/~kazu/proj/pgpdump/

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
