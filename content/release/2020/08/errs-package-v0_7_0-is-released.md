+++
title = "Go 言語用エラーハンドリング・パッケージ errs v0.7.0 をリリースした"
date =  "2020-08-08T10:03:25+09:00"
description = "Go 1.15 のリリースを待って，問題がなければこのまま v1.0 に格上げする予定である。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "package", "module", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Go 言語用エラーハンドリング・パッケージ [`errs`] の v0.7.0 をリリースした。

- [Release v0.7.0 · spiegel-im-spiegel/errs · GitHub](https://github.com/spiegel-im-spiegel/errs/releases/tag/v0.7.0)

「後方互換性？ なにそれおいしいの？」って感じでド派手に書き換えている。
使い方は以下を参照のこと。

- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})

改めて紹介すると， [spiegel-im-spiegel/errs][`errs`] は以下の設計コンセプトで作成・公開した [Go 言語][Go]用エラーハンドリング・パッケージである。

- [Go] 1.13 以降の標準パッケージ `errors` と置き換え可能
- 任意の `error` インスタンスをラッピング可能
- 任意の `error` インスタンスを原因エラーとして埋め込み可能
- 任意のコンテキスト情報を埋め込み可能
    - 既定でエラーが発生した関数名をコンテキスト情報として保持する
- 構造化されたエラー情報を JSON 形式で出力可能
    - `MarshalJSON()` メソッド完備
    - 書式 `%+v` を使って JSON 形式で出力
    - 任意の `error` インスタンスで可能な限り構造を辿って出力

v0.7.0 で個人的に欲しい機能は全て搭載したので， [Go] 1.15 の正式リリースを待って，問題がなければこのまま v1.0 に格上げする予定である。

## ブックマーク

- [Go 1.13 のエラー・ハンドリング]({{< ref "/golang/error-handling-in-go-1_3.md" >}})
- [書式 %v のカスタマイズ]({{< ref "/golang/formatting.md" >}})
- [構造化エラーをログ出力する]({{< ref "/golang/logging-error.md" >}})

[Go]: https://go.dev/
[`errs`]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
