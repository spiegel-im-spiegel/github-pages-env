+++
title = "gpgpdump v0.8.0 をリリースした"
date =  "2020-08-20T11:28:08+09:00"
description = "全面書き直し第2弾（笑） でも元のコードの9割くらいは流用できてるんじゃないのかな。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

貴方は3年前に書いた自分のコードにウンザリしたことはありませんか。
私はあります。

というわけで，全面書き直し第2弾（笑） [OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.8.0 をリリースした。

- [Release v0.8.0 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.8.0)

使い方は以下を参照のこと。

- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

本当はもう少しテストケースを揃えてからリリースしたかったんだけど，[バグを見つけちゃった](https://github.com/spiegel-im-spiegel/gpgpdump/commit/c7caf8c21ac37b9c724ddcd66737e4e45ec1e843)ので，とりあえず出しとこうかと。

具体的な変更点は

1. コマンドライン・ツールとして提供することを前提に構成を変更
2. TOML 形式の出力を廃止

の2つ。

サブパッケージのディレクトリ構成変更が主なので，元のコードの9割くらいは流用できてるんじゃないのかな。
「コマンドライン・ツールとして提供」と書いたが，一応のレイヤ設計はしているので，他の [Go] コードに組み込むことは可能である。

たとえばこんな感じ。

```go
package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/spiegel-im-spiegel/gpgpdump/parse"
    "github.com/spiegel-im-spiegel/gpgpdump/parse/context"
)

const openpgpStr = `
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2

iF4EARMIAAYFAlTDCN8ACgkQMfv9qV+7+hg2HwEA6h2iFFuCBv3VrsSf2BREQaT1
T1ZprZqwRPOjiLJg9AwA/ArTwCPz7c2vmxlv7sRlRLUI6CdsOqhuO1KfYXrq7idI
=ZOTN
-----END PGP SIGNATURE-----
`

func main() {
    p, err := parse.New(
        strings.NewReader(openpgpStr),
        context.New(
            context.Set(context.ARMOR, true),
            context.Set(context.UTC, true),
        ),
    )
    if err != nil {
        fmt.Fprintf(os.Stderr, "%+v", err)
        return
    }
    res, err := p.Parse()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%+v", err)
        return
    }
    fmt.Println(res)
}
```

これを実行すると

```text
$ go run sample.go 
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

となる。

TOML 出力は最初の頃に勢いで組み込んだんだけど，考えてみたら YAML や TOML とかってデータ交換用のフォーマットとしては向いてないよね。
というわけで `--toml` オプションを削除した。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[Go]: https://go.dev/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
