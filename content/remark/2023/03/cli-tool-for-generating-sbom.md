+++
title = "ソフトウェア部品表（SBOM）を生成するツール"
date =  "2023-03-26T16:07:29+09:00"
description = "Go 以外のエコシステムでも使える。 grype と組み合わせて脆弱性のチェックも可能，らいい。"
image = "/images/attention/kitten.jpg"
tags = [ "golang", "tools", "security", "management", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] コードのビルドとリリースを一度にやってくれる [GoReleaser] というツールがあるのだが，これの最近のバージョンはソフトウェア部品表（Software Bill of Materials; SBOM）も生成・リリースできるらしい。
というわけで，[自作ツール](https://github.com/goark/ml/releases/tag/v0.6.6 "Release v0.6.6 · goark/ml")でちょっと試してみた。

設定自体は難しくなく `.goreleaser.yaml` ファイルに以下の記述を追加すればいいだけのようだ。

```yaml
sboms:
  - artifacts: archive
```

早速この記述を追加して手元で動かしてみたのだが...

```text
 • cataloging artifacts
    • cataloging                      artifact=dist/ml_SNAPSHOT-a83f2d0b1db0ade89d839cd70b6870cd90011f55_Windows_ARM64.zip cmd=syft sboms=ml_SNAPSHOT-a83f2d0b1db0ade89d839cd70b6870cd90011f55_Windows_ARM64.zip.sbom
  ⨯ release failed after 1s   error=cataloging artifacts: syft failed: exec: "syft": executable file not found in $PATH: 
```

ふむむ？ `syft` がないって言ってるのか？ 調べてみたら [syft] というのはこれのことらしい。

- [anchore/syft: CLI tool and library for generating a Software Bill of Materials from container images and filesystems](https://github.com/anchore/syft)

[GoReleaser] は内部で [syft] を起動して SBOM を生成しているようだ。
SBOM を生成するための設定を `.goreleaser.yaml` ファイルに記述する際の詳細情報は以下のページが参考になる。

- [Cataloging artifacts - GoReleaser](https://goreleaser.com/customization/sbom/)

つか，最初からマニュアルを読めっての！＞自分

GitHub Actions の [goreleaser-action] を使って SBOM を含むリリースを行う場合は，自前で [syft] をインストールする必要がある。
とはいえ  [goreleaser-action] を使うなら [Go] のコンパイラは事前に導入されているはずなので，簡単に

```yaml { hl_lines=["25-27"] }
name: build

on:
  push:
    tags:
      - v*

permissions:
  contents: write

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
        with:
          go-version-file: 'go.mod'
      -
        name: install syft
        run: go install github.com/anchore/syft/cmd/syft@latest
      -
        name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v4
        with:
          # either 'goreleaser' (default) or 'goreleaser-pro'
          distribution: goreleaser
          version: latest
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          # Your GoReleaser Pro key, if you are using the 'goreleaser-pro' distribution
          # GORELEASER_KEY: ${{ secrets.GORELEASER_KEY }}
```

といった感じに `go install` コマンドでインストールしてしまっても問題なく行けるっぽい。
でも，これをすると SBOM に [syft] のバージョンが入らなくなるんだよなぁ。
[sbom-action] と組み合わせればいいのだろうか。
要検証だな。

[syft] は [Go] 製だが [Go] エコシステム専用というわけではなく，以下のものに対応しているらしい。

{{< fig-quote type="markdown" title="syft/README.md at main · anchore/syft" link="https://github.com/anchore/syft/blob/main/README.md" lang="en" >}}
- Alpine (apk)
- C (conan)
- C++ (conan)
- Dart (pubs)
- Debian (dpkg)
- Dotnet (deps.json)
- Objective-C (cocoapods)
- Elixir (mix)
- Erlang (rebar3)
- Go (go.mod, Go binaries)
- Haskell (cabal, stack)
- Java (jar, ear, war, par, sar, native-image)
- JavaScript (npm, yarn)
- Jenkins Plugins (jpi, hpi)
- PHP (composer)
- Python (wheel, egg, poetry, requirements.txt)
- Red Hat (rpm)
- Ruby (gem)
- Rust (cargo.lock)
- Swift (cocoapods)
{{< /fig-quote >}}

んー。
メジャーどころは網羅してる感じ？

他にも [grype](https://github.com/anchore/grype "anchore/grype: A vulnerability scanner for container images and filesystems") と組み合わせることで脆弱性のチェックとかもできるし，その結果を証明書として作成して発行することもできるそうな。
私が公開しているような小物パッケージではそこまで不要だろうが，企業とかが運用している，それなりに規模の大きなプロジェクトでは重宝するかもしれない。

覚えておこう。

## ブックマーク

- [米国の国家サイバーセキュリティ戦略とインフラとしてのオープンソース – WirelessWire News](https://wirelesswire.jp/2023/03/84355/)

- [Go はどのようにしてサプライチェーン攻撃を低減しているか](https://zenn.dev/spiegel/articles/20220402-how-go-mitigates-supply-chain-attacks)
- [オープンソース製品とソフトウェア部品表]({{< ref "/remark/2022/08/software-bills-of-materials.md" >}})

[Go]: https://go.dev/
[GoReleaser]: https://goreleaser.com/
[goreleaser-action]: https://github.com/goreleaser/goreleaser-action "goreleaser/goreleaser-action: GitHub Action for GoReleaser"
[syft]: https://github.com/anchore/syft "anchore/syft: CLI tool and library for generating a Software Bill of Materials from container images and filesystems"
[sbom-action]: https://github.com/anchore/sbom-action "anchore/sbom-action: GitHub Action for creating software bill of materials using Syft."

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
