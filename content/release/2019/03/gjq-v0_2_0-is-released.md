+++
title = "gjq v0.2.0 をリリースした"
date = "2019-03-17T21:34:01+09:00"
description = "description"
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

`-C` オプションでフィルタリング結果をカラー表示できるようにした。
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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
