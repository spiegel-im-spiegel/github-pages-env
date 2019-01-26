+++
title = "URI エンコーディングについて"
date = "2018-02-04T00:33:35+09:00"
description = "Go 言語ではパーセント・エンコーディングの操作を標準の net/url パッケージで提供している。"
image = "/images/attention/go-code.png"
tags        = [ "golang", "programming", "web" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%BF%E3%82%93%E3%81%AA%E3%81%AEGo%E8%A8%80%E8%AA%9E%E3%80%90%E7%8F%BE%E5%A0%B4%E3%81%A7%E4%BD%BF%E3%81%88%E3%82%8B%E5%AE%9F%E8%B7%B5%E3%83%86%E3%82%AF%E3%83%8B%E3%83%83%E3%82%AF%E3%80%91-%E6%9D%BE%E6%9C%A8%E9%9B%85%E5%B9%B8/dp/477418392X?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=477418392X"><img src="https://images-fe.ssl-images-amazon.com/images/I/61EL3Dc95dL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%BF%E3%82%93%E3%81%AA%E3%81%AEGo%E8%A8%80%E8%AA%9E%E3%80%90%E7%8F%BE%E5%A0%B4%E3%81%A7%E4%BD%BF%E3%81%88%E3%82%8B%E5%AE%9F%E8%B7%B5%E3%83%86%E3%82%AF%E3%83%8B%E3%83%83%E3%82%AF%E3%80%91-%E6%9D%BE%E6%9C%A8%E9%9B%85%E5%B9%B8/dp/477418392X?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=477418392X">みんなのGo言語【現場で使える実践テクニック】</a></dt>
	<dd>松木雅幸, mattn, 藤原俊一郎, 中島大一, 牧 大輔, 鈴木健太</dd>
	<dd>稲葉貴洋 (イラスト)</dd>
    <dd>技術評論社 2016-09-09</dd>
    <dd>Book 大型本</dd>
    <dd>ASIN: 477418392X, EAN: 9784774183923</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">リファレンス本なのに索引が貧弱。これなら Kindle で買ってもよかったか。 1.7 への言及があるのはよかった。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%AB%E3%82%88%E3%82%8BWeb%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E9%96%8B%E7%99%BA-Mat-Ryer/dp/4873117526?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4873117526"><img src="https://images-fe.ssl-images-amazon.com/images/I/51UoREcNrnL._SL160_.jpg" width="125" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%AB%E3%82%88%E3%82%8BWeb%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E9%96%8B%E7%99%BA-Mat-Ryer/dp/4873117526?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4873117526">Go言語によるWebアプリケーション開発</a></dt>
	<dd>Mat Ryer</dd>
	<dd>鵜飼 文敏 (監修), 牧野 聡 (翻訳)</dd>
    <dd>オライリージャパン 2016-01-22</dd>
    <dd>Book 大型本</dd>
    <dd>ASIN: 4873117526, EAN: 9784873117522</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">日本語監訳者による解説（付録 B）が意外に役に立つ感じ。 Web アプリケーションだけでなく，サーバサイドで動く CLI アプリへの言及もある。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
