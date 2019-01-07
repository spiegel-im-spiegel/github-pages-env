+++
date = "2015-09-23T22:05:15+09:00"
description = "PSK を共有するサービス / ブラックリストの共有はうまくいかない"
draft = false
tags = ["wireless", "authentication", "security", "risk", "market"]
title = "今日の戯れ言：SW 終わった"

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
  url = "https://baldanders.info/spiegel/profile/"
+++

シルバー・ウィーク終わりましたね。
渋滞のない快適な通勤も今日で終わりか。
明日は雨って言うし早めに出かけないとなぁ。

1. [PSK を共有するサービス]({{< relref "#wireless" >}})
1. [ブラックリストの共有はうまくいかない]({{< relref "#adblock" >}})

## PSK を共有するサービス{#wireless}

- [周りのWiFiをいつでも利用し放題にするアプリを作成！通信制限からも解放！ |　クラウドファンディング　-　Makuake（マクアケ）](https://www.makuake.com/project/wifishare/)

実はこれを全く知らなくて，最初聞いたときは「[FON](https://corp.fon.com/) みたいなもん？」と考えていたが，そんなレベルじゃなくて，無線 LAN の PSK (Pre-Shared Key) を不特定多数と共有するという恐ろしいサービスだった。

- [検証結果：高木浩光氏がご立腹の「WiFiシェア」、暗号化してるはずのWi-Fiパスワードを平文で保存してた | Geekles](http://geekles.net/gadget/150922-wifi-share-wi-fi-pass-is-plain-text)

いや PSK が平文で保存されていようが暗号化されていようがダメなものはダメだって。
なぜなら PSK は無線 LAN 機器に対して行う認証を兼ねていて，それを不特定多数と共有するということは「War Driving やり放題」と言っているのと同じことだからだ[^a]。

[^a]: ふつう認証というのは（適切でないユーザを排除するために）ユーザに対して行うものだが，無線 LAN の PSK は認証の向きが逆なのである。無線 LAN でユーザに対して認証を行うには IEEE802.1X （EAP は IEEE802.1X の方式のひとつ）を実装する必要がある。

たとえば，無線 LAN は素通しでも上位レイヤで認証を行う [FON](https://corp.fon.com/) のようなやり方や，いっそ全てのアクセスポイントを無線メッシュネットワークでつなぐ[^b] というのならまだ分かる。
[Makuake](https://www.makuake.com/) は無線 LAN へのアクセスを売り買いすることでマネタイズできると考えているようだが， PSK が知れ渡れば誰でもいくらでもサービスを迂回できてしまう（あるいは WIFI シェアを迂回して接続するアプリだって作れるだろう）。

[^b]: 拙文「[“The Shadow Web” — Baldanders.info](https://baldanders.info/spiegel/log2/000599.shtml)」参照。

なんでこんなサービスがうまくいくと考えているのかよく分からないが，クラウド・ファンディングで資金を集めたところを見ると「うまくいく」と思っている人が多いということなのだろう。
やれやれ。

## ブラックリストの共有はうまくいかない{#adblock}

私は広告はあまり気にならないので（広告による Tracking は気になる），最近話題（？）の広告ブロック・アプリについてはほとんど興味が無いのだが，

- [1分で作るiOS9の広告ブロッカー - Qiita](http://qiita.com/kenmaz/items/65cc4a7ca3ef2eae253b)

content blocker って何でもブロックできちゃうのね。
これって広告の問題じゃないよなぁ。

特定のコンテンツをブロックするための「ブラックリスト」を作ってみんなで共有しようという発想は一見よさ気に見えるが，この手の「ブラックリスト」が最終的に上手くいったという話を聞いたことがない。
何故なら，この手の「ブラックリスト」は極めて恣意的なもので，共有できるほどの普遍性はないから。

昔， spam メールをブロックするために「ブラックリスト」を作って共有する仕組みがあったが，最終的に破綻した。
誤報や恣意的な目的で「ブラックリスト」が更新されるケースが跡を絶たなかったからだ。
これを「能動的失敗」という。
能動的失敗が頻出すれば，誰もそれを信用しなくなり，迂回するようになる。

ブラックリストがかろうじて機能しているのは parental control や企業向けファイアウォールくらいだろう。
この場合でもフィルタ設定をかなり慎重にやらないと，やはり「能動的失敗」を引き起こすことになる。

その昔， [iPad を使ってた頃](https://baldanders.info/spiegel/log2/000487.shtml)は微妙に Safari が気に食わなかった。
何故か Safari はコンテンツを勝手に改変する。
たとえばコンテンツ中に数字列があれば電話番号と判断して妙ちきりんなリンクを張ってくさる。
まぁ，もう Apple 製品は懲りたので，今となってはどうでもいい話なのだが。
