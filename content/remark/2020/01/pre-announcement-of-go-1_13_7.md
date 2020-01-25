+++
title = "Go 1.13.7 リリース予告と CVE-2020-0601"
date =  "2020-01-25T16:46:13+09:00"
description = "来週1月28日（日本時間では1月29日かな）に Go 1.13.7 のリリースがあるようだ（予定は未定）。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "security", "vulnerability", "risk", "windows", "cryptography", "ecc", "certification" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

来週1月28日（日本時間では1月29日かな）に [Go] 1.13.7 のリリースがあるようだ（予定は未定）。

- [[security] Go 1.12.16 and Go 1.13.7 pre-announcement - Google group](https://groups.google.com/forum/#!topic/golang-announce/-sdUB4VEQkA)

[Go] 1.13.7 にはセキュリティ・アップデートが含まれる。

{{< fig-quote type="markdown" title="Go 1.12.16 and Go 1.13.7 pre-announcement" link="https://groups.google.com/forum/#!topic/golang-announce/-sdUB4VEQkA" lang="en" >}}
{{< quote >}}These are minor releases that include two security fixes.
One [mitigates the CVE-2020-0601 certificate verification bypass](https://golang.org/cl/215905) on Windows.
The other affects only 32-bit architectures{{< /quote >}}.
{{< /fig-quote >}}

[CVE-2020-0601](https://nvd.nist.gov/vuln/detail/CVE-2020-0601) は Windows CryptoAPI の不備で楕円曲線暗号（ECC）を使った証明書の検証がバイパスされてしまうというもの。
これによって証明書の偽装ができてしまう。
あとは分かるね。

- [CVE-2020-0601 | Windows CryptoAPI Spoofing Vulnerability](https://portal.msrc.microsoft.com/en-US/security-guidance/advisory/CVE-2020-0601)

CVSSv3 の評価は以下の通り。

{{< fig-gen type="markdown" title="CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:U/C:H/I:H/A:N" >}}
| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | なし |
{{< /fig-gen >}}


深刻度は 8.1 (重要)[^cvss1]。

[^cvss1]: CVSSv3 では深刻度7.0以上なら速やかな対応が求められる。アップデートは計画的に（笑）

ただし，この脆弱性は1月の Windows Update で修正されている筈である。
つか，私はもう関係ないので完全にスルーしていた（笑）

- [January 2020 Security Updates: CVE-2020-0601 - Microsoft Security Response Center](https://msrc-blog.microsoft.com/2020/01/14/january-2020-security-updates-cve-2020-0601/)

今回の [CVE-2020-0601] の特徴は NSA が絡んでいる点である。

- {{< pdf-file title="Patch Critical Cryptographic Vulnerability in Microsoft Windows Clients and Servers" link="https://media.defense.gov/2020/Jan/14/2002234275/-1/-1/0/CSA-WINDOWS-10-CRYPT-LIB-20190114.PDF" >}}

NSA の思惑は分からない。
しかし

{{< fig-quote type="markdown" title="Critical Windows Vulnerability Discovered by NSA - Schneier on Security" link="https://www.schneier.com/blog/archives/2020/01/critical_window.html" lang="en" >}}
{{< quote >}}She did not answer when asked how long ago the NSA discovered the vulnerability. She said that this is not the first time the NSA sent Microsoft a vulnerability to fix, but it was the first time it has publicly taken credit for the discovery{{< /quote >}}.
{{< /fig-quote >}}

とある通り（どこぞの馬鹿メディアが言ってたみたいな）別に NSA の「お手柄」でもなんでもなく，政治的な思惑があって色々と天秤にかけた結果「公表」したということだろう。
まっ，国家の諜報機関なのだから当たり前だが（笑）

というわけで，自前で防衛手段を持つというのは悪い話ではない。

## ブックマーク

- [Critical Windows Vulnerability Discovered by NSA - Schneier on Security](https://www.schneier.com/blog/archives/2020/01/critical_window.html)
- [Windowsの暗号化機能に致命的な脆弱性、証明書偽装の恐れ ～米国家安全保障局が警告 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1229173.html)

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[CVE-2020-0601]: https://nvd.nist.gov/vuln/detail/CVE-2020-0601

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015SAFU42" %}} <!-- イミテーション・ゲーム -->
