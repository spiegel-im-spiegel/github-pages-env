+++
title = "Go Playground でお絵描き"
date =  "2020-04-24T14:33:05+09:00"
description = "今回はストレス発散な話。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "image" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

全国緊急事態宣言の折，皆様いかがお過ごしでしょうか。
完全失業中の私は，家族からは白い目で見られ，かといって気晴らしの外出も自粛ムードでままならず，涙目状態です。
どなたか松江で仕事ください。

というわけで，今回はストレス発散な話。
つかですね，

- [The Go Playgroundの実行結果に画像を出力する - My External Storage](https://budougumi0617.github.io/2020/04/22/render-image-on-go-playground/)

という記事を見て自分でも試してみようかな，と。

[Go Playground] で {{< quote lang="en" >}}Display image{{< /quote >}} を選択すると以下のコードが表示され，実行すると favicon 画像データが表示される。

{{< fig-img src="./go-playground-1.png" link="https://play.golang.org/p/XN6x3L23Vok" title="Display image: The Go Playground" width="975" >}}

画像表示の核心部分は以下のコードのようだ。

```go
// displayImage renders an image to the playground's console by
// base64-encoding the encoded image and printing it to stdout
// with the prefix "IMAGE:".
func displayImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	fmt.Println("IMAGE:" + base64.StdEncoding.EncodeToString(buf.Bytes()))
}
```

[`image`]`.Image` インスタンスを PNG 形式でエンコードし，さらに BASE64 形式で文字符号化する。
これに `"IMAGE:"` を頭にくっつけて「表示」すればいいようだ。
当然ながら [Go Playground] 特有の機能なので，他の環境では機能しない（機能されては困る`w`）。

リモートのファイルを表示できれば面白かったのだが，残念ながら [Go Playground] では TCP/IP によるリモートアクセスが潰されている（まぁアクセスできたら踏み台し放題になっちゃうからしないんだろうけど）。

そこで画像データをあらかじめ文字列符号化してソースコードに埋め込めば

```go
var imageFromBase64 = `...`

func getimage() (image.Image, error) {
	img, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageFromBase64)))
	return img, err
}

func main() {
	img, err := getimage()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	displayImage(img)
}
```

てな感じで任意の画像を表示できるだろう[^img1]。
たとえば，こんな感じ。

[^img1]: 大きな画像を [Go Playground] で表示しようとすると失敗（タイムアウト？）する場合がある。また複数の画像は表示できないっぽい。

{{< fig-img src="./go-playground-2.png" link="https://play.golang.org/p/LXmxkAV0z_M" title="Display image: The Go Playground" width="975" >}}

ここで画像データのスケール変換をやってみる。
以前書いた「[Go 言語で画像のサイズ変更：定義済み draw.Scaler の比較]({{< relref "./resize-image-2.md" >}})」を参考に `main()` 関数を以下のように書き換える。

```go
func scale(src image.Image, rect image.Rectangle, scaler draw.Scaler) image.Image {
	dst := image.NewRGBA(rect)
	scaler.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

func main() {
	img, err := getimage()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//scale down
	displayImage(scale(img, image.Rect(0, 0, 80, 80), draw.NearestNeighbor))
}
```

いちばん（速くて）変換性能が悪い [`draw`]`.NearestNeighbor` を使ってみた（笑）

結果はこんな感じ。

{{< fig-img src="./go-playground-3.png" link="https://play.golang.org/p/sJbWR9q2LUJ" title="Display image: The Go Playground" width="975" >}}

よーし，うむうむ，よーし。

## ブックマーク

- [Go 言語で画像のサイズを変更する]({{< relref "./resize-image.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Go Playground]: https://play.golang.org/ "The Go Playground"
[`image`]: https://pkg.go.dev/image "image package · go.dev"
[`draw`]: https://pkg.go.dev/golang.org/x/image/draw "draw package · go.dev"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
