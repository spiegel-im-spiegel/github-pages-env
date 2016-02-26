+++
date = "2016-02-25T16:45:01+09:00"
update = "2016-02-26T17:40:58+09:00"
description = "IPA が提供する “icat for JSON” にアクセスする Go 言語用のパッケージを公開した。"
draft = false
tags = ["golang", "package", "json", "ipa", "security", "vulnerability"]
title = "icat4json 公開"

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

IPA が提供する “[icat for JSON]” にアクセスする [Go 言語]用のパッケージを公開した。

- [github.com/spiegel-im-spiegel/icat4json](https://github.com/spiegel-im-spiegel/icat4json "spiegel-im-spiegel/icat4json: icat for JSON with Golang")

以下のような感じで使える。

```go
package main

import (
	"fmt"
	"log"

	"github.com/spiegel-im-spiegel/icat4json"
)

func main() {
	json, err := icat4json.Get(icat4json.ToolICATW)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Decode()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Title: %v\n", data.Title)
	fmt.Printf("  URL: %v\n", data.Link)
	fmt.Printf(" Date: %v\n", data.Date)
	fmt.Print("Items:\n")
	for _, item := range data.Itemdata {
		fmt.Printf("\t%v: %v (%v)\n", item.Date, item.Title, item.Link)
	}
}
```

```
$ go run sample.go
Title: IPAセキュリティセンター:重要なセキュリティ情報
  URL: https://www.ipa.go.jp/security/vuln/icat.html
 Date: 2016-02-10 11:58:22 +0900 JST
Items:
	2016-02-10 12:00:00 +0900 JST: Microsoft 製品の脆弱性対策について(2016年02月) (http://www.ipa.go.jp/security/ciadr/vul/20160210-ms.html)
	2016-02-10 12:00:00 +0900 JST: Adobe Flash Player の脆弱性対策について(APSB16-04)(CVE-2016-0985等) (http://www.ipa.go.jp/security/ciadr/vul/20160210-adobeflashplayer.html)
	2016-01-20 12:00:00 +0900 JST: Oracle Java の脆弱性対策について(CVE-2016-0494等) (http://www.ipa.go.jp/security/ciadr/vul/20160120-jre.html)
	2016-01-13 12:00:00 +0900 JST: Microsoft 製品の脆弱性対策について(2016年01月) (http://www.ipa.go.jp/security/ciadr/vul/20160113-ms.html)
	2016-01-13 12:00:00 +0900 JST: Adobe Reader および Acrobat の脆弱性対策について(APSB16-02)(CVE-2016-0932等) (http://www.ipa.go.jp/security/ciadr/vul/20160113-adobereader.html)
	2016-01-06 16:40:00 +0900 JST: 【注意喚起】インターネットに接続する複合機等のオフィス機器の再点検を！ (http://www.ipa.go.jp/security/ciadr/vul/20160106-printer.html)
	2016-01-05 14:00:00 +0900 JST: 「DXライブラリ」におけるバッファオーバーフローの脆弱性対策について(JVN#49476817) (http://www.ipa.go.jp/security/ciadr/vul/20160105-jvn.html)
```

“[icat for JSON]” という名前なのに IPA は API の仕様を公開していない。
そこで JavaScript コードの中身を見てみた。

- [icat for JSON について - Qiita](http://qiita.com/spiegel-im-spiegel/items/4acefe47d3dda688a03e)

[`icat4json`] パッケージでは “[icat for JSON]” から取得した JSON データを以下の構造体にデコードする。

```go
//Item - itemdata from icat
type Item struct {
	Title      string    `json:"item_title"`
	Link       string    `json:"item_link"`
	Date       time.Time `json:"item_date"`
	Identifier []string  `json:"item_identifier"`
}

//ICAT - data from icat
type ICAT struct {
	Itemdata []Item    `json:"itemdata"`
	Title    string    `json:"docTitle"`
	Fix      string    `json:"docTitleFix"`
	Link     string    `json:"docLink"`
	Date     time.Time `json:"docDate"`
}
```

JavaScript コードを見ると `item_identifier` 項目は使ってない模様。
`docTitleFix` 項目は中身が `null` かどうかしかチェックしてなくて[^s]， `null` 以外だと `htmlentities()` 関数を通さずに素のまま表示するという恐ろしいことをしている（普通こういうのって boolean 値を使うんじゃないのか？）。
これらの項目は無視でもいいだろう。
本当は IPA が仕様を公開してくれると有難いんだけどねぇ。

[^s]: ところで [Go 言語]における [string] の実体は `[]byte` だが nil 状態はない。 [`icat4json`] パッケージでは `docTitleFix` 項目が `null` の場合は空文字列に展開される。 JSON の `null` 状態を区別したいのであれば [`github.com/guregu/null`](https://github.com/guregu/null "guregu/null: reasonable handling of nullable values") パッケージ等を使う手もある。ちなみに [`github.com/guregu/null`](https://github.com/guregu/null "guregu/null: reasonable handling of nullable values") パッケージの型の実体は [`database/sql`](https://golang.org/pkg/database/sql/ "sql - The Go Programming Language") の `NullString` 型等である。

脆弱性情報をクライアントサイドで取るのなら “[icat for JSON]” ではなく Twitter の @[ICATalerts](https://twitter.com/ICATalerts/) アカウントのタイム・ラインをチェックするほうがオススメ。
JSON データを使うのならサーバサイドでやるべきだよね。
[`icat4json`] パッケージはドメイン・レイヤのエンティティとして使われることを意識している。

{{< fig-img src="/images/icat4json.svg" title="icat4json entity" link="/images/icat4json.svg" >}}

私はたまたま [Go 言語]を勉強中なので [Go 言語]のパッケージとして実装したけど，本当なら Java とか Ruby とか node.js とかサーバサイドの実装があるといいよね。
IPA も jQuery じゃなくて，そういうので実装すればいいのに。

[`icat4json`]: https://github.com/spiegel-im-spiegel/icat4json "spiegel-im-spiegel/icat4json: icat for JSON with Golang"
[icat for JSON]: https://www.ipa.go.jp/security/vuln/icat.html "サイバーセキュリティ注意喚起サービス「icat for JSON」：IPA 独立行政法人 情報処理推進機構"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[string]: http://golang.org/ref/spec#String_types
