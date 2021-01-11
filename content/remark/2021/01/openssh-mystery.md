+++
title = "OpenSSH 完全に理解した（笑）"
date =  "2021-01-11T16:31:17+09:00"
description = "どうやら OpenSSH に対する理解が足りないと気付いたので，知ってること気付いたこと等を書き出してみる。"
image = "/images/attention/kitten.jpg"
tags = [ "openssh", "cryptography", "hash" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

使ってみたかったんだ，このフレーズ（笑）

めでたくも[フィードバック先移行]({{< ref "/remark/2020/12/github-discussions.md" >}} "ようやく Disqus を捨てれるよ")後，はじめてのフィードバックを頂いた。
どうやら [OpenSSH] に対する理解が足りないと気付いたので，知ってること気付いたこと等を書き出してみる。

「ここ違うよ」とかいうのがありましたら[フィードバック](https://github.com/spiegel-im-spiegel/github-pages-env/discussions "Discussions · spiegel-im-spiegel/github-pages-env")にてご指摘いただけると有り難いです。

[OpenSSH] に関する本とかあればよかったんだけどねぇ。
日本語の本はみんな出版年がエラく古いし，しかも「使い方」は書いてあっても中で何してるかについてはあまり言及されてない感じ？ とっかに「Inside OpenSSH」みたいな本があればいいのに（笑）

## [OpenSSH] 暗号化通信の手順

[OpenSSH] におけるサーバ-クライアント間の暗号化通信の手順は大雑把に

1. サーバ-クライアント間で使える暗号スイートの確認
2. 鍵交換アルゴリズムによるシークレットの導出とセッション鍵の生成
3. セッション鍵による暗号化通信の開始
4. クライアント認証

となっているようだ。
以降でひとつづつ見ていこう。

## 鍵交換アルゴリズムとハッシュ・アルゴリズム

暗号化通信を行うためにはサーバ-クライアント間でセッション鍵（共通鍵）を共有する必要がある。
（昔はともかく）今の [OpenSSH] ではセッション鍵の取得に Diffie-Hellman 鍵交換アルゴリズムまたはそのバリエーションを使う。
Diffie-Hellman 鍵交換アルゴリズムは公開鍵暗号の一種だがセッションの開始ごとに使い捨ての鍵ペアを生成する（筈）ので，ユーザがそれを意識することはないだろう。

鍵交換アルゴリズムで得られる共有データを「シークレット」と呼ぶことがある。
[OpenSSH] ではシークレットをハッシュ化した値をセッション鍵として使っているようだ。

したがってセッション鍵の共有には鍵交換アルゴリズムとハッシュ・アルゴリズムの2つが必要となる。
[OpenSSH] でサポートしてる暗号スイートは以下の通り。

| 名称                                   | 鍵交換                  | ハッシュ |         推奨          | 備考               |
| -------------------------------------- | ----------------------- | -------- |:---------------------:| ------------------ |
| `curve25519-sha256`                    | ECDH (curve25519)       | SHA2-256 | {{< icons "check" >}} | [OpenSSH] 7.4 以降 |
| `diffie-hellman-group1-sha1`           | DH (1024 bits)          | SHA1     |                       |                    |
| `diffie-hellman-group14-sha1`          | DH (2048 bits)          | SHA1     |                       |                    |
| `diffie-hellman-group14-sha256`        | DH (2048 bits)          | SHA2-256 | {{< icons "check" >}} | [OpenSSH] 7.3 以降 |
| `diffie-hellman-group16-sha512`        | DH (4096 bits)          | SHA2-512 | {{< icons "check" >}} | [OpenSSH] 7.3 以降 |
| `diffie-hellman-group18-sha512`        | DH (8192 bits)          | SHA2-512 | {{< icons "check" >}} | [OpenSSH] 7.3 以降 |
| `diffie-hellman-group-exchange-sha1`   | Custom DH (?)           | SHA1     |                       |                    |
| `diffie-hellman-group-exchange-sha256` | Custom DH (?)           | SHA2-256 | {{< icons "check" >}} | [OpenSSH] 5.4 以降 |
| `ecdh-sha2-nistp256`                   | ECDH (NIST curve P-256) | SHA2-256 |                       | [OpenSSH] 5.7 以降 |
| `ecdh-sha2-nistp384`                   | ECDH (NIST curve P-384) | SHA2-256 |                       | [OpenSSH] 5.7 以降 |
| `ecdh-sha2-nistp521`                   | ECDH (NIST curve P-521) | SHA2-256 |                       | [OpenSSH] 5.7 以降 |

SHA-1 は危殆化が叫ばれて久しく，既に推奨されていない。
このブログでもあちこちで書いているが Zenn の以下の記事でまとめているので興味のある人はどうぞ。

- [さようなら SHA-1](https://zenn.dev/spiegel/articles/20201025-sayonara-sha1)

ECDH (Elliptic Curve Diffie–Hellman) は Diffie-Hellman 鍵交換アルゴリズムの楕円曲線版と考えてもらって構わない。

ECDH で使用する楕円曲線のうち NIST 推奨パラメータ（curve P-256/384/521）は微妙に評判が悪いようで，推奨しないところがあった。
どうも NIST 発のアルゴリズムは使いたくないということのようだ。
まぁ NSA 絡みで度々やらかしてるから信用がないんだろうな（笑）

これらの暗号スイートの選択と優先順位はサーバ側で設定するのでユーザが意識することはない。
強いて言うなら古いバージョンは使うなってことくらいか。

## 共通鍵暗号アルゴリズムと暗号モード

実際の暗号化通信はセッション鍵を使った共通鍵暗号アルゴリズムで行う。
[OpenSSH] でサポートしてる共通鍵暗号アルゴリズムと暗号モードは以下の通り。

| 名称                            | 共通鍵暗号    | 暗号モード |         推奨          | 備考               |
| ------------------------------- | ------------- | ---------- |:---------------------:| ------------------ |
| `3des-cbc`                      | TripleDES     | CBC        |                       |                    |
| `aes128-cbc`                    | AES-128       | CBC        |                       |                    |
| `aes192-cbc`                    | AES-192       | CBC        |                       |                    |
| `aes256-cbc`                    | AES-256       | CBC        |                       |                    |
| `aes128-ctr`                    | AES-128       | CTR        | {{< icons "check" >}} |                    |
| `aes192-ctr`                    | AES-192       | CTR        | {{< icons "check" >}} |                    |
| `aes256-ctr`                    | AES-256       | CTR        | {{< icons "check" >}} |                    |
| `aes128-gcm`@openssh.com        | AES-128       | GCM        | {{< icons "check" >}} | [OpenSSH] 6.2 以降 |
| `aes256-gcm`@openssh.com        | AES-256       | GCM        | {{< icons "check" >}} | [OpenSSH] 6.2 以降 |
| `arcfour`                       | ARCFOUR (40?) | &mdash;    |                       |                    |
| `arcfour128`                    | ARCFOUR (128) | &mdash;    |                       |                    |
| `arcfour256`                    | ARCFOUR (256) | &mdash;    |                       |                    |
| `blowfish-cbc`                  | Blowfish (64) | CBC        |                       |                    |
| `cast128-cbc`                   | CAST (128)    | CBC        |                       |                    |
| `chacha20-poly1305`@openssh.com | ChaCha20      | Poly1305   | {{< icons "check" >}} | [OpenSSH] 6.5 以降 |

ARCFOUR ってのはいわゆる RC4 ストリーム暗号のこと。
RC4 は公式にはアルゴリズムを公開していないため “Alleged RC FOUR” の意味で ARCFOUR という名称が使われているようだ。

RC4 および CBC モードは危殆化が報告されて久しく，推奨されない。
また TripleDES, Blowfish, CAST はセキュリティ強度不足のため，これも推奨されない。

優先順位としては AEAD (Authenticated Encryption with Associated Data; 認証付き暗号) として機能する GCM や ChaCha20-Poly1305 がより推奨されているらしい。
なお，これらの暗号スイートの選択と優先順位はサーバ側で設定するのでユーザが意識することはない。

## メッセージ認証符号

GCM や ChaCha20-Poly1305 以外の AEAD として機能しない暗号モードの場合は MAC (Message Authentication Code; メッセージ認証符号) を組み合わせる。
組み合わせ方としては

- Encrypt-then-MAC ([OpenSSH] 6.2 以降)
- MAC-then-encrypt
- Encrypt-and-MAC

のいずれかを選択できるらしい。
Encrypt-then-MAC であれば AEAD として機能する。
[OpenSSH] でサポートしてる MAC アルゴリズムとハッシュ・アルゴリズムの組み合わせは以下の通り。

| 名称            | MAC      | ハッシュ         |         推奨          | 備考               |
| --------------- | -------- | ---------------- |:---------------------:| ------------------ |
| `hmac-md5`      | HMAC     | MD5              |                       |                    |
| `hmac-md5-96`   | HMAC     | MD5              |                       |                    |
| `hmac-sha1`     | HMAC     | SHA1             |                       |                    |
| `hmac-sha1-96`  | HMAC     | SHA1             |                       |                    |
| `hmac-sha2-256` | HMAC     | SHA2-256         | {{< icons "check" >}} | [OpenSSH] 5.9 以降 |
| `hmac-sha2-512` | HMAC     | SHA2-512         | {{< icons "check" >}} | [OpenSSH] 5.9 以降 |
| `umac-64`       | UMAC-64  | (Universal Hash) |                       | [OpenSSH] 4.7 以降 |
| `umac-128`      | UMAC-128 | (Universal Hash) |                       | [OpenSSH] 6.2 以降 |

それぞれの名称の後ろに `-etm`@openssh.com と付くと Encrypt-then-MAC で動作する。

MD5 は危殆化が報告されて久しく，推奨されない。 
SHA1 は HMAC に関しては2031年以降も “Acceptable” であるとされているが，他の暗号スイートとの組み合わせで考えると避けたほうがいいかもしれない。

UMAC は [RFC 4418] で規定されている。
この中で

{{< fig-quote type="markdown" title="UMAC: Message Authentication Code using Universal Hashing" link="https://tools.ietf.org/html/rfc4418" lang="en" >}}
{{% quote %}}Likewise, 32-, 96-, and 128-bit tags cannot be forged with more than 1/2^30, 1/2^90, and 1/2^120 probability plus the   probability of a successful attack against AES as a pseudorandom permutation{{% /quote %}}.
{{< /fig-quote >}}

とあるので UMAC-64 や UMAC-128 はセキュリティ強度不足なんじゃないかと思うのだが，どうだろう。

これらの暗号スイートの選択と優先順位はサーバ側で設定するのでユーザが意識することはない。

## 電子署名アルゴリズムとハッシュ・アルゴリズム

クライアント認証で公開鍵暗号を用いる場合は，ユーザ側で認証鍵ペアを生成し，公開鍵をあらかじめサーバ側と共有する必要がある。
また公開鍵暗号を用いたクライアント認証ではハッシュ・アルゴリズムも用いる。
[OpenSSH] でサポートしてるクライアント認証の暗号スイートは以下の通り。

| 名称（サーバ側）                     | 電子署名                 | ハッシュ |         推奨          | 備考               |
| ------------------------------------ | ------------------------ | -------- |:---------------------:| ------------------ |
| `ssh-ed25519`                        | EdDSA (edwards25519)     | SHA2-256 | {{< icons "check" >}} |                    |
| `sk-ssh-ed25519`@openssh.com         | EdDSA (edwards25519)     | SHA2-256 | {{< icons "check" >}} | [OpenSSH] 8.2 以降 |
| `ssh-rsa`                            | RSA                      | SHA1     |                       | 廃止予定           |
| `rsa-sha2-256`                       | RSA                      | SHA2-256 | {{< icons "check" >}} | [OpenSSH] 7.2 以降 |
| `rsa-sha2-512`                       | RSA                      | SHA2-512 | {{< icons "check" >}} | [OpenSSH] 7.2 以降 |
| `ssh-dss`                            | DSA                      | SHA1     |                       |                    |
| `ecdsa-sha2-nistp256`                | ECDSA (NIST curve P-256) | SHA2-256 |                       |                    |
| `ecdsa-sha2-nistp384`                | ECDSA (NIST curve P-384) | SHA2-384 |                       |                    |
| `ecdsa-sha2-nistp521`                | ECDSA (NIST curve P-521) | SHA2-512 |                       |                    |
| `sk-ecdsa-sha2-nistp256`@openssh.com | ECDSA (NIST curve P-256) | SHA2-256 |                       | [OpenSSH] 8.2 以降 |

頭に `sk-` が付いているものは暗号デバイスに対応している。

`ssh-rsa` は将来バージョンで廃止が決まっている。

[OpenSSH] では FIPS 186-3 以降の DSA に対応してないようで，鍵長が1024ビットしか対応してない（これも NIST 嫌悪か？）。
もちろん推奨できない。

NIST 推奨パラメータを使った ECDSA も（ECDH と同じく）微妙に評判が悪い。
あと ECDSA は電子署名時に乱数を使うのだが，この実装をサボると脆弱性の元となる。
実際に，2013年に発覚した[疑似乱数生成器 Dual EC DRBG の脆弱性](https://www.cryptrec.go.jp/topics/cryptrec-er-0001-2013.html "擬似乱数生成アルゴリズム Dual_EC_DRBG について | CRYPTREC")では， NSA が絡んでいたこともあって，一気に ECDSA 忌避が強まったらしい。

以上はサーバ側の設定の話だが，ユーザ側では `ssh-keygen` コマンドを使って認証鍵を生成する必要がある。

認証鍵の生成では鍵種別（`-t` オプション）と鍵長（`-b` オプション）を指定する。
組み合わせは以下の通り。

| 鍵種別       | 鍵長                     | アルゴリズム             |         推奨          | 備考               |
| ------------ | ------------------------ | ------------------------ |:---------------------:| ------------------ |
| `rsa`        | 1024以上<br>（既定3072） | RSA                      | {{< icons "check" >}} | 3072ビット以上推奨 |
| `dsa`        | 1024                     | DSA                      |                       |                    |
| `ed25519`    | &mdash;                  | EdDSA (edwards25519)     | {{< icons "check" >}} |                    |
| `ed25519-sk` | &mdash;                  | EdDSA (edwards25519)     | {{< icons "check" >}} | [OpenSSH] 8.2 以降 |
| `ecdsa`      | 256（既定）              | ECDSA (NIST curve P-256) |                       |                    |
| `ecdsa`      | 384                      | ECDSA (NIST curve P-384) |                       |                    |
| `ecdsa`      | 521                      | ECDSA (NIST curve P-521) |                       |                    |
| `ecdsa-sk`   | &mdash;                  | ECDSA (NIST curve P-256) |                       | [OpenSSH] 8.2 以降 |

これは私が盛大に勘違いしていたのだが，たとえば

```text
$ ssh-keygen -t rsa
Generating public/private rsa key pair.
Enter file in which to save the key (/home/username/.ssh/id_rsa): 
Enter passphrase (empty for no passphrase): 
Enter same passphrase again: 
Your identification has been saved in /home/username/.ssh/id_rsa
Your public key has been saved in /home/username/.ssh/id_rsa.pub
The key fingerprint is:
SHA256:qufsNjgco3QZNjE4eupwQiT6mD8fr2a7nXmU3ybxFHo username@hostname
The key's randomart image is:
+---[RSA 3072]----+
|   .             |
|..o o            |
|+. . o           |
|o.. +       .    |
|.* . +  S. . .   |
|* + =  .o o E    |
|o= +.+.. . *     |
| .+ **+o. o +    |
|   =*XOo   o     |
+----[SHA256]-----+
```

てな感じに鍵を作ったときに表示される `SHA256` は署名時ではなく，鍵指紋のハッシュ・アルゴリズムを指しているらしい。
実際の電子署名でどのハッシュ・アルゴリズムを使うのかはサーバ-クライアント間のネゴシエーションで決まるのかな（？）

## 【おまけ】 各種アルゴリズムのセキュリティ強度

（「[暗号鍵関連の各種変数について]({{< ref "/remark/2017/10/key-parameters.md" >}})」より抜粋）

おまけとして各種アルゴリズムのセキュリティ強度を挙げておく。

### 各種暗号アルゴリズムとセキュリティ強度の関係

最初は種暗号アルゴリズムの鍵長とセキュリティ強度の関係を示す表。
単位は全てビットである。

{{< comparable-security-strengths >}} <!-- 要 MathJax -->

### ハッシュ・アルゴリズムとセキュリティ強度の関係

次はハッシュ・アルゴリズムとセキュリティ強度の関係を示す表。

{{< security-strengths-for-hash >}} <!-- 要 MathJax -->

### セキュリティ強度と有効期限

こちらはセキュリティ強度の有効期限を表したものだ。

{{< security-strength-time-frames >}} <!-- 要 MathJax -->

各用語はそれぞれ

| 用語       | 意味                     |
|:---------- |:------------------------ |
| Applying   | 適用                     |
| Processing | 処理                     |
| Acceptable | 許容                     |
| Legacy use | 許容（レガシー使用のみ） |
| Disallowed | 禁止                     |

という意味だ。
例を挙げると，セキュリティ強度112ビットの暗号スイート（Cipher Suites）を適用する場合は2030年までは許容するけど2031年以降は禁止。
すでに暗号化されているデータを復号したい場合でも2031年以降はレガシー・システムしか許容しない，ということになる。

## ブックマーク

- [OpenSSHの暗号化周りの設定について - Qiita](https://qiita.com/aqmr-kino/items/8c3306ea8022b0d5cbe4)
- [ChaCha20-Poly1305の解説と実装 | 晴耕雨読](https://tex2e.github.io/blog/crypto/chacha20poly1305)
- [OpenSSHの認証に証明書を使う方法｜ConoHa VPSサポート](https://support.conoha.jp/v/openssh/)

[OpenSSH]: https://www.openssh.com/
[RFC 4418]: https://tools.ietf.org/html/rfc4418 "RFC 4418 - UMAC: Message Authentication Code using Universal Hashing"

## 参考図書

{{% review-paapi "B079NL1L9K" %}} <!-- SSH Mastery -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
