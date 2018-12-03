+++
date = "2016-02-07T00:18:50+09:00"
update = "2017-11-07T22:03:13+09:00"
description = "OpenPGP パケットの内容を視覚化する gpgpdump の 0.1.0 をリリースした。"
draft = false
tags = ["tools", "openpgp", "golang", "gpgpdump"]
title = "gpgpdump - OpenPGP packet visualizer"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

余暇でちまちま作っていたが，とりあえず使えるようになったので。

OpenPGP パケットの内容を視覚化する [gpgpdump] の 0.1.0 をリリースした。
名前でピンとくる人もいるだろうが，山本和彦さんの [pgpdump] の翻案である。
特徴は以下のとおり。

- [Go 言語]で作成。特別なパッケージは使用していないので `go get` コマンドのみでビルド可能
- [TOML] （または [JSON]）フォーマットで出力
- [RFC 4880], [RFC 5581] および [RFC 6637] をサポート
- [Apache License Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

たとえば

```
$ cat sig
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2

iF4EARMIAAYFAlTDCN8ACgkQMfv9qV+7+hg2HwEA6h2iFFuCBv3VrsSf2BREQaT1
T1ZprZqwRPOjiLJg9AwA/ArTwCPz7c2vmxlv7sRlRLUI6CdsOqhuO1KfYXrq7idI
=ZOTN
-----END PGP SIGNATURE-----
```

という OpenPGP 署名データがあるとする。
これを [pgpdump] で表示すると

```
$ pgpdump sig
Old: Signature Packet(tag 2)(94 bytes)
        Ver 4 - new
        Sig type - Signature of a canonical text document(0x01).
        Pub alg - Reserved for ECDSA(pub 19)
        Hash alg - SHA256(hash 8)
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Sat Jan 24 11:52:15 東京 (標準時) 2015
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0x31FBFDA95FBBFA18
        Hash left 2 bytes - 36 1f
        Unknown signature(pub 19)
```

となる。
一方， [gpgpdump] の場合は

```
$ gpgpdump sig
[[Packet]]
  name = "Packet"
  value = "Signature Packet (tag 2)"
  note = "94 bytes"

  [[Packet.Item]]
    name = "Version"
    value = "4"
    dump = "04"
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

    [[Packet.Item.Item]]
      name = "Signature Creation Time (sub 2)"
      value = "2015-01-24T11:52:15+09:00"
      dump = "54 c3 08 df"

  [[Packet.Item]]
    name = "Unhashed Subpacket"

    [[Packet.Item.Item]]
      name = "Issuer (sub 16)"
      value = "0x31FBFDA95FBBFA18"

  [[Packet.Item]]
    name = "Hash left 2 bytes"
    dump = "36 1f"

  [[Packet.Item]]
    name = "Multi-precision integer"
    dump = "..."
    note = "ECDSA r (256 bits)"

  [[Packet.Item]]
    name = "Multi-precision integer"
    dump = "..."
    note = "ECDSA s (252 bits)"
```

という感じで同等の内容を [TOML] フォーマットで出力する。
また `-j` オプションを付けると

```
$ gpgpdump -j sig
{
  "Packet": [
    {
      "name": "Packet",
      "value": "Signature Packet (tag 2)",
      "note": "94 bytes",
      "Item": [
        {
          "name": "Version",
          "value": "4",
          "dump": "04",
          "note": "new"
        },
        {
          "name": "Signiture Type",
          "value": "Signature of a canonical text document (0x01)"
        },
        {
          "name": "Public-key Algorithm",
          "value": "ECDSA public key algorithm (pub 19)"
        },
        {
          "name": "Hash Algorithm",
          "value": "SHA256 (hash 8)"
        },
        {
          "name": "Hashed Subpacket",
          "Item": [
            {
              "name": "Signature Creation Time (sub 2)",
              "value": "2015-01-24T11:52:15+09:00",
              "dump": "54 c3 08 df"
            }
          ]
        },
        {
          "name": "Unhashed Subpacket",
          "Item": [
            {
              "name": "Issuer (sub 16)",
              "value": "0x31FBFDA95FBBFA18"
            }
          ]
        },
        {
          "name": "Hash left 2 bytes",
          "dump": "36 1f"
        },
        {
          "name": "Multi-precision integer",
          "dump": "...",
          "note": "ECDSA r (256 bits)"
        },
        {
          "name": "Multi-precision integer",
          "dump": "...",
          "note": "ECDSA s (252 bits)"
        }
      ]
    }
  ]
}
```

という感じに [JSON] 形式で出力する。
だいぶ冗長な表現で申し訳ないが，解析結果を以下の [struct] で正規化している。

```go
//Packets - OpenPGP packets
type Packets struct {
    Packet []*Item
}

//Item - item in Packets
type Item struct {
    Name  string  `toml:"name" json:"name"`
    Value string  `toml:"value,omitempty" json:"value,omitempty"`
    Dump  string  `toml:"dump,omitempty" json:"dump,omitempty"`
    Note  string  `toml:"note,omitempty" json:"note,omitempty"`
    Item  []*Item `toml:"Item,omitempty" json:"Item,omitempty"`
}
```

[gpgpdump] は [Go 言語]の勉強用に作成した。

[`golang.org/x/crypto/openpgp/packet`](https://godoc.org/golang.org/x/crypto/openpgp/packet) というパッケージがあって，これを使えば簡単にできるだろうと思ったのが大間違いで，結局このパッケージで使えたのは [`OpaquePacket`](https://godoc.org/golang.org/x/crypto/openpgp/packet#OpaquePacket) や [`OpaqueSubpacket`](https://godoc.org/golang.org/x/crypto/openpgp/packet#OpaqueSubpacket) くらい。
実際のパケットの解析はゴリゴリとコードを書くはめになった。
いや，これだけでもだいぶ助かったけど。

とはいえ，まだまだ課題はあって

- パケット解析部分のテストが未実装。つか，古いフォーマットのパケットのテストどうしよう
- そもそもパケット解析部分は作りが悪くて，不正なパケットを食わせると簡単に [panic] が起きてしまうので全面的に書きなおす予定
- [Compressed Data Packet (Tag 8)](https://tools.ietf.org/html/rfc4880#section-5.6) が未実装。どうやって実現しようか悩み中
- 実は ECC ([RFC 6637]) がよく分かってない。もしかしたら解釈を間違えているかもしれない
- 最終的には [pgpdump] と同等な出力を目指す

といったあたりを，これからゆっくり手を入れていこうと考えている。

## ブックマーク

- [わかる！ OpenPGP 暗号 — Baldanders.info](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)
- [プログラミング言語 Go — text.Baldanders.info](/golang/)

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: gpgpdump - OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[TOML]: https://github.com/toml-lang/toml "toml-lang/toml: Tom's Obvious, Minimal Language"
[JSON]: https://tools.ietf.org/html/rfc7159 "RFC 7159 - The JavaScript Object Notation (JSON) Data Interchange Format"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 5581]: https://tools.ietf.org/html/rfc5581 "RFC 5581 - The Camellia Cipher in OpenPGP"
[RFC 6637]: https://tools.ietf.org/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[struct]: https://golang.org/ref/spec#Struct_types "Struct types"
[panic]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
