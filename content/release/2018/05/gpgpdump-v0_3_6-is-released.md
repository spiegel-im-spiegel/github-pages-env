+++
title = "gpgpdump 0.3.6 をリリースした"
date = "2018-05-16T22:52:22+09:00"
description = "本家 pgpdump のほうの Issue に便乗する感じでバグ修正版を出してみた。"
image = "/images/attention/tools.png"
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

[OpenPGP] パケットの内容を視覚化する [gpgpdump] の 0.3.6 をリリースした。

- [Release v0.3.6 · spiegel-im-spiegel/gpgpdump](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.3.6)

実は本家 [pgpdump] のほうに [Issue が上がって](https://github.com/kazu-yamamoto/pgpdump/issues/23 "pgpdump fails to dump OpenPGP RSA key generated using Bouncycastle · Issue #23 · kazu-yamamoto/pgpdump")いて，件のパケットをうちでも試してみたら出力がダメダメだったのだ。
なので今回はバグ修正版。

問題のパケットはこれ。

```text
-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: BCPG v1.59

lQOqBFr6hWUBCACK+EjkwiEjxEfphBNNgIitKJ40MuuyU57Ss2TE9On3wF2K5gHh
TwQr8gg4wU1lSTslIvG/GA0xY57VBM9MmVfhqC6gyz+sJQCFAD2qYeK65ePwH/I2
mpcSewuRsMKZtFpqOkDsdnaoGU2L7eC1H7PvZfpVLvKm1C16dtl6Oj7pOdMgsDt9
yLv/oMiIguvPdfSsB0F29mb7JIcAjpRz1yw+mP5GoC51P9HrNY3xe58HtIHk3VYx
H7e+vfLYZKhOuKFNVlh0vDi6drkjWFBJSs61dmsVQLu/JwP1UAm7ByMjLlUHUZGg
7v71sXRRvH94URKL3nEiagsctjNfcq/rJrCLABEBAAEAB/sHM9fJmtZquMihpHmu
NSCymV3Tay9YyKTYkgvKnlBckujT8L0dNPRoCc99tGZ8SAjbo4RdSyrll56K261Q
zBTs3sPOwqDUHWxC+Gm5xiUhEai4Jze2bhCMz3B7/VelNroYDca4WrMmXew2eeMT
WT43zct3cJ2Kxvcjlzu/28SGhAOBf/BU/5J9m4ibnyiLzYV2pttSnI8znjR5b92Z
Lj3KdnQJcWd+vFGfayVIaH09qIv+kkKfxJPGdsK32cJ10a3kKDYmFbihTU2r5fGA
QYJW5Jx0Lu6/uovusgWiuQgvj/6uZI4DbcR5e+PhM1sccQ7ObaKSQbbe/RpdSDwZ
FFhJBAC2lf8ZDXnXctBR5IcI/84DNpsFtSv0Y6rweOB9w9DC5WGINiIwbeSb+YBw
2NKk9I/XO+Vp7+lhUI0273eE/6mSg7HYZU9LYrQ0A8RLgbqjN+K4Bs9Ai3pLA9f6
rLxzYI0cMuVYQYZ+Xh8OPVNC8C9+KJp2QNhqZHZXaf5kaDFweQQAwtjIhbAKG+qO
BqcaOE20by6QIH378k/VW798gRU8S/ohsMNGsU8l6K8uDeDrsZYGowy3Naq9bVV/
xAfQHR9nlXOGumfMrWrzl6H2Zfqq6yTZsjBjo9yeL4n26VIgf3KCmp+gHfYY4ftI
q/NGNEtOijFI0z10MTD4Y3xH9zyS0CMD/irjqzF0E1HEwF617Kw8zMbHaozjsYEG
xZvSJtwM3mduyxANASORLGiQcZo9t8TSZBi1YdF3CtVO2stXZY+3rD1aQWSHxbhN
hhGdr8CXZ57Wq363wi4epxVj5y+DE8+W2RTtHUWkRoMNR7MatO9Fu6fzrhJvQKr9
0miAyNtgibaGfyqQGiOR0hKyZUj2qO4egsGBXVK0GXhtcHA6YWxpY2VAd29uZGVy
bGFuZC5saXSJATQEEwEIAB4FAlr6hWUCmy8FiwkIBwIGlQoJCAsCBJYCAwECngEA
CgkQjj26DsTcw7fvVgf/Y88ABzLXTFgLo4EO1z7eBNT4Y4dJ8zyKNkcwJRVbOup1
lVNwMb65IyqnxbTXVssAHMYD9EmHzQfV1B/BlcH3DhYnVNpOoOreW93rTRcZS2Ew
UtACUBZatc+nZCyJ/2cSm862upKkYSWWBhQj7W0wj+UmnGZbbzQFZhO3zZC8OISv
X48gB9DlvRcjTypdNxfY+LQv5r4oebwABRisq9bqnfRsGt3JaoeeGS//EqppST6T
7niM96sp1dxHUKaUAqS0lxCkfUo2vv7pH8REY3t1VTvVe/d2FpMBs77wHMQSKlGZ
ovSdIgpRjJZ3TZQkMgreW31TSySr/DCZ2HwLiIp1TQ==
=8D8m
-----END PGP PRIVATE KEY BLOCK-----
```

中身は秘密鍵のパケットである。
秘密鍵は暗号化せず生のまま格納されていて2バイトのチェックサムで検証できるようになっているのだが，何故かその部分に20バイトの SHA-1 ハッシュ値が入ってるらしい。
本家の [pgpdump] は「パケットのバグなんだからズレた分だけスキップさせればいいぢゃん」って事で決着してるようだが，うちはこんな感じにした。

```text
$ gpgpdump testdata/dump-fail.asc
Secret-Key Packet (tag 5) (938 bytes)
    Version: 4 (current)
    Public-Key
        Public key creation time: 2018-05-15T15:59:49+09:00
            5a fa 85 65
        Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
        RSA public modulus n (2048 bits)
        RSA public encryption exponent e (17 bits)
    Secret-Key (the secret-key data is not encrypted.)
        Version: 4 (current)
        RSA secret exponent d (2043 bits)
        RSA secret prime value p (1024 bits)
        RSA secret prime value q (p < q) (1024 bits)
        RSA u, the multiplicative inverse of p, mod q (1022 bits)
        SHA-1 hash (20 bytes)
            7f 2a 90 1a 23 91 d2 12 b2 65 48 f6 a8 ee 1e 82 c1 81 5d 52
User ID Packet (tag 13) (25 bytes)
    User ID: xmpp:alice@wonderland.lit
Signature Packet (tag 2) (308 bytes)
    Version: 4 (current)
    Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA2-256 (hash 8)
    Hashed Subpacket (30 bytes)
        Signature Creation Time (sub 2): 2018-05-15T15:59:49+09:00
        Key Flags <critical> (sub 27) (1 bytes)
            Flag: This key may be used to certify other keys.
            Flag: This key may be used to sign data.
            Flag: This key may be used to encrypt communications.
            Flag: This key may be used to encrypt storage.
            Flag: This key may be used for authentication.
        Preferred Symmetric Algorithms <critical> (sub 11) (4 bytes)
            Symmetric Algorithm: AES with 256-bit key (sym 9)
            Symmetric Algorithm: AES with 192-bit key (sym 8)
            Symmetric Algorithm: AES with 128-bit key (sym 7)
            Symmetric Algorithm: TripleDES (168 bit key derived from 192) (sym 2)
        Preferred Hash Algorithms <critical> (sub 21) (5 bytes)
            Hash Algorithm: SHA2-512 (hash 10)
            Hash Algorithm: SHA2-384 (hash 9)
            Hash Algorithm: SHA2-256 (hash 8)
            Hash Algorithm: SHA2-224 (hash 11)
            Hash Algorithm: SHA-1 (hash 2)
        Preferred Compression Algorithms <critical> (sub 22) (3 bytes)
            Compression Algorithm: ZLIB <RFC1950> (comp 2)
            Compression Algorithm: BZip2 (comp 3)
            Compression Algorithm: ZIP <RFC1951> (comp 1)
        Features <critical> (sub 30) (1 bytes)
            Flag: Modification Detection (packets 18 and 19)
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x8e3dba0ec4dcc3b7
    Hash left 2 bytes
        ef 56
    RSA signature value m^d mod n (2047 bits)
```

SHA-1 ハッシュと決めつけてるが，表示の文言については今後変わるかもしれない。
まだデバッグの途中だしね。
いやぁ，バッファ操作が楽な [Go 言語]で良かったよ。

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[OpenPGP]: http://openpgp.org/
[Go 言語]: https://golang.org/ "The Go Programming Language"
