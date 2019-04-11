+++
title = "Git for Windows 2.17.1 がリリース（セキュリティアップデート）"
date = "2018-05-30T20:58:48+09:00"
update = "2018-06-03T17:32:14+09:00"
description = "Git v2.17.1 にはセキュリティ・アップデートがある。"
image = "/images/attention/tools.png"
tags  = [ "tools", "git", "windows", "security", "vulnerability" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Git for Windows] 2.17.1 がリリースされた。

- [Release Git for Windows 2.17.1 · git-for-windows/git](https://github.com/git-for-windows/git/releases/tag/v2.17.1.windows.1)

{{% fig-quote-md title="Release Git for Windows 2.17.1" link="https://github.com/git-for-windows/git/releases/tag/v2.17.1.windows.1" lang="en" %}}
- Comes with [Git v2.17.1](https://github.com/git/git/blob/v2.17.1/Documentation/RelNotes/2.17.1.txt).
- Comes with [Perl v5.26.2](http://search.cpan.org/dist/perl-5.26.2/pod/perldelta.pod).
- The installer [now offers VS Code Insiders as option for Git's default editor](https://github.com/git-for-windows/build-extra/pull/181) if it is installed.
- The vim configuration [was modernized](https://github.com/git-for-windows/build-extra/pull/181).
- Comes with [cURL v7.60.0](https://curl.haxx.se/changes.html#7_60_0).
- Certain errors, e.g. when pushing failed due to a non-fast-forwarding change, [are now colorful](https://github.com/git-for-windows/git/pull/1429).
- Comes with [Git LFS v2.4.1](https://github.com/git-lfs/git-lfs/releases/tag/v2.4.1).
{{% /fig-quote-md %}}

[Git v2.17.1](https://github.com/git/git/blob/v2.17.1/Documentation/RelNotes/2.17.1.txt) にはセキュリティ・アップデートがある。

- [Announcing the May 2018 Git Security Vulnerability – Microsoft DevOps Blog](https://blogs.msdn.microsoft.com/devops/2018/05/29/announcing-the-may-2018-git-security-vulnerability/)

できるだけ早めのアップデートが推奨されているようだ。
詳しい情報は後日追記する。

アップデートは計画的に。

## Git for Windows 2.17.1 (2) がリリース

- [Release Git for Windows 2.17.1(2) · git-for-windows/git](https://github.com/git-for-windows/git/releases/tag/v2.17.1.windows.2)

{{% fig-quote-md title="Release Git for Windows 2.17.1(2)" link="https://github.com/git-for-windows/git/releases/tag/v2.17.1.windows.2" lang="en" %}}
- Comes with [Git Credential Manager v1.16.1](https://github.com/Microsoft/Git-Credential-Manager-for-Windows/releases/tag/v1.16.1).
- Comes with [Git LFS v2.4.2](https://github.com/git-lfs/git-lfs/releases/tag/v2.4.2).
{{% /fig-quote-md %}}

## ブックマーク

- [「Git」に複数の脆弱性、修正版が公開 ～Microsoftは「Git for Windows」の更新を推奨 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1124686.html)

[Git for Windows]: https://gitforwindows.org/
