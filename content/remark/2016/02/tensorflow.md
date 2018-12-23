+++
date = "2016-02-18T12:55:42+09:00"
update = "2016-03-14T01:12:59+09:00"
description = "TensorFlow は Google がオープンソースとして提供する多次元配列（tensor）演算（flow）ライブラリ。いわゆる「機械学習」で威力を発揮する。"
draft = false
tags = ["tensorflow", "bookmark", "engineering", "artificial-intelligence"]
title = "TensorFlow に関するブックマーク"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（この記事は「[プログラミング言語との付き合い方]({{< ref "/remark/2015/programming-language.md" >}})」のおまけの項からの転載です）

[TensorFlow] は Google がオープンソースとして提供する多次元配列（tensor）[演算（flow）](https://en.wikipedia.org/wiki/Flow_%28mathematics%29)ライブラリ。
いわゆる「機械学習（machine learning）」で威力を発揮し， [TensorFlow] の元となっている Google 内製のインフラ [DistBelief](http://research.google.com/pubs/pub40565.html) では既に実績がある。

バックエンドは C++ で構築しているそうだが，フロントエンドでは Python が使える（将来的には他の言語にも対応するそうだ）。
携帯端末から GPU バリバリのワークステーションまでスケーラブルに対応し，簡易な記述で実装できるのが特徴。
可視化ツールもある。

たしかに「[何か作れそうな気がする](https://plus.google.com/+HidekiSaito/posts/EJZgMkANqou)」感じではある。

- [TensorFlow is an Open Source Software Library for Machine Intelligence](http://tensorflow.org/)
    - [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)
- [Google、機械学習システム「TensorFlow」をオープンソースで公開 - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1511/10/news055.html)
- [Googleの公開した人工知能ライブラリTensorFlowを触ってみた - 株式会社ネクスト　エンジニアBlog](http://nextdeveloper.hatenablog.com/entry/2015/11/10/204609)
- [TensorFlowを算数で理解する - Qiita](http://qiita.com/icoxfog417/items/fb5c24e35a849f8e2c5d)
- [TensorFlow 畳み込みニューラルネットワークで手書き認識率99.2%の分類器を構築 - Qiita](http://qiita.com/haminiku/items/36982ae65a770565458d)
- [自然言語処理をはじめたい人のためのゆるい記事 - Qiita](http://qiita.com/kazuhirokomoda/items/a4cd0f6f42eb75c757e4) : [TensorFlow] についても少しだけ言及
- [わざわざTensorFlowの機械学習で$\sqrt{2}$を計算する - Qiita](http://qiita.com/n_kats_/items/73538c7c66559d09f35d)
- [Python - 初めてのTensorFlow - イントロダクションとしての線形回帰 - Qiita](http://qiita.com/TomokIshii/items/f355d8e87d23ee8e0c7a)
- [TensorFlowチュートリアル - マンデルブロ集合（翻訳） - Qiita](http://qiita.com/KojiOhki/items/00ae0297f6809bdbc484)
- [TensorFlowライブラリによる機械学習モデルの、本番アプリケーションへの実装を助けるAPI集TensorFlow ServingをGoogleがリリース | TechCrunch Japan](http://jp.techcrunch.com/2016/02/17/20160216google-makes-it-easier-to-take-machine-learning-models-into-production/)

{{< fig-youtube id="oZikw5k_2FM" title="TensorFlow: Open source machine learning - YouTube" width="500" height="281" >}}

[TensorFlow]: http://tensorflow.org/ "TensorFlow is an Open Source Software Library for Machine Intelligence"
