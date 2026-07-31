+++
title = "Ubuntu 19.10 にアップグレードする"
date =  "2019-10-20T18:49:35+09:00"
description = "特にトラブルはなし。でも，やっぱり Ubuntu はセキュリティ・ツール周りの管理が弱いよなぁ。"
image = "/images/attention/tools.png"
tags  = [ "ubuntu", "tools", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

短期サポート版である [Ubuntu] 19.10 がリリースされた。
2020年7月までサポートが有効である。

- [Ubuntu 19.10 (Eoan Ermine) released](https://lists.ubuntu.com/archives/ubuntu-announce/2019-October/000250.html)
- [EoanErmine/ReleaseNotes - Ubuntu Wiki](https://wiki.ubuntu.com/EoanErmine/ReleaseNotes)
- [Ubuntu 19.10 その27 - Ubuntu 19.10がリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2019/10/ubuntu-1910-27-ubuntu-1910.html)
- [Ubuntu 19.10 その28 - Ubuntu 19.10の新機能と変更点・既知の問題 - kledgeb](https://kledgeb.blogspot.com/2019/10/ubuntu-1910-28-ubuntu-1910.html)
- [Ubuntu 19.04 その149 - Ubuntu 19.10へアップグレードするには・アップグレードの注意事項 - kledgeb](https://kledgeb.blogspot.com/2019/10/ubuntu-1904-149-ubuntu-1910.html)

旧バージョンから 19.10 へのアップグレードは[リリースノート](https://wiki.ubuntu.com/EoanErmine/ReleaseNotes "EoanErmine/ReleaseNotes - Ubuntu Wiki")を参考にするとよい。
手元の環境では特に問題なくアップグレードできた。

アップグレード時にサードパーティの APT リポジトリ（[git]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}} "PPA から Git をインストールする") や [ATOM](https://flight-manual.atom.io/getting-started/sections/installing-atom/ "Installing Atom") など）が外れるので，必要ならアップグレード時に戻して更新すること。

短期サポート版で半年ごとにアップグレードするのは若干面倒だが，ディストリビューションのアップグレードのタイミングでないと更新されないアプリケーションもあるようなので（完全に自前で管理できるなら別だが）きちんとアップグレードに追従しておくほうが長い目で見てお得である。

## 個別のアプリケーションについて

### [GnuPG] および [Libgcrypt]

[GnuPG] および [Libgcrypt] はアップデートされなかった。
はるか古いバージョンのままである。

```text
$ gpg --version
gpg (GnuPG) 2.2.12
libgcrypt 1.8.4
Copyright (C) 2018 Free Software Foundation, Inc.
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

どうも [Ubuntu] は [GnuPG] をまともにメンテできないようだ。
[電子署名 spam の件]({{< ref "/release/2019/07/gnupg-2_2_17-is-released.md" >}} "GnuPG 2.2.17 リリース： 公開鍵サーバ・アクセスに関する過激な変更あり")もあり，特に [Libgcrypt] は[セキュリティ・アップデート]({{< ref "/release/2019/08/libgcrypt-1_8_5-is-released.md" >}} "Libgcrypt 1.8.5 がリリース【セキュリティ・アップデート】")があったばかりなんだけど。
こりゃあ，本格的に何か考えないとな。

- [GnuPG 2.2.17 リリース： 公開鍵サーバ・アクセスに関する過激な変更あり]({{< ref "/release/2019/07/gnupg-2_2_17-is-released.md" >}})
- [Libgcrypt 1.8.5 がリリース【セキュリティ・アップデート】]({{< ref "/release/2019/08/libgcrypt-1_8_5-is-released.md" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Libgcrypt]: https://gnupg.org/software/libgcrypt/

### [OpenSSH] および [OpenSSL]

アップグレード直後の [OpenSSH] および [OpenSSL] のバージョンは以下のとおりだった。

```text
$ ssh -V
OpenSSH_8.0p1 Ubuntu-6build1, OpenSSL 1.1.1c  28 May 2019
```

[OpenSSH] も 8.1 で[セキュリティ・アップデート](https://www.openssh.com/txt/release-8.1)があったばかりだし [OpenSSL] 1.1.1 も[セキュリティ・アップデート](https://www.openssl.org/news/secadv/20190910.txt "OpenSSL Security Advisory [10 September 2019]")があったばかりだが，ちゃんと対応しているのかね。

[OpenSSH]: https://www.openssh.com/
[OpenSSL]: https://www.openssl.org/

### [Thunderbird]

[Thunderbird] は 68 ベースになっていた。
これは嬉しい。

[Thunderbird] に関しては [Enigmail が2020年夏に本体に組み込まれる](https://blog.mozilla.org/thunderbird/2019/10/thunderbird-enigmail-and-openpgp/)という話もあるので，きちんと追従していただきたいところである。

[Thunderbird]: https://www.thunderbird.net/ "Thunderbird — Software made to make email easier. — Mozilla"

### [Lollypop]

（2019-10-24 更新）

音楽プレイヤー [Lollypop の PPA](https://launchpad.net/~gnumdk/+archive/ubuntu/lollypop "Lollypop : Cédric Bellegarde") が 19.10 に対応した。
詳しくは以下の記事をどうぞ。

- [音楽プレイヤー Lollypop を試す]({{< ref "/remark/2019/06/lollypop-music-player.md" >}})

[Lollypop]: https://wiki.gnome.org/Apps/Lollypop?action=show "Apps/Lollypop - GNOME Wiki!"

### [LibreOffice]

[LibreOffice] は 6.3 ベースになった。
もっとも私は[既に APT での管理を止めている]({{< ref "/remark/2019/05/installing-libreoffice-in-ubuntu.md" >}} "Ubuntu に LibreOffice をインストールする3つの方法")のでどうでもいいけど。

[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"

## やっぱり...

[Ubuntu] はセキュリティ・ツール周りの管理が弱いよなぁ。
脆弱性を放置するとか「保守的」では済まないと思うのだが...

## ブックマーク

- [Ubuntu 19.10 その26 - インストーラーからZFS on rootの構築が可能に - kledgeb](https://kledgeb.blogspot.com/2019/10/ubuntu-1910-26-zfs-on-root.html)

- [Ubuntu アプリケーションにおけるセキュリティ・アップデート一覧]({{< ref "/release/vuln-list.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"

### 参考図書

{{% review-paapi "B07VJKJY7M" %}} <!-- 私はどのようにしてLinuxカーネルを学んだか -->
