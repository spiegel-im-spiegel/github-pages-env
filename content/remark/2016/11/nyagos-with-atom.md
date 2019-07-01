+++
description = "やっとこれで Windows 環境でも ATOM にターミナル・エミュレータが導入できるようになった。しかも NYAGOS が動く。素晴らしい！"
tags = [
  "atom",
  "editor",
  "nyagos",
]
draft = false
date = "2016-11-20T16:29:24+09:00"
title = "ATOM × NYAGOS ＝ ♥"

[author]
  flattr = ""
  twitter = "spiegel_2007"
  flickr = "spiegel"
  avatar = "/images/avatar.jpg"
  license = "by-sa"
  tumblr = ""
  instagram = "spiegel_2007"
  url = "https://baldanders.info/profile/"
  facebook = "spiegel.im.spiegel"
  github = "spiegel-im-spiegel"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
+++

[ATOM] 上で動くターミナル・エミュレータはいくつかあるが，大抵は UNIX 互換環境を前提としたもので Windows では動かなかったり特別なインストール手順を要するものが殆どで使う気にならなかったりする。
そんな中で [platformio-ide-terminal] は Windows 環境でも簡単に導入できる優れもののようだ。

- [Atomのterminal-plusが動かないのでplatformio-ide-terminalを入れてみた | spiffieldLabs](http://labs.spiffield.net/archives/508)
- [ATOMエディタではじめるマイナー言語探訪(あんま怖くないNim②) - Qiita](http://qiita.com/kmry2045/items/b61a000ff257c01720e4)

インストールは，他のパッケージと同じように， Settings の Install から選択してインストールすれば OK。
マジ簡単。

{{< fig-img src="https://photo.baldanders.info/flickr/image/31089552246_m.png" title="platformio-ide-terminal (1)" link="https://photo.baldanders.info/flickr/31089552246/" >}}

ふむむ。
どうやら既定では PawerShell が立ち上がるようだ。
ってことは他の shell でもいけるんじゃね？

ということで Settings を見ると

{{< fig-img src="https://photo.baldanders.info/flickr/image/31089552576_m.png" title="platformio-ide-terminal (2)" link="https://photo.baldanders.info/flickr/31089552576/" >}}

おおっ。
やっぱり shell を変えられるんだ。
ほんじゃあ [NYAGOS] を入れてみよっか。
「Shell Override」の項目に [NYAGOS] をフルパスで指定して  [platformio-ide-terminal] を起動する。

{{< fig-img src="https://photo.baldanders.info/flickr/image/31089552456_m.png" title="platformio-ide-terminal (3)" link="https://photo.baldanders.info/flickr/31089552456/" >}}

おおおおっ！ なにこれ素敵。
[NYAGOS] であれば UTF-8 を受け入れるので文字コードがどうとかあまり考えなくてよい。

ん？ 待てよ。
ってことは git bash もいけるのか？ では同じように「Shell Override」の項目に git bash をフルパスで指定指定して  [platformio-ide-terminal] を起動してみる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/31089552546_m.png" title="platformio-ide-terminal (4)" link="https://photo.baldanders.info/flickr/31089552546/" >}}

よいではないか。

[platformio-ide-terminal] を起動する際は， Windows では， `alt-shift-T` を押下すればいいのだが，キーを3つ押さえるのは得意ではないので `ctrl-f1` に振り直した。

```cson
'.platform-win32 atom-workspace':
  'ctrl-f1': 'platformio-ide-terminal:new'
 ```

あぁ。
やっとこれで Windows 環境でも [ATOM] にターミナル・エミュレータが導入できるようになった。
しかも [NYAGOS] が動く。
素晴らしい！

[ATOM]: https://atom.io/ "Atom"
[NYAGOS]: http://www.nyaos.org/index.cgi?p=NYAGOS "NYAOS.ORG - NYAGOS"
[platformio-ide-terminal]: https://atom.io/packages/platformio-ide-terminal
