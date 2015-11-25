+++
date = "2015-11-25T21:49:56+09:00"
description = "今回の Gpg4win のリリースはセキュリティ・アップデートを含み， Advisory も併せてリリースされている。"
draft = false
tags = ["security", "vulnerability", "gnupg", "windows", "cryptography", "openpgp", "tools"]
title = "Security Advisory Gpg4win"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

Windows 用の [GnuPG] を含む暗号ツール群 [Gpg4win] の 2.3.0 がリリースされた。

- [Gpg4win-announce: Gpg4win 2.3.0 released](http://lists.wald.intevation.org/pipermail/gpg4win-announce/2015-November/000067.html)

今回のリリースはセキュリティ・アップデートを含み， Advisory も併せてリリースされている。

- [[Gpg4win-announce] Security Advisory Ggp4win 2015-11-25](http://lists.wald.intevation.org/pipermail/gpg4win-announce/2015-November/000066.html)

内容は以下のとおり。

1. The installer will load and execute other code if it is placed in the same directory as a DLL with the right name.  This "current directory attack" or "dll preloading attack" can be part of a remote exploitation for example if the Gpg4win installer is downloaded to a common Downloads directory and the attacker can previously place files there by tricking a user or other software to download files with a specific name to the same place. If the Gpg4win installer is then executed, the other code may run, while the user believes to run only the Gpg4win installer.
2. There is a "local privilege escalation" during an installer run.  Installer runs can happen during a fresh, an update install or a deinstallation. With Windows Vista or later an administrator can log in as user and give higher privileges to a single process using the User Account Control mechanism (UAC). If the installer is started in this way, there is a time window where an attacker running with user privileges can insert code in a temporary directory of the installer that will be executed with the higher privileges bypassing the UAC.

まぁ Windows ではありがちな脆弱性である。
GnuPG 自体に問題があるわけではなく，[インストーラのバグ](http://sourceforge.net/p/nsis/bugs/1125/ "NSIS: Nullsoft Scriptable Install System / Bugs / #1125 Code execution / Privilege escalation problems with NSIS installers")っぽい。

世の中は[10人にひとりもまともに OpenPGP 製品を使えない]({{< relref "remark/2015/use-the-signal-luke.md" >}})らしいのに [Gpg4win] を使ってる人がどの程度いるのかかなり怪しいのだが，まぁ一応。

影響度はよく分からないが， DLL プリロード攻撃の脆弱性も権限昇格の脆弱性も割とヤバめなので，さっさとアップデートすることをおすすめする。
ひょっとして，これ他の Windows 向けオンラインソフトにも影響出るかなぁ。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"
