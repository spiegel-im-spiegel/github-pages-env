+++
title = "GNKF: NKF ぽいなにか の v0.2.0 をリリースした"
date =  "2020-10-14T15:13:41+09:00"
description = "このバージョンで BASE64 符号化の機能を追加した。これで base64 や openssl コマンドがない環境でも大丈夫。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "tools", "gnkf", "golang", "encoding", "base64" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

文字列処理の習作 [gnkf] の v0.2.0 をリリースした。

- [Release v0.2.0 · spiegel-im-spiegel/gnkf · GitHub](https://github.com/spiegel-im-spiegel/gnkf/releases/tag/v0.2.0)

このバージョンで BASE64 符号化の機能を追加した。

```text
$ gnkf base64 -h
Encode/Decode BASE64.

Usage:
  gnkf base64 [flags]

Aliases:
  base64, b64

Flags:
  -d, --decode          decode BASE64 string
  -f, --file string     path of input text file
  -u, --for-url         encoding/decoding defined in RFC 4648
  -h, --help            help for base64
  -p, --no-padding      no padding
  -o, --output string   path of output file

Global Flags:
      --debug   for debug
```

オリジナルの nkf はコマンドとオプションがごちゃごちゃしてて分かりにくいので `base64` や `openssl base64` のコマンドの形式に合わせることにした。

```text
$ echo Hello World | gnkf base64
SGVsbG8gV29ybGQK

$ echo SGVsbG8gV29ybGQK | gnkf base64 -d
Hello World
```

これで `base64` や `openssl` コマンドがない環境でも大丈夫（笑）

[gnkf]: https://github.com/spiegel-im-spiegel/gnkf "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang"

## ブックマーク

- [GNKF: Network Kanji Filter by Golang]({{< ref "/release/gnkf.md" >}})

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
