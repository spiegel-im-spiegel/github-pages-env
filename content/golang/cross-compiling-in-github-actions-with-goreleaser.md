+++
title = "GitHub Actions でクロス・コンパイル（GoReleaser 編）"
date =  "2020-09-30T15:14:35+09:00"
description = "公式の GitHub Action があるので，それを使えばよい。簡単！"
image = "/images/attention/go-logo_blue.png"
tags = ["golang", "cross-compile", "continuous-integration", "github"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

以前（3年前だ！）に Travis CI と [GoReleaser] でクロス・コンパイル&デプロイを行う方法を[紹介]({{< relref "./cross-compiling-in-travis-ci-with-goreleaser.md" >}} "Travis CI でクロス・コンパイル（GoReleaser 編）")したが，今回は [GitHub] Actions を使う方法を紹介する。
なお [GoReleaser] 自体の説明については（だいぶ内容が古いが）以下の記事を参考にどうぞ。

- [Travis CI でクロス・コンパイル（GoReleaser 編）]({{< relref "./cross-compiling-in-travis-ci-with-goreleaser.md" >}})

紹介と言っても公式の [GitHub] Action があるので，それを使えばよい。

- [goreleaser/goreleaser-action: GitHub Action for GoReleaser](https://github.com/goreleaser/goreleaser-action)

リポジトリの `.github/workflows/` ディレクトリに YAML ファイル（例えば `build.yml`）を置き，以下のように記述する。

```yaml
name: build

on:
  push:
    tags:
      - v*
jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.15
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v2
        with:
          version: latest
          args: release --rm-dist
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

これでバージョンタグを打った際に [GoReleaser] によるクロス・コンパイルとデプロイが走る。
簡単！

## ブックマーク

- [Go で書いた CLI ツールのリリースは GoReleaser と GitHub Actions で個人的には決まり | tellme.tokyo](https://tellme.tokyo/post/2020/02/04/release-go-cli-tool/)
- [goreleaserとGitHub Actionsを使えばGoのCLIはgit tagをpushするだけでGitHubとHomeBrewに自動リリースできる - My External Storage](https://budougumi0617.github.io/2020/10/07/auto_release_by_goreleaser/)

[Go]: https://go.dev/
[GoReleaser]: https://goreleaser.com/ "GoReleaser | Deliver Go binaries as fast and easily as possible"
[GitHub]: https://github.com/ "GitHub"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
