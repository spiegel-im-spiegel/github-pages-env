+++
title = "CI 用の GitHub Actions が諸々アップデートされていた"
date =  "2022-04-24T18:41:02+09:00"
description = "GitHub Actions アップデートまつりw"
image = "/images/attention/go-logo_blue.png"
tags = [ "github", "golang", "continuous-integration" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] で Excel ファイルを扱う [Excelize](https://github.com/qax-os/excelize) パッケージがバージョンアップしていた。

- [Go 言語スプレッドシートライブラリ：Excelize 2.6.0 がリリースされました](https://zenn.dev/xuri/articles/5a22e06f106657)

なので，このパッケージを使っている拙作の [goark/csvdata] パッケージもバージョンを上げた。

- [Release v0.5.1 · goark/csvdata · GitHub](https://github.com/goark/csvdata/releases/tag/v0.5.1)

ぶっちゃけ `go.mod` ファイルだけ更新してもよかったのだが，少し前に公式の GitHub Actions である [actions/setup-go](https://github.com/actions/setup-go) や [actions/checkout](https://github.com/actions/checkout) が v3 系に上がっているのに気付いたこともあり，諸々更新することにした。

- [Release v3.0.0 · actions/setup-go · GitHub](https://github.com/actions/setup-go/releases/tag/v3.0.0)
- [Release v3.0.2 · actions/checkout · GitHub](https://github.com/actions/checkout/releases/tag/v3.0.2)

## [github/codeql-action](https://github.com/github/codeql-action "github/codeql-action: Actions for running CodeQL analysis")

GitHub が買収して手に入れたコードチェッカの GitHub Actions。
Workflow の設定例はこんな感じらしい。

{{< fig-quote class="nobox" type="markdown" title="github/codeql-action: Actions for running CodeQL analysis" link="https://github.com/github/codeql-action" lang="en" >}}
```yaml
name: "Code Scanning - Action"

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
  schedule:
    #        ┌───────────── minute (0 - 59)
    #        │  ┌───────────── hour (0 - 23)
    #        │  │ ┌───────────── day of the month (1 - 31)
    #        │  │ │ ┌───────────── month (1 - 12 or JAN-DEC)
    #        │  │ │ │ ┌───────────── day of the week (0 - 6 or SUN-SAT)
    #        │  │ │ │ │
    #        │  │ │ │ │
    #        │  │ │ │ │
    #        *  * * * *
    - cron: '30 1 * * 0'

jobs:
  CodeQL-Build:
    # CodeQL runs on ubuntu-latest, windows-latest, and macos-latest
    runs-on: ubuntu-latest

    permissions:
      # required for all workflows
      security-events: write

      # only required for workflows in private repositories
      actions: read
      contents: read

    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      # Initializes the CodeQL tools for scanning.
      - name: Initialize CodeQL
        uses: github/codeql-action/init@v2
        # Override language selection by uncommenting this and choosing your languages
        # with:
        #   languages: go, javascript, csharp, python, cpp, java

      # Autobuild attempts to build any compiled languages (C/C++, C#, or Java).
      # If this step fails, then you should remove it and run the build manually (see below).
      - name: Autobuild
        uses: github/codeql-action/autobuild@v2

      # ℹ️ Command-line programs to run using the OS shell.
      # 📚 https://git.io/JvXDl

      # ✏️ If the Autobuild fails above, remove it and uncomment the following
      #    three lines and modify them (or add more) to build your code if your
      #    project uses a compiled language

      #- run: |
      #     make bootstrap
      #     make release

      - name: Perform CodeQL Analysis
        uses: github/codeql-action/analyze@v2
```
{{< /fig-quote >}}


[Go] コードのチェックをするなら “`Initialize CodeQL`” のコメント部分を解除して

```yaml {hl_lines=["5-6"]}
      # Initializes the CodeQL tools for scanning.
      - name: Initialize CodeQL
        uses: github/codeql-action/init@v2
        # Override language selection by uncommenting this and choosing your languages
        with:
          languages: go
```

とすればよい。

CodeQL は v2 系に上がったことでかなり深いところまでチェックするようになったようだ。
たとえば今回の [goark/csvdata] パッケージにはカラムの値を [`sql`][database/sql]`.NullByte` 型に変換して返す

```go
func (r *Rows) ColumnNullByte(s string, base int) (sql.NullByte, error) {
    res, err := r.ColumnNullInt64(s, base)
    if err != nil {
        return sql.NullByte{Valid: false}, errs.Wrap(err)
    }
    return sql.NullByte{Byte: byte(res.Int64), Valid: true}, nil
}
```

というメソッドがあるのだが，最後の

```go
return sql.NullByte{Byte: byte(res.Int64), Valid: true}, nil
```

で「範囲チェックなしで素のまま型変換すんな，ゴラァ（←超意訳）」と怒られてしまった。
素直な私は「なるほど」と納得して

```go {hl_lines=["6-9"]}
func (r *Rows) ColumnNullByte(s string, base int) (sql.NullByte, error) {
    res, err := r.ColumnNullInt64(s, base)
    if err != nil {
        return sql.NullByte{Valid: false}, errs.Wrap(err)
    }
    if res.Valid && (res.Int64 < 0 || res.Int64 > math.MaxUint8) {
        return sql.NullByte{Valid: false}, errs.Wrap(strconv.ErrRange)
    }
    return sql.NullByte{Byte: byte(res.Int64 & 0xff), Valid: true}, nil
}
```

と修正しましたとさ。
今までは何も言われなかったのに。
とほほ

## [golangci/golangci-lint-action](https://github.com/golangci/golangci-lint-action "golangci/golangci-lint-action: Official GitHub action for golangci-lint from its authors")

[golangci-lint](https://golangci-lint.run/) は [Go] 用の複合 linter。
Workflow の設定例はこんな感じ。

{{< fig-quote class="nobox" type="markdown" title="golangci/golangci-lint-action: Official GitHub action for golangci-lint from its authors" link="https://github.com/golangci/golangci-lint-action" lang="en" >}}
```yaml
name: golangci-lint
on:
  push:
    tags:
      - v*
    branches:
      - master
      - main
  pull_request:
permissions:
  contents: read
  # Optional: allow read access to pull request. Use with `only-new-issues` option.
  # pull-requests: read
jobs:
  golangci:
    name: lint
    runs-on: ubuntu-latest
    steps:
      - uses: actions/setup-go@v3
        with:
          go-version: 1.17
      - uses: actions/checkout@v3
      - name: golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          # Optional: version of golangci-lint to use in form of v1.2 or v1.2.3 or `latest` to use the latest version
          version: v1.29

          # Optional: working directory, useful for monorepos
          # working-directory: somedir

          # Optional: golangci-lint command line arguments.
          # args: --issues-exit-code=0

          # Optional: show only new issues if it's a pull request. The default value is `false`.
          # only-new-issues: true

          # Optional: if set to true then the all caching functionality will be complete disabled,
          #           takes precedence over all other caching options.
          # skip-cache: true

          # Optional: if set to true then the action don't cache or restore ~/go/pkg.
          # skip-pkg-cache: true

          # Optional: if set to true then the action don't cache or restore ~/.cache/go-build.
          # skip-build-cache: true
```
{{< /fig-quote >}}

こっちも v3 系に上がっているが， lint は日常的に使ってる（ていうか VS Code ならリアルタイムで走るようにできる）ので特に問題なし。
よかったよかった。

## [sonatype-nexus-community/nancy-github-action](https://github.com/sonatype-nexus-community/nancy-github-action "sonatype-nexus-community/nancy-github-action: Sonatype Nancy for GitHub Actions")

[Sonatype Nancy](https://github.com/sonatype-nexus-community/nancy "sonatype-nexus-community/nancy: A tool to check for vulnerabilities in your Golang dependencies, powered by Sonatype OSS Index") は [Go] の依存パッケージについて既知の脆弱性をチェックしてくれるツール。
こちらは特に変わってなかった。
Workflow の設定例はこんな感じ。

{{< fig-quote class="nobox" type="markdown" title="sonatype-nexus-community/nancy-github-action: Sonatype Nancy for GitHub Actions" link="https://github.com/sonatype-nexus-community/nancy-github-action" lang="en" >}}
```yaml
name: Go Nancy

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: Check out code into the Go module directory
      uses: actions/checkout@v2

    - name: Set up Go 1.x in order to write go.list file
      uses: actions/setup-go@v2
      with:
        go-version: ^1.13
    - name: WriteGoList
      run: go list -json -m all > go.list

    - name: Nancy
      uses: sonatype-nexus-community/nancy-github-action@main
```
{{< /fig-quote >}}

これも個人的に常用しているので無問題。

## [goreleaser/goreleaser-action]

みんな大好き，複数プラットフォームの実行バイナリを同時生成して GitHub のリリースページまで作ってくれる [GoReleaser](https://goreleaser.com/) の GitHub Actions。
こちらは v3 系に上がったね（2022-08-06 修正）。

{{< fig-quote class="nobox" type="markdown" title="goreleaser/goreleaser-action: GitHub Action for GoReleaser" link="https://github.com/goreleaser/goreleaser-action" lang="en" >}}
```yaml
name: goreleaser

on:
  pull_request:
  push:

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      -
        name: Checkout
        uses: actions/checkout@v3
        with:
          fetch-depth: 0
      -
        name: Set up Go
        uses: actions/setup-go@v3
      -
        name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v3
        with:
          # either 'goreleaser' (default) or 'goreleaser-pro'
          distribution: goreleaser
          version: latest
          args: release --rm-dist
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          # Your GoReleaser Pro key, if you are using the 'goreleaser-pro' distribution
          # GORELEASER_KEY: ${{ secrets.GORELEASER_KEY }}```
{{< /fig-quote >}}

この記事を書くのに [goreleaser/goreleaser-action] のページを眺めてて気がついたのだが， OpenPGP 電子署名も生成してくれるんだね。
Secret として隠蔽してくれるとはいえ， OpenPGP の秘密鍵やパスフレーズを GitHub 側に預託（escrow ← 言い方！）するのは抵抗があるなぁ。
まぁ，これは保留ということで。

## ブックマーク

- [golangci-lint を GitHub Actions で使う]({{< ref "/golang/using-golangci-lint-action.md" >}})
- [Go 依存パッケージの脆弱性検査]({{< ref "/golang/check-for-vulns-in-golang-dependencies.md" >}})
- [GitHub Actions でクロス・コンパイル（GoReleaser 編）]({{< ref "/golang/cross-compiling-in-github-actions-with-goreleaser.md" >}})
- [Go のコードでも GitHub Code Scanning が使えるらしい]({{< ref "/remark/2020/10/github-code-scanning-with-golang.md" >}})

[Go]: https://go.dev/
[goark/csvdata]: https://github.com/goark/csvdata "goark/csvdata: Reading CSV Data"
[database/sql]: https://pkg.go.dev/database/sql "sql package - database/sql - pkg.go.dev"
[goreleaser/goreleaser-action]: https://github.com/goreleaser/goreleaser-action "goreleaser/goreleaser-action: GitHub Action for GoReleaser"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
