+++
title = "Go モジュールの依存関係を可視化するツールを作った"
date =  "2019-05-02T18:33:20+09:00"
description = "main.go いっこだけの簡単なお仕事（笑）"
image = "/images/attention/go-logo_blue.png"
tags  = [ "tools", "golang", "module" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] に [graphviz] を[インストールした]({{< ref "/remark/2019/03/using-ubuntu.md#gvz" >}} "Ubuntu で遊ぶ")ので「何か動作確認できるものはないかなぁ」と思い [Go] モジュールの依存関係を可視化するツールを作ってみた。

- [spiegel-im-spiegel/ggm: Graph of Go Modules](https://github.com/spiegel-im-spiegel/ggm)

`main.go` いっこだけの簡単なお仕事（笑）

ダウンロードとインストールは以下の通りでよい（モジュール・モードがオンの場合）。

```text
$ go get github.com/spiegel-im-spiegel/ggm@latest
```

これで作成される `ggm` コマンドへ `go mod graph` の結果を流し込む。

```text
$ go mod graph | ggm
digraph gomod {
	N1->N2;
	N1 [ label="github.com/spiegel-im-spiegel/ggm" ];
	N2 [ label="github.com/awalterschulze/gographviz\nv0.0.0-20190221210632-1e9ccb565bca" ];

}
```

これを更に `dot` コマンドに流し込んで画像ファイルを出力する。

```text
$ go mod graph | ggm | dot -Tpng -o ggm.png
```

結果は以下の通り。

{{< fig-img src="./ggm.png" title="ggm.png" link="ggm.png" >}}

[DOT 言語]への変換は簡単なので直接書いてもよかったが，今回は [awalterschulze/gographviz] パッケージを使ってみた。
これ，色々と応用できるかもねぇ。

ちなみに [Go] のソースコードから UML 図（[PlantUML]）を生成するツールは以下の方が公開しておられる。

- [kazukousen/gouml: Automatically generate PlantUML from Go Code.](https://github.com/kazukousen/gouml)
- [GoのコードからPlantUMLコードを生成する静的解析ツールを作っている - 日記マン](https://i101330.hatenablog.com/entry/2019/04/14/205522)

感謝！

## ブックマーク

- [Graphvizとdot言語でグラフを描く方法のまとめ - Qiita](https://qiita.com/rubytomato@github/items/51779135bc4b77c8c20d)

- [“go mod graph” を視覚化する]({{< ref "/golang/go-mod-graph.md" >}}) : 今回のツールの元ネタ。

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[graphviz]: https://www.graphviz.org/ "Graphviz - Graph Visualization Software"
[DOT 言語]: https://graphviz.gitlab.io/_pages/doc/info/lang.html "The DOT Language"
[awalterschulze/gographviz]: https://github.com/awalterschulze/gographviz "awalterschulze/gographviz: Parses the Graphviz DOT language in golang"
[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[PlantUML]: http://plantuml.com/ja/ "Open-source tool that uses simple textual descriptions to draw beautiful UML diagrams."

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
