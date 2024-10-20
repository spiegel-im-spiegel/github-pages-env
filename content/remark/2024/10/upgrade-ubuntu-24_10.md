+++
title = "Ubuntu 24.10 へのアップグレード"
date =  "2024-10-19T15:17:47+09:00"
description = "OS のアップグレードは特に問題もなく完了。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools", "install", "ppa" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り 2024-10 に [Ubuntu] 24.10 がリリースされた。

- [Canonical Releases Ubuntu 24.10 Oracular Oriole | Ubuntu](https://ubuntu.com/blog/canonical-releases-ubuntu-24-10-oracular-oriole)
- [Oracular Oriole Release Notes - Release - Ubuntu Community Hub](https://discourse.ubuntu.com/t/oracular-oriole-release-notes/44878)

[Ubuntu Japanese Team](https://www.ubuntulinux.jp/) からのアナウンスはまだないな。
24.04 LTS から日本語 Remix のリリースがなくなってるし，日本語固有のローカライズはなくなっていくのかねぇ。

というわけでアップグレードを始めようか。

まずは以下のページから ISO イメージをダウンロードして，転ばぬ先のブータブル USB を作成する。

- [Download Ubuntu Desktop | Ubuntu](https://ubuntu.com/download/desktop)

念のためのバックアップも取ってソフトウェアの更新も行って，以下のダイアログからアップグレードを開始する。

{{< fig-img src="./upgrade-ubuntu-0.png" title="ソフトウェアの更新" link="./upgrade-ubuntu-0.png" width="645" >}}

{{< fig-img src="./upgrade-ubuntu-1.png" title="リリースノート" link="./upgrade-ubuntu-1.png" width="628" >}}

あとは流れに任せる（笑） 特に問題もなく完了。

{{< fig-img src="./upgrade-ubuntu-8.png" title="アップグレード完了" link="./upgrade-ubuntu-8.png" width="694" >}}

## セキュリティ関連ツールのバージョン

まずは [GnuPG]。

```text
$ gpg --version
gpg (GnuPG) 2.4.4
libgcrypt 1.11.0
Copyright (C) 2024 g10 Code GmbH
License GNU GPL-3.0-or-later <https://gnu.org/licenses/gpl.html>
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

[24.04 アップグレード]({{< ref "/remark/2024/06/upgrade-ubuntu-24_04-lts.md" >}} "Ubuntu 24.04 LTS へのアップグレード")から上がっとらんがな！

まぁいいか。
次は [OpenSSH] と [OpenSSL]。

```text
$ ssh -V
OpenSSH_9.7p1 Ubuntu-7ubuntu4, OpenSSL 3.3.1 4 Jun 2024
```

これも古いが... もうどうでもいいや。
バックポートパッチはあたっていると信じよう（投げやり）。

## サードパーティ用 APT リポジトリ定義の確認と更新

OS のアップグレードよりこっちの作業のほうが時間がかかったり。
とりあえず TODO を残しておく（随時更新）。

- [x] [Docker] Engine
  - [Install Docker Engine on Ubuntu | Docker Docs](https://docs.docker.com/engine/install/ubuntu/)
- [x] [PPA] 版 [Git]
  - [Git stable releases : “Ubuntu Git Maintainers” team](https://launchpad.net/~git-core/+archive/ubuntu/ppa)
- [x] [NodeSource] 版 [Node.js]
  - [NodeSource Node.js Binary Distributions](https://github.com/nodesource/distributions/blob/master/README.md)
    - v22 へアップグレードした
- [ ] [pgAdmin 4]
  - [Download](https://www.pgadmin.org/download/pgadmin-4-apt/)
- [ ] [PPA] 版 [KeePassXC]
  - [KeePassXC : Janek Bevendorff](https://launchpad.net/~phoerious/+archive/ubuntu/keepassxc)
- [x] [VS Code]
  - [Running Visual Studio Code on Linux](https://code.visualstudio.com/docs/setup/linux)

### APT バージョン

アップグレード時の APT のバージョンは以下の通り。

```text
$ apt -v
apt 2.9.8 (amd64)
```

[24.04 アップグレード]({{< ref "/remark/2024/06/upgrade-ubuntu-24_04-lts.md" >}} "Ubuntu 24.04 LTS へのアップグレード")時のバージョン（v2.7.14）では RSA 1024 ビット鍵に対してワーニングを出していた。
このバージョンでもワーニングを出すのは変わらないみたいだが（エラーじゃなかった）

```text
$ sudo apt update
...

Warning: http://ppa.launchpad.net/git-core/ppa/ubuntu/dists/oracular/InRelease: Signature by key E1DD270288B4E6030699E45FA1715D88E1DF1F24 uses weak algorithm (rsa1024)
```

`/etc/apt/apt.conf.d/99weakkey-warning` ファイルを使ったワーニングの回避はできなくなったぽい。

### PPA リポジトリの再定義

以前の PPA リポジトリの公開鍵は RSA 1024 ビット鍵だったが，新しいものに替えたようだ。
`add-apt-repository` コマンドでリポジトリ定義を再取得する。
`ppa:git-core/ppa` の場合はこんな感じ。

```text
$ sudo add-apt-repository ppa:git-core/ppa
```

これで `/etc/apt/sources.list.d/` ディレクトリに `git-core-ubuntu-ppa-oracular.sources` ファイルが作成される。
中身はこんな感じ。

```text
Types: deb
URIs: https://ppa.launchpadcontent.net/git-core/ppa/ubuntu/
Suites: oracular
Components: main
Signed-By: -----BEGIN PGP PUBLIC KEY BLOCK-----
 .
 mQINBGYo2OYBEADVRjI+o29u9izslaVr0Xqj8hpmo/2su/Iey1PgoS6A3hMxR4R4
 eZ3u9dRh/gRHXNjxqRMfKj88G6ciqa/7ty8Vfc1eKl3z7yjL1pWOEzcGLKaSB8qd
 MmsxCw31nFNEbzlymgK0+KPubQ5OrIzeSikpfDVGT4HLgO42ppGY+cVy2/bbNv6O
 mmPXcw8gkxRCWFiGAO5jJYG1SyGbhr9Krbf6o+LDUJeDYPTQRMf702IYZ8Bp00ix
 HyK2YOUNM14rr7092o2dw9GKxnJszF4cET+LxRddrREuB3sBlAZav/I0hZtQsKZ3
 QTvpStf2MwIy8Ymj94+BsZaktP0d36wigGn8RWJHrhSVyjujS8YxWRWdjrLUI1ra
 42pyIYZi39IXtsTM71rihQVrsEHbMQ7a5HGRK6eYyHVPrtsXXykNqhVPekLNiFWm
 IekoSH1EwMsQv8y+k3nkCwTCkTHJaueB7awee0Zs5QBwbCm+qPyqCXoVDLEwdQOr
 CHKAfNADtsFQsk/yiTC/+Lmtur/okp38VpJWXg8DphHFjd1KwqQ5E9qZi+tXs/JD
 UXd93VBUdoGGaHK9fj/URxUBOVopGaOXGVYtGFWPn9q8MaNcrESZ48sXDfsVgUVD
 RJ4puKLHjIUtDlCnMQO6lekIhEo4sxtbDnwIUmQp7B9l+U99u+uzBI96cQARAQAB
 tChMYXVuY2hwYWQgUFBBIGZvciBVYnVudHUgR2l0IE1haW50YWluZXJziQJOBBMB
 CgA4FiEE+RGrGEMXYwxZlwlz42PJD48bYhcFAmYo2OYCGwMFCwkIBwIGFQoJCAsC
 BBYCAwECHgECF4AACgkQ42PJD48bYhdwUQ//UFR3i/6zpizJMTA9mpz7hGC9IJV8
 UDoWoaVgl+OR1Ldfz+jvr3K55LZyIMU1o6bbLqbEnoWa2VpRv2za/SCbPqo1igio
 p97EJ2irGytFOhCDd+o3s0djfsXpA7jygAK6COnMx3ejnPhaBA194HbYDhp7KA5b
 gZcqvYeRN1qk9QL99voFYeUWAPnqkLLrNuAcq9qTotmyYZQI61rdAH8P4odgMtU7
 UiS4yLirkAiNCqT+TK07EXvaWSXcqZhgt1HmP+BZhrx/vLT4FlH52CCjamQWeFA2
 mLKX/RSx/JTux1UDroX6L9JdUyMzOrLk22Zz9Tb3FK2ysHy72jRKmUv87ctIMcgD
 u8BY3Mfe/Rw8vPMuBEaAsVfvWl7uYSt81dzjkxtt6ZvIFF5PBR2IhFJVosDbpSnk
 g9/b22Ipjta3ULW8oOaNZjdcRNQ5StpApzZIUoZoP83ZWafwQoSrW7Rz3KvwYAdL
 d3OAfuiW9i7YdLCkvaujBt6KA4tn56fcsp14FzLZrgcj+7XUpW4/yEJFHCCu5f2l
 NWvN37suUTfzbMLZVS6rC1s3qrjOD+C3dvL/dCUlcTfrrsTcs/UnaVXFq0V6NLhA
 ZZxOIVUedn7nGKbaecwTt6taIjpxj0jCBxWy6RcysIkv2xluRcl2UY+HapB8x1Dx
 8giHUSvkXHr4c7I=
 =yxbG
 -----END PGP PUBLIC KEY BLOCK-----
```

ちなみに OpenPGP 公開鍵の部分を取り出すとこんな感じ。

```text
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQINBGYo2OYBEADVRjI+o29u9izslaVr0Xqj8hpmo/2su/Iey1PgoS6A3hMxR4R4
eZ3u9dRh/gRHXNjxqRMfKj88G6ciqa/7ty8Vfc1eKl3z7yjL1pWOEzcGLKaSB8qd
MmsxCw31nFNEbzlymgK0+KPubQ5OrIzeSikpfDVGT4HLgO42ppGY+cVy2/bbNv6O
mmPXcw8gkxRCWFiGAO5jJYG1SyGbhr9Krbf6o+LDUJeDYPTQRMf702IYZ8Bp00ix
HyK2YOUNM14rr7092o2dw9GKxnJszF4cET+LxRddrREuB3sBlAZav/I0hZtQsKZ3
QTvpStf2MwIy8Ymj94+BsZaktP0d36wigGn8RWJHrhSVyjujS8YxWRWdjrLUI1ra
42pyIYZi39IXtsTM71rihQVrsEHbMQ7a5HGRK6eYyHVPrtsXXykNqhVPekLNiFWm
IekoSH1EwMsQv8y+k3nkCwTCkTHJaueB7awee0Zs5QBwbCm+qPyqCXoVDLEwdQOr
CHKAfNADtsFQsk/yiTC/+Lmtur/okp38VpJWXg8DphHFjd1KwqQ5E9qZi+tXs/JD
UXd93VBUdoGGaHK9fj/URxUBOVopGaOXGVYtGFWPn9q8MaNcrESZ48sXDfsVgUVD
RJ4puKLHjIUtDlCnMQO6lekIhEo4sxtbDnwIUmQp7B9l+U99u+uzBI96cQARAQAB
tChMYXVuY2hwYWQgUFBBIGZvciBVYnVudHUgR2l0IE1haW50YWluZXJziQJOBBMB
CgA4FiEE+RGrGEMXYwxZlwlz42PJD48bYhcFAmYo2OYCGwMFCwkIBwIGFQoJCAsC
BBYCAwECHgECF4AACgkQ42PJD48bYhdwUQ//UFR3i/6zpizJMTA9mpz7hGC9IJV8
UDoWoaVgl+OR1Ldfz+jvr3K55LZyIMU1o6bbLqbEnoWa2VpRv2za/SCbPqo1igio
p97EJ2irGytFOhCDd+o3s0djfsXpA7jygAK6COnMx3ejnPhaBA194HbYDhp7KA5b
gZcqvYeRN1qk9QL99voFYeUWAPnqkLLrNuAcq9qTotmyYZQI61rdAH8P4odgMtU7
UiS4yLirkAiNCqT+TK07EXvaWSXcqZhgt1HmP+BZhrx/vLT4FlH52CCjamQWeFA2
mLKX/RSx/JTux1UDroX6L9JdUyMzOrLk22Zz9Tb3FK2ysHy72jRKmUv87ctIMcgD
u8BY3Mfe/Rw8vPMuBEaAsVfvWl7uYSt81dzjkxtt6ZvIFF5PBR2IhFJVosDbpSnk
g9/b22Ipjta3ULW8oOaNZjdcRNQ5StpApzZIUoZoP83ZWafwQoSrW7Rz3KvwYAdL
d3OAfuiW9i7YdLCkvaujBt6KA4tn56fcsp14FzLZrgcj+7XUpW4/yEJFHCCu5f2l
NWvN37suUTfzbMLZVS6rC1s3qrjOD+C3dvL/dCUlcTfrrsTcs/UnaVXFq0V6NLhA
ZZxOIVUedn7nGKbaecwTt6taIjpxj0jCBxWy6RcysIkv2xluRcl2UY+HapB8x1Dx
8giHUSvkXHr4c7I=
=yxbG
-----END PGP PUBLIC KEY BLOCK-----
```

これをクリップボードに入れて拙作 [`gpgpdump`](https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer") で見てみる。

```text
$ gpgpdump --clipboard
Public-Key Packet (tag 6) (525 bytes)
    Version: 4 (current)
    Public key creation time: 2024-04-24T19:03:18+09:00
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (4096 bits)
    RSA public encryption exponent e (17 bits)
User ID Packet (tag 13) (40 bytes)
    User ID: Launchpad PPA for Ubuntu Git Maintainers
Signature Packet (tag 2) (590 bytes)
    Version: 4 (current)
    Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA2-512 (hash 10)
    Hashed Subpacket (56 bytes)
        Issuer Fingerprint (sub 33) (21 bytes)
            Version: 4 (need 20 octets length)
            Fingerprint (20 bytes)
                f9 11 ab 18 43 17 63 0c 59 97 09 73 e3 63 c9 0f 8f 1b 62 17
        Signature Creation Time (sub 2): 2024-04-24T19:03:18+09:00
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to certify other keys.
            Flag: This key may be used to sign data.
        Preferred Symmetric Ciphers for v1 SEIPD (sub 11) (4 bytes)
            Symmetric Algorithm: AES with 256-bit key (sym 9)
            Symmetric Algorithm: AES with 192-bit key (sym 8)
            Symmetric Algorithm: AES with 128-bit key (sym 7)
            Symmetric Algorithm: TripleDES (168 bit key derived from 192) (sym 2)
        Preferred Hash Algorithms (sub 21) (5 bytes)
            Hash Algorithm: SHA2-512 (hash 10)
            Hash Algorithm: SHA2-384 (hash 9)
            Hash Algorithm: SHA2-256 (hash 8)
            Hash Algorithm: SHA2-224 (hash 11)
            Hash Algorithm: SHA-1 (hash 2)
        Preferred Compression Algorithms (sub 22) (3 bytes)
            Compression Algorithm: ZLIB <RFC1950> (comp 2)
            Compression Algorithm: BZip2 (comp 3)
            Compression Algorithm: ZIP <RFC1951> (comp 1)
        Features (sub 30) (1 bytes)
            Flag: Symmetrically Encrypted Integrity Protected Data packet version 1
        Key Server Preferences (sub 23) (1 bytes)
            Flag: No-modify
    Unhashed Subpacket (10 bytes)
        Issuer Key ID (sub 16): 0xe363c90f8f1b6217
    Hash left 2 bytes
        70 51
    RSA signature value m^d mod n (4095 bits)

```

おー。
RSA 4096ビット鍵か。
過剰にデカいな[^c1]。
どうせ [GnuPG] を使うんだから EdDSA 鍵でいいぢゃん，と思うのだけど。
まぁ，いいか。

[^c1]: RSA なら3072ビットもあれば2031年以降も使える。ちなみに EdDSA を含む楕円曲線暗号なら256ビットもあれば必要十分である。詳しくは拙文「[暗号鍵関連の各種変数について]({{< ref "/remark/2017/10/key-parameters.md" >}})」を参照のこと。

という感じで地道にリポジトリ定義を作り直した。

## ブックマーク

- [Ubuntu 24.10 “Oracular Oriole”のリリース | gihyo.jp](https://gihyo.jp/admin/clip/01/ubuntu-topics/202410/11)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: https://www.openssh.com/
[OpenSSL]: https://www.openssl.org/
[Git]: https://git-scm.com/
[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[NodeSource]: https://github.com/nodesource
[Node.js]: https://nodejs.org/
[Docker]: https://www.docker.com/ "Empowering App Development for Developers | Docker"
[pgAdmin 4]: https://www.pgadmin.org/ "pgAdmin - PostgreSQL Tools"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Fcitx 5]: https://fcitx-im.org/wiki/Fcitx_5

## 参考

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B01MS6RPN2" %}} <!-- シリコンパワー USBメモリ 8GB USB3.1 -->
{{% review-paapi "B0DB5N3HJL" %}} <!-- ダンジョンの中の人 マイクロレボリューション TrySail -->
