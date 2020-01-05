+++
title = "Gocli Package v0.8.0 Released"
date = "2018-03-08T22:28:19+09:00"
description = "Ruby の Dir.glob にあるようなディレクトリの再帰検索ができる file.Glob() 関数を作ってみた。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "package", "cli", "filepath" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

この前 [v0.7.0 をリリースした]({{< ref "/release/2018/03/gocli-package-v0_7_0-released.md" >}})ばかりだが，舌の根も乾かぬうちに v0.8.0 をリリースしてしまった。

- [Release v0.8.0 Released · spiegel-im-spiegel/gocli](https://github.com/spiegel-im-spiegel/gocli/releases/tag/v0.8.0)

今回はファイル・ディレクトリ操作用の [`gocli`]`/file` サブパッケージを追加した。

実は [mattn/jvgrep] を使ってみてディレクトリの再帰検索はなかなか便利なことに気づき，自分でも実装してみようと思ったのだ。

[Go 言語]標準の [`filepath`]`.Glob()` 関数はなかなか性能がよくて，パス検索に以下のワイルドカードが使える。

{{% fig-gen type="markdown" title="filepath - The Go Programming Language" link="https://golang.org/pkg/path/filepath/" lang="en" %}}
```
pattern:
    { term }
term:
    '*'         matches any sequence of non-Separator characters
    '?'         matches any single non-Separator character
    '[' [ '^' ] { character-range } ']'
                character class (must be non-empty)
    c           matches character c (c != '*', '?', '\\', '[')
    '\\' c      matches character c

character-range:
    c           matches character c (c != '\\', '-', ']')
    '\\' c      matches character c
    lo '-' hi   matches character c for lo <= c <= hi
```
{{% /fig-gen %}}

しかし，残念ながら Ruby の [`Dir`]`.glob` にあるようなディレクトリの再帰検索（`**/`）は用意されていない。
そこで今回  [`filepath`]`.Glob()` と [`filepath`]`.Walk()` を組み合わせる形で `file.Glob()` 関数を作ってみた。

こんなふうに使える。

```go
package main

import (
    "fmt"
    "os"

    "github.com/spiegel-im-spiegel/gocli/file"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, os.ErrInvalid)
    }

    for _, path := range os.Args[1:] {
        paths := file.Glob(path, file.NewGlobOption())
        if len(paths) > 0 {
            for _, path := range paths {
                fmt.Println(path)
            }
        }
    }
}
```

これを実行するとこんな感じになる。

```text
$ go run /path/to/search.go  **/*.lua
nyagos.d\aliasandset.lua
nyagos.d\aliases.lua
nyagos.d\backquote.lua
nyagos.d\box.lua
nyagos.d\brace.lua
nyagos.d\catalog\autocd.lua
nyagos.d\catalog\autols.lua
nyagos.d\catalog\cho.lua
nyagos.d\catalog\dollar.lua
nyagos.d\catalog\ezoe.lua
nyagos.d\catalog\git.lua
nyagos.d\catalog\gogit.lua
nyagos.d\catalog\neco.lua
nyagos.d\catalog\nyagosini.lua
nyagos.d\catalog\peco.lua
nyagos.d\catalog\subcomplete.lua
nyagos.d\cdlnk.lua
nyagos.d\comspec.lua
nyagos.d\lns.lua
nyagos.d\open.lua
nyagos.d\start.lua
nyagos.d\su.lua
nyagos.d\suffix.lua
nyagos.d\swapstdfunc.lua
nyagos.d\trash.lua
nyagos.d\use.lua
```

[Go 言語]製なので，フォルダの区切り文字には `'/'` と `'\'` の両方が使える。
また出力結果の内，フォルダについては末尾にフォルダ区切り文字（`'/'` または `'\'`）を付加している。
その他，細かい使い方については `glob_test.go` を参照のこと。

本当はもう少しスマートなコードを書きたかったんだけど，行き当たりばったりで書いてたらエラく汚いコードになってしまった。
まぁ，これが今の私の実力，ということで。
そのうち refactoring できる機会もあるかもしれない。

[spiegel-im-spiegel/gocli] パッケージは CLI (Command-Line Interface) を組む際に（主に自分が）便利な細々とした機能を収録している。
他の人には使いにくいかもしれないし大した内容でもないため [CC0](https://creativecommons.org/publicdomain/zero/1.0/ "Creative Commons — CC0 1.0 Universal") で公開している。
一切の権利を放棄しているので自由に持っていって弄っていただいて構わない。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/gocli]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[`gocli`]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[mattn/jvgrep]: https://github.com/mattn/jvgrep "mattn/jvgrep: grep for japanese vimmer"
[`filepath`]: https://golang.org/pkg/path/filepath/ "filepath - The Go Programming Language"
[`Dir`]: https://docs.ruby-lang.org/ja/latest/class/Dir.html "Class: Dir (Ruby 2.5.0)"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
