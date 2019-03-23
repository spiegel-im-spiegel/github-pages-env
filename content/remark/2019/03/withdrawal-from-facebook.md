+++
title = "Facebook はもう駄目か知らん"
date =  "2019-03-23T11:40:26+09:00"
description = "description"
image = "/images/attention/kitten.jpg"
tags = [ "facebook", "privacy", "surveillance-capitalism" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

かつて Google は[「完全なプライバシーは存在しない」と公言](http://www.thesmokinggun.com/documents/internet/google-complete-privacy-does-not-exist)し「[Google はプライバシーに敵対的](http://cloud.watch.impress.co.jp/epw/cda/infostand/2007/06/18/10531.html)」と大きな批判を浴びていた。
しかし「プライバシーに敵対的」な企業・組織・国家は Google だけではない。

[Campbridge Analytica の件]({{< ref "/remark/2018/03/name-identification.md" >}})以来 Facebook は大きな批判にさらされているが，上述したように「プライバシーに敵対的」で「監視資本主義の走狗」たる企業なんて（規模の大小はともかく）いくらでもあるし，特定の企業・サービスを名指しで批判することは問題を矮小化させるだけだろうと控えてきたつもりである。

なので「[Facebookが数億人のパスワードを平文で保存していたと認める](https://jp.techcrunch.com/2019/03/22/2019-03-21-facebook-plaintext-passwords/ "Facebookが数億人のパスワードを平文で保存していたと認める  |  TechCrunch Japan")」という記事のタイトルを見たときも「またやらしてるよぱすわーどへんこうしなきゃ」くらいにしか思わなかったのだが，どうやらもっととんでもない話のようである。

- [Facebookが一部ユーザーのパスワードを平文記録していた問題についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2019/03/22/061503)

インターネットにおける認証の多くはサービスとユーザの間で「秘密」を共有・交換することで成立する。
たとえば生体情報は「秘密」ではないので認証には使えない（識別には使える）。
パスワードはそれが互いの間で「秘密」である限り認証として有効である。

サービス・プロバイダはパスワードという「秘密」を秘密として維持するために苦労している。
理想はサービス・プロバイダの誰もユーザの設定したパスワードを取得できないようにすることである。
故にパスワードが平文のままで保持されている状態はセキュリティ上の重大な瑕疵（defect）となる。

しかし今回の件は瑕疵ではない。
サービス・プロバイダである Facebook によるユーザ情報の「悪用」である。
なぜなら「秘密」を組織内部で「共有」していたのだから。

{{< fig-quote >}}
{{< /fig-quote >}}


## ブックマーク

- [Facebook Stored Hundreds of Millions of User Passwords in Plain Text for Years —  Krebs on Security](https://krebsonsecurity.com/2019/03/facebook-stored-hundreds-of-millions-of-user-passwords-in-plain-text-for-years/)

- [Change your Facebook password. And don't try to remember it. - F-Secure Blog](https://blog.f-secure.com/change-your-facebook-password-and-dont-try-to-remember-it/)
- [Facebookが数億人のパスワードを平文で保存していたと認める  |  TechCrunch Japan](https://jp.techcrunch.com/2019/03/22/2019-03-21-facebook-plaintext-passwords/)
- [Facebookが数億件分のパスワードを暗号化しないままサーバーに保存していたことが明らかに - GIGAZINE](https://gigazine.net/news/20190322-facebook-stored-user-passwords-plain-text/)
