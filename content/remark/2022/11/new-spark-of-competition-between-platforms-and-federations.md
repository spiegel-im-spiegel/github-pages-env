+++
title = "「プラットフォームとフェデレーションとの競争」"
date =  "2022-11-25T18:56:07+09:00"
description = "もし Flickr の ActivityPub 対応が実現するなら SNS 活動の軸足を Fediverse へ移してもいいかもしれない。"
image = "/images/attention/kitten.jpg"
tags = [ "code", "internet", "generativity", "activitypub", "mastodon", "twitter" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

念のため予防線を張っておくと，私は日本で Mastodon が台頭し始めた2017年春に「[GW 過ぎたらみんな忘れてるに100カノッサ]({{< ref "/remark/2017/04/mastodon.md" >}})」と書いて，見事に賭けに負けた人間である。
まぁ，私自身はつい最近まで本当に忘れていたのだが（笑）

というわけで， Mastodon マンセー！ とか言うつもりはないのであしからず。

## ActivityPub 連合

[Mastodon アカウントを復活]({{< relref "./the-return-of-mastodon.md" >}})させてから，あちこち覗いて回っている。
特に以下の記事は2017年よりあとの（特に日本語圏での）状況を知るのに役に立った。
ありがたや {{% emoji "ペコン" %}}

- [【マストドン・Misskeyなど】SNS連合『Fediverse』の遊び方 - DLチャンネル みんなで作る二次元情報サイト！](https://ch.dlsite.com/matome/227051)

この記事にも出てくる [ActivityPub] の W3C 勧告（Recommendation）は 2018-01-23 にリリースされている。

{{< fig-img-quote class="lightmode" src="./ActivityPub-tutorial-image.png" title="ActivityPub-tutorial-image" link="https://en.wikipedia.org/wiki/ActivityPub#/media/File:ActivityPub-tutorial-image.png" lang="en" width="761" >}}

これに先立ち，2017年秋ごろには Mastodon や PeerTube でも実装されていたようだ。
更に2018年以降 Misskey, Pleroma, Pixelfed といったあたりが対応をはじめた。
更に更に，昨今の「Twitter お家騒動」を受けて Tumblr が [ActivityPub] に対応すると表明し，我らが Flickr も検討に入ったらしい。

- [Tumblr to add support for ActivityPub, the social protocol powering Mastodon and other apps • TechCrunch](https://techcrunch.com/2022/11/21/tumblr-to-add-support-for-activitypub-the-social-protocol-powering-mastodon-and-other-apps/)

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">Was just discussing ActivityPub earlier this week internally. Should we add to <a href="https://twitter.com/Flickr?ref_src=twsrc%5Etfw">@Flickr</a>? (No promises, just thinking [err, Tweeting] out loud to gauge interest.) It might be right up our alley, though… <a href="https://t.co/MsrPhlwB3Q">https://t.co/MsrPhlwB3Q</a></p>&mdash; Don MacAskill (@DonMacAskill) <a href="https://twitter.com/DonMacAskill/status/1594945727255699457?ref_src=twsrc%5Etfw">November 22, 2022</a></blockquote>
{{< /fig-gen >}}

正直に言って「Flickr 対応するかも」というのは心が揺れた。
もしこれが実現するなら SNS 活動の軸足を Fediverse へ移してもいいかもしれない。

## EFF による Mastodon/ActivityPub/Fediverse 解説

昨今の状況を受けて EFF (Electronic Frontier Foundation) も解説記事を出した。

- [Leaving Twitter's Walled Garden | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2022/11/leaving-twitters-walled-garden)
  - [Twitterの“囲い込み”から離脱するということ――あるいは Fediverse、Federation、Mastodonとは何か？ | p2ptk[.]org](https://p2ptk.org/freedom-of-speech/4167)
- [The Fediverse Could Be Awesome (If We Don’t Screw It Up) | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2022/11/fediverse-could-be-awesome-if-we-dont-screw-it)
  - [フェディバースは（“私たち”次第で）素晴らしいものになる | p2ptk[.]org](https://p2ptk.org/freedom-of-speech/4165)
- [Is Mastodon Private and Secure? Let’s Take a Look | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2022/11/mastodon-private-and-secure-lets-take-look)

Fediverse とは federation と universe を組み合わせた造語だそうな。
たとえば Mastodon はそれ自体が連合（federation）型の分散システムだが，そうしたサービス同士が更に緩くフラットに結びついた状態を指すらしい。

{{< fig-img-quote src="./Fediverse_small_information.png" title="How-the-Fediverse-connects - Fediverse - Wikipedia" link="https://en.wikipedia.org/wiki/Fediverse#/media/File:How-the-Fediverse-connects.png" lang="en" width="600" >}}

これら連合型システムを結びつける技術要素のひとつが [ActivityPub] というわけだ。
ぶっちゃけ，どこぞの FinTech 山師が叫ぶ Web3 より，こっちのほうがよほど次期 Web ぽいよな（笑）

EFF の解説記事では従来の Facebook や Twitter といったサービスを「プラットフォーム」あるいは「中央集権」型に分類し Fediverse と対置しているのが面白い。
単なる「代替サービス」ではないということだ。

{{< fig-quote type="markdown" title="Twitterの“囲い込み”から離脱するということ――あるいは Fediverse、Federation、Mastodonとは何か？ | p2ptk.org" link="https://p2ptk.org/freedom-of-speech/4167" >}}
連合には中央機関が存在しないため、Twitter独自の認証バッジなどの複数の機能はまったく意味をなさない。「認証（verified）」に近いものとしては、自分のプロフィールに特定のハイパーリンクを掲載し、外部のウェブページやリソースを管理していることをインスタンスに証明するというものがある。

fediverseは分散型であるため、投稿を管理したり、アカウントを削除する単一の権限はなく、すべてはユーザとサーバに委ねられている。Mastodonのユーザは一般に、投稿にコンテンツ警告マークをつけている。これは本当にセンシティブなコンテンツ（たとえば、戦争のニュースに関するコンテンツ警告）のためだけではなく、タイムライン上の投稿の足跡を最小化するためにも用いられている。また、ハッシュタグと合わせて、センシティブではない投稿の分類やキュレーションにも用いられている（たとえば、コンテンツ警告：「うちの猫 #pets」）。
{{< /fig-quote >}}

その上で既存プラットフォームに対して辛辣な評価を下す。

{{< fig-quote type="markdown" title="フェディバースは（“私たち”次第で）素晴らしいものになる" link="https://p2ptk.org/freedom-of-speech/4165" >}}
数週間前まで、ほとんどのソーシャルメディアユーザは“誰に支配されるか”しか[選びようがなかった](https://www.eff.org/deeplinks/2018/06/competition-civil-liberties-and-internet-giants)。現在、オンライン・ユーザが抱える多くの問題は、そうした集中によって生み出されている。

プライバシーを例に取ろう。既存のプラットフォームは、プラットフォーム側が提示する条件を呑むか、アカウントを削除するかというオール・オア・ナッシングの二者択一をユーザに迫っている。設定の奥深くに埋もれたプライバシー・ダッシュボードで多少の調整はできるようになったが、すべてのチェックボックスを外しても[大手サービス](https://www.eff.org/wp/behind-the-one-way-mirror)はあなたのデータの収集をやめようとはしない。主要プラットフォームに依存することによって、我々のプライバシー、セキュリティ、表現の自由の自律性が大幅に損なわれているのだ。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="フェディバースは（“私たち”次第で）素晴らしいものになる" link="https://p2ptk.org/freedom-of-speech/4165" >}}
つい最近まで、インターネットの終焉を想像することは、テック企業の終焉を想像するよりも簡単だった。無責任な巨大企業が支配するシステムのもとで暮らすことの問題は、避けられないように思えわれた。だが、こうした中央集権型プラットフォームは成長を停滞させ、Twitterは無様に崩壊しつつある。Twitterの崩壊、混乱は、これが最後ではないだろう。
{{< /fig-quote >}}

こうしたプラットフォーム型サービスに対する失望感（絶望感？）が今回の「Twitter お家騒動」より派生する諸々の出来事の背景にあるのかもしれない。
そして最後に EFF はこう断ずるのである。

{{< fig-quote type="markdown" title="フェディバースは（“私たち”次第で）素晴らしいものになる" link="https://p2ptk.org/freedom-of-speech/4165" >}}
善き独裁者を選んでも、独裁体制を修正することはできない
{{< /fig-quote >}}

一方で Fediverse に対しても

{{< fig-quote type="markdown" title="フェディバースは（“私たち”次第で）素晴らしいものになる" link="https://p2ptk.org/freedom-of-speech/4165" >}}
はっきりさせておこう。連合（federated ）によって、これまでの問題が魔法のようにたちどころに解決するわけではない。もし連合型ソーシャルメディアが既存の中央集権型ソーシャルメディアよりも優れているのだとしたら、それはそこに集う人たちが良いものを作りあげようと意識的に選択したからであり、**技術的決定論によるものではない**。オープンで分散化されたシステムは、より良いオンライン世界に向けた新たな選択肢を提供する。だが、その選択をするのは我々自身なのである。
{{< /fig-quote >}}

と釘をさしている（強調は私がやりました）。
その上で「連合システムの運用者と利用者に期待される選択肢」として以下を挙げている。

1. コンテンツ・モデレーションに関するサンタクララ原則を採用する
2. コミュニティ／ローカル・コントロール
3. コンテンツ・モデレーションのイノベーション
4. 無数のアプリケーション
5. リミックス可能性
6. 多様な資金援助モデル
7. グローバルなアクセシビリティ
8. 政府からの干渉への抵抗／ユーザの側に立つこと
9. 真の連合
10. 相互運用性と次の囲い込みの阻止
    - 反競争的行為の阻止
    - ポータビリティ
    - 委任可能性（delegability）

詳しくは[元記事を参照](https://p2ptk.org/freedom-of-speech/4165 "フェディバースは（“私たち”次第で）素晴らしいものになる | p2ptk.org")のこと。

最後の「反競争的行為の阻止」「ポータビリティ」「委任可能性」3つは “[Embrace, Extend, and Extinguish" (EEE)](https://en.wikipedia.org/wiki/Embrace,_extend,_and_extinguish)” のカウンタかな。
つか「取り込み、拡張し、抹殺する」とか物騒なフレーズがあるんだな（笑）

## 「プラットフォームとフェデレーションとの競争」

コンピュータおよびコンピュータ・システムの歴史は集中と分散の繰り返しである。
ネットも同じ。
プラットフォーム間の覇権争いに巻き込まれてウンザリしたユーザが緩い連合型システムに流れるのも自然なことのように思える。

しかし，これまで見たように Mastodon は Twitter の代替にはなりそうもない。
Twitter から Mastodon に逃げたユーザが，自分たちの「Twitter 文化」を振りかざして迷惑をまき散らしているという話もちょいちょい聞く[^mst1]。
リアルの「移民問題」と似たような話がネットでも出てくるというのは興味深い。

[^mst1]: 私もちゃんと振る舞えているか自信がない。昔の人は言いました。「3年ROMれ」と。

Fediverse が「素晴らしいもの」になるか否かが「“私たち”次第」ということは，言い方を変えれば Fediverse は属人性の強い壊れやすいシステムであるということだ。


EFF の解説記事はこう締めくくる。

{{< fig-quote type="markdown" title="Twitterの“囲い込み”から離脱するということ――あるいは Fediverse、Federation、Mastodonとは何か？ | p2ptk.org" link="https://p2ptk.org/freedom-of-speech/4167" >}}
既存のソーシャルメディアと、それに変わる連合型メディアの間には常にトレードオフが存在する。プラットフォームとフェデレーションとの競争という新たな火種は、新たなイノベーションとオンラインにおける我々の自律性の向上という2つの可能性を秘めている。
{{< /fig-quote >}}

ここで EFF が「競争（competition）」と言っていることには意味があると思う。

「文化とは文化の{{< ruby "hybrid" >}}混血児{{< /ruby >}}」である。
作られた monolithic な文化は簡単に腐る。
今回の騒動が新たな風を呼び込み generativity に満ちた社会に発展することを期待したい。

## ブックマーク

- [SNS「マストドン」に5年くらい居るコーヒー豆屋がユーザー目線でマストドンの初期おすすめムーブを書いてみる何か｜Nelson Coffee Roaster｜note](https://note.com/ncr/n/n134a1594a6c6)

[ActivityPub]: https://www.w3.org/TR/activitypub/

## 参考文献

{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
{{% review-paapi "4000280872" %}} <!-- イノベーション 悪意なき嘘 -->
{{% review-paapi "B099RTG3J7" %}} <!-- 著作権は文化を発展させるのか: 人権と文化コモンズ -->
