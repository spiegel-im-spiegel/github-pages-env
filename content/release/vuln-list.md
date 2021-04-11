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

[Ubuntu] アプリケーションのチェックは最新ディストリビューション下で行う。
現在有効なディストリビューション・バージョンは以下の通り。

| Uistribution           | Release    | Expire        |
| ---------------------- | ---------- | ------------- |
| 20.10 (Groovy Gorilla) | 2020-10-22 | 2021-07       |
| 20.04 (Focal Fossa)    | 2020-04-23 | 2025-04 (LTS) |

- [Groovy Gorilla Release Notes - Release - Ubuntu Community Hub](https://discourse.ubuntu.com/t/groovy-gorilla-release-notes/15533) ([日本語](https://wiki.ubuntu.com/GroovyGorilla/ReleaseNotes/Ja))
- [FocalFossa/ReleaseNotes - Ubuntu Wiki](https://wiki.ubuntu.com/FocalFossa/ReleaseNotes)
- [EoanErmine/ReleaseNotes - Ubuntu Wiki](https://wiki.ubuntu.com/EoanErmine/ReleaseNotes)

## [GnuPG]

[Ubuntu] 20.10 リリース時の最新バージョンは 2.2.20 (Libgcrypt 1.8.5) である。
現在は[自前でビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")して運用している。
以下は参考情報として記しておく。

| Official | Release    |         Ubuntu | Release    |   Delay |
| -------: | ---------- | -------------: | ---------- | ------: |
|   2.2.26 | 2020-12-21 |                |            |  ∞ days |
|   2.2.25 | 2020-11-23 |        &mdash; | &mdash;    | &mdash; |
|   2.2.24 | 2020-11-17 |        &mdash; | &mdash;    | &mdash; |
|   2.2.23 | 2020-09-03 |        &mdash; | &mdash;    | &mdash; |
|   2.2.22 | 2020-08-27 |        &mdash; | &mdash;    | &mdash; |
|   2.2.21 | 2020-07-09 |        &mdash; | &mdash;    | &mdash; |
|   2.2.20 | 2020-03-20 | (Ubuntu 20.10) | 2020-10-22 | &mdash; |

### Libgcrypt

| Official | Release    |         Ubuntu | Release    |   Delay |
| -------: | ---------- | -------------: | ---------- | ------: |
|    1.8.7 | 2020-10-23 |                |            |  ∞ days |
|    1.8.6 | 2020-07-06 |        &mdash; | &mdash;    | &mdash; |
|    1.8.5 | 2019-08-29 | (Ubuntu 20.10) | 2020-10-22 | &mdash; |

### ブックマーク

- [GnuPG 2.2.26 がリリースされた]({{< ref "/release/2020/12/gnupg-2_2_26-is-released.md" >}})
- [GnuPG 2.2.25 がリリースされた]({{< ref "/release/2020/11/gnupg-2_2_25-is-released.md" >}})
- [GnuPG 2.2.24 がリリースされた]({{< ref "/release/2020/11/gnupg-2_2_24-is-released.md" >}})
- [GnuPG 2.2.23 のリリース【セキュリティ・アップデート】]({{< ref "/release/2020/09/gnupg-2_2_23-is-released.md" >}})
- [GnuPG 2.2.22 がリリースされた]({{< ref "/release/2020/08/gnupg-2_2_22-is-released.md" >}})
- [GnuPG 2.2.21 および Libgcrypt 1.8.6 がリリースされた]({{< ref "/release/2020/07/gnupg-2_2_21-is-released.md" >}})
- [GnuPG 2.2.20 がリリースされた]({{< ref "/release/2020/03/gnupg-2_2_20-is-released.md" >}})
- [GnuPG 2.2.19 がリリースされた]({{< ref "/release/2019/12/gnupg-2_2_19-is-released.md" >}})
- [Libgcrypt 1.8.5 がリリース【セキュリティ・アップデート】]({{< ref "/release/2019/08/libgcrypt-1_8_5-is-released.md" >}})

- [Gnupg-announce Info Page](https://lists.gnupg.org/mailman/listinfo/gnupg-announce)

[GnuPG]: https://www.gnupg.org/ "The GNU Privacy Guard"

## [OpenSSL]

[OpenSSL]: https://www.openssl.org/

[OpenSSL] はディストリビューションがアップグレードされない限り（脆弱性があっても）更新されないようだ（バックポート・パッチで対応？）。
よって[手動で管理する]({{< ref "/remark/2020/06/installing-openssl-in-ubuntu.md" >}} "Ubuntu に最新の OpenSSL を入れる")ことにした。

## [Git]

[Git]: https://git-scm.com/

[Git] については [Ubuntu] 公式リポジトリではなく [PPA のリポジトリ](https://launchpad.net/~git-core/+archive/ubuntu/ppa)からインストールしている。

- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})

## [Firefox]

| Official | Release    |                       Ubuntu | Release    |   Delay |
| -------: | ---------- | ---------------------------: | ---------- | ------: |
|     87.0 | 2021-03-23 |   87.0+build3-0 (USN-4893-1) | 2021-02-26 |  3 days |
|   86.0.1 | 2021-03-11 |                      &mdash; | &mdash;    | &mdash; |
|     86.0 | 2021-02-23 |   86.0+build3-0 (USN-4756-1) | 2021-02-26 |  3 days |
|   85.0.2 | 2021-02-09 |                      &mdash; | &mdash;    | &mdash; |
|   85.0.1 | 2021-02-05 | 85.0.1+build1-0 (USN-4717-2) | 2021-02-08 |  3 days |
|     85.0 | 2021-01-26 |   85.0+build1-0 (USN-4717-1) | 2021-02-01 |  7 days |
|   84.0.2 | 2021-01-06 | 84.0.2+build1-0 (USN-4687-1) | 2021-01-08 |  2 days |
|   84.0.1 | 2020-12-22 |                      &mdash; | &mdash;    | &mdash; |
|     84.0 | 2020-12-15 |   84.0+build3-0 (USN-4671-1) | 2020-12-15 |   0 day |
|     83.0 | 2020-11-17 |   83.0+build2-0 (USN-4637-1) | 2020-11-18 |   1 day |
|   82.0.3 | 2020-11-10 | 82.0.3+build1-0 (USN-4625-1) | 2020-11-10 |   0 day |
|   82.0.2 | 2020-10-28 | 82.0.2+build1-0 (USN-4599-3) | 2020-11-05 |  8 days |
|   82.0.1 | 2020-10-27 |                      &mdash; | &mdash;    | &mdash; |
|     82.0 | 2020-10-20 |   82.0+build2-0 (USN-4599-1) | 2020-10-23 |  3 days |
|   81.0.2 | 2020-10-13 | 81.0.2+build1-0 (USN-4546-2) | 2020-10-15 |  2 days |
|   81.0.1 | 2020-10-01 |                      &mdash; | &mdash;    | &mdash; |
|     81.0 | 2020-09-22 |   81.0+build2-0 (USN-4546-1) | 2020-09-28 |  6 days |
|   80.0.1 | 2020-09-01 | 80.0.1+build1-0 (USN-4474-2) | 2020-09-03 |  2 days |
|     80.0 | 2020-08-25 |   80.0+build2-0 (USN-4474-1) | 2020-08-26 |   1 day |
|     79.0 | 2020-07-28 |   79.0+build1-0 (USN-4443-1) | 2020-07-29 |   1 day |
|   78.0.2 | 2020-07-09 | 78.0.2+build2-0 (USN-4423-1) | 2020-07-14 |  5 days |
|   78.0.1 | 2020-07-01 | 78.0.1+build1-0 (USN-4408-1) | 2020-07-02 |   1 day |
|     78.0 | 2020-06-30 |                      &mdash; | &mdash;    | &mdash; |
|   77.0.1 | 2020-06-03 | 77.0.1+build1-0 (USN-4383-1) | 2020-06-04 |   1 day |
|     77.0 | 2020-06-02 |                      &mdash; | &mdash;    | &mdash; |
|   76.0.1 | 2020-05-08 | 76.0.1+build1-0 (USN-4353-2) | 2020-05-12 |  4 days |
|     76.0 | 2020-05-05 |   76.0+build2-0 (USN-4353-1) | 2020-05-07 |  2 days |
|     75.0 | 2020-04-07 |               (Ubuntu 20.04) | 2020-04-23 | &mdash; |

### ブックマーク

- [Firefox  87.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/87.0/releasenotes/)
- [Firefox  86.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/86.0.1/releasenotes/)
- [Firefox  86.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/86.0/releasenotes/)
- [Firefox  85.0.2, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/85.0.2/releasenotes/)
- [Firefox  85.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/85.0.1/releasenotes/)
- [Firefox  85.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/85.0/releasenotes/)
- [Firefox  84.0.2, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/84.0.2/releasenotes/)
- [Firefox  84.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/84.0.1/releasenotes/)
- [Firefox  84.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/84.0/releasenotes/)
- [Firefox  83.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/83.0/releasenotes/)
- [Firefox  82.0.3, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/82.0.3/releasenotes/)
- [Firefox  82.0.2, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/82.0.2/releasenotes/)
- [Firefox  82.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/82.0.1/releasenotes/)
- [Firefox  82.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/82.0/releasenotes/)
- [Firefox  81.0.2, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/81.0.2/releasenotes/)
- [Firefox  81.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/81.0.1/releasenotes/)
- [Firefox  81.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/81.0/releasenotes/)
- [Firefox  80.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/80.0.1/releasenotes/)
- [Firefox  80.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/80.0/releasenotes/)
- [Firefox  79.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/79.0/releasenotes/)
- [Firefox  78.0.2, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/78.0.2/releasenotes/)
- [Firefox  78.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/78.0.1/releasenotes/)
- [Firefox  78.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/78.0/releasenotes/)
- [Firefox  77.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/77.0.1/releasenotes/)
- [Firefox  77.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/77.0/releasenotes/)
- [Firefox  76.0.1, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/76.0.1/releasenotes/)
- [Firefox  76.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/76.0/releasenotes/)
- [Firefox  75.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/75.0/releasenotes/)

[Firefox]: https://www.mozilla.org/firefox/ "Firefox - Protect your life online with privacy-first products"

## [Thunderbird]

現在は APT を使わず公式バイナリをダウンロードして手動でインストールしているが，参考情報として残しておく。

| Official | Release    |                       Ubuntu | Release    |   Delay |
| -------: | ---------- | ---------------------------: | ---------- | ------: |
|   78.9.1 | 2021-04-08 |                              |            |  ∞ days |
|   78.9.0 | 2021-03-23 |                      &mdash; | &mdash;    |  &mdash |
|   78.8.1 | 2021-03-08 |                      &mdash; | &mdash;    |  &mdash |
|   78.8.0 | 2021-02-23 |                      &mdash; | &mdash;    | &mdash; |
|   78.7.1 | 2021-02-05 | 78.7.1+build1-0 (USN-4736-1) | 2021-02-16 | 11 days |
|   78.7.0 | 2021-01-26 |                      &mdash; | &mdash;    | &mdash; |
|   78.6.1 | 2020-01-11 | 78.6.1+build1-0 (USN-4701-1) | 2021-01-20 |  9 days |
|   78.6.0 | 2020-12-15 |                      &mdash; | &mdash;    | &mdash; |
|   78.5.1 | 2020-12-02 |                      &mdash; | &mdash;    | &mdash; |
|   78.5.0 | 2020-11-17 | 78.5.0+build3-0 (USN-4647-1) | 2020-11-25 |  8 days |
|   78.4.3 | 2020-11-01 |                      &mdash; | &mdash;    | &mdash; |
|   78.4.2 | 2020-11-09 |                      &mdash; | &mdash;    | &mdash; |
|   78.4.1 | 2020-11-06 |                      &mdash; | &mdash;    | &mdash; |
|   78.4.0 | 2020-10-21 |                      &mdash; | &mdash;    | &mdash; |
|   78.3.3 | 2020-10-16 |                      &mdash; | &mdash;    | &mdash; |
|   78.3.2 | 2020-10-07 |                      &mdash; | &mdash;    | &mdash; |
|   78.3.1 | 2020-09-26 |                      &mdash; | &mdash;    | &mdash; |
|   78.3.0 | 2020-09-24 |                      &mdash; | &mdash;    | &mdash; |
|   78.2.2 | 2020-09-10 |                      &mdash; | &mdash;    | &mdash; |
|   78.2.1 | 2020-08-29 |                      &mdash; | &mdash;    | &mdash; |
|   78.2.0 | 2020-08-25 |                      &mdash; | &mdash;    | &mdash; |
|   78.1.1 | 2020-08-06 |                      &mdash; | &mdash;    | &mdash; |
|   78.1.0 | 2020-07-30 |                      &mdash; | &mdash;    | &mdash; |
|   78.0.1 | 2020-07-21 |                      &mdash; | &mdash;    | &mdash; |
|     78.0 | 2020-07-17 |                      &mdash; | &mdash;    | &mdash; |

### ブックマーク

- [Thunderbird — Release Notes (78.9.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.9.1/releasenotes/)
  - [「Thunderbird 78.9.1」が公開 ～「OpenPGP」の受信者エイリアスをサポート - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1317631.html)
- [Thunderbird — Release Notes (78.9.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.9.0/releasenotes/)
- [Thunderbird — Release Notes (78.8.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.8.1/releasenotes/)
- [Thunderbird — Release Notes (78.8.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.8.0/releasenotes/)
- [Thunderbird — Release Notes (78.7.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.7.1/releasenotes/)
- [Thunderbird — Release Notes (78.7.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.7.0/releasenotes/)
- [Thunderbird — Release Notes (78.6.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.6.1/releasenotes/)
- [Thunderbird — Release Notes (78.6.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.6.0/releasenotes/)
- [Thunderbird — Release Notes (78.5.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.5.1/releasenotes/)
- [Thunderbird — Release Notes (78.5.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.5.0/releasenotes/)
- [Thunderbird — Release Notes (78.4.3) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.4.3/releasenotes/)
- [Thunderbird — Release Notes (78.4.2) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.4.2/releasenotes/)
- [Thunderbird — Release Notes (78.4.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.4.1/releasenotes/)
- [Thunderbird — Release Notes (78.4.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.4.0/releasenotes/)
- [Thunderbird — Release Notes (78.3.3) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.3.3/releasenotes/)
- [Thunderbird — Release Notes (78.3.2) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.3.2/releasenotes/)
- [Thunderbird — Release Notes (78.3.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.3.1/releasenotes/)
- [Thunderbird — Release Notes (78.3.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.3.0/releasenotes/)
- [Thunderbird — Release Notes (78.2.2) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.2.2/releasenotes/)
- [Thunderbird — Release Notes (78.2.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.2.1/releasenotes/)
- [Thunderbird — Release Notes (78.2.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.2.0/releasenotes/)
- [Thunderbird — Release Notes (78.1.1) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.1.1/releasenotes/)
- [Thunderbird — Release Notes (78.1.0) — Thunderbird](https://www.thunderbird.net/en-US/thunderbird/78.1.0/releasenotes/)

- [結局 Thunderbird もインストールし直すことにした]({{< ref "/remark/2019/11/reinstalling-thunderbird.md" >}})

[Thunderbird]: https://www.thunderbird.net/ "Thunderbird — There’s nothing here! — Thunderbird"
