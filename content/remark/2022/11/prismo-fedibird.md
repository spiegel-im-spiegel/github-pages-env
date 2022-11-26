+++
title = "Prismo はじめました，他"
date =  "2022-11-26T18:50:25+09:00"
description = "Prismo / fedibird でソーシャル・ブックマークを再開"
image = "/images/attention/kitten.jpg"
tags = [ "bookmark", "mastodon", "activitypub", "flickr" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## Fedibird はじめました

特に深い考えがあったわけではないのだが [Fedibird] に Mastodon アカウントを作った。
こちらは正式に公表することにする。

- [`@spiegel@fedibird.com`](https://fedibird.com/@spiegel)

どぞ，よろしく。

国内サービスはいいねぇ。
私のニックネームはドイツの某雑誌と同じなせいか，海外のサービスで `spiegel` はまず取れないのよ（笑）

mstdn.jp にあるアカウントは予備系として残すことにした。
Mastodon には「引っ越し」機能があってアカウントを別のサーバに移行することができるのだが，機能としては微妙で

- 引っ越し先にあらかじめアカウントエイリアスを作成してから引っ越し作業を行う
- 元のアカウントの削除は（自動では）できない。投稿（toots）できなくなるだけらしい
- 投稿は移動できない（バックアップは可能）
- フォロワーは自動で引き継げるがフォローは（自動では）移動できない
  - 引っ越し元のフォロー・リストをエクスポートして引っ越し先にインポートする
  - ブロック，ミュート，ブックマーク 等も同様

てな具合になっている。
これなら無理に移行しなくても引っ越し先に普通にアカウントを作って，引っ越し元には「引っ越しました。引っ越し先でもよろしく」とでも書いておけば済む話。
私の場合，5年間放置プレイだったせいでフォローしてもらってはいても inactive な人もいるので，無理に引っ越し先にフォロワーを引き継がなくてもいいかなと思ったり。

引っ越し先の [Fedibird] は管理者の [`@noellabo@fedibird.com`](https://fedibird.com/@noellabo) さんがなかなかアクティブな方のようで [Fedibird] 自体が[本家 Mastodon の魔改造](https://ch.dlsite.com/matome/227051 "【マストドン・Misskeyなど】SNS連合『Fediverse』の遊び方 - DLチャンネル みんなで作る二次元情報サイト！")などと評されている。
個人的には「汎用目的のサーバにローカル・タイムラインは要らんよね」と思ってたので [Fedibird] の設計は賛同できる。
まぁ，それ以外にもかなり弄ってあるみたいだが。

実際に

{{< fig-quote type="markdown" title="Fedibirdをフォークとして整備します - Open Collective" link="https://opencollective.com/fedibird-project/updates/fedibirdwofkutoshiteshimasu" >}}
しかしながら、様々な新規機能を実装し、それらによってFedibirdの存在感が増すにつれ、Mastodon本家にマージされる見込みのない様々な違いが拡大を続けていて、小さな差分では管理できなくなってきています。部分的にFedibirdの機能を取り入れるなら、もうFedibirdのコードベースに乗り換えた方が早いという状況です。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Fedibirdをフォークとして整備します - Open Collective" link="https://opencollective.com/fedibird-project/updates/fedibirdwofkutoshiteshimasu" >}}
この状況を鑑みて、Fedibirdをこれまでの私的なものから、広く利用可能なものへと改めるため、Mastodonのフォークの一つとして整備し、役割を果たすべき時が来たと判断しました。
{{< /fig-quote >}}

などと書かれていて，これはもう “[Fedibird]” というプロダクトと考えるほうがいいかもしれない（笑）

- [Fedibird · GitHub](https://github.com/fedibird)
  - [fedibird/mastodon: Mastodonの機能開発への貢献と、様々な理由で標準搭載されていない機能を共有するためのリポジトリです。](https://github.com/fedibird/mastodon)
- [fedibird - scrapbox](https://scrapbox.io/fedibird/)

[Fedibird] では Mastodon の他にも [ActivityPub] 対応のサービスが[用意](https://fedibird.com/about/more)されている。
2022-11 時点ではこんな感じらしい。

- [FediMovie](https://fedimovie.com/) : PeerTube
- [FediSnap](https://fedisnap.com/) : Pixelfed
- [misskey.cloud](https://misskey.cloud/) : Misskey
- [Element](https://element.fedibird.com/) : Matrix (Synapse + Element) によるチャット・通話サービス
- [Prismo / fedibird](https://prismo.fedibird.com/) : Prismo

いずれも独立したサービスでアカウントも個別に取得する必要がある（つまり [Fedibird] の Mastodon ユーザでなくとも利用できる）。

よくもまぁこんなにサービス展開できるな。
やはり[のえるさんは石油王](https://note.com/ncr/n/n134a1594a6c6 "SNS「マストドン」に5年くらい居るコーヒー豆屋がユーザー目線でマストドンの初期おすすめムーブを書いてみる何か｜Nelson Coffee Roaster｜note")なのか？

## Prismo はじめました

というわけで，とりあえず [Prismo] を使ってみようかと。

[Prismo] の説明には “Federated link aggregation powered by ActivityPub” とある。
元々は [Reddit](https://www.reddit.com/ "Reddit - Dive into anything") のカウンタとして作られた連合型プロダクトらしい。
日本だと（連合型という点を除けば）はてブが一番近いだろうか（笑） もちろん自前でサーバを立てることも可能。

まずは [Prismo / fedibird](https://prismo.fedibird.com/) でアカウントを作った。

- [`@spiegel@prismo.fedibird.com`](https://prismo.fedibird.com/@spiegel)

ブックマークを追加するには `[+ストーリーを追加]` ボタンを押す。

{{< fig-img src="./prismo-add-story.png" title="Prismo / fedibird" link="./prismo-add-story.png" width="840" >}}

すると入力画面が開く。

{{< fig-img src="./prismo-add-story-2.png" title="Prismo / fedibird" link="./prismo-add-story-2.png" width="826" >}}

右側のちょっと怖い注意書きにご注意を（笑） 右下にブックマークレットが用意されているので，これをブラウザに設置すれば表示しているページの URL が入った状態でこの画面が表示される。
“Name” にはタイトルを， “Tag list” にはタグ（`#` は付けない）をセットする。
必要に応じて説明（Description）も追加できる。
しかも markdown 記法が使える。

これを登録すると，こんな風にカード形式でブックマークのリストが表示される。

{{< fig-img src="./prismo-home.png" title="Prismo / fedibird" link="./prismo-home.png" width="841" >}}

各カードの下部にある「コメント」とか「投稿」とか書かれているあたりのリンクをクリックすると

{{< fig-img src="./prismo-comment.png" title="Prismo / fedibird" link="./prismo-comment.png" width="838" >}}

という感じに詳細が表示され，コメントを追加することもできる。

Mastodon でアカウントをフォローしていれば，投稿に対するコメントが反映される筈なんだけど... 上手くいかないな。
これはあとで検証しよう。

実は私はかなりのブックマーク魔で，昔は del.icio.us とか使ってたんだけど，今は Instapaper で溜め込んで定期的に自ブログの [bookmarks セクション]({{< rlnk "/bookmarks/" >}}) 吐き出すだけの作業（こうしておけば後で grep で探せるから）。
何処にも繋がらない（笑）

流石に今のペースで [Prismo / fedibird](https://prismo.fedibird.com/) に追加していくのは無理だと思うけど，少しずつソーシャル・ブックマークの比重を上げていければな，と思っている。

## 「おひとりさま」とか無理だから！

最初は「おひとりさま」サーバも視野に入れていたんだよ。
でも，回線の細い自宅に公開サーバを設置するとかありえないし，庶民にも手が届きそうな [Hostdon](https://hostdon.jp/ "Hostdon - Mastodonのホスティングサービス") は新規サーバの受付を停止してるし（2022-11 時点），他はかかるコストの桁が違うから！ 無理！！

まぁ，でも，よく考えたら[結城浩](https://social.hyuki.net/@hyuki "結城浩@social.hyuki.net (@hyuki@social.hyuki.net) - 結城浩のマストドン")さんみたいにセルフ・ブランディングするとか企業・組織ユーザとかなら「おひとりさま」にも意義があるだろうけど，そうでなければ「なし」だよな。
それをするくらいなら mstdn.jp や [Fedibird] みたいな汎用サーバでブロックやミュートを駆使しつつ穏やかなタイムライン運用をしていったほうが平和だろう。

あとはどっか面白そうな活動をしているサーバを見つけたらそっちに移るか，かな。

## フォロー？ 購読？

多分 [Fedibird] の（標準 Mastodon にはない）独自機能だと思うけど，任意のユーザに対して通常の「フォロー」以外に「購読」を指定することができる。

フォローと購読で何が違うのかと一瞬思ったが「購読」は公開された投稿しか読めないので，フォロワー限定の投稿は見れないということらしい。
[`@eff@mastodon.social`](https://mastodon.social/@eff "Electronic Frontier Foundation (@eff@mastodon.social) - Mastodon") などの組織アカウントや前節の [Prismo] アカウントなどは「購読」にしたほうがいいのだろうか。

## 【おまけ】 Flickr の写真を Mastodon TL に（いい感じに）貼り付けたい

Flickr 写真ページの URL，たとえば

```
https://www.flickr.com/photos/spiegel/52396396454/
```

に対して，パスの後ろに `in/dateposted-public/` をくっつけて

```
https://www.flickr.com/photos/spiegel/52396396454/in/dateposted-public/
```

としたものを Mastodon に投稿すると

{{< fig-quote class="nobox center" >}}
<iframe src="https://fedibird.com/@spiegel/109407237432891831/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe><script src="https://fedibird.com/embed.js" async="async"></script>
{{< /fig-quote >}}

などといい感じに表示してくれる。
ありがとう中の人！


[Fedibird]: https://fedibird.com/ "fedibird.com - Fedibird"
[ActivityPub]: https://www.w3.org/TR/activitypub/
[Prismo]: https://gitlab.com/prismosuite "prismosuite · GitLab"
<!-- eof -->
