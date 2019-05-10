+++
title = "Advanced Package Tool に関する覚え書き"
date =  "2019-05-10T18:32:34+09:00"
description = "今回の件で自機でのアプリケーションあるいはパッケージ管理について，ちょっと考えてしまった。"
image = "/images/attention/kitten.jpg"
tags = [ "linux", "ubuntu", "install", "tools", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 Firefox がやらしたじゃないですか。

- [Thread: Add-ons not working due to certificate expiration - Add-on Support - Mozilla Discourse](https://discourse.mozilla.org/t/thread-add-ons-not-working-due-to-certificate-expiration/38968)

証明書の期限切れでアドオンが全部排除されるというバグとしてはありがちなやつなんだけど， Mozilla が改修版（v66.0.4）を出してから [Ubuntu] の [APT] でアップデートできるようになるまで概ね2日ほどかかってるのよ。

今回の件で自機でのアプリケーションあるいはパッケージ管理について，ちょっと考えてしまった。
そこでまずは [APT] について調べ直すところから始めることにした。

## dpkg と Advanced Package Tool

皆さんご存知の通り [Ubuntu] は [Debian] から fork したディストリビューションでアプリケーション管理も [Debian] のものを踏襲している。

アプリケーションのビルド済みバイナリや関連ファイルにバージョン情報や依存情報等のメタデータを付加したものを（ar/tar/gzip/bzip2 などで）ひとまとめにしたものを「[パッケージ](https://debian-handbook.info/browse/ja-JP/stable/packaging-system.html "第 5 章 パッケージシステム、ツールと基本原則")」と呼ぶ。
[Debian]/[Ubuntu] では `.deb` の拡張子が付いたファイルがそのパッケージファイルで，パッケージファイルを利用するためのツールが [dpkg] である。

ただし [dpkg] にはプリミティブな機能しかないため一般のユーザが [dpkg] をそのまま使うことはまずない。
一般ユーザ用のフロントエンド（のひとつ）が [APT](https://debian-handbook.info/browse/ja-JP/stable/apt.html "第 6 章 メンテナンスと更新、APT ツール") (Advanced Package Tool) ということになる。

[APT] はバックエンドにパッケージ管理用のデータベースを持っていて[^sl1]，このデータベースをもとにパッケージ間の依存関係を維持しながら可能な限り自動で導入や削除を行おうとする。

[^sl1]: `/etc/apt/` ディレクトリ以下のファイル群がそれ。このうち `sources.list` が [Ubuntu] 公式のパッケージ・リポジトリを定義したファイルである。またサードパーティのリポジトリは `/etc/apt/sources.list.d/` ディレクトリに `*.list` ファイルで設定可能である。

したがってデータベースにないパッケージは [APT] では導入できない。
この場合は以下の3つの手段をとることができる。

1. サードパーティのリポジトリを登録して [APT] から導入できるようにする
2. deb ファイルを使って直接インストールする
3. ビルド済みバイナリを直接展開して導入する。またはソースファイルから直接ビルドを行う

## [Ubuntu] パッケージのリリースサイクル

[Ubuntu] は概ね半年ごとにアップグレードされ，リリース時の年月がそのままバージョン番号になっている。
たとえば先日2019年4月にリリースされたバージョンには 19.04 が振られている。

[Ubuntu] に収録されるパッケージは OS リリース時にバージョンが固定され重大な不具合や脆弱性が発覚しない限り更新されることはないようだ。
先日の Firefox の件はむしろ例外的に早い対応だったということになる。

しかし，昨今は活況なソフトウェアほどリリースサイクルが短い傾向があり半年というタイムスケールでは追いつかないことも多い。
自身でリスクを引き受けてでも [APT] による管理を離れて自前で最新バージョンを維持したいという要求もあると思う。

つまり [APT] で管理可能なパッケージについても

1. [APT] で導入する
2. deb ファイルを使って直接インストールする
3. ビルド済みバイナリを直接展開して導入する。またはソースファイルから直接ビルドを行う

という3つの手段をとり得るわけだ。

そこで以降からは管理方法毎によく使うパッケージを分類してみる。
なお，この分類は私の独断と偏見に依る部分が大きいので，他の人にはあまり参考にならないであろう点は先に誤っておく。
ゴメンペコン。

## 公式リポジトリから [APT] を使って管理するパッケージ

### セキュリティ関連ツールなのでしょうがない

以下のパッケージはセキュリティ関連ツールで，これらに依存するパッケーも多く，特に保守的な運用になっているようだ。
したがって安定的な運用を優先し [APT] による管理とする。

| 製品名    | パッケージ名     | 備考                                 |
| --------- | ---------------- | ------------------------------------ |
| [GnuPG]   | `gnupg`          | 既定でインストール済                 |
| [OpenSSH] | `openssh-client` | クライアント側。既定でインストール済 |
| [OpenSSL] | `openssl`        | 既定でインストール済                 |

以下も参考にどうぞ

- [Windows 環境で作った GnuPG の鍵束を Ubuntu に移行する]({{< ref "/remark/2019/04/move-gpg-keyring.md" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: https://www.openssh.com/
[OpenSSL]: https://www.openssl.org/

### [APT] に任せて安心なパッケージ

以下は何も考えずに [APT] に任せても大丈夫だろう。

| 製品名        | パッケージ名                    | 備考                                                                                |
| ------------- | ------------------------------- | ----------------------------------------------------------------------------------- |
| [Thunderbird] | `thunderbird`                   | 既定でインストール済                                                                |
| [LibreOffice] | `libreoffice`                   | 既定でインストール済                                                                |
| ifconfig      | `net-tools` [^nt1]              | 何故か既定で入ってなかった                                                          |
| [curl]        | `curl`                          | 何故か既定で入ってなかった                                                          |
| [OpenJDK]     | `openjdk-12-jre`                | JRE のみの場合。バージョンごとにパッケージ名が異なるので注意（左は Java 12 の場合） |
| [KDiff3]      | `kdiff3`                        |                                                                                     |
| [Graphviz]    | `graphviz`                      |                                                                                     |
| [vim]         | `vim`                           | 既定で入ってるのは `vim-tiny` で `vim` を入れると置き換わる                         |
| [BOINC]       | `boinc-client`, `boinc-manager` | クライアント側                                                                      |
| [KeePass]     | `keepass2`                      | あらかじめ [Mono] がインストールされていること                                      |

[^nt1]: パッケージ `net-tools` をインストールすると ifconfig のほかに arp, netstat, rarp, nameif, route といったツールがインストールされる。

[KDiff3]: http://kdiff3.sourceforge.net/
[curl]: https://curl.haxx.se/
[OpenJDK]: http://openjdk.java.net/
[Graphviz]: https://www.graphviz.org/ "Graphviz - Graph Visualization Software"
[BOINC]: https://boinc.berkeley.edu/
[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"
[Thunderbird]: https://www.thunderbird.net/ "Thunderbird — Software made to make email easier. — Mozilla"

### その他

依存関係とか導入手順とかが複雑で自前で管理するのが面倒くさい，て感じのパッケージ。

| 製品名      | パッケージ名              | 備考                                                                                       |
| ----------- | ------------------------- | ------------------------------------------------------------------------------------------ |
| [GCC] 等    | `build-essential`         | 何故か 18.10 には既定で入ってなかった                                                      |
| CIFS Client | `cifs-utils`              | 導入方法は[拙文を参照]({{< ref "/remark/2019/03/common-internet-file-system.md" >}})のこと |
| [ClamAV]    | `clamav`, `clamav-daemon` | 導入方法は[拙文を参照]({{< relref "./clamav-for-ubuntu.md" >}})のこと                      |

[GCC]: https://gcc.gnu.org/ "GCC, the GNU Compiler Collection - GNU Project - Free Software Foundation (FSF)"
[vim]: https://www.vim.org/
[KeePass]: https://keepass.info/ "KeePass Password Safe"
[ClamAV]: https://www.clamav.net/ "ClamavNet"

## サードパーティ・リポジトリから [APT] を使って管理するパッケージ

### [Git]

公式リポジトリでも導入可能だが最新版が欲しかったので [PPA] のリポジトリを使うことにした。
詳しくは

- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})

を参照のこと。

[Git]: https://git-scm.com/
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"

### [Mono]

インストールにはサードパーティ・リポジトリの登録と署名検証用の公開鍵の取得が必要。
詳しくは

- [Ubuntu に Mono を導入する]({{< ref "/remark/2019/04/mono-in-ubuntu.md" >}})

を参照のこと。

[Mono]: https://www.mono-project.com/

## deb ファイルを使って直接インストールする

deb ファイルを使って直接インストールするには gdebi を使うのがオススメである。
導入は [APT] からできる。

```text
$ sudo apt install gdebi-core
```

これで

```text
$ sudo gdebi foo.deb
```

とすれば依存パッケージも含めてインストールしてくれる。
内部で [APT] のデータベースを使ってるのかな。

自前で導入する場合は最新バージョンに常に注意すること。

### [ATOM] エディタ

[リリースページ](https://github.com/atom/atom/releases "Releases · atom/atom") から最新版の `atom-amd64.deb` ファイルをダウンロードしてインストールする。

```text
$ sudo gdebi ./atom-amd64.deb
```

詳しくは

- [Ubuntu に ATOM エディタを導入する]({{< ref "/remark/2019/04/atom-in-ubuntu.md" >}})

を参照のこと。

[ATOM]: https://atom.io/

### [Hugo]

いや，シングルバイナリで依存関係も殆どないので deb ファイルからインストールする必然性は微塵もないのだが，どうも [APT] からインストールできるパッケージが全く追従できてないみたいなので。

[リリースページ](https://github.com/gohugoio/hugo/releases "Releases · gohugoio/hugo")から最新版の deb ファイルを取ってきてインストールすればよい。

```text
$ sudo gdebi ./hugo_0.55.5_Linux-64bit.deb
```

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"

### もうひとつの [LibreOffice]

[LibreOffice] については安定版が [APT] で取得可能で既定で入っているのだが，[ダウンロードページ](https://www.libreoffice.org/download/download/ "Download LibreOffice | LibreOffice - Free Office Suite - Fun Project - Fantastic People")から deb ファイルを取ってきてインストールする方法もある。
この場合 [APT] 管理のバージョンと混在になってしまうので注意が必要である。

実はどうするか悩み中。

## ビルド済みバイナリを直接展開して導入する

自前で導入する場合は最新バージョンに常に注意すること。

### [Go] コンパイラ

[Go] コンパイラ自体は [APT] でも導入可能だが，お互いのリリースタイミングが悪いのか2世代もバージョンが古い。
これでは使いものにならないので（[Go] コンパイラの公式サポートは1世代前まで），[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")から `go1.xx.x.linux-amd64.tar.gz` ファイルを取ってきて任意の場所に展開して使う。

```text
$ cd /usr/local/src
$ sudo curl https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz -O
$ cd ..
$ unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.12.5.linux-amd64.tar.gz
$ sudo mv go go1.12.5
$ sudo ln -s go1.12.5 go
$ ./go/bin/go version
go version go1.12.5 linux/amd64
```

`bin/` ディレクトリにはパスを通しておけば大丈夫。

[Go]: https://golang.org/ "The Go Programming Language"

### [FileZilla]

これも [APT] で導入可能なのだが，だいぶバージョンが古いので，[ダウンロードページ](https://filezilla-project.org/download.php "Download FileZilla Client for Linux")から取ってきたファイルを展開して使うことにする。

```text
$ cd /usr/local/
$ sudo tar xvf src/FileZilla_3.42.1_x86_64-linux-gnu.tar.bz2
$ sudo chown -R root:root FileZilla3
```

`chown` でオーナーを変えるのを忘れないように。

[FileZilla]: https://filezilla-project.org/ "FileZilla - The free FTP solution"

### [Git Extensions]

[Mono] がインストールされていることが前提。
詳しくは

- [Ubuntu に Mono を導入する]({{< ref "/remark/2019/04/mono-in-ubuntu.md" >}})

を参照のこと。

[Git Extensions]: https://gitextensions.github.io/ "Git Extensions | Git Extensions is a graphical user interface for Git that allows you to control Git without using the commandline"

### [TeX Live]

[TeX Live] に関しては（大量のパッケージがあるため）最新の環境が必要ないのであれば [APT] を使うほうがオススメである。

```text
$ apt show texlive
Package: texlive
Version: 2018.20190227-1
Priority: optional
Section: universe/tex
Source: texlive-base
Origin: Ubuntu
...
```

が，やっぱり最新の環境がほしいので手動でインストールすることにした。

手動でインストールする場合はインストーラ `install-tl` を使う。
[TeX Live] 内の各パッケージの更新には tlmgr を使う。

```text
$ sudo tlmgr update --self --all
```

詳しくは

- [TeX Live を Ubuntu に（[APT] を使わずに）導入する]({{< ref "/remark/2019/04/install-texlive-in-ubuntu.md" >}})

を参照のこと。

[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"

## ソースファイルから直接ビルドを行う

自前で導入する場合は最新バージョンに常に注意すること。

### [pgpdump] のビルド

実は [APT] で導入できるっぽいのだが，自作の [gpgpdump] の動作確認用にリファレンス実装として最新版が必要なのよ。
[GCC] 等のツールチェーンがあれば簡単にビルドできる。

[リポジトリ](https://github.com/kazu-yamamoto/pgpdump "kazu-yamamoto/pgpdump: A PGP packet visualizer")からソースコードを取ってきて

```text
$ ./configure 
$ make
```

でビルドできる。
なお圧縮パケットの解凍に bz2 が必要な場合は [APT] でパッケージ `libbz2-dev` をあらかじめインストールしておくこと。

[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

### git-credential-libsecret のビルド

git-credential 用に git-credential-libsecret をビルドする。

libsecret 自体のインストールは [APT] で行う。

```text
$ sudo apt install libsecret-1-dev libgnome-keyring-dev
```

これで展開されるソースコードを適当な場所にコピーしてビルドする。

```text
$ cp -r /usr/share/doc/git/contrib/credential/libsecret ~/work
$ cd ~/work/libsecret
$ make
gcc -g -O2 -Wall  -pthread -I/usr/include/libsecret-1 -I/usr/include/libmount -I/usr/include/blkid -I/usr/include/uuid -I/usr/include/glib-2.0 -I/usr/lib/x86_64-linux-gnu/glib-2.0/include -o git-credential-libsecret.o -c git-credential-libsecret.c
gcc -o git-credential-libsecret  git-credential-libsecret.o -lsecret-1 -lgio-2.0 -lgobject-2.0 -lglib-2.0
```

ビルドした git-credential-libsecret をパスの通ったディレクトリに入れれば完了。

## ブックマーク

- [Ubuntu security notices](https://usn.ubuntu.com/)
- [Debian -- Security Information](https://www.debian.org/security/)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Debian]: https://www.debian.org/ "Debian -- The Universal Operating System"
[dpkg]: https://debian-handbook.info/browse/ja-JP/stable/sect.manipulating-packages-with-dpkg.html "5.4. dpkg を用いたパッケージの操作"
[APT]: https://debian-handbook.info/browse/ja-JP/stable/apt.html "第 6 章 メンテナンスと更新、APT ツール"
