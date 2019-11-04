+++
title = "紙芝居用の簡易 Web サーバを Go 言語で書く"
date = "2018-02-09T21:17:46+09:00"
description = "今は Go 言語が書けるし，今回は Go 言語で組むことにした。めっさ簡単。"
image = "/images/attention/kitten.jpg"
tags        = [ "golang", "web", "engineering" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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
  mathjax = false
  mermaidjs = false
+++

3年ほど前に紙芝居（静的な HTML ファイルを表示するだけの簡単なお仕事）用の簡易 Web サーバを [node.js] で組んだのだが

- [「紙芝居」用の簡単 HTTP サーバを立てる - Qiita](https://qiita.com/spiegel-im-spiegel/items/38b2e0b16ffc4f3548b3)
- [「紙芝居」用の簡単 HTTP サーバを立てる（その2） - Qiita](https://qiita.com/spiegel-im-spiegel/items/1a806ad8a9c38a0d7b70)

今は [Go 言語]が書けるし，今回は [Go 言語]で組むことにした。
こんな感じ。

```go
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("p", 3000, "port number")
	host := flag.String("host", "", "hostname")
	path := "./html"
	flag.Parse()
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}
	addr := fmt.Sprintf("%s:%d", *host, *port)

	if len(*host) == 0 {
		fmt.Printf("Open http://localhost%s/\n", addr)
	} else {
		fmt.Printf("Open http://%s/\n", addr)
	}
	fmt.Println("Press ctrl+c to stop")
	http.Handle("/", http.FileServer(http.Dir(path)))
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
```

これを適当なファイル（たとえば `mock-up.go`）に保存して

```text
$ go build mock-up.go
```

とすればよい。
標準パッケージのみ使ってるので他所からパッケージを取ってくる必要はない。
めっさ簡単。
ブラボー！

起動するとカレントディレクトリ直下の `html` フォルダ以下のファイルを表示できる。

```text
$ mock-up
Open http://localhost:3000/
Press ctrl+c to stop
```

この状態で任意のブラウザから `http://localhost:3000/` へアクセスすればよい。
停止する場合は `Ctrl-C` 押下でおｋ。

ポート番号を変える場合は `-p` オプションでポート番号を指定する。

```text
$ mock-up.exe -p 8080
Open http://localhost:8080/
Press ctrl+c to stop
```

バインドするホスト名を指定する場合は `-host` オプションを使う。

```text
$ mock-up.exe -p 8080 -host hostname
Open http://hostname:8080/
Press ctrl+c to stop
```

ルートフォルダを指定することも可能。

```text
$ mock-up.exe -p 8080 -host hostname C:\path\to\htmlpub
Open http://hostname:8080/
Press ctrl+c to stop
```

存在しないフォルダを指定してもエラーにならず，アクセスしても 404 が返るだけ。
Directory traversal 脆弱性は問題ないと思う。
こっそり他人の PC に仕込んだりしないように（笑）

## ブックマーク

- [Net/httpで安全に静的ファイルを返す - Shogo's Blog](https://shogo82148.github.io/blog/2016/04/13/serving-static-files-in-golang/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[node.js]: https://nodejs.org/en/

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4873117526?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51UoREcNrnL._SL160_.jpg" width="125" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4873117526?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">Go言語によるWebアプリケーション開発</a></dt>
    <dd>Mat Ryer (著), 鵜飼 文敏 (監修), 牧野 聡 (翻訳)</dd>
    <dd>オライリージャパン 2016-01-22</dd>
    <dd>大型本</dd>
    <dd>4873117526 (ASIN), 9784873117522 (EAN), 4873117526 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="3">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">日本語監訳者による解説（付録 B）が意外に役に立つ感じ。 Web アプリケーションだけでなく，サーバサイドで動く CLI アプリへの言及もある。良書だが今となってはちょっと内容が古い。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-08-12">2019-08-12</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
