+++
title = "ストップ!! Cortana 君！"
date =  "2020-10-25T12:14:44+09:00"
description = "こういうのはオプト・インにしてほしい。"
image = "/images/attention/kitten.jpg"
tags = [ "windows" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

半年の無職期間に{{% ruby "なま" %}}鈍{{% /ruby %}}りきった身体と心を徐々に再起動中です。

新しい職場で Windows 10 機を支給されたのだが，いまどき LAN アクセスでピコピコとランプが点滅するマシンらしく，しばらく眺めていたのだが，激しい点滅が止まらんのですよ。

「何してるのかなぁ」とリソースモニタを眺めてみると `SearchUI.exe` なるプロセスがエラい勢いで通信してるみたい。
Windows 10 については全く予備知識がなく `SearchUI.exe` についても知らなかったのだが（最初は malware かと思ったぜ`w`）どうもこれ Cortana とかいう音声入力アシスタントによるプロセスらしい。

ちょろんとググってみたが，バッテリは食うし何もしてなくても激しく通信してくさるので，迷惑に感じてるユーザは多いようだ。
無効化する方法を紹介する記事もあったので，とっとと無効化した。

- [Windows 10でCortanaが起動しないようにする方法 | ライフハッカー［日本版］](https://www.lifehacker.jp/2020/06/how-to-quickly-remove-windows-10s-new-cortana-app.html)
- [SearchUI.exeは停止しても大丈夫？停止方法をご紹介！ | Aprico](https://aprico-media.com/posts/3514)

これで再度リソースモニタを確認してプロセスがないのを確認。
よーし，うむうむ，よーし。

業務によっては便利なのかもしれないが，基本的に迷惑アプリケーションだよな，これ。
こういうのはオプト・インにしてほしい。


{{% review-paapi "B01ID14CC4" %}} <!-- ストップ!!ひばりくん！ -->
{{% review-paapi "B002SUPHGM" %}} <!-- ごとそん -->
