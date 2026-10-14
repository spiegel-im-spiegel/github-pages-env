+++
title = "𝕏 に人間と認められなかったので活動を停止します"
date =  "2026-10-14T16:49:41+09:00"
description = "ブログのコミット履歴を圧縮しました / 𝕏 に人間と認められなかったので活動を停止します"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "site", "hugo", "git", "github", "twitter", "authentication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

今回は，このブログの最近の運用についてお知らせかたがた覚え書きを記しておく。

## ブログのコミット履歴を圧縮しました {#commit-log}

このブログは，[`github-pages-env`](https://github.com/spiegel-im-spiegel/github-pages-env "spiegel-im-spiegel/github-pages-env: Document Environment for spiegel-im-spiegel.github.io") リポジトリで作成環境を，[`spiegel-im-spiegel.github.io`](https://github.com/spiegel-im-spiegel/spiegel-im-spiegel.github.io "spiegel-im-spiegel/spiegel-im-spiegel.github.io: text.Baldanders.info") リポジトリでデプロイを管理している。[Hugo] と [GitHub Pages](https://docs.github.com/ja/pages "GitHub Pages ドキュメント - GitHubドキュメント") を利用。
基本的にブランチは切らずに `master` ブランチに直接コミットしている。
ひとりで運用してるので，こんなもんだろう。

なにせ10年以上もこの体制で運用しているもんで，コミットログも10年分溜まってしまってるのだ。
まぁ，どれだけログが溜まってようと普段は気にしないのだけど，最近は AI ([GitHub Copilot]) といっしょに記事やデプロイのチェックを行っていて，ときどきサンドボックス上に clone して動作テストをしたりするの。
その度に10年分のログがダウンロードされるわけで，遅い遅い。

というわけで「昨年までのコミットだけまとめられる？」と AI に訊いたら「できます」と言うのでサンドボックス上でやらせてみたんだけど，塩漬けのネタ用にブランチを切っていた関係でうまく rebase できなかったので，最終的に今年7月までのコミットをまとめることにした。
サンドボックス上で動作テストまでやらせて，最終的な作業手順をシェルスクリプトにまとめてもらった。
あとはスクリプトを実行して作業完了。

コミットをまとめた弊害もあって，このページの右上にもある "History in GitHub Page" で該当ページのコミット履歴を見れるようにしているのだが，今年7月までの履歴がなくなってしまった。
ゴメンペコン {{% emoji "ペコン" %}}

まぁ，でも，これでリポジトリがスッキリした。

## {{< emoji "X" >}} に人間と認められなかったので活動を停止します {#x}

[Firefox のコンテナ機能](https://support.mozilla.org/ja/kb/how-use-firefox-containers "Firefox コンテナーの使い方 | Firefox ヘルプ")がなかなか便利で最近よく使っているのだが，これの影響で {{< emoji "X" >}} のサインイン状態が外れちゃったのね。
「しょうがない」とサインインし直したのだが

{{< fig-img src="./x-verification-failed-1.png" title="あなたが人間であることを確認します" link="./x-verification-failed-1.png" width="698" >}}

とか出るの。
「は？」とか思ってコンテナやブラウザを変えて試してみるのだが状況は変わらない。
私は Twitter 時代からスマホに公式アプリを入れてなくて，ブラウザからしかアクセスしていない。
なので「スマートフォンでXアプリを開いて続行」とか言われても「知らんがな」となる。

あとで調べて知ったのだが，最近の {{< emoji "X" >}} ってユーザが「人間」であることを確認できないとサインインできなくなってるんだって？ ググったら，これのせいで iOS なんかが困らされてるという記事が挙がっていた。

試しに手持ちの Android タブレットに一時的にアプリを入れてサインインしてみようとしたのだが（スマホに入れるのはどうしても嫌）

{{< fig-img src="./x-verification-failed-2.png" title="あなたが人間であることを確認してください" link="./x-verification-failed-2.png" width="1200" >}}

とか表示が出て生体認証を要求するの。
なんだそれ？ もういいや！ 金払ってるわけじゃないし，サインインしなくても（自 TL が見れないだけで）推し VTuber とかのポストは見れるし，ゲーム情報は Discord でも見れるので実害はない（多分）。

というわけで [{{< emoji "X" >}} 上での活動](https://x.com/spiegel_2007)を完全停止します。
サインインできないから退会すらできないョ。
まぁ，そのうち勝手に BAN されるかもしれんけど（笑） ホンマどこまで粘着質なんだか。

[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[Hugo]: https://gohugo.io/ "The world's fastest framework for building websites"
