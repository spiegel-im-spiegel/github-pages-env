+++
title = "Ubuntu で遊ぶ"
date = "2019-03-31T10:14:36+09:00"
description = "この記事は覚え書きとして随時更新していく予定。"
image = "/images/attention/kitten.jpg"
tags = [ "virtualbox", "linux", "ubuntu", "security" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

さて[インストールも完了した]({{< relref "./the-window-the-window.md" >}} "Windows に Ubuntu を入れる")し，さっそく遊ぼうか。

なお，この記事は覚え書きとして随時更新していく予定なので悪しからず。

## [VirtualBox] との連携

### クリップボードの内容を共有する

[Ubuntu] の動作確認はテキストファイルでメモを取りながら行っているのだが，ホスト OS とゲスト OS との間でクリップボードが共有できないか調べてみたら簡単にできるようだ。

インストール後の [VirtualBox] マネージャは以下の通り。

{{< fig-img src="../the-window-the-window/end-of-setup.png" link="../the-window-the-window/end-of-setup.png" width="976" >}}

この画面から「設定」アイコンをクリックすると以下の画面が開く。

{{< fig-img src="./setting-clipboard.png" link="./setting-clipboard.png" width="831" >}}

この画面の「一般」カテゴリの「高度」タブに「クリップボードの共有」の項目がある。
ここを「双方向」にすればホスト OS とゲスト OS との間でクリップボードを共有できる。

## [Ubuntu] に関する雑多なこと

そもそも [Ubuntu] でターミナルってどうやって開くんだろうと思ったが，デスクトップを右クリックして「端末を開く」でよかった。

### Advanced Package Tool

[Ubuntu] は [Debian] 系のディストリビューションなのでパッケージやシステムの管理には APT (Advanced Package Tool) を使う。
`apt-get` 等が `apt` コマンドに統合されているとは知らなかったよ。
[Debian] 系は古いバージョンしか使ったことなかったからなぁ （笑）

- [第 6 章 メンテナンスと更新、APT ツール](https://debian-handbook.info/browse/ja-JP/stable/apt.html)
- [aptコマンドチートシート - Qiita](https://qiita.com/SUZUKI_Masaya/items/1fd9489e631c78e5b007)

パッケージリストを更新するには `apt update` コマンドを実行する。

```text
$ sudo apt update
[sudo] username のパスワード: 
ヒット:1 http://archive.ubuntulinux.jp/ubuntu cosmic InRelease
ヒット:2 http://jp.archive.ubuntu.com/ubuntu cosmic InRelease
取得:3 http://jp.archive.ubuntu.com/ubuntu cosmic-updates InRelease [88.7 kB]
ヒット:4 http://archive.ubuntulinux.jp/ubuntu-ja-non-free cosmic InRelease
取得:5 http://jp.archive.ubuntu.com/ubuntu cosmic-backports InRelease [74.6 kB]
取得:6 http://security.ubuntu.com/ubuntu cosmic-security InRelease [88.7 kB]
252 kB を 2秒 で取得しました (129 kB/s)
パッケージリストを読み込んでいます... 完了
依存関係ツリーを作成しています
状態情報を読み取っています... 完了
パッケージはすべて最新です。
```

パッケージリストを更新したらパッケージをアップグレードする。
パッケージをアップグレードするには `apt upgrade` コマンドを実行する。

```text
$ sudo apt upgrade
パッケージリストを読み込んでいます... 完了
依存関係ツリーを作成しています
状態情報を読み取っています... 完了
アップグレードパッケージを検出しています... 完了
以下のパッケージが自動でインストールされましたが、もう必要とされていません:
  libncursesw5 libtinfo5
これを削除するには 'sudo apt autoremove' を利用してください。
アップグレード: 0 個、新規インストール: 0 個、削除: 0 個、保留: 0 個。
```

おっと削除を推奨するパッケージがあるようだ。
`apt autoremove` で削除してしまおう。

```text
$ sudo apt autoremove
パッケージリストを読み込んでいます... 完了
依存関係ツリーを作成しています
状態情報を読み取っています... 完了
以下のパッケージは「削除」されます:
  libncursesw5 libtinfo5
アップグレード: 0 個、新規インストール: 0 個、削除: 2 個、保留: 0 個。
この操作後に 856 kB のディスク容量が解放されます。
続行しますか? [Y/n] y
(データベースを読み込んでいます ... 現在 167547 個のファイルとディレクトリがインストールされています。)
libncursesw5:amd64 (6.1+20180210-4ubuntu1) を削除しています ...
libtinfo5:amd64 (6.1+20180210-4ubuntu1) を削除しています ...
libc-bin (2.28-0ubuntu1) のトリガを処理しています ...
```

これでオッケ。

今回はなかったが保留パッケージを含めて更新する場合は `apt full-upgrade` を実行する。

### セキュリティ情報をチェックする

本家の [Debian] ディストリビューションのセキュリティ情報は以下から取得できる。

- [Debian -- Security Information](https://www.debian.org/security/)

[Ubuntu] の方は以下から。

- [Ubuntu security notices](https://usn.ubuntu.com/)

フィードも提供されているので巡回先に加えることにしよう。

### えっ ifconfig って入ってないの？

なにはともあれネットワーク設定を見ようと `ifconfig` コマンドを叩いてみたが

```text
$ ifconfig -a -v

Command 'ifconfig' not found, but can be installed with:

sudo apt install net-tools
```

とか言われる。
まじすか！

まぁ，なければ入れればいい話なので言われたとおりインストールする。

```text
$ sudo apt install net-tools
```

これで

```text
$ ifconfig -a -v
enp0s3: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 10.0.2.15  netmask 255.255.255.0  broadcast 10.0.2.255
        inet6 xxxx::xxxx:xxxx:xxxx:xxxx  prefixlen 64  scopeid 0x20<link>
        ether **:**:**:**:**:**  txqueuelen 1000  (イーサネット)
        RX packets 20  bytes 2764 (2.7 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 82  bytes 8466 (8.4 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (ローカルループバック)
        RX packets 151  bytes 11745 (11.7 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 151  bytes 11745 (11.7 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```

おー，動いた動いた。

ちなみに IP アドレスの `10.0.2.xx` は [VirtualBox] 側の仮想 DHCP ネットワークらしい。



[VirtualBox]: https://www.virtualbox.org/ "Oracle VM VirtualBox"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Debian]: https://www.debian.org/ "Debian -- The Universal Operating System"
