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
- [htaccessで正規表現を使ったリダイレクト - Qiita](https://qiita.com/bass-inu/items/d239cdee54a74ec0d3ef)

- [（不完全ながら） HTTPS 接続に対応した]({{< ref "/remark/2019/01/replace-to-https.md" >}})
- [さくらのレンタルサーバ上で Hugo によるサイト管理を行う]({{< ref "/remark/2019/01/sakura-and-hugo.md" >}})
- [さくらのレンタルサーバの Git が「使える！」ようになっていた]({{< ref "/remark/2019/05/upgrade-sakura.md" >}})

[本家サイト]: https://baldanders.info/ "Baldanders.info"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[`goquery`]: https://github.com/PuerkitoBio/goquery "PuerkitoBio/goquery: A little like that j-thing, only in Go."

## 参考図書

{{% review-paapi "B07TSZZPWN" %}} <!-- フルスクラッチから1日でCMSを作る_シェルスクリプト高速開発手法入門 -->
