+++
date = "2016-03-01T22:02:44+09:00"
update = "2017-03-24T17:57:25+09:00"
description = "いよいよ gcc をインストールする。"
draft = false
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

1. [MSYS2 のインストールから初期化処理まで]({{< relref "remark/2016/03/gcc-msys2-1.md" >}})
2. [gcc パッケージ群の導入]({{< relref "remark/2016/03/gcc-msys2-2.md" >}}) （← イマココ）
3. [pgpdump をビルドする]({{< relref "remark/2016/03/gcc-msys2-3.md" >}})

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

パッケージ (34) db-5.3.28-2  expat-2.1.0-2  gdbm-1.11-3  heimdal-1.5.3-8  libgdbm-1.11-3
                openssh-7.1p2-1  perl-5.22.0-2  perl-Authen-SASL-2.16-2  perl-Convert-BinHex-1.123-2
                perl-Encode-Locale-1.04-1  perl-Error-0.17024-1  perl-File-Listing-6.04-2
                perl-HTML-Parser-3.71-3  perl-HTML-Tagset-3.20-2  perl-HTTP-Cookies-6.01-2
                perl-HTTP-Daemon-6.01-2  perl-HTTP-Date-6.02-2  perl-HTTP-Message-6.06-2
                perl-HTTP-Negotiate-6.01-2  perl-IO-Socket-SSL-2.016-1  perl-IO-stringy-2.111-1
                perl-LWP-MediaTypes-6.02-2  perl-MIME-tools-5.506-1  perl-MailTools-2.14-1
                perl-Net-HTTP-6.09-1  perl-Net-SMTP-SSL-1.02-1  perl-Net-SSLeay-1.70-1
                perl-TermReadKey-2.33-1  perl-TimeDate-2.30-2  perl-URI-1.68-1
                perl-WWW-RobotRules-6.02-2  perl-libwww-6.13-1  vim-7.4.1415-2  git-2.7.1-1

合計ダウンロード容量:   24.75 MiB
合計インストール容量:  114.53 MiB

:: インストールを行いますか？ [Y/n] y
:: パッケージを取得します...
 expat-2.1.0-2-x86_64                                  13.1 KiB  12.8M/s 00:00 [##################] 100%
 vim-7.4.1415-2-x86_64                                  6.1 MiB  1217K/s 00:05 [##################] 100%
 heimdal-1.5.3-8-x86_64                               543.7 KiB   824K/s 00:01 [##################] 100%
 openssh-7.1p2-1-x86_64                               653.4 KiB   312K/s 00:02 [##################] 100%
 db-5.3.28-2-x86_64                                    41.7 KiB  10.2M/s 00:00 [##################] 100%
 libgdbm-1.11-3-x86_64                                 20.4 KiB  9.94M/s 00:00 [##################] 100%
 gdbm-1.11-3-x86_64                                   108.5 KiB   682K/s 00:00 [##################] 100%
 perl-5.22.0-2-x86_64                                  12.4 MiB  1170K/s 00:11 [##################] 100%
 perl-Error-0.17024-1-any                              17.1 KiB  8.33M/s 00:00 [##################] 100%
 perl-Authen-SASL-2.16-2-any                           42.4 KiB  8.28M/s 00:00 [##################] 100%
 perl-Encode-Locale-1.04-1-any                          9.7 KiB  9.46M/s 00:00 [##################] 100%
 perl-HTTP-Date-6.02-2-any                              8.6 KiB  8.42M/s 00:00 [##################] 100%
 perl-File-Listing-6.04-2-any                           7.7 KiB  0.00B/s 00:00 [##################] 100%
 perl-HTML-Tagset-3.20-2-any                           10.3 KiB  10.0M/s 00:00 [##################] 100%
 perl-HTML-Parser-3.71-3-x86_64                        76.9 KiB  10.7M/s 00:00 [##################] 100%
 perl-LWP-MediaTypes-6.02-2-any                        18.0 KiB  8.81M/s 00:00 [##################] 100%
 perl-URI-1.68-1-any                                   75.6 KiB  9.22M/s 00:00 [##################] 100%
 perl-HTTP-Message-6.06-2-any                          71.3 KiB  9.94M/s 00:00 [##################] 100%
 perl-HTTP-Cookies-6.01-2-any                          20.4 KiB  20.0M/s 00:00 [##################] 100%
 perl-HTTP-Daemon-6.01-2-any                           14.2 KiB  0.00B/s 00:00 [##################] 100%
 perl-HTTP-Negotiate-6.01-2-any                        11.4 KiB  0.00B/s 00:00 [##################] 100%
 perl-Net-HTTP-6.09-1-any                              19.8 KiB  19.3M/s 00:00 [##################] 100%
 perl-WWW-RobotRules-6.02-2-any                        12.2 KiB  12.0M/s 00:00 [##################] 100%
 perl-libwww-6.13-1-any                               122.2 KiB   367K/s 00:00 [##################] 100%
 perl-TimeDate-2.30-2-any                              35.9 KiB  8.77M/s 00:00 [##################] 100%
 perl-MailTools-2.14-1-any                             58.4 KiB  9.50M/s 00:00 [##################] 100%
 perl-IO-stringy-2.111-1-any                           52.6 KiB  10.3M/s 00:00 [##################] 100%
 perl-Convert-BinHex-1.123-2-any                       30.1 KiB  9.79M/s 00:00 [##################] 100%
 perl-MIME-tools-5.506-1-any                          180.4 KiB   367K/s 00:00 [##################] 100%
 perl-Net-SSLeay-1.70-1-x86_64                        191.2 KiB   590K/s 00:00 [##################] 100%
 perl-IO-Socket-SSL-2.016-1-any                       112.5 KiB   208K/s 00:01 [##################] 100%
 perl-Net-SMTP-SSL-1.02-1-any                           3.5 KiB  22.2K/s 00:00 [##################] 100%
 perl-TermReadKey-2.33-1-x86_64                        20.9 KiB  18.7K/s 00:01 [##################] 100%
 git-2.7.1-1-x86_64                                     3.6 MiB   448K/s 00:08 [##################] 100%
(34/34) キーリングのキーを確認                                                 [##################] 100%
(34/34) パッケージの整合性をチェック                                           [##################] 100%
(34/34) パッケージファイルのロード                                             [##################] 100%
(34/34) ファイルの衝突をチェック                                               [##################] 100%
(34/34) 空き容量を確認                                                         [##################] 100%
:: パッケージの変更を処理しています...
( 1/34) インストール expat                                                     [##################] 100%
( 2/34) インストール vim                                                       [##################] 100%
( 3/34) インストール heimdal                                                   [##################] 100%
( 4/34) インストール openssh                                                   [##################] 100%
( 5/34) インストール db                                                        [##################] 100%
( 6/34) インストール libgdbm                                                   [##################] 100%
( 7/34) インストール gdbm                                                      [##################] 100%
( 8/34) インストール perl                                                      [##################] 100%
( 9/34) インストール perl-Error                                                [##################] 100%
(10/34) インストール perl-Authen-SASL                                          [##################] 100%
(11/34) インストール perl-Encode-Locale                                        [##################] 100%
(12/34) インストール perl-HTTP-Date                                            [##################] 100%
(13/34) インストール perl-File-Listing                                         [##################] 100%
(14/34) インストール perl-HTML-Tagset                                          [##################] 100%
(15/34) インストール perl-HTML-Parser                                          [##################] 100%
(16/34) インストール perl-LWP-MediaTypes                                       [##################] 100%
(17/34) インストール perl-URI                                                  [##################] 100%
(18/34) インストール perl-HTTP-Message                                         [##################] 100%
(19/34) インストール perl-HTTP-Cookies                                         [##################] 100%
(20/34) インストール perl-HTTP-Daemon                                          [##################] 100%
(21/34) インストール perl-HTTP-Negotiate                                       [##################] 100%
(22/34) インストール perl-Net-HTTP                                             [##################] 100%
(23/34) インストール perl-WWW-RobotRules                                       [##################] 100%
(24/34) インストール perl-libwww                                               [##################] 100%
perl-libwww の提案パッケージ
    perl-LWP-Protocol-HTTPS: for https:// url schemes
(25/34) インストール perl-TimeDate                                             [##################] 100%
(26/34) インストール perl-MailTools                                            [##################] 100%
(27/34) インストール perl-IO-stringy                                           [##################] 100%
(28/34) インストール perl-Convert-BinHex                                       [##################] 100%
module test... pass.
(29/34) インストール perl-MIME-tools                                           [##################] 100%
(30/34) インストール perl-Net-SSLeay                                           [##################] 100%
(31/34) インストール perl-IO-Socket-SSL                                        [##################] 100%
(32/34) インストール perl-Net-SMTP-SSL                                         [##################] 100%
(33/34) インストール perl-TermReadKey                                          [##################] 100%
(34/34) インストール git                                                       [##################] 100%
git の提案パッケージ
    python2: various helper scripts
    subversion: git svn
```

てか，本当に最小限しか入ってないんだな（笑） 提案パッケージについては，とりあえずスルーする。

### MSYS2 のリポジトリ

`/etc/pacman.conf` を見ると MSYS2 では `mingw32`, `mingw64`, `msys` の3つのリポジトリを管理していることが分かる。

| リポジトリ名 | 参照ファイル |
|:-------------|:-------------|
| `mingw32`    | `/etc/pacman.d/mirrorlist.mingw32` |
| `mingw64`    | `/etc/pacman.d/mirrorlist.mingw64` |
| `msys`       | `/etc/pacman.d/mirrorlist.msys` |

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

では実際にリポジトリの中を覗いてみよう。

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

どれを使うかはユーザ次第だが，今回は [MSYS2] 内ではなく Windows 環境で動作するアプリケーションを作りたいので， `mingw32` および `mingw64` リポジトリのものを導入する。

- [Windowsでgccなどを利用できるMSYS2の環境設定など - Qiita](http://qiita.com/chromabox/items/fd07bad3f426101fc4a6)

## gcc のインストール

では gcc のインストールを始めよう。
今回は以下のパッケージグループを導入する。

- `base-devel`
- `mingw-w64-i686-toolchain`
- `mingw-w64-x86_64-toolchain`

ものすごく分かりにくいが `mingw-w64-x86_64-toolchain` が 64bit 用である。
ログがかなり長くなるので最初の部分だけ。

```text
$ pacman -S base-devel mingw-w64-i686-toolchain mingw-w64-x86_64-toolchain
:: 51 のパッケージがグループ base-devel にあります:
:: リポジトリ msys
   1) asciidoc  2) autoconf  3) autoconf2.13  4) autogen  5) automake-wrapper  6) automake1.10
   7) automake1.11  8) automake1.12  9) automake1.13  10) automake1.14  11) automake1.15
   12) automake1.6  13) automake1.7  14) automake1.8  15) automake1.9  16) bison  17) diffstat
   18) diffutils  19) dos2unix  20) file  21) flex  22) gawk  23) gdb  24) gettext  25) gperf
   26) grep  27) groff  28) help2man  29) intltool  30) lemon  31) libtool  32) libunrar  33) m4
   34) make  35) man-db  36) pacman  37) patch  38) patchutils  39) perl  40) pkg-config  41) pkgfile
   42) rcs  43) scons  44) sed  45) swig  46) texinfo  47) texinfo-tex  48) ttyrec  49) unrar
   50) wget  51) xmlto

選択して下さい (デフォルト=all):
警告: file-5.25-1 は最新です -- 再インストール
警告: flex-2.6.0-1 は最新です -- 再インストール
警告: gawk-4.1.3-1 は最新です -- 再インストール
警告: gettext-0.19.6-1 は最新です -- 再インストール
警告: grep-2.22-1 は最新です -- 再インストール
警告: m4-1.4.17-4 は最新です -- 再インストール
警告: pacman-5.0.1.6388.dcb2397-1 は最新です -- 再インストール
警告: perl-5.22.0-2 は最新です -- 再インストール
警告: pkgfile-15-1 は最新です -- 再インストール
警告: sed-4.2.2-2 は最新です -- 再インストール
:: 16 のパッケージがグループ mingw-w64-i686-toolchain にあります:
:: リポジトリ mingw32
   1) mingw-w64-i686-binutils  2) mingw-w64-i686-crt-git  3) mingw-w64-i686-gcc
   4) mingw-w64-i686-gcc-ada  5) mingw-w64-i686-gcc-fortran  6) mingw-w64-i686-gcc-libgfortran
   7) mingw-w64-i686-gcc-libs  8) mingw-w64-i686-gcc-objc  9) mingw-w64-i686-gdb
   10) mingw-w64-i686-headers-git  11) mingw-w64-i686-libmangle-git
   12) mingw-w64-i686-libwinpthread-git  13) mingw-w64-i686-make  14) mingw-w64-i686-pkg-config
   15) mingw-w64-i686-tools-git  16) mingw-w64-i686-winpthreads-git

選択して下さい (デフォルト=all): 1 2 3 7 9 10 11 12 13 14 15 16
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

パッケージ (143) docbook-xml-4.5-2  docbook-xsl-1.78.1-3  glib2-2.44.1-1  libassuan-2.4.2-1
                 libgc-7.2.d-1  libgcrypt-1.6.4-1  libgpg-error-1.21-1  libgpgme-1.6.0-1
                 libguile-2.0.11-3  libltdl-2.4.6-1  libpipeline-1.4.0-1  libunistring-0.9.6-1
                 libxslt-1.1.28-7  mingw-w64-i686-bzip2-1.0.6-5
                 mingw-w64-i686-ca-certificates-20150426-2
                 mingw-w64-i686-expat-2.1.0-6  mingw-w64-i686-gdbm-1.11-3
                 mingw-w64-i686-gettext-0.19.6-2  mingw-w64-i686-gmp-6.1.0-1
                 mingw-w64-i686-isl-0.15-1  mingw-w64-i686-libffi-3.2.1-3
                 mingw-w64-i686-libiconv-1.14-5  mingw-w64-i686-libsystre-1.0.1-2
                 mingw-w64-i686-libtasn1-4.7-1  mingw-w64-i686-libtre-git-r122.c2f5d13-4
                 mingw-w64-i686-mpc-1.0.3-2  mingw-w64-i686-mpfr-3.1.3.p0-2
                 mingw-w64-i686-ncurses-6.0.20160220-1  mingw-w64-i686-openssl-1.0.2.f-1
                 mingw-w64-i686-p11-kit-0.23.1-3  mingw-w64-i686-python2-2.7.11-2
                 mingw-w64-i686-readline-6.3.008-1  mingw-w64-i686-tcl-8.6.4-2
                 mingw-w64-i686-termcap-1.3.1-2  mingw-w64-i686-tk-8.6.4-2
                 mingw-w64-i686-windows-default-manifest-6.4-2  mingw-w64-i686-zlib-1.2.8-9
                 mingw-w64-x86_64-bzip2-1.0.6-5  mingw-w64-x86_64-ca-certificates-20150426-2
                 mingw-w64-x86_64-expat-2.1.0-6  mingw-w64-x86_64-gdbm-1.11-3
                 mingw-w64-x86_64-gettext-0.19.6-2  mingw-w64-x86_64-gmp-6.1.0-1
                 mingw-w64-x86_64-isl-0.15-1  mingw-w64-x86_64-libffi-3.2.1-3
                 mingw-w64-x86_64-libiconv-1.14-5  mingw-w64-x86_64-libsystre-1.0.1-2
                 mingw-w64-x86_64-libtasn1-4.7-1  mingw-w64-x86_64-libtre-git-r122.c2f5d13-4
                 mingw-w64-x86_64-mpc-1.0.3-2  mingw-w64-x86_64-mpfr-3.1.3.p0-2
                 mingw-w64-x86_64-ncurses-6.0.20160220-1  mingw-w64-x86_64-openssl-1.0.2.f-1
                 mingw-w64-x86_64-p11-kit-0.23.1-3  mingw-w64-x86_64-python2-2.7.11-2
                 mingw-w64-x86_64-readline-6.3.008-1  mingw-w64-x86_64-tcl-8.6.4-2
                 mingw-w64-x86_64-termcap-1.3.1-2  mingw-w64-x86_64-tk-8.6.4-2
                 mingw-w64-x86_64-windows-default-manifest-6.4-2  mingw-w64-x86_64-zlib-1.2.8-9
                 perl-Locale-Gettext-1.05-4  perl-Module-Build-0.4212-1  perl-Test-Pod-1.50-1
                 perl-XML-Parser-2.44-1  perl-YAML-Syck-1.29-1  python2-2.7.10-1  tar-1.28-3
                 asciidoc-8.6.9-4  autoconf-2.69-3  autoconf2.13-2.13-2  autogen-5.18.4-2
                 automake-wrapper-10-1  automake1.10-1.10.3-3  automake1.11-1.11.6-3
                 automake1.12-1.12.6-3  automake1.13-1.13.4-4  automake1.14-1.14.1-3
                 automake1.15-1.15-2  automake1.6-1.6.3-2  automake1.7-1.7.9-2  automake1.8-1.8.5-3
                 automake1.9-1.9.6-2  bison-3.0.4-1  diffstat-1.58-1  diffutils-3.3-3
                 dos2unix-7.3.3-1  file-5.25-1  flex-2.6.0-1  gawk-4.1.3-1  gdb-7.9-1
                 gettext-0.19.6-1  gperf-3.0.4-3  grep-2.22-1  groff-1.22.3-1  help2man-1.47.3-1
                 intltool-0.51.0-2  lemon-3.8.7.0-1  libtool-2.4.6-1  libunrar-5.3.7-1  m4-1.4.17-4
                 make-4.1-4  man-db-2.7.4-1  mingw-w64-i686-binutils-2.25.1-2
                 mingw-w64-i686-crt-git-5.0.0.4627.d55f21d-1  mingw-w64-i686-gcc-5.3.0-2
                 mingw-w64-i686-gcc-libs-5.3.0-2  mingw-w64-i686-gdb-7.11-1
                 mingw-w64-i686-headers-git-5.0.0.4627.6baeb9d-1
                 mingw-w64-i686-libmangle-git-5.0.0.4509.2e5a9a2-1
                 mingw-w64-i686-libwinpthread-git-5.0.0.4573.628fdbf-1
                 mingw-w64-i686-make-4.1.2351.a80a8b8-1  mingw-w64-i686-pkg-config-0.29-1
                 mingw-w64-i686-tools-git-5.0.0.4592.90b8472-1
                 mingw-w64-i686-winpthreads-git-5.0.0.4573.628fdbf-1
                 mingw-w64-x86_64-binutils-2.25.1-2  mingw-w64-x86_64-crt-git-5.0.0.4627.03684c4-1
                 mingw-w64-x86_64-gcc-5.3.0-2  mingw-w64-x86_64-gcc-libs-5.3.0-2
                 mingw-w64-x86_64-gdb-7.11-1  mingw-w64-x86_64-headers-git-5.0.0.4627.53be55d-1
                 mingw-w64-x86_64-libmangle-git-5.0.0.4509.2e5a9a2-1
                 mingw-w64-x86_64-libwinpthread-git-5.0.0.4573.628fdbf-1
                 mingw-w64-x86_64-make-4.1.2351.a80a8b8-1  mingw-w64-x86_64-pkg-config-0.29-1
                 mingw-w64-x86_64-tools-git-5.0.0.4592.90b8472-1
                 mingw-w64-x86_64-winpthreads-git-5.0.0.4573.628fdbf-1  pacman-5.0.1.6388.dcb2397-1
                 patch-2.7.5-1  patchutils-0.3.3-2  perl-5.22.0-2  pkg-config-0.28-2  pkgfile-15-1
                 rcs-5.9.4-1  scons-2.4.1-2  sed-4.2.2-2  swig-3.0.7-1  texinfo-6.0-1
                 texinfo-tex-6.0-1  ttyrec-1.0.8-1  unrar-5.3.7-1  wget-1.17.1-2  xmlto-0.0.26-2

合計ダウンロード容量:   186.38 MiB
合計インストール容量:  1221.31 MiB
最終的なアップグレード容量:  1124.41 MiB

:: インストールを行いますか？ [Y/n] y
```

`base-devel` は面倒なので全てインストール。
一部が再インストールされたが問題はなさそう[^bd]。
gcc を含む toolchain については Ada, FORTRAN, Object-C の言語パッケージを除いたものをインストールした。

[^bd]: ていうか，最新版が入ってるのが分かるなら，わざわざ再インストールする必要はないと思うのだが...

### gcc の動作確認

まずは 32bit 版の動作確認。
`mingw32.exe` を起動，または 環境変数 `MSYSTEM` に `MINGW32` をセットして shell を起動する。
[前回]紹介した [ConEmu] を使うのであれば以下のシーケンスで起動できる。

```text
set MSYSTEM=MINGW32 & chcp 65001 & C:\msys64\usr\bin\bash.exe --login -i
```

gcc を起動して動作確認。

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

次に 64bit 版の動作確認。
動作確認には `mingw64.exe` を起動，または 環境変数 `MSYSTEM` に `MINGW64` をセットして shell を起動する。
[前回]紹介した [ConEmu] を使うのであれば以下のシーケンスで起動できる。

```text
set MSYSTEM=MINGW64 & chcp 65001 & C:\msys64\usr\bin\bash.exe --login -i
```

gcc を起動して動作確認。

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

[次回]は実際にビルドを行う。

[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
[前回]: {{< relref "remark/2016/03/gcc-msys2-1.md" >}} "MSYS2 による gcc 開発環境の構築 ― インストールから初期化処理まで"
[次回]: {{< relref "remark/2016/03/gcc-msys2-3.md" >}} "MSYS2 による gcc 開発環境の構築 ― pgpdump をビルドする"
[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
[MinGW-w64]: http://mingw-w64.org/ "Mingw-w64 - GCC for Windows 64 & 32 bits [mingw-w64]"
[TDM-GCC]: http://tdm-gcc.tdragon.net/ "TDM-GCC"
