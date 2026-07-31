+++
title = "LibreOffice 6.3 へのアップグレード"
date =  "2019-08-19T19:51:39+09:00"
description = "6.2 系と 6.3 系はパソコン上で共存できてしまうので 6.3 系のみを使うのであれば 6.2 系は削除すること。"
image = "/images/attention/tools.png"
tags = [ "tools", "libreoffice", "openpgp", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[LibreOffice] 6.2.6 および 6.3 がリリースされた。

- [The Document Foundation announces LibreOffice 6.3 - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/08/08/tdf-announces-libreoffice-63/)
- [LibreOffice 6.2.6 is ready, all users should update for enhanced security - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/08/14/libreoffice-626/)

[LibreOffice] 6.2.6 はセキュリティ・アップデートを含んでいるので早めのアップデートを。

今回は [LibreOffice] 6.3 へのアップグレードを行う。

[Ubuntu] では [LibreOffice] を導入するための手段がいくつか存在するが [OpenPGP] 鍵を使って暗号化や電子署名を行うのであれば[公式サイト]から `*.deb` ファイルを取ってきてインストールするしかない。
やり方は以下の記事を参考にどうぞ。

- [Ubuntu に LibreOffice をインストールする3つの方法]({{< ref "/remark/2019/05/installing-libreoffice-in-ubuntu.md" >}})

6.2 系と 6.3 系はパソコン上で共存できてしまうので 6.3 系のみを使うのであれば 6.2 系は削除すること。

```text
$ sudo apt remove libreoffice6.2* libobasis6.2*
```

細かくチェックしたわけではないが 6.3 でも [OpenPGP] 鍵で暗号化したファイルが開けたので，まぁ問題なかろう。

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[公式サイト]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[OpenPGP]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
