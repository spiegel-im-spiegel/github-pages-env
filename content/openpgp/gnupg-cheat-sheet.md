+++
title = "GnuPG チートシート（鍵作成から失効まで）"
date =  "2017-12-01T17:51:18+09:00"
update =  "2017-12-02T16:20:26+09:00"
description = "ちうわけで GnuPG の使い方に関する簡単な「虎の巻（cheat sheet）」を作ってみることにした。"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

(move from [{{< ref "remark/2017/11/gnupg-sheat-sheet.md" >}}]({{< ref "remark/2017/11/gnupg-sheat-sheet.md" >}} "GnuPG チートシート（鍵作成から失効まで）"))

最近 `git commit` に電子署名する目的などで [GnuPG] を使う記事などをチラホラ見かけるようになったが，やっぱ使い慣れんもんは分からんよねぇ。
しかもバージョンによって微妙に挙動が異なるのが困りものである。

ちうわけで [GnuPG] の使い方に関する簡単な「虎の巻（cheat sheet）」を作ってみることにした。
対象となる  [GnuPG] のバージョンは最新版の 2.2.x とする。

なお，この記事は大変長文なので，あらかじめお茶菓子などを用意した上で読みはじめることをお勧めする。
また Qiita に簡易版を公開した。

- [GnuPG チートシート（簡易版） - Qiita](https://qiita.com/spiegel-im-spiegel/items/079d69282166281eb946)

説明はいいから例示だけ見せろという方はこちらで。

では，ご笑覧あれ。

## コマンドとオプション

[GnuPG] のコマンドラインはちょっと作りが古くて（なんせ初期の PGP の UI を引きずってるので`w`），今時あたり前な「サブコマンド」みたいな構成になっていない。
その代わりオプションの種別が「コマンド」と「オプション」に分かれている。
具体的には `gpg -h` でヘルプを見ると分かる（もちろん `-h` オプションもコマンドである）。

以上を踏まえて，そろそろ本題に入ろう。

## 鍵の作成

鍵の作成コマンドにはいくつか種類がある。

### --generate-key コマンド

`--generate-key` コマンドは対話モードで鍵の作成を行う。
短縮名は `--gen-key`。
あんまり短縮されていないな（笑）

コマンド自体は初期バージョンから存在するが，バージョンによって挙動がかなり違うので要注意だ。

```text
$ gpg --gen-key
gpg (GnuPG) 2.2.3; Copyright (C) 2017 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

注意: 全機能の鍵生成には "gpg --full-generate-key" を使います。

GnuPGはあなたの鍵を識別するためにユーザIDを構成する必要があります。

本名:
```

最新版 2.2.x では暗号アルゴリズムは RSA/2048bit，有効期限は作成日当日で固定されている。
したがって，ユーザが入力するのはユーザID（本名，電子メール・アドレス）とパスフレーズのみとなる（パスフレーズ入力時には Pinentry が起動する）。

`--generate-key` コマンドについては，以下のような設定ファイルを作って `--batch` オプションを付けて起動することで対話モードを回避し，かつアルゴリズム等の詳細な指定をすることもできる。

```text
$ cat alice-key.conf
Key-Type: RSA
Key-Length: 3072
Key-Usage: sign,cert
Subkey-Type: RSA
Subkey-Length: 3072
Subkey-Usage: encrypt
Name-Real: Alice
Name-Email: alice@example.com
Expire-Date: 0
Passphrase: passwd
%commit
%echo done

$ gpg --gen-key --batch alice-key.conf
gpg: 鍵058E5BB44555AF2Cを究極的に信用するよう記録しました
gpg: 失効証明書を 'C:/Users/alice/AppData/Roaming/gnupg/openpgp-revocs.d\DE93A51F5F4EC94847556525058E5BB44555AF2C.rev' に保管しました。
gpg: done

$ gpg --list-keys alice
pub   rsa3072 2017-11-30 [SC]
      DE93A51F5F4EC94847556525058E5BB44555AF2C
uid           [  究極  ] Alice <alice@example.com>
sub   rsa3072 2017-11-30 [E]
```

設定ファイルの書き方は "[Unattended GPG key generation](https://www.gnupg.org/documentation/manuals/gnupg/Unattended-GPG-key-generation.html "Using the GNU Privacy Guard: Unattended GPG key generation")” を参照のこと。
いったん設定ファイルを作ってしまえばこの方法が一番簡単かな。

なお `Passphrase` の項目を削除すれば `--batch` モードでも Pinentry で設定するパスフレーズを訊いてくるので「設定ファイルにパスフレーズを書くのは...」という方も安心である。

### --full-generate-key コマンド

対話モードで暗号アルゴリズムや鍵長を指定したい場合は `--full-generate-key` コマンドを使う。
短縮名は `--full-gen-key`。

具体的にはこんな感じで進行する。

```text
$ gpg --full-gen-key
gpg (GnuPG) 2.2.3; Copyright (C) 2017 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

ご希望の鍵の種類を選択してください:
   (1) RSA と RSA (デフォルト)
   (2) DSA と Elgamal
   (3) DSA (署名のみ)
   (4) RSA (署名のみ)
あなたの選択は? 2
DSA 鍵は 1024 から 3072 ビットの長さで可能です。
鍵長は? (2048) 3072
要求された鍵長は3072ビット
鍵の有効期限を指定してください。
         0 = 鍵は無期限
      <n>  = 鍵は n 日間で期限切れ
      <n>w = 鍵は n 週間で期限切れ
      <n>m = 鍵は n か月間で期限切れ
      <n>y = 鍵は n 年間で期限切れ
鍵の有効期間は? (0)1y
鍵は11/30/18 10:39:03 東京 (標準時)で期限切れとなります
これで正しいですか? (y/N)
```

`--expert` オプションを付けると選択可能なアルゴリズムの組み合わせが増える。
こんな感じ。

```text
$ gpg --full-gen-key --expert
gpg (GnuPG) 2.2.3; Copyright (C) 2017 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

ご希望の鍵の種類を選択してください:
   (1) RSA と RSA (デフォルト)
   (2) DSA と Elgamal
   (3) DSA (署名のみ)
   (4) RSA (署名のみ)
   (7) DSA (機能をあなた自身で設定)
   (8) RSA (機能をあなた自身で設定)
   (9) ECC と ECC
  (10) ECC (署名のみ)
  (11) ECC (機能をあなた自身で設定)
  (13) 既存の鍵
あなたの選択は?
```

### --quick-generate-key コマンド

コマンドライン一発で鍵を作成したい場合は `--quick-generate-key` コマンドでユーザID，アルゴリズム，有効期限を指定できる。
短縮名は `--quick-gen-key`。

```text
Usage: gpg [options] --quick-generate-key user-id [algo [usage [expire]]]
```

`algo` にはアルゴリズムと鍵長を文字列で指定する。
指定可能な文字列は以下の通り。

| 公開鍵暗号アルゴリズム | 名前 |
|:-----------------------|:-----|
| RSA | `default` (= `rsa2048`) |
| RSA (署名のみ) | `rsa` (= `rsa2048`), `rsa1024`, `rsa2048`, `rsa3071`, `rsa4096` |
| DSA (署名のみ) | `dsa` (= `dsa2048`), `dsa1024`, `dsa2048`, `dsa3072` |
| ECDH/EdDSA | `future-default` (= `cv25519`/`ed25519`) |
| EdDSA (署名のみ) | `ed25519` |
| ECDSA (署名のみ) | `nistp256`, `nistp384`, `nistp521`, `brainpoolP256r1`, `brainpoolP384r1`, `brainpoolP512r1`, `secp256k1` |

`usage` には主鍵の機能を文字列で指定する。

| 機能 | 名前 |
|:-----|:-----|
| 署名 | `sign` |
| 証明 | `cert` |
| 認証 | `auth` |
| 暗号化 | `encr` |

主鍵には自動的に `cert` が付与されるため暗号化機能しかないアルゴリズム（ElGamal や ECDH）を主鍵に使うことはできない。
また暗号アルゴリズムと機能がマッチしない場合はエラーになる（電子署名用のアルゴリズムなのに `encr` を指定するなど）。

目的が複数ある場合はカンマで区切って列挙する。
なお `default` または `-` を指定すれば `sign`+`cert` となるので，通常は `default` のままでよい。

`expire` には有効期限を指定する。
1週間なら `7d` または `1w`，1年なら `12m` または `1y` といった感じ。
`0` を指定すると無期限になる。
省略すると作成日当日が有効期限となる。

` --quick-generate-key` コマンドの実行例はこんな感じ。

```text
$ gpg --quick-gen-key "Alice <alice@example.com>" default default 0
たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。

たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。
gpg: 鍵FED63B6C83CE0152を究極的に信用するよう記録しました
gpg: 失効証明書を 'C:/Users/alice/AppData/Roaming/gnupg/openpgp-revocs.d\57D6D370A7E9BA27A02367DAFED63B6C83CE0152.rev' に保管しました。
公開鍵と秘密鍵を作成し、署名しました。

pub   rsa2048 2017-11-30 [SC]
      57D6D370A7E9BA27A02367DAFED63B6C83CE0152
uid                      Alice <alice@example.com>
sub   rsa2048 2017-11-30 [E]
```

### パスフレーズ入力の回避

`--quick-generate-key` コマンドでもパスフレーズの入力は Pinentry から行うことになるが `--pinentry-mode` オプションおよび `--passphrase` オプションを付加することで回避できる。

```text
$ gpg --pinentry-mode loopback --passphrase passwd --quick-gen-key "Alice <alice@example.com>" default default 0
```

ただしコマンドラインの履歴に入力したパスフレーズが残ってしまうのであまりお勧めできないが...

### --quick-add-key コマンドによる副鍵の追加

作成した鍵に `--quick-add-key` コマンドで後から暗号鍵を追加できる。
これは ` --quick-generate-key` コマンドで主鍵のみ作って後から副鍵を加えたい場合などに有効である。

```text
Usage: gpg [options] --quick-add-key key-fingerprint [algo [usage [expire]]]
```

たとえば，以下の鍵に対して

```text
$ gpg --list-keys alice
pub   dsa3072 2017-11-30 [SC]
      B5BF56B346B4D961E6BF25A45CC68B4A317E8E5C
uid           [  究極  ] Alice <alice@example.com>
```

以下のように暗号鍵を副鍵として追加できる。

```text
$ gpg --quick-add-key B5BF56B346B4D961E6BF25A45CC68B4A317E8E5C elg3072 encr
たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。

$ gpg --list-keys alice
pub   dsa3072 2017-11-30 [SC]
      B5BF56B346B4D961E6BF25A45CC68B4A317E8E5C
uid           [  究極  ] Alice <alice@example.com>
sub   elg3072 2017-11-30 [E]
```

ちなみに `B5BF56B346B4D961E6BF25A45CC68B4A317E8E5C` という長ったらしい数字列は鍵指紋（key fingerprint）である。
[OpenPGP] では鍵指紋をそのまま（または下位バイトを）鍵IDとして使っている。

副鍵では機能として `cert` は指定できない。
また暗号アルゴリズムと機能がマッチしない場合はエラーになる。
ただし `default` または `-` を指定すればアルゴリズムに合わせた適切な機能をセットしてくれるみたいなので，大抵の場合は `default` でいいだろう。

暗号化用に使用できるアルゴリズムは以下の通り。

| 公開鍵暗号アルゴリズム | 名前 |
|:-----------------------|:-----|
| RSA | `default` (= `rsa2048`), `rsa` (= `rsa2048`), `rsa1024`, `rsa2048`, `rsa3071`, `rsa4096` |
| ElGamal | `elg` (= `elg2048`), `elg1024`, `elg2048`, `elg3072` |
| ECDH | `cv25519`, `nistp256`, `nistp384`, `nistp521`, `brainpoolP256r1`, `brainpoolP384r1`, `brainpoolP512r1`, `secp256k1` |

電子署名用のアルゴリズムも（署名用の副鍵として指定すれば）もちろん使える。

## 鍵の管理

作成した鍵や配布・受領した公開鍵を管理するためのコマンドを紹介する。

### 鍵束内の公開鍵の検索

鍵束内の公開鍵の検索を検索する場合， `--list-keys` コマンドの引数にユーザID（の一部）または鍵IDを指定することで，条件にマッチする鍵を検索できる。
短縮名は `-k`。

```text
$ gpg -k alice
pub   dsa3072 2017-11-23 [SC]
      3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E
uid           [  充分  ] Alice (root) <alice@example.com>

pub   rsa2048 2017-11-23 [SC]
      A3CEFEEEDA222024F325C403DFFC3F67BBB3C083
uid           [  究極  ] Alice (commit) <alice@example.com>
sub   rsa2048 2017-11-23 [E]
```

引数なしで `--list-keys` コマンドを起動した場合は公開鍵の鍵束（`pubring.kbx`）にある鍵が全て列挙される。

秘密鍵を検索する場合には `--list-secret-keys` コマンドを使う。
短縮名は大文字の `-K`。

```text
$ gpg -K alice
sec   rsa2048 2017-11-23 [SC]
      A3CEFEEEDA222024F325C403DFFC3F67BBB3C083
uid           [  究極  ] Alice (commit) <alice@example.com>
ssb   rsa2048 2017-11-23 [E]
```

### 副鍵の鍵指紋の表示

`--list-keys` コマンドでも主鍵の鍵指紋が表示されるが，副鍵の鍵指紋も表示したい場合は `--fingerprint` コマンドを2つ重ねる。

```text
$ gpg --fingerprint --fingerprint alice
pub   rsa2048 2017-11-30 [SC]
      79FD 2B99 F3C6 2D2D 3B85  0BBC 93B3 5094 7582 0D5D
uid           [  究極  ] Alice <alice@example.com>
sub   rsa2048 2017-11-30 [E]
      476E 9EA7 D703 F0BB 01B6  FA44 9278 B060 D202 3C53
```

### パスフレーズの変更

秘密鍵のパスフレーズを変更する場合には `--change-passphrase` コマンドを使う。
短縮名は `--passwd`。

```text
$ gpg --passwd alice
```

引数にはユーザID（の一部）を指定できる。
パスワードの入力は Pinentry で行う。

### 有効期限の変更

自身の鍵の有効期限を変更する場合には `--quick-set-expire` コマンドを使う。

鍵の鍵指紋がが以下の場合

```text
$ gpg --fingerprint --fingerprint alice
pub   rsa2048 2017-11-30 [SC]
      79FD 2B99 F3C6 2D2D 3B85  0BBC 93B3 5094 7582 0D5D
uid           [  究極  ] Alice <alice@example.com>
sub   rsa2048 2017-11-30 [E]
      476E 9EA7 D703 F0BB 01B6  FA44 9278 B060 D202 3C53
```

有効期限を2年（`2y`）に指定するなら，操作は以下の通り。

```text
$ gpg --quick-set-expire 79FD2B99F3C62D2D3B850BBC93B3509475820D5D 2y

$ gpg --list-keys alice
pub   rsa2048 2017-11-30 [SC] [有効期限: 2019-11-30]
      79FD2B99F3C62D2D3B850BBC93B3509475820D5D
uid           [  究極  ] Alice <alice@example.com>
sub   rsa2048 2017-11-30 [E]

$ gpg --quick-set-expire 79FD2B99F3C62D2D3B850BBC93B3509475820D5D 2y 476E9EA7D703F0BB01B6FA449278B060D2023C53

$ gpg --list-keys alice
pub   rsa2048 2017-11-30 [SC] [有効期限: 2019-11-30]
      79FD2B99F3C62D2D3B850BBC93B3509475820D5D
uid           [  究極  ] Alice <alice@example.com>
sub   rsa2048 2017-11-30 [E] [有効期限: 2019-11-30]
```

前半は主鍵，後半は（主鍵に紐づく）副鍵の有効期限を変更している。
このことから分かるとおり，主鍵と副鍵は個別に有効期限を設定することが可能である。

### 公開鍵をエクスポートする

公開鍵のエクスポートには `--export` コマンドを使う。

```text
$ gpg --armor --export alice
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQENBFofiskBCADjUvPHA3PNscg0K74/Uwxj46+oLsyIy7fYIp/4C4dHejcbbPjx
VFeic9wQ4aQFp3VKjYgONgQrRo/9p40Ei1+PtMAV7D6Oy6dxlV8zyCJcSf74ahpB
B15GyA7v4uvTf0Py+Ujyt241ik0fXeLEuwt7p4SIbgJnQs1Fb+61wo8UcCFOLJO5
An6HjXNgNs6fFoiTad+T4PfaTbRHLfFPkoqmDUKWy40hjWl+Ui0QborXH+PUeUm9
vgHbqZzS0QRDGI7rO9AeJ6LweBkP1A2qbDLyexS/F+WUEcY0b76IQM5XH0txwnnl
uCPYcQfIGWce3US1GWJhChF9s/bMGVXOEJbvABEBAAG0GUFsaWNlIDxhbGljZUBl
eGFtcGxlLmNvbT6JAVQEEwEIAD4CGwMFCwkIBwIGFQgJCgsCBBYCAwECHgECF4AW
IQR5/SuZ88YtLTuFC7yTs1CUdYINXQUCWh+LMAUJA8JnZwAKCRCTs1CUdYINXcKT
B/4tLFaPRe289GcX91yLJ/yPS0JvvJKyZzjpNqLbKHuQHPEqGromMGlP4LcaGdFL
rVZ36W3kVk+75q8JFkld0eRS22vftjz6lA9lyb3W9lU1CayF5s3IsC/Ehj55uaHc
OHnp6rl7zEeIdvca6yV0gwySs3j9VPHy58zNrpN/clHoB4Zozy6vCXFMShyLc/wF
brPySf/5LP/642Uro92M2lbkIvZpDhZCVG7s7Ilz3BzsTTNMPkPd5yvdGa5lHQzK
OmXHaxydOYbEWBgqRGqzEIIoLbEd8KHxJVIVDfcAQCjSWRUjAUSDLpBokGsKoQfp
41NjWwjkIsfyJ2tDUeRPGYRbuQENBFofiskBCACzyYfIB+/ZwJBJXw7WMDlEKdnz
L4abwVpw9rBGAWGXjaC/cu7l0svNilXyTgZNq4uKddJ6aYjs7of0SaBl20I8aj5G
nbw0pG+KkoYhfpZaAZc+bcb+6SprSbAsRhrZ810XNIBUMa8XWsUDn1uv70vGBWBv
keKZZ7FJ4kuQe0nTONmvQ4EwFekV+IXT5LwdgmPWF0QR7cO8jqeb6psHYauktuzZ
2ul4nMLmLLf/m4DwiCAbEdToBXqRA30KshtgBYYQwL1YkWYgknnAdhHyeu6ybJvv
Y57JYzotjFOlnFhtcGITESEWv+pnj0RJUUrlVwLkJhUOKMwL+sbhw0s5+m27ABEB
AAGJATwEGAEIACYCGwwWIQR5/SuZ88YtLTuFC7yTs1CUdYINXQUCWh+LhAUJA8Jn
uwAKCRCTs1CUdYINXXuvB/9IKK3SLgJ6lOc2Vq73rGYsrDqfjYt5rCDXhjIaFRE7
LYmFJcGL5CHJTae438XtAixa+mu6PYG28eknjZs58Cx/bSj9uS6NiLAPCgyTAtvg
ao6usECOm9Y0xf2+ZcZ9Uji+wsCAFmxRC9je0yUErVyuyQRqzNtdqytnszoTzvb9
iOP8sX/YNrjC83BtZ4Vg3fzAu8qvwbObgSbws5M8TBwIKd4WFTjOtSU6F8aioJ1g
mpfd8KGljHkzC0oG8l8fZiTNYqkIMbfyfPpVwsSqsysLKofifFT+mNs79DJdqNFO
HA2W4WzekYmWWmgK7J8kXHYkxUJA6VpSmNAKwUKqXbNV
=hneF
-----END PGP PUBLIC KEY BLOCK-----
```

`--armor` オプションを指定すると，上記のように， ASCII Armor 形式のテキストを出力する。
短縮名は `-a`。
`--armor` オプションを付けないとバイナリを吐く。

秘密鍵をエクスポートする場合は `--export-secret-key` コマンドを使う（パスフレーズ入力あり）。

公開鍵をファイル等で配布する場合は `--export` コマンドの出力をファイルに落とせばよい。

```text
$ gpg -a --export alice > alice-key.asc
```

### 公開鍵をインポートする

公開鍵のインポートには `--import` コマンドを使う。

```text
$ gpg --import alice-key.asc
```

インポートではパイプが使えるので

```text
cat alice-key.asc | gpg --import
```

などとできる。
また Web ページ上に公開鍵のファイルを置いている場合は `--fetch-keys` コマンドで直接インポートすることもできる。

```text
$ gpg --fetch-keys http://www.baldanders.info/spiegel/pubkeys/spiegel.asc
```



インポートする鍵が既に鍵束にある場合でも，単純な上書きではなく， [GnuPG] がいい感じにマージしてくれる。

### 公開鍵を鍵サーバに送信する

鍵束にある公開鍵を鍵サーバに送信するには `--send-keys` コマンドを使う。

```text
$ gpg --keyserver keys.gnupg.net --send-keys 7E20B81C
```

鍵の指定には鍵Dを使う。

上記のように `--keyserver` オプションで鍵サーバを指定する。
または鍵束フォルダにある `gpg.conf` ファイルに既定の鍵サーバを指定できる。

```text
keyserver  keys.gnupg.net
```

鍵サーバは，基本的には互いに同期しているので，どのサーバを指定してもいいのだが，有名なところでは

- [keys.gnupg.net](http://keys.gnupg.net/ "Nebraska Wesleyan University - OpenPGP Keyserver")
- [pgp.mit.edu](https://pgp.mit.edu/ "MIT PGP Key Server")
- [pgp.nic.ad.jp](http://pgp.nic.ad.jp/ "PGP KEYSERVER")

あたりだろうか。

### 公開鍵を鍵サーバから受信する

鍵サーバから公開鍵を受信する場合は `--receive-keys` コマンドを使う。
短縮名は `--recv-keys`。

```text
$ gpg --keyserver keys.gnupg.net --recv-keys 7E20B81C
```

送信のときと同じく，こちらも鍵の指定には鍵Dを使う。
あらかじめ鍵IDがわからない場合は `--search-keys` コマンドで検索できる。

```text
$ gpg --keyserver keys.gnupg.net --search-keys alice@example.com
(1)     by Teemob <alice@example.com>
        by Teemob <lockstar2017@gmail.com>
          3072 bit RSA key 966893ECDA2FD3EC, 作成: 2017-11-15
(2)     Alice (Alice's key) <Alice@example.com>
          1024 bit DSA key A251C75C6213F841, 作成: 2017-11-12, 有効期限: 2018-11-12
(3)     Alice <alice@example.com>
          3072 bit RSA key 8EDAABFF277776F3, 作成: 2017-11-03
(4)     Alice <alice@example.com>
          3072 bit RSA key 37FC4F26B92A3964, 作成: 2017-10-18
(5)     Alice <alice@example.com>
          3072 bit RSA key 09AB44CAA589D7A2, 作成: 2017-10-04
(6)     Alice <alice@example.com>
          3072 bit RSA key D3205D5A68E02E2B, 作成: 2017-10-04
(7)     Alice <alice@example.com>
          2048 bit RSA key 29FD3D6668D47FA1, 作成: 2017-09-14
(8)     Alice <alice@example.com>
          3072 bit RSA key 25B9727BCE238CDE, 作成: 2017-08-10
(9)     Alice <alice@example.com>
          3072 bit RSA key 2604F9169E9C4E37, 作成: 2017-08-05 (失効)
(10)    Alice <alice@example.com>
          3072 bit RSA key B2251B1A2B632A2E, 作成: 2017-07-17
(11)    Alice <alice-example-187723@mailismagic.com>
          2048 bit RSA key FF99048E395DC7E7, 作成: 2017-04-20, 有効期限: 2019-04-20
Keys 1-11 of 103 for "alice@example.com".  番号(s)、N)次、またはQ)中止を入力してください >
```

検索結果に対して番号を指定すればそのままインポートしてくれる。

### 公開鍵に署名する

インポートした公開鍵が有効であることを確認したら，公開鍵に電子署名して有効化しよう。
公開鍵への電子署名には `--sign-key` コマンドまたは `--quick-sign-key` コマンドを使う。

`--sign-key` コマンドは対話モードで複数の鍵にひとつずつ署名することができる。
`--quick-sign-key` コマンドは鍵指紋を指定して一気に処理を行う（パスフレーズ入力あり）。

```text
gpg --quick-sign-key 1B5202DB4A3EC776F1E0AD18B4DA3BAE7E20B81C
```

電子署名が可能な秘密鍵を複数所持している場合は `--local-user` オプションで電子署名に使う鍵を指定する。
短縮名は `-u`。

```text
gpg -u alice --quick-sign-key 1B5202DB4A3EC776F1E0AD18B4DA3BAE7E20B81C
```

または電子署名に使う鍵を鍵束フォルダにある `gpg.conf` ファイルで指定することもできる。

```text
default-key alice
```

この電子署名は公開鍵のエクスポート時にも付加されて配布される。
電子署名を配布されては困る場合は `--lsign-key` コマンドまたは `--quick-lsign-key` コマンドを使う。

## データの暗号化

[GnuPG] の暗号化は概ね2種類ある。

### ハイブリッド暗号

ハイブリッド暗号は [OpenPGP] の基本機能で，平文を暗号化する「セッション鍵」とセッション鍵を暗号化する公開鍵で構成される。

{{< fig-img src="http://www.baldanders.info/spiegel/archive/pgpdump/hybrid-enc.svg" width="715" title="「わかる！ OpenPGP 暗号」より" link="http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml" >}}

セッション鍵は [GnuPG] が自動的にし生成するのでコマンドラインではセッション鍵を暗号化する公開鍵を指定する。

暗号化を行うには `--encrypt` コマンドを使う。
短縮名は `-e`。

```text
$ gpg -a --recipient alice -e plain-data
-----BEGIN PGP MESSAGE-----

hQIOAzn9g6TGzi9EEAgAsjEZxs4vutjg1U6BooimWrX3immaTL958Uheqcr7risr
2MCfVzzHuZtOpS4/lk/K4zk2xCxR3/NreZKlGrWZ205RCUJEY6Hy8GtrjJ2yilC6
ZV2U/ICRrpoTJm/J/R7W+99arXhP3zDSD2k5Fx9AMJ+OaKuHaBTxJQUtESV8J7Uk
RhyxQvJPXIfNRG4ZrfTzFzVOV1s9EeFR4UQHhQnp2q+7LA3qrfEh/y/sj4fs6o3G
KYcRVvUeAYsC1NGGcmpK6Q33oWJxN9vxl+NYlLebtCDS6GYl/bMw+YCXtfMh+cA2
aLiGUqXZT0Nhb/zVX8zlnP6CZE2kxS60LmTWv11DMQf/UCjdnIM80GKFvvy7/Vas
OlAwzQv2sWgI4ayL/VvslGVixSATsLD9DREjGo2/RfyDX/aLRsykK7H+Lr/+a3kp
LQzviY0ogYem1jCcqJs6wKMh1B+M+Ukkk9kVrgXelM6bmPT93Sb54LW9VVCf0GFK
ntVqfAkhOSOt3p+mHGH0hAmzGGVA9FGU5dIpvWUrMRdoBBJXj3akFVfLFv81QU9H
j3CCVHvCnGxBDtXWJV9CqVYWARit72R8FOLonpkFTRJ/IvFpePTccsMfsVvvBxS0
jt88EQAZ7bpdoJZ9qklr7LPMcNzXfZHdZLzNihbLhgEpVkfxI1vfflS5B5p1fIrK
/NJIAXJc8rgTJ0uI6MyYsgmJS2IVDXzwlsZWDLE9D3cbB8Xa53mlnPvmHgHwxAEn
Ic3OL8vsjZz9IcRksLr38/nbWhsHIUOrCovj
=jIJs
-----END PGP MESSAGE-----
```

あるいはパイプを使って

```text
$ echo Hello world | gpg -a --recipient alice -e
-----BEGIN PGP MESSAGE-----

hQIOAzn9g6TGzi9EEAgAsjEZxs4vutjg1U6BooimWrX3immaTL958Uheqcr7risr
2MCfVzzHuZtOpS4/lk/K4zk2xCxR3/NreZKlGrWZ205RCUJEY6Hy8GtrjJ2yilC6
ZV2U/ICRrpoTJm/J/R7W+99arXhP3zDSD2k5Fx9AMJ+OaKuHaBTxJQUtESV8J7Uk
RhyxQvJPXIfNRG4ZrfTzFzVOV1s9EeFR4UQHhQnp2q+7LA3qrfEh/y/sj4fs6o3G
KYcRVvUeAYsC1NGGcmpK6Q33oWJxN9vxl+NYlLebtCDS6GYl/bMw+YCXtfMh+cA2
aLiGUqXZT0Nhb/zVX8zlnP6CZE2kxS60LmTWv11DMQf/UCjdnIM80GKFvvy7/Vas
OlAwzQv2sWgI4ayL/VvslGVixSATsLD9DREjGo2/RfyDX/aLRsykK7H+Lr/+a3kp
LQzviY0ogYem1jCcqJs6wKMh1B+M+Ukkk9kVrgXelM6bmPT93Sb54LW9VVCf0GFK
ntVqfAkhOSOt3p+mHGH0hAmzGGVA9FGU5dIpvWUrMRdoBBJXj3akFVfLFv81QU9H
j3CCVHvCnGxBDtXWJV9CqVYWARit72R8FOLonpkFTRJ/IvFpePTccsMfsVvvBxS0
jt88EQAZ7bpdoJZ9qklr7LPMcNzXfZHdZLzNihbLhgEpVkfxI1vfflS5B5p1fIrK
/NJIAXJc8rgTJ0uI6MyYsgmJS2IVDXzwlsZWDLE9D3cbB8Xa53mlnPvmHgHwxAEn
Ic3OL8vsjZz9IcRksLr38/nbWhsHIUOrCovj
=jIJs
-----END PGP MESSAGE-----
```

とすることもできる。

`--recipient` がセッション鍵の暗号化を行う公開鍵を指定するオプションである。
短縮名は `-r`。
`--recipient` オプションは複数指定できる。
また鍵束フォルダにある `gpg.conf` ファイルで常に使用する公開鍵を指定することもできる[^dr1]。

[^dr1]: `default-recipient-self` の指定は自身の鍵で復号できるよう設定するためのものである。相手の公開鍵のみで暗号化してしまうと，暗号化した本人が復号できないことになってしまうため。

```text
default-key alice
default-recipient-self
```

### セッション鍵のみで暗号化する

公開鍵は使わずセッション鍵のみで暗号化を行う場合は `--symmetric` コマンドを使う。
短縮名は `-c`。

```text
$ echo Hello world | gpg -a -c
-----BEGIN PGP MESSAGE-----

jA0EBwMChZ5yarrU9aTF0kIBioFpcLD/laFWIMDVz7AzkzQl+Xwnao+iKpE+yaGo
sWe2GdB8IGA0O+CAqQYqwQTLKFVtWmAJKMi1hXsb/fuPpzU=
=5pGP
-----END PGP MESSAGE-----
```

コマンド起動時にパスフレーズの入力を要求され，パスフレーズからセッション鍵を生成して暗号化を行う。
したがって，何らかの方法で暗号データの受け手とパスフレーズを共有する必要がある。

## 暗号データの復号

暗号データのf区号には `--decrypt` コマンドを使う。
短縮名は `-d`。

```text
$ echo Hello world | gpg -r alice -e -a > alice-enc.asc

$ gpg -d alice-enc.asc
gpg: 2048-ビットELG鍵, ID 39FD83A4C6CE2F44, 日付2017-11-30に暗号化されました
      "Alice <alice@example.com>"
Hello world
```

（パスフレーズ入力あり）

復号したデータはファイルにリダイレクトすればいいのだが， Windows の場合は安全のため `--output` オプションを使うことをお勧めする。
短縮名は `-o`。

```text
$ gpg -o out.txt -d alice-enc.asc
gpg: 2048-ビットELG鍵, ID 39FD83A4C6CE2F44, 日付2017-11-30に暗号化されました
      "Alice <alice@example.com>"
Hello world

$ cat out.txt
Hello world
```

セッション鍵のみで暗号化した場合も同じコマンドで復号できる。

```text
$ echo Hello world | gpg -a -c > alice-sym-enc.asc

$ gpg -d alice-sym-enc.asc
gpg: AES暗号化済みデータ
gpg: 1 個のパスフレーズで暗号化
Hello world
```

（パスフレーズ入力あり）

## データへの電子署名と検証

データへの電子署名にも幾つかの方法がある。

### クリア署名

まずデータがテキストの場合は「クリア署名」という方法が使える。
クリア署名には `--clear-sign` コマンドを使う。

```text
$ echo Hello world | gpg -u alice --clear-sign
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA256

Hello world
-----BEGIN PGP SIGNATURE-----

iIgEAREIADAWIQTvp6IB8w6ZkSW4Whx6nWuc4jC66QUCWh/IZhIcYWxpY2VAZXhh
bXBsZS5jb20ACgkQep1rnOIwuuns2QD/RWTidtZjon5cPaiGJHM6oYnYx4HpQXNw
/xABYweyKdgA/3ArBLWmGhGq1aB8au7bixK91IdIRyhLC0DDJhXG2vM/
=sLc9
-----END PGP SIGNATURE-----
```

このようにクリア署名は元になるテキストと電子署名がくっついた状態で出力される。
なお，クリア署名の場合は必ず ASCII Armor 形式の出力になるため `--armor` オプションは不要である。

署名の検証には `--verify` コマンドを使う。
少し横着して署名と検証をパイプで繋いでしまおう。

```text
$ echo Hello world | gpg -u alice --clear-sign | gpg --verify
gpg: 11/30/17 18:02:44 東京 (標準時)に施された署名
gpg:                DSA鍵EFA7A201F30E999125B85A1C7A9D6B9CE230BAE9を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

まぁ自分で署名して，出力をそのまま自分で検証してるんだから正しくてあたり前なのだが，流れは分かると思う。

実は `--verify` コマンドは `--decrypt` コマンドで代替えできる。

```text
$ echo Hello world | gpg -u alice --clear-sign | gpg -d
Hello world
gpg: 11/30/17 18:05:11 東京 (標準時)に施された署名
gpg:                DSA鍵EFA7A201F30E999125B85A1C7A9D6B9CE230BAE9を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

このように `--decrypt` コマンドを使うと署名対象のテキストを抽出して出力してくれるのが利点である。

### 分離署名

次はファイルへの電子署名をやってみる。
まず署名対象のファイルを用意する。

```text
$ echo Hello world> hello.txt

$ cat hello.txt
Hello world
```

このファイルを配布する際に途中で改竄がないか知りたい。
こういう場合は「分離署名」にする。
分離署名には `--detach-sign` コマンドを使う。
短縮名は `-b`。

```text
$ gpg -u alice -b hello.txt
```

この結果 `hello.txt` ファイルと同じ場所に `hello.txt.sig` ファイルが作成される。
中身はバイナリデータである。

この`hello.txt` ファイルと `hello.txt.sig` ファイルをセットで配布するのである。
どちらかのファイルが改竄されていれば署名の検証が NG になるはずである。

```text
$ gpg --verify hello.txt.sig
gpg: 署名されたデータが'hello.txt'にあると想定します
gpg: 11/30/17 18:31:17 東京 (標準時)に施された署名
gpg:                DSA鍵EFA7A201F30E999125B85A1C7A9D6B9CE230BAE9を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

ここでは署名対象のファイルを `hello.txt` と推測して署名の検証を行っている。
署名対象のファイルを明示して指定するには

```text
$ gpg --verify hello.txt.sig hello.txt
gpg: 11/30/17 18:31:17 東京 (標準時)に施された署名
gpg:                DSA鍵EFA7A201F30E999125B85A1C7A9D6B9CE230BAE9を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

とする。
分離署名の検証でも `--decrypt` コマンドが使える。
ただし出力は全く同じ。

```text
$ gpg -d hello.txt.sig
gpg: 署名されたデータが'hello.txt'にあると想定します
gpg: 11/30/17 18:31:17 東京 (標準時)に施された署名
gpg:                DSA鍵EFA7A201F30E999125B85A1C7A9D6B9CE230BAE9を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

なお，署名対象のファイルがテキスト・ファイルの場合は `--textmode` オプションを付けて電子署名を行ったほうが安全である。

```text
$ gpg -u alice --textmode -b hello.txt
```

テキスト・ファイルの場合，配布経路によっては改行コードが変えられたりするため（電子メールや FTP 転送など），電子署名を行ったり署名検証を行ったりする前にテキストを正規化しているのである。

### 署名データに署名対象のデータを含める

[GnuPG] ではもうひとつ電子署名の形式がある。
電子署名データの中に署名対象のデータを埋め込んでしまうのである。
これを行うには `--sign` コマンドを使う。
短縮名は `-s`。

```text
$ gpg -u alice -s hello.txt
```

この結果 `hello.txt` ファイルと同じ場所に `hello.txt.gpg` ファイルが作成される。
中身はバイナリデータである。

さて，できたファイルを検証しよう。

```text
$ gpg --verify hello.txt.gpg
gpg: 11/30/17 19:21:22 東京 (標準時)に施された署名
gpg:                RSA鍵0200ACBA190D0FB1D4EC7C02DC1D13A6CCA86E21を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

検証は OK だが署名対象のデータが取り出せない。
そこでまた `--decrypt` コマンドを使う。。

```text
$ gpg -d hello.txt.gpg
Hello world
gpg: 11/30/17 19:21:22 東京 (標準時)に施された署名
gpg:                RSA鍵0200ACBA190D0FB1D4EC7C02DC1D13A6CCA86E21を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

これでデータの抽出ができた。
署名対象のデータを埋め込む方式の何が嬉しいかというと，暗号化と組み合わせる事ができるのである。

### 暗号化と電子署名を同時に行う

暗号化と電子署名を同時に行うには `--encrypt` コマンドと`--sign` コマンドを同時に指定する。
短縮名は `-se` または `-es`。

```text
$ gpg -u alice -r bob -se hello.txt
```

ここでは Bob の公開鍵で暗号化して Alice の鍵で電子署名するようにしてみた。
この結果 `hello.txt` ファイルと同じ場所に `hello.txt.gpg` ファイルが作成される。
中身はバイナリデータである。

ではこれを復号してみよう。

```text
$ gpg -d hello.txt.gpg
gpg: 2048-ビットRSA鍵, ID D74C71530446FD66, 日付2017-11-30に暗号化されました
      "Bob <bob@example.com>"
Hello world
gpg: 11/30/17 19:34:06 東京 (標準時)に施された署名
gpg:                RSA鍵0200ACBA190D0FB1D4EC7C02DC1D13A6CCA86E21を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

このように復号と署名検証が同時に行われる。

ちなみにセッション鍵のみの暗号化と電子署名を組み合わせることもできる。

```text
$ gpg -u alice -sc hello.txt
gpg: AES暗号化を使用します

$ gpg -d hello.txt.gpg
gpg: AES暗号化済みデータ
gpg: 1 個のパスフレーズで暗号化
Hello world
gpg: 11/30/17 19:47:57 東京 (標準時)に施された署名
gpg:                RSA鍵0200ACBA190D0FB1D4EC7C02DC1D13A6CCA86E21を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

署名・暗号化ではパスフレーズ入力が最大3回（暗号化で確認を入れて2回，電子署名で1回）発生するので注意すること。

## 鍵の失効

パスフレーズの漏洩や暗号アルゴリズムの危殆化などによって鍵を失効しなければならない場合がある。

鍵を作成する際に鍵束フォルダの `openpgp-revocs.d` フォルダに失効証明書が作成される。
中身はこんな感じ。

```text
これは失効証明書でこちらのOpenPGP鍵に対するものです:

pub   rsa2048 2017-11-30 [SC]
      0200ACBA190D0FB1D4EC7C02DC1D13A6CCA86E21
uid          Alice <alice@example.com>

失効証明書は "殺すスイッチ" のようなもので、鍵がそれ以上使えない
ように公に宣言するものです。一度発行されると、そのような失効証明書は
撤回することはできません。

秘密鍵のコンプロマイズや紛失の場合、これを使ってこの鍵を失効させます。
しかし、秘密鍵がまだアクセス可能である場合、新しい失効証明書を生成し、
失効の理由をつける方がよいでしょう。詳細は、GnuPGマニュアルのgpgコマン
ド "--generate-revocation"の記述をご覧ください。

このファイルを誤って使うのを避けるため、以下ではコロンが5つのダッシュ
の前に挿入されます。この失効証明書をインポートして公開する前に、テク
スト・エディタでこのコロンを削除してください。

:-----BEGIN PGP PUBLIC KEY BLOCK-----
Comment: This is a revocation certificate

iQE2BCABCAAgFiEEAgCsuhkND7HU7HwC3B0TpsyobiEFAlof25ICHQAACgkQ3B0T
psyobiHRgwf/cNwI01IlXP1dw6op6IgIv3r8nT9XXU4S1WjCvT7yoNs0u+BLHELU
1V16vY9FcnaiNzz/xkSaAVpY+X1O1G7RZ7oYUMA6yMmeUH2fdP7eh4RFM2RZtlq+
HQAoyJb6PVu3uIsfqZh2uMH5v3cUIpRI0dwAZG9hQkg0uZ2a1SGKuSjN9voC9vsE
T55v2WSAtOeleMsNxmywcYGGQBm8YV1F8AC+7K5oc+dmciTBX1IpVHMHkxccObfy
yrpaQGEWJ39Bp8aR+W6Ywe2Bcpbz1tKWmXmXh4iMYEXDBqs/tnpA30dWJYAiLdCA
OYcNJtm9leku3UYJGiTSlxZWmImOEgT8ng==
=xiB6
-----END PGP PUBLIC KEY BLOCK-----
```

このファイルをインポートすることで鍵が失効される。
なお失効証明書を使用の際には

```text
:-----BEGIN PGP PUBLIC KEY BLOCK-----
```

の先頭のコロン（`:`）を削除して使うこと。

```text
$ gpg --import openpgp-revocs.d/0200ACBA190D0FB1D4EC7C02DC1D13A6CCA86E21.rev
gpg: 鍵DC1D13A6CCA86E21:"Alice <alice@example.com>"失効証明書をインポートしました
gpg:           処理数の合計: 1
gpg:         新しい鍵の失効: 1

$ gpg -k alice
pub   rsa2048 2017-11-30 [SC] [失効: 2017-11-30]
      0200ACBA190D0FB1D4EC7C02DC1D13A6CCA86E21
uid           [  失効  ] Alice <alice@example.com>

$ gpg -a --export alice > alice-rev.asc
```

**失効した公開鍵を配布するのを忘れずに！**

失効証明書は `--generate-revocation` コマンドで作成することもできる。
短縮名は `--gen-revoke`。

```test
$ gpg --gen-revoke alice

sec  rsa2048/F3B15FCBA57934CF 2017-11-30 Alice <alice@example.com>

この鍵に対する失効証明書を作成しますか? (y/N) y
失効の理由を選択してください:
  0 = 理由は指定されていません
  1 = 鍵(の信頼性)が損なわれています
  2 = 鍵がとりかわっています
  3 = 鍵はもはや使われていません
  Q = キャンセル
(ここではたぶん1を選びたいでしょう)
あなたの決定は? 1
予備の説明を入力。空行で終了:
>
失効理由: 鍵(の信頼性)が損なわれています
(説明はありません)
よろしいですか? (y/N) y
ASCII外装出力を強制します。
gpg: AllowSetForegroundWindow(11408) failed: アクセスが拒否されました。

-----BEGIN PGP PUBLIC KEY BLOCK-----
Comment: This is a revocation certificate

iQE2BCABCAAgFiEErVfNdsSCr6H+LTqu87Ffy6V5NM8FAlof5W0CHQIACgkQ87Ff
y6V5NM+rwgf/dhNTJlYaDdt52CkS8ckSjhrwK3t56mei+sXaic89mYG6RZsJJeAg
+/KAZbruQZqcYAYYw9jOM0UZpysBvZRRfHj7v44FbcJX7GJORDv3lgtQ0nANwHVN
DXzjpuxBTXGHkBKaOkJ/K5FKGxzFCg+uxJbFh8S710UgS7eg499X+wuKUYuC5orT
n8qdTvehxLf6hfznCA8fgkSP4VFh1X9NWXBcuH1kogAdOTfTcveY/qC2km/i4SfY
6x/s4pQvwAIS682dGaqXro0pODsi5Am43xIZeOJaNui7Ear98zB6S/I0Cbp/knzr
kAc/Jx5aYcyrXqcZtxNwHF+oflpRWyd0KA==
=wLMC
-----END PGP PUBLIC KEY BLOCK-----
失効証明書を作成しました。

みつからないように隠せるような媒体に移してください。もし_悪者_がこの証明書への
アクセスを得ると、あなたの鍵を使えなくすることができます。
媒体が読出し不能になった場合に備えて、この証明書を印刷して保管するのが賢明です。
しかし、ご注意ください。あなたのマシンの印字システムは、他の人がアクセスできる
場所にデータをおくことがあります!
```

このうち ASCII Armor 形式の部分をコピペして使えばよい。

鍵作成時に作られた失効証明書は別の場所に補完しておくことをお勧めする。
もし失効が必要になった時に時間的な余裕があれば `--generate-revocation` コマンドで失効証明書を（失効理由も含める形で）作成し，即失効，配布を行うのがいいと思う。

## ブックマーク

- [Using the GNU Privacy Guard: OpenPGP Key Management](https://www.gnupg.org/documentation/manuals/gnupg/OpenPGP-Key-Management.html)
- [Using the GNU Privacy Guard: Unattended GPG key generation](https://www.gnupg.org/documentation/manuals/gnupg/Unattended-GPG-key-generation.html)
- [Using the GNU Privacy Guard: GPG Configuration Options](https://www.gnupg.org/documentation/manuals/gnupg/GPG-Configuration-Options.html)

- [わかる！ OpenPGP 暗号 — Baldanders.info](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)
- [OpenPGP 鍵管理に関する考察]({{< relref "remark/2017/11/openpgp-key-management.md" >}})

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
