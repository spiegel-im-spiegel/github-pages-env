+++
title = "Struct2pflag v0.2.0 をリリースした"
date =  "2026-05-29T18:13:56+09:00"
description = "最近ようやくこれを使う機会があったが，いくつか機能が足りなかったので，それらを追加して v0.2.0 としてリリースした。"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "package", "module", "cli", "struct" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

昨年 [`goark/struct2pflag`] の[最初のリリース]({{< ref "/release/2025/11/struct2pflag-is-released.md" >}} "hymkor/struct2flag から fork した goark/struct2pflag をリリースした")を行ったが，その後，骨折やら何やらでずうっと止まっていた。
最近ようやくこれを使う機会があったが，いくつか機能が足りなかったので，それらを追加して v0.2.0 としてリリースした。

- [Release v0.2.0 · goark/struct2pflag](https://github.com/goark/struct2pflag/releases/tag/v0.2.0)

追加機能のひとつは環境変数をバインドする `BindEnv()` 関数である。
こんな感じに使う。

```go {hl_lines=["2-3", 9]}
type Env struct {
    TopN int           `env:"TOP_N" pflag:"top-n,n,number of top tags"`
    Wait time.Duration `env:"WAIT"  pflag:"wait,w,wait duration"`
}

func main() {
    env := Env{TopN: 10, Wait: time.Second}

    if err := struct2pflag.BindEnv(&env); err != nil {
        _ = fmt.Fprintln(os.Stderr, err)
        return
    }
    struct2pflag.BindDefault(&env)
    pflag.Parse()
}
```

本当は [`hymkor/struct2env`] を使いたかったのだが， `string` 型しかサポートしてないため泣く泣く自作した。
コードを書いたのは AI だけど（笑）

もうひとつは `Bind()` および `BindEnv()` 関数で対応する型を増やしたこと。
今のところ，対応状況は以下のようになっている。

| Type | `Bind` | `BindEnv` |
|------|:------:|:---------:|
| `bool` | ✓ | ✓ |
| `int` | ✓ | ✓ |
| `int8` / `int16` / `int32` / `int64` | | ✓ |
| `uint` | ✓ | ✓ |
| `uint8` / `uint16` / `uint32` / `uint64` | | ✓ |
| `float32` | ✓ | ✓ |
| `float64` | ✓ | ✓ |
| `string` | ✓ | ✓ |
| `time.Duration` | ✓ | ✓ |

本当は `[]string` もサポートしたかったが，キリがないので見送った。
必要にかられれば対応する。
[`spf13/pflag`] がサポートしている型って物凄くあるので，全部カバーするのは面倒なのだ。
いや，どうせコードを書くのは AI なんだけど。

[`hymkor/struct2flag`]: https://github.com/hymkor/struct2flag "hymkor/struct2flag: `struct2flag` automatically registers struct fields as flags for your Go command-line tools."
[`hymkor/struct2env`]: https://github.com/hymkor/struct2env "hymkor/struct2env: `struct2env` binds environment variables to struct fields according to their env tags."
[`goark/struct2pflag`]: https://github.com/goark/struct2pflag "goark/struct2pflag: automatically registers struct fields as flags for your Go command-line tools."
[`spf13/pflag`]: https://github.com/spf13/pflag "spf13/pflag: Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags."

## 参考図書

{{% review-paapi "4814401191" %}} <!-- 初めてのGo言語 第2版 -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
