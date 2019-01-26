+++
title = "Gocli Package v0.8.0 Released"
date = "2018-03-08T22:28:19+09:00"
description = "Ruby の Dir.glob にあるようなディレクトリの再帰検索ができる file.Glob() 関数を作ってみた。"
image = "/images/attention/tools.png"
tags  = [ "golang", "package", "cli" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

この前 [v0.7.0 をリリースした]({{< ref "/release/2018/03/gocli-package-v0_7_0-released.md" >}})ばかりだが，舌の根も乾かぬうちに v0.8.0 をリリースしてしまった。

- [Release v0.8.0 Released · spiegel-im-spiegel/gocli](https://github.com/spiegel-im-spiegel/gocli/releases/tag/v0.8.0)

今回はファイル・ディレクトリ操作用の [`gocli`]`/file` サブパッケージを追加した。

実は [mattn/jvgrep] を使ってみてディレクトリの再帰検索はなかなか便利なことに気づき，自分でも実装してみようと思ったのだ。

[Go 言語]標準の [`filepath`]`.Glob()` 関数はなかなか性能がよくて，パス検索に以下のワイルドカードが使える。

{{% fig-quote title="filepath - The Go Programming Language" link="https://golang.org/pkg/path/filepath/" lang="en" %}}
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
{{% /fig-quote %}}

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
