+++
title = "OpenPGP パケットを可視化する gpgpdump"
date =  "2019-09-10T22:07:44+09:00"
description = "gpgpdump は OpenPGP パッケットの内容を human-readable な形式で可視化する CLI ツールである。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "golang", "gpgpdump", "pgpdump"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [goark/gpgpdump: OpenPGP packet visualizer][gpgpdump]

[gpgpdump] は [OpenPGP] パッケットの内容を human-readable な形式で可視化するコマンドライン・ツールである。
[山本和彦](http://www.mew.org/~kazu/)さんによる [pgpdump] を参考デザインとし， [Go 言語]で組み直している。

[gpgpdump] は [pgpdump] と比較して以下の特徴がある。

- 平文テキストによる結果出力のほか [JSON] 形式による出力もできる
- 現行仕様である [RFC 4880] に追加して [RFC 5581] および [RFC 6637] にも対応している
- 次期 [OpenPGP] ドラフト案である [RFC 4880bis] にも一部対応している
- [HKP (HTTP Keyserver Protocol)][HKP] を用いて [OpenPGP] 鍵サーバから直接公開鍵を取得して検証できる
- [GitHub] に登録している OpenPGP 公開鍵を直接取得して検証できる

[![check vulns](https://github.com/goark/gpgpdump/workflows/vulns/badge.svg)](https://github.com/goark/gpgpdump/actions)
[![lint status](https://github.com/goark/gpgpdump/workflows/lint/badge.svg)](https://github.com/goark/gpgpdump/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/gpgpdump/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/gpgpdump.svg)](https://github.com/goark/gpgpdump/releases/latest)

## ダウンロードとビルド

[gpgpdump] は以下の [Go] コマンドでダウンロードとビルドができる。

```
$ go get github.com/goark/gpgpdump@latest
```

なおビルドには [Go] 1.13 以降が要件となるのでご注意を。

64ビット版のみであるが，各プラットフォーム用のバイナリも用意している。
[最新バイナリはリリースページから取得](https://github.com/goark/gpgpdump/releases/latest)できる。

## 簡単な使い方

`-h` オプションで簡単なヘルプを表示できる。

```text
$ gpgpdump -h
OpenPGP (RFC 4880) packet visualizer by golang.

Usage:
  gpgpdump [flags]
  gpgpdump [command]

Available Commands:
  completion  Generate completion script
  fetch       Dumps OpenPGP packets form the Web
  github      Dumps OpenPGP keys registered on GitHub
  help        Help about any command
  hkp         Dumps OpenPGP packets from the key server
  version     Print the version number

Flags:
  -a, --armor         accepts ASCII armor text only
  -c, --cert          dumps attested certification in signature packets (tag 2)
      --clipboard     input from clipboard (ASCII armor text only)
      --debug         for debug
  -f, --file string   path of OpenPGP file
  -h, --help          help for gpgpdump
      --indent int    indent size for output text
  -i, --int           dumps multi-precision integers
  -j, --json          output with JSON format
  -l, --literal       dumps literal packets (tag 11)
  -m, --marker        dumps marker packets (tag 10)
  -p, --private       dumps private packets (tag 60-63)
  -u, --utc           output with UTC time
  -v, --version       output version of gpgpdump

Use "gpgpdump [command] --help" for more information about a command.
```

たとえば以下のような [OpenPGP] 電子署名データファイルがあるとする。

```text
$ cat testdata/eccsig.asc
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2

iF4EARMIAAYFAlTDCN8ACgkQMfv9qV+7+hg2HwEA6h2iFFuCBv3VrsSf2BREQaT1
T1ZprZqwRPOjiLJg9AwA/ArTwCPz7c2vmxlv7sRlRLUI6CdsOqhuO1KfYXrq7idI
=ZOTN
-----END PGP SIGNATURE-----
```

これを [gpgpdump] で表示するとこんな感じの出力になる。

```text
$ gpgpdump -u -f testdata/eccsig.asc
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

パイプ等を使って標準入力からの入力も可能である。

```text
$ cat testdata/eccsig.asc | gpgpdump -u
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

`-j` オプションを指定すれば JSON フォーマットで出力される[^jq1]。

[^jq1]: [jq](https://stedolan.github.io/jq/) は JSON 整形コマンドで，データ構造の一部を抽出することもできる。

```text
$ cat testdata/eccsig.asc | gpgpdump -u -j | jq .
{
  "Packet": [
    {
      "name": "Signature Packet (tag 2)",
      "note": "94 bytes",
      "Item": [
        {
          "name": "Version",
          "value": "4",
          "note": "current"
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
          "value": "SHA2-256 (hash 8)"
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
          "name": "ECDSA value r",
          "note": "256 bits"
        },
        {
          "name": "ECDSA value s",
          "note": "252 bits"
        }
      ]
    }
  ]
}
```

[gpgpdump] で使える主なオプションは以下の通り。

| オプション名 | 短縮名 | 内容                                                        |
| ------------ |:------:| ----------------------------------------------------------- |
| `armor`      |  `a`   | ASCII armor テキストのみ受け入れる                          |
| `cert`       |  `c`   | 署名パケット（tag 2）内の証明された認証を16進ダンプ表示する |
| `int`        |  `i`   | MPI データを16進ダンプ表示する                              |
| `literal`    |  `l`   | リテラル・パケット（tag 11）を16進ダンプ表示する            |
| `marker`     |  `m`   | マーカー・パケット（tag 10）を16進ダンプ表示する            |
| `private`    |  `p`   | プライベート用パケット（tag 60-63）を16進ダンプ表示する     |
| `utc`        |  `u`   | 時刻を UTC で表示する                                       |
| `file`       |  `f`   | [OpenPGP] データファイルのパスを指定する                    |
| `json`       |  `j`   | [JSON] 形式で出力する                                       |
| `indent`     |        | 平文テキスト出力時のインデント幅を指定する                  |
| `debug`      |        | デバッグ用                                                  |

## HKP モード

```text
$ gpgpdump hkp -h
Dumps OpenPGP packets from the key server.

Usage:
  gpgpdump hkp [flags] {userID | keyID}

Aliases:
  hkp, h

Flags:
  -h, --help               help for hkp
      --keyserver string   OpenPGP key server (default "keys.gnupg.net")
      --port int           port number of OpenPGP key server (default 11371)
      --raw                output raw text from OpenPGP key server
      --secure             enable HKP over HTTPS

Global Flags:
  -a, --armor        accepts ASCII armor text only
  -c, --cert         dumps attested certification in signature packets (tag 2)
      --debug        for debug
      --indent int   indent size for output text
  -i, --int          dumps multi-precision integers
  -j, --json         output with JSON format
  -l, --literal      dumps literal packets (tag 11)
  -m, --marker       dumps marker packets (tag 10)
  -p, --private      dumps private packets (tag 60-63)
  -u, --utc          output with UTC time
```

`hkp` サブコマンドを指定することで [OpenPGP] 鍵サーバから [HKP] により直接公開鍵を取得して中身を検証することができる。

```text
$ gpgpdump hkp 0x44ce6900e2b307a4 -u
Public-Key Packet (tag 6) (269 bytes)
    Version: 4 (current)
    Public key creation time: 2009-11-08T15:20:55Z
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (2048 bits)
    RSA public encryption exponent e (17 bits)
User ID Packet (tag 13) (25 bytes)
    User ID: Alice <alice@example.com>
...
```

以下のオプションを使って [OpenPGP] 鍵サーバを指定可能。

| オプション名 | 既定値           | 内容                             |
| ------------ | ---------------- | -------------------------------- |
| `keyserver`  | `keys.gnupg.net` | [OpenPGP] 鍵サーバ名（ホスト名） |
| `port`       | 11371            | [OpenPGP] 鍵サーバのポート番号   |
| `secure`     | false            | [HKP] over HTTPS を有効にする      |

さらに `--raw` オプションを使うとダンプ表示はせず，鍵サーバから取得したデータ（ASCII armor テキスト）をそのまま表示する。

```text
$ gpgpdump hkp --raw 0x44ce6900e2b307a4
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: SKS 1.1.6
Comment: Hostname: sks.pod02.fleetstreetops.com

mQENBEr24dcBCADQeCxUo1pNF33ytHuzLn4vK9Z8LWXCUoZsQAZ9+cMKAzbQ9ncO+LfMleDz
RpjsBxYWDaTnn6a8OySveDcw9/CZ9Wu0ND0+uHErdNk5qh+z81x15sOAfN9xj4pUm0iH092Z
wuILrLjWWqgKMZYmB8HKaHXDkQmSfQmhx7oyZ4tWHfMN/VqBWLyUt0RaU0X+s4zLrdJSsTaf
ECZRo/2OJecpyBzLBc45Tzv3RJAXTyv31MLDYn38bS0EiShRoqaGIZthC7ZnX9EoaS2trg1K
uZtv6NeScRU4TqS21q/kYnE6HBnAMg7mI7dtFbg8x20TB2rTA5v8o/8cqZ3MLQukqjZ1ABEB
AAG0GUFsaWNlIDxhbGljZUBleGFtcGxlLmNvbT6IjAQQFggANBYhBDvMx8/SWX5TRN2WSnKb
Uj0R86jXBQJe8BASFhSAAAAAAA0AAHJlbUBnbnVwZy5vcmcACgkQcptSPRHzqNexFQEA3wFC
8PN9jOyFJak06/OWplZpQCMvBEBKJl+hJZYLNdIBAPEZay004L/HD0CA6O8l9emQyDCglYkT
y2AIzzpeFvABiQE4BBMBAgAiBQJK9uHXAhsDBgsJCAcDAgYVCAIJCgsEFgIDAQIeAQIXgAAK
CRBEzmkA4rMHpJNiB/0bWxus4CYj1fdRmRTIJmSVNiuqrohX3c1DZry8j34v5fEFLsAwHPL8
54uw2FhQVgjhZN75vKPzghsZoUh7dbtC+KZuACnAqmsHV/Nx9D0Ac7x8tVEvt90glG9jkp8Q
SMLA8SElUmfPQoXvjugc93ZdZs7A6J8Nxxlcu9zsrKQqH+60aTIUs03F5D/PaQFeZOCFkoOt
dj5QTV8Kwkow4nnMdQ55dJCnD1Ze7RFZmMEqd+jAQ6N3Vg41f6+qsmBew+t7aqC30tWpVw+s
6XSdIkbFfLN8yPiRARn1r8U2ZzsLDs1O6ftdcBNaQTOnl/4zXNu+R1skFwWfDML/xkcW3pVx
uQENBEr24dcBCADZ+26/F9bLQ92XPiCeCwPG2rwzg2o4a5kHkpX9lR6HLwDKbHpXZIjyEIFR
eu1oefIGPmnlpdVuCh8ulaE7574vU3fEg6B/QoSTVz6mAKeLuMjx0qth02Gots/U/sixx/Nn
V5epDVuR/exH6egunpzDvEg+UD6Rkib86LIL8CmQXq38ZZVfd/Px0rObF7YyUbWUidqKW6+l
2lj/X6svQdx3B65PWGwBuoxv3j4eLP7zJouyGf7h+2H9v3eIcPTyTOgiT50YwztdalMFllkP
71Lf3F/6L5zo4yZ4/YMyCkHV+2LevwjnDpjGSHyXQ603WPjQtOwybcgCkV0N+KuYwvinABEB
AAGJAR8EGAECAAkFAkr24dcCGwwACgkQRM5pAOKzB6Rm8wgAyw09KTy98Y1lMvuS4sr+IHPt
RCLg8/JKHfNut8N+W7z9B+7q1bB9XDg1KxtYhuEtwVr393xfY8v9f1saUUt4bh8hwrciU5gt
cGP7MhgQT5xztAkM1JChtT2+SKRF4NU9El4oUao4lbTU/jB8ZSF2jJAhUnpcHgCSzPzsqyye
XMGK/CQPmGRWr96c9L2efBZvV0aLrziI/ftIl+TJ68P9oV0RwjoA6z6dWgwMfaZ8GN6MsuXQ
KH/KNBqYms9+E0UXWd70hQZaKKV3Tch/ldnAVMSwMaKmkD6zFvULJqkcP+QNWIqxLqc7He8/
P1FRChSHTwsazDVjP0AKzttqhGYxNw==
=QFnp
-----END PGP PUBLIC KEY BLOCK-----
```

## [GitHub] に登録されている OpenPGP 公開鍵を検証する

```text
$ gpgpdump github -h
Dumps OpenPGP keys registered on GitHub.

Usage:
  gpgpdump github [flags] GitHubUserID

Aliases:
  github, gh, g

Flags:
  -h, --help           help for github
      --keyid string   OpenPGP key ID
      --raw            output raw text (ASCII armor text)

Global Flags:
  -a, --armor        accepts ASCII armor text only
  -c, --cert         dumps attested certification in signature packets (tag 2)
      --debug        for debug
      --indent int   indent size for output text
  -i, --int          dumps multi-precision integers
  -j, --json         output with JSON format
  -l, --literal      dumps literal packets (tag 11)
  -m, --marker       dumps marker packets (tag 10)
  -p, --private      dumps private packets (tag 60-63)
  -u, --utc          output with UTC time
```

[GitHub] はユーザごとに OpenPGP 公開鍵を登録することができ，コミット等に付与された電子署名を検証することができる。
`github` サブコマンドでこの公開鍵を [GitHub] から直接取得して中身を検証することができる。

```text
$ gpgpdump github spiegel-im-spiegel -u
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

複数の OpenPGP 公開鍵を登録している場合は，全ての鍵が連結された状態で取得される。
特定の公開鍵だけを見たい場合は `--keyid` オプションを使って

```text
$ gpgpdump github spiegel-im-spiegel -u --keyid 3b460ba9a59048c9
Public-Key Packet (tag 6) (51 bytes)
    Version: 4 (current)
    Public key creation time: 2020-10-27T06:20:19Z
    Public-key Algorithm: EdDSA (pub 22)
    ECC Curve OID: ed25519 (256bits key size)
        2b 06 01 04 01 da 47 0f 01
    EdDSA EC point (Native point format of the curve follows) (263 bits)
...
```

などとすればよい。

`--raw` オプションを使うとダンプ表示はせず， [GitHub] から取得したデータ（ASCII armor テキスト）をそのまま表示する。

```text
$ gpgpdump github spiegel-im-spiegel --keyid 3b460ba9a59048c9 --raw
-----BEGIN PGP PUBLIC KEY BLOCK-----

mDMEX5e8IxYJKwYBBAHaRw8BAQdAsbKSTsqzMQSUdWv/UQBfYf+HD9/+dzCT+zVN
IvZRQ/e0L1NwaWVnZWwgKGdpdGh1YikgPHNwaWVnZWwuaW0uc3BpZWdlbEBnbWFp
bC5jb20+iJAEExYIADgWIQRK8EwUs6FchRaOGmc7RguppZBIyQUCX5e8IwIbAwUL
CQgHAgYVCgkICwIEFgIDAQIeAQIXgAAKCRA7RguppZBIyVTWAPoChzM/FBkpqBVl
Uv2v3cE3UIfPENm5qWegUmmBg3yTLgD+MXyI6zter+DDUT9e0UWeaOZrsJvCpBUh
9eLxUMpiKg64OARfl7wjEgorBgEEAZdVAQUBAQdALSF4rszFFAhisXD/7ZKytW1M
LSWGPxW4sLNezkhKfX0DAQgHiHgEGBYIACAWIQRK8EwUs6FchRaOGmc7RguppZBI
yQUCX5e8IwIbDAAKCRA7RguppZBIyUzLAQCDNMwKtq1zR2KHm9F0Fvp66vv2grX1
j0TRb5sEayG9YQEA5ap7JltKc4iVThzu7bY1D+sZ3KbXetecf+QQv6YrwAE=
=n3Jr
-----END PGP PUBLIC KEY BLOCK-----
```

## 任意の URL を指定して Web 上にある OpenPGP パケット・データを検証する

```text
$ gpgpdump fetch -h
Dumps OpenPGP packets form the Web.

Usage:
  gpgpdump fetch [flags] URL

Aliases:
  fetch, fch, f

Flags:
  -h, --help   help for fetch
      --raw    output raw data

Global Flags:
  -a, --armor        accepts ASCII armor text only
  -c, --cert         dumps attested certification in signature packets (tag 2)
      --debug        for debug
      --indent int   indent size for output text
  -i, --int          dumps multi-precision integers
  -j, --json         output with JSON format
  -l, --literal      dumps literal packets (tag 11)
  -m, --marker       dumps marker packets (tag 10)
  -p, --private      dumps private packets (tag 60-63)
  -u, --utc          output with UTC time
```

`fetch` サブコマンドで Web 上にある OpenPGP パケット・データ（公開鍵や電子署名など）を直接取得して中身を検証することができる。

```text
$ gpgpdump fetch https://github.com/spiegel-im-spiegel.gpg -u
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

`--raw` オプションを使うとダンプ表示はせず， 取得したデータをそのまま標準出力に出力する。

```text
$ gpgpdump fetch https://github.com/spiegel-im-spiegel.gpg --raw
```

中身がテキストとは限らないのでご注意を。

## [gpgpdump] を [Go] パッケージとして使う

使いどころが思い浮かばないのだが，一応 [gpgpdump] は [Go 言語]用のパッケージまたはモジュールとして組み込むことができる。

たとえばこんな感じ。

```go
package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/goark/gpgpdump/parse"
    "github.com/goark/gpgpdump/parse/context"
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
        context.New(
            context.Set(context.ARMOR, true),
            context.Set(context.UTC, true),
        ),
        strings.NewReader(openpgpStr),
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

これを実行すると以下のような出力を得られる。

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

`context.New()` 関数で CLI 版と同等のオプションを設定できる。
`context.New()` 関数の引数として `context.Set()` 関数を0個以上複数指定できる。
指定可能なオプションは以下の通り。

| オプション        | 既定値 | 対応する CLI オプション |
| ----------------- |:------:| ----------------------- |
| `context.ARMOR`   | false  | `armor`                 |
| `context.CERT`    | false  | `cert`                  |
| `context.INTEGER` | false  | `int`                   |
| `context.LITERAL` | false  | `literal`               |
| `context.MARKER`  | false  | `marker`                |
| `context.PRIVATE` | false  | `private`               |
| `context.UTC`     | false  | `utc`                   |
| `context.DEBUG`   | false  | `debug`                 |

詳しくはコードを見てね♡ ってことで（笑）

## 【おまけ】 各 Shell 用の自動補完スクリプトを生成する

[gpgpdump] では bash などの shell 用の自動補完スクリプトを生成できる。
対応している shell は bash, zsh, fish, PowerShell の4つ。
自動補完の設定方法は

```text
$ gpgpdump completion -h
```

で表示されるヘルプを見ていただくとして，ここでは Linux の bash について紹介する。

Bash 用の自動補完スクリプト自体は以下のコマンドで標準出力に出力される。

```text
$ gpgpdump completion bash
```

ユーザごとに仕込むのであれば

```text
$ source <(gpgpdump completion bash)
```

でOK。

システム全体として設定するのであれば以下のディレクトリのいずれかにスクリプトを放りこむ（後者がオススメ）。

- `/etc/bash_completion.d/`
- `/usr/share/bash-completion/completions/`

コマンドラインで書くとこんな感じ。

```text
sudo sh -c "gpgpdump completion bash > /usr/share/bash-completion/completions/gpgpdump"
```

これでログインし直せば有効になっているはずである。
たとえば

```text
$ gpgpdump
```

の状態で [`TAB`] キーを2回押下すれば

```text
$ gpgpdump [TAB][TAB]
completion  fetch       github      help        hkp         version
```

と利用可能なサブコマンドを列挙してくれる。
また

```text
$ gpgpdump -
```

とハイフンのみ入力した状態で [`TAB`] キーを2回押下すれば

```text
$ gpgpdump -[TAB][TAB]
--armor      --file=      --literal    -a           -l
--cert       --indent     --marker     -c           -m
--clipboard  --indent=    --private    -f           -p
--debug      --int        --utc        -i           -u
--file       --json       --version    -j           -v
```

てな感じに使えるオプションを表示してくれる。
便利！

## ブックマーク

- [OpenPGP の実装]({{< rlnk  "/openpgp/" >}})
- [GitHub に登録した OpenPGP 公開鍵を取り出す](https://zenn.dev/spiegel/articles/20201014-openpgp-pubkey-in-github)

[OpenPGP]: http://openpgp.org/
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/
[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[JSON]: https://tools.ietf.org/html/rfc7159 "RFC 7159 - The JavaScript Object Notation (JSON) Data Interchange Format"
[TOML]: https://github.com/toml-lang/toml "toml-lang/toml: Tom's Obvious, Minimal Language"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 5581]: http://tools.ietf.org/html/rfc5581
[RFC 6637]: http://tools.ietf.org/html/rfc6637
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[HKP]: https://tools.ietf.org/html/draft-shaw-openpgp-hkp-00 "draft-shaw-openpgp-hkp-00 - The OpenPGP HTTP Keyserver Protocol (HKP)"
[GitHub]: https://github.com/ "GitHub"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
