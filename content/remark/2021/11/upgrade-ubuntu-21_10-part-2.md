+++
title = "Ubuntu 21.10 へのアップグレード（本番）"
date =  "2021-11-06T20:49:34+09:00"
description = "日本語入力は問題なし / mount.cifs で手こずる / GnuPG が古すぎる / OpenSSH と OpenSSL はこのままでいいか"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "nas", "samba", "gnupg", "firefox", "thunderbird", "libreoffice", "password", "management", "tools", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]の続き。

[Ubuntu 21.10 日本語 Remix が登場](https://www.ubuntulinux.jp/News/ubuntu2110-ja-remix "Ubuntu 21.10 日本語 Remix リリース | Ubuntu Japanese Team")したので，自宅のサブマシンに真っさらからインストールしてみたが，[前回]アレだった日本語周りも特に問題なく動いてる感じなので，いよいよメインマシンをアップグレードすることにした。

アップグレード作業自体は滞りなく完了。

1. 日本語入力周りは問題なく動いている。 [Ubuntu] コミュニティのご尽力に感謝
2. [前回]書いたとおり Firefox は Snap 版になっていたが，こちらも問題なし。日本語フォント周りも壊れることなくちゃんと設定された
3. Thunderbird が 91 系にアップグレードされている。今のところは特に問題なし
4. LibreOffice が 7.2 系にアップグレードされている。 OpenPGP で暗号化したファイルもちゃんと開けるし，問題ないだろう
5. [KeePassXC] は 21.10 リリース以降アップデートされてないため APT リポジトリを impish に切り替えれない（`apt update` でワーニングが出る）しばらくは様子見か
6. 同じく [pgAdmin4] も 21.10 リリース以降アップデートされてないため APT リポジトリを impish に切り替えれない。こちらもしばらくは様子見

## mount.cifs で手こずる

アップグレードして再起動したら NAS に接続できなくなったのには焦った。

もしかしたら [Ubuntu] 21.04 以降で [Active Directory に対応](https://kledgeb.blogspot.com/2021/04/ubuntu-2104-21-ubuntu-2104microsoft.html "Ubuntu 21.04 その21 - Ubuntu 21.04登場・Microsoft Active Directoryとの連携やMicrosoft SQL Serverのサポート改善など - kledgeb")したことが影響してるのかも知れないが， SMB でマウントする際の認証情報で `domain` または `workgroup` オプションでドメイン名またはワークグループ名を明示しないとダメぽいみたい。
でも，なんで前のバージョンで上手くいってたのか...

まぁ，でも，今は `/etc/fstab` ファイルで設定するのは時代遅れみたいなので，見直しが必要かなぁ。

- [[root tip] systemd mount unit samples - Tutorials - Manjaro Linux Forum](https://forum.manjaro.org/t/root-tip-systemd-mount-unit-samples/1191)

## GnuPG が古すぎる

[GnuPG] は

```text
$ gpg --version
gpg (GnuPG) 2.2.20
libgcrypt 1.8.8
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

と，相変わらず古すぎる。
2.3 系にしないのは仕方ないとしても，せめて [2.2.32]({{< ref "/release/2021/10/gnupg-2_2_32-lts-is-released.md" >}} "GnuPG 2.2.32 (LTS) がリリースされた") 相当には上げておかないと拙いと思うのだが，その気配はないよね。

## OpenSSH と OpenSSL はこのままでいいか

OpenSSH と OpenSSL は

```text
$ ssh -V
OpenSSH_8.4p1 Ubuntu-6ubuntu2, OpenSSL 1.1.1l  24 Aug 2021
```

となっている。
バックポート・パッチが当たっているはずなので，あまり気にしてない。
次の LTS バージョンである [Ubuntu] 22.04 では [OpenSSL 3.0 に上げる](https://kledgeb.blogspot.com/2021/10/ubuntu-2204-3-openssl-30.html "Ubuntu 22.04 その3 - OpenSSL 3.0への移行計画 - kledgeb")らしいので，おとなしく待つことにしよう。

やっぱ問題は [GnuPG] だよな。
まぁ，サブマシンを真っさらの [Ubuntu] 機にできたので，これを使って [GnuPG] 2.3 系の実験を始めるのもいいかも知れない。

## ブックマーク

- [CIFS 経由で NAS に接続する]({{< ref "/remark/2019/03/common-internet-file-system.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[前回]: {{< ref "/remark/2021/10/upgrade-ubuntu-21_10.md" >}} "Ubuntu 21.10 へのアップグレードを試してみた"
[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[pgAdmin4]: https://www.pgadmin.org/ "pgAdmin - PostgreSQL Tools"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
<!-- eof -->
