+++
title = "Ubuntu 20.10 のアップグレードに失敗していた"
date =  "2020-11-23T11:02:44+09:00"
description = "ま，まぁ，ともかく，アップグレードは完了した。"
image = "/images/attention/kitten.jpg"
tags  = [ "ubuntu", "tools", "install", "openssl", "openssh", "gnupg" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2020年10月に [Ubuntu] 20.10 がリリースされたが，今回は日本語 Remix 版が出るまで呑気に待つつもりだった。
その後ちょっと忙しくなったので完全に忘れていたのだが

- [[ubuntu-jp:6373] Ubuntu 20.10 日本語 Remix リリース](https://lists.ubuntu.com/archives/ubuntu-jp/2020-November/006372.html)
- [Ubuntu 20.10 その50 - Ubuntu 20.10 日本語 Remixがリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2020/11/ubuntu-2010-50-ubuntu-2010-remix.html)

2週間前に出てるぢゃん `orz`

あれ？ そういえば GUI の update が何も言ってこないな？ と思って少し調べてみた。
まず GUI 側の設定は問題ない（新バージョンの通知が「すべての新バージョン」になっていればOK）。

{{< fig-img src="./settings-update.png" link="./settings-update.png" width="1009" >}}

念のため `/etc/update-manager/release-upgrades` ファイルも見てみるが

```toml
[DEFAULT]
# Default prompting and upgrade behavior, valid options:
#
#  never  - Never check for, or allow upgrading to, a new release.
#  normal - Check to see if a new release is available.  If more than one new
#           release is found, the release upgrader will attempt to upgrade to
#           the supported release that immediately succeeds the
#           currently-running release.
#  lts    - Check to see if a new LTS release is available.  The upgrader
#           will attempt to upgrade to the first LTS release available after
#           the currently-running one.  Note that if this option is used and
#           the currently-running release is not itself an LTS release the
#           upgrader will assume prompt was meant to be normal.
Prompt=normal
```

ってなってた。
これも問題ない。

しょうがないので手動でターミナル・エミュレータから `do-release-upgrade` を起動してみる。
したら

```text
$ sudo do-release-upgrade
Checking for a new Ubuntu release
Failed to connect to https://changelogs.ubuntu.com/meta-release. Check your Internet connection or proxy settings
No new release found.
```

は？ ネット接続設定に問題はねーよ。
どーりでアップグレードが動かないわけだよ。
ちなみにブラウザから上述の URL を叩いてみたがちゃんと開くし，末尾に Groovy Gorilla の項目もある。

何じゃこりゃ，とググってみたら `/etc/update-manager/meta-release` ファイル

```toml
[METARELEASE]
URI = https://changelogs.ubuntu.com/meta-release
URI_LTS = https://changelogs.ubuntu.com/meta-release-lts
URI_UNSTABLE_POSTFIX = -development
URI_PROPOSED_POSTFIX = -proposed
```

の `URI` 項目を

```toml
URI = http://changelogs.ubuntu.com/meta-release
```

に書き換えろ，とあった。
つまり HTTPS 接続できないらしい。
実際に書き換えたらようやくアップグレードが走り出した。

これ，大丈夫か？

ま，まぁ，ともかく，途中で `mozc` や GNOME の日本語パッケージが勝手に消されるアクシデントもあったが設定の「地域と言語」から復元できたし，アップグレードは完了した。

## 個別のアプリケーションについて

### [GnuPG] および [Libgcrypt]

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Libgcrypt]: https://gnupg.org/software/libgcrypt/

```text
$ gpg --version
gpg (GnuPG) 2.2.20
libgcrypt 1.8.5
Copyright (C) 2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: /home/username/.gnupg
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256,
      TWOFISH, CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

[GnuPG] の遅れっぷりは相変わらず。
ちなみに現時点（2020-11-23）の最新版は

```text {hl_lines=["2-3"]}
$ gpg --version
gpg (GnuPG) 2.2.24
libgcrypt 1.8.7
Copyright (C) 2020 g10 Code GmbH
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: ********
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
    CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

である。
おっ， 2.2.24 から著作権者表示が代わってるな。
いや，開発やメンテは一貫して “g10 Code GmbH” がやってるんだけど。

### [OpenSSH] および [OpenSSL]

[OpenSSH]: https://www.openssh.com/
[OpenSSL]: https://www.openssl.org/

```text
$ ssh -V
OpenSSH_8.3p1 Ubuntu-1, OpenSSL 1.1.1h  22 Sep 2020
```

[OpenSSH] は 2020-09-27 に 8.4 が出てるのでこれも周回遅れ。

[OpenSSL] については APT では

```text
$ apt show openssl
Package: openssl
Version: 1.1.1f-1ubuntu4
...
```

とバックポート・パッチのせいで相変わらずよく分からないのだが，私が

- [Ubuntu に最新の OpenSSL を入れる]({{< ref "/remark/2020/06/installing-openssl-in-ubuntu.md" >}})

という感じに手動でアップデートを行っている。
ちなみに，現時点（2020-11-23）での [OpenSSL] 1.1.1 系最新版は 2020-09-22 にリリースされた 1.1.1h である。

### Libsecret のアップデート

[git]: https://git-scm.com/

[Git][git] credential helper である GNOME/libsecret もバージョンが上がってる？ のかな。

```text
$ apt show libsecret-1-dev
Package: libsecret-1-dev
Version: 0.20.3-1
...
```

というわけで，これも手動でビルドする。

```text
$ mkdor ~/work
$ cd ~/work
$ cp -r /usr/share/doc/git/contrib/credential/libsecret .
$ cd libsecret
$ make
gcc -g -O2 -Wall  -pthread -I/usr/include/libsecret-1 -I/usr/include/libmount -I/usr/include/blkid -I/usr/include/glib-2.0 -I/usr/lib/x86_64-linux-gnu/glib-2.0/include -o git-credential-libsecret.o -c git-credential-libsecret.c
gcc -o git-credential-libsecret  git-credential-libsecret.o -lsecret-1 -lgio-2.0 -lgobject-2.0 -lglib-2.0
```

これで生成した `git-credential-libsecret` を `$PATH` の通ったディレクトリに放り込んでおけばよい。
[Git][git] の設定に組み込むには以下のコマンドでOK。

```text
$ git config --global credential.helper libsecret
```

## ブックマーク

- [Ubuntu Fridge | Ubuntu 20.10 (Groovy Gorilla) released](http://ubuntu-news.org/2020/10/22/ubuntu-20-10-groovy-gorilla-released/)
- [Groovy Gorilla Release Notes - Release - Ubuntu Community Hub](https://discourse.ubuntu.com/t/groovy-gorilla-release-notes/15533) ([日本語](https://wiki.ubuntu.com/GroovyGorilla/ReleaseNotes/Ja))
- [OpenSSL Security Advisory [09 September 2020]](https://www.openssl.org/news/secadv/20200909.txt)
- [[SOLVED] Cannot upgrade 20.10 Failed to connect to https://changelogs.ubuntu.com/meta-release](https://ubuntuforums.org/showthread.php?t=2452515&s=b5cbbeb5143c46914981366cb0eec2ef)
- [Ubuntu 20.10 その12 - Ubuntu 20.10がリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2020/10/ubuntu-2010-12-ubuntu-2010.html)
- [Ubuntu20.04から20.10にアップグレードする方法 | ネットでお金を稼ぐ男のブログ](https://net.youhei02.com/ubuntu13/)

- [Ubuntu インストール用のブート可能 USB メモリを作成する]({{< ref "/remark/2019/04/bootable-usb.md" >}})
- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"

## 参考

{{% review-paapi "B01NBU1OS5" %}} <!-- シリコンパワー USBメモリ 32GB USB3.1 -->
