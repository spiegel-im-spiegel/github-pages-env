+++
date = "2015-09-26T21:47:01+09:00"
description = "というわけで， Microsoft が Windows 10 で自信を持ってお送りする Wi-Fi Sense についていくつか記事を紹介する。"
draft = false
tags = ["security", "risk", "wireless", "windows"]
title = "Windows 10 で PSK を共有する"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

[この前]({{< relref "remark/2015/0923-diary.md#wireless" >}})，「WiFi シェア」を盛大に dis ったわけだが，「Windows 10のWi-fi Senseも仲間に入れてあげてください」とフィードバックを頂いた。
確かにそのとおりだ！

というわけで， Microsoft が Windows 10 で自信を持ってお送りする {{% quote lang="en" %}}Wi-Fi Sense{{% /quote %}} についていくつか記事を紹介する。
なお，私は仕事でも私用でも Windows 10 機は持ってない[^a] ので悪しからず。

[^a]: 古参の Windows ユーザある私が言うのもナニだが，現状で Windows 7 or 8.x に不満がないユーザは無理に Windows 10 にする必要はない。 XP や Vista のユーザなら多少はメリットは有るかもしれないが。 Windows 7 でも2020年まではサポートされるわけで，その間にゆっくり評価すればいいと思う。というか，5年後も Windows がメジャーであり続けるかどうかすら現時点ではわからない。パワーユーザ以外の多くのユーザがスマホやタブレットをメインに使うようになるなら， Windows に固執する必要は全くない。

- [友人とWi-Fiパスワードの自動共有ができるWindows 10新機能「Wi-Fi Sense」 - GIGAZINE](http://gigazine.net/news/20150703-windows-10-wifi-sense/)
- [Windows 10: Wi-Fi センサー (Wi-Fi Sense) とは? 正しく理解して使用してみよう - 日本のセキュリティチーム - Site Home - TechNet Blogs](http://blogs.technet.com/b/jpsecurity/archive/2015/08/21/windows-10-wifi-sense.aspx)
- [社内 Wi-Fi を Windows 10 の Wi-Fi 共有対象から外してパスワードを変更する](https://r2.ag/counter-windows-10-wifi-sense/)
- [Windows 10 uses your bandwidth to share updates](http://thenextweb.com/microsoft/2015/07/30/windows-10-steals-your-bandwidth-to-send-other-people-updates/)

{{< fig-quote title="社内 Wi-Fi を Windows 10 の Wi-Fi 共有対象から外してパスワードを変更する" link="https://r2.ag/counter-windows-10-wifi-sense/" >}}
<q>Windows 10 には Wi-Fi Sense という新機能があり、これは Wi-Fi スポットを自分の連絡先に自動的に共有できるものです。一度有効にすると、 Microsoft のサーバーに保存された Wi-Fi パスワードが連絡先に晒されてしまいます。</q>
{{< /fig-quote >}}

{{< fig-gen >}}
<blockquote class="twitter-tweet" lang="ja"><p lang="ja" dir="ltr">MSのWifi Senseは貴方がWin10を使っていないくても誰かWin10ユーザにPWを教えてWifiを使わせるとそのユーザの連絡帳のユーザは全く知らない人でもあなたのwifiを利用可能になる。止めるにはSSID変更のみ RE <a href="https://t.co/vZcHDhAqED">https://t.co/vZcHDhAqED</a></p>&mdash; 高梨陣平 (@jingbay) <a href="https://twitter.com/jingbay/status/628343560291622912">2015, 8月 3</a></blockquote>
{{< /fig-gen >}}

まぁ「WiFi シェア」が安全だというのなら {{% quote lang="en" %}}Wi-Fi Sense{{% /quote %}} も同程度には安全（笑）だと思うので，そっちを使ったほうがいいと思うよ。
{{% quote lang="en" %}}Wi-Fi Sense{{% /quote %}} を使うだけならお金取られないみたいだし。

繰り返して言うが PSK 方式はユーザに対する認証ではない。
ユーザに対する認証が必要なら IEEE802.1X の認証システムが必要だが，家庭用の無線 LAN ルータでこれを備えているものはほぼ存在しない[^b]。
OS ベンダやアプリ・ベンダが PSK 共有を「安全」だなどと言いはるなら，ユーザ側は自衛のためにも，家族にも誰にも PSK を教えてはいけない。
1人1ルータ時代の到来ですな（笑）

[^b]: 法人向けにはいくつか製品があるが，当然高い。
