+++
date = "2017-06-21T12:59:52+09:00"
description = "よーやくですよ！"
draft = false
tags = ["tools", "editor", "atom", "git", "github"]
title = "ATOM 1.18 stable リリースで公式に Git 機能に対応"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

先週 [ATOM] 1.18 の stable 版がリリースされたが，公式に [git] 機能が組み込まれたようだ。

- [Atom 1.18 | Atom Blog](http://blog.atom.io/2017/06/13/atom-1-18.html)
    - [Integrating Git in Atom | GitHub Engineering](https://githubengineering.com/integrating-git-in-atom/)
    - [Git and GitHub Integration comes to Atom | Atom Blog](http://blog.atom.io/2017/05/16/git-and-github-integration-comes-to-atom.html)
- [「Git」と“GitHub”を統合した「Atom」v1.18が正式版に ～GitHub製の無償エディター - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1065638.html)

もともと core package として [git-diff] は組み込まれていたのだが，実際のリポジトリ操作には [git-plus] 等の外部パッケージを使わざるを得なかった。
これが [github] として core package に組み込まれ，リポジトリ操作が GUI で提供されることになったわけだ。
よーやくですよ！

{{< fig-img src="http://blog.atom.io/img/posts/github-package-git.png" title="Atom 1.18" link="http://blog.atom.io/2017/06/13/atom-1-18.html" width="1560" >}}

Stage や commit/amend や fetch/pull/push といった基本操作はもちろん，hunk[^h] を選択して stage すること（`git add -p` 相当）も GUI で可能なようだ。
よしよし。
ただし stash や cherry-pick といった細かい操作はできなさそうっぽい？ 

[^h]: 変更箇所のひとかたまりを [git] では hunk と呼ぶ。 Hunk の概念は cherry-pick 時にも出てくるので覚えておくとお得。

ただ，良くも悪くも GUI なので「マウスやトラックパッドなんて飾りです。偉い人には...」な方々には従来通り [git-plus] のほうがお勧めである。
余談だが，私は command palette を F1 キーに割り当てているが（秀丸を使っていた頃の名残）， [git-plus] 専用のメニューは shift-F1 キーに割り当てている。
こんな感じ。

```cson
'.platform-win32':
  'shift-f1': 'git-plus:menu'
```

もうひとつの機能である [GitHub](https://github.com/) との連携（今のところ pull request の表示のみ？）であるが，これを使うためには access token を取得して [ATOM] に登録する必要がある。

- [GitHub「Personal access tokens」の設定方法 - Qiita](http://qiita.com/kz800/items/497ec70bff3e555dacd0)

では，たのしくお仕事しましょう！

## ブックマーク

- [ATOM Editor に関するメモ]({{< ref "/remark/2015/atom-editor.md" >}})
- [あまり使わないけど，たまに使おうとすると忘れてる Git コマンド集]({{< ref "/remark/2015/git-commands.md" >}})

[ATOM]: https://atom.io/ "Atom"
[git]: https://git-scm.com/ "Git"
[git-diff]: https://atom.io/packages/git-diff
[git-plus]: https://atom.io/packages/git-plus
[github]: https://atom.io/packages/github
