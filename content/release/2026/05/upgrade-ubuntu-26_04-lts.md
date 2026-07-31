+++
title = "Ubuntu 26.04 LTS へのアップグレード"
date =  "2026-05-19T12:40:42+09:00"
description = "システムがなにか言ってくるまで様子見を決め込んでいたのだが，今朝ようやく通知が来た。"
isCJKLanguage = true
image = "/images/attention/tools.png"
tags = [ "ubuntu", "linux", "tools", "install", "ppa" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先月の4月に [Ubuntu] 26.04 LTS がリリースされた。

- [Ubuntu 26.04 ("Resolute Raccoon") LTS released](https://lists.ubuntu.com/archives/ubuntu-announce/2026-April/000323.html)
- [Ubuntu 26.04 LTS release notes](https://documentation.ubuntu.com/release-notes/26.04/)

急ぐ必要はないし，システムがなにか言ってくるまで様子見を決め込んでいたのだが，今朝ようやく通知が来た。

{{< fig-img src="./01-upgrade-ubuntu-26-04-lts.png" link="./01-upgrade-ubuntu-26-04-lts.png" width="733" >}}

とりあえずパッケージを最新状態にしてから

{{< fig-img src="./02-upgrade-ubuntu-26-04-lts.png" link="./02-upgrade-ubuntu-26-04-lts.png" width="582" >}}

アップグレード開始。

{{< fig-img src="./03-releasenote.png" link="./03-releasenote.png" width="613" >}}

このウィンドウが出たときに脳内 BGM があの[名曲](https://www.youtube.com/watch?v=aZ6wQG9Rhdk "ロックリバーへ｜『あらいぐまラスカル』オープニングテーマ - YouTube")になった（笑）

サードパーティのパッケージについてのアラート。

{{< fig-img src="./04-foreign-packages.png" link="./04-foreign-packages.png" width="609" >}}

今はスクショ撮って AI に「何書いてあんの？」って聞けば教えてくれるので，ホンマ助かる。
ここは大人しく「はい」を選択して続行する。
サードパーティ・パッケージのリポジトリは一旦無効化されるが，アップグレード後に手動で回復させる。

アップグレード自体は特にトラブルなく完了した。

{{< fig-img src="./05-reboot.png" link="./05-reboot.png" width="675" >}}

再起動も問題ないかな。

{{< fig-img src="./06-reboot.png" link="./06-reboot.png" width="1031" >}}

このウィンドウが出たときに脳内 BGM が（以下略）

このあと簡単な設定を行う。

{{< fig-img src="./07-setting.png" link="./07-setting.png" width="1031" >}}
{{< fig-img src="./08-setting.png" link="./08-setting.png" width="1031" >}}
{{< fig-img src="./09-setting.png" link="./09-setting.png" width="1031" >}}
{{< fig-img src="./10-setting.png" link="./10-setting.png" width="1031" >}}

んー。
こんな感じかな。
バージョン等を確認しておく。

```text
$ cat /etc/os-release
PRETTY_NAME="Ubuntu 26.04 LTS"
NAME="Ubuntu"
VERSION_ID="26.04"
VERSION="26.04 LTS (Resolute Raccoon)"
VERSION_CODENAME=resolute
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=resolute
LOGO=ubuntu-logo
```

セキュリティ関連ツールもチェックしておくか。

```text
$ gpg --version
gpg (GnuPG) 2.4.8
libgcrypt 1.12.0
Copyright (C) 2025 g10 Code GmbH
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

うーん。
[GnuPG] を 2.5 系に上げる気はなさそうだな。
今ちょっと[面倒くさい](https://kris.fail/posts/opgpvslpgp/ "OpenPGPとLibrePGP―GnuPGとそれ以外の実装での対立")ことになってるし，しょうがないか。

```text
$ ssh -V
OpenSSH_10.2p1 Ubuntu-2ubuntu3.2, OpenSSL 3.5.5 27 Jan 2026
```

[OpenSSH] も [OpenSSL] も微妙にバージョンが古い気がするが... まぁいいか。

動作確認で Snap 版 [Remmina] を起動したら，こんなのが表示される。

{{< fig-img src="./remmina-snap.png" link="./remmina-snap.png" width="643" >}}

あれ？ これ設定がリセットされてる？ 念のためもう一度設定しておこう。

```text
$ sudo snap connect remmina:audio-record :audio-record
$ sudo snap connect remmina:avahi-observe :avahi-observe
$ sudo snap connect remmina:cups-control :cups-control
$ sudo snap connect remmina:mount-observe :mount-observe
$ sudo snap connect remmina:password-manager-service :password-manager-service
$ sudo snap connect remmina:ssh-keys :ssh-keys
$ sudo snap connect remmina:ssh-public-keys :ssh-public-keys
```

サードパーティのパッケージを作業の覚え書きとして挙げておく。

- [x] [PPA] 版 [Git]
  - [Git stable releases : “Ubuntu Git Maintainers” team](https://launchpad.net/~git-core/+archive/ubuntu/ppa)
- [x] [GitHub CLI]
    - [Installing gh on Linux and BSD](https://github.com/cli/cli/blob/trunk/docs/install_linux.md)
- [x] [NodeSource] 版 [Node.js]
  - [Node.js Distributions](https://nodesource.com/products/distributions)
    - [v24]({{< ref "/release/2025/05/nodejs-v24-is-released.md" >}} "Node.js v24 がリリースされた") LTS 版を選択
- [x] [pgAdmin 4]
  - [Download](https://www.pgadmin.org/download/pgadmin-4-apt/)
- [x] [VS Code]
  - [Running Visual Studio Code on Linux](https://code.visualstudio.com/docs/setup/linux)
- [x] [xtradeb applications](https://launchpad.net/~xtradeb/+archive/ubuntu/apps "xtradeb applications : “xtradeb packaging” team")
  - [Ungoogled Chromium]
  - [KeePassXC]

おー。
全部 `resolute` バージョンに対応してるな。
待ってた甲斐があったな（笑）

## 「ソフトウェアとアップデート」をインストールする

[Ubuntu] 26.04 LTS から APT リポジトリを GUI で管理できる「ソフトウェアとアップデート」がなくなってしまった。
ただ，リポジトリにはあるらしいのでインストール可能。

「ソフトウェアとアップデート」のパッケージ名って何だっけ？ よく分からないのでアプリセンターで検索したらあったよ。

{{< fig-img src="./app-center.png" link="./app-center.png" width="1342" >}}

ここからインストールして事なきを得た。

多分 [Ubuntu] 側はアプリセンターに統一したいんだろうけど APT リポジトリ管理が隠蔽されてるのは困るのよ。
サードパーティのリポジトリ設定が正しいか確認するのに GUI の「ソフトウェアとアップデート」を使ってるし OS アップグレード設定もこれで行っている。

もうちょっと考えてほしい。
少なくとも LTS 版でやることじゃないよなぁ。

## ブックマーク

- [Ubuntu 26.04 LTS “Resolute Raccoon”のリリース | gihyo.jp](https://gihyo.jp/admin/clip/01/ubuntu-topics/202604/24)
- [第908回　Ubuntu 26.04 LTSの変更点 ［25.10→26.04 LTS編］ | gihyo.jp](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0908)
- [第909回　Ubuntu 26.04 LTSの変更点 ［24.04→26.04 LTS編］ | gihyo.jp](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0909)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: https://www.openssh.com/
[OpenSSL]: https://www.openssl.org/
[Git]: https://git-scm.com/
[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[NodeSource]: https://github.com/nodesource
[Node.js]: https://nodejs.org/
[pgAdmin 4]: https://www.pgadmin.org/ "pgAdmin - PostgreSQL Tools"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Ungoogled Chromium]: https://github.com/ungoogled-software/ungoogled-chromium "ungoogled-software/ungoogled-chromium: Google Chromium, sans integration with Google"
[Remmina]: https://www.remmina.org/ "Remmina remote desktop client - Remmina"
[GitHub CLI]: https://cli.github.com/ "GitHub CLI"

## 参考

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B01MS6RPN2" %}} <!-- シリコンパワー USBメモリ 8GB USB3.1 -->
