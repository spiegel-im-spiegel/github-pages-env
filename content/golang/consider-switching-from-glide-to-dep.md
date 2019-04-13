+++
title = "Glide から Dep への移行を検討する"
date = "2017-10-10T18:02:56+09:00"
update = "2018-03-06T11:00:35+09:00"
description = "つまり「依存関係（Vendoring）管理ツールとしては dep を推奨するけど移行できない人のために当面はサポートを続けるよ（でも将来は分からん）」という解釈でいいのだろうか。"
tags = ["golang", "engineering", "package", "vendoring", "tools", "glide", "dep", "testing"]

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

（この記事は [Qiita とのマルチポスト](https://qiita.com/spiegel-im-spiegel/items/e931ad1a7565d02d179e)です。
まぁ，向こうは草稿版だけど）

久しぶりに [glide] を使おうと最新版（[0.13.0](https://github.com/Masterminds/glide/releases/tag/v0.13.0 "Release 0.13.0 · Masterminds/glide")）を見に行ったら “**Consider switching to [dep]**” とか書いてあるじゃない。

{{< fig-quote title="Release 0.13.0" link="https://github.com/Masterminds/glide/releases/tag/v0.13.0" lang="en" >}}
<q>Glide is used by a great number of projects and will continue to get support for some time.
But, the near future is likely in dep.
dep can handle importing Glide config files.
Please consider trying dep on your project or converting to dep.</q>
{{< /fig-quote >}}

まじすか。

つまり「依存関係（Vendoring）管理ツールとしては [dep] を推奨するけど移行できない人のために当面はサポートを続けるよ（でも将来は分からん）」という解釈でいいのだろうか。

[dep] は [Go 言語]開発プロジェクトの公式ツールで，2017年の始めくらいに日本でも話題になったような気がするが，私は [glide] で完全に満足していたのでスルーしていた。
こんなことになるなんて。
ならもう [dep] に移行するしかないぢゃん。

とはいえ，いきなり本番環境に投入するのは怖いので，なにか適当なテストケースはないか，と自分のリポジトリを漁ってたら丁度いいのがあったよ。

- [spiegel-im-spiegel/pi: Estimate of Pi with Monte Carlo method.](https://github.com/spiegel-im-spiegel/pi)

これって[モンテカルロ法で遊んでた]({{< relref "./estimate-of-pi.md" >}})ときに作ったものだ。
最悪ぶっ壊れてもいいので，これ使って試してみるか。

## [dep] の取得

まず [dep] の取得から始めないとだが，リポジトリ自体は `go get` コマンドで取得できる。

```text
$ go get -u github.com/golang/dep/cmd/dep
```

これをこのまま使ってもいいのだが，[リリースページ](https://github.com/golang/dep/releases "Releases · golang/dep")にビルド済みのモジュールが置かれているので，ありがたくこれを使わせてもらおう。

最新版（現時点で [v0.4.1](https://github.com/golang/dep/releases/tag/v0.4.1 "Release v0.4.1 · golang/dep")）には Windows 用のモジュール `dep-windows-amd64.exe` もある。
これを `dep.exe` にリネームして使う。

万が一があっては困るのでモジュールの SHA256 ハッシュ値を確認しておく（こういうのこそ OpenPGP を使ってくれないものか）。
Windows ユーザで Windows 8.1 以降であれば PowerShell（4.0 以上）で [`Get-FileHash`] コマンドレットが使える[^ps1]。

[^ps1]: Windows 7 の場合は “[Windows Management Framework 4.0](https://www.microsoft.com/ja-jp/download/details.aspx?id=40855)” をインストールすることで PowerShell 4.0 にアップグレードできる。

```powershell
PS C:\Users\username\Downloads> Get-FileHash dep-windows-amd64.exe -Algorithm SHA256 | Format-List

Algorithm : SHA256
Hash      : F6E6A872C54D5AE7536AC71FD5BCAC9F4E7B8A1DAFA1EF7C23866E2F3069FE4E
Path      : C:\Users\username\Downloads\dep-windows-amd64.exe
```

これを `dep-windows-amd64.exe.sha256` に記載されている値と比較する。
改竄されてなければ同じ値になるはずである。
目視は辛いのでテキストエディタ等の検索機能を使えばいいだろう。

{{% div-box type="md" %}}
**【2017-10-31 追記】** だんだん面倒になってきたので[ハッシュ値を計算するツール](https://github.com/spiegel-im-spiegel/hash "spiegel-im-spiegel/hash: Calculating Hash Value")を作った。
詳しくは「[Hash 値を計算するパッケージを作ってみた]({{< relref "calculating-hash-value.md" >}})」を参照のこと。
{{% /div-box %}}

<!--
Windows ユーザには（`sha256sum` といった）標準ツールがないのが痛いのだが， [7-Zip] があるなら，これを使ってハッシュ値を確認できる。

```text
$ 7z.exe h -scrcSHA256 dep-windows-amd64

7-Zip [64] 16.04 : Copyright (c) 1999-2016 Igor Pavlov : 2016-10-04

Scanning
1 file, 7696896 bytes (7517 KiB)

SHA256                                                                    Size  Name
---------------------------------------------------------------- -------------  ------------
D4BF3EC10B1808CAB883C6AB2901C396CF463E684FDA350199E93E31806C194A       7696896  dep-windows-amd64
---------------------------------------------------------------- -------------  ------------
D4BF3EC10B1808CAB883C6AB2901C396CF463E684FDA350199E93E31806C194A       7696896

Size: 7696896

SHA256 for data:              D4BF3EC10B1808CAB883C6AB2901C396CF463E684FDA350199E93E31806C194A

Everything is Ok
```

これを `dep-windows-amd64.sha256` に記載されている値と比較する。
改竄されてなければ同じ値になるはずである。
目視は辛いのでテキストエディタ等の検索機能を使えばいいだろう。

{{% div-box type="md" %}}
[NYAGOS](http://www.nyaos.org/index.cgi?p=NYAGOS "NYAOS.ORG - NYAGOS") を使っている人なら `.nyagos` ファイルに

```lua
alias {
    sha1sum='%COMSPEC% /c \"%PROGRAMFILES%/7-Zip/7z.exe\" h -scrcSHA1 $*',
    sha256sum='%COMSPEC% /c \"%PROGRAMFILES%/7-Zip/7z.exe\" h -scrcSHA256 $*',
}
```

とか記述しておけば

```text
$ sha256sum dep-windows-amd64
```

で同じ結果が得られる。
改竄の有無を確認するためにファイルのハッシュ値を調べることはよくあるので準備しておくとよい。
{{% /div-box %}}
-->

実行モジュールの動作確認もしておく。

```text
$ dep
Dep is a tool for managing dependencies for Go projects

Usage: "dep [command]"

Commands:

  init     Set up a new Go project, or migrate an existing one
  status   Report the status of the project's dependencies
  ensure   Ensure a dependency is safely vendored in the project
  prune    Pruning is now performed automatically by dep ensure.
  version  Show the dep version information

Examples:
  dep init                               set up a new project
  dep ensure                             install the project's dependencies
  dep ensure -update                     update the locked versions of all dependencies
  dep ensure -add github.com/pkg/errors  add a dependency to the project

Use "dep help [command]" for more information about a command.

$ dep version
dep:
 version     : v0.4.1
 build date  : 2018-01-24
 git hash    : 37d9ea0a
 go version  : go1.9.1
 go compiler : gc
 platform    : windows/amd64
```

## [glide] から [dep] への移行

お試し用の [spiegel-im-spiegel/pi] をビルド可能な適当な場所に置く。

このパッケージの `glide.yaml` はこんな感じになっている。

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

```yaml
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

この状態で `dep init` コマンドを実行する。

```text
$ dep init
Importing configuration from glide. These are only initial constraints, and are further refined during the solve process.                                                                             
Detected glide configuration files...                                                              
Converting from glide.yaml and glide.lock...                                                       
  Trying v0.3.0 (5929f04) as initial lock for imported dep github.com/spiegel-im-spiegel/gocli     
  Trying * (6b74a60) as initial lock for imported dep github.com/spf13/cobra                       
  Trying * (248dadf) as initial lock for imported dep github.com/pkg/errors                        
  Trying master (98c0ea5) as initial lock for imported dep github.com/seehuhn/mt19937              
  Trying * (189780b) as initial lock for imported dep github.com/davidminor/gorand                 
  Trying master (5745f1b) as initial lock for imported dep github.com/davidminor/uint128           
  Trying v1.0 (76626ae) as initial lock for imported dep github.com/inconshreveable/mousetrap      
  Trying * (5ccb023) as initial lock for imported dep github.com/spf13/pflag                       
  Locking in  (248dadf) for direct dep github.com/pkg/errors                                       
  Locking in  (6b74a60) for direct dep github.com/spf13/cobra                                      
  Using master as constraint for direct dep github.com/seehuhn/mt19937                             
  Locking in master (98c0ea5) for direct dep github.com/seehuhn/mt19937                            
  Using ^0.3.0 as constraint for direct dep github.com/spiegel-im-spiegel/gocli                    
  Locking in v0.3.0 (5929f04) for direct dep github.com/spiegel-im-spiegel/gocli                   
  Locking in  (189780b) for direct dep github.com/davidminor/gorand                                
```

実は [spiegel-im-spiegel/gocli] パッケージの最新版は v0.5.0 だが， `glide.lock` の内容を読み取って，ちゃんと v0.3.0 のものを取ってきているようだ。
偉いぞ！

`dep init` コマンドにより [`Gopkg.toml`] および [`Gopkg.lock`] の2つのファイルと `vendor/` フォルダが作成される。
このうち [`Gopkg.toml`] の内容は以下の通り。

```toml
[[constraint]]
  branch = "master"
  name = "github.com/seehuhn/mt19937"

[[constraint]]
  name = "github.com/spiegel-im-spiegel/gocli"
  version = "0.3.0"

[prune]
  go-tests = true
  unused-packages = true
```

そして [`Gopkg.lock`] の内容は以下の通り。

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
  inputs-digest = "3f9a0c0024e81ba251efaa0cb0014694f8315add84c7e8044a346f370e3e088e"
  solver-name = "gps-cdcl"
  solver-version = 1
```

`glide.lock` と [`Gopkg.lock`] の内容がマッチしているのが分かると思う。

{{% div-box type="md" %}}
**【2018-02-02 追記】** dep v0.4 から挙動が変わった？
どうやら `dep init` 時点でリビジョン管理が必要と判断されるパッケージのみ `Gopkg.toml` に記載される感じ。
それ以外でリビジョン管理が必要なものは手動で `Gopkg.toml` に記述する必要があるかも。
{{% /div-box %}}

念のため `dep status` も見ておこう。

```text
$ dep status
PROJECT                               CONSTRAINT     VERSION        REVISION  LATEST   PKGS USED
github.com/davidminor/gorand          *                             189780b            1
github.com/davidminor/uint128         branch master  branch master  5745f1b   5745f1b  1
github.com/inconshreveable/mousetrap  v1.0           v1.0           76626ae   v1.0     1
github.com/pkg/errors                 *                             248dadf            1
github.com/seehuhn/mt19937            branch master  branch master  98c0ea5   98c0ea5  1
github.com/spf13/cobra                *                             6b74a60            1
github.com/spf13/pflag                *                             5ccb023            1
github.com/spiegel-im-spiegel/gocli   ^0.3.0         v0.3.0         5929f04   v0.3.0   1
```

ビルドもちゃんと通る。

```text
$ go build -v .
github.com/spiegel-im-spiegel/pi/vendor/github.com/davidminor/uint128
github.com/spiegel-im-spiegel/pi/vendor/github.com/spf13/pflag
github.com/spiegel-im-spiegel/pi/vendor/github.com/seehuhn/mt19937
github.com/spiegel-im-spiegel/pi/vendor/github.com/inconshreveable/mousetrap
github.com/spiegel-im-spiegel/pi/vendor/github.com/spiegel-im-spiegel/gocli
github.com/spiegel-im-spiegel/pi/vendor/github.com/davidminor/gorand/lcg
github.com/spiegel-im-spiegel/pi/vendor/github.com/pkg/errors
github.com/spiegel-im-spiegel/pi/gencmplx
github.com/spiegel-im-spiegel/pi/qq
github.com/spiegel-im-spiegel/pi/genpi
github.com/spiegel-im-spiegel/pi/plot
github.com/spiegel-im-spiegel/pi/vendor/github.com/spf13/cobra
github.com/spiegel-im-spiegel/pi/estmt
github.com/spiegel-im-spiegel/pi/cmd
github.com/spiegel-im-spiegel/pi
```

というわけで， `glide.yaml` と `glide.lock` が正しい状態で残っていれば問題なく [dep] に移行できそうだ。

## 依存関係の管理

[`Gopkg.toml`] を以下のように修正してみる。

```toml
[[constraint]]
  name = "github.com/davidminor/gorand"
  branch = "master"

[[constraint]]
  name = "github.com/pkg/errors"
  version = "0.8.*"

[[constraint]]
  name = "github.com/seehuhn/mt19937"
  branch = "master"

[[constraint]]
  name = "github.com/spf13/cobra"
  version = "0.0.*"

[[constraint]]
  name = "github.com/spiegel-im-spiegel/gocli"
  version = "0.3.*"

[prune]
  go-tests = true
  unused-packages = true
```

この状態で `dep ensure` コマンドを実行しステータスを見ると，以下のような感じになる。

```text
$ dep status
PROJECT                               CONSTRAINT     VERSION        REVISION  LATEST   PKGS USED
github.com/davidminor/gorand          branch master  branch master  283446f   283446f  1
github.com/davidminor/uint128         branch master  branch master  5745f1b   5745f1b  1
github.com/inconshreveable/mousetrap  v1.0           v1.0           76626ae   v1.0     1
github.com/pkg/errors                 ^0.8.0         v0.8.0         645ef00   v0.8.0   1
github.com/seehuhn/mt19937            branch master  branch master  98c0ea5   98c0ea5  1
github.com/spf13/cobra                ^0.0.1         v0.0.1         7b2c5ac   v0.0.1   1
github.com/spf13/pflag                *                             5ccb023            1
github.com/spiegel-im-spiegel/gocli   ^0.3.0         v0.3.0         5929f04   v0.3.0   1
```

たとえば

```toml
[[constraint]]
  name = "github.com/davidminor/gorand"
  branch = "master"
```

であれば [github.com/davidminor/gorand] パッケージで master ブランチの最新コミットを取ってくる。
また

```toml
[[constraint]]
  name = "github.com/spiegel-im-spiegel/gocli"
  version = "0.3.*"
```

であれば [spiegel-im-spiegel/gocli] パッケージで v0.3.x の最新バージョンを取ってくる[^v1]。

[^v1]: [`Gopkg.toml`] のバージョンの考え方は “[Semantic Versioning]” に従っている。ワイルドカード等を使ったバージョン指定については [Masterminds/semver] パッケージを参照するとよい。

`[prune]` 指定では `vendor/` フォルダから除外するパッケージやファイルを指定する。

```toml
[prune]
  go-tests = true
  unused-packages = true
```

`go-tests` はテスト用のファイル（`*_test.go`）を `unused-packages` は未使用のパッケージを指す。
なお，値は `true` 以外はエラーになるようだ。
たとえば未使用パッケージも含めたいのであれば `unused-packages = false` とするのではなく記述自体を削除する。

```toml
[prune]
  go-tests = true
```

## 依存関係の視覚化

`dep status` コマンドには結果を DOT 言語で吐き出すオプションがあるようだ。
[Graphviz] があれば，この出力結果を画像データに変換できる。

```text
$ dep status -dot | dot -Tpng -o pi-dependency.png
```

結果はこんな感じ。

{{< fig-img src="./pi-dependency.png" link="./pi-dependency.png" title="pi-dependency.png" width="1275" >}}

ブラボー！

## リポジトリへのパスを直接指定する

GitHub みたいな有名 SaaS に置いてあるパッケージなら [`Gopkg.toml`] に

```toml
[[constraint]]
  name = "github.com/spiegel-im-spiegel/gocli"
  version = "0.3.*"
```

とか書けば適切に処理してくれるけど，有名でない SaaS ディレクトリや職場 LAN のリポジトリ上のパッケージではこうはいかないこともある。
こういう場合には，以下に示す通り，直接リポジトリへの（プロトコルを含めた）パスを指定する。

{{< highlight toml "hl_lines=3" >}}
[[constraint]]
  name = "github.com/spiegel-im-spiegel/gocli"
  source = "git@github.com:spiegel-im-spiegel/gocli.git"
  version = "0.3.* "
{{< /highlight >}}

これで `dep ensure` すれば

{{< highlight text "hl_lines=5" >}}
$ dep ensure -v

...

(1/8) Wrote github.com/spiegel-im-spiegel/gocli (from git@github.com:spiegel-im-spiegel/gocli.git)@v0.3.0
(2/8) Wrote github.com/pkg/errors@248dadf4e9068a0b3e79f02ed0a610d935de5302
(3/8) Wrote github.com/davidminor/gorand@189780b8053a44a111339a4248394fd844c1da40
(4/8) Wrote github.com/spf13/pflag@5ccb023bc27df288a957c5e994cd44fd19619465
(5/8) Wrote github.com/inconshreveable/mousetrap@v1.0
(6/8) Wrote github.com/davidminor/uint128@master
(7/8) Wrote github.com/spf13/cobra@6b74a60562f5c1c920299b8f02d153e16f4897fc
(8/8) Wrote github.com/seehuhn/mt19937@master
{{< /highlight >}}

といった感じになる。
ちゃんと指定したリポジトリからパッケージを取得してきているのが分かるだろう。

## Go 1.9 から glide novendor は必要なくなった

Vendoring で一番あつかいに困るのがテストで，たとえば安直に

```text
$ go test -v ./...
```

とかやると `vendor/` フォルダ以下のパッケージまでテスト・シーケンスが走ってしまうのが困りものであった。
このため [glide] にはこれを回避する `glide novendor` コマンドがあって

```text
$ go test -v $(glide novendor)
```

とすることで `vendor/` フォルダへのテストを回避できるようになっていたのだ[^tst1]。

[^tst1]: [glide] を使わない場合は `go test -v $(go list ./... | grep -v /vendor/)` とかする。どのみち Windows のコマンドプロンプトでは無理だけど（笑）

ところがところが！！

[Go 言語] 1.9 からは `./...` の扱いが変更になり

{{% fig-quote type="md" title="Go 1.9 Release Notes" link="https://golang.org/doc/go1.9#vendor-dotdotdot" lang="en" %}}
“By popular request, `./...` no longer matches packages in `vendor` directories in tools accepting package names, such as `go test`.”
{{% /fig-quote %}}

ということで `./...` に `vendor/` フォルダ以下が含まれないことになったのだ。
たとえば [spiegel-im-spiegel/pi] パッケージの場合は

```text
$ go list ./...
github.com/spiegel-im-spiegel/pi
github.com/spiegel-im-spiegel/pi/cmd
github.com/spiegel-im-spiegel/pi/estmt
github.com/spiegel-im-spiegel/pi/gencmplx
github.com/spiegel-im-spiegel/pi/genpi
github.com/spiegel-im-spiegel/pi/plot
github.com/spiegel-im-spiegel/pi/qq
```

となる。
逆に `vendor/` フォルダも含めたいなら

```text
$ go list ./... ./vendor/...
github.com/spiegel-im-spiegel/pi
github.com/spiegel-im-spiegel/pi/cmd
github.com/spiegel-im-spiegel/pi/estmt
github.com/spiegel-im-spiegel/pi/gencmplx
github.com/spiegel-im-spiegel/pi/genpi
github.com/spiegel-im-spiegel/pi/plot
github.com/spiegel-im-spiegel/pi/qq
github.com/spiegel-im-spiegel/pi/vendor/github.com/davidminor/gorand/lcg
github.com/spiegel-im-spiegel/pi/vendor/github.com/davidminor/gorand/pcg
github.com/spiegel-im-spiegel/pi/vendor/github.com/davidminor/uint128
github.com/spiegel-im-spiegel/pi/vendor/github.com/inconshreveable/mousetrap
github.com/spiegel-im-spiegel/pi/vendor/github.com/pkg/errors
github.com/spiegel-im-spiegel/pi/vendor/github.com/seehuhn/mt19937
github.com/spiegel-im-spiegel/pi/vendor/github.com/spf13/cobra
github.com/spiegel-im-spiegel/pi/vendor/github.com/spf13/cobra/cobra
github.com/spiegel-im-spiegel/pi/vendor/github.com/spf13/cobra/cobra/cmd
github.com/spiegel-im-spiegel/pi/vendor/github.com/spf13/cobra/doc
github.com/spiegel-im-spiegel/pi/vendor/github.com/spf13/pflag
github.com/spiegel-im-spiegel/pi/vendor/github.com/spiegel-im-spiegel/gocli
```

とすればよい。
こっちのほうが遥かに扱いやすいよね。

これでまたひとつ [glide] が「要らない子」になる理由が増えてしまったのだった。

## ブックマーク

- [Goオフィシャルチーム作成の依存関係管理ツール dep を試してみた ｜ Developers.IO](https://dev.classmethod.jp/go/dep/)
- [Big Sky :: golang オフィシャル謹製のパッケージ依存解決ツール「dep」](https://mattn.kaoriya.net/software/lang/go/20170125023240.htm)
- [[go]depでブランチ指定 - Qiita](https://qiita.com/noppefoxwolf/items/49bd460034a5c84e1956)
    - [Models and Mechanisms · dep](https://golang.github.io/dep/docs/ensure-mechanics.html)

- [GOPATH 汚染問題]({{< relref "gopath-pollution.md" >}})
- [パッケージの依存状況の視覚化]({{< relref "package-visualization-tool.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[glide]: https://github.com/Masterminds/glide "Masterminds/glide"
[dep]: https://golang.github.io/dep/ "dep · Dependency management for Go"
[`Gopkg.toml`]: https://golang.github.io/dep/docs/Gopkg.toml.html "Gopkg.toml · dep"
[`Gopkg.lock`]: https://golang.github.io/dep/docs/Gopkg.lock.html "Gopkg.lock · dep"
[7-Zip]: http://www.7-zip.org/
[`Get-FileHash`]: http://technet.microsoft.com/en-us/library/dn520872.aspx
[Graphviz]: http://graphviz.org/ "Graphviz - Graph Visualization Software"
[spiegel-im-spiegel/pi]: https://github.com/spiegel-im-spiegel/pi "spiegel-im-spiegel/pi: Estimate of Pi with Monte Carlo method."
[spiegel-im-spiegel/gocli]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Command line interface"
[github.com/davidminor/gorand]: https://github.com/davidminor/gorand "davidminor/gorand: Basic golang implementation of a permuted congruential generator for pseudorandom number generation"
[Semantic Versioning]: http://semver.org/ "Semantic Versioning 2.0.0 | Semantic Versioning"
[Masterminds/semver]: https://github.com/Masterminds/semver "Masterminds/semver: Work with Semantic Versions in Go"
