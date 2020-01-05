+++
title = "NAS を導入した"
date = "2018-04-30T15:32:40+09:00"
description = "Win7 のサポートが切れる2020年までに Windows は捨てる予定なので，今回の NAS 導入はその前準備というところである。"
image = "/images/attention/kitten.jpg"
tags        = [ "engineering", "nas", "samba" ]

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

{{% review-paapi "B079YJS1J1" %}} <!-- ［試して理解］Linuxのしくみ -->
