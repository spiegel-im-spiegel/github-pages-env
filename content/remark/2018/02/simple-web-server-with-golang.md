+++
title = "紙芝居用の簡易 Web サーバを Go 言語で書く"
date = "2018-02-09T21:17:46+09:00"
description = "今は Go 言語が書けるし，今回は Go 言語で組むことにした。めっさ簡単。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "web", "engineering" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

1. [紙芝居用の簡易 Web サーバを Go 言語で書く]({{< ref "/remark/2018/02/simple-web-server-with-golang.md" >}}) ←イマココ
2. [紙芝居用の簡易サーバを書く【Go 1.16 版】]({{< ref "/golang/embeded-filesystem.md" >}})
3. [紙芝居用の簡易 Web サーバを Go 言語で書く【叱られ編】]({{< ref "/golang/simple-web-server-with-golang-3.md" >}})

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

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->

{{% review-paapi "4873117526" %}} <!-- Go言語によるWebアプリケーション開発 -->
