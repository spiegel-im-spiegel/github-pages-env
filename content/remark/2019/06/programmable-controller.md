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

## 追記

その後 Chrome, Firefox, WebView がアップデートされたが，これ以降，端末の発熱等はピタリとおさまった。
これらのうちのどれか（または全部）が犯人なのは間違いないだろう。

今回の経験を経て，ケータイを使わないときは節電モードにしてデータ通信を行わせないようにした。

## 追記2

- [Google、「古いAndroid端末で異常な大量通信」問題認める　「修正プログラム展開中」 - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/1906/27/news089.html)

上の件も併せれば犯人は WebView でほぼ確定かな。
やっぱ WebView は滅びるべきだな。
ついでに WebKit も（笑）

## ブックマーク

- [技術的負債としての Web ブラウザ]({{< ref "/remark/2018/06/web-browser-as-the-technical-debt.md" >}})

[No Coin]: https://github.com/keraf/NoCoin/ "keraf/NoCoin: No Coin is a tiny browser extension aiming to block coin miners such as Coinhive."

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4757143044?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/413qoSjODUL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4757143044?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">信頼と裏切りの社会</a></dt>
    <dd>ブルース・シュナイアー (著), 山形 浩生 (翻訳)</dd>
    <dd>NTT出版 2013-12-24</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4757143044 (ASIN), 9784757143043 (EAN), 4757143044 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">社会における「信頼」とは。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-23">2019-03-23</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
