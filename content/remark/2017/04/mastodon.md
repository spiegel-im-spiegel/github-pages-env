+++
date = "2017-04-16T23:48:58+09:00"
update = "2017-05-02T12:36:16+09:00"
title = "GW 過ぎたらみんな忘れてるに100カノッサ"
tags = ["mastodon", "twitter", "communication"]
description = "んー。要するに今回の「お祭り」は，どこぞの院生が Mastodon のインスタンスを立ち上げたらうっかりユーザが殺到しちゃって，さくらや Azure の中の人なんかも巻き込んで今に至る。ということでおっけ？"
draft = false

[scripts]
  mathjax = false
  mermaidjs = false
+++

んー。
要するに今回の「お祭り」は，どこぞの院生が [Mastodon] のインスタンスを立ち上げたらうっかりユーザが殺到しちゃって，さくらや Azure の中の人なんかも巻き込んで今に至る。
ということでおっけ？

とりあえず [mstdn.jp] にアカウントを作ってみた。

- [@spiegel@mstdn.jp](https://mstdn.jp/@spiegel "Der spiegel im spiegel - mstdn.jp")

「世界一」とか言う割にまだ “spiegel" が取られてなかったことからしてそれほどは流行ってない感じ？
500文字なんて日本語なら小説が書けるレベルだよな。

分散型メッセージングってのは夢だよねぇ。
私は Twitter より Jabber/XMPP を連想したよ。
まぁ [Mastodon] は（Twitter と同じく）分類としてはマイクロブロギング・サービスになるんだろうけど。

そもそもインターネットが（現状はともかく理想としては）分散型のネットワークだし，電子メールだって Web だって分散型のサービスだ。
今世紀に入ってからなら Tor や Blockchain を含めてもいいかもしれない。
いわゆる「[技術的ゲートキーパ](https://baldanders.info/blog/000490/ "監視をコントロールする — Baldanders.info")」を迂回できるというのは大きい。

でも実際には分散型サービスって成功例があんまり無いんだよね。
みんな思いつくけど（そして作ろうと思えば作れるけど）続かない。
続けていくインセンティブが（サービス運用側に）ないから。
よほどのキラーコンテンツが登場しない限りはね（それか JK の間で流行りだすとか）。
あるいは犯罪者の巣窟になってケーサツに潰されるのもお約束（[Mastodon] って C&C サーバにうってつけだよね`w`）。

ちうわけで様子見。
GW 過ぎたらみんな忘れてるに100カノッサ。
まぁでも私のお予想はたいがい外れるので，外れたら [mstdn.jp] に寄付してもいいかもしれない。

自ドメインに自分専用インスタンスを設置するのはパス。
計算機リソース借りるのだって無料じゃないし Docker イメージがあるとはいえ明らかに面倒くさい（レンタルサーバにおまけでついてくるというのならともかく）。
費用対効果の観点から少なくとも個人がそこまでしてするメリットはないだろう。

一方でサービスベンダが自前でインスタンスを設置するのはあり。
Twitter と違って「名前」の問題で悩む必要がないだけでもメリットはある。
まぁ [Mastodon] をみんなが使っているというのが前提だけどね。

## ブックマーク

- [ASCII.jp：Twitterのライバル？　実は、新しい「マストドン」（Mastodon）とは！｜遠藤諭のプログラミング＋日記](http://ascii.jp/elem/000/001/465/1465842/)
- [ポストTwitterのSNS！？「Mastodon（マストドン）」の初期設定やスマホで使うアプリを紹介するぞ！ | むねさだブログ](http://munesada.com/2017/04/13/blog-9885)
- [Android用「マストドン」アプリ「Tusky」の利用方法とエラー対策](http://did2memo.net/2017/04/14/mastodon-android-app/)
- [Mastodon は自分のドメインでIDを持つことが大事。「リモートフォロー」の価値を最大化するべし。 ｜ 諸多日記](https://isid.ai/diary/2017/04/14/1179/)
- [OStatusの仕様をかいつまんで適当に和訳するよ - hito_asaの日記](http://hitoasa.hateblo.jp/entry/20101013/1286950786)
- [本の虫: ここらでもう一度マストドンについて語っておくか](https://cpplover.blogspot.jp/2017/04/blog-post_20.html)
- [日本のネットを騒がせる「マストドン」、その課題と可能性をえふしん氏に聞いた | TechCrunch Japan](https://jp.techcrunch.com/2017/04/20/mastodon/)
- [企業は安易なMastodonインスタンスの運用を避けるべきでは? - ただのにっき(2017-04-21)](http://sho.tdiary.net/20170421.html)
- [gnusocial や mastodon の哲学 - 何とは言わない天然水飲みたさ](https://blog.cardina1.red/2017/04/13/federated-social-web/)
- [【さくらのクラウドで】マストドンおひとり様用インスタンスを作ってみた【スタートアップスクリプト】 - 脱SEして文筆家になった人](http://yotsumao.hatenablog.com/entry/2017/04/19/%E3%80%90%E3%81%95%E3%81%8F%E3%82%89%E3%81%AE%E3%82%AF%E3%83%A9%E3%82%A6%E3%83%89%E3%81%A7%E3%80%91%E3%83%9E%E3%82%B9%E3%83%88%E3%83%89%E3%83%B3%E3%81%8A%E3%81%B2%E3%81%A8%E3%82%8A%E6%A7%98%E7%94%A8)
- [MastodonをAWSでシュッと動かすやつ - ✘╹◡╹✘](http://r7kamura.hatenablog.com/entry/2017/04/20/014606)
- [Mastodon API gemを使って投稿する - Qiita](http://qiita.com/takahashim/items/a8c0eb3a75d366cfe87b)
- [日本のマストドン インスタンス一覧 - Qiita](http://qiita.com/cv_k/items/8ecafea3ce7dd720cec6)
- [Mastodon インスタンスの画像や動画の保存先をクラウドストレージ （Amazon S3） に移行した話 | WWW WATCH](https://hyper-text.org/archives/2017/04/mastodon-instance-with-amazon-s3.shtml)
- [Masto.Host - Hosting for Mastodon Instances](https://masto.host/) : ホスティングサービス。5ユーロ/月より
- [マストドンのタイムラインをgo-mastodon のWebSocketを使用し取得する - Qiita](http://qiita.com/S-YOU/items/cf677ae282bd6f38fbbb)
- [Mastodon を CentOS にインストールする (Docker未使用) - Qiita](http://qiita.com/bezeklik/items/1a8530d530613acd665c)
- [ゼロからはじめるMastodon - さくらのナレッジ](http://knowledge.sakura.ad.jp/knowledge/8591/)

[Mastodon]: https://github.com/tootsuite/mastodon "tootsuite/mastodon: A GNU Social-compatible microblogging server"
[mstdn.jp]: https://mstdn.jp/
