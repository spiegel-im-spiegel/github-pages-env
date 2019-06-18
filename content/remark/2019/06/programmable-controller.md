+++
title = "【憶測記事】 マイニング・コードを仕込まれたかもしれない"
date =  "2019-06-18T22:22:58+09:00"
description = "今回のようなケースでは「まずブラウザを疑え」ってのが分かっただけでもよかったことにしよう。"
image = "/images/attention/kitten.jpg"
tags = [ "k-tai", "chromium", "firefox" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

あらかじめ予防線を張っておくと，今回の記事はかなりの部分が憶測でいつも以上に独断で偏見に満ちているので，そのへん割り引いて読んでほしい。

いや，そもそも先週末くらいから手持ちのスマートフォンが異様に熱くなって急速にバッテリが減衰する症状が断続的にあらわれたのが始まり。
思い当たるフシが全く無いので似たような事例がないかネットで探しつつ，監視用のアプリを一時的に導入して色々調べてみるが，プロセッサの全コアがフルでブン回っている事以外に具体的に誰がどんな悪さをしているのかは分からなかった。

それでも，今朝起床したら（電源を繋いでいるにも関わらず）バッテリ残量が残り1パーセントになってるのを見てこれは本格的にヤバいと思い始めた。
幸いなことに節電モードにして通話機能等の主要なアプリ以外をオフにしてしまえばバッテリ消費が抑えられるのは分かっていたので，とりあえず今日は（家の充電器に繋いだまま）ケータイを持たずに仕事に出掛けた（デスクワークなので仕事にケータイは要らないけど出退勤のバス移動中に音楽が聴けないのは耳が寂しかった）。

んで，自宅に帰ってからふと思いついて Chrome と Firefox のデータを完全初期化してみたら嘘みたいにおさまった。

**原因はお前か！**

これってアレかね。
最近噂のマイニング・コードってやつを仕込まれちゃったかね。
ここのところヤバげなサイトも含めてあちこちネットサーフィン（古語`w`）してたからなぁ。
今となっては確かめる術はない。
一応 Firefox には [No Coin] を入れてる。
Chrome は基本放置で（拡張機能を入れるには Google アカウントにサインインしないといけないため）， Firefox では行きたくないサイトのみプライベートブラウジングで閲覧するようにしてたんだけど，やっぱ Chrome は使わないほうがいいかな。

最近 Mozilla が調子こいて

{{< fig-quote title="Firefoxがユーザーを追跡から護る機能を強化しパスワードマネージャーをデスクトップに導入 | TechCrunch Japan" link="https://jp.techcrunch.com/2019/06/06/2019-06-04-firefox-gets-enhanced-tracking-protection-desktop-password-manager-and-more/" >}}
<q>Firefoxではつねに、口先以上のことをしている。本当に人々を護るには、プライバシーを再優先する新しいスタンダードの確立が必要だ</q>
{{< /fig-quote >}}

とか言ってるそうな。
どの口が言うか！ って話だよ。
本気で「口先以上のことをしている」というのなら検索窓の既定を DuckDuckGo にしろっての！ 話はそこからだ。
あと Firefox の利用は明らかに携帯端末のほうが多いのに携帯端末版 Firefox のセキュリティ・プライバシー設定は旧態依然としたままなのはどうしてなのかね？ ユーザをナメるのも大概にしていただきたいものである。
これならば潔くあからさまにユーザ軽視の Safari や Chrome のほうがまだマシかもしれない（ユーザが正しく警戒するという意味で）。

今回思ったけど，こういうイレギュラーが起きたときにユーザ自身で調べる手段が少ないってのは致命的だな。
Android 用の開発環境でもあればよかったのかもしれないが，私はもうエンジニアではないのでそんなもん入れたくないし，たぶんケータイ・ショップに行っても分かる人いないよね。

つくづくケータイってのは Personal Computer じゃなくて Programmable Controller なんだなぁ，と改めて実感する。
まぁ最近は Windows を始めとした所謂「パソコン」も Programmable Controller に成り下がってるけど（笑）

まぁ，今回のようなケースでは「まずブラウザを疑え」ってのが分かっただけでもよかったことにしよう。

## ブックマーク

- [技術的負債としての Web ブラウザ]({{< ref "/remark/2018/06/web-browser-as-the-technical-debt.md" >}})

[No Coin]: https://github.com/keraf/NoCoin/ "keraf/NoCoin: No Coin is a tiny browser extension aiming to block coin miners such as Coinhive."

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
