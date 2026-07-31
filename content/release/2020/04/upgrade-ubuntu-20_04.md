+++
title = "Ubuntu 20.04 にアップグレードする"
date =  "2020-04-26T16:58:37+09:00"
description = "よーやく GnuPG/Libgcrypt のバージョンが上がったよ。"
image = "/images/attention/tools.png"
tags  = [ "ubuntu", "tools", "install", "openssl", "openssh", "gnupg" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] 20.04 がリリースされた。
20.04 は長期サポート版で2025年4月までサポートが有効である。

- [Ubuntu 20.04 LTS arrives | Ubuntu](https://ubuntu.com/blog/ubuntu-20-04-lts-arrives)
- [FocalFossa/ReleaseNotes - Ubuntu Wiki](https://wiki.ubuntu.com/FocalFossa/ReleaseNotes)

19.10 または長期サポート版の 18.04 から 20.04 へのアップグレードは[リリースノート](https://wiki.ubuntu.com/FocalFossa/ReleaseNotes "FocalFossa/ReleaseNotes - Ubuntu Wiki")を参考にするとよい。
基本的には

```text
$ update-manager -c -d
```

とアップグレード・モードで GUI を起動すれば，あとはよろしくやってくれる。
ただし 20.04 では32bitアーキテクチャをサポートしなくなったので，この場合は 18.04 からアップグレードしてはいけない[^u1]。

[^u1]: [Ubuntu] 18.04 のサポート期間は2023年4月まで。 18.04 サポート満了を以って32bit版 [Ubuntu] のサポートは終了する。ので，早めに64bitアーキテクチャへのリプレースを検討しませう。

なお，アップグレード時にサードパーティの APT リポジトリ（[git]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}} "PPA から Git をインストールする") や [ATOM](https://flight-manual.atom.io/getting-started/sections/installing-atom/ "Installing Atom") など）が外れるので，必要ならアップグレード時に戻して更新すること。

## 個別のアプリケーションについて

### [GnuPG] および [Libgcrypt]

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Libgcrypt]: https://gnupg.org/software/libgcrypt/

よーやく [GnuPG]/[Libgcrypt] のバージョンが上がったよ。

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

まぁ，現時点での[最新版は 2.2.20]({{< ref "/release/2020/03/gnupg-2_2_20-is-released.md" >}} "GnuPG 2.2.20 がリリースされた") だけどね。
ちなみに [GnuPG] 2.2.17 から SHA-1 電子署名の扱いが変わったのでご注意を。

- [GnuPG 2.2.17 リリース： 公開鍵サーバ・アクセスに関する過激な変更あり]({{< ref "/release/2019/07/gnupg-2_2_17-is-released.md" >}})
- [GnuPG 2.2.18 リリース： さようなら SHA-1]({{< ref "/release/2019/11/gnupg-2_2_18-is-released.md" >}})

### [OpenSSH] および [OpenSSL]

[OpenSSH]: https://www.openssh.com/
[OpenSSL]: https://www.openssl.org/

アップグレード直後の [OpenSSH] および [OpenSSL] のバージョンは以下のとおりだった。

```text
$ ssh -V
OpenSSH_8.2p1 Ubuntu-4, OpenSSL 1.1.1f  31 Mar 2020
```

[OpenSSH] は [2020-02-14 にリリースされた最新版](https://www.openssh.com/txt/release-8.2)が入っているようだ。
ただ [OpenSSL] はこの前セキュリティ・アップデートがあったばかりなんだよねぇ。

- [OpenSSL の脆弱性対策について(CVE-2020-1967) ：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/alert20200423.html)
- [OpenSSL の脆弱性 (CVE-2020-1967) に関する注意喚起](https://www.jpcert.or.jp/at/2020/at200018.html)

まぁ，間に合わなかったんだろうけど。
[CVSSv3 のスコアが 7.5](https://access.redhat.com/security/cve/CVE-2020-1967) で深刻度が高いので早めの対応をお願いしたい。

### 保留されているパッケージ

アップグレード後に `libc++1` が保留状態で残っていた。
保留状態のパッケージがある場合は

```text
$ sudo apt full-upgrade
```

で保留分をまとめてアップグレードするか，個別に

```text
$ sudo apt install libc++1
```

などとすればいいようだ。

### [GCC] が消えとるがな

[GCC]: https://gcc.gnu.org/ "GCC, the GNU Compiler Collection - GNU Project - Free Software Foundation (FSF)"
[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

アップグレードしたら何故か [GCC] が削除されてた `orz`

しょうがないので

```text
$ sudo apt install build-essential
```

したですよ。
[Rust] のコンパイル・リンクも問題なく動くみたいだし，いっかな[^rust1]。

[^rust1]: [Rust] も APT で導入できるようだが，バージョンが古いので，最新版が必要なら [APT は使わない]({{< ref "/rust-lang/hello.md" >}} "みんな大好き Hello World")ほうがよい。

ちなみに [GCC] は 9.3 が導入される。

### [ATOM] も削除される

[ATOM]: https://atom.io/

サードパーティ・リポジトリから [ATOM] をインストールしている場合，アップグレード時に APT のリストからリポジトリが外されるだけだが，今回の 20.04 へのアップグレードでは [ATOM] 自体がまるっと削除される。

ただし設定やパッケージはそのまま残っているので，リポジトリ設定を戻して `apt install` し直せばおっけ。

- [Installing Atom](https://flight-manual.atom.io/getting-started/sections/installing-atom/)

### [Lollypop]

[Lollypop]: https://wiki.gnome.org/Apps/Lollypop?action=show "Apps/Lollypop - GNOME Wiki!"

どうやら [Lollypop] は標準の APT リポジトリに入ったらしい。
いつからだ？

```text
$ apt show lollypop
Package: lollypop
Version: 1.2.35-1
Priority: optional
Section: universe/gnome
Origin: Ubuntu
...
```

ちうわけで，もう [PPA リポジトリ](https://launchpad.net/~gnumdk/+archive/ubuntu/lollypop "Lollypop : Cédric Bellegarde")を使わなくともよさそうである，多分。

### 【追記 2020-04-28】 Libsecret のアップデート

[git]: https://git-scm.com/
[Git]: https://git-scm.com/

[Git] credential helper である GNOME/libsecret もバージョンが上がっているようだ。

```text
$ apt show libsecret-1-dev
Package: libsecret-1-dev
Version: 0.20.2-1
Priority: optional
Section: libdevel
Source: libsecret
Origin: Ubuntu
...
```

GNOME/libsecret は `apt upgrade` しただけではダメで，手動でビルドする必要がある。

```text
$ mkdir ~/work
$ cd ~/work
$ cp -r /usr/share/doc/git/contrib/credential/libsecret .
$ cd libsecret
$ make
gcc -g -O2 -Wall  -pthread -I/usr/include/libsecret-1 -I/usr/include/libmount -I/usr/include/blkid -I/usr/include/glib-2.0 -I/usr/lib/x86_64-linux-gnu/glib-2.0/include -o git-credential-libsecret.o -c git-credential-libsecret.c
gcc -o git-credential-libsecret  git-credential-libsecret.o -lsecret-1 -lgio-2.0 -lgobject-2.0 -lglib-2.0
```

これで生成した `git-credential-libsecret` を `$PATH` の通ったディレクトリに放り込んでおけばよい。

- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})

## 【2020-05-15 追記】 日本語 Remix のリリースと Bootable USB メモリの作成

Japanese Team による [Ubuntu] 20.04 LTS 日本語 Remix がリリースされた。
感謝！

- [Ubuntu 20.04 LTS 日本語 Remix リリース | Ubuntu Japanese Team](https://www.ubuntulinux.jp/News/ubuntu2004-ja-remix)

Japanese Team による追加パッケージのリポジトリを導入するには以下の手順を実行する。

```text
$ wget -q https://www.ubuntulinux.jp/ubuntu-ja-archive-keyring.gpg -O- | sudo apt-key add -
$ wget -q https://www.ubuntulinux.jp/ubuntu-jp-ppa-keyring.gpg -O- | sudo apt-key add -
$ sudo wget https://www.ubuntulinux.jp/sources.list.d/focal.list -O /etc/apt/sources.list.d/ubuntu-ja.list
$ sudo apt update
```

また ISO イメージ・ファイルも公開されているのでこれを使って bootable USB メモリを作成しておく。
詳しくは以下の拙文を参考にどうぞ。

- [Ubuntu インストール用のブート可能 USB メモリを作成する]({{< ref "/remark/2019/04/bootable-usb.md" >}})

## ブックマーク

- [Ubuntu 20.04 その23 - Ubuntu 20.04 LTSがリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2020/04/ubuntu-2004-23-ubuntu-2004-lts.html)
- [Ubuntu 20.04 その24 - Ubuntu 20.04 LTSの新機能と変更点 - kledgeb](https://kledgeb.blogspot.com/2020/04/ubuntu-2004-24-ubuntu-2004-lts.html)
- [Ubuntu 20.04 その25 - Ubuntu 20.04 LTSの既知の問題 - kledgeb](https://kledgeb.blogspot.com/2020/04/ubuntu-2004-25-ubuntu-2004-lts.html)
- [Ubuntu 20.04 その26 - Ubuntu 20.04.1 LTSのリリーススケジュール・Ubuntu 18.04 LTSユーザーにアップグレードパスの提供 - kledgeb](https://kledgeb.blogspot.com/2020/04/ubuntu-2004-26-ubuntu-20041-ltsubuntu.html)
- [Ubuntu 20.04 その27 - Qt 5.14.2の採用見送りとQt 5.12.8の採用 - kledgeb](https://kledgeb.blogspot.com/2020/04/ubuntu-2004-27-qt-5142qt-5128.html)
- [Ubuntu 20.04 その28 - Ubuntu Desktopの新機能と魅力・様々な新機能と改良点の紹介 - kledgeb](https://kledgeb.blogspot.com/2020/04/ubuntu-2004-28-ubuntu-desktop.html)
- [Ubuntu 20.04 その29 - Linux kernel 5.4の新機能 - kledgeb](https://kledgeb.blogspot.com/2020/04/ubuntu-2004-29-linux-kernel-54.html)
- [Ubuntu 20.04 LTSインストールガイド【スクリーンショットつき解説】 | LFI](https://linuxfan.info/ubuntu-20-04-install-guide)
- [(今はまだ)WSL1にUbuntu 20.04を入れるな - Qiita](https://qiita.com/mmns/items/eaf42dd3345a2285ff9e)
- [Ubuntu 20.04 その48 - Ubuntu 20.04 LTS 日本語 Remixがリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2020/05/ubuntu-2004-48-ubuntu-2004-lts-remix.html)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"

## 参考

{{% review-paapi "B01NBU1OS5" %}} <!-- シリコンパワー USBメモリ 32GB USB3.1 -->
