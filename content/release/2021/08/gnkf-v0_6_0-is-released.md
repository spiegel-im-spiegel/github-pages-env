+++
title = "GNKF: NKF ぽいなにか の v0.6.0 をリリースした"
date =  "2021-08-23T23:07:22+09:00"
description = "指定した文字列を BCrypt アルゴリズムで符号化する機能を追加した。悪用しないように（笑）"
image = "/images/attention/tools.png"
tags  = [ "tools", "gnkf", "golang", "hash" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語][Go]における文字列処理の習作 [gnkf] の v0.6.0 をリリースした。

- [Release v0.6.0 · spiegel-im-spiegel/gnkf · GitHub](https://github.com/spiegel-im-spiegel/gnkf/releases/tag/v0.6.0)

このバージョンで指定した文字列を BCrypt アルゴリズムで符号化する機能を追加した。

```text
$ gnkf bcrypt -h
Hash and compare by BCrypt.

Usage:
  gnkf bcrypt [flags] string [string...]

Aliases:
  bcrypt, bc

Flags:
      --compare string   compare to BCrypt hashed string
  -c, --cost int         BCrypt cost (4-31) (default 10)
  -h, --help             help for bcrypt

Global Flags:
      --debug   for debug
```

こんな感じで使う。

```text
$ gnkf bc password
$2a$10$ES0KxMf9p.t0FEMp8WB6we8X43rMzfXb9r5WvFeUSk8Q2z3wdjrCS
```

符号化した文字列を検証することもできる。

```text
$ gnkf bc --compare '$2a$10$ES0KxMf9p.t0FEMp8WB6we8X43rMzfXb9r5WvFeUSk8Q2z3wdjrCS' foo bar password
compare BCrypt hashed string '$2a$10$ES0KxMf9p.t0FEMp8WB6we8X43rMzfXb9r5WvFeUSk8Q2z3wdjrCS' to...
foo : crypto/bcrypt: hashedPassword is not the hash of the given password
bar : crypto/bcrypt: hashedPassword is not the hash of the given password
password : match!
```

いやね。
最初は Java のパスワード処理を検証するミニツールを作ってたんだけど，符号化するところだけチェックすればいいのなら [gnkf] に組み込んじゃえばいいぢゃん，と思いついて「えいやっ」で組み込んでしまった。
反省はしない。

悪用しないように（笑）

## ブックマーク

- [GNKF: Network Kanji Filter by Golang]({{< ref "/release/gnkf.md" >}})

[Go]: https://go.dev/
[gnkf]: https://github.com/spiegel-im-spiegel/gnkf "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
