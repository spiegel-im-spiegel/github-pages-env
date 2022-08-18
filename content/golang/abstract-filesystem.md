+++
title = "fs.FS を使ってディレクトリ・ファイルを参照する"
date =  "2022-08-18T20:22:38+09:00"
description = "忘れないうちにメモっておく。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

仕事で簡単なファイル操作のコードを書いたのだが [`io/fs`][`fs`] パッケージの使い方をド忘れして往生したので，忘れないうちにメモっておく。

## まずは基本形

まず，以下のディレクトリ・ファイルがあるとする。

```text
$ tree testdata
testdata
├── hello1.txt
├── hello2.txt
└── hello3.txt
```

この `testdata/*.txt` ファイルの内容を標準出力に片っ端から出力するという簡単なお仕事。
なお，表示順は考慮しないものとする。

やりかたは色々あると思うが，今回は [`io/fs`][`fs`] パッケージを使ったやり方でやってみる。

```go
package main

import (
    "fmt"
    "io"
    "io/fs"
    "os"
)

func output(fileSystem fs.FS) error {
    paths, err := fs.Glob(fileSystem, "*.txt")
    if err != nil {
        return err
    }
    for _, path := range paths {
        if _, err := func(path string) (int64, error) {
            f, err := fileSystem.Open(path)
            if err != nil {
                return 0, err
            }
            defer f.Close()
            return io.Copy(os.Stdout, f)
        }(path); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    if err := output(os.DirFS("testdata")); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
```

手順としては

1. [`os`]`.DirFS()` 関数で `testdata` ディレクトリの [`fs`]`.FS` 型インスタンスを取得して `output()` 関数に渡す
2. `output()` 関数では [`fs`]`.Glob()` 関数で `"*.txt"` にマッチするファイルの一覧を取得し，順にファイル内容を標準出力にコピーする

という感じ。
各ファイルのオープンからクローズまでの処理を無名関数[^c1] として囲っているのは `defer f.Close()` を使いたかったから。
深い意味はない。

[^c1]: [Go] の無名関数は関数閉包（closure）として機能する，念のため。

さて，この `output()` 関数を変えることなく通常のファイル以外のファイルシステムで応用してみる。

## 埋め込みファイルシステム

次は [`embed`] パッケージでディレクトリ・ファイルを埋め込んだ場合。
`main` 関数のみ書いておく。
パッケージのインポートは適当に補完してね（標準パッケージしか使ってないので無問題）。

```go
//go:embed testdata/*.txt
var assets embed.FS

func main() {
    fileSystem, err := fs.Sub(assets, "testdata")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    if err := output(fileSystem); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
```

上のコードでは `assets` ファイルシステムのルートが `testdata` ディレクトリになるので [`fs`]`.Sub()` 関数でサブディレクトリまで下りてから `output()` 関数に渡している。

## Zip ファイルシステム

今度は `testdata` ディレクトリを zip 圧縮した `testdata.zip` を読み込む場合。
これも `main` 関数のみ書いておく。

```go
func main() {
    rc, err := zip.OpenReader("./testdata.zip")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer rc.Close()

    fileSystem, err := fs.Sub(rc, "testdata")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    if err := output(fileSystem); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
```

[`zip`]`.OpenReader()` 関数は `*`[`zip`]`.ReadCloser` 型のインスタンスを返すのだが，これが [`fs`]`.FS` interface 型と互換があるため，そのまま [`io/fs`][`fs`] パッケージの関数・メソッドに適用することができる。
最初の頃はこれに気付かなくて，しばらくソースコードとにらめっこしてしまった（笑）

## ブックマーク

- [紙芝居用の簡易サーバを書く【Go 1.16 版】]({{< relref "./embeded-filesystem.md" >}})

[Go]: https://go.dev/
[`fs`]: https://pkg.go.dev/io/fs "fs package - io/fs - Go Packages"
[`os`]: https://pkg.go.dev/os "os package - os - Go Packages"
[`zip`]: https://pkg.go.dev/archive/zip "zip package - archive/zip - Go Packages"
[`embed`]: https://pkg.go.dev/embed "embed package - embed - Go Packages"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
