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

### gcc もないのかよ！

[Ubuntu] って gcc とかって後から入れるの？ 最近の UNIX 系ディストリビューションってそんな感じ？ まぁいいや。
ないなら入れればいいよね。

```text
$ sudo apt install build-essential
```

これで gcc と make 等の開発ツールチェーンが入った。

```text
$ gcc -v
Using built-in specs.
COLLECT_GCC=gcc
COLLECT_LTO_WRAPPER=/usr/lib/gcc/x86_64-linux-gnu/8/lto-wrapper
OFFLOAD_TARGET_NAMES=nvptx-none
OFFLOAD_TARGET_DEFAULT=1
Target: x86_64-linux-gnu
Configured with: ../src/configure -v --with-pkgversion='Ubuntu 8.2.0-7ubuntu1' --with-bugurl=file:///usr/share/doc/gcc-8/README.Bugs --enable-languages=c,ada,c++,go,brig,d,fortran,objc,obj-c++ --prefix=/usr --with-gcc-major-version-only --program-suffix=-8 --program-prefix=x86_64-linux-gnu- --enable-shared --enable-linker-build-id --libexecdir=/usr/lib --without-included-gettext --enable-threads=posix --libdir=/usr/lib --enable-nls --with-sysroot=/ --enable-clocale=gnu --enable-libstdcxx-debug --enable-libstdcxx-time=yes --with-default-libstdcxx-abi=new --enable-gnu-unique-object --disable-vtable-verify --enable-libmpx --enable-plugin --enable-default-pie --with-system-zlib --with-target-system-zlib --enable-objc-gc=auto --enable-multiarch --disable-werror --with-arch-32=i686 --with-abi=m64 --with-multilib-list=m32,m64,mx32 --enable-multilib --with-tune=generic --enable-offload-targets=nvptx-none --without-cuda-driver --enable-checking=release --build=x86_64-linux-gnu --host=x86_64-linux-gnu --target=x86_64-linux-gnu
Thread model: posix
gcc version 8.2.0 (Ubuntu 8.2.0-7ubuntu1) 

$ make -v
GNU Make 4.2.1
このプログラムは x86_64-pc-linux-gnu 用にビルドされました
Copyright (C) 1988-2016 Free Software Foundation, Inc.
ライセンス GPLv3+: GNU GPL バージョン 3 以降 <http://gnu.org/licenses/gpl.html>
これはフリーソフトウェアです: 自由に変更および配布できます.
法律の許す限り、　無保証　です.
```

### [KDiff3] の導入

[KDiff3] は普通に APT からインストールできる。

```text
$ sudo apt show kdiff3
Package: kdiff3
Version: 0.9.98-4
Priority: optional
Section: universe/kde
Origin: Ubuntu
```

よし，最新版があるな。
ではインストールしてしまおう。

```text
$ sudo apt install kdiff3
```

インストールが完了するとドックのアプリボタンで表示されるアプリ一覧にアイコンが追加される。

ちなみに `/usr/bin/kdiff3` にインストールされるので，たとえば `~/.gitconfig` に

```text
[merge]
    tool = kdiff3
[diff]
    guitool = kdiff3
[difftool "kdiff3"]
    path = /usr/bin/kdiff3
```

と設定しておけば [git] のマージ等で [KDiff3] が呼び出される。

[VirtualBox]: https://www.virtualbox.org/ "Oracle VM VirtualBox"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Debian]: https://www.debian.org/ "Debian -- The Universal Operating System"
[git]: https://git-scm.com/
[KDiff3]: http://kdiff3.sourceforge.net/

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/Ubuntu-18-04-Remix-%E4%BD%BF%E3%81%84%E6%96%B9%E3%81%8C%E5%85%A8%E9%83%A8%E3%82%8F%E3%81%8B%E3%82%8B%E6%9C%AC-%E6%97%A5%E7%B5%8CBP%E3%83%91%E3%82%BD%E3%82%B3%E3%83%B3%E3%83%99%E3%82%B9%E3%83%88%E3%83%A0%E3%83%83%E3%82%AF/dp/4296100742?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4296100742"><img src="https://images-fe.ssl-images-amazon.com/images/I/519EJD1-MXL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/Ubuntu-18-04-Remix-%E4%BD%BF%E3%81%84%E6%96%B9%E3%81%8C%E5%85%A8%E9%83%A8%E3%82%8F%E3%81%8B%E3%82%8B%E6%9C%AC-%E6%97%A5%E7%B5%8CBP%E3%83%91%E3%82%BD%E3%82%B3%E3%83%B3%E3%83%99%E3%82%B9%E3%83%88%E3%83%A0%E3%83%83%E3%82%AF/dp/4296100742?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4296100742">Ubuntu 18.04 LTS 日本語 Remix 使い方が全部わかる本 (日経BPパソコンベストムック)</a></dt>
    <dd>日経Linux (編集)</dd>
    <dd>日経BP社 2018-10-02 (Release 2018-10-02)</dd>
    <dd>Book ムック</dd>
    <dd>ASIN: 4296100742, EAN: 9784296100743</dd>
  </dl>
  <p class="description">Kindle 版もあるが紙のほうなら DVD が付録でついてくる。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-31">2019-03-31</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
