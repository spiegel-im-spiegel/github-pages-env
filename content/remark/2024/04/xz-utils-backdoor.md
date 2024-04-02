+++
title = "XZ Utils に仕組まれたバックドアに関する覚え書き"
date =  "2024-04-01T20:18:07+09:00"
description = "起こったことはしょうがないので，ユーザとしては粛々と対応するしかない。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "vulnerability", "malware", "linux" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日から騒ぎになっている XZ Utils に仕組まれたバックドアについて覚え書きとしてブックマークをまとめておく。
最初は単なるセキュリティ脆弱性か悪意のコードかみたいな話が飛び交ってた気がするが，完全に malicious code として認知されたみたいね。
まぁ，詐欺としてはよくある話。
悪人顔の詐欺師なんかフィクションである（笑）

起こったことはしょうがないので，ユーザとしては粛々と対応するしかない。
政治的な話は知らんが，願わくばまっとうなエンジニアを非難し活動を萎縮させるような行為は止めていただきたいものである。

- [XZ Utils backdoor](https://tukaani.org/xz-backdoor/)
- [NVD - CVE-2024-3094](https://nvd.nist.gov/vuln/detail/CVE-2024-3094)
  - `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:H`
  - 深刻度: 緊急 (Score: 10)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更あり |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |

ただの圧縮ソフトだろ，と思うかもしれないが

{{< fig-quote type="markdown" title="XZ Utilsの脆弱性 CVE-2024-3094 についてまとめてみた - piyolog" link="https://piyolog.hatenadiary.jp/entry/2024/04/01/035321" >}}
影響を受けたライブラリをリンクしているsshdを使用していた場合、liblzmaの実装が改ざんされ、任意のソフトウエアを利用できるようになり、結果として外部から不正アクセスをされる恐れがある
{{< /fig-quote >}}

ということで，近年話題になっているサプライチェーン攻撃としても結構ヤバい話である。
今のところ（2024-04-01 時点），対応としては影響を受けないバージョン（5.4.6など）にダウングレードするしかない。

ただし，影響を受ける Linux ディストリビューションのバージョンが少ないのは不幸中の幸いだろう。
Ubuntu は，自前で最新版を入れていない限り，現在リリースされているバージョンではいずれも[影響がない](https://ubuntu.com/security/CVE-2024-3094 "CVE-2024-3094 | Ubuntu")ようだ。

詳しい情報は以下が参考になる。

- [XZ Utilsに悪意のあるコードが挿入された問題（CVE-2024-3094）について](https://www.jpcert.or.jp/newsflash/2024040101.html)
- [XZ Utilsの脆弱性 CVE-2024-3094 についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2024/04/01/035321)

その他，参考情報：

- [Reported Supply Chain Compromise Affecting XZ Utils Data Compression Library, CVE-2024-3094 | CISA](https://www.cisa.gov/news-events/alerts/2024/03/29/reported-supply-chain-compromise-affecting-xz-utils-data-compression-library-cve-2024-3094)
- [Urgent security alert for Fedora 41 and Fedora Rawhide users](https://www.redhat.com/en/blog/urgent-security-alert-fedora-41-and-rawhide-users) : Red Hat
- [[SECURITY] [DSA 5649-1] xz-utils security update](https://lists.debian.org/debian-security-announce/2024/msg00057.html) : Debian
- [xz-utils backdoor situation (CVE-2024-3094) · GitHub](https://gist.github.com/thesamesam/223949d5a074ebc3dce9ee78baad9e27)
- [amlweems/xzbot: notes, honeypot, and exploit demo for the xz backdoor (CVE-2024-3094)](https://github.com/amlweems/xzbot)
- [Xz Utils Backdoor - Schneier on Security](https://www.schneier.com/blog/archives/2024/04/xz-utils-backdoor.html)
- [広く使用されている「xz」にssh接続を突破するバックドアが仕込まれていた事が判明。重大度はクリティカルでLinuxのほかmacOSにも影響 | ソフトアンテナ](https://softantenna.com/blog/xz-backdoor/)

## ブックマーク

- [オープンソース製品とソフトウェア部品表]({{< ref "/remark/2022/08/software-bills-of-materials.md" >}})

## 参考図書

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "4296001574" %}} <!-- ハッキング思考 -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
