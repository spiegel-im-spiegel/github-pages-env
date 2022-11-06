+++
title = "Ubuntu 22.10 へのアップグレード"
date =  "2022-10-30T16:08:08+09:00"
description = "サードパーティ・パッケージの管理に注意"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] 22.10 がリリースされた。
約半年の短期サポート。

- [Kinetic Kudu Release Notes - Release - Ubuntu Community Hub](https://discourse.ubuntu.com/t/kinetic-kudu-release-notes/27976)

LTS 版ではないので，長期運用を望む場合は [22.04 LTS]({{< ref "/remark/2022/05/upgrade-ubuntu-22_04-lts.md" >}} "Ubuntu 22.04 LTS へのアップグレード") のままでOK。

22.04 LTS を入れていると「Ubuntuの新バージョンの通知」が「長期サポート(LTS)版」になっているので，これを「すべての新バージョン」に変更する。

{{< fig-img src="./update-settings-1.png" title="ソフトウェアとアップデート" link="./update-settings-1.png" width="1041" >}}

これで「ソフトウェアの更新」を開くと（最新パッケージへの更新の後） `[アップグレード...]` ボタンが表示される。

{{< fig-img src="./check-update.png" title="ソフトウェアの更新" link="./check-update.png" width="645" >}}

このボタンを押下するとアップグレードが始まる[^cmd1]。

[^cmd1]: コマンドラインでアップグレードする場合は `sudo apt dist-upgrade` でアップグレードが開始される。

{{< fig-img src="./upgrade-1.png" title="アップグレードの開始" link="./upgrade-1.png" width="620" >}}

今回も特に大きな問題はなかった。

## いつものように [GnuPG] が古い

```text
$ gpg --version
gpg (GnuPG) 2.2.35
libgcrypt 1.10.1
Copyright (C) 2022 g10 Code GmbH
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

ちなみに 2022-10-30 時点で [GnuPG] v2.2 系の最新は 2.2.40 である。
相変わらずだなぁ。
まぁ Libgcrypt が 1.10 系になっているだけマシと思うことにしよう。

はやく 2.3 系にアップグレードしてくれよ。

## [OpenSSH] と [OpenSSL]

```text
$ ssh -V
OpenSSH_9.0p1 Ubuntu-1ubuntu7, OpenSSL 3.0.5 5 Jul 2022
```

んー。
ちょっと古い気がするが，バックポートパッチがあたっていると信じよう（笑）

ところで [OpenSSL] は v3.0.6 がリリースされたあと一度引っ込めている。
どうも 2022-11-01 UTC に改めてセキュリティ・アップデートを出すようだが，今回の脆弱性はちょっとヤバいかも？ みたいな話もある。
ご注意を。

- [OpenSSL Security Advisory [11 October 2022]](https://www.openssl.org/news/secadv/20221011.txt)
- [OpenSSL warns of critical security vulnerability with upcoming patch | ZDNET](https://www.zdnet.com/article/openssl-warns-of-critical-security-vulnerability-with-upcoming-patch/)

{{< div-box type="markdown" >}}
**【2022-11-06 追記】**

- [OpenSSL の脆弱性対策について(CVE-2022-3602、CVE-2022-3786)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/alert20221102.html)
- [OpenSSLの脆弱性（CVE-2022-3602、CVE-2022-3786）に関する注意喚起](https://www.jpcert.or.jp/at/2022/at220030.html)
- [脆弱性を修正した「OpenSSL 3.0.7」が予告通り公開 ～ただし、深刻度評価は引き下げ - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1452569.html)
{{< /div-box >}}

ちなみに [GnuPG] の `gpg-agent` との連携は問題なかった。
よーし，うむうむ，よーし。

- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})
- [gpg-agent の設定： GnuPG in Ubuntu]({{< ref "/openpgp/gpg-agent-in-ubuntu.md" >}})
- [Ubuntu で OpenSSH の鍵管理を gpg-agent に委譲する【たぶん決定版】](https://zenn.dev/spiegel/articles/20210109-gpg-agent)

## APT でインストールされているサードパーティ・パッケージの管理に注意

今回のアップグレードで何がありがたいって `/etc/apt/sources.list.d/` ディレクトリで管理されるサードパーティ・パッケージでも可能な限り `kinetic` リポジトリに対応してくれるみたい。
特に [PPA (Personal Package Archives)][PPA] で提供されるパッケージについては全く問題なく移行できていた。
素晴らしい。

でも，もちろん上手くいかないパッケージもあるわけで，そういうのはちまちまと手動で書き換える必要がある。

注意しないといけないのは「ソフトウェアとアップデート」の「他のソフトウェア」で表示されなくなったサードパーティ・パッケージである。
どうも APT の標準で管理されていない電子署名鍵を使うパッケージは「他のソフトウェア」の一覧に表示されないようだ。

「APT の標準で管理されていない電子署名鍵」については以下の拙文を参考にどうぞ。

- [ついに apt-key コマンドに「非推奨」のワーニングが]({{< ref "/remark/2022/05/apt-key-is-deprecated.md" >}})
- [apt-key が非推奨になったので](https://zenn.dev/spiegel/articles/20220508-apt-key-is-deprecated)

というわけで，この機会に `/etc/apt/sources.list.d/` ディレクトリの中身を棚卸しして整理することをお勧めする。

ちなみに「他のソフトウェア」の管理外になった [pgAdmin4] は 2022-10-30 時点で `kinetic` リポジトリがまだ用意されていない `orz`

## ブックマーク

- [Ubuntu 22.10 日本語 Remix リリース | Ubuntu Japanese Team](ubuntu2210-ja-remix)
- [Ubuntu 22.04 その331 - Ubuntu 22.10へアップグレードするには・アップグレードの注意事項 - kledgeb](https://kledgeb.blogspot.com/2022/10/ubuntu-2204-331-ubuntu-2210.html)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[pgAdmin4]: https://www.pgadmin.org/ "pgAdmin - PostgreSQL Tools"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Node.js]: https://nodejs.org/
[OpenSSL]: https://www.openssl.org/
[OpenSSH]: https://www.openssh.com/
[Docker]: https://www.docker.com/ "Empowering App Development for Developers | Docker"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"

## 参考図書

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B0BG8J5QJ1" %}} <!-- ［試して理解］Linuxのしくみ 増補改訂版 -->
{{% review-paapi "B01NBU1OS5" %}} <!-- シリコンパワー USBメモリ 32GB USB3.1 -->
