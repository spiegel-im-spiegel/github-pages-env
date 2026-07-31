+++
title = "Flickr 短縮 URL"
date =  "2022-11-24T21:18:39+09:00"
description = "Base58 って標準化されてるんじゃないのか orz"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "flickr" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Flickr 写真ページの URL を短縮 URL に変換したいなぁ，と思いついた。

たとえば

{{< fig-img src="./52462249619_a35b2073ff_e.jpg" title="Stand-Up Paddleboard (SUP) in 宍道湖 | Flickr" link="https://www.flickr.com/photos/spiegel/52462249619/">}}

という写真を置いている Web ページの URL は

```text
https://www.flickr.com/photos/spiegel/52462249619
```

となっている。
一般化すると

{{< fig-quote class="nobox" type="markdown" title="Web Page URLs" link="https://www.flickr.com/services/api/misc.urls.html" lang="en" >}}
```text
https://www.flickr.com/photos/{user-id}/{photo-id} - individual photo
```
{{< /fig-quote >}}

という形式。
これを短縮 URL で表すには

{{< fig-quote class="nobox" type="markdown" title="Short URLs" link="https://www.flickr.com/services/api/misc.urls.html#short" lang="en" >}}
```text
https://flic.kr/p/{base58-photo-id}
```
{{< /fig-quote >}}

とすればいのだが（つまり `photo-id` を Base58 で符号化する），この Base58 ってのが曲者で，どうやらアプリケーションによって微妙に[実装が異なる](https://ja.wikipedia.org/wiki/Base58 "Base58 - Wikipedia")らしい。
なんだよ，それ。
標準化されてるんじゃないのか `orz`

ちなみに Flickr 版 Base58 の仕様はこちら：

- [Flickr: Discussing manufacturing flic.kr style photo URLs in Flickr API](https://www.flickr.com/groups/api/discuss/72157616713786392/)

幸いなことに Flickr 版にも対応した Base58 パッケージを公開している人がおられた。

- [itchyny/base58-go: Base58 encoding/decoding package and command written in Go](https://github.com/itchyny/base58-go)

ありがたや {{% emoji "ペコン" %}}

というわけで，簡単なサンプルコードを書いてみる。

```go
package main

import (
    "flag"
    "fmt"
    "net/url"
    "os"
    "strings"

    "github.com/itchyny/base58-go"
)

func FlickrShortURL(s string) string {
    u, err := url.Parse(s)
    if err != nil {
        return s
    }
    if !strings.HasSuffix(u.Hostname(), "flickr.com") {
        return s
    }
    var photoId string
    elms := strings.Split(strings.TrimSuffix(u.EscapedPath(), "/"), "/")[1:]
    if len(elms) == 3 && strings.EqualFold(elms[0], "photos") {
        photoId = elms[2]
    }
    if len(photoId) == 0 || !isDigitString(photoId) {
        return s
    }
    b, err := base58.FlickrEncoding.Encode([]byte(photoId))
    if err != nil {
        return s
    }
    return "https://flic.kr/p/" + string(b)
}

func isDigitString(s string) bool {
    return strings.IndexFunc(s, func(c rune) bool {
        return c < '0' || '9' < c
    }) < 0
}

func main() {
    flag.Parse()
    args := flag.Args()
    if len(args) == 0 {
        fmt.Fprintln(os.Stderr, os.ErrInvalid)
        return
    }
    fmt.Println(FlickrShortURL(args[0]))
}
```

`FlickrShortURL()` が変換関数。
引数の URL 文字列を簡単にチェックして Flickr の写真ページの形式なら短縮 URL に変換して返している。
それ以外なら素通し。

これを実行すると

```text
$ go run sample.go https://www.flickr.com/photos/spiegel/52462249619/
https://flic.kr/p/2nVUK4a
```

と出力される。
念のために検算しておこう。

件の写真ページにあるシェアボタン {{< icons "share-alt" >}} をクリックすると，以下のダイアログが表示される。

{{< fig-img src="./flickr-share.png" link="./flickr-share.png">}}

ここに表示されている短縮 URL が先程のサンプルコードの実行結果と同じなら無問題。

よーし，うむうむ，よーし。

[Go]: https://go.dev/

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
