+++
title = "Go モジュールの依存関係を可視化するツール ggm v0.2.0 をリリースした"
date =  "2019-05-04T18:09:55+09:00"
description = "前回作ったツールが思ったより有用なことに気がついて「これはちゃんとコードを書こう」と思い立った。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "tools", "golang", "module" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

前回「[Go モジュールの依存関係を可視化するツールを作った]({{< relref "./ggm.md" >}})」のだが，思ったより有用なことに気がついて「これはちゃんとコードを書こう」と思い立った。

ちうわけで，リリースしました。

- [Release v0.2.0 · spiegel-im-spiegel/ggm](https://github.com/spiegel-im-spiegel/ggm/releases/tag/v0.2.0)

今回はちゃんとバイナリも用意しているぞ（笑）

使い方はこんな感じ。

```text
$ ggm -h
Usage:
  ggm [flags] [input file]

Flags:
  -c, --config string   Configuration file
      --debug           Debug flag
  -h, --help            help for ggm
  -v, --version         Output version of ggm
```

[前回]と異なるのは `-c` オプションを付けたことか。
たとえば以下の内容で `ggm.toml` というファイルを作って

```toml
[node]
  fontname = "Inconsolata"
```

以下のコマンドラインで DOT ファイルを生成すると

```text
$ go mod graph | ggm -c ggm.toml
digraph G {
	ID = "G";
	
	n1[fontname="Inconsolata",label="github.com/spiegel-im-spiegel/ggm"];
	n2[fontname="Inconsolata",label="github.com/BurntSushi/toml\nv0.3.1"];
	n3[fontname="Inconsolata",label="github.com/emicklei/dot\nv0.9.3"];
	n4[fontname="Inconsolata",label="github.com/spf13/cobra\nv0.0.3"];
	n5[fontname="Inconsolata",label="github.com/spf13/pflag\nv1.0.3"];
	n6[fontname="Inconsolata",label="github.com/spiegel-im-spiegel/gocli\nv0.9.5"];
	n7[fontname="Inconsolata",label="golang.org/x/xerrors\nv0.0.0-20190410155217-1f06c39b4373"];
	n1->n2;
	n1->n3;
	n1->n4;
	n1->n5;
	n1->n6;
	n1->n7;
	
}
```

てな感じで `fontname` 属性を仕込むことができる。
これを `dot` コマンドに流し込むと

```text
$ go mod graph | ggm -c ggm.toml | dot -Tpng -o ggm.png
```

{{< fig-img src="./ggm.png" title="ggm.png" link="ggm.png" width="2380" >}}

てな風にフォントを [Inconsolata] にすることができるのですよ（勿論あらかじめフォントがインストールされていることが条件ね）。

`-c` オプションで指定するファイルは [TOML] 形式で `node` と `edge` の属性を設定することができる。
ただし手抜き実装で属性名やその値の正しさについてはノーチェックなのでご注意を（つまり `key=value` で表されるものなら何でも入る`w`）。

今回 [DOT 言語]用のビルダ・パッケージには [github.com/emicklei/dot](https://github.com/emicklei/dot "emicklei/dot: Go package for writing descriptions using the Graphviz DOT language") を利用している。
シンプルな設計で（簡単な図であれば）使い勝手がよい。

本当は本家の [Go] コンパイラが [DOT 言語]で吐いてくれればこんなの要らないんだけどねぇ。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< relref "./ggm.md" >}} "Go モジュールの依存関係を可視化するツールを作った"
[Inconsolata]: https://www.levien.com/type/myfonts/inconsolata.html
[TOML]: https://github.com/toml-lang/toml "toml-lang/toml: Tom's Obvious, Minimal Language"
[DOT 言語]: https://graphviz.gitlab.io/_pages/doc/info/lang.html "The DOT Language"

## ブックマーク

- [“go mod graph” を視覚化する]({{< ref "/golang/go-mod-graph.md" >}})

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
