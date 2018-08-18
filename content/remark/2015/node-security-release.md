+++
date = "2015-12-05T21:51:15+09:00"
update = "2016-12-26T21:07:45+09:00"
description = "前々から予告があったが node.js に関する複数の脆弱性に対するアップデートが行われた。"
draft = false
tags = ["security", "vulnerability", "nodejs"]
title = "Node.js : Security Update"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

前々から予告があったが node.js に関する複数の脆弱性に対するアップデートが行われた。

- [December Security Release Summary | Node.js](https://nodejs.org/en/blog/vulnerability/december-2015-security-releases/)
- [Node v0.10.41 (Maintenance) | Node.js](https://nodejs.org/en/blog/release/v0.10.41/)
- [Node v0.12.9 (LTS) | Node.js](https://nodejs.org/en/blog/release/v0.12.9/)
- [Node v4.2.3 "Argon" (LTS) | Node.js](https://nodejs.org/en/blog/release/v4.2.3/)
- [Node v5.1.1 (Stable) | Node.js](https://nodejs.org/en/blog/release/v5.1.1/)

なお， CVSS については[解説ページ]({{< ref "/remark/2015/cvss-v3-metrics-in-jvn.md" >}})を参照のこと。

## CVE-2015-8027 Denial of Service Vulnerability

**CVSSv3 基本評価値 7.5 (`CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 低（L）           |
| 必要な特権レベル（PR）                  | 不要（N）         |
| ユーザ関与レベル（UI）                  | 不要（N）         |
| スコープ（S）                           | 変更なし (U)      |
| 情報漏えいの可能性（機密性への影響, C） | なし（N）         |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）         |
| 業務停止の可能性（可用性への影響, A）   | 高（H）           |

## CVE-2015-6764 V8 Out-of-bounds Access Vulnerability

**CVSSv3 基本評価値 4.4 (`CVSS:3.0/AV:N/AC:H/PR:H/UI:N/S:U/C:N/I:N/A:H`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 高（H）           |
| 必要な特権レベル（PR）                  | 高（H）           |
| ユーザ関与レベル（UI）                  | 不要（N）         |
| スコープ（S）                           | 変更なし (U)      |
| 情報漏えいの可能性（機密性への影響, C） | なし（N）         |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）         |
| 業務停止の可能性（可用性への影響, A）   | 高（H）           |

## CVE-2015-3193 OpenSSL BN_mod_exp may produce incorrect results on x86_64

**CVSSv2 基本評価値 2.6 (`AV:N/AC:H/Au:N/C:P/I:N/A:N`)**
（[Redhat による評価](https://access.redhat.com/security/cve/cve-2015-3193)）

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 高（H）           |
| 攻撃前の認証要否（Au）                  | 不要（N）         |
| 情報漏えいの可能性（機密性への影響, C） | 部分的（P）       |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）         |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |

## CVE-2015-3194 OpenSSL Certificate verify crash with missing PSS parameter

**CVSSv2 基本評価値 4.3 (`AV:N/AC:M/Au:N/C:N/I:N/A:P`)**
（[Redhat による評価](https://access.redhat.com/security/cve/cve-2015-3194)）

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 中（M）           |
| 攻撃前の認証要否（Au）                  | 不要（N）         |
| 情報漏えいの可能性（機密性への影響, C） | なし（N）         |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）         |
| 業務停止の可能性（可用性への影響, A）   | 部分的（P）       |

なお， OpenSSL に関しては他にも複数の脆弱性が報告されている。

- [OpenSSL Security Advisory [3 Dec 2015] - Updated [4 Dec 2015]](http://openssl.org/news/secadv/20151203.txt)
- [OpenSSLに4つの脆弱性が存在 - US-CERT | マイナビニュース](http://news.mynavi.jp/news/2015/12/05/124/)

## 影響のあるバージョン

| node.js バージョン | CVE-2015-8027 | CVE-2015-6764 | CVE-2015-3193 | CVE-2015-3193 |
|-------------------:|:-------------:|:-------------:|:-------------:|:-------------:|
| Node v0.10.41 未満 | ✕            | ✕            | ✕            | ◯            |
| Node v0.12.9 未満  | ◯            | ✕            | ✕            | ◯            |
| Node v4.2.3 未満   | ◯            | ◯            | ◯            | ◯            |
| Node v5.1.1 未満   | ◯            | ◯            | ◯            | ◯            |

この機会に v5.x に上げようか...
