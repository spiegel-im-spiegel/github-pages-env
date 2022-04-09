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

| Package |  Bin  | [goark][github.com/goark] | Blog  |
| ------- | :---: | :-----------------------: | :---: |
| [gpgpdump: OpenPGP packet visualizer](https://github.com/goark/gpgpdump) | {{% icons "check" %}} | {{% icons "square-check" %}} | [{{% icons "square-check" %}}]({{< ref "/release/gpgpdump.md" >}}) |
| [depm: Visualize depndency packages and modules](https://github.com/goark/depm) | {{% icons "check" %}} | {{% icons "square-check" %}} | [{{% icons "square-check" %}}]({{< ref "/release/dependency-graph-for-golang-modules.md" >}}) |
| [books-data: Search for Books Data](https://github.com/goark/books-data) | {{% icons "check" %}} | {{% icons "square-check" %}} | [{{% icons "square-check" %}}]({{< ref "/release/books-data.md" >}}) |
| [gnkf: Network Kanji Filter by Golang](https://github.com/goark/gnkf) | {{% icons "check" %}} | {{% icons "square-check" %}} | [{{% icons "square-check" %}}]({{< ref "/release/gnkf.md" >}}) |
| [gimei-cli: 姓名・住所データ生成ツール](https://github.com/goark/gimei-cli) | {{% icons "check" %}} | {{% icons "square-check" %}} |  |
| [cov19jpn: COVID-2019 in Japan; Importing Google COVID-19 Public Forecasts](https://github.com/goark/cov19jpn) | {{% icons "check" %}} | {{% icons "square-check" %}} |  |
| [ml: Make Link with Markdown Format](https://github.com/goark/ml) | {{% icons "check" %}} | {{% icons "square-check" %}} | |
| [pa-api: APIs for Amazon Product Advertising API v5 by Golang](https://github.com/goark/pa-api) | | {{% icons "square-check" %}} | [{{% icons "square-check" %}}]({{< ref "/release/pa-api-v5.md" >}}) |
| [aozora-api: APIs for Aozora-bunko RESTful Service by Golang](https://github.com/goark/aozora) | | {{% icons "square-check" %}} | [{{% icons "square-check" %}}]({{< ref "/release/aozora-api-package-for-golang.md" >}}) |
| [openbd-api: APIs for openBD by Golang](https://github.com/goark/openbd-api) | | {{% icons "square-check" %}} | [{{% icons "square-check" %}}]({{< ref "/release/openbd-api-package-for-golang.md" >}}) |
| [errs: Error handling for Golang](https://github.com/goark/errs) | | {{% icons "square-check" %}} | [{{% icons "square-check" %}}]({{< ref "/release/errs-package-for-golang.md" >}}) |
| [gocli: Minimal Packages for Command-Line Interface](https://github.com/goark/gocli) | | {{% icons "square-check" %}} | [{{% icons "square-check" %}}]({{< ref "/release/gocli-package-for-golang.md" >}}) |
| [mt: Mersenne Twister; Pseudo Random Number Generator, Implemented by Golang](https://github.com/goark/mt) | | {{% icons "square-check" %}} | [{{% icons "square-check" %}}]({{< ref "/release/mersenne-twister-by-golang.md" >}}) |
| [krconv: Convert kana-character to roman-alphabet](https://github.com/goark/krconv) | | {{% icons "square-check" %}} |  |
| [kkconv: Hiragana-Katakana Conversion](https://github.com/goark/kkconv) | | {{% icons "square-check" %}} |  |
| [go-cvss: Common Vulnerability Scoring System (CVSS)](https://github.com/goark/go-cvss) | | {{% icons "square-check" %}} |  |
| [cov19data: Importing WHO COVID-2019 Cases Global Data](https://github.com/goark/cov19data) | | {{% icons "square-check" %}} |  |
| [csvdata: Reading CSV Data](https://github.com/goark/csvdata) | | {{% icons "square-check" %}} |  |
| [fetch: Fetch Data from URL](https://github.com/goark/fetch) | | {{% icons "square-check" %}} |  |
| [koyomi: 日本のこよみ](https://github.com/goark/koyomi) | | {{% icons "square-check" %}} | |

## 移行しないパッケージ（多分）

最終的に archive にする（多分）

| パッケージ | 理由 | Archive |
| ---------- | ---- | :---: |
| [logf](https://github.com/spiegel-im-spiegel/logf) | [rs/zerolog](https://github.com/rs/zerolog) に乗り換えた | {{% icons "square-check" %}} |
| [writers](https://github.com/spiegel-im-spiegel/writers) | 全く使ってない | {{% icons "square-check" %}} |
| [gcavoc](https://github.com/spiegel-im-spiegel/gcavoc) | 全く使ってない | {{% icons "square-check" %}} |
| [gprompt](https://github.com/spiegel-im-spiegel/gprompt) | [nyaosorg/go-readline-ny](https://github.com/nyaosorg/go-readline-ny) に乗り換えた | {{% icons "square-check" %}} |
| [file](https://github.com/spiegel-im-spiegel/file) | 全く使ってない | {{% icons "square-check" %}} |
| [jpera](https://github.com/spiegel-im-spiegel/jpera) | [koyomi](https://github.com/goark/koyomi) パッケージに移行・統合した | {{% icons "square-check" %}} |
| [jzodiac](https://github.com/spiegel-im-spiegel/jzodiac) | [koyomi](https://github.com/goark/koyomi) パッケージに移行・統合した | {{% icons "square-check" %}} |
| [gjq](https://github.com/spiegel-im-spiegel/gjq) | 全く使ってない | {{% icons "square-check" %}} |
| [jvnman](https://github.com/spiegel-im-spiegel/jvnman) | 全く使ってない | {{% icons "square-check" %}} |
| [go-myjvn](https://github.com/spiegel-im-spiegel/go-myjvn) | 全く使ってない | {{% icons "square-check" %}} |
| [godump](https://github.com/spiegel-im-spiegel/godump) | [gnkf](https://github.com/goark/gnkf) に組込み済み | {{% icons "square-check" %}} |
| [gocodic](https://github.com/spiegel-im-spiegel/gocodic) | 全く使ってない | {{% icons "square-check" %}} |
| [icat4json](https://github.com/spiegel-im-spiegel/icat4json) | 全く使ってない | {{% icons "square-check" %}} |
| [gitioapi](https://github.com/spiegel-im-spiegel/gitioapi) | 全く使ってない | {{% icons "square-check" %}} |
| [xls2csv](https://github.com/spiegel-im-spiegel/xls2csv) | 外部パッケージのパスとバージョンを更新したものをリリースして凍結 | {{% icons "square-check" %}} |
| [emojis](https://github.com/spiegel-im-spiegel/emojis) | 全く使ってない | {{% icons "square" %}} |

## その他の TODO

{{% icons "square-check" %}} [github.com/spiegel-im-spiegel](https://github.com/spiegel-im-spiegel) の Readme を修正する

{{% icons "square-check" %}} Zenn の「[Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang)」の改訂

## ブックマーク

- [独り GitHub Team を作ってみた]({{< ref "/remark/2022/03/github-team.md" >}})

[github.com/goark]: https://github.com/goark "Playing with Go Language"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"
[depm]: https://github.com/goark/depm "goark/depm: Visualize depndency packages and modules"
[Go]: https://go.dev/

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
