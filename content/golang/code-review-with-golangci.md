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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4542503461?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51CAFNAdZPL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4542503461?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">組込み開発者におくるMISRA‐C:2004―C言語利用の高信頼化ガイド</a></dt>
    <dd>MISRA‐C研究会 (編集)</dd>
    <dd>日本規格協会 2006-10-01</dd>
    <dd>単行本</dd>
    <dd>4542503461 (ASIN), 9784542503465 (EAN), 4542503461 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">私が持っているのはこれよりひとつ古い版だが，まぁいいか。むかし，車載用の組み込みエンジニアをやっていた頃は必読書として読まされました。今はもっと包括的な内容のものがあるはず。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-02-06">2019-02-06</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
