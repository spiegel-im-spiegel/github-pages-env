# GnuPG チートシート（簡易版）

[GnuPG] の使い方に関する簡単な「虎の巻（cheat sheet）」を作ってみることにした。
なお詳細版は以下の記事にまとめている。

- [GnuPG チートシート（鍵作成から失効まで） — しっぽのさきっちょ | text.Baldanders.info](https://text.baldanders.info/remark/2017/11/gnupg-sheat-sheet/)

## 鍵の生成

### 対話モードでの作成

```
$ gpg --gen-key
```

最新版 2.2.x では暗号アルゴリズムは RSA/2048bit，有効期限は作成日当日で固定されている。

暗号アルゴリズムや鍵長などを選択したい場合は `--full-gen-key` コマンドを使う。

```
$ gpg --full-gen-key
```

- さらに `--expert` オプションを付けると ECC 暗号が選択可能になる

```
$ gpg --full-gen-key --expert
```

### バッチ処理

`--gene-key` コマンドについては，以下のような設定ファイルを作って `--batch` オプションを付けて起動することで対話モードを回避し，かつアルゴリズム等の詳細な指定をすることもできる。

```
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

### コマンドラインでアルゴリズムを指定

```
Usage: gpg [options] --quick-generate-key user-id [algo [usage [expire]]]
```

実行例はこんな感じ。

```
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

`--quick-gen-key` コマンドでは（`default` 指定以外では）主鍵のみの作成となるので，`--quick-add-key` コマンドで副鍵を追加する。

```
Usage: gpg [options] --quick-add-key key-fingerprint [algo [usage [expire]]]
```

実行例はこんな感じ。

```
$ gpg --list-keys alice
pub   dsa3072 2017-11-30 [SC]
      B5BF56B346B4D961E6BF25A45CC68B4A317E8E5C
uid           [  究極  ] Alice <alice@example.com>

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

## 鍵の管理

### 鍵束内の公開鍵の検索

```
$ gpg -k alice
pub   dsa3072 2017-11-23 [SC]
      3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E
uid           [  充分  ] Alice (root) <alice@example.com>

pub   rsa2048 2017-11-23 [SC]
      A3CEFEEEDA222024F325C403DFFC3F67BBB3C083
uid           [  充分  ] Alice (commit) <alice@example.com>
sub   rsa2048 2017-11-23 [E]
```

秘密鍵を検索する場合には `--list-secret-key` コマンドを使う。

### 副鍵の鍵指紋の表示

`--list-key` コマンドでも主鍵の鍵指紋が表示されるが，副鍵の鍵指紋も表示したい場合は `--fingerprint` コマンドを2つ重ねる。

```
$ gpg --fingerprint --fingerprint alice
pub   rsa2048 2017-11-30 [SC]
      79FD 2B99 F3C6 2D2D 3B85  0BBC 93B3 5094 7582 0D5D
uid           [  究極  ] Alice <alice@example.com>
sub   rsa2048 2017-11-30 [E]
      476E 9EA7 D703 F0BB 01B6  FA44 9278 B060 D202 3C53
```

### パスフレーズの変更

```
$ gpg --passwd alice
```

パスワードの入力は Pinentry で行う。

### 有効期限の変更

```
$ gpg --fingerprint --fingerprint alice
pub   rsa2048 2017-11-30 [SC]
      79FD 2B99 F3C6 2D2D 3B85  0BBC 93B3 5094 7582 0D5D
uid           [  究極  ] Alice <alice@example.com>
sub   rsa2048 2017-11-30 [E]
      476E 9EA7 D703 F0BB 01B6  FA44 9278 B060 D202 3C53

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

### 公開鍵をエクスポートする

```
$ gpg -a --export alice > alice-key.asc
```

### 公開鍵をインポートする

```
$ gpg --import alice-key.asc
```

### 公開鍵を鍵サーバに送信する

```
$ gpg --keyserver keys.gnupg.net --send-keys 7E20B81C
```

鍵束フォルダにある `gpg.conf` ファイルに既定の鍵サーバを指定できる。

```
keyserver  keys.gnupg.net
```

主な鍵サーバ

- [keys.gnupg.net](http://keys.gnupg.net/ "Nebraska Wesleyan University - OpenPGP Keyserver")
- [pgp.mit.edu](https://pgp.mit.edu/ "MIT PGP Key Server")
- [pgp.nic.ad.jp](http://pgp.nic.ad.jp/ "PGP KEYSERVER")

### 公開鍵を鍵サーバから受信する

```
$ gpg --keyserver keys.gnupg.net --recv-keys 7E20B81C
```

あらかじめ鍵IDがわからない場合は `--search-keys` コマンドで検索できる。

```
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

### 公開鍵に署名する

alice の鍵で署名する場合。

```
gpg -u alice --quick-sign-key 1B5202DB4A3EC776F1E0AD18B4DA3BAE7E20B81C
```

電子署名を配布されては困る場合は `--quick-lsign-key` コマンドを使う。

電子署名に使う既定の鍵を鍵束フォルダにある `gpg.conf` ファイルで指定することができる。

```
default-key alice
```

## データの暗号化

### ハイブリッド暗号

alice の公開鍵で暗号化する場合。

```
$ echo Hello world | gpg -a -r alice -e
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

`--recipient` オプションは複数指定できる。

鍵束フォルダにある `gpg.conf` ファイルで常に使用する公開鍵を指定することもできる。

```
default-key alice
default-recipient-self
```

### セッション鍵のみで暗号化する

```
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

```
$ echo Hello world | gpg -r alice -e -a > alice-enc.asc

$ gpg -d alice-enc.asc
gpg: 2048-ビットELG鍵, ID 39FD83A4C6CE2F44, 日付2017-11-30に暗号化されました
      "Alice <alice@example.com>"
Hello world
```

セッション鍵のみで暗号化した場合も同じコマンドで復号できる。

```
$ echo Hello world | gpg -a -c > alice-sym-enc.asc

$ gpg -d alice-sym-enc.asc
gpg: AES暗号化済みデータ
gpg: 1 個のパスフレーズで暗号化
Hello world
```

Windows の場合は安全のため復号データの出力先を `--output` オプションで指定することをお勧めする。

```
$ gpg -o out.txt -d alice-sym-enc.asc
gpg: AES暗号化済みデータ
gpg: 1 個のパスフレーズで暗号化

$ type out.txt
Hello world
```

## データへの電子署名と検証

### クリア署名

```
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

署名検証は `--decrypt` コマンドが使える。

```
$ echo Hello world | gpg -u alice --clear-sign | gpg -d
Hello world
gpg: 11/30/17 18:05:11 東京 (標準時)に施された署名
gpg:                DSA鍵EFA7A201F30E999125B85A1C7A9D6B9CE230BAE9を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

### 分離署名

まず署名対象のファイルを用意する。

```
$ echo Hello world> hello.txt

$ cat hello.txt
Hello world
```

このファイルを配布した際に途中で改竄がないかどうか知りたい。
こういう場合は「分離署名」を使う。
分離署名には `--detach-sign` コマンドを使う。

```
$ gpg -u alice -b hello.txt
```

この結果 `hello.txt` ファイルと同じ場所に `hello.txt.sig` ファイルが作成される。
この`hello.txt` ファイルと `hello.txt.sig` ファイルをセットで配布するのである。
どちらかのファイルが改竄されていれば署名の検証が NG になるはずである。

```
$ gpg --verify hello.txt.sig
gpg: 署名されたデータが'hello.txt'にあると想定します
gpg: 11/30/17 18:31:17 東京 (標準時)に施された署名
gpg:                DSA鍵EFA7A201F30E999125B85A1C7A9D6B9CE230BAE9を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

署名対象のファイルがテキスト・ファイルの場合は `--textmode` オプションを付けて電子署名を行ったほうが安全。

```
$ gpg -u alice --textmode -b hello.txt
```

### 署名データに署名対象のデータを埋め込む

```text
$ gpg -u alice -s hello.txt
```

```text
$ gpg -d hello.txt.gpg
Hello world
gpg: 11/30/17 19:21:22 東京 (標準時)に施された署名
gpg:                RSA鍵0200ACBA190D0FB1D4EC7C02DC1D13A6CCA86E21を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

これでデータの検証と抽出ができた。
署名対象のデータを埋め込む方式の何が嬉しいかというと，暗号化と組み合わせる事ができるのである。

### 暗号化と電子署名を同時に行う

Bob の公開鍵で暗号化して Alice の鍵で電子署名する。

```
$ gpg -u alice -r bob -se hello.txt
```

```
$ gpg -d hello.txt.gpg
gpg: 2048-ビットRSA鍵, ID D74C71530446FD66, 日付2017-11-30に暗号化されました
      "Bob <bob@example.com>"
Hello world
gpg: 11/30/17 19:34:06 東京 (標準時)に施された署名
gpg:                RSA鍵0200ACBA190D0FB1D4EC7C02DC1D13A6CCA86E21を使用
gpg:                発行者"alice@example.com"
gpg: "Alice <alice@example.com>"からの正しい署名 [究極]
```

セッション鍵のみで暗号化する場合も同様にできる。

```
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

## 鍵の失効

鍵を作成する際に鍵束フォルダ以下の `openpgp-revocs.d` フォルダに失効証明書が作成される。

```
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

```
:-----BEGIN PGP PUBLIC KEY BLOCK-----
```

の先頭のコロン（`:`）を削除して使うこと。

```
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

失効証明書の作成。

```
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

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
