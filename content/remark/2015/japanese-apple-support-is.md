+++
date = "2015-10-03T11:26:15+09:00"
update = "2015-10-22T10:46:00+09:00"
description = "絶対に Apple は日本のユーザを「極東のカモネギ」と思ってるよ。いいの？"
draft = false
tags = ["security", "vulnerability", "risk", "apple"]
title = "Apple の日本語サポートがクズ過ぎるので"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

今週出た Apple 製品の一連の脆弱性情報。

- [About the security content of iOS 9.0.2 - Apple Support](https://support.apple.com/en-us/HT205284)
- [About the security content of Safari 9 - Apple Support](https://support.apple.com/en-us/HT205265)
- [About the security content of OS X El Capitan v10.11 - Apple Support](https://support.apple.com/en-us/HT205267)

まぁ，私はもう Apple ユーザではないし，ましてや信者でもないので，グダグダ言ってもしょうがないんだけど。
でも，一応仕事に絡むかもしれないからセキュリティ情報だけはチェックしているわけですよ。

[iOS がロックスクリーンの脆弱性](http://news.mynavi.jp/news/2015/09/29/338/)に迅速に対処したのは（何かのついでなんだろうけど）素晴らしいんだけどね。
OS X のほう。
これ酷すぎないかい。
たとえば bash の CVE-2014-6277, CVE-2014-7186, CVE-2014-7187 ってこれ1年前の [shellshock 脆弱性](http://www.baldanders.info/spiegel/log2/000743.shtml)じゃん。
あんなに大騒ぎになったのに，今頃かよ！

ほんでもってこれらの脆弱性は Yosemite （およびそれ以前の）ユーザには手当てしないんだよね。
相変わらず鬼畜だなぁ。
みんなよく信者やってるよ。
ユーザ企業の側も，こんなセキュリティ・マネジメントの甘い企業の製品とかなんで使うの？ IBM が Apple と提携するとか何かの冗談としか思えん。

で，いちばんクズなのは日本語サポート。
Apple 製品のセキュリティ情報は [`https://support.apple.com/kb/HT201222`](https://support.apple.com/kb/HT201222) から取得できるんだけど，おそらく IP アドレスのリージョンかブラウザの言語情報を見て [`https://support.apple.com/ja-jp/HT201222`](https://support.apple.com/ja-jp/HT201222) の日本語ページにリダイレクトされちゃうのですよ[^a]。
でも日本語ページは10月5日の今になっても古い情報のまま。
いつも何故かちょうど1世代前の情報になっている。
これ絶対わざとだよね。
日本語圏のユーザはこれでいいわけ？

[^a]: 1年以上前から？ こんな感じになってしまった。

まぁ [`https://support.apple.com/en-us/HT201222`](https://support.apple.com/en-us/HT201222) のほうを直で開けばいいので実質的に困ることはないんだけどさ（そもそも私は Apple ユーザじゃないし），なんで日本のユーザは怒らないのかね。
絶対に Apple は日本のユーザを「極東のカモネギ」と思ってるよ。
いいの？
