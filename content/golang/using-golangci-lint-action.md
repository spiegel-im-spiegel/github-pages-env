+++
title = "golangci-lint を GitHub Actions で使う"
date =  "2020-09-29T11:20:19+09:00"
description = "これで pull request 時に golangci-lint が走る。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "lint", "programming", "engineering", "github", "continuous-integration" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[golangci-lint] は `go vet` をはじめ複数の lint を集約して結果を表示してくれる優れものである。
かつては GolangCI.com で [GitHub] と連携できていたのだが，[2020年4月でサービスが停止](https://medium.com/golangci/golangci-com-is-closing-d1fc1bd30e0e "GolangCI.com is closing. Dear customers of GolangCI.com, | by Denis Isaev | golangci | Medium")してしまい，寂しい限り。

と思っていたのだが，いつの間にか公式の [GitHub] Actions が用意されていた。
気付かなんだよ。
不覚。

- [golangci/golangci-lint-action: Official GitHub action for golangci-lint from it's authors](https://github.com/golangci/golangci-lint-action)

使い方は簡単。
リポジトリの `.github/workflows/` ディレクトリに YAML ファイル（例えば `lint.yml`）を置き，以下のように記述する。

```yaml
name: golangci-lint
on:
  push:
    tags:
      - v*
    branches:
      - master
  pull_request:
jobs:
  golangci:
    strategy:
      matrix:
        go-version: [1.15.x]
        os: [ubuntu-latest, macos-latest, windows-latest]
    name: lint
    runs-on: ${{ matrix.os }}
    steps:
      - uses: actions/checkout@v2
      - name: golangci-lint
        uses: golangci/golangci-lint-action@v2
        with:
          # Required: the version of golangci-lint is required and must be specified without patch version: we always use the latest patch version.
          version: v1.31

          # Optional: working directory, useful for monorepos
          # working-directory: somedir

          # Optional: golangci-lint command line arguments.
          # args: --issues-exit-code=0

          # Optional: show only new issues if it's a pull request. The default value is `false`.
          # only-new-issues: true
```

また，リポジトリ直下の `.gitattributes` ファイルに以下の記述を追加する。

```text
*.go text eol=lf
```

これで pull request 時， `master` ブランチ[^br1] への push 時，およびバージョンタグを打った際に [golangci-lint] が走る。
[golangci-lint] は `matrix` の組み合わせで並列処理されるようだ。

[^br1]: 2020年10月から [GitHub の新規リポジトリの既定ブランチ名が `main` になるらしい]({{< ref "/remark/2020/08/renaming-default-branch-name-in-github-repositries.md" >}} "GitHub リポジトリの既定ブランチ名が main になるらしい")。ご注意を。

{{< fig-img src="./reviews-in-pr.png" link="./reviews-in-pr.png" width="867" >}}

よーし，うむうむ，よーし。

まぁ，プラットフォーム依存のコードでもない限り [Go] 最新バージョンの `ubuntu-latest` だけでいいと思うけどね。

## ブックマーク

- [golangci/golangci-lint: Fast linters Runner for Go](https://github.com/golangci/golangci-lint)
- [reviewdog-golangci-lint を使う](https://zenn.dev/ikawaha/articles/57384e8fc69c7b057f7f)

- [golangci-lint に叱られる]({{< ref "/golang/donot-sleep-through-life.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[golangci-lint]: https://golangci-lint.run/
[GitHub]: https://github.com/

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
