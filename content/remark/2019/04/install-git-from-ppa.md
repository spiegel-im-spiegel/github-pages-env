+++
title = "PPA から Git をインストールする"
date = "2019-04-08T00:12:49+09:00"
description = "Git に関しては PPA (Personal Package Archive) で最新版のビルドを提供されている方がいるらしい。ありがたや。"
image = "/images/attention/kitten.jpg"
tags = [ "linux", "ubuntu", "ppa", "git", "ssh", "credential" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]で [OpenSSH] と [GnuPG] を連携させたのでようやく [git] で遊べるようになった。

しかし，実際にコマンドを叩いてみたら

```text
$ git version

Command 'git' not found, but can be installed with:

sudo apt install git
```

またかよ。
嫌な予感しかしない。

```text
$ sudo apt show git
Package: git
Version: 1:2.19.1-1ubuntu1.1
Priority: optional
Section: vcs
Origin: Ubuntu
...
```

やっぱり。
2.21.0 が出たのってだいぶ前だよな。

いろいろ調べてみたら [git] に関しては [PPA (Personal Package Archive)](https://launchpad.net/ubuntu/+ppas) で最新版のビルドを提供されている方がいるらしい。
ありがたや。

- [Git stable releases : “Ubuntu Git Maintainers” team](https://launchpad.net/~git-core/+archive/ubuntu/ppa)

何故か `add-apt-repository` は既に入ってるぽかったので，まずはリポジトリを導入しよう。

```text
$ sudo add-apt-repository ppa:git-core/ppa
 The most current stable version of Git for Ubuntu.

For release candidates, go to https://launchpad.net/~git-core/+archive/candidate .
 詳しい情報: https://launchpad.net/~git-core/+archive/ubuntu/ppa
[ENTER] を押すと続行します。Ctrl-c で追加をキャンセルできます。

ヒット:1 http://jp.archive.ubuntu.com/ubuntu cosmic InRelease
取得:2 http://jp.archive.ubuntu.com/ubuntu cosmic-updates InRelease [88.7 kB]
ヒット:3 http://archive.ubuntulinux.jp/ubuntu cosmic InRelease
取得:4 http://jp.archive.ubuntu.com/ubuntu cosmic-backports InRelease [74.6 kB]
ヒット:5 http://archive.ubuntulinux.jp/ubuntu-ja-non-free cosmic InRelease
取得:6 http://security.ubuntu.com/ubuntu cosmic-security InRelease [88.7 kB]
取得:7 http://ppa.launchpad.net/git-core/ppa/ubuntu cosmic InRelease [20.7 kB]
取得:8 http://ppa.launchpad.net/git-core/ppa/ubuntu cosmic/main i386 Packages [3,032 B]
取得:9 http://ppa.launchpad.net/git-core/ppa/ubuntu cosmic/main amd64 Packages [3,032 B]
取得:10 http://ppa.launchpad.net/git-core/ppa/ubuntu cosmic/main Translation-en [2,248 B]
281 kB を 3秒 で取得しました (81.9 kB/s)
パッケージリストを読み込んでいます... 完了

$ sudo apt update
ヒット:1 http://archive.ubuntulinux.jp/ubuntu cosmic InRelease
ヒット:2 http://jp.archive.ubuntu.com/ubuntu cosmic InRelease
ヒット:3 http://archive.ubuntulinux.jp/ubuntu-ja-non-free cosmic InRelease
取得:4 http://jp.archive.ubuntu.com/ubuntu cosmic-updates InRelease [88.7 kB]
取得:5 http://jp.archive.ubuntu.com/ubuntu cosmic-backports InRelease [74.6 kB]
取得:6 http://security.ubuntu.com/ubuntu cosmic-security InRelease [88.7 kB]
ヒット:7 http://ppa.launchpad.net/git-core/ppa/ubuntu cosmic InRelease
252 kB を 2秒 で取得しました (155 kB/s)
パッケージリストを読み込んでいます... 完了
依存関係ツリーを作成しています
状態情報を読み取っています... 完了
パッケージはすべて最新です。

$ sudo apt upgrade
パッケージリストを読み込んでいます... 完了
依存関係ツリーを作成しています
状態情報を読み取っています... 完了
アップグレードパッケージを検出しています... 完了
アップグレード: 0 個、新規インストール: 0 個、削除: 0 個、保留: 0 個。

$ sudo apt show git
Package: git
Version: 1:2.21.0-0ppa1~ubuntu18.10.1
Priority: optional
Section: vcs
```

よし。
最新版が入ってるな。
インストールっと。

```text
$ sudo apt install git
```

これでインストール完了。
動作確認してみる。

```text
$ git version
git version 2.21.0
```

よしよし。
おっと，初期設定しないと。

```text
$ git config --global user.name "Alice"
$ git config --global user.email "alice@example.com"
$ git config --global user.signingkey 697CDD8A
$ git config --global commit.gpgsign true
```

## SSH 経由で git clone する

まずは適当なリポジトリを SSH 経由で git clone してみる。

```text
$ git clone git@github.com:spiegel-im-spiegel/gocli.git github.com/spiegel-im-spiegel/gocli
Cloning into 'github.com/spiegel-im-spiegel/gocli'...
The authenticity of host 'github.com (192.30.255.112)' can't be established.
RSA key fingerprint is SHA256:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added 'github.com,192.30.255.112' (RSA) to the list of known hosts.
remote: Enumerating objects: 71, done.
remote: Counting objects: 100% (71/71), done.
remote: Compressing objects: 100% (52/52), done.
remote: Total 246 (delta 31), reused 53 (delta 16), pack-reused 175
Receiving objects: 100% (246/246), 53.99 KiB | 431.00 KiB/s, done.
Resolving deltas: 100% (107/107), done.
```

[前回]で [OpenSSH] と [GnuPG] を連携させたので認証時に Pinentry が起動する。

{{< fig-img src="../move-gpg-keyring/ssh-login.png" link="../move-gpg-keyring/ssh-login.png" >}}

よーし，うむうむ，よーし。

## git-credential による認証管理

SSH 経由でリポジトリにアクセスする場合はこれでいいのだが HTTPS 経由でアクセスする場合はアクセスするたびに毎回パスワードを訊かれるため鬱陶しい。
この場合は git-credential による認証管理を行うのがいいようだ。

[Ubuntu] では gnome-keyring を使った認証管理が使える。

- [shugo/git-credential-gnomekeyring: A git credential helper for GNOME keyring](https://github.com/shugo/git-credential-gnomekeyring)

これによると

{{< fig-quote title="shugo/git-credential-gnomekeyring" link="https://github.com/shugo/git-credential-gnomekeyring" lang="en" >}}
<q>There is a git package that includes all of contrib sources in /usr/share/doc/git/contrib/. Users who want to use git-credential-gnomekeyring should copy source from /usr/share/doc/git/contrib/credential/gnome-keyring to their working directory, then users can make and install it.</q>
{{< /fig-quote >}}

ということらしい。
さっそくやってみよう。

```text
$ cp -r /usr/share/doc/git/contrib/credential/gnome-keyring ~/work
$ cd ~/work/gnome-keyring
$ make
gcc -c -g -O2 -Wall -o git-credential-gnome-keyring.o git-credential-gnome-keyring.c
git-credential-gnome-keyring.c:28:10: fatal error: glib.h: そのようなファイルやディレクトリはありません
```

まじすか。
まぁ，入れればいいか。

```text
$ sudo apt install libglib2.0-dev
```

では再開。

```text
$ make
gcc -c -g -O2 -Wall -I/usr/include/glib-2.0 -I/usr/lib/x86_64-linux-gnu/glib-2.0/include -o git-credential-gnome-keyring.o git-credential-gnome-keyring.c
git-credential-gnome-keyring.c:29:10: fatal error: gnome-keyring.h: そのようなファイルやディレクトリはありません
```

ええい。
これも入れればいいのね。

```text
$ sudo apt install libgnome-keyring-dev
```

```text
$ make
gcc -g -O2 -Wall  -I/usr/include/gnome-keyring-1 -I/usr/include/glib-2.0 -I/usr/lib/x86_64-linux-gnu/glib-2.0/include -o git-credential-gnome-keyring.o -c git-credential-gnome-keyring.c
```

ワーニングいっぱい出たけどようやく通ったよ。
これで作成された `git-credential-gnome-keyring` をパスの通ったディレクトリに入れれば完了。
[git] 側が認識しているか確認してみよう。

```text
$ git help -a | grep credential-
   credential-cache     Helper to temporarily store passwords in memory
   credential-store     Helper to store credentials on disk
   credential-gnome-keyring
```

よしよし。
認識しているな。

これで [git] 設定に以下を加えれば完了となる。

```text
$ git config --global credential.helper gnome-keyring
```

適当なリポジトリで確かめてみよう。

```text
$ git clone https://github.com/spiegel-im-spiegel/private-repos.git aaa
Cloning into 'aaa'...
Username for 'https://github.com': alice@example.com
Password for 'https://alice@example.com@github.com': 
remote: Enumerating objects: 51, done.
remote: Counting objects: 100% (51/51), done.
remote: Compressing objects: 100% (41/41), done.
remote: Total 51 (delta 8), reused 47 (delta 8), pack-reused 0
Unpacking objects: 100% (51/51), done.

$ git clone https://github.com/spiegel-im-spiegel/private-repos.git bbb
Cloning into 'bbb'...
remote: Enumerating objects: 51, done.
remote: Counting objects: 100% (51/51), done.
remote: Compressing objects: 100% (41/41), done.
remote: Total 51 (delta 8), reused 47 (delta 8), pack-reused 0
Unpacking objects: 100% (51/51), done.
```

よし。
2回目からは訊いてこないな。

ちなみにこのときのパスワードは Web ページにサインインするときのパスワードじゃなくて，設定画面で振り出した personal access token を使うのでご注意を（つか，私がすっかり忘れててハマったのだが`w`）。

## ブックマーク

- [UbuntuのPPAて何？ [Linuxの使い方] All About](https://allabout.co.jp/gm/gc/438675/)
- [ubuntuのapt-getで最新版のgitをインストールする方法 - spangled shalalala blog](http://spangled-shalalala.hatenablog.com/entry/2017/09/05/060106)
- [Git - Gitのインストール](https://git-scm.com/book/ja/v1/%E4%BD%BF%E3%81%84%E5%A7%8B%E3%82%81%E3%82%8B-Git%E3%81%AE%E3%82%A4%E3%83%B3%E3%82%B9%E3%83%88%E3%83%BC%E3%83%AB)

- [Git Commit で OpenPGP 署名を行う]({{< ref "/openpgp/git-commit-with-openpgp-signature.md" >}})

[前回]: {{< relref "./move-gpg-keyring.md" >}} "Windows 環境で作った GnuPG の鍵束を Ubuntu に移行する"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[git]: https://git-scm.com/
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
