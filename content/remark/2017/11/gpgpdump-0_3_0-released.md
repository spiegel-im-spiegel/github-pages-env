+++
title = "gpgpdump 0.3.0 をリリースした"
date =  "2017-11-26T19:21:26+09:00"
update =  "2017-12-04T07:44:18+09:00"
description = "OpenPGP パケットの内容を視覚化する gpgpdump の 0.3.0 をリリースした。このバージョンでようやく pgpdump に近い出力ができるようになった。"
image = "/images/attention/remark.jpg"
tags = ["tools", "openpgp", "golang", "gpgpdump"]

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

[OpenPGP] パケットの内容を視覚化する [gpgpdump] の 0.3.0 をリリースした。

- [Release v0.3.0 · spiegel-im-spiegel/gpgpdump](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.3.0)

このバージョンでようやく [pgpdump] に近い出力ができるようになった。

```text
$ cat sig
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2

iF4EARMIAAYFAlTDCN8ACgkQMfv9qV+7+hg2HwEA6h2iFFuCBv3VrsSf2BREQaT1
T1ZprZqwRPOjiLJg9AwA/ArTwCPz7c2vmxlv7sRlRLUI6CdsOqhuO1KfYXrq7idI
=ZOTN
-----END PGP SIGNATURE-----

$ cat sig | gpgpdump -u
Signature Packet (tag 2) (94 bytes)
    Version: 4 (new)
    Signiture Type: Signature of a canonical text document (0x01)
    Public-key Algorithm: ECDSA public key algorithm (pub 19)
    Hash Algorithm: SHA256 (hash 8)
    Hashed Subpacket (6 bytes)
        Signature Creation Time (sub 2): 2015-01-24T02:52:15Z
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x31fbfda95fbbfa18
    Hash left 2 bytes
        36 1f
    ECDSA r (256 bits)
    ECDSA s (252 bits)
```

今までの TOML フォーマット出力を行うには `-t` または `--toml` オプションを付ければよい。

```text
$ cat sig | gpgpdump -t -u
[[Packet]]
  name = "Signature Packet (tag 2)"
  note = "94 bytes"

  [[Packet.Item]]
    name = "Version"
    value = "4"
    note = "new"

  [[Packet.Item]]
    name = "Signiture Type"
    value = "Signature of a canonical text document (0x01)"

  [[Packet.Item]]
    name = "Public-key Algorithm"
    value = "ECDSA public key algorithm (pub 19)"

  [[Packet.Item]]
    name = "Hash Algorithm"
    value = "SHA256 (hash 8)"

  [[Packet.Item]]
    name = "Hashed Subpacket"
    note = "6 bytes"

    [[Packet.Item.Item]]
      name = "Signature Creation Time (sub 2)"
      value = "2015-01-24T02:52:15Z"

  [[Packet.Item]]
    name = "Unhashed Subpacket"
    note = "10 bytes"

    [[Packet.Item.Item]]
      name = "Issuer (sub 16)"
      value = "0x31fbfda95fbbfa18"

  [[Packet.Item]]
    name = "Hash left 2 bytes"
    dump = "36 1f"

  [[Packet.Item]]
    name = "ECDSA r"
    note = "256 bits"

  [[Packet.Item]]
    name = "ECDSA s"
    note = "252 bits"
```

あとは “[Compressed Data Packet (Tag 8)](https://tools.ietf.org/html/rfc4880#section-5.6)” に対応したことか。
これもようやく。

```text
$ gpgpdump -u sig2
Compressed Data Packet (tag 8) (149 bytes)
    Compression Algorithm: ZIP <RFC1951> (comp 1)
        One-Pass Signature Packet (tag 4) (13 bytes)
            Version: 3 (new)
            Signiture Type: Signature of a binary document (0x00)
            Hash Algorithm: SHA256 (hash 8)
            Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
            Key ID: 0xb4da3bae7e20b81c
            Encrypted session key: other than one pass signature (01)
        Literal Data Packet (tag 11) (19 bytes)
            Literal data format: b (binary)
            File name: <null>
            Modification time of a file: 2017-11-25T06:29:56Z
            Literal data (13 bytes)
        Signature Packet (tag 2) (117 bytes)
            Version: 4 (new)
            Signiture Type: Signature of a binary document (0x00)
            Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
            Hash Algorithm: SHA256 (hash 8)
            Hashed Subpacket (29 bytes)
                Issuer Fingerprint (sub 33) (21 bytes)
                    Version: 4 (need 20 octets length)
                    Fingerprint (20 bytes)
                        1b 52 02 db 4a 3e c7 76 f1 e0 ad 18 b4 da 3b ae 7e 20 b8 1c
                Signature Creation Time (sub 2): 2017-11-25T06:29:56Z
            Unhashed Subpacket (10 bytes)
                Issuer (sub 16): 0xb4da3bae7e20b81c
            Hash left 2 bytes
                73 3c
            DSA r (256 bits)
            DSA s (255 bits)
```

[RFC 1951] を [Go 言語]でどうやって実装するのか分からなくてさ。
[`compress/flate`] パッケージを使えばいいと気づくまで試行錯誤しちゃったよ（自作しなくて済んだ）。
この辺は[別の記事]({{< ref "/golang/compress-data.md" >}} "Go 言語でデータ圧縮と解凍")で紹介する。

現在のオプションはこんな感じ。

```text
$ gpgpdump -h
Usage:
  gpgpdump [flags] [OpenPGP file]

Flags:
  -a, --armor     accepts ASCII input only
      --debug     for debug
  -h, --help      help for gpgpdump
  -i, --int       dumps multi-precision integers
  -j, --json      output with JSON format
  -l, --literal   dumps literal packets (tag 11)
  -m, --marker    dumps marker packets (tag 10)
  -p, --private   dumps private packets (tag 60-63)
  -t, --toml      output with TOML format
  -u, --utc       output with UTC time
  -v, --version   output version of gpgpdump
```

`--debug` オプションは主に私がテストするためののもの。
以前は隠し機能にしてたんだけど，正式なオプションとして昇格させた。
これを付けると出力がかなりウザいことになるので，普段はお勧めしない。

残りの TODO はこんな感じ。
欲しい機能はだいたい実装できたかな。

- パケット解析の未テスト部分を埋める（古いフォーマットのパケットのテストどうしよう）
- 実は ECC ([RFC 6637]) がよく分かってない。もしかしたら解釈を間違えているかもしれない

個人的には，これで [pgpdump] から置き換えてもいいかなぁ，ってところまできたかな。
テストが圧倒的に不足してるのでしばらくはデバッグに専念するけど。
でも，ようやく専念できるようになったよ（笑）

## ブックマーク

- [gpgpdump - OpenPGP packet visualizer]({{< ref "/remark/2016/02/gpgpdump-released.md" >}})
- [gpgpdump 0.2.0 をリリースした]({{< ref "/remark/2017/11/gpgpdump-0_2_0-released.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[TOML]: https://github.com/toml-lang/toml "toml-lang/toml: Tom's Obvious, Minimal Language"
[JSON]: https://tools.ietf.org/html/rfc7159 "RFC 7159 - The JavaScript Object Notation (JSON) Data Interchange Format"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 5581]: https://tools.ietf.org/html/rfc5581 "RFC 5581 - The Camellia Cipher in OpenPGP"
[RFC 6637]: https://tools.ietf.org/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[RFC 1951]: https://tools.ietf.org/html/rfc1951 "RFC 1951 - DEFLATE Compressed Data Format Specification version 1.3"
[`compress/flate`]: https://golang.org/pkg/compress/flate/ "flate - The Go Programming Language"
[OpenPGP]: http://openpgp.org/
[Go 言語]: https://golang.org/ "The Go Programming Language"
