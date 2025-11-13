+++
title = "Ubuntu 25.10 へのアップグレード"
date =  "2025-11-13T11:24:45+09:00"
description = "ホンマに X.Org 関連のパッケージが削除対象になってるよ。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools", "install", "ppa" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

予定通り 2025-10 に [Ubuntu] 25.10 がリリースされ，アップグレードしろと頻繁に通知が来るのだが

{{< fig-img src="./01-upgrade-ubuntu-25-10.png" link="./01-upgrade-ubuntu-25-10.png" width="774" >}}

しばらく様子見していた。
でも，そろそろアップグレードしておかないとな。

- [Questing Quokka Release Notes - Release - Ubuntu Community Hub](https://discourse.ubuntu.com/t/questing-quokka-release-notes/59220)
- [Download Ubuntu Desktop | Ubuntu](https://ubuntu.com/download/desktop)

とりあえずパッケージを最新状態にしておいてから

{{< fig-img src="./01-upgrade-ubuntu-25-10-2.png" link="./01-upgrade-ubuntu-25-10-2.png" width="645" >}}

アップグレード開始。

{{< fig-img src="./02-releasenote.png" link="./02-releasenote.png" width="722" >}}

サードパーティのパッケージについてのアラート。

{{< fig-img src="./04-foreign-packages.png" link="./04-foreign-packages.png" width="621" >}}

試しに「いいえ (N)」を押したらアップグレード自体が中止された。
ここは素直に「はい (Y)」を押して続行。

パッケージの更新は滞りなく終わって，使わなくなったパッケージを削除するか訊いてくるのだが

{{< fig-img src="./07-delete-packages-2.png" link="./07-delete-packages-2.png" width="1102" >}}

おー。
ホンマに X.Org 関連のパッケージが削除対象になってるよ。

- [Ubuntu 25.10 その12 - Xorg セッションの削除と Wayland セッションに一本化 - kledgeb](https://kledgeb.blogspot.com/2025/09/ubuntu-2510-12-xorg-wayland.html)

残しておく理由はないので「削除 (R)」を実行。

というわけで，特にトラブルなく完了した。

{{< fig-img src="./08-reboot.png" link="./08-reboot.png" width="788" >}}

再起動後も特に問題はなさそうだ。
バージョン等を確認しておく。

```text
$ cat /etc/os-release
PRETTY_NAME="Ubuntu 25.10"
NAME="Ubuntu"
VERSION_ID="25.10"
VERSION="25.10 (Questing Quokka)"
VERSION_CODENAME=questing
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=questing
LOGO=ubuntu-logo
```

サードパーティのパッケージも手動でアップデートする。

- [x] [Docker] Engine
  - [Install Docker Engine on Ubuntu | Docker Docs](https://docs.docker.com/engine/install/ubuntu/)
- [x] [PPA] 版 [Git]
  - [Git stable releases : “Ubuntu Git Maintainers” team](https://launchpad.net/~git-core/+archive/ubuntu/ppa)
- [x] [NodeSource] 版 [Node.js]
  - [Node.js Distributions](https://nodesource.com/products/distributions)
    - [v24]({{< ref "/release/2025/05/nodejs-v24-is-released.md" >}} "Node.js v24 がリリースされた") LTS 版を選択
- [x] [pgAdmin 4]
  - [Download](https://www.pgadmin.org/download/pgadmin-4-apt/)
- [ ] [PPA] 版 [KeePassXC]
  - [KeePassXC : Janek Bevendorff](https://launchpad.net/~phoerious/+archive/ubuntu/keepassxc)
- [x] [VS Code]
  - [Running Visual Studio Code on Linux](https://code.visualstudio.com/docs/setup/linux)

今回も楽でよかった。

## ブックマーク

- [Ubuntu 25.10 “Questing Quokka”のリリース | gihyo.jp](https://gihyo.jp/admin/clip/01/ubuntu-topics/202510/10)
- [Ubuntu 25.10 その32 - Ubuntu 25.10 がリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2025/10/ubuntu-2510-32-ubuntu-2510.html)
  - [Ubuntu 25.10 その33 - Ubuntu 25.10 の新機能と変更点・既知の問題 - kledgeb](https://kledgeb.blogspot.com/2025/10/ubuntu-2510-33-ubuntu-2510.html)
  - [Ubuntu 25.10 その34 - Ubuntu 25.10 の注目機能 - kledgeb](https://kledgeb.blogspot.com/2025/10/ubuntu-2510-34-ubuntu-2510.html)
  - [Ubuntu 25.10 その37 - Ubuntu 25.10 で導入されたセキュリティーの強化とまとめ - kledgeb](https://kledgeb.blogspot.com/2025/10/ubuntu-2510-37-ubuntu-2510.html)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: https://www.openssh.com/
[OpenSSL]: https://www.openssl.org/
[Git]: https://git-scm.com/
[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[NodeSource]: https://github.com/nodesource
[Node.js]: https://nodejs.org/
[Docker]: https://www.docker.com/ "Empowering App Development for Developers | Docker"
[pgAdmin 4]: https://www.pgadmin.org/ "pgAdmin - PostgreSQL Tools"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"

## 参考

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B01MS6RPN2" %}} <!-- シリコンパワー USBメモリ 8GB USB3.1 -->
