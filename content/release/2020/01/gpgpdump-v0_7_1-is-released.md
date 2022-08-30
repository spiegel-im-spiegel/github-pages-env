+++
title = "gpgpdump v0.7.1 をリリースした"
date =  "2020-01-31T22:49:01+09:00"
description = "変更点は大きく2つ。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.7.1 をリリースした。

- [Release v0.7.0 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.7.0)
- [Release v0.7.1 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.7.1)

いや，すまん。
v0.7.0 をリリースしてすぐ不備に気付いて v0.7.1 でリリースし直したのよ。

変更点は大きく2つ。

ひとつは [RFC 4880bis] で追加された署名サブパケット Sub 37 (Attested Certifications) に対応したこと。
これに伴い起動オプションに `-c` (`--cert`) を追加した。
このオプションを付けて起動すると sub37 の内容を16進数ダンプで表示する。
まぁ，今のところ具体的な実装がないので試しようがないんだけどね。

もうひとつは Literal Data Packet (Tag 11) におけるタイムスタンプについて。
実は [pgpdump/issues/28](https://github.com/kazu-yamamoto/pgpdump/issues/28 "Potential Confusing Output in \"Literal Data Packet\" Dump · Issue #28 · kazu-yamamoto/pgpdump") に

{{< fig-quote type="markdown" title="Potential Confusing Output in Literal Data Packet Dump" link="https://github.com/kazu-yamamoto/pgpdump/issues/28" lang="en" >}}
{{% quote %}}RFC 4880 states the following: "A four-octet number that indicates a date associated with the literal data. Commonly, the date might be the modification date of a file, or the time the packet was created, or a zero that indicates no specific time."; I assume there's no way to distinguish between an actual file creation date vs. the packet creation date, so the label should be open to both interpretations, I guess{{% /quote %}}.
{{< /fig-quote >}}

とあって「なるほど」と思ったのだ。
新しいバージョンの [gpgpdump] では Literal Data Packet は

```text {hl_lines=["12-17"]}
$ gpgpdump -f testdata/comp-sig.asc --indent 2 -l -u
Compressed Data Packet (tag 8) (149 bytes)
  Compression Algorithm: ZIP <RFC1951> (comp 1)
  Compressed data (148 bytes)
    One-Pass Signature Packet (tag 4) (13 bytes)
      Version: 3 (current)
      Signiture Type: Signature of a binary document (0x00)
      Hash Algorithm: SHA2-256 (hash 8)
      Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
      Key ID: 0xb4da3bae7e20b81c
      Encrypted session key: other than one pass signature (flag 0x01)
    Literal Data Packet (tag 11) (19 bytes)
      Literal data format: b (binary)
      File name (0 byte)
      Creation time: 2017-11-25T06:29:56Z
      Literal data (13 bytes)
        48 65 6c 6c 6f 20 77 6f 72 6c 64 0d 0a
    Signature Packet (tag 2) (117 bytes)
      Version: 4 (current)
      Signiture Type: Signature of a binary document (0x00)
      Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
      Hash Algorithm: SHA2-256 (hash 8)
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
      DSA value r (256 bits)
      DSA value s (255 bits)
```

のように展開される。

あとは，まぁ，細々とした修正をしている。

## ブックマーク

- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
