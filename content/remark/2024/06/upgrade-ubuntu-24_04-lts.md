+++
title = "Ubuntu 24.04 LTS へのアップグレード（追記あり）"
date =  "2024-06-09T16:28:10+09:00"
description = "サードパーティの APT リポジトリ定義を全部書き直した。あとは遅まきながら Fcitx 5 の導入"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools", "install", "ppa" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

すっかり忘れていたが [Ubuntu] 24.04 LTS がリリースされているのでアップグレードしなきゃ。

- [Ubuntu 24.04 LTS (Noble Numbat) Release Notes - Release - Ubuntu Community Hub](https://discourse.ubuntu.com/t/ubuntu-24-04-lts-noble-numbat-release-notes/39890)
- [Ubuntu 24.04 LTS リリース | Ubuntu Japanese Team](/News/ubuntu2404)
- [Ubuntu 24.04 LTS “Noble Numbat”のリリース | gihyo.jp](https://gihyo.jp/admin/clip/01/ubuntu-topics/202404/26)

本当は日本語 Remix が出るまで待つつもりだったのだが，今回は出そうもない（2024-06-14 追記：今回は[日本語 Remix は出ない](https://lists.ubuntu.com/archives/ubuntu-jp/2024-June/006668.html "[ubuntu-jp:6669] Ubuntu 24.04 LTSの日本語Remixについて")らしい）。

## 前準備

まずは現行版を最新化しておく。

```text
$ sudo apt update
$ sudo apt upgrade
$ sudo snap refresh
```
以下のページから ISO イメージをダウンロードして，転ばぬ先のブータブル USB を作成する。

- [Download Ubuntu Desktop | Ubuntu](https://ubuntu.com/download/desktop)

うへ。
6GB もあるのか。

## アップグレード

念のため諸々のバックアップをとって「ソフトウェアの更新」からアップグレードを起動する。

{{< fig-img src="./upgrade-ubuntu-0.png" title="ソフトウェアの更新" link="./upgrade-ubuntu-0.png" width="645" >}}

{{< fig-img src="./upgrade-ubuntu-1.png" title="リリースノート" link="./upgrade-ubuntu-1.png" width="628" >}}

では状況を始めよう。

{{< fig-img src="./upgrade-ubuntu-2.png" title="アップグレードを開始しますか？" link="./upgrade-ubuntu-2.png" width="1121" >}}

ん？ git を削除するの？ なんで？ まぁいいや。
とりあえずアップグレードを開始しよう。

{{< fig-img src="./upgrade-ubuntu-3.png" title="Upgrade to the Thunderbird snap" link="./upgrade-ubuntu-3.png" width="675" >}}

おっと。
そういえば 24.04 から Thunderbird は snap で提供されるんだった。

このあとは大きな問題もなくアップグレードは完了した。

{{< fig-img src="./upgrade-ubuntu-4.png" title="Upgrade to the Thunderbird snap" link="./upgrade-ubuntu-4.png" width="694" >}}

再起動して確認していこう。

## ようやく [GnuPG] が 2.4 系に

```text
$ gpg --version
gpg (GnuPG) 2.4.4
libgcrypt 1.10.3
Copyright (C) 2024 g10 Code GmbH
License GNU GPL-3.0-or-later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: /home/username/.gnupg
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256,
      TWOFISH, CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

ようやくかい。
まぁ年末には 2.2 系のサポートが切れるからな。
ちなみに 2024-06 時点の [GnuPG] 最新版は 2.4.5 である。
この辺もうちょっとなんとかならんもんかねぇ。

## [OpenSSH] と [OpenSSL]

```text
$ ssh -V
OpenSSH_9.6p1 Ubuntu-3ubuntu13, OpenSSL 3.0.13 30 Jan 2024
```

いつものことではあるが，これも微妙に古い。
どうしてセキュリティ製品のバージョンが古いのだろう。
妙なバックポートパッチじゃなくて公式の最新版を使えよ。

## APT (Advanced Package Tool)

```text
$ apt -v
apt 2.7.14 (amd64)
```

24.04 から APT の仕様がが変わった。

- [第812回 aptの新機能あれこれ ［Ubuntu 24.04 LTS版］ | gihyo.jp](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0812)

具体的には以下の通り。

1. rsa1024 の電子署名公開鍵を使っているとワーニングが出る
2. deb822 形式への移行

特に1番目については 2.8 以降は rsa1024 はワーニングではなくエラーになるらしい。
公式のリポジトリなら問題ないだろうがサードパーティのリポジトリを使う場合には問題になる。

### [PPA] 版 [Git] の再インストール

24.04 へのアップグレードで [PPA] 版 [git][Git] が削除されてしまったので改めてインストールする（公式リポジトリの [git][Git] はバージョンが古いため）。
まっさらな状態からインストールする場合は

```text
$ sudo add-apt-repository ppa:git-core/ppa
```

でリポジトリを登録できる。
また既にリポジトリを登録している場合は「ソフトウェアとアップデート」の「他のソフトウェア」タブを開くと

{{< fig-img src="./software-and-update-1.png" title="ソフトウェアとアップデート" link="./software-and-update-1.png" width="1052" >}}

[PPA] 版 [git][Git] が既に ready 状態になっているのでチェックを入れて有効にする。

#### ワーニングの回避 {#nowarn}

この状態で `apt update` すると

```text
$ sudo apt update
...
W: http://ppa.launchpad.net/git-core/ppa/ubuntu/dists/noble/InRelease: Signature by key E1DD270288B4E6030699E45FA1715D88E1DF1F24 uses weak algorithm (rsa1024)
```

などと最後にワーニングが出る。
[PPA] のサイトを見ると

{{< fig-quote type="markdown" title="Git stable releases : “Ubuntu Git Maintainers” team" link="https://launchpad.net/~git-core/+archive/ubuntu/ppa" lang="en" >}}
There’s nothing that PPA owners can do about this for now, as the key is fully controlled internally by Launchpad.
{{< /fig-quote >}}

とあり Launchpad 運営側で対処してくれないとどうにもならないようだ。
とりあえずワーニングを消したいのであれば

- [[Workaround] Apt Warning: Signature Key Uses Weak Algorithm | UbuntuHandbook](https://ubuntuhandbook.org/index.php/2024/04/workaround-apt-warning-signature-key-uses-weak-algorithm/)
- [Ubuntu 24.04: PPA: "Signature by key 98703123E0F52B2BE16D586EF13930B14BB9F05F uses weak algorithm (rsa1024)" · Issue #2845 · sabnzbd/sabnzbd · GitHub](https://github.com/sabnzbd/sabnzbd/issues/2845)

に書かれている方法で対処できる。
すなわち `/etc/apt/apt.conf.d/99weakkey-warning` を開いて（ない場合は作成）

```text
APT::Key::Assert-Pubkey-Algo ">=rsa1024";
```

と書き込めばよい。
これで rsa1024 鍵でもワーニングが出なくなる。
APT 2.8 になったらどうか分からんけど。

あとは `sudo apt install git` で再インストールできる。

### [KeePassXC] のアップデート

[KeePassXC] も [PPA] で提供されている。

```text
$ sudo add-apt-repository ppa:phoerious/keepassxc
```

既にリポジトリを登録している場合は「ソフトウェアとアップデート」の「他のソフトウェア」タブを開くと

{{< fig-img src="./software-and-update-2.png" title="ソフトウェアとアップデート" link="./software-and-update-2.png" width="1052" >}}

[KeePassXC] が既に ready 状態になっているのでチェックを入れて有効にする。
あとは普通に `sudo apt update & sudo apt upgrade` でアップデートできる。

### [NodeSource] 版 [Node.js] のアップデート

[NodeSource] 版 [node.js][Node.js] のインストール方法は以下を参照のこと。

- [NodeSource Node.js Binary Distributions](https://github.com/nodesource/distributions/blob/master/README.md)

上の方法でリポジトリを登録すると以下のファイルができてるはず。

- `/usr/share/keyrings/nodesource.gpg`
- `/etc/apt/sources.list.d/nodesource.list`

`nodesource.list` の内容が以下とすると

```text
deb [arch=amd64 signed-by=/usr/share/keyrings/nodesource.gpg] https://deb.nodesource.com/node_20.x nodistro main
```

新たに `/etc/apt/sources.list.d/nodesource.sources` ファイルを作成し，以下の内容とする。

```text
Types: deb
URIs: https://deb.nodesource.com/node_20.x
Suites: nodistro
Components: main
Signed-By: /usr/share/keyrings/nodesource.gpg
Architectures: amd64
```

`nodesource.sources` を作成したら `nodesource.list` ファイルは削除してよい。
あとは普通に `sudo apt update & sudo apt upgrade` でアップデートできる。

### [Docker] Engine のアップデート

- [Install Docker Engine on Ubuntu | Docker Docs](https://docs.docker.com/engine/install/ubuntu/)

既に以下のファイルがあるとする。

- `/etc/apt/keyrings/docker.asc`
- `/etc/apt/sources.list.d/docker.list`

また `docker.list` の内容は以下の通りとする。

```text
deb [arch=amd64 signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu noble stable
```

これに対し `/etc/apt/sources.list.d/docker.sources` ファイルを作成し，以下の内容とする。

```text
Types: deb
URIs: https://download.docker.com/linux/ubuntu
Suites: noble
Components: stable
Signed-By: /etc/apt/keyrings/docker.asc
Architectures: amd64
```

`docker.sources` を作成したら `docker.list` ファイルは削除してよい。
あとは普通に `sudo apt update & sudo apt upgrade` でアップデートできる。

### [pgAdmin 4] のアップデート

- [Download](https://www.pgadmin.org/download/pgadmin-4-apt/)

既に以下のファイルがあるとする。

- `/etc/apt/keyrings/packages-pgadmin-org.gpg`
- `/etc/apt/sources.list.d/pgadmin4.list`

また `pgadmin4.list` の内容は以下の通りとする。

```text
deb [signed-by=/usr/share/keyrings/packages-pgadmin-org.gpg] https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/noble pgadmin4 main
```

これに対し `/etc/apt/sources.list.d/pgadmin4.sources` ファイルを作成し，以下の内容とする。

```text
Types: deb
URIs: https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/noble
Suites: pgadmin4
Components: main
Signed-By: /etc/apt/keyrings/packages-pgadmin-org.gpg
```

`pgadmin4.sources` を作成したら `pgadmin4.list` ファイルは削除してよい。
あとは普通に `sudo apt update & sudo apt upgrade` でアップデートできる。

### [VS Code] のアップデート

- [Running Visual Studio Code on Linux](https://code.visualstudio.com/docs/setup/linux)

既に以下のファイルがあるとする。

- `/etc/apt/keyrings/packages.microsoft.gpg`
- `/etc/apt/sources.list.d/vscode.list`

また `vscode.list` の内容は以下の通りとする。

```text
deb [arch=amd64,arm64,armhf signed-by=/etc/apt/keyrings/packages.microsoft.gpg] https://packages.microsoft.com/repos/code stable main
```

これに対し `/etc/apt/sources.list.d/vscode.sources` ファイルを作成し，以下の内容とする。

```text
Types: deb
URIs: https://packages.microsoft.com/repos/code
Suites: stable
Components: main
Signed-By: /etc/apt/keyrings/packages.microsoft.gpg
Architectures: amd64,arm64,armhf
```

`vscode.sources` を作成したら `vscode.list` ファイルは削除してよい。
あとは普通に `sudo apt update & sudo apt upgrade` でアップデートできる。

### 【2024-06-14 追記】 Ubuntu 24.04 LTS 向け Ubuntu Japanese Team パッケージ {#jp}

- [[ubuntu-jp:6669] Ubuntu 24.04 LTSの日本語Remixについて](https://lists.ubuntu.com/archives/ubuntu-jp/2024-June/006668.html)

Ubuntu 24.04 LTS では日本語 Remix は出ないらしい。
一方で Ubuntu Japanese Team パッケージについては deb822 形式に対応したものを利用できる。

```text
$ sudo wget https://www.ubuntulinux.jp/sources.list.d/noble.sources -O /etc/apt/sources.list.d/ubuntu-ja.sources
$ sudo apt -U upgrade # update して upgrade する
```

まっさらから使う場合は `ubuntu-defaults-ja` をインストールする。

```text
$ sudo apt install ubuntu-defaults-ja
```

`/etc/apt/sources.list.d/ubuntu-ja.sources` ファイルに公開鍵も埋め込まれているが rsa1024 鍵なのでワーニングが出る。
とりあえずワーニングを回避したい場合は「[ワーニングの回避](#nowarn)」の節を参照のこと。

将来バージョンでは [PPA] のリポジトリを直接利用できるようになるらしい。

## [Fcitx 5] を導入する

今更だが input method を [Fcitx 5] に変更する。

- [第689回　Ubuntu 21.10でFcitx 5を使用する | gihyo.jp](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0689)
- [GUIでの日本語入力環境の整備（Ubuntu 24.04）](https://www.aise.ics.saitama-u.ac.jp/~gotoh/InputMethodOnUbuntu2404.html)

導入は APT で可能。

```text
$ sudo apt install fcitx5-mozc
$ im-config -n fcitx5
```

`~/.profile` ファイルに以下の記述を追記する。

```bash
while true; do
  dbus-update-activation-environment --systemd DBUS_SESSION_BUS_ADDRESS DISPLAY XAUTHORITY 2> /dev/null && break
done

export GTK_IM_MODULE=fcitx
export QT_IM_MODULE=fcitx
export XMODIFIERS="@im=fcitx"
export DefaultIMModule=fcitx5
if [ $SHLVL = 1 ] ; then
  (fcitx5 --disable=wayland -d --verbose '*'=0 &)
  xset -r 49  > /dev/null 2>&1
fi
```

再起動して `fcitx5` が起動していればOK。

```text
$ ps -ef | grep fcitx
username    2964       1  0 15:12 ?        00:00:00 fcitx5 --disable=wayland -d --verbose *=0
username    6384    4812  0 15:14 pts/1    00:00:00 jvgrep fcitx
```

全体設定は `fcitx5-configtool` を起動する（GUI 画面）。

{{< fig-img src="./fcitx5-configtool-1.png" title="fcitx5-configtool" link="./fcitx5-configtool-1.png" width="806" >}}

私は日本語キーボードを使ってるので，キーボードの設定を「キーボード - 日本語」に変更する。

{{< fig-img src="./fcitx5-configtool-2.png" title="fcitx5-configtool" link="./fcitx5-configtool-2.png" width="806" >}}

あとはお好みで。

## ブックマーク

- [第811回　ゴールデンウィーク特別企画 新学生・新社会人向けのUbuntuデスクトップ講座2024 | gihyo.jp](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0811)
- [Ubuntu 20.04にNode.jsをインストールする方法  | DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-20-04-ja)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: https://www.openssh.com/
[OpenSSL]: https://www.openssl.org/
[Git]: https://git-scm.com/
[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[NodeSource]: https://github.com/nodesource
[Node.js]: https://nodejs.org/
[Docker]: https://www.docker.com/ "Empowering App Development for Developers | Docker"
[pgAdmin 4]: https://www.pgadmin.org/ "pgAdmin - PostgreSQL Tools"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Fcitx 5]: https://fcitx-im.org/wiki/Fcitx_5

## 参考

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B01NBU1OS5" %}} <!-- シリコンパワー USBメモリ 32GB USB3.1 -->
