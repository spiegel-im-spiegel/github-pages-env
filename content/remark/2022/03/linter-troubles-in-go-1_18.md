+++
title = "Go 1.18 と Generics と Linter"
date =  "2022-03-19T15:16:40+09:00"
description = "Go 1.18 に組み込まれた Generics と既存の lint の間でトラブルが続出しているらしい。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "generics", "lint" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予測して然るべきであったが [Go] 1.18 に組み込まれた Generics と既存の lint の間でトラブルが続出しているらしい。

私も大変お世話になっている [Go] の代表的な総合 linter である [golangci-lint] は [v1.45](https://github.com/golangci/golangci-lint/releases/tag/v1.45.0 "Release v1.45.0 · golangci/golangci-lint") で暫定的な対応を行ったようだ。

{{< fig-quote type="markdown" title="Support of go1.18 · Issue #2649 · golangci/golangci-lint" link="https://github.com/golangci/golangci-lint/issues/2649#issuecomment-1070528726" lang="en" >}}
So I added a new configuration option inside the run section.<br>
We will not have to remove this option in the future because we will be able to use it for some linters that have a Go version option.

An example:

```yaml
run:
    go: 1.18
```

By default, the supported Go version will be go1.17 because it allows us to run all our tests without a huge rewrite.

If you set the version to go1.18, the following linters will be inactive:

- `bodyclose`
- `contextcheck`
- `gosimple`
- `nilerr`
- `noctx`
- `rowserrcheck`
- `sqlclosecheck`
- `staticcheck`
- `stylecheck`
- `tparallel`
- `unparam`
- `unused`
- `wastedassign`
{{< /fig-quote >}}

この説明にあるように `.golangci.yaml` ファイルに

```yaml
run:
    go: 1.18
```

の記述を入れることで [Go] 1.18 に対応していない lint を無効にしてくれるようだ。
あとは各 lint パッケージが 1.18 に対応してくれることを気長に待つしかないか。

なお Generics の機能を使わないのなら今回の件は気にしなくてよい（笑）

あと [golangci-lint-action] もいつの間にか v3 に上がってるな。
見落としてたかなぁ。
この辺はおいおい検証しよう。

[Go]: https://go.dev/
[golangci-lint]: https://github.com/golangci/golangci-lint/ "golangci/golangci-lint: Fast linters Runner for Go"
[golangci-lint-action]: https://github.com/golangci/golangci-lint-action "golangci/golangci-lint-action: Official GitHub action for golangci-lint from its authors"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
