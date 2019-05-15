+++
title = "ATOM v1.37 がリリースされた"
date =  "2019-05-15T21:36:43+09:00"
description = "Ubuntu 環境では ATOM を自動更新してくれないようだ。"
image = "/images/attention/tools.png"
tags = ["atom", "editor", "ubuntu", "install"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[ATOM] エディタ v.1.37.0 がリリースされている。

- [Release 1.37.0 · atom/atom · GitHub](https://github.com/atom/atom/releases/tag/v1.37.0)

ところで [Ubuntu] 環境では [ATOM] を自動更新してくれないようで， “About Atom” の画面に以下のメッセージが表示されている。

{{< fig-img src="./about-atom.png" link="./about-atom.png" width="572" >}}

[以前紹介した]({{< ref "/remark/2019/04/atom-in-ubuntu.md" >}} "Ubuntu に ATOM エディタを導入する")ように deb ファイルをダウンロードして gdebi コマンドで上書きインストールしてもいいのだが，今回は以下のページで紹介されている方法を試してみよう。

- [Installing Atom](https://flight-manual.atom.io/getting-started/sections/installing-atom/)

まず，以前に gdebi コマンドでインストールしていたものは念の為に削除しておこう。

```text
$ sudo apt remove atom
```

それから [ATOM] 用のリポジトリを登録する。

最初に [ATOM] 用のリポジトリのための公開鍵をインポートする。

```text
$ wget -qO - https://packagecloud.io/AtomEditor/atom/gpgkey | sudo apt-key add -
```

次に `/etc/apt/sources.list.d/` ディレクトリに [ATOM] 用のリポジトリを定義した `atom.list` ファイルを作成する。

```text
$ sudo sh -c 'echo "deb [arch=amd64] https://packagecloud.io/AtomEditor/atom/any/ any main" > /etc/apt/sources.list.d/atom.list'
```

これでリポジトリの登録が完了する。
APT のデータベースをアップデートしておこう。

```text
$ sudo apt update
```

[ATOM] の最新版がインストールできるか確認しておこう。

```text
$ apt show atom
Package: atom
Version: 1.37.0
Priority: optional
Section: devel
Maintainer: GitHub <atom@github.com>
...
```

よしよし。
これで [ATOM] の最新版がインストールできる。

```text
$ sudo apt install atom
```

ちなみにベータ版は

```text
$ sudo apt install atom-beta
```

でインストールできるらしい。
いや，入れないよ。

[ATOM]: https://atom.io/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
