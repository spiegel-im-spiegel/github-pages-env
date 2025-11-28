+++
title = "Web ページのメタデータを取得する webinfo/webinfo をリリースした"
date =  "2025-11-28T19:26:26+09:00"
description = "完全に自分用のパッケージなので汎用性はあまり考慮していない。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "package", "module", "web" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

私は自分が作ってる各プロダクトの中で特定の機能が3つ以上重複した場合は別パッケージとして切り出すことにしている。
というわけで Web ページのメタデータを取得する機能をパッケージとして切り出した。

- [goark/webinfo: Extract metadata and structured information from web pages](https://github.com/goark/webinfo)

要するに，完全に自分用のパッケージなので汎用性はあまり考慮していない。

使い方はこんな感じ。

```go
package main

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "os"

    "github.com/goark/webinfo"
)

func errorOut(err error) {
    fmt.Fprintf(os.Stderr, "%+v\n", err)
}

func main() {
    ctx := context.Background()
    // Fetch metadata for a page
    info, err := webinfo.Fetch(ctx, "https://text.baldanders.info/", "") // empty UA uses default
    if err != nil {
        errorOut(err)
        return
    }
    // output metadata as pretty JSON
    b, err := json.MarshalIndent(info, "", "  ")
    if err != nil {
        errorOut(err)
        return
    }
    if _, err := io.Copy(os.Stdout, bytes.NewBuffer(b)); err != nil {
        errorOut(err)
        return
    }
    fmt.Println()
    fmt.Println()

    // Download main image to permanent file in current directory
    imgPath, err := info.DownloadImage(ctx, ".", false)
    if err != nil {
        errorOut(err)
        return
    }
    fmt.Println("image saved:", imgPath)

    // Download thumbnail to permanent file in current directory
    thumbPath, err := info.DownloadThumbnail(ctx, ".", 150, false) // 150px width
    if err != nil {
        errorOut(err)
        return
    }
    fmt.Println("thumbnail saved:", thumbPath)
}
```

[`encoding/json/v2`] パッケージを使えば `MarshalWrite()` 関数で直接標準出力に出せるんだけど，まだ実験的実装なので，今回は従来の [`encoding/json`] パッケージを使っている。

これを実行するとこんな感じ。

```text
$ go run sample.go
{
  "url": "https://text.baldanders.info/",
  "location": "https://text.baldanders.info/",
  "canonical": "https://text.baldanders.info/",
  "title": "text.Baldanders.info",
  "description": "帰ってきた「しっぽのさきっちょ」",
  "image_url": "https://text.baldanders.info/images/attention/site.jpg",
  "user_agent": "Mozilla/5.0 (Windows NT 6.1; rv:11.0) Gecko/20100101 Firefox/11.0"
}

image saved: site.jpg
thumbnail saved: site-thumb.jpg
```

ダウンロードした `site.jpg` がこれ。

{{< fig-img src="./site.jpg" title="site.jpg" link="./site.jpg" width="2230" >}}

`site-thumb.jpg` はこれ。

{{< fig-img src="./site-thumb.jpg" title="site-thumb.jpg" link="./site-thumb.jpg" >}}

`site-thumb.jpg` は `site.jpg` を幅150ピクセルに[リサイズ]({{< ref "/golang/resize-image.md" >}} "Go 言語で画像のサイズを変更する")して保存している。

[`README.md`](https://github.com/goark/webinfo/blob/main/README.md) ファイルや [godoc ページ][`github.com/goark/webinfo`]を見ると分かるけど，私にしてはドキュメントが煩い（笑） 実はこれ，ほぼ [GitHub Copilot] (model: GPT-5 mini を選択) に書かせたものなのよね。
今回，新規に書いたメソッドとテストおよびドキュメントは Copilot に丸投げしてみた。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

[前にも書いた]({{< ref "/remark/2025/11/released-the-translation-of-the-cc-doc.md" >}} "“Pay-To-Crawl Issue Brief” の翻訳を公開した")けど，プロンプトさえ上手くハマれば7割以上は使えるコードを書いてくれる。
でも結局のところ，残りの3割なり2割なり1割なりがインパクトの大きい部分だったりするのだ。
案の定 lint にかけると，特に [`gosec`][`securego/gosec`] あたりにがっつり怒られた。
そんでチマチマ直したりする。

あとこれは感覚的なもの言いで申し訳ないんだけど，記述に一貫性がないんだよね。
なんというか，面識のない複数の他人がバラバラに書き込んだような感じ。
だからコードがちょっと読みにくいし，全体で見て合理的ではない記述もある。
でも，記述が間違ってないか確認するためにはきちんと精査しないといけないわけで，そこも微妙なんだよな。

テストコードについては最初「動けばいいや」ということで，大きく間違ってない限り Copilot が書いたコードをそのまま採用していた。
Copilot に「○○のケースを追加して」って書くと，記述が面倒くさいケースでも律儀に書いてくれるんだけど（これは有り難い），書いてる様子を VSCode 上で眺めてると「あぁ，これは正しく動かんだろうな」みたいなテストコードを最初書いて，実際にテストを動かしたらやっぱりエラーになって，そんでまた修正を入れて... ってのを何度か繰り返してて「ずいぶん人間臭いことするなぁ」と思ったりした。

だいたい収束したところで PR を出して Copilot にレビューさせたのだが，テストコードをけちょんけちょんに貶された（笑） いや，同じ Copilot が書いたコードやろがい。
で，やっぱり人間である私が改めてテストコードを精査し指摘部分を書き直す作業をチマチマと。

コメントやドキュメントについてはかなり助かった。
最初に書いたように「ここまで要るか？」ってくらい詳細に書いてくれるので，逆に Copilot が書いたコメントやドキュメントが設計書代わりになっている。
特に [`.github/copilot-instructions.md`](https://github.com/goark/webinfo/blob/main/.github/copilot-instructions.md) を Copilot が書き出してくれたのは助かった。
私はそれを見て実装の過不足を確認できるし Copilot も作業の精度を上げることができる。
仕様・設計上の変更についても，人間がコードを修正して Copilot に言えば実装に合わせてドキュメントも直してくれるので，実装とドキュメントが乖離するといった「あるある」なことは，ある程度緩和できるのだ。

Copilot の使い方も少しずつ分かってきたし，今回作ったパッケージを既存のプロダクトに組み込んだら，いよいよ新しいツールの制作に入るか。

[`encoding/json`]: https://pkg.go.dev/encoding/json "json package - encoding/json - Go Packages"
[`encoding/json/v2`]: https://pkg.go.dev/encoding/json/v2 "json package - encoding/json/v2 - Go Packages"
[`github.com/goark/webinfo`]: https://pkg.go.dev/github.com/goark/webinfo "https://pkg.go.dev/github.com/goark/webinfo"
[`securego/gosec`]: https://github.com/securego/gosec "securego/gosec: Go security checker"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"

## 参考図書

{{% review-paapi "4814401191" %}} <!-- 初めてのGo言語 第2版 -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
