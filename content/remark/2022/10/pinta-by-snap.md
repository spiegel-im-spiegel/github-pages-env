+++
title = "Linux 版 Pinta は Snap でインストールしろ！"
date =  "2022-10-30T18:22:06+09:00"
description = "別途 Mono 入れなくても動いた！"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Windows の「ペイント」みたいな，機能はそんなに要らないけど簡単に描ける Ubuntu で動くお絵かきソフトが欲しい！ と常々思ってたけど Pinta ってのがいいらしい。

- [Linuxでもペイントは使えるんやで - Qiita](https://qiita.com/kaitaku/items/dd20c292c903b4b62b17)

この記事に従って PPA から apt install しようとしたんだけど

```text
$ sudo add-apt-repository ppa:pinta-maintainers/pinta-stable
リポジトリ: 'deb https://ppa.launchpadcontent.net/pinta-maintainers/pinta-stable/ubuntu/ kinetic main'
概要：
Note: this PPA only provides GTK2 builds of Pinta (1.7.x) and is no longer updated for Pinta 2.0 and later. Installing the flatpak or snap package is recommended instead.

It is strongly recommended to also ensure that Mono 6.10 or higher is installed on your system (run `mono --version` to check), as earlier 6.x versions can produce random crashes. You can install the latest Mono version from https://www.mono-project.com/download/stable/#download-lin
より詳しい情報: https://launchpad.net/~pinta-maintainers/+archive/ubuntu/pinta-stable
リポジトリを追加しています。
続けるには「Enter」キーを、中止するにはCtrl-cを押してください。
```

とりあえず `Ctrl+C` で中止（笑）

なるほど，最新バージョンは apt じゃなくて snap から入れろと。
ほんで Mono 6 が要るのか。
ふむむー。

もう少しググってみたら，こんな記事を見つけた。

- [Pinta 2.0 Released, Completes Port to GTK3 & .NET 6 - OMG! Ubuntu!](https://www.omgubuntu.co.uk/2022/01/pinta-2-0-released-completes-port-to-gtk3-net-6)

{{< fig-quote type="markdown" title="Pinta 2.0 Released, Completes Port to GTK3 & .NET 6" link="https://www.omgubuntu.co.uk/2022/01/pinta-2-0-released-completes-port-to-gtk3-net-6" lang="en" >}}
And on Windows and macOS (Pinta is cross-platform, remember) the installers now bundle all of the necessary dependencies meaning users no longer need to head off to hunt down installers for GTK or .NET/Mono themselves. The app is now notarised on macOS too.
{{< /fig-quote >}}

んー？ ひょっとして Linux 版も Mono を別途入れなくても大丈夫とか？ 試してみるか。

```text
$ snap install pinta
pinta 2.0.2 from James Carroll✪ installed
```

おっ，普通にインストールできた（ちなみに，今の私の環境に Mono は入れてない）。

{{< fig-img src="./pinta.png" title="Pinta 2.0.2" link="./pinta.png" width="1204" >}}

[こいつ，動くぞ](https://dic.nicovideo.jp/a/%E3%81%93%E3%81%84%E3%81%A4%E3%83%BB%E3%83%BB%E3%83%BB%E5%8B%95%E3%81%8F%E3%81%9E%21)！

ちょっと触ってみたが，日本語入力に若干難はあるものの，問題なく動いてる。
凄いな .NET 6。
別途 devel とか runtime とか入れなくていいんだ。
仕事でも .NET 6 へのリプレイス案件とかチラホラ聞くようになったし，そろそろ真面目に勉強したほうがいいか知らん。

画面のスナップショットに印をつけたりコメントを入れたりする程度で GIMP 起動するのもなぁ，と思ってたので，これから便利に使わせていただきます。

<!-- eof -->
