+++
title = "次期 Go 言語で導入される（かもしれない） io/fs パッケージについて予習する"
date =  "2020-09-06T16:18:45+09:00"
description = "ツリー型のディレクトリ・ファイル構成のファイルシステムを操作するパッケージに対して統一した interface 型を提供して互換性を高めようというわけだ。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "engineering", "package", "file-system" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日行われた “[Go 1.15 Release Party in Japan](https://gocon.connpass.com/event/186317/)” で[紹介](https://gist.github.com/tenntenn/fe8995c347a5e1000832d3c6942f1fbe "Draft designを読む · GitHub")されていた File System Interfaces のドラフト案について予習がてら覚え書きとして記しておく。

- [File System Interfaces for Go — Draft Design](https://go.googlesource.com/proposal/+/master/design/draft-iofs.md)

{{< fig-youtube id="yx7lmuwUNv8" title="io/fs draft design - YouTube" >}}

たとえば `/path/to/filename.txt` のようにツリー型のディレクトリ・ファイル構成のファイルシステムは多い。
メジャーな OS のファイルシステムは大抵そうだし Web のパスや書庫ファイル（`*.tar` や `*.zip` など）もツリー型のディレクトリ・ファイル構成になっている。

たとえば [Go] の標準パッケージ

- [`archive/zip`]
- [`html/template`]
- [`net/http`][`http`]
- [`os`]
- [`text/template`]

などは（ほぼ）同じツリー型だが使い方やメソッド名などが微妙に異なっている。
またサードパーティ製のパッケージでは， [`rakyll/statik`] のように，実行モジュールにディレクトリ・ファイルをまるっと埋め込めるものもあったりする[^efs1]。

[^efs1]: 実行モジュールにディレクトリ・ファイルを埋め込めるパッケージとしては [`jteeuwen/go-bindata`] や [`jessevdk/go-assets`] が有名だが，これらは最早メンテナンスされていないので使わないほうがいい。

こういったパッケージに対して統一した interface 型を提供して互換性を高めようというわけだ。
したら，テストとかもやり易くなるしね（笑）

## fs.FS 型と fs.File 型

[ドラフト案]では `io/fs` パッケージを新たに作ってファイルシステムの汎化を定義するようだ。

まず，ファイルシステムの汎化型 `fs.FS` は以下のように定義するらしい。

```go
type FS interface {
	Open(name string) (File, error)
}
```

また `fs.FS.Open()` メソッドの返り値になっている `fs.File` 型は

```go
type File interface {
	Stat() (os.FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}
```

と定義される。

たとえば，通常のファイルの読み書きについて

```go
type myFS struct{}

func NewFS() fs.FS {
	return &myFS{}
}

func (fsys *myFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}
```

みたいに定義すれば

```go
func main() {
	f, err := NewFS().Open("no-exist.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
    //Output:
    //open no-exist.txt: no such file or directory
}
```

てな感じに書ける。
どやさ！

ちなみにディレクトリ区切り文字はスラッシュ “`/`” で（実際のファイルシステムに関わらず）統一するらしい。
また相対パス指定で “`.`” や “`..`” は使えないようにするようだ。
まぁ，実際にはパス変換関数とか必要になるかもしれないね。

## ファイルシステム・インタフェースの拡張

上述の説明だと「[`http`]`.FileSystem` 型を使えばええんちゃうん？」となる。
実際 [`http`]`.FileSystem` 型は

```go
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

type FileSystem interface {
	Open(name string) (File, error)
}
```

と定義されているため `fs.FS` / `fs.File` 型とほぼ変わらない[^efs2]。

[^efs2]: たとえば [`rakyll/statik`] パッケージではファイルシステムの生成・初期化関数 `fs.New()` の返り値は [`http`]`.FileSystem` 型である。

駄菓子菓子。

`io/fs` パッケージでは拡張機能を定義した型も用意するらしい。

たとえばファイル情報を取得する `Stat()` メソッドを含む

```go
type StatFS interface {
	FS
	Stat(name string) (os.FileInfo, error)
}
```

や，ディレクトリエントリを読む機能を含む

```go
type ReadDirFS interface {
	FS
	ReadDir(name string) ([]os.FileInfo, error)
}
```

といった interface 型も用意されている。

他にもファイルの内容を一括で読み込める

```go
type ReadFileFS interface {
	FS
	ReadFile(name string) ([]byte, error)
}
```

や `Glob()` メソッドが使える

```go
type GlobFS interface {
	FS
	Glob(pattern string) ([]string, error)
}
```

も用意するようだ。
実際にはこれらの interface 型を組み合わせて使うことになると思われる。

## 【2020-12-14 追記】

2021年2月リリース予定の [Go] 1.16 で実行モジュールへの任意のファイルの埋め込み機能が公式にサポートされるらしい。

- [Big Sky :: Go に go:embed が入った。](https://mattn.kaoriya.net/software/lang/go/20201030092005.htm)
- [Go1.16で追加されるembedとio/fsパッケージについてざっと調べた - Qiita](https://qiita.com/convto/items/4b43072b05e6efdf8dd7)
- [go:embed 詳解 - 使用編 - - Qiita](https://qiita.com/cia_rana/items/e5758816393498d2c50f)

うまくすればこの記事の `io/fs` パッケージとも将来的に統合されるかもね。
楽しみ！

## ブックマーク

- [Draft designを読む · GitHub](https://gist.github.com/tenntenn/fe8995c347a5e1000832d3c6942f1fbe)
- [rakyll/statik でシングルバイナリにまとめる]({{< relref "./using-statik-package.md" >}})

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[ドラフト案]: https://go.googlesource.com/proposal/+/master/design/draft-iofs.md "File System Interfaces for Go — Draft Design"
[`archive/zip`]: https://pkg.go.dev/archive/zip "zip package · pkg.go.dev"
[`html/template`]: https://pkg.go.dev/html/template "template package · pkg.go.dev"
[`http`]: https://pkg.go.dev/net/http "http package · pkg.go.dev"
[`os`]: https://pkg.go.dev/os "os package · pkg.go.dev"
[`text/template`]: https://pkg.go.dev/text/template "template package · pkg.go.dev"
[`rakyll/statik`]: https://github.com/rakyll/statik "rakyll/statik: Embed files into a Go executable"
[`jteeuwen/go-bindata`]: https://github.com/jteeuwen/go-bindata "jteeuwen/go-bindata: Hard fork of jteeuwen/go-bindata because it disappeared, Please refer to issues#5 for details."
[`jessevdk/go-assets`]: https://github.com/jessevdk/go-assets "jessevdk/go-assets: Simple embedding of assets in go"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
