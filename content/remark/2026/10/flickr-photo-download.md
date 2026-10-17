+++
title = "Flickr の写真ファイルを取得できなくなった"
date =  "2026-10-17T22:24:19+09:00"
description = "自作ツールの不具合を修正していく"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "flickr", "atproto", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## Flickr の写真ファイルを取得できなくなった

私は[自作ツール][`goark/toolbox`]（公開してるけど，自分専用ツール）を使って [Flickr にアップしている写真](https://www.flickr.com/photos/spiegel/ "Yasuhiro ARAKAWA | Flickr")を Mastodon と Bluesky に自動投稿してるんだけど，昨日くらいからエラーで投稿できなくなった。
どうも Flickr サイトから写真データをダウンロードする際に [429 (Too Many Requests)](https://developer.mozilla.org/ja/docs/Web/HTTP/Reference/Status/429 "429 Too Many Requests - HTTP | MDN") が返っているようだ。

とりあえず自動更新処理を止めて調べ始めたのだが `curl` コマンドだと写真ファイルを普通に取れるのね。
ブラウザで URL を直接叩いても問題ない。
一度に複数処理するのがダメなのかと意図的に遅延させてみたけどダメ。
他にも色々弄ってみたんだけど，結局よく分からなかったので，自力で調べるのは諦め [GitHub Copilot] に調べさせてみた。
以下が調査結果：

{{< div-ai type="markdown" >}}
可能性として一番濃いのは「Go のHTTP/2クライアント実装の指紋で弾かれている」ケースです。

今回こちらで確認した事実は以下です。

1. `sample.go` をそのまま `go run sample.go` すると `HTTP/2.0 429`。
2. 同じコードでも `GODEBUG=http2client=0`（HTTP/2無効）だと成功。
3. `net/http` のカスタム `Transport`（実質HTTP/1.1）でも `200 OK`。
4. `curl` は HTTP/2 でも `200`。

この組み合わせから、単純なURL間違いや `fetch` のバグよりも、次が考えやすいです。

1. CDN/WAFのbot判定差<br>
GoのHTTP/2（TLS/ALPN, SETTINGS, HPACK, ヘッダー順序など）のシグネチャが `curl` と異なり、レート制限・Bot制限に触れることがあります。

2. `User-Agent` 偽装との不整合<br>
`sample.go` で `User-Agent: curl/8.18.0` を送っていますが、実際のTLS/HTTP2実装はcurlではないため、判定が厳しい環境だと不自然と見なされることがあります。

3. 経路差・POP差による429<br>
同じURLでも、接続先POPやタイミング次第で429/200が揺れることがあります。`go` と `curl` でDNS解決・接続経路が微妙に変わる場合もあります。

実運用での対策候補です。

1. 対象ホストだけHTTP/1.1にフォールバックする（最も効く可能性が高い）
2. `User-Agent` を偽装せず、実アプリ名にする
3. 429時に `Retry-After` を見て指数バックオフ再試行
4. 必要なら `Referer` など相手側が期待するヘッダーを追加（規約順守前提）
{{< /div-ai >}}

ちなみに上のやり取りで使ってる `sample.go` の中身はこんな感じ。

```go {hl_lines=[16]}
package main

import (
	"context"
	"fmt"

	"github.com/goark/fetch"
)

func main() {
	u, err := fetch.URL("https://live.staticflickr.com/65535/55582060788_9bf7829404_b.jpg")
	if err != nil {
		fmt.Printf("%+v\n", err) // %+v でエラー詳細を出力
		return
	}
	resp, err := fetch.New().GetWithContext(
		context.Background(),
		u,
		fetch.WithRequestHeaderSet("User-Agent", "curl/8.18.0"),
		fetch.WithRequestHeaderSet("Accept", "*/*"),
	)
	if err != nil {
		fmt.Printf("%+v\n", err) // %+v ステータスコードを含むエラー詳細を出力
		return
	}
	defer resp.Close()

	fmt.Println("HTTP Protocol:", resp.Proto())
}
```

これを実行すると以下の出力になる。

```text
$ go run sample.go | jq .
{
  "Type": "*errs.Error",
  "Err": {
    "Type": "*errors.errorString",
    "Msg": "invalid HTTP request"
  },
  "Context": {
    "function": "github.com/goark/fetch.(*client).GetWithContext",
    "url": "https://live.staticflickr.com/65535/55582060788_9bf7829404_b.jpg"
  },
  "Cause": {
    "Type": "*errs.Error",
    "Err": {
      "Type": "*fmt.wrapError",
      "Msg": "bad HTTP status: status 429",
      "Cause": {
        "Type": "*errors.errorString",
        "Msg": "bad HTTP status"
      }
    },
    "Context": {
      "function": "github.com/goark/fetch.(*client).fetch",
      "protocol": "HTTP/2.0",
      "status": 429
    }
  }
}
```

429 が返っているのがお分かりだろうか。

[`goark/fetch`] は標準の [`net/http`] をラップした自作パッケージで，`fetch.New()` で引数を省略するとデフォルトの `http.Client{}` インスタンスを使うようになっている。
ここを，たとえば

```go {hl_lines=["1-8"]}
resp, err := fetch.New(fetch.WithHTTPClient(
    &http.Client{
        Transport: &http.Transport{
            ForceAttemptHTTP2: false, // HTTP/2 の無効化
            TLSNextProto:      map[string]func(string, *tls.Conn) http.RoundTripper{}, // TLSNextProto を空にする
        },
    },
)).GetWithContext(
    context.Background(),
    u,
    fetch.WithRequestHeaderSet("User-Agent", "curl/8.18.0"),
    fetch.WithRequestHeaderSet("Accept", "*/*"),
)
```

とすれば HTTP/1.1 専用のクライアントになる。
これを使ってリクエストを発行し，ステータス [200 (OK)](https://developer.mozilla.org/ja/docs/Web/HTTP/Reference/Status/200 "200 OK - HTTP | MDN") で写真ファイルがダウンロードできることを確認できた。

強制的に HTTP/1.1 を使うのはイケてない気がするが，素の [`net/http`] の HTTP/2 挙動を調査するのは大変だし，今回は対症療法ということで HTTP/1.1 専用クライアントを生成する方法を採用した。

毎回構造体をゴリゴリ書くのは面倒なので [`goark/fetch`] に HTTP/1.1 専用のクライアントを生成するオプションを追加して対応した。
さらにこれを[自作ツール][`goark/toolbox`]に組み込んで修正完了。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:4kzexuq4xieuarc3q6lmkeou/app.bsky.feed.post/3my2xdddgao25" data-bluesky-cid="bafyreihnrho6bxnm4zadlpenf6kt2rqh74aiwpwb3uyyxrcssybsuvoejm" data-bluesky-embed-color-mode="system"><p lang="">#flickr #photography 街路樹
https://www.flickr.com/photos/spiegel/55594508941/
<br><br><a href="https://bsky.app/profile/did:plc:4kzexuq4xieuarc3q6lmkeou/post/3my2xdddgao25?ref_src=embed">[image or embed]</a></p>&mdash; Spiegel with お散歩カメラ (<a href="https://bsky.app/profile/did:plc:4kzexuq4xieuarc3q6lmkeou?ref_src=embed">@photos.baldanders.info</a>) <a href="https://bsky.app/profile/did:plc:4kzexuq4xieuarc3q6lmkeou/post/3my2xdddgao25?ref_src=embed">2026年10月17日 20:10</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

よしよし。
ちゃんと動いてるな。

## Go 1.27 から http.Response.Body の中身を読み捨てなくてよくなった

- [Better HTTP/1 connection reuse - The Go Programming Language](https://go.dev/blog/http-connection-reuse)

昔から [Go] 特有のイディオムとして [`http`][`net/http`]`.Response.Body` を `Close()` する前に中身を読み捨てる必要があったが，1.27 からは [`http`][`net/http`]`.Transport` がいい感じに処理してくれるため，明示的に読み捨てる必要がなくなった。

{{< fig-quote type="markdown" title="Better HTTP/1 connection reuse" link="https://go.dev/blog/http-connection-reuse" lang="en" >}}
Starting in Go 1.27, [`net/http.Transport`](https://go.dev/pkg/net/http#Transport) automatically does everything we have implemented in the code above: it will asynchronously drain an unread HTTP/1 response body for up to 50 ms or 256 KB. It also implements additional optimizations; for example, when a response contains the `Connection: close` header, draining the body is guaranteed to be wasted work, and it therefore skips doing so. Note that this behavior might evolve in future versions of Go, as other possible enhancements become apparent.

As a general best practice, we now recommend that users simply close their HTTP bodies once they no longer need them, like in our original example:

```go
func getStatus(url string) (int, error) {
    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()
    return resp.StatusCode, nil
}
```

Go is all about simplicity. Let your code stay clean and simple, and let [net/http](https://go.dev/pkg/net/http) deal with the complexities on your behalf.
{{< /fig-quote >}}

これを [`goark/fetch`] に対応しようか悩んだのだが，このパッケージはまだ 1.26 対応のままなので，今回は諦めた。
1.28 が出て 1.26 がサポートから外れたら，改めて対応することにしよう。

[Go] はシンプルが正義！

## ATProto ライブラリ indigo の古い API が Deprecated になってた

[自作ツール][`goark/toolbox`]で Bluesky に投稿するために [Go] 用の [ATProto][atproto] ライブラリ [`bluesky-social/indigo`] パッケージを使っているのだが，今の最新版（頼むからバージョンタグを付けてくれ！）をインポートしたら初期から使ってる API が deprecated になっていた。

今のところは見なかったことにしてバージョンを戻している（笑）

でも，近いうちに時間を取って対応しないと。
他にやりたいことが山積みなんだがな。

[Go]: https://go.dev/
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[`goark/toolbox`]: https://github.com/goark/toolbox "goark/toolbox: A collection of miscellaneous commands"
[`goark/fetch`]: https://github.com/goark/fetch "goark/fetch: Fetch Data from URL"
[`net/http`]: https://pkg.go.dev/net/http "http package - net/http - Go Packages"
[`bluesky-social/indigo`]: https://github.com/bluesky-social/indigo "bluesky-social/indigo: Go source code for Bluesky's atproto services."
[atproto]: https://atproto.com/ "The AT Protocol"

## 参考図書

{{< linkcard "9e8d8f717b1d1d85b0da7f07cb10a83b78d4227f" >}} <!-- プログラミング言語Go -->
{{< review-paapi "B0CFL1DK8Q" >}} <!-- Go言語 100Tips -->
{{< review-paapi "B0DNYMMBBQ" >}} <!-- Go言語で学ぶ並行プログラミング -->
