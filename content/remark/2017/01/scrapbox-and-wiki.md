+++
draft = false
tags = ["wiki", "slack", "tools"]
description = "で，当面の問題は私が Wiki 記法をほとんど忘れてるってことなんだよなぁ。"
date = "2017-01-10T20:21:13+09:00"
update = "2017-04-04T17:05:40+09:00"
title = "Scrapbox または Wiki で再び遊ぶ"

[author]
  facebook = "spiegel.im.spiegel"
  url = "http://www.baldanders.info/spiegel/profile/"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  flattr = "spiegel"
  name = "Spiegel"
  flickr = "spiegel"
  instagram = "spiegel_2007"
  tumblr = "spiegel-im-spiegel"
  github = "spiegel-im-spiegel"
  avatar = "/images/avatar.jpg"
  twitter = "spiegel_2007"
+++

私事だが [Scrapbox] のアカウントを取った。
[Scrapbox] に興味を持ったのは以下の記事がきっかけ。

- [Wiki脳だってTumblrしたい - Scrapboxの話をしよう｜塚本 牧生｜note](https://note.mu/tsukamoto/n/n9c6a0ea7030b)

ゼロ年代後半にあれだけハマった Twitter や Tumblr について，私は全く関心をなくしてしまった（アカウントはまだあるしボット経由で記事というか activity も上げてるけどね）。
それは多分「ネットのテレビ化」と連動している気がする。
たしかに Twitter や Tumblr は「消費行動」としては面白いが，それ以上がない。

で，ここ数年は Facebook に引きこもっている。
はっきり言って私が書き散らしたものでもっとも反応があるのは Facebook の TL に書いたものである。
それは単に「いいね」だけじゃなく反論やお叱りも込みで反応をいただける。
有り難いことである。

ただ，個人的に Facebook での発言は公開範囲を限定していることと TL 上の記事は過去のものほど検索が難しくなることがネックになっていて，どうやってアウトプットしようかと悩んではいた。
TL をもつサービスでは，上げたものはただ流れ去るだけで上げた本人もそのことを忘れてしまう。
忘れてしまったものは検索できない。

そこで [Scrapbox] の出番である。

まずは公開プロジェクト “[Spiegel's Branch]” を作成した。
当面はブログ記事にするまでもない戯れ言や Facebook TL に書いた発言で一般公開してもいいものを “[Spiegel's Branch]” に上げていくことにする。
使い勝手が悪かったり飽きたりしたら止めるかもだけど。

[Scrapbox] は Wiki サービスである。

[Scrapbox] には複数の Project があり，各 Project ごとにユーザ管理やページの管理ができる。
また Project は（名前が被らない限り）ユーザが自由に作成でき，作成した Project は公開にも非公開にもできる（作成直後は非公開）。

なお，ヘルプは [help Project](https://scrapbox.io/help/ "Scrapbox ヘルプ - Scrapbox") として公開されている。

- Wiki なので当然ながら寄ってたかってページを作成・編集できる。複数人で編集する場合は Project ごとに Team を編成する。[招待制](https://scrapbox.io/help/%E3%83%A1%E3%83%B3%E3%83%90%E3%83%BC%E3%82%92%E8%BF%BD%E5%8A%A0%E3%81%99%E3%82%8B "メンバーを追加する - Scrapbox ヘルプ - Scrapbox")。
- 入力文法はいわゆる Wiki 記法だけど，色々と拡張されていて，[ソースコードの記述](https://scrapbox.io/help/%E3%82%B3%E3%83%BC%E3%83%89%E3%83%96%E3%83%AD%E3%83%83%E3%82%AF%E8%A8%98%E6%B3%95 "コードブロック記法 - Scrapbox ヘルプ - Scrapbox")もできるし[アイコンの設置](https://scrapbox.io/help/%E3%82%A2%E3%82%A4%E3%82%B3%E3%83%B3%E8%A8%98%E6%B3%95 "アイコン記法 - Scrapbox ヘルプ - Scrapbox")もできる。
- [Scrapbox] には[リンクが2種類ある](https://scrapbox.io/help/%E3%83%9A%E3%83%BC%E3%82%B8%E3%82%92%E3%83%AA%E3%83%B3%E3%82%AF%E3%81%99%E3%82%8B "ページをリンクする - Scrapbox ヘルプ - Scrapbox")ようだ。ひとつは Wiki 記法による通常のリンク。もうひとつは `#tags` と記述すると自動作成されるリンク。 `#tags` 記法でタグ付けしておくとプロジェクト内の記事を横串で集約してくれる。これめっちゃ便利。
- 作成したページは [JSON 形式で Export](https://scrapbox.io/help/%E3%82%A4%E3%83%B3%E3%83%9D%E3%83%BC%E3%83%88%E3%83%BB%E3%82%A8%E3%82%AF%E3%82%B9%E3%83%9D%E3%83%BC%E3%83%88%E3%81%99%E3%82%8B "インポート・エクスポートする - Scrapbox ヘルプ - Scrapbox") できる。同様に JSON 形式での Import もできる。これでいつでも撤退できる。
- [Slack と連動可能](https://scrapbox.io/help/Slack%E3%81%AB%E6%9B%B4%E6%96%B0%E3%82%92%E9%80%9A%E7%9F%A5%E3%81%99%E3%82%8B "Slackに更新を通知する - Scrapbox ヘルプ - Scrapbox")。 [Slack] のチーム/チャネルと [Scrapbox] のプロジェクトを組み合わせれば便利かもしれない。
- [ヘルプ](https://scrapbox.io/help/ "Scrapbox ヘルプ - Scrapbox")には記載が見つからないがフィードを吐く。たとえば [Spiegel's Branch] （ https://scrapbox.io/spiegel-branch/ ）なら https://scrapbox.io/api/feed/spiegel-branch がフィードの URL になるようだ。
- [オンプレミス版](https://scrapbox.io/help/%E3%82%AA%E3%83%B3%E3%83%97%E3%83%AC%E7%89%88%E3%81%AE%E5%88%A9%E7%94%A8%E3%82%AC%E3%82%A4%E3%83%89 "オンプレ版の利用ガイド - Scrapbox ヘルプ - Scrapbox")（[エンタープライズ版](https://scrapbox.io/enterprise "Scrapbox - Enterprise")）がある！

で，当面の問題は私が Wiki 記法をほとんど忘れてるってことなんだよなぁ[^w]。
Markdown に慣れすぎてしまった。

[^w]: 昔は[本家サイト](http://www.baldanders.info/ "Baldanders.info")に Wiki を設置して遊んでいたが， spam の対処に困って閉鎖してしまった。ユーザ認証とかしてなかったからなぁ。

## ブックマーク

- [Scrapboxページの文章をMarkdownに変換するBookmarkletを書いた - #daiizメモ](http://daiiz.hatenablog.com/entry/2017/02/17/074508)

[Scrapbox]: https://scrapbox.io/ "Scrapbox - A new style of team wiki"
[Spiegel's Branch]: https://scrapbox.io/spiegel-branch/ "Spiegel's Branch - Scrapbox"
[Slack]: https://slack.com/ "Slack: Be less busy"
