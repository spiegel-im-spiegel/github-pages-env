+++
title = "gjq v0.2.0 をリリースした"
date = "2019-03-17T21:34:01+09:00"
description = "2つほどオプションを追加した。"
image = "/images/attention/tools.png"
tags  = [ "tools", "json", "golang", "cli" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[昨日]の今日で恐縮だが [spiegel-im-spiegel/gjq] v0.2.0 をリリースした。

- [Release v0.2.0 · spiegel-im-spiegel/gjq · GitHub](https://github.com/spiegel-im-spiegel/gjq/releases/tag/v0.2.0)

2つほどオプションを追加した。

{{< highlight text "hl_lines=6-7" >}}
$ gjq -h
Usage:
  gjq [flags] <filter string>

Flags:
  -c, --clipboard     copy to clipboard in interactive mode
  -C, --color         colorize JSON
      --debug         for debug
  -f, --file string   JSON data (file path)
  -h, --help          help for gjq
  -i, --indent int    indent size for formatted JSON string (default 2)
  -I, --interactive   interactive mode
  -t, --tab           use tabs for indentation
  -u, --url string    JSON data (URL)
  -v, --version       output version of gjq
{{< /highlight >}}

小文字の `-c` をつけると対話モード時にフィルタリング結果をクリップボードにコピーする。
逆にこのオプションを付けないとクリップボードへのコピーは行わないようにした。
対話モードは `Ctrl+C` を入力するまで続くが，この割り込みタイミングで直前の結果をクリップボードにコピーする。

大文字の `-C` オプションでフィルタリング結果をカラー表示できるようにした。
ファイルやクリップボードへの出力には影響しない（筈）。

JSON 文字列のカラー化には以下の2つのパッケージを利用している。

- [fatih/color: Color package for Go (golang)](https://github.com/fatih/color)
- [nwidger/jsoncolor: Colorized JSON output for Go https://godoc.org/github.com/nwidger/jsoncolor](https://github.com/nwidger/jsoncolor)

いやぁ，おかげさまで実装がめっさ簡単だったわ（笑）

あと Windows とそれ以外のプラットフォームとの差を吸収するために

- [mattn/go-colorable](https://github.com/mattn/go-colorable)

を利用している。
たとえば

```go
fmt.Fprintln(colorable.NewColorableStdout(), string(coloredJSONbytes))
```

とかすればプラットフォームに関係なく [ANSI 標準のエスケープシーケンス](https://en.wikipedia.org/wiki/ANSI_escape_code "ANSI escape code - Wikipedia")を使って文字に色を付けることができる。

これで「jq ぽい何か」でやりたいことはだいたい実装できたので本来の作業へ戻るとしよう。

## ブックマーク

- [jq ぽい何か を書いてみた]({{< relref "./gjq.md" >}})

[昨日]: {{< relref "./gjq.md" >}} "jq ぽい何か を書いてみた"
[spiegel-im-spiegel/gjq]: https://github.com/spiegel-im-spiegel/gjq "spiegel-im-spiegel/gjq: Another Implementation of "jq" by golang"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
