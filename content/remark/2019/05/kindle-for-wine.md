+++
title = "Ubuntu でも Kindle 本が読みたい"
date =  "2019-05-28T21:41:58+09:00"
description = "Linux プラットフォームで動くネイティブな Kindle アプリケーションはないようで Wine 上で Windows 用の Kindle for PC を起動するしかないようだ。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "tools", "ubuntu", "kindle", "wine", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

まぁマンガやラノベを読む程度なら Kindle 端末でも十分だし一部のマンガ等は [Kindle Cloud Reader](https://read.amazon.co.jp/) を利用する手もあるのだが，やはり [Ubuntu] のデスクトップ上でも読みたいのである。

ただ Linux プラットフォームで動くネイティブな Kindle アプリケーションはないようで，今のところ [Wine] 上で Windows 用の Kindle for PC を起動するしかない。

まぁ [Wine] はそのうち試してみる予定だったので，このさい試してみるとしよう。

## [Wine] のインストール

[Wine] は APT で導入可能である。

```text
$ apt show wine
Package: wine
Version: 4.0-1
Built-Using: khronos-api (= 4.6+git20180514-1), unicode-data (= 11.0.0-1)
Priority: optional
Section: universe/otherosfs
Origin: Ubuntu
Maintainer: Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
Original-Maintainer: Debian Wine Party <debian-wine@lists.debian.org>
...
```

ちなみに `wine-stable` だと v3 系， `wine-development` だと v4.2 がインストールされるようだ（2019-05 時点）。
ここは素直に `wine` をインストールしておこう。

```text
$ sudo apt install wine
```

一応，動作確認しておく。

```text
$ wine --version
wine-4.0 (Ubuntu 4.0-1)
```

よしよし。

### 日本語フォントのインストール

当然ながら素の [Wine] には Windows アプリケーション用の日本語フォントが入っていない。
日本語フォントを入れるに [Winetricks] をインストールする。

```text
$ sudo apt install winetricks
```

[Ubuntu] のメニューに [Winetricks] のアイコンが追加されるのでクリックして起動する。
すると以下の画面が起動する。

{{< fig-img src="./winetricks.png" link="./winetricks.png" width="1010" >}}

ここで “Select the default wineprefix” を選択して *[OK]* を押すと以下の画面に遷移する。

{{< fig-img src="./winetricks-install-font.png" link="./winetricks-install-font.png" width="1010" >}}

更にここで “Install a font” を選択して *[OK]* を押してフォントの選択画面に遷移する。

{{< fig-img src="./winetricks-install-cjkfont.png" link="./winetricks-install-cjkfont.png" width="1010" >}}

“cjkfonts” にチェックを付けて *[OK]* を押せばフォントのインストールが開始される。

[Winetricks] ではワーニングやエラーのプロンプトが出まくるが全部無視しても問題ない（なんだかなぁ）。

## Kindle for PC のインストール

あとは Kindle for PC のインストーラを取ってきて [Wine] 上で起動すればいいのだが，どうやら [Amazon にある最新版の Kindle for PC](https://www.amazon.co.jp/exec/obidos/ASIN/B011UEHYWQ/baldandersinf-22/) はうまく動かないっぽい。

しょうがないので v1.17 の Kindle for PC のインストーラを使う。

- [ダウンロード kindle for pc 無料](https://kindle-for-pc.jp.uptodown.com/windows/download)

こんなん使って大丈夫なのか不安になるが，どの解説ページでもこれを使えって言ってくるんだよなぁ。
頻繁に使うもんじゃないし，いっかぁ（投げやり）。

インストーラの起動は以下の通り[^w1]。

[^w1]: 古いインストーラを使うせいかもしれないが [Wine] 内の Windows バージョンを “Windows 10” にするとインストーラがうまく動かないらしい。 [Wine] 内の Windows バージョンを設定するには `winecfg` コマンドで設定ダイアログを起動する。

```text
$ wine kindle-for-pc-1-17-44183.exe
```

これでインストールが行われて Kindle for PC が起動する。
[Ubuntu] のメニューに Kindle アイコンが登録されるので，他のアプリケーションと同じようにアイコンのクリックでいつでも起動できる。

### [Wine] 内のアプリケーションの削除

[Wine] 内のアプリケーションを削除するには以下のコマンドを起動する。

```text
$ wine uninstaller &
```

{{< fig-img src="./wine-uninstaller.png" link="./wine-uninstaller.png" >}}

## ブックマーク

- [Ubuntu - WineHQ Wiki](https://wiki.winehq.org/Ubuntu)
- [WineHQ  - Amazon Kindle for PC](https://appdb.winehq.org/objectManager.php?sClass=application&iId=10597)
- [Ubuntu18.04にKindle for PCをインストールした - Qiita](https://qiita.com/sakai39e/items/75b2c95bc4c3cab13849)
- [Ubuntu18.04.1でWineをインストールして、Kindleを使う方法 | Awesome Blog](https://awesome-linus.com/2019/04/11/ubuntu18-wine-install-kindle/)
- [Ubuntu1604LTSでKindleを使う(PlayOnLinux) - Qiita](https://qiita.com/giwagiwa/items/d2e447af5225c1ce9800) : 私の環境ではなぜか PlayOnLinux がうまく動かなくて諦めた
- [How to uninstall WINE applications](https://www.dedoimedo.com/computers/wine-uninstall-apps.html)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Wine]: https://www.winehq.org/ "WineHQ - Run Windows applications on Linux, BSD, Solaris and macOS"
[Winetricks]: https://github.com/Winetricks

{{% review-paapi "B07HCSL6BN" %}} <!-- Kindle Paperwhite -->
