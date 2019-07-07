+++
title = "7iD リスク"
date =  "2019-07-07T16:26:46+09:00"
description = "現金決済であれキャッシュレス決済であれ，基盤となるシステムを信用するからこそ成り立つものであり，信用がなくなれば「そこで試合終了」なんだけど。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "defect", "authentication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

7pay の件は Twitter の TL に流れてくる情報を中心に眺めていたが，2010年代も終わろうかという頃に「サイバーノーガード戦法」とかあまりに馬鹿すぎて他山の石にもならないようだ。
リンクは以下の2つだけ張っておけば十分だろう。

- [7payの不正利用についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2019/07/04/065925)
- [7payを使った不正購入事案についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2019/07/05/055548)

被害にあった方々にはお見舞い申し上げるとともに（被害額の返金だけではなく）賠償請求を行ってきちんとペナルティを支払わせることを強くお勧めする。
向こうは「謝罪して返金すりゃいいんだろ」って感じだし。
バブル崩壊以降，ホンマに「謝罪」が安くなったよなぁ。
まぁお金が絡む話に謝罪とか一銭の価値もない。

今回の件は脆弱性（vulnerability）というより設計上の欠陥（defect）と位置づけるのが適切だろう。
そうなると，その認証基盤である [7iD（旧 omni7）](https://www.omni7.jp/)ってどうなん？ という話になる。

で，[サイトを覗いてみた](https://www.omni7.jp/)ら提携企業の多さに「これヤバくね？」って感じなのだが，それらの企業は今回の件をどう思っているのかね。
現金決済であれキャッシュレス決済であれ，基盤となるシステムを信用するからこそ成り立つものであり，信用がなくなれば「[そこで試合終了](https://dic.nicovideo.jp/a/%E3%81%82%E3%81%8D%E3%82%89%E3%82%81%E3%81%9F%E3%82%89%E3%81%9D%E3%81%93%E3%81%A7%E8%A9%A6%E5%90%88%E7%B5%82%E4%BA%86%E3%81%A0%E3%82%88)」なんだけど。

今後も情報を集めるとしたら，その辺を重点的に見る必要があるかもなぁ。

そもそも「QR決済」以前に「QRコード」そのものを信用していないので[^qr1]，「QRコード」が絡むサービスは極力使わないようにしているのだが，今後もしばらくその状態は続きそうである。

[^qr1]: だってQRコードのあの図形を見ても何についての情報かすら分からないし（human-readable でない），汎用のスキャナを使って復号した情報を確認してから使うのならまだしも，最近のQRコードを使ったサービスはそれすら省いてる感じだし，使いたくない。

あと Twitter 眺めてて気になったのだが「2段階認証（2-step authentication）」と「2要素認証（2-factor authentication）」は意味が違うし，昔は「2段階認証」と称してパスワードを2つ登録させる馬鹿なサービスもあって安全性に対する印象が悪いのだけど，みんながみんな「2段階認証」を連呼するのはその辺も踏まえての話なのかね。

## ブックマーク

- [ひさしぶり「サイバーノーガード戦法」 | Baldanders.info](https://baldanders.info/blog/000470/)

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E4%BF%A1%E9%A0%BC%E3%81%A8%E8%A3%8F%E5%88%87%E3%82%8A%E3%81%AE%E7%A4%BE%E4%BC%9A-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC/dp/4757143044?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4757143044"><img src="https://images-fe.ssl-images-amazon.com/images/I/413qoSjODUL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E4%BF%A1%E9%A0%BC%E3%81%A8%E8%A3%8F%E5%88%87%E3%82%8A%E3%81%AE%E7%A4%BE%E4%BC%9A-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC/dp/4757143044?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4757143044">信頼と裏切りの社会</a></dt>
	<dd>ブルース・シュナイアー</dd>
	<dd>山形 浩生 (翻訳)</dd>
    <dd>エヌティティ出版 2013-12-24</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4757143044, EAN: 9784757143043</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">社会における「信頼」とは。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-23">2019-03-23</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
