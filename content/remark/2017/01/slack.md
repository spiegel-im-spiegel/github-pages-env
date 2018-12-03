+++
description = "個人で Slack を使ってみる。"
draft = false
tags = ["slack", "tools", "feed"]
date = "2017-01-09T17:24:54+09:00"
update = "2018-01-13T14:11:18+09:00"
title = "いまさら聞けない Slack の使い方"

[author]
  github = "spiegel-im-spiegel"
  facebook = "spiegel.im.spiegel"
  tumblr = ""
  license = "by-sa"
  instagram = "spiegel_2007"
  avatar = "/images/avatar.jpg"
  name = "Spiegel"
  flattr = "spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  linkedin = "spiegelimspiegel"
  flickr = "spiegel"
+++

いや， [Slack] 自体のアカウントは持ってるのよ（2015年の夏には取ってたし）。
で， [aozorahack] （[join](https://aozoraslackin.herokuapp.com/)） と [creativecommons] （[join](https://slack-signup.creativecommons.org/)） には入ってるんだけど，それっきり放置プレイ状態。
まぁボッチが使うツールじゃないよねぇ，と思ってたら昨年末にこんな記事を見かけた。

- [個人 Slack のススメ - Qiita](http://qiita.com/saitotak/items/ac0eb7ddc0d8d83cbe91)

おおっなるほど！ そういう使い方もあるのか。
ということで私もやってみることにした。

## Slack とは

そのまえに [Slack] について簡単に。

[Slack] は基本的にはチャットサービスである。

チーム（team）およびチャネル（channel）毎にやりとりできる。
チーム単位で[無料プランと有料プランを選べる](https://slack.com/pricing)が，小規模なコミュニケーションなら無料プランで充分賄える。
有料プランは大規模コミュニティまたは企業向けといえるかもしれない。

ユーザは必ず何らかのチームに所属しなければならない。
新たにアカウントを作成する場合は，事前に何処かのチームから招待（invitation）を受けるか自分でチームを作成する
（ただし[招待の発行を自動化するツール](http://qiita.com/homata/items/d809e7489b4d14f0768b "手軽にチームのSlackに参加してもらう方法 - Qiita")はいくつか存在する）。

で， [Slack] 最大の特徴は他サービスと連携するための API およびアプリが用意されていることである。
アプリはボット（bot）と呼ばれることもある[^co]。

[^co]: これに関連して “ChatOps” という考え方がある。これはチャットサービスをベースにしたシステム運用あるいは業務改善といったニュアンスで使われるようだ。多くの業務タスクが SaaS で提供されるようになった結果，業務タスク間の効率的な連携が問題になり始めた。これに対する解として2015年辺りから ChatOps を導入するところが増えている。また ChatOps を前提にした業務システムを構築することで各タスクを疎結合にできるため柔軟な業務改善が期待できる。この ChatOps の手段として [Slack] を含む hackable なチャットサービスが注目されている。

## フィードを読み込んでチャネルに出力する

では早速試してみよう。

フィード（feed）を読み込むには [RSS integration](https://slack.com/apps/A0F81R7U7-rss) をインストールする。
あっ，その前にフィード情報を出力するチャネルを作成しておくとよい。

{{< fig-img src="https://c2.staticflickr.com/1/335/32179569785_c3f97eeaeb.jpg" title="Slack RSS integration 1"  link="https://www.flickr.com/photos/spiegel/32179569785/" >}}

続けて参照するフィードの URL と出力先のチャネルを設定する（チャネルはその時点で存在するものしか指定できないので注意）。

{{< fig-img src="https://c5.staticflickr.com/1/524/31804596620_fb500bec88.jpg" title="Slack RSS integration 2"  link="https://www.flickr.com/photos/spiegel/31804596620/" >}}

フィードはいくらでも登録できるようなので，よく見るサイトのフィードを片っ端から登録しておけばいい。
とりあえず私は以下のフィードを登録した。

- #astronomy
    - [星の情報.jp](https://news.local-group.jp/rss.rdf)
    - [月探査情報ステーション](http://moonstation.jp/feed)
    - [国立天文台 天文情報センター 暦計算室](http://eco.mtk.nao.ac.jp/koyomi/site/koyomi.rdf)
    - [国立天文台(NAOJ)](http://www.nao.ac.jp/atom.xml)
    - [SciencePortal](http://scienceportal.jst.go.jp/rss/news.rdf)
- #feed
    - [Big Sky](http://mattn.kaoriya.net/index.rss)
    - [POSTD](http://postd.cc/feed/)
    - [Publickey](http://www.publickey1.jp/atom.xml)
    - [aozorablog](http://www.aozora.gr.jp/aozorablog/?feed=atom)
    - [ぽじとろんの竹本泉観察記](http://rss.exblog.jp/rss/exblog/positron/index.xml)
    - [中東・イスラーム学の風姿花伝](http://ikeuchisatoshi.com/feed/)
    - [お知らせ | さくらインターネット](https://www.sakura.ad.jp/rss/info.rdf)
    - [Blog – Creative Commons](https://creativecommons.org/blog/feed/)
    - [「平和構築」を専門にする国際政治学者](http://shinodahideaki.blog.jp/atom.xml)
    - [The Go Programming Language Blog](https://blog.golang.org/feed.atom)
    - [Foresight](http://www.fsight.jp/list/feed/rss)
    - [Hugo](https://gohugo.io/index.xml "The world’s fastest framework for building websites | Hugo")
    - [MathJax](https://www.mathjax.org/feed.xml)
- #security
    - [IPAセキュリティセンター:重要なセキュリティ情報](https://www.ipa.go.jp/security/rss/alert.rdf)
    - [情報セキュリティ新着情報RSS](https://www.ipa.go.jp/security/rss/info.rdf)
    - [JPCERT/CC RSS Feed](https://www.jpcert.or.jp/rss/jpcert.rdf)
    - [JVNRSS Feed - Update Entry](http://jvn.jp/rss/jvn.rdf)
    - [US-CERT Alerts](http://www.us-cert.gov/ncas/alerts.xml)
    - [piyolog](http://d.hatena.ne.jp/Kango/rss)
    - [Schneier on Security](https://www.schneier.com/blog/atom.xml)
    - [wizSafe Security Signal -安心・安全への道標- IIJ](https://wizsafe.iij.ad.jp/feed/)
    - [エフセキュアブログ](http://blog.f-secure.jp/atom.xml)

メールや Facebook TL で取れる情報は除いたつもりだが，ちょっと多かったかな。
調整しながら運用してみよう。

## Slack でスターを付けた発言を Pocket に登録する

この機能を設定するには [Zapier] が必要なため，あらかじめサインアップしておく。
[Zapier] は [IFTTT] と同じように trigger と action を繋いで処理を行う API マッシュアップサービス。
無料版では最大5つまで設定できる（それ以上は[有料](https://slack.com/pricing)）。

準備ができたら以下のページから Zap を作成し設定を行う。
ちなみに [Pcoket] のアカウントはあらかじめ作成済みとする。

- [Save new starred Slack messages with links to Pocket](https://zapier.com/app/editor/template/1702)

これで「フィードを読み込んでチャネルに出力する」で出力された発言にスターを付ければ自動的に [Pcoket] に転送される。
たとえば [Slack] で以下のように発言が出力されるとすると

{{< fig-img src="https://c7.staticflickr.com/1/316/31361991614_8ec3890765.jpg" title="Feed to Slack"  link="https://www.flickr.com/photos/spiegel/31361991614/" >}}

[Zapier] 側はは以下のように `Massage Attachments From Url` および `Massage Attachments Title` をテンプレートにセットすればきれいに [Pcoket] に出力される。

{{< fig-img src="https://c7.staticflickr.com/1/680/31827692990_61360c3ac8.jpg" title="Zap in Zapier"  link="https://www.flickr.com/photos/spiegel/31827692990/" >}}

## その他

もし追加したアプリ等あればこの記事に追記していくことにする。

## ブックマーク

- [ワークスペースに新規メンバーを招待する – Slack](https://get.slack.help/hc/ja/articles/201330256-Invite-new-members-to-your-Slack-team)

- [We're on Slack! Join us! - Creative Commons](https://creativecommons.org/2016/10/18/slack-announcement/)
- [Qiitaの特定タグへの新着投稿をRSS受信する - Qiita](http://qiita.com/takecy/items/7e587e91aa2574926299)
- [手軽にチームのSlackに参加してもらう方法 - Qiita](http://qiita.com/homata/items/d809e7489b4d14f0768b)
- [ChatOpsで加速するOps効率化 | NTTデータ](http://www.nttdata.com/jp/ja/insights/trend_keyword/2016070701.html)
- [ChatOpsでなにができるのかまとめ - Qiita](http://qiita.com/m_mizutani/items/f7fa7b1d1c077b139f98)
- [SlackとGoogleが戦略的連携を発表―SlackとGoogleドライブが緊密化する | TechCrunch Japan](http://jp.techcrunch.com/2016/12/08/20161207slack-and-google-announce-partnership-focused-on-better-integrating-their-services/)

[Slack]: https://slack.com/ "Slack: Be less busy"
[Zapier]: https://zapier.com/ "The best apps. Better together. - Zapier"
[IFTTT]: https://ifttt.com/ "Learn how IFTTT works - IFTTT"
[Pcoket]: https://getpocket.com/
[aozorahack]: https://aozorahack.slack.com/
[creativecommons]: https://creativecommons.slack.com/
