+++
title = "とりあえず ATOM エディタ内ターミナルを x-terminal に乗り換えた"
date =  "2020-05-28T11:23:04+09:00"
description = "description"
image = "/images/attention/kitten.jpg"
tags = ["atom", "editor", "ubuntu", "terminal"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先週， [ATOM] エディタの [1.47 がリリース](https://github.com/atom/atom/releases/tag/v1.47.0 "Release 1.47.0 · atom/atom")され [Electron] も v5 ベースに上がったのだが，またもや [platformio-ide-terminal] が動かなくなった。
今回はちゃんとリビルドできてるのに。

一週間待ったが [platformio-ide-terminal] がアップデートされる気配がないので，諦めて他のパッケージを物色することにした。
まぁ，今は Linux/[Ubuntu] なので Windows の頃よりは選択肢が多いだろう。

ちうわけで幾つか試した結果，以下のパッケージに乗り換えた。

- [x-terminal]

このパッケージは atom-xterm からの fork らしいのだが，本家はメンテされなくなってから久しいようだ。
[x-terminal] の方はマメにアップデートされているようなので，しばらくはこちらを使ってみることにする。

しかし [ATOM] における [Electron] の外れっぷりはどうにかならないのかなぁ。
v5 なんかとっくの昔にサポートから外れてるっちうねん。
この機会に真剣に VS Code への乗り換えを検討したほうがええかもしらん[^edt1]。

[^edt1]: 5年前に秀丸から [ATOM] に[乗り換えた]({{< ref "/remark/2015/atom-editor.md" >}} "ATOM Editor に関するメモ")のだが，当時 VS Code は登場したばかりでチューニングの仕方とかよく分からなかったのだ。まぁ GitHub は Microsoft に買収されちゃったし [ATOM] はそのうちなくなって VS Code に一本化されちゃうのかもねぇ。

## ブックマーク

- [アプリフレームワーク「Electron 9」が公開 ～「PDFium」ベースのPDFリーダーが有効化 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1254642.html)

[ATOM]: https://atom.io/
[Electron]: https://electronjs.org/ "Electron | Build cross platform desktop apps with JavaScript, HTML, and CSS."
[platformio-ide-terminal]: https://atom.io/packages/platformio-ide-terminal
[x-terminal]: https://atom.io/packages/x-terminal
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
