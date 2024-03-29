+++
date = "2015-10-30T00:56:49+09:00"
update = "2018-03-07T10:29:56+09:00"
description = "そんなわけで，えんやらやっと ConEmu を導入することにした。ついでに NYAGOS も入れなおすことに。"
tags = ["windows", "tools", "conemu", "nyagos", "terminal", "shell", "putty"]
title = "ようやく ConEmu と NYAGOS を導入した"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いやぁ， [Hugo] 使い出したらコマンドプロンプトがいくらあっても足りないのさ。
そんなわけで，えんやらやっと [ConEmu] を導入することにした。
ついでに [NYAGOS] も入れなおすことに。

## NYAGOS の導入

- [NYAOS.ORG - NYAGOS](http://www.nyaos.org/index.cgi?p=NYAGOS)
    - [nyaosorg/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows](https://github.com/nyaosorg/nyagos)

以前 [NYAGOS] を導入した時は[どえら苦労した](http://qiita.com/spiegel-im-spiegel/items/9c11acf72fa38ef379f8)が[^a]，[最近のバージョン](https://github.com/nyaosorg/nyagos/releases)は64ビット版のビルド済みのものがある。
ありがたくこのまま使わせて頂く[^nyole1]。

[^a]: まぁこれは [MSYS2] を試すきっかけになったので，結果的にはよかったのだが。  [MSYS2] の記事はそのうちちゃんと書かないといけないなぁ。
[^nyole1]: 最近のバージョンでは [zetamatta/nyole](https://github.com/zetamatta/nyole "zetamatta/nyole: Nihong Yet another OLE extension for lua") は不要になった。

zip ファイル内のファイル群をフォルダ構成ごと適当な場所に展開すればよい。
展開できたら動作確認。

```text
C:\program\nyagos>nyagos.exe
Nihongo Yet Another GOing Shell 4.2.3_3-amd64 by go1.9.3 & Lua 5.3
(c) 2014-2018 NYAOS.ORG <http://www.nyaos.org>
<hostname:~>
$ ls
Doc/          _nyagos       makeicon.cmd* nyagos.exe*   readme_ja.md
LICENSE       lua53.dll     nyagos.d/     readme.md
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

{{< fig-img src="https://photo.baldanders.info/flickr/image/22385157049_m.png" title="ConEmu Setting" link="https://photo.baldanders.info/flickr/22385157049/" >}}

フォントは既定のままで問題ない（もちろん好きなフォントに変えてもよい）。
ただし，日本語を使う場合は “Monospace” のチェックは必ず外すこと。

{{< fig-img src="https://photo.baldanders.info/flickr/image/21949162924_m.png" title="ConEmu Setting" link="https://photo.baldanders.info/flickr/21949162924/" >}}

“Center console in ConEmu workspace” にチェックを入れて数ピクセル程度アキを入れるとウィンドウ境界付近の窮屈感が解消される。
おススメ。

{{< fig-img src="https://photo.baldanders.info/flickr/image/22583116011_m.png" title="ConEmu Setting" link="https://photo.baldanders.info/flickr/22583116011/" >}}

Windows Explorer の context menu 設定。
上段の “ConEmu Here” を登録（register）すると， context menu を開いたフォルダで [ConEmu] を起動してくれる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/22583116021_m.png" title="ConEmu Setting" link="https://photo.baldanders.info/flickr/22583116021/" >}}

[ConEmu] 起動時の状態。
“Auto save/restore opened tabs” を選択すると，タブの状態を保持してくれる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/22575095305_m.png" title="ConEmu Setting" link="https://photo.baldanders.info/flickr/22575095305/" >}}

[ConEmu] 起動時の環境設定。
コードページの設定（`chcp` コマンド）もここでできるし alias の設定もできたりする。
ただし，環境設定は shell 呼び出し時にも個別にできるので，ここでは [ConEmu] 全体の設定に限定すべきだろう。

{{< fig-img src="https://photo.baldanders.info/flickr/image/22388493089_m.png" title="ConEmu Setting" link="https://photo.baldanders.info/flickr/22388493089/" >}}

Shell 呼び出しの設定例。
ここでは [MSYS2] の bash を呼んでいる。
起動時の calling sequence は以下のとおり。

```text
set MSYSTEM=MSYS & C:\msys64\usr\bin\bash.exe --login -i
```

前半で環境変数を設定し，後半で実際に bash を呼び出している。
やぁ，これで mintty を使わなくて済むよ。

[PuTTY] の場合は以下のようになる[^b]。

[^b]: ~~ただし [ConEmu] から [PuTTY] を呼び出す場合，サーバ側が UTF-8 だと文字化けするっぽい。ググると「`chcp 65001` に設定しろ」みたいなことが書いてあるが， [PuTTY] に対しては効いてない感じ。~~ （11月11日追記） 左記を訂正。オリジナルの [PuTTY] じゃなくて，日本語に対応している [PuTTYjp] なら問題なく日本語で表示される。ブラボー！

```text
C:\PATH\TO\PuTTY\putty.exe -load "mysession" -new_console
```

`mysession` には [PuTTY] であらかじめ作成したセッション名をセットする。

ちなみに [NYAGOS] の場合は `nyagos.exe` をフルパスで指定すれば OK。

ところで設定ダイアログの左下に [Flattr] ボタンがあるのにお気づきだろうか。
[Flattr] アカウントのある人は是非。

{{% div-box type="markdown" %}}
**追記 2018-03-07**

[ConEmu](https://conemu.github.io/ "ConEmu - Handy Windows Terminal") の画面の桁数を大きくすると画面が乱れる場合がある。
この場合は `Setting` → `Main` → `Compress long string to fit space` を無効にすることで改善するようだ。
{{% /div-box %}}

## ブックマーク

- [NYAGOS、かんたん設定 - Qiita](https://qiita.com/zetamatta/items/99feb1d74e36ea5848cd)
- [NYAGOS で環境設定系バッチファイルが期待どおり動かない原因と対策 - Qiita](https://qiita.com/zetamatta/items/f62bafd711755a4cf8d7)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
[NYAGOS]: http://www.nyaos.org/index.cgi?p=NYAGOS "NYAOS.ORG - NYAGOS"
[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
[PuTTY]: http://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free telnet/ssh client"
[PuTTYjp]: http://hp.vector.co.jp/authors/VA024651/PuTTYkj.html "hdk の自作ソフトの紹介 | PuTTYjp"
[Flattr]: https://flattr.com/
