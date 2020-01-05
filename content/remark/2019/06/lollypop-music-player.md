+++
title = "音楽プレイヤー Lollypop を試す"
date =  "2019-06-26T12:00:37+09:00"
description = "私の場合は持ってるアルバムの中からテキトーなのを見繕って鳴らしてくれれば機能としては必要十分。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools", "music", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いま使ってる [Ubuntu] の音楽プレイヤーの既定は Rhythmbox なんだけど，イマイチ使い方が分からないんだよね。
私の場合は持ってるアルバム[^ma1] の中からテキトーなのを見繕って鳴らしてくれれば機能としては必要十分なので，もっとシンプルな GUI はないのか，と思ってたら [Lollypop] がよさげである。

[^ma1]: CD は学生の頃に買っていたものも全部含めて MP3 に切り刻んで NAS に置いている。

[Lollypop] は GNOME 用の GUI アプリケーションで，もちろん [Ubuntu] でも使える。
それどころか [Ubuntu] 既定のアプリケーションを Rhythmbox から [Lollypop] へ替えようかという議論まであったらしい。

- [Ubuntu 18.04 その37 - デフォルトのミュージックプレーヤーをLollypopに変更する提案 - kledgeb](https://kledgeb.blogspot.com/2018/01/ubuntu-1804-37-lollypop.html)

現在 [Ubuntu] 向けの [Lollypop] パッケージは [PPA (Personal Package Archive)](https://launchpad.net/ubuntu/+ppas) で提供されている。

- [Lollypop : Cédric Bellegarde](https://launchpad.net/~gnumdk/+archive/ubuntu/lollypop)

[Lollypop] 用の APT リポジトリの追加手順は以下の通り。

```text
$ sudo add-apt-repository ppa:gnumdk/lollypop
$ sudo apt update
```

これで [Lollypop] がインストール可能になった。

```text
$ apt show lollypop
Package: lollypop
Version: 1.1.1-1~disco
Priority: extra
Section: gnome
Maintainer: cedric.bellegarde@adishatz.org
```

ここまで来れば APT コマンド一発である。

```text
$ sudo apt install lollypop
```

これでアイコンがメニューに追加されたので起動してみる。

{{< fig-img src="./lollypop.png" link="./lollypop.png" width="827" >}}

おおっ，ちゃんと日本語で表示されるぢゃん。
設定や操作も直感的ですぐ分かった。

よーし，うむうむ，よーし。

[Lollypop]: https://wiki.gnome.org/Apps/Lollypop?action=show "Apps/Lollypop - GNOME Wiki!"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"

{{% review-paapi "B07STKKM3D" %}} <!-- 賢者の孫EDテーマ｢圧倒的Vivid Days｣ -->
