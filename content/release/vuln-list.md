+++
title = "Ubuntu アプリケーションにおけるセキュリティ・アップデート一覧"
date =  "2019-07-31T07:43:30+09:00"
description = "Ubuntu の各アプリケーションにおける脆弱性発覚時のセキュリティ・アップデートがヒッジョーに分かりにくいためメモを作って管理することにした。"
image = "/images/attention/tools.png"
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

| Uistribution        | Release    |
| ------------------- | ---------- |
| 19.10 (Eoan Ermine) | 2019-10-17 |

- [Ubuntu 19.10 (Eoan Ermine) released](https://lists.ubuntu.com/archives/ubuntu-announce/2019-October/000250.html)
    - [EoanErmine/ReleaseNotes - Ubuntu Wiki](https://wiki.ubuntu.com/EoanErmine/ReleaseNotes)

## GnuPG

GnuPG は [Ubuntu] 19.04 リリース時の 2.2.12 (libgcrypt 1.8.4) から動いていない。
ちなみに [Ubuntu] 19.10 リリース時の公式版最新バージョンは 2.2.17 (libgcrypt 1.8.5) である。

| Official | Release    |  Ubuntu | Release |   delay |
| --------:| ---------- | -------:| ------- | -------:|
|   2.2.19 | 2019-12-07 |         |         | ∞ days |
|   2.2.18 | 2019-11-25 | &mdash; | &mdash; | &mdash; |
|   2.2.17 | 2019-07-09 | &mdash; | &mdash; | &mdash; |

### Libgcrypt

| Official | Release    |               Ubuntu | Release    |      delay |
| --------:| ---------- | --------------------:| ---------- | ----------:|
|    1.8.5 | 2019-08-29 | 1.8.4-5 (USN-4236-1) | 2020-01-13 | 4.5 months |

### ブックマーク

- [GnuPG 2.2.19 がリリースされた]({{< ref "/release/2019/12/gnupg-2_2_19-is-released.md" >}})
- [GnuPG 2.2.18 リリース： さようなら SHA-1]({{< ref "/release/2019/11/gnupg-2_2_18-is-released.md" >}})
- [Libgcrypt 1.8.5 がリリース【セキュリティ・アップデート】]({{< ref "/release/2019/08/libgcrypt-1_8_5-is-released.md" >}})
- [GnuPG 2.2.17 リリース： 公開鍵サーバ・アクセスに関する過激な変更あり]({{< ref "/release/2019/07/gnupg-2_2_17-is-released.md" >}})

- [Gnupg-announce Info Page](https://lists.gnupg.org/mailman/listinfo/gnupg-announce)

## Git

Git については [Ubuntu] 公式リポジトリではなく PPA のリポジトリからインストールしている。

| Official | Release    |        Ubuntu | Release    |  delay |
| --------:| ---------- | -------------:| ---------- | ------:|
|   2.25.0 | 2020-01-13 | 2.25.0-1~ppa0 | 2020-01-14 |  1 day |
|   2.24.1 | 2019-12-10 | 2.24.1-1~ppa0 | 2019-12-11 |  1 day |
|   2.24.0 | 2019-11-04 | 2.24.0-1~ppa0 | 2019-11-04 |  0 day |
|   2.23.0 | 2019-08-16 |  2.23.0-0ppa1 | 2019-08-18 | 2 days |

### ブックマーク

- [[ANNOUNCE] Git v2.25.0 - Junio C Hamano](https://lore.kernel.org/git/xmqqtv4zjgv5.fsf@gitster-ct.c.googlers.com/)
- [[ANNOUNCE] Git v2.24.1 and others](https://public-inbox.org/git/xmqqr21cqcn9.fsf@gitster-ct.c.googlers.com/T/)
- [[ANNOUNCE] Git v2.24.0](https://public-inbox.org/git/20191104173426.GA68471@syl.local/T/)
- [[ANNOUNCE] Git v2.23.0 - Junio C Hamano](https://public-inbox.org/git/xmqqy2zszuz7.fsf@gitster-ct.c.googlers.com/)

- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})

## Firefox

| Official | Release    |                       Ubuntu | Release    |   delay |
| --------:| ---------- | ----------------------------:| ---------- | -------:|
|   73.0.1 | 2020-02-18 | 73.0.1+build1-0 (USN-4278-3) | 2020-02-26 |  8 days |
|     73.0 | 2020-02-11 |   73.0+build3-0 (USN-4278-1) | 2020-02-13 |  2 days |
|   72.0.2 | 2020-01-20 | 72.0.2+build1-0 (USN-4234-2) | 2020-01-30 | 10 days |
|   72.0.1 | 2020-01-08 | 72.0.1+build1-0 (USN-4234-1) | 2020-01-09 |   1 day |
|     72.0 | 2020-01-07 |                      &mdash; | &mdash;    | &mdash; |
|     71.0 | 2019-12-03 |   71.0+build5-0 (USN-4216-1) | 2019-12-09 |  6 days |
|     70.0 | 2019-10-22 |   70.0+build2-0 (USN-4165-1) | 2019-10-23 |   1 day |
|   69.0.3 | 2019-10-10 |               (Ubuntu 19.10) | 2019-10-17 |  7 days |

### ブックマーク

- [Firefox  73.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/73.0.1/releasenotes/)
- [Firefox  73.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/73.0/releasenotes/)
- [Firefox  72.0.2, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/72.0.2/releasenotes/)
- [Firefox  72.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/72.0.1/releasenotes/)
    - [Firefox の脆弱性 (CVE-2019-17026) に関する注意喚起](https://www.jpcert.or.jp/at/2020/at200005.html)
- [Firefox  72.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/72.0/releasenotes/)
- [Firefox  71.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/71.0/releasenotes/)
- [Firefox  70.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/70.0.1/releasenotes/)
- [Firefox  70.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/70.0/releasenotes/)
- [Firefox  69.0.3, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/69.0.3/releasenotes/)

## Thunderbird

現在は APT を使わず公式バイナリをダウンロードして手動でインストールしているが，参考情報として残しておく。

| Official | Release    |                       Ubuntu | Release    |   delay |
| --------:| ---------- | ----------------------------:| ---------- | -------:|
|   68.5.0 | 2020-02-11 |                              |            | ∞ days |
|   68.4.2 | 2020-01-24 |                      &mdash; | &mdash;    | &mdash; |
|   68.4.1 | 2020-01-09 | 68.4.1+build1-0 (USN-4241-1) | 2020-01-16 |  7 days |
|   68.3.1 | 2019-12-16 |                      &mdash; | &mdash;    | &mdash; |
|   68.3.0 | 2019-12-03 |                      &mdash; | &mdash;    | &mdash; |
|   68.2.2 | 2019-11-07 | 68.2.2+build1-0 (USN-4202-2) | 2019-12-10 | 1 month |
|   68.2.1 | 2019-10-31 | 68.2.1+build1-0 (USN-4202-1) | 2019-11-26 | 4 weeks |
|   68.2.0 | 2019-10-22 |                      &mdash; | &mdash;    | &mdash; |
|   68.1.2 | 2019-10-10 |               (Ubuntu 19.10) | 2019-10-17 |  7 days |
|   68.1.1 | 2019-09-25 |                      &mdash; | &mdash;    | &mdash; |
|     68.1 | 2019-09-11 |                      &mdash; | &mdash;    | &mdash; |
|     68.0 | 2019-08-27 |                      &mdash; | &mdash;    | &mdash; |

### ブックマーク

- [Thunderbird — Release Notes (68.5.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/68.5.0/releasenotes/)
- [Thunderbird — Release Notes (68.4.2) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/68.4.2/releasenotes/)
- [Thunderbird — Release Notes (68.4.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/68.4.1/releasenotes/)
- [Thunderbird — Release Notes (68.3.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/68.3.1/releasenotes/)
- [Thunderbird — Release Notes (68.3.0) — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.3.0/releasenotes/)
- [Thunderbird — Release Notes (68.2.2) — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.2.2/releasenotes/)
- [Thunderbird — Release Notes (68.2.1) — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.2.1/releasenotes/)
- [Thunderbird — Release Notes (68.2.0) — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.2.0/releasenotes/)
- [Thunderbird — Release Notes (68.1.2) — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.1.2/releasenotes/)
- [Thunderbird — Release Notes (68.1.1) — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.1.1/releasenotes/)
- [Thunderbird — Release Notes (68.1.0) — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.1.0/releasenotes/)
- [Thunderbird — Release Notes (68.0) — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.0/releasenotes/)

- [結局 Thunderbird もインストールし直すことにした]({{< ref "/remark/2019/11/reinstalling-thunderbird.md" >}})

## LibreOffice

現在は APT を使わず公式バイナリをダウンロードして手動でインストールしているが，参考情報として残しておく。

| Official | Release    |         Ubuntu | Release    |   delay |
| --------:| ---------- | --------------:| ---------- | -------:|
|    6.3.4 | 2019-12-12 |                |            | ∞ days |
|    6.3.3 | 2019-10-31 |        &mdash; | &mdash;    | &mdash; |
|    6.3.2 | 2019-09-26 | (Ubuntu 19.10) | 2019-10-17 | 3 weeks |
|    6.3.1 | 2019-09-05 |        &mdash; | &mdash;    | &mdash; |
|      6.3 | 2019-08-08 |        &mdash; | &mdash;    | &mdash; |

6.2 以前のバージョンは公式サポートから外れているため使うべきではない。

### ブックマーク

- [LibreOffice 6.3.4 available for download - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/12/12/libreoffice-6-3-4/)
- [The Document Foundation releases LibreOffice 6.3.3 - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/10/31/tdf-releases-libreoffice-633/)
- [LibreOffice 6.2.8 is available, the last release of the 6.2 family - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/10/17/libreoffice-628/)
- [LibreOffice 6.3.1 and LibreOffice 6.2.7 announced, focusing on security - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/09/05/lo-6-3-1-and-lo-6-2-7-announced/)
- [The Document Foundation announces LibreOffice 6.3 - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/08/08/tdf-announces-libreoffice-63/)

- [Ubuntu に LibreOffice をインストールする3つの方法]({{< ref "/remark/2019/05/installing-libreoffice-in-ubuntu.md" >}})

## Java ([OpenJDK])

現在は APT を使わず公式バイナリをダウンロードして手動でインストールしている。

Java というか [OpenJDK] のセキュリティ管理は LTS (Long Term Support) バージョンのみ行っているらしく，ショートサイクルのバージョン（12, 13, ...）は管理がグダグダのようだ。
面倒くさくなったので LTS バージョンのみ参考情報として残しておく。

[OpenJDK]: http://openjdk.java.net/

| Official | Release    |                   Ubuntu | Release    |    delay |
| --------:| ---------- | ------------------------:| ---------- | --------:|
|   11.0.6 | 2020-01-14 | 11.0.6+10-1 (USN-4257-1) | 2020-01-28 |  2 weeks |
|   11.0.5 | 2019-10-15 | 11.0.5+10-0 (USN-4223-1) | 2019-12-17 | 2 months |
|   11.0.4 | 2019-07-16 | 11.0.4+11-1 (USN-4083-1) | 2019-07-31 |  3 weeks |
|   11.0.3 | 2019-04-16 |  11.0.3+7-1 (USN-3975-1) | 2019-05-14 |  4 weeks |

### ブックマーク

- [OpenJDK Vulnerability Advisory: 2020/01/14](https://openjdk.java.net/groups/vulnerability/advisories/2020-01-14)
- [OpenJDK Vulnerability Advisory: 2019/10/15](https://openjdk.java.net/groups/vulnerability/advisories/2019-10-15)
- [OpenJDK Vulnerability Advisory: 2019/7/16](https://openjdk.java.net/groups/vulnerability/advisories/2019-07-16)
    - [Oracle Java の脆弱性対策について(CVE-2019-7317等)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20190717-jre.html)
    - [USN-4083-1: OpenJDK 11 vulnerabilities | Ubuntu security notices](https://usn.ubuntu.com/4083-1/)

- [結局 OpenJDK をインストールし直すことにした]({{< ref "/remark/2019/07/reinstalling-openjdk.md" >}})
