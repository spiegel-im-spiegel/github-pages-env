+++
title = "SCM ツールの脆弱性"
date =  "2017-08-19T21:02:07+09:00"
description = "改修されたバージョンが公開されている。更新作業は計画的に。"
tags = [
  "security",
  "vulnerability",
  "tools",
  "git",
]

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
  flattr    = "spiegel"
+++

がっつり見落としてました。

- [Compromise On Checkout - Vulnerabilities in SCM Tools · The Recurity Lablog](http://blog.recurity-labs.com/2017-08-10/scm-vulns)
- [[ANNOUNCE] Git v2.14.1, v2.13.5, and others](https://www.mail-archive.com/linux-kernel@vger.kernel.org/msg1466490.html)
- [GitLab 9.4.4, 9.3.10, 9.2.10, 9.1.10, 9.0.13, and 8.17.8 Critical Security Release | GitLab](https://about.gitlab.com/2017/08/10/gitlab-9-dot-4-dot-4-released/)
- [Mercurial 4.3 / 4.3.1 (2017-08-10)](https://www.mercurial-scm.org/wiki/WhatsNew#Mercurial_4.3_.2F_4.3.1_.282017-08-10.29)
- [Arbitrary code execution on clients through malicious svn+ssh URLs in svn:externals and svn:sync-from-url](https://subversion.apache.org/security/CVE-2017-9800-advisory.txt)
- [「Git」「Mercurial」「Subversion」などにコマンドインジェクションの脆弱性 - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1075909.html)

## 脆弱性の内容

{{< fig-quote title="「Git」「Mercurial」「Subversion」などにコマンドインジェクションの脆弱性" link="http://forest.watch.impress.co.jp/docs/news/1075909.html" >}}
<q>同社のブログ記事によると、「Git LFS」の旧バージョンにはURLの解釈処理に問題があり、たとえば“ssh://-oProxyCommand=some-command”というURLの場合、ホスト名を“-o ProxyCommand=some-command”と解釈してしまうため、“some-command”が実行されてしまう。同様の問題は、「Git」「Mercurial」「Subversion」にも存在する。</q>
{{< /fig-quote >}}

## 影響度（CVSS）

- [CVE-2017-9800 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2017-9800)
- [CVE-2017-1000116 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2017-1000116)
- [CVE-2017-1000117 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2017-1000117)

**CVSSv3 基本評価値 6.3 (`CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:U/C:L/I:L/A:L`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 低（L）           |
| 必要な特権レベル（PR）                  | 不要（N）         |
| ユーザ関与レベル（UI）                  | 要（R）           |
| スコープ（S）                           | 変更なし（U）     |
| 情報漏えいの可能性（機密性への影響, C） | 低（L）           |
| 情報改ざんの可能性（完全性への影響, I） | 低（L）           |
| 業務停止の可能性（可用性への影響, A）   | 低（L）           |

CVSS については[解説ページ]({{< ref "/remark/2015/cvss-v3-metrics-in-jvn.md" >}})を参照のこと。

## 影響を受ける製品

- Git v2.7.6, v2.8.6, v2.9.5, v2.10.4, v2.11.3, v2.12.4, and v2.13.5
- Mercurial 4.3 以前
- Subversion
    - 1.0.0 through 1.8.18
    - 1.9.0 through 1.9.6
    - 1.10.0-alpha3
- GitLab （[CVE-2017-12426](http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-12426) を含む）
    - 7.9.0 through 8.17.7
    - 9.0.0 through 9.0.12
    - 9.1.0 through 9.1.9
    - 9.2.0 through 9.2.9
    - 9.3.0 through 9.3.9
    - 9.4.0 through 9.4.3

## 対策・回避策

改修されたバージョンが公開されている。
更新作業は計画的に。

- Git v2.14.1
- Mercurial 4.3.1
- Subversion 1.8.19 and 1.9.7
- GitLab 8.17.8, 9.0.13, 9.1.10, 9.2.10, 9.3.10, and 9.4.4

## ブックマーク

- [CVE-2017-1000117 対策のメモ - Qiita](http://qiita.com/bells17/items/b8a21b1ef8d9ec36a151)
