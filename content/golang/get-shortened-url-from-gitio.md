+++
date = "2015-11-26T20:29:30+09:00"
update = "2018-01-21T11:19:13+09:00"
description = "これって curl で書けるんなら Go 言語で表現できるんじゃね？"
draft = false
tags = ["golang", "github", "curl", "programming", "api"]
title = "Git.io から短縮 URL を取得するコード"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（初出： [Git.io 短縮 URL を golang コードで取得してみる - Qiita](http://qiita.com/spiegel-im-spiegel/items/042751d98e315e4e3382)）

## Git.io で短縮 URL が取得できるらしい

- [Git・Githubに隠された便利な機能 | GitHub Cheat Sheet（日本語訳） - Qiita](http://qiita.com/sotayamashita/items/1cf05f2a2be3d6fb3388)

これを読んでたら後ろの方に [Git.io] の話が出ていた。
このサイトで短縮 URL を生成できるらしい。

- [Git.io: GitHub URL Shortener](https://github.com/blog/985-git-io-github-url-shortener)

API が [curl] で掲載されていて，例えば私の [https://github.com/spiegel-im-spiegel](https://github.com/spiegel-im-spiegel) なら

```bash
$ curl -i "https://git.io" -F "url=https://github.com/spiegel-im-spiegel"
HTTP/1.1 201 Created
Server: Cowboy
Connection: keep-alive
Date: Sat, 08 Aug 2015 02:42:16 GMT
Status: 201 Created
Content-Type: text/html;charset=utf-8
Location: http://git.io/vOj52
Content-Length: 37
X-Xss-Protection: 1; mode=block
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Runtime: 0.210952
X-Node: 871d903e-a8e0-46ff-a96f-ef424385e5ed
X-Revision: b1d9ce07ccb700fc90398edafd397beb8d3bd772
Via: 1.1 vegur

https://github.com/spiegel-im-spiegel
```

てな感じで，ヘッダの Location 要素に短縮 URL が返ってくる仕組みらしい。
って，これって [curl] で書けるんなら [Go 言語]で表現できるんじゃね？

## cURL as DSL

- [cURL as DSL — cURL as DSL 1.0 documentation](https://shibukawa.github.io/curl_as_dsl/)
- [Shibu's Diary: cURL as DSLとは何だったのか。あるいは細かすぎて伝わらないcURL as DSL。](http://blog.shibu.jp/article/115602749.html)

[cURL as DSL] とは [curl] の構文を任意のコード[^a] に変換してくれるもので，どういうことかというと「[Web API は curl で表現すればいいんじゃね？](http://qiita.com/Hiraku/items/dfda2f8a5353b0742271)」ということらしい。

[^a]: 今のところは [Go 言語]のほかに Python3, PHP, JavaScript (node.js/XMLHttpRequest), Java, Objective-C (NSURL_Session/NSURLConnection), Vim Script (WebAPI-vim) に対応している。

## さっそく curl を Go 言語に変換してみる

では早速，上述の curl コマンドを [cURL as DSL] を使って [Go 言語]に変換してみる（ただし `-i` オプションは付けない）。
結果はこんな感じ。

```go
package main

import (
    "bytes"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
)

func main() {
    var buffer bytes.Buffer
    writer := multipart.NewWriter(&buffer)
    writer.WriteField("url", "https://github.com/spiegel-im-spiegel")
    writer.Close()

    resp, err := http.Post("https://git.io", "multipart/form-data; boundary="+writer.Boundary(), &buffer)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(string(body))
}
```

出力のみちょっと弄って

```go
package main

import (
    "bytes"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
)

func main() {
    var buffer bytes.Buffer
    writer := multipart.NewWriter(&buffer)
    writer.WriteField("url", "https://github.com/spiegel-im-spiegel")
    writer.Close()

    resp, err := http.Post("https://git.io", "multipart/form-data; boundary="+writer.Boundary(), &buffer)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("  Status: ", resp.Header.Get("Status"))
    log.Println("Location: ", resp.Header.Get("Location"))
    log.Println("    Body: ", string(body))
}
```

として実行すると

```bash
C:>go run gitio.go
2015/08/08 12:00:00   Status: 201 Created
2015/08/08 12:00:00 Location: http://git.io/vOj52
2015/08/08 12:00:00     Body: https://github.com/spiegel-im-spiegel
```

となり，めでたく短縮 URL が取得できた。

ちなみに最初の [curl] コマンドの `-F` を `-d` に替えて [cURL as DSL] にかけると[^b]

[^b]: [初出の記事](http://qiita.com/spiegel-im-spiegel/items/042751d98e315e4e3382)のコメントで [cURL as DSL] の作者の方に教えていただいた。感謝！

```go
package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
)

func main() {
    values := url.Values{
        "url": {"https://github.com/spiegel-im-spiegel"},
    }

    resp, err := http.PostForm("https://git.io", values)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(string(body))
}
```

のように変換される。
[`http`].`Post()` 関数から [`http`].`PostForm()` 関数に替わってかなりコードがすっきりした。
今回のような単純な request ならこちらの方がいいだろう。

## ついでにパッケージも作ってみた

- [spiegel-im-spiegel/gitioapi](https://github.com/spiegel-im-spiegel/gitioapi)

これを称して「他人の褌で相撲を取る」という[^c]。
なるほど。
こうやって API を実装していくんだね。

[^c]: 実はジェネレータで生成したコードの著作権は誰に帰属するのか，とかいろいろ思うところはあるのだけど，それはまた別の機会に。

## ブックマーク

- [よく使うcurlコマンドのオプションまとめ（12個） - Qiita](https://qiita.com/shtnkgm/items/45b4cd274fa813d29539)
- [curlコマンドをPythonやnode.jsのコードに変換する方法 - Qiita](https://qiita.com/tottu22/items/9112d30588f0339faf9b) : Python への変換は[こちらのサービス](https://curl.trillworks.com/ "Convert cURL command syntax to Python requests, Node.js code")のほうが簡潔なコードを吐いてくれるらしい

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[Git.io]: http://git.io/ "git.io"
[curl]: http://curl.haxx.se/ "curl and libcurl"
[cURL as DSL]: https://shibukawa.github.io/curl_as_dsl/ "cURL as DSL — cURL as DSL 1.0 documentation"
[`http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"
