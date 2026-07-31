+++
date = "2017-02-01T23:06:06+09:00"
title = "NIST SP800-52 Rev.1 の邦訳が登場"
tags = ["security", "cryptography", "tls", "nist"]
description = "NIST SP800-52 Rev.1 の邦訳が登場したようだ。"

[scripts]
  mathjax = false
  mermaidjs = false
+++

IPA は [NIST のセキュリティ関連文書の邦訳を積極的に行っている](http://www.ipa.go.jp/security/publications/nist/ "セキュリティ関連NIST文書：IPA 独立行政法人 情報処理推進機構")が，SP800-52 Rev.1 の邦訳が登場したようだ。

- [Guidelines for the Selection, Configuration, and Use of Transport Layer Security (TLS) Implementations | NIST](https://www.nist.gov/node/562891?pub_id=915295)
- {{< pdf-file title="NIST Special Publication 800-52 Revision 1 トランスポート層セキュリティ (TLS) 実装の選択、設定、および使用のためのガイドライン" link="http://www.ipa.go.jp/files/000057084.pdf" >}}

古い話になるが，2013年までに [RC4 の危殆化](https://baldanders.info/blog/000626/ "RC4 終了のお知らせ — Baldanders.info")や SSL/TLS の攻略コードがいくつか「開発」されたことにより TLS 1.2 への移行が強く推奨されることになった。
それを受けての SP800-52 改訂だったのだが，その後の SSL/TLS やその実装である OpenSSL 等のソフトウェアへの攻撃の激しさはみなさんご存じのとおりである。

そうそう。
IPA と言えば最近になってヤバい注意喚起が上がっている。

- [【注意喚起】SQLインジェクションをはじめとしたウェブサイトの脆弱性の再点検と速やかな改修を：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/announce/website_vuln.html)

また2014年に大騒ぎになった Heartbleed 脆弱性をいまだに放置しているサイトもあるようだ。

- [「Heartbleed」脆弱性、多くのサイトやサーバでいまだに存在--Shodan Report - ZDNet Japan](https://japan.zdnet.com/article/35095570/)

攻撃者は既知の攻撃は当然のように試す。
先延ばししていいことは何もない。
いや，マジでお願いしますよ，サイト運用者の方々。

## ブックマーク

- [CRYPTREC Report 2013 — Baldanders.info](https://baldanders.info/blog/000740/)
- [パスワード変更は計画的に — Baldanders.info](https://baldanders.info/blog/000682/)
- [Prohibiting RC4 — Baldanders.info](https://baldanders.info/blog/000810/)

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
