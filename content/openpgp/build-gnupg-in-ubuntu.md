+++
title = "最新版 GnuPG をビルドする： GnuPG in Ubuntu"
date = "2021-01-01T13:18:27+09:00"
description = "念のために警告しておくと，今回の自力ビルドは積極的にはお勧めしない。"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "gnupg", "ubuntu", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Linux 環境では [GnuPG] は（ユーザは意識しないかもしれないが）セキュリティ中核部品のひとつだが， [Ubuntu] はかなり保守的な運用をしているそうで，滅多なことでは [GnuPG] をアップデートしないようだ。
かといって CVE ID が割り振られるような脆弱性まで長期間放置するのはいかがなものかと思うが。

そこで，かねてから懸案だった [GnuPG] の自力ビルドを試してみることにした。

念のために警告しておくと，今回の自力ビルドは積極的にはお勧めしない。
[GnuPG] の最新機能を試したいとか，何がなんでも最新版が欲しいとかいうのでない限り APT で配布されるバージョンを使うほうがよい。
まぁ「最悪はこういう手段も取れるよ」という感じで記憶の片隅にでも置いていただければ（笑）

## ビルド対象のパッケージ

[GnuPG] は複数のパッケージで構成されていて，個別にビルド&インストールしていく必要がある。
今回ビルド対象となるパッケージは以下の通り。

|   # | パッケージ名                                             | バージョン | 公開日     |
| ---:| -------------------------------------------------------- | ---------- | ---------- |
|   1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.41       | 2020-12-21 |
|   2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.7      | 2020-10-23 |
|   3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.4      | 2020-10-23 |
|   4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.5.0      | 2020-11-18 |
|   5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6        | 2018-07-16 |
|   6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0      | 2020-08-27 |
|   7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.26     | 2020-12-21 |

ビルド&インストールの順番は [Libgpg-error](https://gnupg.org/software/libgpg-error/) を最初にして [GnuPG](https://gnupg.org/software/) を最後にすることさえ気をつければいいと思うが，特にこだわりがないのであれば上の順番でビルドしていくのがいいだろう。

## 前準備

最初に，上に挙げた以外でビルドに必要なパッケージを APT でインストールしておく。

```text
$ sudo apt install libgnutls28-dev bzip2 make gettext texinfo gnutls-bin build-essential libbz2-dev zlib1g-dev libncurses5-dev libsqlite3-dev libldap2-dev
```

適当にググって見繕ったものなので，もしかしたらこんなに要らないかもしれないが，精査するのが面倒くさかったので今回はこれで（笑）

ビルド用にダウンロードしたファイルの[完全性（integrity）をチェック](https://gnupg.org/download/integrity_check.html "GnuPG - Integrity Check")するために電子署名を検証する。
たとえばこんな感じ。

```text
gpg --verify gnupg-2.2.26.tar.bz2.sig gnupg-2.2.26.tar.bz2
```

署名チェックに必要な [OpenPGP] 公開鍵は以下のページにある。

- [GnuPG - Signature Key](https://gnupg.org/signature_key.html)

クリップボード操作ができる xsel または xclip コマンドがあるなら，上のページに貼り付けられている公開鍵（ASCII armor 形式）をコピって

```text
$ xsel | gpg --import
```

などとすれば公開鍵を取り込める。

## 各パッケージのビルド

ではパッケージを順にビルドしていこう。
ビルド用に適当なディレクトリをあらかじめ掘っておくとよい。

```text
$ sudo mkdir /var/local/gnupg-build
$ cd /var/local/gnupg-build
```

### [Libgpg-error](https://gnupg.org/software/libgpg-error/) のビルド

以下に手順だけ示しておく。

```text
$ sudo curl -L https://gnupg.org/ftp/gcrypt/libgpg-error/libgpg-error-1.41.tar.bz2 -O
$ sudo curl -L https://gnupg.org/ftp/gcrypt/libgpg-error/libgpg-error-1.41.tar.bz2.sig -O
$ gpg -d libgpg-error-1.41.tar.bz2.sig # integrity check
$ sudo tar xvf libgpg-error-1.41.tar.bz2
$ pushd libgpg-error-1.41/
$ ./configure 
$ make
$ sudo make install
$ popd
```

これで `/usr/local/` ディレクトリ以下の各ディレクトリにビルド結果がインストールされる。
`configure` コマンドは引数なしでも問題なさそうだ。

### [Libgcrypt](https://gnupg.org/software/libgcrypt/) のビルド

同様に手順を示しておく。

```text
$ sudo curl -L https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.8.7.tar.bz2 -O
$ sudo curl -L https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.8.7.tar.bz2.sig -O
$ gpg -d libgcrypt-1.8.7.tar.bz2.sig # integrity check
$ sudo tar xvf libgcrypt-1.8.7.tar.bz2
$ pushd libgcrypt-1.8.7/
$ ./configure 
$ make
$ sudo make install
$ popd
```

これで `/usr/local/` ディレクトリ以下の各ディレクトリにビルド結果がインストールされる。

### [Libassuan](https://gnupg.org/software/libassuan/) のビルド

同様に手順を示しておく。

```text
$ sudo curl -L https://gnupg.org/ftp/gcrypt/libassuan/libassuan-2.5.4.tar.bz2 -O
$ sudo curl -L https://gnupg.org/ftp/gcrypt/libassuan/libassuan-2.5.4.tar.bz2.sig -O
$ gpg -d libassuan-2.5.4.tar.bz2.sig # integrity check
$ sudo tar xvf libassuan-2.5.4.tar.bz2
$ pushd libassuan-2.5.4
$ ./configure 
$ make
$ sudo make install
$ popd
```

これで `/usr/local/` ディレクトリ以下の各ディレクトリにビルド結果がインストールされる。

### [Libksba](https://gnupg.org/software/libksba/) のビルド

同様に手順を示しておく。

```text
$ sudo curl -L https://gnupg.org/ftp/gcrypt/libksba/libksba-1.5.0.tar.bz2 -O
$ sudo curl -L https://gnupg.org/ftp/gcrypt/libksba/libksba-1.5.0.tar.bz2.sig -O
$ gpg -d libksba-1.5.0.tar.bz2.sig # integrity check
$ sudo tar xvf libksba-1.5.0.tar.bz2
$ pushd libksba-1.5.0
$ ./configure 
$ make
$ sudo make install
$ popd
```

これで `/usr/local/` ディレクトリ以下の各ディレクトリにビルド結果がインストールされる。

### [nPth](https://gnupg.org/software/npth/) のビルド

同様に手順を示しておく。

```text
$ sudo curl -L https://gnupg.org/ftp/gcrypt/npth/npth-1.6.tar.bz2 -O
$ sudo curl -L https://gnupg.org/ftp/gcrypt/npth/npth-1.6.tar.bz2.sig -O
$ gpg -d npth-1.6.tar.bz2.sig # integrity check
$ sudo tar xvf npth-1.6.tar.bz2
$ pushd npth-1.6
$ ./configure 
$ make
$ sudo make install
$ popd
```

これで `/usr/local/` ディレクトリ以下の各ディレクトリにビルド結果がインストールされる。

### [ntbTLS](https://gnupg.org/software/ntbtls/) のビルド

同様に手順を示しておく。

```text
$ sudo curl -L https://gnupg.org/ftp/gcrypt/ntbtls/ntbtls-0.2.0.tar.bz2 -O
$ sudo curl -L https://gnupg.org/ftp/gcrypt/ntbtls/ntbtls-0.2.0.tar.bz2.sig -O
$ gpg -d ntbtls-0.2.0.tar.bz2.sig # integrity check
$ sudo tar xvf ntbtls-0.2.0.tar.bz2
$ pushd ntbtls-0.2.0
$ ./configure 
$ make
$ sudo make install
$ popd
```

これで `/usr/local/` ディレクトリ以下の各ディレクトリにビルド結果がインストールされる。

### [GnuPG](https://gnupg.org/software/) のビルド

ようやく本命。
今までと同様に手順を示しておく。

```text
$ sudo curl -L https://gnupg.org/ftp/gcrypt/gnupg/gnupg-2.2.26.tar.bz2 -O
$ sudo curl -L https://gnupg.org/ftp/gcrypt/gnupg/gnupg-2.2.26.tar.bz2.sig -O
$ gpg -d gnupg-2.2.26.tar.bz2.sig # integrity check
$ sudo tar xvf gnupg-2.2.26.tar.bz2
$ pushd gnupg-2.2.26
$ ./configure 
$ make
$ sudo make install
$ popd
```

これで `/usr/local/` ディレクトリ以下の各ディレクトリにビルド結果がインストールされる。
そうそう，最後に

```text
$ sudo ldconfig
```

としておくのを忘れずに。
これで

```text
$ gpg --version
gpg (GnuPG) 2.2.26
libgcrypt 1.8.7
Copyright (C) 2020 Free Software Foundation, Inc.
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

と最新版で起動できる。
念のため設定状況もチェックしておこう。

```text
$ gpgconf --list-components
gpg:OpenPGP:/usr/local/bin/gpg
gpg-agent:プライベート鍵:/usr/local/bin/gpg-agent
scdaemon:スマートカード:/usr/local/libexec/scdaemon
gpgsm:S/MIME:/usr/local/bin/gpgsm
dirmngr:ネットワーク:/usr/local/bin/dirmngr
pinentry:パスフレーズ入力:/usr/local/bin/pinentry

$ gpgconf --list-dirs
sysconfdir:/usr/local/etc/gnupg
bindir:/usr/local/bin
libexecdir:/usr/local/libexec
libdir:/usr/local/lib/gnupg
datadir:/usr/local/share/gnupg
localedir:/usr/local/share/locale
socketdir:/run/user/1000/gnupg
dirmngr-socket:/run/user/1000/gnupg/S.dirmngr
agent-ssh-socket:/run/user/1000/gnupg/S.gpg-agent.ssh
agent-extra-socket:/run/user/1000/gnupg/S.gpg-agent.extra
agent-browser-socket:/run/user/1000/gnupg/S.gpg-agent.browser
agent-socket:/run/user/1000/gnupg/S.gpg-agent
homedir:/home/username/.gnupg
```

ふむむ。
上の設定に合わせて，もう少し弄っておこうか。

```text
$ cd /usr/local/bin/
$ sudo ln -s /etc/alternatives/pinentry pinentry
```

こんな感じかな。
Pinentry は下手に弄ると絶対にドツボにはまるので今回は既存のものを使う。

## ソケットががが【2021-01-10 変更】

この節の内容は「[gpg-agent の設定]{{< relref "./gpg-agent-in-ubuntu.md" >}}」に移動した。


<!--
上の `gpgconf --list-dirs` で示されるソケットについて

```text
socketdir:/run/user/1000/gnupg
dirmngr-socket:/run/user/1000/gnupg/S.dirmngr
agent-ssh-socket:/run/user/1000/gnupg/S.gpg-agent.ssh
agent-extra-socket:/run/user/1000/gnupg/S.gpg-agent.extra
agent-browser-socket:/run/user/1000/gnupg/S.gpg-agent.browser
agent-socket:/run/user/1000/gnupg/S.gpg-agent
```

そもそも `/run/user/1000/` ディレクトリ[^uid1] はブート時（？）に毎回初期化されるのだが，どうやら `/usr/bin/` ディレクトリ以下のオリジナル・コンポーネントを使って初期化してるみたいで，そのままだと `/usr/bin/gpg-agent` のほうが常駐してしまう。

[^uid1]: ちなみに `1000` はユーザIDを指す。ログインしている自身のユーザIDを知るには `id -u` で調べられる。

これでは面白くないし `gpg-agent` 経由でなにか操作をする度に `gpg-agent` のバージョンが古いと怒られるので，以下のコマンドでソケットを再作成する。

```text
$ gpgconf --kill all
$ gpgconf --remove-socketdir
$ gpgconf --create-socketdir
$ gpg-connect-agent --quiet /bye
```

これで `/usr/local/` ディレクトリ以下のコンポーネントに対応したソケットを再作成してくれる。
下３行はまとめて `~/.bashrc` 等に書いておくといいかもしれない。
いずれにせよ，かなりダサいやり方という自覚はありますよ。

これ，最初のソケット初期化時にやってくれないかなぁ。
[GnuPG] のサイトにはそもそもあまり情報がないし，ググっても情報が上手く見つけられなくて困っている。

どなたか情報をご存知でしたら教えてください {{< emoji "ペコン" >}}
-->

## 最後に動作確認

```text
$ echo hello world | gpg -a -s -u mykey
-----BEGIN PGP MESSAGE-----

owGbwMvMwCG45Zb1ujqFHTKMp4WSGOLfTbPOSM3JyVcozy/KSeHqmMzCIMjBYC2m
yCIdxHTby+542ccHayVgeliZQBpkZYoLMlPTU3P0MnP1oEyH9NzEzBy95PxcBi5O
AZj6RwcY/scanE3cwbVrv0v/zA7njR1J8d7vTjwP28EmLvzWVuPhqnaG/4kBz24f
634aYBi/NW79/1WcDusyco91C7N6zWwVC114PRsA
=/6XI
-----END PGP MESSAGE-----
```

署名時に Pinentry が起動してパスフレーズが通ればOK。
よーし，うむうむ，よーし。

## ブックマーク

- [Nitrokey/gnupg-docker: Build and use specific GnuPG version using Ubuntu image within Docker's container](https://github.com/Nitrokey/gnupg-docker) : Docker 環境で [GnuPG] をビルド&セットアップする

- [GnuPG チートシート（鍵作成から失効まで）]({{< ref "/openpgp/gnupg-cheat-sheet.md" >}})
- [Windows Terminal × NYAGOS × Scoop ＝ ♥]({{< ref "/remark/2020/10/windows-terminal-and-nyagos-and-scoop.md" >}}) : Windows 版 [GnuPG] は Scoop 経由でインストールするのがオススメ

[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
