+++
title = "gpgpdump v0.5.0 をリリースした"
description = "--indent オプションを追加しV5 パケットへの対応を更に進めた。"
date = "2019-03-24T10:57:27+09:00"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "golang", "gpgpdump"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を視覚化する [gpgpdump] の v0.5.0 をリリースした。

- [Release v0.5.0 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.5.0)

まず `--indent` オプションを追加した。

{{< highlight text "hl_lines=9" >}}
$ gpgpdump -h
Usage:
  gpgpdump [flags] [OpenPGP file]

Flags:
  -a, --armor        accepts ASCII input only
      --debug        for debug
  -h, --help         help for gpgpdump
      --indent int   indent size for output string
  -i, --int          dumps multi-precision integers
  -j, --json         output with JSON format
  -l, --literal      dumps literal packets (tag 11)
  -m, --marker       dumps marker packets (tag 10)
  -p, --private      dumps private packets (tag 60-63)
  -t, --toml         output with TOML format
  -u, --utc          output with UTC time
  -v, --version      output version of gpgpdump
{{< /highlight >}}

このオプションで出力テキストのインデント幅を調整できる。

```text
$ cat testdata/eccsig.asc
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2

iF4EARMIAAYFAlTDCN8ACgkQMfv9qV+7+hg2HwEA6h2iFFuCBv3VrsSf2BREQaT1
T1ZprZqwRPOjiLJg9AwA/ArTwCPz7c2vmxlv7sRlRLUI6CdsOqhuO1KfYXrq7idI
=ZOTN
-----END PGP SIGNATURE-----

$ cat testdata/eccsig.asc | gpgpdump -u --indent 2
Signature Packet (tag 2) (94 bytes)
  Version: 4 (current)
  Signiture Type: Signature of a canonical text document (0x01)
  Public-key Algorithm: ECDSA public key algorithm (pub 19)
  Hash Algorithm: SHA2-256 (hash 8)
  Hashed Subpacket (6 bytes)
    Signature Creation Time (sub 2): 2015-01-24T02:52:15Z
  Unhashed Subpacket (10 bytes)
    Issuer (sub 16): 0x31fbfda95fbbfa18
  Hash left 2 bytes
    36 1f
  ECDSA value r (256 bits)
  ECDSA value s (252 bits)
```

`--indent` オプションがない場合には今までどおりインデントにタブ文字を使う。

`--indent` オプションは JSON (`-j`) または TOML (`-t`) でも有効となる。
前バージョンからの変更点として JSON 出力時に `--indent` オプションがないとコンパクト出力となる。

```text
$ cat testdata/eccsig.asc | gpgpdump -u -j
{"Packet":[{"name":"Signature Packet (tag 2)","note":"94 bytes","Item":[{"name":"Version","value":"4","note":"current"},{"name":"Signiture Type","value":"Signature of a canonical text document (0x01)"},{"name":"Public-key Algorithm","value":"ECDSA public key algorithm (pub 19)"},{"name":"Hash Algorithm","value":"SHA2-256 (hash 8)"},{"name":"Hashed Subpacket","note":"6 bytes","Item":[{"name":"Signature Creation Time (sub 2)","value":"2015-01-24T02:52:15Z"}]},{"name":"Unhashed Subpacket","note":"10 bytes","Item":[{"name":"Issuer (sub 16)","value":"0x31fbfda95fbbfa18"}]},{"name":"Hash left 2 bytes","dump":"36 1f"},{"name":"ECDSA value r","note":"256 bits"},{"name":"ECDSA value s","note":"252 bits"}]}]}
```

前にリリースした「[jq のようなもの]({{< ref "/release/2019/03/gjq-v0_2_0-is-released.md" >}})」と組み合わせれば

```json
$ cat testdata/eccsig.asc | gpgpdump -u -j | gjq -C .Packet.[0].Item.[0:2]
[
  {
    "name": "Version",
    "value": "4",
    "note": "current"
  },
  {
    "name": "Signiture Type",
    "value": "Signature of a canonical text document (0x01)"
  },
  {
    "name": "Public-key Algorithm",
    "value": "ECDSA public key algorithm (pub 19)"
  }
]
```

とかできるので問題ないだろう。

もうひとつ。
[RFC 4880bis] の V5 パケットへの対応を更に進めた。

具体的には[先日公開されたサンプル]({{< ref "/remark/2019/03/openpgp-v5-key-sample.md" >}} "OpenPGP V5 鍵のサンプル")について

{{< highlight text "hl_lines=11-14 18-57 70-73 75-91" >}}
$ cat testdata/v5/sample.txt | gpgpdump -u --indent 2
Secret-Key Packet (tag 5) (97 bytes)
  Version: 5 (draft)
  Public-Key
    Public key creation time: 2019-03-20T08:08:04Z
      5c 91 f4 e4
    Public-key Algorithm: EdDSA (pub 22)
    ECC Curve OID: ed25519 (256bits key size)
      2b 06 01 04 01 da 47 0f 01
    EdDSA EC point (40 || compressd format) (263 bits)
  Secret-Key (s2k usage 0; plain secret-key material)
    EdDSA secret key (256 bits)
    2-octet checksum
      0e 5c
User ID Packet (tag 13) (24 bytes)
  User ID: emma.goldman@example.net
Signature Packet (tag 2) (150 bytes)
  Version: 5 (draft)
  Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
  Public-key Algorithm: EdDSA (pub 22)
  Hash Algorithm: SHA2-256 (hash 8)
  Hashed Subpacket (72 bytes)
    Issuer Fingerprint (sub 33) (33 bytes)
      Version: 5 (need 32 octets length)
      Fingerprint (32 bytes)
        19 34 7b c9 87 24 64 02 5f 99 df 3e c2 e0 00 0e d9 88 48 92 e1 f7 b3 ea 4c 94 00 91 59 56 9b 54
    Signature Creation Time (sub 2): 2019-03-20T08:08:04Z
    Key Flags (sub 27) (1 bytes)
      Flag: This key may be used to certify other keys.
      Flag: This key may be used to sign data.
    Preferred Symmetric Algorithms (sub 11) (4 bytes)
      Symmetric Algorithm: AES with 256-bit key (sym 9)
      Symmetric Algorithm: AES with 192-bit key (sym 8)
      Symmetric Algorithm: AES with 128-bit key (sym 7)
      Symmetric Algorithm: TripleDES (168 bit key derived from 192) (sym 2)
    Preferred AEAD Algorithms (sub 34) (2 bytes)
      AEAD Algorithm: OCB mode <RFC7253> (aead 2)
      AEAD Algorithm: EAX mode (aead 1)
    Preferred Hash Algorithms (sub 21) (5 bytes)
      Hash Algorithm: SHA2-512 (hash 10)
      Hash Algorithm: SHA2-384 (hash 9)
      Hash Algorithm: SHA2-256 (hash 8)
      Hash Algorithm: SHA2-224 (hash 11)
      Hash Algorithm: SHA-1 (hash 2)
    Preferred Compression Algorithms (sub 22) (3 bytes)
      Compression Algorithm: ZLIB <RFC1950> (comp 2)
      Compression Algorithm: BZip2 (comp 3)
      Compression Algorithm: ZIP <RFC1951> (comp 1)
    Features (sub 30) (1 bytes)
      Flag: Modification Detection (packets 18 and 19)
      Flag: Unknown flag1(0x06)
    Key Server Preferences (sub 23) (1 bytes)
      Flag: No-modify
  Hash left 2 bytes
    f5 c0
  EdDSA compressed value r (254 bits)
  EdDSA compressed value s (255 bits)
Secret-Subkey Packet (tag 7) (102 bytes)
  Version: 5 (draft)
  Public-Key
    Public key creation time: 2019-03-20T08:08:04Z
      5c 91 f4 e4
    Public-key Algorithm: ECDH public key algorithm (pub 18)
    ECC Curve OID: cv25519 (256bits key size)
      2b 06 01 04 01 97 55 01 05 01
    ECDH EC point (40 || X) (263 bits)
    KDF parameters (3 bytes)
      Hash Algorithm: SHA2-256 (hash 8)
      Symmetric Algorithm: AES with 128-bit key (sym 7)
  Secret-Key (s2k usage 0; plain secret-key material)
    ECDH secret key (255 bits)
    2-octet checksum
      11 47
Signature Packet (tag 2) (122 bytes)
  Version: 5 (draft)
  Signiture Type: Subkey Binding Signature (0x18)
  Public-key Algorithm: EdDSA (pub 22)
  Hash Algorithm: SHA2-256 (hash 8)
  Hashed Subpacket (44 bytes)
    Issuer Fingerprint (sub 33) (33 bytes)
      Version: 5 (need 32 octets length)
      Fingerprint (32 bytes)
        19 34 7b c9 87 24 64 02 5f 99 df 3e c2 e0 00 0e d9 88 48 92 e1 f7 b3 ea 4c 94 00 91 59 56 9b 54
    Signature Creation Time (sub 2): 2019-03-20T08:08:04Z
    Key Flags (sub 27) (1 bytes)
      Flag: This key may be used to encrypt communications.
      Flag: This key may be used to encrypt storage.
  Hash left 2 bytes
    39 24
  EdDSA compressed value r (256 bits)
  EdDSA compressed value s (256 bits)
{{< /highlight >}}

と表示できるようになった。

よーし，うむうむ，よーし。

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[`gpgpdump`]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

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
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
