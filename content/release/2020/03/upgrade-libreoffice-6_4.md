+++
title = "LibreOffice 6.4 へのアップグレード"
date =  "2020-03-01T10:46:04+09:00"
description = "前回と同じく今回も公式サイトからファイルを取ってきてインストールした。"
image = "/images/attention/tools.png"
tags = [ "tools", "libreoffice", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[LibreOffice] 6.4 がリリースされた。

- [Performance-focused LibreOffice 6.4 is available for download - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2020/01/29/libreoffice-6-4/)
- [The Document Foundation announces LibreOffice 6.4.1 - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2020/02/27/libreoffice-641/)

[Ubuntu] で [LibreOffice] を導入する方法はいくつかあるが，[前回]({{< ref "/release/2019/08/upgrade-libreoffice-6_3.md" >}} "LibreOffice 6.3 へのアップグレード")と同じく今回も[公式サイト]から `*.deb` ファイルを取ってきてインストールした。
細かくチェックしたわけではないが 6.4 でも [OpenPGP] 鍵で暗号化したファイルが開けたので，まぁ問題なかろう。

6.3 系と 6.4 系はパソコン上で共存できてしまうので 6.4 系のみを使うのであれば以前の 6.3 系は削除する。

```text
$ sudo apt remove libreoffice6.3* libobasis6.3*
```

6.4 系を削除する場合も

```text
$ sudo apt remove libreoffice6.4* libobasis6.4*
```

でよい。

## ブックマーク

- [Ubuntu に LibreOffice をインストールする3つの方法]({{< ref "/remark/2019/05/installing-libreoffice-in-ubuntu.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[公式サイト]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[OpenPGP]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
