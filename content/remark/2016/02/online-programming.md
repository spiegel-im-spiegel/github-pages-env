+++
date = "2016-02-14T20:19:41+09:00"
description = "paiza.IO を試してみる。"
draft = false
tags = ["programming"]
title = "オンラインでプログラミング"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

先ほどの「[週末スペシャル]({{< relref "remark/2016/02/14-stories.md" >}})」でも紹介したが， [paiza.IO] を試してみる。

- [ブラウザでプログラミング・実行ができる「オンライン実行環境」| paiza.IO](https://paiza.io/)
- [GistをPaiza.ioで使ってみる - Qiita](http://qiita.com/omochiiiY/items/b3b3f7ece1dedca1d4e1)

コードを試すだけならアカウントはいらないが，コードを再利用したり他のサービスと連携する場合はサインアップが必要。
[GitHub] アカウントと連携できるのでその辺は問題ないだろう。

まず「[スタック追跡とパニック・ハンドリング](http://localhost:1313/golang/stack-trace-and-panic-handling/)」のコード例を [paiza.IO] で[書いてみた](https://paiza.io/projects/MencTrqIn3FYdg6u53xNSg)。

{{< fig-gen >}}
<iframe src="https://paiza.io/projects/e/MencTrqIn3FYdg6u53xNSg?theme=github" width="100%" height="500" scrolling="no" seamless="seamless"></iframe>
{{< /fig-gen >}}

こんな感じでコードと実行結果を埋め込むこともできる[^cke]。
また埋め込まれたコードはその場で編集して実行し直すこともできる。
試しに for 文の中を以下のコードに差し替えて実行してみて欲しい。

[^cke]: ただし埋め込んだコードを表示するにはブラウザ設定で third-party cookie を有効にする必要がある。 Firefox であれば例外設定をすることもできる。いや，こういうのマジ勘弁して欲しいんですけど。今時 third-party cookie とかありえない。

```go
for depth := 0; ; depth++ {
    pc, _, line, ok := runtime.Caller(depth)
    if !ok {
        break
    }
    fmt.Fprintf(log, " -> %d: %s: (%d)\n", depth, runtime.FuncForPC(pc).Name(), line)
}
```

[paiza.IO] で書いたコードは Gist と連携できる。
たとえば上のコードは

{{< fig-gist "https://gist.github.com/spiegel-im-spiegel/89526909cc206f31c1d7" >}}

という感じで Gist に保存されている。
Gist との同期は自動ではなく明示的に指定する必要がある。

[paiza.IO] で書いたコードで他の（Twitter API などの）サービスと接続することも可能。
"[The Go Playground](https://play.golang.org/)” では外部との通信はできないので，これは嬉しい機能である。
たとえば「[Git.io から短縮 URL を取得するコード]({{< relref "golang/get-shortened-url-from-gitio.md" >}})」を参考に書いたコードがこれ。

{{< fig-gen >}}
<iframe src="https://paiza.io/projects/e/uUG8z-Teb45q4RZIBSSAeg?theme=github" width="100%" height="500" scrolling="no" seamless="seamless"></iframe>
{{< /fig-gen >}}

実行をスケジューリングすることも可能なので，簡単なサーバ監視プログラムとか，色々な使い方ができそうである。

更に [Twitter](https://twitter.com/) でコードを書くことも可能。
あらかじめ @[paiza_run](https://twitter.com/paiza_run) を follow しておいて

{{< fig-gen >}}
<blockquote class="twitter-tweet" data-lang="ja"><p lang="en" dir="ltr"><a href="https://twitter.com/paiza_run">@paiza_run</a> go: package main<br>import &quot;fmt&quot;<br>func main() {<br>    fmt.Println(&quot;Hello World!&quot;)<br>}</p>&mdash; Spiegel im Spiegel (@spiegel_2007) <a href="https://twitter.com/spiegel_2007/status/698818840533184512">2016, 2月 14</a></blockquote>
{{< /fig-gen >}}

とコードを投げれば

{{< fig-gen >}}
<blockquote class="twitter-tweet" data-lang="ja"><p lang="en" dir="ltr"><a href="https://twitter.com/spiegel_2007">@spiegel_2007</a> Hello World! <a href="https://twitter.com/hashtag/paiza_run_result?src=hash">#paiza_run_result</a></p>&mdash; paiza_run (@paiza_run) <a href="https://twitter.com/paiza_run/status/698818848921776128">2016, 2月 14</a></blockquote>
{{< /fig-gen >}}

と結果が返ってくる。
これはスクリプト言語でやったほうが面白いかな。

更に更に，他のユーザが書いたコードを fork することも可能。
またチャット等でリアルタイムにやり取りしながらコードを書くこともできるみたい。
遠隔ペアプログラミングとかできそうだよね。

ただし， [paiza.IO] では実行時間と使えるメモリに制限があって，大体どの言語でも2秒で512MBが制限になっている模様。
また [Go 言語]の場合は標準以外の外部パッケージが使えない。
たとえば以下は「[文字エンコーディング変換]({{< relref "golang/transform-character-encoding.md" >}})」で書いたコードだが

```go
package main

import (
    "fmt"
    "io"
    "os"

    "golang.org/x/text/encoding/japanese"
    "golang.org/x/text/transform"
)

func main() {
    reader := NewDecoder(os.Stdin)
    writer := NewEncoder(os.Stdout)
    if _, err := io.Copy(writer, reader); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}

func NewDecoder(reader io.Reader) *transform.Reader {
    return transform.NewReader(reader, japanese.ShiftJIS.NewDecoder())
}

func NewEncoder(writer io.Writer) *transform.Writer {
    return transform.NewWriter(writer, japanese.EUCJP.NewEncoder())
}
```

[paiza.IO] ではコンパイル時にエラーになる。

欲を言えば "[The Go Playground](https://play.golang.org/)” にあるようなコードの整形機能が欲しいところである。

[paiza.IO]: https://paiza.io/ "ブラウザでプログラミング・実行ができる「オンライン実行環境」| paiza.IO"
[GitHub]: https://github.com/ "GitHub"
[Go 言語]: https://golang.org/ "The Go Programming Language"
