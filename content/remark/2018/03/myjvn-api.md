+++
title = "MyJVN API に関する覚え書き"
date = "2018-03-12T22:47:09+09:00"
update = "2018-03-17T02:05:27+09:00"
description = "MyJVN API は JVN が提供している「脆弱性対策情報共有フレームワーク」のひとつである。"
image = "/images/attention/kitten.jpg"
tags        = [ "security", "risk", "management", "Vulnerability", "cvss", "curl", "jvn", "xml" ]

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

[MyJVN API] は [JVN] が提供している「[脆弱性対策情報共有フレームワーク]」のひとつである。
[MyJVN API] ではメインサービスである RESTful API のほか， Twitter での情報配信も行っている。

- [@JVNiPedia](https://twitter.com/JVNiPedia) : 脆弱性情報
- [@MyJVN](https://twitter.com/MyJVN) : バージョン更新情報

今回は RESTful API に絞り，覚え書きとして残しておく（随時加筆予定）。
なお「[脆弱性対策情報共有フレームワーク]」では [MyJVN API] と併せて以下のツールも提供している（どれも機能がイマイチなのが...）。

- MyJVN バージョンチェッカ（Windows 専用）
    - [JRE 版](https://jvndb.jvn.jp/apis/myjvn/vccheck.html)
    - [.NET 版](https://jvndb.jvn.jp/apis/myjvn/vccheckdotnet.html)
    - [CLI 版](https://jvndb.jvn.jp/apis/myjvn/vccheckcmd.html) （要 JRE）
- [MyJVN 脆弱性対策情報収集ツール](https://jvndb.jvn.jp/apis/myjvn/mjcheck.html)
    - [MyJVN 脆弱性対策情報フィルタリング収集ツール](https://jvndb.jvn.jp/apis/myjvn/mjcheck3.html) （要 Adobe AIR）

## [MyJVN API] のライセンス

[MyJVN API] 自体の利用については以下に記載がある。

- [MyJVN - API: 利用上の留意事項](https://jvndb.jvn.jp/apis/termsofuse.html)

これを見れば分かるが， [MyJVN API] の利用には制限がありオープンでもフリーでもない点は注意が必要である。

また [JVN] が提供しているデータにはデータベースの著作権が発生する筈だが，データの利用許諾範囲が明示されていないため，このままでは（著作権法上は）利用できない[^c1]。
データの取り扱いについて一定のリスクがある点も注意すべきだろう。

[^c1]: [MyJVN API] で取得したデータの利用については別途許諾を得る必要がある？

## [MyJVN API] のバージョンアップ

[MyJVN API] は 2018-02-21 にバージョンアップしたが，これに伴い旧バージョンの API のほうは 2018-04-02 から使えなくなる。

{{< fig-quote title="SSL暗号化通信への対応に伴う注意事項（MyJVN API、MyJVNバージョンチェッカ等の仕様変更について）" link="https://jvndb.jvn.jp/apis/myjvn/svc-change.html" >}}
<q>2018年2月20日以前の旧サービス（MyJVN API, MyJVNバージョンチェッカ等）は、2018年4月2日（月）以降は継続したサービス利用ができなくなります。旧サービスを継続して利用する場合には以下の案内を参考にして、MyJVN APIを使用する利用者プログラムを改修する、あるいは新ツールの再ダウンロードによる利用、などの対処を実施してください。</q>
{{< /fig-quote >}}

RESTful API についてはスキーマが HTTP から HTTPS へ変更となり，場合によってはバージョン名（HND/ITM）を付加する必要がある。

深刻度（severity）の評価に注意すること。
旧バージョンでは CVSSv2 ベースだったが，新バージョンでは CVSSv3 ベースになっている。
両者は以下のように異なっている。

| 深刻度 | CVSSv2 Base スコア |
| ------ | ------------------ |
| 危険   | 7.0 - 10.0         |
| 警告   | 4.0 - 6.9          |
| 注意   | 0.0 - 3.9          |

| 深刻度           | CVSSv3 Base スコア |
| ---------------- | ------------------ |
| 緊急（Critical） | 9.0 - 10.0         |
| 重要（High）     | 7.0 - 8.9          |
| 警告（Medium）   | 4.0 - 6.9          |
| 注意（Low）      | 0.1 - 3.9          |
| なし（None）     | 0                  |

以降から個々の API について概説する。

## 脆弱性対策概要情報一覧の取得

- [getVulnOverviewList (ver. HND)](https://jvndb.jvn.jp/apis/getVulnOverviewList_api_hnd.html)

### [cURL as DSL]

とりあえず必須オプションのみ[^curl1]。
過去1週間の情報を取得する。

[^curl1]: `-G` オプションは GET プロトコルを指す。明示するなら `-X GET` とするのがいいかも。

```
curl -G -d "method=getVulnOverviewList" -d "feed=hnd" -d "lang=ja" https://jvndb.jvn.jp/myjvn
```

必須オプションの他に期間やフィルタリング条件などを設定できる。
1回の API 発行で取得できる最大件数は50。

出力フォーマットは XML で RSS 1.0 形式。
ただし [mod_sec] と呼ばれる [JVN] 独自のスキーマを含んでいて，通常のフィード情報の他に脆弱性情報も記述されている。
なお，詳細を取得したいのであれば次の `getVulnDetailInfo` を使うほうがよい。

## 脆弱性対策詳細情報の取得

- [getVulnDetailInfo (ver. HND)](https://jvndb.jvn.jp/apis/getVulnDetailInfo_api_hnd.html)

### [cURL as DSL]

とりあえず必須オプションのみ。
以下は JVNDB-2018-000024 の情報を取得する場合。    

```
curl -G -d "method=getVulnDetailInfo" -d "feed=hnd" -d "lang=ja" -d "vulnId=JVNDB-2018-000024" https://jvndb.jvn.jp/myjvn
```

脆弱性対策情報ID（`vulnId`）が複数ある場合は `vulnId=JVNDB-2018-000024+JVNDB-2018-000022` という感じに `+` で繋いで指定できる。
1回の API 発行で取得できる最大件数は10。

出力フォーマットは XML で [VULDEF] と呼ばれる [JVN] 独自のスキーマを使っている。
HTML ページ（たとえば [JVNDB-2018-000024](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000024.html "JVNDB-2018-000024 - JVN iPedia - 脆弱性対策情報データベース")）に記載されている内容はほぼ網羅されているため HTML ページをわざわざ scraping する必要はない。

## ブックマーク

- [JVNRSS: Specification of Japan Vulnerability Notes RSS](https://jvndb.jvn.jp/schema/jvnrss.html)
    - [JVNRSS: Qualified Security Advisory Reference (mod_sec)](https://jvndb.jvn.jp/schema/mod_sec.html)
- [VULDEF: The VULnerability Data publication and Exchange Format data model](http://jvnrss.ise.chuo-u.ac.jp/jtg/vuldef/index.ja.html)
- [XMLの処理 · Build web application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/ja/07.1.html)

- [Vuls · Agentless Vulnerability Scanner for Linux/FreeBSD](https://vuls.io/)
    - [future-architect/vuls: Vulnerability scanner for Linux/FreeBSD, agentless, written in Go](https://github.com/future-architect/vuls/)

- [“JVN iPedia”がHTTPS対応などのリニューアル、「MyJVNバージョンチェッカ」は要更新 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1107654.html)

- [JVN が CVSSv3 による脆弱性評価を開始]({{< relref "remark/2015/cvss-v3-metrics-in-jvn.md" >}})
- [go-myjvn パッケージを作ってみた]({{< relref "release/2018/03/go-myjvn-v0_1_0-released.md" >}})

[JVN]: https://jvn.jp/ "Japan Vulnerability Notes"
[脆弱性対策情報共有フレームワーク]: https://jvndb.jvn.jp/apis/myjvn/ "脆弱性対策情報共有フレームワーク - MyJVN"
[MyJVN API]: https://jvndb.jvn.jp/apis/
[cURL as DSL]: https://shibukawa.github.io/curl_as_dsl/ "cURL as DSL — cURL as DSL 1.0 documentation"
[mod_sec]: https://jvndb.jvn.jp/schema/mod_sec.html "JVNRSS: Qualified Security Advisory Reference (mod_sec)"
[VULDEF]: http://jvnrss.ise.chuo-u.ac.jp/jtg/vuldef/index.ja.html "VULDEF: The VULnerability Data publication and Exchange Format data model"
