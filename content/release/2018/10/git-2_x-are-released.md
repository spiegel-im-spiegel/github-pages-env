+++
title = "【セキュリティ・アップデート】 Git 2.x 最新版がリリース"
date = "2018-10-10T16:49:56+09:00"
description = "サブモジュールに関する脆弱性で，再帰的な git clone において任意のコードが実行される可能性がある。深刻度高めなのでご注意を。"
image = "/images/attention/tools.png"
tags  = [ "tools", "git", "security", "vulnerability" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

[git] 2.x の各マイナーバージョンに脆弱性が見つかったため修正版がリリースされている。

- ['[Announce] Git 2.14.5, 2.15.3, 2.16.5, 2.17.2, 2.18.1, and 2.19.1' - MARC](https://marc.info/?l=git&m=153875888916397&w=2)
- [Remediating the October 2018 Git Security Vulnerability – Microsoft DevOps Blog](https://blogs.msdn.microsoft.com/devops/2018/10/05/remediating-the-october-2018-git-security-vulnerability/)
- [NVD - CVE-2018-17456](https://nvd.nist.gov/vuln/detail/CVE-2018-17456)

サブモジュールに関する脆弱性で，再帰的な `git clone` において任意のコードが実行される可能性がある。

{{< fig-quote title="[Announce] Git 2.14.5, 2.15.3, 2.16.5, 2.17.2, 2.18.1, and 2.19.1" link="https://marc.info/?l=git&m=153875888916397&w=2" lang="en" >}}
<q>When running "git clone --recurse-submodules", Git parses the supplied .gitmodules file for a URL field and blindly passes it as an argument to a "git clone" subprocess.  If the URL field is set to a string that begins with a dash, this "git clone" subprocess interprets the URL as an option.  This can lead to executing an arbitrary script shipped in the superproject as the user who ran "git clone".</q>
{{< /fig-quote >}}

サブモジュールの脆弱性は前にもなかったっけ。
まぁ，この手の脆弱性はありがちではある。
TeX とか Go 言語の cgo とかでも似たようなことがあったよなぁ。

深刻度高めみたいなのでご注意を。

- [CVE-2018-17456 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2018-17456)

**深刻度: 重要 (8.8)** : `CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:U/C:H/I:H/A:H`

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 要           |
| スコープ         | 変更なし     |
| 機密性への影響   | 高           |
| 完全性への影響   | 高           |
| 可用性への影響   | 高           |


各マイナーバージョンの最新版がリリースされているので更新すること。

なお [Git for Windows] には影響がないとのこと。

{{< fig-quote title="Remediating the October 2018 Git Security Vulnerability" link="https://blogs.msdn.microsoft.com/devops/2018/10/05/remediating-the-october-2018-git-security-vulnerability/" lang="en" >}}
<q>Git for Windows is uniquely not vulnerable to this security issue: this vulnerability requires writing a file to disk, and that filename must be particularly formatted and include a colon. Since colons are not permitted characters on Windows filesystems, Git for Windows will refuse to write the file.</q>
{{< /fig-quote >}}

このせいなのか [Git for Windows] にアナウンスがなくて気が付かなかったんだよねぇ（笑）

アップデートは計画的に。

## ブックマーク

- [バージョン管理システム「Git」に任意コード実行の脆弱性、修正版が公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1146/869/index.html)

[git]: https://git-scm.com/
[Git for Windows]: https://gitforwindows.org/
