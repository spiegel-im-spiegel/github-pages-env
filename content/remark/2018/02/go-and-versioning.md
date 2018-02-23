+++
title = "Vgo (Versioned Go) に関する覚え書き"
date =  "2018-02-23T13:43:41+09:00"
description = "description"
image = "/images/attention/kitten.jpg"
tags        = [ "golang", "engineering", "versioning" ]
draft = true

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

[Go 言語]の次のバージョン（v1.11）から vgo (Versioned Go) を実装する計画があるようで， vgo 関連のドキュメントが公開されている。

- [research!rsc: Go & Versioning](https://research.swtch.com/vgo)
    - [research!rsc: Go += Package Versioning (Go & Versioning, Part 1)](https://research.swtch.com/vgo-intro)
    - [research!rsc: A Tour of Versioned Go (vgo) (Go & Versioning, Part 2)](https://research.swtch.com/vgo-tour)
    - [research!rsc: Semantic Import Versioning (Go & Versioning, Part 3)](https://research.swtch.com/vgo-import)
    - [research!rsc: Minimal Version Selection (Go & Versioning, Part 4)](https://research.swtch.com/vgo-mvs)
    - [research!rsc: Reproducible, Verifiable, Verified Builds (Go & Versioning, Part 5)](https://research.swtch.com/vgo-repro)
    - [research!rsc: Defining Go Modules (Go & Versioning, Part 6)](https://research.swtch.com/vgo-module)

vgo は新しいパッケージのバージョン管理機能で，vendoring 機能を使った [dep] や [glide] のような仕組みとは異なるアプローチらしい。
まず v1.11 で試験的に導入し， v1.12 で正式に導入することを目指しているようだ。
最終的には `go get` を置き換えることを考えてるみたい。

{{< fig-quote title="Go += Package Versioning" link="" lang="en" >}}
<q>I would like Go 1.11 to ship with preliminary support for Go modules, as a kind of technology preview, and then I'd like Go 1.12 to ship with official support. In some later release, we'll remove support for the old, unversioned go get. That's an aggressive schedule, though, and if getting the functionality right means waiting for later releases, we will. </q>
{{< /fig-quote >}}








## ブックマーク

- [Go & Versioning(vgo)を読んで大きな変更が入ったなと思った - Qiita](https://qiita.com/lufia/items/67701e2f927c77a75d6e)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[glide]: https://github.com/Masterminds/glide "Masterminds/glide"
[dep]: https://golang.github.io/dep/ "dep · Dependency management for Go"
