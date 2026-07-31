+++
title = "Git Extensions に関する覚え書き"
date =  "2018-01-21T16:07:15+09:00"
update =  "2018-02-03T18:31:50+09:00"
description = "毎度のことながらインストールしたときのことを忘れているため，Git Extensions のインストールに絡むあれこれについて覚え書きを記しておく。"
image = "/images/attention/remark.jpg"
tags        = [ "git", "tools", "git-extensions", "conemu", "putty" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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

私は [Git] の GUI フロントエンドとして [Git Extensions] を愛用しているのだが，毎度のことながらインストールしたときのことを忘れているため，インストールに絡むあれこれについて覚え書きを記しておく。

## [Git Extensions] のインストール

現在 [Git Extensions] は以下のサイトで運用されている。

- [Git Extensions | Git Extensions is a graphical user interface for Git that allows you to control Git without using the commandline](http://gitextensions.github.io/)

[Git Extensions] は .NET Framework （4.6.1 以上）で動作するが Mono （5.0 以上）でも動くので Linux や macOS 等でも利用可能らしい。
Windows 用なら [Git Extensions] 本体の他に以下のツールが同梱されている[^inst1]。

[^inst1]: 最近の [Git Extensions] には [GCM (Git Credential Manager for Windows)](https://github.com/Microsoft/Git-Credential-Manager-for-Windows "Microsoft/Git-Credential-Manager-for-Windows: Secure Git credential storage for Windows with support for Visual Studio Team Services, GitHub, and Bitbucket multi-factor authentication.") は同梱されていない。 [GCM] は [Git for Windows] の方に同梱されているため， [Git for Windows] インストール時に [GCM] も有効にすること。

- [Git for Windows]
- [KDiff3]
- [PuTTY]
- [ConEmu]

このうち [Git for Windows] と [KDiff3] は SetupComplete ファイル[^sc1] にのみ同梱されているが，バージョンが古いため，別途インストールしてから（SetupComplete ファイルではなく） Setup ファイルでインストールしたほうがいいかもしれない。

[^sc1]: [Git Extensions] 2.50.02 であれば `GitExtensions-2.50.02-SetupComplete.msi` が SetupComplete ファイルである。

[PuTTY] を別途インストールしている場合は設定でそちらに差し替えできる。
[PuTTY] は時々セキュリティ・アップデートが行われるが [Git Extensions] 側で追従できないことも多いので（SSH クライアントとして [PuTTY] を使うのであれば）自前でインストールして運用した方がいいだろう。

[ConEmu] はそのままでOK。

## [Git Extensions] を日本語化したいけど...

最近の [Git Extensions] はインストールした状態では日本語にできない。
[リポジトリ](https://github.com/gitextensions/gitextensions)に日本語化モジュールはあるのだが Setup ファイルに同梱されていないようだ。

そこで無理矢理ではあるが，[リリースページ](https://github.com/gitextensions/gitextensions/releases/latest)から “Source Code” をダウンロードし，その中にある以下のファイルを [Git Extensions] インストール先フォルダ以下の `Translation` に入れてみる。

- `gitextensions-2.xx.xx/GitUI/Translation/Japanese.gif`
- `gitextensions-2.xx.xx/GitUI/Translation/Japanese.xlf`
- `gitextensions-2.xx.xx/GitUI/Translation/Japanese.Plugins.xlf`

これで言語を日本語にできる。
言語の変更は Settings ダイアログの以下の部分をクリックする。

{{< fig-img src="https://photo.baldanders.info/flickr/image/24939315857_m.png" title="Setting of Git Extensions" link="https://photo.baldanders.info/flickr/24939315857/" >}}

すると以下のウィンドウが表示されるので日本の国旗を選択すれば日本語に切り替わる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/39100077644_m.png" title="Choose language in Git Extensions Setting" link="https://photo.baldanders.info/flickr/39100077644/" >}}

やってみるとわかると思うが，けっこう英語のままである（そのせいで Setup ファイルに含まれないのかな）。
`Japanese.xlf` および `Japanese.Plugins.xlf` はテキストファイルなので，ここをいじれば日本語化できそうだ。
もし英語得手の方がいれば翻訳に貢献できるかもしれない。
私は... まぁ英語でも別に困らないか。

## ブックマーク

- [私家版 Git For Windowsのインストール手順 | OPC Diary](https://opcdiary.net/?page_id=27065)
- [KDiff3 の日本語化 – 各種パソコン技術情報](http://w3w.nnn2.com/?p=874)
- [しれっと、GitExtensionを日本語化する。 - Sinatraのあらかると](http://takkii.hatenablog.com/entry/2017/11/06/195156)

- [GnuPG for Windows : gpg-agent について]({{< ref "/openpgp/using-gnupg-for-windows-2.md" >}}) : SSH の鍵管理も gpg-agent で行う
- [Git Extensions 2.49 × ConEmu ＝ ♥]({{< ref "/remark/2016/11/git-extensions-2_49.md" >}})
- [Git Extensions と LibreOffice 6 と OpenPGP]({{< ref "/openpgp/git-extensions-and-libreoffice-6-with-openpgp.md" >}})

[Git]: https://git-scm.com/
[Git for Windows]: http://gitforwindows.org/
[KDiff3]: http://kdiff3.sourceforge.net/
[PuTTY]: https://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free SSH and Telnet client"
[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
[GCM]: https://github.com/Microsoft/Git-Credential-Manager-for-Windows "Git Credential Manager for Windows"
[Git Extensions]: http://gitextensions.github.io/ "Git Extensions | Git Extensions is a graphical user interface for Git that allows you to control Git without using the commandline"
