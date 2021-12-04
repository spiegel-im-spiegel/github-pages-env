+++
title = "spiegel-im-spiegel/ml v0.3.0 をリリースした"
date =  "2021-01-03T16:12:38+09:00"
description = "ひょっとして nyaosorg/go-readline-ny パッケージ使ったら CUI の簡易プロンプトがもっと簡単に実装できるんちゃうん？"
image = "/images/attention/tools.png"
tags  = [ "tools", "golang" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

...ついカッとなってやった。
反省はしない...

みなさま，あけましておめでとうございます。
本年も「書きたくないときには書かない」「他人の評価など気にしない」をモットーに，ゆるゆるとやっていく所存です。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}，正月早々 Zenn で

- [「コマンドラインシェル？？？　誰でも作れますよ」](https://zenn.dev/zetamatta/articles/d7b76ff6535d7d)

という記事を眺めてて

{{< div-box type="markdown" >}}
「あれ？ ひょっとして [nyaosorg/go-readline-ny][go-readline-ny] パッケージ使ったら CUI の簡易プロンプトがもっと簡単に実装できるんちゃうん？」

[go-readline-ny]: https://github.com/nyaosorg/go-readline-ny "nyaosorg/go-readline-ny: Readline library for golang , used in nyagos"
{{< /div-box >}}

と思いついてしまい，（たぶん私しか使ってないであろう）[Markdown 形式のリンクを生成するツール][ml]に手を入れることにした。
その結果を v0.3.0 としてリリースしている。

- [Release v0.3.0 · spiegel-im-spiegel/ml · GitHub](https://github.com/spiegel-im-spiegel/ml/releases/tag/v0.3.0)

今回の変更は以下の3つ。
 
1. リポジトリ名を `mklink` から `ml` に変更する
2. [spiegel-im-spiegel/ml][ml] リポジトリ直下に `main.go` を移動し，コマンドライン・ツールとして構成し直す
3. 対話モードでの入力を [nyaosorg/go-readline-ny][go-readline-ny] パッケージで書き直す

最初のはメインマシンを Ubuntu に換装した関係で先延ばしにしていたやつ。
Windows の内部コマンドに `mklink` てのがあって（シンボリックリンク操作のコマンド），それと名前が被ってたのだ。
実際には GitHub のリポジトリ設定から簡単にリネームできた（簡単ならとっととやれってば）。
旧 URL からのリダイレクトもやってくれている。
感謝。

2番めは，近年私がよくやる変更。
昔は「[Go] コード用の外部パッケージとして書いて，その実装例としてコマンドラインツールも書く」というスタンスで構成していたが，機能がニッチ過ぎて汎用で再利用しづらいため「じゃあ最初からコマンドラインツールとして構成すればいいぢゃん」と考えを改めた。

一応

```go
package main

import (
    "context"
    "fmt"
    "io"
    "os"

    "github.com/spiegel-im-spiegel/ml/makelink"
)

func main() {
    lnk, err := makelink.New(context.Background(), "https://git.io/vFR5M")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    _, _ = io.Copy(os.Stdout, lnk.Encode(makelink.StyleMarkdown))
    // Output:
    // [GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/ml)
}
```

のように [Go] のコードに組み込むことも可能。

3番目の [nyaosorg/go-readline-ny][go-readline-ny] パッケージの組み込みはマジでやってよかった。
ざっと見た感じ，このパッケージの特徴は以下の通り。

- エコーバック（？）の `Writer` を指定できる
- 簡易ヒストリ機能を付けられる
- `Ctrl+C` および `Ctrl+D` を正しく拾ってエラー（`readline.CtrlC` および `io.EOF`）として返してくれる（上位レイヤでの SIGNAL 操作が不要）

たとえばこんな感じに書けるらしい。

{{< fig-quote type="markdown" class="nobox" title="nyaosorg/go-readline-ny: Readline library for golang , used in nyagos" link="https://github.com/nyaosorg/go-readline-ny" lang="en" >}}
```go
package main

import (
    "context"
    "fmt"
    "os"
    "os/exec"
    "strings"

    "github.com/mattn/go-colorable"

    "github.com/nyaosorg/go-readline-ny"
    "github.com/nyaosorg/go-readline-ny/simplehistory"
)

func main() {
    history := simplehistory.New()

    editor := readline.Editor{
        Prompt:  func() (int, error) { return fmt.Print("$ ") },
        Writer:  colorable.NewColorableStdout(),
        History: history,
    }
    fmt.Println("Tiny Shell. Type Ctrl-D to quit.")
    for {
        text, err := editor.ReadLine(context.Background())

        if err != nil {
            fmt.Printf("ERR=%s\n", err.Error())
            return
        }

        fields := strings.Fields(text)
        if len(fields) <= 0 {
            continue
        }
        cmd := exec.Command(fields[0], fields[1:]...)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Stdin = os.Stdin

        cmd.Run()

        history.Add(text)
    }
}
```
{{< /fig-quote >}}

というわけで，拙作の [spiegel-im-spiegel/ml][ml] でも対話モードのプロンプトに簡易ヒストリが使えるようにした。
めっさ便利！

[nyaosorg/go-readline-ny][go-readline-ny] パッケージは元々 [NYAGOS] をターゲットに書かれたもののようだが，私の Ubuntu 環境でも問題なく動作している。

うんうん。
よかったよかった。

## ブックマーク

- [Markdown 形式のリンクを生成するツールを作ってみた]({{< ref "/golang/make-link-with-markdown-format.md" >}})

[Go]: https://go.dev/
[ml]: https://github.com/spiegel-im-spiegel/ml "spiegel-im-spiegel/ml: Make Link with Markdown Format"
[go-readline-ny]: https://github.com/nyaosorg/go-readline-ny "nyaosorg/go-readline-ny: Readline library for golang , used in nyagos"
[NYAGOS]: https://github.com/nyaosorg/nyagos "nyaosorg/nyagos: NYAGOS - The hybrid Commandline Shell betweeeeeeen UNIX & DOS"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
