+++
title = "Go 言語で画像のサイズ変更：定義済み draw.Scaler の比較"
date = "2019-02-03T11:30:42+09:00"
description = "golang.org/x/image/draw パッケージで定義済みの4つの draw.Scaler について性能比較を試みる。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "image" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]({{< relref "./resize-image.md" >}} "Go 言語で画像のサイズを変更する") [Go 言語]で画像データのサイズ変換を行う手順を紹介したが，その際に使われる [golang.org/x/image/draw](https://godoc.org/golang.org/x/image/draw) パッケージで定義済みの4つの [`draw`]`.Scaler` について性能比較を試みる。

[golang.org/x/image/draw](https://godoc.org/golang.org/x/image/draw) パッケージでは以下の4つの [`draw`]`.Scaler` が定義されている。

```go
var (
    // NearestNeighbor is the nearest neighbor interpolator. It is very fast,
    // but usually gives very low quality results. When scaling up, the result
    // will look 'blocky'.
    NearestNeighbor = Interpolator(nnInterpolator{})

    // ApproxBiLinear is a mixture of the nearest neighbor and bi-linear
    // interpolators. It is fast, but usually gives medium quality results.
    //
    // It implements bi-linear interpolation when upscaling and a bi-linear
    // blend of the 4 nearest neighbor pixels when downscaling. This yields
    // nicer quality than nearest neighbor interpolation when upscaling, but
    // the time taken is independent of the number of source pixels, unlike the
    // bi-linear interpolator. When downscaling a large image, the performance
    // difference can be significant.
    ApproxBiLinear = Interpolator(ablInterpolator{})

    // BiLinear is the tent kernel. It is slow, but usually gives high quality
    // results.
    BiLinear = &Kernel{1, func(t float64) float64 {
        return 1 - t
    }}

    // CatmullRom is the Catmull-Rom kernel. It is very slow, but usually gives
    // very high quality results.
    //
    // It is an instance of the more general cubic BC-spline kernel with parameters
    // B=0 and C=0.5. See Mitchell and Netravali, "Reconstruction Filters in
    // Computer Graphics", Computer Graphics, Vol. 22, No. 4, pp. 221-228.
    CatmullRom = &Kernel{2, func(t float64) float64 {
        if t < 1 {
            return (1.5*t-2.5)*t*t + 1
        }
        return ((-0.5*t+2.5)*t-4)*t + 2
    }}
)
```

では，早速コードを組んでみよう。
こんな感じでどうだろう。

```go
package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"golang.org/x/image/draw"
)

func scale(src image.Image, rect image.Rectangle, scaler draw.Scaler) image.Image {
	dst := image.NewRGBA(rect)
	scaler.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

func exportFile(fname string, img image.Image) error {
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return err
	}
	return nil
}

func main() {
	scalers := []struct {
		fname  string
		scaler draw.Scaler
	}{
		{"NearestNeighbor.png", draw.NearestNeighbor},
		{"ApproxBiLinear.png", draw.ApproxBiLinear},
		{"BiLinear.png", draw.BiLinear},
		{"CatmullRom.png", draw.CatmullRom},
	}

	//open original image file
	src, err := os.Open("circle.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer src.Close()

	//original image
	imgSrc, _, err := image.Decode(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//scale down
	rect := image.Rect(0, 0, 320, 320) //size of scaled image
	for _, scaler := range scalers {
		start := time.Now()
		imgDst := scale(imgSrc, rect, scaler.scaler)
		end := time.Now()
		fmt.Println(scaler.fname, end.Sub(start))
		if err := exportFile(scaler.fname, imgDst); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
}
```

元画像の `circle.png` は [github.com/nfnt/resize] のサンプル画像を拝借した。

{{< fig-img src="./circle.png" title="circle.png" link="./circle.png" width="1000" >}}

ブラウザ経由だと自動でリサイズしてしまうので分かりづらいだろうけど 1,000×1,000 の同心円を描いた画像である。
これを 320×320 に変換する。

では実行してみよう（私の PC はかなり遅いので変換時間は参考程度でお願いします）。

```text
$ go run scaler.go
NearestNeighbor.png 2.0001ms
ApproxBiLinear.png 2.0001ms
BiLinear.png 16.001ms
CatmullRom.png 21.0012ms
```

変換結果は以下の通り。

{{< fig-img src="./NearestNeighbor.png" title="NearestNeighbor.png" link="./NearestNeighbor.png" >}}

{{< fig-img src="./ApproxBiLinear.png" title="ApproxBiLinear.png" link="./ApproxBiLinear.png" >}}

{{< fig-img src="./BiLinear.png" title="BiLinear.png" link="./BiLinear.png" >}}

{{< fig-img src="./CatmullRom.png" title="CatmullRom.png" link="./CatmullRom.png" >}}

[`draw`]`.NearestNeighbor` と [`draw`]`.ApproxBiLinear` は変換速度は早いが品質としては実用に耐える感じではなさそうである。
サムネイルとかを作成するならこちらでもいいかも知れない。

品質でいうならやはり [`draw`]`.BiLinear` や [`draw`]`.CatmullRom` が妥当かな。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`draw`]: https://godoc.org/golang.org/x/image/draw "draw - GoDoc"
[github.com/nfnt/resize]: https://github.com/nfnt/resize "nfnt/resize: Pure golang image resizing"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
