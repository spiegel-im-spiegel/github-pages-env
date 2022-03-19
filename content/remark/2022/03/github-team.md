+++
title = "独り GitHub Team を作ってみた"
date =  "2022-03-13T20:31:26+09:00"
description = "ついカッとなってやった。今はちょっとだけ反省している。"
image = "/images/attention/kitten.jpg"
tags = [ "github", "golang", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

GitHub のリポジトリで諸々を公開するようになって後悔（駄洒落だよ）してるのは「ユーザ名が長すぎる」ことである。
ユーザ名は（名前が被らなければ）後から変更することもできるのだが（新しい方の URL にリダイレクトされる）， [Go] の場合はリポジトリの URI がそのままモジュール・パッケージのパスになるので移行が面倒そうだし，このブログ（GitHub Page）の CNAME 変更も面倒が起きそうな気がしたので諦めていた。

でも，半年ほど前に

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">nyaosorg という organization を作って、そちらへ nyagos のレポジトリを引っ越しました。nyagos が依存するパッケージも順次そちらへ移動する予定です<a href="https://t.co/yFwrGxdHiP">https://t.co/yFwrGxdHiP</a></p>&mdash; 𝒏𝒚𝒂𝒐𝒔.𝒐𝒓𝒈 (@NyaosOrg) <a href="https://twitter.com/NyaosOrg/status/1436892273489301506?ref_src=twsrc%5Etfw">September 12, 2021</a></blockquote>
{{< /fig-gen >}}

という tweet を見て「なるほど！」と思ったわけよ。
もっとも本業がだんだん忙しくなって何となく先延ばしにしてたのだが，昨日ついカッとなってやってもうた。

- [github.com/goark Playing with Go Language](https://github.com/goark)

忙しいときほど要らんことをしたくなるよね（部屋の掃除とか）。
今はちょっとだけ反省している。

“ark” は archive の駄洒落である[^d1]。
なので「ごーあーく」とでも呼んでいただければ幸いである。
もしくは更に縮めて「ごらく」とか（読めねーよ）

[^d1]: エンジニアに必要なのはダジャレ力だって [Go の偉い人が言ってた](https://twitter.com/mattn_jp/status/1502146966855495682)（笑）

用途としては，タイトル通り，個人的な「[Go] の遊び場」として利用することを考えとりやす。
元の [github.com/spiegel-im-spiegel](https://github.com/spiegel-im-spiegel) から [Go] パッケージを徐々に移行する予定。
とりあえず [gpgpdump] と [depm] は個人的に使うので（依存パッケージと併せて）真っ先に移行した。
他のアクティブなパッケージもそのうち移行するつもりである。

移行自体は簡単で，リポジトリの Settings ページの下の方に “Danger Zone” てのがあるのだが

{{< fig-img src="./danger-zone.png" title="Danger Zone" link="./danger-zone.png" width="794" lang="en" >}}

その中の “Transfer ownership” で移行先のユーザ名（または組織名）を指定すればよい。

{{< fig-img src="./transfer.png" title="Transfer ownership" link="./transfer.png" lang="en" >}}

移行するためには移行先にリポジトリの作成権限があることが必要。

更に [Go] パッケージについてはソースコード内のインポート・パスを書き換えないといけないのだが，実際にやってみると思ったほど大変ではなかった。
`go.mod` ファイルに依存パッケージが列挙されているので，これを見ながら一括置換してしまえばいいのだ。

注意する点としては移行後にバージョンタグを付けてバージョンを上げておくこと。
そうしないと，ひとつのバージョンに対して複数のパスが存在することになり `go mod tidy` とかでがっつり怒られる。

そうそう。
今回は Free ではなく，有料の Team にした。
だって Codespace が使いたかったんじゃもん。

{{< fig-img-quote src="./pricing.png" title="Pricing · Plans for every developer" link="https://github.com/pricing" width="858" lang="en" >}}

今や IDE やエディタも XaaS の時代ですよ。
まぁ，飽きたらフリーに戻すかも知れんけど。

## ブックマーク

- [GitHub のアカウントの名前を変更したら起きること - 標準愚痴出力](https://zetamatta.hatenablog.com/entry/2022/02/02/183120)
  - [続・GitHub のアカウントの名前を変更したら起きること 〈AppVeyor編〉 - 標準愚痴出力](https://zetamatta.hatenablog.com/entry/2022/03/17/192943) : 結局 GitHub のユーザ名を変えて一番面倒くさいのは OAuth 周りなんだな。やっぱ下手に変えないほうがいいか
- [パッケージ引っ越し大作戦](https://zenn.dev/zetamatta/scraps/e622959b4c34eb)
- [GitHubセキュリティ Organization運用のベストプラクティス](https://zenn.dev/tmknom/books/github-organization-security)
- [GitHub開発チームでのCodespacesの利用 - GitHubブログ](https://github.blog/jp/2021-08-30-githubs-engineering-team-moved-codespaces/)
- [GitHub Codespaces · GitHub](https://github.com/features/codespaces)
- [GitHub Codespaces Documentation - GitHub Docs](https://docs.github.com/en/codespaces)
- [Github CodeSpace 触ってみた - Qiita](https://qiita.com/Alt225/items/5d904fafc779e6505768)
- [GitHub Codespaces をつかって 3分で始めるサービス開発 | Wantedly, Inc.](https://www.wantedly.com/companies/wantedly/post_articles/355862)

- [Go パッケージの移行状況]({{< ref "/release/2022/03/status-of-migrations.md" >}})

[github.com/goark]: https://github.com/goark "Playing with Go Language"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"
[depm]: https://github.com/goark/depm "goark/depm: Visualize depndency packages and modules"
[Go]: https://go.dev/

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
