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

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B079YJS1J1/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51CPvtuv%2BwL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B079YJS1J1/baldandersinf-22/">［試して理解］Linuxのしくみ ～実験と図解で学ぶOSとハードウェアの基礎知識</a></dt><dd>武内 覚 </dd><dd>技術評論社 2018-02-23</dd><dd>評価<abbr class="rating" title="4"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B079TLW41L/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B079TLW41L.09._SCTHUMBZZZ_.jpg"  alt="エンジニアリング組織論への招待　～不確実性に向き合う思考と組織のリファクタリング"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B07C3JFK3V/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B07C3JFK3V.09._SCTHUMBZZZ_.jpg"  alt="前処理大全［データ分析のためのSQL/R/Python実践テクニック］"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B07BKVP9QY/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B07BKVP9QY.09._SCTHUMBZZZ_.jpg"  alt="独学プログラマー Python言語の基本から仕事のやり方まで"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B078J4TNT1/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B078J4TNT1.09._SCTHUMBZZZ_.jpg"  alt="低レベルプログラミング"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B07CM2YNVD/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B07CM2YNVD.09._SCTHUMBZZZ_.jpg"  alt="まんがでわかるLinux シス管系女子3"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B07B8T1F4R/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B07B8T1F4R.09._SCTHUMBZZZ_.jpg"  alt="江添亮の詳説C++17 (アスキードワンゴ)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B075ST51Y5/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B075ST51Y5.09._SCTHUMBZZZ_.jpg"  alt="ふつうのLinuxプログラミング 第2版　Linuxの仕組みから学べるgccプログラミングの王道"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B079Q6S12G/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B079Q6S12G.09._SCTHUMBZZZ_.jpg"  alt="コンテナ・ベース・オーケストレーション Docker/Kubernetesで作るクラウド時代のシステム基盤"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B07BBTSX65/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B07BBTSX65.09._SCTHUMBZZZ_.jpg"  alt="クラウドエンジニア養成読本［クラウドを武器にするための知識＆実例満載！］ Software Design plus"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B07B8GH29F/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B07B8GH29F.09._SCTHUMBZZZ_.jpg"  alt="C言語本格入門 ～基礎知識からコンピュータの本質まで"  /></a> </p>
<p class="description">コンテナ全盛のこの時代にかなり硬派な内容の Linux 解説書。コンピュータの教科書としても使えそう。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-04-30">2018-04-30</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
