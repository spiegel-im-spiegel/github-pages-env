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

「[**計算機・ネットワークは有益な目的に使おう**](https://baldanders.info/blog/000581/)」

という教えを胸に刻みつつ，今後も活動を継続していく所存である。

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[BOINC]: https://boinc.berkeley.edu/

{{% review-paapi "4898140866" %}} <!-- SETI@homeファンブック -->
