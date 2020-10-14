+++
title = "GitHub に登録されている OpenPGP 公開鍵の情報を取得する"
date =  "2019-10-21T17:29:15+09:00"
description = "公開鍵パケットデータを base64 コマンドで復号し，更に拙作の gpgpdump で可視化すれば中身を見ることができる。"
image = "/images/attention/openpgp.png"
tags = [ "cryptography", "openpgp", "github", "json", "gpgpdump" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

{{< div-box type="markdown" >}}
**【2020-10-14】**
同じ話を Zenn で書いてしまった。
併せてどうぞ orz

- [GitHub に登録した OpenPGP 公開鍵を取り出す](https://zenn.dev/spiegel/articles/20201014-openpgp-pubkey-in-github)
{{< /div-box >}}

ネットで見かけた小ネタで（笑）

[GitHub] の git リポジトリにアクセスする際に使う [SSH 公開鍵を取得する方法](https://qiita.com/zackey2/items/429c77e5780ba8bc1bf9 "もう「公開鍵送ってください」というやり取りは不要だった - Qiita")は割と知られているようだが

```text
$ curl -s https://github.com/spiegel-im-spiegel.keys
```

似た感じで登録している [OpenPGP] 公開鍵も取得できる。
ただし出力は JSON 形式で，こんな感じ。

```text
$ curl -s https://api.github.com/users/spiegel-im-spiegel/gpg_keys
[
  {
    "id": 305745,
    "primary_key_id": null,
    "key_id": "2287557885231C76",
    "raw_key": "-----BEGIN PGP PUBLIC KEY BLOCK-----\r\n\r\n ... \r\n-----END PGP PUBLIC KEY BLOCK-----",
    "public_key": "...",
    "emails": [
      {
        "email": "...",
        "verified": true
      }
    ],
    "subkeys": [
      {
        "id": 305746,
        "primary_key_id": 305745,
        "key_id": "5B07C6DBBBDAB020",
        "raw_key": null,
        "public_key": "...",
        "emails": [
        ],
        "subkeys": [
        ],
        "can_sign": false,
        "can_encrypt_comms": true,
        "can_encrypt_storage": true,
        "can_certify": false,
        "created_at": "2018-02-15T11:34:33.000Z",
        "expires_at": null
      }
    ],
    "can_sign": true,
    "can_encrypt_comms": false,
    "can_encrypt_storage": false,
    "can_certify": true,
    "created_at": "2018-02-15T11:34:33.000Z",
    "expires_at": "2020-02-15T00:22:09.000Z"
  },
  {
    ...
  },
  ...
]
```

いろいろ端折ってるが，あしからず。

この中で `raw_key` 項目に GitHub へ登録した ASCII armor 形式の公開鍵がそのまま入っている。
取り出しは [jq] コマンドを使って以下のようにできる。

```text
$ curl -s https://api.github.com/users/spiegel-im-spiegel/gpg_keys | jq -r .[0].raw_key
-----BEGIN PGP PUBLIC KEY BLOCK-----

mFIEWoTPwRMIKoZIzj0DAQcCAwRBr6HVaUrhEBxBcty/ToFv3aJyC+yojwVG84CL
...
JgA=
=gjzz
-----END PGP PUBLIC KEY BLOCK-----
```

ただし `raw_key` 項目は登録した公開鍵によっては `null` になっているようだ。
なんでだろう。

登録した鍵の公開鍵パケットのみであれば `public_key` 項目にセットされている[^pk1]。

[^pk1]: ちなみに公開鍵パケットのみで自己署名もない状態では暗号化も署名検証もできない。少なくとも [GnuPG] はそうなっている。

```text
$ curl -s https://api.github.com/users/spiegel-im-spiegel/gpg_keys | jq -r .[0].public_key
xlIEWoTPwRMIKoZIzj0DAQcCAwRBr6HVaUrhEBxBcty/ToFv3aJyC+yojwVG84CLs/XUsT7TUUxrrME+RrzbCs4PMYjdZ9B9nCcD1ni2Bjk+GI9/
```

これはバイナリデータを Base64 形式で符号化したもののようだ。
なので base64 コマンドで復号し，更に拙作の [gpgpdump] で可視化すれば

```text
$ curl -s https://api.github.com/users/spiegel-im-spiegel/gpg_keys | jq -r .[0].public_key | base64 -d | gpgpdump 
Public-Key Packet (tag 6) (82 bytes)
    Version: 4 (current)
    Public key creation time: 2018-02-15T09:09:37+09:00
        5a 84 cf c1
    Public-key Algorithm: ECDSA public key algorithm (pub 19)
    ECC Curve OID: nistp256 (256bits key size)
        2a 86 48 ce 3d 03 01 07
    ECDSA EC point (uncompressed format) (515 bits)
```

という感じに中身を見ることができる。

以上，広告記事でした（笑）

## ブックマーク

- [GitHub API v3 | GitHub Developer Guide](https://developer.github.com/v3/)
    - [User GPG Keys | GitHub Developer Guide](https://developer.github.com/v3/users/gpg_keys/)

- [もう「公開鍵送ってください」というやり取りは不要だった - Qiita](https://qiita.com/zackey2/items/429c77e5780ba8bc1bf9)

- [Git Commit で OpenPGP 署名を行う]({{< ref "openpgp/git-commit-with-openpgp-signature.md" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[GitHub]: https://github.com/
[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[jq]: https://stedolan.github.io/jq/
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
