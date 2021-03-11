+++
title = "OpenSSH 8.5 のリリース【セキュリティ・アップデート】"
date =  "2021-03-11T19:11:29+09:00"
description = "アップデートは計画的に"
image = "/images/attention/tools.png"
tags  = [ "openssh", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先週の話で恐縮だが， 2021-03-03 にリリースされた [OpenSSH] 8.5 で [CVE-2021-28041] の脆弱性の改修がされていたらしい。

{{< fig-quote type="markdown" title="NVD - CVE-2021-28041" link="https://nvd.nist.gov/vuln/detail/CVE-2021-28041" lang="en" >}}
{{% quote %}}ssh-agent in OpenSSH before 8.5 has a double free that may be relevant in a few less-common scenarios, such as unconstrained agent-socket access on a legacy operating system, or the forwarding of an agent to an attacker-controlled host{{% /quote %}}.
{{< /fig-quote >}}

ssh-agent のソケット周りとかヤバいぢゃん。

CVSSv3 評価は以下の通り。

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H`
- 深刻度: 緊急 (Score: 9.8)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |

実は [Ubuntu] で，これを修正したらしい [OpenSSH] 8.3 バックポート・パッチが出てたので気付いた次第。

- [USN-4762-1: OpenSSH vulnerability | Ubuntu security notices | Ubuntu](https://ubuntu.com/security/notices/USN-4762-1)

まぁ，アップデートは計画的に，ということで。

[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-28041]: https://nvd.nist.gov/vuln/detail/CVE-2021-28041

## 参考図書

{{% review-paapi "B079NL1L9K" %}} <!-- SSH Mastery -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
