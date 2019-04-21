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

とりあえず [VirtualBox] を USB メモリからのインストールを試そうと思ったが， USB からの仮想ブートの仕方が分からない。
どうやらちょっと特殊な方法があるようだ。

- [Boot your USB Drive in VirtualBox - AgniPulse](http://agnipulse.com/2009/07/boot-your-usb-drive-in-virtualbox/)
- [Windows版VirtualBoxでUSBブートをする方法 - Qiita](https://qiita.com/ta_b0_/items/6946f6e62b6a8c5bb4bc)

この記事のとおりに USB ドライブに接続する vmdk ファイルを作って仮想ディスクとして追加すれば，この仮想ディスクからブートできる。

で，実際に試してみたら上手く行ったんだけどインストールの最後の最後で失敗するんだよなぁ。
まぁ DVD よりはアクセスが速いし本番で使えればいいか。

さて， GW に入ったらいよいよ OS を換装するんだ（フラグ）

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[VirtualBox]: https://www.virtualbox.org/ "Oracle VM VirtualBox"

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%B7%E3%83%AA%E3%82%B3%E3%83%B3%E3%83%91%E3%83%AF%E3%83%BC-USB%E3%83%A1%E3%83%A2%E3%83%AA-%E4%BA%9C%E9%89%9B%E5%90%88%E9%87%91%E3%83%9C%E3%83%87%E3%82%A3-PS4%E5%8B%95%E4%BD%9C%E7%A2%BA%E8%AA%8D%E6%B8%88-SP032GBUF3J80V1TJA/dp/B01NBU1OS5?psc=1&SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01NBU1OS5"><img src="https://images-fe.ssl-images-amazon.com/images/I/317fsDSqG8L._SL160_.jpg" width="160" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%B7%E3%83%AA%E3%82%B3%E3%83%B3%E3%83%91%E3%83%AF%E3%83%BC-USB%E3%83%A1%E3%83%A2%E3%83%AA-%E4%BA%9C%E9%89%9B%E5%90%88%E9%87%91%E3%83%9C%E3%83%87%E3%82%A3-PS4%E5%8B%95%E4%BD%9C%E7%A2%BA%E8%AA%8D%E6%B8%88-SP032GBUF3J80V1TJA/dp/B01NBU1OS5?psc=1&SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01NBU1OS5">シリコンパワー USBメモリ 32GB USB3.0  亜鉛合金ボディ 防水 防塵 耐衝撃  永久保証 PS4動作確認済 Jewel J80 SP032GBUF3J80V1TJA</a></dt>
    <dd>シリコンパワー (Release 2017-02-01)</dd>
    <dd>Personal Computer Personal Computers</dd>
    <dd>ASIN: B01NBU1OS5, EAN: 4712702655292</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">OS 移行用に購入。ひたすらデザインで決めた（笑） よく考えたら32GBも要らなかった。まぁ，何か使い道があるだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-04-21">2019-04-21</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
