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

- [spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer](https://github.com/spiegel-im-spiegel/gpgpdump)

[gpgpdump] は [OpenPGP] パッケットの内容を human-readable な形式で可視化する CLI (Command-Line Interface) ツールである。
[山本和彦](http://www.mew.org/~kazu/)さんによる [pgpdump] を参考デザインとし [Go 言語]で組み直している。

[gpgpdump] は [pgpdump] と比較して以下の特徴がある。

- 平文テキストによる結果出力のほか [JSON] や [TOML] といった構造化テキスト・フォーマットによる出力もできる
- 現行仕様である [RFC 4880] に追加して [RFC 5581] および [RFC 6637] にも対応している
- 次期 [OpenPGP] ドラフト案である [RFC 4880bis] にも一部対応している
- HKP プロトコルを用いて [OpenPGP] 鍵サーバから直接公開鍵を取得して検証できる

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/gpgpdump.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/gpgpdump)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/gpgpdump/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/gpgpdump.svg)](https://github.com/spiegel-im-spiegel/gpgpdump/releases/latest)

## ダウンロードとビルド

[gpgpdump] は以下の [Go] コマンドでダウンロードとビルドができる。

```
$ go get github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump@latest
```

なおビルドには [Go] 1.13 以上が要件となる。
ご注意を。

各プラットフォーム用のバイナリも用意している。
[最新バイナリはリリースページから取得](https://github.com/spiegel-im-spiegel/gpgpdump/releases/latest)できる。

## 簡単な使い方

`-h` オプションで簡単なヘルプを表示できる。

```text
$ gpgpdump -h
Usage:
  gpgpdump [flags]
  gpgpdump [command]

Available Commands:
  help        Help about any command
  hkp         Dumps from OpenPGP key server
  version     Print the version number

Flags:
  -a, --armor         accepts ASCII input only
  -c, --cert          dumps attested certification in signiture packets (tag 2)
      --debug         for debug
  -f, --file string   path of OpenPGP file
  -h, --help          help for gpgpdump
      --indent int    indent size for output string
  -i, --int           dumps multi-precision integers
  -j, --json          output with JSON format
  -l, --literal       dumps literal packets (tag 11)
  -m, --marker        dumps marker packets (tag 10)
  -p, --private       dumps private packets (tag 60-63)
  -t, --toml          output with TOML format
  -u, --utc           output with UTC time
  -v, --version       output version of gpgpdump

Use "gpgpdump [command] --help" for more information about a command.
```

たとえば以下のような [OpenPGP] 電子署名データファイルがあるとする。

```text
$ cat sig.asc 
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2

iF4EARMIAAYFAlTDCN8ACgkQMfv9qV+7+hg2HwEA6h2iFFuCBv3VrsSf2BREQaT1
T1ZprZqwRPOjiLJg9AwA/ArTwCPz7c2vmxlv7sRlRLUI6CdsOqhuO1KfYXrq7idI
=ZOTN
-----END PGP SIGNATURE-----
```

これを [gpgpdump] で表示するとこんな感じの出力になる。

```text
$ gpgpdump -f sig.asc 
Signature Packet (tag 2) (94 bytes)
    Version: 4 (current)
    Signiture Type: Signature of a canonical text document (0x01)
    Public-key Algorithm: ECDSA public key algorithm (pub 19)
    Hash Algorithm: SHA2-256 (hash 8)
    Hashed Subpacket (6 bytes)
        Signature Creation Time (sub 2): 2015-01-24T11:52:15+09:00
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x31fbfda95fbbfa18
    Hash left 2 bytes
        36 1f
    ECDSA value r (256 bits)
    ECDSA value s (252 bits)
```

入力は標準入力からも可能である。

```text
$ cat sig.asc | gpgpdump
Signature Packet (tag 2) (94 bytes)
    Version: 4 (current)
    Signiture Type: Signature of a canonical text document (0x01)
    Public-key Algorithm: ECDSA public key algorithm (pub 19)
    Hash Algorithm: SHA2-256 (hash 8)
    Hashed Subpacket (6 bytes)
        Signature Creation Time (sub 2): 2015-01-24T11:52:15+09:00
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x31fbfda95fbbfa18
    Hash left 2 bytes
        36 1f
    ECDSA value r (256 bits)
    ECDSA value s (252 bits)
```

`-j` または `--json` オプションを指定すれば JSON フォーマットで出力される。

```text
$ cat sig.asc | gpgpdump -j | jq .
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
              "value": "2015-01-24T11:52:15+09:00"
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

| オプション名 | 内容                                                        |
| ------------ | ----------------------------------------------------------- |
| `armor`      | ASCII armor テキストのみ受け入れる                          |
| `cert`       | 署名パケット（tag 2）内の証明された認証を16進ダンプ表示する |
| `int`        | MPI データを16進ダンプ表示する                              |
| `literal`    | リテラル・パケット（tag 11）を16進ダンプ表示する            |
| `marker`     | マーカー・パケット（tag 10）を16進ダンプ表示する            |
| `private`    | プライベート用パケット（tag 60-63）を16進ダンプ表示する     |
| `utc`        | 時刻を UTC で表示する                                       |
| `indent`     | 平文テキスト出力時のインデント幅を指定する                  |
| `json`       | [JSON] 形式で出力する                                       |
| `toml`       | [TOML] 形式で出力する                                       |
| `debug`      | デバッグ用                                                  |

### HKP モード

```text
$ gpgpdump hkp -h
Dumps from OpenPGP key server

Usage:
  gpgpdump hkp [flags] <user ID or key ID>

Flags:
  -h, --help               help for hkp
      --keyserver string   OpenPGP key server (default "keys.gnupg.net")
      --port int           port number of OpenPGP key server (default 11371)
      --raw                output raw text from OpenPGP key server
      --secure             enable HKP over HTTPS

Global Flags:
  -a, --armor        accepts ASCII input only
  -c, --cert         dumps attested certification in signature packets (tag 2)
      --debug        for debug
      --indent int   indent size for output string
  -i, --int          dumps multi-precision integers
  -j, --json         output with JSON format
  -l, --literal      dumps literal packets (tag 11)
  -m, --marker       dumps marker packets (tag 10)
  -p, --private      dumps private packets (tag 60-63)
  -t, --toml         output with TOML format
  -u, --utc          output with UTC time
```

`hkp` コマンドを指定することで [OpenPGP] 鍵サーバから HKP プロトコルを使い，直接公開鍵を取得して中身を検証することができる。

```text
$ gpgpdump hkp 0x44ce6900e2b307a4
Public-Key Packet (tag 6) (269 bytes)
    Version: 4 (current)
    Public key creation time: 2009-11-09T00:20:55+09:00
        4a f6 e1 d7
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (2048 bits)
    RSA public encryption exponent e (17 bits)
User ID Packet (tag 13) (25 bytes)
    User ID: Alice <alice@example.com>
Signature Packet (tag 2) (312 bytes)
    Version: 4 (current)
    Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (34 bytes)
        Signature Creation Time (sub 2): 2009-11-09T00:20:55+09:00
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
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x44ce6900e2b307a4
    Hash left 2 bytes
        93 62
    RSA signature value m^d mod n (2045 bits)
Public-Subkey Packet (tag 14) (269 bytes)
    Version: 4 (current)
    Public key creation time: 2009-11-09T00:20:55+09:00
        4a f6 e1 d7
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (2048 bits)
    RSA public encryption exponent e (17 bits)
Signature Packet (tag 2) (287 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (9 bytes)
        Signature Creation Time (sub 2): 2009-11-09T00:20:55+09:00
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to encrypt communications.
            Flag: This key may be used to encrypt storage.
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x44ce6900e2b307a4
    Hash left 2 bytes
        66 f3
    RSA signature value m^d mod n (2048 bits)
```

`--raw` オプションを使うとダンプ表示はせず HKP リクエストの結果をそのまま表示する。

```text
$ gpgpdump hkp 0x44ce6900e2b307a4 --raw
<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd" >
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title>Public Key Server -- Get "0x44ce6900e2b307a4 "</title>
<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
<style type="text/css">
/*<![CDATA[*/
 .uid { color: green; text-decoration: underline; }
 .warn { color: red; font-weight: bold; }
/*]]>*/
</style></head><body><h1>Public Key Server -- Get "0x44ce6900e2b307a4 "</h1>
<pre>
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: SKS 1.1.6
Comment: Hostname: pgpkeys.uk

mQENBEr24dcBCADQeCxUo1pNF33ytHuzLn4vK9Z8LWXCUoZsQAZ9+cMKAzbQ9ncO+LfMleDz
RpjsBxYWDaTnn6a8OySveDcw9/CZ9Wu0ND0+uHErdNk5qh+z81x15sOAfN9xj4pUm0iH092Z
wuILrLjWWqgKMZYmB8HKaHXDkQmSfQmhx7oyZ4tWHfMN/VqBWLyUt0RaU0X+s4zLrdJSsTaf
ECZRo/2OJecpyBzLBc45Tzv3RJAXTyv31MLDYn38bS0EiShRoqaGIZthC7ZnX9EoaS2trg1K
uZtv6NeScRU4TqS21q/kYnE6HBnAMg7mI7dtFbg8x20TB2rTA5v8o/8cqZ3MLQukqjZ1ABEB
AAG0GUFsaWNlIDxhbGljZUBleGFtcGxlLmNvbT6JATgEEwECACIFAkr24dcCGwMGCwkIBwMC
BhUIAgkKCwQWAgMBAh4BAheAAAoJEETOaQDiswekk2IH/RtbG6zgJiPV91GZFMgmZJU2K6qu
iFfdzUNmvLyPfi/l8QUuwDAc8vzni7DYWFBWCOFk3vm8o/OCGxmhSHt1u0L4pm4AKcCqawdX
83H0PQBzvHy1US+33SCUb2OSnxBIwsDxISVSZ89Che+O6Bz3dl1mzsDonw3HGVy73OyspCof
7rRpMhSzTcXkP89pAV5k4IWSg612PlBNXwrCSjDiecx1Dnl0kKcPVl7tEVmYwSp36MBDo3dW
DjV/r6qyYF7D63tqoLfS1alXD6zpdJ0iRsV8s3zI+JEBGfWvxTZnOwsOzU7p+11wE1pBM6eX
/jNc275HWyQXBZ8Mwv/GRxbelXG5AQ0ESvbh1wEIANn7br8X1stD3Zc+IJ4LA8bavDODajhr
mQeSlf2VHocvAMpseldkiPIQgVF67Wh58gY+aeWl1W4KHy6VoTvnvi9Td8SDoH9ChJNXPqYA
p4u4yPHSq2HTYai2z9T+yLHH82dXl6kNW5H97Efp6C6enMO8SD5QPpGSJvzosgvwKZBerfxl
lV938/HSs5sXtjJRtZSJ2opbr6XaWP9fqy9B3HcHrk9YbAG6jG/ePh4s/vMmi7IZ/uH7Yf2/
d4hw9PJM6CJPnRjDO11qUwWWWQ/vUt/cX/ovnOjjJnj9gzIKQdX7Yt6/COcOmMZIfJdDrTdY
+NC07DJtyAKRXQ34q5jC+KcAEQEAAYkBHwQYAQIACQUCSvbh1wIbDAAKCRBEzmkA4rMHpGbz
CADLDT0pPL3xjWUy+5Liyv4gc+1EIuDz8kod8263w35bvP0H7urVsH1cODUrG1iG4S3BWvf3
fF9jy/1/WxpRS3huHyHCtyJTmC1wY/syGBBPnHO0CQzUkKG1Pb5IpEXg1T0SXihRqjiVtNT+
MHxlIXaMkCFSelweAJLM/OyrLJ5cwYr8JA+YZFav3pz0vZ58Fm9XRouvOIj9+0iX5Mnrw/2h
XRHCOgDrPp1aDAx9pnwY3oyy5dAof8o0Gpiaz34TRRdZ3vSFBloopXdNyH+V2cBUxLAxoqaQ
PrMW9QsmqRw/5A1YirEupzsd7z8/UVEKFIdPCxrMNWM/QArO22qEZjE3
=1Cgb
-----END PGP PUBLIC KEY BLOCK-----
</pre>
</body></html>
```

以下のオプションを使って [OpenPGP] 鍵サーバを指定することもできる。

| オプション名 | 既定値           | 内容                             |
| ------------ | ---------------- | -------------------------------- |
| `keyserver`  | `keys.gnupg.net` | [OpenPGP] 鍵サーバ名（ホスト名） |
| `port`       | 11371            | [OpenPGP] 鍵サーバのポート番号   |
| `secure`     | false            | HKP over HTTPS を有効にする      |

## [gpgpdump] を [Go] パッケージとして使う

使いどころが思い浮かばないのだが（笑），一応 [gpgpdump] は [Go 言語]用のパッケージまたはモジュールとして組み込むことができる。

たとえばこんな感じ。

```go
package main

import (
    "fmt"

    "github.com/spiegel-im-spiegel/gpgpdump"
    "github.com/spiegel-im-spiegel/gpgpdump/options"
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
    info, err := gpgpdump.ParseByte(
        []byte(openpgpStr),
        options.New(
            options.Set(options.ARMOR, true),
            options.Set(options.UTC, true),
        ),
    )
    if err != nil {
        fmt.Fprintf(os.Stderr, "%+v", err)
        return
    }
    fmt.Println(info)
}
```

あるいは [`io`]`.Reader` インタフェースを使うのであれば

```go
info, err := gpgpdump.Parse(
    bytes.NewBufferString(openpgpStr),
    options.New(
        options.Set(options.ARMOR, true),
        options.Set(options.UTC, true),
    ),
)
```

などとしてもよい。

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

`options.New()` 関数で CLI 版と同等のオプションを設定できる。
`options.New()` 関数の引数として `options.Set()` 関数を0個以上複数指定できる。
指定可能なオプションは以下の通り。

| オプション        | 既定値 | 対応する CLI オプション |
| ----------------- | ------ | ----------------------- |
| `options.ARMOR`   | false  | `armor`                 |
| `options.CERT`    | false  | `cert`                  |
| `options.INTEGER` | false  | `int`                   |
| `options.LITERAL` | false  | `literal`               |
| `options.MARKER`  | false  | `marker`                |
| `options.PRIVATE` | false  | `private`               |
| `options.UTC`     | false  | `utc`                   |
| `options.DEBUG`   | false  | `debug`                 |


`gpgpdump.Parse()` または `gpgpdump.ParseByte()` 関数の返り値として構造体 `info.Info` を取得できる。
`info.Info` 構造体の構成は以下の通り。

```go
//Info is information class for OpenPGP packets
type Info struct {
    Packets []*Item `toml:"Packet,omitempty" json:"Packet,omitempty"`
}

//Item is information item class
type Item struct {
    Name  string  `toml:"name" json:"name"`
    Value string  `toml:"value,omitempty" json:"value,omitempty"`
    Dump  string  `toml:"dump,omitempty" json:"dump,omitempty"`
    Note  string  `toml:"note,omitempty" json:"note,omitempty"`
    Items []*Item `toml:"Item,omitempty" json:"Item,omitempty"`
}
```

たとえば，先ほどの `main()` 関数を

{{< highlight go "hl_lines=13-15" >}}
func main() {
    info, err := gpgpdump.Parse(
        bytes.NewBufferString(openpgpStr),
        options.New(
            options.Set(options.ARMOR, true),
            options.Set(options.UTC, true),
        ),
    )
    if err != nil {
        fmt.Fprintf(os.Stderr, "%+v", err)
        return
    }
    if json, err := info.JSON(2); err == nil {
        fmt.Printf("%s", json)
    }
}
{{< /highlight >}}

のように変えれば

```text
$ go run sample.go 
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

のような出力になる。
これなら `info.Info` 構造体の構成が分かりやすいかな。

## ブックマーク

- [OpenPGP の実装]({{< rlnk  "/openpgp/" >}})

[OpenPGP]: http://openpgp.org/
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/
[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[JSON]: https://tools.ietf.org/html/rfc7159 "RFC 7159 - The JavaScript Object Notation (JSON) Data Interchange Format"
[TOML]: https://github.com/toml-lang/toml "toml-lang/toml: Tom's Obvious, Minimal Language"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 5581]: http://tools.ietf.org/html/rfc5581
[RFC 6637]: http://tools.ietf.org/html/rfc6637
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4314009071?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4314009071?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">暗号化 プライバシーを救った反乱者たち</a></dt>
    <dd>スティーブン・レビー (著), 斉藤 隆央 (翻訳)</dd>
    <dd>紀伊國屋書店 2002-02-16</dd>
    <dd>単行本</dd>
    <dd>4314009071 (ASIN), 9784314009072 (EAN), 4314009071 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-16">2018-12-16</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/B015643CPE?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/B015643CPE?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">暗号技術入門 第3版　秘密の国のアリス</a></dt>
    <dd>結城 浩 (著)</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>Kindle版</dd>
    <dd>B015643CPE (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
