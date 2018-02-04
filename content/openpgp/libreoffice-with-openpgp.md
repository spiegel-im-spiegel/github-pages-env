+++
title = "LibreOffice と OpenPGP （仕切り直し）"
date = "2018-02-04T14:58:31+09:00"
description = "原因が分かったので仕切り直しの記事を書くことにする。"
image = "/images/attention/openpgp.png"
tags        = [ "tools", "libreoffice", "openpgp", "gnupg" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

「[Git Extensions と LibreOffice 6 と OpenPGP]({{< relref "openpgp/git-extensions-and-libreoffice-6-with-openpgp.md" >}})」で [LibreOffice] 文書に [OpenPGP] 署名できなくて「何がいけないんだろう。 誰かやり方を教えてください」と締めくくっていたが，原因が分かったので仕切り直しの記事を書くことにする。
情報を下さった方には感謝を捧げます。

## [OpenPGP] 署名について

まず [LibreOffice] 6.0 のリリースノートはこちら。

- [LibreOffice 6.0 リリースノート - The Document Foundation Wiki](https://wiki.documentfoundation.org/ReleaseNotes/6.0/ja)

Linux 版では [LibreOffice] 5.4 から [OpenPGP] 署名とその検証が可能だったようだが， 6.0 で Windows 版および macOS 版でも可能になった。
他のプラットフォーム（Android 版のビュアーとか）で対応しているかどうかは不明。

[LibreOffice] ドキュメントへのデジタル署名は X.509 鍵では以前から可能だった。
デジタル署名を付与することにより以下が担保される。

- 完全性（Integrity）： 改竄の有無
- 認証（Authentication）： なりすましの有無
- 否認防止（Non-repudiation）： 文書作成の否認を防止（完全性と認証により達成可能）

デジタル署名可能な [LibreOffice] ドキュメントとして以下を確認している。

- Writer （文書ドキュメント）
- Calc （表計算ドキュメント）
- Impress （プレゼンテーション）
- Draw （図形描画）
- Math （数式）

また [LibreOffice] では PDF ドキュメントへのデジタル署名も可能だが， [OpenPGP] 署名には対応していないようだ（X.509 のみ）。
残念（まぁ，他の PDF ツールが対応してないと意味がないし）。

この記事では以下の環境を前提に話をすすめる。

- Windows 7 Professional SP1 64ビット版（Win10 でも大丈夫っぽい）
- LibreOffice 6.0.0 Windows 64ビット版（[Portable 版](https://portableapps.com/apps/office/libreoffice_portable "LibreOffice Portable | PortableApps.com - Portable software for USB, portable and cloud drives")でもいけるらしい）
- [GnuPG] 2.2.x

理由はすぐに分かる。

## [GnuPG] のインストール

[OpenPGP] 署名を行うために [GnuPG] をインストールする。
インストールの手順については以下を参照のこと。

- [GnuPG for Windows インストール編]({{< relref "openpgp/using-gnupg-for-windows-1.md" >}})

そしてインストール先だが，**必ず既定の “`C:\Program Files (x86)\gnupg`” フォルダにインストールすること**。

どうやら [LibreOffice] は実行本体である `gpg.exe` の場所を “`C:\Program Files (x86)\gnupg\bin`” フォルダに決め打ちしているらしく，他のフォルダにインストールしても [LibreOffice] が認識しないのだ[^gpg1]。
ただ最近のインストーラはインストール先を指定する画面をすっ飛ばすみたい（？）なので，新規にインストールしている場合はあまり考えなくていいかもしれない。

[^gpg1]: いや，これダメだろ。何でこんなダサい作りになってるんだ？ 普通にインストールすれば PATH が張られるし，オプションに OpenPGP の項目があるんだから `gpg.exe` までのフルパスを指定できるようにすればいいぢゃん。

なお [LibreOffice] が推奨している [Gpg4win] は不要で [GnuPG] の最小構成のみインストールできていれば問題ない。
フロント・エンドの [Kleopatra] も不要[^k1]。

[^k1]: 色々試行錯誤して [Gpg4win] も [Kleopatra] も不要と分かったときには椅子から転げ落ちそうになったよ。何だったんだ[前回]のアレは `orz`

## [LibreOffice] ドキュメントへの [OpenPGP] 署名

まずメニューから「ファイル(F)」→「デジタル署名(Q)」→「デジタル署名(R)...」と辿っていく。

{{< fig-img src="https://farm5.staticflickr.com/4657/26181610768_acaa5d2e1c_o.png" title="LibreOffice: digital sign" link="https://www.flickr.com/photos/spiegel/26181610768/" >}}

すると以下のウィンドウが表示される。

{{< fig-img src="https://farm5.staticflickr.com/4627/40071712341_13c8f8c2b7.jpg" title="List of digital sigs (1)" link="https://www.flickr.com/photos/spiegel/40071712341/" >}}

ここで「ドキュメントに署名（D）...」ボタンを押すと署名可能な鍵のリストが表示される。

{{< fig-img src="https://farm5.staticflickr.com/4747/26198104318_ced43cb690.jpg" title="Select key" link="https://www.flickr.com/photos/spiegel/26198104318/" >}}

[OpenPGP] 鍵であれば「支払期日」（何で？）の項目が “OpenPGP” と表記されている筈である（X.509 鍵なら “X.509” と表記される）。
鍵を選択して「署名」ボタンを押すとデジタル署名が付与される[^pe1]。

[^pe1]: 必要に応じて `gpg-agent` が Pinentry を起動して PIN またはパスフレーズの入力を促す。

{{< fig-img src="https://farm5.staticflickr.com/4741/40071712521_f1dbbfb8c3.jpg" title="List of digital sigs (2)" link="https://www.flickr.com/photos/spiegel/40071712521/" >}}

デジタル署名は複数付与できる。
署名が完了したら「閉じる」ボタンを押して作業を終える。

正しいデジタル署名が付与されていれば，編集画面が以下のようになる筈である。

{{< fig-img src="https://farm5.staticflickr.com/4674/39172914465_05e85ab216.jpg" title="Valid sig" link="https://www.flickr.com/photos/spiegel/39172914465/" >}}

またこの状態で編集作業を行うと以下の表示になる。

{{< fig-img src="https://farm5.staticflickr.com/4724/40038567142_c99e4f0848.jpg" title="Editing" link="https://www.flickr.com/photos/spiegel/40038567142/" >}}

ドキュメントの保存を実行するとデジタル署名が削除されるのでご注意を。

## 保存時に [OpenPGP] 暗号化オプションがあるんだけど...

文書を「名前を付けて保存」すると以下のオプションが表示される。

{{< fig-img src="https://farm5.staticflickr.com/4650/40072008511_1b320717d5.jpg" title="Save as..." link="https://www.flickr.com/photos/spiegel/40072008511/" >}}

ここで「GPGキーで暗号化する」をチェックして保存すると暗号化可能な鍵のリストが表示されて暗号化できるっぽいんだけど，実際にやってみると

{{< fig-img src="https://farm5.staticflickr.com/4669/40072008591_44d719908c.jpg" title="Encryption error" link="https://www.flickr.com/photos/spiegel/40072008591/" >}}

となって上手くいかない。
自分の鍵でもダメだったので，まだ機能が組み込まれてないとか？

[OpenPGP] 鍵でドキュメントが暗号化できると取り回しが凄く楽になるんだけどねぇ。
メールにパスワード書くとかマヌケなことをしなくて良くなるし（笑）










## ブックマーク

- [GnuPG チートシート（鍵作成から失効まで）]({{< relref "openpgp/gnupg-cheat-sheet.md" >}})
- [OpenPGP 鍵管理に関する考察]({{< relref "openpgp/openpgp-key-management.md" >}})

[前回]: {{< relref "openpgp/git-extensions-and-libreoffice-6-with-openpgp.md" >}} "Git Extensions と LibreOffice 6 と OpenPGP"
[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[Kleopatra]: https://www.kde.org/applications/utilities/kleopatra/ "KDE - Kleopatra - Certificate Manager and Unified Crypto GUI"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"
