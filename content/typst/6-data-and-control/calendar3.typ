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
					[#it.at(1)],
				)
			}
		}),
	)
}

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
