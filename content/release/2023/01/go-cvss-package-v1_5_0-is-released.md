+++
title = "go-cvss パッケージ v1.5.0 をリリースした"
date =  "2023-01-31T16:36:54+09:00"
description = " 一連の変更でようやく CVSSv2 処理系は v3/metric パッケージと互換になった。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "package", "module", "golang", "cvss" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

バグ報告が止まらない。
ごめんなさい
ごめんなさい
ごめんなさい...

ということで [`github.com/goark/go-cvss`] パッケージの v1.5.0 をリリースした。

- [Release v1.5.0 · goark/go-cvss · GitHub](https://github.com/goark/go-cvss/releases/tag/v1.5.0)

このバージョンで CVSSv2 のデコードを行うサブパッケージを `v2/base` から `v2/metric` に移行した。
こんな感じ。

```go
//go:build run
// +build run

package main

import (
    "fmt"
    "os"

    "github.com/goark/go-cvss/v2/metric"
)

func main() {
    bm, err := metric.NewBase().Decode("AV:N/AC:L/Au:N/C:N/I:N/A:C") //CVE-2002-0392
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Printf("Severity: %v (%v)\n", bm.Severity(), bm.Score())
    // Output:
    // Severity: Severity: High (7.8)
}
```

これに伴い `v2/base` パッケージ以下の型定義およびメソッドには `Deprecated` マークを付けている。
一連の変更でようやく CVSSv2 処理系は `v3/metric` パッケージと互換になった。
まぁ今どき CVSSv2 を使うことはないと思うけどね。

そもそも [`github.com/goark/go-cvss`] パッケージ自体がかなり試行錯誤していて，随分とカオスな状態になっていたので，この機会に整理できてよかった。
もうバグはないよなぁ。
ないと思いたい。

## 2022-01-31 追記

その後のバグ報告で，上述の `Deprecated` マークを付けた `v2/base` パッケージはデータ構造上 CVSSv2 の仕様を満たせないことが分かったので v1.6.0 で drop した。

- [Release v1.6.0 · goark/go-cvss · GitHub](https://github.com/goark/go-cvss/releases/tag/v1.6.0)

とほほ `orz`

## ブックマーク

- [CVSS v2 Complete Documentation](https://www.first.org/cvss/v2/guide)

[`github.com/goark/go-cvss`]: https://github.com/goark/go-cvss "goark/go-cvss: Common Vulnerability Scoring System (CVSS)"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
