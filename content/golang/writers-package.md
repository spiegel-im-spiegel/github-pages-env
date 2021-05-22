+++
title = "Writers パッケージを作ってみた"
date =  "2020-03-28T20:40:05+09:00"
description = "言語的に面白いトピックはないし，手遊びということで。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "golang", "package", "logger", "regular-expression" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Twitter で

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">できればLogレベルでファイル分割したいところですが、Goの情報はなかなか見つからなかった...。<br><br>Trace/Debug/Info/Warning<br>Error/Fatal<br>とかいう好きな単位でログファイルを切り替えられれば、別ツイのtail -f時のノイズ削減になりそう。<a href="https://t.co/MTyoOqyvgQ">https://t.co/MTyoOqyvgQ</a><br>(分割について記載なすび)</p>&mdash; SIG (@sig_246) <a href="https://twitter.com/sig_246/status/1243135650322243584?ref_src=twsrc%5Etfw">March 26, 2020</a></blockquote>
{{< /fig-gen >}}

というのを見かけたので，試しに作ってみた。

[^tee1]: ちなみに [Go] の標準パッケージにも [`io`]`.TeeReader()` 関数ってのがあって `tee` コマンドと同等のことができる。

- [spiegel-im-spiegel/writers: Filtering Writer](https://github.com/spiegel-im-spiegel/writers)

いや `tee` および `grep` コマンドを組み合わせれば出力の分割はできるんだけどね[^tee1]。
まぁ，言語的に面白いトピックはないし，手遊びということで。

たとえば，拙作の [`logf`] パッケージを使ってこんなログ出力を考えてみる。

```go
package main

import (
    "os"

    "github.com/spiegel-im-spiegel/logf"
)

func main() {
    logf.SetOutput(os.Stdout)
    for i := 0; i < 6; i++ {
        logf.SetMinLevel(logf.TRACE + logf.Level(i))
        logf.Tracef("Traceing: No. %d\n", i+1)
        logf.Debugf("Debugging: No. %d\n", i+1)
        logf.Printf("Information: No. %d\n", i+1)
        logf.Warnf("Warning: No. %d\n", i+1)
        logf.Errorf("Erroring: No. %d\n", i+1)
        logf.Fatalf("Fatal Erroring: No. %d\n", i+1)
    }
}
```

これを実行すると，こんな感じになる。

```text
$ go run sample.go
2020/03/28 14:44:44 [TRACE] Traceing: No. 1
2020/03/28 14:44:44 [DEBUG] Debugging: No. 1
2020/03/28 14:44:44 [INFO] Information: No. 1
2020/03/28 14:44:44 [WARN] Warning: No. 1
2020/03/28 14:44:44 [ERROR] Erroring: No. 1
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 1
2020/03/28 14:44:44 [DEBUG] Debugging: No. 2
2020/03/28 14:44:44 [INFO] Information: No. 2
2020/03/28 14:44:44 [WARN] Warning: No. 2
2020/03/28 14:44:44 [ERROR] Erroring: No. 2
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 2
2020/03/28 14:44:44 [INFO] Information: No. 3
2020/03/28 14:44:44 [WARN] Warning: No. 3
2020/03/28 14:44:44 [ERROR] Erroring: No. 3
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 3
2020/03/28 14:44:44 [WARN] Warning: No. 4
2020/03/28 14:44:44 [ERROR] Erroring: No. 4
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 4
2020/03/28 14:44:44 [ERROR] Erroring: No. 5
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 5
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 6
```

これを出発点とする。

出力を多重化するには [`io`]`.MultiWriter()` 関数を使うとよい。
こんな感じ。

```go {hl_lines=[5, "19-22"]}
package main

import (
    "fmt"
    "io"
    "os"

    "github.com/spiegel-im-spiegel/logf"
)

func main() {
    file, err := os.Create("log.txt")
    if err != nil {
        fmt.Printf("%#v\n", err)
        return
    }
    defer file.Close()

    ws := io.MultiWriter(
        file,
        os.Stdout,
    )
    logf.SetOutput(ws)
    for i := 0; i < 6; i++ {
        logf.SetMinLevel(logf.TRACE + logf.Level(i))
        logf.Tracef("Traceing: No. %d\n", i+1)
        logf.Debugf("Debugging: No. %d\n", i+1)
        logf.Printf("Information: No. %d\n", i+1)
        logf.Warnf("Warning: No. %d\n", i+1)
        logf.Errorf("Erroring: No. %d\n", i+1)
        logf.Fatalf("Fatal Erroring: No. %d\n", i+1)
    }
}
```

これで標準出力と `log.txt` ファイルに全く同じ内容が出力される。

次に，標準出力には `[ERROR]` と `[FATAL]` のログのみ出力したい。
そこでこんな型を考える。

```go
package writers

//FilterWriter type is Writer with filter
type FilterWriter struct {
    word   []byte
    writer io.Writer
}
```

この型に対して以下の `Write()` メソッド

```go
//Write function writes bytes data.
func (w *FilterWriter) Write(b []byte) (int, error) {
    if w.match(b) {
        return w.writer.Write(b)
    }
    return len(b), nil
}

func (w *FilterWriter) match(b []byte) bool {
    if len(b) == 0 {
        return false
    }
    if w.word == nil {
        return true
    }
    return bytes.Contains(b, w.word)
}
```

を組み込めば，設定したキーワードを含んでいる場合のみ書き込みを行うようになる。

[`writers`]`.FilterWriter` を使って先程のコードを書き換えてみよう。
こんな感じ。

```go {hl_lines=[9, "20-24"]}
package main

import (
    "fmt"
    "io"
    "os"

    "github.com/spiegel-im-spiegel/logf"
    "github.com/spiegel-im-spiegel/writers"
)

func main() {
    file, err := os.Create("log.txt")
    if err != nil {
        fmt.Printf("%#v\n", err)
        return
    }
    defer file.Close()

    ws := io.MultiWriter(
        file,
        writers.Filter(os.Stdout, []byte("[ERROR]")),
		writers.Filter(os.Stdout, []byte("[FATAL]")),
    )
    logf.SetOutput(ws)
    for i := 0; i < 6; i++ {
        logf.SetMinLevel(logf.TRACE + logf.Level(i))
        logf.Tracef("Traceing: No. %d\n", i+1)
        logf.Debugf("Debugging: No. %d\n", i+1)
        logf.Printf("Information: No. %d\n", i+1)
        logf.Warnf("Warning: No. %d\n", i+1)
        logf.Errorf("Erroring: No. %d\n", i+1)
        logf.Fatalf("Fatal Erroring: No. %d\n", i+1)
    }
}
```

これで標準出力が

```
$ go run sample.go
2020/03/28 14:44:44 [ERROR] Erroring: No. 1
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 1
2020/03/28 14:44:44 [ERROR] Erroring: No. 2
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 2
2020/03/28 14:44:44 [ERROR] Erroring: No. 3
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 3
2020/03/28 14:44:44 [ERROR] Erroring: No. 4
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 4
2020/03/28 14:44:44 [ERROR] Erroring: No. 5
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 5
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 6
```

となった。

単純な比較のみだと複雑なパターンを構成し辛いので，正規表現バージョンも作ってみた。

```go
//RegexpWriter type is Writer with regular expression filter
type RegexpWriter struct {
    re     *regexp.Regexp
    writer io.Writer
}

//WriteString function writes string.
func (w *RegexpWriter) Write(b []byte) (int, error) {
    if w.match(b) {
        return w.writer.Write(b)
    }
    return len(b), nil
}

func (w *RegexpWriter) match(b []byte) bool {
    if len(b) == 0 {
        return false
    }
    if w.re == nil {
        return true
    }
    return w.re.Match(b)
}
```

これを使えば，先程のコードはこんな感じにできる。

```go {hl_lines=[7, "21-24"]}
package main

import (
    "fmt"
    "io"
    "os"
    "regexp"

    "github.com/spiegel-im-spiegel/logf"
    "github.com/spiegel-im-spiegel/writers"
)

func main() {
    file, err := os.Create("log.txt")
    if err != nil {
        fmt.Printf("%#v\n", err)
        return
    }
    defer file.Close()

    ws := io.MultiWriter(
        file,
        writers.FilterRegexp(os.Stdout, regexp.MustCompile(`\[(ERROR|FATAL)\]`)),
    )
    logf.SetOutput(ws)
    for i := 0; i < 6; i++ {
        logf.SetMinLevel(logf.TRACE + logf.Level(i))
        logf.Tracef("Traceing: No. %d\n", i+1)
        logf.Debugf("Debugging: No. %d\n", i+1)
        logf.Printf("Information: No. %d\n", i+1)
        logf.Warnf("Warning: No. %d\n", i+1)
        logf.Errorf("Erroring: No. %d\n", i+1)
        logf.Fatalf("Fatal Erroring: No. %d\n", i+1)
    }
}
```

これで同じ結果が得られる。

今回はベースの `Writer` に出力多重化やらフィルタやらの機能を被せているだけなので，色々と応用が効くだろう。
効くといいな（笑）

## ブックマーク

- [Log パッケージで遊ぶ]({{< relref "./logger.md" >}})
- [正規表現に関する戯れ言]({{< relref "./regular-expression.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`io`]: https://pkg.go.dev/io "io package · go.dev"
[`logf`]: https://github.com/spiegel-im-spiegel/logf "spiegel-im-spiegel/logf: Simple leveled logging package by Golang"
[`writers`]: https://github.com/spiegel-im-spiegel/writers "spiegel-im-spiegel/writers: Writer Collection"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
