+++
title = "Ubuntu で遊ぶ"
date = "2019-03-31T10:14:36+09:00"
description = "この記事は覚え書きとして随時更新していく予定。"
image = "/images/attention/kitten.jpg"
tags = [ "virtualbox", "linux", "ubuntu", "security", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

さて[インストールも完了した]({{< relref "./the-window-the-window.md" >}} "Windows に Ubuntu を入れる")し，さっそく遊ぼうか。

なお，この記事は覚え書きとして随時更新していく予定なので悪しからず。

- [VirtualBox との連携]({{< relref "#vb" >}})
    - [クリップボードの内容を共有する]({{< relref "#clipboard" >}})
- [Ubuntu に関する雑多なこと]({{< relref "#ubuntu" >}})
    - [Advanced Package Tool]({{< relref "#apt" >}})
    - [セキュリティ情報をチェックする]({{< relref "#secinfo" >}})
    - [えっ ifconfig って入ってないの？]({{< relref "#ifconfig" >}})
    - [gcc もないのかよ！]({{< relref "#gcc" >}})
    - [KDiff3 の導入]({{< relref "#kdiff3" >}})
    - [Go コンパイラを導入する]({{< relref "#golang" >}})

## VirtualBox との連携{#vb}

### クリップボードの内容を共有する{#clipboard}

[Ubuntu] の動作確認はテキストファイルでメモを取りながら行っているのだが，ホスト OS とゲスト OS との間でクリップボードが共有できないか調べてみたら簡単にできるようだ。

インストール後の [VirtualBox] マネージャは以下の通り。

{{< fig-img src="../the-window-the-window/end-of-setup.png" link="../the-window-the-window/end-of-setup.png" width="976" >}}

この画面から「設定」アイコンをクリックすると以下の画面が開く。

{{< fig-img src="./setting-clipboard.png" link="./setting-clipboard.png" width="831" >}}

この画面の「一般」カテゴリの「高度」タブに「クリップボードの共有」の項目がある。
ここを「双方向」にすればホスト OS とゲスト OS との間でクリップボードを共有できる。

## Ubuntu に関する雑多なこと{#ubuntu}

そもそも [Ubuntu] でターミナルってどうやって開くんだろうと思ったが，デスクトップを右クリックして「端末を開く」でよかった。

### Advanced Package Tool{#apt}

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

### セキュリティ情報をチェックする{#secinfo}

本家の [Debian] ディストリビューションのセキュリティ情報は以下から取得できる。

- [Debian -- Security Information](https://www.debian.org/security/)

[Ubuntu] の方は以下から。

- [Ubuntu security notices](https://usn.ubuntu.com/)

フィードも提供されているので巡回先に加えることにしよう。

### えっ ifconfig って入ってないの？{#ifconfig}

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

### gcc もないのかよ！{#gcc}

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

### KDiff3 の導入{#kdiff3}

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

### Go コンパイラを導入する{#golang}

APT で [Go 言語]コンパイラを入れれるかなぁ，と思ったが

```text
$ sudo apt show golang
[sudo] spiegel のパスワード: 
Package: golang
Version: 2:1.10~4ubuntu1
Priority: optional
Section: devel
Source: golang-defaults
Origin: Ubuntu
...
```

なんだこれ。
2世代も前じゃん。
やる気がなさすぎる `orz`

[Go] コンパイラは APT を使わず自前で入れたほうがいいようだ。
それなら `/usr/local` とかじゃなくてユーザのホームディレクトリ下に入れようか。
私は標準ライブラリのソースコードも頻繁に見るし，そのほうがいいだろう。

```text
$ mkdir ~/local
$ cd local
$ curl https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz -O

Command 'curl' not found, but can be installed with:

sudo apt install curl
```

まじすか。
[curl] が入ってないとか。
しょうがない入れるか。

```text
$ sudo apt install curl

$ curl -V
curl 7.61.0 (x86_64-pc-linux-gnu) libcurl/7.61.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.5 libpsl/0.20.2 (+libidn2/2.0.4) nghttp2/1.32.1 librtmp/2.3
Release-Date: 2018-07-11
Protocols: dict file ftp ftps gopher http https imap imaps ldap ldaps pop3 pop3s rtmp rtsp smb smbs smtp smtps telnet tftp 
Features: AsynchDNS IDN IPv6 Largefile GSS-API Kerberos SPNEGO NTLM NTLM_WB SSL libz TLS-SRP HTTP2 UnixSockets HTTPS-proxy PSL 
```

うーん？ 微妙にバージョンが古い気がするが... まぁいいか。

では気を取り直して。

```text
$ curl https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz -O
$ tar -xvf go1.12.4.linux-amd64.tar.gz
$ cd go/bin
$ ./go version
go version go1.12.4 linux/amd64
```

よし。
ちゃんと動く。

`go env` で環境周りを見てみると

```text
$ go env
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/username/.cache/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/username/go"
GOPROXY=""
GORACE=""
GOROOT="/home/username/local/go"
...
```

というわけで環境変数 `GOPATH` には `~/go` ディレクトリがセットされているようだ。
設定はこのまま使うとして `~/.profile` に以下の記述を追加しておく。

```text
# set PATH so it includes user's private bin if it exists
if [ -d "$HOME/local/go/bin" ] ; then
    PATH="$PATH:$HOME/local/go/bin:$HOME/go/bin"
    export GO111MODULE=on
fi
```

これで再ログインすれば

```text
$ go version
go version go1.12.4 linux/amd64

$ env | grep GO111MODULE
GO111MODULE=on
```

と設定が反映されているのが分かる。

一応動作確認をしておこう。

```text
$ go get github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump@latest
go: finding github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump latest
go: finding github.com/spiegel-im-spiegel/gpgpdump/cli latest
go: finding github.com/spiegel-im-spiegel/gpgpdump v0.5.0
go: downloading github.com/spiegel-im-spiegel/gpgpdump v0.5.0
go: extracting github.com/spiegel-im-spiegel/gpgpdump v0.5.0
go: finding github.com/inconshreveable/mousetrap v1.0.0
go: finding github.com/BurntSushi/toml v0.3.1
go: finding github.com/spf13/pflag v1.0.3
go: finding github.com/spiegel-im-spiegel/gocli v0.9.4
go: finding github.com/spf13/cobra v0.0.3
go: finding golang.org/x/crypto v0.0.0-20190320223903-b7391e95e576
go: finding golang.org/x/xerrors v0.0.0-20190315151331-d61658bd2e18
go: finding github.com/mattn/go-isatty v0.0.7
go: finding golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a
go: finding golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223
go: downloading github.com/spiegel-im-spiegel/gocli v0.9.4
go: downloading github.com/spf13/cobra v0.0.3
go: downloading golang.org/x/xerrors v0.0.0-20190315151331-d61658bd2e18
go: downloading golang.org/x/crypto v0.0.0-20190320223903-b7391e95e576
go: downloading github.com/BurntSushi/toml v0.3.1
go: extracting github.com/spiegel-im-spiegel/gocli v0.9.4
go: extracting golang.org/x/xerrors v0.0.0-20190315151331-d61658bd2e18
go: extracting github.com/BurntSushi/toml v0.3.1
go: extracting github.com/spf13/cobra v0.0.3
go: downloading github.com/spf13/pflag v1.0.3
go: extracting github.com/spf13/pflag v1.0.3
go: extracting golang.org/x/crypto v0.0.0-20190320223903-b7391e95e576

$ ~/go/bin/gpgpdump -v
gpgpdump dev-version
Copyright 2016-2019 Spiegel (based on pgpdump by kazu-yamamoto)
Licensed under Apache License, Version 2.0
```

よしよし。
ちゃんとビルドできてるな。

### OpenJDK を入れる{#jdk}

[OpenJDK] は3月と9月の半年毎にバージョンアップが行われる。
リリース時期とサポート期間は以下の通り（ディストリビューションによっては延長サポートあり）。

|         | Oracle Java           | [OpenJDK]             |
| ------- | --------------------- | --------------------- |
| Java 9  | 2017年9月 - 2018年3月 | 2017年9月 - 2018年3月 |
| Java 10 | 2018年3月 - 2018年9月 | 2018年3月 - 2018年9月 |
| Java 11 | 2018年9月 - 2026年9月 | 2018年9月 - 2019年3月 |
| Java 12 | -                     | 2019年3月 - 2019年9月 |
| Java 13 | -                     | 2019年9月 - 2020年3月 |
| Java 14 | -                     | 2020年3月 - 2020年9月 |
| Java 15 | -                     | 2020年9月 - 2021年3月 |
| Java 16 | -                     | 2021年3月 - 2021年9月 |
| Java 17 | 2021年9月 - 2029年9月 | 2021年9月 - 2022年3月 |

[OpenJDK] のインストールは APT で普通にできるが複数のバージョンがあるので注意が必要だ。
まずは提供しているバージョンを調べてみる（[Ubuntu] 19.04 の場合）。

```text
$ sudo apt search openjdk-\(\.\)\+-jre$
ソート中... 完了
全文検索... 完了  
nvidia-openjdk-8-jre/disco 9.+8u77~10.1.105-0ubuntu1 amd64
  NVIDIA provided OpenJDK Java runtime, using Hotspot JIT

openjdk-11-jre/disco-updates 11.0.3+7-1ubuntu1 amd64
  OpenJDK Java ランタイム - Hotspot JIT 版

openjdk-12-jre/disco,now 12.0.1+12-1 amd64
  OpenJDK Java ランタイム - Hotspot JIT 版

openjdk-13-jre/disco 13~13-0ubunt1 amd64
  OpenJDK Java ランタイム - Hotspot JIT 版
```

開発環境を含めるのであれば `jdk` で探す。

```text
$ sudo apt search openjdk-\(\.\)\+-jdk$
ソート中... 完了
全文検索... 完了  
openjdk-11-jdk/disco-updates 11.0.3+7-1ubuntu1 amd64
  OpenJDK Development Kit (JDK)

openjdk-12-jdk/disco 12.0.1+12-1 amd64
  OpenJDK Development Kit (JDK)

openjdk-13-jdk/disco 13~13-0ubunt1 amd64
  OpenJDK Development Kit (JDK)
```

今回は [OpenJDK] 12 のランタイムのみを入れたいので

```text
$ sudo apt install openjdk-12-jre
```

とすればよい。
上手くインストールできていれば

```text
$ java -version
openjdk version "12.0.1" 2019-04-16
OpenJDK Runtime Environment (build 12.0.1+12-Ubuntu-1)
OpenJDK 64-Bit Server VM (build 12.0.1+12-Ubuntu-1, mixed mode, sharing)
```

てな感じになる。
よしよし，[先日のアップデート](https://forest.watch.impress.co.jp/docs/news/1180549.html "Oracle、「Java SE 12.0.1」「Java SE 8 Update 211」を公開 ～新元号“令和”に対応 - 窓の杜")も反映されているな。

自分の環境に
複数バージョンの Java が入っている場合は

```
$ sudo update-alternatives --config java
```

とする。

```text
There are 2 choices for the alternative java (providing /usr/bin/java).  
Selection Path Priority Status 
———————————————————— 
* 0 /usr/lib/jvm/java-6-openjdk/jre/bin/java 1061 auto mode 
1 /usr/lib/jvm/jre1.7.0/jre/bin/java 3 manual mode  

Press enter to keep the current choice[*], or type selection number:
```

みたいな感じになって切り替えができるようだ。

- [OpenJDK（Java）を最新のUbuntuにインストール - Qiita](https://qiita.com/terappy/items/537c069923144a9d9755)
- [Java - Community Help Wiki](https://help.ubuntu.com/community/Java)

- [Java 環境のリリースとサポートに関する覚え書き]({{<ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})

## ブックマーク

- [Installing LibreOffice on Linux - The Document Foundation Wiki](https://wiki.documentfoundation.org/Documentation/Install/Linux)
- [Ubuntu 19.04 その17 - Xorgセッションで分数スケーリング（Fractional Scaling）を有効にするには - kledgeb](https://kledgeb.blogspot.com/2019/04/ubuntu-1904-17-xorgfractional-scaling.html)

[VirtualBox]: https://www.virtualbox.org/ "Oracle VM VirtualBox"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Debian]: https://www.debian.org/ "Debian -- The Universal Operating System"
[git]: https://git-scm.com/
[KDiff3]: http://kdiff3.sourceforge.net/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Go]: https://golang.org/ "The Go Programming Language"
[curl]: https://curl.haxx.se/
[OpenJDK]: http://openjdk.java.net/

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
