+++
date = "2015-10-24T01:13:56+09:00"
update = "2016-03-15T00:55:46+09:00"
description = "GitHub-flow を捨てた ATOM Editor / 暗号プロトコルのセキュリティ評価 / VoLTE の脆弱性 / 「おかげさまで半世紀も生きちゃったぜ」記念"
draft = false
tags = ["atom", "editor", "tools", "git", "github", "security", "cryptography", "authentication", "risk", "vulnerability", "ntp", "wireless"]
title = "今日の戯れ言： GitHub-flow を捨てた ATOM Editor"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1. [GitHub-flow を捨てた ATOM Editor]({{< relref "#atom" >}})
1. [暗号プロトコルのセキュリティ評価]({{< relref "#cipher" >}})
1. [VoLTE の脆弱性]({{< relref "#volte" >}})
1. [「おかげさまで半世紀も生きちゃったぜ」記念]({{< relref "#birth" >}})

## GitHub-flow を捨てた ATOM Editor{#atom}

- [Introducing the Atom Beta Channel](http://blog.atom.io/2015/10/21/introducing-the-atom-beta-channel.html)
- [オープンソースのテキストエディター「Atom」にベータチャンネルが登場 - 窓の杜](http://www.forest.impress.co.jp/docs/news/20151022_726976.html)

{{< fig-quote title="オープンソースのテキストエディター「Atom」にベータチャンネルが登場" link="http://www.forest.impress.co.jp/docs/news/20151022_726976.html" >}}
<q>これまで「Atom」は“master”ブランチ（ソースコードの大元）を直接リリースする形態がとられていた。この方法はシンプルで、機能の追加や不具合の修正をそのままユーザーの元へ届けることができたため、開発当初はうまくいっていたという。しかし、機能が増えるにつれて修正済みの不具合が再発する“リグレッション”が多くなり、ワークフローに混乱が生じることがたびたび発生するようになったようだ。<br>
そこで、GitHubは“master”ブランチを直接リリースする方法をやめ、「Google Chrome」などでお馴染みの“リリースチャンネル”を設ける方針をとることにした。「Atom」の場合、開発を行う“master”ブランチから切り離されたベータ版と、ベータ版でテストされた新機能や不具合修正を盛り込んだ正式版（stable）の2つのリリースチャンネルが設けられる。</q>
{{< /fig-quote >}}

まぁ要するに [GitHub] は， [ATOM] に関してはお家芸の GitHub-flow を捨てて Git-flow に切り替えたということらしい。

- [GitHub Flow – Scott Chacon](http://scottchacon.com/2011/08/31/github-flow.html)
- [GitHub Flow (Japanese translation)](https://gist.github.com/Gab-km/3705015)
- [A successful Git branching model » nvie.com](http://nvie.com/posts/a-successful-git-branching-model/)
- [git-flow cheatsheet](http://danielkummer.github.io/git-flow-cheatsheet/) （[日本語](http://danielkummer.github.io/git-flow-cheatsheet/index.ja_JP.html)）
- [git flowとgithub flowざっくりまとめ | KentaKomai Blog](http://komaken.me/blog/2013/09/09/git-flow%E3%81%A8github-flow%E3%81%96%E3%81%A3%E3%81%8F%E3%82%8A%E3%81%BE%E3%81%A8%E3%82%81/)

どちらがいいのかについては何とも言えないが，複数バージョンを同時に管理していくのであれば Git-flow にせざるを得ないだろう。
個人的には 1.0 系でそれほど不満はないので，敢えてベータ版を使う必要はないかな。

[ATOM]: https://atom.io/ "Atom"
[GitHub]: https://github.com/ "GitHub"

## 暗号プロトコルのセキュリティ評価{#cipher}

- [プレスリリース | 暗号プロトコルのセキュリティ評価結果をリスト化・公開 | NICT-情報通信研究機構](http://www.nict.go.jp/press/2015/10/20-2.html)
- [Cryptographic Protocol Verification Portal](http://crypto-protocol.nict.go.jp/)
- [ASCII.jp：58個の暗号プロトコルをセキュリティ評価！NICTがリスト公開](http://ascii.jp/elem/000/001/068/1068218/) : 10月23日現在ダウン中[^a]。

[^a]: [イルカ漁への抗議で Anonymous が絶賛攻撃中](http://japanese.engadget.com/2015/10/22/ascii-jp-ddos-anonymous-it-ascii-jp/)らしい。 ASCII.jp 関係ないじゃん。迷惑な話。ちなみにイルカは美味い。

各種の認証および鍵交換プロトコルの評価一覧。
分かりやすくまとめられてるし技術文書へのリンクもあってとても参考になる。
なんで [NICT] なのかは分からないけど。
[NICT] の活動って chaotic でイマイチよく分からないんだよなぁ。

そうそう。
[NICT] で思い出したけど，また ntpd の脆弱性が見つかったみたい。

- [Cisco Identifies Multiple Vulnerabilities in Network Time Protocol daemon (ntpd)](http://blogs.cisco.com/security/talos/2015-10-ntpd-vulnerabilities)
- [時刻同期のNTPに複数の脆弱性--HTTPS接続のバイパスなどを誘発 - ZDNet Japan](http://japan.zdnet.com/article/35072380/)

やれやれ。

[NICT]: http://www.nict.go.jp/ "NICT-情報通信研究機構"

## VoLTE の脆弱性{#volte}

脆弱性といえばもうひとつ。
いや4つ。

- [Breaking and Fixing VoLTE: Exploiting Hidden Data Channels and Mis-implementations](http://dl.acm.org/citation.cfm?id=2813718)
- [Vulnerability Note VU#943167 - Voice over LTE implementations contain multiple vulnerabilities](https://www.kb.cert.org/vuls/id/943167)
- [JVNVU#93463833: Voice over LTE (VoLTE) の実装に複数の脆弱性](https://jvn.jp/vu/JVNVU93463833/index.html)
- [VoLTEの脆弱性、携帯3社は検証の上で「問題なし」 | マイナビニュース](http://news.mynavi.jp/news/2015/10/23/085/)

ちうことで日本の3大キャリアでは問題なさそうだ。

{{< fig-quote title="ニュースコメント[VoLTEの実装で複数の脆弱性と報告] | 無線にゃん" link="http://wnyan.jp/4204" >}}
<q>GoogleとしてはIMSベアラに直接触れないようにOSを修繕する必要があるので「対策します」ってことになったのでしょうが、基本的にはシステム全体としてはとっくに対策されてるんですよね。あまり知られていませんがVoLTEも実はモデムチップに依存したシステムで、主要チップ屋の実装として端末への直接攻撃も防がれるようになっているみたいですし。まあ、そういう情報は普通はあまり外に漏れないので（漏れたらそれはそれでセキュリティリスクになるし）、理屈上はこうやったらハックできちゃうぜ、って発表されることは悪いことではないと思うのですが、少なくとも、日本やその他情報通信先進国でまともな装置を使っているような国では、このリスクは問題ないと思っていただいて大丈夫だと思いますよ。</q>
{{< /fig-quote >}}

慌てない慌てない。

## 「おかげさまで半世紀も生きちゃったぜ」記念{#birth}

今月で五十路に突入します。
いやぁ，半世紀も生きちゃったよ。

半世紀も生きてるのに何かを悟ったとか全然ないし，財や名誉を得たとかも全然ないし，何やってるんだろうね，私は。
50代になったらもっとこうナイスミドル（笑）になってるとかありそうなもんだけど，言動が（歳の割に）幼いのか，いまだに生物年齢相当に見られないのは困ったもんである。

お祝いを下さった友人方々，本当に有難うございます。
みなさんのおかげで明日からも生きていけると思います。
