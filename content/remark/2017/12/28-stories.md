+++
title = "2017年最後の戯れ言（多分）"
date =  "2017-12-28T17:54:26+09:00"
description = "おひとりさまクリスマス / 祝♪ 「 続・情報共有の未来」書籍化 / サイトオーナーがページの広告掲載の代わりにマイニング・コードを仕込むのはヤクザの「みかじめ料」と同じ"
image = "/images/attention/remark.jpg"
tags        = [ "code", "internet", "security" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

これが多分2017年最後の記事になると思われ。

1. [おひとりさまクリスマス]({{< relref "#xmas" >}})
1. [祝♪ 「 続・情報共有の未来」書籍化]({{< relref "#infoshare2" >}})
1. [サイトオーナーがページの広告掲載の代わりにマイニング・コードを仕込むのはヤクザの「みかじめ料」と同じ]({{< relref "#gray" >}})
1. [というわけで]({{< relref "#bye" >}})

## おひとりさまクリスマス{#xmas}

別に特定の連れ合いがいるわけではないのだが，考えてみたらクリスマス・イブおよびクリスマスの夜を自宅で独りで過ごすとか何年ぶりだろう。
昔のタイムラインとか見ると大抵朝まで飲んだくれてたからなぁ。

というわけで，クリスマスな感じの記事を紹介。

- [クリスマス・キャロル――海賊版が意味をなさなくなる日 – P2Pとかその辺のお話R](http://p2ptk.org/copyright/698)
- [数式って言葉なんだよ（結城浩ミニ文庫）｜結城浩](https://mm.hyuki.net/n/n64e93dd7c35b)
- [aozorablog » 小人のくつ屋さん（グリム兄弟・編／大久保ゆう・訳）](http://www.aozora.gr.jp/aozorablog/?p=4074)

「よかった探しリース」が開催されなかったこともあるが，今年は「よかった探し」記事は止めにした。
自身の昨年と今年を振り返ると，あまりに悲惨過ぎて盛大にヘコむので。
来年はちょっとでもいい年になりますように。

## 祝♪ 「 続・情報共有の未来」書籍化{#infoshare2}

- [もうすぐ絶滅するという開かれたウェブについて  続・情報共有の未来 - 達人出版会](https://tatsu-zine.com/books/infoshare2)
- [『もうすぐ絶滅するという開かれたウェブについて 続・情報共有の未来』サポートページ公開 - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20171226/openweb)

ゴメンナサイ。
今は購入できない。
PayPal が使えないのよ。
いや，使えるようにすることはできるんだけど（アカウントはある），いまは使いたくないというか...

でも[ボーナストラックが面白いらしい](http://d.hatena.ne.jp/yomoyomo/20171228/openweb)ので来年の何処かで買う算段をつけます（とかいって忘れるんだよなぁ）。

## サイトオーナーがページの広告掲載の代わりにマイニング・コードを仕込むのはヤクザの「みかじめ料」と同じ{#gray}

（この記事は [Scrapbox の記事](https://scrapbox.io/spiegel-branch/%E3%82%B5%E3%82%A4%E3%83%88%E3%82%AA%E3%83%BC%E3%83%8A%E3%83%BC%E3%81%8C%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%AE%E5%BA%83%E5%91%8A%E6%8E%B2%E8%BC%89%E3%81%AE%E4%BB%A3%E3%82%8F%E3%82%8A%E3%81%AB%E3%83%9E%E3%82%A4%E3%83%8B%E3%83%B3%E3%82%B0JavaScript%E3%82%92%E4%BB%95%E8%BE%BC%E3%82%80%E3%81%AE%E3%81%AF%E3%83%A4%E3%82%AF%E3%82%B6%E3%81%AE%E3%80%8C%E3%81%BF%E3%81%8B%E3%81%98%E3%82%81%E6%96%99%E3%80%8D%E3%81%A8%E5%90%8C%E3%81%98%E3%80%82)の再構成です）

- [ニュース解説 - FinTech？マルウエア？無断でスマホCPU使う謎のサービス：ITpro](http://itpro.nikkeibp.co.jp/atcl/column/14/346926/110801194/)

私はブラウザに関してはとりあえず [keraf/NoCoin](https://github.com/keraf/NoCoin "keraf/NoCoin: No Coin is a tiny browser extension aiming to block coin miners such as Coinhive.") を入れて自衛している。
どこまで効くのかは分からないが...

こういうのは線引きが難しい。
ウイルス対策ソフトを提供してるセキュリティ企業はこれを「グレーウェア」とか呼んでるらしい。

{{< fig-quote title="FinTech？マルウエア？無断でスマホCPU使う謎のサービス" link="http://itpro.nikkeibp.co.jp/atcl/column/14/346926/110801194/" >}}
<q>しかし、現実としてはトレンドマイクロのウイルスバスターはCoinhiveをブロックしている。CoinhiveのJavaScriptが埋め込まれたWebサイトにアクセスすると、警告メッセージを表示してスクリプトの実行を止める。シマンテックなどほかのベンダーのセキュリティソフトも同様だ。<br>
マルウエアには当たらないが、好ましくない動作を行う可能性がある「グレーウエア」に分類されているためだ。</q>
{{< /fig-quote >}}

しかし提供元の言い分

{{< fig-quote title="FinTech？マルウエア？無断でスマホCPU使う謎のサービス" link="http://itpro.nikkeibp.co.jp/atcl/column/14/346926/110801194/" >}}
<q>Coinhiveの運営チームは自分たちのサービスを、「押しつけがましい広告の代わりになる」と主張している。少なからぬスマホアプリやWebサイトで、操作を邪魔するように画面に重なる広告や、強制的に視聴させる動画広告を表示している。こうした広告の代わりに、仮想通貨を採掘した見返りとしてコンテンツを提供すればいい、というのがCoinhiveの主張だ。</q>
{{< /fig-quote >}}

というのは，むかし繁華街でよく見かけた「ヤクザが店の用心棒をする代わりにみかじめ料を要求する」構図そっくりそのままである。
当の企業やそれを利用するサイトオーナーがそこまで意識してやっているかどうかは知らないが。

分散型コンピューティング・プロジェクトではユーザ側の（消費電力を含めた）計算資源をどう管理するかが常に問題となる。
[SETI@home](http://setiathome.berkeley.edu/) の古参ユーザなら「カリフォルニア電力危機」当時を思い出す人もいるかもしれない。
ちなみに，現在の [SETI@home](http://setiathome.berkeley.edu/) でも使われている学術用分散コンピューティング・プラットフォーム [BOINC](http://boinc.berkeley.edu/) では実行時のプロセッサ占有率をチューニングすることができる。

そういったものに比べると今回のやり方は全くもって筋が悪い。
かつての「[SETI@home日本語情報ページ](https://web.archive.org/web/*/http://www.planetary.or.jp/setiathome/)」には「計算機による信号傍受と解析は人類のためにこそ使いましょう」と書かれていたが，不特定多数がアクセスする Web ページにマイニング・スクリプトを仕込むことが「誰得」なのか真剣に考えるべきだと思う。

### その他の記事

- [ハッカーがあなたのブラウザを利用して、仮想通貨をマイニングしているかもしれない | ライフハッカー［日本版］](https://www.lifehacker.jp/2017/10/171031-how-to-stop-sites-from-harvesting-cryptocurrency-from.html)
- [あなたのブラウザーは、誰かのために暗号通貨を勝手に「採掘」しているかもしれない｜WIRED.jp](https://wired.jp/2017/12/25/cryptojacking/)
- [仮想通貨のマイニングでスマホのバッテリを破壊するトロイの木馬「Loapi」  - PC Watch](https://pc.watch.impress.co.jp/docs/news/1098829.html)
- [Operaのニューバージョンはユーザーのマシンで勝手に暗号通貨がマイニングされるのを防ぐ  |  TechCrunch Japan](https://techcrunch.com/2018/01/03/opera-now-protects-you-from-cryptojacking-attacks/)

## というわけで{#bye}

今年はおしまい。
来年お会いしましょう。

## 参考図書

{{% review-paapi "B00C41BSHM" %}} <!-- 超人ロック　ミラーリング -->
