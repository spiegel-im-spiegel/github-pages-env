+++
date = "2015-10-15T20:32:32+09:00"
description = "今回はちょっと横道にそれて，パッケージの依存状況を視覚化するツールをご紹介。"
draft = false
tags = ["golang", "package", "tools"]
title = "パッケージの依存状況の視覚化"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

今回はちょっと横道にそれて，パッケージの依存状況を視覚化するツールをご紹介。

- [hirokidaichi/goviz]

## goviz のインストール

導入は `go get` でOK。

```bash
C:>go get -v github.com/hirokidaichi/goviz
github.com/hirokidaichi/goviz (download)
github.com/jessevdk/go-flags (download)
github.com/hirokidaichi/goviz/dotwriter
github.com/jessevdk/go-flags
github.com/hirokidaichi/goviz/goimport
github.com/hirokidaichi/goviz/metrics
github.com/hirokidaichi/goviz
```

## パッケージ依存状況の視覚化

では，早速動かしてみましょう。

```bash
C:>goviz.exe -i github.com/hirokidaichi/goviz
digraph main{
	edge[arrowhead=vee]
	graph [rankdir=LR,compound=true,ranksep=1.0];
	/* plot github.com/hirokidaichi/goviz */
	"github.com/hirokidaichi/goviz"[shape="record",label="main|github.com/hirokidaichi/goviz|goviz.go",style="solid"]
	"github.com/hirokidaichi/goviz" -> "github.com/hirokidaichi/goviz/dotwriter"[dir=forward]
	/* plot github.com/hirokidaichi/goviz/dotwriter */
	"github.com/hirokidaichi/goviz/dotwriter"[shape="record",label="dotwriter|github.com/hirokidaichi/goviz/dotwriter|dotwriter.go",style="solid"]
	"github.com/hirokidaichi/goviz" -> "github.com/hirokidaichi/goviz/goimport"[dir=forward]
	/* plot github.com/hirokidaichi/goviz/goimport */
	"github.com/hirokidaichi/goviz/goimport"[shape="record",label="goimport|github.com/hirokidaichi/goviz/goimport|import.go\nimport_factory.go\nsource.go",style="solid"]
	"github.com/hirokidaichi/goviz/goimport" -> "github.com/hirokidaichi/goviz/dotwriter"[dir=forward]
	"github.com/hirokidaichi/goviz" -> "github.com/hirokidaichi/goviz/metrics"[dir=forward]
	/* plot github.com/hirokidaichi/goviz/metrics */
	"github.com/hirokidaichi/goviz/metrics"[shape="record",label="metrics|github.com/hirokidaichi/goviz/metrics|metrics.go",style="solid"]
	"github.com/hirokidaichi/goviz/metrics" -> "github.com/hirokidaichi/goviz/dotwriter"[dir=forward]
	"github.com/hirokidaichi/goviz" -> "github.com/jessevdk/go-flags"[dir=forward]
	/* plot github.com/jessevdk/go-flags */
	"github.com/jessevdk/go-flags"[shape="record",label="flags|github.com/jessevdk/go-flags|arg.go\nclosest.go\ncommand.go\ncommand_private.go\ncompletion.go\nconvert.go\nerror.go\nflags.go\ngroup.go\ngroup_private.go\nhelp.go\nini.go\nini_private.go\nman.go\nmultitag.go\noption.go\noption_private.go\noptstyle_other.go\noptstyle_windows.go\nparser.go\nparser_private.go\ntermsize.go\ntermsize_linux.go\ntermsize_nosysioctl.go\ntermsize_other.go\ntermsize_unix.go",style="solid"]
}
```

おー。
なんだか凄いコードが出力されました。
じつはこれ [DOT というデータ記述言語](https://ja.wikipedia.org/wiki/DOT%E8%A8%80%E8%AA%9E)で書かれたものです。
なので，この出力を [Graphviz] のツールに通すことで最終的な出力を得ます。

```bash
C:>goviz.exe -i github.com/hirokidaichi/goviz | dot.exe -Tpng -o goviz.png
```

{{< fig-img flickr="true" src="https://farm6.staticflickr.com/5782/21563262573_630b0eed8a.jpg" alt="output by goviz + graphviz" title="output by goviz + graphviz" link="https://www.flickr.com/photos/spiegel/21563262573/" >}}

## パッケージ依存度の評価

`-m` オプションを指定するとパッケージ依存度（結合度）の評価ができます。

```bash
C:>goviz.exe -i github.com/hirokidaichi/goviz -m
Inst:1.000 Ca(  0) Ce(  4)      github.com/hirokidaichi/goviz
Inst:0.500 Ca(  1) Ce(  1)      github.com/hirokidaichi/goviz/goimport
Inst:0.500 Ca(  1) Ce(  1)      github.com/hirokidaichi/goviz/metrics
Inst:0.000 Ca(  3) Ce(  0)      github.com/hirokidaichi/goviz/dotwriter
Inst:0.000 Ca(  1) Ce(  0)      github.com/jessevdk/go-flags
```

`Inst` は Instability， `Ca` は Afferent Couplings， `Ce` は Efferent Couplings かな。

Afferent Couplings は，そのパッケージに依存しているパッケージがいくつあるか，を示すものです。
Efferent Couplings は，逆にそのパッケージが依存しているパッケージがいくつあるか，を示すものです。
Instability は $Inst = Ce / (Ce + Ca)$ で算出される値で，この値が大きいほど他パッケージへの依存度が高いと評価できます。
共通ライブラリとして運用したいパッケージはなるべく Instability を低く抑えたいところです。
またソースコードを読む場合は Instability の高いパッケージから優先的にみるといいかもしれません。

## goviz の起動オプション

[hirokidaichi/goviz] のオプションは以下の通り。

```bash
C:>goviz.exe -h
Usage:
  goviz.exe [OPTIONS]

Application Options:
  /i, /input:    intput ploject name
  /o, /output:   output file (default: STDOUT)
  /d, /depth:    max plot depth of the dependency tree (default: 128)
  /f, /focus:    focus on the specific module
  /s, /search:   top directory of searching
  /l, /leaf      whether leaf nodes are plotted (default: false)
  /m, /metrics   display module metrics (default: false)

Help Options:
  /?             Show this help message
  /h, /help      Show this help message
```

## ブックマーク{#bookmark}

- [そろそろ理解しておきたいのでDockerのソースコードをビジュアルに読む！ - Qiita](http://qiita.com/hirokidaichi/items/52fc6286c9e432792a07) : [hirokidaichi/goviz] : 作者による解説
- [データのビジュアル化を最少の労力で: Graphviz](http://www.showa-corp.jp/special/graphtools/graphviz.html)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[hirokidaichi/goviz]: https://github.com/hirokidaichi/goviz "hirokidaichi/goviz"
[Graphviz]: http://www.graphviz.org/ "Graphviz | Graphviz - Graph Visualization Software"
