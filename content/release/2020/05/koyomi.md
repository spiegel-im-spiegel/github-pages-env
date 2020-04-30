+++
title = "日本の暦情報を取得するパッケージを作ってみた"
date =  "2020-05-01T05:52:53+09:00"
description = "「国立天文台 天文情報センター 暦計算室」より日本の暦情報を取得する Go 言語用パッケージです。"
image = "/images/attention/tools.png"
tags = [ "astronomy", "ephemeris", "japanese", "calendar", "golang", "package" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いやね。

- [国立天文台の暦要項データを取得する - Qiita](https://qiita.com/hadzimme/items/a9460bec4b2c044da2ba)

ちう記事を見かけて「[Go 言語]でも似たようなコードを書けばいいぢゃん」と軽く考えたわけですよ。

実は Google Calendar を操作する [Go 言語]用パッケージとしては Google 公式の

- [google.golang.org/api/calendar/v3](https://pkg.go.dev/google.golang.org/api/calendar/v3 "calendar package · go.dev")

ってのがあるのだが，これって認証とか含めたガチなやつなのよ。
でも欲しいのは国立天文台から「公開」されているただの暦情報なので，こんなガチなやつは（面倒くさいだけだし）要らないわけ。

もっとお気楽に使える iCal パーサがないかなぁ，と思ったらありました。

- [PuloV/ics-golang: Golang ics parser](https://github.com/PuloV/ics-golang)

ありがたや。

早速，これを使って日本の暦情報を取得するパッケージを作ってみた。

- [spiegel-im-spiegel/koyomi: 日本のこよみ](https://github.com/spiegel-im-spiegel/koyomi)

これを使った以下のコードは，2020年5月の祝日・休日と二十四節気・雑節を CSV 形式で出力するもの。

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/koyomi"
)

func main() {
	start, _ := koyomi.DateFrom("2020-05-01")
	end, _ := koyomi.DateFrom("2020-05-31")
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(koyomi.Holiday, koyomi.SolarTerm),
		koyomi.WithStartDate(start),
		koyomi.WithEndDate(end),
	).Get()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	b, err := k.EncodeCSV()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	io.Copy(os.Stdout, bytes.NewReader(b))
}
```

これを実行すると

```text
$ go run sample.go 
"Date","Title"
"2020-05-01","八十八夜"
"2020-05-03","憲法記念日"
"2020-05-04","みどりの日"
"2020-05-05","こどもの日"
"2020-05-05","立夏"
"2020-05-06","休日"
"2020-05-20","小満"
```

てな感じになる。
また

```go
b, err := k.EncodeCSV()
```

の部分を

```go
b, err := k.EncodeJSON()
```

と書き換えれば

```json
$ go run sample.go | jq .
[
  {
    "Date": "2020-05-01",
    "Title": "八十八夜"
  },
  {
    "Date": "2020-05-03",
    "Title": "憲法記念日"
  },
  {
    "Date": "2020-05-04",
    "Title": "みどりの日"
  },
  {
    "Date": "2020-05-05",
    "Title": "こどもの日"
  },
  {
    "Date": "2020-05-05",
    "Title": "立夏"
  },
  {
    "Date": "2020-05-06",
    "Title": "休日"
  },
  {
    "Date": "2020-05-20",
    "Title": "小満"
  }
]
```

てな感じに JSON 形式でも出力できる。

[`koyomi`]`.WithCalendarID()`, [`koyomi`]`.WithStartDate()`, [`koyomi`]`.WithEndDate()` 各関数は [Functional Option]({{< ref "/golang/functional-options-pattern.md" >}} "インスタンスの生成と Functional Options パターン") なので省略可能である。
まぁ，全部省略したら何も取れないけど（笑）

[`koyomi`]`.WithCalendarID()` 関数には1個以上の [`koyomi`]`.CalendarID` を指定できる。
指定できる  [`koyomi`]`.CalendarID`  は以下の通り。

```go
const (
    Holiday   CalendarID = iota + 1 //国民の祝日および休日
    MoonPhase                       //朔弦望
    SolarTerm                       //二十四節気・雑節
    Eclipse                         //日食・月食・日面経過
    Planet                          //惑星現象
)
```

取得できるイベントは日本時間がベースになっていて，しかも（終日イベントなので）日付のみ有効である。
時刻情報はカットされているのであしからず。

あと [PuloV/ics-golang](https://github.com/PuloV/ics-golang "PuloV/ics-golang: Golang ics parser") パッケージの仕様の問題で，リモートにある iCal ファイルを一時ファイルに落とし込むようだ。
落とし込み先ディレクトリの既定がカレントの `tmp/` になっている（ない場合は `tmp/` ディレクトリを作成しようとする）。
このディレクトリを指定するのであれば

```go {hl_lines=[5]}
k, err := koyomi.NewSource(
    koyomi.WithCalendarID(koyomi.Holiday, koyomi.SolarTerm),
    koyomi.WithStartDate(start),
    koyomi.WithEndDate(end),
    koyomi.WithTempDir("/home/username/.cache/"),
).Get()
```

などとすればよい。

ぶっちゃけ遅いです。
まぁ Google Calndar から iCal ファイルをフィルタリングなしでまるっと取ってくるのだから遅いに決まってるのだけど。
実務で使うならバッチ処理でデータベース等に保持っておくのがよろしいかと思われ。

## ブックマーク

- [暦要項 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/yoko/)
- [暦Wiki - 国立天文台暦計算室](http://eco.mtk.nao.ac.jp/koyomi/wiki/)
    - [暦Wiki/月の満ち欠け/いろいろな月たち - 国立天文台暦計算室](http://eco.mtk.nao.ac.jp/koyomi/wiki/B7EEA4CECBFEA4C1B7E7A4B12FA4A4A4EDA4A4A4EDA4CAB7EEA4BFA4C1.html)
    - [暦Wiki/季節/二十四節気とは？ - 国立天文台暦計算室](http://eco.mtk.nao.ac.jp/koyomi/wiki/B5A8C0E12FC6F3BDBDBBCDC0E1B5A4A4C8A4CFA1A9.html)
    - [暦Wiki/季節/雑節とは？ - 国立天文台暦計算室](http://eco.mtk.nao.ac.jp/koyomi/wiki/B5A8C0E12FBBA8C0E1A4C8A4CFA1A9.html)
    - [暦Wiki/日面経過 - 国立天文台暦計算室](http://eco.mtk.nao.ac.jp/koyomi/wiki/C6FCCCCCB7D0B2E1.html)
- [こよみ用語解説 天象 - 国立天文台暦計算室](http://eco.mtk.nao.ac.jp/koyomi/faq/phenomena.html)
- [カレンダーに祝日を入れたいなら国立天文台へ行けばいいじゃない]({{< ref "/remark/2019/05/google-ephemeris.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`koyomi`]: https://github.com/spiegel-im-spiegel/koyomi "spiegel-im-spiegel/koyomi: 日本のこよみ"
