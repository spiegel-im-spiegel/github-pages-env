+++
title = "URI エンコーディングについて"
date = "2018-02-04T00:33:35+09:00"
description = "Go 言語ではパーセント・エンコーディングの操作を標準の net/url パッケージで提供している。"
image = "/images/attention/go-code.png"
tags        = [ "golang", "programming", "web" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
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

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/477418392X/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/61EL3Dc95dL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/477418392X/baldandersinf-22/">みんなのGo言語【現場で使える実践テクニック】</a></dt><dd>松木雅幸 mattn 藤原俊一郎 中島大一 牧 大輔 鈴木健太 稲葉貴洋 </dd><dd>技術評論社 2016-09-09</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774184322/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774184322.09._SCTHUMBZZZ_.jpg"  alt="WEB+DB PRESS Vol.95"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621300253.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274219151/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274219151.09._SCTHUMBZZZ_.jpg"  alt="プログラミングElixir"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798147400/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798147400.09._SCTHUMBZZZ_.jpg"  alt="詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド (NEXT ONE)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774182869/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774182869.09._SCTHUMBZZZ_.jpg"  alt="WEB+DB PRESS Vol.94"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117763/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117763.09._SCTHUMBZZZ_.jpg"  alt="Docker"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/477418361X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/477418361X.09._SCTHUMBZZZ_.jpg"  alt="オブジェクト指向設計実践ガイド ~Rubyでわかる 進化しつづける柔軟なアプリケーションの育て方"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4295000337/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4295000337.09._SCTHUMBZZZ_.jpg"  alt="WebデベロッパーのためのReact開発入門 JavaScript UIライブラリの基本と活用"  /></a> </p>
<p class="description">リファレンス本なのに索引が貧弱。これなら Kindle で買ってもよかったか。 1.7 への言及があるのはよかった。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-11-17">2016-11-17</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51UoREcNrnL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/">Go言語によるWebアプリケーション開発</a></dt><dd>Mat Ryer 鵜飼 文敏 </dd><dd>オライリージャパン 2016-01-22</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621300253.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語Go"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117607/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117607.09._SCTHUMBZZZ_.jpg"  alt="マイクロサービスアーキテクチャ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774178667/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774178667.09._SCTHUMBZZZ_.jpg"  alt="nginx実践入門 (WEB+DB PRESS plus)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4863541783/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4863541783.09._SCTHUMBZZZ_.jpg"  alt="改訂2版 基礎からわかる Go言語"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774179930/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774179930.09._SCTHUMBZZZ_.jpg"  alt="サーバ/インフラエンジニア養成読本 DevOps編 [Infrastructure as Code を実践するノウハウが満載! ] (Software Design plus)"  /></a> </p>
<p class="description">日本語監訳者による解説（付録 B）が意外に役に立つ感じ。 Web アプリケーションだけでなく，サーバサイドで動く CLI アプリへの言及もある。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-03-13">2016-03-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
