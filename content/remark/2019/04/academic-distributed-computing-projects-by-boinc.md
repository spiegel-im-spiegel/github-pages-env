+++
title = "BOINC による学術分散コンピューティング・プロジェクトでの活動を再開した"
date =  "2019-04-30T20:13:42+09:00"
description = "「計算機・ネットワークは有益な目的に使おう」"
image = "/images/attention/kitten.jpg"
tags = [ "boinc", "internet", "linux", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

OS 移行の準備のためにこの一ヶ月ほど [BOINC (Berkeley Open Infrastructure for Network Computing)](https://boinc.berkeley.edu/) での活動を中断していたのだが，換装した [Ubuntu] も安定的に稼働しているようだし，再開することにした。

{{< fig-img src="./boinc-manager.jpg" title="OS 移行のため中断していたが BOINC による学術分散コンピューティング・プロジェクトの活動を再開した" link="https://www.flickr.com/photos/spiegel/33863325058/" >}}

Linux 版の BOINC クライアントのインストールは[パッケージ管理ツールを使うことを推奨](https://boinc.berkeley.edu/wiki/Installing_BOINC "Installing BOINC - BOINC]")しているようで [Ubuntu] でも APT (Advanced Package Tool) を使うのがいいようだ。

```text
$ apt show boinc-client
Package: boinc-client
Version: 7.14.2+dfsg-3
Priority: optional
Section: universe/net
Source: boinc
Origin: Ubuntu
```

というわけで早速インストール。

```text
$ sudo apt install boinc-client boinc-manager
パッケージリストを読み込んでいます... 完了
依存関係ツリーを作成しています                
状態情報を読み取っています... 完了
以下の追加パッケージがインストールされます:
  libboinc7 libwxbase3.0-0v5 libwxgtk-webview3.0-gtk3-0v5 libwxgtk3.0-gtk3-0v5
提案パッケージ:
  boinc-client-opencl boinc-client-nvidia-cuda libgl1-mesa-glx
以下のパッケージが新たにインストールされます:
  boinc-client boinc-manager libboinc7 libwxbase3.0-0v5 libwxgtk-webview3.0-gtk3-0v5 libwxgtk3.0-gtk3-0v5
アップグレード: 0 個、新規インストール: 6 個、削除: 0 個、保留: 0 個。
7,421 kB のアーカイブを取得する必要があります。
この操作後に追加で 29.5 MB のディスク容量が消費されます。
```

まずは計算リソースの割り当てを少なめにして様子を見る予定だ。

私としては [SETI@home](https://setiathome.berkeley.edu/) を通じて学んだ

「[**計算機・ネットワークは有益な目的に使おう**](https://baldanders.info/spiegel/log2/000581.shtml)」

という教えを胸に刻みつつ，今後も活動を継続していく所存である。

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[BOINC]: https://boinc.berkeley.edu/

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/SETI-home%E3%83%95%E3%82%A1%E3%83%B3%E3%83%96%E3%83%83%E3%82%AF%E2%80%95%E3%81%8A%E3%81%86%E3%81%A1%E3%81%AE%E3%83%91%E3%82%BD%E3%82%B3%E3%83%B3%E3%81%A7%E5%AE%87%E5%AE%99%E4%BA%BA%E6%8E%A2%E3%81%97-%E9%87%8E%E5%B0%BB-%E6%8A%B1%E4%BB%8B/dp/4898140866?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4898140866"><img src="https://images-fe.ssl-images-amazon.com/images/I/51A74XV7MDL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/SETI-home%E3%83%95%E3%82%A1%E3%83%B3%E3%83%96%E3%83%83%E3%82%AF%E2%80%95%E3%81%8A%E3%81%86%E3%81%A1%E3%81%AE%E3%83%91%E3%82%BD%E3%82%B3%E3%83%B3%E3%81%A7%E5%AE%87%E5%AE%99%E4%BA%BA%E6%8E%A2%E3%81%97-%E9%87%8E%E5%B0%BB-%E6%8A%B1%E4%BB%8B/dp/4898140866?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4898140866">SETI@homeファンブック―おうちのパソコンで宇宙人探し</a></dt>
	<dd>野尻 抱介</dd>
    <dd>ローカス 2000-01</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4898140866, EAN: 9784898140864</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">内容は古いけど当時の「熱」を伝えた名著だと思うけどなぁ。著者の方が自己出版で Kindle で出してくれたらいいのに。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-28">2019-03-28</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
