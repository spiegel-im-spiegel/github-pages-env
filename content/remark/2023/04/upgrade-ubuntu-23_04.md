+++
title = "Ubuntu 23.04 へのアップグレード"
date =  "2023-04-29T14:55:14+09:00"
description = "もはや Ubuntu は GnuPG をメンテナンスする気がないと判断する。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] 23.04 がリリースされた。
約半年の短期サポート版。

- [Canonical releases Ubuntu 23.04 Lunar Lobster | Canonical](https://canonical.com/blog/canonical-releases-ubuntu-23-04-lunar-lobster)
- [Ubuntu Fridge | Ubuntu 23.04 (Lunar Lobster) released](https://ubuntu-news.org/2023/04/20/ubuntu-23-04-lunar-lobster-released/)
- [LunarLobster/ReleaseNotes - Ubuntu Wiki](https://wiki.ubuntu.com/LunarLobster/ReleaseNotes)
- [Ubuntu 23.04 その24 - Ubuntu 23.04の新機能と変更点・既知の問題 - kledgeb](https://kledgeb.blogspot.com/2023/04/ubuntu-2304-24-ubuntu-2304.html)

「ソフトウェアの更新」を開くと（最新パッケージへの更新の後） `[アップグレード...]` ボタンが表示される。

{{< fig-img src="./check-upgrade.png" title="ソフトウェアの更新" link="./check-upgrade.png" width="645" >}}

このボタンを押下するとアップグレードが始まる[^cmd1]。

[^cmd1]: コマンドラインでアップグレードする場合は `sudo apt dist-upgrade` でアップグレードが開始される。

{{< fig-img src="./upgrade-lunar.png" title="アップグレードの開始" link="./upgrade-lunar.png" width="620" >}}

あとはなりゆきでOK。
今回も特に大きな問題はなかった。

## いつものように [GnuPG] が古い（怒）

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

ちなみに 2023-04-28 に [GnuPG] v2.4.1 (安定版) が出ている。
v2.3 系までは我慢してたが，もはや [Ubuntu] は [GnuPG] をメンテナンスする気がないと判断する。

というわけで，近いうちに自前で v2.4 系にアップデートする予定。

- [Installing GnuPG 2.4 on Ubuntu 22.04 | Pro Custodibus](https://www.procustodibus.com/blog/2023/02/gpg-2-4-on-ubuntu-22-04/)

## [OpenSSH] と [OpenSSL]

```text
$ ssh -V
OpenSSH_9.0p1 Ubuntu-1ubuntu8, OpenSSL 3.0.8 7 Feb 2023
```

これも微妙に古いなぁ。
ちなみに [OpenSSL] は 2023-03-14 に v3.1.0 が出ているが，流石に間に合わなかったようだ。
[OpenSSH] は 9.x でいくつか脆弱性の修正があったはずだが，ちゃんとバックポート・パッチを当てているのだろうか。
大丈夫？ 多分？

## 今後の TODO

- [ ] Lunar 版の日本語 Remix が出たらインストールメディアを作成する
- [ ] [pgAdmin4] には Lunar 版のリポジトリがない。出たら対応する
- [ ] [NodeSource](https://github.com/nodesource) の LTS 版 v18.x を利用しているのだが，まだ Lunar 版のリポジトリがない。出たら対応する
  - [NodeSource Node.js Binary Distributions](https://github.com/nodesource/distributions/blob/master/README.md)
  - [Ubuntu 20.04にNode.jsをインストールする方法  | DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-20-04-ja)

## ブックマーク

- [Ubuntu 23.04 その23 - Ubuntu 23.04がリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2023/04/ubuntu-2304-23-ubuntu-2304.html)
- [Ubuntu 22.10 その84 - Ubuntu 23.04へアップグレードするには・アップグレードの注意事項 - kledgeb](https://kledgeb.blogspot.com/2023/04/ubuntu-2210-84-ubuntu-2304.html)

- [apt-key が非推奨になったので](https://zenn.dev/spiegel/articles/20220508-apt-key-is-deprecated)

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

## 参考図書

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B0BG8J5QJ1" %}} <!-- ［試して理解］Linuxのしくみ 増補改訂版 -->
{{% review-paapi "B01NBU1OS5" %}} <!-- シリコンパワー USBメモリ 32GB USB3.1 -->
