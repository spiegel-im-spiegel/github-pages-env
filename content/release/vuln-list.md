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

## GnuPG

GnuPG は [Ubuntu] 19.04 リリース時の 2.2.12 (libgcrypt 1.8.4) から動いていない。
ちなみに [Ubuntu] 19.04 リリース時の公式版最新バージョンは 2.2.15 である。

| Official | Release    | Ubuntu | Release |   delay |
| --------:| ---------- | ------:| ------- | -------:|
|   2.2.17 | 2019-07-09 |        |         | ∞ days |

### Libgcrypt

| Official | Release    | Ubuntu | Release |   delay |
| --------:| ---------- | ------:| ------- | -------:|
|    1.8.5 | 2019-08-29 |        |         | ∞ days |

### ブックマーク

- [GnuPG 2.2.17 リリース： 公開鍵サーバ・アクセスに関する過激な変更あり]({{< ref "/release/2019/07/gnupg-2_2_17-is-released.md" >}})

- [Gnupg-announce Info Page](https://lists.gnupg.org/mailman/listinfo/gnupg-announce)

## Git

Git については [Ubuntu] 公式リポジトリではなく PPA のリポジトリからインストールしている。

| Official | Release    |       Ubuntu | Release    |  delay |
| --------:| ---------- | ------------:| ---------- | ------:|
|     2.23 | 2019-08-16 | 2.23.0-0ppa1 | 2019-08-18 | 2 days |

### ブックマーク

- [[ANNOUNCE] Git v2.23.0 - Junio C Hamano](https://public-inbox.org/git/xmqqy2zszuz7.fsf@gitster-ct.c.googlers.com/)

- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})

## Firefox

| Official | Release    |                       Ubuntu | Release    |   delay |
| --------:| ---------- | ----------------------------:| ---------- | -------:|
|   69.0.1 | 2019-09-18 |                              |            | ∞ days |
|     69.0 | 2019-09-03 |   69.0+build2-0 (USN-4122-1) | 2019-09-03 |  1 days |
|   68.0.2 | 2019-08-14 | 68.0.2+build1-0 (USN-4101-1) | 2019-08-16 |  2 days |
|   68.0.1 | 2019-07-18 | 68.0.1+build1-0 (USN-4054-1) | 2019-07-25 |  7 days |
|     68.0 | 2019-07-09 |   68.0+build3-0 (USN-4054-1) | 2019-07-12 |  3 days |
|   67.0.4 | 2019-06-20 | 67.0.4+build1-0 (USN-4032-1) | 2019-06-24 |  4 days |
|   67.0.3 | 2019-06-18 | 67.0.3+build1-0 (USN-4020-1) | 2019-07-20 |  2 days |

### ブックマーク

- [Firefox  69.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/69.0.1/releasenotes/)
- [Firefox  69.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/69.0/releasenotes/)
- [Firefox  68.0.2, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/68.0.2/releasenotes/)
- [Firefox  68.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/68.0.1/releasenotes/)
- [Firefox  68.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/68.0/releasenotes/)
- [Firefox  67.0.4, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/67.0.4/releasenotes/)
- [Firefox  67.0.3, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/67.0.3/releasenotes/)
    - [Firefox の脆弱性 (CVE-2019-11707) に関する注意喚起](https://www.jpcert.or.jp/at/2019/at190027.html)

## Thunderbird

| Official | Release    |                       Ubuntu | Release    |   delay |
| --------:| ---------- | ----------------------------:| ---------- | -------:|
|     68.0 | 2019-08-27 |                              |            | ∞ days |
|   60.8.0 | 2019-07-09 | 60.8.0+build1-0 (USN-4064-1) | 2019-07-17 |  8 days |
|   60.7.2 | 2019-06-20 | 60.7.2+build2-0 (USN-4045-1) | 2019-07-01 | 11 days |

### ブックマーク

- [Thunderbird — Release Notes (68.0) — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.0/releasenotes/)
- [フリーのメールソフト「Thunderbird」v60.8.0が正式公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1195473.html)
- [「Thunderbird」v60.7.2が公開 ～「Firefox」で悪用されているゼロデイ脆弱性に対処 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1191820.html)

## LibreOffice

現在は APT を使わず公式バイナリをダウンロードして手動でインストールしているが，参考情報として残しておく。
なお [Ubuntu] 19.04 では 6.2 系のみサポートされているようだ。

| Official | Release    | Ubuntu | Release |   delay |
| --------:| ---------- | ------:| ------- | -------:|
|      6.3 | 2019-08-08 |        |         | ∞ days |

| Official | Release    |               Ubuntu | Release    |  delay |
| --------:| ---------- | --------------------:| ---------- | ------:|
|    6.2.6 | 2019-08-14 | 6.2.6-0 (USN-4102-1) | 2019-08-19 | 5 days |
|    6.2.5 | 2019-07-04 | 6.2.5-0 (USN-4063-1) | 2019-07-12 | 8 days |

6.1 以前のバージョンは公式サポートから外れているため使うべきではない。

### ブックマーク

- [The Document Foundation announces LibreOffice 6.3 - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/08/08/tdf-announces-libreoffice-63/)

- [The Document Foundation announces LibreOffice 6.2.5 - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/07/04/tdf-announces-libreoffice-625/)
- [LibreOffice 6.2.6 is ready, all users should update for enhanced security - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/08/14/libreoffice-626/)

- [Ubuntu に LibreOffice をインストールする3つの方法]({{< ref "/remark/2019/05/installing-libreoffice-in-ubuntu.md" >}})

## Java ([OpenJDK])

現在は APT を使わず公式バイナリをダウンロードして手動でインストールしている。

Java というか [OpenJDK] のセキュリティ管理は LTS (Long Term Support) バージョンのみ行っているらしく，ショートサイクルのバージョン（12, 13, ...）は管理がグダグダのようだ。
面倒くさくなったので LTS バージョンのみ参考情報として残しておく。

[OpenJDK]: http://openjdk.java.net/

| Official | Release    |                   Ubuntu | Release    |   delay |
| --------:| ---------- | ------------------------:| ---------- | -------:|
|   11.0.4 | 2019-07-16 | 11.0.4+11-1 (USN-4083-1) | 2019-07-31 | 15 days |
|   11.0.3 | 2019-04-16 |  11.0.3+7-1 (USN-3975-1) | 2019-05-14 | 28 days |

### ブックマーク

- [OpenJDK Vulnerability Advisory: 2019/7/16](https://openjdk.java.net/groups/vulnerability/advisories/2019-07-16)
    - [Oracle Java の脆弱性対策について(CVE-2019-7317等)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20190717-jre.html)
    - [USN-4083-1: OpenJDK 11 vulnerabilities | Ubuntu security notices](https://usn.ubuntu.com/4083-1/)

- [結局 OpenJDK をインストールし直すことにした]({{< ref "/remark/2019/07/reinstalling-openjdk.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
