+++
title = "gpgpdump 0.2.0 をリリースした"
date =  "2017-11-07T22:35:32+09:00"
description = "OpenPGP パケットの内容を視覚化する gpgpdump の 0.2.0 をリリースした。"
tags = ["tools", "openpgp", "golang", "gpgpdump"]

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

[OpenPGP] パケットの内容を視覚化する [gpgpdump] の 0.2.0 をリリースした。
名前でピンとくる人もいるだろうが，山本和彦さんの [pgpdump] の翻案である[^adpt1]。
特徴は以下のとおり。

[^adpt1]: [Go 言語]で書いた [pgpdump] だから [gpgpdump]。 gpg-pgp-dump という意図もある。

- [Go 言語]用のパッケージおよびコマンドライン・インタフェースを提供
- [TOML] （または [JSON]）フォーマットで出力
- [RFC 4880], [RFC 5581] および [RFC 6637] をサポート
- [Apache License Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

[これまでの 0.1.x]({{< ref "/remark/2016/02/gpgpdump-released.md" >}} "gpgpdump - OpenPGP packet visualizer") から全面的に書き換えた。
一応 60% くらいは元のコードは残ってるかな。
不正パケットで [panic] になる状況はだいぶ減ったと思う。

[Go 言語]のパッケージとして使う場合は以下のようなコードになる。

```go
const openpgpStr = `
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2

iF4EARMIAAYFAlTDCN8ACgkQMfv9qV+7+hg2HwEA6h2iFFuCBv3VrsSf2BREQaT1
T1ZprZqwRPOjiLJg9AwA/ArTwCPz7c2vmxlv7sRlRLUI6CdsOqhuO1KfYXrq7idI
=ZOTN
-----END PGP SIGNATURE-----
`

info, err := gpgpdump.Parse(strings.NewReader(openpgpStr), options.NewOptions())
if err != nil {
	return
}
fmt.Println(info.Packets[0].Value)
// Output:
// Signature Packet (tag 2)

```

コマンドライン・インタフェースは 0.1.x とほぼ変わらず。

```text
$ gpgpdump -h
Usage:
  gpgpdump [flags] [PGPfile]

Flags:
  -a, --armor     accepts ASCII input only
  -h, --help      help for gpgpdump
  -i, --int       dumps multi-precision integers
  -j, --json      output with JSON format
  -l, --literal   dumps literal packets (tag 11)
  -m, --marker    dumps marker packets (tag 10)
  -p, --private   dumps private packets (tag 60-63)
  -u, --utc       output with UTC time
  -v, --version   output version of gpgpdump

$ cat sig
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2

iF4EARMIAAYFAlTDCN8ACgkQMfv9qV+7+hg2HwEA6h2iFFuCBv3VrsSf2BREQaT1
T1ZprZqwRPOjiLJg9AwA/ArTwCPz7c2vmxlv7sRlRLUI6CdsOqhuO1KfYXrq7idI
=ZOTN
-----END PGP SIGNATURE-----

$ cat sig | gpgpdump -u
[[Packet]]
  name = "Packet"
  value = "Signature Packet (tag 2)"
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
    name = "Multi-precision integer"
    note = "ECDSA r (256 bits)"

  [[Packet.Item]]
    name = "Multi-precision integer"
    note = "ECDSA s (252 bits)"

$ cat sig | gpgpdump -j -u
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
          "note": "6 bytes",
          "Item": [
            {
              "name": "Signature Creation Time (sub 2)",
              "value": "2015-01-24T02:52:15Z"
            }
          ]
        },
        {
          "name": "Unhashed Subpacket",
          "note": "10 bytes",
          "Item": [
            {
              "name": "Issuer (sub 16)",
              "value": "0x31fbfda95fbbfa18"
            }
          ]
        },
        {
          "name": "Hash left 2 bytes",
          "dump": "36 1f"
        },
        {
          "name": "Multi-precision integer",
          "note": "ECDSA r (256 bits)"
        },
        {
          "name": "Multi-precision integer",
          "note": "ECDSA s (252 bits)"
        }
      ]
    }
  ]
}
```

残りの TODO はこんな感じかな。

- パケット解析の未テスト部分を埋める（古いフォーマットのパケットのテストどうしよう）
- [Compressed Data Packet (Tag 8)](https://tools.ietf.org/html/rfc4880#section-5.6) が未実装。どうやって実現しようか悩み中
- 実は ECC ([RFC 6637]) がよく分かってない。もしかしたら解釈を間違えているかもしれない
- 最終的には [pgpdump] と同等な出力を目指す

余暇でやってるので，まぁボチボチやります。

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[TOML]: https://github.com/toml-lang/toml "toml-lang/toml: Tom's Obvious, Minimal Language"
[JSON]: https://tools.ietf.org/html/rfc7159 "RFC 7159 - The JavaScript Object Notation (JSON) Data Interchange Format"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 5581]: https://tools.ietf.org/html/rfc5581 "RFC 5581 - The Camellia Cipher in OpenPGP"
[RFC 6637]: https://tools.ietf.org/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[OpenPGP]: http://openpgp.org/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[panic]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
