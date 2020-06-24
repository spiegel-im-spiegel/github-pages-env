+++
title = "Ubuntu に最新の OpenSSL を入れる"
date =  "2020-06-24T20:37:29+09:00"
description = "セキュリティの中核部品なので手動で入れるのはどうかと思ったが，調べたところ簡単に行きそうだ。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "install", "openssl", "openssh" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近ようやく [Ubuntu] 20.04 で [OpenSSH] 8.3 に対応するバックポート・パッチが出たようだ？

```text
$ ssh -V
OpenSSH_8.2p1 Ubuntu-4ubuntu0.1, OpenSSL 1.1.1f  31 Mar 2020
```

ただし，上に見るように [OpenSSL] に関しては 1.1.1f のまま，[バックポート・パッチもなく](https://usn.ubuntu.com/ "Ubuntu security notices")，脆弱性が放置されている模様。

- [OpenSSL の脆弱性対策について(CVE-2020-1967) ：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/alert20200423.html)
- [OpenSSL の脆弱性 (CVE-2020-1967) に関する注意喚起](https://www.jpcert.or.jp/at/2020/at200018.html)

私としてはそろそろ我慢の限界なので，手動で最新バージョンの [OpenSSL] を入れることにした。
セキュリティの中核部品なので手動で入れるのはどうかと思ったが，調べたところ簡単に行きそうだ。

自分で試す場合は自己責任でね☆

## 既定状態

まず [OpenSSL] を手動インストールする前の状態をメモっておく。

```text
$ which openssl
/usr/bin/openssl

$ ldd /usr/bin/openssl
	linux-vdso.so.1 (0x00007ffe186db000)
	libssl.so.1.1 => /usr/lib/x86_64-linux-gnu/libssl.so.1.1 (0x00007f5d77c11000)
	libcrypto.so.1.1 => /usr/lib/x86_64-linux-gnu/libcrypto.so.1.1 (0x00007f5d7793b000)
	libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f5d77918000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f5d77726000)
	libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007f5d77720000)
	/lib64/ld-linux-x86-64.so.2 (0x00007f5d77d7e000)

$ openssl version
OpenSSL 1.1.1f  31 Mar 2020
```

この状態を覚えておく。

## [OpenSSL] 最新版のダウンロード

2020-06-24 時点の最新バージョンは 1.1.1g なので，これを[ダウンロードする](https://www.openssl.org/source/)ところから始める。

```text
$ cd /usr/local/src/
$ sudo wget https://www.openssl.org/source/openssl-1.1.1g.tar.gz
$ sudo wget https://www.openssl.org/source/openssl-1.1.1g.tar.gz.asc
$ gpg -d openssl-1.1.1g.tar.gz.asc
gpg: 署名されたデータが'openssl-1.1.1g.tar.gz'にあると想定します
gpg: 2020年04月21日 21時22分45秒 JSTに施された署名
gpg:                RSA鍵8657ABB260F056B1E5190839D9C4D26D0E604491を使用
gpg: "Matt Caswell <matt@openssl.org>"からの正しい署名 [不明の]
gpg:                 別名"Matt Caswell <frodo@baggins.org>" [不明の]
gpg: *警告*: この鍵は信用できる署名で証明されていません!
gpg:          この署名が所有者のものかどうかの検証手段がありません。
主鍵フィンガープリント: 8657 ABB2 60F0 56B1 E519  0839 D9C4 D26D 0E60 4491
```

`openssl-1.1.1g.tar.gz.asc` ファイルは `openssl-1.1.1g.tar.gz` ファイルに対する [OpenPGP] 電子署名ファイルで，上述の [GnuPG] コマンドで検証できる[^gpg1]。

[^gpg1]: インポートした公開鍵を有効化してないと署名検証時に「この鍵は信用できる署名で証明されていません」と表示される。まぁ，有効化しなくても検証自体はできるけど。 [OpenPGP] 鍵の操作については拙文「[GnuPG チートシート（鍵作成から失効まで）]({{< ref "/openpgp/gnupg-cheat-sheet.md" >}})」を参考にどうぞ。

ダウンロードしたファイルの検証は必ず行うこと。
検証用の公開鍵がない場合には

```text
$ gpg --recv-keys 0xD9C4D26D0E604491
```

で鍵サーバからインポートできる。

## [OpenSSL] 最新版をビルド&インストールする

では，インストールの続きを。

```text
$ pwd
/usr/local/src

$ sudo tar xvf openssl-1.1.1g.tar.gz 
$ cd openssl-1.1.1g/
$ sudo ./config
$ sudo make
$ sudo make install
```

`make` コマンドが始まると長いため，あらかじめお茶菓子を用意しておきましょう（笑）

`make install` により各ファイル（マニュアル以外）は以下のディレクトリに配置される。

- `/usr/local/bin/`
    - `openssl`
- `/usr/local/lib/`
    - `libcrypto.so.1.1`
    - `libssl.so.1.1`
- `/usr/local/ssl/`
    - `openssl.cnf` 等[^cnf1]

[^cnf1]: 既定では `openssl.cnf` ファイルは `/etc/ssl/` ディレクトリに置かれているものを読む。環境変数 `OPENSSL_CONF` で変更することも可能。

## ldconfig を忘れずに

ほんじゃあ動かしてみよっか。

```text
$ which openssl
/usr/local/bin/openssl

$ ldd /usr/local/bin/openssl
	linux-vdso.so.1 (0x00007ffcb01b7000)
	libssl.so.1.1 => /usr/lib/x86_64-linux-gnu/libssl.so.1.1 (0x00007f1d88d55000)
	libcrypto.so.1.1 => /usr/lib/x86_64-linux-gnu/libcrypto.so.1.1 (0x00007f1d88a7f000)
	libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f1d88a5c000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f1d8886a000)
	libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007f1d88864000)
	/lib64/ld-linux-x86-64.so.2 (0x00007f1d88ec6000)

$ openssl version
openssl: symbol lookup error: openssl: undefined symbol: EVP_mdc2, version OPENSSL_1_1_0
```

おっと。
`libssl.so.1.1` と `libcrypto.so.1.1` のリンク情報が古いままだったね。
ちうわけで

```text {hl_lines=["5-6"]}
$ sudo ldconfig

$ ldd /usr/local/bin/openssl
	linux-vdso.so.1 (0x00007fff9e6f0000)
	libssl.so.1.1 => /usr/local/lib/libssl.so.1.1 (0x00007fcbf44ab000)
	libcrypto.so.1.1 => /usr/local/lib/libcrypto.so.1.1 (0x00007fcbf41c1000)
	libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007fcbf419e000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fcbf3fac000)
	libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007fcbf3fa6000)
	/lib64/ld-linux-x86-64.so.2 (0x00007fcbf4620000)

$ openssl version
OpenSSL 1.1.1g  21 Apr 2020
```

`ldconfig` コマンドでキャッシュを更新するのを忘れずに[^lc1]（笑）
これで [OpenSSH] も

[^lc1]: 万が一 `ldconfig` で新しいファイルに切り替わらない場合は `/etc/ld.so.conf.d/` ディレクトリにある `*.conf` ファイルをチェックして `/usr/local/lib/` ディレクトリもキャッシュの対象に含めるよう設定する。

```text
$ ssh -V
OpenSSH_8.2p1 Ubuntu-4ubuntu0.1, OpenSSL 1.1.1g  21 Apr 2020
```

と新しい方の [OpenSSL] を見てくれる。

ちなみにオリジナルの `/usr/bin/openssl` はそのまま残ってて

```text
$ /usr/bin/openssl version
OpenSSL 1.1.1f  31 Mar 2020 (Library: OpenSSL 1.1.1g  21 Apr 2020)
```

てな感じになる。
APT で [OpenSSL] を更新する場合は注意すること。

## ブックマーク

- [UbuntuでOpenSSLを最新バージョンにする | Trusted Design](http://www.trusted-design.net/archives/933/)
- [OpenSSLが error while loading shared libraries libssl.so.1.1 - セキュリティ](https://kaworu.jpn.org/security/OpenSSL%E3%81%8C_error_while_loading_shared_libraries_libssl.so.1.1)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[OpenSSH]: https://www.openssh.com/
[OpenSSL]: https://www.openssl.org/
[OpenPGP]: https://www.openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
<!-- eof -->
