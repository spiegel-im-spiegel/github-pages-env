+++
title = "LibreOffice 7.0 へのアップグレード"
date =  "2020-08-11T12:21:15+09:00"
description = "前回と同じく今回も公式サイトからファイルを取ってきてインストールした。"
image = "/images/attention/tools.png"
tags = [ "tools", "libreoffice", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[LibreOffice] 7.0 がリリースされた。

- [Announcement of LibreOffice 7.0 - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2020/08/05/announcement-of-libreoffice-7-0/)
- [Microsoft Officeとの互換性も向上した「LibreOffice 7.0」が正式版に ～Skia/Vulkanにも対応 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1269685.html)

しばらく静観していたが，緊急パッチもなさそうなので，このままアップグレードすることにした。

[Ubuntu] で [LibreOffice] を導入する方法はいくつかあるが，[前回]({{< ref "/release/2020/03/upgrade-libreoffice-6_4.md" >}} "LibreOffice 6.4 へのアップグレード")と同じく今回も[公式サイト]から `*.deb` ファイルを取ってきてインストールした。
細かくチェックしたわけではないが 7.0 でも [OpenPGP] 鍵で暗号化したファイルが開けたので，まぁ問題なかろう。

[LibreOffice] は異なるバージョンと共存できてしまうので 7.0 系のみを使うのであれば以前のバージョンは削除する。

```text
$ sudo apt purge libreoffice6.4* libobasis6.4*
```

7.0 系を削除する場合も

```text
$ sudo apt purge libreoffice7.0* libobasis7.0*
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
つか， [Ubuntu] 20.04 にアップグレードする際に [LibreOffice] 6.4 の APT パッケージ情報が上書きされてしまったみたいで，手動でひとつづつ削除するしかなかった `orz`

これから [Ubuntu] をアップグレードする際には気をつけないとなー

## ブックマーク

- [Announcement of LibreOffice 6.4.6 - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2020/08/13/announcement-of-libreoffice-6-4-6/)

- [Ubuntu に LibreOffice をインストールする3つの方法]({{< ref "/remark/2019/05/installing-libreoffice-in-ubuntu.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[公式サイト]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[OpenPGP]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
