+++
title = "Go パッケージの移行状況"
date =  "2022-03-19T21:15:27+09:00"
description = "ダラダラやってるといつまで経っても終わらないので，移行状況を記しておく"
image = "/images/attention/go-logo_blue.png"
tags = [ "github", "golang", "package", "module" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

現在 [github.com/goark] に [Go] パッケージを移行中だが，ダラダラやってるといつまで経っても終わらないので，覚え書きとして移行状況を記しておく。
この記事は移行が完了するまで随時更新される。

| Package | Bin | [goark][github.com/goark] | Usage |
| ---------- | :------: | :---: | :---: |
| [gpgpdump: OpenPGP packet visualizer](https://github.com/goark/gpgpdump) | {{% icons "check" %}} | {{% icons "square-check" %}} | [{{% icons "square" %}}]({{< ref "release/gpgpdump.md" >}}) |
| [depm: Visualize depndency packages and modules](https://github.com/goark/depm) | {{% icons "check" %}} | {{% icons "square-check" %}} | [{{% icons "square" %}}]({{< ref "release/dependency-graph-for-golang-modules.md" >}}) |
| [books-data: Search for Books Data](https://github.com/spiegel-im-spiegel/books-data) | {{% icons "check" %}} | {{% icons "square" %}} | [{{% icons "square" %}}]({{< ref "release/books-data.md" >}}) |
| [gnkf: Network Kanji Filter by Golang](https://github.com/spiegel-im-spiegel/gnkf) | {{% icons "check" %}} | {{% icons "square" %}} | [{{% icons "square" %}}]({{< ref "release/gnkf.md" >}}) |
| [gimei-cli: 姓名・住所データ生成ツール](https://github.com/goark/gimei-cli) | {{% icons "check" %}} | {{% icons "square-check" %}} |  |
| [cov19jpn: COVID-2019 in Japan; Importing Google COVID-19 Public Forecasts](https://github.com/goark/cov19jpn) | {{% icons "check" %}} | {{% icons "square-check" %}} |  |
| [xls2csv: xport CSV Text from Excel Data](https://github.com/spiegel-im-spiegel/xls2csv) | {{% icons "check" %}} | {{% icons "square" %}} | |
| [ml: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/ml) | {{% icons "check" %}} | {{% icons "square" %}} | |
| [pa-api: APIs for Amazon Product Advertising API v5 by Golang](https://github.com/goark/pa-api) | | {{% icons "square-check" %}} | [{{% icons "square" %}}]({{< ref "release/pa-api-v5.md" >}}) |
| [aozora-api: APIs for Aozora-bunko RESTful Service by Golang](https://github.com/spiegel-im-spiegel/aozora-api) | | {{% icons "square" %}} | [{{% icons "square" %}}]({{< ref "release/aozora-api-package-for-golang.md" >}}) |
| [openbd-api: APIs for openBD by Golang](https://github.com/spiegel-im-spiegel/openbd-api) | | {{% icons "square" %}} | [{{% icons "square" %}}]({{< ref "release/openbd-api-package-for-golang.md" >}}) |
| [errs: Error handling for Golang](https://github.com/goark/errs) | | {{% icons "square-check" %}} | [{{% icons "square" %}}]({{< ref "release/errs-package-for-golang.md" >}}) |
| [gocli: Minimal Packages for Command-Line Interface](https://github.com/goark/gocli) | | {{% icons "square-check" %}} | [{{% icons "square" %}}]({{< ref "release/gocli-package-for-golang.md" >}}) |
| [mt: Mersenne Twister; Pseudo Random Number Generator, Implemented by Golang](https://github.com/goark/mt) | | {{% icons "square-check" %}} | [{{% icons "square" %}}]({{< ref "release/mersenne-twister-by-golang.md" >}}) |
| [krconv: Convert kana-character to roman-alphabet](https://github.com/goark/krconv) | | {{% icons "square-check" %}} |  |
| [kkconv: Hiragana-Katakana Conversion](https://github.com/goark/kkconv) | | {{% icons "square-check" %}} |  |
| [go-cvss: Common Vulnerability Scoring System (CVSS)](https://github.com/goark/go-cvss) | | {{% icons "square-check" %}} |  |
| [cov19data: Importing WHO COVID-2019 Cases Global Data](https://github.com/goark/cov19data) | | {{% icons "square-check" %}} |  |
| [csvdata: Reading CSV Data](https://github.com/goark/csvdata) | | {{% icons "square-check" %}} |  |
| [fetch: Fetch Data from URL](https://github.com/goark/fetch) | | {{% icons "square-check" %}} |  |
| [koyomi: 日本のこよみ](https://github.com/spiegel-im-spiegel/koyomi) | | {{% icons "square" %}} | |
| [jzodiac: Japanese Zodiac](https://github.com/spiegel-im-spiegel/jzodiac) | | {{% icons "square" %}} | |

## 移行しないパッケージ（多分）

最終的に archive にする（多分）

| パッケージ | 理由 | Archive |
| ---------- | ---- | :---: |
| [logf](https://github.com/spiegel-im-spiegel/logf) | [rs/zerolog](https://github.com/rs/zerolog) に乗り換えたため | {{% icons "square" %}} |
| [writers](https://github.com/spiegel-im-spiegel/writers) | 全く使ってない | {{% icons "square" %}} |
| [gcavoc](https://github.com/spiegel-im-spiegel/gcavoc) | 全く使ってない | {{% icons "square" %}} |
| [emojis](https://github.com/spiegel-im-spiegel/emojis) | 全く使ってない | {{% icons "square" %}} |
| [gprompt](https://github.com/spiegel-im-spiegel/gprompt) | [nyaosorg/go-readline-ny](https://github.com/nyaosorg/go-readline-ny) に乗り換えたため | {{% icons "square" %}} |
| [file](https://github.com/spiegel-im-spiegel/file) | 全く使ってない | {{% icons "square" %}} |
| [jpera](https://github.com/spiegel-im-spiegel/jpera) | 全く使ってない | {{% icons "square" %}} |
| [gjq](https://github.com/spiegel-im-spiegel/gjq) | 全く使ってない | {{% icons "square" %}} |
| [jvnman](https://github.com/spiegel-im-spiegel/jvnman) | 全く使ってない | {{% icons "square" %}} |
| [go-myjvn](https://github.com/spiegel-im-spiegel/go-myjvn) | 全く使ってない | {{% icons "square" %}} |
| [godump](https://github.com/spiegel-im-spiegel/godump) | [gnkf](https://github.com/spiegel-im-spiegel/gnkf) に組込み済み | {{% icons "square" %}} |
| [gocodic](https://github.com/spiegel-im-spiegel/gocodic) | 全く使ってない | {{% icons "square" %}} |
| [icat4json](https://github.com/spiegel-im-spiegel/icat4json) | 全く使ってない | {{% icons "square" %}} |
| [gitioapi](https://github.com/spiegel-im-spiegel/gitioapi) | 全く使ってない | {{% icons "square" %}} |

## その他の TODO

{{% icons "square" %}} [github.com/spiegel-im-spiegel](https://github.com/spiegel-im-spiegel) の Readme を修正する

{{% icons "square" %}} Zenn の「[Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang)」の改訂

## ブックマーク

- [独り GitHub Team を作ってみた]({{< ref "/remark/2022/03/github-team.md" >}})

[github.com/goark]: https://github.com/goark "Playing with Go Language"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"
[depm]: https://github.com/goark/depm "goark/depm: Visualize depndency packages and modules"
[Go]: https://go.dev/

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
