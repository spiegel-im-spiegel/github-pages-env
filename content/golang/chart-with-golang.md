+++
title = "Go 言語でもグラフを描きたい！"
date =  "2020-04-25T20:48:53+09:00"
description = " Gonum は本来，数値計算パッケージなのだが plot パッケージを使って簡単なグラフを描くこともできるらしい。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "chart", "math" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いや，決まったデータに対して棒グラフとか折れ線グラフとか簡単に描けないかなぁ，と思っただけなんスけどね。
ぶっちゃけ「グラフを描く」だけなら，それこそ [gnuplot] で必要十分なんだけど，今回もお遊びということで。

とはいえ自前でコードを書くのはアレなので，なにか手頃なパッケージはないかなぁ，とググってみたら [Gonum] ってのがいいらしい。
[Gonum] は

{{< fig-quote type="markdown" title="Gonum" link="https://www.gonum.org/" lang="en" >}}
{{< quote >}}Gonum contains libraries for [matrices and linear algebra](https://godoc.org/gonum.org/v1/gonum/mat); [statistics](https://godoc.org/gonum.org/v1/gonum/stat), [probability](https://godoc.org/gonum.org/v1/gonum/stat/distuv) [distributions](https://godoc.org/gonum.org/v1/gonum/stat/distmv), and [sampling](https://godoc.org/gonum.org/v1/gonum/stat/sampleuv); tools for [function differentiation](https://godoc.org/gonum.org/v1/gonum/diff/fd), [integration](https://godoc.org/gonum.org/v1/gonum/integrate/quad), and [optimization](https://godoc.org/gonum.org/v1/gonum/optimize); [network](https://godoc.org/gonum.org/v1/gonum/graph) creation and analysis; and more{{< /quote >}}.
{{< /fig-quote >}}

とある通り，本来は数値計算パッケージなのだが， [`plot`] パッケージを使って簡単なグラフを描くこともできるらしい。
どんなグラフが描けるかについては以下が参考になる。

- [Example plots · gonum/plot Wiki · GitHub](https://github.com/gonum/plot/wiki/Example-plots)

試しに何か描いてみる。

## 棒グラフを描いてみる

まずは普通に棒グラフを描いてみよう。
ちょっと長いが，こんな感じのコードでどうだろう。

```go
package main

import (
	"fmt"
	"math"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type CaseData struct {
	Date     string
	NewCases int
}

var Dataset = []CaseData{
	{"2020-04-18", 628},
	{"2020-04-19", 566},
	{"2020-04-20", 390},
	{"2020-04-21", 368},
	{"2020-04-22", 377},
	{"2020-04-23", 423},
	{"2020-04-24", 469},
}

func main() {
	//import data
	labelX := []string{}
	dataY := plotter.Values{}
	for _, d := range Dataset {
		labelX = append(labelX, d.Date)
		dataY = append(dataY, (float64)(d.NewCases))
	}

	//new plot
	p, err := plot.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//new bar chart
	bar, err := plotter.NewBarChart(dataY, vg.Points(20))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(2) //plotutil.SoftColors[2]
	bar.Offset = 0
	bar.Horizontal = false
	p.Add(bar)

	//legend
	p.Legend.Add("New confirmed cases by date", bar)
	p.Legend.Top = true   //top
	p.Legend.Left = false //right
	p.Legend.XOffs = 0
	p.Legend.YOffs = -10

	//labels of X
	p.NominalX(labelX...)
	p.X.Label.Text = "Date"
	p.X.Padding = 0
	p.X.Width = p.Y.Width
	p.X.Tick.Label.Rotation = math.Pi / 2.5
	p.X.Tick.Label.XAlign = draw.XRight
	p.X.Tick.Label.YAlign = draw.YCenter

	//labels of Y
	p.Y.Label.Text = "Cases"
	p.Y.Padding = 0
	p.Y.Min = 0
	p.Y.Max = 800

	//title
	p.Title.Text = "Confirmed COVID-2019 Cases in Japan"

	//output image
	if err := p.Save(15*vg.Centimeter, 15*vg.Centimeter, "bar-chart-1.png"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
```

実行結果はこんな感じ。

{{< fig-img src="./bar-chart-1.png" title="bar-chart-1.png" link="./bar-chart-1.png" width="567" >}}

## [Gonum] でも日本語を使いたい

[`plot/vg`] パッケージにはいくつかのフリー・フォントが含まれている。
どのようなフォントが使えるかについては [`vg`]`.FontMap` を参照すればよい。

```go
package main

import (
	"fmt"

	"gonum.org/v1/plot/vg"
)

func main() {
	for k, v := range vg.FontMap {
		fmt.Println(k, ":", v)
	}
    //Output:
    //Helvetica-Oblique : LiberationSans-Italic
    //Helvetica-BoldOblique : LiberationSans-BoldItalic
    //Times-Italic : LiberationSerif-Italic
    //Courier : LiberationMono-Regular
    //Courier-Bold : LiberationMono-Bold
    //Courier-Oblique : LiberationMono-Italic
    //Helvetica : LiberationSans-Regular
    //Helvetica-Bold : LiberationSans-Bold
    //Times-BoldItalic : LiberationSerif-BoldItalic
    //Courier-BoldOblique : LiberationMono-BoldItalic
    //Times-Roman : LiberationSerif-Regular
    //Times-Bold : LiberationSerif-Bold
}
```

たとえば，先程のコードに

```go
//default font
plot.DefaultFont = "Courier"
plotter.DefaultFont = "Courier"
```

を追加すれば

{{< fig-img src="./bar-chart-2.png" title="bar-chart-2.png" link="./bar-chart-2.png" width="567" >}}

という感じに既定フォントを入れ換えることができる。

[`plot/vg`] パッケージに含まれるフォントは当然ながら日本語の字体を含んでいない。
日本語が使いたいのであれば，別途日本語フォントを読み込む必要がある[^fnt1]。
こんな感じ。

[^fnt1]: OS のフォント・キャッシュが使えればいいんだけど [`plot/vg`] パッケージ単体では無理みたい。どなたかいい方法があれば教えてください。

```go
//import japanese fonts
b, err := ioutil.ReadFile("/usr/local/texlive/2020/texmf-dist/fonts/truetype/public/ipaex/ipaexg.ttf")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
ft, err := truetype.Parse(b)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
fontName := "ipaex"
vg.AddFont(fontName, ft)

//default font
plot.DefaultFont = fontName
plotter.DefaultFont = fontName
```

今回は [TeX Live] 2020 に収録されている [IPA]ex フォントを流用してみた。
これを使って，更にラベルのいくつかを日本語に置き換えたのがこちら。

{{< fig-img src="./bar-chart-3.png" title="bar-chart-3.png" link="./bar-chart-3.png" width="567" >}}

またグラフの各要素に対して個別にフォントを指定するのであれば

```go
//set fonts
font, err := vg.MakeFont(fontName, 10)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
p.Title.Font = font
p.X.Label.Font = font
p.Y.Label.Font = font
p.X.Tick.Label.Font = font
p.Y.Tick.Label.Font = font
p.Legend.Font = font
```

のように書ける。

フォントの読み込みには [`github.com/golang/freetype/truetype`] パッケージを使っているのだが，どうも TTC (TrueType Collections) ファイルには対応していない模様。
たとえば [Ubuntu] 19.10 に収録されている NOTO フォントを使おうとしたが

```go
b, err := ioutil.ReadFile("/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
ft, err := truetype.Parse(b)
if err != nil {
    fmt.Fprintln(os.Stderr, err) //Output: freetype: invalid TrueType format: bad TTF version
    return
}
```

実行時エラーになってしまった。
ちなみに OTF ファイルにも対応してない。
まぁ，文章を書くんじゃないから [IPA] フォントでも十分だよね。

## ブックマーク

- [gonum · GitHub](https://github.com/gonum)
- [Golangで棒グラフを描く - 逆さまにした](https://cipepser.hatenablog.com/entry/2017/03/25/104301)
- [Golangでローソク足チャートを描くパッケージを書いた - 逆さまにした](https://cipepser.hatenablog.com/entry/2017/07/16/113619)
- [Golangでグラフを描く - Qiita](https://qiita.com/yutsuki/items/7de97e09289a915f86b9)
- [Go言語で折れ線グラフや棒グラフを描く - Qiita](https://qiita.com/RuyPKG/items/0a569953e9e24870f527)
- [TeX Live 2020 へのアップグレード]({{< ref "/remark/2020/04/upgrade-texlive-2020.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[gnuplot]: http://www.gnuplot.info/ "gnuplot homepage"
[Gonum]: https://www.gonum.org/
[`plot`]: https://pkg.go.dev/gonum.org/v1/plot "plot package · go.dev"
[`plot/vg`]: https://pkg.go.dev/gonum.org/v1/plot/vg "vg package · go.dev"
[`vg`]: https://pkg.go.dev/gonum.org/v1/plot/vg "vg package · go.dev"
[`github.com/golang/freetype/truetype`]: https://godoc.org/github.com/golang/freetype/truetype "truetype - GoDoc"
[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[IPA]: https://ipafont.ipa.go.jp/ "IPAexフォント/IPAフォント"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
