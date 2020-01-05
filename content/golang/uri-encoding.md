+++
title = "URI エンコーディングについて"
date = "2018-02-04T00:33:35+09:00"
description = "Go 言語ではパーセント・エンコーディングの操作を標準の net/url パッケージで提供している。"
image = "/images/attention/go-logo_blue.png"
tags        = [ "golang", "programming", "web" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

今回は軽く小ネタで。

「パーセント・エンコーディング（percent-encoding）」というのがある。
これは文字符号を “`%xx`” 形式でオクテット単位の16進数コードに展開して記述することを指して言うらしい。
たとえば「こんにちは」は UTF-8 で「`%E3%81%93%E3%82%93%E3%81%AB%E3%81%A1%E3%81%AF`」と表記する。

パーセント・エンコーディングは Web などの URI 表記で使われる。
詳しいことは [RFC 3986] 辺りを参照してもらうとして， [Go 言語]ではこれを標準の `net/`[`url`] パッケージで提供している。
ただし以下の2種類ある。

- [`url`]`.QueryEscape()` / [`url`]`.QueryUnescape()`
- [`url`]`.PathEscape()` / [`url`]`.PathUnescape()` [^go18]

[^go18]: [`url`]`.PathEscape()` / [`url`]`.PathUnescape()` は [Go 1.8 から追加](https://golang.org/doc/go1.8#net_url "Go 1.8 Release Notes - The Go Programming Language")された。

どう違うかは実際にコードを組んでみたほうが早いだろう。

まずは [`url`]`.QueryEscape()` / [`url`]`.QueryUnescape()` のほうから

```go
package main

import (
    "fmt"
    "net/url"
    "os"
)

func main() {
    str0 := "こんにちは 世界"

    str1 := url.QueryEscape(str0)
    fmt.Println(str1)
    // %E3%81%93%E3%82%93%E3%81%AB%E3%81%A1%E3%81%AF+%E4%B8%96%E7%95%8C

    str2, err := url.QueryUnescape(str1)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println(str2)
    // こんにちは 世界
}
```

次に [`url`]`.PathEscape()` / [`url`]`.PathUnescape()` のほう

```go
package main

import (
    "fmt"
    "net/url"
    "os"
)

func main() {
    str0 := "こんにちは 世界"

    str1 := url.PathEscape(str0)
    fmt.Println(str1)
    // %E3%81%93%E3%82%93%E3%81%AB%E3%81%A1%E3%81%AF%20%E4%B8%96%E7%95%8C

    str2, err := url.PathUnescape(str1)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println(str2)
    // こんにちは 世界
}
```

パッと見で分かりやすいのは空白文字の扱いだろう。
両者の違いは関数の名前の通り URI の query 部分（part）で使うか path 部分で使うかの違いである。

ちなみに URI を

```html
foo://example.com:8042/over/there?name=ferret#nose
```

とするなら，構成要素（component）は

| 部分名    | 部分               |
| --------- | ------------------ |
| scheme    | `foo`              |
| authority | `example.com:8042` |
| path      | `/over/there`      |
| query     | `name=ferret`      |
| fragment  | `nose`             |

という感じに分類される。
上手く使い分けよう。

## ブックマーク

- [Golangでパーセントエンコーディング - 逆さまにした](http://cipepser.hatenablog.com/entry/2017/07/29/083729)
    - [Golangでパーセントエンコーディングをデコードする - 逆さまにした](http://cipepser.hatenablog.com/entry/2017/08/05/095807)
- [encodeURIComponentが世界基準だと誤解してた話 - Qiita](https://qiita.com/shibukawa/items/c0730092371c0e243f62)
- [URLエンコードについておさらいしてみた - Qiita](https://qiita.com/sisisin/items/3efeb9420cf77a48135d)
- [Big Sky :: Golang で物理ファイルの操作に path/filepath でなく path を使うと爆発します。](https://mattn.kaoriya.net/software/lang/go/20171024130616.htm) : URL の操作には `http.ServeFile` を使うとかあるらしい

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`url`]: https://golang.org/pkg/net/url/ "url - The Go Programming Language"
[RFC 3986]: https://tools.ietf.org/html/rfc3986 "RFC 3986 - Uniform Resource Identifier (URI): Generic Syntax"

## 参考図書

{{% review-paapi "B07VPSXF6N" %}} <!-- 改訂2版 みんなのGo言語 -->

{{% review-paapi "4873117526" %}} <!-- Go言語によるWebアプリケーション開発 -->
