+++
date = "2016-03-01T22:02:47+09:00"
update = "2016-03-08T13:57:12+09:00"
description = "前回で gcc を導入できたので，実際にビルドを試してみる。今回はターゲットとして pgpdump を用いる。"
draft = false
tags = ["msys2", "gcc", "tools"]
title = "MSYS2 による gcc 開発環境の構築 ― pgpdump をビルドする"

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
2. [gcc パッケージ群の導入]({{< relref "remark/2016/03/gcc-msys2-2.md" >}})
3. [pgpdump をビルドする]({{< relref "remark/2016/03/gcc-msys2-3.md" >}}) （← イマココ）

[前回]で gcc を導入できたので，実際にビルドを試してみる。
今回はターゲットとして [pgpdump] を用いる。

## pgpdump

[pgpdump] は山本和彦さんによる [OpenPGP](http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format") パケットの[視覚化ツール](http://www.mew.org/~kazu/proj/pgpdump/ja/)。
ソースコードは [GitHub で公開](https://github.com/kazu-yamamoto/pgpdump "kazu-yamamoto/pgpdump: A PGP packet visualizer")されているが， UNIX 系のプラットフォームを前提に作られているため [MSYS2] 上でビルドを行う。

まずは [pgpdump] のソースコードをダウンロードする。

```text
$ git clone https://github.com/kazu-yamamoto/pgpdump.git
Cloning into 'pgpdump'...
remote: Counting objects: 492, done.
Receiving objects:  59% (291remote: Total 492 (delta 0), reused 0 (delta 0), pack-reused 492 92
Receiving objects: 100% (492/492), 180.29 KiB | 0 bytes/s, done.
Resolving deltas: 100% (320/320), done.
Checking connectivity... done.
```

## 32bit 版のビルド

[pgpdump] のビルド手順は `configure` を実行した後 make を実行する[^conf]。
まずは何も考えずに `configure` の実行してみる。

[^conf]: [前回]インストールした `base-devel` パッケージグループは今回の [pgpdump] ビルドには必要なく，インストールしなくても問題ない。まぁ今回は単純なビルドだし。

```text
$ cd pgpdump/

$ ./configure
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.exe
checking for suffix of executables... .exe
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking for inflate in -lz... yes
checking for BZ2_bzBuffToBuffDecompress in -lbz2... yes
checking how to run the C preprocessor... gcc -E
checking for grep that handles long lines and -e... /usr/bin/grep
checking for egrep... /usr/bin/grep -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for unistd.h... (cached) yes
checking sys/time.h usability... yes
checking sys/time.h presence... yes
checking for sys/time.h... yes
checking unixlib/local.h usability... no
checking unixlib/local.h presence... no
checking for unixlib/local.h... no
checking whether time.h and sys/time.h may both be included... yes
checking whether struct tm is in sys/time.h or time.h... time.h
checking for struct tm.tm_zone... no
checking whether tzname is declared... yes
checking for tzname... yes
configure: creating ./config.status
config.status: creating Makefile
config.status: WARNING:  'Makefile.in' seems to ignore the --datarootdir setting
config.status: creating config.h
```

[pgpdump] ではパケット内の圧縮データを扱うため `libz` および `libbz2` が必要となるが，ちゃんと認識しているようだ。
これによって作成された `Makefile` がこれ。

```text
prefix = /usr/local
exec_prefix = ${prefix}
bindir = ${exec_prefix}/bin
mandir = ${prefix}/share/man
LIBS = -lbz2 -lz
CFLAGS  = -g -O2 -O -Wall
LDFLAGS =
VERSION = `git tag | tail -1 | sed -e 's/v//'`

RM = rm -f
INSTALL  = install

INCS = pgpdump.h
SRCS = pgpdump.c types.c tagfuncs.c packet.c subfunc.c signature.c keys.c \
       buffer.c uatfunc.c
OBJS = pgpdump.o types.o tagfuncs.o packet.o subfunc.o signature.o keys.o \
       buffer.o uatfunc.o
PROG = pgpdump

MAN  = pgpdump.1

CNF = config.h config.status config.cache config.log
MKF = Makefile

.c.o:
	$(CC) -c $(CFLAGS) $<

all: $(PROG)

$(PROG): $(OBJS)
	$(CC) $(CFLAGS) -o $(PROG) $(OBJS) $(LIBS) $(LDFLAGS)

clean:
	$(RM) $(OBJS) $(PROG)

distclean:
	$(RM) $(OBJS) $(PROG) $(CNF) $(MKF)

install: all
	$(INSTALL) -d $(DESTDIR)$(bindir)
	$(INSTALL) -cp -pm755 $(PROG) $(DESTDIR)$(bindir)
	$(INSTALL) -d $(DESTDIR)$(mandir)/man1
	$(INSTALL) -cp -pm644 $(MAN) $(DESTDIR)$(mandir)/man1

archive:
	git archive master -o ~/pgpdump-$(VERSION).tar --prefix=pgpdump-$(VERSION)/
	gzip ~/pgpdump-$(VERSION).tar
```

この時点での問題は以下のとおり。

1. `prefix` が `/usr/local` になっている。このままでもエラーにはならないが，今回は 32bit 版と 64bit 版を分けたいので `/mingw32` としたい
1. リンク時のオプション（`LDFLAGS`）に `-static` がないため，このままビルドすると DLL に依存する構成になってしまう

これらについては `configure` に情報を渡せばよい。

```text
$ ./configure --prefix=/mingw32 LDFLAGS=-static
```

これで `Makefile` は以下のようになる。

```text
prefix = /mingw32
exec_prefix = ${prefix}
bindir = ${exec_prefix}/bin
mandir = ${prefix}/share/man
LIBS = -lbz2 -lz
CFLAGS  = -g -O2 -O -Wall
LDFLAGS = -static
VERSION = `git tag | tail -1 | sed -e 's/v//'`

RM = rm -f
INSTALL  = install

INCS = pgpdump.h
SRCS = pgpdump.c types.c tagfuncs.c packet.c subfunc.c signature.c keys.c \
       buffer.c uatfunc.c
OBJS = pgpdump.o types.o tagfuncs.o packet.o subfunc.o signature.o keys.o \
       buffer.o uatfunc.o
PROG = pgpdump

MAN  = pgpdump.1

CNF = config.h config.status config.cache config.log
MKF = Makefile

.c.o:
	$(CC) -c $(CFLAGS) $<

all: $(PROG)

$(PROG): $(OBJS)
	$(CC) $(CFLAGS) -o $(PROG) $(OBJS) $(LIBS) $(LDFLAGS)

clean:
	$(RM) $(OBJS) $(PROG)

distclean:
	$(RM) $(OBJS) $(PROG) $(CNF) $(MKF)

install: all
	$(INSTALL) -d $(DESTDIR)$(bindir)
	$(INSTALL) -cp -pm755 $(PROG) $(DESTDIR)$(bindir)
	$(INSTALL) -d $(DESTDIR)$(mandir)/man1
	$(INSTALL) -cp -pm644 $(MAN) $(DESTDIR)$(mandir)/man1

archive:
	git archive master -o ~/pgpdump-$(VERSION).tar --prefix=pgpdump-$(VERSION)/
	gzip ~/pgpdump-$(VERSION).tar
```

では make を実行しよう。

```text
$ make
cc -c -g -O2 -O -Wall pgpdump.c
cc -c -g -O2 -O -Wall types.c
cc -c -g -O2 -O -Wall tagfuncs.c
cc -c -g -O2 -O -Wall packet.c
cc -c -g -O2 -O -Wall subfunc.c
cc -c -g -O2 -O -Wall signature.c
cc -c -g -O2 -O -Wall keys.c
cc -c -g -O2 -O -Wall buffer.c
cc -c -g -O2 -O -Wall uatfunc.c
cc -g -O2 -O -Wall -o pgpdump pgpdump.o types.o tagfuncs.o packet.o subfunc.o signature.o keys.o buffer.o uatfunc.o -lbz2 -lz  -static

$ strip pgpdump.exe
```

できた実行ファイルをコマンドプロンプトから起動してみる。

```text
C:>pgpdump.exe -v
pgpdump.exe version 0.29, Copyright (C) 1998-2014 Kazu Yamamoto
```

うまくいったようである。

## 64bit 版のビルド

32bit 版で要領は分かったので一気に終わらせてしまおう。

```text
$ git clone https://github.com/kazu-yamamoto/pgpdump.git
Cloning into 'pgpdump'...
remote: Counting objects: 492, done.
Receiving objects:  62% (306remote: Total 492 (delta 0), reused 0 (delta 0), pack-reused 492/92
Receiving objects: 100% (492/492), 180.29 KiB | 0 bytes/s, done.
Resolving deltas: 100% (320/320), done.
Checking connectivity... done.

$ cd pgpdump/

$ ./configure --prefix=/mingw64 LDFLAGS=-static
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.exe
checking for suffix of executables... .exe
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking for inflate in -lz... yes
checking for BZ2_bzBuffToBuffDecompress in -lbz2... yes
checking how to run the C preprocessor... gcc -E
checking for grep that handles long lines and -e... /usr/bin/grep
checking for egrep... /usr/bin/grep -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for unistd.h... (cached) yes
checking sys/time.h usability... yes
checking sys/time.h presence... yes
checking for sys/time.h... yes
checking unixlib/local.h usability... no
checking unixlib/local.h presence... no
checking for unixlib/local.h... no
checking whether time.h and sys/time.h may both be included... yes
checking whether struct tm is in sys/time.h or time.h... time.h
checking for struct tm.tm_zone... no
checking whether tzname is declared... yes
checking for tzname... yes
configure: creating ./config.status
config.status: creating Makefile
config.status: WARNING:  'Makefile.in' seems to ignore the --datarootdir setting
config.status: creating config.h

$ make
cc -c -g -O2 -O -Wall pgpdump.c
cc -c -g -O2 -O -Wall types.c
cc -c -g -O2 -O -Wall tagfuncs.c
cc -c -g -O2 -O -Wall packet.c
cc -c -g -O2 -O -Wall subfunc.c
cc -c -g -O2 -O -Wall signature.c
cc -c -g -O2 -O -Wall keys.c
cc -c -g -O2 -O -Wall buffer.c
cc -c -g -O2 -O -Wall uatfunc.c
cc -g -O2 -O -Wall -o pgpdump pgpdump.o types.o tagfuncs.o packet.o subfunc.o signature.o keys.o buffer.o uatfunc.o -lbz2 -lz  -static

$ strip pgpdump.exe
```

これもコマンドプロンプト上で起動してみる。

```text
C:>pgpdump.exe -v
pgpdump.exe version 0.29, Copyright (C) 1998-2014 Kazu Yamamoto
```

問題なし。

## 動作確認

実際にちゃんと動くかどうか [JPCERT/CC の公開鍵](https://www.jpcert.or.jp/jpcert-pgp.html "JPCERT コーディネーションセンター PGP公開鍵")をを使って確認してみる。

```text
C:>pgpdump.exe info-0x69ECE048.asc
Old: Public Key Packet(tag 6)(269 bytes)
        Ver 4 - new
        Public key creation time - Tue Jun 02 14:43:57 東京 (標準時) 2009
        Pub alg - RSA Encrypt or Sign(pub 1)
        RSA n(2048 bits) - ...
        RSA e(17 bits) - ...
Old: User ID Packet(tag 13)(29 bytes)
        User ID - JPCERT/CC <info@jpcert.or.jp>
Old: Signature Packet(tag 2)(316 bytes)
        Ver 4 - new
        Sig type - Generic certification of a User ID and Public Key packet(0x10).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA1(hash 2)
        Hashed Sub: preferred symmetric algorithms(sub 11)(3 bytes)
                Sym alg - AES with 256-bit key(sym 9)
                Sym alg - CAST5(sym 3)
                Sym alg - Triple-DES(sym 2)
        Hashed Sub: key server preferences(sub 23)(4 bytes)
                Flag - No-modify
        Hashed Sub: key flags(sub 27)(4 bytes)
                Flag - This key may be used to certify other keys
                Flag - This key may be used to sign data
                Flag - This key may be used to encrypt communications
                Flag - This key may be used to encrypt storage
                Flag - The private component of this key may be in the possession of more than one person
        Hashed Sub: preferred compression algorithms(sub 22)(2 bytes)
                Comp alg - ZLIB <RFC1950>(comp 2)
                Comp alg - ZIP <RFC1951>(comp 1)
        Hashed Sub: features(sub 30)(4 bytes)
                Flag - Modification detection (packets 18 and 19)
        Hashed Sub: preferred hash algorithms(sub 21)(3 bytes)
                Hash alg - SHA256(hash 8)
                Hash alg - SHA384(hash 9)
                Hash alg - SHA512(hash 10)
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Tue Jun 16 12:51:22 東京 (標準時) 2009
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0x317D97A469ECE048
        Hash left 2 bytes - cd 79
        RSA m^d mod n(2047 bits) - ...
                -> PKCS-1
Old: Signature Packet(tag 2)(277 bytes)
        Ver 3 - old
        Hash material(5 bytes):
                Sig type - Generic certification of a User ID and Public Key packet(0x10).
                Creation time - Tue Jun 02 14:43:57 東京 (標準時) 2009
        Key ID - 0xE7734FA60C7BDE12
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA1(hash 2)
        Hash left 2 bytes - e9 53
        RSA m^d mod n(2047 bits) - ...
                -> PKCS-1
Old: Signature Packet(tag 2)(156 bytes)
        Ver 4 - new
        Sig type - Generic certification of a User ID and Public Key packet(0x10).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA1(hash 2)
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Mon Jun 15 14:51:27 東京 (標準時) 2009
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0x8C756B2E2C94D4ED
        Hash left 2 bytes - 35 fd
        RSA m^d mod n(1022 bits) - ...
                -> PKCS-1
Old: Public Subkey Packet(tag 14)(269 bytes)
        Ver 4 - new
        Public key creation time - Tue Jun 02 14:43:57 東京 (標準時) 2009
        Pub alg - RSA Encrypt or Sign(pub 1)
        RSA n(2048 bits) - ...
        RSA e(17 bits) - ...
Old: Signature Packet(tag 2)(577 bytes)
        Ver 4 - new
        Sig type - Subkey Binding Signature(0x18).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA1(hash 2)
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Tue Jun 02 14:43:58 東京 (標準時) 2009
        Hashed Sub: key flags(sub 27)(4 bytes)
                Flag - This key may be used to encrypt communications
                Flag - This key may be used to encrypt storage
        Hashed Sub: embedded signature(sub 32)(284 bytes)
        Ver 4 - new
        Sig type - Primary Key Binding Signature(0x19).
        Pub alg - RSA Encrypt or Sign(pub 1)
        Hash alg - SHA256(hash 8)
        Hashed Sub: signature creation time(sub 2)(4 bytes)
                Time - Tue Jun 02 14:43:57 東京 (標準時) 2009
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0x09D704B753BA1622
        Hash left 2 bytes - 71 2d
        RSA m^d mod n(2048 bits) - ...
                -> PKCS-1
        Sub: issuer key ID(sub 16)(8 bytes)
                Key ID - 0x317D97A469ECE048
        Hash left 2 bytes - 1d e2
        RSA m^d mod n(2046 bits) - ...
                -> PKCS-1
```

うむ，問題ないようだな。
ちなみに Windows バイナリは[うちのサイトでも公開](http://www.baldanders.info/spiegel/archive/pgpdump/)しているのでご自由にどうぞ。

[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
[前回]: {{< relref "remark/2016/03/gcc-msys2-2.md" >}} "MSYS2 による gcc 開発環境の構築 ― gcc パッケージ群の導入"
[pgpdump]: https://github.com/kazu-yamamoto/pgpdump "kazu-yamamoto/pgpdump: A PGP packet visualizer"
