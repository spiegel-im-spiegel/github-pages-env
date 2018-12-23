+++
date = "2016-03-01T22:02:44+09:00"
update = "2018-03-31T17:40:56+09:00"
description = "いよいよ gcc をインストールする。"
draft = false
tags = ["msys2", "mingw", "gcc", "tools", "windows", "pacman"]
title = "MSYS2 による gcc 開発環境の構築 ― gcc パッケージ群の導入"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1. [MSYS2 のインストールから初期化処理まで]({{< ref "/remark/2016/03/gcc-msys2-1.md" >}})
2. [gcc パッケージ群の導入]({{< ref "/remark/2016/03/gcc-msys2-2.md" >}}) （← イマココ）
3. [pgpdump をビルドする]({{< ref "/remark/2016/03/gcc-msys2-3.md" >}})

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

### MSYS2 のリポジトリ

`/etc/pacman.conf` を見ると MSYS2 では `mingw32`, `mingw64`, `msys` の3つのリポジトリを管理していることが分かる。

| リポジトリ名 | 参照ファイル                       |
|:------------ |:---------------------------------- |
| `mingw32`    | `/etc/pacman.d/mirrorlist.mingw32` |
| `mingw64`    | `/etc/pacman.d/mirrorlist.mingw64` |
| `msys`       | `/etc/pacman.d/mirrorlist.msys`    |

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
mingw32 mingw-w64-i686-gcc 7.3.0-1
mingw32 mingw-w64-i686-gcc-ada 7.3.0-1
mingw32 mingw-w64-i686-gcc-fortran 7.3.0-1
mingw32 mingw-w64-i686-gcc-libgfortran 7.3.0-1
mingw32 mingw-w64-i686-gcc-libs 7.3.0-1
mingw32 mingw-w64-i686-gcc-objc 7.3.0-1
mingw64 mingw-w64-x86_64-gcc 7.3.0-1
mingw64 mingw-w64-x86_64-gcc-ada 7.3.0-1
mingw64 mingw-w64-x86_64-gcc-fortran 7.3.0-1
mingw64 mingw-w64-x86_64-gcc-libgfortran 7.3.0-1
mingw64 mingw-w64-x86_64-gcc-libs 7.3.0-1
mingw64 mingw-w64-x86_64-gcc-objc 7.3.0-1
msys gcc 6.4.0-3
msys gcc-fortran 6.4.0-3
msys gcc-libs 6.4.0-3 [インストール済み]
msys mingw-w64-cross-gcc 6.4.0-2
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
コマンドラインは以下の通り（ログがかなり長くなるので省略）。

```text
$ pacman -S base-devel mingw-w64-i686-toolchain mingw-w64-x86_64-toolchain
```

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
COLLECT_LTO_WRAPPER=C:/msys64/mingw32/bin/../lib/gcc/i686-w64-mingw32/7.3.0/lto-wrapper.exe
Target: i686-w64-mingw32
Configured with: ../gcc-7.3.0/configure --prefix=/mingw32 --with-local-prefix=/mingw32/local --build=i686-w64-mingw32 --host=i686-w64-mingw32 --target=i686-w64-mingw32 --with-native-system-header-dir=/mingw32/i686-w64-mingw32/include --libexecdir=/mingw32/lib --enable-bootstrap --with-arch=i686 --with-tune=generic --enable-languages=c,lto,c++,objc,obj-c++,fortran,ada --enable-shared --enable-static --enable-libatomic --enable-threads=posix --enable-graphite --enable-fully-dynamic-string --enable-libstdcxx-time=yes --enable-libstdcxx-filesystem-ts=yes --disable-libstdcxx-pch --disable-libstdcxx-debug --disable-isl-version-check --enable-lto --enable-libgomp --disable-multilib --enable-checking=release --disable-rpath --disable-win32-registry --disable-nls --disable-werror --disable-symvers --with-libiconv --with-system-zlib --with-gmp=/mingw32 --with-mpfr=/mingw32 --with-mpc=/mingw32 --with-isl=/mingw32 --with-pkgversion='Rev1, Built by MSYS2 project' --with-bugurl=https://sourceforge.net/projects/msys2 --with-gnu-as --with-gnu-ld --disable-sjlj-exceptions --with-dwarf2
Thread model: posix
gcc version 7.3.0 (Rev1, Built by MSYS2 project)
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
COLLECT_LTO_WRAPPER=C:/msys64/mingw64/bin/../lib/gcc/x86_64-w64-mingw32/7.3.0/lto-wrapper.exe
Target: x86_64-w64-mingw32
Configured with: ../gcc-7.3.0/configure --prefix=/mingw64 --with-local-prefix=/mingw64/local --build=x86_64-w64-mingw32 --host=x86_64-w64-mingw32 --target=x86_64-w64-mingw32 --with-native-system-header-dir=/mingw64/x86_64-w64-mingw32/include --libexecdir=/mingw64/lib --enable-bootstrap --with-arch=x86-64 --with-tune=generic --enable-languages=c,lto,c++,objc,obj-c++,fortran,ada --enable-shared --enable-static --enable-libatomic --enable-threads=posix --enable-graphite --enable-fully-dynamic-string --enable-libstdcxx-time=yes --enable-libstdcxx-filesystem-ts=yes --disable-libstdcxx-pch --disable-libstdcxx-debug --disable-isl-version-check --enable-lto --enable-libgomp --disable-multilib --enable-checking=release --disable-rpath --disable-win32-registry --disable-nls --disable-werror --disable-symvers --with-libiconv --with-system-zlib --with-gmp=/mingw64 --with-mpfr=/mingw64 --with-mpc=/mingw64 --with-isl=/mingw64 --with-pkgversion='Rev1, Built by MSYS2 project' --with-bugurl=https://sourceforge.net/projects/msys2 --with-gnu-as --with-gnu-ld
Thread model: posix
gcc version 7.3.0 (Rev1, Built by MSYS2 project)
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
[前回]: {{< ref "/remark/2016/03/gcc-msys2-1.md" >}} "MSYS2 による gcc 開発環境の構築 ― インストールから初期化処理まで"
[次回]: {{< ref "/remark/2016/03/gcc-msys2-3.md" >}} "MSYS2 による gcc 開発環境の構築 ― pgpdump をビルドする"
[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
[MinGW-w64]: http://mingw-w64.org/ "Mingw-w64 - GCC for Windows 64 & 32 bits [mingw-w64]"
[TDM-GCC]: http://tdm-gcc.tdragon.net/ "TDM-GCC"
