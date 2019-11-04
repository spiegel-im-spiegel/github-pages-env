+++
title = "Go 言語の環境変数管理"
date = "2019-09-01T15:31:18+09:00"
description = "Go 1.13 からは go env コマンドに -w オプションを付けて環境変数を設定できる。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "environment" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.13 より環境変数の管理の仕方が変わった。

## 環境変数の設定・削除

まずは `go env` コマンドで [Go 言語]関連の環境変数を表示してみる（一部だけね）。
ちなみに私の作業環境は Linux/[Ubuntu] である。

```text
$ go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/username/.cache/go-build"
GOENV="/home/username/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/username/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
...
```

たとえば `GO111MODULE` の値を `on` にしたければ

```text
$ export GO111MODULE=on

$ go env GO111MODULE
on
```

などとする。
[Go] 1.12 まではこれで OK。

[Go] 1.13 からは `go env` コマンドに `-w` オプションを付けて環境変数を設定できる。

```text
$ export -n GO111MODULE

$ go env -w GO111MODULE=auto

$ go env GO111MODULE
auto
```

また `-u` オプションで設定を削除できる。

```text
$ go env -u GO111MODULE

$ go env | grep GO111MODULE
GO111MODULE=""
```

ちなみに shell で設定している環境変数と `go env -w` コマンドで設定する環境変数が被る場合

```text
$ export GO111MODULE=on

$ go env -w GO111MODULE=auto
warning: go env -w GO111MODULE=... does not override conflicting OS environment variable

$ go env | grep GO111MODULE
GO111MODULE="on"
```

となり shell 側の設定のほうが優先されるようだ。
一時的な変更の場合は shell 側の環境変数を使えということやね。

## 環境変数設定の置き場所

`go env -w` コマンドで設定した環境変数の値は `GOENV` で指示されるファイルに格納される。
`GOENV` の既定値は以下の通り（Linux/[Ubuntu] の場合）。

```text
$ go env GOENV
/home/username/.config/go/env
```

ちなみに `$HOME/.config/` ディレクトリは [XDG Base Directory] でユーザごとにアプリケーション設定を格納するディレクトリとして規定されている。
また `XDG_CONFIG_HOME` 環境変数が定義されている場合は，こちらの値が設定ディレクトリとして優先される。

[Go] 1.13 からは [`os`]`.UserConfigDir()` 関数で [XDG Base Directory] に対応した設定ディレクトリを取得できる。
[`os`]`.UserConfigDir()` 関数は他のプラットフォームにも対応していて，プラットフォームごとに適切なパスを返すようだ。

```go
// UserConfigDir returns the default root directory to use for user-specific
// configuration data. Users should create their own application-specific
// subdirectory within this one and use that.
//
// On Unix systems, it returns $XDG_CONFIG_HOME as specified by
// https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html if
// non-empty, else $HOME/.config.
// On Darwin, it returns $HOME/Library/Application Support.
// On Windows, it returns %AppData%.
// On Plan 9, it returns $home/lib.
//
// If the location cannot be determined (for example, $HOME is not defined),
// then it will return an error.
func UserConfigDir() (string, error) {
	var dir string

	switch runtime.GOOS {
	case "windows":
		dir = Getenv("AppData")
		if dir == "" {
			return "", errors.New("%AppData% is not defined")
		}

	case "darwin":
		dir = Getenv("HOME")
		if dir == "" {
			return "", errors.New("$HOME is not defined")
		}
		dir += "/Library/Application Support"

	case "plan9":
		dir = Getenv("home")
		if dir == "" {
			return "", errors.New("$home is not defined")
		}
		dir += "/lib"

	default: // Unix
		dir = Getenv("XDG_CONFIG_HOME")
		if dir == "" {
			dir = Getenv("HOME")
			if dir == "" {
				return "", errors.New("neither $XDG_CONFIG_HOME nor $HOME are defined")
			}
			dir += "/.config"
		}
	}

	return dir, nil
}
```

以前から [`os`]`.UserCacheDir()` 関数で [XDG Base Directory] 対応のキャッシュ・ディレクトリは取得可能だったが，これでまたひとつ対応が進んだわけだ。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[XDG Base Directory]: https://standards.freedesktop.org/basedir-spec/latest/ "XDG Base Directory Specification"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
