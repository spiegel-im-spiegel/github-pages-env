+++
title = "ぼちぼち Mastodon への移住を進めようかと"
date =  "2022-12-26T13:20:38+09:00"
description = "テレビ並みにつまらない Twitter に固執する必要なくね？"
image = "/images/attention/kitten.jpg"
tags = [ "mastodon", "activitypub", "disease" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Mastodon アカウントを復活させた]({{< ref "/remark/2022/11/the-return-of-mastodon.md" >}} "Mastodon の復活")ときとは言い草がエラい違うな，という自覚はあるです。
きっかけは大きく2つ。

## 入院生活と Twitter

ひとつは先日の[入院騒ぎ]({{< relref "./heart-attack.md" >}} "ハライタだと思った？ 残念！ 心筋梗塞でした")。

いや，入院中はホンマに暇なのよ。
身体がしんどいときは考える余裕もないけど，そういう状況が過ぎると「飯を食う」と「寝る」以外に殆どイベントがない入院生活は苦痛ですらある。
心臓リハビリとか暇つぶしとして嬉々としてやってたくらいだし。

一応，病室にはテレビがあって金を払えば見れるんだけど，金払ってまでテレビを見たいとは思わんぢゃん。
NHK 受信料すら払いたくないのに（同居者がテレビっ子なので払ってるけど）。
そうなると，必然的にネットに没入するしかないんだけど，パソコンを持ち込むわけにはいかないから，ネットを覗き見る窓はスマホしかないわけ。
しかも入れてるアプリも制限してるので，常時見てるのは Twitter と Mastodon くらい。

五十路過ぎたオッサンが暇を持て余して日がな一日スマホで Twitter TL を眺めている姿を想像してくれ。
そしてつくづく思った。

**今の Twitter はテレビ並みにつまらない！**
{ .center }

このままではテレビに向かってツッコミをいれる独居老人みたいになってしまうではないか！ しかも実際に TL にツッコめる。
ヤバい匂いしかしない（笑）

## 「離脱の自由」

もうひとつのきっかけは Cory Doctorow による以下の記事：

- [Pluralistic: Freedom of reach IS freedom of speech (10 Dec 2022) – Pluralistic: Daily links from Cory Doctorow](https://pluralistic.net/2022/12/10/e2e/)
  - [“リーチの自由”は“言論の自由”である | p2ptk[.]org](https://p2ptk.org/freedom-of-speech/4211)
- [Pluralistic: Better failure for social media (19 Dec 2022) – Pluralistic: Daily links from Cory Doctorow](https://pluralistic.net/2022/12/19/better-failure/)
  - [我々が「離脱の自由」を必要とする理由――あるいはソーシャルメディアの失敗をマシにする方法 | p2ptk[.]org](https://p2ptk.org/freedom-of-speech/4214)

日本語訳もキレッキレで（4人部屋の病室なので一応は心の中で）爆笑しつつも脳を揺さぶられるような衝撃があった。
特に

{{< fig-quote type="markdown" title="“リーチの自由”は“言論の自由”である" link="https://p2ptk.org/freedom-of-speech/4211" >}}
コンテンツモデレーションは、情報セキュリティの中で唯一、隠蔽によるセキュリティ（security through obscurity）が有効だと考えられている領域なのである。

https://doctorow.medium.com/como-is-infosec-307f87004563
{{< /fig-quote >}}

は，マジで目からウロコだった。

もちろん Mastodon が最適解というわけではない。
Cory Doctorow は

{{< fig-quote type="markdown" title="“リーチの自由”は“言論の自由”である" link="https://p2ptk.org/freedom-of-speech/4211" >}}
オンラインの言論の自由をめぐる議論はあまりに見当違いで、度し難いほどに愚しい。

- アルゴリズムの改善にばかり注目し、見たいと求めたものをフィード上に表示させられるかどうかはまったく気にしない
- 未承諾メッセージが配信されないことばかりに注目して、承諾メッセージが読者に届くかどうかは気にかけない
- アルゴリズムの透明性ばかりに注目し、アルゴリズムの学習データを生成する行動追跡をオプトアウトできるかどうかを見逃す
- 社会的・仕事上・個人的なつながりを失うことなくプラットフォームを離脱できるかどうかには興味を持たず、プラットフォームがユーザを十分に取り締まっているかばかりを気にする
{{< /fig-quote >}}

と批判し，その上で Mastodon のメリットについて

{{< fig-quote type="markdown" title="我々が「離脱の自由」を必要とする理由――あるいはソーシャルメディアの失敗をマシにする方法" link="https://p2ptk.org/freedom-of-speech/4214" >}}
一方、Mastodonは他のどの巨大ソーシャルメディアも真剣に試みなかった２つのことを正しく行っている。

I. Mastodonで誰かをフォローすると、その人が投稿したものをすべて見ることができる。

II. Mastodonのサーバを離脱しても、自分のフォロワーとフォローしているアカウントの双方を持ち出すことができる。
{{< /fig-quote >}}

と述べる。
そして

{{< fig-quote type="markdown" title="我々が「離脱の自由」を必要とする理由――あるいはソーシャルメディアの失敗をマシにする方法" link="https://p2ptk.org/freedom-of-speech/4214" >}}
イーロン・マスクは邪悪な天才などではない。彼は運良く強大な権力を手にしたが、説明責任をほとんど果たすことのできないただの凡人である。Mastodonの運営者にはマスクのような傾向を持つ人物がいて、そこにユーザが放たれることになる。だが、そこには明確な違いがある。ユーザは2つのリンクをクリックして、別のインスタンスに移行できるということだ。そんじゃーねーーーー！
{{< /fig-quote >}}

と言い放つ（この記述が一番ウケた）。
所属している Mastodon サーバの管理者が某マスク氏のように[暗黒面に堕ちた](https://yamdas.hatenablog.com/entry/20221226/elon-musk-henry-ford-extremism "イーロン・マスクはヘンリー・フォードの轍を踏み、過激思想にいたる暗黒面に堕ちつつある？ - YAMDAS現更新履歴")としても，私達には「離脱の自由」があるじゃないか！

ここまで読んで，はたと気がつく。

**テレビ並みにつまらない Twitter に固執する必要なくね？**
{ .center }

## そして Mastodon へ

大勢のオーディエンスを人質に取られている有名人ならスイッチング・コストは高くつくだろうが，私のようなネットの辺境にいる人間は今の Twitter ソーシャルグラフを全部チャイして Mastodon に完全移行したとしても大して支障はない。
実際には「[ちょうぜつエンジニアめもりーちゃん](https://twitter.com/search?q=%23%E3%81%A1%E3%82%87%E3%81%86%E3%81%9C%E3%81%A4%E3%82%A8%E3%83%B3%E3%82%B8%E3%83%8B%E3%82%A2%E3%82%81%E3%82%82%E3%82%8A%E3%83%BC%E3%81%A1%E3%82%83%E3%82%93)」とか楽しみにしている Twitter 連載もあるし，誤配されてくるツイートの中にも「おおっ！」と思うものもあるので（少なくとも当面は）全く見なくなるということはないだろうが，なにせ Twitter TL は（プロモーションツイートも含め）ノイズが多すぎる。
どっち側に軸足を置くかと言われれば，今後は Mastodon になっていくと思う。

というわけで，様子見の期間は終わったと思うので，来年からは Mastodon/Fediverse に軸足を移すべく色々とやっていきたいと思う。

## ブックマーク

- [Keyoxide](https://keyoxide.org/) : 来年はこれで遊びたい
- [コリイ・ドクトロウの新刊はビッグテックや巨大メディアの権力へのクリエイターの対抗を呼びかける「チョークポイント資本主義」 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20220829/chokepoint-capitalism)
- [Mastodonインスタンス運営者のための法律入門 | p2ptk[.]org](https://p2ptk.org/freedom-of-speech/4220)
- [友だち、家族、顧客、コミュニティを失わずにFacebookを離脱する方法 | p2ptk[.]org](https://p2ptk.org/monopoly/antitrust/4226)
- [Facebookをどれほど嫌いになってもFacebookをやめられないのはなぜか | p2ptk[.]org](https://p2ptk.org/monopoly/antitrust/4231)
