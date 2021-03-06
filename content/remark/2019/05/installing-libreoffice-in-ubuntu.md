+++
title = "Ubuntu に LibreOffice をインストールする3つの方法"
date =  "2019-05-17T00:32:45+09:00"
description = "結論から言うと LibreOffice の全ての機能を支障なく使いたいなら公式サイトのパッケージを使いなはれ，である。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "libreoffice", "openpgp", "ubuntu", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いやぁ，まいった。
私はただ LibreOffice の文書ファイルに[電子署名を付与]({{< ref "/openpgp/libreoffice-with-openpgp.md" >}} "LibreOffice と OpenPGP （仕切り直し）")したかっただけなのに，大事になってしまった。
忘れないうちにメモしておく。

[Ubuntu] に [LibreOffice] を導入する方法としては，主に以下の3つがある。

1. [APT] を使ってインストールする
2. [Snap] を使ってインストールする
3. [公式サイト]から deb パッケージ・ファイルを取ってきてインストールする。

これらはそれぞれ管理が独立しているため（あまりオススできないが）混在させることが可能である。

結論から言うと [LibreOffice] の全ての機能を支障なく使いたいなら[公式サイト]のパッケージを使いなはれ，である。

## [APT] を使ってインストールする

というか [Ubuntu] をインストールすると最初から入っている [LibreOffice] がこれである。
[Ubuntu] のアップデート・ポリシーに則って運用されるため保守的なバージョン管理になっている。
それでも安定的な運用を求めるのであれば，こちらを選択するのは悪くない。

[LibreOffice] 6.0 以降では文書ファイルにデジタル署名と暗号化を施すことができるのだが， [APT] 版は [OpenPGP] の鍵をうまく認識でいないようだ。
[Ubuntu] では「パスワードと鍵（Seahorse）」というツールで鍵やパスワードを統合的に管理しているのだが，これとの連携がうまく行っていないように見える。

ちなみに [APT] 版を削除するには

```text
$ sudo apt remove libreoffice*
```

でいけるようだ。
パッケージをワイルドカードで指定するのがポイントである。

## [Snap] を使ってインストールする

[Snap] は最近流行りのパッケージ管理ツールでディストリビューションから独立している点と，きめ細かいバージョン管理が可能である（パッケージによっては nightly ビルド等にも対応している）点が評価されている。
[APT] のようにローカルにデータベースを保持していないのでシンプルに運用できる。
[Ubuntu] には [Snap] が既定で導入されている。

[Snap] 版をインストールするには

```text
$ sudo snap install libreoffice
```

とすればよい。
削除するには

```text
$ sudo snap remove libreoffice
```

で OK。
ワイルドカード指定は不要である。

こちらも OpenPGP 鍵を認識できず，電子署名ができなかった。
[APT] 版も [Snap] 版も Canonical が運用・管理しているらしいのだが，セキュリティ方面には関心が薄いのだろう。
残念な話である。

## [公式サイト]から deb パッケージ・ファイルを取ってきてインストールする

[LibreOffice] の公式サイトでは deb パッケージ形式でインストールファイルを公開していて，これをダウンロードしてインストールできる。

まず[公式サイト]から deb ファイルが同梱された `*.tar.gz` ファイルをダウンロードする。
たとえばこんなファイル名になっている。

- `LibreOffice_7.1.0_Linux_x86-64_deb.tar.gz`

このファイルを適当な場所で展開すると `LibreOffice_7.1.0_Linux_x86-64_deb/` ディレクトリが作成される。
`LibreOffice_7.1.0_Linux_x86-64_deb/DEBS/` ディレクトリに降りると複数の deb ファイルがあることが分かるだろう。
基本的にはこれを全てインストールする。
全てのファイルに対していちいち `gdebi` コマンドを使うのはかったるいので

```text
$ sudo apt install ./*.deb
```

で一気にインストールしてしまおう。
てか，この手が使えたのか。

これで導入した [LibreOffice] では英語とフランス語くらいしか対応していないため日本語パックとヘルプも導入する。
同じく[公式サイト]から以下の `*.tar.gz` ファイルをダウンロードする。

- `LibreOffice_7.1.0_Linux_x86-64_deb_langpack_ja.tar.gz`
- `LibreOffice_7.1.0_Linux_x86-64_deb_helppack_ja.tar.gz`

これらのファイルを適当な場所で展開して中に入っている deb ファイルを片っ端からインストールしていく。

削除する場合は

```text
$ sudo apt remove libreoffice7.1* libobasis7.1*
```

のようにバージョンとワイルドカードを指定すること。

分かってますよ。
めがっさ面倒くさいです。
しかしこの公式版ならちゃんと [OpenPGP] 鍵を認識して電子署名も暗号化もできる。

日本でもメールにパスワード書くような悪習は早くなくしていきたいものである。

## ブックマーク

- [第507回　さまざまなLibreOfficeのインストール方法：Ubuntu Weekly Recipe｜gihyo.jp … 技術評論社](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0507)

- [LibreOffice と OpenPGP （仕切り直し）]({{< ref "/openpgp/libreoffice-with-openpgp.md" >}}) : Windows 版の [LibreOffice] 6.0 を対象にしているが機能は同じなので適宜読み替えてください

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[APT]: https://debian-handbook.info/browse/ja-JP/stable/apt.html "第 6 章 メンテナンスと更新、APT ツール"
[Snap]: https://github.com/snapcore/snapd "snapcore/snapd: The snapd and snap tools enable systems to work with .snap files."
[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[公式サイト]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[OpenPGP]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
