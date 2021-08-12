+++
title = "Go 言語で画像のサイズを変更する"
date = "2019-02-02T17:54:21+09:00"
description = " ちょっと画像データを弄る機会があったので，やったことを忘れないうちに記しておく。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "image" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は小ネタ。
ちょっと[画像データを弄る機会があった]({{< ref "/remark/2019/01/escape-from-flickr.md" >}} "Flickr から写真を引き揚げました")ので，やったことを忘れないうちに記しておく。

## 画像の Width/Height の値を取得する

画像のサイズを変更するにはまず元の画像のサイズが分かっていないといけない。
これは標準の [`image`] パッケージを使って実装できる。
いきなりコードを書くとこんな感じ。

```go
package main

import (
    "flag"
    "fmt"
    "image"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
    "os"
)

func main() {
    flag.Parse()
    if flag.NArg() != 1 {
        fmt.Fprintln(os.Stderr, os.ErrInvalid)
        return
    }
    
    //open image file
    file, err := os.Open(flag.Arg(0)) //maybe file path
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer file.Close()

    //decode image
    img, t, err := image.Decode(file)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println("Type of image:", t)

    //rectange info of image
    rct := img.Bounds()
    fmt.Println("Width:", rct.Dx())
    fmt.Println("Height:", rct.Dy())
}
```

[`image`]`.Decode()` 関数で画像データをデコードし [`image`]`.Image.Bounds()` 関数で画像全体の矩形情報（[`image`]`.Rectangle`）を取得する[^img1]。
ポイントは [`image`]`.Decode()` 関数が factory method のように機能していて画像フォーマットを意識する必要はないという点。
ではどこで依存（dependency）を注入しているかというと，パッケージのインポート

[^img1]: [`image`]`.Image` は interface 型で抽象化されている。

```go
import (
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
)
```

の部分。
この場合は GIF, JPEG, PNG 形式に対応していることを示す。
もちろんこれらと互換性のある独自パッケージを使ってもよい[^register1]。

[^register1]: 厳密に言うと [`image`]`.RegisterFormat()` 関数で登録する。例えば [`image/jpeg`] パッケージの場合は `init()` 関数内で [`image`]`.RegisterFormat()` 関数が呼ばれている。このようなデザイン・パターンは [Go 言語]ではよく見られるので覚えておくとよいだろう（参考：[Hash 値を計算するパッケージを作ってみた]({{< relref "./calculating-hash-value.md" >}})）。

では早速

{{< fig-img src="https://photo.baldanders.info/flickr/image/7444837_o.jpg" title="Frog" link="https://photo.baldanders.info/flickr/7444837/" width="640" >}}

の画像を読み込んでみよう。

```text
$ go run sample1.go 7444837_o.jpg
Type of image: jpeg
Width: 640
Height: 480
```

よーし，うむうむ，よーし。

## 画像のサイズを変更する

次に画像のサイズを変更してみよう。
ググってみると [github.com/nfnt/resize] を使った例をよく見かけるが，件のリポジトリを見てみると

{{< fig-quote title="nfnt/resize: Pure golang image resizing" link="https://github.com/nfnt/resize" lang="en" >}}
<q>This package is no longer being updated! Please look for alternatives if that bothers you.</q>
{{< /fig-quote >}}

と書かれていて2018年2月を最後に活動を停止しているようである。
代わりのパッケージとして

- [golang.org/x/image/draw](https://godoc.org/golang.org/x/image/draw)
- [gopkg.in/gographics/imagick.v2/imagick](https://godoc.org/gopkg.in/gographics/imagick.v2/imagick)

が[紹介](https://github.com/nfnt/resize/issues/63 "I'm no longer updating this package. · Issue #63 · nfnt/resize")されている。

そこで今回は [golang.org/x/image/draw](https://godoc.org/golang.org/x/image/draw) パッケージを使って画像のサイズ変更を行ってみる。
これもいきなりコードを書くとこんな感じ。

```go
package main

import (
    "flag"
    "fmt"
    "image"
    "image/gif"
    "image/jpeg"
    "image/png"
    "os"

    "golang.org/x/image/draw"
)

func main() {
    flag.Parse()
    if flag.NArg() != 2 {
        fmt.Fprintln(os.Stderr, os.ErrInvalid)
        return
    }

    //open original image file
    src, err := os.Open(flag.Arg(0)) //maybe src file path
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer src.Close()

    //decode image
    imgSrc, t, err := image.Decode(src)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println("Type of image:", t)

    //rectange of image
    rctSrc := imgSrc.Bounds()
    fmt.Println("Width:", rctSrc.Dx())
    fmt.Println("Height:", rctSrc.Dy())

    //scale down by 1/2
    imgDst := image.NewRGBA(image.Rect(0, 0, rctSrc.Dx()/2, rctSrc.Dy()/2))
    draw.CatmullRom.Scale(imgDst, imgDst.Bounds(), imgSrc, rctSrc, draw.Over, nil)

    //create resized image file
    dst, err := os.Create(flag.Arg(1)) //maybe dst file path
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer dst.Close()

    //encode resized image
    switch t {
    case "jpeg":
        if err := jpeg.Encode(dst, imgDst, &jpeg.Options{Quality: 100}); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
    case "gif":
        if err := gif.Encode(dst, imgDst, nil); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
    case "png":
        if err := png.Encode(dst, imgDst); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
    default:
        fmt.Fprintln(os.Stderr, "format error")
    }
}
```

このコードでは Width/Height を半分のサイズにしたイメージをファイル出力している。

前半の画像の読み込み部分は先程と同じ。
その次に [`image`]`.NewRGBA()` 関数で Width/Height を半分のサイズにした空のインスタンスを生成し [`draw`]`.CatmullRom.Scale()` 関数で変換した画像データを流し込んでいる。
その後ファイルに出力するのだが，出力するフォーマットごとに異なる `Encode()` 関数になっている。

[`draw`]`.CatmullRom` はサイズ変換のための [`draw`]`.Scaler` インタフェースを持つ [`draw`]`.Kernel` 型インスタンスで，以下のように定義されている。

```go
var (
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

[`draw`]`.CatmullRom` はどちらかというと品質優先の [`draw`]`.Scaler` のようだが，他にもいくつか定義済みのインスタンスがある。
性能について詳しくは[次回の記事]({{< relref "./resize-image-2.md" >}} "Go 言語で画像のサイズ変更：定義済み draw.Scaler の比較")を参考にしてほしい。

では，実際に動かしてみよう。

```text
$ go run sample2.go 7444837_o.jpg out.jpg
Type of image: jpeg
Width: 640
Height: 480
```

出力結果は以下の通り。

{{< fig-img src="./out.jpg" title="out.jpg" link="./out.jpg" >}}

うむうむ。
いい感じである。

## ブックマーク

- [GolangでアニメーションGifをリサイズする - Qiita](https://qiita.com/from_Unknown/items/40d1947292c53fe7ea74)
- [Scale down by 2x using x/image/draw · GitHub](https://gist.github.com/aclements/599107a2e3f187f8a2c0)
- [Golang resize png images using different interpolations · GitHub](https://gist.github.com/logrusorgru/570d64fd6a051e0441014387b89286ca)
- [大量のPNG画像を標準パッケージのみで2.4倍高速に処理する - Qiita](https://qiita.com/cia_rana/items/a56269b11e909b3a246d)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`image`]: https://golang.org/pkg/image/ "image - The Go Programming Language"
[`image/jpeg`]: https://golang.org/pkg/image/jpeg/ "jpeg - The Go Programming Language"
[`draw`]: https://godoc.org/golang.org/x/image/draw "draw - GoDoc"
[github.com/nfnt/resize]: https://github.com/nfnt/resize "nfnt/resize: Pure golang image resizing"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
