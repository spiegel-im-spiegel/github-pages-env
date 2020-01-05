+++
title = "Ubuntu インストール用のブート可能 USB メモリを作成する"
date = "2019-04-21T22:20:33+09:00"
description = "まずは USB メモリを買うところから（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "virtualbox", "linux", "ubuntu", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

自宅 PC を [Ubuntu] に換装するためにインストール用の DVD を作ったのだが絶望的に遅いので USB メモリを使おうと思ったのだが，ここ10年くらい USB メモリなんか使ったことがなかったので，まずは USB メモリを買うところから（笑）

{{< fig-img src="./usb-memory.jpg" link="https://www.flickr.com/photos/spiegel/40690049803/" title="OS 移行用に購入。" >}}

ブート可能な USB メモリってどうやって作ろうかと思ってたら [Ubuntu] にツールがあった。

USB メモリを挿してメニューから「ブータブルUSBの作成」を起動する。

{{< fig-img src="./tool-make-bootable-usb-memory.png" link="./tool-make-bootable-usb-memory.png" width="818" >}}

「書き込み元のディスクイメージ」に [ISO イメージファイル](http://www.ubuntulinux.jp/download/ja-remix "Ubuntu Desktop 日本語 Remixのダウンロード | Ubuntu Japanese Team")を指定して「ブータブルUSBの作成」ボタンを押す。
確認プロンプトが出るので「はい」で書き込みを開始する。

{{< fig-img src="./tool-confirm.png" link="./tool-confirm.png" width="611" >}}

書き込みが完了したら以下のプロンプトが出ておしまい。

{{< fig-img src="./tool-conplete.png" link="./tool-conplete.png" width="1010" >}}

## [VirtualBox] を使って USB メモリからのインストールを試す

とりあえず [VirtualBox] で USB メモリからのインストールを試そうと思ったが， USB からの仮想ブートの仕方が分からない。
どうやらちょっと特殊な方法があるようだ。

- [Boot your USB Drive in VirtualBox - AgniPulse](http://agnipulse.com/2009/07/boot-your-usb-drive-in-virtualbox/)
- [Windows版VirtualBoxでUSBブートをする方法 - Qiita](https://qiita.com/ta_b0_/items/6946f6e62b6a8c5bb4bc)

この記事のとおりに USB ドライブに接続する vmdk ファイルを作って仮想ディスクとして追加すれば，この仮想ディスクからブートできる。

で，実際に試してみたら上手く行ったんだけどインストールの最後の最後で失敗するんだよなぁ。
まぁ DVD よりはアクセスが速いし本番で使えればいいか。

さて， GW に入ったらいよいよ OS を換装するんだ（フラグ）

## ブックマーク

- [Ubuntu 18.04 その85 - UbuntuのライブUSBメモリーを作成するには（Ubuntu/GNOME Disks編） - kledgeb](https://kledgeb.blogspot.com/2018/04/ubuntu-1804-85-ubuntuusbubuntugnome.html) : 「ディスク」ツールを使ってブータブル USB を作成する方法。 19.04 でもできることを確認済み

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[VirtualBox]: https://www.virtualbox.org/ "Oracle VM VirtualBox"

{{% review-paapi "B01NBU1OS5" %}} <!-- シリコンパワー USBメモリ 32GB USB3.1 -->
