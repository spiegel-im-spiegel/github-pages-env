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
