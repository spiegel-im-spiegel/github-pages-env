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

## GnuPG & Libgcrypt

[Ubuntu] 19.04 リリース時の 2.2.12 (libgcrypt 1.8.4) から動いていない...
ちなみに [Ubuntu] 19.04 リリース時の公式版最新バージョンは 2.2.15 である。

| Official Version | Release Date | Update by Ubuntu | Release Date |   delay |
| ----------------:| ------------ | ----------------:| ------------ | -------:|
|           2.2.17 | 2019-07-09   |                  |              | ∞ days |

### ブックマーク

- [GnuPG 2.2.17 リリース： 公開鍵サーバ・アクセスに関する過激な変更あり]({{< ref "/release/2019/07/gnupg-2_2_17-is-released.md" >}})

- [Gnupg-announce Info Page](https://lists.gnupg.org/mailman/listinfo/gnupg-announce)

## Firefox

| Official Version | Release Date |             Update by Ubuntu | Release Date |  delay |
| ----------------:| ------------ | ----------------------------:| ------------ | ------:|
|           68.0.1 | 2019-07-18   | 68.0.1+build1-0 (USN-4054-1) | 2019-07-25   | 7 days |
|             68.0 | 2019-07-09   |   68.0+build3-0 (USN-4054-1) | 2019-07-12   | 3 days |
|           67.0.4 | 2019-06-20   | 67.0.4+build1-0 (USN-4032-1) | 2019-06-24   | 4 days |
|           67.0.3 | 2019-06-18   | 67.0.3+build1-0 (USN-4020-1) | 2019-07-20   | 2 days |

### ブックマーク

- [Mozilla、「Firefox 68.0.1」を公開 ～不具合の修正を中心としたメンテナンス更新 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1196976.html)
- [「Firefox 68」が正式公開 ～アドオン管理を改善、新しい拡張機能との出会いの場に - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1195193.html)
- [Firefox の脆弱性 (CVE-2019-11707) に関する注意喚起](https://www.jpcert.or.jp/at/2019/at190027.html)

## Thunderbird

| Official Version | Release Date |             Update by Ubuntu | Release Date |   delay |
| ----------------:| ------------ | ----------------------------:| ------------ | -------:|
|           60.8.0 | 2019-07-09   | 60.8.0+build1-0 (USN-4064-1) | 2019-07-17   |  8 days |
|           60.7.2 | 2019-06-20   | 60.7.2+build2-0 (USN-4045-1) | 2019-07-01   | 11 days |

### ブックマーク

- [フリーのメールソフト「Thunderbird」v60.8.0が正式公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1195473.html)
- [「Thunderbird」v60.7.2が公開 ～「Firefox」で悪用されているゼロデイ脆弱性に対処 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1191820.html)

## LibreOffice

現在は APT を使わず公式バイナリをダウンロードして手動でインストールしているが，参考情報として残しておく。

| Official Version | Release Date |     Update by Ubuntu | Release Date |  delay |
| ----------------:| ------------ | --------------------:| ------------ | ------:|
|            6.2.5 | 2019-07-04   | 6.2.5-0 (USN-4063-1) | 2019-07-12   | 8 days |

### ブックマーク

- [The Document Foundation announces LibreOffice 6.2.5 - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2019/07/04/tdf-announces-libreoffice-625/)

## Java

現在は APT を使わず公式バイナリをダウンロードして手動でインストールしているが，参考情報として残しておく。

| Official Version | Release Date | Update by Ubuntu | Release Date |   delay |
| ----------------:| ------------ | ----------------:| ------------ | -------:|
|           12.0.2 | 2019-07-16   |                  |              | ∞ days |

### ブックマーク

- [Oracle Java の脆弱性対策について(CVE-2019-7317等)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20190717-jre.html)

- [結局 OpenJDK をインストールし直すことにした]({{< ref "/remark/2019/07/reinstalling-openjdk.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
