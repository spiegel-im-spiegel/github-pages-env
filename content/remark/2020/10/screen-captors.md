+++
title = "スクリーン・キャプチャのキホン"
date =  "2020-10-29T20:56:22+09:00"
description = "それ，知らんかっとってんちんとんしゃん"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "windows", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

最初にものっそい言い訳してしまうと，普段スクリーン・キャプチャ機能なんか使わんのですよ，私は。
それでも `[Print Screen]` キーを押せば画面が撮れることを覚えていただけでも褒めて欲しい（笑）

で，今までは `[Print Screen]` キーでスクリーン全体を撮るか `[Alt] + [Print Screen]` でアクティブなウィンドウ領域を撮るかくらいしか知らなかったのだが， Windows 10 では `[`{{< icons "windows">}}`] + [Shift] + [s]` で任意の矩形領域が撮れると教えてもらって驚愕した。

`[`{{< icons "windows">}}`] + [Shift] + [s]` 押下で画面上部に以下のようなコントロールが表示される。

{{< fig-img src="./screen-captors.png" link="./screen-captors.png" >}}

左端のアイコンを選択しているこの状態が任意の矩形領域を切り取れるモードだ。

更にこの機能を `[Print Screen]` キーに割り当てることができるようだ。
詳しい手順は以下の記事を参考にどうぞ。

- [［Windows］+［Shift］+［S］で画面領域切り取り（PrintScreenに割り当ても可） | Windows 10 | 初心者のためのOffice講座](https://hamachan.info/win10-win-winss/)

更に更に，セルフタイマー付きスクリーン・キャプチャ機能や切り取った画像を加工できる “Snipping Tool” が Windows 標準のアクセサリとして搭載されていた。
それ，知らんかっとってんちんとんしゃん。

今までペイントツールでちまちまとトリミングしてた私の苦労は何だったのか `orz`

ところで [Ubuntu] でも簡単に矩形領域を撮れないの？ と軽くググってみたが， `[Shift] + [Print Screen]` 押下で矩形領域を切り取れるようだ。
Windows より簡単だったよ。

今まで Shotwell Viewer でちまちまとトリミングしてた私の苦労は何だったのか `orz`

[Ubuntu] の場合，スクリーン・キャプチャ機能の実体は `gnome-screenshot` コマンドで，これはコマンドラインからも起動できる。
たとえば，5秒間の遅延後にアクティブなウィンドウを撮りたければ

```text
$ gnome-screenshot --window --delay=5
```

とすればいいようだ。

## ブックマーク

- [Copy the window or screen contents - Office Support](https://support.microsoft.com/en-us/office/copy-the-window-or-screen-contents-98c41969-51e5-45e1-be36-fb9381b32bb7)
- [Snipping Tool を使ってスクリーン ショットをキャプチャする](https://support.microsoft.com/ja-jp/windows/snipping-tool-%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%BC%E3%83%B3-%E3%82%B7%E3%83%A7%E3%83%83%E3%83%88%E3%82%92%E3%82%AD%E3%83%A3%E3%83%97%E3%83%81%E3%83%A3%E3%81%99%E3%82%8B-00246869-1843-655f-f220-97299b865f6b)
- [Ubuntuで端末からスクリーンショットを撮る方法まとめ - Qiita](https://qiita.com/yas-nyan/items/80f2db8c4bdf4c8e87b8)
- [2020年10月現在スクリーンキャプチャおすすめツール(Windows/Mac)](https://zenn.dev/junki555/articles/982b5f3f124da1fb5548)
- [Snagit = The Best Screen Capture Software (Free Trial) | TechSmith](https://www.techsmith.com/screen-capture.html) : 有料だがかなり「使える！」らしい

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4257904623" %}} <!-- ちまりまわるつ -->
{{% review-paapi "B004JSTW5S" %}} <!-- カードキャプターさくら 1期+2期+3期 コンプリート DVD-BOX -->
