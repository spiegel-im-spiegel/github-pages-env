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
