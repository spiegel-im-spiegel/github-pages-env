+++
title = "誰が「パスワードは複雑なものにしろ」と言ったのか"
date =  "2024-12-02T23:06:54+09:00"
description = "ルールによってユーザの行動がどう変わるかみたいな話を1979年の論文で考慮しろってのが無理筋である。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "authentication", "password", "risk", "management", "nist" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いつものように yomoyomo さんの記事を起点に小咄を。

- [長年の誤ったパスワードポリシーが推奨された原因はあの偉人の論文だった？ - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20241202/bad-password-policies)

この記事の元ネタが更にあって，ボクらの [Bruce Schneier 先生の記事](https://www.schneier.com/blog/archives/2024/11/good-essay-on-the-history-of-bad-password-policies.html "Good Essay on the History of Bad Password Policies - Schneier on Security")で紹介されている以下の記事。

- [How some of the world's most brilliant computer scientists got password policies so wrong | Mildly-Aggrieved (not mad!) Scientist](https://stuartschechter.org/posts/password-history/)

これは古のパスワード生成ルール「パスワード文字列は英数字と記号の三種盛りにしろ」とか「パスワードは定期的に変更しませう」とかの元ネタが何かって話らしい。
記事によると Robert Morris さんと Ken Thompson さんが1979年11月に公開した “{{< pdf-file link="https://dl.acm.org/doi/pdf/10.1145/359168.359172" title="Password Security: A Case History">}}” という論文が大元なんだとか。
Ken Thompson さんは Go 言語開発者のひとりなんだね。

{{< fig-quote type="markdown" title="How some of the world's most brilliant computer scientists got password policies so wrong" link="https://stuartschechter.org/posts/password-history/" lang="en" >}}
First, was Morris and Thompson’s confidence that their solution, a password policy, would fix the underlying problem of weak passwords. They incorrectly assumed that if they prevented the specific categories of weakness that they had noted, that the result would be something strong.
{{< /fig-quote >}}

この記事では言及されてないが，更に悪いことに，この考え方で NIST SP 800-63 のほうも実装されてしまったらしい。
このせいで悪名高きパスワード生成ルールが世に蔓延ることになったわけだ。

以下は改訂された NIST SP 800-63-3 が正式リリースされた2017年の ZDNET Japan の記事。

{{< fig-quote type="markdown" title="あの「面倒なパスワード作成ルール」、作った人も後悔していた - ZDNET Japan" link="https://japan.zdnet.com/article/35105725/" >}}
 　大文字、小文字、数字、記号を使った覚えにくいパスワードを作らされ（しかもそれを定期的に変更することを義務づけられ）て、どうしてこんなルールがあるのかと憤ったことがある人は多いだろう。それはおそらく、どこかの開発者が、米国立標準技術研究所（NIST）が2003年に作成した文書に従ったためだ。

　この8ページの「NIST Special Publication 800-63別表A」は、NISTの元管理職であるBill Burr氏によって作成された。同氏はすでに引退しており、今では72歳になっている。

　Burr氏は、The Wall Street Journalの取材に対して、「今では、わたしが決めたことの大部分について後悔している」と語った。
{{< /fig-quote >}}

“[How some of the world's most brilliant computer scientists got password policies so wrong](https://stuartschechter.org/posts/password-history/ "How some of the world's most brilliant computer scientists got password policies so wrong | Mildly-Aggrieved (not mad!) Scientist")” では更に

{{< fig-quote type="markdown" title="How some of the world's most brilliant computer scientists got password policies so wrong" link="https://stuartschechter.org/posts/password-history/" lang="en" >}}
That mistake was their recommendations on how passwords should be stored. They recommended that systems should not store passwords, but instead assign each user a random “hash” function used to compute a number (the hash) from that users’ password.

[...]

Storing numeric hashes instead of the passwords can protect users whose passwords are hard to guess, but it also prevents scientists from examining those passwords to determine if there might be categories of common (weak) passwords that users should be discouraged, or prevented, from choosing. While Morris and Thompson did not invent password hashing 3, they implemented it into Unix, strongly recommended it, and their paper would be the one most cited to support the necessity of password hashing.
{{< /fig-quote >}}

などと書いている。
ユーザによるパスワード生成の実態をサービスプロバイダ側が把握できなかったせいで「間違い」が放置されたままだったというのだ。
少なくとも論文が登場した1979年には公開鍵暗号は発明されてたんだからハッシュ関数を使うのではなく公開鍵暗号を使えばよかったとか言ってるようだ。

まぁ「研究者」としてならこれはこれで正しいのかも知れないが，エンジニアとしてはパスワード情報の秘匿に公開鍵暗号を使うのが「善」なのか首をひねるところがある。
そもそも記事で挙げている RSA は20世紀当時は特許でガチガチに固められていて，しかも米国外では（軍事的な理由で）使用不可なアルゴリズムだったので，実質的には無理な話だったと思う。

いや，パスワードを（ハッシュ化ではなく）暗号化するのはいいけど，その鍵をどうやって管理するのかって話ですよ。
厳格な管理が要求される Bitcoin 交換所でもたまに漏洩事件が起きたりするんだよ。

もしもの話として，パスワードを公開鍵暗号の公開鍵で暗号化するとして，ひとつの鍵ペアで全部のパスワードを暗号化するわけないし（そんなことして万が一秘密鍵が解読されたり漏れりしたら全ユーザのパスワードが晒されてしまう），そうなるとユーザごとに鍵ペアが必要ってことになるだろう（無数の鍵ペアを使うにしても秘密鍵をひとつところに置いていれば同じことだがw）[^p1]。
つか，そんなことするくらいならパスワード認証なんかやめて公開鍵暗号を使って認証すればいいぢゃん。
ssh みたいに。

[^p1]: さらに別のもしもの話として，生のパスワード情報を暗号化してどっかに保持って（実際の認証ではなく）統計情報にのみ使うとして，その結果をユーザの行動にどう反映させるのか，という問題もある。メディアで定期的にヤバいパスワードランキングが公表されるが，あれはそういったパスワード情報が既に攻撃手段として用いられている（だろう）と分かってるからできることだ。拙文「[「パスワードのベストプラクティス」が変わる]({{< ref "/remark/2017/10/changes-in-password-best-practices.md" >}})」でも書いてるが「**パスワードを覚えるなんて脳みその無駄使い**」である。人間は複雑というか予測不能な文字列をいくつも思いつけるようにはできていない。パスワード認証を使うならパスワード管理ツールで生成も管理も任せてしまったほうが安全ってことだと思う。

パスワード生成のルール化はどちらかというとユーザ体験（UX; User eXperience）の問題と言える。
パスワードの複雑化や定期更新を強制するのは，そのほうがシステム管理側が楽だからだ。
そして面倒な部分をユーザ側に「転嫁」しているのだ。
セキュリティ・マネジメントの観点で「転嫁」戦略は悪いことではないが，不特定のユーザにそれを強いてもユーザはただ迂回するだけである。
これは「悪い UX」の典型と言える。

そういえば『[はじめて学ぶ ビデオゲームの心理学]』に UX の起源みたいな話があって

{{< fig-quote type="markdown" title="『はじめて学ぶ ビデオゲームの心理学』 p.34" link="https://www.amazon.co.jp/dp/4571210450?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
ドナルド・ノーマン（大きな影響力をもつ『誰のためのデザイン？ &mdash;&mdash; 認知科学者のデザイン原論』（Norman, 1990）の著者）は「ユーザー体験（UX：user experience）」という単語を提唱し、製品やそのエコシステム（マーケティング、ウェブサイト、顧客サービスなど）にユーザーが関与するときの体験全般を考えるという方針を示しました。そのため、企業やグローバル戦略も UX に含まれます。
{{< /fig-quote >}}

という感じに UX 自体が1990年代以降の考え方らしい。
ルールによってユーザの行動がどう変わるかみたいな話を1979年の論文で考慮しろってのが無理筋である。

まぁ，歴史に「たられば」はないっちうことやね。

ところで，最初に挙げた yomoyomo さんの記事を読んで気がついたのだが NIST SP 800-63-4 のドラフト版って2022年に公開されてるんだね。
しかもパブコメの募集は先々月の10月で締め切られているらしい。

- [NIST SP 800-63 Digital Identity Guidelines](https://pages.nist.gov/800-63-4/)

まぁでもここから正式版が出るまでが長いならなぁ NIST は。
気長に待ちましょうか。

[はじめて学ぶ ビデオゲームの心理学]: https://www.amazon.co.jp/dp/B0C9Z7KGRN?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "はじめて学ぶ ビデオゲームの心理学 脳のはたらきとユーザー体験（UX） | セリア ホデント, 山根 信二（監修）, 山根 信二, 成田 啓行 | 工学 | Kindleストア | Amazon"

## ブックマーク

- [「パスワードのベストプラクティス」が変わる]({{< ref "/remark/2017/10/changes-in-password-best-practices.md" >}})
- [Authenticator と AAL]({{< ref "/remark/2020/09/authenticator-and-aal.md" >}})
- [『はじめて学ぶ ビデオゲームの心理学』は読んどけ！]({{< ref "/remark/2023/04/the-psychology-of-video-games.md" >}})

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "B0C9Z7KGRN" %}} <!-- はじめて学ぶ ビデオゲームの心理学 Kindle 版 -->
