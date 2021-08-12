+++
title = "画像データを連結してみる"
date =  "2020-05-04T13:20:55+09:00"
description = "Go 言語っぽいトピックはなし。今回もお遊びの小ネタで。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "image" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回もお遊びの小ネタで。
複数の画像データを連結してひとつの画像データにすることを考えてみる。

具体的には `image-1.png` と `image-2.png` の2つの画像データを使って

1. 元の画像データから各々 [`image`]`.Image` を取得する
2. 各 [`image`]`.Image` から矩形情報を抽出し，空の結合 [`image`]`.Image` を生成する
3. 空の結合 [`image`]`.Image` に元の [`image`]`.Image` を貼り付ける
4. 結合 [`image`]`.Image` を PNG データとして出力する

といった手順。
図にすると

{{< fig-img src="./concatenate-images.png" link="./concatenate-images.png" width="718" >}}

といった感じか。

それでは順にコードを書いてみよう。

## 画像データから image.Image を取得する

ファイルから [`image`]`.Image` を取得する関数はこんな感じでどうだろう。

```go
import (
	"image"
	_ "image/png"
	"os"
)

func imageFrom(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
```

今回は PNG データのみ取り扱うので [`image`]`/png` パッケージのみインポートしているが，他の形式も取り扱うのであれば各形式のパッケージを（暗黙的に）インポートして「依存の注入」を行えばよい。

ちなみに [`image`]`.Image` は interface 型で

```go
// Image is a finite rectangular grid of color.Color values taken from a color
// model.
type Image interface {
	// ColorModel returns the Image's color model.
	ColorModel() color.Model
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	Bounds() Rectangle
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	At(x, y int) color.Color
}
```

と定義されている。
ここで定義される [`image`]`.Image.Bounds()` メソッドを使えば矩形情報  [`image`]`.Rectangle` が取れるので，ここから画像の幅や高さも分かるというわけ。 

```go {hl_lines=[7,8]}
func main() {
	img, err := imageFrom("image-1.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	rct := img.Bounds()
	fmt.Println("Width:", rct.Dx(), ", height:", rct.Dy())
	//Output:
	//Width: 352 , height: 219
}
```

## 空の結合 image.Image を生成する

まずは元の画像データの [`image`]`.Image` を保持っておくところから始めよう。
こんな感じ。

```go
srcImages := make([]image.Image, 0, len(srcPaths))
for _, path := range srcPaths {
    img, err := imageFrom(path)
    if err != nil {
        return err
    }
    srcImages = append(srcImages, img)
}
```

このタイミングで結合 [`image`]`.Image` の幅と高さも計算してしまおう。

```go {hl_lines=[2,"8-10"]}
srcImages := make([]image.Image, 0, len(srcPaths))
width, height := 0, 0
for _, path := range srcPaths {
    img, err := imageFrom(path)
    if err != nil {
        return err
    }
    rct := img.Bounds()
    width = max(width, rct.Dx())
    height += rct.Dy()
    srcImages = append(srcImages, img)
}
```

ちなみに `max()` 関数は

```go
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```

と定義している[^max1]。

[^max1]: [Go の標準ライブラリには整数型の `Min`/`Max` 関数は用意されていない](https://www.pixelstech.net/article/1559993656-Why-no-max-min-function-for-integer-in-GoLang "Why no max/min function for integer in GoLang | Pixelstech.net")ので，必要に応じて自前で用意する必要がある。

これで，算出した `width`, `height` を使って，空の [`image`]`.Image` を生成できる。
こんな感じ。

```go
dstImage := image.NewRGBA(image.Rect(0, 0, width, height))
```

## 空の結合 image.Image に元の image.Image を貼り付ける

ここまでくれば，あとは機械的な繰り返し作業。

```go
offset := 0
for _, img := range srcImages {
    srcRect := img.Bounds()
    draw.Draw(
        dstImage,
        image.Rect(0, offset, srcRect.Dx(), offset+srcRect.Dy()),
        img,
        image.Point{0, 0},
        draw.Over,
    )
    offset += srcRect.Dy()
}
```

## 結合 image.Image を PNG データとして出力する

結合 [`image`]`.Image` をファイルに出力するにはこんな感じにすればよい。

```go
file, err := os.Create(dstPath)
if err != nil {
    return err
}
defer file.Close()
if err := png.Encode(file, dstImage); err != nil {
    return err
}
```

上のコードは PNG 形式で出力する場合。
各形式へのエンコーディングは（[`image`]`.Decode()` 関数のように）抽象化されていないので，それぞれの形式のパッケージが提供しているエンコーダを使う必要がある（[`image`] パッケージに準拠していれば自作も可能）。

## 実行結果

一連の手順を関数化してみる。
こんな感じ。

```go {hl_lines=["32-69"]}
package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func imageFrom(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func concatImageFiles(dstPath string, srcPaths ...string) error {
	srcImages := make([]image.Image, 0, len(srcPaths))
	width, height := 0, 0
	for _, path := range srcPaths {
		img, err := imageFrom(path)
		if err != nil {
			return err
		}
		rct := img.Bounds()
		width = max(width, rct.Dx())
		height += rct.Dy()
		srcImages = append(srcImages, img)
	}

	dstImage := image.NewRGBA(image.Rect(0, 0, width, height))
	offset := 0
	for _, img := range srcImages {
		srcRect := img.Bounds()
		draw.Draw(
			dstImage,
			image.Rect(0, offset, srcRect.Dx(), offset+srcRect.Dy()),
			img,
			image.Point{0, 0},
			draw.Over,
		)
		offset += srcRect.Dy()
	}

	file, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := png.Encode(file, dstImage); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := concatImageFiles("out.png", "image-1.png", "image-2.png"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
```

これを実行すると以下の画像データが出力される。

{{< fig-img src="./out.png" link="./out.png" title="out.png" >}}

よーし，うむうむ，よーし。

## ブックマーク

- [Go言語で複数の画像を縦に連結する - Qiita](https://qiita.com/tarotaro0/items/7d139b09a0a1ac01d42f)
- [Go 言語で画像のサイズを変更する]({{< relref "./resize-image.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`image`]: https://pkg.go.dev/image "image package · go.dev"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
