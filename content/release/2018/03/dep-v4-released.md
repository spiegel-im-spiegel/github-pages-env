+++
title = "Go 言語の依存性管理ツール Dep v0.4 がリリースされていた"
date = "2018-03-04T14:18:16+09:00"
description = "2018年1月に dep の v0.4 がリリースされた。 v0.4 では Gopkg.toml に prune 項目が追加された。"
image = "/images/attention/tools.png"
tags  = [ "golang", "tools", "dep", "package", "vendoring" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

随分経っているが，忘れないように覚え書きとして残しておく。

2018年1月に [dep] の v0.4 がリリースされた（2018-03-04 時点で v0.4.1）。
随分前の話だが，忘れないように覚え書きとして残しておく。

- [Release v0.4.0 · golang/dep](https://github.com/golang/dep/releases/tag/v0.4.0)
- [Release v0.4.1 · golang/dep](https://github.com/golang/dep/releases/tag/v0.4.1)

v0.4 では [`Gopkg.toml`] に `prune` 項目が追加された。
`prune` は `vendor/` フォルダにパッケージを展開する際に除外するファイルを指定できる。
以下のように記述する。

```toml
[prune]
  unused-packages = true
  go-tests = true
```

`unused-packages` は未使用パッケージを除外する。
`go-tests` はテストコードを除外する。
他にも `non-go` を指定すると [Go 言語]コード以外のファイルを除外してくれる。

また，パッケージ単位でも

```toml
[prune]
  non-go = true

  [[prune.project]]
    name = "github.com/project/name"
    go-tests = true
    non-go = false
```

という感じに指定できるらしい。

ちなみに [dep] を [Travis CI] で使うには `.travis.yml` で以下のように記述する。

```yaml
env:
  - DEP_VERSION="0.4.1"

before_install:
  # Download the binary to bin folder in $GOPATH
  - curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o $GOPATH/bin/dep
  # Make the binary executable
  - chmod +x $GOPATH/bin/dep
```

[dep] のバージョンが上がった場合には `DEP_VERSION` の値だけ変えればいける（はず）。

## ブックマーク

- [Glide から Dep への移行を検討する]({{< ref "/golang/consider-switching-from-glide-to-dep.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[dep]: https://golang.github.io/dep/ "dep · Dependency management for Go"
[`Gopkg.toml`]: https://golang.github.io/dep/docs/Gopkg.toml.html "Gopkg.toml · dep"
[Travis CI]: https://travis-ci.org/ "Travis CI - Test and Deploy Your Code with Confidence"
