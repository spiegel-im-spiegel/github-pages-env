+++
date = "2015-09-19T22:41:30+09:00"
update = "2015-09-21T10:43:00+09:00"
description = "今後のことを考えてパッケージ化の作業を行うことにします。"
draft = false
tags = ["golang", "package", "github"]
title = "機能のパッケージ化"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（初出： [はじめての Go 言語 (on Windows) その6 - Qiita](http://qiita.com/spiegel-im-spiegel/items/404871d2bafd22bdbb90)）

[前回]の続き。
なんだけど，今回はパッケージのお話。
ユリウス日の計算なんて簡単なので今まで `main()` 関数の中にゴリゴリ書いてましたが，今後のことを考えて，これを使ってパッケージ化の作業を行うことにします。

## ユリウス日計算のパッケージ化

まずは，[前回]のコードから計算処理部分をきちんと分離します。

```go
package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "time"
)

func ModifiedJulianDayNumber(t time.Time) int64 {
    if t.Sub(time.Unix(0, 0)) >= 0 {
        return mjdnUnix(t)
    } else {
        return mjdnGregorian(t)
    }
}

func mjdnGregorian(t time.Time) int64 {
    y := int64(t.Year())
    m := int64(t.Month())
    if m < 3 {
        y -= 1
        m += 9
    } else {
        m -= 3
    }
    d := int64(t.Day()) - 1
    return (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
}

func mjdnUnix(t time.Time) int64 {
    const (
        onday   = int64(86400) //seconds
        baseDay = int64(40587) //Modified Julian Date at January 1, 1970
    )
    return (t.Unix())/onday + baseDay
}

func main() {
    //引数のチェック
    flag.Parse()
    argsStr := flag.Args()
    if len(argsStr) < 3 {
        fmt.Fprintln(os.Stderr, "年月日を指定してください")
        return
    }
    args := make([]int, 3)
    for i := 0; i < 3; i++ {
        num, err := strconv.Atoi(argsStr[i])
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        } else {
            args[i] = num
        }
    }
    tm := time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.UTC)
    fmt.Printf("%v\n", tm)
    fmt.Printf("MJD = %d日\n", ModifiedJulianDayNumber(tm))
}
```

```shell
C:>go run julian-day-4.go 1969 12 31
1969-12-31 00:00:00 +0000 UTC
MJD = 40586日

C:>go run julian-day-4.go 1970 1 1
1970-01-01 00:00:00 +0000 UTC
MJD = 40587日

C:>go run julian-day-4.go 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

ユリウス日の端数が鬱陶しいので修正ユリウス日の整数部分のみ計算しています。
あと1970年1月1日を境界として計算方法を変えています。
本当はユリウス暦の場合の計算も含めるべきなんでしょうけど，今回は割愛します。

さて，上のコードのうち修正ユリウス日計算関数を別ファイルにしてパッケージ化します。

```go
package modjulian

import "time"

func DayNumber(t time.Time) int64 {
    if t.Sub(time.Unix(0, 0)) >= 0 {
        return dnUnix(t)
    } else {
        return dnGregorian(t)
    }
}

func dnGregorian(t time.Time) int64 {
    y := int64(t.Year())
    m := int64(t.Month())
    if m < 3 {
        y -= 1
        m += 9
    } else {
        m -= 3
    }
    d := int64(t.Day()) - 1
    return (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
}

func dnUnix(t time.Time) int64 {
    const (
        onday   = int64(86400) //seconds
        baseDay = int64(40587) //Modified Julian Date at January 1, 1970
    )
    return (t.Unix())/onday + baseDay
}
```

パッケージ名は `modjulian` としました。
ちなみにパッケージ内の関数等は名前の先頭が大文字のものだけが外部から参照可能です。

### Go 言語における名前の問題

Go 言語およびそのコミュニティは名前にうるさいようです。
たとえば「[実践Go言語](http://golang.jp/effective_go)[^1]」によると

[^1]: オリジナルは “[Effective Go](https://golang.org/doc/effective_go.html)”

> 慣例では，パッケージ名は小文字でひとつの単語です。アンダースコアや大文字が混ざって(mixedCaps)はいけません。パッケージ使用者がその名前をタイプすることを考慮して，簡潔すぎるぐらいにしてください。

なんだそうで，本当はパッケージ名を `julianDate` としたかったんだけど，それでは筋が悪いらしい。

もうひとつ。

もう一つの慣例は，パッケージ名がそのソースディレクトリのベース名であるということです。たとえば`src/pkg/encoding/base64`に置かれているパッケージは，“`encoding/base64`”としてインポートし，名前は`base64`となります。`encoding_base64`や`encodingBase64`とはなりません。

なのでパッケージの指定はハンガリアン記法とかではなく単語をディレクトリで区切って階層化することで整理できそうです。
これって Java とかに慣れてる人には比較的とっつきやすい仕組みかもしれません。

> 長い名前は，慣れたとしても読みやすくなることはありません。複雑もしくは微妙なニュアンスを持つものに名前をつけるときは，すべての情報を名前で表現しようとするより，通常は役立つドキュメントコメントを書いたほうがよいでしょう。

## パッケージの配置と呼び出し

パッケージを呼び出すのには（標準のパッケージと同じく） `import` を使えばいいのですが，記述によってパッケージをどこに配置するかが変わります。

たとえば

```go
import "./modjulian"
```

と相対パスで記述した場合は，呼び出し元のファイルの場所にある `modjulian` フォルダを探します。
以下はパッケージが見つからなくてエラーになってる例です。

```
C:\workspace\jd\src\julian-day-4b>go build julian-day-4b.go
julian-day-4b.go:10:2: open C:\workspace\jd\src\julian-day-4b\modjulian: The system cannot find the file specified.
```

一方

```go
import "modjulian"
```

と記述した場合には，環境変数 `GOROOT` および `GOPATH` で指定されるフォルダ以下を探すようです。
以下もパッケージが見つからなくてエラーになってる例です。

```shell
C:\workspace\jd\src\julian-day-4b>go build julian-day-4b.go
julian-day-4b.go:10:2: cannot find package "modjulian" in any of:
        C:\Go\src\modjulian (from $GOROOT)
        C:\Gopath\src\modjulian (from $GOPATH)
```

[Go 言語]ではパッケージを相対パスで指定するのは（デバッグ時などを除いて）良くないと言われています。
これは `go get` コマンドでパッケージをビルドする際，相対パスを解釈しないようにしているからのようです[^2]。

[^2]: これについてはいろいろな意見があるようですが，妥当な割り切りだと思います。特にコードを CI (Continuous Integration) によって管理している場合は重要なポイントです。

### パッケージを GitHub 上に配置する

「[go get コマンドでパッケージを管理する]({{< relref "go-get-package.md" >}})」でも紹介しましたが，インターネット上の repository にあるパッケージを

```go
import "github.com/username/package"
```

のように指定することで，任意に取り込むことが可能です。
そこで今回のパッケージを [GitHub] 上に公開しました。

- [`github.com/spiegel-im-spiegel/astrocalc`]`/modjulian`

このパッケージを使って書いてみます。

```go
package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "time"

    "github.com/spiegel-im-spiegel/astrocalc/modjulian"
)

func main() {
    //引数のチェック
    flag.Parse()
    argsStr := flag.Args()
    if len(argsStr) < 3 {
        fmt.Fprintln(os.Stderr, "年月日を指定してください")
        return
    }
    args := make([]int, 3)
    for i := 0; i < 3; i++ {
        num, err := strconv.Atoi(argsStr[i])
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        } else {
            args[i] = num
        }
    }
    tm := time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.UTC)
    fmt.Printf("%v\n", tm)
    fmt.Printf("MJD = %d日\n", modjulian.DayNumber(tm))
}
```

このソースファイル `julian-day-4b.go` と `modjulian` パッケージを以下のように配置してビルドします。

```
C:>SET GOPATH=C:\workspace\jd

C:>cd C:\workspace\jd

C:\workspace\jd>tree /f .
C:\WORKSPACE\JD
│
└─src
    └─julian-day-4b
            julian-day-4b.go

C:\workspace\jd>go get -v github.com/spiegel-im-spiegel/astrocalc/modjulian
github.com/spiegel-im-spiegel/astrocalc (download)
github.com/spiegel-im-spiegel/astrocalc/modjulian

C:\workspace\jd>go install ./...

C:\workspace\jd>tree /f .
C:\WORKSPACE\JD
│
├─bin
│      julian-day-4b.exe
│
├─pkg
│  └─windows_amd64
│      └─github.com
│          └─spiegel-im-spiegel
│              └─astrocalc
│                      modjulian.a
│
└─src
    ├─github.com
    │  └─spiegel-im-spiegel
    │      └─astrocalc
    │          │  .editorconfig
    │          │  .gitignore
    │          │  .travis.yml
    │          │  LICENSE
    │          │  README.md
    │          │
    │          └─modjulian
    │                  example_test.go
    │                  LICENSE
    │                  modjulian.go
    │                  modjulian_test.go
    │
    └─julian-day-4b
            julian-day-4b.go

C:\workspace\jd>bin\julian-day-4b.exe 1969 12 31
1969-12-31 00:00:00 +0000 UTC
MJD = 40586日

C:\workspace\jd>bin\julian-day-4b.exe 1970 1 1
1970-01-01 00:00:00 +0000 UTC
MJD = 40587日

C:\workspace\jd>bin\julian-day-4b.exe 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

まだまだ続くよ。
[次回]はテストについて。

## ブックマーク

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< relref "julian-day-number.md" >}} "「ユリウス日」で遊ぶ"
[次回]: {{< relref "testing.md" >}} "Go 言語のテスト・フレームワーク"
[GitHub]: https://github.com/
[`github.com/spiegel-im-spiegel/astrocalc`]: https://github.com/spiegel-im-spiegel/astrocalc
