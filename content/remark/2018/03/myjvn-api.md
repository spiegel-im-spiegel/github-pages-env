+++
title = "MyJVN API に関する覚え書き"
date = "2018-03-12T22:47:09+09:00"
update = "2018-05-13T21:31:14+09:00"
description = "MyJVN API は JVN が提供している「脆弱性対策情報共有フレームワーク」のひとつである。"
image = "/images/attention/kitten.jpg"
tags        = [ "security", "risk", "management", "vulnerability", "cvss", "curl", "jvn", "xml" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
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

[MyJVN API] は [JVN] が提供している「[脆弱性対策情報共有フレームワーク]」のひとつである。
[MyJVN API] ではメインサービスである RESTful API のほか， Twitter での情報配信も行っている。

- [@JVNiPedia](https://twitter.com/JVNiPedia) : 脆弱性情報
- [@MyJVN](https://twitter.com/MyJVN) : バージョン更新情報

今回は RESTful API に絞り，覚え書きとして残しておく（随時加筆予定）。
なお「[脆弱性対策情報共有フレームワーク]」では [MyJVN API] と併せて以下のツールも提供している（どれも機能がイマイチなのが...）。

- MyJVN バージョンチェッカ（Windows 専用）
    - [Java JRE 版](https://jvndb.jvn.jp/apis/myjvn/vccheck.html) [^jre1]
        - [CLI 版](https://jvndb.jvn.jp/apis/myjvn/vccheckcmd.html)
    - [.NET 版](https://jvndb.jvn.jp/apis/myjvn/vccheckdotnet.html) : GUI 版と CUI 版がある
- [MyJVN 脆弱性対策情報収集ツール](https://jvndb.jvn.jp/apis/myjvn/mjcheck.html)
    - [MyJVN 脆弱性対策情報フィルタリング収集ツール](https://jvndb.jvn.jp/apis/myjvn/mjcheck3.html) （要 Adobe AIR）

[^jre1]: Java JRE には Windows アカウント名に全角文字を含む場合にエラーになる不具合があるらしい。 JRE 実装に起因するものなので JRE 側で修正されない限り対応できないようだ。 Windows 専用ツールでわざわざ Java を使うメリットはないので .NET 版を使うことを強くお勧めする。

## [MyJVN API] のライセンス

[MyJVN API] 自体の利用については以下に記載がある。

- [MyJVN - API: 利用上の留意事項](https://jvndb.jvn.jp/apis/termsofuse.html)

これを見れば分かるが， [MyJVN API] の利用には制限がありオープンでもフリーでもない点は注意が必要である。

また [JVN] が提供しているデータにはデータベースの著作権が発生する筈だが，データの利用許諾範囲が明示されていないため，このままでは（著作権法上は）利用できない[^c1]。
データの取り扱いについて一定のリスクがある点も注意すべきだろう。

[^c1]: 念のために説明しておくと著作権では著作物の「使用」と「利用」を区別し「利用」について制限をかける（「使用」については著作権は関知しない）。著作物の利用とは，大まかに *複製（フォーマット変換・機械翻訳を含む）・公開（上演や展示など）・配布（共有）・翻案（改変）* を指す。 [MyJVN API] で取得したデータの利用については別途許諾を得る必要がある？

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
curl -G -d "method=getVulnOverviewList" -d "feed=hnd" https://jvndb.jvn.jp/myjvn
```

### 必須オプション

必須オプションは以下の通り

| オプション名 | 値                    |
| ------------ | --------------------- |
| `method`     | `getVulnOverviewList` |
| `feed`       | `hnd`                 |

### 取得のハンドリング

脆弱性対策概要情報一覧は一度に50件までしか取得できない。
取得可能件数が50件以上ある場合はオプションで「n件目から取得」と指定すればよい。

| オプション名   | 内容                             | 既定値      |
| -------------- | -------------------------------- | ----------- |
| `startItem`    | 取得可能なエントリの取得開始位置 | 1           |
| `maxCountItem` | 一度に取得できるエントリ数       | 50 (max 50) |

取得可能件数は，返却された XML データの `<status:Status>` 要素に記述されている。

### ベンダおよび製品情報によるフィルタリング

ベンダおよび製品情報によるフィルタリングは以下の通り。

| オプション名 | 内容                                                        | 既定値 |
| ------------ | ----------------------------------------------------------- | ------ |
| `cpeName`    | CPE製品名。ワイルドカード使用可。複数指定可（`+` で区切る） | なし   |
| `vendorId`   | ベンダID。複数指定可（`+` で区切る）                        | なし   |
| `productId`  | 製品ID。複数指定可（`+` で区切る）                          | なし   |

`cpeName`, `vendorId`, `productId` はいずれかひとつのみ指定可能。

もっと大雑把にキーワードを指定してフィルタリングすることもできる。

| オプション名 | 内容                                                                             | 既定値 |
| ------------ | -------------------------------------------------------------------------------- | ------ |
| `keyword`    | 脆弱性概要情報の部分一致。ワイルドカード使用不可。大文字小文字の区別なし。 UTF-8 | なし   |

### CVSS 評価値によるフィルタリング

| オプション名 | 内容                                                | 既定値 |
| ------------ | --------------------------------------------------- | ------ |
| `severity`   | CVSS 深刻度（`n/l/m/h/c`）                          | なし   |
| `vector`     | CVSS 基本評価基準。ベクタ値（`CVSS:3.0/...`）で指定 | なし   |

CVSS は Version 3 のみ対応のようだ。

### 期間指定によるフィルタリング

期間指定によるフィルタリングは発見日・発行日・更新日に対応している。

まずは発見日によるフィルタリング。

| オプション名               | 内容                                                                     | 既定値      |
| -------------------------- | ------------------------------------------------------------------------ | ----------- |
| `rangeDatePublic`          | 発見日の範囲（n/w/m）                                                    | `w`         |
| `datePublicStartY`         | 発見日開始年（整数4桁）                                                  | なし        |
| `datePublicStartM`         | 発見日開始月（整数 1-12）                                                | なし        |
| `datePublicStartD`         | 発見日開始日（整数1-31）                                                 | なし        |
| `datePublicEndY`           | 発見日終了年（整数4桁）                                                  | なし        |
| `datePublicEndM`           | 発見日終了月（整数 1-12）                                                | なし        |
| `datePublicEndD`           | 発見日終了日（整数 1-31）                                                | なし        |

`rangeDatePublic` で `w` を指定すると当日を含む直近一週間を， `m` を指定すると直近１ヶ月の情報を取得する。
期間を指定する場合は `rangeDatePublic` に `n` を指定して開始および終了条件をセットする。
開始および終了条件は年・年月・年月日で指定できる。

更新日によるフィルタリングは以下の通り。

| オプション名               | 内容                                                                     | 既定値      |
| -------------------------- | ------------------------------------------------------------------------ | ----------- |
| `rangeDatePublished`       | 更新日の範囲（n/w/m）                                                    | `w`         |
| `datePublishedStartY`      | 更新日開始年（整数4桁）                                                  | なし        |
| `datePublishedStartM`      | 更新日開始月（整数 1-12）                                                | なし        |
| `datePublishedStartD`      | 更新日開始日（整数 1-31）                                                | なし        |
| `datePublishedEndY`        | 更新日終了年（整数4桁）                                                  | なし        |
| `datePublishedEndM`        | 更新日終了月（整数 1-12）                                                | なし        |
| `datePublishedEndD`        | 更新日終了日（整数 1-31）                                                | なし        |

指定の仕方は発見日と同じ要領でできる。

発行日によるフィルタリングは以下の通り。

| オプション名               | 内容                                                                     | 既定値      |
| -------------------------- | ------------------------------------------------------------------------ | ----------- |
| `rangeDateFirstPublished`  | 発行日の範囲（n/w/m）                                                    | `w`         |
| `dateFirstPublishedStartY` | 発行日開始年（整数4桁）                                                  | なし        |
| `dateFirstPublishedStartM` | 発行日開始月（整数 1-12）                                                | なし        |
| `dateFirstPublishedStartD` | 発行日開始日（整数 1-31）                                                | なし        |
| `dateFirstPublishedEndY`   | 発行日終了年（整数4桁）                                                  | なし        |
| `dateFirstPublishedEndM`   | 発行日終了月（整数 1-12）                                                | なし        |
| `dateFirstPublishedEndD`   | 発行日終了日（整数 1-31）                                                | なし        |

指定の仕方は発見日と同じ要領でできる。

### 表示言語の指定

表示言語は日本語と英語のみ対応している。

| オプション名 | 内容                | 既定値 |
| ------------ | ------------------- | ------ |
| `lang`       | 表示言語（`ja/en`） | `ja`   |


### 返却地

返却データのフォーマットは XML で RSS 1.0 形式。
ただし [mod_sec] と呼ばれる [JVN] 独自のスキーマを含んでいて，通常のフィード情報の他に脆弱性情報の記述もある。
なお，詳細を取得したいのであれば次節の `getVulnDetailInfo` を使うほうがよい。

XML の大まかな構成は以下の通り。

```xml
<rdf:RDF>
  <channel>
    ....
  </channel>
  <item>
    <title>...</title>
    <link>>https://jvndb.jvn.jp/...</link>
    <description>...</description>
    <sec:identifier>JVNDB-2018-XXXXXX</sec:identifier>
    <sec:references source="JVN" id="JVN#XXXXXXXX">https://jvn.jp/jp/JVNXXXXXXXX/index.html</sec:references>
    <sec:references source="CVE" id="CVE-2018-XXXXX">https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-XXXXX</sec:references>
    <sec:references id="CWE-XX" title="...">https://cwe.mitre.org/data/definitions/XX.html</sec:references>
    <sec:references>...</sec:references>
    <sec:cpe version="2.2" vendor="...." product="....">cpe:/.....</sec:cpe>
    <sec:cpe>...</sec:cpe>
    <sec:cvss score="8.8" severity="High" vector="CVSS:3.0/..." version="3.0" type="Base"/>
    <sec:cvss score="5.8" severity="Medium" vector="..." version="2.0" type="Base"/>
    <dc:date>2006-01-02T15:04:05+09:00</dc:date>
    <dcterms:issued>2006-01-02T15:04:05+09:00</dcterms:issued>
    <dcterms:modified>2006-01-02T15:04:05+09:00</dcterms:modified>
  </item>
  <item>
    ....
  </item>
  <status:Status version="3.3" method="getVulnOverviewList" lang="ja" retCd="0" retMax="X" errCd="" errMsg="" totalRes="X" totalResRet="X" firstRes="X" feed="hnd" ... />
</rdf:RDF>
```

`<status:Status>` 要素には API 処理時のステータスが返る。
正常終了であれば `retCd` に `0` がセットされる。
また取得可能件数は `totalRes` 属性にセットされる。

## 脆弱性対策詳細情報の取得

- [getVulnDetailInfo (ver. HND)](https://jvndb.jvn.jp/apis/getVulnDetailInfo_api_hnd.html)

### [cURL as DSL]

とりあえず必須オプションのみ。
以下は JVNDB-2018-000024 の情報を取得する場合。    

```
curl -G -d "method=getVulnDetailInfo" -d "feed=hnd" -d "vulnId=JVNDB-2018-000024" https://jvndb.jvn.jp/myjvn
```

### 必須オプション

必須オプションは以下の通り

| オプション名 | 値                                            |
| ------------ | --------------------------------------------- |
| `method`     | `getVulnDetailInfo`                           |
| `feed`       | `hnd`                                         |
| `vulnId`     | 脆弱性対策情報ID 。複数指定可（`+` で区切る） |

### 取得のハンドリング

脆弱性対策詳細情報は一度に10件までしか取得できない。
取得可能件数が10件以上ある場合はオプションで「n件目から取得」と指定すればよい。

| オプション名   | 内容                             | 既定値      |
| -------------- | -------------------------------- | ----------- |
| `startItem`    | 取得可能なエントリの取得開始位置 | 1           |
| `maxCountItem` | 一度に取得できるエントリ数       | 50 (max 50) |

ただし取得可能件数は `vulnId` オプションでコントロールできるため，あまり意味のあるオプションではないだろう。

### 表示言語の指定

表示言語は日本語と英語のみ対応している。

| オプション名 | 内容                | 既定値 |
| ------------ | ------------------- | ------ |
| `lang`       | 表示言語（`ja/en`） | `ja`   |


### 返却地

返却データのフォーマットは XML で [VULDEF] と呼ばれる [JVN] 独自のスキーマを使っている。
HTML ページ（たとえば [JVNDB-2018-000024](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000024.html "JVNDB-2018-000024 - JVN iPedia - 脆弱性対策情報データベース")）に記載されている内容はほぼ網羅されているため HTML ページをわざわざ scraping する必要はない。

XML の大まかな構成は以下の通り。

```xml
<VULDEF-Document>
  <Vulinfo>
    <VulinfoID>JVNDB-2018-XXXXXXX</VulinfoID>
    <VulinfoData>
      <Title>...</Title>
      <VulinfoDescription>
        <Overview>...</Overview>
      </VulinfoDescription>
      <Affected>
        <AffectedItem>
          <Name>...</Name>
          <ProductName>...</ProductName>
          <Cpe version="2.2">...</Cpe>
          <VersionNumber>...</VersionNumber>
        </AffectedItem>
        <AffectedItem>
          ...
        </AffectedItem>
      </Affected>
      <Impact>
        <Cvss version="2.0">
          <Severity type="Base">...</Severity>
          <Base>X.X</Base>
          <Vector>...</Vector>
        </Cvss>
        <Cvss version="3.0">
          <Severity type="Base">...</Severity>
          <Base>X.X</Base>
          <Vector>CVSS:3.0/...</Vector>
        </Cvss>
        <ImpactItem>
          <Description>...</Description>
        </ImpactItem>
      </Impact>
      <Solution>
        <SolutionItem>
          <Description>...</Description>
        </SolutionItem>
        <SolutionItem>
          ...
        </SolutionItem>
      </Solution>
      <Related>
        <RelatedItem type="...">
          <Name>...</Name>
          <VulinfoID>...</VulinfoID>
          <Title>...</Title>
          <URL>...</URL>
        </RelatedItem>
        <RelatedItem type="...">
          ...
        </RelatedItem>
      </Related>
      <History>
        <HistoryItem>
          <HistoryNo>1</HistoryNo>
          <DateTime>2006-01-02T15:04:05+09:00</DateTime>
          <Description>...</Description>
        </HistoryItem>
        <HistoryItem>
          ...
        </HistoryItem>
      </History>
      <DateFirstPublished>2006-01-02T15:04:05+09:00</DateFirstPublished>
      <DateLastUpdated>2006-01-02T15:04:05+09:00</DateLastUpdated>
      <DatePublic>2006-01-02T15:04:05+09:00</DatePublic>
    </VulinfoData>
  </Vulinfo>
  <Vulinfo>
    ....
  </Vulinfo>
  <sec:handling>
    <marking:Marking>
      <marking:Marking_Structure ... />
    </marking:Marking>
  </sec:handling>
  <status:Status version="3.3" method="getVulnDetailInfo" lang="ja" retCd="0" retMax="X" errCd="" errMsg="" totalRes="X" totalResRet="X" firstRes="X" feed="hnd" vulnId="..." ... />
</VULDEF-Document>
```

`<status:Status>` 要素には API 処理時のステータスが返る。
正常終了であれば `retCd` に `0` がセットされる。
また取得可能件数は `totalRes` 属性にセットされる。

## ブックマーク

- [JVNRSS: Specification of Japan Vulnerability Notes RSS](https://jvndb.jvn.jp/schema/jvnrss.html)
    - [JVNRSS: Qualified Security Advisory Reference (mod_sec)](https://jvndb.jvn.jp/schema/mod_sec.html)
- [VULDEF: The VULnerability Data publication and Exchange Format data model](http://jvnrss.ise.chuo-u.ac.jp/jtg/vuldef/index.ja.html)
- [XMLの処理 · Build web application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/ja/07.1.html)

- [Vuls · Agentless Vulnerability Scanner for Linux/FreeBSD](https://vuls.io/)
    - [future-architect/vuls: Vulnerability scanner for Linux/FreeBSD, agentless, written in Go](https://github.com/future-architect/vuls/)

- [“JVN iPedia”がHTTPS対応などのリニューアル、「MyJVNバージョンチェッカ」は要更新 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1107654.html)

- [JVN が CVSSv3 による脆弱性評価を開始]({{< ref "/remark/2015/cvss-v3-metrics-in-jvn.md" >}})
- [go-myjvn パッケージを作ってみた]({{< ref "/release/2018/03/go-myjvn-v0_1_0-released.md" >}})

[JVN]: https://jvn.jp/ "Japan Vulnerability Notes"
[脆弱性対策情報共有フレームワーク]: https://jvndb.jvn.jp/apis/myjvn/ "脆弱性対策情報共有フレームワーク - MyJVN"
[MyJVN API]: https://jvndb.jvn.jp/apis/
[cURL as DSL]: https://shibukawa.github.io/curl_as_dsl/ "cURL as DSL — cURL as DSL 1.0 documentation"
[mod_sec]: https://jvndb.jvn.jp/schema/mod_sec.html "JVNRSS: Qualified Security Advisory Reference (mod_sec)"
[VULDEF]: http://jvnrss.ise.chuo-u.ac.jp/jtg/vuldef/index.ja.html "VULDEF: The VULnerability Data publication and Exchange Format data model"
