+++
date = "2015-09-20T18:33:08+09:00"
update = "2015-09-23T19:57:00+09:00"
description = "仕事部屋が広くなった orz / 国勢調査を片付けた / Git for Windows 2.5.3"
draft = false
tags = ["book", "comic", "phishing", "git"]
title = "今日の戯れ言：週末は掃除三昧"

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

世間ではシルバーウィークとかなんとか言うらしいですが，私には関わりのねぇこってござんす。

## 仕事部屋が広くなった orz

ついに実家の親から仕事（魔窟）部屋をかたせ，と命令が下りまして（他にも色々大変だったんだけど），泣く泣く本を手放した。

今の場所に引っ越してきたのは16年ほど前だけど，主に漫画本が溜まりまくってて。
軽トラ1台分の単行本を処分。
まぁ，最近は Kindle で買えるし，いいんだけどね。
15年くらい前は4コマで人気だったのに，今ではさっぱり名前を聞かない作家さんの本とかもあって，惜しむ気持ちはあったんだけど，命令には逆らえません。

でも好きな作家さんの本はどうしても捨てられませんでした。
というわけで

{{< fig-img src="./21549048532_o.jpg" alt="竹本泉エリア。整理したら３段ぶち抜きになってしまった" title="竹本泉エリア。整理したら３段ぶち抜きになってしまった" link="./21549048532_o.jpg" width="1080" >}}

漫画本の本棚が大変なことに。

仕事関係の本は今回は手付かず。
（Kindle 含め） E ブックで買えるなら，これこそ紙で買う必要はないんだけどね。

ラノベで紙の本で買ってたものは全て処分した。
Kindle で買えるし。

## 国勢調査を片付けた

国勢調査の回答をネットで片付けた。

- [国勢調査オンライン窓口 - 総務省統計局](http://www.e-kokusei.go.jp/)

Symantic の証明書かぁ。
いや，まぁ，いいんだけどね。
そういや，偽サイト作って怒られてた人いたな。

- [国勢調査の“偽サイト”作った意図は？　総務省から削除依頼……「騒ぎになり深く反省」と制作者 (1/3) - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1509/15/news083.html)

常識だと思うが， Phishing サイト自体は誰でも作れる。
昔は日本語が怪しかったりして，ひと目でわかるようなものが多かったが，今の職業犯罪者は明確に「憐れな日本人」ユーザをターゲットにしているので，見た目ではわからないことも多い。
従って，そのサイトまたはページが本物かどうかは URI と証明書で判断するしかない。

幸い今回は調査員がマニュアルとアカウント情報を手渡しでくれるので（調査員が信用できるのなら）それほど問題ではないだろう。
怪しい電子メールに書いてある URL や 検索サービスの検索結果に表示されているページをそのまま開くのでないなら，ね。

## Git for Windows 2.5.3

- [Release Git for Windows 2.5.3 · git-for-windows/git](https://github.com/git-for-windows/git/releases/tag/v2.5.3.windows.1)

Windows 版 git が MSYS2 ベースになって，本家のアップデートに素早く対応できるようになったのはいいんだけど，アップデートのペースが速いよ。
もうちょっとなんとかならんのん，これ。
