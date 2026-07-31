+++
title = "gpgpdump v0.11.0 をリリースした"
date =  "2020-12-20T19:59:58+09:00"
description = "クリップボードに読み込まれている OpenPGP パケットを直接ダンプできるようにした。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.11.0 をリリースした。

- [Release v0.11.0 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.11.0)

`--clipboard` オプションを追加し，クリップボードに読み込まれている [OpenPGP] パケットを直接ダンプできるようにした。
公開鍵をメール本文に直接貼り付けたり Web ページに `<pre>` 要素で貼り付けたりとか今だにあるようなので，そろそろ私が面倒くさくなってきたのですよ（笑）

たとえば

```text
$ cat testdata/eccsig.asc
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2

iF4EARMIAAYFAlTDCN8ACgkQMfv9qV+7+hg2HwEA6h2iFFuCBv3VrsSf2BREQaT1
T1ZprZqwRPOjiLJg9AwA/ArTwCPz7c2vmxlv7sRlRLUI6CdsOqhuO1KfYXrq7idI
=ZOTN
-----END PGP SIGNATURE-----
```

のようなテキストがクリップボードに読み込まれているとすると

```text
$ gpgpdump --clipboard --utc
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

てな感じで処理できる。

なお `--clipboard` オプションによる入力は ASCII armor テキストのみ受け入れる。
また `--file` オプションと同時に指定できない（エラーになる）。

```text
$ gpgpdump --clipboard -f testdata/eccsig.asc 
Error: cannot set --clipborad and --file options at onece
```

更に UNIX 系の環境では `xclip` または `xsel` コマンドがインストールされていることが前提となる。
macOS では `pbpaste` てのを使うみたいだが，これは標準で入ってるのかな？

とまぁ制約が多いが，悪しからず。

## ブックマーク

- [atotto/clipboard: clipboard for golang](https://github.com/atotto/clipboard)

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[Go]: https://go.dev/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
