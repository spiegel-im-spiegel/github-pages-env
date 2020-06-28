+++
title = "GnuPG チートシート（鍵作成から失効まで）"
date =  "2017-12-01T17:51:18+09:00"
description = "ちうわけで GnuPG の使い方に関する簡単な「虎の巻（cheat sheet）」を作ってみることにした。"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近 `git commit` に電子署名する目的などで [GnuPG] を使う記事などをチラホラ見かけるようになったが，やっぱ使い慣れんもんは分からんよねぇ。
しかもバージョンによって微妙に挙動が異なるのが困りものである。

ちうわけで [GnuPG] の使い方に関する簡単な「虎の巻（cheat sheet）」を作ってみることにした。
対象となる  [GnuPG] のバージョンは最新版の 2.2.x とする。

なお，この記事は大変長文なので，あらかじめお茶菓子などを用意した上で読みはじめることをお勧めする。
また Qiita に簡易版を公開した。

- [GnuPG チートシート（簡易版） - Qiita](https://qiita.com/spiegel-im-spiegel/items/079d69282166281eb946)

説明はいいから例示だけ見せろという方はこちらで。

では，ご笑覧あれ。

## 目次 {#toc}

1. [はじめに]({{< relref "#intro" >}})
    - [GnuPG とは]({{< relref "#about-gpg" >}})
    - [コマンドとオプション]({{< relref "#cmd-opt" >}})
1. [鍵の作成]({{< relref "#create" >}})
    - [`--generate-key` コマンド]({{< relref "#generate-key" >}})
    - [`--full-generate-key` コマンド]({{< relref "#full-generate-key" >}})
    - [`--quick-generate-key` コマンド]({{< relref "#quick-generate-key" >}})
    - [パスフレーズ入力の回避]({{< relref "#passphrase" >}})
    - [`--quick-add-key` コマンドによる副鍵の追加]({{< relref "#quick-add-key" >}})
1. [鍵の管理]({{< relref "#manage" >}})
    - [鍵束内の公開鍵の検索]({{< relref "#list-keys" >}})
    - [副鍵の鍵指紋の表示]({{< relref "#fingerprint" >}})
    - [パスフレーズの変更]({{< relref "#change-passphrase" >}})
    - [有効期限の変更]({{< relref "#quick-set-expire" >}})
    - [公開鍵をエクスポートする]({{< relref "#export" >}})
    - [公開鍵をインポートする]({{< relref "#import" >}})
    - [公開鍵を鍵サーバに送信する]({{< relref "#send-keys" >}})
    - [公開鍵を鍵サーバから受信する]({{< relref "#receive-keys" >}})
    - [公開鍵に署名する]({{< relref "#sign-key" >}})
    - [`--edit-key` コマンド]({{< relref "#edit-key" >}})
1. [データの暗号化]({{< relref "#encrypt" >}})
    - [ハイブリッド暗号]({{< relref "#hybrid" >}})
    - [セッション鍵のみで暗号化する]({{< relref "#symmetric" >}})
1. [暗号データの復号]({{< relref "#decrypt" >}})
1. [データへの電子署名と検証]({{< relref "#sign" >}})
    - [クリア署名]({{< relref "#clear-sign" >}})
    - [分離署名]({{< relref "#detach-sign" >}})
    - [署名データに署名対象のデータを含める]({{< relref "#data-in-sign" >}})
    - [暗号化と電子署名を同時に行う]({{< relref "#encrypt-and-sign" >}})
1. [鍵の失効]({{< relref "#revocs" >}})
    - [`--generate-revocation` コマンド]({{< relref "#gen-revoke" >}})

## はじめに {#intro}

### [GnuPG] とは {#about-gpg}

[GnuPG (GNU Privacy Guard)][GnuPG] は [OpenPGP] 実装のひとつで，現在は以下の RFC に準拠している。

- [RFC 4880 - OpenPGP Message Format][RFC 4880]
- [RFC 5581 - The Camellia Cipher in OpenPGP][RFC 5581]
- [RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP][RFC 6637]

また [RFC 4880] の後継として現在ドラフト案が公開されている [RFC 4880bis] の一部も取り込んでいる。
更にいうと，今回は言及しないが [OpenPGP] 以外にも X.509 や OpenSSH の鍵管理も可能である。

[OpenPGP] の源流は Phil Zimmermann 氏によって1991年に発表された暗号ツール PGP にまで遡る。
当時の PGP の仕様は [RFC 1991] として残されている。
ただ Phil Zimmermann 氏が暗号特許について迂闊だったため PGP は長く不遇な扱いを受けることになる。

当時の PGP の仕様を知財的に安全なものとして再構成したのが [OpenPGP] である。
更に，その実装である [GnuPG] は GNU プロジェクトの一部として [FSF (Free Software Foundation)](https://www.fsf.org/) が著作権を保有し GPLv3 の下にライセンスされている。

```text
$ gpg --version
gpg (GnuPG) 2.2.19
libgcrypt 1.8.5
Copyright (C) 2019 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: /home/username/.gnupg
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256,
      TWOFISH, CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

ちなみに [OpenPGP] の実装としては他にも [OpenPGP.js](http://openpgpjs.org/ "OpenPGP.js | OpenPGP JavaScript Implementation") や [Sequoia-PGP](https://sequoia-pgp.org/) など様々なものがある。
現在の PGP も [OpenPGP] 実装のひとつとして Symantec 社が[販売](https://www.symantec.com/products/encryption)している。

### コマンドとオプション {#cmd-opt}

[GnuPG] は基本的にコマンドラインツールだが，作りが古くて（なんせ初期の PGP の UI を引きずってるので`w`）今時あたり前な「サブコマンド」みたいな構成になっていない。
その代わりオプションの種別が「コマンド」と「オプション」に分かれている。
具体的には `gpg -h` でヘルプを見ると分かる（もちろん `-h` オプションもコマンドである）。

以上を踏まえて，そろそろ本題に入ろう。

## 鍵の作成 {#create}

鍵の作成コマンドにはいくつか種類がある。

### --generate-key コマンド {#generate-key}

`--generate-key` コマンドは対話モードで鍵の作成を行う。
短縮名は `--gen-key`。
あんまり短縮されていないな（笑）

コマンド自体は初期バージョンから存在するが，バージョンによって挙動がかなり違うので要注意だ。

```text
$ gpg --gen-key
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

注意: 全機能の鍵生成には "gpg --full-generate-key" を使います。

GnuPGはあなたの鍵を識別するためにユーザIDを構成する必要があります。

本名: 
```

最新版 2.2.x では暗号アルゴリズムは RSA 3072bit，有効期限は作成日当日で固定されている。
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
gpg: 鍵AEE3D12F6D2F7F92を究極的に信用するよう記録しました
gpg: 失効証明書を '/home/username/.gnupg/openpgp-revocs.d/0F99F15C200B30194855B93AAEE3D12F6D2F7F92.rev' に保管しました。
gpg: done

$ gpg --list-keys alice
pub   rsa3072 2020-06-27 [SC]
      0F99F15C200B30194855B93AAEE3D12F6D2F7F92
uid           [  究極  ] Alice <alice@example.com>
sub   rsa3072 2020-06-27 [E]
```

設定ファイルの書き方は "[Unattended GPG key generation](https://www.gnupg.org/documentation/manuals/gnupg/Unattended-GPG-key-generation.html "Using the GNU Privacy Guard: Unattended GPG key generation")” を参照のこと。
いったん設定ファイルを作ってしまえばこの方法が一番簡単かな。

なお `Passphrase` の項目を削除すれば `--batch` モードでも Pinentry で設定するパスフレーズを訊いてくるので「設定ファイルにパスフレーズを書くのは...」という方も安心である。

### --full-generate-key コマンド {#full-generate-key}

対話モードで暗号アルゴリズムや鍵長を指定したい場合は `--full-generate-key` コマンドを使う。
短縮名は `--full-gen-key`。

具体的にはこんな感じで進行する。

```text
$ gpg --full-gen-key
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

ご希望の鍵の種類を選択してください:
   (1) RSA と RSA (デフォルト)
   (2) DSA と Elgamal
   (3) DSA (署名のみ)
   (4) RSA (署名のみ)
  (14) カードに存在する鍵
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
鍵の有効期間は? (0) 1y
鍵は2021年06月27日 09時18分48秒 JSTで期限切れとなります
これで正しいですか? (y/N) 
```

`--expert` オプションを付けると選択可能なアルゴリズムの組み合わせが増える。
こんな感じ。

```text
$ gpg --full-gen-key --expert
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
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
  (14) カードに存在する鍵
あなたの選択は? 
```

ECC (Elliptic Curve Cryptography; 楕円曲線暗号) 鍵の取り扱いについては以下の記事を参照のこと。

- [そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{<ref "/openpgp/using-ecc-with-gnupg.md" >}})

### --quick-generate-key コマンド {#quick-generate-key}

コマンドライン一発で鍵を作成したい場合は `--quick-generate-key` コマンドでユーザID，アルゴリズム，有効期限を指定できる。
短縮名は `--quick-gen-key`。

```text
Usage: gpg [options] --quick-generate-key user-id [algo [usage [expire]]]
```

`algo` にはアルゴリズムと鍵長を文字列で指定する。
指定可能な文字列は以下の通り。

| 公開鍵暗号アルゴリズム | 名前 |
|:-----------------------|:-----|
| RSA | `default` (= `rsa3071`) |
| RSA (署名鍵のみ) | `rsa` (= `rsa3071`), `rsa1024`, `rsa2048`, `rsa3071`, `rsa4096` |
| DSA (署名のみ) | `dsa` (= `dsa3072`), `dsa1024`, `dsa2048`, `dsa3072` |
| ECDH/EdDSA | `future-default` (= `cv25519`/`ed25519`) |
| EdDSA (署名鍵のみ) | `ed25519` |
| ECDSA (署名鍵のみ) | `nistp256`, `nistp384`, `nistp521`, `brainpoolP256r1`, `brainpoolP384r1`, `brainpoolP512r1`, `secp256k1` |

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
gpg --quick-gen-key "Alice <alice@example.com>" default default 0
たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。
たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。
gpg: 鍵732A737992A2A7A5を究極的に信用するよう記録しました
gpg: 失効証明書を '/home/username/.gnupg/openpgp-revocs.d/F3546931BCBECA3D73E5E60E732A737992A2A7A5.rev' に保管しました。
公開鍵と秘密鍵を作成し、署名しました。

pub   rsa3072 2020-06-27 [SC]
     F3546931BCBECA3D73E5E60E732A737992A2A7A5
uid                      Alice <alice@example.com>
sub   rsa3072 2020-06-27 [E]
```

### パスフレーズ入力の回避 {#passphrase}

`--quick-generate-key` コマンドでもパスフレーズの入力は Pinentry から行うことになるが `--pinentry-mode` オプションおよび `--passphrase` オプションを付加することで回避できる。

```text
$ gpg --pinentry-mode loopback --passphrase passwd --quick-gen-key "Alice <alice@example.com>" default default 0
```

ただしコマンドラインの履歴に入力したパスフレーズが残ってしまうのであまりお勧めできないが...

### --quick-add-key コマンドによる副鍵の追加 {#quick-add-key}

作成した鍵に `--quick-add-key` コマンドで後から暗号鍵を追加できる。
これは ` --quick-generate-key` コマンドで主鍵のみ作って後から副鍵を加えたい場合などに有効である。

```text
Usage: gpg [options] --quick-add-key key-fingerprint [algo [usage [expire]]]
```

たとえば，以下の鍵に対して

```text
$ gpg --list-keys alice
pub   dsa3072 2020-06-27 [SC] [有効期限: 2021-06-27]
      B8AF88FC3448859FCF65810C20EDB41D5093CA8A
uid           [  究極  ] Alice <alice@example.com>
```

以下のように暗号鍵を副鍵として追加できる。

```text
$ gpg --quick-add-key B8AF88FC3448859FCF65810C20EDB41D5093CA8A elg3072 encr
たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。

$ gpg --list-keys alice
pub   dsa3072 2020-06-27 [SC] [有効期限: 2021-06-27]
      B8AF88FC3448859FCF65810C20EDB41D5093CA8A
uid           [  究極  ] Alice <alice@example.com>
sub   elg3072 2020-06-27 [E]
```

ちなみに `B8AF88FC3448859FCF65810C20EDB41D5093CA8A` という長ったらしい数字列は鍵指紋（key fingerprint）である。
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

## 鍵の管理 {#manage}

作成した鍵や配布・受領した公開鍵を管理するためのコマンドを紹介する。

### 鍵束内の公開鍵の検索 {#list-keys}

鍵束内の公開鍵を検索する場合， `--list-keys` コマンドの引数にユーザID（の一部）または鍵IDを指定することで，条件にマッチする鍵を検索できる。
短縮名は `-k`。

```text
$ gpg -k alice
pub   dsa3072 2020-06-27 [SC] [有効期限: 2021-06-27]
      B8AF88FC3448859FCF65810C20EDB41D5093CA8A
uid           [  究極  ] Alice <alice@example.com>
sub   elg3072 2020-06-27 [E]
```

ちなみに `[SC]` や `[E]` は鍵の機能を示していて， `C` は証明， `S` は署名， `E` は暗号化を指す。

引数なしで `--list-keys` コマンドを起動した場合は公開鍵の鍵束（`pubring.kbx`）にある鍵が全て列挙される。

秘密鍵を検索する場合には `--list-secret-keys` コマンドを使う。
短縮名は大文字の `-K`。

```text
$ gpg -K alice
sec   dsa3072 2020-06-27 [SC] [有効期限: 2021-06-27]
      B8AF88FC3448859FCF65810C20EDB41D5093CA8A
uid           [  究極  ] Alice <alice@example.com>
ssb   elg3072 2020-06-27 [E]
```

### 副鍵の鍵指紋の表示 {#fingerprint}

`--list-keys` コマンドでも主鍵の鍵指紋が表示されるが，副鍵の鍵指紋も表示したい場合は `--fingerprint` コマンドを2つ重ねる。

```text
$ gpg --fingerprint --fingerprint alice
pub   dsa3072 2020-06-27 [SC] [有効期限: 2021-06-27]
      B8AF 88FC 3448 859F CF65  810C 20ED B41D 5093 CA8A
uid           [  究極  ] Alice <alice@example.com>
sub   elg3072 2020-06-27 [E]
      7BB9 F625 7053 4A0F B297  ECA6 7792 0F61 EFE7 15CC
```

### パスフレーズの変更 {#change-passphrase}

秘密鍵のパスフレーズを変更する場合には `--change-passphrase` コマンドを使う。
短縮名は `--passwd`。

```text
$ gpg --passwd alice
```

引数にはユーザID（の一部）を指定できる。
パスワードの入力は Pinentry で行う。

### 有効期限の変更 {#quick-set-expire}

自身の鍵の有効期限を変更する場合には `--quick-set-expire` コマンドを使う。

鍵の鍵指紋がが以下の場合

```text
$ gpg --fingerprint --fingerprint alice
pub   dsa3072 2020-06-27 [SC] [有効期限: 2021-06-27]
      B8AF 88FC 3448 859F CF65  810C 20ED B41D 5093 CA8A
uid           [  究極  ] Alice <alice@example.com>
sub   elg3072 2020-06-27 [E]
      7BB9 F625 7053 4A0F B297  ECA6 7792 0F61 EFE7 15CC
```

有効期限を2年（`2y`）に指定するなら，操作は以下の通り。

```text
$ gpg --quick-set-expire B8AF88FC3448859FCF65810C20EDB41D5093CA8A 2y

$ gpg --list-keys alice
pub   dsa3072 2020-06-27 [SC] [有効期限: 2022-06-27]
      B8AF88FC3448859FCF65810C20EDB41D5093CA8A
uid           [  究極  ] Alice <alice@example.com>
sub   elg3072 2020-06-27 [E]

$ gpg --quick-set-expire B8AF88FC3448859FCF65810C20EDB41D5093CA8A 2y 7BB9F62570534A0FB297ECA677920F61EFE715CC

$ gpg --list-keys alice
pub   dsa3072 2020-06-27 [SC] [有効期限: 2022-06-27]
      B8AF88FC3448859FCF65810C20EDB41D5093CA8A
uid           [  究極  ] Alice <alice@example.com>
sub   elg3072 2020-06-27 [E] [有効期限: 2022-06-27]
```

前半は主鍵，後半は（主鍵に紐づく）副鍵の有効期限を変更している。
このことから分かるとおり，主鍵と副鍵は個別に有効期限を設定することが可能である。

### 公開鍵をエクスポートする {#export}

公開鍵のエクスポートには `--export` コマンドを使う。

```text
$ gpg --armor --export alice
-----BEGIN PGP PUBLIC KEY BLOCK-----

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
=nDyQ
-----END PGP PUBLIC KEY BLOCK-----
```

`--armor` オプションを指定すると，上記のように， ASCII Armor 形式のテキストを出力する。
短縮名は `-a`。
`--armor` オプションを付けないとバイナリを吐く。

秘密鍵をエクスポートする場合は `--export-secret-key` コマンドを使う（パスフレーズ入力あり）。

公開鍵をファイル等で配布する場合は `--export` コマンドの出力をファイルにリダイレクトすればよい。

```text
$ gpg -a --export alice > alice-key.asc
```

ASCII Armor 形式のテキスト・ファイルの拡張子は慣習的に `.asc` とすることが多い。

### 公開鍵をインポートする {#import}

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
以下は [JPCERT/CC の公開鍵](https://www.jpcert.or.jp/jpcert-pgp.html "JPCERT コーディネーションセンター PGP公開鍵")をサイトからインポートしたところ。

```text
$ gpg --fetch-keys https://www.jpcert.or.jp/keys/info-0x69ECE048.asc
```

インポートする鍵が既に鍵束にある場合でも，単純な上書きではなく， [GnuPG] がいい感じにマージしてくれるのでご安心を。

### 公開鍵を鍵サーバに送信する {#send-keys}

鍵束にある公開鍵を鍵サーバに送信するには `--send-keys` コマンドを使う。

```text
$ gpg --keyserver keys.gnupg.net --send-keys 69ECE048
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

### 公開鍵を鍵サーバから受信する {#receive-keys}

鍵サーバから公開鍵を受信する場合は `--receive-keys` コマンドを使う。
短縮名は `--recv-keys`。

```text
$ gpg --keyserver keys.gnupg.net --recv-keys 69ECE048
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

### 公開鍵に署名する {#sign-key}

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

### --edit-key コマンド {#edit-key}

`--edit-key` コマンドは対話モードで鍵の編集を行う。

```text
$ gpg --edit-key alice
```

対話モードに入るとプロンプトが `gpg>` に変わる。
対話モードで使えるコマンドは以下の通り。

```text
gpg> help
quit        このメニューを終了
save        保存して終了
help        このヘルプを表示
fpr         鍵のフィンガープリントを表示
grip        keygripを表示
list        鍵とユーザIDの一覧
uid         ユーザID Nの選択
key         副鍵Nの選択
check       署名の確認
sign        選択したユーザIDに署名する [* 以下の関連コマンドを参照 ]
lsign       選択したユーザIDにローカルに署名
tsign       選択したユーザIDに信用署名を署名する
nrsign      選択したユーザIDに失効不可の署名をする
adduid      ユーザIDの追加
addphoto    フォトIDの追加
deluid      選択したユーザIDの削除
addkey      副鍵を追加
addcardkey  スマートカードへ鍵の追加
keytocard   鍵をスマートカードへ移動
bkuptocard  バックアップ鍵をスマートカードへ移動
delkey      選択した副鍵の削除
addrevoker  失効鍵の追加
delsig      選択したユーザIDから署名を削除する
expire      鍵または選択した副鍵の有効期限を変更する
primary     選択したユーザIDを主にする
pref        優先指定の一覧 (エキスパート)
showpref    優先指定の一覧 (冗長)
setpref     選択したユーザIDに優先指定リストを設定
keyserver   選択したユーザIDに優先鍵サーバのURLを設定
notation    選択したユーザIDに注釈を設定する
passwd      パスフレーズの変更
trust       所有者信用の変更
revsig      選択したユーザIDの署名を失効
revuid      選択したユーザIDの失効
revkey      鍵の失効または選択した副鍵の失効
enable      鍵を有効にする
disable     鍵を無効にする
showphoto   選択したフォトIDを表示
clean       使えないユーザIDをコンパクトにし、使えない署名を鍵から除去
minimize    使えないユーザIDをコンパクトにし、すべての署名を鍵から除去

* 'sign' コマンドは 'l' で始まると、ローカルの署名で (lsign)、
  't' で始まると信用署名 (tsign)、'nr' で始まると失効不可署名
  (nrsign)、もしくはこれらの組み合わせ (ltsign, tnrsign, など)となります。
```

本記事では，対話モードの詳細については割愛する。
基本的には `list`, `uid`, `key` を使ってでユーザIDや副鍵を指定し，それぞれに対して操作を行う感じ。
編集結果を保存して終了するには `save` を，単に対話モードから抜けるには `quit` を入力する。

## データの暗号化 {#encrypt}

[GnuPG] の暗号化は概ね2種類ある。

### ハイブリッド暗号 {#hybrid}

ハイブリッド暗号は [OpenPGP] の基本機能で，平文を暗号化する「セッション鍵」とセッション鍵を暗号化する公開鍵で構成される。

{{< fig-img src="https://baldanders.info/spiegel/openpgp/hybrid-enc.svg" width="715" title="「わかる！ OpenPGP 暗号」より" link="https://baldanders.info/spiegel/openpgp/" >}}

セッション鍵は [GnuPG] が自動的にし生成するのでコマンドラインではセッション鍵を暗号化する公開鍵を指定する。

暗号化を行うには `--encrypt` コマンドを使う。
短縮名は `-e`。

```text
$ gpg -a --recipient alice -e plain-data
-----BEGIN PGP MESSAGE-----

hQMOA3eSD2Hv5xXMEAwAkpcgXFe/zIlUyibCRLxQAK5KNYvzBxcoaxyk0lEsoUGt
D7F1bLF64ha9J/FuzN3B63LsmqTw1I6tBIjq2OrSMvLAN10B6bfUtclZGRnK/H8J
oRMYQUuteUUOcVR2v3MjBkRDrxMQDBfJwIx5N13R4bJrD/k+unfwo82JqiT/LuET
cO7KLgiyf+mBpu1ln90QfYsNtQIZMOPdxutqsIAnZxM0gFjKLBmHBhBG85v4QonS
KUqGZfPzxQfDW3t6gMfbkeDNqnldDMJyY0nXqdkaSznAXIJAT9TU7D2NlUaF7C8p
51O5NnK/+8NYSWsrzkcL+taw3i83jcZFYzcI2oqjuC6wxrNW9SQ1Jr6QqNSanWot
ieLtMGnfxXizEhYaZfiU+y2ebZFxOlS1SigPc+CfEiKKYlQ0DrnAyjxkVKFtp5oG
UwA3n1rnsJZO7m7vwnpYlu25nlkM5y+AqjT6zqtPACsjBzzWG8nsMq6kMWvKYva/
Fj0oDxtMJw2jAoIGITKzC/94tFTxs1WoLrMl1Lgjrr07xLqfp0mzFL1cBLK2YaRn
rB2w0MOYFFjyMxPXPx4sUv22U8bfz3Vu/bD7I+0Jea6CUKzDjemdZlxYENPDvrX3
XoKUEwdslrq98DdhH7xcBFC64gv2DyyP6TSpnkFb8sV3DUnQ+RiFkYZN0HKaqHx6
EUlzMekSl+wUoxD/sGA7tgkjzYg+VHjdXWiCRpDTDIC9wpInP2il40BBWf9KCLCq
K5hzERr5qepCg0RUs0kgjOjeAngEuAjMrHJJo3iQpiyFJq2Q0b/NAMJiBP0j88MY
uM4sr8DJz1rfvQyLEfI2Satx8CjdXEkSwekxwBUzGSyij5NPJVLYFswTxhLrVUfq
GSZAO1tbD++EU9ViSc0pcsuDNTcWPsba1fcotAVtCIZXqUGCNULuOZdOX4n/1oSt
eu2sWmgrF8z1qTCobcYSiC8F1t9DkbOtot49izVqbE8uRkZ8U1kn5VbDgPCGVrtU
DmOsVsgqcj8+MioyTDVMYqXSRwFLJEbBJJegDyjuVdVwSN9lqROPsuqHP/l8mLPs
7KTQ/+Vi94pFitZmp7jrMc9BgDQwZgox81uNI2c2H9UsEl60sKqWYyIk
=wqIs
-----END PGP MESSAGE-----
```

あるいはパイプを使って

```text
$ echo Hello world | gpg -a --recipient alice -e
-----BEGIN PGP MESSAGE-----

hQMOA3eSD2Hv5xXMEAwAkpcgXFe/zIlUyibCRLxQAK5KNYvzBxcoaxyk0lEsoUGt
D7F1bLF64ha9J/FuzN3B63LsmqTw1I6tBIjq2OrSMvLAN10B6bfUtclZGRnK/H8J
oRMYQUuteUUOcVR2v3MjBkRDrxMQDBfJwIx5N13R4bJrD/k+unfwo82JqiT/LuET
cO7KLgiyf+mBpu1ln90QfYsNtQIZMOPdxutqsIAnZxM0gFjKLBmHBhBG85v4QonS
KUqGZfPzxQfDW3t6gMfbkeDNqnldDMJyY0nXqdkaSznAXIJAT9TU7D2NlUaF7C8p
51O5NnK/+8NYSWsrzkcL+taw3i83jcZFYzcI2oqjuC6wxrNW9SQ1Jr6QqNSanWot
ieLtMGnfxXizEhYaZfiU+y2ebZFxOlS1SigPc+CfEiKKYlQ0DrnAyjxkVKFtp5oG
UwA3n1rnsJZO7m7vwnpYlu25nlkM5y+AqjT6zqtPACsjBzzWG8nsMq6kMWvKYva/
Fj0oDxtMJw2jAoIGITKzC/94tFTxs1WoLrMl1Lgjrr07xLqfp0mzFL1cBLK2YaRn
rB2w0MOYFFjyMxPXPx4sUv22U8bfz3Vu/bD7I+0Jea6CUKzDjemdZlxYENPDvrX3
XoKUEwdslrq98DdhH7xcBFC64gv2DyyP6TSpnkFb8sV3DUnQ+RiFkYZN0HKaqHx6
EUlzMekSl+wUoxD/sGA7tgkjzYg+VHjdXWiCRpDTDIC9wpInP2il40BBWf9KCLCq
K5hzERr5qepCg0RUs0kgjOjeAngEuAjMrHJJo3iQpiyFJq2Q0b/NAMJiBP0j88MY
uM4sr8DJz1rfvQyLEfI2Satx8CjdXEkSwekxwBUzGSyij5NPJVLYFswTxhLrVUfq
GSZAO1tbD++EU9ViSc0pcsuDNTcWPsba1fcotAVtCIZXqUGCNULuOZdOX4n/1oSt
eu2sWmgrF8z1qTCobcYSiC8F1t9DkbOtot49izVqbE8uRkZ8U1kn5VbDgPCGVrtU
DmOsVsgqcj8+MioyTDVMYqXSRwFLJEbBJJegDyjuVdVwSN9lqROPsuqHP/l8mLPs
7KTQ/+Vi94pFitZmp7jrMc9BgDQwZgox81uNI2c2H9UsEl60sKqWYyIk
=wqIs
-----END PGP MESSAGE-----
```

とすることもできる。

ちなみに，この暗号データの中身を拙作 [gpgpdump] で覗くとこんな感じになる。

```text
$ echo Hello world | gpg -a --recipient alice -e | gpgpdump --indent 2
Public-Key Encrypted Session Key Packet (tag 1) (782 bytes)
  Version: 3 (current)
  Key ID: 0x77920f61efe715cc
  Public-key Algorithm: Elgamal (Encrypt-Only) (pub 16)
  ElGamal g^k mod p (3072 bits)
  ElGamal m * y^k mod p (3071 bits)
Sym. Encrypted Integrity Protected Data Packet (tag 18) (71 bytes)
  Encrypted data: sym alg is specified in pub-key encrypted session key (71 bytes)
```

`--recipient` がセッション鍵の暗号化を行う公開鍵を指定するオプションである。
短縮名は `-r`。
`--recipient` オプションは複数指定できる。
また鍵束フォルダにある `gpg.conf` ファイルで常に使用する公開鍵を指定することもできる[^dr1]。

[^dr1]: `default-recipient-self` の指定は自身の鍵で復号できるよう設定するためのものである。相手の公開鍵のみで暗号化してしまうと，暗号化した本人が復号できないことになってしまうため。

```text
default-key alice
default-recipient-self
```

### セッション鍵のみで暗号化する {#symmetric}

公開鍵は使わずパスフレーズから生成したセッション鍵のみで暗号化を行う場合は `--symmetric` コマンドを使う。
短縮名は `-c`。

```text
$ echo Hello world | gpg -a -c
-----BEGIN PGP MESSAGE-----

jA0EBwMCgMkCBCmvi5D/0kEBHskD5D+oRASdhy04sJQ6EX6hObuqT170+wr7wB4f
XTDWFf7c6XXN0LABKN6z3kYxz2MTRC3cqo2ExRF30RHpeA==
=xvgI
-----END PGP MESSAGE-----
```

コマンド起動時にパスフレーズの入力を要求され，パスフレーズからセッション鍵を生成して暗号化を行う。
したがって，何らかの方法で暗号データの受け手とパスフレーズを共有する必要がある。

この暗号データの中身を拙作 [gpgpdump] で覗くとこんな感じになる。

```text
$ echo Hello world | gpg -a -c | gpgpdump --indent 2
Symmetric-Key Encrypted Session Key Packet (tag 3) (13 bytes)
  Version: 4 (current)
  Symmetric Algorithm: AES with 128-bit key (sym 7)
  String-to-Key (S2K) Algorithm: Iterated and Salted S2K (s2k 3)
    Hash Algorithm: SHA-1 (hash 2)
    Salt
      80 c9 02 04 29 af 8b 90
    Count: 65011712
      ff
Sym. Encrypted Integrity Protected Data Packet (tag 18) (65 bytes)
  Encrypted data: sym alg is specified in sym-key encrypted session key; plain text + MDC SHA1(20 bytes) (65 bytes)
```

## 暗号データの復号 {#decrypt}

暗号データのf区号には `--decrypt` コマンドを使う（パスフレーズ入力あり）。
短縮名は `-d`。

```text
$ echo Hello world | gpg -r alice -e -a > alice-enc.asc

$ gpg -d alice-enc.asc
gpg: 3072-ビットELG鍵, ID 77920F61EFE715CC, 日付2020-06-27に暗号化されました
      "Alice <alice@example.com>"
Hello world
```

復号したデータはファイルにリダイレクトすればいいのだが， Windows の場合は安全のため `--output` オプションを使うことをお勧めする。
短縮名は `-o`。

```text
$ gpg -o out.txt -d alice-enc.asc
gpg: 3072-ビットELG鍵, ID 77920F61EFE715CC, 日付2020-06-27に暗号化されました
      "Alice <alice@example.com>"

$ cat out.txt
Hello world
```

セッション鍵のみで暗号化した場合も同じコマンドで復号できる（パスフレーズ入力あり）。

```text
$ echo Hello world | gpg -a -c > alice-sym-enc.asc

$ gpg -d alice-sym-enc.asc
gpg: AES暗号化済みデータ
gpg: 1 個のパスフレーズで暗号化
Hello world
```

## データへの電子署名と検証 {#sign}

データへの電子署名にも幾つかの方法がある。

### クリア署名 {#clear-sign}

まずデータがテキストの場合は「クリア署名」という方法が使える。
クリア署名には `--clear-sign` コマンドを使う。

```text
$ echo Hello world | gpg -u alice --clear-sign
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA256

Hello world
-----BEGIN PGP SIGNATURE-----

iIgEAREIADAWIQS4r4j8NEiFn89lgQwg7bQdUJPKigUCXvaY4hIcYWxpY2VAZXhh
bXBsZS5jb20ACgkQIO20HVCTyorOTQEAyrm2FVfF3SMA34mh0+B8iRDj+hRr9fL5
WXI/DhtRkFsBAOLTv1riIFb3yeodDX6zUgNkISGWQM1dSLDILeEpVOTt
=XyJW
-----END PGP SIGNATURE-----
```

このようにクリア署名は元になるテキストと電子署名がくっついた状態で出力される。
なお，クリア署名の場合は必ず ASCII Armor 形式の出力になるため `--armor` オプションは不要である。

署名の検証には `--verify` コマンドを使う。
少し横着して署名と検証をパイプで繋いでしまおう。

```text
$ echo Hello world | gpg -u alice --clear-sign | gpg --verify
gpg: 2020年06月27日 09時55分43秒 JSTに施された署名
gpg:                DSA鍵B8AF88FC3448859FCF65810C20EDB41D5093CA8Aを使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

まぁ自分で署名して，出力をそのまま自分で検証してるんだから正しくてあたり前なのだが，流れは分かると思う。

実は `--verify` コマンドは `--decrypt` コマンドで代替えできる。

```text
$ echo Hello world | gpg -u alice --clear-sign | gpg -d
Hello world
gpg: 2020年06月27日 09時56分17秒 JSTに施された署名
gpg:                DSA鍵B8AF88FC3448859FCF65810C20EDB41D5093CA8Aを使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

このように `--decrypt` コマンドを使うと署名対象のテキストを抽出して出力してくれるのが利点である。

### 分離署名 {#detach-sign}

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
gpg: 2020年06月27日 09時58分05秒 JSTに施された署名
gpg:                DSA鍵B8AF88FC3448859FCF65810C20EDB41D5093CA8Aを使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

ここでは署名対象のファイルを `hello.txt` と推測して署名の検証を行っている。
署名対象のファイルを明示して指定するには

```text
$ gpg --verify hello.txt.sig hello.txt
gpg: 2020年06月27日 09時58分05秒 JSTに施された署名
gpg:                DSA鍵B8AF88FC3448859FCF65810C20EDB41D5093CA8Aを使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

とする。
分離署名の検証でも `--decrypt` コマンドが使える。
ただし出力は全く同じ。

```text
$  gpg -d hello.txt.sig
gpg: 署名されたデータが'hello.txt'にあると想定します
gpg: 2020年06月27日 09時58分05秒 JSTに施された署名
gpg:                DSA鍵B8AF88FC3448859FCF65810C20EDB41D5093CA8Aを使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

なお，署名対象のファイルがテキスト・ファイルの場合は `--textmode` オプションを付けて電子署名を行ったほうが安全である。

```text
$ gpg -u alice --textmode -b hello.txt
```

テキスト・ファイルの場合，配布経路によっては改行コードが変えられたりするため（電子メールや FTP 転送など），電子署名を行ったり署名検証を行ったりする前にテキストを正規化しているのである。

### 署名データに署名対象のデータを含める {#data-in-sign}

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
gpg: 2020年06月27日 10時00分45秒 JSTに施された署名
gpg:                DSA鍵B8AF88FC3448859FCF65810C20EDB41D5093CA8Aを使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

検証は OK だが署名対象のデータが取り出せない。
そこでまた `--decrypt` コマンドを使う。。

```text
$ gpg -d hello.txt.gpg
Hello world
gpg: 2020年06月27日 10時00分45秒 JSTに施された署名
gpg:                DSA鍵B8AF88FC3448859FCF65810C20EDB41D5093CA8Aを使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

これでデータの抽出ができた。
署名対象のデータを埋め込む方式の何が嬉しいかというと，暗号化と組み合わせる事ができるのである。

### 暗号化と電子署名を同時に行う {#encrypt-and-sign}

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
gpg: 3072-ビットRSA鍵, ID 02F4454D89E9C0D6, 日付2020-06-27に暗号化されました
      "Bob <bob@bob@example.com>"
Hello world
gpg: 2020年06月27日 10時14分56秒 JSTに施された署名
gpg:                DSA鍵B8AF88FC3448859FCF65810C20EDB41D5093CA8Aを使用
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
gpg: 2020年06月27日 10時16分16秒 JSTに施された署名
gpg:                DSA鍵B8AF88FC3448859FCF65810C20EDB41D5093CA8Aを使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

署名・暗号化ではパスフレーズ入力が最大3回（暗号化で確認を入れて2回，電子署名で1回）発生するので注意すること。

## 鍵の失効 {#revocs}

パスフレーズの漏洩や暗号アルゴリズムの危殆化などによって鍵を失効しなければならない場合がある。

鍵を作成する際に鍵束フォルダの `openpgp-revocs.d` フォルダに失効証明書が作成される。
中身はこんな感じ。

```text
これは失効証明書でこちらのOpenPGP鍵に対するものです:

pub   dsa3072 2020-06-27 [S] [有効期限: 2021-06-27]
      B8AF88FC3448859FCF65810C20EDB41D5093CA8A
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

iHgEIBEIACAWIQS4r4j8NEiFn89lgQwg7bQdUJPKigUCXvaS7wIdAAAKCRAg7bQd
UJPKihg/AP9enGjHPngyJLQXJmQhygq7FQ6GmqAmDrvPRwYf5eLLQQD+IAgyE+ho
0nWE84Q2/HosO6a8as+0CFyJWzS1RkXUa5g=
=Zck3
-----END PGP PUBLIC KEY BLOCK-----
```

このファイルをインポートすることで鍵が失効される。
なお失効証明書を使用の際には

```text
:-----BEGIN PGP PUBLIC KEY BLOCK-----
```

の先頭のコロン（`:`）を削除して使うこと。

```text
$ gpg --import openpgp-revocs.d/B8AF88FC3448859FCF65810C20EDB41D5093CA8A.rev
gpg: 鍵20EDB41D5093CA8A:"Alice <alice@example.com>"失効証明書をインポートしました
gpg: 処理数の合計: 1
gpg:    新しい鍵の失効: 1

$ gpg -k alice
pub   dsa3072 2020-06-27 [SC] [失効: 2020-06-27]
      B8AF88FC3448859FCF65810C20EDB41D5093CA8A
uid           [  失効  ] Alice <alice@example.com>

$ gpg --send-keys B8AF88FC3448859FCF65810C20EDB41D5093CA8A
```

{{< div-gen type="markdown" class="center caution" >}}
失効した公開鍵を配布するのを忘れずに！
{{< /div-gen >}}

### --generate-revocation コマンド {#gen-revoke}

失効証明書は `--generate-revocation` コマンドで作成することもできる。
短縮名は `--gen-revoke`。

```test
$ gpg --gen-revoke alice

sec  dsa3072/20EDB41D5093CA8A 2020-06-27 Alice <alice@example.com>

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
> for example
> 
失効理由: 鍵(の信頼性)が損なわれています
for example
よろしいですか? (y/N) y
ASCII外装出力を強制します。
-----BEGIN PGP PUBLIC KEY BLOCK-----
Comment: This is a revocation certificate

iIMEIBEIACsWIQS4r4j8NEiFn89lgQwg7bQdUJPKigUCXvafEQ0dAmZvciBleGFt
cGxlAAoJECDttB1Qk8qKbacA+wd+SaX+khrq6mz4G9NlZio8jcmtHvF5RFncq4wl
UZ/MAP42/Kkuy249M9ZDHQAnTd7Yb3NbuBcphEqJ1iPIZ6niYg==
=AIvD
-----END PGP PUBLIC KEY BLOCK-----
失効証明書を作成しました。

みつからないように隠せるような媒体に移してください。もし_悪者_がこの証明書への
アクセスを得ると、あなたの鍵を使えなくすることができます。
媒体が読出し不能になった場合に備えて、この証明書を印刷して保管するのが賢明です。
しかし、ご注意ください。あなたのマシンの印字システムは、他の人がアクセスできる
場所にデータをおくことがあります!
```

このうち ASCII Armor 形式の部分をコピペして使えばよい。
ちなみに，この失効証明書の中身を拙作 [gpgpdump] で覗くとこんな感じになる。

```text
$ gpgpdump -f rev.txt --indent 2
Signature Packet (tag 2) (131 bytes)
  Version: 4 (current)
  Signiture Type: Key revocation signature (0x20)
  Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
  Hash Algorithm: SHA2-256 (hash 8)
  Hashed Subpacket (43 bytes)
    Issuer Fingerprint (sub 33) (21 bytes)
      Version: 4 (need 20 octets length)
      Fingerprint (20 bytes)
        b8 af 88 fc 34 48 85 9f cf 65 81 0c 20 ed b4 1d 50 93 ca 8a
    Signature Creation Time (sub 2): 2020-06-27T10:21:21+09:00
    Reason for Revocation (sub 29): Key material has been compromised (key revocations) (2) (12 bytes)
      Additional information: for example
  Unhashed Subpacket (10 bytes)
    Issuer (sub 16): 0x20edb41d5093ca8a
  Hash left 2 bytes
    6d a7
  DSA value r (251 bits)
  DSA value s (254 bits)
```

鍵作成時に作られた失効証明書は別の場所に補完しておくことをお勧めする。
もし失効が必要になった時に時間的な余裕があれば `--generate-revocation` コマンドで失効証明書を（失効理由も含める形で）作成し，即失効，配布を行うのがいいと思う。

## ブックマーク

- [Using the GNU Privacy Guard: OpenPGP Key Management](https://www.gnupg.org/documentation/manuals/gnupg/OpenPGP-Key-Management.html)
- [Using the GNU Privacy Guard: Unattended GPG key generation](https://www.gnupg.org/documentation/manuals/gnupg/Unattended-GPG-key-generation.html)
- [Using the GNU Privacy Guard: GPG Configuration Options](https://www.gnupg.org/documentation/manuals/gnupg/GPG-Configuration-Options.html)

- [わかる！ OpenPGP 暗号 — Baldanders.info](https://baldanders.info/spiegel/cc-license/)
- [OpenPGP 鍵管理に関する考察]({{< ref "/remark/2017/11/openpgp-key-management.md" >}})

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[RFC 5581]: https://tools.ietf.org/html/rfc5581 "RFC 5581 - The Camellia Cipher in OpenPGP"
[RFC 6637]: https://tools.ietf.org/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[RFC 1991]: https://tools.ietf.org/html/rfc1991 "RFC 1991 - PGP Message Exchange Formats"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
