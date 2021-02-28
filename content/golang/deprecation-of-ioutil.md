+++
title = "io/ioutil の非推奨化について"
date =  "2021-02-21T12:06:59+09:00"
description = "Refactoring は計画的に。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "package" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日リリースされた [Go] 1.16 における大きな変更のひとつとして [`io/ioutil`][`ioutil`] パッケージの非推奨化（deprecation）が挙げられる。
Deprecation といっても近い将来に（少なくともバージョン 1.x の間は）廃止されるわけではないのだが， [`io/ioutil`][`ioutil`] パッケージは徐々にメンテナンスされなくなる可能性があるため，互換関数（または変数）へ置き換えていくことが推奨されている。

対象となる変数・関数は以下の通り。

| 1.16 以降 非推奨         | 1.16 以降 推奨        |
| ------------------------ | --------------------- |
| [`ioutil`]`.Discard`     | [`io`]`.Discard`      |
| [`ioutil`]`.NopCloser()` | [`io`]`.NopCloser()`  |
| [`ioutil`]`.ReadAll()`   | [`io`]`.ReadAll()`    |
| [`ioutil`]`.ReadDir()`   | [`os`]`.ReadDir()`    |
| [`ioutil`]`.ReadFile()`  | [`os`]`.ReadFile()`   |
| [`ioutil`]`.TempDir()`   | [`os`]`.MkdirTemp()`  |
| [`ioutil`]`.TempFile()`  | [`os`]`.CreateTemp()` |
| [`ioutil`]`.WriteFile()` | [`os`]`.WriteFile()`  |

このうち [`ioutil`]`.Discard`, [`ioutil`]`.NopCloser()`, [`ioutil`]`.ReadAll()`, [`ioutil`]`.ReadFile()`, [`ioutil`]`.WriteFile()` については置き換え後の変数・関数のラッパーとして再実装されているので，特に気にする必要はないだろう。
何かのついでに refactoring していけばよい。

```go
package ioutil

import (
	"io"
	"io/fs"
	"os"
	"sort"
)

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
//
// As of Go 1.16, this function simply calls io.ReadAll.
func ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

// ReadFile reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF. Because ReadFile
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
//
// As of Go 1.16, this function simply calls os.ReadFile.
func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm
// (before umask); otherwise WriteFile truncates it before writing, without changing permissions.
//
// As of Go 1.16, this function simply calls os.WriteFile.
func WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

// NopCloser returns a ReadCloser with a no-op Close method wrapping
// the provided Reader r.
//
// As of Go 1.16, this function simply calls io.NopCloser.
func NopCloser(r io.Reader) io.ReadCloser {
	return io.NopCloser(r)
}

// Discard is an io.Writer on which all Write calls succeed
// without doing anything.
//
// As of Go 1.16, this value is simply io.Discard.
var Discard io.Writer = io.Discard
```

[`ioutil`]`.TempDir()` と [`ioutil`]`.TempFile()` においては，入出力のインタフェースは [`os`]`.MkdirTemp()`, [`os`]`.CreateTemp()` と完全互換と言えるが，内部のロジックが微妙に異なる。
もしかしたら（並行処理下で使う場合など）何らかの検証が必要かもしれない。

まぁ [`os`]`.MkdirTemp()`, [`os`]`.CreateTemp()` のほうが出来がいいと思うけど。
可能なら早めに置き換えたほうがいいだろう。

問題は [`ioutil`]`.ReadDir()` と [`os`]`.ReadDir()` の差異だ。

[Go] 1.16 の [`ioutil`]`.ReadDir()` は以下のように実装されている。

```go
// ReadDir reads the directory named by dirname and returns
// a list of fs.FileInfo for the directory's contents,
// sorted by filename. If an error occurs reading the directory,
// ReadDir returns no directory entries along with the error.
//
// As of Go 1.16, os.ReadDir is a more efficient and correct choice:
// it returns a list of fs.DirEntry instead of fs.FileInfo,
// and it returns partial results in the case of an error
// midway through reading a directory.
func ReadDir(dirname string) ([]fs.FileInfo, error) {
    f, err := os.Open(dirname)
    if err != nil {
        return nil, err
    }
    list, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        return nil, err
    }
    sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
    return list, nil
}
```

ここで [`fs`]`.FileInfo` 型は [Go] 1.16 で追加された [`io/fs`][`fs`] パッケージで定義されているが，中身は 1.15 までの  [`os`]`.FileInfo` 型と全く同じである。

```go
// A FileInfo describes a file and is returned by Stat.
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
}

// A FileMode represents a file's mode and permission bits.
// The bits have the same definition on all systems, so that
// information about files can be moved from one system
// to another portably. Not all bits apply to all systems.
// The only required bit is ModeDir for directories.
type FileMode uint32
```

なお [Go] 1.16 の [`os`]`.FileInfo` 型は [`fs`]`.FileInfo` の [type alias]({{< relref "./go-1_9-and-type-alias.md" >}} "Go 1.9 と Type Alias") として再定義されている。

```go
package os

import (
	"io/fs"
	"syscall"
)

// A FileInfo describes a file and is returned by Stat and Lstat.
type FileInfo = fs.FileInfo
```

一方 [`os`]`.ReadDir()` は以下のように実装されている。

```go
// A DirEntry is an entry read from a directory
// (using the ReadDir function or a File's ReadDir method).
type DirEntry = fs.DirEntry

// ReadDir reads the named directory,
// returning all its directory entries sorted by filename.
// If an error occurs reading the directory,
// ReadDir returns the entries it was able to read before the error,
// along with the error.
func ReadDir(name string) ([]DirEntry, error) {
    f, err := Open(name)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    dirs, err := f.ReadDir(-1)
    sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })
    return dirs, err
}
```

ここで [`fs`]`.DirEntry` 型は以下のように定義されている。

```go
// A DirEntry is an entry read from a directory
// (using the ReadDir function or a ReadDirFile's ReadDir method).
type DirEntry interface {
    // Name returns the name of the file (or subdirectory) described by the entry.
    // This name is only the final element of the path (the base name), not the entire path.
    // For example, Name would return "hello.go" not "/home/gopher/hello.go".
    Name() string

    // IsDir reports whether the entry describes a directory.
    IsDir() bool

    // Type returns the type bits for the entry.
    // The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
    Type() FileMode

    // Info returns the FileInfo for the file or subdirectory described by the entry.
    // The returned FileInfo may be from the time of the original directory read
    // or from the time of the call to Info. If the file has been removed or renamed
    // since the directory read, Info may return an error satisfying errors.Is(err, ErrNotExist).
    // If the entry denotes a symbolic link, Info reports the information about the link itself,
    // not the link's target.
    Info() (FileInfo, error)
}
```

つまり返り値から [`fs`]`.FileInfo` の情報を取り出すには [`fs`]`.DirEntry.Info()` メソッドを使い，さらに返り値の error を評価する必要がある。

このように [`ioutil`]`.ReadDir()` と [`os`]`.ReadDir()` では返り値とその評価方法が異なるため，単純な置き換えではなく，若干のコードの変更が必要となる。

Refactoring は計画的に。

## ブックマーク

- [Go1.16で追加されたio#ReadAll関数から読むストリーミング中のバッファ拡張の仕方 - My External Storage](https://budougumi0617.github.io/2021/02/22/update_capacity/)
- [#golang io/ioutil の非推奨化（deprecation）について](https://zenn.dev/spiegel/scraps/13158f793611df)

[Go]: https://golang.org/ "The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`ioutil`]: https://golang.org/pkg/io/ioutil/ "ioutil - The Go Programming Language"
[`fs`]: https://golang.org/pkg/io/fs/ "fs - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
