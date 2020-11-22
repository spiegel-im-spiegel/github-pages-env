+++
title = "gpgpdump v0.10.0 をリリースした"
date =  "2020-11-22T13:33:47+09:00"
description = "このバージョンで github および fetch サブコマンドを追加した。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.10.0 をリリースした。

- [Release v0.10.0 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.10.0)

このバージョンで `github` および `fetch` サブコマンドを追加した。

たとえば [GitHub] リポジトリのコミットやタグに

{{< fig-img src="./verified-signature.png" link="./verified-signature.png" width="547" >}}

のような署名情報があるときに

```text
$ gpgpdump github spiegel-im-spiegel --keyid B4DA3BAE7E20B81C -u
Public-Key Packet (tag 6) (1198 bytes)
    Version: 4 (current)
    Public key creation time: 2013-04-28T10:29:43Z
    Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
    DSA p (3072 bits)
    DSA q (q is a prime divisor of p-1) (256 bits)
    DSA g (3070 bits)
    DSA y (= g^x mod p where x is secret) (3067 bits)
...
```

などとすれば，公開鍵の中身が見れる。

ついでに `--raw` オプションを付ければ，公開鍵そのものを取得することもできる。
まぁ [GitHub] に登録してある公開鍵を [GnuPG] にまとめてインポートするなら，シンプルに

```text
$ gpg --fetch-keys https://github.com/spiegel-im-spiegel.gpg
```

とかすればいいのだが（笑）

`fetch` サブコマンドは URL を直接指定して OpenPGP パケット・データを取得できる。
たとえば [JPCERT/CC の公開鍵](https://www.jpcert.or.jp/jpcert-pgp.html "JPCERT コーディネーションセンター PGP公開鍵")なら

```text
$ gpgpdump fetch https://www.jpcert.or.jp/keys/info-0x69ECE048.asc -u
Public-Key Packet (tag 6) (269 bytes)
    Version: 4 (current)
    Public key creation time: 2009-06-02T05:43:57Z
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (2048 bits)
    RSA public encryption exponent e (17 bits)
User ID Packet (tag 13) (29 bytes)
    User ID: JPCERT/CC <info@jpcert.or.jp>
Signature Packet (tag 2) (316 bytes)
    Version: 4 (current)
    Signiture Type: Generic certification of a User ID and Public-Key packet (0x10)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (38 bytes)
        Preferred Symmetric Algorithms (sub 11) (3 bytes)
            Symmetric Algorithm: AES with 256-bit key (sym 9)
            Symmetric Algorithm: CAST5 (128 bit key, as per) (sym 3)
            Symmetric Algorithm: TripleDES (168 bit key derived from 192) (sym 2)
        Key Server Preferences (sub 23) (4 bytes)
            Flag: No-modify
        Key Flags (sub 27) (4 bytes)
            Flag: This key may be used to certify other keys.
            Flag: This key may be used to sign data.
            Flag: This key may be used to encrypt communications.
            Flag: This key may be used to encrypt storage.
            Flag: The private component of this key may be in the possession of more than one person.
        Preferred Compression Algorithms (sub 22) (2 bytes)
            Compression Algorithm: ZLIB <RFC1950> (comp 2)
            Compression Algorithm: ZIP <RFC1951> (comp 1)
        Features (sub 30) (4 bytes)
            Flag: Modification Detection (packets 18 and 19)
        Preferred Hash Algorithms (sub 21) (3 bytes)
            Hash Algorithm: SHA2-256 (hash 8)
            Hash Algorithm: SHA2-384 (hash 9)
            Hash Algorithm: SHA2-512 (hash 10)
        Signature Creation Time (sub 2): 2009-06-16T03:51:22Z
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x317d97a469ece048
    Hash left 2 bytes
        cd 79
    RSA signature value m^d mod n (2047 bits)
Signature Packet (tag 2) (277 bytes)
    Version: 3 (old)
    Hashed material (5 bytes)
        Signiture Type: Generic certification of a User ID and Public-Key packet (0x10)
        Signature creation time: 2009-06-02T05:43:57Z
    Key ID: 0xe7734fa60c7bde12
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hash left 2 bytes
        e9 53
    RSA signature value m^d mod n (2047 bits)
Signature Packet (tag 2) (156 bytes)
    Version: 4 (current)
    Signiture Type: Generic certification of a User ID and Public-Key packet (0x10)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (6 bytes)
        Signature Creation Time (sub 2): 2009-06-15T05:51:27Z
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x8c756b2e2c94d4ed
    Hash left 2 bytes
        35 fd
    RSA signature value m^d mod n (1022 bits)
Public-Subkey Packet (tag 14) (269 bytes)
    Version: 4 (current)
    Public key creation time: 2009-06-02T05:43:57Z
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (2048 bits)
    RSA public encryption exponent e (17 bits)
Signature Packet (tag 2) (577 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (299 bytes)
        Signature Creation Time (sub 2): 2009-06-02T05:43:58Z
        Key Flags (sub 27) (4 bytes)
            Flag: This key may be used to encrypt communications.
            Flag: This key may be used to encrypt storage.
        Embedded Signature (sub 32) (284 bytes)
            Signature Packet (tag 2) (284 bytes)
                Version: 4 (current)
                Signiture Type: Primary Key Binding Signature (0x19)
                Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
                Hash Algorithm: SHA2-256 (hash 8)
                Hashed Subpacket (6 bytes)
                    Signature Creation Time (sub 2): 2009-06-02T05:43:57Z
                Unhashed Subpacket (10 bytes)
                    Issuer (sub 16): 0x09d704b753ba1622
                Hash left 2 bytes
                    71 2d
                RSA signature value m^d mod n (2048 bits)
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x317d97a469ece048
    Hash left 2 bytes
        1d e2
    RSA signature value m^d mod n (2046 bits)
```

てな感じに中身を見ることができる。
うわっ。
電子署名を埋め込んでるよ。

さて，[次](https://zenn.dev/spiegel/scraps/4ced7e0004c6fa83f037 "gpgpdump のネタ帳")はどうするか。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})
- [GitHub に登録した OpenPGP 公開鍵を取り出す](https://zenn.dev/spiegel/articles/20201014-openpgp-pubkey-in-github)

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[Go]: https://golang.org/ "The Go Programming Language"
[GitHub]: https://github.com/ "GitHub"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
