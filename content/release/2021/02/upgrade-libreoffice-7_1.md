+++
title = "LibreOffice 7.1 へのアップグレード"
date =  "2021-02-07T16:48:31+09:00"
description = "いつもどおり，手動でアップグレード。"
image = "/images/attention/tools.png"
tags = [ "tools", "libreoffice", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[LibreOffice] 7.1 がリリースされた。

- [LibreOffice 7.1 Community released by The Document Foundation - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2021/02/03/libreoffice-7-1-community/)
- [無償のオフィス総合ソフト「LibreOffice 7.1 Community」がリリース - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1304425.html)

[Ubuntu] で [LibreOffice] を導入する方法はいくつかあるが，[前回]({{< ref "/release/2020/08/upgrade-libreoffice-7_0.md" >}})と同じく今回も[公式サイト]から `*.deb` ファイルを取ってきてインストールした。
細かくチェックしたわけではないが 7.0 でも [OpenPGP] 鍵で暗号化したファイルが開けたので，まぁ問題なかろう。

[LibreOffice] は異なるバージョンと共存できてしまうので 7.1 系のみを使うのであれば以前のバージョンは削除する。

```text
$ sudo apt purge libreoffice7.0* libobasis7.0*
```

7.1 系を削除する場合も

```text
$ sudo apt purge libreoffice7.1* libobasis7.1*
```

でよい。

もし上の APT コマンドで上手く行かない場合は

```text
$ dpkg -l | grep libreoffice
```

あるいは

```text
$ dpkg -l | grep libobasis
```

で一覧を出してひとつづつ `sudo pkg -r` コマンドで削除していくしかない（依存関係に注意）。

## ブックマーク

- [Ubuntu に LibreOffice をインストールする3つの方法]({{< ref "/remark/2019/05/installing-libreoffice-in-ubuntu.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[公式サイト]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[OpenPGP]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
