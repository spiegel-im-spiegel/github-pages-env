+++
title = "エラーハンドリング・パッケージ errs v1.0.0 他のリリース【Go 1.15 リリース記念】"
date =  "2020-08-13T09:48:44+09:00"
description = "機能追加なし。不具合修正もなし。 go.mod と Travis CI 設定を変更しただけの簡単なお仕事デス。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "package", "module", "error", "pa-api", "openbd-api", "aozora-api", "tools", "gpgpdump", "books-data", "gnkf" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

この前 v0.7.0 を出したばかりだが，[予告した]({{< relref "./errs-package-v0_7_0-is-released.md" >}} "Go 言語用エラーハンドリング・パッケージ errs v0.7.0 をリリースした")とおり， [Go 1.15 リリース]({{< relref "./go-1_15-is-released.md" >}} "")を記念して [spiegel-im-spiegel/errs] パッケージのメジャー・バージョンを上げることにした。

- [Release v1.0.0 · spiegel-im-spiegel/errs · GitHub](https://github.com/spiegel-im-spiegel/errs/releases/tag/v1.0.0)

機能追加なし。
不具合修正もなし。
`go.mod` と [Travis CI] の設定を変更して [Go] 1.15 に対応させただけの簡単なお仕事デス。

今後は後方互換性をきちんと考えて改修していく予定です `m(_ _)m`

ついでに [spiegel-im-spiegel/errs] パッケージを利用している主なパッケージやツールもバージョンを上げた。

- [Release v0.7.1 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.7.1)
- [Release v0.2.5 · spiegel-im-spiegel/openbd-api · GitHub](https://github.com/spiegel-im-spiegel/openbd-api/releases/tag/v0.2.5)
- [Release v0.2.5 · spiegel-im-spiegel/aozora-api · GitHub](https://github.com/spiegel-im-spiegel/aozora-api/releases/tag/v0.2.5)
- [Release v0.7.2 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.7.2)
- [Release v0.5.6 · spiegel-im-spiegel/books-data · GitHub](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.5.6)
- [Release v0.1.2 · spiegel-im-spiegel/gnkf · GitHub](https://github.com/spiegel-im-spiegel/gnkf/releases/tag/v0.1.2)
- [Release v0.1.20 · spiegel-im-spiegel/mklink · GitHub](https://github.com/spiegel-im-spiegel/mklink/releases/tag/v0.1.20)

基本的には [spiegel-im-spiegel/errs] パッケージのバージョンを上げただけ。
まぁ [v0.7.0 で後方互換性をぶっ壊している]({{< relref "./errs-package-v0_7_0-is-released.md" >}} "Go 言語用エラーハンドリング・パッケージ errs v0.7.0 をリリースした")ので，手直しが多かったけど（笑）

## 32ビット・バイナリの提供を止めました

[gpgpdump], [books-data], [gnkf], [mklink] の各ツールはバイナリでの提供も行っているが，今回から32ビット・バイナリの提供を止めることにした（ARM アーキテクチャを除く）。

きっかけは [Go] 1.15 で `darwin/386` のサポートが終了したこと。

{{< fig-quote type="markdown" title="Go 1.15 Release Notes - The Go Programming Language" link="https://golang.org/doc/go1.15" lang="en" >}}
{{% quote %}}As [announced](https://golang.org/doc/go1.14#darwin) in the Go 1.14 release notes, Go 1.15 drops support for 32-bit binaries on macOS, iOS, iPadOS, watchOS, and tvOS (the `darwin/386` and `darwin/arm` ports). Go continues to support the 64-bit `darwin/amd64` and `darwin/arm64` ports{{% /quote %}}.
{{< /fig-quote >}}

クロスコンパイルは [GoReleaser] を使ってるのでそんなに面倒ではないんだけど... もう32ビット流行らんべ。
Windows で古い32ビット機を（しかも拙作を）使ってる人は多分いないだろうし [Ubuntu] も最新 LTS では遂に32ビット・アーキテクチャのサポートを終了したしね。

まぁ Pure [Go] で書いてるのでダウンロード&ビルドは難しくないし，たいした問題ではないだろう，うんうん。

## ブックマーク

<ul>
  <li><a href="{{< ref "/release/gpgpdump.md" >}}">OpenPGP パケットを可視化する gpgpdump</a></li>
  <li><a href="{{< ref "/release/books-data.md" >}}">書籍データ取得ツール books-data</a></li>
  <li><a href="{{< ref "/release/gnkf.md" >}}">GNKF: Network Kanji Filter by Golang</a></li>
  <li><a href="{{< ref "/release/aozora-api-package-for-golang.md" >}}">Go 言語用青空文庫 API クライアント・パッケージ</a></li>
  <li><a href="{{< ref "/release/openbd-api-package-for-golang.md" >}}">Go 言語用 openBD クライアント・パッケージ</a></li>
  <li><a href="{{< ref "/release/pa-api-v5.md" >}}">Go 言語用 PA-API v5 クライアント・パッケージ</a></li>
  <li><a href="{{< ref "/release/errs-package-for-golang.md" >}}">Go 言語用エラーハンドリング・パッケージ</a></li>
</ul>

- [Markdown 形式のリンクを生成するツールを作ってみた]({{< ref "/golang/make-link-with-markdown-format.md" >}})
- [Travis CI でクロス・コンパイル（GoReleaser 編）]({{< ref "/golang/cross-compiling-in-travis-ci-with-goreleaser.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/errs]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"
[Travis CI]: https://travis-ci.org/ "Travis CI - Test and Deploy Your Code with Confidence"
[GoReleaser]: https://goreleaser.com/
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump/ "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[books-data]: https://github.com/spiegel-im-spiegel/books-data/ "spiegel-im-spiegel/books-data: Search for Books Data"
[gnkf]: https://github.com/spiegel-im-spiegel/gnkf/ "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang"
[mklink]: https://github.com/spiegel-im-spiegel/mklink/ "spiegel-im-spiegel/mklink: Make Link with Markdown Format"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
