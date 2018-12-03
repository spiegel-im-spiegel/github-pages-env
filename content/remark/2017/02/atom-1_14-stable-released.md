+++
tags = ["atom", "editor", "tools", "conemu"]
description = "今回の目玉は，なんといっても Windows 版において64ビット版が登場したこと。ぶらぼー！"
date = "2017-02-09T21:01:02+09:00"
update = "2017-02-15T10:04:51+09:00"
title = "ATOM 1.14 stable がリリース"
draft = false

[author]
  flickr = "spiegel"
  twitter = "spiegel_2007"
  flattr = "spiegel"
  linkedin = "spiegelimspiegel"
  tumblr = ""
  github = "spiegel-im-spiegel"
  avatar = "/images/avatar.jpg"
  url = "http://www.baldanders.info/spiegel/profile/"
  facebook = "spiegel.im.spiegel"
  license = "by-sa"
  name = "Spiegel"
  instagram = "spiegel_2007"
+++

[ATOM] 1.14 の stable がリリースされている。

- [Release 1.14.0 · atom/atom](https://github.com/atom/atom/releases/tag/v1.14.0)
- [Atom 1.14 - Atom Blog](http://blog.atom.io/2017/02/08/atom-1-14.html)

今回の目玉は，なんといっても Windows 版において64ビット版が登場したこと。
ぶらぼー！

これまでの32ビット版から64ビット版に換装するには上書きインストールでよさげ。
Portable 版を使っている人はインストール先を分けて「とりあえず試してみる」でもいいかもしれない。

（最初 `AtomSetup-x64.msi` ファイルでインストールしようとしたのだがうまくいかなかった。 `AtomSetup-x64.exe` ファイルならうまくいく）

そしてもうひとつの目玉は，巨大ファイルのパフォーマンス向上だ。
基本コンポーネントを C++ で実装し直したようだ。

{{< fig-quote title="Atom 1.14" link="http://blog.atom.io/2017/02/08/atom-1-14.html" lang="en" >}}
<q>A fundamental component of the text editor called the display layer has been redesigned to rely on a new data structure that is implemented in C++. These changes enable Atom to open larger files more quickly while using much less memory. Improvements in this area are ongoing, so expect more in upcoming releases.</q>
{{< /fig-quote >}}

ちうことで，以降のバージョンでは更によくなるらしい。
期待しましょう。

## [platformio-ide-terminal] が動かない

ところで Windows 64ビット版に換装したら [platformio-ide-terminal] が動かなくなった。
ターミナル・ウィンドウは開くけど shell が立ち上がってこない。
困るなぁ。
まぁ，最悪ターミナルは別に [ConEmu] を使うからいいんだけど...

（[Issue](https://github.com/platformio/platformio-atom-ide-terminal/issues/155 "Not working with Atom 1.14 · Issue #155 · platformio/platformio-atom-ide-terminal") として既に上がっているようだ。はやく直ってぇ）

（追記）

[platformio-ide-terminal] 2.2.3 で修正された。
ありがとう！

[ATOM]: https://atom.io/ "Atom"
[platformio-ide-terminal]: https://atom.io/packages/platformio-ide-terminal
[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
