+++
title = "GnuPG の HOME はどこにある？"
date = "2020-11-02T18:36:08+09:00"
description = "GnuPG の HOME ディレクトリって .gnupg フォルダじゃねーの？"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "gnupg", "windows" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近また Windows パソコンを使うことになって，色々と[環境を整える]({{< ref "/remark/2020/10/windows-terminal-and-nyagos-and-scoop.md" >}} "Windows Terminal × NYAGOS × Scoop ＝ ♥")ついでに [GnuPG] もインストールしたのだが...

今回 [Scoop] というパッケージ・マネージャを導入したのだが，これを使って [GnuPG] をインストールするとちょっと面白いことになる。

[Scoop] の導入方法は割愛して，いきなり [GnuPG] のインストールから話を始める。

```text
$ scoop install gnupg
Installing 'gnupg' (2.2.23) [64bit]

...

Linking ~\scoop\apps\gnupg\current => ~\scoop\apps\gnupg\2.2.23
Persisting home
'gnupg' (2.2.23) was installed successfully!
```

これで [GnuPG] v2.2.23 は `%USERPROFILE%\scoop\apps\gnupg\2.2.23` フォルダにインストールされ `%USERPROFILE%\scoop\apps\gnupg\current` にシンボリック・リンクが張られる。
さらに環境変数 `PATH` に `%USERPROFILE%\scoop\apps\gnupg\current` が追加される。

[Scoop] の特徴でありメリットのひとつはインストールした実行バイナリを `%USERPROFILE%\scoop\shims` フォルダに集めることで環境変数 `PATH` を無駄に汚さないことにあるが， [GnuPG] の場合はシングル・バイナリではなく，インストールした複数のバイナリと協調して動作するため，単純に `shims` フォルダにコピるのは難しかったようだ。

[Scoop] によってインストールされた [GnuPG] のディレクトリ構成は以下のようになっている[^env1]。

[^env1]: ちなみに私は [Windows Terminal ＋ NYAGOS を愛用]({{< ref "/remark/2020/10/windows-terminal-and-nyagos-and-scoop.md" >}} "Windows Terminal × NYAGOS × Scoop ＝ ♥")している。例に出ている `ll` は [NYAGOS] の組み込みコマンド（の alias）である。

```text
$ pwd
C:\Users\username

$ cd scoop\apps\gnupg\current

$ ll
-rw-a--  33K Sep  4 00:39:00 README.txt
drwx---    0 Oct 27 13:59:32 bin/
dr-x---    0 Oct 27 13:59:32 home@ -> C:\Users\username\scoop\persist\gnupg\home
drwx---    0 Oct 27 13:59:32 include/
-rw-a--   58 Oct 27 13:59:32 install.json
drwx---    0 Oct 27 13:59:32 lib/
-rw-a-- 1.4K Oct 27 13:36:14 manifest.json
drwx---    0 Oct 27 13:59:32 share/
```

実際に [GnuPG] を動かしてみると

```text
$ gpg --version
gpg (GnuPG) 2.2.23
libgcrypt 1.8.6
Copyright (C) 2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: C:/Users/username/scoop/apps/gnupg/current/home
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
    CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

となる。
HOME ディレクトリが `%USERPROFILE%\scoop\apps\gnupg\current\home` になっているのがお分かりだろうか。
ただし，この `home` ディレクトリはシンボリック・リンクなので，実体は `%USERPROFILE%\scoop\persist\gnupg\home` ということになる。

って [GnuPG] の HOME ディレクトリって `%USERPROFILE%` 直下の `.gnupg` フォルダじゃねーの？

おやや，と思って調べてみたら Windows 版の [GnuPG] はちょっと違うらしい。

{{< fig-quote type="markdown" title="Agent Options (Using the GNU Privacy Guard)" link="https://www.gnupg.org/documentation/manuals/gnupg/Agent-Options.html" lang="en" >}}
{{% quote %}}To install GnuPG as a portable application under Windows, create an empty file named `gpgconf.ctl` in the same directory as the tool `gpgconf.exe`. The root of the installation is then that directory; or, if `gpgconf.exe` has been installed directly below a directory named `bin`, its parent directory{{% /quote %}}.
{{< /fig-quote >}}

つまり（`gpgconf.exe` のある） `%USERPROFILE%\scoop\apps\gnupg\current\bin` フォルダに空の `gpgconf.ctl` ファイルを置けば その親ディレクトリ（`%USERPROFILE%\scoop\apps\gnupg\current`）がインストール・ルートとして認識され，その直下の `home` フォルダが [GnuPG] の `HOME` ディレクトリとなる。

Windows 版の場合  `gpgconf.ctl` ファイルによる HOME の指定が優先され，環境変数 `GNUPGHOME` やレジストリを使った指定は無視されるらしい。
ただし `--homedir` オプションは効くとのこと。
これはこれで合理的。

試しに [Ubuntu] で `gpgconf.ctl` を試してみたが効かなかった。
まぁ，当たり前か（笑）

それにしても [Scoop] はよい。
端末エミュレータでコマンドライン中心に作業している人には [Scoop] のほうが設計がシンプルで使いやすいんじゃないだろうか。
まぁ，私は作業の中心が（仕事以外では）もはや Windows ではなくなっているので，余計にそう思うのかもしれないが（笑）

## ブックマーク

- [scoop / nyagos で始めるコマンドライン生活](https://zenn.dev/zetamatta/books/5ac80a9ddb35fef9a146)

[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Scoop]: https://scoop.sh/ "Scoop"
[Chocolatey]: https://chocolatey.org/ "Chocolatey Software | Chocolatey - The package manager for Windows"
[NYAGOS]: https://github.com/zetamatta/nyagos "zetamatta/nyagos: NYAGOS - The hybrid Commandline Shell betweeeeeeen UNIX & DOS"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書
