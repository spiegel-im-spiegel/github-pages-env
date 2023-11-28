+++
title = "Ubuntu 23.10 へのアップグレード"
date =  "2023-10-31T18:25:10+09:00"
description = "GUI 版のアップグレードは問題があるらしい。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いつまで経っても自宅機に Ubuntu 23.10 へのアップグレード通知が来ないのですっかり忘れていたが， 2023-10-12 には出てたんだな。

- [Ubuntu Fridge | Ubuntu 23.10 (Mantic Minotaur) released](https://ubuntu-news.org/2023/10/16/ubuntu-23-10-mantic-minotaur-released/)
- [Canonical releases Ubuntu 23.10 Mantic Minotaur](https://canonical.com/blog/canonical-releases-ubuntu-23-10-mantic-minotaur)

どうやらアップグレードする際に画面の表示が文字化けする問題があるみたい。

- [Ubuntu 23.04 その94 - まだUbuntu 23.10へアップグレードできない理由 - kledgeb](https://kledgeb.blogspot.com/2023/10/ubuntu-2304-94-ubuntu-2310.html)

ただし Terminal から直接コマンド叩けば行けるみたいなので，これでアップグレードを行った。

```text
$ sudo do-release-upgrade -d
```

OS のアップグレードを行う前に APT と Snap のパッケージを最新にアップデートしておいてね。

```text
$ sudo apt update
$ sudo apt upgrade
$ sudo snap refresh
```

まぁ，これでアップグレード自体は問題なく完了した。

## いつものように [GnuPG] が古い

```text
$ gpg --version
gpg (GnuPG) 2.2.40
libgcrypt 1.10.1
Copyright (C) 2022 g10 Code GmbH
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

ちなみに 2023-07-04 に [GnuPG] v2.4.3 (安定版) が出ている。
もはや [Ubuntu] は [GnuPG] をメンテナンスする気がないと判断しているが，自前でのアップデートはまだやってない。
そのうち... と思いつつズルズルと。

## [OpenSSH] と [OpenSSL]

```text
$ ssh -V
OpenSSH_9.3p1 Ubuntu-1ubuntu3, OpenSSL 3.0.10 1 Aug 2023
```

これも微妙に古いなぁ。
ちなみに [OpenSSL] は 2023-10-24 に v3.0.12 と v3.1.4 が出ている。
バックポートとかの状況は知らない。
v3.2 はベータ版が出ている。
ちゃんと追従してくれてるのかねぇ...

## 今後の TODO

とりあえず予定だけ立てておく。

- [x] 日本語 Remix のインストールメディアを作成する
  - [Ubuntu 23.10 日本語 Remix リリース | Ubuntu Japanese Team](https://www.ubuntulinux.jp/News/ubuntu2310-ja-remix) : 「公式版では新しいインストーラーが採用されていますが、日本語Remixは従来のインストーラーを利用したものとなっています。これは、[新しいインストーラーには、どの言語を選択してもライブセッションが英語になるという問題](https://bugs.launchpad.net/ubuntu-release-notes/+bug/2013329)が確認されているためです」とのこと
- [x] Git PPA 版
- [x] [KeePassXC] PPA 版
- [x] [NodeSource](https://github.com/nodesource)
  - [NodeSource Node.js Binary Distributions](https://github.com/nodesource/distributions/blob/master/README.md)
    - nodistro を指定することでメンテナンスフリーになった。めでたし！
  - [Ubuntu 20.04にNode.jsをインストールする方法  | DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-20-04-ja)
- [x] [pgAdmin4]
  - [Download - pgAdmin 4 (APT)](https://www.pgadmin.org/download/pgadmin-4-apt/)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[pgAdmin4]: https://www.pgadmin.org/ "pgAdmin - PostgreSQL Tools"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Node.js]: https://nodejs.org/
[OpenSSL]: https://www.openssl.org/
[OpenSSH]: https://www.openssh.com/
[Docker]: https://www.docker.com/ "Empowering App Development for Developers | Docker"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"

## ブックマーク

- [Ubuntu 23.10 その44 - Ubuntu 23.10がリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2023/10/ubuntu-2310-44-ubuntu-2310.html)
- [Ubuntu 23.10 その45 - Ubuntu 23.10の新機能と変更点・既知の問題 - kledgeb](https://kledgeb.blogspot.com/2023/10/ubuntu-2310-45-ubuntu-2310.html)
- [Ubuntu 23.10 その46 - Ubuntu 23.10の注目機能 - kledgeb](https://kledgeb.blogspot.com/2023/10/ubuntu-2310-46-ubuntu-2310.html)

## 参考図書

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B0BG8J5QJ1" %}} <!-- ［試して理解］Linuxのしくみ 増補改訂版 -->
{{% review-paapi "B01NBU1OS5" %}} <!-- シリコンパワー USBメモリ 32GB USB3.1 -->
