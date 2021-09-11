+++
title = "gpgpdump v0.12.5 をリリースした"
date =  "2021-09-11T21:05:29+09:00"
description = "鍵の有効期限日が上手く表示されてなかったので修正した。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.12.2 をリリースした。

- [Release v0.12.5 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.12.5)

実は，この前の [TeX Live 2021 のインストール]({{< ref "/remark/2021/09/install-texlive-in-ubuntu-again.md" >}} "改めて TeX Live を Ubuntu に（APT を使わずに）導入する")時に気が付いたのだが [TUG の公開鍵](http://www.tug.org/texlive/files/texlive.asc)のように有効期限を電子署名を追記する形で更新している場合に新しいほうの有効期限日が上手く表示されてなかった。
ついでに，有効期限日数の表示がおかしなことになっていたので小数点以下を切り上げて表示することにした。

最新バージョンで件の公開鍵を見ると以下のようになる。

```text
$ gpgpdump fetch http://www.tug.org/texlive/files/texlive.asc -u
Public-Key Packet (tag 6) (269 bytes)
    Version: 4 (current)
    Public key creation time: 2016-03-19T01:48:04Z
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (2048 bits)
    RSA public encryption exponent e (17 bits)
User ID Packet (tag 13) (40 bytes)
    User ID: TeX Live Distribution <tex-live@tug.org>
Signature Packet (tag 2) (318 bytes)
    Version: 4 (current)
    Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (40 bytes)
        Signature Creation Time (sub 2): 2016-03-19T01:48:04Z
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to certify other keys.
            Flag: This key may be used to sign data.
        Key Expiration Time (sub 9): 540 days after (2017-09-10T01:48:04Z)
        Preferred Symmetric Algorithms (sub 11) (5 bytes)
            Symmetric Algorithm: AES with 256-bit key (sym 9)
            Symmetric Algorithm: AES with 192-bit key (sym 8)
            Symmetric Algorithm: AES with 128-bit key (sym 7)
            Symmetric Algorithm: CAST5 (128 bit key, as per) (sym 3)
            Symmetric Algorithm: TripleDES (168 bit key derived from 192) (sym 2)
        Preferred Hash Algorithms (sub 21) (5 bytes)
            Hash Algorithm: SHA2-256 (hash 8)
            Hash Algorithm: SHA-1 (hash 2)
            Hash Algorithm: SHA2-384 (hash 9)
            Hash Algorithm: SHA2-512 (hash 10)
            Hash Algorithm: SHA2-224 (hash 11)
        Preferred Compression Algorithms (sub 22) (3 bytes)
            Compression Algorithm: ZLIB <RFC1950> (comp 2)
            Compression Algorithm: BZip2 (comp 3)
            Compression Algorithm: ZIP <RFC1951> (comp 1)
        Features (sub 30) (1 bytes)
            Flag: Modification Detection (packets 18 and 19)
        Key Server Preferences (sub 23) (1 bytes)
            Flag: No-modify
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x0d5e5d9106bab6bc
    Hash left 2 bytes
        89 fd
    RSA signature value m^d mod n (2047 bits)
Signature Packet (tag 2) (312 bytes)
    Version: 4 (current)
    Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (34 bytes)
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to certify other keys.
            Flag: This key may be used to sign data.
        Preferred Symmetric Algorithms (sub 11) (5 bytes)
            Symmetric Algorithm: AES with 256-bit key (sym 9)
            Symmetric Algorithm: AES with 192-bit key (sym 8)
            Symmetric Algorithm: AES with 128-bit key (sym 7)
            Symmetric Algorithm: CAST5 (128 bit key, as per) (sym 3)
            Symmetric Algorithm: TripleDES (168 bit key derived from 192) (sym 2)
        Preferred Hash Algorithms (sub 21) (5 bytes)
            Hash Algorithm: SHA2-256 (hash 8)
            Hash Algorithm: SHA-1 (hash 2)
            Hash Algorithm: SHA2-384 (hash 9)
            Hash Algorithm: SHA2-512 (hash 10)
            Hash Algorithm: SHA2-224 (hash 11)
        Preferred Compression Algorithms (sub 22) (3 bytes)
            Compression Algorithm: ZLIB <RFC1950> (comp 2)
            Compression Algorithm: BZip2 (comp 3)
            Compression Algorithm: ZIP <RFC1951> (comp 1)
        Features (sub 30) (1 bytes)
            Flag: Modification Detection (packets 18 and 19)
        Key Server Preferences (sub 23) (1 bytes)
            Flag: No-modify
        Signature Creation Time (sub 2): 2016-03-19T04:07:41Z
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x0d5e5d9106bab6bc
    Hash left 2 bytes
        01 d6
    RSA signature value m^d mod n (2048 bits)
Signature Packet (tag 2) (540 bytes)
    Version: 4 (current)
    Signiture Type: Generic certification of a User ID and Public-Key packet (0x10)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA2-512 (hash 10)
    Hashed Subpacket (6 bytes)
        Signature Creation Time (sub 2): 2016-03-20T01:57:31Z
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x6caca448860cdc13
    Hash left 2 bytes
        5a ec
    RSA signature value m^d mod n (4096 bits)
Signature Packet (tag 2) (284 bytes)
    Version: 4 (current)
    Signiture Type: Generic certification of a User ID and Public-Key packet (0x10)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (6 bytes)
        Signature Creation Time (sub 2): 2016-03-20T23:53:54Z
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x0716748a30d155ad
    Hash left 2 bytes
        21 c4
    RSA signature value m^d mod n (2047 bits)
Public-Subkey Packet (tag 14) (269 bytes)
    Version: 4 (current)
    Public key creation time: 2016-03-19T01:48:04Z
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (2048 bits)
    RSA public encryption exponent e (17 bits)
Signature Packet (tag 2) (293 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (15 bytes)
        Signature Creation Time (sub 2): 2016-03-19T01:48:04Z
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to encrypt communications.
            Flag: This key may be used to encrypt storage.
        Key Expiration Time (sub 9): 540 days after (2017-09-10T01:48:04Z)
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x0d5e5d9106bab6bc
    Hash left 2 bytes
        a3 a8
    RSA signature value m^d mod n (2047 bits)
Signature Packet (tag 2) (287 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (9 bytes)
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to encrypt communications.
            Flag: This key may be used to encrypt storage.
        Signature Creation Time (sub 2): 2016-03-19T04:09:15Z
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x0d5e5d9106bab6bc
    Hash left 2 bytes
        d7 48
    RSA signature value m^d mod n (2047 bits)
Public-Subkey Packet (tag 14) (269 bytes)
    Version: 4 (current)
    Public key creation time: 2016-03-19T01:49:00Z
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (2048 bits)
    RSA public encryption exponent e (17 bits)
Signature Packet (tag 2) (580 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (15 bytes)
        Signature Creation Time (sub 2): 2016-03-19T01:49:00Z
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to sign data.
        Key Expiration Time (sub 9): 540 days after (2017-09-10T01:49:00Z)
    Unhashed Subpacket (297 bytes)
        Issuer (sub 16): 0x0d5e5d9106bab6bc
        Embedded Signature (sub 32) (284 bytes)
            Signature Packet (tag 2) (284 bytes)
                Version: 4 (current)
                Signiture Type: Primary Key Binding Signature (0x19)
                Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
                Hash Algorithm: SHA-1 (hash 2)
                Hashed Subpacket (6 bytes)
                    Signature Creation Time (sub 2): 2016-03-19T01:49:00Z
                Unhashed Subpacket (10 bytes)
                    Issuer (sub 16): 0x4ce1877e19438c70
                Hash left 2 bytes
                    f3 8d
                RSA signature value m^d mod n (2046 bits)
    Hash left 2 bytes
        91 40
    RSA signature value m^d mod n (2047 bits)
Signature Packet (tag 2) (603 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (38 bytes)
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to sign data.
        Issuer Fingerprint (sub 33) (21 bytes)
            Version: 4 (need 20 octets length)
            Fingerprint (20 bytes)
                c7 8b 82 d8 c7 95 12 f7 9c c0 d7 c8 0d 5e 5d 91 06 ba b6 bc
        Signature Creation Time (sub 2): 2017-09-06T23:56:27Z
        Key Expiration Time (sub 9): 902 days after (2018-09-06T23:56:27Z)
    Unhashed Subpacket (297 bytes)
        Embedded Signature (sub 32) (284 bytes)
            Signature Packet (tag 2) (284 bytes)
                Version: 4 (current)
                Signiture Type: Primary Key Binding Signature (0x19)
                Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
                Hash Algorithm: SHA-1 (hash 2)
                Hashed Subpacket (6 bytes)
                    Signature Creation Time (sub 2): 2016-03-19T01:49:00Z
                Unhashed Subpacket (10 bytes)
                    Issuer (sub 16): 0x4ce1877e19438c70
                Hash left 2 bytes
                    f3 8d
                RSA signature value m^d mod n (2046 bits)
        Issuer (sub 16): 0x0d5e5d9106bab6bc
    Hash left 2 bytes
        0d 16
    RSA signature value m^d mod n (2047 bits)
Signature Packet (tag 2) (603 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (38 bytes)
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to sign data.
        Issuer Fingerprint (sub 33) (21 bytes)
            Version: 4 (need 20 octets length)
            Fingerprint (20 bytes)
                c7 8b 82 d8 c7 95 12 f7 9c c0 d7 c8 0d 5e 5d 91 06 ba b6 bc
        Signature Creation Time (sub 2): 2018-08-31T16:19:24Z
        Key Expiration Time (sub 9): 1261 days after (2019-08-31T16:19:12Z)
    Unhashed Subpacket (297 bytes)
        Embedded Signature (sub 32) (284 bytes)
            Signature Packet (tag 2) (284 bytes)
                Version: 4 (current)
                Signiture Type: Primary Key Binding Signature (0x19)
                Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
                Hash Algorithm: SHA-1 (hash 2)
                Hashed Subpacket (6 bytes)
                    Signature Creation Time (sub 2): 2016-03-19T01:49:00Z
                Unhashed Subpacket (10 bytes)
                    Issuer (sub 16): 0x4ce1877e19438c70
                Hash left 2 bytes
                    f3 8d
                RSA signature value m^d mod n (2046 bits)
        Issuer (sub 16): 0x0d5e5d9106bab6bc
    Hash left 2 bytes
        d4 df
    RSA signature value m^d mod n (2044 bits)
Signature Packet (tag 2) (603 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (38 bytes)
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to sign data.
        Issuer Fingerprint (sub 33) (21 bytes)
            Version: 4 (need 20 octets length)
            Fingerprint (20 bytes)
                c7 8b 82 d8 c7 95 12 f7 9c c0 d7 c8 0d 5e 5d 91 06 ba b6 bc
        Signature Creation Time (sub 2): 2019-03-13T00:38:24Z
        Key Expiration Time (sub 9): 1479 days after (2020-04-06T00:38:24Z)
    Unhashed Subpacket (297 bytes)
        Embedded Signature (sub 32) (284 bytes)
            Signature Packet (tag 2) (284 bytes)
                Version: 4 (current)
                Signiture Type: Primary Key Binding Signature (0x19)
                Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
                Hash Algorithm: SHA-1 (hash 2)
                Hashed Subpacket (6 bytes)
                    Signature Creation Time (sub 2): 2016-03-19T01:49:00Z
                Unhashed Subpacket (10 bytes)
                    Issuer (sub 16): 0x4ce1877e19438c70
                Hash left 2 bytes
                    f3 8d
                RSA signature value m^d mod n (2046 bits)
        Issuer (sub 16): 0x0d5e5d9106bab6bc
    Hash left 2 bytes
        0f 14
    RSA signature value m^d mod n (2047 bits)
Signature Packet (tag 2) (603 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA2-512 (hash 10)
    Hashed Subpacket (38 bytes)
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to sign data.
        Issuer Fingerprint (sub 33) (21 bytes)
            Version: 4 (need 20 octets length)
            Fingerprint (20 bytes)
                c7 8b 82 d8 c7 95 12 f7 9c c0 d7 c8 0d 5e 5d 91 06 ba b6 bc
        Signature Creation Time (sub 2): 2020-04-06T02:40:10Z
        Key Expiration Time (sub 9): 1960 days after (2021-07-30T02:40:10Z)
    Unhashed Subpacket (297 bytes)
        Embedded Signature (sub 32) (284 bytes)
            Signature Packet (tag 2) (284 bytes)
                Version: 4 (current)
                Signiture Type: Primary Key Binding Signature (0x19)
                Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
                Hash Algorithm: SHA-1 (hash 2)
                Hashed Subpacket (6 bytes)
                    Signature Creation Time (sub 2): 2016-03-19T01:49:00Z
                Unhashed Subpacket (10 bytes)
                    Issuer (sub 16): 0x4ce1877e19438c70
                Hash left 2 bytes
                    f3 8d
                RSA signature value m^d mod n (2046 bits)
        Issuer (sub 16): 0x0d5e5d9106bab6bc
    Hash left 2 bytes
        69 4e
    RSA signature value m^d mod n (2047 bits)
Signature Packet (tag 2) (603 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA2-512 (hash 10)
    Hashed Subpacket (38 bytes)
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to sign data.
        Issuer Fingerprint (sub 33) (21 bytes)
            Version: 4 (need 20 octets length)
            Fingerprint (20 bytes)
                c7 8b 82 d8 c7 95 12 f7 9c c0 d7 c8 0d 5e 5d 91 06 ba b6 bc
        Signature Creation Time (sub 2): 2021-07-02T22:45:27Z
        Key Expiration Time (sub 9): 2322 days after (2022-07-27T22:45:27Z)
    Unhashed Subpacket (297 bytes)
        Embedded Signature (sub 32) (284 bytes)
            Signature Packet (tag 2) (284 bytes)
                Version: 4 (current)
                Signiture Type: Primary Key Binding Signature (0x19)
                Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
                Hash Algorithm: SHA-1 (hash 2)
                Hashed Subpacket (6 bytes)
                    Signature Creation Time (sub 2): 2016-03-19T01:49:00Z
                Unhashed Subpacket (10 bytes)
                    Issuer (sub 16): 0x4ce1877e19438c70
                Hash left 2 bytes
                    f3 8d
                RSA signature value m^d mod n (2046 bits)
        Issuer (sub 16): 0x0d5e5d9106bab6bc
    Hash left 2 bytes
        89 0c
    RSA signature value m^d mod n (2048 bits)
```

長ったらしくなって申し訳ないが，毎年1年ずつ有効期限を延長しているのが分かる。
私もこういう鍵運用にしようかなぁ。
やっぱ鍵そのものを度々替えるのはダメだよねぇ。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[Go]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
