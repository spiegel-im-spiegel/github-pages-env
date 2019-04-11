+++
title = "Git Extensions と LibreOffice 6 と OpenPGP"
date =  "2018-02-03T18:18:31+09:00"
update =  "2018-02-04T15:05:38+09:00"
description = "LibreOffice での OpenPGP 署名が上手くできないんだけど！"
image = "/images/attention/openpgp.png"
tags        = [ "tools", "git-extensions", "libreoffice", "openpgp", "gnupg" ]

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

## [Git Extensions] 2.51 と [OpenPGP] 署名

先日 [Git Extensions] 2.51 がリリースされた。

- [Release Version 2.51 · gitextensions/gitextensions · GitHub](https://github.com/gitextensions/gitextensions/releases/tag/v2.51)

2.51 では色々と機能追加・改善があるが，個人的な目玉は [OpenPGP] 署名に対応したことだ。
メインのリポジトリ・ウィンドウでは “GPG” タブが追加され，コミットの電子署名の検証ができるようになった。
こんな感じ[^v1]。

[^v1]: 一部文字化けなのはご容赦。文字化けしてるのは時差表示の部分なのだが，うちの Windows 環境ではこの部分を Shift-JIS で吐き出すので文字化けしてしまうのだ。 `display-charset` オプションを弄ってもダメぽい。コマンドプロンプトや NYAGOS で使うぶんには問題ないんだけどねぇ。

{{< fig-img src="https://photo.baldanders.info/flickr/image/39343646184_m.png" title="GitEx: gpg validation" link="https://photo.baldanders.info/flickr/39343646184/" >}}

“GPG” タブを表示させるには設定で「GPG 情報を表示」にチェックを入れる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/25182287037_m.png" title="GitEx: setting" link="https://photo.baldanders.info/flickr/25182287037/" >}}

更にコミットやタグ作成時に電子署名を付与することもできる。
コミット時はこんな感じ。

{{< fig-img src="https://photo.baldanders.info/flickr/image/40021720412_m.png" title="GitEx: commit" link="https://photo.baldanders.info/flickr/40021720412/" >}}

タグ作成時はこんな感じ。

{{< fig-img src="https://photo.baldanders.info/flickr/image/25182287227_m.png" title="GitEx: create tag" link="https://photo.baldanders.info/flickr/25182287227/" >}}

[Git] と [GnuPG] を連携する方法については以下を参照のこと。

- [Git Commit で OpenPGP 署名を行う]({{< ref "/openpgp/git-commit-with-openpgp-signature.md" >}})

## [LibreOffice] 6 と [OpenPGP] 署名

[LibreOffice] 6 がリリースされた。

- [LibreOffice 6.0 リリースノート - The Document Foundation Wiki](https://wiki.documentfoundation.org/ReleaseNotes/6.0/ja)
- [「LibreOffice 6.0」が公開 ～多数の改善を盛り込んだメジャーバージョンアップ - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1104230.html)

これも色々と機能追加・改善があるが，この記事ではやはり OpenPGP 署名について。
Linux 版では以前からできてたそうだがバージョン 6 で Windows や macOS でもできるようになった。

{{< fig-img src="https://photo.baldanders.info/flickr/image/26181610768_m.png" title="LibreOffice: digital sign" link="https://photo.baldanders.info/flickr/26181610768/" >}}

つか，うまくいかないんだけど！

どうも [GnuPG] 本体だけではダメみたいで [Kleopatra] などのフロントエンド・ツールが必要なようだ。
Windows 環境では [Gpg4win] を導入することで関連ツールを一度にインストールできる（macOS では [GPG Suite](https://gpgtools.org/) を導入すればいいらしい）。
GnuPG 側と上手く連携できていれば「証明書の選択」ダイアログに電子署名可能な鍵が表示されるらしい。

でもダメなんだよねー。

{{< fig-img src="https://photo.baldanders.info/flickr/image/40055187561_m.png" title="Empty" link="https://photo.baldanders.info/flickr/40055187561/" >}}

何がいけないんだろう。
誰かやり方を教えてください。
はっきり言って [Gpg4win] 要らないので使えないなら削除したい。

{{% md-box %}}
**【2018-02-04 追記】** 原因が分かったので仕切り直した。

- [LibreOffice と OpenPGP （仕切り直し）]({{< ref "/openpgp/libreoffice-with-openpgp.md" >}})
{{% /md-box %}}

## ブックマーク

- [Git Extensions に関する覚え書き]({{< ref "/remark/2018/01/git-extensions-in-japanese.md" >}})

[Git Extensions]: http://gitextensions.github.io/ "Git Extensions | Git Extensions is a graphical user interface for Git that allows you to control Git without using the commandline"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Git]: https://git-scm.com/ "Git"
[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[Kleopatra]: https://www.kde.org/applications/utilities/kleopatra/ "KDE - Kleopatra - Certificate Manager and Unified Crypto GUI"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"
