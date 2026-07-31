+++
title = "Go 1.12 がリリースされた"
date = "2019-02-26T23:03:40+09:00"
description = "主な変更点としては TLS 1.3 の暫定的なサポートとモジュール・モードの挙動の一部が変わったことだろうか。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "engineering", "module", "tls", "regular-expression", "concurrency" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Go 1.12 がリリースされた。
そういえば2月ももう終わりか。

- [Go 1.12 is released - The Go Blog](https://blog.golang.org/go1.12)
- [Go 1.12 Release Notes - The Go Programming Language](https://golang.org/doc/go1.12)

主な変更点としては TLS 1.3 の暫定的なサポート（有効にするには環境変数の設定が必要）と[モジュール・モード]({{< ref "/golang/go-module-aware-mode.md" >}} "モジュール対応モードへの移行を検討する")の挙動の一部が変わったことだろうか。

たとえば環境変数 `$GO111MODULE` を `on` にしておけば [mattn](https://github.com/mattn) さんの [jvgrep](https://github.com/mattn/jvgrep "mattn/jvgrep: grep for japanese vimmer") をインストールする際にも任意のフォルダで

```text
$ go get github.com/mattn/jvgrep@latest
```

とすればよい。
ダミーの `go.mod` ファイルを作る必要はなくなった。
ブラボー！

Go 1.13 からはモジュール・モードが既定になるようで，「GOPATH モードじゃないと困る」とかじゃなければ環境変数 `$GO111MODULE` は `on` のままにしてしまえばいいんじゃないのかなぁ。

あと地味だが嬉しい変更としては，並行処理下で正規表現パッケージ [`regexp`] を使う際に [`regexp`]`.Regexp.Copy()` でクローンを作らなくてもブロッキングが起きないようになった。

[Go 言語] はこの 1.12 から 1.13 にかけて大きく変わる予感がする（今のところ後方互換性は確保されるだろうが）。
色々と試していって慣れていくのがいいかもしれない。

## ブックマーク

- [モジュール対応モードへの移行を検討する]({{< ref "/golang/go-module-aware-mode.md" >}})
- [正規表現に関する戯れ言]({{< ref "/golang/regular-expression.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`regexp`]: https://golang.org/pkg/regexp/ "regexp - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
