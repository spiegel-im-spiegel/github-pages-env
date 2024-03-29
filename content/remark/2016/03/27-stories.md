+++
date = "2016-03-27T02:38:30+09:00"
update = "2016-04-28T21:52:18+09:00"
description = "安全と安心をいっしょくたにする輩を信用してはいけない / 漫画家小山田いくさん死去 / RFC 7763 The text/markdown Media Type / その他の気になる記事"
draft = false
tags = ["security", "risk", "management", "comic", "markdown", "rfc"]
title = "週末スペシャル： 安全と安心をいっしょくたにする輩を信用してはいけない"

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
  url = "https://baldanders.info/profile/"
+++

1. [安全と安心をいっしょくたにする輩を信用してはいけない]({{< relref "#sec" >}})
1. [漫画家小山田いくさん死去]({{< relref "#iku" >}})
1. [RFC 7763 The text/markdown Media Type]({{< relref "#md" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## 安全と安心をいっしょくたにする輩を信用してはいけない{#sec}

一昨年前に Facebook で呟いていた。

> 安心は思い込みさえすればお手軽に手に入るが、安全を手に入れるには少なくないコストがかかるし、誰かから貰えたりもしない。

どんな文脈で言ったか憶えてないが，まぁいつものことだろう。

安全でなくても安心（感）は手に入る。
そして人が素朴に求めるのは安全ではなく安心のほうだ。
これが悪いほうに作用した典型例が Apple 製品である。

{{< fig-youtube id="bmpSJwZ5xqA" width="500" height="375" title="セキュリティ MacCM - YouTube" >}}

当時， Apple はユーザに対して安心感を「売る」キャンペーンを大々的に展開した。
また App Store からセキュリティ関連アプリを排除したりしたこともあった（今はどうなってるか知らないが）。
Apple の失敗はユーザに安心感を与えるために安全を omit してしまったことだ。
そのせいで Apple は大勢のユーザを危険に晒すことになる。

- [FBIとAppleの対立が急転、サン電子子会社がiPhoneロック解除で協力か | マイナビニュース](http://news.mynavi.jp/news/2016/03/24/162/)
- [iPhoneの暗号化をAppleの協力なしで解除する7つの方法 - GIGAZINE](http://gigazine.net/news/20160324-fbi-unlock-iphone/)
- [司法省、対Apple訴訟を取り下げ―テロ容疑者のiPhoneはFBIがアンロックに成功 | TechCrunch Japan](https://techcrunch.com/2016/03/28/justice-department-drops-lawsuit-against-apple-over-iphone-unlocking-case/)

今まで何度も言っているが「警察にできることは犯罪者にもできる」。
自力で iPhone を突破できなかった FBI は外部の企業を使うようだ。
つまり警察は犯罪者以下だと証明してしまったわけだ。

IPA も気をつけなよ。

- [「つながる世界の開発指針」を公開：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/sec/reports/20160324.html)

### 追加ブックマーク

- [米政府によるスマホデータ取り出しの協力要請、ACLUが実態調査 - CNET Japan](http://japan.cnet.com/news/society/35080404/)
- [FBIのiPhoneロック解除方法、Appleに知らされない可能性も (1/3) - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1604/01/news114.html)
- [FBI長官、「購入したロック解除ツールはiPhone 5sでは機能しなかった」 - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1604/08/news060.html)
- [FBIは銃撃犯のiPhoneをハッキングしたツールの詳細をAppleに開示しない意向 | TechCrunch Japan](https://techcrunch.com/2016/04/26/fbi-to-keep-apple-guessing-on-san-bernardino-iphone-hack/)

## 漫画家小山田いくさん死去{#iku}

{{< fig-gen >}}
<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">小山田いく先生。お亡くなりになりました。<br>お疲れ様でした。</p>&mdash; 田上 喜久 (@HgKjDWcYgmlilNs) <a href="https://twitter.com/HgKjDWcYgmlilNs/status/713252253725904896">2016年3月25日</a></blockquote>
{{< /fig-gen >}}

- [漫画家・小山田いくさん亡くなる　代表作に出身地・長野県小諸市を舞台にした「すくらっぷ・ブック」など - ねとらぼ](http://nlab.itmedia.co.jp/nl/articles/1603/25/news127.html)

「すくらっぷ・ブック」の主要キャラは私より1歳下（連載当時）で思春期真っ盛りの私はシンクロしまくりだった。
社会人になって読み返したらこっ恥ずかしい台詞のオンパレードで赤面するのだが，今でも1980年代までに出た単行本は実家に全部ある。

手塚治虫さんとか石森章太郎さんとかの作品って私からはだいぶ上の世代なんだよね。
私等くらいの世代だと小山田いくさんとかご兄弟のたがみよしひささんとか竹本泉さんとかあたりが思春期リアルなんじゃないだろうか[^a]。

[^a]: まぁ他にも新谷かおるさんとか聖悠紀さんとか人によって色々あるだろうけど。

その時代の作家さんが亡くなるのはやっぱくるねぇ。
合掌。

## RFC 7763 The text/markdown Media Type{#md}

Markdown が RFC になったらしい。

- [RFC 7763 - The text/markdown Media Type](https://tools.ietf.org/html/rfc7763)

文法についても定義されているけど markdown って既に方言バリバリだからなぁ。

## その他の気になる記事{#other}

- [AtomとPlantUMLで爆速UMLモデリング - Qiita](http://qiita.com/nakahashi/items/3d88655f055ca6a2617c)
    - [Atom+PlantUMLで見た目もいい感じのシーケンス図を作成する - Qiita](http://qiita.com/k_nakayama/items/77ca73753ebd049a66de)
- [【意訳】たった今、npmのパッケージを解放しました。 - Qiita](http://qiita.com/chuck0523/items/ee23293f2645d40cb317)
    - [The npm Blog — kik, left-pad, and npm](http://blog.npmjs.org/post/141577284765/kik-left-pad-and-npm)
- [モバイル最若年層では、Eメールが死につつある | TechCrunch Japan](https://techcrunch.com/2016/03/24/email-is-dying-among-mobiles-youngest-users/)
- [MacOSとWindowsのネイティブ仮想化を用いたDocker純正ツール「Docker for Mac/Windows」登場、VirtualBoxは不要に － Publickey](http://www.publickey1.jp/blog/16/docker_for_macwindows.html)

## 参考図書

{{% review-paapi "483544213X" %}} <!-- すくらっぷ・ブック -->
