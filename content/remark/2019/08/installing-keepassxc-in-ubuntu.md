+++
title = "Ubuntu に KeePassXC を導入する"
date =  "2019-08-25T12:05:20+09:00"
description = "APT または Snap から導入できるのだが Snap 版を利用したほうがいいみたい。 "
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "password", "management", "tools", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

以前に [Mono 版の KeePass を導入]({{< ref "/remark/2019/04/mono-in-ubuntu.md" >}} "Ubuntu に Mono を導入する")する手順を紹介したが Linux 用であればサードパーティの [KeePassXC] がよさげである。

[KeePassXC] は [KeePass] のサードパーティ版である KeePassX の fork でマルチプラットフォームに対応している。
Windows や macOS の他，各種 Linux ディストリビューションにも対応しているようだ。

[Ubuntu] では APT または [Snap] から導入できるのだが[^snp1]

[^snp1]: [Ubuntu] 18.04 以降は標準で [Snap] が入ってるらしい。

```text
$ apt show keepassxc
Package: keepassxc
Version: 2.3.4+dfsg.1-1
Priority: optional
Section: universe/utils
Origin: Ubuntu
Maintainer: Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
Original-Maintainer: Julian Andres Klode <jak@debian.org>
...
```

```text
$ snap find keepassxc
Name             Version  Publisher       Notes  Summary
keepassxc        2.4.3    keepassxreboot  -      community driven port of the windows application “Keepass Password Safe”
keepassx-elopio  2.0.2    elopio          -      KeePassX is a cross platform password safe
```

ということで [Snap] 版を利用したほうがいいみたい。
というわけで，早速！

```
$ sudo snap install keepassxc
keepassxc 2.4.3 from Jonathan White (keepassxreboot) installed
```

これで [Ubuntu] のメニューから [KeePassXC] を起動できる。

{{< fig-img src="./keepassxc.png" link="./keepassxc.png" width="800" >}}

おおっ，日本語だ（笑）

使い方は本家の [KeePass] とだいたい同じようだ。
ブラウザ連携は今のところ怖くて使う気にならない。
まぁ，そのうち気が向いたらね。

削除する場合は

```
$ sudo snap remove keepassxc
```

でおｋ。
ちなみにアップグレードは

```
$ sudo snap refresh
```

で [Snap] 管理下にある全てのパッケージを一括でアップグレードできる。

[Snap] はローカルにリポジトリ・データベースを持たないので（APT に比べて）扱いが簡単なのはいいのだが，こういうアプリ・ストア型のパッケージ管理システムは，どうしてもサービス提供者側の統制が強くなるので，好かんのだよなぁ。
一応 [PPA] にも[リポジトリがあるみたい](https://launchpad.net/~phoerious/+archive/ubuntu/keepassxc "KeePassXC : Janek Bevendorff")なのだが，うーん...

## ブックマーク

- [GitHub - keepassxreboot/keepassxc: KeePassXC is a cross-platform community-driven port of the Windows application “Keepass Password Safe”.](https://github.com/keepassxreboot/keepassxc)
- [Ubuntu 19.04 その146 - Snap Storeからインストール可能なセキュリティーアプリ5選 - kledgeb](https://kledgeb.blogspot.com/2019/08/ubuntu-1904-146-snap-store5.html)
- [Snapから使いやすいセキュリティソフトウェア5選 | マイナビニュース](https://news.mynavi.jp/article/20190825-883052/)

[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[KeePass]: https://keepass.info/ "KeePass Password Safe"
[Snap]: https://github.com/snapcore/snapd "snapcore/snapd: The snapd and snap tools enable systems to work with .snap files."
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
