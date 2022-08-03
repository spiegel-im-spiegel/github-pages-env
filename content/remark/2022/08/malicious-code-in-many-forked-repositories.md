+++
title = "悪意のコードを含む多数の分岐リポジトリが見つかった話"
date =  "2022-08-04T08:12:52+09:00"
description = "今回の件は spam の一種とみなすこともできる"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "spam", "engineering", "managemant" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2022-08-03 頃の Twitter TL でちょっとした騒ぎを見かけたのだが，詳細記事が出たらしい。

- [35,000 code repos not hacked—but clones flood GitHub to serve malware](https://www.bleepingcomputer.com/news/security/35-000-code-repos-not-hacked-but-clones-flood-github-to-serve-malware/)

最初は各リポジトリに悪意のコードがねじ込まれたのかと思って，コード事態も問題だけど，リポジトリ・アクセスへの認証周りで何か問題があったのかと思ったが，実はそうではなく

{{< fig-quote type="markdown" title="35,000 code repos not hacked—but clones flood GitHub to serve malware" link="https://www.bleepingcomputer.com/news/security/35-000-code-repos-not-hacked-but-clones-flood-github-to-serve-malware/" lang="en" >}}
Rather, the thousands of backdoored projects are copies (forks or clones) of legitimate projects purportedly made by threat actors to push malware.
{{< /fig-quote >}}

ということらしい。

知らない人のために一応解説しておくと GitHub には pull request という仕組みがあり，他者のリポジトリにコードを貢献したい場合に自身のリポジトリに分岐（fork）させた上でコードを変更・追加し，そのコードを対象のリポジトリにマージするよう提案を行うことができる。
提案を受けた側は，そのコードのレビューを行った上で明示的な操作でコードを受け入れることができる。
もちろんダメなら拒否もできる。

Pull request は GitHub アカウントを持つユーザなら誰でも可能であり，その過程で悪意のコードを忍ばせることは形式上は可能である。
まぁ，普通は「そういう PR はレビューで拒否しましょうね」となるだろうし，一度そんなコードを送りつけた相手を二度と信用することはないだろう。
なので今回の件は spam の一種とみなすこともできる。
実際，アホみたいな数だしね。

気をつける点があるとすれば「悪意の PR」を送りつけるために作った分岐リポジトリのコードをうっかり取り込んでしまう場合だろう。
最初に紹介した記事では

{{< fig-quote type="markdown" title="35,000 code repos not hacked—but clones flood GitHub to serve malware" link="https://www.bleepingcomputer.com/news/security/35-000-code-repos-not-hacked-but-clones-flood-github-to-serve-malware/" lang="en" >}}
As a best practice, remember to consume software from the official project repos and watch out for potential typosquats or repository forks/clones that may appear identical to the original project but hide malware.
{{< /fig-quote >}}

と述べている。
言われんでも（笑）

それに続けて

{{< fig-quote type="markdown" title="35,000 code repos not hacked—but clones flood GitHub to serve malware" link="https://www.bleepingcomputer.com/news/security/35-000-code-repos-not-hacked-but-clones-flood-github-to-serve-malware/" lang="en" >}}
Open source code commits signed with GPG keys of authentic project authors are one way of verifying the authenticity of code.
{{< /fig-quote >}}

と書かれているが，ぶっちゃけプロジェクトの外側にいる人から見てコミットに OpenPGP 署名があることは大した保証にはならない。
この辺は以前書いた[拙文]({{< ref "/openpgp/web-of-trust.md" >}} "OpenPGP の電子署名は「ユーザーの身元を保証し」ない")を参考にしてほしい。

{{< fig-quote type="markdown" title="OpenPGP の電子署名は「ユーザーの身元を保証し」ない" link="https://text.baldanders.info/openpgp/web-of-trust/" >}}
じゃあ git commit で OpenPGP 署名を付与することにどんな意義があるかというと，それはチーム運営で威力を発揮する。
つまり公開鍵や電子署名で「ユーザーの身元を保証」するのではなく「身元の保証されたユーザ」同士で鍵と電子署名を運用するのである。
これでチーム以外からのなりすまし commit を検知（防止ではない）しやすくなる。
{{< /fig-quote >}}

オープンソース・プロジェクトであれば pull request を受け付けるメンテナの技量と判断が試されるところだろう。

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
