+++
title = "カレンダーデータの解析パッケージをリリースした"
date =  "2026-11-09T15:25:38+09:00"
description = "apognu/gocal をコアとした並行処理対応版を作ってみた"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "package", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今までカレンダーデータ（iCalendar あるいは iCal フォーマット）の解析は [`PuloV/ics-golang`] を使っていた。
ただし，オリジナルは2019年を最後に更新されていないため [Go] Modules にも対応しておらず，並行処理部分で不具合が多いため自分のリポジトリに fork して使っていた。

以前から「こりゃあ自前でいちから書き直したほうが早いか？」と思っていたのだが， GitHub の他人様のリポジトリを眺めてたらソースとして [`io`]`.Reader` 型を受け付けるシンプルな [`apognu/gocal`] がよさげな感じなので，これをコアとした並行処理版を作ってみた。

- [goark/concal: Concurrent iCalendar parser wrapper for Go](https://github.com/goark/concal)

まずは単一ソースの場合：

```go
parser := concal.NewParser(
    concal.WithStartTime(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
    concal.WithEndTime(time.Date(2026, time.December, 31, 23, 59, 59, 0, time.UTC)),
)
events, err := parser.ParseOne(context.Background(), "./calendar.ics")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
for _, event := range events {
    fmt.Println(event.Summary)
}
```

[`concal`][`goark/concal`]`.NewParser()` 関数の引数は [Functional Options パターン]({{< ref "/golang/functional-options-pattern.md" >}} "インスタンスの生成と Functional Options パターン")で与える。

- [`concal`][`goark/concal`]`.WithStartTime()` はイベント取得開始時点を指定する
- [`concal`][`goark/concal`]`.WithEndTime()` はイベント取得終了時点を指定する
- [`concal`][`goark/concal`]`.WithLocationAllDayEvent()` は終日イベントを取得する際の時差を指定する
- [`concal`][`goark/concal`]`.WithLocationMap()` は時差の独自名（`MyTZ` など）を定義する（複数指定可）
- [`concal`][`goark/concal`]`.WithHTTP11Client()` は，強制的に HTTP/1.1 でアクセスする

開始・終了時点を指定しない場合は現時点から90日間を既定値としてイベントを取得する。

`ParseOne()` メソッドで指定するソースにはローカルファイルまたは Web 上のリソース（`http(s)://` で始まる URL）を指定できる（文字列から自動判定）。

解析結果は [`gocal`][`apognu/gocal`]`.Event` 型の slice で返ってくる。

次に，複数ソースを一度に処理する場合：

```go
parser := concal.NewParser(
    concal.WithStartTime(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
    concal.WithEndTime(time.Date(2026, time.December, 31, 23, 59, 59, 0, time.UTC)),
)
conc := concal.NewConcurrent(context.Background(), parser)
defer func() { _ = conc.Close() }() // TODO: Error handling is required

for _, source := range []string{
    "./calendar-a.ics",
    "https://example.com/calendar.ics",
} {
    if err := conc.Submit(source); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
go conc.Wait() // start waiting in a separate goroutine

for event := range conc.Events() {
    fmt.Println(event.Summary)
}
```

手順としては：

1. 並行処理用のインスタンスを生成
2. ソースを `Submit()` メソッドで登録（複数指定可）
3. `Wait()` メソッドを別ゴルーチンで実行して解析終了を待ち合わせる（`Wait()` 後は `Submit()` できない）
4. `Events()` からイベント出力チャネルを取得

という感じ。
Fork/Join パターンに似た構成になっているが，ソースの数が不確定なので `Wait()` メソッドのタイミングで入力を打ち切っている。
イベント出力チャネルからの出力順はソートされていないので，取り出した側で整理する必要がある。

「とりあえず組んでみた」な状態なので，バグ潰しやリファクタリングを続けてブラッシュアップしていく予定である。

[Go]: https://go.dev/
[`PuloV/ics-golang`]: https://github.com/PuloV/ics-golang "PuloV/ics-golang: Golang ics parser"
[`apognu/gocal`]: https://github.com/apognu/gocal "apognu/gocal: ICS file parser in Golang"
[`goark/concal`]: https://github.com/goark/concal "goark/concal: Concurrent iCalendar parser wrapper for Go"
[`io`]: https://pkg.go.dev/io "io package - io - Go Packages"

## 参考図書

{{< linkcard "9e8d8f717b1d1d85b0da7f07cb10a83b78d4227f" >}} <!-- プログラミング言語Go -->
{{< review-paapi "B0CFL1DK8Q" >}} <!-- Go言語 100Tips -->
{{< review-paapi "B0DNYMMBBQ" >}} <!-- Go言語で学ぶ並行プログラミング -->
