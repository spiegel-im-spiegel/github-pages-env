+++
title = "“go mod graph” を視覚化する"
date = "2019-02-12T23:49:55+09:00"
description = "やっつけでコードを書いてみた。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "module" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回も小ネタで。

`go mod graph` コマンドを使うと以下のようにモジュール間の依存関係を表示してくれるのだが

```text
$ go mod graph
github.com/spiegel-im-spiegel/gpgpdump github.com/BurntSushi/toml@v0.3.1
github.com/spiegel-im-spiegel/gpgpdump github.com/inconshreveable/mousetrap@v1.0.0
github.com/spiegel-im-spiegel/gpgpdump github.com/pkg/errors@v0.8.1
github.com/spiegel-im-spiegel/gpgpdump github.com/spf13/cobra@v0.0.3
github.com/spiegel-im-spiegel/gpgpdump github.com/spf13/pflag@v1.0.3
github.com/spiegel-im-spiegel/gpgpdump github.com/spiegel-im-spiegel/gocli@v0.9.1
github.com/spiegel-im-spiegel/gpgpdump golang.org/x/crypto@v0.0.0-20190208162236-193df9c0f06f
github.com/spiegel-im-spiegel/gocli@v0.9.1 github.com/mattn/go-isatty@v0.0.4
```

パッと見で分かりにくいのでこれを [PlantUML] の図に変換することを考える。
やっつけでこんなコードを書いてみた。

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mapPackage = map[string]string{}

type pair struct {
	left  string
	right string
}

var graph = []pair{}

func split(text string, count int) int {
	ss := strings.Split(text, " ")
	for _, s := range ss {
		if _, ok := mapPackage[s]; !ok {
			count++
			mapPackage[s] = fmt.Sprintf("P%d", count)
		}
	}
	p := pair{left: mapPackage[ss[0]], right: mapPackage[ss[1]]}
	graph = append(graph, p)
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		count = split(scanner.Text(), count)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("@startuml")
	fmt.Println("hide circle")
	fmt.Println("hide empty members")
	for k, v := range mapPackage {
		fmt.Printf("class \"%v\" as %v\n", strings.Replace(k, "@", "\\n", -1), v)
	}
	for _, p := range graph {
		fmt.Printf("\"%v\" ..> \"%v\"\n", p.left, p.right)
	}
	fmt.Println("@enduml")

	return
}
```

難しいことは何もしてないので，まぁ見れば分かるであろう。
これを実行バイナリにコンパイルしておく。

```text
$ go build mod-graph.go
```

では，実際に変換してみよう。

```text
$ go mod graph | mod-graph
@startuml
hide circle
hide empty members
class "github.com/spiegel-im-spiegel/gpgpdump" as P1
class "github.com/BurntSushi/toml\nv0.3.1" as P2
class "github.com/pkg/errors\nv0.8.1" as P4
class "github.com/spiegel-im-spiegel/gocli\nv0.9.1" as P7
class "github.com/mattn/go-isatty\nv0.0.4" as P9
class "github.com/inconshreveable/mousetrap\nv1.0.0" as P3
class "github.com/spf13/cobra\nv0.0.3" as P5
class "github.com/spf13/pflag\nv1.0.3" as P6
class "golang.org/x/crypto\nv0.0.0-20190208162236-193df9c0f06f" as P8
"P1" ..> "P2"
"P1" ..> "P3"
"P1" ..> "P4"
"P1" ..> "P5"
"P1" ..> "P6"
"P1" ..> "P7"
"P1" ..> "P8"
"P7" ..> "P9"
@enduml
```

変換結果を適当なファイルにリダイレクトし PNG ファイルへとコンパイルする。

```text
$ go mod graph | mod-graph > gpgpdump.puml
java -jar plantuml.jar -nometadata -charset utf-8 -tpng *.puml
```

結果はこんな感じ。

{{< fig-img src="./gpgpdump.png" title="gpgpdump.png" link="gpgpdump.png" width="1419" >}}

んー。
こんなもんかな。
本当は `go mod graph` コマンドが DOT 言語とかで出力してくれれば一番いいんだけどねぇ。

## ブックマーク

- [真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})
- [Go モジュールの依存関係を可視化するツールを作った]({{< ref "/release/2019/05/ggm-0_2_0-is-released.md" >}})
- [Go モジュールの依存関係を可視化するツール ggm v0.2.0 をリリースした]({{< ref "/release/2019/05/ggm-0_2_0-is-released.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

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
