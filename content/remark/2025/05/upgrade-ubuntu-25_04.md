+++
title = "Ubuntu 25.04 へのアップグレード"
date =  "2025-05-16T20:15:43+09:00"
description = "description"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools", "install", "ppa" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

予定通り 2025-04 に [Ubuntu] 25.04 がリリースされた。

- [Ubuntu 25.04 (Plucky Puffin) released](https://lists.ubuntu.com/archives/ubuntu-announce/2025-April/000311.html)
- [Plucky Puffin Release Notes - Release - Ubuntu Community Hub](https://discourse.ubuntu.com/t/plucky-puffin-release-notes/48687)
- [Download Ubuntu Desktop | Ubuntu](https://ubuntu.com/download/desktop)
- [Ubuntu 25.04 “Plucky Puffin”のリリース | gihyo.jp](https://gihyo.jp/admin/clip/01/ubuntu-topics/202504/18)

まぁ急がないのでシステムがなにか言ってくるまで待機してた。

{{< fig-img src="./01-upgrade-ubuntu-25-04.png" link="./01-upgrade-ubuntu-25-04.png" width="774" >}}

よしよし。
じゃあ始めるか。

{{< fig-img src="./02-releasenote.png" link="./02-releasenote.png" width="722" >}}

おっ。
サードパーティのパッケージについてアラートが表示されるようになったのか。

{{< fig-img src="./04-foreign-packages.png" link="./04-foreign-packages.png" width="621" >}}

助かるなぁ。
とりあえずこのまま続行する。

というわけで，特にトラブルなく完了した。

{{< fig-img src="./08-reboot.png" link="./08-reboot.png" width="788" >}}

バージョン等を確認。

```text
$ cat /etc/os-release
PRETTY_NAME="Ubuntu 25.04"
NAME="Ubuntu"
VERSION_ID="25.04"
VERSION="25.04 (Plucky Puffin)"
VERSION_CODENAME=plucky
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=plucky
LOGO=ubuntu-logo
```

サードパーティのパッケージも手動でアップデートする。

- [x] [Docker] Engine
  - [Install Docker Engine on Ubuntu | Docker Docs](https://docs.docker.com/engine/install/ubuntu/)
- [x] [PPA] 版 [Git]
  - [Git stable releases : “Ubuntu Git Maintainers” team](https://launchpad.net/~git-core/+archive/ubuntu/ppa)
- [x] [NodeSource] 版 [Node.js]
  - [NodeSource Node.js Binary Distributions](https://github.com/nodesource/distributions/blob/master/README.md)
    - [v24]({{< ref "/release/2025/05/nodejs-v24-is-released.md" >}} "Node.js v24 がリリースされた") へアップグレードした。秋に LTS 版になる予定
- [x] [pgAdmin 4]
  - [Download](https://www.pgadmin.org/download/pgadmin-4-apt/)
- [x] [PPA] 版 [KeePassXC]
  - [KeePassXC : Janek Bevendorff](https://launchpad.net/~phoerious/+archive/ubuntu/keepassxc)
- [x] [VS Code]
  - [Running Visual Studio Code on Linux](https://code.visualstudio.com/docs/setup/linux)

いやぁ，今回は楽でよかった。

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
[Fcitx 5]: https://fcitx-im.org/wiki/Fcitx_5

## 参考

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B01MS6RPN2" %}} <!-- シリコンパワー USBメモリ 8GB USB3.1 -->
