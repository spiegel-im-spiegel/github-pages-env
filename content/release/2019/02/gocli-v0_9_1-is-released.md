+++
title = "spiegel-im-spiegel/gocli v0.9.1 をリリースした"
date = "2019-02-09T22:00:33+09:00"
description = "対話モードで動かしたいツールがあって，それ用に使えるサブパッケージ gocli/prompt を追加してみた。"
image = "/images/attention/tools.png"
tags  = [ "package", "golang", "cli" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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

CLI で使える細々した機能を詰め合わせた自作パッケージ [spiegel-im-spiegel/gocli] の v0.9.1 をリリースした。

- [spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface](https://github.com/spiegel-im-spiegel/gocli)

対話モードで動かしたいツールがあって，それ用に使えるサブパッケージ [`gocli`]`/prompt` を追加してみた。

こんな感じに書ける。

```go
package main

import (
    "fmt"
    "os"

    "github.com/spiegel-im-spiegel/gocli/prompt"
    "github.com/spiegel-im-spiegel/gocli/rwi"
)

func main() {
    p := prompt.New(
        rwi.New(
            rwi.WithReader(os.Stdin),
            rwi.WithWriter(os.Stdout),
        ),
        func(s string) (string, error) {
            if s == "q" || s == "quit" {
                return "quit prompt", prompt.ErrTerminate
            }
            runes := []rune(s)
            for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
                runes[i], runes[j] = runes[j], runes[i]
            }
            return string(runes), nil
        },
        prompt.WithPromptString("Sample> "),
        prompt.WithHeaderMessage("Input 'q' or 'quit' to stop"),
    )
    if !p.IsTerminal() {
        fmt.Fprintln(os.Stderr, "not terminal (or pipe?)")
        return
    }
    if err := p.Run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
```

このコードでは入力した文字列を逆順に並べ替えて出力する。
“`q`” または “`quit`” を入力するとプログラムを終了する。 

これを動かすとこんな感じ。

```text
$ go run prompt/sample/sample.go
Input 'q' or 'quit' to stop
Sample> abcdef
fedcba
Sample> quit
quit prompt
```

ちなみに [github.com/atotto/clipboard] と組み合わせて

{{< highlight go "hl_lines=7 26-27" >}}
package main

import (
    "fmt"
    "os"

    "github.com/atotto/clipboard"
    "github.com/spiegel-im-spiegel/gocli/prompt"
    "github.com/spiegel-im-spiegel/gocli/rwi"
)

func main() {
    p := prompt.New(
        rwi.New(
            rwi.WithReader(os.Stdin),
            rwi.WithWriter(os.Stdout),
        ),
        func(s string) (string, error) {
            if s == "q" || s == "quit" {
                return "quit prompt", prompt.ErrTerminate
            }
            runes := []rune(s)
            for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
                runes[i], runes[j] = runes[j], runes[i]
            }
            res := string(runes)
            return res, clipboard.WriteAll(res)
        },
        prompt.WithPromptString("Sample> "),
        prompt.WithHeaderMessage("Input 'q' or 'quit' to stop"),
    )
    if !p.IsTerminal() {
        fmt.Fprintln(os.Stderr, "not terminal (or pipe?)")
        return
    }
    if err := p.Run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
{{< /highlight >}}

とか書けば標準出力への出力と同時にクリップボードにも出力される。

[spiegel-im-spiegel/gocli] パッケージは CLI を組む際に（主に自分が）便利な細々とした機能を収録している。
他の人には使いにくいかもしれないし大した内容でもないため [CC0](https://creativecommons.org/publicdomain/zero/1.0/ "Creative Commons — CC0 1.0 Universal") で公開している。
一切の権利を放棄しているので自由に持っていって弄っていただいて構わない。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/gocli]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[`gocli`]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[github.com/atotto/clipboard]: https://github.com/atotto/clipboard "atotto/clipboard: clipboard for golang"

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
