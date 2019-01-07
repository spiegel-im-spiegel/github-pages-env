+++
title = "Travis CI でクロス・コンパイル（GoReleaser 編）"
date =  "2017-11-02T14:01:06+09:00"
update = "2018-02-01T16:52:35+09:00"
description = "クロス・コンパイルと GitHub への deploy をまとめてやってくれる GoReleaser というツールがあるらしい。"
tags = ["golang", "cross-compile", "continuous-integration", "github", "travis-ci", "tools"]

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

以前「[Travis CI でクロス・コンパイル]({{< relref "cross-compiling-in-travis-ci.md" >}})」で [mitchellh/gox] を使ったクロス・コンパイルと [tcnksm/ghr] を使った [GitHub] への deploy 手順を紹介したが，これらをまとめてやってくれる [GoReleaser] というツールがあるらしい。

- [GoReleaser | Deliver Go binaries as fast and easily as possible](https://goreleaser.com/)
- [goreleaser/goreleaser: Deliver Go binaries as fast and easily as possible](https://github.com/goreleaser/goreleaser)

すでにあるプロジェクトで試すのはどうかと思ったので，まずは以下のデモ用のリポジトリを作って試してみることにした。

- [spiegel-im-spiegel/reldemo](https://github.com/spiegel-im-spiegel/reldemo)

ちなみに中身は，このまえ[ついカッとなって](https://qiita.com/spiegel-im-spiegel/items/272c1b8c01eb287059e0)作った [spiegel-im-spiegel/godump](https://github.com/spiegel-im-spiegel/godump) のコードを流用している。
おおっ，役に立ったじゃないか（笑）

## [GoReleaser] の導入

当然ながら [GoReleaser] 自身は [GitHub] 上でバイナリを配布しているので，そちらを使うのが手っ取り早い。

- [Releases · goreleaser/goreleaser](https://github.com/goreleaser/goreleaser/releases/latest)

ヘルプを見るとこんな感じ。

```text
$ goreleaser help
NAME:
   goreleaser - Deliver Go binaries as fast and easily as possible

USAGE:
   goreleaser.exe [global options] command [command options] [arguments...]

VERSION:
   0.35.5, commit 11fee22a2edf211caec98f2bee97576d3160bdb7, built at 2017-10-25T11:49:23Z

COMMANDS:
     init, i  generate .goreleaser.yml
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, --file FILE, -c FILE, -f FILE  Load configuration from FILE (default: ".goreleaser.yml")
   --release-notes FILE                          Load custom release notes from a markdown FILE
   --skip-validate                               Skip all the validations against the release
   --skip-publish                                Skip all publishing pipes of the release
   --snapshot                                    Generate an unversioned snapshot release
   --rm-dist                                     Remove ./dist before building
   --parallelism value, -p value                 Amount of builds launch in parallel (default: 4)
   --debug                                       Enable debug mode
   --help, -h                                    show help
   --version, -v                                 print the version
```

[GoReleaser] では `.goreleaser.yml` ファイルでビルドや deploy を制御しているようだ。
`goreleaser init` コマンドで雛形を生成してくれるみたいなので試してみる。

```text
$ goreleaser init
   • config created; please edit accordingly to your needs file=.goreleaser.yml
```

作成された `.goreleaser.yml` の中身はこんな感じ。

```yaml
project_name: reldemo
release:
  github:
    owner: spiegel-im-spiegel
    name: reldemo
  name_template: '{{.Tag}}'
brew:
  commit_author:
    name: goreleaserbot
    email: goreleaser@carlosbecker.com
  install: bin.install "reldemo"
builds:
- goos:
  - linux
  - darwin
  goarch:
  - amd64
  - "386"
  goarm:
  - "6"
  main: .
  ldflags: -s -w -X main.version={{.Version}} -X main.commit={{.Commit}} -X main.date={{.Date}}
  binary: reldemo
archive:
  format: tar.gz
  name_template: '{{ .Binary }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}{{ if .Arm }}v{{
    .Arm }}{{ end }}'
  files:
  - licence*
  - LICENCE*
  - license*
  - LICENSE*
  - readme*
  - README*
  - changelog*
  - CHANGELOG*
snapshot:
  name_template: SNAPSHOT-{{ .Commit }}
checksum:
  name_template: '{{ .ProjectName }}_{{ .Version }}_checksums.txt'
```

これをベースにアレンジしていくわけだ。

## 設定ファイルの調整とクロス・コンパイル

さて，修正した `.goreleaser.yml` がこれ。

```yaml
project_name: reldemo
release:
  github:
    owner: spiegel-im-spiegel
    name: reldemo

builds:
- goos:
  - linux
  - darwin
  - windows
  goarch:
  - amd64
  - "386"
  - arm
  - arm64
  goarm:
  - "6"
  main: ./cli/reldemo/
  ldflags: -s -w -X github.com/spiegel-im-spiegel/reldemo/cli/reldemo/facade.Version={{ .Version }}
  binary: reldemo

archive:
  format: tar.gz
  format_overrides:
    - goos: windows
      format: zip
  name_template: '{{ .Binary }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}{{ if .Arm }}v{{ .Arm }}{{ end }}'
  replacements:
    amd64: 64bit
    386: 32bit
    arm: ARM
    arm64: ARM64
    darwin: macOS
    linux: Linux
    windows: Windows
  files:
  - LICENSE*
  - README*

snapshot:
  name_template: SNAPSHOT-{{ .Commit }}

checksum:
  name_template: '{{ .ProjectName }}_{{ .Version }}_checksums.txt'
```

主な変更点は以下の通り。

- `brew` 項目はバッサリ捨てた
- コンパイル対象の OS に Windows を加えた。更にアーキテクチャに ARM を加えた
- `ldflags` を現状のものに合わせた
- 圧縮フォーマットで Windows の場合は zip 圧縮にした。また名前の置き換えも行った

これで実際にビルドを行ってみる。

```text
$ goreleaser --snapshot --skip-publish
   • running goreleaser 0.35.5
   • loading config file       file=.goreleaser.yml
   • publishing disabled in snapshot mode
   • SETTING DEFAULTS
   • GETTING AND VALIDATING GIT STATE
   • releasing v0.0.9, commit 6b96405452b3b9af8817157629fce00acb81564e
   • GENERATING CHANGELOG
   • skipped                   reason=not available for snapshots
   • LOADING ENVIRONMENT VARIABLES
   • skipped                   reason=publishing is disabled
   • CHECKING ./DIST
   • BUILDING BINARIES
   • building                  binary=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_ARM64\reldemo
   • building                  binary=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_32bit\reldemo
   • building                  binary=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_64bit\reldemo
   • building                  binary=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_ARMv6\reldemo
   • building                  binary=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_macOS_64bit\reldemo
   • building                  binary=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_macOS_32bit\reldemo
   • building                  binary=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Windows_64bit\reldemo.exe
   • building                  binary=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Windows_32bit\reldemo.exe
   • CREATING ARCHIVES
   • creating                  archive=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_ARM64.tar.gz
   • creating                  archive=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_32bit.tar.gz
   • creating                  archive=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_64bit.tar.gz
   • creating                  archive=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_macOS_32bit.tar.gz
   • creating                  archive=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_ARMv6.tar.gz
   • creating                  archive=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Windows_64bit.zip
   • creating                  archive=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Windows_32bit.zip
   • creating                  archive=dist\reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_macOS_64bit.tar.gz
   • new release artifact      artifact=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_64bit.tar.gz
   • new release artifact      artifact=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Windows_64bit.zip
   • new release artifact      artifact=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Windows_32bit.zip
   • new release artifact      artifact=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_32bit.tar.gz
   • new release artifact      artifact=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_ARM64.tar.gz
   • new release artifact      artifact=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_ARMv6.tar.gz
   • new release artifact      artifact=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_macOS_32bit.tar.gz
   • new release artifact      artifact=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_macOS_64bit.tar.gz
   • CREATING LINUX PACKAGES WITH FPM
   • skipped                   reason=no output formats configured
   • CREATING LINUX PACKAGES WITH SNAPCRAFT
   • skipped                   reason=no summary nor description were provided
   • CALCULATING CHECKSUMS
   • checksumming              file=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_macOS_64bit.tar.gz
   • checksumming              file=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_32bit.tar.gz
   • checksumming              file=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_64bit.tar.gz
   • checksumming              file=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_ARMv6.tar.gz
   • checksumming              file=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Windows_64bit.zip
   • checksumming              file=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Windows_32bit.zip
   • checksumming              file=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_Linux_ARM64.tar.gz
   • checksumming              file=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_macOS_32bit.tar.gz
   • new release artifact      artifact=reldemo_SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e_checksums.txt
   • CREATING DOCKER IMAGES
   • skipped                   reason=docker section is not configured
   • RELEASING TO GITHUB
   • skipped                   reason=--skip-publish is set
   • CREATING HOMEBREW FORMULA
   • skipped                   reason=--skip-publish is set
   • SUCCESS!
```

これでビルドと（`README.md` や `LICENSE` ファイルを同梱した）圧縮ファイルの生成までできた。
fpm とか snapcraft とか Homebrew とか Docker イメージとか設定がないのでスキップしてるけど，今回はスルーの方向で。

まだバージョンタグを打ってないので `--snapshot` で。
また現時点では [GitHub] に deploy して欲しくないので `--skip-publish` にしている。
OS とアーキテクチャの組み合わせで出来ないものはビルドされないのが分かるだろうか。

実際に生成された実行モジュールを起動してみると

```text
$ reldemo -h
Usage:
  reldemo [flags] [binary file]

Flags:
  -h, --help          help for reldemo
  -n, --name string   value name (default "dumpList")
  -v, --vaersion      output version

$ reldemo -v
reldemo SNAPSHOT-6b96405452b3b9af8817157629fce00acb81564e
```

という感じでバージョン番号にスナップショット情報が入っているのがわかると思う。

なお [GoReleaser] でビルドする際は全てのファイルがコミットされている必要がある。
したがって `.goreleaser.yml` ファイルの調整に手こずるとコミット履歴がアレなことになる。

## [Travis CI] との連携と [GitHub] への Deploy

「[Travis CI でクロス・コンパイル]({{< relref "cross-compiling-in-travis-ci.md" >}})」でも書いたが [Travis CI] から [GitHub] へ Deploy するためには [GitHub] のアクセス・トークンを取得して [Travis CI] の環境変数としてセットする必要がある。

[GitHub] のアクセス・トークンは "Settings” の “Developer settings” → "Personal access tokens” のページで取得できる。

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1626/24367702843_e72366313f.jpg" title="Get access token in GitHub" link="https://www.flickr.com/photos/spiegel/24367702843/" >}}

repo の権限のみを付けること。
この access token を [Travis CI] で参照するには， "Settings” の "Environment Variables” でセットすればよい。
Build log にこの access token が表示されないようにすること。

今回の `.travis.yml` の内容はこんな感じ[^dp1]。

[^dp1]: [Travis CI] への [dep] インストール手順については [FAQ.md](https://golang.github.io/dep/docs/FAQ.html "FAQ · dep") を参考にした。

```yaml
language: go

go:
  - 1.9.*

env:
  - DEP_VERSION="0.4.1"

before_install:
  # Download the binary to bin folder in $GOPATH
  - curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o $GOPATH/bin/dep
  # Make the binary executable
  - chmod +x $GOPATH/bin/dep

install:
  - $GOPATH/bin/dep ensure -v

script:
  - go test -v ./...

after_success:
  - test -n "$TRAVIS_TAG" && curl -sL https://git.io/goreleaser | bash
```

最後の行が [GoReleaser] 起動に関する記述である。
[GitHub] リポジトリにタグが打たれている場合は `$TRAVIS_TAG` にタグの値が入る。
したがってタグがない場合は [GoReleaser] は起動しない。

`https://git.io/goreleaser` は短縮 URL で，中身は [goreleaser/get] の `get` ファイルでシェル・スクリプトになっている。
こんな感じ。

```bash
#!/bin/sh
set -e

TAR_FILE="/tmp/goreleaser.tar.gz"
RELEASES_URL="https://github.com/goreleaser/goreleaser/releases"
test -z "$TMPDIR" && TMPDIR="$(mktemp -d)"

last_version() {
  curl -sL -o /dev/null -w %{url_effective} "$RELEASES_URL/latest" | 
    rev | 
    cut -f1 -d'/'| 
    rev
}

download() {
  test -z "$VERSION" && VERSION="$(last_version)"
  test -z "$VERSION" && {
    echo "Unable to get goreleaser version." >&2
    exit 1
  }
  rm -f "$TAR_FILE"
  curl -s -L -o "$TAR_FILE" \
    "$RELEASES_URL/download/$VERSION/goreleaser_$(uname -s)_$(uname -m).tar.gz"
}

download
tar -xf "$TAR_FILE" -C "$TMPDIR"
"${TMPDIR}/goreleaser" "$@"
```

要するに [GoReleaser] の最新バージョンを取ってきて実行しているのである。

これで全ての準備が整ったので，コミットして origin/master にマージし，タグを討つ。
しばらくして [Travis CI] 側の処理が終われば [Releases](https://github.com/spiegel-im-spiegel/reldemo/releases "Releases · spiegel-im-spiegel/reldemo") ページに反映される。

{{< fig-img src="https://farm5.staticflickr.com/4477/26327727169_039d07e64d.jpg" title="Release page in GitHub" link="https://www.flickr.com/photos/spiegel/26327727169/" >}}

Changelog も [GoReleaser] が生成している。
コミット・ログを元に生成しているので，ログがショボいと，上の図のように， Changelog もショボくなる。
まぁ，最悪は手直しすればいいんだけどね[^cl1]。

[^cl1]: 特定の単語を含むログをフィルタリングする設定も `.goreleaser.yml` でできる。

こんなところかな。

`.goreleaser.yml` を自分の気に入るように調整していく作業は悩ましいが，一度出来てしまえば他プロジェクトでも使い回しがし易いと思う。そうなればリリース管理はかなり楽になるはずである。

## ブックマーク

- [goreleaser と Travis CI で Golang のバイナリ配布を自動化する - /storage/tummy.log](http://rnitame.hatenablog.com/entry/automate-golang-binary-distribution)
- [Go言語: ビルド時にバージョン情報を埋め込みたい - Qiita](https://qiita.com/suin/items/d643a0ccb6270e8e3734)
- [goreleaserを使ってGoで書いたツールのバイナリをGithub Releasesで配布する - $shibayu36->blog;](http://blog.shibayu36.org/entry/2017/10/04/193000)
- [Travis CIでdepを使う - くりにっき](http://sue445.hatenablog.com/entry/2017/01/30/214345)
- [goreleaserでHomebrewのFormulaを自動生成する - Qiita](https://qiita.com/knqyf263/items/53dd0d0916afc5472281)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[mitchellh/gox]: https://github.com/mitchellh/gox "mitchellh/gox: A dead simple, no frills Go cross compile tool"
[tcnksm/ghr]: https://github.com/tcnksm/ghr "tcnksm/ghr: Create Github Release and upload artifacts in parallel"
[goreleaser/get]: https://github.com/goreleaser/get "goreleaser/get: Get the latest goreleaser binary"
[GoReleaser]: https://goreleaser.com/ "GoReleaser | Deliver Go binaries as fast and easily as possible"
[Travis CI]: https://travis-ci.org/ "Travis CI - Test and Deploy Your Code with Confidence"
[GitHub]: https://github.com/ "GitHub"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"
