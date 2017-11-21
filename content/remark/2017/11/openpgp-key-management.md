+++
title = "OpenPGP 鍵管理に関する考察"
date =  "2017-11-21T22:34:52+09:00"
description = "OpenPGP 鍵の管理について考えてみるテスト。"
image = "/images/attention/remark.jpg"
tags = [
  "cryptography",
  "openpgp",
  "management",
  "pki",
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
  mermaidjs = true
+++

たまたま以下の記事を見かけたのだが

- [gpg (GNU Privacy Guard)の使い方 - Qiita](https://qiita.com/moutend/items/5c22d6e57a74845578f6)

このやり方も良さそうだけど，もう少し簡単に運用できないか考えてみる。
なお [OpenPGP] の鍵管理は目的別にアドホック（ad hoc）な運用も可能なので「これ！」という正解はない。
自分にあったやり方を考えるのも面白いと思う。

## [OpenPGP] の信用モデル

鍵の管理について考える前に [OpenPGP] の信用モデルについておさらいしておこう。

最初の登場人物は Alice と Bob。
2人はそれぞれ [OpenPGP] 鍵を持っていて，これを使って秘密のやり取りをしようと考えている。
持っている鍵が信用できることを証明するために，お互い相手の公開鍵に自身の鍵で電子署名を施した。

{{< fig-mermaid >}}
graph LR
  Alice["Alice's Public Key"]
  Bob["Bob's Public Key"]

  Alice-- Digital Sign -->Bob
  Bob-- Digital Sign -->Alice
{{< /fig-mermaid >}}

鍵に施されている電子署名を確認することでコンテンツに対する暗号文や電子署名が正しい鍵で行われていることが証明できるわけだ。
これが [OpenPGP] の基本。
公開鍵への電子署名を使って peer-to-peer で信頼関係を結ぶ[^ksp1]。

[^ksp1]: 相手の公開鍵に電子署名するには (1) 相手の公開鍵を貰って `--import` （または鍵サーバから `--recv-keys`）する (2) インポートした鍵に `--sign-key` する (3) 署名した公開鍵を `--export` して相手に返す（または鍵サーバへ `--send-keys`） といった手順を踏む。これをひとりひとりやるのは割と面倒なので，複数人が一度に電子署名を交わすために「[キーサイン・パーティ（key signing party）](https://en.wikipedia.org/wiki/Key_signing_party "Key signing party - Wikipedia")」が行われることがある。日本ではあまり聞かないけど。

ここで3人目の人物 Chris に登場してもらおう。
Bob と Chris は面識があり既にお互いの公開鍵で電子署名を交わしている。
しかし Alice と Chris は面識がなく電子署名を交わす機会がないとする。
Alice から見て Chris の鍵は信用できるだろうか？

（念のために言うと，ここで言う「信用」は「あなたは人として信用できる」の信用ではなく「この鍵は正しく本人のものであると信用できる」の信用である）

{{< fig-mermaid >}}
graph LR
  Alice["Alice's Public Key"]
  Bob["Bob's Public Key"]
  Chris["Chris's Public Key"]

  Alice-- Digital Sign -->Bob
  Alice-. trust? .->Chris
  Bob-- Digital Sign -->Alice
  Bob-- Digital Sign -->Chris
  Chris-- Digital Sign -->Bob
{{< /fig-mermaid >}}

まず Alice から見て直接電子署名を交わした Bob の鍵が信用できるのは明らかである。

{{< fig-mermaid >}}
 graph LR
   Alice["Alice's Public Key"]
   Bob["Bob's Public Key"]
   Chris["Chris's Public Key"]

   Alice-- Digital Sign -->Bob
   Alice-. trust .->Bob
   Bob-- Digital Sign -->Alice
   Bob-- Digital Sign -->Chris
   Chris-- Digital Sign -->Bob

{{< /fig-mermaid >}}

Alice は Chris の公開鍵に信用できる Bob の公開鍵による電子署名を見つけたため， Bob の公開鍵と同じく Chris の公開鍵も有効であると見なすことができる。

{{< fig-mermaid >}}
graph LR
  Alice["Alice's Public Key"]
  Bob["Bob's Public Key"]
  Chris["Chris's Public Key"]

  Alice-- Digital Sign -->Bob
  Alice-. validate! .->Chris
  Alice-. trust .->Bob
  Bob-- Digital Sign -->Alice
  Bob-- Digital Sign -->Chris
  Chris-- Digital Sign -->Bob
{{< /fig-mermaid >}}

これが [OpenPGP] の代表的な信用モデル “web of trust” の基本的な考え方である[^tm1]。
このことから [OpenPGP] の鍵管理ににおいて「信用できる鍵」の条件は

[^tm1]: もう少し詳しく言うと，公開鍵への電子署名の際に「信用度」を設定し，集まった「信用度」の累積から公開鍵の「有効性」を機械的に判定する。なので（信用度が分からない）全く未知の人の電子署名がいくらあっても「有効性」は上がりにくい。また公開鍵に電子署名を施すことは相手の鍵をある程度以上信用していることになるため，よく分からない鍵に対して安易に電子署名を行うことは避けるべきである。ちなみに，現行バージョンの [GnuPG] では信用度は `trustdb.gpg` ファイルを使って鍵束とは独立に管理される。また web of trust 以外にも X.509 や [Tofu (Trust on first use)](https://en.wikipedia.org/wiki/Trust_on_first_use "Trust on first use - Wikipedia") といった信用モデルもサポートしている。

- 多くの電子署名（とできれば信用）が集まっていること

だと分かる。
更にこのことから派生的に

- 同じ鍵が永続的に使われ続けていること

も「信用できる鍵」の条件となる。
何故なら，鍵が頻繁に変わるとその度に電子署名をやり直すことになり，鍵に署名（＝信用）が集まりにくくなるからである。

## [OpenPGP] の鍵管理

以上を踏まえて [OpenPGP] 鍵の管理について考えてみよう。

### ひとつの鍵で運用する場合

一番簡単なケースで，ひとつの [OpenPGP] 鍵で全てをまかなう場合を考える。
たとえば，ふだん暗号化ツールなんて全然使わないけど git commit に電子署名するためだけに [GnuPG] を使いたい，など。

[OpenPGP] の代表的な実装である [GnuPG] の最新バージョンでは，以下のように，鍵の作成処理が（昔と比べて）大幅に簡略化されている。

```text
$ gpg --gen-key
gpg (GnuPG) 2.2.3; Copyright (C) 2017 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

gpg: keybox'C:/Users/username/AppData/Roaming/gnupg/pubring.kbx'が作成されました
注意: 全機能の鍵生成には "gpg --full-generate-key" を使います。

GnuPGはあなたの鍵を識別するためにユーザIDを構成する必要があります。

本名: Alice
電子メール・アドレス: alice@examle.com
次のユーザIDを選択しました:
    "Alice <alice@examle.com>"

名前(N)、電子メール(E)の変更、またはOK(O)か終了(Q)? o
たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。

たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。
gpg: C:/Users/username/AppData/Roaming/gnupg/trustdb.gpg: 信用データベースができました
gpg: 鍵9E5A7D3B1F249DF9を究極的に信用するよう記録しました
gpg: ディレクトリ'C:/Users/username/AppData/Roaming/gnupg/openpgp-revocs.d'が作成されました
gpg: 失効証明書を 'C:/Users/username/AppData/Roaming/gnupg/openpgp-revocs.d\FAAD28B393CAA55FCA6791149E5A7D3B1F249DF9.rev' に保管しました。
公開鍵と秘密鍵を作成し、署名しました。

pub   rsa2048 2017-11-20 [SC] [有効期限: 2019-11-20]
      FAAD28B393CAA55FCA6791149E5A7D3B1F249DF9
uid                      Alice <alice@examle.com>
sub   rsa2048 2017-11-20 [E] [有効期限: 2019-11-20]
```

1. `--gen-key` オプションで鍵作成モードを起動
1. 本名，電子メール・アドレス，コメントを指定して `o` で確定
1. パスフレーズを設定（Pinentry で2回入力）

こんな感じで「本名」と「電子メール・アドレス」とパスフレーズの入力だけで生成できるよぷになった[^gk1]。
ただし，このままでは有効期限が当日のままなので `--edit-key` オプションで修正する必要がある。

[^gk1]: 暗号アルゴリズムや鍵長などを細かく指定したい場合は `--full-gen-key` オプションを使う。また流行りの ECC を使いたいのであれば `--expert` オプションも付加する。更に `--batch` オプションを使う方法もある。 `--batch` オプションの使い方は “[Unattended GPG key generation](https://www.gnupg.org/documentation/manuals/gnupg/Unattended-GPG-key-generation.html)” が参考になる。

```text
$ gpg --edit-key alice
gpg (GnuPG) 2.2.3; Copyright (C) 2017 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

gpg: 信用データベースの検査
gpg: marginals needed: 3  completes needed: 1  trust model: pgp
gpg: 深さ: 0  有効性:   1  署名:   0  信用: 0-, 0q, 0n, 0m, 0f, 1u
gpg: 次回の信用データベース検査は、2019-11-20です
秘密鍵が利用できます。

sec  rsa2048/9E5A7D3B1F249DF9
     作成: 2017-11-20  有効期限: 2019-11-20  利用法: SC
     信用: 究極          有効性: 究極
ssb  rsa2048/B4E4E1C7747239A2
     作成: 2017-11-20  有効期限: 2019-11-20  利用法: E
[  究極  ] (1). Alice <alice@examle.com>

gpg> key 0

sec  rsa2048/9E5A7D3B1F249DF9
     作成: 2017-11-20  有効期限: 2019-11-20  利用法: SC
     信用: 究極          有効性: 究極
ssb  rsa2048/B4E4E1C7747239A2
     作成: 2017-11-20  有効期限: 2019-11-20  利用法: E
[  究極  ] (1). Alice <alice@examle.com>

gpg> expire
主鍵の有効期限を変更します。
鍵の有効期限を指定してください。
         0 = 鍵は無期限
      <n>  = 鍵は n 日間で期限切れ
      <n>w = 鍵は n 週間で期限切れ
      <n>m = 鍵は n か月間で期限切れ
      <n>y = 鍵は n 年間で期限切れ
鍵の有効期間は? (0)0
鍵の有効期間は? (0)は無期限です
これで正しいですか? (y/N) y

sec  rsa2048/9E5A7D3B1F249DF9
     作成: 2017-11-20  有効期限: 無期限      利用法: SC
     信用: 究極          有効性: 究極
ssb  rsa2048/B4E4E1C7747239A2
     作成: 2017-11-20  有効期限: 2019-11-20  利用法: E
[  究極  ] (1). Alice <alice@examle.com>

gpg> key 1

sec  rsa2048/9E5A7D3B1F249DF9
     作成: 2017-11-20  有効期限: 無期限      利用法: SC
     信用: 究極          有効性: 究極
ssb* rsa2048/B4E4E1C7747239A2
     作成: 2017-11-20  有効期限: 2019-11-20  利用法: E
[  究極  ] (1). Alice <alice@examle.com>

gpg> expire
副鍵の有効期限を変更します。
鍵の有効期限を指定してください。
         0 = 鍵は無期限
      <n>  = 鍵は n 日間で期限切れ
      <n>w = 鍵は n 週間で期限切れ
      <n>m = 鍵は n か月間で期限切れ
      <n>y = 鍵は n 年間で期限切れ
鍵の有効期間は? (0)0
鍵の有効期間は? (0)は無期限です
これで正しいですか? (y/N) y

sec  rsa2048/9E5A7D3B1F249DF9
     作成: 2017-11-20  有効期限: 無期限      利用法: SC
     信用: 究極          有効性: 究極
ssb* rsa2048/B4E4E1C7747239A2
     作成: 2017-11-20  有効期限: 無期限      利用法: E
[  究極  ] (1). Alice <alice@examle.com>

gpg> save
```

操作手順としては

1. `--edit-key` オプションで編集モードを起動
1. `key 0` コマンドで主鍵を選択
1. `expire` コマンドで有効期限の設定モードにする
1. 無期限（= `0`）を指定して `y` で確定
1. `key 1` コマンドで副鍵を選択
1. `expire` コマンドで有効期限の設定モードにする
1. 無期限（= `0`）を指定して `y` で確定
1. `save` コマンドで保存と終了

となっている[^k1]。

[^k1]: [OpenPGP] の鍵は1つの主鍵（primary key）と0個以上の副鍵（subkey）で構成されている。主鍵は必ず電子署名用の鍵になっていて，この主鍵にユーザID（とその自己署名）や他の鍵からの電子署名が付与される。副鍵は暗号化または電子署名用の鍵である。たとえば，データの暗号化と復号は実際にはこの副鍵を使って行う。副鍵は個別に追加したり失効したりできる。ちなみに [GnuPG] では通常のやり方では暗号化機能のみの鍵は作れない。電子署名機能のみの鍵は作ることができる。

これでRSA/2048ビット，無期限の鍵の生成が完了した。
公開鍵を配布するには

```text
$ gpg --export -a alice > alice-key.asc
```

として `alice-key.asc` を直接配布するか

```text
$ gpg --keyserver keys.gnupg.net --send-key alice
```

として鍵サーバ（ここでは [`keys.gnupg.net`](http://keys.gnupg.net/)）にアップロードすればいい。

注意する点としてはパスフレーズと失効証明書を紛失・漏洩しないよう管理することであろう。
できれば失効証明書は普段使う PC や携帯端末とは別に管理しておくのが望ましい。
ちなみに [OpenPGP] ではパスフレーズは何処にも保存されないので，パスフレーズを忘れてしまうと全くリカバリできなくなる（だからといってパスフレーズを設定しないというのは通常運用ではあり得ないが）。

ノートPCや携帯端末には常に紛失・盗難のリスクが付きまとうが，予防に注力しすぎて現実的でない対策を執るよりも，これはもう「起こり得ること」として「事後」がスムーズに行われるようバックアップ等の準備しておくほうが賢明である。

#### 鍵の失効

秘密鍵が漏洩するなどして鍵の失効が必要になった場合には，鍵生成時に作られた失効証明書を使って失効処理を行う。
失効証明書の中身はこんな感じになっている。

```text
$ cd C:/Users/username/AppData/Roaming/gnupg/openpgp-revocs.d

$ cat FAAD28B393CAA55FCA6791149E5A7D3B1F249DF9.rev
これは失効証明書でこちらのOpenPGP鍵に対するものです:

pub   rsa2048 2017-11-20 [SC] [有効期限: 2019-11-20]
      FAAD28B393CAA55FCA6791149E5A7D3B1F249DF9
uid          Alice <alice@examle.com>

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

iQE2BCABCAAgFiEE+q0os5PKpV/KZ5EUnlp9Ox8knfkFAloSt6YCHQAACgkQnlp9
Ox8knflrFwf/Xzb7CPOMO5uZph2QoEJDzo34plliJ+GB/kpmubzZ/jf5CSTZfKUF
OeioKwgYzduKV4UorZDe3BUV0GwnzzVlDR9xu2E/WV/J621NtPmqbyIU7Q6iYoKd
3UiZc/P4vR5JSjgTJD4fryTIIeTZNAGng37/kFjVhwAUyvzyEdi+eFCMz13LiZAk
XFMCdSHEifGWJnI6LJkt1rpnUAaSoLjQIuwVtpw1ZyI5oob8chFTGZkrNQ0XURlP
vNBcD8kMA/e0pJguDinpNnBcaBPWuKpJllu0nl54oD1x2qykNH8ApWIWUpsM9eXL
v4LDSpTyA2kmLH+L05DyWoiwfgE4rz7JeA==
=pj4U
-----END PGP PUBLIC KEY BLOCK-----
```

実際に使う場合は

```text
:-----BEGIN PGP PUBLIC KEY BLOCK-----
```

の先頭のコロンを削除して使う。
失効処理を行うには `--import` オプションで失効証明書を読み込めばよい。

```text
$ gpg --import FAAD28B393CAA55FCA6791149E5A7D3B1F249DF9.rev
gpg: 鍵9E5A7D3B1F249DF9:"Alice <alice@examle.com>"失効証明書をインポートしました
gpg:           処理数の合計: 1
gpg:         新しい鍵の失効: 1
gpg: marginals needed: 3  completes needed: 1  trust model: pgp
gpg: 深さ: 0  有効性:   1  署名:   0  信用: 0-, 0q, 0n, 0m, 0f, 1u
```

これで先程作成した鍵は

```text
$ gpg --list-keys alice
pub   rsa2048 2017-11-20 [SC] [失効: 2017-11-20]
      FAAD28B393CAA55FCA6791149E5A7D3B1F249DF9
uid           [  失効  ] Alice <alice@examle.com>
```

となり失効したことが分かる。
**失効した公開鍵を配布するのを忘れずに！**

### ひとつの鍵に複数のユーザIDを付与する場合

ユーザIDというのは鍵の名前 “`Alice <alice@examle.com>`” の部分である。

たとえば，以下の鍵があったとき

```text
$ gpg --list-keys alice
pub   rsa2048 2017-11-20 [SC] [有効期限: 2019-11-20]
      8DF6263068D826D5691E2D85E86157A2B2117F02
uid           [  究極  ] Alice <alice@examle.com>
sub   rsa2048 2017-11-20 [E] [有効期限: 2019-11-20]
```

これに新しいユーザID “`Alice <alice@examle2.com>`” を付加するには

```text
$ gpg --edit-key alice
gpg (GnuPG) 2.2.3; Copyright (C) 2017 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

秘密鍵が利用できます。

sec  rsa2048/E86157A2B2117F02
     作成: 2017-11-20  有効期限: 2019-11-20  利用法: SC
     信用: 究極          有効性: 究極
ssb  rsa2048/D1565C702022D41D
     作成: 2017-11-20  有効期限: 2019-11-20  利用法: E
[  究極  ] (1). Alice <alice@examle.com>

gpg> adduid
本名: Alice
電子メール・アドレス: alice@examle2.com
コメント: second
次のユーザIDを選択しました:
    "Alice (second) <alice@examle2.com>"

名前(N)、コメント(C)、電子メール(E)の変更、またはOK(O)か終了(Q)? o

sec  rsa2048/E86157A2B2117F02
     作成: 2017-11-20  有効期限: 2019-11-20  利用法: SC
     信用: 究極          有効性: 究極
ssb  rsa2048/D1565C702022D41D
     作成: 2017-11-20  有効期限: 2019-11-20  利用法: E
[  究極  ] (1)  Alice <alice@examle.com>
[  不明  ] (2). Alice (second) <alice@examle2.com>

gpg> save
```

1. `--edit-key` オプションで編集モードを起動
1. `adduid` コマンドでユーザID追加モードにする
1. 本名，電子メール・アドレス，コメントを指定して `o` で確定
1. `save` コマンドで保存と終了

とする。
これで新しいユーザIDが付加された。

```text
$ gpg --list-keys alice
pub   rsa2048 2017-11-20 [SC] [有効期限: 2019-11-20]
      8DF6263068D826D5691E2D85E86157A2B2117F02
uid           [  究極  ] Alice (second) <alice@examle2.com>
uid           [  究極  ] Alice <alice@examle.com>
sub   rsa2048 2017-11-20 [E] [有効期限: 2019-11-20]
```

ちなみに，この鍵を [pgpdump] にかけると以下のようになる。

```text
$ gpg --export -a alice | pgpdump -u
Old: Public Key Packet(tag 6)(269 bytes)
        Ver 4 - new
        Public key creation time - Mon Nov 20 12:23:12 UTC 2017
        Pub alg - RSA Encrypt or Sign(pub 1)
        RSA n(2048 bits) - ...
        RSA e(17 bits) - ...
Old: User ID Packet(tag 13)(24 bytes)
        User ID - Alice <alice@examle.com>
Old: Signature Packet(tag 2)(340 bytes)
        Ver 4 - new
        Sig type - Positive certification of a User ID and Public Key packet(0x13).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA256(hash 8)
        Hashed Sub: issuer fingerprint(sub 33)(21 bytes)
         v4 -   Fingerprint - 8d f6 26 30 68 d8 26 d5 69 1e 2d 85 e8 61 57 a2 b2 11 7f 02
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Mon Nov 20 12:23:12 UTC 2017
        Hashed Sub: key flags(sub 27)(1 bytes)
                Flag - This key may be used to certify other keys
                Flag - This key may be used to sign data
        Hashed Sub: key expiration time(sub 9)(4 bytes)
                Time - Wed Nov 20 12:23:12 UTC 2019
        Hashed Sub: preferred symmetric algorithms(sub 11)(4 bytes)
                Sym alg - AES with 256-bit key(sym 9)
                Sym alg - AES with 192-bit key(sym 8)
                Sym alg - AES with 128-bit key(sym 7)
                Sym alg - Triple-DES(sym 2)
        Hashed Sub: preferred hash algorithms(sub 21)(5 bytes)
                Hash alg - SHA256(hash 8)
                Hash alg - SHA384(hash 9)
                Hash alg - SHA512(hash 10)
                Hash alg - SHA224(hash 11)
                Hash alg - SHA1(hash 2)
        Hashed Sub: preferred compression algorithms(sub 22)(3 bytes)
                Comp alg - ZLIB <RFC1950>(comp 2)
                Comp alg - BZip2(comp 3)
                Comp alg - ZIP <RFC1951>(comp 1)
        Hashed Sub: features(sub 30)(1 bytes)
                Flag - Modification detection (packets 18 and 19)
        Hashed Sub: key server preferences(sub 23)(1 bytes)
                Flag - No-modify
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0xE86157A2B2117F02
        Hash left 2 bytes - 44 0d
        RSA m^d mod n(2048 bits) - ...
                -> PKCS-1
Old: User ID Packet(tag 13)(34 bytes)
        User ID - Alice (second) <alice@examle2.com>
Old: Signature Packet(tag 2)(340 bytes)
        Ver 4 - new
        Sig type - Positive certification of a User ID and Public Key packet(0x13).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA256(hash 8)
        Hashed Sub: issuer fingerprint(sub 33)(21 bytes)
         v4 -   Fingerprint - 8d f6 26 30 68 d8 26 d5 69 1e 2d 85 e8 61 57 a2 b2 11 7f 02
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Mon Nov 20 12:26:49 UTC 2017
        Hashed Sub: key flags(sub 27)(1 bytes)
                Flag - This key may be used to certify other keys
                Flag - This key may be used to sign data
        Hashed Sub: key expiration time(sub 9)(4 bytes)
                Time - Wed Nov 20 12:23:12 UTC 2019
        Hashed Sub: preferred symmetric algorithms(sub 11)(4 bytes)
                Sym alg - AES with 256-bit key(sym 9)
                Sym alg - AES with 192-bit key(sym 8)
                Sym alg - AES with 128-bit key(sym 7)
                Sym alg - Triple-DES(sym 2)
        Hashed Sub: preferred hash algorithms(sub 21)(5 bytes)
                Hash alg - SHA256(hash 8)
                Hash alg - SHA384(hash 9)
                Hash alg - SHA512(hash 10)
                Hash alg - SHA224(hash 11)
                Hash alg - SHA1(hash 2)
        Hashed Sub: preferred compression algorithms(sub 22)(3 bytes)
                Comp alg - ZLIB <RFC1950>(comp 2)
                Comp alg - BZip2(comp 3)
                Comp alg - ZIP <RFC1951>(comp 1)
        Hashed Sub: features(sub 30)(1 bytes)
                Flag - Modification detection (packets 18 and 19)
        Hashed Sub: key server preferences(sub 23)(1 bytes)
                Flag - No-modify
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0xE86157A2B2117F02
        Hash left 2 bytes - 28 ee
        RSA m^d mod n(2046 bits) - ...
                -> PKCS-1
Old: Public Subkey Packet(tag 14)(269 bytes)
        Ver 4 - new
        Public key creation time - Mon Nov 20 12:23:12 UTC 2017
        Pub alg - RSA Encrypt or Sign(pub 1)
        RSA n(2048 bits) - ...
        RSA e(17 bits) - ...
Old: Signature Packet(tag 2)(316 bytes)
        Ver 4 - new
        Sig type - Subkey Binding Signature(0x18).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA256(hash 8)
        Hashed Sub: issuer fingerprint(sub 33)(21 bytes)
         v4 -   Fingerprint - 8d f6 26 30 68 d8 26 d5 69 1e 2d 85 e8 61 57 a2 b2 11 7f 02
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Mon Nov 20 12:23:12 UTC 2017
        Hashed Sub: key flags(sub 27)(1 bytes)
                Flag - This key may be used to encrypt communications
                Flag - This key may be used to encrypt storage
        Hashed Sub: key expiration time(sub 9)(4 bytes)
                Time - Wed Nov 20 12:23:12 UTC 2019
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0xE86157A2B2117F02
        Hash left 2 bytes - 05 a3
        RSA m^d mod n(2047 bits) - ...
                -> PKCS-1
```

ユーザID毎に電子署名（自己署名）が付与されているのがお分かりだろうか。

ひとつの鍵に複数のユーザIDを付与することに関しては昔から賛否あるのだが，手段が提供されていることは覚えておいて損はないだろう。

### 用途によって鍵を使い分けたい場合

たとえば，暗号化メール，リリースファイルの電子署名， git commit 時の電子署名といった用途毎に異なる鍵を使いたいことがある。
その場合でもそれぞれ鍵を生成して運用すればいいのだが，新しい鍵を作る度にそれぞれの鍵とやり取りを行うユーザが毎度電子署名を行うのは相当に煩雑な作業である。

そこで「ルート鍵」と「運用鍵」の2種類の鍵を作って運用する。
具体的にはルート鍵と各運用鍵との間で電子署名を交わす。

{{< fig-mermaid >}}
graph LR
  rt[Root Key]-- Digital Sign -->op1[Operation Key 1]
  op1-- Digital Sign -->rt

  rt-- Digital Sign -->op2[Operation Key 2]
  op2-- Digital Sign -->rt
{{< /fig-mermaid >}}

運用鍵とやり取りするユーザは，各運用鍵ではなく，ルート鍵と署名を交わし信用度を設定することによって各運用鍵の有効性を確認できる。

{{< fig-mermaid >}}
graph LR
  usr[User's Key]
  rt[Root Key]
  op1[Operation Key 1]
  op2[Operation Key 2]

  usr-- Digital Sign -->rt
  usr-. trust .->rt
  rt-- Digital Sign -->usr

  usr-. validate .->op1
  rt-- Digital Sign -->op1

  rt-- Digital Sign -->op2
  usr-. validate .->op2

{{< /fig-mermaid >}}

この方法ならユーザも各運用鍵もルート鍵とのみ信頼関係を構築すればいいので柔軟な運用が可能になる。
欠点としてはルート鍵の管理の手間が多く責務も重くなるため，かなり慎重な運用が求められることであろう。

では実際にやってみよう。

ルート鍵は電子署名を行うだけの鍵なので，以下のように署名専用鍵として作成する。

```text
$ gpg --full-gen-key
gpg (GnuPG) 2.2.3; Copyright (C) 2017 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

gpg: keybox'C:/Users/username/AppData/Roaming/gnupg/pubring.kbx'が作成されました
ご希望の鍵の種類を選択してください:
   (1) RSA と RSA (デフォルト)
   (2) DSA と Elgamal
   (3) DSA (署名のみ)
   (4) RSA (署名のみ)
あなたの選択は? 3
DSA 鍵は 1024 から 3072 ビットの長さで可能です。
鍵長は? (2048) 3072
要求された鍵長は3072ビット
鍵の有効期限を指定してください。
         0 = 鍵は無期限
      <n>  = 鍵は n 日間で期限切れ
      <n>w = 鍵は n 週間で期限切れ
      <n>m = 鍵は n か月間で期限切れ
      <n>y = 鍵は n 年間で期限切れ
鍵の有効期間は? (0)0
鍵の有効期間は? (0)は無期限です
これで正しいですか? (y/N) y

GnuPGはあなたの鍵を識別するためにユーザIDを構成する必要があります。

本名: Alice
電子メール・アドレス: alice@examle.com
コメント: root
次のユーザIDを選択しました:
    "Alice (root) <alice@examle.com>"

名前(N)、コメント(C)、電子メール(E)の変更、またはOK(O)か終了(Q)? o
たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。
gpg: *警告*: いくつかのOpenPGPプログラムはこのダイジェスト長のDSA鍵を扱うことができません

gpg: C:/Users/username/AppData/Roaming/gnupg/trustdb.gpg: 信用データベースができました
gpg: 鍵FFA6AEB318E21E7Aを究極的に信用するよう記録しました
gpg: ディレクトリ'C:/Users/username/AppData/Roaming/gnupg/openpgp-revocs.d'が作成されました
gpg: 失効証明書を 'C:/Users/username/AppData/Roaming/gnupg/openpgp-revocs.d\2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A.rev' に保管
しました。
公開鍵と秘密鍵を作成し、署名しました。

この鍵は暗号化には使用できないことに注意してください。暗号化を行うには、
"--edit-key"コマンドを使って副鍵を生成してください。
pub   dsa3072 2017-11-21 [SC]
      2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A
uid                      Alice (root) <alice@examle.com>
```

1. `--full-gen-key` オプションで鍵作成モードを起動
1. 公開鍵のアルゴリズムとして「(3) DSA (署名のみ)」を選択
1. 鍵長を3072ビットで指定
1. 有効期限を無期限（= `0`）に設定して `y` で確定
1. 本名，電子メール・アドレス，コメントを指定して `o` で確定
1. パスフレーズを設定（Pinentry で2回入力）

`--full-gen-key` オプションを使ってるので操作手順が少し多いのが難である。
次に運用鍵を用意する。
最初の `--gen-key` オプションを使ったやり方でこんな公開鍵を作ってみた。

```text
$ gpg --list-keys 990BB11E9219BA613733A5829DBB5E987E95A3EF
pub   rsa2048 2017-11-21 [SC] [有効期限: 2019-11-21]
      990BB11E9219BA613733A5829DBB5E987E95A3EF
uid           [  究極  ] Alice (commit) <alice@examle.com>
sub   rsa2048 2017-11-21 [E] [有効期限: 2019-11-21]
```

この鍵にルート鍵で電子署名を施す。
こんな感じ。

```text
$ gpg -u 2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A --sign-key 990BB11E9219BA613733A5829DBB5E987E95A3EF

sec  rsa2048/9DBB5E987E95A3EF
     作成: 2017-11-21  有効期限: 2019-11-21  利用法: SC
     信用: 究極          有効性: 究極
ssb  rsa2048/1E497E54FD1C6C87
     作成: 2017-11-21  有効期限: 2019-11-21  利用法: E
[  究極  ] (1). Alice (commit) <alice@examle.com>


sec  rsa2048/9DBB5E987E95A3EF
     作成: 2017-11-21  有効期限: 2019-11-21  利用法: SC
     信用: 究極          有効性: 究極
  主鍵フィンガープリント: 990B B11E 9219 BA61 3733  A582 9DBB 5E98 7E95 A3EF

     Alice (commit) <alice@examle.com>

この鍵は2019-11-21で期限が切れます。
本当にこの鍵にあなたの鍵"Alice (root) <alice@examle.com>"で署名してよいですか
(FFA6AEB318E21E7A)

本当に署名しますか? (y/N) y
```

電子署名の状態は感じ。

```text
$ gpg --list-sigs 990BB11E9219BA613733A5829DBB5E987E95A3EF
pub   rsa2048 2017-11-21 [SC] [有効期限: 2019-11-21]
      990BB11E9219BA613733A5829DBB5E987E95A3EF
uid           [  究極  ] Alice (commit) <alice@examle.com>
sig 3        9DBB5E987E95A3EF 2017-11-21  Alice (commit) <alice@examle.com>
sig          FFA6AEB318E21E7A 2017-11-21  Alice (root) <alice@examle.com>
sub   rsa2048 2017-11-21 [E] [有効期限: 2019-11-21]
sig          9DBB5E987E95A3EF 2017-11-21  Alice (commit) <alice@examle.com>
```

主鍵にルート鍵の署名が付与されているのが分かるだろうか。

同様にルート鍵に対しても運用鍵で署名を施しておく。
結果はこんな感じ。

```text
$ gpg --list-sigs 2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A
pub   dsa3072 2017-11-21 [SC]
      2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A
uid           [  究極  ] Alice (root) <alice@examle.com>
sig 3        FFA6AEB318E21E7A 2017-11-21  Alice (root) <alice@examle.com>
sig          9DBB5E987E95A3EF 2017-11-21  Alice (commit) <alice@examle.com>
```

今度は別の環境を用意する。

まずは Bob の公開鍵をこんな感じで作ってみた。

```text
$ gpg --list-keys bobby
pub   rsa2048 2017-11-21 [SC] [有効期限: 2019-11-21]
      A7E072D768DABA25E7F5AAB00A27C4ADAD0ECEF6
uid           [  究極  ] Bobby <bob@example.com>
sub   rsa2048 2017-11-21 [E] [有効期限: 2019-11-21]
```

この環境に先程の Alice の公開鍵（ルート鍵と運用鍵）をインポートしてみる。
インポート直後の公開鍵はこんな感じになっている。

```text
$ gpg --list-keys alice
pub   dsa3072 2017-11-21 [SC]
      2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A
uid           [  不明  ] Alice (root) <alice@examle.com>

pub   rsa2048 2017-11-21 [SC] [有効期限: 2019-11-21]
      990BB11E9219BA613733A5829DBB5E987E95A3EF
uid           [  不明  ] Alice (commit) <alice@examle.com>
sub   rsa2048 2017-11-21 [E] [有効期限: 2019-11-21]
```

まだ署名されていない未知の鍵なので有効性はどちらも「不明」である。

では，この Alice の公開鍵のうち，ルート鍵に先ほど作った Bob の鍵で署名しよう。
結果はこんな感じ。

```text
$ gpg --list-sigs 2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A
pub   dsa3072 2017-11-21 [SC]
      2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A
uid           [  充分  ] Alice (root) <alice@examle.com>
sig 3        FFA6AEB318E21E7A 2017-11-21  Alice (root) <alice@examle.com>
sig          9DBB5E987E95A3EF 2017-11-21  Alice (commit) <alice@examle.com>
sig          0A27C4ADAD0ECEF6 2017-11-21  Bobby <bob@example.com>
```

この時点で Bob （の環境）から見た Alice の公開鍵の有効性は以下の通り。

```text
$ gpg --list-keys alice
pub   dsa3072 2017-11-21 [SC]
      2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A
uid           [  充分  ] Alice (root) <alice@examle.com>

pub   rsa2048 2017-11-21 [SC] [有効期限: 2019-11-21]
      990BB11E9219BA613733A5829DBB5E987E95A3EF
uid           [ 未定義 ] Alice (commit) <alice@examle.com>
sub   rsa2048 2017-11-21 [E] [有効期限: 2019-11-21]
```

署名したルート鍵は「充分」だが運用鍵の方は「未定義」なのが分かるだろうか。
ここで更に Alice のルート鍵の信用度も設定する。
信用度の設定は `--edit-key` オプションを使う。

```text
$ gpg --edit-key 2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A
gpg (GnuPG) 2.2.3; Copyright (C) 2017 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.


pub  dsa3072/FFA6AEB318E21E7A
     作成: 2017-11-21  有効期限: 無期限      利用法: SC
     信用: 不明の        有効性: 充分
[  充分  ] (1). Alice (root) <alice@examle.com>

gpg> trust
pub  dsa3072/FFA6AEB318E21E7A
     作成: 2017-11-21  有効期限: 無期限      利用法: SC
     信用: 不明の        有効性: 充分
[  充分  ] (1). Alice (root) <alice@examle.com>

他のユーザの鍵を正しく検証するために、このユーザの信用度を決めてください
(パスポートを見せてもらったり、他から得たフィンガープリントを検査したり、などなど)

  1 = 知らない、または何とも言えない
  2 = 信用しない
  3 = まぁまぁ信用する
  4 = 充分に信用する
  5 = 究極的に信用する
  m = メーン・メニューに戻る

あなたの決定は? 4

pub  dsa3072/FFA6AEB318E21E7A
     作成: 2017-11-21  有効期限: 無期限      利用法: SC
     信用: 充分          有効性: 充分
[  充分  ] (1). Alice (root) <alice@examle.com>
プログラムを再起動するまで、表示された鍵の有効性は正しくないかもしれない、
ということを念頭においてください。

gpg> save
鍵は無変更なので更新は不要です。
```

1. `--edit-key` オプションで編集モードを起動
1. `trust` コマンドで信用度設定モードにする
1. 「4 = 充分に信用する」をセットする
1. 念のため `save` コマンドで保存と終了

この状態でもう一度 Alice の公開鍵を見てみよう。

```text
$ gpg --list-keys alice
pub   dsa3072 2017-11-21 [SC]
      2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A
uid           [  充分  ] Alice (root) <alice@examle.com>

pub   rsa2048 2017-11-21 [SC] [有効期限: 2019-11-21]
      990BB11E9219BA613733A5829DBB5E987E95A3EF
uid           [  充分  ] Alice (commit) <alice@examle.com>
sub   rsa2048 2017-11-21 [E] [有効期限: 2019-11-21]
```

運用鍵の有効性が変わったことが分かると思う。
ちなみにルート鍵の信頼度を「まぁまぁ信用する」にすると，こんな感じになる。

```text
$ gpg --list-keys alice
pub   dsa3072 2017-11-21 [SC]
      2D5A462785C5D42ADE2AEBA9FFA6AEB318E21E7A
uid           [  充分  ] Alice (root) <alice@examle.com>

pub   rsa2048 2017-11-21 [SC] [有効期限: 2019-11-21]
      990BB11E9219BA613733A5829DBB5E987E95A3EF
uid           [まぁまぁ] Alice (commit) <alice@examle.com>
sub   rsa2048 2017-11-21 [E] [有効期限: 2019-11-21]
```

以上がルート鍵を使った鍵管理である。
独りでこれをやるメリットは殆どないが，プロジェクト・チーム等で一括して鍵管理を行いたい場合は有効な手段だと思う。

## 有効期限について

この記事ではすべての鍵を「無期限」で設定している。
公開鍵の有効期限をどのようにするかは意見が別れるところだと思うが，私個人の意見としては原則として「無期限」にすることをお薦めする。
何故なら [OpenPGP] 鍵は永続性と一貫性が重要だからである。

公開鍵に有効期限を設ける理由としては

- 期間の決まったプロジェクト内でのみ使用する鍵である
- チーム・メンバの出入りが激しく無期限では却って管理が煩雑になる
- 鍵のセキュリティ強度の問題から期限を切って運用したい（たとえば RSA/2048ビット鍵が acceptable なのは2030年までである）

といったところだろうか。
これならば，これまで述べたようにルート鍵と運用鍵を分けて，ルート鍵の方で永続性と一貫性を担保するように運用していくのがよいだろう。

自分の鍵であれば有効期限はいつでも変更できるのだが，有効期限を随時延長していく運用は鍵のオーナーもそれを使うユーザも手間が増えるだけであまりメリットがない[^own1]。
[OpenPGP] 鍵は（今のところ）利用するユーザに公開鍵の変更を自動的に通知・配信する仕組みがないので（それとも cron で鍵サーバへアクセスする？），ユーザ側の手間の多い運用では取りこぼしが出る可能性が大きくなる。

[^own1]: 公開鍵の更新を鍵オーナーの存在証明に使うのは，あまり筋のいい運用とは思えない。

## ブックマーク

- [わかる！ OpenPGP 暗号](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)
- [GnuPG for Windows ― インストール編]({{< relref "remark/2016/03/using-gnupg-modern-version-1.md" >}})
- [GnuPG for Windows ― gpg-agent について]({{< relref "remark/2016/03/using-gnupg-modern-version-2.md" >}})

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
