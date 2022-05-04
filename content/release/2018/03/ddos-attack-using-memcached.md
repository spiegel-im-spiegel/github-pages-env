+++
title = "Memcached を踏み台にしたとみられる DDoS 攻撃に関する覚え書き"
date = "2018-03-03T19:39:12+09:00"
update = "2018-03-11T14:28:33+09:00"
description = "DDoS 攻撃の踏み台として使われるのは memcached だけではない。ソフトウェアを常に最新に保つとともにアドレス制限を含む適切な設定を行う必要がある。"
image = "/images/attention/tools.png"
tags = [
  "security",
  "vulnerability",
  "ddos",
]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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

## 脆弱性の内容

[memcached] は，名前の通り，分散型メモリキャッシュ・システムである。
この [memcached] を踏み台にしたとみられる大規模な DDoS 攻撃（Distributed Denial of Service attack）が観測されている。

- [memcached のアクセス制御に関する注意喚起](http://www.jpcert.or.jp/at/2018/at180009.html)
- [【重要】memcachedのアクセス制御に関する注意喚起 | さくらインターネット](https://www.sakura.ad.jp/news/sakurainfo/newsentry.php?id=1885)
- [memcached を悪用したDDoS攻撃についてまとめてみた - piyolog](http://d.hatena.ne.jp/Kango/20180301/1519939259)

1.2.7 以降の [memcached] をデフォルト設定のまま利用している場合， 11211/tcp および 11211/udp のポートがアクセス可能な状態になっている場合があるそうだ。
このポートを利用して [memcached] が保持する情報へアクセスしたり他サーバへの攻撃の踏み台にされる可能性がある。

## 影響を受ける製品

- [memcached] 1.2.7 以降 1.5.6 未満でデフォルト設定のまま利用している場合

## 対策・回避策

[memcached] へアクセスできるアドレスやポートを制限する。
[memcached] のセキュリティ設定については以下のページが参考になる。

- [CentOSにMemcachedをインスト－ル | suka4's memo](http://suka4.blogspot.jp/2011/02/centosmemcached.html)

なお 2018-02-27 にリリースされた 1.5.6 ではデフォルト設定でポートをオープンしないようになっているようだ。

- [ReleaseNotes156 · memcached/memcached Wiki](https://github.com/memcached/memcached/wiki/ReleaseNotes156)

DDoS 攻撃の踏み台として使われるのは [memcached] だけではない。
ソフトウェアを常に最新に保つとともにアドレス制限を含む適切な設定を行う必要がある。

- [UDP-Based Amplification Attacks](https://www.us-cert.gov/ncas/alerts/TA14-017A)

## ブックマーク

- [Memcached脆弱性でDDoS踏み台に使われて、プロバイダによるネットワーク制限発生 - Qiita](https://qiita.com/flyjay/items/b9a379ab4ec0f5c0c96e)
- [GitHubに最大1.35TbpsのDDoS攻撃発生。断続的にサービス停止も、短時間で復旧果たす  |  TechCrunch Japan](https://jp.techcrunch.com/2018/03/02/engadget-github-1-35tbps-ddos/)
- [第2回　memcachedのセキュリティと脆弱性：memcachedの活用と運用 実践編｜gihyo.jp … 技術評論社](http://gihyo.jp/dev/feature/01/memcached_advanced/0002)
- [memcachedの開放ポート(11211/tcp, 11211/udp)をサクっと確認する - ろば電子が詰まっている](http://d.hatena.ne.jp/ozuma/20180228/1519828918)
- [memcachedを用いたUDP Amplification攻撃 – wizSafe Security Signal -安心・安全への道標- IIJ](https://wizsafe.iij.ad.jp/2018/03/269/)
- [In-the-wild DDoSes use new way to achieve unthinkable sizes | Ars Technica](https://arstechnica.com/information-technology/2018/02/in-the-wild-ddoses-use-new-way-to-achieve-unthinkable-sizes/)
    - [New DDoS Reflection-Attack Variant - Schneier on Security](https://www.schneier.com/blog/archives/2018/03/new_ddos_reflec.html)
- [【重要】さくらのVPS／クラウド／専用サーバにおけるポート使用状況確認のお願い | さくらインターネット](https://www.sakura.ad.jp/news/sakurainfo/newsentry.php?id=1890)
- [DDoS攻撃、過去最大の1.7Tbpsを記録 - ZDNet Japan](https://japan.zdnet.com/article/35115722/)
- [memcachedを悪用する攻撃、「キルスイッチ」で抑制できる可能性 - ZDNet Japan](https://japan.zdnet.com/article/35115876/)

[memcached]: https://memcached.org/ "memcached - a distributed memory object caching system"
