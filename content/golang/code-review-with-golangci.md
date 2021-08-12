+++
title = "GolangCI でコード・レビューを自動化する"
date =  "2019-12-15T14:05:25+09:00"
description = "リポジトリ全体をチェックしてくれるレビュー・サービスの存在はありがたい。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "lint", "programming", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

{{< div-box type="markdown" >}}
この記事で紹介した [GolangCI] のレビュー・サービスは 2020-04-15 で終了するらしい。

- [GolangCI.com is closing - golangci - Medium](https://medium.com/golangci/golangci-com-is-closing-d1fc1bd30e0e)

まぁ [Go 言語]自体がもっと普及しないと営利サービスとしては難しいのかもね。
残念ではあるが lint ツール自体は公開を続けるそうな。
うんうん，よかったよかった。

[GolangCI]: https://golangci.com/ "Automated code review for Go"
[Go 言語]: https://golang.org/ "The Go Programming Language"
{{< /div-box >}}

以前「[golangci-lint に叱られる]({{< relref "./donot-sleep-through-life.md" >}})」で

> {{% quote %}}[GolangCI](https://golangci.com/ "Automated code review for Go") も気になるが，それはまたいつか{{% /quote %}}

と書いたが，今回はその話。

[GolangCI] は [GitHub] と連携して機能するコード・レビュー・サービスで， [GitHub] 上のリポジトリにある [Go 言語]コードに lint をかけて結果を報告してくれる。
Pull request とも連携してレビュー結果を上げてくれるので，レビューにかかる労力をかなり引き下げることができる[^lint1]。

[^lint1]: とはいえ [GolangCI] がチェックしてくれるのは lint レベルのコード・チェックなので，ビジネスロジック等の妥当性は人間が判断する必要がある。言い方を変えれば [GolangCI] で lint レベルのチェックを事前に行っておけば，人間はビジネスロジック等のチェックに専念することができる。

サインアップは [GitHub] アカウントで行うことができる。
サインアップに成功したらリポジトリ一覧画面に行けるようになる[^repos1]。

[^repos1]: 既定では公開リポジトリのみが対象となる。お金を払えばプライベート・リポジトリもチェックできるらしいが試していない。

{{< fig-img src="./list-of-reps.jpg" link="./list-of-reps.jpg" width="850" >}}

連携したいリポジトリに `[Connect]` するとレビューを開始するのだが，レビュー結果が表示されるまで結構時間がかかるみたい。
実は随分前に [GolangCI] に登録していくつかのパッケージと連携させていたのだが，なかなか処理が終わらないので，そのまま綺麗サッパリ忘れ去っていたのだった（笑）

レビューが終わったリポジトリの `[Report]` を開くとこんな感じの画面になる。
これは問題がなかったリポジトリの場合：

{{< fig-img src="./result-review-no-error.jpg" link="./result-review-no-error.jpg" width="850" >}}

こっちは問題が発見されたリポジトリ：

{{< fig-img src="./result-review-error.jpg" link="./result-review-error.jpg" width="850" >}}

このレポート結果を基にコードを修正する。
当然ながら master ブランチ上で作業をしないこと。

修正を commit & push し，修正を行ったブランチから pull request をかける。
[GitHub] 側は pull request したコードに対して連携しているサービスを呼び出して事前チェックを行う。
チェックにパスすればこんな感じになる。

{{< fig-img src="./pull-request.jpg" link="./pull-request.jpg" width="850" >}}

問題があれば “Details” で問題箇所が示されるので，修正を行って再度 commit & push する。

[GolangCI] には[コマンドライン・ツール](https://github.com/golangci/golangci-lint "golangci/golangci-lint: Linters Runner for Go. 5x faster than gometalinter. Nice colored output. Can report only new issues. Fewer false-positives. Yaml/toml config.")も用意されていて IDE やテキストエディタなどとも連携可能なのだが[^atom1]，どうしても見落としがあるみたいで，リポジトリ全体をチェックしてくれるレビュー・サービスの存在はありがたい。

[^atom1]: [ATOM] エディタの場合は [go-plus](https://atom.io/packages/go-plus) パッケージで Linter に [golangci-lint] を指定できる。

さて，他のパッケージも修正するか。

## ブックマーク

- [無料で使える Go 言語の CI サービス『GolangCI』を使ってみる ｜ Developers.IO](https://dev.classmethod.jp/go/golangci/)

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[GolangCI]: https://golangci.com/ "Automated code review for Go"
[golangci-lint]: https://github.com/golangci/golangci-lint "golangci/golangci-lint: Linters Runner for Go. 5x faster than gometalinter. Nice colored output. Can report only new issues. Fewer false-positives. Yaml/toml config."
[GitHub]: https://github.com/ "The world’s leading software development platform · GitHub"
[ATOM]: https://atom.io/

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->

{{% review-paapi "4542503461" %}} <!-- 組込み開発者におくるMISRA‐C:2004 -->
