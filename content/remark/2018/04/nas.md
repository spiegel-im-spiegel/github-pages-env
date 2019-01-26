+++
title = "NAS を導入した"
date = "2018-04-30T15:32:40+09:00"
description = "Win7 のサポートが切れる2020年までに Windows は捨てる予定なので，今回の NAS 導入はその前準備というところである。"
image = "/images/attention/kitten.jpg"
tags        = [ "engineering", "nas", "samba" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

GW に突入して最初の作業は NAS (Network-Attached Storage) の導入だった。
つっても，今までパソコンの「外付けハードディスク」として使っていたものを家庭内 LAN 上で使えるようにしただけなのだが。
ストレージは増設して 2TB に。
つか，安い簡易 NAS に件の外付け HDD を組み込んだという方が正しいかな。

大昔は職場で減価償却の終わった古い PC を安く譲ってもらって Linux をぶち込んでたりしたのだが，そういうのもすっかり面倒くさくなって（今は PC を捨てるのにもお金かかるしねぇ），というかサーバ立てて実験するならクラウド・サービスを使うほうが早い。
実験用としても自宅でフルスクラッチでサーバを組むメリットはなくなっている（ラズパイ等で遊ぶのならありだろうが）。

とりあえず Windows 7 と Android 携帯端末から [Samba] 経由で NAS にアクセスできることは確認したので問題ないかな。
日本語ファイル名のモノがかなりあって心配したのだが，最近のは優秀なんだねぇ。
アクセス制御を含む各種設定も Web ブラウザでできるし。
Box や Dropbox といったクラウドストレージとも連携できるし。

Win7 のサポートが切れる2020年までに Windows は捨てる予定なので[^win7]，今回の NAS 導入はその前準備というところである。
ここのところの Windows 関連の脆弱性情報を見るに Microsoft は既に Win7 にリソースを割いてない（割けない）印象を受ける。

[^win7]: Windows を捨てる理由は個人的な（Win8 以降に馴染めない）ものなので，今後も Windows を使い続けたいなら，なるべく早く Windows 10 に移行することを強くお勧めする。なお Windows はメジャー・バージョンを「アップグレード」する際にトラブルが起きやすいので，この際マシンごと買い換えるほうがいいと思う。

- [「KB4099950」が公開、3月パッチでWindows 7のネットワーク設定が失われる問題を修正 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1114999.html)
- [IEに脆弱性、ゼロデイ攻撃に悪用と指摘 - ZDNet Japan](https://japan.zdnet.com/article/35118231/)

ひょっとして予定を前倒ししたほうがいいかも。

## ブックマーク

- [File Explorer Pro - file manager and file transfer - Google Play](https://play.google.com/store/apps/details?id=com.skyjos.apps.fileexplorer)

[Samba]: https://www.samba.org/ "Samba - opening windows to a wider world"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%EF%BC%BB%E8%A9%A6%E3%81%97%E3%81%A6%E7%90%86%E8%A7%A3%EF%BC%BDLinux%E3%81%AE%E3%81%97%E3%81%8F%E3%81%BF-%EF%BD%9E%E5%AE%9F%E9%A8%93%E3%81%A8%E5%9B%B3%E8%A7%A3%E3%81%A7%E5%AD%A6%E3%81%B6OS%E3%81%A8%E3%83%8F%E3%83%BC%E3%83%89%E3%82%A6%E3%82%A7%E3%82%A2%E3%81%AE%E5%9F%BA%E7%A4%8E%E7%9F%A5%E8%AD%98-%E6%AD%A6%E5%86%85-%E8%A6%9A-ebook/dp/B079YJS1J1?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B079YJS1J1"><img src="https://images-fe.ssl-images-amazon.com/images/I/51CPvtuv%2BwL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%EF%BC%BB%E8%A9%A6%E3%81%97%E3%81%A6%E7%90%86%E8%A7%A3%EF%BC%BDLinux%E3%81%AE%E3%81%97%E3%81%8F%E3%81%BF-%EF%BD%9E%E5%AE%9F%E9%A8%93%E3%81%A8%E5%9B%B3%E8%A7%A3%E3%81%A7%E5%AD%A6%E3%81%B6OS%E3%81%A8%E3%83%8F%E3%83%BC%E3%83%89%E3%82%A6%E3%82%A7%E3%82%A2%E3%81%AE%E5%9F%BA%E7%A4%8E%E7%9F%A5%E8%AD%98-%E6%AD%A6%E5%86%85-%E8%A6%9A-ebook/dp/B079YJS1J1?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B079YJS1J1">［試して理解］Linuxのしくみ ～実験と図解で学ぶOSとハードウェアの基礎知識</a></dt>
	<dd>武内 覚</dd>
    <dd>技術評論社 2018-02-23 (Release 2018-02-23)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B079YJS1J1</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">コンテナ全盛のこの時代にかなり硬派な内容の Linux 解説書。コンピュータの教科書としても使えそう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-01-07">2019-01-07</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
