+++
title = "Typst におけるデータと制御"
date =  "2025-03-02T21:46:52+09:00"
description = "データと制御が分離しやすくコードが（比較的）書きやすいというのは Typst の利点だと思う。 "
image = "/images/attention/tools.png"
tags  = [ "typst", "programming" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[Typst] は CSV や JSON などのテキストベースのデータ（ファイル）を読み込んで使うことができる。
簡単な例をいくつか挙げてみる。

## CSV データの読み込みと表示

以下の内容の CSV ファイルがあるとする。

```csv
"日付","曜日","名称"
"2025年5月3日","土","憲法記念日"
"2025年5月4日","日","みどりの日"
"2025年5月5日","月","こどもの日"
"2025年5月6日","火","休日"
```

これを読み込んで表にすることを考える。
[Typst] のコードはこんな感じ。

```typst {hl_lines=["12-16"]}
#show raw: body => {
    set text(font: (
      (
        name: "Inconsolata",
        covers: "latin-in-cjk",
      ),
      "Noto Sans CJK JP"
    ))
    body
}

#let holidays = csv(
  "./holidays.csv",
  delimiter: ",",
  row-type: dictionary,
)

#holidays
```

フォントの指定については今回はスルーで[^r1]（笑） CSV データの読み込みには [`raw`] 関数を使う。
今回のように1行目がヘッダ情報になっている場合は `row-type` に [`dictionary`] を指定する。
ヘッダ情報がない場合は既定の [`array`] でOK。

[^r1]: データのダンプ表示時のフォント指定については「[Typst のドキュメント要素]({{< relref "./3-typst-document-elements-1.md" >}})」の [`raw`] の説明を参照のこと。

これを PDF に出力すると以下のような内容になる。

{{< fig-img src="./csv-1.png" title="CSV データの読み込み" link="./csv-1.pdf" width="677" >}}

見ての通り連想配列（[`dictionary`]）の配列（[`array`]）という構造になっている。

次にこれをヘッダ情報とデータに分離する。
[Typst] のコードはこんな感じ。

```typst {hl_lines=["17-18"]}
#show raw: body => {
    set text(font: (
      (
        name: "Inconsolata",
        covers: "latin-in-cjk",
      ),
      "Noto Sans CJK JP"
    ))
    body
}

#let holidays = csv(
  "./holidays.csv",
  delimiter: ",",
  row-type: dictionary,
)
#let header = holidays.first().keys()
#let data = holidays.map(holiday => holiday.values())

#header

#data
```

PDF への出力結果は以下の通り。

{{< fig-img src="./csv-2.png" title="ヘッダとデータを分離" link="./csv-2.pdf" width="677" >}}

`data` は2次元配列になっている点に注意。
一応，元データの並び順のままヘッダ情報もデータも取れるんだね。
[`array`] にはコンテナ操作ではお馴染みの `filter`, `map`, `fold` といったメソッドが使える。
ありがたや。

これで CSV データは取れたので [`table`] へ展開してみる。

```typst
#set text(font: "NOTO Serif CJK JP", lang: "ja")

//関数定義
#let tableOfHolidays(path) = {
  let holidays = csv( //CSV ファイルの読み込み
    path,
    delimiter: ",",
    row-type: dictionary,
  )

  if holidays.len() > 0 { //データがある場合のみテーブルを表示
    let header = holidays.first().keys() //ヘッダ情報の抽出
    table(
      columns: header.len(), //ヘッダ情報の要素数
      align: header.map(it => {
          if it == "日付" {
              right
          } else if it == "曜日" {
              center
          } else {
              left
          }
        }
      ), //ヘッダ情報の名前によって文字列の寄せを設定
      fill: (x, y) => if y == 0 {
        green.lighten(80%)
      }, //ヘッダ部の背景色を設定
      table.header(..header.map(it => {
          set text(font: "NOTO Sans CJK JP", weight: "bold")
          it
      })), //ヘッダ情報，文字コードも併せて設定している
      ..holidays.map(holiday => holiday.values()).flatten() //データを一次元のデータの並びに展開
    )
  }
}

//CSV ファイルの読み込んでテーブルを表示
#tableOfHolidays("./holidays.csv")
```

まず `tableOfHolidays` 関数を定義して CSV ファイルへのパスを引数とする（CSV 形式の文字列でも可）。
`tableOfHolidays` 関数内では CSV ファイルからデータを取得して [`table`] へ展開している。
最後に `tableOfHolidays` 関数に CSV ファイルへのパスを渡して実行する。

配列に対する `flatten` 関数は多次元配列を一次元配列に展開する。

配列の頭に付いている `..` は配列を要素の並びに展開する。
関数の引数で `min(..nums)` みたいな感じでよく使われる。

あとはヘッダ部の装飾のためにごちゃごちゃ書いているが，詳細は割愛する。
そんなもんと思って眺めていただければ（笑）

PDF への出力結果は以下の通り。

{{< fig-img src="./csv-3.png" title="CSV からテーブル生成" link="./csv-3.pdf" width="677" >}}

まぁ，こんなもんかな。

## JSON データを読み込んでカレンダーを作ろう

もうひとつ。
練習問題としてカレンダーを作ってみる。

[Typst] には日時情報を操作する型として [`datetime`] があるのだが，今回の用途には向かないので他の手段を考える。

ある月のカレンダーを組む際に必要な情報としては以下のものがあればいいだろう。

- 年
- 月
- 月初日の曜日（`0` 〜 `6`， `0` が日曜日）
- 月の最終日

[`datetime`] では「月の最終日」を取得するのが難しい（もしくは面倒くさい）んだよね。

というわけで，まずは `#let calendar(year, month, first_weekday, lastday) = { ... }` という関数を定義してみる。

```typst {hl_lines=["5-13"]}
#set text(font: "NOTO Serif CJK JP", lang: "ja")

//カレンダーを作成
#let calendar(year, month, first_weekday, lastday) = {
  let days = ()
  let i = 0
  while i < first_weekday { //初日の曜日まで空白を追加
    days.push("")
    i = i + 1
  }
  days = days + range(1, lastday + 1).map(day => { //日付を追加
    [#day]
  })
  //カレンダーを作成
  table(
    stroke: (x, y) => if y == 1 {//罫線を設定
      (bottom: 0.7pt + black)
    },
    align: (x, y) => ( //文字の位置を設定
      if y > 1 { right }
      else { center }
    ),
    columns: 7, //列数を設定
    table.header( //ヘッダーを設定
      table.cell( //年月を設定
        colspan: 7,
        [
          #set text(font: "NOTO Sans CJK JP", weight: "bold")
          #year 年 #month 月
        ]
      ),
      ..(text(red)[日], [月], [火], [水], [木], [金], text(blue)[土]).map(it => { //曜日を設定
        set text(font: "NOTO Sans CJK JP", weight: "bold")
        it
      })
    ),
    ..days.enumerate(start:0).map(it => {
      if calc.rem(it.at(0), 7) == 0 { //日曜日の場合
        table.cell(
          [
            #set text(red)
            #it.at(1)
          ]
        )
      } else if calc.rem(it.at(0), 7) == 6 { //土曜日の場合
        table.cell(
          [
            #set text(blue)
            #it.at(1)
          ]
        )
      } else { //その他の場合
        table.cell(
          [#it.at(1)]
        )
      }
    }),
  )
}

#calendar(2025, 5, 4, 31) //2025年5月のカレンダーを表示
```

実際に曜日を考慮した日付情報を生成している部分を強調している。
他はほぼテーブルの装飾のためのコードである。

これの組版結果は以下の通り。

{{< fig-img src="./calendar2.png" title="カレンダーを生成（2025年5月）" link="./calendar2.pdf" width="1005" >}}

次は1月から12月までの年間カレンダーを作ってみよう。

要は作成した `calendar` 関数を12回呼び出せばいいのだが，必要な情報をいちいち手入力するのも不毛なので `calendar.json` ファイルに外出しする。
`calendar.json` ファイルの作成は [Go] で以下のように組んでみた。

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
)

// Month represents a calendar month with its associated year, month number,
// the first weekday of the month, and the last day of the month.
// Year is the year of the month.
// Month is the month number (1-12).
// FirstWeekday is the weekday of the first day of the month (0-6, where 0 is Sunday).
// Lastday is the last day of the month.
type Month struct {
    Year         int `json:"year"`
    Month        int `json:"month"`
    FirstWeekday int `json:"first_weekday"`
    Lastday      int `json:"lastday"`
}

func main() {
    year := 2025
    months := make([]Month, 0, 12)
    for month := 1; month <= 12; month++ {
        m := Month{
            Year:         year,
            Month:        month,
            FirstWeekday: int(time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).Weekday()),
            Lastday:      time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day(),
        }
        months = append(months, m)
    }
    b, err := json.Marshal(months)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println(string(b))
}
```

このコードの実行結果は以下の通り（途中を端折っている）。

```text
$ go run months.go | jq .
[
  {
    "year": 2025,
    "month": 1,
    "first_weekday": 3,
    "lastday": 31
  },
  {
    "year": 2025,
    "month": 2,
    "first_weekday": 6,
    "lastday": 28
  },

  ...

  {
    "year": 2025,
    "month": 12,
    "first_weekday": 1,
    "lastday": 31
  }
]
```

この出力を `calendar.json` ファイルにリダイレクトすればOK。

[Typst] のコードについては `calendar` 関数を呼び出してる部分を以下のように書き換える。

```typst
#{
  let calendars = ()
  for month in json("./months.json") { //月ごとにカレンダーを作成
    calendars.push(calendar(month.year, month.month, month.first_weekday, month.lastday))
  }
  //カレンダーを3列×4行で表示
  grid(
    stroke: none,
    gutter: 0.5em,
    columns: (1fr, 1fr, 1fr),
    rows: (1fr, 1fr, 1fr, 1fr),
    ..calendars,
  )
}
```

ここでは [`table`] ではなく [`grid`] を使っている。
機能的には両者に殆ど違いはないが，ページ内をいくつか仕切って配置するという用途であれば [`grid`] を使ったほうがいいだろうか。

組版結果は以下の通り。

{{< fig-img src="./calendar3.png" title="年間カレンダーを生成（2025年）" link="./calendar3.pdf" >}}

次。
この年間カレンダーに対して祝日・休日の日に色を付けてみよう。

祝日・休日情報の収集についても [Go] で以下のコードを組んでみる。

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
    "time"

    "github.com/goark/koyomi"
    "github.com/goark/koyomi/value"
)

// Holiday represents a holiday with its date and title.
// Year is the year of the holiday.
// Month is the month of the holiday (1-12).
// Day is the day of the holiday (1-31).
// Weekday is the day of the week of the holiday.
// Title is the name or description of the holiday.
type Holiday struct {
    Year    int          `json:"year"`
    Month   int          `json:"month"`
    Day     int          `json:"day"`
    Weekday time.Weekday `json:"weekday"`
    Title   string       `json:"title"`
}

func main() {
    start, _ := value.DateFrom("2025-01-01")
    end, _ := value.DateFrom("2025-12-31")
    td, err := os.MkdirTemp(os.TempDir(), "blog")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer func() { _ = os.RemoveAll(td) }()
    k, err := koyomi.NewSource(
        koyomi.WithCalendarID(koyomi.Holiday),
        koyomi.WithStartDate(start),
        koyomi.WithEndDate(end),
        koyomi.WithTempDir(td),
    ).Get()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    holidays := make([]Holiday, 0, len(k.Events()))
    for _, e := range k.Events() {
        holidays = append(holidays, Holiday{
            Year:    e.Date.Year(),
            Month:   int(e.Date.Month()),
            Day:     e.Date.Day(),
            Weekday: e.Date.Weekday(),
            Title:   e.Title,
        })
    }
    b, err := json.Marshal(holidays)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println(string(b))
}
```

このコードの実行結果は以下の通り（途中を端折っている）。

```text
$ go run holidays.go | jq .
[
  {
    "year": 2025,
    "month": 1,
    "day": 1,
    "weekday": 3,
    "title": "元日"
  },
  {
    "year": 2025,
    "month": 1,
    "day": 13,
    "weekday": 1,
    "title": "成人の日"
  },

  ...

  {
    "year": 2025,
    "month": 11,
    "day": 24,
    "weekday": 1,
    "title": "休日"
  }
]
```

この出力を `holidays.json` ファイルにリダイレクトすればOK。

[Typst] のコードについては `calendar` 関数周りを以下のように書き換える。

```typst {hl_lines=["1-9","46-47","49-57",64,"66-74",81,"83-88",92]}
//祝日・休日の取得
#let holidays = json("./holidays.json")

//指定した年月日が祝日・休日かどうかを判定
#let containHoliday(year, month, day) = {
  holidays.find(holiday => {
    holiday.year == year and holiday.month == month and holiday.day == day
  }) != none
}

//カレンダーを作成
#let calendar(year, month, first_weekday, lastday) = {
  let days = ()
  let i = 0
  while i < first_weekday { //初日の曜日まで空白を追加
    days.push("")
    i = i + 1
  }
  days = days + range(1, lastday + 1).map(day => { //日付を追加
    [#day]
  })
  //カレンダーを作成
  table(
    stroke: (x, y) => if y == 1 {//罫線を設定
      (bottom: 0.7pt + black)
    },
    align: (x, y) => ( //文字の位置を設定
      if y > 1 { right }
      else { center }
    ),
    columns: 7, //列数を設定
    table.header( //ヘッダーを設定
      table.cell( //年月を設定
        colspan: 7,
        [
          #set text(font: "NOTO Sans CJK JP", weight: "bold")
          #year 年 #month 月
        ]
      ),
      ..(text(red)[日], [月], [火], [水], [木], [金], text(blue)[土]).map(it => { //曜日を設定
        set text(font: "NOTO Sans CJK JP", weight: "bold")
        it
      })
    ),
    ..days.enumerate(start:0).map(it => {
      let day = it.at(0)-first_weekday+1 //日付
      let hflag = day > 0 and day <= lastday and containHoliday(year, month, day) //祝日・休日かどうか
      if calc.rem(it.at(0), 7) == 0 { //日曜日の場合
        if hflag { //祝日・休日の場合
          table.cell(
            fill: red.lighten(90%),
            [
              #set text(red)
              #it.at(1)
            ]
          )
        } else { //祝日・休日でない場合
          table.cell(
            [
              #set text(red)
              #it.at(1)
            ]
          )
        }
      } else if calc.rem(it.at(0), 7) == 6 { //土曜日の場合
        if hflag { //祝日・休日の場合
          table.cell(
            fill: red.lighten(90%),
            [
              #set text(blue)
              #it.at(1)
            ]
          )
        } else { //祝日・休日でない場合
          table.cell(
            [
              #set text(blue)
              #it.at(1)
            ]
          )
        }
      } else { //その他の場合
        if hflag { //祝日・休日の場合
          table.cell(
            fill: red.lighten(90%),
            [#it.at(1)],
          )
        } else { //祝日・休日でない場合
          table.cell(
            [#it.at(1)],
          )
        }
      }
    }),
  )
}
```

組版結果は以下の通り。

{{< fig-img src="./calendar4.png" title="年間カレンダーを生成（2025年）" link="./calendar4.pdf" >}}

こんな感じで [Typst] のコードモードには得手も不得手もあるが，ある程度データを整えて与えてあげればそこそこの制御ができそうである。
ぶっちゃけ $\mathrm{\TeX}$/$\mathrm{\LaTeX}$ のマクロは触る気にもならないが [Typst] のコードモードは今どきのスクリプト言語が操れる人なら違和感少なくイケそうな気がする。

## データファイルをコマンドラインで指定する

前節でつくった年間カレンダーはデータファイル名をコードに埋め込んでいるが，これをコマンドラインで指定できるようにしてみる。

[Typst] コード側は [`sys`] を使って以下のようにコマンドラインの情報を読み込むように書き換える。

```typst {hl_lines=["2-6"]}
//祝日・休日の取得
#let hfile = "./holidays.json"
#if "holidays" in sys.inputs {
	hfile = sys.inputs.at("holidays")
}
#let holidays = json(hfile)
```

`calendar` 関数は変更がないので割愛する。

```typst {hl_lines=["2-5"]}
#{
	let months = "./months.json"
	if "months" in sys.inputs {
		months = sys.inputs.at("months")
	}
	let calendars = ()
	for month in json(months) { //月ごとにカレンダーを作成
		calendars.push(calendar(month.year, month.month, month.first_weekday, month.lastday))
	}
	//カレンダーを3列×4行で表示
	grid(
		stroke: none,
		gutter: 0.5em,
		columns: (1fr, 1fr, 1fr),
		rows: (1fr, 1fr, 1fr, 1fr),
		..calendars,
	)
}
```

一方，コマンドライン側は以下のように指定する。

```text
$ typst compile --input holidays=holidays2025.json --input months=months2025.json calendar5.typ
```

これで `holidays.json` や `months.json` ではなく `holidays2025.json` や `months2025.json` を読み込む。
`--input` オプションで指定しない場合はデフォルトのファイル名を使う。

コマンドラインで変数を指定する方法については「[変数をコマンドライン引数で指定する]({{< relref "./x-miscellaneous.md#input" >}} "Typst に関する雑多な話")」で少し詳しく紹介している。

## 余談だが

今回も [VS Code] 上で作業しているのだが，コーディングに関しては GitHub Copilot に大変お世話になっている。
[Go] のコードに関してはほぼ完璧に働いてくれるのだが， [Typst] のコードに関しては，どうも TypeScript と混乱してるっぽく，しょっちゅう嘘をついてくれるのが困りものである（笑）

データと制御が分離しやすくコードが（比較的）書きやすいというのは [Typst] の利点だと思う。
これなら業務にも組み込みやすいのではないだろうか。

## ブックマーク

ブックマークは「[Typst に関するブックマーク]({{< relref "./0-bookmark.md" >}})」にてまとめています。

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[Typst Documentation]: https://typst.app/docs/ "Typst Documentation"
[Tutorial]: https://typst.app/docs/tutorial "Tutorial – Typst Documentation"
[`raw`]: https://typst.app/docs/reference/text/raw/ "Raw Text / Code Function – Typst Documentation"
[`csv`]: https://typst.app/docs/reference/data-loading/csv/ "CSV Function – Typst Documentation"
[`array`]: https://typst.app/docs/reference/foundations/array/ "Array Type – Typst Documentation"
[`dictionary`]: https://typst.app/docs/reference/foundations/dictionary/ "Dictionary Type – Typst Documentation"
[`datetime`]: https://typst.app/docs/reference/foundations/datetime/ "Datetime Type – Typst Documentation"
[`sys`]: https://typst.app/docs/reference/foundations/sys/ "System Functions – Typst Documentation"
[`table`]: https://typst.app/docs/reference/model/table/ "Table Function – Typst Documentation"
[`grid`]: https://typst.app/docs/reference/layout/grid/ "Grid Function – Typst Documentation"

[Go]: https://go.dev/
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Tinymist Typst]: https://marketplace.visualstudio.com/items?itemName=myriad-dreamin.tinymist "Tinymist Typst - Visual Studio Marketplace"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
