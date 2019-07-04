+++
title = "本家サイトをリニューアルしました"
date =  "2019-07-04T22:35:58+09:00"
description = "さて，今年の目標は粗かた達成しちゃったな。後半はもう少しノンビリするか。 "
image = "/images/attention/kitten.jpg"
tags = [ "site", "hugo" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

リニューアルした本家サイトの公開を行った。

- [サイトをリニューアルしました — お知らせ | Baldanders.info](https://baldanders.info/announce/site-renewal/)

最初はできるだけ URL を変えないようにと思っていたが，そもそも[本家サイト]を開設した理由は[あちこちのサイト・サービスに書き散らしたコンテンツを集約するため](https://baldanders.info/blog/000005/ "Baldanders.info 正式オープン")だったので，ひとつの CMS に統合するのは無理だと分かった。
分かってしまった。

そこで無理に URL を旧版に合わせることはせず [Hugo] で組みやすいよう再構成することにした。
実は知らなかったのだが（もの知らずでゴメン）[さくらのレンサバ](https://www.sakura.ne.jp/ "さくらのレンタルサーバ")では `.htaccess` を使ったリダイレクトに正規表現が使えるそうなので

```text
RedirectMatch permanent /spiegel/log2/(\d{6}?).shtml /blog/$1/
```

旧 URL の多くは自動でリダイレクトしてくれるはずである。
できなかったらゴメンペコン。

あと [Hugo セクション]({{< rlnk "/hugo/" >}})は[本家サイトに移転](https://baldanders.info/hugo/ "ゼロから始める Hugo | Baldanders.info")させた。
内容が古すぎて手のつけようがないので[本家サイト]のアーカイブとして余生を送らせることにする。
[Hugo] の最新情報については引き続きこちらでフォローする予定なので，今後ともよろしく！

さて，[今年の目標]({{< ref "/remark/2019/01/replace-to-https.md#2019" >}})は粗かた達成しちゃったな。
[本家サイト]のリニューアルはもっとかかると思っていたが [`goquery`] パッケージを使ったページデータの scraping が思ったより簡単にできてしまったので，あとは Go で変換ツールを組んで shell スクリプトでバッチ化すれば完了という（笑）

後半はもう少しノンビリするか。
実は今年の隠れ目標は「政治的無関心を装う」なので参院選も（投票には行くが）知らん顔する予定である。
せっかく田舎に帰郷ったんだから，溢れかえる政治広告に振り回されることなく心静かに生きたいものである。

## ブックマーク

- [rinopo/sakura-init: さくらのレンタルサーバを借りたとき最初にすること](https://github.com/rinopo/sakura-init)

- [（不完全ながら） HTTPS 接続に対応した]({{< ref "/remark/2019/01/replace-to-https.md" >}})
- [さくらのレンタルサーバ上で Hugo によるサイト管理を行う]({{< ref "/remark/2019/01/sakura-and-hugo.md" >}})
- [さくらのレンタルサーバの Git が「使える！」ようになっていた]({{< ref "/remark/2019/05/upgrade-sakura.md" >}})

[本家サイト]: https://baldanders.info/ "Baldanders.info"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[`goquery`]: https://github.com/PuerkitoBio/goquery "PuerkitoBio/goquery: A little like that j-thing, only in Go."

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%95%E3%83%AB%E3%82%B9%E3%82%AF%E3%83%A9%E3%83%83%E3%83%81%E3%81%8B%E3%82%891%E6%97%A5%E3%81%A7CMS%E3%82%92%E4%BD%9C%E3%82%8B_%E3%82%B7%E3%82%A7%E3%83%AB%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%97%E3%83%88%E9%AB%98%E9%80%9F%E9%96%8B%E7%99%BA%E6%89%8B%E6%B3%95%E5%85%A5%E9%96%80-%E6%94%B9%E8%A8%822%E7%89%88-%E3%82%A2%E3%82%B9%E3%82%AD%E3%83%BC%E3%83%89%E3%83%AF%E3%83%B3%E3%82%B4-%E4%B8%8A%E7%94%B0-%E9%9A%86%E4%B8%80-ebook/dp/B07TSZZPWN?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07TSZZPWN"><img src="https://images-fe.ssl-images-amazon.com/images/I/51H%2B4kUhbFL._SL160_.jpg" width="121" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%95%E3%83%AB%E3%82%B9%E3%82%AF%E3%83%A9%E3%83%83%E3%83%81%E3%81%8B%E3%82%891%E6%97%A5%E3%81%A7CMS%E3%82%92%E4%BD%9C%E3%82%8B_%E3%82%B7%E3%82%A7%E3%83%AB%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%97%E3%83%88%E9%AB%98%E9%80%9F%E9%96%8B%E7%99%BA%E6%89%8B%E6%B3%95%E5%85%A5%E9%96%80-%E6%94%B9%E8%A8%822%E7%89%88-%E3%82%A2%E3%82%B9%E3%82%AD%E3%83%BC%E3%83%89%E3%83%AF%E3%83%B3%E3%82%B4-%E4%B8%8A%E7%94%B0-%E9%9A%86%E4%B8%80-ebook/dp/B07TSZZPWN?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07TSZZPWN">フルスクラッチから1日でCMSを作る_シェルスクリプト高速開発手法入門 改訂2版 (アスキードワンゴ)</a></dt>
	<dd>上田 隆一, 後藤 大地</dd>
	<dd>ＵＳＰ研究所 (監修)</dd>
    <dd>ドワンゴ 2019-07-05 (Release 2019-07-05)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B07TSZZPWN</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">既存の常識に凝り固まったソフトウェア・エンジニアに「痛恨の一撃」を加える快書もしくは怪書。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2014-09-21">2014-09-21</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
