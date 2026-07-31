+++
title = "Git v2.26.1 のリリース【セキュリティ・アップデート】"
date =  "2020-04-15T10:18:03+09:00"
description = "対象となるのは 2.17.x から 2.26.x までの各マイナーバージョン (CVE-2020-5260)"
image = "/images/attention/tools.png"
tags  = [ "git", "tools", "security", "vulnerability", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Git] v2.26.1 を含む複数のバージョンがリリースされた。

- [[Announce] Git v2.26.1 and others](https://lore.kernel.org/git/xmqqy2qy7xn8.fsf@gitster.c.googlers.com/T/)

対象となるのは 2.17.x から 2.26.x までの各マイナーバージョン。
今回はセキュリティ・アップデートを含むため必ず対応すること。

{{< fig-quote type="markdown" title="Git v2.26.1 and others" link="https://lore.kernel.org/git/xmqqy2qy7xn8.fsf@gitster.c.googlers.com/T/" lang="en" >}}
{{< quote >}}These releases address the security issue CVE-2020-5260, which
allowed a crafted URL to trick a Git client to send credential
information for a wrong host to the attacker's site{{< /quote >}}.
{{< /fig-quote >}}

[Ubuntu] の APT の標準リポジトリは最新版を提供してないが， [USN-4329-1](https://usn.ubuntu.com/4329-1/ "USN-4329-1: Git vulnerability | Ubuntu security notices") によると各バージョンの [git] でアップデートを行っているようだ。
なお [PPA 版リポジトリ](https://launchpad.net/~git-core/+archive/ubuntu/ppa "Git stable releases : “Ubuntu Git Maintainers” team")では既に v2.26.1 がリリースされている。

## [CVE-2020-5260](https://nvd.nist.gov/vuln/detail/CVE-2020-5260)

{{< fig-quote type="markdown" title="CVE-2020-5260" link="https://nvd.nist.gov/vuln/detail/CVE-2020-5260" lang="en" >}}
{{< quote >}}Affected versions of Git have a vulnerability whereby Git can be tricked into sending private credentials to a host controlled by an attacker. Git uses external "credential helper" programs to store and retrieve passwords or other credentials from secure storage provided by the operating system. Specially-crafted URLs that contain an encoded newline can inject unintended values into the credential helper protocol stream, causing the credential helper to retrieve the password for one server (e.g., good.example.com) for an HTTP request being made to another server (e.g., evil.example.com), resulting in credentials for the former being sent to the latter. There are no restrictions on the relationship between the two, meaning that an attacker can craft a URL that will present stored credentials for any host to a host of their choosing. The vulnerability can be triggered by feeding a malicious URL to git clone. However, the affected URLs look rather suspicious; the likely vector would be through systems which automatically clone URLs not visible to the user, such as Git submodules, or package systems built around Git{{< /quote >}}.
{{< /fig-quote >}}

説明がなげーよ `orz`

 要するに {{< quote lang="en" >}}an attacker can craft a URL that will present stored credentials for any host to a host of their choosing{{< /quote >}} ってところが重要。


- `CVSS:3.1/AV:N/AC:L/PR:N/UI:R/S:C/C:H/I:H/A:N` (GitHub, Inc.)
- 深刻度: 緊急 (9.3)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 要           |
| スコープ         | 変更あり     |
| 機密性への影響   | 高           |
| 完全性への影響   | 高           |
| 可用性への影響   | なし         |

## [Git] Credential Helper

ちなみに credential helper のひとつ `GNOME/libsecret` だが， [Ubuntu] 19.10 で用意されている APT 最新版は

```text
$ sudo apt show libsecret-1-dev
Package: libsecret-1-dev
Version: 0.18.8-2
Priority: optional
Section: libdevel
Source: libsecret
Origin: Ubuntu
...
```

だった。
ちょっと古いっぽい（？）気もするが... Linux 系独特の意味不明な backport patch はどうにかならないのだろうか。

まぁ，いいや。
[以前も書いた]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}} "PPA から Git をインストールする")がインストール手順は以下の通り（`/usr/share/...` 以下を汚したくなかったので）。

```text
$ sudo apt install libsecret-1-dev
$ mkdir ~/work
$ cp -r /usr/share/doc/git/contrib/credential/libsecret ~/work
$ cd ~/work/libsecret
$ make
gcc -g -O2 -Wall  -pthread -I/usr/include/libsecret-1 -I/usr/include/libmount -I/usr/include/blkid -I/usr/include/glib-2.0 -I/usr/lib/x86_64-linux-gnu/glib-2.0/include -o git-credential-libsecret.o -c git-credential-libsecret.c
gcc -o git-credential-libsecret  git-credential-libsecret.o -lsecret-1 -lgio-2.0 -lgobject-2.0 -lglib-2.0
```

これでビルドした `git-credential-libsecret` を `$PATH` の通ったディレクトリに放り込んでおけばよい。
確認は以下の通り。

```text {hl_lines=[4]}
$ git help -a | grep credential-
   credential-cache     Helper to temporarily store passwords in memory
   credential-store     Helper to store credentials on disk
   credential-libsecret
```

よーし，うむうむ，よーし。

[Git] 設定は以下の通り。

```text
$ git config --global credential.helper libsecret
```

## アップデートは...

アップデートは計画的に。

## ブックマーク

- [malicious URLs may cause Git to present stored credentials to the wrong server · Advisory · git/git · GitHub](https://github.com/git/git/security/advisories/GHSA-qm7j-c969-7j4q)
- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[git]: https://git-scm.com/
[Git]: https://git-scm.com/
[Git for Windows]: https://gitforwindows.org/
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
