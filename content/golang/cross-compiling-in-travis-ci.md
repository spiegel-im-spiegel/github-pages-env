+++
date = "2016-02-13T19:37:11+09:00"
description = "今回は Travis CI から GitHub へ mitchellh/gox で生成した実行バイナリを deploy することを考える。"
draft = false
tags = ["golang", "cross-compile", "continuous-integration"]
title = "Travis CI でクロス・コンパイル"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

## Gox でまとめてクロス・コンパイル

[Go 言語]の特徴のひとつにクロス・コンパイルの容易さがあるが，複数プラットフォームのビルドをまとめて行う [mitchellh/gox] を使うと便利である。

```text
$ go get github.com/mitchellh/gox

$gox -h
Usage: gox [options] [packages]

  Gox cross-compiles Go applications in parallel.

  If no specific operating systems or architectures are specified, Gox
  will build for all pairs supported by your version of Go.

Options:

  -arch=""            Space-separated list of architectures to build for
  -build-toolchain    Build cross-compilation toolchain
  -cgo                Sets CGO_ENABLED=1, requires proper C toolchain (advanced)
  -gcflags=""         Additional '-gcflags' value to pass to go build
  -ldflags=""         Additional '-ldflags' value to pass to go build
  -tags=""            Additional '-tags' value to pass to go build
  -os=""              Space-separated list of operating systems to build for
  -osarch=""          Space-separated list of os/arch pairs to build for
  -osarch-list        List supported os/arch pairs for your Go version
  -output="foo"       Output path template. See below for more info
  -parallel=-1        Amount of parallelism, defaults to number of CPUs
  -rebuild            Force rebuilding of package that were up to date
  -verbose            Verbose mode

Output path template:

  The output path for the compiled binaries is specified with the
  "-output" flag. The value is a string that is a Go text template.
  The default value is "{{.Dir}}_{{.OS}}_{{.Arch}}". The variables and
  their values should be self-explanatory.

Platforms (OS/Arch):

  The operating systems and architectures to cross-compile for may be
  specified with the "-arch" and "-os" flags. These are space separated lists
  of valid GOOS/GOARCH values to build for, respectively. You may prefix an
  OS or Arch with "!" to negate and not build for that platform. If the list
  is made up of only negations, then the negations will come from the default
  list.

  Additionally, the "-osarch" flag may be used to specify complete os/arch
  pairs that should be built or ignored. The syntax for this is what you would
  expect: "darwin/amd64" would be a valid osarch value. Multiple can be space
  separated. An os/arch pair can begin with "!" to not build for that platform.

  The "-osarch" flag has the highest precedent when determing whether to
  build for a platform. If it is included in the "-osarch" list, it will be
  built even if the specific os and arch is negated in "-os" and "-arch",
  respectively.
```

オプションは色々あるが，とりあえずパッケージを指定して起動すれば全てのプラットフォームについてビルドを行う[^bt]。
私が今つくってる [gpgpdump] を例にすると

[^bt]: ちなみに [Go 言語]の 1.4 までは [mitchellh/gox] インストール後に `gox -build-toolchain` でクロス環境を生成する必要があったが， 1.5 からは不要になった。めでたい！

```text
$ go get get github.com/spiegel-im-spiegel/gpgpdump

$ gox -output "goxdist/{{.OS}}_{{.Arch}}/{{.Dir}}" github.com/spiegel-im-spiegel/gpgpdump

-->      netbsd/arm: github.com/spiegel-im-spiegel/gpgpdump
-->   windows/amd64: github.com/spiegel-im-spiegel/gpgpdump
-->   freebsd/amd64: github.com/spiegel-im-spiegel/gpgpdump
-->      darwin/386: github.com/spiegel-im-spiegel/gpgpdump
-->    darwin/amd64: github.com/spiegel-im-spiegel/gpgpdump
-->       linux/386: github.com/spiegel-im-spiegel/gpgpdump
-->     linux/amd64: github.com/spiegel-im-spiegel/gpgpdump
-->       linux/arm: github.com/spiegel-im-spiegel/gpgpdump
-->     freebsd/386: github.com/spiegel-im-spiegel/gpgpdump
-->      netbsd/386: github.com/spiegel-im-spiegel/gpgpdump
-->     freebsd/arm: github.com/spiegel-im-spiegel/gpgpdump
-->    netbsd/amd64: github.com/spiegel-im-spiegel/gpgpdump
-->   openbsd/amd64: github.com/spiegel-im-spiegel/gpgpdump
-->     openbsd/386: github.com/spiegel-im-spiegel/gpgpdump
-->     windows/386: github.com/spiegel-im-spiegel/gpgpdump

$ ls -l goxdist
drwx 0 2016-02-13 17:41 darwin_386/
drwx 0 2016-02-13 17:41 darwin_amd64/
drwx 0 2016-02-13 17:42 freebsd_386/
drwx 0 2016-02-13 17:41 freebsd_amd64/
drwx 0 2016-02-13 17:42 freebsd_arm/
drwx 0 2016-02-13 17:41 linux_386/
drwx 0 2016-02-13 17:41 linux_amd64/
drwx 0 2016-02-13 17:42 linux_arm/
drwx 0 2016-02-13 17:42 netbsd_386/
drwx 0 2016-02-13 17:42 netbsd_amd64/
drwx 0 2016-02-13 17:41 netbsd_arm/
drwx 0 2016-02-13 17:42 openbsd_386/
drwx 0 2016-02-13 17:42 openbsd_amd64/
drwx 0 2016-02-13 17:42 windows_386/
drwx 0 2016-02-13 17:41 windows_amd64/

$ ls -l goxdist/windows_amd64
-rw- 5712896 2016-02-13 17:41 gpgpdump.exe
```

といい感じに出力してくれる。

OS を指定する場合は `-os "linux windows"` のように指定する。
アーキテクチャは `-arch` オプションを， OS とアーキテクチャを組み合わせる場合は `-osarch linux/arm` などとする。

`-output` オプションの `"goxdist/{{.OS}}_{{.Arch}}/{{.Dir}}"` は出力先のテンプレートで `{{ }}` で囲まれている部分に実際の値が埋め込まれる。
たとえば OS が windows でアーキテクチャが amd64 なら `goxdist/windows_amd64/gpgpdump` と展開されるわけだ[^tpl]。

[^tpl]: `{{ }}` でテンプレートをハンドリングするには [`text/template`] パッケージを使う。静的サイト・ジェネレータの [Hugo] でもこのテンプレート・パッケージを使っている。

このようにクロス・コンパイルが非常に簡単なので [Travis CI] などで複数プラットフォームのバイナリを生成するのも難しくない。

## ghr を使って GitHub に Deploy する

今回は [Travis CI] から [GitHub] へ [mitchellh/gox] で生成した実行バイナリを deploy することを考える。

[mitchellh/gox] で生成した実行バイナリをそのままアップしてもいいのだが，ちょっと気持ち悪いので[^s]，まずは zip で固めてしまおう。
以下のような簡単な shell script を書いてみた。

[^s]: 企業などのネット環境では Web から exe ファイルなどの実行バイナリを直接ダウンロードすることを禁止している場合もある。

```bash
#!/bin/bash

DIR=`pwd`
mkdir ./goxdist/dist
for PLATFORM in $(find ./goxdist -mindepth 1 -maxdepth 1 -type d); do
    PLATFORM_NAME=$(basename ${PLATFORM})

    if [ ${PLATFORM_NAME} = "dist" ]; then
        continue
    fi

    cd ${PLATFORM}
    zip ${DIR}/goxdist/dist/${PLATFORM_NAME}.zip ./*
    cd ${DIR}
done
```

先ほど `goxdist` フォルダ以下に生成した各バイナリをひとつづつ zip 圧縮して `goxdist/dist` フォルダに置くだけの簡単なお仕事。
これで `goxdist/dist` フォルダの中身を [GitHub] に deploy すればよい。
[GitHub] への deploy には [tcnksm/ghr] が便利である。

```text
$ go get github.com/tcnksm/ghr

$ ghr -h

Usage: ghr [options] TAG PATH

  ghr is a tool to create Release on Github and upload your artifacts to
  it. ghr parallelizes upload multiple artifacts.

  You can use ghr on GitHub Enterprise. Change URL by GITHUB_API env var.

Options:

  --username, -u        GitHub username. By default, ghr extracts user
                        name from global gitconfig value.

  --repository, -r      GitHub repository name. By default, ghr extracts
                        repository name from current directory's .git/config
                        value.

  --token, -t           GitHub API Token. To use ghr, you will first need
                        to create a GitHub API token with an account which
                        has enough permissions to be able to create releases.
                        You can set this value via GITHUB_TOKEN env var.

  --parallel=-1         Parallelization factor. This option limit amount
                        of parallelism of uploading. By default, ghr uses
                        number of logic CPU of your PC.

  --delete              Delete release if it already created. If you want
                        to recreate release itself from begining, use
                        this. Just want to upload same artifacts to same
                        release again, use --replace option.

  --replace             Replace artifacts if it is already uploaded. Same
                        artifact measn, same release and same artifact
                        name.

  --stat=false          Show number of download of each release and quit.
                        This is special command.

Examples:

  $ ghr v1.0 dist/     Upload all artifacts which are in dist directory
                       with version v1.0.

  $ ghr --stat         Show download number of each relase and quit.

$ ghr --username spiegel-im-spiegel --token $GITHUB_TOKEN v0.1.2 goxdist/dist/
--> Uploading: windows_amd64_v0.1.2.zip
--> Uploading: linux_amd64_v0.1.2.zip
--> Uploading: darwin_386_v0.1.2.zip
--> Uploading: darwin_amd64_v0.1.2.zip
--> Uploading: freebsd_386_v0.1.2.zip
--> Uploading: freebsd_amd64_v0.1.2.zip
--> Uploading: freebsd_arm_v0.1.2.zip
--> Uploading: linux_386_v0.1.2.zip
--> Uploading: netbsd_arm_v0.1.2.zip
--> Uploading: linux_arm_v0.1.2.zip
--> Uploading: netbsd_386_v0.1.2.zip
--> Uploading: netbsd_amd64_v0.1.2.zip
--> Uploading: openbsd_amd64_v0.1.2.zip
--> Uploading: openbsd_386_v0.1.2.zip
--> Uploading: windows_386_v0.1.2.zip
```

てな感じで deploy できる。
ちなみに `$GITHUB_TOKEN` には [GitHub] の access token をセットする。
Access token は [GitHub] の "Settings” の "Personal access tokens” のページで取得できる。

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1626/24367702843_e72366313f.jpg" title="Get access token in GitHub" link="https://www.flickr.com/photos/spiegel/24367702843/" >}}

repo の権限を付けること。
この access token を [Travis CI] で参照するには， "Settings” の "Environment Variables” でセットすればよい。
Build log にこの access token が表示されないようにすること。

最終的な `.travis.yml` の内容はこんな感じ。

```yaml
language: go

go:
  - 1.5.3

branches:
  only: master

before_install:
  - go get -v github.com/mitchellh/gox
  - go get -v github.com/tcnksm/ghr

script:
  - go test -v ./...

after_success:
  - gox -output "goxdist/{{.OS}}_{{.Arch}}_`git tag | tail -1`/{{.Dir}}" -ldflags "-X main.Version=`git tag | tail -1`"
  - sh scripts/package.sh
  - ghr --username spiegel-im-spiegel --token $GITHUB_TOKEN `git tag | tail -1` goxdist/dist/
```

`git tag | tail -1` で最新のタグを取得して，そこに deploy するようにしている。
[GitHub] に何か push するたびにビルドが走るのはウザいので， master ブランチのみテスト & ビルドの対象とした。

まっ，こんなもんかな。

## ブックマーク

- [Golang + Raspberry Pi + LPS331AP で気圧・温度を測定してみた - Qiita](http://qiita.com/yanolab/items/5a6dfb3c07c94f7c760d) : arm アーキテクチャでいけるらしい
- [Go1.5はクロスコンパイルがより簡単 | SOTA](http://deeeet.com/writing/2015/07/22/go1_5-cross-compile/)
- [Raspberry PI ２ 用の consul を作る (201512版 - Qiita](http://qiita.com/rerofumi/items/d6a8ba08270acb61b31c) : Raspberry PI 上でビルドするより Linux のクロス環境を使ったほうが速いらしい
- [CI-as-a-ServiceでGo言語プロジェクトの最新ビルドを継続的に提供する | SOTA](http://deeeet.com/writing/2014/10/16/golang-in-ci-as-a-service/)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[mitchellh/gox]: https://github.com/mitchellh/gox "mitchellh/gox: A dead simple, no frills Go cross compile tool"
[tcnksm/ghr]: https://github.com/tcnksm/ghr "tcnksm/ghr: Create Github Release and upload artifacts in parallel"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: gpgpdump - OpenPGP packet visualizer"
[Travis CI]: https://travis-ci.org/ "Travis CI - Test and Deploy Your Code with Confidence"
[GitHub]: https://github.com/ "GitHub"
[Hugo]: http://gohugo.io/ "Hugo :: A fast and modern static website engine"
[`text/template`]: https://golang.org/pkg/text/template/ "template - The Go Programming Language"
