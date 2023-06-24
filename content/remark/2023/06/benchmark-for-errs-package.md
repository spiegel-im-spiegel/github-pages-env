+++
title = "拙作 gorak/errs パッケージの性能評価をしてもらった"
date =  "2023-06-24T10:06:56+09:00"
description = "ありがとうございます 🙇"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "package", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

こんなマイナーなパッケージの性能評価をしていただいてありがとうございます {{< emoji "ペコン" >}} いや，マジで。

- [次なる`pkg/errors`を探して - カンムテックブログ](https://tech.kanmu.co.jp/entry/2023/06/19/150000)
- [次なる pkg/errors を探してを読んで - 薄いブログ](https://orisano.hatenablog.com/entry/2023/06/21/231349)

[`pkg/errors`] は昔から人気の高いエラーハンドリング・パッケージで，私も随分お世話になった。
このパッケージの更新が止まって read-only になったのに伴い代替となるパッケージが望まれていたのは知っている。
で，登場したのが [`cockroachdb/errors`] パッケージなわけだ。

[`cockroachdb/errors`] パッケージは，おそらく CockroachDB などのデータベース操作に向いたエラーハンドリング・パッケージと思われ， [`pkg/errors`] との互換性を維持したまま PII (Personally Identifiable Information) のマスキングもできる優れものである。
[`pkg/errors`] からの乗り換えを考えるなら [`cockroachdb/errors`] パッケージはアリな選択だと思うし個人的にもお勧めである。

一方で拙作の [`goark/errs`] はもう少し違うところを目指していて

- 任意の `error` インスタンスをラッピングすることに主眼を置く
  - 任意の error インスタンスを原因エラーとして埋め込み可能
- 任意のコンテキスト情報を埋め込み可能
  - 既定でエラーが発生した関数名をコンテキスト情報として保持する
- 構造化されたエラー情報を JSON 形式で出力可能
  - MarshalJSON() メソッド完備
  - 書式 `%+v` を使って JSON 形式で出力
  - 任意の `error` インスタンスで（`Unwrap` メソッドの挙動に従い）可能な限り構造を辿って出力

といった機能を有している。
もちろんこれは [`pkg/errors`] パッケージに対するささやかな不満から来ている。

私は[「スタック情報は9割以上がノイズ」「藁束の中から金の針を探すようなもの」](https://zenn.dev/spiegel/books/error-handling-in-golang/viewer/error-logging "ぼくがかんがえたさいきょうのえらーろぐ｜Go のエラーハンドリング")という危険思想の持ち主なので，どういう形であれスタック情報を丸ごとどーんと出力することはしないようにしている。
他人様が書いた Java コードのデバッグでカジュアルにスタックトレースを吐き出しやがる（しかもそれを見ても結局分からずデバッガを動かす羽目になる）のに辟易してたというのもある。

それならスタック情報はエラーを吐き出した関数名を保持するのみとし，あとはエラーに至る「文脈（context）」をできる限りかき集めてエラー・インスタンスに突っ込むという戦略のほうが幾らかマシだろう，と考えたのだ。
どうしても関数呼び出しの構造が欲しければ，エラーを検出した時点で都度ラッピングしていけばいいという考え方である。

ところで最初に挙げた記事では

{{< fig-quote type="markdown" link="https://orisano.hatenablog.com/entry/2023/06/21/231349" title="次なる pkg/errors を探してを読んで" >}}
出力の処理は `json.Marshal` と `fmt.Sprintf` を使っています。 json.Marshal を高速化するために [`goccy/go-json`](https://github.com/goccy/go-json) に変えるのもありかもしれません。
{{< /fig-quote >}}

と評価をいただいていて， JSON の marshalling については（一瞬心が揺らいだが）性能がよくてもサードパーティのパッケージには依存したくないというのがあるので，パスさせていただくが， [`fmt`]`.Sprintf` に関しては正直に言って実装をサボっているだけなので，少し改善してみることにした。

まずは [miyataka/benchmark_pkg_errors_alternatives](https://github.com/miyataka/benchmark_pkg_errors_alternatives) を拝借して改めてベンチマークをとってみる。
他のパッケージと比べても仕方がないので [`goark/errs`] を使った結果のみ示すと

| benchmark                                           | ns/op | B/op | allocs/op |
| --------------------------------------------------- | ----: | ---: | --------: |
| BenchmarkErrors/goark/errs-stack-10-12              |  2746 |  648 |         7 |
| BenchmarkErrors/goark/errs-stack-100-12             |  3278 |  648 |         7 |
| BenchmarkErrors/goark/errs-stack-1000-12            |  6810 |  648 |         7 |
| BenchmarkStackFormatting/goark/errs-%s-stack-10-12  | 167.3 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%v-stack-10-12  | 185.0 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%+v-stack-10-12 |  8680 | 1401 |        33 |
| BenchmarkStackFormatting/goark/errs-%s-stack-30-12  | 174.8 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%v-stack-30-12  | 180.4 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%+v-stack-30-12 |  8826 | 1401 |        33 |
| BenchmarkStackFormatting/goark/errs-%s-stack-60-12  | 170.0 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%v-stack-60-12  | 160.5 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%+v-stack-60-12 |  8636 | 1401 |        33 |

うっ，アロケート回数が33回とか `orz`

凹みつつも JSON データ生成部分でなるべく [`fmt`]`.Sprintf` を使わないようにした v1.2.3 をリリースした。

- [Release v1.2.3 · goark/errs · GitHub](https://github.com/goark/errs/releases/tag/v1.2.3)

これを使って同じ条件でベンチマークをとってみたのだが

| benchmark                                           | ns/op | B/op | allocs/op |
| --------------------------------------------------- | ----: | ---: | --------: |
| BenchmarkErrors/goark/errs-stack-10-12              |  2850 |  648 |         7 |
| BenchmarkErrors/goark/errs-stack-100-12             |  3344 |  648 |         7 |
| BenchmarkErrors/goark/errs-stack-1000-12            |  6365 |  648 |         7 |
| BenchmarkStackFormatting/goark/errs-%s-stack-10-12  | 167.7 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%v-stack-10-12  | 164.6 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%+v-stack-10-12 |  7098 | 1385 |        31 |
| BenchmarkStackFormatting/goark/errs-%s-stack-30-12  | 171.8 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%v-stack-30-12  | 171.5 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%+v-stack-30-12 |  6974 | 1385 |        31 |
| BenchmarkStackFormatting/goark/errs-%s-stack-60-12  | 173.9 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%v-stack-60-12  | 164.8 |    8 |         1 |
| BenchmarkStackFormatting/goark/errs-%+v-stack-60-12 |  7097 | 1385 |        31 |

ちょっとしか変わらん `orz` やっぱ [`json`][`encoding/json`]`.Marshal` を使ってるのがあかんのか？ そもそも改善になってない？

...というわけで諦めました。
こんなのでよろしければ使ってやってください。

そうそう [`errors`]`.Join` 互換の関数ってあったほうがいいのかなぁ。
それをするにはマルチエラー用の型を作らないといけないのだが... これはちょっと考えてみてもいいかも。

## ブックマーク

- [Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang)
- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})
- [Zap と go-log を試す]({{< ref "/golang/zap-and-golog.md" >}})

[`pkg/errors`]: https://github.com/pkg/errors "pkg/errors: Simple error handling primitives"
[`cockroachdb/errors`]: https://github.com/cockroachdb/errors "cockroachdb/errors: Go error library with error portability over the network"
[`goark/errs`]: https://github.com/goark/errs "goark/errs: Error handling for Golang"
[`fmt`]: https://pkg.go.dev/fmt "fmt package - fmt - Go Packages"
[`encoding/json`]: https://pkg.go.dev/encoding/json "json package - encoding/json - Go Packages"
[`errors`]: https://pkg.go.dev/errors "errors package - errors - Go Packages"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
