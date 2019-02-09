+++
title = "Flickr から写真を引き揚げました"
date = "2019-01-20T19:16:46+09:00"
update = "2019-02-09T11:17:42+09:00"
description = "ただの「写真置き場」なら自分のところに置いておけばいいぢゃん，と気がついた。"
image = "/images/attention/kitten.jpg"
tags = [ "site", "web", "photography", "image", "flickr", "creative-commons" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

私にとって Web 2.0 そして SaaS を象徴するのは blog でも SNS でもなくて [Flickr] なのですよ。
更に言えば [Flickr] は Creative Commons の「[ブートストラップ問題](https://mag.osdn.jp/03/09/29/0955208 "クリエイティヴ・コモンズに関する悲観的な見解 | OSDN Magazine")」を解消した功労者のひとつでもある。

当時は私も Pro 会員だった。
それも今は昔の話。

昨年11月に [Flickr] から無料会員に対し，アップできる写真を最大1,000枚に制限する旨のアナウンスがあった（CC Licenses 下で既に公開しているものについては執行猶予あり）。

- [Why we’re changing Flickr  free accounts | Flickr Blog](https://blog.flickr.net/en/2018/11/01/changing-flickr-free-accounts-1000-photos/)
- [Flickrの新しいビジネスモデルでCreative Commonsの作品は消されるのか  |  TechCrunch Japan](https://jp.techcrunch.com/2018/11/03/2018-11-02-flickrs-new-business-model-could-see-works-deleted-from-creative-commons/)
- [Flickr says it won’t delete Creative Commons photos – TechCrunch](https://techcrunch.com/2018/11/07/flickr-says-it-wont-delete-creative-commons-photos/)

正直に言って Pro 会員に戻ろうか悩んだ。

確かに今の私は（色々あって）ビンボーだけど，ゼロ年代からお世話になっているサービスだし，必要であるならお金を払うことに吝かではない。
でも現時点で [Flickr] にお金を払って何が得られるか考えたときに，写真の枚数以外に何もないことに気がついた。

現在の殆どの Web サービスは，人と情報をサービス内部に囲い込むように構成されている。
でも写真共有サービスはサービスを横断して，息をするように共有できなければただの「写真置き場」になってしまう。
ただの「写真置き場」なら自分のところに置いておけばいいぢゃんか（今更`w`）。

ちうわけで [Flickr] の写真を引き揚げて自サイトに写真置き場を作りました。

- [photo.Baldanders.info](https://photo.baldanders.info/)

写真の画像データとメタデータ（JSON 形式）については昨年のうちにバックアップ済みだったのだが，メタデータからページをおこすための変換ツール作成（なにせ2,600枚以上あるのだ）に着手するまでに私事で色々あり過ぎてすっかり遅くなってしまった。

変換ツールについても最初はなるべく汎用的に作ってオープンソースで公開しようと思ったが個人レベルのチューニングを上手く外だしにできなかった。
幸いなことに GitHub が無料ユーザにもプライベート・リポジトリを開放してくれたので，そっち側にリポジトリを作って作業した。

- [New year, new GitHub: Announcing unlimited free private repos and unified Enterprise offering | The GitHub Blog](https://blog.github.com/2019-01-07-new-year-new-github/)
- [朗報、GitHub無料ユーザーも無制限にプライベートリポジトリを使えるようになる  |  TechCrunch Japan](https://jp.techcrunch.com/2019/01/08/2019-01-07-github-free-users-now-get-unlimited-private-repositories/)
- [GitHub、無料ユーザーもプライベートリポジトリを無制限に作成可能に － Publickey](https://www.publickey1.jp/blog/19/github_4.html)
- [“GitHub”の非公開リポジトリ、無償プランでも無制限に ～新しい料金プランが発表 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1161195.html)

なので変換ツールは非公開です。
ゴメンペコン。

これでようやく肩の荷が下りたよ。
今後は呑気に本家サイトの再構築に取り掛かるとしよう。
今回の作業で再構築に必要なものとか大体わかったので，今度はもう少し上手くやる。

## ブックマーク

- [Flickrのバックアップを急げ、今日から1000枚+の写真削除開始  |  TechCrunch Japan](https://jp.techcrunch.com/2019/02/06/2019-02-05-flickr-deleting-photos-how-to-backup/)
- [Flickrは写真の削除を3月12日までに延期  |  TechCrunch Japan](https://jp.techcrunch.com/2019/02/07/2019-02-06-flickr-will-hold-off-on-deleting-photos-until-march-12th/)

[Flickr]: https://www.flickr.com/
