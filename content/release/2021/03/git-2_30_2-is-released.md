+++
title = "Git v2.30.2 のリリース【セキュリテイ・アップデート】"
date =  "2021-03-10T18:41:10+09:00"
description = "アップデートは計画的に。"
image = "/images/attention/tools.png"
tags  = [ "tools", "git", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Git][git] v2.30.2 がリリースされた。

- [[ANNOUNCE] Git v2.30.2 and below for CVE-2021-21300 - Junio C Hamano](https://lore.kernel.org/git/xmqqim6019yd.fsf@gitster.c.googlers.com/)

タイトルにもあるとおり，今回は [CVE-2021-21300] の修正を含んでいる。

{{< fig-quote type="markdown" title="malicious repositories can execute remote code while cloning · Advisory · git/git · GitHub" link="https://github.com/git/git/security/advisories/GHSA-8prw-h3cq-mghm" lang="en" >}}
A specially crafted repository that contains symbolic links as well as files using a clean/smudge filter such as Git LFS, may cause just-checked out script to be executed while cloning onto a case-insensitive file system such as NTFS, HFS+ or APFS (i.e. the default file systems on Windows and macOS).

Note that clean/smudge filters have to be configured for that. Git for Windows configures Git LFS by default, and is therefore vulnerable.
{{< /fig-quote >}}

流石に[ク◯ったれな HFS+]({{< ref "/golang/unicode-normalization.md" >}} "Go 言語と Unicode 正規化") は駆逐されていると思うが， NTFS と APFS をファイルシステムとして使っているプラットフォームは要注意である。
Windows や macOS パソコンだけじゃなくて， NTFS で NAS を構成している場合も結構ヤバいんじゃないかな？

[Git for Windows] に関しては既にアップデートが出ている。

- [Release MinGit v2.29.2.windows.4 · git-for-windows/git · GitHub](https://github.com/git-for-windows/git/releases/tag/v2.29.2.windows.4)
- [Release Git for Windows 2.30.2 · git-for-windows/git · GitHub](https://github.com/git-for-windows/git/releases/tag/v2.30.2.windows.1)

他のプラットフォームでも順次アップデートしていくのがいいだろう。
今回の条件に該当するけどすぐにアップデートできない場合は，とりあえず

```text
$ git config --global core.symlinks false
```

としてシンボリック・リンクを無効にすることで回避できるらしい。

アップデートは計画的に。

## [CVE-2021-21300]

GitHub による評価：

- `CVSS:3.1/AV:N/AC:H/PR:N/UI:R/S:C/C:H/I:H/A:N`
- 深刻度: 重要 (Score: 8)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 高 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 要 |
| スコープ | 変更あり |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | なし |

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[git]: https://git-scm.com/
[Git for Windows]: https://gitforwindows.org/
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
[CVE-2021-21300]: https://nvd.nist.gov/vuln/detail/CVE-2021-21300 "CVE-2021-21300"
