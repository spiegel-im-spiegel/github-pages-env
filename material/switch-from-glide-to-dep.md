# Glide から dep に移行せよ

久しぶりに [glide] を使おうと最新版（[0.13.0](https://github.com/Masterminds/glide/releases/tag/v0.13.0 "Release Release 0.13.0 · Masterminds/glide")）を見に行ったら “Consider switching to [dep]” とか書いてあるじゃないですか。

> Glide is used by a great number of projects and will continue to get support for some time. But, the near future is likely in dep.

まじすか。
つまり「依存関係（Vendoring）管理ツールとしては [dep] を推奨するけど移行できない人のために当面はサポートを続けるよ（でも将来は分からん）」という解釈でいいのだろうか。
じゃあ [dep] に移行するしかないぢゃん。

とはいえ，いきなり本番環境に投入するのは怖いので，なにか適当なテストケースはないか，と自分のリポジトリを漁ってたら丁度いいのがありました。

- [spiegel-im-spiegel/pi: Estimate of Pi with Monte Carlo method.](https://github.com/spiegel-im-spiegel/pi)

これは[モンテカルロ法で遊んでた](http://text.baldanders.info/golang/estimate-of-pi/)ときに作ったものなので，最悪ぶっ壊れてもいいのである。
では移行を始めよう。

## [dep] の取得

[dep] のリポジトリは go get コマンドで取得できる。

```text
$ go get -v github.com/golang/dep
github.com/golang/dep (download)
github.com/golang/dep/vendor/github.com/pkg/errors
github.com/golang/dep/vendor/github.com/Masterminds/semver
github.com/golang/dep/vendor/github.com/armon/go-radix
github.com/golang/dep/vendor/github.com/Masterminds/vcs
github.com/golang/dep/vendor/github.com/boltdb/bolt
github.com/golang/dep/internal/fs
github.com/golang/dep/vendor/github.com/golang/protobuf/proto
github.com/golang/dep/internal/gps/paths
github.com/golang/dep/internal/gps/pkgtree
github.com/golang/dep/vendor/github.com/nightlyone/lockfile
github.com/golang/dep/vendor/github.com/sdboyer/constext
github.com/golang/dep/vendor/github.com/jmank88/nuts
github.com/golang/dep/vendor/golang.org/x/net/context
github.com/golang/dep/vendor/github.com/pelletier/go-buffruneio
github.com/golang/dep/vendor/golang.org/x/sync/errgroup
github.com/golang/dep/vendor/github.com/pelletier/go-toml
github.com/golang/dep/internal/gps/internal/pb
github.com/golang/dep/internal/gps
github.com/golang/dep
```

これをこのままインストールしてもいいのだが

```text
$ go install -v github.com/golang/dep/cmd/dep
github.com/golang/dep/vendor/github.com/go-yaml/yaml
github.com/golang/dep/internal/feedback
github.com/golang/dep/internal/importers/base
github.com/golang/dep/internal/importers/godep
github.com/golang/dep/internal/importers/vndr
github.com/golang/dep/internal/importers/gvt
github.com/golang/dep/internal/importers/glide
github.com/golang/dep/internal/importers/govend
github.com/golang/dep/internal/importers
github.com/golang/dep/cmd/dep
```

[リリースページ](https://github.com/golang/dep/releases "Releases · golang/dep")にビルド済みのモジュールが置かれているので，ありがたくこれを使わせてもらう。
最新版（現時点で [v0.3.1](https://github.com/golang/dep/releases/tag/v0.3.1 "Release v0.3.1 · golang/dep")）には Windows 用のモジュール `dep-windows-amd64` もある。
Windows ユーザはなんのファイルかと思うかもしれないが，実はこれ実行ファイルなので， `dep.exe` にリネームしてそのまま使える。
一応，動作確認をしておこう。

```text
$ dep
dep is a tool for managing dependencies for Go projects

Usage: dep <command>

Commands:

  init     Initialize a new project with manifest and lock files
  status   Report the status of the project's dependencies
  ensure   Ensure a dependency is safely vendored in the project
  prune    Prune the vendor tree of unused packages
  version  Show the dep version information

Examples:
  dep init                               set up a new project
  dep ensure                             install the project's dependencies
  dep ensure -update                     update the locked versions of all dependencies
  dep ensure -add github.com/pkg/errors  add a dependency to the project

Use "dep help [command]" for more information about a command.

$ dep version
dep:
 version     : v0.3.1
 build date  : 2017-09-19
 git hash    : 83789e2
 go version  : go1.9
 go compiler : gc
 platform    : windows/amd64
```

## [glide] から [dep] への移行

[spiegel-im-spiegel/pi] をビルド可能な適当な場所に置く。

この時の `glide.yaml` はこんな感じになっている。

```yaml
package: github.com/spiegel-im-spiegel/pi
import:
- package: github.com/spiegel-im-spiegel/gocli
- package: github.com/spf13/cobra
- package: github.com/pkg/errors
- package: github.com/seehuhn/mt19937
- package: github.com/davidminor/gorand
- package: github.com/davidminor/uint128
```

また `glide.lock` はこんな感じ。

```text
hash: d570123d6231810c51dd17e415673df221fb2dec7ef6ab45cd34093002a87cbb
updated: 2016-11-16T17:28:38.2997832+09:00
imports:
- name: github.com/davidminor/gorand
  version: 189780b8053a44a111339a4248394fd844c1da40
  subpackages:
  - lcg
- name: github.com/davidminor/uint128
  version: 5745f1bf80414e0ad2670e85d6aece8c58031def
- name: github.com/inconshreveable/mousetrap
  version: 76626ae9c91c4f2a10f34cad8ce83ea42c93bb75
- name: github.com/pkg/errors
  version: 248dadf4e9068a0b3e79f02ed0a610d935de5302
- name: github.com/seehuhn/mt19937
  version: 98c0ea580d2f3c5a171acf4d4f15321b72209d08
- name: github.com/spf13/cobra
  version: 6b74a60562f5c1c920299b8f02d153e16f4897fc
- name: github.com/spf13/pflag
  version: 5ccb023bc27df288a957c5e994cd44fd19619465
- name: github.com/spiegel-im-spiegel/gocli
  version: 5929f04fb8e4a19ac29fdf658866f9441f339cd9
testImports: []
```

この状態で dep init を実行する。

```text
$ dep init
Importing configuration from glide. These are only initial constraints, and are further refined during the solve process.
Detected glide configuration files...
Converting from glide.yaml and glide.lock...
  Using * as initial constraint for imported dep github.com/spiegel-im-spiegel/gocli
  Trying v0.3.0 (5929f04) as initial lock for imported dep github.com/spiegel-im-spiegel/gocli
  Using * as initial constraint for imported dep github.com/spf13/cobra
  Trying * (6b74a60) as initial lock for imported dep github.com/spf13/cobra
  Using * as initial constraint for imported dep github.com/pkg/errors
  Trying * (248dadf) as initial lock for imported dep github.com/pkg/errors
  Using * as initial constraint for imported dep github.com/seehuhn/mt19937
  Trying master (98c0ea5) as initial lock for imported dep github.com/seehuhn/mt19937
  Using * as initial constraint for imported dep github.com/davidminor/gorand
  Trying * (189780b) as initial lock for imported dep github.com/davidminor/gorand
  Using * as initial constraint for imported dep github.com/davidminor/uint128
  Trying master (5745f1b) as initial lock for imported dep github.com/davidminor/uint128
  Using * as initial constraint for imported dep github.com/inconshreveable/mousetrap
  Trying v1.0 (76626ae) as initial lock for imported dep github.com/inconshreveable/mousetrap
  Using * as initial constraint for imported dep github.com/spf13/pflag
  Trying * (5ccb023) as initial lock for imported dep github.com/spf13/pflag
```

これにより `Gopkg.toml` および `Gopkg.lock` の2つのファイルと `vendor/` フォルダが作成される。
このうち [spiegel-im-spiegel/gocli] の最新版は v0.4.0 だが， `glide.lock` の内容を読み取って，ちゃんと v0.3.0 のものを取ってきているようだ。

`Gopkg.toml` の内容は以下の通り。

```toml
[[constraint]]
  name = "github.com/davidminor/gorand"

[[constraint]]
  name = "github.com/pkg/errors"

[[constraint]]
  name = "github.com/seehuhn/mt19937"

[[constraint]]
  name = "github.com/spf13/cobra"

[[constraint]]
  name = "github.com/spiegel-im-spiegel/gocli"
```

`Gopkg.lock` の内容は以下の通りだ。

```toml
[[projects]]
  name = "github.com/davidminor/gorand"
  packages = ["lcg"]
  revision = "189780b8053a44a111339a4248394fd844c1da40"

[[projects]]
  branch = "master"
  name = "github.com/davidminor/uint128"
  packages = ["."]
  revision = "5745f1bf80414e0ad2670e85d6aece8c58031def"

[[projects]]
  name = "github.com/inconshreveable/mousetrap"
  packages = ["."]
  revision = "76626ae9c91c4f2a10f34cad8ce83ea42c93bb75"
  version = "v1.0"

[[projects]]
  name = "github.com/pkg/errors"
  packages = ["."]
  revision = "248dadf4e9068a0b3e79f02ed0a610d935de5302"

[[projects]]
  branch = "master"
  name = "github.com/seehuhn/mt19937"
  packages = ["."]
  revision = "98c0ea580d2f3c5a171acf4d4f15321b72209d08"

[[projects]]
  name = "github.com/spf13/cobra"
  packages = ["."]
  revision = "6b74a60562f5c1c920299b8f02d153e16f4897fc"

[[projects]]
  name = "github.com/spf13/pflag"
  packages = ["."]
  revision = "5ccb023bc27df288a957c5e994cd44fd19619465"

[[projects]]
  name = "github.com/spiegel-im-spiegel/gocli"
  packages = ["."]
  revision = "5929f04fb8e4a19ac29fdf658866f9441f339cd9"
  version = "v0.3.0"

[solve-meta]
  analyzer-name = "dep"
  analyzer-version = 1
  inputs-digest = "4a7cc1799d386351173ccdf8266d22ebe2971ce7ba417395a0b63ca267ea9267"
  solver-name = "gps-cdcl"
  solver-version = 1
```

`glide.lock` と `Gopkg.lock` の内容がきちんとマッチしているのが分かると思う。
念のため status も見ておこう。

```text
$ dep status
PROJECT                               CONSTRAINT  VERSION        REVISION  LATEST   PKGS USED
github.com/davidminor/gorand          *                          189780b            1
github.com/davidminor/uint128         *           branch master  5745f1b   5745f1b  1
github.com/inconshreveable/mousetrap  *           v1.0           76626ae   76626ae  1
github.com/pkg/errors                 *                          248dadf            1
github.com/seehuhn/mt19937            *           branch master  98c0ea5   98c0ea5  1
github.com/spf13/cobra                *                          6b74a60            1
github.com/spf13/pflag                *                          5ccb023            1
github.com/spiegel-im-spiegel/gocli   *           v0.3.0         5929f04   ce636bb  1
```

ビルドもちゃんと通る。

```text
$ go build -v ./...
github.com/spiegel-im-spiegel/pi/vendor/github.com/spiegel-im-spiegel/gocli
github.com/spiegel-im-spiegel/pi/vendor/github.com/spf13/pflag
github.com/spiegel-im-spiegel/pi/vendor/github.com/inconshreveable/mousetrap
github.com/spiegel-im-spiegel/pi/vendor/github.com/davidminor/uint128
github.com/spiegel-im-spiegel/pi/vendor/github.com/seehuhn/mt19937
github.com/spiegel-im-spiegel/pi/vendor/github.com/pkg/errors
github.com/spiegel-im-spiegel/pi/vendor/github.com/davidminor/gorand/lcg
github.com/spiegel-im-spiegel/pi/gencmplx
github.com/spiegel-im-spiegel/pi/qq
github.com/spiegel-im-spiegel/pi/genpi
github.com/spiegel-im-spiegel/pi/plot
github.com/spiegel-im-spiegel/pi/estmt
github.com/spiegel-im-spiegel/pi/vendor/github.com/spf13/cobra
github.com/spiegel-im-spiegel/pi/cmd
github.com/spiegel-im-spiegel/pi
```

というわけで， `glide.yaml` と `glide.lock` が正しい状態で残っていれば問題なく [dep] に移行できそうだ。

## ブックマーク

- [Big Sky :: golang オフィシャル謹製のパッケージ依存解決ツール「dep」](https://mattn.kaoriya.net/software/lang/go/20170125023240.htm)
- [Goオフィシャルチーム作成の依存関係管理ツール dep を試してみた ｜ Developers.IO](https://dev.classmethod.jp/go/dep/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[glide]: https://github.com/Masterminds/glide "Masterminds/glide"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"
[cointoss]: https://github.com/spiegel-im-spiegel/cointoss "spiegel-im-spiegel/cointoss: コインを10回投げたとき、表は何回出るだろう。"
[spiegel-im-spiegel/pi]: https://github.com/spiegel-im-spiegel/pi "spiegel-im-spiegel/pi: Estimate of Pi with Monte Carlo method."
[spiegel-im-spiegel/gocli]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Command line interface"
