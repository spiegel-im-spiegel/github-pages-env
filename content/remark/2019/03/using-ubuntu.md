+++
title = "Ubuntu で遊ぶ"
date = "2019-03-31T10:14:36+09:00"
description = "この記事は覚え書きとして随時更新していく予定。"
image = "/images/attention/kitten.jpg"
tags = [ "virtualbox", "linux", "ubuntu", "security", "java", "install" ]
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
    - [Snap でパッケージ管理をする]({{< ref "#snap" >}})
    - [セキュリティ情報をチェックする]({{< relref "#secinfo" >}})
    - [えっ ifconfig って入ってないの？]({{< relref "#ifconfig" >}})
    - [OpenJDK を入れる]({{< relref "#jdk" >}})
    - [GUI な SFTP クライアントを導入する]({{< relref "#sftp" >}})
    - [GNOME Shell をリスタートする]({{< relref "#restart" >}})
    - [スワップの解放]({{< relref "#swap" >}})

{{< div-box type="md" >}}
いくつかの記事（インストール情報だけ書いた節）は「[Advanced Package Tool に関する覚え書き]({{< ref "/remark/2019/05/advanced-package-tool.md" >}})」へ移動した。
併せてどうぞ。
{{< /div-box >}}

## VirtualBox との連携{#vb}

### クリップボードの内容を共有する{#clipboard}

[Ubuntu] の動作確認はテキストファイルでメモを取りながら行っているのだが，ホスト OS とゲスト OS との間でクリップボードが共有できないか調べてみたら簡単にできるようだ。

インストール後の [VirtualBox] マネージャは以下の通り。

{{< fig-img src="../the-window-the-window/end-of-setup.png" link="../the-window-the-window/end-of-setup.png" width="976" >}}

この画面から「設定」アイコンをクリックすると以下の画面が開く。

{{< fig-img src="./setting-clipboard.png" link="./setting-clipboard.png" width="831" >}}

この画面の「一般」カテゴリの「高度」タブに「クリップボードの共有」の項目がある。
ここを「双方向」にすればホスト OS とゲスト OS との間でクリップボードを共有できる。

[VirtualBox]: https://www.virtualbox.org/ "Oracle VM VirtualBox"

## Ubuntu に関する雑多なこと{#ubuntu}

そもそも [Ubuntu] でターミナルってどうやって開くんだろうと思ったが，デスクトップを右クリックして「端末を開く」でよかった。

### Advanced Package Tool{#apt}

[Ubuntu] は [Debian] 系のディストリビューションなのでパッケージやシステムの管理には APT (Advanced Package Tool) を使う。
`apt-get` 等が `apt` コマンドに統合されているとは知らなかったよ。
[Debian] 系は古いバージョンしか使ったことなかったからなぁ（笑）

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

今回はなかったが保留パッケージを含めて更新する場合は

```text
$ sudo apt full-upgrade
```

を実行する。

ちなみに APT のログは `/var/log/apt/` ディレクトリにある。
あと `/var/log/dpkg.log` ファイルも参考になるだろう。

### Snap でパッケージ管理をする{#snap}

APT 以外では [Snap] を使ったパッケージ管理もあるらしい。
[Snap] は [Ubuntu] 18.04 以降は既定で入っているらしい。

```text
$ apt show snapd
Package: snapd
Version: 2.38+19.04
Built-Using: apparmor (= 2.12-4ubuntu10), golang-1.10 (= 1.10.4-2ubuntu1), libcap2 (= 1:2.25-2), libseccomp (= 2.3.3-3ubuntu2)
Priority: optional
Section: devel
Origin: Ubuntu
...
```

ただし APT とは別管理になるため取り扱いには注意が必要だろう。
たとえば LibreOffice や Firefox は APT 版と [Snap] 版が混在してしまう。
また Thunderbird のように [Snap] 版のほうがバージョンが古かったりする場合もある。

[Snap] で自機に導入されているパッケージは

```text
$ snap list
```

で見れる。
また [Snap] が提供しているパッケージは

```text
$ snap find packagename
```

で検索でき

```text
$ sudo snap install packagename
```

でインストールできる。
削除は

```text
$ sudo snap remove packagename
```

で OK。
更新は

```text
$ sudo snap refresh
```

で一括更新できるようだ。

- [uApp Explorer](https://uappexplorer.com/snaps)
- [Ubuntu 18.04 LTSでSnapパッケージをデフォルト同梱提案 | マイナビニュース](https://news.mynavi.jp/article/20180214-582997/)
- [snapdインストール方法まとめ【Ubuntu・Linux Mint・Debian・Fedoraなど】 | LFI](https://linuxfan.info/snapd)
- [第515回　Ubuntu 18.04 LTSとSnapパッケージ：Ubuntu Weekly Recipe｜gihyo.jp … 技術評論社](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0515)
- [UbuntuにVSCodeをインストールする3つの方法 - Qiita](https://qiita.com/yoshiyasu1111/items/e21a77ed68b52cb5f7c8) : VS Code は [Snap] でインストールしたほうがいいかも

[Snap]: https://github.com/snapcore/snapd "snapcore/snapd: The snapd and snap tools enable systems to work with .snap files."

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

### OpenJDK を入れる{#jdk}

{{< div-box type="md" >}}
**【追記 2019-07-31】**
どうも Java 12 は放置プレイっぽいので [APT は使わない方向]({{< ref "/remark/2019/07/reinstalling-openjdk.md" >}} "結局 OpenJDK をインストールし直すことにした")で。
{{< /div-box >}}

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

[OpenJDK]: http://openjdk.java.net/

### GUI な SFTP クライアントを導入する{#sftp}

[Ubuntu] で GUI な SFTP クライアントというと [gFTP] か [FileZilla] になりそうだが [gFTP] の最終リリースが2008年というのを見て [FileZilla] を選択した。
いや古いのが悪いってわけじゃないんだけど，10年以上セキュリティ・アップデートもないってのはちょっと信じられない。
やはりソフトウェアを取り巻くコミュニティが活況でないということなんだろう。

[FileZilla] ドキュメントには

{{< fig-quote title="Client Installation - FileZilla Wiki" link="https://wiki.filezilla-project.org/Client_Installation#Installing_on_GNU.2FLinux_and_other_Unix.28-like.29_systems" lang="en" >}}
<q>It is recommended that you use the package manager of your distribution.<br>
If you're using GNU/Linux, you can also try using the precompiled binaries. After extracting the files to any location (location does not matter, FileZilla can detect its own installation prefix), you can start the program using the filezilla executable in the bin/ subdirectory. Please note that due to differences in distributions, the provided binaries might not work on your system. </q>
{{< /fig-quote >}}

と書かれていて，とりあえず APT で調べてみると

```text
$ apt show filezilla
Package: filezilla
Version: 3.39.0-2
Priority: optional
Section: universe/net
Origin: Ubuntu
...
```

いや，古すぎるだろ。
こんなんばっかだな APT は。

ちうわけで[ダウンロードページ](https://filezilla-project.org/download.php "Download FileZilla Client for Linux")から取ってきたファイルを展開して使うことにする。

```text
$ sudo mv ./FileZilla_3.42.1_x86_64-linux-gnu.tar.bz2 /usr/local/src/
$ cd /usr/local/
$ sudo tar xvf src/FileZilla_3.42.1_x86_64-linux-gnu.tar.bz2
$ sudo chown -R root:root FileZilla3
```

展開時のファイル・ディレクトリ構成を維持していれば [FileZilla] は起動時のパスを見て適切に起動できるようだ。
なので，別にパスを通さなくても `.bashrc` などで

```bash
alias fz='/usr/local//FileZilla3/bin/filezilla'
```

とか定義しておけば

```text
$ fz &
```

で起動できる。

使い方は見れば分かる（笑）
「サイト マネージャー」で接続先を設定して「接続」すればよい。
SFTP で公開鍵認証をを行う場合はログオンタイプを「インタラクティブ」にしておけば接続時に `SSH_AUTH_SOCK` 環境変数で指定したソケットからエージェントが起動する。

エージェントに `gpg-agent` を利用する場合は[拙文]({{< ref "/remark/2019/04/move-gpg-keyring.md" >}} "Windows 環境で作った GnuPG の鍵束を Ubuntu に移行する")を参考にどうぞ。
[FileZilla] に限らず，セキュリティ製品でもないソフトウェアにパスワードや秘密鍵を登録するのはオススメしない。

[gFTP]: https://www.gftp.org/
[FileZilla]: https://filezilla-project.org/ "FileZilla - The free FTP solution"

### GNOME Shell をリスタートする{#restart}

現在 [Ubuntu] の GUI Shell である GNOME Shell はコマンド一発で簡単にリスタートできるらしい。

- [GNOME Shellを「リスタート」して軽くしよう！ | LFI](https://linuxfan.info/restart-gnome-shell)

手順は以下の通り

1. `alt-F2` で「コマンドを入力」を表示させる
2. コマンド「`r`」を入力してリターンキーを押す

以上。

システムモニタを見ると `gnome-shell` プロセスの使用メモリが減っているのが分かる。
再起動や再ログインなしでできるのは嬉しい。

### スワップの解放{#swap}

マシンを稼働させっぱなしにしていると少しずつスワップ領域が食われていくので以下のコマンドで解放する。

```text
$ sudo swapoff -a
$ sudo swapon -a
```

`swapoff`コマンドでスワップアウトされているデータをメモリに戻しているのだが，当然ながらメモリに空きがないと失敗する。
この場合は不要なプロセスを終了させるなどして空きを作ればいいのだが，無理そうなら諦めて再起動したほうが早いかもしれない。
まぁ，サーバ機では簡単に再起動とはいかないだろうが，デスクトップならね。
メモリに空きを作る手順は以下を参考にそうぞ。

- [[Linux]Swap領域をクリアする方法 ·  DQNEO起業日記](http://dqn.sakusakutto.jp/2014/01/linux_swap.html)

## その他ブックマーク

- [Ubuntu 19.04 その17 - Xorgセッションで分数スケーリング（Fractional Scaling）を有効にするには - kledgeb](https://kledgeb.blogspot.com/2019/04/ubuntu-1904-17-xorgfractional-scaling.html)
- [Debian Linuxで7zファイルを圧縮・解凍する / p7zipの使い方 -- ぺけみさお](https://www.xmisao.com/2014/09/25/debian-linux-extract-7z.html)
- [Amazon DriveをUbuntu（Linux）で自動マウントする | web net FORCE](https://webnetforce.net/amazon-drive-auto-mount-for-ubuntu/)
- [gFTP――多機能で便利なLinuxファイル転送の万能ナイフ (1/2) - ITmedia エンタープライズ](https://www.itmedia.co.jp/enterprise/articles/0705/29/news010.html)
- [Ubuntu 16.04 で MinGW を使い C++11 プログラムの Windows バイナリをクロスコンパイルする - Qiita](https://qiita.com/syoyo/items/678f9c3c82d5f6f5607c)
- [GNOMEユーザー必見！「GSConnect」でAndroidと常時連携！ブラウザ拡張も登場！ | LFI](https://linuxfan.info/gsconnect)
- [Linuxでネットワークトラフィックを監視する方法 | マイナビニュース](https://news.mynavi.jp/article/20100730-linux-freebsd-net-traffic-tools/)
    - [Wireshark · Go Deep.](https://www.wireshark.org/) : Wireshark って元々 Linux 用のツールだったんだな。もの知らずでゴメンペコン

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Debian]: https://www.debian.org/ "Debian -- The Universal Operating System"

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
