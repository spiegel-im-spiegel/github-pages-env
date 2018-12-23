+++
date = "2015-12-02T19:18:01+09:00"
description = "当面（2017年度末まで）は CVSSv2 と CVSSv3 の併記で運用するらしい。"
draft = false
tags = ["security", "vulnerability", "risk", "management", "cvss", "ipa", "jpcert"]
title = "JVN が CVSSv3 による脆弱性評価を開始"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

## JVN が CVSSv3 による脆弱性評価を開始

- [共通脆弱性評価システムCVSS v3 (新バージョン)での評価の開始について：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/vuln/SeverityLevel3.html)
- [JVN が共通脆弱性評価システム CVSS v3 による脆弱性評価を開始](https://www.jpcert.or.jp/press/2015/20151201-CVSSv3.html)

当面（2017年度末まで）は CVSSv2 と CVSSv3 の併記で運用するらしい。

たとえば

- [JVNVU#97647301: RSI Video Technologies の Videofied Frontel がセキュアでない独自プロトコルを使用する問題](http://jvn.jp/vu/JVNVU97647301/)

では

**CVSSv2 基本評価値 5.0 (`AV:N/AC:L/Au:N/C:P/I:N/A:N`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 低 (L)            |
| 攻撃前の認証要否（Au）                  | 不要（N）         |
| 情報漏えいの可能性（機密性への影響, C） | 部分的（P）       |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）         |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |


**CVSSv3 基本評価値 7.5 (`CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 低 (L)            |
| 必要な特権レベル（PR）                  | 不要（N）         |
| ユーザ関与レベル（UI）                  | 不要（N）         |
| スコープ（S）                           | 変更なし (U)      |
| 情報漏えいの可能性（機密性への影響, C） | 高（H）           |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）         |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |

という感じで CVSSv2 と CVSSv3 では評価の観点も最終的な評価点も異なることがわかると思う[^a]。

[^a]: CVSSv3 では基本評価値が 7.0 以上で「重要」レベル，9.0 以上で「緊急」レベルとしている。基本評価値が 7.0 以上あるならすぐに対策に向けたアクションを起こすべきだろう。

余談だが CERT/CC では

- [Vulnerability Note VU#792004 - RSI Video Technologies Videofied security system Frontel software uses an insecure custom protocol](http://www.kb.cert.org/vuls/id/792004)

**CVSSv2 基本評価値 9.4 (`AV:N/AC:L/Au:N/C:C/I:C/A:N`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 低 (L)            |
| 攻撃前の認証要否（Au）                  | 不要（N）         |
| 情報漏えいの可能性（機密性への影響, C） | 全面的（C）       |
| 情報改ざんの可能性（完全性への影響, I） | 全面的（C）       |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |

と完全性への影響が高いと評価している[^b]。

[^b]: 仮に完全性への影響が高いと考えると先ほどの CVSSv3 は `CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:N` で，評価は 9.1 となる。なお，この脆弱性についてはベンダ企業から修正版がリリースされている。

## CVSSv3 の変更点

CVSSv2 から CVSSv3 への変更点は以下の通り。
（なおこれは，以前書いた「[CVSS に関するメモ 3](http://www.baldanders.info/spiegel/log2/000864.shtml)」の再掲載である）

バージョン2からの大きな違いは深刻度をコンポーネント単位で評価できるようになったことだろう。
以前は攻撃対象となるシステムやホストマシンが対象だったので，更に細かい評価ができるようになったと言える。

これに合わせて「スコープ（Scope）」という評価軸が追加された。
厳密には「管理権限の範囲（Authorization Scope）」。
脆弱性のあるコンポーネントが他のコンポーネントに影響をおよぼす場合に，「他のコンポーネント」が管理権限の範囲の内側か外側かで深刻度が異なる。
脆弱性が管理権限の範囲の外側に影響を及ぼす具体例としてはクロスサイトスクリプティング（XSS）脆弱性などが挙げられるだろう。

基本評価基準（Base Metrics）で新たに追加された評価項目としては


- 必要な特権レベル（Privileges Required）
- ユーザ関与レベル（User Interaction）
- スコープ（Scope）

がある。
一方，バージョン2にあった「攻撃前の認証要否（Authentication）」項目は廃止され，「必要な特権レベル」に含まれる形になった。

一番大きく変わったのは環境評価基準（Environmental Metrics）だろう。
環境評価基準では対策後の状況を再評価し，評価が低いものについては再度対策を行えるようになっている。
CVSS を使ってセキュリティ対策のプロセスをきちんと回せるようになったわけだ。

再評価の観点は以下のとおり。


- 緩和策後の攻撃元区分（Modified Attack Vector）
- 緩和策後の攻撃条件の複雑さ（Modified Attack Complexity）
- 緩和策後の必要な特権レベル（Modified Privileges Required）
- 緩和策後のユーザ関与レベル（Modified User Interaction）
- 緩和策後のスコープ（Modified Scope）
- 緩和策後の機密性への影響（Modified Confidentiality Impact）
- 緩和策後の完全性への影響（Modified Integrity Impact）
- 緩和策後の可用性への影響（Modified Availability Impact）

つか基本評価基準（Base Metrics）の評価観点で再評価するって感じかな。

最近の「セキュリティ対策」がこれまでと異なっているのは，目の前の脆弱性に対処すれば「完了」とはならない点である。
（私は今までもずうっと言ってきているけど）セキュリティは「ハザード（hazard）」ではなく「リスク（risk）」で考えなければならない。
リスクをゼロにするには無限のリソース（人的資源やもっと端的にお金）が必要だが，私たちの持っているリソースは無限ではない。
つまりリスクをゼロにすることはできないのだ。
したがって，業務プロセスの中で改善を行いながら，最適解を探っていくしかない。

「セキュリティ対策」をコストと考えるなら果てしなく不毛な話であるが，「セキュリティ対策」を投資と考え，きちんと PDCA サイクルを回していくならば，それほど不毛ではないはずである。

## 関連リンク

- [CVSS v3.0 Specification Document](https://www.first.org/cvss/specification-document)
- [共通脆弱性評価システムCVSS v3概説：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/vuln/CVSSv3.html)
- [CVSSv3 用の node.js モジュールを作ってみた - Qiita](http://qiita.com/spiegel-im-spiegel/items/d6fe10d3df92b9d8556b) : 記事では「基本評価基準のみ実装」と書いているが，一応フルセット実装している
    - [spiegel-im-spiegel/cvss3](https://github.com/spiegel-im-spiegel/cvss3) : node.js 0.12 で動作確認。 4.x や 5.x で動作するかどうかは分からない
- [node.js の CVSS v3 モジュールを使ってデモページを作ってみた - Qiita](http://qiita.com/spiegel-im-spiegel/items/f2db3759b957206d4521)
    - [Demo for CVSS v3](http://www.baldanders.info/spiegel/archive/cvss/cvss3.html)
