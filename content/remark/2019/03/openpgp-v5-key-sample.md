+++
title = "OpenPGP V5 鍵のサンプル"
date = "2019-03-21T10:11:37+09:00"
description = "よーし，書くぞ書くぞ書くぞ！"
image = "/images/attention/openpgp.png"
tags = [
  "cryptography",
  "openpgp",
  "gpgpdump",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP メーリングリスト](https://mailarchive.ietf.org/arch/browse/openpgp/)に次期 OpenPGP ([RFC 4880bis]) の v5 パケットに対応する秘密鍵のサンプルが上がっている。

- [[openpgp] v5 sample key](https://mailarchive.ietf.org/arch/msg/openpgp/9SheW_LENE0Kxf7haNllovPyAdY)

全文引用ご容赦。

{{% fig-gen type="md" title="v5 sample key" link="https://mailarchive.ietf.org/arch/msg/openpgp/9SheW_LENE0Kxf7haNllovPyAdY" lang="en" %}}
```text
Hi!

I recently implemented v5 key support in gpg.  I include a sample secret
key and would be interested to see whether other implementations can
parse it, verify the signatures and that the fingerprints match.  The
key is not protected but the dash lines are indented to stop MUAs from
processing the key.


Shalom-Salam,

   Werner


--8<---------------cut here---------------start------------->8---

sec   ed25519 2019-03-20 [SC]
      19347BC9872464025F99DF3EC2E0000ED9884892E1F7B3EA4C94009159569B54
uid                      emma.goldman@example.net
ssb   cv25519 2019-03-20 [E]
      E4557C2B02FFBF4B04F87401EC336AF7133D0F85BE7FD09BAEFD9CAEB8C93965

  -----BEGIN PGP PRIVATE KEY BLOCK-----

lGEFXJH05BYAAAAtCSsGAQQB2kcPAQEHQFhZlVcVVtwf+21xNQPX+ecMJJBL0MPd
fj75iux+my8QAAAAAAAiAQCHZ1SnSUmWqxEsoI6facIVZQu6mph3cBFzzTvcm5lA
Ng5ctBhlbW1hLmdvbGRtYW5AZXhhbXBsZS5uZXSIlgUTFggASCIhBRk0e8mHJGQC
X5nfPsLgAA7ZiEiS4fez6kyUAJFZVptUBQJckfTkAhsDBQsJCAcCAyICAQYVCgkI
CwIEFgIDAQIeBwIXgAAA9cAA/jiR3yMsZMeEQ40u6uzEoXa6UXeV/S3wwJAXRJy9
M8s0AP9vuL/7AyTfFXwwzSjDnYmzS0qAhbLDQ643N+MXGBJ2BZxmBVyR9OQSAAAA
MgorBgEEAZdVAQUBAQdA+nysrzml2UCweAqtpDuncSPlvrcBWKU0yfU0YvYWWAoD
AQgHAAAAAAAiAP9OdAPppjU1WwpqjIItkxr+VPQRT8Zm/Riw7U3F6v3OiBFHiHoF
GBYIACwiIQUZNHvJhyRkAl+Z3z7C4AAO2YhIkuH3s+pMlACRWVabVAUCXJH05AIb
DAAAOSQBAP4BOOIR/sGLNMOfeb5fPs/02QMieoiSjIBnijhob2U5AQC+RtOHCHx7
TcIYl5/Uyoi+FOvPLcNw4hOv2nwUzSSVAw==
=IiS2
  -----END PGP PRIVATE KEY BLOCK-----

--8<---------------cut here---------------end--------------->8---
```
{{% /fig-gen %}}

とりあえず [gpgpdump] で中身を見てみる[^pd1]。

[^pd1]: 本家 [pgpdump] は [RFC 4880] よりあとの仕様追加・変更にはほとんど対応していないため v5 パケットも見れない。ただし[リポジトリ](https://github.com/kazu-yamamoto/pgpdump "kazu-yamamoto/pgpdump: A PGP packet visualizer")では割と積極的に pull request を受け入れているようなので，興味があればコードを寄与するのも悪くないだろう。個人的には Haskell 版 [pgpdump] が出れば面白いのに，と思ったりする。テストは既に Haskell で書いてるみたいだし。

```text
$ gpgpdump testdata/v5/sample.txt
Secret-Key Packet (tag 5) (97 bytes)
        Version: 5 (draft)
        Public-Key
                Public key creation time: 2019-03-20T17:08:04+09:00
                        5c 91 f4 e4
                Public-key Algorithm: EdDSA (pub 22)
                ECC Curve OID: ed25519 (256bits key size)
                        2b 06 01 04 01 da 47 0f 01
                EdDSA EC point (40 || compressd format) (263 bits)
        Secret-Key (the secret-key data is not encrypted.)
                EdDSA secret key (0 bits)
                Checksum
                        00 00
        Unknown data (37 bytes)
User ID Packet (tag 13) (24 bytes)
        User ID: emma.goldman@example.net
Signature Packet (tag 2) (150 bytes)
        Version: 5 (unknown)
        Unknown data (149 bytes)
Secret-Subkey Packet (tag 7) (102 bytes)
        Version: 5 (draft)
        Public-Key
                Public key creation time: 2019-03-20T17:08:04+09:00
                        5c 91 f4 e4
                Public-key Algorithm: ECDH public key algorithm (pub 18)
                ECC Curve OID: cv25519 (256bits key size)
                        2b 06 01 04 01 97 55 01 05 01
                ECDH EC point (40 || X) (263 bits)
                KDF parameters (3 bytes)
                        Hash Algorithm: SHA2-256 (hash 8)
                        Symmetric Algorithm: AES with 128-bit key (sym 7)
        Secret-Key (the secret-key data is not encrypted.)
                ECDH secret key (0 bits)
                Checksum
                        00 00
        Unknown data (37 bytes)
Signature Packet (tag 2) (122 bytes)
        Version: 5 (unknown)
        Unknown data (121 bytes)
```

まだ v5 パケットに対応してないので妙な表示になっているが入れ物は出来てるので，あとはコードを書くだけだ。

書くぞ書くぞ書くぞ！

[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## ブックマーク

- [次期 OpenPGP と AEAD]({{< ref "/remark/2018/01/next-openpgp-with-aead.md" >}})
- [OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）]({{< ref "/openpgp/algorithms-for-openpgp.md" >}})

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
