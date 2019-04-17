+++
title = "Ubuntu に ATOM エディタを導入する"
date = "2019-04-15T23:32:39+09:00"
description = "この記事は随時更新します。 悪しからずご了承の程を。"
image = "/images/attention/kitten.jpg"
tags = ["atom", "editor", "ubuntu", "linux", "font"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

さて，そろそろ [ATOM] を [Ubuntu] に入れようか。

なお，この記事は随時更新します。
悪しからずご了承の程を。

## 前準備

[ATOM] の[リポジトリ](https://github.com/atom/atom/ "atom/atom: The hackable text editor")に各バージョンの `atom-amd64.deb` ファイルが公開されているので，これを使う。

deb ファイルを使ってインストールを行うには `gdebi` というコマンドを使うらしい。
これは ATP で取得可能だ。

```text
$ sudo apt show gdebi-core
Package: gdebi-core
Version: 0.9.5.7+nmu2
Priority: optional
Section: universe/admin
Source: gdebi
Origin: Ubuntu
Maintainer: Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
Original-Maintainer: Ubuntu Developers <ubuntu-dev-team@lists.alioth.debian.org>
Bugs: https://bugs.launchpad.net/ubuntu/+filebug
Installed-Size: 876 kB
Depends: python3:any (>= 3.3.2-2~), python3-apt, python3-debian, file
Suggests: xz-utils | xz-lzma
Task: ubuntu-mate-core, ubuntu-mate-desktop
Download-Size: 116 kB
APT-Sources: http://jp.archive.ubuntu.com/ubuntu cosmic/universe amd64 Packages
Description: simple tool to install deb files
 gdebi を使うと、ローカルの deb パッケージをその依存関係を解決しながらインス トールできます。apt
 は同じ動作をしますが、インストールできるのはリモート (http, ftp) にあるパッケージのみです。
 .
 It can also resolve build-depends of local debian/control files.
 .
 This package contains the libraries and command-line utility.
```

ふむむ。
なるほど。

ではインストールしてしまおう。

```text
$ sudo apt install gdebi-core
```

インストールできたら動作確認しておく。

```text
$ gdebi --version
0.9.5.7+nmu2
```

よしよし。

## [ATOM] の導入

[リリースページ](https://github.com/atom/atom/releases "Releases · atom/atom") から最新版の `atom-amd64.deb` ファイルをダウンロードし `gdebi` コマンドでインストールする。

```text
$ sudo gdebi atom-amd64.deb
Reading package lists... Done
Building dependency tree        
Reading state information... Done
Reading state information... Done
以下のパッケージのインストールが必要です: gconf-service gconf-service-backend gconf2 gconf2-common libgconf-2-4 libpython-stdlib libpython2-stdlib python python-minimal python2 python2-minimal python2.7 python2.7-minimal 

A hackable text editor for the 21st Century.
 Atom is a free and open source text editor that is modern, approachable, and hackable to the core.
ソフトウェアパッケージをインストールしますか? [y/N]:y

[...snip...]

atom-amd64.deb を展開する準備をしています ...
atom (1.36.0) を展開しています...
atom (1.36.0) を設定しています ...
desktop-file-utils (0.23-3ubuntu3) のトリガを処理しています ...
gnome-menus (3.13.3-11ubuntu2) のトリガを処理しています ...
mime-support (3.60ubuntu1) のトリガを処理しています ...
```

インストールが完了するとドックのアプリボタンで表示されるアプリ一覧にアイコンが追加される。
もちろんターミナルからも起動できる。
1.36 から単一のファイルのみを指定して起動することが可能になった。

```text
$ atom ~/work/foo.txt 
```

{{< fig-img src="./open-single-file-by-atom.png" link="./open-single-file-by-atom.png" width="790" >}}

善き哉。

## [ATOM] 用にフォントを導入する。

個人的にテキストエディタのフォントは [Inconsolata]，そして日本語は明朝体に限る。
日本語は標準の NOTO フォントを使うとして [Inconsolata] はネットから取ってこないと。

[Ubuntu] ではフォントの置き場所は以下の3箇所になるようだ。

- `/usr/share/fonts` : システムフォント
- `/usr/local/share/fonts` : 追加フォントをマシンで共有する場合
- `~/.fonts` : 個人で導入する場合

これらのディレクトリのいずれかに入れたらキャッシュを更新しておく。

```text
$ fc-cache -fv
```

ちゃんと導入できたか確認しておこう。

```text
$ fc-list | grep Inconsolata
/home/username/.fonts/Inconsolata.otf: Inconsolata:style=Medium
```

よーし，うむうむ，よーし。

[ATOM] でフォントを指定する場合には Settings → Editor Settings の Font Family の項目で

```
Inconsolata, Noto Serif CJK JP
```

などと指定しておけばよい。
ちなみに NOTO フォントは以下のように配置されている。

```text
$ fc-list | grep Noto | grep JP
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans CJK JP,Noto Sans CJK JP Bold:style=Bold,Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-Bold.ttc: Noto Serif CJK JP:style=Bold
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans Mono CJK JP,Noto Sans Mono CJK JP Bold:style=Bold,Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Medium.ttc: Noto Sans CJK JP,Noto Sans CJK JP Medium:style=Medium,Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-Black.ttc: Noto Serif CJK JP,Noto Serif CJK JP Black:style=Black,Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans CJK JP,Noto Sans CJK JP Regular:style=Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-Light.ttc: Noto Serif CJK JP,Noto Serif CJK JP Light:style=Light,Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-Regular.ttc: Noto Serif CJK JP:style=Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans Mono CJK JP,Noto Sans Mono CJK JP Regular:style=Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-DemiLight.ttc: Noto Sans CJK JP,Noto Sans CJK JP DemiLight:style=DemiLight,Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Thin.ttc: Noto Sans CJK JP,Noto Sans CJK JP Thin:style=Thin,Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Light.ttc: Noto Sans CJK JP,Noto Sans CJK JP Light:style=Light,Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-SemiBold.ttc: Noto Serif CJK JP,Noto Serif CJK JP SemiBold:style=SemiBold,Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Black.ttc: Noto Sans CJK JP,Noto Sans CJK JP Black:style=Black,Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-ExtraLight.ttc: Noto Serif CJK JP,Noto Serif CJK JP ExtraLight:style=ExtraLight,Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-Medium.ttc: Noto Serif CJK JP,Noto Serif CJK JP Medium:style=Medium,Regular
```

## ブックマーク

- [Install Atom on Ubuntu 18.04 Bionic Beaver Linux - LinuxConfig.org](https://linuxconfig.org/install-atom-on-ubuntu-18-04-bionic-beaver-linux)
- [ubuntu18.04にAtomをインストール – v0.0.1](https://www.leo-leo.uno/2018/08/18/392/)
- [Ubuntu 18.04 LTSの日本語フォントを変更する！【詳細解説】 | LFI](https://linuxfan.info/ubuntu-18-04-change-ja-font)
- [UbuntuTips/Desktop/InstallFont - Ubuntu Japanese Wiki](https://wiki.ubuntulinux.jp/UbuntuTips/Desktop/InstallFont)
- [CentOS/Ubuntuでフォントを追加する方法 | 俺的備忘録 〜なんかいろいろ〜](https://orebibou.com/2017/01/centosubuntu%E3%81%A7%E3%83%95%E3%82%A9%E3%83%B3%E3%83%88%E3%82%92%E8%BF%BD%E5%8A%A0%E3%81%99%E3%82%8B%E6%96%B9%E6%B3%95/)
- [GitHub、フリーのコードエディター「Atom 1.36」を正式公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1180090.html)
- [Items in the Menu Bar are invisible on Ubuntu 18.10 when atom window is focused. · Issue #18535 · atom/atom · GitHub](https://github.com/atom/atom/issues/18535)

- [ATOM Editor に関するメモ]({{< ref "/remark/2015/atom-editor.md" >}})

[ATOM]: https://atom.io/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Inconsolata]: https://www.levien.com/type/myfonts/inconsolata.html
