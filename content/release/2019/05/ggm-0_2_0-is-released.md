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

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< relref "./ggm.md" >}} "Go モジュールの依存関係を可視化するツールを作った"
[Inconsolata]: https://www.levien.com/type/myfonts/inconsolata.html
[TOML]: https://github.com/toml-lang/toml "toml-lang/toml: Tom's Obvious, Minimal Language"
[DOT 言語]: https://graphviz.gitlab.io/_pages/doc/info/lang.html "The DOT Language"

## ブックマーク

- [“go mod graph” を視覚化する]({{< ref "/golang/go-mod-graph.md" >}})

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
