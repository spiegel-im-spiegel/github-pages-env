+++
title = "脆弱性情報の収集・管理ツール jvnman v0.2.0 をリリース，他"
date = "2018-05-15T21:53:54+09:00"
description = "脆弱性情報の収集・管理ツール jvnman の v0.2.0 をリリースした。 また jvnman に関連するパッケージもリリースした。"
image = "/images/attention/tools.png"
tags  = [ "tools", "security", "vulnerability", "risk", "management", "jvn", "cvss" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

脆弱性情報の収集・管理ツール [jvnman] の v0.2.0 をリリースした。
また [jvnman] に関連するパッケージもリリースした。

- [Release v0.2.0 · spiegel-im-spiegel/jvnman · GitHub](https://github.com/spiegel-im-spiegel/jvnman/releases/tag/v0.2.0)
- [Release v0.4.0 · spiegel-im-spiegel/go-myjvn · GitHub](https://github.com/spiegel-im-spiegel/go-myjvn/releases/tag/v0.4.0)
- [Release v0.1.0 · spiegel-im-spiegel/go-cvss · GitHub](https://github.com/spiegel-im-spiegel/go-cvss/releases/tag/v0.1.0)

[spiegel-im-spiegel/go-cvss] は [spiegel-im-spiegel/go-myjvn] 内のサブパッケージを独立させたもので今回が初リリース。
これでようやく CVSS をまともにハンドリングできるようになった。

たとえば [JVNDB-2018-003148](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003148.html "JVNDB-2018-003148 - JVN iPedia - 脆弱性対策情報データベース") に関する情報がほしければ

```text
$ jvnman info JVNDB-2018-003148
```

とすれば標準出力に

```markdown
# Mercurial におけるアクセス制御に関する脆弱性

脆弱性対策情報ID: [JVNDB-2018-003148](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003148.html)

Mercurial には、アクセス制御に関する脆弱性が存在します。

## 想定される影響

情報を取得される、および情報を改ざんされる可能性があります。

### 影響を受ける製品

- Debian / Debian GNU/Linux 
- Mercurial / Mercurial 4.5 およびそれ以前


### 深刻度

緊急: CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:N（9.1）

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | なし |


## 対策

ベンダより正式な対策が公開されています。ベンダ情報を参照して適切な対策を実施してください。

## 関連情報

- Common Vulnerabilities and Exposures (CVE) [CVE-2018-1000132](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-1000132) 
- Debian [[SECURITY] [DLA 1331-1] mercurial security update](https://lists.debian.org/debian-lts-announce/2018/03/msg00034.html) 
- JVNDB [CWE-284](https://cwe.mitre.org/data/definitions/284.html) 不適切なアクセス制御
- National Vulnerability Database (NVD) [CVE-2018-1000132](https://nvd.nist.gov/vuln/detail/CVE-2018-1000132) 
- Release Note [Mercurial 4.5.1 / 4.5.2 (2018-03-06)](https://www.mercurial-scm.org/wiki/WhatsNew#Mercurial_4.5.1_.2F_4.5.2_.282018-03-06.29) 


## 更新情報

- 発見日 2018年3月6日
- 公開日 2018年5月15日
- 最終更新日 2018年5月15日

(Powerd by [JVN](https://jvn.jp/))
```

と出力される。
CVSS のベクタもちゃんと表に展開されるぞ。

よーし，うむうむ，よーし。

とりあえず，これで欲しい情報は取れるようになったので，あとはのんびりテストをするか。
つか， [gpgpdump](https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer") のほうも作業を進めないといけないのだが...

[jvnman]: https://github.com/spiegel-im-spiegel/jvnman "spiegel-im-spiegel/jvnman: JVN Vulnerability Data Management"
[spiegel-im-spiegel/go-myjvn]: https://github.com/spiegel-im-spiegel/go-myjvn "spiegel-im-spiegel/go-myjvn: Handling MyJVN RESTful API by Golang"
[spiegel-im-spiegel/go-cvss]: https://github.com/spiegel-im-spiegel/go-cvss "spiegel-im-spiegel/go-cvss: Common Vulnerability Scoring System (CVSS) Version 3"
