+++
title = "Typst におけるデータと制御"
date =  "2025-02-28T17:27:26+09:00"
description = "description"
image = "/images/attention/tools.png"
tags  = [ "typst", "pdf" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
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

フォントの指定については今はスルーで[^r1]（笑） CSV データの読み込みには [`raw`] 関数を使う。
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


## ブックマーク

ブックマークは「[Typst に関するブックマーク]({{< relref "./0-bookmark.md" >}})」にてまとめています。

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[Typst Documentation]: https://typst.app/docs/ "Typst Documentation"
[Tutorial]: https://typst.app/docs/tutorial "Tutorial – Typst Documentation"
[`raw`]: https://typst.app/docs/reference/text/raw/ "Raw Text / Code Function – Typst Documentation"
[`csv`]: https://typst.app/docs/reference/data-loading/csv/ "CSV Function – Typst Documentation"
[`array`]: https://typst.app/docs/reference/foundations/array/ "Array Type – Typst Documentation"
[`dictionary`]: https://typst.app/docs/reference/foundations/dictionary/ "Dictionary Type – Typst Documentation"
[`table`]: https://typst.app/docs/reference/model/table/ "Table Function – Typst Documentation"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
