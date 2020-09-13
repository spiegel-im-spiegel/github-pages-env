+++
title = "GitHub プロファイルを（ちょっとだけ）カッコよくしてみる"
date =  "2020-09-12T21:17:02+09:00"
description = "自前でポートフォリオを構成できる。"
image = "/images/attention/kitten.jpg"
tags = [ "github", "profile" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

何となく他所様の [GitHub] プロファイルを眺めてたら既定の表示と違うものがチラホラと見受けられる。

どうやら，自分のユーザ名（私の場合は `spiegel-im-spiegel`）のリポジトリを作って中の `README.md` をプロファイルページの Overview タブに表示できるらしい。

たとえばこんな感じ。

{{< fig-img src="./github-profile-readme.png" title="GitHub Profile Overview" link="https://github.com/spiegel-im-spiegel" width="896" >}}

Markdown や HTML のほか，拡張子を変えれば AsciiDoc とかも使えるみたいなので，割と自由に記述できる。
要するに，自前でポートフォリオを構成できるわけだ。

取り敢えず，それなりの画面を作りたいなら以下のサービスを利用するのがお勧め。

- [GitHub Profile Readme Generator | GitHub Profile Readme Generator](https://rahuldkjain.github.io/gh-profile-readme-generator/)
- [rahuldkjain/github-profile-readme-generator: Generate github profile README easily with latest add-ons like visitors count, github stats, etc using minimal UI.](https://github.com/rahuldkjain/github-profile-readme-generator)

項目を適当に埋めていって，アイコンやアドオンを選択して *[Generate README]* ボタンを押せば雛形となる Markdown コード（中身はほぼ HTML だけど`w`）を出力してくれるので，そのまま使うもよし，アレンジして使うもよしである。

## ブックマーク

- [How to create Github Profile-README | by Pratik Kumar | Jul, 2020 | Towards Data Science](https://towardsdatascience.com/explore-new-github-readme-feature-7d5cc21bf02f?gi=eb8dd4afb703)
- [Shields.io: Quality metadata badges for open source projects](https://shields.io/)
- [alexandresanlim/Badges4-README.md-Profile: Improve your README.md profile with these amazing badges.](https://github.com/alexandresanlim/Badges4-README.md-Profile)
- [anuraghazra/github-readme-stats: Dynamically generated stats for your github readmes](https://github.com/anuraghazra/github-readme-stats)
- [Blog Post Workflow · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/blog-post-workflow)
    - [gautamkrishnar/blog-post-workflow: Show your latest blog posts from any sources or StackOverflow activity or Youtube Videos on your GitHub profile/project readme automatically using the RSS feed](https://github.com/gautamkrishnar/blog-post-workflow)
- [GitHub Readme Stats を利用してGitHubプロフィールをカッコよくする - Qiita](https://qiita.com/zizi4n5/items/f8076cb25bbf64a9bc1c)

[GitHub]: https://github.com/
