+++
date = "2015-10-03T00:24:53+09:00"
update = "2015-10-22T08:35:00+09:00"
description = "さて予定通り，正式オープンしました。"
draft = false
tags = ["site", "policy", "creative-commons", "license", "hugo", "github"]
title = "text.Baldanders.info 正式オープン"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

さて[予定]({{< ref "/remark/2015/todo-replace-blog.md" >}})通り，正式オープンしました。

## かえってきた「しっぽのさきっちょ」

「[しっぽのさきっちょ](https://baldanders.info/spiegel/log/)」は，むかし運営していた日記コンテンツの名前です。
「しっぽ」というのはもちろん {{< quote lang="en" >}}long tail{{< /quote >}} のことで，つまり「しっぽのさきっちょ」とは Web における「辺境」を意味します。
辺境へようこそ。

[以前も書いた](https://baldanders.info/spiegel/log/200604.html#d27_t2)けど，「しっぽ」の部分で谷山浩子ちっくな感受性を，「さきっちょ」の部分でそこはかとないエロスを感じていただければ幸いです。

私もいい歳ですし，辺境に帰って，これからは慎ましく生きていく所存です。

## このサイトについて

このサイトの運営方針は以下を参照して下さい。

- [text.Baldanders.info について]({{< relref "/site-policy.md" >}})

基本的には[本家サイト](https://baldanders.info/policy.shtml)と同じです。

### ホスティングについて

このサイトは [GitHub Pages](https://pages.github.com/) でホスティングしています。
アクセス履歴等は取れないので，現在も将来もこのサイトが流行ってるのか過疎ってるのか分かりません（多分）。
[Google Analytics](https://www.google.com/analytics/) 等はやってて虚しくなってきたので本家サイトでもとっくの昔に止めました。
こっちでもやりません。

ただし Twitter とか Facebook とか Amazon.co.jp などとは連携しているので，そういったサービスに Tracking されている可能性はあります。
あとホスティング・サービスを行っている [GitHub](https://github.com/) がプライバシーに関して evil ではないと断言できませんので，その辺も悪しからず。
心配な方は Firefox の[新しい private browsing 機能](http://www.mozilla.jp/blog/entry/10504/)を使って下さい。

### ライセンスについて

個々の記事については，原則として， [CC License](https://baldanders.info/spiegel/archive/cc-license/) の [by-sa] を設定することにしました。
ここは本家との違い。

[前にも紹介](https://baldanders.info/spiegel/log2/000796.shtml)しましたが， [Creative Commons] が提供しているライセンスツールのうち， Free Culture License に分類されるもの[^a] が過半数を超え，更にその中では [by-sa] が一番多いようです。
これにはいろんな理由があるんでしょうけど，私は {{< quote lang="en" >}}copyleft{{< /quote >}} がプログラム・コード以外にも浸透しつつあると好意的に解釈しています。

[^a]: [cc0], [by], [by-sa]。

だから，そろそろ自分のコンテンツも copyleft にしてもいいんじゃないかなぁ，ということで。
ただし，古い記事を再掲載する場合などは今まで通り [by] でいきます。

念の為に書いておきますが，ライセンスの許諾範囲内であれば事前事後の連絡は無用です。
また「リンク」や「引用」といったそもそも許可する必要の無いことについてもうだうだ言わなくていいです[^b]。
それ以外のことで問題点や不明な点があればフィードバックをいただけると助かります。

[^b]: 「シェアさせていただきました」とか意味不明の連絡も要りません。リンクも引用もシェアもリブログも勝手にどうぞ。

## リポジトリの公開

前述したように，このサイトは [GitHub Pages](https://pages.github.com/) でホスティングしているため，サイトのリポジトリは公開されています。

- [spiegel-im-spiegel/spiegel-im-spiegel.github.io](https://github.com/spiegel-im-spiegel/spiegel-im-spiegel.github.io)

### 作業環境の公開

このサイトは [Hugo] で生成しています。
[Hugo] を使った作業環境についてもリポジトリを公開しています。

- [spiegel-im-spiegel/github-pages-env](https://github.com/spiegel-im-spiegel/github-pages-env)

またこのサイト用に theme を作成して，こちらも公開しています。

- [spiegel-im-spiegel/hugo-theme-text](https://github.com/spiegel-im-spiegel/hugo-theme-text)

Theme のほうは [cc0] で公開しているため自由に利用していただいて構いません。
なお， [Hugo] については Section を立てて[連載中](/hugo)です。


[Creative Commons]: https://creativecommons.org/ "Creative Commons"
[cc0]: https://creativecommons.org/publicdomain/zero/1.0/deed.ja "Creative Commons — CC0 1.0 Universal"
[by]: https://creativecommons.org/licenses/by/4.0/deed.ja "Creative Commons — Attribution 4.0 International — CC BY 4.0"
[by-sa]: https://creativecommons.org/licenses/by-sa/4.0/deed.ja "Creative Commons — Attribution-ShareAlike 4.0 International — CC BY-SA 4.0"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
