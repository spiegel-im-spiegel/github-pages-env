+++
title = "今日はいつ？ — goark/today v0.1.0 をリリース"
date =  "2026-09-09T20:27:24+09:00"
description = "コマンドで today って打ったら今日の日付を西暦と和暦で表示するようにした。"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "calendar", "tools", "cli" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いまだにお役所関連の書類（の一部）は元号での記載を要求してくる。
ほんで「今日って令和何年だっけ？」ってなる。
今年ももう9月になるというのに和暦が覚えられないよ。

大抵は西暦-和暦変換表をググって確認するのだが，今回はちょっとキれた（笑）

{{< div-gen class="center" type="markdown" >}}
**やってらんねー！！！**
{{< /div-gen >}}

というわけでコマンドで `today` って打ったら今日の日付を西暦と和暦で表示するようにした。

{{< linkcard "04daa8670b109900c866a4a3343b4dc25ce04522" >}} <!-- https://github.com/goark/today goark/today: Show information for today -->

[Go] 製シングルバイナリのコマンドラインツールで Linux/Windows/macOS 用のビルド済みバイナリを[用意](https://github.com/goark/today/releases "Releases · goark/today")している。

バイナリを `PATH` が通ってるディレクトリに放り込んで，こんな感じに使う。

```text
$ today
2026-09-09 Wednesday
令和8年 (丙午) 9月9日 水曜日 (丙戌)
```

引数で日付を指定することもできる。

```text
$ today 2019-04-30
2019-04-30 Tuesday
平成31年 (己亥) 4月30日 火曜日 (丁酉)

$ today 2019-05-01
2019-05-01 Wednesday
令和1年 (己亥) 5月1日 水曜日 (戊戌)
```

オプションはこんな感じ。

```text
$ today --help
Usage of today

Usage:
  today [flags] [yyyy-mm-dd]

Flags:
  -a, --all-events          show all events information
      --debug               enable debug mode
      --event-file string   path to the event file (default "/home/username/.config/today/event.json")
  -h, --holiday             show holiday information
  -j, --json                output information in JSON format
  -e, --other-events        show other events information
  -s, --solar-term          show solar term information
      --temp-dir string     path to the temporary directory (default "/home/username/.cache/today")
  -v, --version             show version information
```

たとえば，以下のように `-h` オプションを付けて祝日を表示する。

```text
$ today -h 2026-09-23
2026-09-23 Wednesday
令和8年 (丙午) 9月23日 水曜日 (庚子)
秋分の日
```

他のオプションは以下の通り[^o1]。

[^o1]: 祝日・休日，二十四節気＋雑節，朔望月の情報は[国立天文台のデータ]({{< ref "/remark/2019/05/google-ephemeris.md" >}} "カレンダーに祝日を入れたいなら国立天文台へ行けばいいじゃない")から取得している。

|     | 意味 |
| --- | --- |
| `-h, --holiday` | 祝日・休日を表示 |
| `-s, --solar-term` | 二十四節気＋雑節および朔望月を表示 |
| `-e, --other-events` | その他のイベント情報を表示（イベントファイルで定義） |
| `-a, --all-events` | すべてのイベント情報を表示 |

イベントファイルは `--event-file` で指定できる（指定がない場合はデフォルトのファイルを開こうとする）。
イベントファイルが見つからない場合はイベント情報がないものと見なす。

たとえばイベントファイルの中身が

```json
[
  {
    "start": "2026-09-25",
    "end": "2026-09-25",
    "title": "中秋の名月"
  }
]
```

であれば

```text
$ today -e 2026-09-25
2026-09-25 Friday
令和8年 (丙午) 9月25日 金曜日 (壬寅)
中秋の名月
```

という感じに表示される。

んー。
作ったはいいけど，あまり使わんか？ でもイベントファイルは色々と使い道がありそう？ そういえば CSV や JSON などのテキストデータを SQL で操作するパッケージがあったな。

- [GitHub - nao1215/filesql: loads CSV, TSV, LTSV, JSON, JSONL, Parquet, XLSX, ACH, and Fedwire files into SQLite; includes prep and frame for cleanup and in-memory transforms](https://github.com/nao1215/filesql)
- [【Golang】CSV, TSV, LTSVをsql.DBで操作するfilesqlパッケージを作った話](https://debimate.jp/post/ja/2025-08-28-golangcsv-tsv-ltsv%E3%82%92sql-db%E3%81%A7%E6%93%8D%E4%BD%9C%E3%81%99%E3%82%8Bfilesql%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E3%82%92%E4%BD%9C%E3%81%A3%E3%81%9F%E8%A9%B1/)
- [【Golang】CSV／JSON／Excel／Parquet に SQL を実行する sqly を約4年かけて v1.0.0 にした](https://debimate.jp/post/ja/2026-08-09-golangcsvjsonexcelparquet%E3%81%ABsql%E3%82%92%E5%AE%9F%E8%A1%8C%E3%81%99%E3%82%8Bsqly%E3%82%92v1.0.0%E3%81%AB%E3%81%99%E3%82%8B%E8%A9%B1/)

面白そうだし，ちょっと検討してみようかな。

今回もあまり AI のお世話にならず（初期ファイルの配置とレビューとドキュメントは AI にやってもらった）楽しくコードを書いて遊んだ。
ちょっとは気が済んだかな（笑）

[Go]: https://go.dev/

## 参考

{{< linkcard "9e8d8f717b1d1d85b0da7f07cb10a83b78d4227f" >}} <!-- プログラミング言語Go -->
{{< review-paapi "B0CFL1DK8Q" >}} <!-- Go言語 100Tips -->
{{< review-paapi "B0DNYMMBBQ" >}} <!-- Go言語で学ぶ並行プログラミング -->
