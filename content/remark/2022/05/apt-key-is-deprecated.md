+++
title = "ついに apt-key コマンドに「非推奨」のワーニングが"
date =  "2022-05-06T18:10:02+09:00"
description = "しょうがない。 この機会に対応するか。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "management", "tools", "install", "openpgp" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]({{< relref "./upgrade-ubuntu-22_04-lts.md" >}} "Ubuntu 22.04 LTS へのアップグレード")の続き。
[Docker] を [Ubuntu] にインストールするには以下の拙文が参考になる。

- [Ubuntu に Docker を入れる]({{< ref "/remark/2020/12/installing-docker-in-ubuntu.md" >}})

簡単に説明すると

```text
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

と電子署名検証用の OpenPGP 鍵をインポートした上で

```text
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```

とすれば ([Ubuntu] 22.04 LTS であれば) jammy 用のリポジトリが追加される。
実際には `/etc/apt/sources.list.d/` ディレクトリにある `archive_uri-https_download_docker_com_linux_ubuntu-jammy.list` ファイルに

```text
deb [arch=amd64] https://download.docker.com/linux/ubuntu jammy stable
```

という内容が書かれている（筈）。

ところが今回 [Ubuntu] 22.04 LTS では，この状態で `apt update` すると

```text {hl_lines="9"}
$ sudo apt update
...
ヒット:5 https://download.docker.com/linux/ubuntu jammy InRelease
...
パッケージリストを読み込んでいます... 完了
依存関係ツリーを作成しています... 完了
状態情報を読み取っています... 完了
パッケージはすべて最新です。
W: https://download.docker.com/linux/ubuntu/dists/jammy/InRelease: Key is stored in legacy trusted.gpg keyring (/etc/apt/trusted.gpg), see the DEPRECATION section in apt-key(8) for details.
```

とワーニングが出る。
あぁ，ついに [Ubuntu] でもこのワーニングが出るようになったか。
今まで先延ばしにしてたからなぁ。
しょうがない。
この機会に対応するか。

ちなみに `apt-key(8)` のマニュアルには

{{< fig-quote type="markdown" title="apt-key(8)" lang="en" >}}
Except for using `apt-key del` in maintainer scripts, the use of `apt-key` is deprecated. This section shows how to replace existing use of `apt-key`.

If your existing use of `apt-key add` looks like this:

```
wget -qO- https://myrepo.example/myrepo.asc | sudo apt-key add -
```

Then you can directly replace this with (though note the recommendation below):

```
wget -qO- https://myrepo.example/myrepo.asc | sudo tee /etc/apt/trusted.gpg.d/myrepo.asc
```

Make sure to use the "`asc`" extension for ASCII armored keys and the "`gpg`" extension for the binary OpenPGP format (also known as "GPG key public ring"). The binary OpenPGP format works for all apt versions, while the ASCII armored format works for apt version >= 1.4.

Recommended: Instead of placing keys into the `/etc/apt/trusted.gpg.d` directory, you can place them anywhere on your filesystem by using the `Signed-By` option in your `sources.list` and pointing to the filename of the key. See sources.list(5) for details. Since APT 2.4, `/etc/apt/keyrings` is provided as the recommended location for keys not managed by packages. When using a deb822-style sources.list, and with apt version >= 2.4, the `Signed-By` option can also be used to include the full ASCII armored keyring directly in the `sources.list` without an additional file.
{{< /fig-quote >}}

などと書かれている。
要するに `/etc/apt/keyrings` ディレクトリに公開鍵を ASCII Armor 形式のまま放り込んでしまえばいいのか？

んー。
こんな感じかな。

```text
$ cd /etc/apt/keyrings
$ sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o docker-key.asc
```

ファイル名は適当。
念のため出力した `docker-key.asc` ファイルの中身を拙作 [gpgpdump] で覗いてみよう。

```text
$ gpgpdump -u -f docker-key.asc
Public-Key Packet (tag 6) (525 bytes)
  Version: 4 (current)
  Public key creation time: 2017-02-22T18:36:26Z
  Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
  RSA public modulus n (4096 bits)
  RSA public encryption exponent e (17 bits)
User ID Packet (tag 13) (43 bytes)
  User ID: Docker Release (CE deb) <docker@docker.com>
Signature Packet (tag 2) (567 bytes)
  Version: 4 (current)
  Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
  Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
  Hash Algorithm: SHA2-512 (hash 10)
  Hashed Subpacket (33 bytes)
    Signature Creation Time (sub 2): 2017-02-22T19:34:24Z
    Key Flags (sub 27) (1 bytes)
      Flag: This key may be used to certify other keys.
      Flag: This key may be used to sign data.
      Flag: This key may be used to encrypt communications.
      Flag: This key may be used to encrypt storage.
      Flag: This key may be used for authentication.
    Preferred Symmetric Algorithms (sub 11) (4 bytes)
      Symmetric Algorithm: AES with 256-bit key (sym 9)
      Symmetric Algorithm: AES with 192-bit key (sym 8)
      Symmetric Algorithm: AES with 128-bit key (sym 7)
      Symmetric Algorithm: CAST5 (128 bit key, as per) (sym 3)
    Preferred Hash Algorithms (sub 21) (4 bytes)
      Hash Algorithm: SHA2-512 (hash 10)
      Hash Algorithm: SHA2-384 (hash 9)
      Hash Algorithm: SHA2-256 (hash 8)
      Hash Algorithm: SHA2-224 (hash 11)
    Preferred Compression Algorithms (sub 22) (4 bytes)
      Compression Algorithm: ZLIB <RFC1950> (comp 2)
      Compression Algorithm: BZip2 (comp 3)
      Compression Algorithm: ZIP <RFC1951> (comp 1)
      Compression Algorithm: Uncompressed (comp 0)
    Features (sub 30) (1 bytes)
      Flag: Modification Detection (packets 18 and 19)
    Key Server Preferences (sub 23) (1 bytes)
      Flag: No-modify
  Unhashed Subpacket (10 bytes)
    Issuer (sub 16): 0x8d81803c0ebfcd88
  Hash left 2 bytes
    b2 c9
  RSA signature value m^d mod n (4094 bits)
Public-Subkey Packet (tag 14) (525 bytes)
  Version: 4 (current)
  Public key creation time: 2017-02-22T18:36:26Z
  Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
  RSA public modulus n (4096 bits)
  RSA public encryption exponent e (17 bits)
Signature Packet (tag 2) (1086 bytes)
  Version: 4 (current)
  Signiture Type: Subkey Binding Signature (0x18)
  Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
  Hash Algorithm: SHA2-256 (hash 8)
  Hashed Subpacket (9 bytes)
    Signature Creation Time (sub 2): 2017-02-22T18:36:26Z
    Key Flags (sub 27) (1 bytes)
      Flag: This key may be used to sign data.
  Unhashed Subpacket (553 bytes)
    Issuer (sub 16): 0x8d81803c0ebfcd88
    Embedded Signature (sub 32) (540 bytes)
      Signature Packet (tag 2) (540 bytes)
        Version: 4 (current)
        Signiture Type: Primary Key Binding Signature (0x19)
        Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
        Hash Algorithm: SHA2-256 (hash 8)
        Hashed Subpacket (6 bytes)
          Signature Creation Time (sub 2): 2017-02-22T18:36:26Z
        Unhashed Subpacket (10 bytes)
          Issuer (sub 16): 0x7ea0a9c3f273fcd8
        Hash left 2 bytes
          d5 60
        RSA signature value m^d mod n (4095 bits)
  Hash left 2 bytes
    f2 b8
  RSA signature value m^d mod n (4095 bits)
```

もともと `/etc/apt/trusted.gpg` ファイルに入ってる公開鍵は

```text
$ apt-key list
Warning: apt-key is deprecated. Manage keyring files in trusted.gpg.d instead (see apt-key(8)).
/etc/apt/trusted.gpg
--------------------
pub   rsa4096 2017-02-22 [SCEA]
      9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
uid           [  不明  ] Docker Release (CE deb) <docker@docker.com>
sub   rsa4096 2017-02-22 [S]
...
```

となっているので，同じ鍵ということでいいよね。

ここで `/etc/apt/sources.list.d/` ディレクトリにある `archive_uri-https_download_docker_com_linux_ubuntu-jammy.list` ファイルの中身を以下のように書き換える。

```text
deb [arch=amd64 signed-by=/etc/apt/keyrings/docker-key.asc] https://download.docker.com/linux/ubuntu jammy stable
```

ポイントは `signed-by` オプション。
このオプションに先程の公開鍵ファイルをフルパスで指定する。
`add-apt-repository` コマンドを使うなら

```text
sudo add-apt-repository "deb [arch=amd64 signed-by=/etc/apt/keyrings/docker-key.asc] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```

という感じだろうか。これでもう一度 `apt update` すると

```text
$ sudo apt update
...
ヒット:5 https://download.docker.com/linux/ubuntu jammy InRelease
...
パッケージリストを読み込んでいます... 完了
依存関係ツリーを作成しています... 完了
状態情報を読み取っています... 完了
パッケージはすべて最新です。
```

よし。
ワーニングは出なくなったな。

なんでこんな面倒くさいことをするかというと，もし秘密鍵が漏洩したサードパーティの公開鍵がひとつでも APT の鍵束に入っていると，その鍵を使って malware 入れ放題になる可能性があるからだ[^key1]。
これによってサードパーティの公開鍵が作用するパッケージを限定することができる。
簡単に言うと「もうサードパーティの公開鍵は APT で管理しないから自分でやれ！ 管理する仕組みは提供するから」ということである。
まぁ，苦肉の策って感じだけどね（笑）

[^key1]: OpenPGP 鍵の信用モデルは「信用の輪（web of trust）」と呼ばれるもので，ユーザ同士の peer で直接的な関係が鍵管理の前提になっている。このため不特定のユーザの間ではどうしても鍵管理が緩くなってしまう。この辺は不特定ユーザを前提にした X.509 とは思想が異なる。 OpenPGP 鍵の信用モデルについては拙文「[OpenPGP の電子署名は「ユーザーの身元を保証し」ない]({{< ref "/openpgp/web-of-trust.md" >}})」あたりを参考にしていただけるとありがたい。

最後に `/etc/apt/trusted.gpg` ファイルを削除するか他所に退避させておけば問題ないであろう。

## ブックマーク

- [Ubuntu Manpage: apt-key - APT キー管理ユーティリティ](https://manpages.ubuntu.com/manpages/trusty/ja/man8/apt-key.8.html)
- [非推奨となったapt-keyの代わりにsigned-byとgnupgを使う方法 - 2021-05-05 - ククログ](https://www.clear-code.com/blog/2021/5/5.html)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Docker]: https://www.docker.com/ "Empowering App Development for Developers | Docker"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "4900900028" %}} <!-- PGP―暗号メールと電子署名 -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
