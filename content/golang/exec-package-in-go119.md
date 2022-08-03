+++
title = "Go 1.19 で os/exec パッケージの挙動が変わった話"
date =  "2022-08-03T21:28:36+09:00"
description = "Windows 環境でコマンドを起動する際の脆弱性の回避"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "security", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.19 の[リリースノート](https://go.dev/doc/go1.19 "Go 1.19 Release Notes - The Go Programming Language")を眺めてみると

{{< fig-quote type="markdown" title="Go 1.19 Release Notes - The Go Programming Language" link="https://go.dev/doc/go1.19" lang="en" >}}
[`Command`](https://go.dev/pkg/os/exec/#Command) and [`LookPath`](https://go.dev/pkg/os/exec/#LookPath) no longer allow results from a PATH search to be found relative to the current directory. This removes a [common source of security problems](https://go.dev/blog/path-security) but may also break existing programs that depend on using, say, `exec.Command("prog")` to run a binary named `prog` (or, on Windows, `prog.exe`) in the current directory. See the [`os/exec`](https://go.dev/pkg/os/exec/) package documentation for information about how best to update such programs.
{{< /fig-quote >}}

とある。
さっそく試してみよう。

まず Windows 環境で [`gpgpdump.exe`](https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer") コマンドを PATH で指定されたフォルダ以外，具体的には以下のソースファイルと同じフォルダに置く。

```go
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := "gpgpdump.exe"
    out, err := exec.Command(cmd, "version").CombinedOutput()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("output by %v:\n%v\n", cmd, string(out))
}
```

これを [Go] 1.19 コンパイル環境下で実行すると

```text
> go run sample.go
exec: "gpgpdump.exe": cannot run executable found relative to current directory
```

「カレントディレクトリに指定の実行ファイルあるけど起動してやらん（←超意訳）」とエラーになった。

Windows ではパス指定なしでコマンドを起動する際に，カレントフォルダに同名の実行ファイルが存在すると優先的にそれを起動してしまう。
[Go] 標準の [`os/exec`][`exec`] パッケージもこの挙動に合わせていたのだが，2020年の [CVE-2020-27955] で問題になった。
この挙動を悪用して悪意のコマンドを実行される可能性があるというわけだ。

この脆弱性を回避するために，様々な試行錯誤が行われたが [Go] 1.19 の改修が決定打になるだろう。
カレントフォルダにある同名の実行ファイルを無視するのではなく，エラーとして「起動させない」というのがポイント。

なお，今まで通りパスなしのコマンド指定時にカレントフォルダの実行ファイルを起動したいなら [`exec`]`.ErrDot` エラーを明示的に潰すことで実現できる。
こんな感じ。

```go {hl_lines=["11-17"]}
package main

import (
    "errors"
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("gpgpdump.exe", "version")
    if cmd.Err != nil {
        fmt.Println(cmd.Err)
        if !errors.Is(cmd.Err, exec.ErrDot) {
            return
        }
        cmd.Err = nil
    }
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("output by %v:\n%v\n", cmd, string(out))
}
```

これを実行すると

```text
> go run sample2.go
exec: "gpgpdump.exe": cannot run executable found relative to current directory
output by .\gpgpdump.exe version:
gpgpdump v0.14.0
repository: https://github.com/goark/gpgpdump
```

となる。
エラーを無視してカレントディレクトリ `.` を付加した状態で実行されているのがお分かりだろうか。

ちなみに，同じコードを Windows 以外の環境で実行すると（`.exe` の拡張子は外してね）

```text
$ go run sample2b.go 
exec: "gpgpdump": executable file not found in $PATH
```

と PATH 上に実行ファイルが見つからない旨の普通のエラーが表示される。
これでアプリケーション側は OS ごとに処理を分ける必要がなくなったわけだ。
めでたい！

ところで Windows には `NoDefaultCurrentDirectoryInExePath` なる環境変数があるそうで，これが有効になっているとパスなしのコマンド指定時にカレントフォルダの同名実行ファイルを無視するらしい。

で [`os/exec`][`exec`] パッケージは律儀にこの環境変数にも対応している。

{{< fig-quote type="markdown" title="Go 1.19 Release Notes - The Go Programming Language" link="https://go.dev/doc/go1.19" lang="en" >}}
On Windows, `Command` and `LookPath` now respect the `NoDefaultCurrentDirectoryInExePath` environment variable, making it possible to disable the default implicit search of “`.`” in PATH lookups on Windows systems.
{{< /fig-quote >}}


標準パッケージのソースコード `os/exec/lp_windows.go` を眺めると

```go {hl_lines=["49-56"]}
// LookPath searches for an executable named file in the
// directories named by the PATH environment variable.
// LookPath also uses PATHEXT environment variable to match
// a suitable candidate.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// Otherwise, on success, the result is an absolute path.
//
// In older versions of Go, LookPath could return a path relative to the current directory.
// As of Go 1.19, LookPath will instead return that path along with an error satisfying
// errors.Is(err, ErrDot). See the package documentation for more details.
func LookPath(file string) (string, error) {
    var exts []string
    x := os.Getenv(`PATHEXT`)
    if x != "" {
        for _, e := range strings.Split(strings.ToLower(x), `;`) {
            if e == "" {
                continue
            }
            if e[0] != '.' {
                e = "." + e
            }
            exts = append(exts, e)
        }
    } else {
        exts = []string{".com", ".exe", ".bat", ".cmd"}
    }

    if strings.ContainsAny(file, `:\/`) {
        f, err := findExecutable(file, exts)
        if err == nil {
            return f, nil
        }
        return "", &Error{file, err}
    }

    // On Windows, creating the NoDefaultCurrentDirectoryInExePath
    // environment variable (with any value or no value!) signals that
    // path lookups should skip the current directory.
    // In theory we are supposed to call NeedCurrentDirectoryForExePathW
    // "as the registry location of this environment variable can change"
    // but that seems exceedingly unlikely: it would break all users who
    // have configured their environment this way!
    // https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-needcurrentdirectoryforexepathw
    // See also go.dev/issue/43947.
    var (
        dotf   string
        dotErr error
    )
    if _, found := syscall.Getenv("NoDefaultCurrentDirectoryInExePath"); !found {
        if f, err := findExecutable(filepath.Join(".", file), exts); err == nil {
            if godebug.Get("execerrdot") == "0" {
                return f, nil
            }
            dotf, dotErr = f, &Error{file, ErrDot}
        }
    }

    path := os.Getenv("path")
    for _, dir := range filepath.SplitList(path) {
        if f, err := findExecutable(filepath.Join(dir, file), exts); err == nil {
            if dotErr != nil {
                // https://go.dev/issue/53536: if we resolved a relative path implicitly,
                // and it is the same executable that would be resolved from the explicit %PATH%,
                // prefer the explicit name for the executable (and, likely, no error) instead
                // of the equivalent implicit name with ErrDot.
                //
                // Otherwise, return the ErrDot for the implicit path as soon as we find
                // out that the explicit one doesn't match.
                dotfi, dotfiErr := os.Lstat(dotf)
                fi, fiErr := os.Lstat(f)
                if dotfiErr != nil || fiErr != nil || !os.SameFile(dotfi, fi) {
                    return dotf, dotErr
                }
            }

            if !filepath.IsAbs(f) && godebug.Get("execerrdot") != "0" {
                return f, &Error{file, ErrDot}
            }
            return f, nil
        }
    }

    if dotErr != nil {
        return dotf, dotErr
    }
    return "", &Error{file, ErrNotFound}
}
```

と `NoDefaultCurrentDirectoryInExePath` 環境変数がない場合だけカレントディレクトリ `.` を付加してチェックしているのが分かる。
ご苦労さんなことである。

やっぱ Windows は面倒くさいな（笑）

## ブックマーク

- [Go でサブプロセスを起動する際は LookPath に気をつけろ！](https://zenn.dev/spiegel/articles/20201107-lookpath-by-golang)
- [github.com/cli/safeexec パッケージを使った外部コマンドの安全な起動]({{< relref "./safeexec-packge.md" >}})

[Go]: https://go.dev/
[`exec`]: https://pkg.go.dev/os/exec "exec package - os/exec - Go Packages"
[CVE-2020-27955]: https://nvd.nist.gov/vuln/detail/CVE-2020-27955

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
