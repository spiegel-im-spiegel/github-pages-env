+++
title = "私の Mastodon 遍歴"
date =  "2024-05-26T14:54:28+09:00"
description = "私の Mastodon 遍歴を纏めておくのは多少の意義はあるかな。あと愚痴とかも（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "mastodon", "activitypub", "twitter", "bluesky", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[fedibird.com] の管理者であるのえるさんが

{{< fig-quote type="markdown" title="MastodonやMisskeyは設置されたサーバごとに違うSNSだと思った方がいい、という話をしましたが、それだけに選ぶのは難しい..." link="https://fedibird.com/@noellabo/112493023994664254" >}}
MastodonやMisskeyは設置されたサーバごとに違うSNSだと思った方がいい、という話をしましたが、それだけに選ぶのは難しい。

fedibird.comから生きたサーバとして認識されているMastodon、Misskey、Pleroma系サーバの合計が約1万2千あります。このうち、500ユーザー以上登録されているサーバにしぼっても、698サーバあります。

とはいえ、いくつかの共通のソフトウェアが使われていて、サーバ同士が連合していて、どのサーバから参加してもお互いにつながることができる、というのもActivityPubの分散SNSの売り文句ではあります。

つまり、ざっくりした情報からとりあえずどこかに参加してみて、そこからいろいろ調べて確かめて、自分がどこからどうやってActivityPubネットワークに参加するか決めていく、という過程を経るしかないのかなと思います。

最初のざっくりした選択肢は、デフォルトのmastodon.socialや、joinmastodon.org掲載のmstdn.jpやmastodon-japan.net、あとは個人の紹介しているいくつかのサーバになるでしょうか。MisskeyにはMisskeyHubがありますが、基本的にmisskey.ioに直接集まるかな。
{{< /fig-quote >}}

と書いておられて，さらに続けて

{{< fig-quote type="markdown" title="そういうわけで、最初のサーバを探す人のことはあえて意識から外して..." link="https://fedibird.com/@noellabo/112493167200165085" >}}
そういうわけで、最初のサーバを探す人のことはあえて意識から外して、

もうこちらにある程度慣れた人に向けた情報が増えていく必要があるのではないかと思う今日この頃です。

そういう詳しい情報、最初のサーバを探してる人にも役立つんですけどね。
{{< /fig-quote >}}

と書かれている。
んじゃあ，私の Mastodon 遍歴を纏めておくのは多少の意義はあるかな。
ということで書いてみる。
あと愚痴とかも（笑） まぁ，言うほど大した内容じゃないが。

## Mastodon に転ぶきっかけ

最初に予防線を張っておくと，私は2017年春に日本で Mastodon が流行し始めた頃に「[GW 過ぎたらみんな忘れてるに100カノッサ]({{< ref "/remark/2017/04/mastodon.md" >}})」と書いて最終的に賭けに負けた人間である。
もっとも私自身は（今はなき [Qiitadon を試した]({{< ref "/remark/2017/11/qiitadon.md" >}} "Qiita って Mastodon やってたのか")こともあったが）2017年が終わる頃には Mastodon について綺麗サッパリ忘れ去っていた。

再び Mastodon を思い出したのは，言うまでもなく，旧 Twitter を買い取って{{% ruby "メタクソ" %}}私物{{% /ruby %}}化した某マスク氏の暴虐が始まった2022年後半。
具体的な[トリガー]({{< ref "/remark/2022/11/the-return-of-mastodon.md">}} "Mastodon の復活")としては，暗号化メッセージングアプリの Signal が Mastodon アカウントを公開したことだ。

{{< fig-gen lang="en" >}}
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">Hello, Mastodon - signalapp@mastodon.world</p>&mdash; Signal (@signalapp) <a href="https://twitter.com/signalapp/status/1593678164319997953?ref_src=twsrc%5Etfw">November 18, 2022</a></blockquote>
{{< /fig-gen >}}

もっとも，このときはまだ Mastodon に懐疑的で，放置していた mstdn.jp サーバのアカウントを流用して様子見の状態だった。
その後 [fedibird.com] へ拠点を移している。

本格的に Twitter から Mastodon への移住を考え始めたのは2022年が終わる頃。
2022年末は心筋梗塞で[入院騒ぎ]({{< ref "/remark/2022/12/heart-attack.md" >}} "ハライタだと思った？ 残念！ 心筋梗塞でした")をやらかしていて，このとき

{{< fig-quote type="markdown" title="ぼちぼち Mastodon への移住を進めようかと" link="/remark/2022/12/move-to-mastodon/" >}}
今の Twitter はテレビ並みにつまらない！
{{< /fig-quote >}}

と感じたのが移住を決意するポイントだった。
最終的に {{% emoji "X" %}} (旧 Twitter) アカウントは2023年9月に[休眠状態]({{< ref "/remark/2023/09/suspend-activity-on-twitter.md" >}} "𝕏 (旧 Twitter) の活動を休止します（期間未定）")としている。
実際には私の関心領域にかかる人が {{% emoji "X" %}} で活動されているので完全放置というわけでもないのだが。

mstdn.jp から [fedibird.com] に拠点を移した理由だが，標準機能をかなり拡張して細かくカスタマイズできるから。
misskey.io 等の絵文字スタンプもきれいに表示してくれるし引用ポストもできるし全文検索機能もある。
ただし（2024年5月時点で）バージョン 3.x をベースにしているようで[^f1]，ポストの変更などには対応してない。

[^f1]: [fedibird.com] では，セキュリティ脆弱性に対してバックポート・パッチを当てることで対応しているようだ。

## ActivityPub と Fediverse

話は変わるが Mastodon 分散プロトコルのベースになっている [ActivityPub] は Mastodon が発祥ではないらしい。
Mastodon に [ActivityPub] が実装され始めたのは [v1.6](https://github.com/mastodon/mastodon/releases/tag/v1.6.0 "Release v1.6.0 · mastodon/mastodon") からだそうなので2017年9月頃になる。
ちなみに [ActivityPub] が標準仕様として W3C 勧告になったのは2018年1月である。

[ActivityPub] の実装は本当に多岐に渡っている。

{{< fig-img-quote src="/bookmarks/fediverse/Fediverse_branches_1.2.png" title="File:Fediverse branches 1.2.png - Wikipedia" link="https://en.wikipedia.org/wiki/File:Fediverse_branches_1.2.png" lang="en" >}}

Web ベースの SNS 分散プロトコルとしては [ActivityPub] がおそらく最強だろう（最適解とまでは言わないが）。
WordPress の[拡張機能](https://wordpress.org/plugins/activitypub/ "ActivityPub – WordPress plugin | WordPress.org")として組み込むこともできるし Cloudflare は [Wildebeest] を使って Mastodon 互換サーバを構築するサービスを展開しているそうだ。

- [Welcome to Wildebeest: the Fediverse on Cloudflare](https://blog.cloudflare.com/welcome-to-wildebeest-the-fediverse-on-cloudflare)
- [Cloudflare wants to help you set up your own Mastodon-compatible server in ‘minutes’ - The Verge](https://www.theverge.com/2023/2/10/23593966/cloudflare-mastodon-server-wildebeest-instance-fediverse)
- [CloudflareがMastodonに対応したActivityPub実装、Wildebeestを作ってたので紹介します。一部有料です](https://zenn.dev/tkithrta/articles/9069279e1a3a1e)

最初に紹介したポストにもあるとおり Mastodon サーバだけでもかなりの数があるし，広く [ActivityPub] ベースの Fediverse という括りであれば更にバリエーションがある。
今のところ私は [ActivityPub]/Mastodon は {{% emoji "X" %}} の代替ではなく，全く別の生態系（ecosystem）の一部であると認識している。

## Bluesky はどうなん？

2023年に入り Nostr や AT-protocol/Bluesky といった分散プロトコルやそのサービスが台頭してくる。

個人の感想として Nostr は（使い捨て ID を使ったデータ交換としては面白いが）人間同士のコミュニケーション・ツールとしては筋が悪いと[評価]({{< ref "/remark/2023/02/plant-your-flag-mark-your-nostr-territory.md" >}} "Nostr に旗を立てろ！")していて，早々に見限った。

一方 Bluesky はプロトコルとプラットフォームの開発を同時進行するというなかなかに難しい運営をしているが，最近になってようやく分散化のロードマップが明確になってきて，実際に [bsky.social](https://bsky.social/) といったサーバも登場しているようだ。
ただ AT-protocol についてはマジで「なにもわからない」状態なので，そろそろ真面目に勉強すべきかなぁ。

私としては {{% emoji "X" %}} の代替サービスを探しているのなら Mastodon より Bluesky のほうがいいんじゃないかと思っている。
Bluesky のほうが優れていると言うつもりはないが Mastodon というか [ActivityPub] は仕組み上 Fediverse 全体を俯瞰することが難しい。
あくまで各ユーザ間のフォロー・フォロワー関係の中でデータのやり取りとその保持を行うので，特にユーザ数の少ない小規模サーバは「離島」状態になりやすい。
これでは {{% emoji "X" %}} で大勢のオーディエンスを抱えているユーザは安易に [ActivityPub]/Mastodon に移行し辛いんじゃないだろうか。

Bluesky が実際に分散化し「連合（federation）」を組むようになったときにこういった問題を上手く捌けるかどうかは分からないが期待感はある。

## mastodon.social ボット・アカウントを BAN された話

2023年に入って {{% emoji "X" %}} とIFTTT との連携が[有料になった](https://ifttt.com/explore/updates-to-twitter-2023 "Updates to Twitter access - IFTTT")ため，自サイトのコンテンツ情報を {{% emoji "X" %}} に自動でアップできなくなった（{{% emoji "X" %}} にも IFTTT にもお金を払う気はない）。
そこで Web 上のコンテンツ情報を自動で収集し Mastodon や Bluesky にポストするためのツールを[自作](https://github.com/goark/toolbox "goark/toolbox: A collection of miscellaneous commands")することにした[^t1]。
ちなみに IFTTT は完全撤退した。

[^t1]: [自作ツール](https://github.com/goark/toolbox "goark/toolbox: A collection of miscellaneous commands")は完全に自分用に特化していて，OSS で公開はしているけど，他人が使うことを想定してない。あしからず。汎用ツールとしては [mattn/go-mastodon](https://github.com/mattn/go-mastodon "mattn/go-mastodon: mastodon client for golang") や [mattn/bsky](https://github.com/mattn/bsky "mattn/bsky: A cli application for bluesky social") あたりを使って cron 等で運用するのがオススメである。

更にこれを使って自サイトのコンテンツ情報のみならず他サイトのコンテンツ情報を収集・ポストすることも検討し始めた。
要するに feed reader を Mastodon/Bluesky で代替する仕組みが欲しかったのだ。
そうすれば色んなサービスを渡り歩かなくても済むかな，と思って。

「他サイトのコンテンツ情報を収集・ポストする」アカウントは別に用意した。
Bluesky はまだ招待制だったが紹介状（invitation）が有り余ってたし Mastodon は折角だから [mastodon.social] にボットアカウントを作った[^m1]。

[^m1]: Mastodon でボットアカウントを作る際にはプロファイルでボットであることを示すフラグを立てる。言い換えると，そのフラグさえ立てておけばボットアカウントは許容されている。

しばらくはこれで上手く行ってたんだけど [mastodon.social] のボット・アカウントのほうがいきなり BAN されてしまった。
理由は示されず不明。
気がついたらデータのエクスポート以外何もできなくなってたのだ。

この経験から自身の活動を大手 Mastodon サーバで行うのはかなりのリスクを伴うことに気がついた。
そこで個人活動用に Mastodon サーバを立ち上げ，最終的に Mastodon での活動拠点を全てそこに移すことにした。

- [個人用 Mastodon サーバを立てた]({{< ref "/remark/2023/12/personal-mastodon-server.md" >}})
- [個人用 Mastodon サーバに活動拠点を移す]({{< ref "/remark/2024/03/moving-to-personal-mastodon-server.md" >}})

ただ個人サーバはホンマに「離島」生活みたいだよな。
フォローしたアカウントしか聞こえてこないから。
予備系として大きめのサーバにアカウントを確保しておいたほうがいいかも知れない。
私は [fedibird.com] のアカウントは今も消さずに予備系として残している。

## 封建時代のソーシャル・ネットワーク

今は本当に色んな Mastodon サーバがある。
たとえば [W3C](https://w3c.social/explore "w3c.social") や [Internet Archive](https://mastodon.archive.org/ "mastodon.archive.org") や [BBC](https://social.bbc "social.bbc") は自前でサーバを立ち上げてるし，FOSS 関連の [Fosstodon](https://fosstodon.org/) や [FLOSS.social](https://floss.social/about) といったサーバもある。
Mastodon ベースと言われる Threads は最近ようやく [ActivityPub] 連携を unlock した（ただしオプトインらしい）。

ただ，海外のサーバはルールとして英語での会話を要求しているところが多く，私のように母国語が日本語で英語不得手の人間は利用し辛い。
あとモデレーションの基準がどうしても厳しいしよく分からない。
たとえば前節で挙げた [mastodon.social] は言語の縛りはないようだが misskey.io をまるっと拒否している。
まぁ misskey.io 側は気にしちゃいないだろうが（笑） 単一プラットフォームのようにユーザが支配されていないからといって好き勝手絶頂に振る舞えるわけではないのだ。

前々節で「ユーザ数の少ない小規模サーバは「離島」状態になりやすい」と書いたが，ユーザ規模の大きいサーバは縛りが大きくて窮屈と言える。
まぁ規模が大きければ大きいほどユーザを「量」でしか評価できなくなるからね。
[ActivityPub]/Mastodon と言えども「[デジタル封建主義](https://yamdas.hatenablog.com/entry/20230828/neo-feudalism "時代はデジタル封建主義？ ジョエル・コトキン『新しい封建制がやってくる』が出るぞ - YAMDAS現更新履歴")」からは逃れられないのだ。
いや [ActivityPub]/Mastodon の場合は「デジタル氏族制度」とでも言うべきか？

EFF が言っていたが

{{< fig-quote type="markdown" title="フェディバースは（“私たち”次第で）素晴らしいものになる | p2ptk[.]org" link="https://p2ptk.org/freedom-of-speech/4165" >}}
善き独裁者を選んでも、独裁体制を修正することはできない
{{< /fig-quote >}}

のである。
かつてリベラルが「多文化主義」という名の墓穴を掘ったように Fediverse が「多様化」の墓穴にならないことを祈ろう。

## ブックマーク

- [Fediverse 関連のブックマーク]({{< ref "/bookmarks/fediverse.md" >}})

[fedibird.com]: https://fedibird.com/ "Fedibird"
[mastodon.social]: https://mastodon.social/
[ActivityPub]: https://www.w3.org/TR/activitypub/
[Wildebeest]: https://github.com/cloudflare/wildebeest "cloudflare/wildebeest: Wildebeest is an ActivityPub and Mastodon-compatible server"

## 参考図書

{{% review-paapi "B0CG5R6N2G" %}} <!-- 新しい封建制がやってくる ネオ封建制 -->
{{% review-paapi "4903127044" %}} <!-- 排除型社会 -->
