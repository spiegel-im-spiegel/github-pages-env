+++
title = "Git for Windows 2.17.0 がリリース（OpenSSL セキュリティアップデートを含む）"
date = "2018-04-08T11:32:36+09:00"
description = "HTTPS アクセスに OpenSSL を選択している人は早めのアップデートをお勧めする。"
image = "/images/attention/tools.png"
tags  = [ "tools", "git", "windows", "openssl", "security", "vulnerability" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Git for Windows] 2.17.0 がリリースされた。

- [Release Git for Windows 2.17.0 · git-for-windows/git](https://github.com/git-for-windows/git/releases/tag/v2.17.0.windows.1)

{{% fig-quote title="Release Git for Windows 2.17.0" link="https://github.com/git-for-windows/git/releases/tag/v2.17.0.windows.1" lang="en" %}}
- Comes with [Git v2.17.0](https://github.com/git/git/blob/v2.17.0/Documentation/RelNotes/2.17.0.txt).
- Comes with [OpenSSL v1.0.2o](https://www.openssl.org/news/openssl-1.0.2-notes.html).
- Comes with [Git Credential Manager v1.15.2](https://github.com/Microsoft/Git-Credential-Manager-for-Windows/releases/tag/v1.15.2).
- Comes with [OpenSSH v7.7p1](https://www.openssh.com/txt/release-7.7).
{{% /fig-quote %}}

OpenSSL v1.0.2o はセキュリティ・アップデートを含んでいる（深刻度：CVSSv3 Base 評価 で 5.9 【警告】）。

- [JVNVU#93502675: OpenSSL に複数の脆弱性](https://jvn.jp/vu/JVNVU93502675/)
- [JVNDB-2017-011252 - JVN iPedia - 脆弱性対策情報データベース](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-011252.html)

HTTPS アクセスに OpenSSL を選択している人は早めのアップデートをお勧めする。

アップデートは計画的に。

[Git for Windows]: https://gitforwindows.org/
