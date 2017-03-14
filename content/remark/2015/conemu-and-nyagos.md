+++
date = "2015-10-30T00:56:49+09:00"
update = "2017-02-27T09:38:44+09:00"
description = "そんなわけで，えんやらやっと ConEmu を導入することにした。ついでに NYAGOS も入れなおすことに。"
draft = false
tags = ["windows", "tools", "conemu", "nyagos", "terminal", "shell", "putty"]
title = "ようやく ConEmu と NYAGOS を導入した"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

いやぁ， [Hugo] 使い出したらコマンドプロンプトがいくらあっても足りないのさ。
そんなわけで，えんやらやっと [ConEmu] を導入することにした。
ついでに [NYAGOS] も入れなおすことに。

## NYAGOS の導入

- [NYAOS.ORG - NYAGOS](http://www.nyaos.org/index.cgi?p=NYAGOS)
    - [zetamatta/nyagos](https://github.com/zetamatta/nyagos)
    - [zetamatta/nyole](https://github.com/zetamatta/nyole)

以前 [NYAGOS] を導入した時は[どえら苦労した](http://qiita.com/spiegel-im-spiegel/items/9c11acf72fa38ef379f8)が[^a]，[最近のバージョン](https://github.com/zetamatta/nyagos/releases)は64ビット版のビルド済みのものがある。
ありがたくこのまま使わせて頂く。
[nyole](https://github.com/zetamatta/nyole) も同梱されているので無問題。

[^a]: まぁこれは [MSYS2] を試すきっかけになったので，結果的にはよかったのだが。  [MSYS2] の記事はそのうちちゃんと書かないといけないなぁ。

zip ファイル内のファイル群をフォルダ構成ごと適当な場所に展開すればよい。
展開できたら動作確認。

```
C:\program\nyagos>nyagos.exe
Nihongo Yet Another GOing Shell 4.0.9_10-amd64 Powered by go1.5.1 & Lua 5.3
Copyright (c) 2014,2015 HAYAMA_Kaoru and NYAOS.ORG
C:/program/nyagos>ls
Doc/               lua53.dll          nyagos.lua*        specialfolders.js*
catalog.d/         makeicon.cmd*      nyole.dll
license.txt        nyagos.d/          readme.md
lnk.js*            nyagos.exe*        readme_ja.md
```

ほい，おっけ。

## ConEmu の導入

[ConEmu] は Windows 用のターミナルでタブごとに異なる shell を呼び出せるのが特徴。
コマンドプロンプトや上述の [NYAGOS] はもちろん， [MSYS2] の bash や [PuTTY] も呼び出せてしまう優れもの。

- [ConEmu - Handy Windows Terminal](https://conemu.github.io/)
    - [Windows：コマンドプロンプト代替をConsole2からConEmuに変更](http://kenpg.bitbucket.org/blog/201506/07.html)
    - [ConEmu 突っ込んだら Git for Windows の Git Bash がカッコよくなった - てっく煮ブログ](http://tech.nitoyon.com/ja/blog/2014/03/07/fancy-git-bash/)
    - [ConEmu + nyagos で Windows ターミナル環境を作る - Qiita](http://qiita.com/1000k/items/4a2f9419b19fdc9ed5f4)
    - [Configuring ConEmu and Putty | theCRUMB](http://thecrumb.com/2013/03/04/configuring-conemu-and-putty/)
    - [ConEmu + PuTTYでSSHクライアントを快適に使う（Windows版） | Black Everyday Company](http://kuroeveryday.blogspot.jp/2015/10/ConEmu-PuTTY.html)

以降，覚え書きで [ConEmu] の設定を晒しておく。

{{< fig-img flickr="true" src="https://farm1.staticflickr.com/716/22385157049_4c7af6ef0a.jpg" alt="ConEmu Setting" title="ConEmu Setting" link="https://www.flickr.com/photos/spiegel/22385157049/" >}}

フォントは既定のままで問題ない（もちろん好きなフォントに変えてもよい）。
ただし，日本語を使う場合は “Monospace” のチェックは必ず外すこと。

{{< fig-img flickr="true" src="https://farm6.staticflickr.com/5792/21949162924_ba2742c857.jpg" alt="ConEmu Setting" title="ConEmu Setting" link="https://www.flickr.com/photos/spiegel/21949162924/" >}}

“Center console in ConEmu workspace” にチェックを入れて数ピクセル程度アキを入れるとウィンドウ境界付近の窮屈感が解消される。
おススメ。

{{< fig-img flickr="true" src="https://farm1.staticflickr.com/591/22583116011_9e109f185e.jpg" alt="ConEmu Setting" title="ConEmu Setting" link="https://www.flickr.com/photos/spiegel/22583116011/" >}}

Windows Explorer の context menu 設定。
上段の “ConEmu Here” を登録（register）すると， context menu を開いたフォルダで [ConEmu] を起動してくれる。

{{< fig-img flickr="true" src="https://farm6.staticflickr.com/5630/22583116021_222c1acaee.jpg" alt="ConEmu Setting" title="ConEmu Setting" link="https://www.flickr.com/photos/spiegel/22583116021/" >}}

[ConEmu] 起動時の状態。
“Auto save/restore opened tabs” を選択すると，タブの状態を保持してくれる。

{{< fig-img flickr="true" src="https://farm1.staticflickr.com/644/22575095305_4d0c7752e6.jpg" alt="ConEmu Setting" title="ConEmu Setting" link="https://www.flickr.com/photos/spiegel/22575095305/" >}}

[ConEmu] 起動時の環境設定。
コードページの設定（`chcp` コマンド）もここでできるし alias の設定もできたりする。
ただし，環境設定は shell 呼び出し時にも個別にできるので，ここでは [ConEmu] 全体の設定に限定すべきだろう。

{{< fig-img flickr="true" src="https://farm1.staticflickr.com/698/22388493089_73bb752b56.jpg" alt="ConEmu Setting" title="ConEmu Setting" link="https://www.flickr.com/photos/spiegel/22388493089/" >}}

Shell 呼び出しの設定例。
ここでは [MSYS2] の bash を呼んでいる。
起動時の calling sequence は以下のとおり。

```
set MSYSTEM=MSYS & C:\msys64\usr\bin\bash.exe --login -i
```

前半で環境変数を設定し，後半で実際に bash を呼び出している。
やぁ，これで mintty を使わなくて済むよ。

[PuTTY] の場合は以下のようになる[^b]。

[^b]: ~~ただし [ConEmu] から [PuTTY] を呼び出す場合，サーバ側が UTF-8 だと文字化けするっぽい。ググると「`chcp 65001` に設定しろ」みたいなことが書いてあるが， [PuTTY] に対しては効いてない感じ。~~ （11月11日追記） 左記を訂正。オリジナルの [PuTTY] じゃなくて，日本語に対応している [PuTTYjp] なら問題なく日本語で表示される。ブラボー！

```
C:\PATH\TO\PuTTY\putty.exe -load "mysession" -new_console
```

`mysession` には [PuTTY] であらかじめ作成したセッション名をセットする。

ちなみに [NYAGOS] の場合は `nyagos.exe` をフルパスで指定すれば OK。

ところで設定ダイアログの左下に [Flattr] ボタンがあるのにお気づきだろうか。
[Flattr] アカウントのある人は是非。

[Hugo]: https://gohugo.io/ "Hugo :: A fast and modern static website engine"
[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
[NYAGOS]: http://www.nyaos.org/index.cgi?p=NYAGOS "NYAOS.ORG - NYAGOS"
[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
[PuTTY]: http://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free telnet/ssh client"
[PuTTYjp]: http://hp.vector.co.jp/authors/VA024651/PuTTYkj.html "hdk の自作ソフトの紹介 | PuTTYjp"
[Flattr]: https://flattr.com/
