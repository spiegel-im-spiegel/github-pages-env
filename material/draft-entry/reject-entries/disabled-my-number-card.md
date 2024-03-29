+++
title = "【提案】無効化マイナンバーカードを作るべき"
date =  "2020-08-04T09:06:45+09:00"
description = "description"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "identification", "my-number" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近 ID 窃盗に絡む事件が立て続けにあったらしい。

- [日本郵便のe転居を悪用したストーカー事件についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2020/07/17/174628)
- [なりすまし申請による特別定額給付金詐欺事件についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2020/08/02/014951)

いずれも他者になりすましてサービス・アカウントを詐取し悪用するというものだ。

この手の詐欺パターンは実は昔からよくある手だが，これに対抗する手段は「相手より先に ID を実効支配する」しかない。
だが，普段使わないサービスやそもそも存在を知らないサービスについて「相手より先に ID を実効支配する」のは簡単ではない。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}，個人番号カードの話。

個人番号および個人番号カードの最大の問題は

{{< div-gen type="markdown" class="center" >}}
**「識別」と「許可」が未分離**
{{< /div-gen >}}

な点である。
これは住基ネットの頃から問題視されていたが，全く改善されることなく「個人番号」に移行・統合されてしまった。

住基番号や個人番号を使ったオンライン・サービスといえば，今まではせいぜい税務申告をオンライン申請する程度で，番号そのものにさしたる価値はなかったのだが[^mn1]，ここのところの「ばら撒き政策」で個人番号に金銭的価値が発生してしまった。
こうなると，今後は上述したような ID 窃盗の事例は増えていくと予想できる。

[^mn1]: 個人番号のせいで「隠れ副業」がし辛くなって「夜の街」の兼業ホステスさんが大量に引退してしまった，とかいった副作用は割愛する。


個人番号の理想はシステムに関わる actor が全て個人番号を知らないことだ。
そして個人番号が必要であれば，その都度ドメインと期限を定めた「許可トークン」を発行すればいいのだ（発行と破棄は対にしなければならない点に注意）。



これがオンラインでできるなら「個人番号カード」を作る意義は十分ある。





<!-- eof -->
