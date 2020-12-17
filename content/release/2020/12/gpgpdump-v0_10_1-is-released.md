+++
title = "gpgpdump v0.10.1 のリリース【もしかしたらセキュリティ・アップデート】"
date =  "2020-12-17T21:09:49+09:00"
description = "ついでなので golang.org/x/crypto モジュールを更新した。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang", "security", "vulnerability"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.10.1 をリリースした。

- [Release v0.10.1 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.10.1)

今回は，以前に紹介した

- [io.Copy を上限付きで実行する]({{< ref "/golang/copying-with-upper-limit.md" >}})

に絡む修正。

[OpenPGP] のパケット・フォーマットには “Compressed Data Packet” というのがあって，実際に zlib 等で圧縮がかかっている。
これを解凍する際に上限を 1GB までとし，超える場合には，エラーにはしないが

```text {hl_lines=[4]}
$ gpgpdump -f ecompression-bomb.asc 
Compressed Data Packet (tag 8) (149 bytes)
    Compression Algorithm: ZIP <RFC1951> (comp 1)
    Compressed data: <too laege decompressed data> (148 bytes)
```

などとして，圧縮データの解析を行わないようにした。

ちうわけで「もしかしたらセキュリティ・アップデート」みたいな？

あと，ついでなので [golang.org/x/crypto][crypto] モジュールを更新した。

- [golang.org/x/crypto/ssh パッケージのセキュリティ・アップデート]({{< relref "./updated-golang-x-crypto-module.md" >}})

まぁ [gpgpdump] が使ってるのは [golang.org/x/crypto/openpgp](https://pkg.go.dev/golang.org/x/crypto/openpgp "openpgp · pkg.go.dev") パッケージの方なので直接は関係ないんだけどね。
気分だよ，気分。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[Go]: https://golang.org/ "The Go Programming Language"
[crypto]: https://pkg.go.dev/golang.org/x/crypto "crypto · pkg.go.dev"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
