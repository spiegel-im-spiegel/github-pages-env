+++
title = "Ubuntu アプリケーションにおけるセキュリティ・アップデート一覧"
date =  "2019-07-31T07:43:30+09:00"
description = "Ubuntu の各アプリケーションにおける脆弱性発覚時のセキュリティ・アップデートがヒッジョーに分かりにくいためメモを作って管理することにした。"
tags = [ "security", "vulnerability", "risk", "tools", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] の各アプリケーションにおける脆弱性発覚時のセキュリティ・アップデートがヒッジョーに分かりにくいためメモを作って管理することにした。
なお，ここのメモは私がよく使うツールに限定しているため全く網羅的ではないが悪しからず。
また kernel やライブラリのセキュリティ・アップデートは（今のところは）除外する。

- [Ubuntu security notices](https://usn.ubuntu.com/)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## [Ubuntu] ディストリビューションのリリース

| Uistribution        | Release    | Expire  |
| ------------------- | ---------- | ------- |
| 20.04 (Focal Fossa) | 2020-04-23 | 2025-04 |
| 19.10 (Eoan Ermine) | 2019-10-17 | 2020-07 |

- [FocalFossa/ReleaseNotes - Ubuntu Wiki](https://wiki.ubuntu.com/FocalFossa/ReleaseNotes)
- [EoanErmine/ReleaseNotes - Ubuntu Wiki](https://wiki.ubuntu.com/EoanErmine/ReleaseNotes)

## GnuPG

[Ubuntu] 20.04 LTS リリース時の最新バージョンは 2.2.19 (Libgcrypt 1.8.5) である。

| Official | Release    |         Ubuntu | Release    |      Delay |
| --------:| ---------- | --------------:| ---------- | ----------:|
|   2.2.20 | 2020-03-20 |                |            |    ∞ days |
|   2.2.19 | 2019-12-07 | (Ubuntu 19.10) | 2020-04-23 | 4.5 months |

### Libgcrypt

| Official | Release    |               Ubuntu | Release    |      Delay |
| --------:| ---------- | --------------------:| ---------- | ----------:|
|    1.8.5 | 2019-08-29 |       (Ubuntu 19.10) | 2020-04-23 |    &mdash; |
|    1.8.5 | 2019-08-29 | 1.8.4-5 (USN-4236-1) | 2020-01-13 | 4.5 months |

### ブックマーク

- [GnuPG 2.2.20 がリリースされた]({{< ref "/release/2020/03/gnupg-2_2_20-is-released.md" >}})
- [GnuPG 2.2.19 がリリースされた]({{< ref "/release/2019/12/gnupg-2_2_19-is-released.md" >}})
- [Libgcrypt 1.8.5 がリリース【セキュリティ・アップデート】]({{< ref "/release/2019/08/libgcrypt-1_8_5-is-released.md" >}})

- [Gnupg-announce Info Page](https://lists.gnupg.org/mailman/listinfo/gnupg-announce)

## Git

Git については [Ubuntu] 公式リポジトリではなく [PPA のリポジトリ](https://launchpad.net/~git-core/+archive/ubuntu/ppa)からインストールしている。

| Official | Release    |        Ubuntu | Release    |  Delay |
| --------:| ---------- | -------------:| ---------- | ------:|
|   2.26.2 | 2020-04-20 |  2.26.2-0ppa1 | 2020-04-21 | 1 days |
|   2.26.1 | 2020-04-14 |  2.26.1-0ppa1 | 2020-04-15 | 1 days |
|   2.26.0 | 2020-03-22 | 2.26.0-1~ppa1 | 2020-03-24 | 2 days |

### ブックマーク

- [[Announce] Git v2.26.2 and others](https://lore.kernel.org/git/xmqq4kterq5s.fsf@gitster.c.googlers.com/T/)
- [[Announce] Git v2.26.1 and others](https://lore.kernel.org/git/xmqqy2qy7xn8.fsf@gitster.c.googlers.com/T/)
- [[ANNOUNCE] Git v2.26.0 - Junio C Hamano](https://lore.kernel.org/git/xmqqa7477u6j.fsf@gitster.c.googlers.com/)

- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})

## Firefox

| Official | Release    |                       Ubuntu | Release    |   Delay |
| --------:| ---------- | ----------------------------:| ---------- | -------:|
|     75.0 | 2020-04-07 |               (Ubuntu 20.04) | 2020-04-23 | &mdash; |
|     75.0 | 2020-04-07 |   75.0+build3-0 (USN-4323-1) | 2020-04-07 |   0 day |

### ブックマーク

- [Firefox  75.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/75.0/releasenotes/)

## Thunderbird

現在は APT を使わず公式バイナリをダウンロードして手動でインストールしているが，参考情報として残しておく。

| Official | Release    |                       Ubuntu | Release    |   Delay |
| --------:| ---------- | ----------------------------:| ---------- | -------:|
|   68.8.0 | 2020-05-05 |                              |            | ∞ days |
|   68.7.0 | 2020-04-08 |               (Ubuntu 20.04) | 2020-04-23 | &mdash; |
|   68.7.0 | 2020-04-08 | 68.7.0+build1-0 (USN-4328-1) | 2020-04-13 |  5 days |

### ブックマーク

- [Thunderbird — Release Notes (68.8.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/68.8.0/releasenotes/)
- [Thunderbird — Release Notes (68.7.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/68.7.0/releasenotes/)

- [結局 Thunderbird もインストールし直すことにした]({{< ref "/remark/2019/11/reinstalling-thunderbird.md" >}})

## Java ([OpenJDK])

現在は APT を使わず公式バイナリをダウンロードして手動でインストールしている。

Java というか [OpenJDK] のセキュリティ管理は LTS (Long Term Support) バージョンのみ行っているらしく，ショートサイクルのバージョン（12, 13, ...）は管理がグダグダのようだ。
面倒くさくなったので LTS バージョンのみ参考情報として残しておく。

[OpenJDK]: http://openjdk.java.net/

| Official | Release    |                   Ubuntu | Release    |    Delay |
| --------:| ---------- | ------------------------:| ---------- | --------:|
|   11.0.7 | 2020-04-14 |           (Ubuntu 20.04) | 2020-04-23 |  &mdash; |
|   11.0.7 | 2020-04-14 | 11.0.7+10-2 (USN-4337-1) | 2020-04-22 |   8 days |

### ブックマーク

- [OpenJDK Vulnerability Advisory: 2020/04/14](https://openjdk.java.net/groups/vulnerability/advisories/2020-04-14)

- [結局 OpenJDK をインストールし直すことにした]({{< ref "/remark/2019/07/reinstalling-openjdk.md" >}})
