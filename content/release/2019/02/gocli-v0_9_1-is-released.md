+++
title = "spiegel-im-spiegel/gocli v0.9.1 をリリースした"
date = "2019-02-09T22:00:33+09:00"
description = "対話モードで動かしたいツールがあって，それ用に使えるサブパッケージ gocli/prompt を追加してみた。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "package", "golang", "cli" ]

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

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
