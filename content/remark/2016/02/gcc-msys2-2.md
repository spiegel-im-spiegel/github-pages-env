+++
date = "2016-02-28T18:13:15+09:00"
description = "いよいよ gcc をインストールする。"
draft = true
tags = ["msys2", "gcc", "tools", "pacman"]
title = "MSYS2 による gcc 開発環境の構築 ― gcc パッケージ群の導入"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1. [MSYS2 による gcc 開発環境の構築 ― インストールから初期化処理まで]({{< relref "remark/2016/02/gcc-msys2-1.md" >}})
2. [MSYS2 による gcc 開発環境の構築 ― gcc パッケージ群の導入]({{< relref "remark/2016/02/gcc-msys2-2.md" >}})
3. [MSYS2 による gcc 開発環境の構築 ― pgpdump をビルドする]({{< relref "remark/2016/02/gcc-msys2-3.md" >}})

いよいよ gcc をインストールする。
その前に...

## pacman によるパッケージ管理

[前回]の初期化でもちょっとだけ出てきたが， [MSYS2] のパッケージ管理には pacman コマンドを使う。
pacman は元々 Arch Linux のパッケージ管理ツールで，使い方もこれと同じようだ。

- [pacman - ArchWiki](https://archlinuxjp.kusakata.com/wiki/Pacman)

よく使うコマンドとしては

| コマンドライン | 処理内容 |
|:---------------|:---------|
| `pacman -S <package_name1> [package_name2...]` | 指定したパッケージをインストールする |
| `pacman -Sl [repository]` | パッケージを一覧表示する（リポジトリを指定可能） |
| `pacman -Ss [regex...]` | パッケージを一覧・検索する（キーワードを指定可能） |
| `pacman -Su` | インストール済みのパッケージを更新する |
| `pacman -Sy` | パッケージのデータベースを更新する |
| `pacman -Syu` | `-Su`, `-Sy` の同時指定 |
| `pacman -Q [package_name...]` | インストール済みパッケージを一覧・検索する（パッケージを指定可能） |
| `pacman -Qs [regex...]` | インストール済みパッケージを一覧・検索する（をキーワードを指定可能） |
| `pacman -R <package_name>` | 指定したパッケージを削除する |
| `pacman -Rs <package_name>` | 依存関係を含めて指定したパッケージを削除する（指定したパッケージにのみ依存するパッケージを削除） |

あたりか。
試しに git をインストールしてみよう。

```text
$ pacman -S git
依存関係を解決しています...
衝突するパッケージがないか確認しています...

パッケージ (34) db-5.3.28-2  expat-2.1.0-2  gdbm-1.11-3  heimdal-1.5.3-8
                libgdbm-1.11-3  openssh-7.1p2-1  perl-5.22.0-2
                perl-Authen-SASL-2.16-2  perl-Convert-BinHex-1.123-2
                perl-Encode-Locale-1.04-1  perl-Error-0.17024-1
                perl-File-Listing-6.04-2  perl-HTML-Parser-3.71-3
                perl-HTML-Tagset-3.20-2  perl-HTTP-Cookies-6.01-2
                perl-HTTP-Daemon-6.01-2  perl-HTTP-Date-6.02-2
                perl-HTTP-Message-6.06-2  perl-HTTP-Negotiate-6.01-2
                perl-IO-Socket-SSL-2.016-1  perl-IO-stringy-2.111-1
                perl-LWP-MediaTypes-6.02-2  perl-MIME-tools-5.506-1
                perl-MailTools-2.14-1  perl-Net-HTTP-6.09-1
                perl-Net-SMTP-SSL-1.02-1  perl-Net-SSLeay-1.70-1
                perl-TermReadKey-2.33-1  perl-TimeDate-2.30-2  perl-URI-1.68-1
                perl-WWW-RobotRules-6.02-2  perl-libwww-6.13-1  vim-7.4.1415-2
                git-2.7.1-1

合計ダウンロード容量:   24.75 MiB
合計インストール容量:  114.53 MiB

:: インストールを行いますか？ [Y/n] y
:: パッケージを取得します ...
 expat-2.1.0-2-x86_64                 13.1 KiB  4.26M/s 00:00 [##################] 100%
 vim-7.4.1415-2-x86_64                 6.1 MiB  1248K/s 00:05 [##################] 100%
 heimdal-1.5.3-8-x86_64              543.7 KiB   840K/s 00:01 [##################] 100%
 openssh-7.1p2-1-x86_64              653.4 KiB   832K/s 00:01 [##################] 100%
 db-5.3.28-2-x86_64                   41.7 KiB  8.14M/s 00:00 [##################] 100%
 libgdbm-1.11-3-x86_64                20.4 KiB  6.63M/s 00:00 [##################] 100%
 gdbm-1.11-3-x86_64                  108.5 KiB  8.83M/s 00:00 [##################] 100%
 perl-5.22.0-2-x86_64                 12.4 MiB  1231K/s 00:10 [##################] 100%
 perl-Error-0.17024-1-any             17.1 KiB  16.7M/s 00:00 [##################] 100%
 perl-Authen-SASL-2.16-2-any          42.4 KiB  10.4M/s 00:00 [##################] 100%
 perl-Encode-Locale-1.04-1-any         9.7 KiB  9.46M/s 00:00 [##################] 100%
 perl-HTTP-Date-6.02-2-any             8.6 KiB  0.00B/s 00:00 [##################] 100%
 perl-File-Listing-6.04-2-any          7.7 KiB  0.00B/s 00:00 [##################] 100%
 perl-HTML-Tagset-3.20-2-any          10.3 KiB  0.00B/s 00:00 [##################] 100%
 perl-HTML-Parser-3.71-3-x86_64       76.9 KiB  10.7M/s 00:00 [##################] 100%
 perl-LWP-MediaTypes-6.02-2-any       18.0 KiB  17.6M/s 00:00 [##################] 100%
 perl-URI-1.68-1-any                  75.6 KiB  12.3M/s 00:00 [##################] 100%
 perl-HTTP-Message-6.06-2-any         71.3 KiB  9.94M/s 00:00 [##################] 100%
 perl-HTTP-Cookies-6.01-2-any         20.4 KiB  20.0M/s 00:00 [##################] 100%
 perl-HTTP-Daemon-6.01-2-any          14.2 KiB  13.9M/s 00:00 [##################] 100%
 perl-HTTP-Negotiate-6.01-2-any       11.4 KiB  11.2M/s 00:00 [##################] 100%
 perl-Net-HTTP-6.09-1-any             19.8 KiB  9.65M/s 00:00 [##################] 100%
 perl-WWW-RobotRules-6.02-2-any       12.2 KiB  12.0M/s 00:00 [##################] 100%
 perl-libwww-6.13-1-any              122.2 KiB   255K/s 00:00 [##################] 100%
 perl-TimeDate-2.30-2-any             35.9 KiB  57.5K/s 00:01 [##################] 100%
 perl-MailTools-2.14-1-any            58.4 KiB  74.5K/s 00:01 [##################] 100%
 perl-IO-stringy-2.111-1-any          52.6 KiB   112K/s 00:00 [##################] 100%
 perl-Convert-BinHex-1.123-2-any      30.1 KiB  96.4K/s 00:00 [##################] 100%
 perl-MIME-tools-5.506-1-any         180.4 KiB   128K/s 00:01 [##################] 100%
 perl-Net-SSLeay-1.70-1-x86_64       191.2 KiB   244K/s 00:01 [##################] 100%
 perl-IO-Socket-SSL-2.016-1-any      112.5 KiB   358K/s 00:00 [##################] 100%
 perl-Net-SMTP-SSL-1.02-1-any          3.5 KiB  0.00B/s 00:00 [##################] 100%
 perl-TermReadKey-2.33-1-x86_64       20.9 KiB  10.2M/s 00:00 [##################] 100%
 git-2.7.1-1-x86_64                    3.6 MiB  1128K/s 00:03 [##################] 100%
(34/34) キーリングのキーを確認                                [##################] 100%
(34/34) パッケージの整合性をチェック                          [##################] 100%
(34/34) パッケージファイルのロード                            [##################] 100%
(34/34) ファイルの衝突をチェック                              [##################] 100%
(34/34) 空き容量を確認                                        [##################] 100%
:: パッケージの変更を処理しています...
( 1/34) インストール expat                                    [##################] 100%
( 2/34) インストール vim                                      [##################] 100%
( 3/34) インストール heimdal                                  [##################] 100%
( 4/34) インストール openssh                                  [##################] 100%
( 5/34) インストール db                                       [##################] 100%
( 6/34) インストール libgdbm                                  [##################] 100%
( 7/34) インストール gdbm                                     [##################] 100%
( 8/34) インストール perl                                     [##################] 100%
( 9/34) インストール perl-Error                               [##################] 100%
(10/34) インストール perl-Authen-SASL                         [##################] 100%
(11/34) インストール perl-Encode-Locale                       [##################] 100%
(12/34) インストール perl-HTTP-Date                           [##################] 100%
(13/34) インストール perl-File-Listing                        [##################] 100%
(14/34) インストール perl-HTML-Tagset                         [##################] 100%
(15/34) インストール perl-HTML-Parser                         [##################] 100%
(16/34) インストール perl-LWP-MediaTypes                      [##################] 100%
(17/34) インストール perl-URI                                 [##################] 100%
(18/34) インストール perl-HTTP-Message                        [##################] 100%
(19/34) インストール perl-HTTP-Cookies                        [##################] 100%
(20/34) インストール perl-HTTP-Daemon                         [##################] 100%
(21/34) インストール perl-HTTP-Negotiate                      [##################] 100%
(22/34) インストール perl-Net-HTTP                            [##################] 100%
(23/34) インストール perl-WWW-RobotRules                      [##################] 100%
(24/34) インストール perl-libwww                              [##################] 100%
perl-libwww の提案パッケージ
    perl-LWP-Protocol-HTTPS: for https:// url schemes
(25/34) インストール perl-TimeDate                            [##################] 100%
(26/34) インストール perl-MailTools                           [##################] 100%
(27/34) インストール perl-IO-stringy                          [##################] 100%
(28/34) インストール perl-Convert-BinHex                      [##################] 100%
module test... pass.
(29/34) インストール perl-MIME-tools                          [##################] 100%
(30/34) インストール perl-Net-SSLeay                          [##################] 100%
(31/34) インストール perl-IO-Socket-SSL                       [##################] 100%
(32/34) インストール perl-Net-SMTP-SSL                        [##################] 100%
(33/34) インストール perl-TermReadKey                         [##################] 100%
(34/34) インストール git                                      [##################] 100%
git の提案パッケージ
    python2: various helper scripts
    subversion: git svn
```

てか，本当に最小限しか入ってないんだな（笑） 提案パッケージはとりあえずスルーしても問題ない。

### MSYS2 のリポジトリ

`/etc/pacman.conf` を見ると MSYS2 では `mingw32`, `mingw64`, `msys` の3つのリポジトリを管理していることが分かる。

| リポジトリ名 | 参照ファイル |
|:-------------|:-------------|
| `mingw32` | `/etc/pacman.d/mirrorlist.mingw32` |
| `mingw64` | `/etc/pacman.d/mirrorlist.mingw64` |
| `msys` | `/etc/pacman.d/mirrorlist.msys` |

実際には「参照ファイル」にリポジトリへの URI が書かれている。
たとえば `/etc/pacman.d/mirrorlist.mingw32` ならこんな感じ。

```text
##
## 32-bit Mingw-w64 repository mirrorlist
##

## Primary
## msys2.org
Server = http://repo.msys2.org/mingw/i686
Server = http://downloads.sourceforge.net/project/msys2/REPOS/MINGW/i686
Server = http://www2.futureware.at/~nickoe/msys2-mirror/i686/
```

では実際にリポジトリの中を除いてみよう。

```text
$ pacman -Sl | grep gcc
mingw32 mingw-w64-i686-gcc 5.3.0-2
mingw32 mingw-w64-i686-gcc-ada 5.3.0-2
mingw32 mingw-w64-i686-gcc-fortran 5.3.0-2
mingw32 mingw-w64-i686-gcc-libgfortran 5.3.0-2
mingw32 mingw-w64-i686-gcc-libs 5.3.0-2
mingw32 mingw-w64-i686-gcc-objc 5.3.0-2
mingw64 mingw-w64-x86_64-gcc 5.3.0-2
mingw64 mingw-w64-x86_64-gcc-ada 5.3.0-2
mingw64 mingw-w64-x86_64-gcc-fortran 5.3.0-2
mingw64 mingw-w64-x86_64-gcc-libgfortran 5.3.0-2
mingw64 mingw-w64-x86_64-gcc-libs 5.3.0-2
mingw64 mingw-w64-x86_64-gcc-objc 5.3.0-2
msys gcc 4.9.2-6
msys gcc-fortran 4.9.2-6
msys gcc-libs 4.9.2-6 [インストール済み]
msys mingw-w64-cross-gcc 4.9.2-3
```

どれを使うかはユーザ次第だが， `msys` リポジトリの方は開発用というより [MSYS2] ビルド用という感じである。
gcc を 32bit 版と 64bit 版で使い分ける場合は `msys` 版は入れないほうがいいかもしれない。

- [Windowsでgccなどを利用できるMSYS2の環境設定など - Qiita](http://qiita.com/chromabox/items/fd07bad3f426101fc4a6)

## gcc のインストール

では gcc のインストールを始めよう。

### mingw32 版 gcc のインストール

mingw32 版 gcc をインストールするには pacman コマンドで `mingw-w64-i686-toolchain` パッケージグループを導入すればよい。

```text
$ pacman -S mingw-w64-i686-toolchain
:: 16 のパッケージがグループ mingw-w64-i686-toolchain にあります:
:: リポジトリ mingw32
   1) mingw-w64-i686-binutils  2) mingw-w64-i686-crt-git  3) mingw-w64-i686-gcc
   4) mingw-w64-i686-gcc-ada  5) mingw-w64-i686-gcc-fortran  6) mingw-w64-i686-gcc-libgfortran
   7) mingw-w64-i686-gcc-libs  8) mingw-w64-i686-gcc-objc  9) mingw-w64-i686-gdb
   10) mingw-w64-i686-headers-git  11) mingw-w64-i686-libmangle-git
   12) mingw-w64-i686-libwinpthread-git  13) mingw-w64-i686-make  14) mingw-w64-i686-pkg-config
   15) mingw-w64-i686-tools-git  16) mingw-w64-i686-winpthreads-git

選択して下さい (デフォルト=all): 1 2 3 7 9 10 11 12 13 14 15 16
依存関係を解決しています...
衝突するパッケージがないか確認しています...

パッケージ (36) mingw-w64-i686-bzip2-1.0.6-4  mingw-w64-i686-ca-certificates-20150426-2
                mingw-w64-i686-expat-2.1.0-6  mingw-w64-i686-gdbm-1.11-3
                mingw-w64-i686-gettext-0.19.6-2  mingw-w64-i686-gmp-6.1.0-1
                mingw-w64-i686-isl-0.15-1  mingw-w64-i686-libffi-3.2.1-3
                mingw-w64-i686-libiconv-1.14-5  mingw-w64-i686-libsystre-1.0.1-2
                mingw-w64-i686-libtasn1-4.7-1  mingw-w64-i686-libtre-git-r122.c2f5d13-4
                mingw-w64-i686-mpc-1.0.3-2  mingw-w64-i686-mpfr-3.1.3.p0-2  
                mingw-w64-i686-ncurses-6.0.20150627-2  mingw-w64-i686-openssl-1.0.2.f-1
                mingw-w64-i686-p11-kit-0.23.1-3  mingw-w64-i686-python2-2.7.11-1
                mingw-w64-i686-readline-6.3.008-1  mingw-w64-i686-tcl-8.6.4-2
                mingw-w64-i686-termcap-1.3.1-2  mingw-w64-i686-tk-8.6.4-2
                mingw-w64-i686-windows-default-manifest-6.4-2  mingw-w64-i686-zlib-1.2.8-9
                mingw-w64-i686-binutils-2.25.1-2  mingw-w64-i686-crt-git-5.0.0.4622.38d327f-1
                mingw-w64-i686-gcc-5.3.0-2  mingw-w64-i686-gcc-libs-5.3.0-2
                mingw-w64-i686-gdb-7.10.1-1  mingw-w64-i686-headers-git-5.0.0.4622.38d327f-1
                mingw-w64-i686-libmangle-git-5.0.0.4509.2e5a9a2-1
                mingw-w64-i686-libwinpthread-git-5.0.0.4573.628fdbf-1
                mingw-w64-i686-make-4.1.2351.a80a8b8-1  mingw-w64-i686-pkg-config-0.29-1
                mingw-w64-i686-tools-git-5.0.0.4592.90b8472-1
                mingw-w64-i686-winpthreads-git-5.0.0.4573.628fdbf-1

合計ダウンロード容量:   71.93 MiB
合計インストール容量:  456.81 MiB

:: インストールを行いますか？ [Y/n] y
:: パッケージを取得します ...
 mingw-w64-i686-libiconv-1.14-5-any                   602.2 KiB   929K/s 00:01 [##################] 100%
 mingw-w64-i686-bzip2-1.0.6-4-any                      67.2 KiB  9.37M/s 00:00 [##################] 100%
 mingw-w64-i686-zlib-1.2.8-9-any                      153.2 KiB   317K/s 00:00 [##################] 100%
 mingw-w64-i686-binutils-2.25.1-2-any                  10.3 MiB  1114K/s 00:09 [##################] 100%
 mingw-w64-i686-headers-git-5.0.0.4622.38d327f-1-any    5.0 MiB  1114K/s 00:05 [##################] 100%
 mingw-w64-i686-crt-git-5.0.0.4622.38d327f-1-any     1717.0 KiB  1081K/s 00:02 [##################] 100%
 mingw-w64-i686-isl-0.15-1-any                        539.7 KiB   851K/s 00:01 [##################] 100%
 mingw-w64-i686-gmp-6.1.0-1-any                       459.3 KiB   943K/s 00:00 [##################] 100%
 mingw-w64-i686-mpfr-3.1.3.p0-2-any                   258.0 KiB   801K/s 00:00 [##################] 100%
 mingw-w64-i686-mpc-1.0.3-2-any                        62.3 KiB  10.1M/s 00:00 [##################] 100%
 mingw-w64-i686-libwinpthread-git-5.0.0.4573.628...    27.1 KiB  8.84M/s 00:00 [##################] 100%
 mingw-w64-i686-gcc-libs-5.3.0-2-any                  609.8 KiB   761K/s 00:01 [##################] 100%
 mingw-w64-i686-windows-default-manifest-6.4-2-any   1472.0   B  0.00B/s 00:00 [##################] 100%
 mingw-w64-i686-winpthreads-git-5.0.0.4573.628fd...    35.1 KiB  11.4M/s 00:00 [##################] 100%
 mingw-w64-i686-gcc-5.3.0-2-any                        24.6 MiB  1285K/s 00:20 [##################] 100%
 mingw-w64-i686-expat-2.1.0-6-any                     107.1 KiB  11.6M/s 00:00 [##################] 100%
 mingw-w64-i686-gettext-0.19.6-2-any                    3.0 MiB  1012K/s 00:03 [##################] 100%
 mingw-w64-i686-gdbm-1.11-3-any                       150.1 KiB   310K/s 00:00 [##################] 100%
 mingw-w64-i686-libffi-3.2.1-3-any                     36.3 KiB   230K/s 00:00 [##################] 100%
 mingw-w64-i686-libtre-git-r122.c2f5d13-4-any          69.1 KiB   427K/s 00:00 [##################] 100%
 mingw-w64-i686-libsystre-1.0.1-2-any                   9.1 KiB  0.00B/s 00:00 [##################] 100%
 mingw-w64-i686-ncurses-6.0.20150627-2-any           1734.6 KiB   685K/s 00:03 [##################] 100%
 mingw-w64-i686-libtasn1-4.7-1-any                    171.3 KiB   352K/s 00:00 [##################] 100%
 mingw-w64-i686-p11-kit-0.23.1-3-any                  198.0 KiB   613K/s 00:00 [##################] 100%
 mingw-w64-i686-ca-certificates-20150426-2-any        382.4 KiB   795K/s 00:00 [##################] 100%
 mingw-w64-i686-openssl-1.0.2.f-1-any                   2.7 MiB   929K/s 00:03 [##################] 100%
 mingw-w64-i686-termcap-1.3.1-2-any                    12.3 KiB  6.01M/s 00:00 [##################] 100%
 mingw-w64-i686-readline-6.3.008-1-any                327.2 KiB   988K/s 00:00 [##################] 100%
 mingw-w64-i686-tcl-8.6.4-2-any                         2.8 MiB  1150K/s 00:03 [##################] 100%
 mingw-w64-i686-tk-8.6.4-2-any                       1874.2 KiB  1082K/s 00:02 [##################] 100%
 mingw-w64-i686-python2-2.7.11-1-any                   10.8 MiB  1273K/s 00:09 [##################] 100%
 mingw-w64-i686-gdb-7.10.1-1-any                        2.8 MiB   985K/s 00:03 [##################] 100%
 mingw-w64-i686-libmangle-git-5.0.0.4509.2e5a9a2...    16.5 KiB  16.1M/s 00:00 [##################] 100%
 mingw-w64-i686-make-4.1.2351.a80a8b8-1-any           105.3 KiB  10.3M/s 00:00 [##################] 100%
 mingw-w64-i686-pkg-config-0.29-1-any                 235.8 KiB   732K/s 00:00 [##################] 100%
 mingw-w64-i686-tools-git-5.0.0.4592.90b8472-1-any    235.1 KiB   721K/s 00:00 [##################] 100%
(36/36) キーリングのキーを確認                                                 [##################] 100%
(36/36) パッケージの整合性をチェック                                           [##################] 100%
(36/36) パッケージファイルのロード                                             [##################] 100%
(36/36) ファイルの衝突をチェック                                               [##################] 100%
(36/36) 空き容量を確認                                                         [##################] 100%
:: パッケージの変更を処理しています...
( 1/36) インストール mingw-w64-i686-libiconv                                   [##################] 100%
( 2/36) インストール mingw-w64-i686-bzip2                                      [##################] 100%
( 3/36) インストール mingw-w64-i686-zlib                                       [##################] 100%
( 4/36) インストール mingw-w64-i686-binutils                                   [##################] 100%
( 5/36) インストール mingw-w64-i686-headers-git                                [##################] 100%
( 6/36) インストール mingw-w64-i686-crt-git                                    [##################] 100%
( 7/36) インストール mingw-w64-i686-isl                                        [##################] 100%
( 8/36) インストール mingw-w64-i686-gmp                                        [##################] 100%
( 9/36) インストール mingw-w64-i686-mpfr                                       [##################] 100%
(10/36) インストール mingw-w64-i686-mpc                                        [##################] 100%
(11/36) インストール mingw-w64-i686-libwinpthread-git                          [##################] 100%
(12/36) インストール mingw-w64-i686-gcc-libs                                   [##################] 100%
(13/36) インストール mingw-w64-i686-windows-default-manifest                   [##################] 100%
(14/36) インストール mingw-w64-i686-winpthreads-git                            [##################] 100%
(15/36) インストール mingw-w64-i686-gcc                                        [##################] 100%
(16/36) インストール mingw-w64-i686-expat                                      [##################] 100%
(17/36) インストール mingw-w64-i686-gettext                                    [##################] 100%
(18/36) インストール mingw-w64-i686-gdbm                                       [##################] 100%
(19/36) インストール mingw-w64-i686-libffi                                     [##################] 100%
(20/36) インストール mingw-w64-i686-libtre-git                                 [##################] 100%
(21/36) インストール mingw-w64-i686-libsystre                                  [##################] 100%
(22/36) インストール mingw-w64-i686-ncurses                                    [##################] 100%
(23/36) インストール mingw-w64-i686-libtasn1                                   [##################] 100%
(24/36) インストール mingw-w64-i686-p11-kit                                    [##################] 100%
(25/36) インストール mingw-w64-i686-ca-certificates                            [##################] 100%
(26/36) インストール mingw-w64-i686-openssl                                    [##################] 100%
(27/36) インストール mingw-w64-i686-termcap                                    [##################] 100%
(28/36) インストール mingw-w64-i686-readline                                   [##################] 100%
(29/36) インストール mingw-w64-i686-tcl                                        [##################] 100%
(30/36) インストール mingw-w64-i686-tk                                         [##################] 100%
(31/36) インストール mingw-w64-i686-python2                                    [##################] 100%
(32/36) インストール mingw-w64-i686-gdb                                        [##################] 100%
(33/36) インストール mingw-w64-i686-libmangle-git                              [##################] 100%
(34/36) インストール mingw-w64-i686-make                                       [##################] 100%
(35/36) インストール mingw-w64-i686-pkg-config                                 [##################] 100%
(36/36) インストール mingw-w64-i686-tools-git                                  [##################] 100%
```

ここでは Ada, FORTRAN, Object-C の言語パッケージを除いたものをインストールした。

動作確認には `mingw32_shell.bat` の起動，または 環境変数 `MSYSTEM` に `MINGW32` をセットして shell を起動する。
[前回]紹介した [ConEmu] を使うのであれば以下のシーケンスで起動できる。

```text
set MSYSTEM=MINGW32 & C:\msys64\usr\bin\bash.exe --login -i
```

mingw32 版の shell が起動したところで gcc の動作確認を行う。

```text
$ gcc -v
Using built-in specs.
COLLECT_GCC=C:\msys64\mingw32\bin\gcc.exe
COLLECT_LTO_WRAPPER=C:/msys64/mingw32/bin/../lib/gcc/i686-w64-mingw32/5.3.0/lto-wrapper.exe
Target: i686-w64-mingw32
Configured with: ../gcc-5.3.0/configure --prefix=/mingw32 --with-local-prefix=/mingw32/local --build=i686-w64-mingw32 --host=i686-w64-mingw32 --target=i686-w64-mingw32 --with-native-system-header-dir=/mingw32/i686-w64-mingw32/include --libexecdir=/mingw32/lib --with-gxx-include-dir=/mingw32/include/c++/5.3.0 --enable-bootstrap --with-arch=i686 --with-tune=generic --enable-languages=c,lto,c++,objc,obj-c++,fortran,ada --enable-shared --enable-static --enable-libatomic --enable-threads=posix --enable-graphite --enable-fully-dynamic-string --enable-libstdcxx-time=yes --disable-libstdcxx-pch --disable-libstdcxx-debug --enable-version-specific-runtime-libs --disable-isl-version-check --enable-lto --enable-libgomp --disable-multilib --enable-checking=release --disable-rpath --disable-win32-registry --disable-nls --disable-werror --disable-symvers --with-libiconv --with-system-zlib --with-gmp=/mingw32 --with-mpfr=/mingw32 --with-mpc=/mingw32 --with-isl=/mingw32 --with-pkgversion='Rev2, Built by MSYS2 project' --with-bugurl=http://sourceforge.net/projects/msys2 --with-gnu-as --with-gnu-ld --disable-sjlj-exceptions --with-dwarf2
Thread model: posix
gcc version 5.3.0 (Rev2, Built by MSYS2 project)
```

### mingw64 版 gcc のインストール

同じように mingw64 版 gcc をインストールするには pacman コマンドで `mingw-w64-x86_64-toolchain` パッケージグループを導入すればよい。

```text
$ pacman -S mingw-w64-x86_64-toolchain
:: 16 のパッケージがグループ mingw-w64-x86_64-toolchain にあります:
:: リポジトリ mingw64
   1) mingw-w64-x86_64-binutils  2) mingw-w64-x86_64-crt-git  3) mingw-w64-x86_64-gcc
   4) mingw-w64-x86_64-gcc-ada  5) mingw-w64-x86_64-gcc-fortran  6) mingw-w64-x86_64-gcc-libgfortran
   7) mingw-w64-x86_64-gcc-libs  8) mingw-w64-x86_64-gcc-objc  9) mingw-w64-x86_64-gdb
   10) mingw-w64-x86_64-headers-git  11) mingw-w64-x86_64-libmangle-git
   12) mingw-w64-x86_64-libwinpthread-git  13) mingw-w64-x86_64-make  14) mingw-w64-x86_64-pkg-config
   15) mingw-w64-x86_64-tools-git  16) mingw-w64-x86_64-winpthreads-git

選択して下さい (デフォルト=all): 1 2 3 7 9 10 11 12 13 14 15 16
依存関係を解決しています...
衝突するパッケージがないか確認しています...

パッケージ (36) mingw-w64-x86_64-bzip2-1.0.6-4  mingw-w64-x86_64-ca-certificates-20150426-2
                mingw-w64-x86_64-expat-2.1.0-6  mingw-w64-x86_64-gdbm-1.11-3
                mingw-w64-x86_64-gettext-0.19.6-2  mingw-w64-x86_64-gmp-6.1.0-1
                mingw-w64-x86_64-isl-0.15-1  mingw-w64-x86_64-libffi-3.2.1-3
                mingw-w64-x86_64-libiconv-1.14-5  mingw-w64-x86_64-libsystre-1.0.1-2
                mingw-w64-x86_64-libtasn1-4.7-1  mingw-w64-x86_64-libtre-git-r122.c2f5d13-4
                mingw-w64-x86_64-mpc-1.0.3-2  mingw-w64-x86_64-mpfr-3.1.3.p0-2
                mingw-w64-x86_64-ncurses-6.0.20150627-2  mingw-w64-x86_64-openssl-1.0.2.f-1
                mingw-w64-x86_64-p11-kit-0.23.1-3  mingw-w64-x86_64-python2-2.7.11-1
                mingw-w64-x86_64-readline-6.3.008-1  mingw-w64-x86_64-tcl-8.6.4-2
                mingw-w64-x86_64-termcap-1.3.1-2  mingw-w64-x86_64-tk-8.6.4-2
                mingw-w64-x86_64-windows-default-manifest-6.4-2  mingw-w64-x86_64-zlib-1.2.8-9
                mingw-w64-x86_64-binutils-2.25.1-2  mingw-w64-x86_64-crt-git-5.0.0.4622.38d327f-1
                mingw-w64-x86_64-gcc-5.3.0-2  mingw-w64-x86_64-gcc-libs-5.3.0-2
                mingw-w64-x86_64-gdb-7.10.1-1  mingw-w64-x86_64-headers-git-5.0.0.4622.38d327f-1
                mingw-w64-x86_64-libmangle-git-5.0.0.4509.2e5a9a2-1
                mingw-w64-x86_64-libwinpthread-git-5.0.0.4573.628fdbf-1
                mingw-w64-x86_64-make-4.1.2351.a80a8b8-1  mingw-w64-x86_64-pkg-config-0.29-1
                mingw-w64-x86_64-tools-git-5.0.0.4592.90b8472-1
                mingw-w64-x86_64-winpthreads-git-5.0.0.4573.628fdbf-1

合計ダウンロード容量:   75.67 MiB
合計インストール容量:  489.96 MiB

:: インストールを行いますか？ [Y/n] y
:: パッケージを取得します ...
 mingw-w64-x86_64-libiconv-1.14-5-any                 600.2 KiB   913K/s 00:01 [##################] 100%
 mingw-w64-x86_64-bzip2-1.0.6-4-any                    69.4 KiB  9.68M/s 00:00 [##################] 100%
 mingw-w64-x86_64-zlib-1.2.8-9-any                    148.6 KiB   306K/s 00:00 [##################] 100%
 mingw-w64-x86_64-binutils-2.25.1-2-any                11.9 MiB  1178K/s 00:10 [##################] 100%
 mingw-w64-x86_64-headers-git-5.0.0.4622.38d327f...     5.0 MiB  1072K/s 00:05 [##################] 100%
 mingw-w64-x86_64-crt-git-5.0.0.4622.38d327f-1-any      2.7 MiB  1002K/s 00:03 [##################] 100%
 mingw-w64-x86_64-isl-0.15-1-any                      524.3 KiB   823K/s 00:01 [##################] 100%
 mingw-w64-x86_64-gmp-6.1.0-1-any                     477.1 KiB   752K/s 00:01 [##################] 100%
 mingw-w64-x86_64-mpfr-3.1.3.p0-2-any                 265.2 KiB   824K/s 00:00 [##################] 100%
 mingw-w64-x86_64-mpc-1.0.3-2-any                      62.3 KiB  10.1M/s 00:00 [##################] 100%
 mingw-w64-x86_64-libwinpthread-git-5.0.0.4573.6...    24.2 KiB  7.87M/s 00:00 [##################] 100%
 mingw-w64-x86_64-gcc-libs-5.3.0-2-any                541.9 KiB   681K/s 00:01 [##################] 100%
 mingw-w64-x86_64-windows-default-manifest-6.4-2-any 1484.0   B  1449K/s 00:00 [##################] 100%
 mingw-w64-x86_64-winpthreads-git-5.0.0.4573.628...    33.2 KiB  10.8M/s 00:00 [##################] 100%
 mingw-w64-x86_64-gcc-5.3.0-2-any                      25.1 MiB  1257K/s 00:20 [##################] 100%
 mingw-w64-x86_64-expat-2.1.0-6-any                   106.7 KiB  9.47M/s 00:00 [##################] 100%
 mingw-w64-x86_64-gettext-0.19.6-2-any                  3.0 MiB  1191K/s 00:03 [##################] 100%
 mingw-w64-x86_64-gdbm-1.11-3-any                     151.8 KiB   311K/s 00:00 [##################] 100%
 mingw-w64-x86_64-libffi-3.2.1-3-any                   34.5 KiB  11.2M/s 00:00 [##################] 100%
 mingw-w64-x86_64-libtre-git-r122.c2f5d13-4-any        69.2 KiB   427K/s 00:00 [##################] 100%
 mingw-w64-x86_64-libsystre-1.0.1-2-any                 9.3 KiB  9.11M/s 00:00 [##################] 100%
 mingw-w64-x86_64-ncurses-6.0.20150627-2-any         1794.9 KiB   704K/s 00:03 [##################] 100%
 mingw-w64-x86_64-libtasn1-4.7-1-any                  171.4 KiB   264K/s 00:01 [##################] 100%
 mingw-w64-x86_64-p11-kit-0.23.1-3-any                193.5 KiB   584K/s 00:00 [##################] 100%
 mingw-w64-x86_64-ca-certificates-20150426-2-any      382.1 KiB   789K/s 00:00 [##################] 100%
 mingw-w64-x86_64-openssl-1.0.2.f-1-any                 3.3 MiB   978K/s 00:04 [##################] 100%
 mingw-w64-x86_64-termcap-1.3.1-2-any                  12.6 KiB  12.3M/s 00:00 [##################] 100%
 mingw-w64-x86_64-readline-6.3.008-1-any              327.4 KiB   986K/s 00:00 [##################] 100%
 mingw-w64-x86_64-tcl-8.6.4-2-any                       2.8 MiB  1195K/s 00:02 [##################] 100%
 mingw-w64-x86_64-tk-8.6.4-2-any                     1869.4 KiB   891K/s 00:02 [##################] 100%
 mingw-w64-x86_64-python2-2.7.11-1-any                 10.8 MiB  1226K/s 00:09 [##################] 100%
 mingw-w64-x86_64-gdb-7.10.1-1-any                      2.7 MiB   969K/s 00:03 [##################] 100%
 mingw-w64-x86_64-libmangle-git-5.0.0.4509.2e5a9...    16.3 KiB  7.95M/s 00:00 [##################] 100%
 mingw-w64-x86_64-make-4.1.2351.a80a8b8-1-any         103.2 KiB  10.1M/s 00:00 [##################] 100%
 mingw-w64-x86_64-pkg-config-0.29-1-any               226.1 KiB   704K/s 00:00 [##################] 100%
 mingw-w64-x86_64-tools-git-5.0.0.4592.90b8472-1-any  247.0 KiB   772K/s 00:00 [##################] 100%
(36/36) キーリングのキーを確認                                                 [##################] 100%
(36/36) パッケージの整合性をチェック                                           [##################] 100%
(36/36) パッケージファイルのロード                                             [##################] 100%
(36/36) ファイルの衝突をチェック                                               [##################] 100%
(36/36) 空き容量を確認                                                         [##################] 100%
:: パッケージの変更を処理しています...
( 1/36) インストール mingw-w64-x86_64-libiconv                                 [##################] 100%
( 2/36) インストール mingw-w64-x86_64-bzip2                                    [##################] 100%
( 3/36) インストール mingw-w64-x86_64-zlib                                     [##################] 100%
( 4/36) インストール mingw-w64-x86_64-binutils                                 [##################] 100%
( 5/36) インストール mingw-w64-x86_64-headers-git                              [##################] 100%
( 6/36) インストール mingw-w64-x86_64-crt-git                                  [##################] 100%
( 7/36) インストール mingw-w64-x86_64-isl                                      [##################] 100%
( 8/36) インストール mingw-w64-x86_64-gmp                                      [##################] 100%
( 9/36) インストール mingw-w64-x86_64-mpfr                                     [##################] 100%
(10/36) インストール mingw-w64-x86_64-mpc                                      [##################] 100%
(11/36) インストール mingw-w64-x86_64-libwinpthread-git                        [##################] 100%
(12/36) インストール mingw-w64-x86_64-gcc-libs                                 [##################] 100%
(13/36) インストール mingw-w64-x86_64-windows-default-manifest                 [##################] 100%
(14/36) インストール mingw-w64-x86_64-winpthreads-git                          [##################] 100%
(15/36) インストール mingw-w64-x86_64-gcc                                      [##################] 100%
(16/36) インストール mingw-w64-x86_64-expat                                    [##################] 100%
(17/36) インストール mingw-w64-x86_64-gettext                                  [##################] 100%
(18/36) インストール mingw-w64-x86_64-gdbm                                     [##################] 100%
(19/36) インストール mingw-w64-x86_64-libffi                                   [##################] 100%
(20/36) インストール mingw-w64-x86_64-libtre-git                               [##################] 100%
(21/36) インストール mingw-w64-x86_64-libsystre                                [##################] 100%
(22/36) インストール mingw-w64-x86_64-ncurses                                  [##################] 100%
(23/36) インストール mingw-w64-x86_64-libtasn1                                 [##################] 100%
(24/36) インストール mingw-w64-x86_64-p11-kit                                  [##################] 100%
(25/36) インストール mingw-w64-x86_64-ca-certificates                          [##################] 100%
(26/36) インストール mingw-w64-x86_64-openssl                                  [##################] 100%
(27/36) インストール mingw-w64-x86_64-termcap                                  [##################] 100%
(28/36) インストール mingw-w64-x86_64-readline                                 [##################] 100%
(29/36) インストール mingw-w64-x86_64-tcl                                      [##################] 100%
(30/36) インストール mingw-w64-x86_64-tk                                       [##################] 100%
(31/36) インストール mingw-w64-x86_64-python2                                  [##################] 100%
(32/36) インストール mingw-w64-x86_64-gdb                                      [##################] 100%
(33/36) インストール mingw-w64-x86_64-libmangle-git                            [##################] 100%
(34/36) インストール mingw-w64-x86_64-make                                     [##################] 100%
(35/36) インストール mingw-w64-x86_64-pkg-config                               [##################] 100%
(36/36) インストール mingw-w64-x86_64-tools-git                                [##################] 100%
```

動作確認には `mingw64_shell.bat` の起動，または 環境変数 `MSYSTEM` に `MINGW64` をセットして shell を起動する。
[前回]紹介した [ConEmu] を使うのであれば以下のシーケンスで起動できる。

```text
set MSYSTEM=MINGW64 & C:\msys64\usr\bin\bash.exe --login -i
```

mingw64 版の shell が起動したところで gcc の動作確認を行う。

```text
$ gcc -v
Using built-in specs.
COLLECT_GCC=C:\msys64\mingw64\bin\gcc.exe
COLLECT_LTO_WRAPPER=C:/msys64/mingw64/bin/../lib/gcc/x86_64-w64-mingw32/5.3.0/lto-wrapper.exe
Target: x86_64-w64-mingw32
Configured with: ../gcc-5.3.0/configure --prefix=/mingw64 --with-local-prefix=/mingw64/local --build=x86_64-w64-mingw32 --host=x86_64-w64-mingw32 --target=x86_64-w64-mingw32 --with-native-system-header-dir=/mingw64/x86_64-w64-mingw32/include --libexecdir=/mingw64/lib --with-gxx-include-dir=/mingw64/include/c++/5.3.0 --enable-bootstrap --with-arch=x86-64 --with-tune=generic --enable-languages=c,lto,c++,objc,obj-c++,fortran,ada --enable-shared --enable-static --enable-libatomic --enable-threads=posix --enable-graphite --enable-fully-dynamic-string --enable-libstdcxx-time=yes --disable-libstdcxx-pch --disable-libstdcxx-debug --enable-version-specific-runtime-libs --disable-isl-version-check --enable-lto --enable-libgomp --disable-multilib --enable-checking=release --disable-rpath --disable-win32-registry --disable-nls --disable-werror --disable-symvers --with-libiconv --with-system-zlib --with-gmp=/mingw64 --with-mpfr=/mingw64 --with-mpc=/mingw64 --with-isl=/mingw64 --with-pkgversion='Rev2, Built by MSYS2 project' --with-bugurl=http://sourceforge.net/projects/msys2 --with-gnu-as --with-gnu-ld
Thread model: posix
gcc version 5.3.0 (Rev2, Built by MSYS2 project)
```

ターゲット名が `x86_64-w64-mingw32` となっていて非常に分かりにくいが， mingw32 版では `i686-w64-mingw32` となっているので，なんとか区別して欲しい。

## gcc のスレッドモデルと例外処理

[MSYS2] で提供される gcc のスレッドモデルと例外処理は以下のようになっているらしい。

{{< fig-quote title="MSYS2 / Tickets / #24 Mingw64 downloads through msys2?" link="http://sourceforge.net/p/msys2/tickets/24/" lang="en" >}}
<q>On MSYS2 I maintain only one configuration:<br>
32-bit: threads=posix, exceptions=dwarf<br>
64-bit: threads=posix, exceptions=seh</q>
{{< /fig-quote >}}

[MSYS2] のベースになっている [MinGW-w64] のスレッドモデルには win32 と posix が提供されている。 win32 は Windows ネイティブなスレッド関連関数を使っているため速いが， C++11 の thread, mutex, future はサポートしていないらしい。
一方，例外処理には SJLJ（SetJump/LongJump）， [DWARF](http://ja.wikipedia.org/wiki/DWARF)2， SEH（Structured Exception Handling）があり，この中では Windows ネイティブな SEH が一番速い。ただし SEH は 64bit gcc でしか対応していないため 32bit 環境では使えない[^w64]。

[^w64]: ちなみにオリジナルの [MinGW-w64] ではインストール時にスレッドモデルと例外処理を選択できる。

[MinGW-w64] のもうひとつの実装である [TDM-GCC] では，スレッドモデルには posix，例外処理には SJLJ（32bit）または SEH（64bit）が設定されているようだ。

{{< fig-quote title="TDM-GCC : Quirks" link="http://tdm-gcc.tdragon.net/quirks" lang="en" >}}
<q>TDM-GCC includes a pthreads emulation layer for Microsoft Windows systems, called "winpthreads". This lets you use std::thread and other C++11 concurrency features in your programs, out of the box.<br>
[...]<br>
64-bit binaries use "SEH", which stands for "Structured Exception Handling".<br>
[...]<br>
If you create a 32-bit binary with the TDM64 edition, it will use SJLJ exception handling. DW2 exception handling is not available in this edition.</q>
{{< /fig-quote >}}

[次回]は実際にコンパイルを行う。

[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
[前回]: {{< relref "remark/2016/02/gcc-msys2-1.md" >}} "MSYS2 による gcc 開発環境の構築 ― インストールから初期化処理まで"
[次回]: {{< relref "remark/2016/02/gcc-msys2-3.md" >}} "MSYS2 による gcc 開発環境の構築 ― pgpdump をビルドする"
[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
[MinGW-w64]: http://mingw-w64.org/ "Mingw-w64 - GCC for Windows 64 & 32 bits [mingw-w64]"
[TDM-GCC]: http://tdm-gcc.tdragon.net/ "TDM-GCC"
