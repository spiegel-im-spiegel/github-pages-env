+++
title = "Ubuntu 22.04 LTS へのアップグレード"
date =  "2022-05-06T18:07:48+09:00"
description = "今回は特に大きな問題はなし。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "gnupg", "openssh", "firefox", "libreoffice", "password", "management", "tools", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] が「22.04 LTS にアップグレードせんのん？（← 超意訳）」と煩いので GW で余裕のある間にアップグレードすることにした。
今回は特に（少なくとも最近の機械では）日本語周りで不具合の話も聞かなかったし。

{{< fig-img src="./release-note.png" title="Ubuntu 22.04 リリースノート" link="./release-note.png" width="620" >}}

今回は特に大きな問題はなし。

1. 日本語入力周りは問題なく動いている。 [Ubuntu] コミュニティのご尽力に感謝
2. 21.10 から Firefox は Snap 版になっているのに何故か言語パッケージが APT で提供されているというチグハグな状態だったが，今回で統一された？
3. Thunderbird は相変わらず [GnuPG] の鍵束を認識してくれない。そろそろ捨てようか
4. LibreOffice は 7.3 系。 OpenPGP 鍵で暗号化したファイルもちゃんと開けるし，問題ないだろう。 Snap 版は相変わらず評判が悪いようだ（笑）
5. [KeePassXC] は既に jammy リポジトリがある。依存ライブラリのバージョンが違うらしい
6. [pgAdmin4] は jammy リポジトリがまだ用意されてない模様。こちらはしばらく様子見か
7. [Node.js] は [NodeSource](https://github.com/nodesource/distributions) で v18 系をインストールしてみた。今のところは問題なし。なにかあれば v16 (LTS) に戻すつもり

## いつものように GnuPG が古い

```text
$ gpg --version
gpg (GnuPG) 2.2.27
libgcrypt 1.9.4
Copyright (C) 2021 Free Software Foundation, Inc.
License GNU GPL-3.0-or-later <https://gnu.org/licenses/gpl.html>
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

まぁ v2.3 系でないのはしょうがないにしても，せめて今の LTS 最新版（[v2.2.34]({{< ref "/release/2022/02/gnupg-2_2_34-lts-is-released.md" >}} "GnuPG 2.2.34 (LTS) がリリースされた")）にしてくれよ `orz`

## [OpenSSH] と [OpenSSL]

```text
$ ssh -V
OpenSSH_8.9p1 Ubuntu-3, OpenSSL 3.0.2 15 Mar 2022
```

[OpenSSL] は v3.0.3 の[セキュリティ・アップデート](https://www.openssl.org/news/secadv/20220503.txt)に対応した[バックポート・パッチ](https://ubuntu.com/security/notices/USN-5402-1 "USN-5402-1: OpenSSL vulnerabilities | Ubuntu security notices | Ubuntu")が出ているようだ。
相変わらず分かりにくい。
普通にバージョンを上げてくれんもんかねぇ。

ちなみに [GnuPG] の `gpg-agent` との連携は問題なかった。

- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})
- [gpg-agent の設定： GnuPG in Ubuntu]({{< ref "/openpgp/gpg-agent-in-ubuntu.md" >}})

## [Docker] と apt-key

これに関しては記事を分けた。
[後半]({{< relref "./apt-key-is-deprecated.md" >}})へ続く（笑）

## ブックマーク

- [Ubuntu 22.04 その58 - Ubuntu 22.04 LTSがリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2022/04/ubuntu-2204-58-ubuntu-2204-lts.html)
- [Ubuntu 22.04 その59 - Ubuntu 22.04 LTSの新機能と変更点 - kledgeb](https://kledgeb.blogspot.com/2022/04/ubuntu-2204-59-ubuntu-2204-lts.html)
- [Ubuntu 22.04 その60 - Ubuntu 22.04 LTSの既知の問題 - kledgeb](https://kledgeb.blogspot.com/2022/04/ubuntu-2204-60-ubuntu-2204-lts.html)
- [Ubuntu 21.10 その89 - Ubuntu 22.04 LTSへのアップグレードは数日中に提供予定 - kledgeb](https://kledgeb.blogspot.com/2022/04/ubuntu-2110-89-ubuntu-2204-lts.html)
- [Ubuntu 20.04 その240 - Ubuntu 22.04 LTSへのアップグレード提供時期について - kledgeb](https://kledgeb.blogspot.com/2022/04/ubuntu-2004-240-ubuntu-2204-lts.html)
- [Canonical、「Ubuntu 22.04 LTS」の提供を開始 ～WSLパッケージもストアに登場 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1405022.html)
- [Ubuntu 22.04 をインストールしたら設定すること 10 ヶ条](https://zenn.dev/sprout2000/articles/8ea4a77d81583a)
- [Ubuntu 20.04にNode.jsをインストールする方法  | DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-20-04-ja)

- [Ubuntu 21.10 その90 - Ubuntu 22.04 LTSへアップグレードするには・アップグレードの注意事項 - kledgeb](https://kledgeb.blogspot.com/2022/04/ubuntu-2110-90-ubuntu-2204-lts.html)
- [Ubuntu 22.04 その63 - Ubuntu 22.04 LTSで行われたRaspberry Piのサポート改善 - kledgeb](https://kledgeb.blogspot.com/2022/04/ubuntu-2204-63-ubuntu-2204-ltsraspberry.html)
- [Ubuntu 22.04 その104 - 日付やタイムゾーンを自動設定するには・時間の12時間制と24時間制を切り替えるには - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-104-1224.html)
- [Ubuntu 22.04 その102 - ユーザーアカウントを追加するには・ユーザーアカウントを削除するには - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-102.html)
- [Ubuntu 22.04 その101 - ユーザーアカウントの各種設定を変更するには・ユーザーアカウントのパスワードや自動ログイン設定を変更するには - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-101.html)
- [Ubuntu 22.04 その100 - UIの言語を変更するには・言語環境に足りないソフトウェアがないかチェックするには - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-100-ui.html)
- [Ubuntu 22.04 その99 - 日本語入力（Mozc）の設定をカスタマイズして作業効率を上げよう - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-99-mozc.html)
- [Ubuntu 22.04 その95 - ディスプレイサーバーを切り替えるには・XorgとWayland - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-95-xorgwayland.html)
- [Ubuntu 22.04 その91 - 目への負担を軽減する夜間モードを設定するには - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-91.html)
- [Ubuntu 22.04 その90 - ディスプレイの解像度やHiDPIの設定をカスタマイズする・マルチディスプレイを設定する - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-90-hidpi.html)
- [Ubuntu 22.04 その89 - デバイスのバッテリー残量の確認やパフォーマンスを調整するには・その他電源に関する設定を変更するには - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-89.html)
- [Ubuntu 22.04 その67 - Active Directoryのサポート強化とADsys - kledgeb](https://kledgeb.blogspot.com/2022/04/ubuntu-2204-67-active-directoryadsys.html)
- [Ubuntu 22.04 LTSのインストール その7 - UEFIブートマネージャーの確認 〜 MOKの登録 - kledgeb](https://kledgeb.blogspot.com/2022/04/ubuntu-2204-lts-7-uefi-mok.html)
- [Ubuntu 22.04 LTS 日本語 Remix リリース | Ubuntu Japanese Team](ubuntu2204-ja-remix)
  - [Ubuntu 22.04 その106 - Ubuntu 22.04 LTS 日本語 Remixがリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-106-ubuntu-2204-lts-remix.html)
- [Ubuntu 22.04 その120 - UbuntuのブートローダーをBoot Repairで修復するには・Ubuntuが起動しないトラブルを解決 - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-120-ubuntuboot-repairubuntu.html)
- [Ubuntu 22.04 その125 - 強化されたスクリーンキャプチャー機能・スクリーンキャストも利用可能に - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-125.html)
- [Ubuntu 22.04 その133 - ファイルアプリとDropboxを連携させてDropboxを便利に活用しよう - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-133-dropboxdropbox.html)
- [Ubuntu 22.04 その136 - 画面からはみ出して移動できなくなったウィンドウを移動するには・サイズが小さすぎてサイズ調整できなくなったウィンドウのサイズを調整するには - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-136.html)
- [Ubuntu 22.04 その138 - インジケーターからGNOME Shell拡張を管理するには・Extension Listの使い方 - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-138-gnome-shellextension.html)
- [Ubuntu 22.04 その139 - インジケーターからマウントされたボリュームを開いたりアンマウントするには・Removable Drive Menuの使い方 - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-139-removable-drive-menu.html)
- [Ubuntu 22.04 その140 - インジケーターからサウンドデバイスやオーディオプロファイルを切り替えるには・Sound Input & Output Device Chooserの使い方 - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-140-sound-input-output.html)
- [Ubuntu 22.04 その141 - クリップボードの履歴を取りクリップボードの内容を再利用するには・Clipboard Indicatorの使い方 - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-141-clipboard-indicator.html)
- [Ubuntu 22.04 その144 - デスクトップアイコンをもっと細かくカスタマイズするには・ボリュームやアイコン表示のカスタマイズなど - kledgeb](https://kledgeb.blogspot.com/2022/05/ubuntu-2204-144.html)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[pgAdmin4]: https://www.pgadmin.org/ "pgAdmin - PostgreSQL Tools"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Node.js]: https://nodejs.org/
[OpenSSL]: https://www.openssl.org/
[OpenSSH]: https://www.openssh.com/
[Docker]: https://www.docker.com/ "Empowering App Development for Developers | Docker"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"
