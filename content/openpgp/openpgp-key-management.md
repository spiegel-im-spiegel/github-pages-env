+++
title = "OpenPGP 鍵管理に関する考察"
date =  "2017-12-01T17:51:59+09:00"
update =  "2017-12-02T16:20:26+09:00"
description = "OpenPGP 鍵の管理について考えてみるテスト。"
image = "/images/attention/openpgp.png"
tags = [
  "cryptography",
  "openpgp",
  "management",
  "pki",
  "trust",
  "gnupg",
]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = true
+++

(move from [{{< ref "/remark/2017/11/openpgp-key-management.md" >}}]({{< ref "/remark/2017/11/openpgp-key-management.md" >}} "OpenPGP 鍵管理に関する考察"))

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
公開鍵への電子署名を使って peer-to-peer で信用関係を結ぶ[^ksp1]。

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

[^tm1]: もう少し詳しく言うと，公開鍵への電子署名の際に「信用度」を設定し，集まった「信用度」の累積から公開鍵の「有効性」を機械的に判定する。なので（信用度が分からない）全く未知の人の電子署名がいくらあっても「有効性」は上がりにくい。また公開鍵に電子署名を施すことは相手の鍵をある程度以上信用していることになるため，よく分からない鍵に対して安易に電子署名を行うことは避けるべきである。なお，現行バージョンの [GnuPG] では web of trust 以外にも X.509 や [Tofu (Trust on first use)](https://en.wikipedia.org/wiki/Trust_on_first_use "Trust on first use - Wikipedia") といった信用モデルもサポートしている。

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

[OpenPGP] の代表的な実装である [GnuPG] の最新バージョンでは，以下に示すように，鍵の作成処理が（昔と比べて）大幅に簡略化できる。

```text
$ gpg --batch --passphrase pass123 --quick-generate-key "Alice <alice@example.com>" default default 0
gpg: 鍵02C94DC57527A786を究極的に信用するよう記録しました
gpg: 失効証明書を 'C:/Users/alice/AppData/Roaming/gnupg/openpgp-revocs.d\9416E477EBA825CD1694573102C94DC57527A786.rev' に保管しました。
```

- `--batch` オプション[^gk1]（または `--pinentry-mode` オプションに `loopback` を指定）と `--passphrase` オプションを組み合わせて Pinentry によるパスフレーズ入力を回避できる
- `--quick-generate-key` コマンドの第1引数にユーザID，第2引数にアルゴリズム，第3引数に使用目的，第4引数に有効期限を指定する
    - アルゴリズムに `default` を指定するか指定しない場合は既定のアルゴリズム（RSA/2048ビット）で鍵を作成する
    - 使用目的には主鍵[^k1] の種類を指定する。通常は `default` のまま（署名と証明）でよい（指定しなければ `default`）
    - 有効期限には期間（1週間なら `7d` または `1w`，1年なら `12m` または `1y` など）を指定する。 `0` を指定すると無期限になる（指定しないと有効期限が当日になる）

[^gk1]: `--gen-key` コマンドに `--batch` オプションを組み合わせて設定ファイルから鍵を作成することも可能である。詳しい方法は “[Unattended GPG key generation](https://www.gnupg.org/documentation/manuals/gnupg/Unattended-GPG-key-generation.html)” が参考になるだろう。
[^k1]: [OpenPGP] の鍵は1つの主鍵（primary key）と0個以上の副鍵（subkey）で構成されている。主鍵は必ず電子署名用の鍵になっていて，この主鍵にユーザID（とその自己署名）や他の鍵からの電子署名が付与される。副鍵は暗号化または電子署名用の鍵である。たとえば，データの暗号化と復号は実際にはこの副鍵を使って行う。副鍵は個別に追加したり失効したりできる。ちなみに [GnuPG] では通常のやり方では暗号化機能のみの鍵は作れない。電子署名機能のみの鍵は作ることができる。

生成された鍵の状態は以下の通り。

```text
$ gpg --list-keys alice
pub   rsa2048 2017-11-23 [SC]
      9416E477EBA825CD1694573102C94DC57527A786
uid           [  究極  ] Alice <alice@example.com>
sub   rsa2048 2017-11-23 [E]
```

作成した鍵の公開鍵を配布するには

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

### ひとつの鍵に複数のユーザIDを付与する場合

ユーザIDというのは鍵の名前 “`Alice <alice@examle.com>`” の部分である。

たとえば，以下の鍵があったとき

```text
$ gpg --list-keys alice
pub   rsa2048 2017-11-23 [SC]
      B686F36CA9FDC10057EFC5D58D7E04B8CE947112
uid           [  究極  ] Alice <alice@example.com>
sub   rsa2048 2017-11-23 [E]
```

これに新しいユーザID “`Alice <alice@examle2.com>`” を付加するには `--quick-add-uid` コマンドを使って

```text
$ gpg --quick-add-uid alice "Alice <alice@examle2.com>"
$ gpg --update-trustdb
```

とする[^updtd1]。
これで新しいユーザIDが付加された。

[^updtd1]: `--update-trustdb` コマンドは信用データベース（trustdb）の更新コマンドである。現行バージョンの [GnuPG] では信用度は `trustdb.gpg` ファイルを使って鍵束とは独立に管理される。通常は鍵の状態が変わると自動的に信用データベースが更新されるのだが，自動更新しない場合は `--update-trustdb` コマンドで更新できる。なお，他ユーザの公開鍵に電子署名を施した場合は `--update-trustdb` コマンドを起動して署名した鍵の信用度を設定する。信用度の設定は `--edit-key` コマンドの編集モードでも設定・変更が可能である。

```text
$ gpg --list-keys alice
pub   rsa2048 2017-11-23 [SC]
      B686F36CA9FDC10057EFC5D58D7E04B8CE947112
uid           [  究極  ] Alice <alice@examle2.com>
uid           [  究極  ] Alice <alice@example.com>
sub   rsa2048 2017-11-23 [E]
```

ちなみに，この鍵を [pgpdump] にかけると以下のようになる。

```text
$ gpg --export -a alice | pgpdump -u
Old: Public Key Packet(tag 6)(269 bytes)
        Ver 4 - new
        Public key creation time - Thu Nov 23 06:22:56 UTC 2017
        Pub alg - RSA Encrypt or Sign(pub 1)
        RSA n(2048 bits) - ...
        RSA e(17 bits) - ...
Old: User ID Packet(tag 13)(25 bytes)
        User ID - Alice <alice@example.com>
Old: Signature Packet(tag 2)(334 bytes)
        Ver 4 - new
        Sig type - Positive certification of a User ID and Public Key packet(0x13).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA256(hash 8)
        Hashed Sub: issuer fingerprint(sub 33)(21 bytes)
         v4 -   Fingerprint - b6 86 f3 6c a9 fd c1 00 57 ef c5 d5 8d 7e 04 b8 ce 94 71 12
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Thu Nov 23 06:22:56 UTC 2017
        Hashed Sub: key flags(sub 27)(1 bytes)
                Flag - This key may be used to certify other keys
                Flag - This key may be used to sign data
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
                Key ID - 0x8D7E04B8CE947112
        Hash left 2 bytes - 05 21
        RSA m^d mod n(2047 bits) - ...
                -> PKCS-1
Old: User ID Packet(tag 13)(25 bytes)
        User ID - Alice <alice@examle2.com>
Old: Signature Packet(tag 2)(334 bytes)
        Ver 4 - new
        Sig type - Positive certification of a User ID and Public Key packet(0x13).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA256(hash 8)
        Hashed Sub: issuer fingerprint(sub 33)(21 bytes)
         v4 -   Fingerprint - b6 86 f3 6c a9 fd c1 00 57 ef c5 d5 8d 7e 04 b8 ce 94 71 12
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Thu Nov 23 06:33:28 UTC 2017
        Hashed Sub: key flags(sub 27)(1 bytes)
                Flag - This key may be used to certify other keys
                Flag - This key may be used to sign data
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
                Key ID - 0x8D7E04B8CE947112
        Hash left 2 bytes - 7d 5a
        RSA m^d mod n(2048 bits) - ...
                -> PKCS-1
Old: Public Subkey Packet(tag 14)(269 bytes)
        Ver 4 - new
        Public key creation time - Thu Nov 23 06:22:56 UTC 2017
        Pub alg - RSA Encrypt or Sign(pub 1)
        RSA n(2048 bits) - ...
        RSA e(17 bits) - ...
Old: Signature Packet(tag 2)(310 bytes)
        Ver 4 - new
        Sig type - Subkey Binding Signature(0x18).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA256(hash 8)
        Hashed Sub: issuer fingerprint(sub 33)(21 bytes)
         v4 -   Fingerprint - b6 86 f3 6c a9 fd c1 00 57 ef c5 d5 8d 7e 04 b8 ce 94 71 12
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Thu Nov 23 06:22:56 UTC 2017
        Hashed Sub: key flags(sub 27)(1 bytes)
                Flag - This key may be used to encrypt communications
                Flag - This key may be used to encrypt storage
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0x8D7E04B8CE947112
        Hash left 2 bytes - 3a a7
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

この方法ならユーザも各運用鍵もルート鍵とのみ信用関係を構築すればいいので柔軟な運用が可能になる。
欠点としてはルート鍵の管理が煩雑になりがちで信用に関する責務も重くなるため，かなり慎重な運用が求められることであろう。

#### Alice のルート鍵と運用鍵

では実際にやってみよう。

まず Alice がルート鍵と運用鍵の運用を行うとする。
ルート鍵は電子署名を行うだけの鍵なので，以下のように署名専用鍵として作成する（アルゴリズムは DSA/3072ビット にしてみた）。

```text
$ gpg --batch --passphrase pass123 --quick-generate-key "Alice (root) <alice@example.com>" dsa3072 default 0
gpg: *警告*: いくつかのOpenPGPプログラムはこのダイジェスト長のDSA鍵を扱うことができません
gpg: 鍵B965D53DB907EF0Eを究極的に信用するよう記録しました
gpg: 失効証明書を 'C:/Users/spiegel/AppData/Roaming/gnupg/openpgp-revocs.d\3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E.rev' に保管しました。
```

もうひとつ。
運用鍵も作っておく。

```text
$ gpg --batch --passphrase pass123 --quick-generate-key "Alice (commit) <alice@example.com>" default default 0
gpg: 鍵DFFC3F67BBB3C083を究極的に信用するよう記録しました
gpg: 失効証明書を 'C:/Users/spiegel/AppData/Roaming/gnupg/openpgp-revocs.d\A3CEFEEEDA222024F325C403DFFC3F67BBB3C083.rev' に保管しました。
```

この2つの鍵でお互いに電子署名を交わす（パスフレースの入力あり）。

```text
$ gpg -u 3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E --quick-sign-key A3CEFEEEDA222024F325C403DFFC3F67BBB3C083

sec  rsa2048/DFFC3F67BBB3C083
     作成: 2017-11-23  有効期限: 無期限      利用法: SC
     信用: 究極          有効性: 究極
  主鍵フィンガープリント: A3CE FEEE DA22 2024 F325  C403 DFFC 3F67 BBB3 C083

     Alice (commit) <alice@example.com>

$ gpg -u A3CEFEEEDA222024F325C403DFFC3F67BBB3C083 --quick-sign-key 3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E

sec  dsa3072/B965D53DB907EF0E
     作成: 2017-11-23  有効期限: 無期限      利用法: SC
     信用: 究極          有効性: 究極
  主鍵フィンガープリント: 3F8E FC47 7F9D 4D49 AA6C  308F B965 D53D B907 EF0E

     Alice (root) <alice@example.com>
```

これで Alice の2つの鍵の署名状態はこんな感じになった。

```text
$ gpg --list-sigs alice
pub   dsa3072 2017-11-23 [SC]
      3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E
uid           [  究極  ] Alice (root) <alice@example.com>
sig 3        B965D53DB907EF0E 2017-11-23  Alice (root) <alice@example.com>
sig          DFFC3F67BBB3C083 2017-11-23  Alice (commit) <alice@example.com>

pub   rsa2048 2017-11-23 [SC]
      A3CEFEEEDA222024F325C403DFFC3F67BBB3C083
uid           [  究極  ] Alice (commit) <alice@example.com>
sig 3        DFFC3F67BBB3C083 2017-11-23  Alice (commit) <alice@example.com>
sig          B965D53DB907EF0E 2017-11-23  Alice (root) <alice@example.com>
sub   rsa2048 2017-11-23 [E]
sig          DFFC3F67BBB3C083 2017-11-23  Alice (commit) <alice@example.com>
```

一方の主鍵に他方の鍵の電子署名が付与されているのが分かるだろうか。

#### Bob 鍵で Alice のルート鍵に電子署名する

今度は Bob の側である。
まずは Bob の公開鍵をこんな感じで作ってみた（作成操作は省略）。

```text
$ gpg --list-keys bob
pub   rsa2048 2017-11-23 [SC]
      B4E708652A1E81445B328A3D93F496726CBE8335
uid           [  究極  ] Bob <bob@example.com>
sub   rsa2048 2017-11-23 [E]
```

この環境に先程の Alice の公開鍵をインポートしてみる。
まずはルート鍵から。

```text
$ gpg --import alice-root.asc
gpg: key B965D53DB907EF0E: 鍵がないため1個の署名は検査しません
gpg: 鍵B965D53DB907EF0E: 公開鍵"Alice (root) <alice@example.com>"をインポートしました
gpg:           処理数の合計: 1
gpg:             インポート: 1
gpg: marginals needed: 3  completes needed: 1  trust model: pgp
gpg: 深さ: 0  有効性:   1  署名:   0  信用: 0-, 0q, 0n, 0m, 0f, 1u
```

当然ながら，この時点では読み込んだルート鍵の有効性は不明である。

```text
$ gpg --list-keys alice
pub   dsa3072 2017-11-23 [SC]
      3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E
uid           [  不明  ] Alice (root) <alice@example.com>
```

では，ルート鍵に Bob の鍵で署名しよう[^lsign1]（パスフレースの入力あり）。

[^lsign1]: 公開鍵に電子署名したことを公開したくない場合は `--lsign-key` コマンドで署名する。 `--lsign-key` コマンドで付与した署名はエクスポートされないため他ユーザには公開されない。公開鍵に関する確証はないけど取り敢えず使いたいという場合には有効な手だろう。

```text
$ gpg --quick-sign-key 3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E

pub  dsa3072/B965D53DB907EF0E
     作成: 2017-11-23  有効期限: 無期限      利用法: SC
     信用: 不明の        有効性: 不明の
  主鍵フィンガープリント: 3F8E FC47 7F9D 4D49 AA6C  308F B965 D53D B907 EF0E

     Alice (root) <alice@example.com>
```

これでルート鍵の有効性は「充分」になった。

```text
$ gpg --list-keys alice
pub   dsa3072 2017-11-23 [SC]
      3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E
uid           [  充分  ] Alice (root) <alice@example.com>
```

さらに信用データベースを更新して信用度も設定する。

```text
$ gpg --update-trustdb
gpg: marginals needed: 3  completes needed: 1  trust model: pgp
gpg: 深さ: 0  有効性:   1  署名:   1  信用: 0-, 0q, 0n, 0m, 0f, 1u
信用度が指定されていません:
pub   dsa3072 2017-11-23 [SC]
      "Alice (root) <alice@example.com>"
  主鍵フィンガープリント: 3F8E FC47 7F9D 4D49 AA6C  308F B965 D53D B907 EF0E

他のユーザの鍵を正しく検証するために、このユーザの信用度を決めてください
(パスポートを見せてもらったり、他から得たフィンガープリントを検査したり、などなど)

  1 = 知らない、または何とも言えない
  2 = 信用し ない
  3 = まぁまぁ信用する
  4 = 充分に信用する
  s = この鍵はとばす
  q = 終了

あなたの決定は? 4
gpg: 深さ: 1  有効性:   1  署名:   0  信用: 0-, 0q, 0n, 0m, 1f, 0u
      3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E
```

ここでは「充分に信用する」を選択した。

次に運用鍵もインポートする（操作は同じなので省略）。
この時点で Alice の公開鍵の状態を見ると

```text
$ gpg --list-keys alice
pub   dsa3072 2017-11-23 [SC]
      3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E
uid           [  充分  ] Alice (root) <alice@example.com>

pub   rsa2048 2017-11-23 [SC]
      A3CEFEEEDA222024F325C403DFFC3F67BBB3C083
uid           [  充分  ] Alice (commit) <alice@example.com>
sub   rsa2048 2017-11-23 [E]
```

運用鍵の有効性も既に「充分」になっていることが分かると思う。
ちなみにルート鍵の信用度を「まぁまぁ信用する」にすると

```text
$ gpg --list-keys alice
pub   dsa3072 2017-11-23 [SC]
      3F8EFC477F9D4D49AA6C308FB965D53DB907EF0E
uid           [  充分  ] Alice (root) <alice@example.com>

pub   rsa2048 2017-11-23 [SC]
      A3CEFEEEDA222024F325C403DFFC3F67BBB3C083
uid           [まぁまぁ] Alice (commit) <alice@example.com>
sub   rsa2048 2017-11-23 [E]
```

運用鍵の有効性が「まぁまぁ」になる。

独りでこうした運用をやるメリットは殆どないが，プロジェクト・チーム等で一括して鍵管理を行いたい場合は有効な手段だと思う。

## 有効期限について

この記事ではすべての鍵を「無期限」で設定している。
公開鍵の有効期限をどのようにするかは意見が別れるところだと思うが，私個人としては原則として「無期限」にすることをお薦めする。
何故なら [OpenPGP] 鍵は永続性と一貫性が重要だからである。

公開鍵に有効期限を設ける理由としては

- 期間の決まったプロジェクト内でのみ使用する鍵である
- チーム・メンバの出入りが激しく無期限では却って管理が煩雑になる
- 鍵のセキュリティ強度の問題から期限を切って運用したい（たとえば RSA/2048ビット鍵が acceptable なのは2030年までである）

といったところだろうか。
これならば，これまで述べたようにルート鍵と運用鍵を分けて，ルート鍵の方で永続性と一貫性を担保するように運用していくのがよいだろう。

自分の鍵であれば有効期限はいつでも変更できる（変更時にパスフレーズ入力あり）。

```text
$ gpg --list-keys A3CEFEEEDA222024F325C403DFFC3F67BBB3C083
pub   rsa2048 2017-11-23 [SC]
      A3CEFEEEDA222024F325C403DFFC3F67BBB3C083
uid           [  究極  ] Alice (commit) <alice@example.com>
sub   rsa2048 2017-11-23 [E]

$ gpg --quick-set-expire A3CEFEEEDA222024F325C403DFFC3F67BBB3C083 2y

$ gpg --list-keys A3CEFEEEDA222024F325C403DFFC3F67BBB3C083
pub   rsa2048 2017-11-23 [SC] [有効期限: 2019-11-23]
      A3CEFEEEDA222024F325C403DFFC3F67BBB3C083
uid           [  究極  ] Alice (commit) <alice@example.com>
sub   rsa2048 2017-11-23 [E] [有効期限: 2019-11-23]
```

しかし，有効期限を随時延長していく運用は鍵のオーナーもそれを使うユーザも手間が増えるだけであまりメリットがない[^own1]。
[OpenPGP] 鍵は（今のところ）利用するユーザに公開鍵の変更を自動的に通知・配信する仕組みがないので（それとも cron で鍵サーバへアクセスする？），ユーザ側の手間の多い運用では取りこぼしが出る可能性が大きくなる。

[^own1]: 公開鍵の更新を鍵オーナーの存在証明に使うのは，あまり筋のいい運用とは思えない。

## 鍵を失効させる

秘密鍵が漏洩するなどして鍵の失効が必要になった場合には，鍵作成時に作られた失効証明書を使って失効させる[^rvk1]。
鍵作成時に作られた失効証明書の中身はこんな感じになっている。

[^rvk1]: または `--gen-revoke` コマンドで失効証明書を作成する。 `--gen-revoke` コマンドで作成した場合は失効理由も含めることができる。

```text
$ cd C:/Users/alice/AppData/Roaming/gnupg/openpgp-revocs.d

$ cat 9416E477EBA825CD1694573102C94DC57527A786.rev
これは失効証明書でこちらのOpenPGP鍵に対するものです:

pub   rsa2048 2017-11-23 [SC]
      9416E477EBA825CD1694573102C94DC57527A786
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

iQE2BCABCAAgFiEElBbkd+uoJc0WlFcxAslNxXUnp4YFAloWYkcCHQAACgkQAslN
xXUnp4aG1wf/XoyxQPks2JlJ93ghQALdqCIxFPh015q21K53u0rVwTsMocwdGowR
l+/UppyBxnGyU1uIba68D787INlruMzsOyUTuruCUZ5XJpiuYYVXcRuovUmCREWF
EbW1DGd1lzmrO8cr70qu3yVCMYjGOQ6NA0fh1lpKTpFqHc3GC+ue19RDoVY1KnCC
YsWuNom4PAuUyHlH3uJLM9+F9J2Qec+0PIedxHyxuIStWOSg+/TGjD4cP3FEItIt
giRxx6qLWcK+bfg6WXv7I3+FA2J8eRKjLoD/vsZX+FpxG7T+c4mvfTUgxn0+bZD9
gxTKlFWg2bhKTcxi0EsJ9XAEmocvOolwPQ==
=ShPY
-----END PGP PUBLIC KEY BLOCK-----
```

実際に使う場合は

```text
:-----BEGIN PGP PUBLIC KEY BLOCK-----
```

の先頭のコロンを削除して使う。
失効処理を行うには `--import` コマンドで失効証明書をインポートすればよい。

```text
$ gpg --import 9416E477EBA825CD1694573102C94DC57527A786.rev
gpg: 鍵02C94DC57527A786:"Alice <alice@example.com>"失効証明書をインポートしました
gpg:           処理数の合計: 1
gpg:         新しい鍵の失効: 1
gpg: marginals needed: 3  completes needed: 1  trust model: pgp
gpg: 深さ: 0  有効性:   1  署名:   0  信用: 0-, 0q, 0n, 0m, 0f, 1u
```

これで鍵の状態は

```text
$ gpg --list-keys alice
pub   rsa2048 2017-11-23 [SC] [失効: 2017-11-23]
      9416E477EBA825CD1694573102C94DC57527A786
uid           [  失効  ] Alice <alice@example.com>
```

となり失効したことが分かる。
**失効した公開鍵を配布するのを忘れずに！**

## ブックマーク

- [Using the GNU Privacy Guard: OpenPGP Key Management](https://www.gnupg.org/documentation/manuals/gnupg/OpenPGP-Key-Management.html)

- [わかる！ OpenPGP 暗号](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)
- [GnuPG for Windows ― インストール編]({{< ref "/remark/2016/03/using-gnupg-modern-version-1.md" >}})
- [GnuPG for Windows ― gpg-agent について]({{< ref "/remark/2016/03/using-gnupg-modern-version-2.md" >}})

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
