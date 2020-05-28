+++
title = "Ubuntu に KeePassXC を導入する"
date =  "2019-08-25T12:05:20+09:00"
description = "Snap 版は日本語入力ができないため PPA リポジトリを使って最新版をインストールする。"
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

**【2020-05-28 修正】**
[Ubuntu] では APT または [Snap] から導入できるのだが， [Snap] 版は日本語入力ができない（日本語の文字列のコピペは可能）という致命的な欠点があるため [PPA] リポジトリを使って最新版をインストールする。

まずは [PPA] リポジトリの登録から。

```text
$ sudo add-apt-repository ppa:phoerious/keepassxc
$ sudo apt update
```

これで [KeePassXC] がインストール可能になった。

```text
$ apt show keepassxc
Package: keepassxc
Version: 2.5.4-1ppa1~focal1
Priority: optional
Section: utils
Maintainer: KeePassXC Team <team@keepassxc.org>
...
```

あとは普通に `apt install` すればよい。


```text
$ sudo apt install keepassxc
```

これで [Ubuntu] のメニューから [KeePassXC] を起動できる。

{{< fig-img src="./keepassxc.png" link="./keepassxc.png" width="800" >}}

おおっ，日本語だ（笑）

使い方は本家の [KeePass] とだいたい同じようだ。
ブラウザ連携は今のところ怖くて使う気にならない。
まぁ，そのうち気が向いたらね。

## ブックマーク

- [GitHub - keepassxreboot/keepassxc: KeePassXC is a cross-platform community-driven port of the Windows application “Keepass Password Safe”.](https://github.com/keepassxreboot/keepassxc)
- [Ubuntu 19.04 その146 - Snap Storeからインストール可能なセキュリティーアプリ5選 - kledgeb](https://kledgeb.blogspot.com/2019/08/ubuntu-1904-146-snap-store5.html)
- [Snapから使いやすいセキュリティソフトウェア5選 | マイナビニュース](https://news.mynavi.jp/article/20190825-883052/)

[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[KeePass]: https://keepass.info/ "KeePass Password Safe"
[Snap]: https://github.com/snapcore/snapd "snapcore/snapd: The snapd and snap tools enable systems to work with .snap files."
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
