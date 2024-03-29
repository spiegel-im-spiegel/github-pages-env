+++
title = "VTuber に満たない Bluesky"
date =  "2023-07-23T14:56:09+09:00"
description = "70万ユーザ登録おめでとうございます。"
image = "/images/attention/kitten.jpg"
tags = [ "communication", "engineering", "mastodon", "bluesky", "activitypub", "cloud" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いやぁ，夏っスねぇ。
九州はまだ梅雨が明けてないのに，本州は明けてしまったらしい。
相変わらずテキトーだな，気象庁，いや地方気象台か？

昨日も自転車で遊びに出かけたのだが，日中が暑すぎて図書館から出れなくなったり。
しょうがないので昼飯食いに路線バスで移動したですよ。

{{< fig-img src="./53063471038_a2982980e4_e.jpg" title="夏雲 | Flickr" link="https://www.flickr.com/photos/spiegel/53063471038/">}}

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

Threads のリリースですっかり影が薄くなった Bluesky は昨日時点でユーザ総数が34万人を超えたらしい。
同じ日， VTuber の周防パトラの[「70万人いくまでギターを弾き続ける耐久！」ライブ][耐久ライブ]が行われていた。

{{< fig-youtube id="xdDw9YGApLE" title="70万人いくまでギターを弾き続ける耐久！ ロック＆メタル！... - YouTube" >}}

念のために言っておくが， Bluesky は今（2023-07 時点で）なおクローズドベータ状態で，サインアップするためには wait list に登録するか誰かから招待コードを貰うしかない。
でも，単純に言って Bluesky のユーザ規模の倍以上のオーディエンスがいるわけですよ，彼女には。
もちろん世の中にはもっと上の桁のオーディエンスを抱える有名人もいるわけで。

先日，またもや Twitter がやらかしてかなりのユーザが Mastodon や Bluesky 等に流れ込んだわけだが，自前でどうとでも調達できる（クラウドという名の）計算資源を持ってる「ビッグテック」と違って「その他大勢」のサービスでは，とつぜん難民が押し寄せたからといってリアルタイムにスケールアップできるわけじゃない。
実際，新規登録を制限してるはずの Bluesky でさえ一時的にサインアップを止めざるを得なかった。
おそらくサーバ資源の増強をしてたんだろう。

Mastodon サーバのひとつで私がメインで利用している fedibird.com のユーザ総数は現在33K人を超えたあたりで，アクティブユーザに限れば10K人ほどらしい[^fb1]。
そんで，ランニングコストは145K円/月ほどだそうな。
この規模なら，単純計算で各アクティブユーザが300円/年ほど寄付すれば（少なくともランニングコスト分は）賄えるとのこと。

[^fb1]: fedibird.com は 2023-02 から新規登録を制限している。サインアップするには既存ユーザからの招待が必要。なので fedibird.com ユーザの増加ペースはゆっくりしている。

{{< fig-quote class="nobox center" >}}
<iframe src="https://fedibird.com/@noellabo/110758504812079195/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe><script src="https://fedibird.com/embed.js" async="async"></script>
{{< /fig-quote >}}

この中に人件費は含まれてないだろうし，累積赤字がけっこうあるらしいので「お金なんてナンボあってもいいですからね」って感じだろうけど。

実は昨年末に Hostdon で（何かに使えるかなと思って）ホスティングサービスを契約していたのだが，何もしないまま先月解約した。
ホスティングサービスに500円/月払うくらいなら fedibird.com に寄付するほうが全然安上がりだろうという判断。
不特定多数相手の汎用サーバだとコンプライアンスが云々とかあるので，特殊用途（TRUTH Social とかw）ならホスティングのほうがいいんだろうけど（あとはセルフブランディングで独自ドメインを使いたいとか），結局私には用がなかった（笑）

Mastodon のサーバで最大規模のユーザを抱えているのが mastodon.social で，2023-07時点で1.5M人に満たないくらい。
ちなみに mastodon.social 以外に1M人を超えるユーザを有しているサーバ（Threads のような連合しないサーバは除く）は今のところ存在しない。

もし Mastodon 連合が全体で1億人を超えるユーザを夢見るなら，ものすごく単純に考えても mastodon.social 規模のサーバが100は必要になる。
実際には ActivityPub で「連合」する [Mastodon サーバ](https://instances.social/list/advanced "Mastodon instances")の総数は18Kほどだが[全体のユーザ総数](https://mastodon.social/@mastodonusercount "Mastodon Users (@mastodonusercount@mastodon.social) - Mastodon")は13M人を超えたあたりのようだ[^u1]。

[^u1]: Facebook のユーザ数が29.9億人で世界最大。続いて YouTube は20億人， Instagram が10億人， TikTok が同じく10億人という感じ[らしい](https://growthseed.jp/experts/sns/number-of-users/ "【2023年6月最新】SNSの利用者数とユーザー属性や特徴まとめ")。ちなみに Twitter は3.3億人だそうな。トラブルで数万ユーザが逃げたところで大したことないし，喉元すぎれば熱さを忘れて戻ってくるユーザも多かろう（笑）

こうしてみると Threads が単体で初日で10M人ユーザを獲得したのは計算資源の調達からして並大抵ではないし，これが将来 ActivitiyPub 連携で Fediverse に加わるかもしれない考えるとぞんぞがさばる（出雲弁）。

クラウド資源あるいはサービスの利用はすぐに始められるので，個人利用やスタートアップでは使い勝手がいいが，スケールについて真剣に考え始めると足枷になることも多い。
クラウドで無制限にスケールできるのはクラウドの所有者（社）だけだろう。
あるいは札束で明かりとりをする成金か（笑） 喩えるならクラウドは砂漠の水売りみたいなもんで，ゼロ年代に言われてた「情報ダム」のような公共イメージとは程遠い。

...てなことを[耐久ライブ]を見ながら考えてた。
そうそう，70万ユーザ登録おめでとうございます。

結局のところ Twitter が本当に沈没しても代わりになるようなサービスは存在せず，某タイタニック号のごとく全てを道連れにするしかないのだ。
これは他の単一プラットフォームの大規模 SNS でも同じ。
たとえば私は広島時代の友人の近況を見るためだけに Facebook を使ってるが，仮に Facebook がサービスをシャットダウンしても代わりになるものがなく，旧友との関係も「ハイそれまでョ」になりかねない。
これはそういうものだと割り切るしかないだろう。

ただ，自分自身がネットから切り離されないよう，特定の企業・サービスの思惑に巻き込まれることなく，依存をできるだけ少なくするようにしないと。

## ブックマーク

- [我々が「離脱の自由」を必要とする理由――あるいはソーシャルメディアの失敗をマシにする方法 | p2ptk[.]org](https://p2ptk.org/freedom-of-speech/4214)
- [メタクソ化するTiktok：プラットフォームが生まれ、成長し、支配し、滅びるまで | p2ptk[.]org](https://p2ptk.org/monopoly/4366)
- [クラウドネイティブから見たクラウドの小史｜塚本 牧生](https://note.com/tsukamoto/n/n8056f3562d91)

- [Threads isn’t for news and politics, says Instagram’s boss - The Verge](https://www.theverge.com/2023/7/7/23787334/instagram-threads-news-politics-adam-mosseri-meta-facebook)
- [Twitter対抗の分散型SNS「Threads」が登場することでMastodonは変わるのか？をMastodonのCEOが解説 - GIGAZINE](https://gigazine.net/news/20230706-mastodon-ceo-comments-about-threads/)
- [Threads、JASRACと利用許諾契約を結んでいた　Twitterとの差別化点に - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/2307/06/news174.html)

- [Fediverse 関連のブックマーク]({{< ref "bookmarks/fediverse.md" >}})

[耐久ライブ]: https://www.youtube.com/watch?v=xdDw9YGApLE "【ギター/ESP FRX】70万人いくまでメタルをギターを弾き続ける耐久！ロック＆メタル！ フーファイターズ/メガデス/紅/Helloween/メガロバニア and more...【周防パトラ】 - YouTube"

## 参考

{{% review-paapi "B011LC4D58" %}} <!-- クレイジー・キャッツ・デラックス ハイそれまでョ -->
