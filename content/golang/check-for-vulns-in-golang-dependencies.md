+++
title = "Go 依存パッケージの脆弱性検査"
date =  "2020-09-30T12:49:21+09:00"
description = "nancy を使うのがよさげである。"
image = "/images/attention/go-logo_blue.png"
tags = [ "security", "vulnerability", "risk", "management", "golang", "package", "module", "github", "continuous-integration" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

どの言語でも同じだけど，インポートする外部パッケージが安全かどうかを調べるのはけっこう大変である。
で， [Go 言語][Go]の場合は [nancy] を使うのがよさげである。

- [sonatype-nexus-community/nancy: A tool to check for vulnerabilities in your Golang dependencies, powered by Sonatype OSS Index](https://github.com/sonatype-nexus-community/nancy)

[nancy] は “[Sonatype OSS Index]” の情報を使って依存パッケージ／モジュールの検査をしてくれる。
ツール自体は Apache-2.0 でライセンスされている。
提供されているデータについては

{{< fig-quote type="markdown" title="Sonatype OSS Index" link="https://ossindex.sonatype.org/" lang="en" >}}
{{% quote %}}OSS Index and the associated tools are and always will be free to the community. The data we gather is derived from public sources, and does not include human curated intelligence nor expert remediation guidance{{% /quote %}}.
{{< /fig-quote >}}

とあるので，オープンな場で使うなら問題ないかな。

使い方は簡単で，開発中のパッケージのリポジトリ上で

```text
$ go list -json -m all | nancy sleuth
```

とすればよい。
問題なければ

{{< fig-gen type="markdown" lang="en" >}}
```text
$ go list -json -m all | nancy sleuth -n
┏━━━━━━━━━━━━━━━┓
┃ Summary                      ┃
┣━━━━━━━━━━━━━┳━┫
┃ Audited Dependencies     ┃ 9┃
┣━━━━━━━━━━━━━╋━┫
┃ Vulnerable Dependencies  ┃ 0┃
┗━━━━━━━━━━━━━┻━┛
```
{{< /fig-gen >}}

みたいな感じで結果を返してくれる。
問題のあるパッケージ／モジュールを含んでると，ものすごい勢いで叱られるけど（笑）

## [GitHub] Actions でも使える

[nancy] には [GitHub] Actions も用意されている。
ありがたや。

- [sonatype-nexus-community/nancy-github-action: Sonatype Nancy for GitHub Actions](https://github.com/sonatype-nexus-community/nancy-github-action)

設定は簡単。
リポジトリの `.github/workflows/` ディレクトリに YAML ファイル（例えば `vulns.yml`）を置き，以下のように記述する。

```yaml
name: vulns
on:
  push:
    tags:
      - v*
    branches:
      - master
  pull_request:
jobs:
  vulns:
    name: Vulnerability scanner
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: ^1.13
      - name: WriteGoList
        run: go list -json -m all > go.list
      - name: Nancy
        uses: sonatype-nexus-community/nancy-github-action@main
```

これで pull request 時， `master` ブランチ[^br1] への push 時，およびバージョンタグを打った際に脆弱性検査が走る。

[^br1]: 2020年10月から [GitHub の新規リポジトリの既定ブランチ名が `main` になるらしい]({{< ref "/remark/2020/08/renaming-default-branch-name-in-github-repositries.md" >}} "GitHub リポジトリの既定ブランチ名が main になるらしい")。ご注意を。

## 依存の依存パッケージに脆弱性がある

直接インポートするパッケージに脆弱性があるなら無害なバージョンに差し替えればいいけど，依存パッケージが依存しているパッケージに脆弱性がある場合はどうするか。

とりあえず，そのパッケージ宛てには issue を投げておくとして，それまでの継ぎとしては `go.mod` ファイルの `replace` ディレクティブを使って凌ぐことができそうだ。

たとえば，依存パッケージが [github.com/coreos/etcd] v3.3.13 に依存してるんだけど v3.3.13 に脆弱性がある場合，

```text
replace (
	github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
)
```

などとして無害なバージョンに差し替えできる。

`require` ディレクティブで

```text
require (
	github.com/coreos/etcd v3.3.25+incompatible
)
```

とか書いても同じ効果があるけど，名目だけの依存関係で実際にはインポートしないパッケージは `go mod tidy` コマンドで記述が削除されちゃうのでオススメできない。

`go list -m all` って，実際にはリンクしない名目上の依存関係も全部拾ってリストアップしちゃうので，凄い面倒くさいんだよねぇ。
実際にリンクするパッケージだけリストアップしてくれないものだろうか...

## 【2021-02-25 追記】 [depm] との連携

拙作の [Go 言語用モジュール依存関係可視化ツール]({{< ref "/release/dependency-graph-for-golang-modules.md" >}}) [depm] を使って

```text
$ depm list --json | nancy sleuth -n
```

とすることで名目だけの依存パッケージの誤検知を回避できる。
ただ，私としては [depm] の信頼性にイマイチ確信が持てないので，ご利用は自己責任で（笑）

[GitHub] Actions を使うのであれば

```yaml {hl_lines=[18,19,21]}
name: vulns
on:
  push:
    tags:
      - v*
    branches:
      - master
  pull_request:
jobs:
  vulns:
    name: Vulnerability scanner
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: ^1.16
      - name: install depm
        run: go install github.com/spiegel-im-spiegel/depm@latest
      - name: WriteGoList
        run: depm list --json > go.list
      - name: Nancy
        uses: sonatype-nexus-community/nancy-github-action@main
```
などとすればOK。
[Go] 1.16 以降で有効なのでご注意を。

## ブックマーク

- [CI 用の GitHub Actions が諸々アップデートされていた]({{< ref "/golang/update-github-actions.md" >}})

[Go]: https://go.dev/
[nancy]: https://github.com/sonatype-nexus-community/nancy "sonatype-nexus-community/nancy: A tool to check for vulnerabilities in your Golang dependencies, powered by Sonatype OSS Index"
[Sonatype OSS Index]: https://ossindex.sonatype.org/
[GitHub]: https://github.com/
[github.com/coreos/etcd]: https://github.com/etcd-io/etcd "etcd-io/etcd: Distributed reliable key-value store for the most critical data of a distributed system"
[depm]: https://github.com/spiegel-im-spiegel/depm "spiegel-im-spiegel/depm: Visualize depndency packages and modules"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
