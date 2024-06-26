+++
title = "Git v2.26.2 のリリース【セキュリティ・アップデート】"
date =  "2020-04-21T09:12:19+09:00"
description = " 前回の CVE-2020-5260 が直りきってなかったってことでいいのかな？"
image = "/images/attention/tools.png"
tags  = [ "git", "tools", "security", "vulnerability", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Git] v2.26.2 を含む複数のバージョンがリリースされた。

- [[Announce] Git v2.26.2 and others](https://lore.kernel.org/git/xmqq4kterq5s.fsf@gitster.c.googlers.com/T/)

対象となるのは 2.17.x から 2.26.x までの各マイナーバージョン。
前回の [CVE-2020-5260] が直りきってなかったってことでいいのかな？

{{< fig-quote type="markdown" title="Git v2.26.2 and others" link="https://lore.kernel.org/git/xmqq4kterq5s.fsf@gitster.c.googlers.com/T/" lang="en" >}}
{{< quote >}}These releases address the security issue CVE-2020-11008, which is similar to the recently addressed CVE-2020-5260{{< /quote >}}.
{{< /fig-quote >}}

[CVE-2020-5260] と同様ということは深刻度も「緊急」ちうことかな。
[CVE-2020-5260] については，[前のリリース記事]({{< relref "./git-2_26_1-is-released.md" >}} "Git v2.26.1 のリリース【セキュリティ・アップデート】")を参考にどうぞ。

## [CVE-2020-11008](https://nvd.nist.gov/vuln/detail/CVE-2020-11008)

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N`
- 深刻度: 重要 (7.5)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更なし     |
| 機密性への影響   | 高           |
| 完全性への影響   | なし         |
| 可用性への影響   | なし         |

一方 GitHub は

- `CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:C/C:L/I:N/A:N`
- 深刻度: 警告 (4.0)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 高           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更あり     |
| 機密性への影響   | 低           |
| 完全性への影響   | なし         |
| 可用性への影響   | なし         |

と見積もっている。

## アップデートは...

[Ubuntu] であれば [PPA 版リポジトリ](https://launchpad.net/~git-core/+archive/ubuntu/ppa "Git stable releases : “Ubuntu Git Maintainers” team")を利用することをおすすめする。

- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})

アップデートは計画的に。

## ブックマーク

- [Release Git for Windows 2.26.2 · git-for-windows/git · GitHub](https://github.com/git-for-windows/git/releases/tag/v2.26.2.windows.1)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[git]: https://git-scm.com/
[Git]: https://git-scm.com/
[Git for Windows]: https://gitforwindows.org/
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
[CVE-2020-5260]: https://nvd.nist.gov/vuln/detail/CVE-2020-5260
