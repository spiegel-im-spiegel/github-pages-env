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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E8%B3%A2%E8%80%85%E3%81%AE%E5%AD%ABED%EF%BE%83%EF%BD%B0%EF%BE%8F%EF%BD%A2%E5%9C%A7%E5%80%92%E7%9A%84Vivid-Days%EF%BD%A3-%E5%90%89%E4%B8%83%E5%91%B3%E3%80%82/dp/B07STKKM3D?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07STKKM3D"><img src="https://images-fe.ssl-images-amazon.com/images/I/610RL7gTRYL._SL160_.jpg" width="160" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E8%B3%A2%E8%80%85%E3%81%AE%E5%AD%ABED%EF%BE%83%EF%BD%B0%EF%BE%8F%EF%BD%A2%E5%9C%A7%E5%80%92%E7%9A%84Vivid-Days%EF%BD%A3-%E5%90%89%E4%B8%83%E5%91%B3%E3%80%82/dp/B07STKKM3D?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07STKKM3D">賢者の孫EDﾃｰﾏ｢圧倒的Vivid Days｣</a></dt>
	<dd>吉七味。 (メインアーティスト)</dd>
    <dd>avex pictures (Release 2019-06-19)</dd>
    <dd>Digital Music Album MP3 ダウンロード</dd>
    <dd>ASIN: B07STKKM3D</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">アニメ「賢者の孫」の ED 曲を含む。ヘヴィ・ローテーション中。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-06-26">2019-06-26</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
